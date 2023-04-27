package employees

import (
	"context"
	"errors"
)

var mockStoreErr = errors.New("error occurred")

type MockEmployeeStore struct {
	Store                 []*EmployeeModel
	RaiseErrorOnOperation bool
}

// Ensure v follows EmployeeRepository
var _ EmployeeRepository = (*MockEmployeeStore)(nil)

func NewMockEmployeeStore() MockEmployeeStore {
	s := MockEmployeeStore{}
	s.Store = append(s.Store, &EmployeeModel{
		ID:         1,
		CompanyId:  1,
		FirstName:  "Yig",
		LastName:   "Sad",
		Email:      "yig.sad@google.com",
		BirthDay:   15,
		BirthMonth: 4,
	})

	return s
}

func (s MockEmployeeStore) FetchAll(ctx context.Context, companyId int) ([]*EmployeeModel, error) {
	if s.RaiseErrorOnOperation {
		return nil, mockStoreErr
	}

	return s.Store, nil
}

func (s MockEmployeeStore) FindOne(
	ctx context.Context,
	companyId int,
	employeeId int,
) (*EmployeeModel, error) {
	if s.RaiseErrorOnOperation {
		return nil, mockStoreErr
	}

	return s.Store[0], nil
}

func (s MockEmployeeStore) Create(
	ctx context.Context,
	companyId int,
	dto EmployeeDto,
) (*EmployeeModel, error) {
	if s.RaiseErrorOnOperation {
		return nil, mockStoreErr
	}

	return s.Store[0], nil
}

func (s MockEmployeeStore) Update(
	ctx context.Context,
	companyId int,
	employeeId int,
	dto EmployeeDto,
) (*EmployeeModel, error) {
	if s.RaiseErrorOnOperation {
		return nil, mockStoreErr
	}

	return s.Store[0], nil
}

func (s MockEmployeeStore) Delete(ctx context.Context, companyId int, employeeId int) error {
	if s.RaiseErrorOnOperation {
		return mockStoreErr
	}

	return nil
}
