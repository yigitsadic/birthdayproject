package pg_employees

import (
	"context"
	"database/sql"

	"github.com/yigitsadic/birthday-app-api/internal/employees"
)

type PgEmployeesStore struct {
	DB *sql.DB
}

// Ensure v follows EmployeeRepository
var _ employees.EmployeeRepository = (*PgEmployeesStore)(nil)

func NewPgEmployeeStore(db *sql.DB) *PgEmployeesStore {
	return &PgEmployeesStore{db}
}

func (store *PgEmployeesStore) FetchAll(ctx context.Context, companyId int) ([]*employees.EmployeeModel, error) {
	rows, err := store.DB.QueryContext(ctx, fetchAllQuery, companyId)

	if err != nil {
		return []*employees.EmployeeModel{}, err
	}

	var models []*employees.EmployeeModel

	for rows.Next() {
		model := employees.EmployeeModel{}

		rows.Scan(
			&model.ID, &model.CompanyId, &model.FirstName, &model.LastName, &model.Email, &model.BirthDay, &model.BirthMonth,
		)

		models = append(models, &model)
	}

	return models, nil
}

func (store *PgEmployeesStore) FindOne(ctx context.Context, companyId int, employeeId int) (*employees.EmployeeModel, error) {
	row := store.DB.QueryRowContext(ctx, findOneQuery, companyId, employeeId)
	model := employees.EmployeeModel{}
	err := row.Scan(&model.ID, &model.CompanyId, &model.FirstName, &model.LastName, &model.Email, &model.BirthDay, &model.BirthMonth)

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (store *PgEmployeesStore) Create(ctx context.Context, companyId int, dto employees.EmployeeDto) (*employees.EmployeeModel, error) {
	row := store.DB.QueryRowContext(ctx, createQuery,
		companyId, dto.FirstName, dto.LastName, dto.Email, dto.BirthDay, dto.BirthMonth,
	)

	model := employees.EmployeeModel{}
	err := row.Scan(
		&model.ID, &model.CompanyId,
		&model.FirstName, &model.LastName,
		&model.Email, &model.BirthDay, &model.BirthMonth,
	)

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (store *PgEmployeesStore) Update(ctx context.Context, companyId int, employeeId int, dto employees.EmployeeDto) (*employees.EmployeeModel, error) {
	row := store.DB.QueryRowContext(ctx, updateQuery,
		dto.FirstName, dto.LastName, dto.Email, dto.BirthDay, dto.BirthMonth,
		companyId, employeeId,
	)

	model := employees.EmployeeModel{}
	err := row.Scan(
		&model.ID, &model.CompanyId,
		&model.FirstName, &model.LastName,
		&model.Email, &model.BirthDay, &model.BirthMonth,
	)

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (store *PgEmployeesStore) Delete(ctx context.Context, companyId int, employeeId int) error {
	_, err := store.DB.ExecContext(ctx, deleteQuery, companyId, employeeId)

	if err != nil {
		return err
	}

	return nil
}
