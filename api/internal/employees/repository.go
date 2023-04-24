package employees

import "context"

type EmployeeRepository interface {
	FetchAll(ctx context.Context, companyId int) ([]*EmployeeModel, error)
	FindOne(ctx context.Context, companyId int, employeeId int) (*EmployeeModel, error)
	Create(ctx context.Context, companyId int, dto EmployeeDto) (*EmployeeModel, error)
	Update(ctx context.Context, companyId int, employeeId int, dto EmployeeDto) (*EmployeeModel, error)
	Delete(ctx context.Context, companyId int, employeeId int) error
}
