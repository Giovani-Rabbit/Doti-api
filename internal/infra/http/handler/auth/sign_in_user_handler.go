package authhandler

import (
	"net/http"

	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	authdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/auth/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

func (ah *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	var userCredentials authdto.SignInDTO
	if !res.DecodeJSONBody(r, &userCredentials) {
		return
	}

	userModel := userdomain.NewSignInUser(
		userCredentials.Email,
		userCredentials.Password,
	)

	userAuth, token, err := ah.SignInUseCase.Execute(r.Context(), userModel)

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
