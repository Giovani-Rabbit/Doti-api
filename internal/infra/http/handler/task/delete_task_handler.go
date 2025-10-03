package taskhdl

import (
	"net/http"
	"strconv"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type Delete interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type delete struct {
	deletecase taskcase.Delete
}

func NewDeleteTaskHandler(
	deletecase taskcase.Delete,
) Delete {
	return &delete{
		deletecase: deletecase,
	}
}

func (dt *delete) Execute(w http.ResponseWriter, r *http.Request) {
	taskIdStr := r.PathValue("id")
	res := resp.NewHttpJSONResponse(w)

	taskId, err := strconv.ParseInt(taskIdStr, 10, 32)
	if err != nil {
		res.Error(resp.NewInvalidPathValue(err))
		return
	}

	err = dt.deletecase.Execute(r.Context(), int32(taskId))
	if err != nil {
		res.Error(err)
		return
	}

	res.Write(http.StatusNoContent)
}
