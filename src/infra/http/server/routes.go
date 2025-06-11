package server

import (
	"database/sql"
	"net/http"

	"github.com/Giovani-Coelho/Doti-API/src/infra/container"
)

func Routes(DB *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	appContainer := container.NewContainer(DB)

	userController := appContainer.NewUserContainer()

	router.HandleFunc("POST /users/", userController.CreateUser)

	return router
}
