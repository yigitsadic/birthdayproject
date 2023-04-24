package pg_companies

const (
	companyDetailQuery = `
select id, name from companies
where companies.id = $1
limit 1
`

	companyUpdateQuery = `
update companies
set name = $1
where companies.id = $2
returning id, name
`
)
