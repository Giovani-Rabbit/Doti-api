package modulehandler

import (
	"net/http"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type rename struct {
	renameModule modulecase.Rename
}

type Rename interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

func NewRenameHandler(
	renameModuleUseCase modulecase.Rename,
) Rename {
	return &rename{
		renameModule: renameModuleUseCase,
	}
}

type NewModuleNameDTO struct {
	Name string `json:"name"`
}

func (rm *rename) Execute(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	res := resp.NewHttpJSONResponse(w)

	var moduleName NewModuleNameDTO
	if !res.DecodeJSONBody(r, &moduleName) {
		return
	}

	err := rm.renameModule.Execute(r.Context(), id, moduleName.Name)

	if err != nil {
		res.Error(err)
		return
	}

	res.Write(204)
}
