package container

import (
	modulecase "github.com/Giovani-Coelho/Doti-API/src/core/app/module"
	modulehandler "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/module"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
)

func (c *Container) NewModuleContainer() modulehandler.IModuleHandler {
	moduleRepo := repository.NewModuleRepository(c.DB)

	createModulecase := modulecase.NewCreateModuleUseCase(moduleRepo)

	return modulehandler.New(
		createModulecase,
	)
}
