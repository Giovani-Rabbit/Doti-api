package server

import (
	"database/sql"

	userServices "github.com/Giovani-Coelho/Doti-API/src/application/user/services/createUser"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/repository"
	userController "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user"
	"github.com/gorilla/mux"
)

func Routes(DB *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Repository
	userR := repository.NewUserRepository(DB)

	// UseCase
	createUserService := userServices.NewCreateUserService(userR)

	// Controller
	createUserController := userController.NewCreateUserController(createUserService)

	router.HandleFunc("/user", createUserController.Execute)

	return router
}
