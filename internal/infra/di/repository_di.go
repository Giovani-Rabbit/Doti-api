package di

import (
	"database/sql"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type repositoryContainer struct {
	module repository.ModuleRepository
	user   repository.UserRepository
}

func newRepositoryContainer(db *sql.DB) *repositoryContainer {
	moduleRepository := repository.NewModuleRepository(db)
	userRepository := repository.NewUserRepository(db)

	return &repositoryContainer{
		module: moduleRepository,
		user:   userRepository,
	}
}
