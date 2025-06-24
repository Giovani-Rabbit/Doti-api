package userhdl

import (
	"net/http"

	userUseCase "github.com/Giovani-Coelho/Doti-API/src/core/app/user/usecases"
)

type UserHandler struct {
	CreateUserUseCase userUseCase.ICreateUserUseCase
}

type IUserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(
	createUserUseCase userUseCase.ICreateUserUseCase,
) *UserHandler {
	return &UserHandler{
		CreateUserUseCase: createUserUseCase,
	}
}
