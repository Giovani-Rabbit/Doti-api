package server

import (
	"database/sql"

	authServices "github.com/Giovani-Coelho/Doti-API/src/core/app/auth"
	userServices "github.com/Giovani-Coelho/Doti-API/src/core/app/user"
	userController "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/http/middleware"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
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

	router.Use(middleware.CORSMiddleware)

	router.HandleFunc("/account", userController.CreateUser)
	router.HandleFunc("/login", userController.LoginUser)

	return router
}
