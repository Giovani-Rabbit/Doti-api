package container

import modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"

type ModuleHandler struct {
	Create modulehandler.Create
	Delete modulehandler.Delete
	Get    modulehandler.Get
	Rename modulehandler.Rename
}

func newModuleHandler(modulecase *ModuleCase) *ModuleHandler {
	return &ModuleHandler{
		Create: modulehandler.NewCreateModuleHandler(modulecase.create),
		Delete: modulehandler.NewDeleteHandler(modulecase.delete),
		Get:    modulehandler.NewGetHandler(modulecase.get),
		Rename: modulehandler.NewRenameHandler(modulecase.rename),
	}
}
