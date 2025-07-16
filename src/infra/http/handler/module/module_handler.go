package modulehandler

import (
	"net/http"

	modulecase "github.com/Giovani-Coelho/Doti-API/src/core/app/module"
)

type ModuleHandler struct {
	CreateModuleUseCase modulecase.ICreateModuleUseCase
}

type IModuleHandler interface {
	CreateModule(w http.ResponseWriter, r *http.Request)
}

func New(
	createModulecase modulecase.ICreateModuleUseCase,
) IModuleHandler {
	return &ModuleHandler{
		CreateModuleUseCase: createModulecase,
	}
}
