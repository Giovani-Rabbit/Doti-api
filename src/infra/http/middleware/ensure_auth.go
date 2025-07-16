package middleware

import (
	"fmt"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/src/pkg/auth"
)

type EnsureAuth struct {
	http http.Handler
}

func (ea *EnsureAuth) ServerHttp(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetAuthenticatedUser(r)

	if err != nil {
		return
	}

	fmt.Println(user)
}

func NewEnsureAuth(handlerToWrap http.Handler) *EnsureAuth {
	return &EnsureAuth{handlerToWrap}
}
