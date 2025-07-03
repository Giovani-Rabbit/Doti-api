package userhandler

import (
	"context"
	"net/http"

	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	userdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	httphdl "github.com/Giovani-Coelho/Doti-API/src/infra/http/httphdl"
)

func (uc *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto userdto.CreateUserDTO

	if err := httphdl.DecodeJSONBody(r, &userDto); err != nil {
		httphdl.HandleError(w, err)
		return
	}

	ctx := context.Background()

	userDomain := userdomain.NewCreateUserDomain(
		userDto.Name,
		userDto.Email,
		userDto.Password,
	)

	res, err := uc.CreateUserUseCase.Execute(ctx, userDomain)

	if err != nil {
		httphdl.HandleError(w, err)
		return
	}

	httphdl.ResponseHttpJson(
		w, http.StatusCreated,
		userdto.NewUserCreatedResponse(res),
	)
}
