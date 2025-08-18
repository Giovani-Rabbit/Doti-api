package dicase

import (
	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type ModuleUseCases struct {
	Create    modulecase.Create
	GetByUser modulecase.GetByUser
	Rename    modulecase.Rename
	Delete    modulecase.Delete
}

func NewModuleCases(moduleRepo repository.ModuleRepository) *ModuleUseCases {
	return &ModuleUseCases{
		Create:    modulecase.NewCreateModuleUseCase(moduleRepo),
		GetByUser: modulecase.NewGetModulesUseCase(moduleRepo),
		Rename:    modulecase.NewRenameModuleUseCase(moduleRepo),
		Delete:    modulecase.NewDeleteModuleUseCase(moduleRepo),
	}
}
