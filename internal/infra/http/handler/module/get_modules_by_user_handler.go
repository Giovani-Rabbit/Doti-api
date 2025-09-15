package modulehandler

import (
	"net/http"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
)

type get struct {
	getModulesByUser modulecase.GetByUser
}

type Get interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

func NewGetHandler(
	getModulesUseCase modulecase.GetByUser,
) Get {
	return &get{
		getModulesByUser: getModulesUseCase,
	}
}

func (gh *get) Execute(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	userAuth, err := auth.GetAuthenticatedUserFromContext(r.Context())

	if err != nil {
		res.Error(err)
		return
	}

	modules, err := gh.getModulesByUser.Execute(r.Context(), userAuth.ID)

	if err != nil {
		res.Error(err)
		return
	}

	modulesListResponse := moduledto.NewModuleListDTO(modules)
	res.AddBody(modulesListResponse)
	res.Write(200)
}
