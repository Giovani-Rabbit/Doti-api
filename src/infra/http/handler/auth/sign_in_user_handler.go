package authhandler

import (
	"context"
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

	res := httphdl.NewHttpJSONResponse(w)

	if err := httphdl.DecodeJSONBody(r, &userCredentials); err != nil {
		res.Error(err, 400)
		return
	}

	ctx := context.Background()

	user := userdomain.NewSignInUserDomain(
		userCredentials.Email,
		userCredentials.Password,
	)

	userDomain, token, err := ah.SignInUseCase.Execute(ctx, user)

	if err != nil {
		res.Error(err, 500)
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
