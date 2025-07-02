package container

import (
	usercase "github.com/Giovani-Coelho/Doti-API/src/core/app/user/usecases"
	userhandler "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user"
	userRepository "github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
)

func (c *Container) NewUserContainer() *userhandler.UserHandler {
	userRepository := userRepository.NewUserRepository(c.DB)

	createUserUseCase := usercase.NewCreateUserUseCase(userRepository)

	return userhandler.NewUserHandler(
		createUserUseCase,
	)
}
