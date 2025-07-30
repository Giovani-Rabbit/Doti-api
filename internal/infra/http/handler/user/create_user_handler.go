package userhandler

import (
	"net/http"

	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	userdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/user/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

func (uc *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	var userDto userdto.CreateUserDTO
	if !res.DecodeJSONBody(r, &userDto) {
		return
	}

	userDomain := userdomain.NewCreateUser(
		userDto.Name,
		userDto.Email,
		userDto.Password,
	)

	user, err := uc.CreateUserUseCase.Execute(r.Context(), userDomain)

	if err != nil {
		res.Error(err, 400)
		return
	}

	res.AddBody(userdto.NewUserCreatedResponse(user))
	res.Write(201)
}
