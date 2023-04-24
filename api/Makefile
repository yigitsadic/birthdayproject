# Runs application.
.PHONY: run
run:
	@echo "Running application..."
	go run ./cmd/api

# Run tests.
.PHONY: test
test:
	@echo "Running tests..."
	SKIP_DB_TEST=YES go test -cover -coverprofile=./tmp/cover.out ./...
	@echo "Total test coverage: "
	@go tool cover -func=./tmp/cover.out | grep total | awk '{print $$3}'

# Run tests with database.
.PHONY: test/db
test/db:
	@echo "Running tests with database..."
	go test -cover -coverprofile=./tmp/all_test_cover.out ./...
	@echo "Total test coverage: "
	@go tool cover -func=./tmp/all_test_cover.out | grep total | awk '{print $$3}'

.PHONY: tidy
tidy:
	@echo "Tidy project..."
	go mod tidy
	staticcheck ./...
