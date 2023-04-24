package pg_companies

import (
	"context"
	"database/sql"

	"github.com/yigitsadic/birthday-app-api/internal/companies"
)

type PgCompaniesStore struct {
	DB *sql.DB
}

// Ensure v follows CompanyRepository
var _ companies.CompanyRepository = (*PgCompaniesStore)(nil)

func NewPgCompanyStore(db *sql.DB) *PgCompaniesStore {
	return &PgCompaniesStore{db}
}

func (store *PgCompaniesStore) FetchOne(ctx context.Context, id int) (*companies.CompanyModel, error) {
	row := store.DB.QueryRowContext(ctx, companyDetailQuery, id)
	model := companies.CompanyModel{}
	err := row.Scan(&model.ID, &model.Name)

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (store *PgCompaniesStore) Update(ctx context.Context, id int, dto companies.CompanyUpdateDto) (*companies.CompanyModel, error) {
	row := store.DB.QueryRowContext(ctx, companyUpdateQuery,
		dto.Name, id,
	)
	model := companies.CompanyModel{}
	err := row.Scan(&model.ID, &model.Name)

	if err != nil {
		return nil, err
	}

	return &model, nil
}
