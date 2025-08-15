package server

import (
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/container"
)

func UserRoutes(mux *http.ServeMux, user *container.UserHandler) (m *http.ServeMux) {
	mux.HandleFunc("POST /users", user.Create.Execute)
	mux.HandleFunc("POST /sign-in", user.SignIn.Execute)

	return
}
