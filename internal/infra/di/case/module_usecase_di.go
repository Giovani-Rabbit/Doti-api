package dicase

import (
	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type ModuleUseCases struct {
	Create     modulecase.Create
	GetByUser  modulecase.GetByUser
	GetTasks   modulecase.GetTasks
	Rename     modulecase.Rename
	Delete     modulecase.Delete
	UpdateIcon modulecase.UpdateIcon
}

func NewModuleCases(
	mr repository.ModuleRepository,
	tr repository.TaskRepository,
) *ModuleUseCases {
	return &ModuleUseCases{
		Create:     modulecase.NewCreateModuleUseCase(mr),
		GetByUser:  modulecase.NewGetModulesUseCase(mr),
		GetTasks:   modulecase.NewGetTasksByModuleId(tr, mr),
		Rename:     modulecase.NewRenameModuleUseCase(mr),
		Delete:     modulecase.NewDeleteModuleUseCase(mr),
		UpdateIcon: modulecase.NewUpdateModuleIconUseCase(mr),
	}
}
