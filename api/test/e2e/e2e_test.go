package e2e

import (
	"database/sql"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/yigitsadic/birthday-app-api/cmd/api/server"
	"github.com/yigitsadic/birthday-app-api/internal/auth"
	"github.com/yigitsadic/birthday-app-api/internal/postgres_store/pg_companies"
	"github.com/yigitsadic/birthday-app-api/internal/postgres_store/pg_employees"
	"github.com/yigitsadic/birthday-app-api/internal/postgres_store/pg_sessions"
	"github.com/yigitsadic/birthday-app-api/internal/postgres_store/pg_users"
	"github.com/yigitsadic/birthday-app-api/test/dbtestconfig"
	"go.uber.org/zap"
)

type IntegrationTestSuite struct {
	suite.Suite

	Pool       *dockertest.Pool
	Resource   *dockertest.Resource
	DB         *sql.DB
	JWTStore   auth.JWTStore
	TestServer *httptest.Server
}

func (testSuite *IntegrationTestSuite) BeforeTest(suiteName, testName string) {
	_ = dbtestconfig.ExecuteInitialSchema(testSuite.DB)
}

func (testSuite *IntegrationTestSuite) SetupSuite() {
	pool, resource, db, err := dbtestconfig.ConnectTestDockerContainer()

	require.Nil(testSuite.T(), err)

	jwtStore := auth.NewJWT("12345")

	testSuite.JWTStore = jwtStore
	testSuite.Pool = pool
	testSuite.Resource = resource
	testSuite.DB = db

	sessionsStore := pg_sessions.NewPgSessionsStore(db)
	usersStore := pg_users.NewPgUserStore(db)
	employeesStore := pg_employees.NewPgEmployeeStore(db)
	companiesStore := pg_companies.NewPgCompanyStore(db)

	logger, err := zap.NewProduction()
	require.Nil(testSuite.T(), err)

	srv := server.Server{
		JWTStore: jwtStore,
		Logger:   logger,

		SessionRepository:  sessionsStore,
		UserRepository:     usersStore,
		EmployeeRepository: employeesStore,
		CompanyRepository:  companiesStore,
	}
	r := chi.NewMux()

	srv.MountRoutes(r)

	testSuite.TestServer = httptest.NewServer(r)
}

func (testSuite *IntegrationTestSuite) TearDownSuite() {
	testSuite.TestServer.Close()

	err := dbtestconfig.PurgeResources(testSuite.Pool, testSuite.Resource)

	require.Nil(testSuite.T(), err)
}

func Test_E2ETests(t *testing.T) {
	if os.Getenv("SKIP_DB_TEST") == "YES" {
		return
	}

	suite.Run(t, new(IntegrationTestSuite))
}
