package modulehandler

import (
	"net/http"

	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type NewModuleNameDTO struct {
	Name string `json:"name"`
}

func (mh *moduleHandler) RenameModule(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	res := resp.NewHttpJSONResponse(w)

	var moduleName NewModuleNameDTO
	if !res.DecodeJSONBody(r, &moduleName) {
		return
	}

	err := mh.RenameModuleUseCase.Execute(r.Context(), id, moduleName.Name)

	if err != nil {
		res.Error(err, 400)
		return
	}

	res.Write(200)
}
