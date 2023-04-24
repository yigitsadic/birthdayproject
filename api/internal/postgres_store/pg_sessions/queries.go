package pg_sessions

const (
	findUserByIdQuery = `
select id, email, password_hash, company_id from users
where users.id = $1
limit 1
	`

	findUserByEmail = `
	select id, email, password_hash, company_id from users
where users.email = $1
limit 1
	`
)
