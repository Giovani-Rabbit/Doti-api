package container

import (
	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type ModuleCase struct {
	create modulecase.CreateModuleUseCase
	get    modulecase.GetModulesUseCase
	rename modulecase.RenameModuleUseCase
	delete modulecase.DeleteModuleUseCase
}

func newModuleCase(moduleRepo repository.ModuleRepository) *ModuleCase {
	return &ModuleCase{
		create: modulecase.NewCreateModuleUseCase(moduleRepo),
		get:    modulecase.NewGetModulesUseCase(moduleRepo),
		rename: modulecase.NewRenameModuleUseCase(moduleRepo),
		delete: modulecase.NewDeleteModuleUseCase(moduleRepo),
	}
}
