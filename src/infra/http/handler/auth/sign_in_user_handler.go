package authhandler

import (
	"context"
	"fmt"
	"net/http"

	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	authdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/auth/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/http/httphdl"
)

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

	user, token, err := ah.SignInUseCase.Execute(ctx, user)

	fmt.Println(user)
	fmt.Println(token)
	fmt.Println(err)
}
