package dicase

import (
	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type ModuleUseCases struct {
	Create modulecase.CreateModuleUseCase
	Get    modulecase.GetModulesUseCase
	Rename modulecase.RenameModuleUseCase
	Delete modulecase.DeleteModuleUseCase
}

func NewModuleCases(moduleRepo repository.ModuleRepository) *ModuleUseCases {
	return &ModuleUseCases{
		Create: modulecase.NewCreateModuleUseCase(moduleRepo),
		Get:    modulecase.NewGetModulesUseCase(moduleRepo),
		Rename: modulecase.NewRenameModuleUseCase(moduleRepo),
		Delete: modulecase.NewDeleteModuleUseCase(moduleRepo),
	}
}
