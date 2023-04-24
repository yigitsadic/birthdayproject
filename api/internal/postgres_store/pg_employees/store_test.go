package pg_employees

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/birthday-app-api/internal/employees"
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

func Test_PgEmployeesStore(t *testing.T) {
	if os.Getenv("SKIP_DB_TEST") == "YES" {
		t.Skip()
	}

	store := NewPgEmployeeStore(db)

	dto := employees.EmployeeDto{
		CompanyId:  1,
		FirstName:  "Test",
		LastName:   "User",
		Email:      "testuser@google.com",
		BirthDay:   14,
		BirthMonth: 7,
	}

	t.Run("it should successfully handle FetchAll", func(t *testing.T) {
		require.Nil(t, dbtestconfig.ExecuteInitialSchema(db))
		results, err := store.FetchAll(context.TODO(), 1)

		assert.Nil(t, err)
		assert.Equal(t, "yigit", results[0].FirstName)
		assert.Equal(t, "sadic", results[0].LastName)
		assert.Equal(t, "yigit@google.com", results[0].Email)
		assert.Equal(t, 13, results[0].BirthDay)
		assert.Equal(t, 2, results[0].BirthMonth)
	})

	t.Run("it should successfully handle FindOne", func(t *testing.T) {
		require.Nil(t, dbtestconfig.ExecuteInitialSchema(db))
		result, err := store.FindOne(context.TODO(), 1, 1)

		assert.Nil(t, err)
		assert.Equal(t, "yigit", result.FirstName)
		assert.Equal(t, "sadic", result.LastName)
		assert.Equal(t, "yigit@google.com", result.Email)
		assert.Equal(t, 13, result.BirthDay)
		assert.Equal(t, 2, result.BirthMonth)
	})

	t.Run("it should successfully handle Create", func(t *testing.T) {
		require.Nil(t, dbtestconfig.ExecuteInitialSchema(db))
		result, err := store.Create(context.TODO(), 1, dto)

		assert.Nil(t, err)
		assert.Equal(t, dto.FirstName, result.FirstName)
		assert.Equal(t, dto.LastName, result.LastName)
		assert.Equal(t, dto.Email, result.Email)
		assert.Equal(t, dto.CompanyId, result.CompanyId)
		assert.Equal(t, dto.BirthDay, result.BirthDay)
		assert.Equal(t, dto.BirthMonth, result.BirthMonth)
	})

	t.Run("it should successfully handle Update", func(t *testing.T) {
		require.Nil(t, dbtestconfig.ExecuteInitialSchema(db))

		result, err := store.Update(context.TODO(), 1, 1, dto)

		assert.Nil(t, err)
		assert.Equal(t, dto.FirstName, result.FirstName)
		assert.Equal(t, dto.LastName, result.LastName)
		assert.Equal(t, dto.Email, result.Email)
		assert.Equal(t, dto.CompanyId, result.CompanyId)
		assert.Equal(t, dto.BirthDay, result.BirthDay)
		assert.Equal(t, dto.BirthMonth, result.BirthMonth)
	})

	t.Run("it should successfully handle Delete", func(t *testing.T) {
		require.Nil(t, dbtestconfig.ExecuteInitialSchema(db))

		err := store.Delete(context.TODO(), 1, 1)

		assert.Nil(t, err)
	})
}
