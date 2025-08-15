package server

import (
	"net/http"

	dihdl "github.com/Giovani-Coelho/Doti-API/internal/infra/di/handler"
)

func UserRoutes(mux *http.ServeMux, user *dihdl.UserHandlers) (m *http.ServeMux) {

	mux.HandleFunc("POST /users", user.Create.Execute)
	mux.HandleFunc("POST /sign-in", user.SignIn.Execute)

	return
}
