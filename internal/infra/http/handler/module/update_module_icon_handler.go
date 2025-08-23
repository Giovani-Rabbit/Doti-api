package modulehandler

import (
	"net/http"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type updateIcon struct {
	updateIconUseCase modulecase.UpdateIcon
}

type UpdateIcon interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

func NewUpdateIconHandler(
	updateIconCase modulecase.UpdateIcon,
) UpdateIcon {
	return &updateIcon{
		updateIconUseCase: updateIconCase,
	}
}

func (ui *updateIcon) Execute(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	res := resp.NewHttpJSONResponse(w)

	var input *moduledto.UpdateIconDTO
	if !res.DecodeJSONBody(r, &input) {
		return
	}

	err := ui.updateIconUseCase.Execute(r.Context(), id, input.Icon)

	if err != nil {
		res.Error(err)
		return
	}

	res.Write(204)
}
