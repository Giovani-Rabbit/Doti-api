package container

import (
	authcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/auth"
	usercase "github.com/Giovani-Coelho/Doti-API/internal/core/app/user"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type UserCase struct {
	create usercase.CreateUserUseCase
	signIn authcase.SignInUseCase
}

func newUserCase(userRepo repository.UserRepository) *UserCase {
	return &UserCase{
		create: usercase.NewCreateUserUseCase(userRepo),
		signIn: authcase.NewLoginUseCase(userRepo),
	}
}
