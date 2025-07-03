package authhandler

import (
	"context"
	"encoding/json"
	"net/http"

	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	authdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/auth/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/http/httphdl"
)

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (ah *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var userCredentials authdto.SignInDTO

	if err := httphdl.DecodeJSONBody(r, &userCredentials); err != nil {
		httphdl.HandleError(w, err)
		return
	}

	ctx := context.Background()

	user := userdomain.NewSignInUserDomain(
		userCredentials.Email,
		userCredentials.Password,
	)

	userDomain, token, err := ah.SignInUseCase.Execute(ctx, user)

	if err != nil {
		httphdl.HandleError(w, err)
		return
	}

	response := UserResponse{
		ID:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Authorization", token)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(response)
}
