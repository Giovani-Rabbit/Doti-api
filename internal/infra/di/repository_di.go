package di

import (
	"database/sql"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type repositoryContainer struct {
	module repository.ModuleRepository
	task   repository.TaskRepository
	user   repository.UserRepository
}

func newRepositoryContainer(db *sql.DB) *repositoryContainer {
	moduleRepository := repository.NewModuleRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userRepository := repository.NewUserRepository(db)

	return &repositoryContainer{
		module: moduleRepository,
		task:   taskRepository,
		user:   userRepository,
	}
}
