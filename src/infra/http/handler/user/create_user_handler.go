package user

import (
	"context"
	"net/http"

	userdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	httphdl "github.com/Giovani-Coelho/Doti-API/src/infra/http/httphdl"
)

func (uc *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user userdto.CreateUserDTO

	if err := httphdl.DecodeJSONBody(r, &user); err != nil {
		httphdl.HandleError(w, err)
		return
	}

	ctx := context.Background()

	res, err := uc.CreateUserUseCase.Execute(ctx, user)
	if err != nil {
		httphdl.HandleError(w, err)
		return
	}

	httphdl.ResponseHttpJson(
		w, http.StatusCreated,
		userdto.NewUserCreatedResponse(res),
	)
}
