package container

import (
	usercase "github.com/Giovani-Coelho/Doti-API/src/core/app/user"
	userhandler "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
)

func (c *Container) NewUserContainer() *userhandler.UserHandler {
	userRepository := repository.NewUserRepository(c.DB)

	createUserUseCase := usercase.NewCreateUserUseCase(userRepository)

	return userhandler.NewUserHandler(
		createUserUseCase,
	)
}
