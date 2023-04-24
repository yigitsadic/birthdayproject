package pg_users

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/birthday-app-api/internal/users"
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

	store := NewPgUserStore(db)

	t.Run("it should successfully handle GetUser", func(t *testing.T) {
		t.Run("it should retrieve user", func(t *testing.T) {
			res, err := store.GetUser(context.TODO(), 1)

			assert.Nil(t, err)
			assert.Equal(t, 1, res.ID)
		})

		t.Run("it should handle not found", func(t *testing.T) {
			res, err := store.GetUser(context.TODO(), 2)

			assert.NotNil(t, err)
			assert.Equal(t, sql.ErrNoRows, err)
			assert.Nil(t, res)
		})
	})

	t.Run("it should successfully handle UpdateUser", func(t *testing.T) {
		dto := users.UserDto{
			FirstName: "Gandalf",
			LastName:  "the Gray",
		}

		t.Run("it should update", func(t *testing.T) {
			res, err := store.UpdateUser(context.TODO(), 1, dto)

			assert.Nil(t, err)
			assert.Equal(t, "Gandalf", res.FirstName)
			assert.Equal(t, "the Gray", res.LastName)
		})

		t.Run("it should give an error if no user found", func(t *testing.T) {
			res, err := store.UpdateUser(context.TODO(), 2, dto)

			assert.NotNil(t, err)
			assert.Equal(t, sql.ErrNoRows, err)
			assert.Nil(t, res)
		})
	})
}
