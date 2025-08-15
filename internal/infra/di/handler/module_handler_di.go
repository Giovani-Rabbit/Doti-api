package dihdl

import (
	dicase "github.com/Giovani-Coelho/Doti-API/internal/infra/di/case"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
)

type ModuleHandlers struct {
	Create modulehandler.Create
	Delete modulehandler.Delete
	Get    modulehandler.Get
	Rename modulehandler.Rename
}

func NewModuleHandlers(modulecase *dicase.ModuleUseCases) *ModuleHandlers {
	return &ModuleHandlers{
		Create: modulehandler.NewCreateModuleHandler(modulecase.Create),
		Delete: modulehandler.NewDeleteHandler(modulecase.Delete),
		Get:    modulehandler.NewGetHandler(modulecase.Get),
		Rename: modulehandler.NewRenameHandler(modulecase.Rename),
	}
}
