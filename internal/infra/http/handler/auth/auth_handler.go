package authhandler

import (
	"net/http"

	authcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/auth"
)

type AuthHandler struct {
	SignInUseCase authcase.SignInUseCase
}

type IAuthHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(
	signInUseCase authcase.SignInUseCase,
) *AuthHandler {
	return &AuthHandler{
		SignInUseCase: signInUseCase,
	}
}
