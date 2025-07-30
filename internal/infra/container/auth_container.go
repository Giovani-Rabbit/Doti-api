package container

import (
	authcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/auth"
	authhandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/auth"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

func (c *container) NewAuth() authhandler.AuthHandler {
	userRepository := repository.NewUserRepository(c.DB)

	signIncase := authcase.NewLoginUseCase(userRepository)

	return authhandler.NewAuthHandler(
		signIncase,
	)
}
