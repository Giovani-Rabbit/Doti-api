package server

import (
	"database/sql"

	authServices "github.com/Giovani-Coelho/Doti-API/src/application/auth/services"
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
	authServices := authServices.NewAuthServices(userRepository)

	// Controller
	userController := userController.NewUserControllers(userServices, authServices)

	router.HandleFunc("/user", userController.CreateUser).Methods("POST")
	router.HandleFunc("/login", userController.LoginUser).Methods("POST")

	return router
}
