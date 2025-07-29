package modulehandler

import (
	"net/http"

	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
)

func (mh *ModuleHandler) CreateModule(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	var createModuleDTO moduledto.CreateModuleDTO
	if !res.DecodeJSONBody(r, &createModuleDTO) {
		return
	}

	userAuth, err := auth.GetAuthenticatedUserFromContext(r.Context())

	if err != nil {
		res.Error(err, 400)
		return
	}

	moduleEntity := moduledomain.NewCreateModule(
		userAuth.ID,
		createModuleDTO.Name,
		createModuleDTO.Icon,
	)

	module, err := mh.CreateModuleUseCase.Execute(r.Context(), moduleEntity)

	if err != nil {
		res.Error(err, 400)
		return
	}

	moduleResponse := moduledto.NewModuleCreatedResponse(module)
	res.AddBody(moduleResponse)
	res.Write(201)
}
