package server

import (
	"database/sql"
	"net/http"

	authServices "github.com/Giovani-Coelho/Doti-API/src/core/app/auth"
	userServices "github.com/Giovani-Coelho/Doti-API/src/core/app/user"
	userController "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
)

func Routes(DB *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	// Repository
	userRepository := repository.NewUserRepository(DB)

	// UseCase
	userServices := userServices.NewUserServices(userRepository)
	authServices := authServices.NewAuthServices(userRepository)

	// Controller
	userController := userController.NewUserControllers(userServices, authServices)

	router.HandleFunc("/account", userController.CreateUser)
	router.HandleFunc("/login", userController.LoginUser)

	return router
}
