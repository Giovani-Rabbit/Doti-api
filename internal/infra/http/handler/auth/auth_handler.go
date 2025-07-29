package authhandler

import (
	"net/http"

	authcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/auth"
)

type AuthHandler struct {
	SignInUseCase authcase.ISignInUseCase
}

type IAuthHandler interface {
	SignIn(w http.ResponseWriter, r *http.Request)
}

func NewAuthHandler(
	signInUseCase authcase.ISignInUseCase,
) *AuthHandler {
	return &AuthHandler{
		SignInUseCase: signInUseCase,
	}
}
