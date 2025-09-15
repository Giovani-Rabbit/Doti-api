package dihdl

import (
	dicase "github.com/Giovani-Coelho/Doti-API/internal/infra/di/case"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
)

type ModuleHandlers struct {
	Create     modulehandler.Create
	Delete     modulehandler.Delete
	Get        modulehandler.Get
	GetTasks   modulehandler.GetTasksByModule
	Rename     modulehandler.Rename
	UpdateIcon modulehandler.UpdateIcon
}

func NewModuleHandlers(modulecase *dicase.ModuleUseCases) *ModuleHandlers {
	return &ModuleHandlers{
		Create:     modulehandler.NewCreateModuleHandler(modulecase.Create),
		Delete:     modulehandler.NewDeleteHandler(modulecase.Delete),
		Get:        modulehandler.NewGetHandler(modulecase.GetByUser),
		GetTasks:   modulehandler.NewGetTasksByModuleHandler(modulecase.GetTasks),
		Rename:     modulehandler.NewRenameHandler(modulecase.Rename),
		UpdateIcon: modulehandler.NewUpdateIconHandler(modulecase.UpdateIcon),
	}
}
