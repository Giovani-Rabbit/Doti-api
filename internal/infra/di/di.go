package di

import (
	"database/sql"

	dicase "github.com/Giovani-Coelho/Doti-API/internal/infra/di/case"
	dihdl "github.com/Giovani-Coelho/Doti-API/internal/infra/di/handler"
)

type DI struct {
	Module *dihdl.ModuleHandlers
	Task   *dihdl.TaskHandlers
	User   *dihdl.UserHandlers
}

func New(db *sql.DB) *DI {
	repo := newRepositoryContainer(db)

	moduleCases := dicase.NewModuleCases(repo.module, repo.task)
	taskCases := dicase.NewTaskUseCase(repo.task, repo.module)
	userCases := dicase.NewUserCases(repo.user)

	return &DI{
		Module: dihdl.NewModuleHandlers(moduleCases),
		Task:   dihdl.NewTaskHandler(taskCases),
		User:   dihdl.NewUserHandlers(userCases),
	}
}
