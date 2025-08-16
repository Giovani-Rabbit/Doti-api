package modulehandler

import (
	"net/http"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
)

type Create interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type create struct {
	createModuleUseCase modulecase.Create
}

func NewCreateModuleHandler(
	createModulecase modulecase.Create,
) Create {
	return &create{
		createModuleUseCase: createModulecase,
	}
}

func (mh *create) Execute(w http.ResponseWriter, r *http.Request) {
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

	module, err := mh.createModuleUseCase.Execute(r.Context(), moduleEntity)

	if err != nil {
		res.Error(err, 400)
		return
	}

	moduleResponse := moduledto.NewModuleCreatedResponse(module)
	res.AddBody(moduleResponse)
	res.Write(201)
}
