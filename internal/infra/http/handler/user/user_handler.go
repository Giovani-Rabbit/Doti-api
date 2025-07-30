package userhandler

import (
	"net/http"

	usercase "github.com/Giovani-Coelho/Doti-API/internal/core/app/user"
)

type UserHandler struct {
	CreateUserUseCase usercase.CreateUserUseCase
}

type IUserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(
	createUserUseCase usercase.CreateUserUseCase,
) *UserHandler {
	return &UserHandler{
		CreateUserUseCase: createUserUseCase,
	}
}
