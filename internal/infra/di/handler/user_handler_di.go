package dihdl

import (
	dicase "github.com/Giovani-Coelho/Doti-API/internal/infra/di/case"
	authhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/auth"
	userhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/user"
)

type UserHandlers struct {
	Create userhandler.Create
	SignIn authhandler.SignIn
}

func NewUserHandlers(userCase *dicase.UserUseCases) *UserHandlers {
	return &UserHandlers{
		Create: userhandler.NewCreateHandler(userCase.Create),
		SignIn: authhandler.NewSignInHandler(userCase.SignIn),
	}
}
