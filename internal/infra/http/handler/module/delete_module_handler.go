package modulehandler

import (
	"net/http"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	httperr "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type delete struct {
	deleteCase modulecase.DeleteModuleUseCase
}

type Delete interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

func NewDeleteHandler(
	deleteUseCase modulecase.DeleteModuleUseCase,
) Delete {
	return &delete{
		deleteCase: deleteUseCase,
	}
}

func (dm *delete) Execute(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	res := resp.NewHttpJSONResponse(w)

	err := dm.deleteCase.Execute(r.Context(), id)

	if err != nil {
		resterr, ok := err.(*httperr.RestErr)

		if !ok {
			res.Error(err, 500)
			return
		}

		res.Error(resterr, resterr.Code)
		return
	}

	res.Write(200)
}
