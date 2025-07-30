package modulehandler

import (
	"net/http"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
)

type moduleHandler struct {
	CreateModuleUseCase modulecase.CreateModuleUseCase
}

type ModuleHandler interface {
	CreateModule(w http.ResponseWriter, r *http.Request)
}

func New(
	createModulecase modulecase.CreateModuleUseCase,
) ModuleHandler {
	return &moduleHandler{
		CreateModuleUseCase: createModulecase,
	}
}
