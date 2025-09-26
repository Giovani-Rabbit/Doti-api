package taskhdl

import (
	"net/http"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type UpdatePosition interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type updatePosition struct {
	updatePosition taskcase.UpdatePosition
}

func NewUpdatePositionHandler(
	updatePositionCase taskcase.UpdatePosition,
) UpdatePosition {
	return &updatePosition{
		updatePosition: updatePositionCase,
	}
}

func (up *updatePosition) Execute(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	var tasks taskdto.UpdatePositionDTO
	if !res.DecodeJSONBody(r, &tasks) {
		return
	}

	err := up.updatePosition.Execute(r.Context(), &tasks)
	if err != nil {
		res.Error(err)
		return
	}

	res.Write(http.StatusNoContent)
}
