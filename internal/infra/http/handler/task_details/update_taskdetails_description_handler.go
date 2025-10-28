package detailshdl

import (
	"net/http"
	"strconv"

	detailscase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task_details"
	detailsdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task_details/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type UpdateDescription interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

type updateDescription struct {
	updateDescription detailscase.UpdateDescription
}

func NewUpdateDescriptionHandler(
	updateDescriptioncase detailscase.UpdateDescription,
) UpdateDescription {
	return &updateDescription{
		updateDescription: updateDescriptioncase,
	}
}

func (ud *updateDescription) Execute(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)
	taskIdStr := r.PathValue("id")

	taskId, err := strconv.ParseInt(taskIdStr, 10, 32)
	if err != nil {
		res.Error(resp.NewInvalidPathValue(err))
		return
	}

	var body *detailsdto.UpdateDescriptionRequest
	if res.DecodeJSONBody(r, &body) {
		return
	}

	err = ud.updateDescription.Execute(r.Context(), int32(taskId), body.Description)
	if err != nil {
		res.Error(err)
		return
	}

	res.Write(http.StatusNoContent)
}
