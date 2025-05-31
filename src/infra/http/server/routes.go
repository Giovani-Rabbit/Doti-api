package server

import (
	"database/sql"
	"net/http"

	userUseCase "github.com/Giovani-Coelho/Doti-API/src/core/app/user/usecases"
	userController "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
)

func Routes(DB *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	// Repository
	userRepository := repository.NewUserRepository(DB)

	// UseCase
	createUserUseCase := userUseCase.NewCreateUserUseCase(userRepository)

	// Controller
	createUserController := userController.NewCreateUserController(createUserUseCase)

	router.HandleFunc("POST /users/", createUserController.Execute)

	return router
}
