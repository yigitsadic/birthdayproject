.PHONY: migrate
migrate:
	@echo "Migrating database..."
	@migrate -path db -database "postgres://birthdayproject:birthdayproject@localhost:5435/birthdayproject?sslmode=disable" up

.PHONY: rollback
rollback:
	@echo "Rolling back database one step..."
	@migrate -path db -database "postgres://birthdayproject:birthdayproject@localhost:5435/birthdayproject?sslmode=disable" down 1
