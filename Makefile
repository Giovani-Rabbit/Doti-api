run:
	@go run cmd/api/main.go

test: 
	@go test ./...

migration:
	@migrate create -ext sql -dir internal/infra/persistence/db/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down