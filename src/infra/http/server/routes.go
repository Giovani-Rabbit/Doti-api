package server

import (
	"database/sql"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/src/infra/container"
	"github.com/Giovani-Coelho/Doti-API/src/infra/http/middleware"
)

func Routes(DB *sql.DB) (mux *http.ServeMux) {
	mux = http.NewServeMux()

	appContainer := container.NewContainer(DB)

	userHandler := appContainer.NewUserContainer()
	authHandler := appContainer.NewAuthContainer()
	moduleHandler := appContainer.NewModuleContainer()

	// USER
	mux.HandleFunc("POST /users", userHandler.CreateUser)

	// AUTH
	mux.HandleFunc("POST /sign-in", authHandler.SignIn)

	// MODULE
	mux.Handle("POST /module",
		middleware.EnsureAuth(http.HandlerFunc(
			moduleHandler.CreateModule,
		)),
	)

	return
}
