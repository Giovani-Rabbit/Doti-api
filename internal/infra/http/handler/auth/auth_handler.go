package authhandler

import (
	"net/http"

	authcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/auth"
)

type authHandler struct {
	SignInUseCase authcase.SignInUseCase
}

type AuthHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(
	signInUseCase authcase.SignInUseCase,
) AuthHandler {
	return &authHandler{
		SignInUseCase: signInUseCase,
	}
}
