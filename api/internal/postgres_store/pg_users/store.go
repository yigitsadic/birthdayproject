package pg_users

import (
	"context"
	"database/sql"

	"github.com/yigitsadic/birthday-app-api/internal/users"
)

type PgUsersStore struct {
	DB *sql.DB
}

// Ensure PgUsersStore follows UserRepository
var _ users.UserRepository = (*PgUsersStore)(nil)

func NewPgUserStore(db *sql.DB) *PgUsersStore {
	return &PgUsersStore{db}
}

func (store *PgUsersStore) GetUser(ctx context.Context, id int) (*users.UserModel, error) {
	row := store.DB.QueryRowContext(ctx, userDetailQuery, id)
	model := users.UserModel{}
	err := row.Scan(
		&model.ID,
		&model.FirstName,
		&model.LastName,
		&model.Email,
	)

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (store *PgUsersStore) UpdateUser(ctx context.Context, id int, dto users.UserDto) (*users.UserModel, error) {
	row := store.DB.QueryRowContext(ctx, userUpdateQuery,
		id, dto.FirstName, dto.LastName,
	)
	model := users.UserModel{}
	err := row.Scan(
		&model.ID,
		&model.FirstName,
		&model.LastName,
		&model.Email,
	)

	if err != nil {
		return nil, err
	}

	return &model, nil
}
