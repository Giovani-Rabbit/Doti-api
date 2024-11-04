package useCase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database/repository"
)

type AddUserUseCase struct {
	UserRepository repository.IUserRepository
}

func NewCreateUserUseCase(
	userRepository repository.IUserRepository,
) *AddUserUseCase {
	return &AddUserUseCase{
		UserRepository: userRepository,
	}
}

func (a *AddUserUseCase) Execute(ctx context.Context) error {
	err := a.UserRepository.Create(ctx)

	if err != nil {
		panic(err)
	}

	return nil
}
