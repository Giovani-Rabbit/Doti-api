package modulehandler

import (
	"context"
	"net/http"

	moduledomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/module"
	moduledto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/module/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/src/infra/http/responder"
)

func (mh *ModuleHandler) CreateModule(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	var createModuleDTO moduledto.CreateModuleDTO
	if !res.DecodeJSONBody(r, &createModuleDTO) {
		return
	}

	moduleEntity := moduledomain.NewCreateModule(
		"146681af-2cee-493a-a145-d23609ae056d", // get the user Foreign key from jwt
		createModuleDTO.Name,
		createModuleDTO.Icon,
	)

	ctx := context.Background()
	module, err := mh.CreateModuleUseCase.Execute(ctx, moduleEntity)

	if err != nil {
		res.Error(err, 400)
		return
	}

	moduleResponse := moduledto.NewModuleCreatedResponse(module)
	res.AddBody(moduleResponse)
	res.Write(201)
}
