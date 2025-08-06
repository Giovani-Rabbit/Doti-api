package container

import (
	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

func (c *container) NewModule() modulehandler.ModuleHandler {
	moduleRepo := repository.NewModuleRepository(c.DB)

	createModulecase := modulecase.NewCreateModuleUseCase(moduleRepo)
	getModulescase := modulecase.NewGetModulesUseCase(moduleRepo)
	renameModulecase := modulecase.NewRenameModuleUseCase(moduleRepo)

	return modulehandler.New(
		createModulecase,
		getModulescase,
		renameModulecase,
	)
}
