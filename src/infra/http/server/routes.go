package server

import (
	"database/sql"

	userServices "github.com/Giovani-Coelho/Doti-API/src/application/user/services"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/repository"
	userController "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user"
	"github.com/gorilla/mux"
)

func Routes(DB *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Repository
	userRepository := repository.NewUserRepository(DB)

	// UseCase
	userServices := userServices.NewUserServices(userRepository)

	// Controller
	userController := userController.NewUserControllers(userServices)

	router.HandleFunc("/user", userController.CreateUser).Methods("POST")

	return router
}
