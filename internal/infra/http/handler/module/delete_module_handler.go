package modulehandler

import (
	"net/http"

	httperr "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

func (mh *moduleHandler) DeleteModule(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	res := resp.NewHttpJSONResponse(w)

	err := mh.DeleteModuleUseCase.Execute(r.Context(), id)

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
