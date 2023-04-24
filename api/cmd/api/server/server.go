package server

import (
	"github.com/yigitsadic/birthday-app-api/internal/auth"
	"github.com/yigitsadic/birthday-app-api/internal/companies"
	"github.com/yigitsadic/birthday-app-api/internal/employees"
	"github.com/yigitsadic/birthday-app-api/internal/sessions"
	"github.com/yigitsadic/birthday-app-api/internal/users"
	"go.uber.org/zap"
)

// Server is dependency injection stores repositories, logger etc.
type Server struct {
	CompanyRepository  companies.CompanyRepository
	EmployeeRepository employees.EmployeeRepository
	SessionRepository  sessions.SessionRepository
	UserRepository     users.UserRepository

	JWTStore auth.JWTStore
	Logger   *zap.Logger
}
