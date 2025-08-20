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
	createModule modulecase.Create
}

func NewCreateModuleHandler(
	createModuleUseCase modulecase.Create,
) Create {
	return &create{
		createModule: createModuleUseCase,
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
		res.Error(err)
		return
	}

	moduleEntity := moduledomain.NewCreateModule(
		userAuth.ID,
		createModuleDTO.Name,
		createModuleDTO.Icon,
	)

	module, err := mh.createModule.Execute(r.Context(), moduleEntity)

	if err != nil {
		res.Error(err)
		return
	}

	moduleResponse := moduledto.NewModuleCreatedResponse(module)
	res.AddBody(moduleResponse)
	res.Write(201)
}
