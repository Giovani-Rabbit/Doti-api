migration:
	@migrate create -ext sql -dir src/infra/database/store/pgstore/migrations $(filter-out $@, $(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down