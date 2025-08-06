package server

import (
	"database/sql"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/internal/infra/container"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/http/middleware"
)

func Routes(DB *sql.DB) (mux *http.ServeMux) {
	mux = http.NewServeMux()

	appContainer := container.NewContainer(DB)

	userHandler := appContainer.NewUser()
	authHandler := appContainer.NewAuth()
	moduleHandler := appContainer.NewModule()

	// USER
	mux.HandleFunc("POST /users", userHandler.CreateUser)

	// AUTH
	mux.HandleFunc("POST /sign-in", authHandler.SignIn)

	// MODULE
	mux.Handle("POST /module/",
		middleware.EnsureAuth(http.HandlerFunc(
			moduleHandler.CreateModule,
		)),
	)

	mux.Handle("GET /module/",
		middleware.EnsureAuth(http.HandlerFunc(
			moduleHandler.GetModules,
		)),
	)

	mux.Handle("PATCH /module/{id}",
		middleware.EnsureAuth(http.HandlerFunc(
			moduleHandler.RenameModule,
		)),
	)

	return
}
