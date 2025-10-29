package detailshdl

import (
	"net/http"
	"strconv"

	detailscase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task_details"
	detailsdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task_details/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type UpdatePomodoroTarget interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type updatePomodoroTarget struct {
	updatePomodoroTargetCase detailscase.UpdatePomodoroTarget
}

func NewUpdatePomodoroTargetHandler(
	updatePomodoroTargetCase detailscase.UpdatePomodoroTarget,
) UpdatePomodoroTarget {
	return &updatePomodoroTarget{
		updatePomodoroTargetCase: updatePomodoroTargetCase,
	}
}

func (ud *updatePomodoroTarget) Execute(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)
	taskIdStr := r.PathValue("id")

	taskId, err := strconv.ParseInt(taskIdStr, 10, 32)
	if err != nil {
		res.Error(resp.NewInvalidPathValue(err))
		return
	}

	var body *detailsdto.UpdatePomodoroTargetRequest
	if !res.DecodeJSONBody(r, &body) {
		return
	}

	err = ud.updatePomodoroTargetCase.Execute(
		r.Context(), int32(taskId), body.PomodoroTarget,
	)
	if err != nil {
		res.Error(err)
		return
	}

	res.Write(http.StatusNoContent)
}
