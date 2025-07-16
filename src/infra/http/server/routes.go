package server

import (
	"database/sql"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/src/infra/container"
)

func Routes(DB *sql.DB) (router *http.ServeMux) {
	router = http.NewServeMux()

	appContainer := container.NewContainer(DB)

	userHandler := appContainer.NewUserContainer()
	authHandler := appContainer.NewAuthContainer()
	moduleHandler := appContainer.NewModuleContainer()

	// USER
	router.HandleFunc("POST /users", userHandler.CreateUser)

	// AUTH
	router.HandleFunc("POST /sign-in", authHandler.SignIn)

	// MODULE
	router.HandleFunc("POST /module", moduleHandler.CreateModule)

	return
}
