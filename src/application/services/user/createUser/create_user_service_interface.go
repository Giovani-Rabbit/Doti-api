package userServices

import (
	"context"

	dtos "github.com/Giovani-Coelho/Doti-API/src/application/dtos/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/repository"
)

type CreateUserService struct {
	UserRepository repository.IUserRepository
}

type ICreateUserService interface {
	CreateUser(ctx context.Context, userDTO dtos.CreateUserDto) error
}

func NewCreateUserService(
	userRepository repository.IUserRepository,
) ICreateUserService {
	return &CreateUserService{
		UserRepository: userRepository,
	}
}
