package userhandler

import (
	"net/http"

	usercase "github.com/Giovani-Coelho/Doti-API/internal/core/app/user"
)

type userHandler struct {
	CreateUserUseCase usercase.CreateUserUseCase
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(
	createUserUseCase usercase.CreateUserUseCase,
) UserHandler {
	return &userHandler{
		CreateUserUseCase: createUserUseCase,
	}
}
