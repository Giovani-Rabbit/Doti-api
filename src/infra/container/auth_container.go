package container

import (
	authcase "github.com/Giovani-Coelho/Doti-API/src/core/app/auth"
	authhandler "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/auth"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
)

func (c *Container) NewAuthContainer() *authhandler.AuthHandler {
	userRepository := repository.NewUserRepository(c.DB)

	signIncase := authcase.NewLoginUseCase(userRepository)

	return authhandler.NewAuthHandler(
		signIncase,
	)
}
