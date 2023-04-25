.PHONY: migrate
migrate:
	@echo "Migrating database..."
	@migrate -path db -database "postgres://birthdayproject:birthdayproject@localhost:5435/birthdayproject?sslmode=disable" up

.PHONY: rollback
rollback:
	@echo "Rolling back database one step..."
	@migrate -path db -database "postgres://birthdayproject:birthdayproject@localhost:5435/birthdayproject?sslmode=disable" down 1

.PHONY: pgdump
pgdump:
	@echo "Updating database schema..."
	@docker run -it -v "$(PWD):/tmp" -e "PGPASSWORD=birthdayproject" --network=birthdayproject_backend postgres:15.2-alpine pg_dump --username=birthdayproject --dbname=birthdayproject --host=database --port=5432 --file=/tmp/schema.sql --schema-only

