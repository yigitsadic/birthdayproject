package pg_sessions

import (
	"context"
	"database/sql"

	"github.com/yigitsadic/birthday-app-api/internal/sessions"
)

type PgSessionsStore struct {
	DB *sql.DB
}

// Ensure PgSessionsStore follows SessionRepository
var _ sessions.SessionRepository = (*PgSessionsStore)(nil)

func NewPgSessionsStore(db *sql.DB) *PgSessionsStore {
	return &PgSessionsStore{db}
}

func (store *PgSessionsStore) FindUser(ctx context.Context, email string) (*sessions.SessionUser, error) {
	row := store.DB.QueryRowContext(ctx, findUserByEmail, email)
	model := sessions.SessionUser{}
	err := row.Scan(&model.ID, &model.Email, &model.PasswordHash, &model.CompanyId)

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (store *PgSessionsStore) FindById(ctx context.Context, id int) (*sessions.SessionUser, error) {
	row := store.DB.QueryRowContext(ctx, findUserByIdQuery, id)
	model := sessions.SessionUser{}
	err := row.Scan(&model.ID, &model.Email, &model.PasswordHash, &model.CompanyId)

	if err != nil {
		return nil, err
	}

	return &model, nil
}
