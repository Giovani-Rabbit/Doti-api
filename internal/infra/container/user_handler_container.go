package container

import (
	authhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/auth"
	userhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/user"
)

type UserHandler struct {
	Create userhandler.Create
	SignIn authhandler.SignIn
}

func newUserHandler(userCase *UserCase) *UserHandler {
	return &UserHandler{
		Create: userhandler.NewCreateHandler(userCase.create),
		SignIn: authhandler.NewSignInHandler(userCase.signIn),
	}
}
