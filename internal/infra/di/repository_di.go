package di

import (
	"database/sql"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type repositoryContainer struct {
	module      repository.ModuleRepository
	task        repository.TaskRepository
	taskDetails repository.TaskDetailsRepository
	user        repository.UserRepository
}

func newRepositoryContainer(db *sql.DB) *repositoryContainer {
	moduleRepository := repository.NewModuleRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	taskDetailsRepository := repository.NewTaskDetailsRepository(db)
	userRepository := repository.NewUserRepository(db)

	return &repositoryContainer{
		module:      moduleRepository,
		task:        taskRepository,
		taskDetails: taskDetailsRepository,
		user:        userRepository,
	}
}
