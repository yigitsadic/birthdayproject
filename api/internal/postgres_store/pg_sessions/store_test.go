package pg_sessions

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/birthday-app-api/test/dbtestconfig"
)

var db *sql.DB

func TestMain(m *testing.M) {
	pool, postgres, newDb, err := dbtestconfig.ConnectTestDockerContainer()
	if err != nil {
		log.Fatalf("Error occurred initializing test postgres docker image. Err=%s\n", err)
	}

	db = newDb

	code := m.Run()

	_ = dbtestconfig.PurgeResources(pool, postgres)
	os.Exit(code)
}

func Test_PgUserStore(t *testing.T) {
	if os.Getenv("SKIP_DB_TEST") == "YES" {
		t.Skip()
	}

	store := NewPgSessionsStore(db)

	t.Run("it should successfully handle FindUser", func(t *testing.T) {
		t.Run("it should retrieve user", func(t *testing.T) {
			res, err := store.FindUser(context.TODO(), "johndo@google.com")

			assert.Nil(t, err)
			assert.Equal(t, "johndo@google.com", res.Email)
		})

		t.Run("it should handle not found", func(t *testing.T) {
			res, err := store.FindUser(context.TODO(), "noexists@google.com")

			assert.NotNil(t, err)
			assert.Equal(t, sql.ErrNoRows, err)
			assert.Nil(t, res)
		})
	})

	t.Run("it should successfully handle FindById", func(t *testing.T) {
		t.Run("it should update", func(t *testing.T) {
			res, err := store.FindById(context.TODO(), 1)

			assert.Nil(t, err)
			assert.Equal(t, 1, res.ID)
		})

		t.Run("it should give an error if no user found", func(t *testing.T) {
			res, err := store.FindById(context.TODO(), 2)

			assert.NotNil(t, err)
			assert.Equal(t, sql.ErrNoRows, err)
			assert.Nil(t, res)
		})
	})
}
