package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/yigitsadic/birthday-app-api/internal/auth"
	"github.com/yigitsadic/birthday-app-api/internal/common"
	"github.com/yigitsadic/birthday-app-api/internal/postgres_store/pg_companies"
	"github.com/yigitsadic/birthday-app-api/internal/postgres_store/pg_employees"
	"github.com/yigitsadic/birthday-app-api/internal/postgres_store/pg_sessions"
	"github.com/yigitsadic/birthday-app-api/internal/postgres_store/pg_users"
	"github.com/yigitsadic/birthday-app-api/internal/server"
	"go.uber.org/zap"
)

func main() {
	connectionString := os.Getenv("POSTGRES_DSL")
	if connectionString == "" {
		connectionString = "postgres://birthdayproject:birthdayproject@localhost:5435/birthdayproject?sslmode=disable"
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// Linter error.
	defer func() {
		_ = logger.Sync()
	}()

	dbClient, err := sql.Open("postgres", connectionString)
	if err != nil {
		logger.Fatal("error occurred initializing db object", zap.Error(err))
	}

	if err = dbClient.Ping(); err != nil {
		logger.Fatal("error occurred pinging database", zap.Error(err))
	}

	defer dbClient.Close()

	dbClient.SetMaxOpenConns(25)
	dbClient.SetMaxIdleConns(25)
	dbClient.SetConnMaxLifetime(5 * time.Minute)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7755"
	}

	sessionsStore := pg_sessions.NewPgSessionsStore(dbClient)
	usersStore := pg_users.NewPgUserStore(dbClient)
	employeesStore := pg_employees.NewPgEmployeeStore(dbClient)
	companiesStore := pg_companies.NewPgCompanyStore(dbClient)

	// Override default chi logger with zap logger.
	middleware.DefaultLogger = common.DefaultLogger(logger)

	// FIX: This value should be taken from env variable.
	jwtStore := auth.NewJWT("top secret")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/heartbeat"))
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           60 * 10, // 10 minutes, maximum value not ignored by any of major browsers
	}))

	srv := server.Server{
		CompanyRepository:  companiesStore,
		EmployeeRepository: employeesStore,
		SessionRepository:  sessionsStore,
		UserRepository:     usersStore,

		JWTStore: jwtStore,
		Logger:   logger,
	}

	srv.MountRoutes(r)

	logger.Info("Server up and running", zap.String("port", port))

	logger.Fatal("server shutdown",
		zap.Error(
			http.ListenAndServe(":"+port, r),
		),
	)
}
