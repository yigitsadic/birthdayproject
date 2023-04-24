package companies

import "context"

type CompanyRepository interface {
	FetchOne(context.Context, int) (*CompanyModel, error)
	Update(context.Context, int, CompanyUpdateDto) (*CompanyModel, error)
}
