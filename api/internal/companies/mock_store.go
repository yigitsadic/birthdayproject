package companies

import (
	"context"
	"errors"
)

type MockCompanyStore struct {
	Store              []*CompanyModel
	RaiseErrorOnUpdate bool
	RaiseErrorOnFind   bool
}

// Ensure MockCompanyStore follows CompanyRepository
var _ CompanyRepository = (*MockCompanyStore)(nil)

func NewMockCompanyStore() *MockCompanyStore {
	s := MockCompanyStore{}
	s.Store = append(s.Store, &CompanyModel{
		ID:   16,
		Name: "Evil Co.",
	})
	s.Store = append(s.Store, &CompanyModel{
		ID:   17,
		Name: "Good Co.",
	})

	return &s
}

func (m *MockCompanyStore) FetchOne(_ context.Context, _ int) (*CompanyModel, error) {
	if m.RaiseErrorOnFind {
		return nil, errors.New("record not found")
	}

	return m.Store[0], nil
}

func (m *MockCompanyStore) Update(_ context.Context, _ int, _ CompanyUpdateDto) (*CompanyModel, error) {
	if m.RaiseErrorOnUpdate {
		return nil, errors.New("an error")
	}

	return m.Store[0], nil
}
