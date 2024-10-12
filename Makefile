migration:
	@migrate create -ext sql -dir src/infra/database/store/pgstore/migrations $(filter-out $@, $(MAKECMDGOALS))