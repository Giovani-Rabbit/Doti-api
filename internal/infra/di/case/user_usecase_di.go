package dicase

import (
	authcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/auth"
	usercase "github.com/Giovani-Coelho/Doti-API/internal/core/app/user"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
)

type UserUseCases struct {
	Create usercase.Create
	SignIn authcase.SignIn
}

func NewUserCases(userRepo repository.UserRepository) *UserUseCases {
	return &UserUseCases{
		Create: usercase.NewCreateUserUseCase(userRepo),
		SignIn: authcase.NewLoginUseCase(userRepo),
	}
}
