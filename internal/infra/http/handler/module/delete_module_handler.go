package modulehandler

import (
	"net/http"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"

	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type delete struct {
	deleteModule modulecase.Delete
}

type Delete interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

func NewDeleteHandler(
	deleteModuleUseCase modulecase.Delete,
) Delete {
	return &delete{
		deleteModule: deleteModuleUseCase,
	}
}

func (dm *delete) Execute(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	res := resp.NewHttpJSONResponse(w)

	err := dm.deleteModule.Execute(r.Context(), id)

	if err != nil {
		resterr, ok := err.(*resp.RestErr)

		if !ok {
			res.Error(err, 500)
			return
		}

		res.Error(resterr, resterr.Code)
		return
	}

	res.Write(204)
}
