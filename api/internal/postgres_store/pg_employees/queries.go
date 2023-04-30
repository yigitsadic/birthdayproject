package pg_employees

const (
	fetchAllQuery = `
select id, company_id, first_name, last_name, email, birth_day, birth_month
from employees
where employees.company_id = $1
order by employees.id DESC
`

	findOneQuery = `
select id, company_id, first_name, last_name, email, birth_day, birth_month
from employees
where employees.company_id = $1 and employees.id = $2
	`

	createQuery = `
insert into employees(company_id, first_name, last_name, email, birth_day, birth_month, created_at, updated_at)
values ($1, $2, $3, $4, $5, $6, now(), now())
returning id, company_id, first_name, last_name, email, birth_day, birth_month
	`

	updateQuery = `
update employees set first_name = $1, last_name = $2, email = $3, birth_day = $4, birth_month = $5, updated_at=now()
where employees.company_id = $6 and employees.id = $7
returning id, company_id, first_name, last_name, email, birth_day, birth_month
`
	deleteQuery = `
	delete from employees where employees.company_id = $1 and employees.id = $2
`
)
