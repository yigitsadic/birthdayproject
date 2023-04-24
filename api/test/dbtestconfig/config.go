package dbtestconfig

import (
	"database/sql"
	"fmt"
	"os"

	_ "embed"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

//go:embed schema.sql
var DB_SCHEMA string

const (
	PG_VERSION       = "15.2-alpine"
	PG_TEST_USER     = "api_tester"
	PG_TEST_PASSWORD = "loremipsum"
)

var ConfigEnvVariables = []string{
	"POSTGRES_USER=" + PG_TEST_USER,
	"POSTGRES_PASSWORD=" + PG_TEST_PASSWORD,
	"listen_addresses = '*'",
}

func GetConnectionString(port string) string {
	return fmt.Sprintf(
		"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
		port,
		PG_TEST_USER,
		PG_TEST_PASSWORD,
		PG_TEST_USER,
	)
}

func ExecuteInitialSchema(db *sql.DB) error {
	_, err := db.Exec(DB_SCHEMA)

	return err
}

// ConnectTestDockerContainer Connects to test postgres docker image.
func ConnectTestDockerContainer() (*dockertest.Pool, *dockertest.Resource, *sql.DB, error) {
	if os.Getenv("SKIP_DB_TEST") == "YES" {
		return nil, nil, nil, nil
	}

	var db *sql.DB

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, nil, err
	}

	err = pool.Client.Ping()
	if err != nil {
		return nil, nil, nil, err
	}
	postgres, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        PG_VERSION,
		Env:        ConfigEnvVariables,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		return nil, nil, nil, err
	}

	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", GetConnectionString(postgres.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		return nil, nil, nil, err
	}

	if err = ExecuteInitialSchema(db); err != nil {
		return nil, nil, nil, err
	}

	return pool, postgres, db, nil
}

// PurgeResources purges test postgres docker image.
func PurgeResources(pool *dockertest.Pool, postgres *dockertest.Resource) error {
	if os.Getenv("SKIP_DB_TEST") == "YES" {
		return nil
	}

	if err := pool.Purge(postgres); err != nil {
		return err
	}

	return nil
}
