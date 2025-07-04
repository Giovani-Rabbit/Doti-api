package authhandler

import (
	"context"
	"net/http"

	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	authdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/auth/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/src/infra/http/responder"
)

func (ah *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	var userCredentials authdto.SignInDTO
	if !res.DecodeJSONBody(r, &userCredentials) {
		return
	}

	userModel := userdomain.NewSignInUserDomain(
		userCredentials.Email,
		userCredentials.Password,
	)

	ctx := context.Background()

	userAuth, token, err := ah.SignInUseCase.Execute(ctx, userModel)

	if err != nil {
		res.Error(err, 400)
		return
	}

	response := authdto.SignInResponseDTO{
		ID:    userAuth.GetID(),
		Name:  userAuth.GetName(),
		Email: userAuth.GetEmail(),
	}

	res.AddHeader("authorization", token)
	res.AddBody(response)
	res.Write(200)
}
