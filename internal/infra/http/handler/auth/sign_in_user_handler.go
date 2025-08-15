package authhandler

import (
	"net/http"

	authcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/auth"
	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	authdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/auth/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

type signIn struct {
	signInCase authcase.SignInUseCase
}

type SignIn interface {
	Execute(w http.ResponseWriter, r *http.Request)
}

func NewSignInHandler(
	signInUseCase authcase.SignInUseCase,
) SignIn {
	return &signIn{
		signInCase: signInUseCase,
	}
}

func (sh *signIn) Execute(w http.ResponseWriter, r *http.Request) {
	res := resp.NewHttpJSONResponse(w)

	var userCredentials authdto.SignInDTO
	if !res.DecodeJSONBody(r, &userCredentials) {
		return
	}

	userModel := userdomain.NewSignInUser(
		userCredentials.Email,
		userCredentials.Password,
	)

	userAuth, token, err := sh.signInCase.Execute(r.Context(), userModel)

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
