package container

import (
	"database/sql"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type RepositoryContainer struct {
	moduleRepo repository.ModuleRepository
	userRepo   repository.UserRepository
}

func newRepository(db *sql.DB) *RepositoryContainer {
	moduleRepository := repository.NewModuleRepository(db)
	userRepository := repository.NewUserRepository(db)

	return &RepositoryContainer{
		moduleRepo: moduleRepository,
		userRepo:   userRepository,
	}
}
