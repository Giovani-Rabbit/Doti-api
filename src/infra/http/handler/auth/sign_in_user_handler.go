package authhandler

import (
	"context"
	"net/http"

	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	authdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/auth/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/src/infra/http/responder"
)

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (ah *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	var userCredentials authdto.SignInDTO
	if !res.DecodeJSONBody(r, &userCredentials) {
		return
	}

	user := userdomain.NewSignInUserDomain(
		userCredentials.Email,
		userCredentials.Password,
	)

	ctx := context.Background()

	userDomain, token, err := ah.SignInUseCase.Execute(ctx, user)

	if err != nil {
		res.Error(err, 400)
		return
	}

	response := UserResponse{
		ID:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
	}

	res.AddHeader("authorization", token)
	res.AddBody(response)
	res.Write(200)
}
