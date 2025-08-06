package modulehandler

import (
	"net/http"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
)

type moduleHandler struct {
	CreateModuleUseCase modulecase.CreateModuleUseCase
	GetModulesUseCase   modulecase.GetModulesUseCase
	RenameModuleUseCase modulecase.RenameModuleUseCase
}

type ModuleHandler interface {
	CreateModule(w http.ResponseWriter, r *http.Request)
	GetModules(w http.ResponseWriter, r *http.Request)
	RenameModule(w http.ResponseWriter, r *http.Request)
}

func New(
	createModulecase modulecase.CreateModuleUseCase,
	getModulesCase modulecase.GetModulesUseCase,
	renameModuleCase modulecase.RenameModuleUseCase,
) ModuleHandler {
	return &moduleHandler{
		CreateModuleUseCase: createModulecase,
		GetModulesUseCase:   getModulesCase,
		RenameModuleUseCase: renameModuleCase,
	}
}
