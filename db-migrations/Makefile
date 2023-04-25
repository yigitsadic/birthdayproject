include ../.env

.PHONY: migrate
migrate:
	@echo "Migrating database..."
	@migrate -path db -database "${POSTGRES_DSL}" up

.PHONY: rollback
rollback:
	@echo "Rolling back database one step..."
	@migrate -path db -database "${POSTGRES_DSL}" down 1

.PHONY: pgdump
pgdump:
	@echo "Updating database schema..."
	@docker run -it -v "$(PWD):/tmp" -e "PGPASSWORD=${DBPASS}" --network=birthdayproject_backend postgres:15.2-alpine pg_dump --username=${DBNAME} --dbname=${DBNAME} --host=database --port=5432 --file=/tmp/schema.sql --schema-only

