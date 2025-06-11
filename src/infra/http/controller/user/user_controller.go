package userController

import (
	"net/http"

	userUseCase "github.com/Giovani-Coelho/Doti-API/src/core/app/user/usecases"
)

type UserController struct {
	CreateUserUseCase userUseCase.ICreateUserUseCase
}

type IUserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserController(
	createUserUseCase userUseCase.ICreateUserUseCase,
) *UserController {
	return &UserController{
		CreateUserUseCase: createUserUseCase,
	}
}
