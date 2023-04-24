package pg_companies

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/yigitsadic/birthday-app-api/internal/companies"
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

func Test_PgCompanyStore(t *testing.T) {
	if os.Getenv("SKIP_DB_TEST") == "YES" {
		t.Skip()
	}

	store := NewPgCompanyStore(db)

	t.Run("it should successfully handle FetchOne", func(t *testing.T) {
		t.Run("it should retrieve company", func(t *testing.T) {
			res, err := store.FetchOne(context.TODO(), 1)

			assert.Nil(t, err)
			assert.Equal(t, 1, res.ID)
		})

		t.Run("it should handle not found", func(t *testing.T) {
			res, err := store.FetchOne(context.TODO(), 2)

			assert.NotNil(t, err)
			assert.Equal(t, sql.ErrNoRows, err)
			assert.Nil(t, res)
		})
	})

	t.Run("it should successfully handle Update", func(t *testing.T) {
		dto := companies.CompanyUpdateDto{
			Name: "Evil Company",
		}

		t.Run("it should update", func(t *testing.T) {
			res, err := store.Update(context.TODO(), 1, dto)

			assert.Nil(t, err)
			assert.Equal(t, "Evil Company", res.Name)
		})

		t.Run("it should give an error if no company found", func(t *testing.T) {
			res, err := store.Update(context.TODO(), 2, dto)

			assert.NotNil(t, err)
			assert.Equal(t, sql.ErrNoRows, err)
			assert.Nil(t, res)
		})
	})
}
