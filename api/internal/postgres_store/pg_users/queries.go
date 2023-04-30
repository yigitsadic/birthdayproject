package pg_users

const (
	userDetailQuery = `
select id, first_name, last_name, email from users
where users.id = $1
limit 1
`

	userUpdateQuery = `
update users
set first_name = $2, last_name = $3, updated_at = now()
where users.id = $1
returning id, first_name, last_name, email
`
)
