package modulehandler

import (
	"net/http"

	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
)

func (mh *moduleHandler) GetModules(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	userAuth, err := auth.GetAuthenticatedUserFromContext(r.Context())

	if err != nil {
		res.Error(err, 400)
		return
	}

	modules, err := mh.GetModulesUseCase.Execute(r.Context(), userAuth.ID)

	if err != nil {
		res.Error(err, 400)
	}

	modulesListResponse := moduledto.NewModuleListDTO(modules)
	res.AddBody(modulesListResponse)
	res.Write(200)
}
