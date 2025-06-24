package container

import (
	userUseCase "github.com/Giovani-Coelho/Doti-API/src/core/app/user/usecases"
	userController "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user"
	userRepository "github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
)

func (c *Container) NewUserContainer() *userController.UserHandler {
	userRepository := userRepository.NewUserRepository(c.DB)

	createUserUseCase := userUseCase.NewCreateUserUseCase(userRepository)

	return userController.NewUserHandler(
		createUserUseCase,
	)
}
