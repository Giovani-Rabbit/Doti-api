package userhandler

import (
	"net/http"

	usercase "github.com/Giovani-Coelho/Doti-API/internal/core/app/user"
	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	userdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/user/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type create struct {
	createUser usercase.Create
}

type Create interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

func NewCreateHandler(
	createUserUseCase usercase.Create,
) *create {
	return &create{
		createUser: createUserUseCase,
	}
}

func (cu *create) Execute(w http.ResponseWriter, r *http.Request) {
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

	user, err := cu.createUser.Execute(r.Context(), userDomain)

	if err != nil {
		res.Error(err)
		return
	}

	res.AddBody(userdto.NewUserCreatedResponse(user))
	res.Write(201)
}
