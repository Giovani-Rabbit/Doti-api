package userServices

import (
	"context"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/application/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/repository"
)

type CreateUserService struct {
	UserRepository repository.IUserRepository
}

type ICreateUserService interface {
	CreateUser(ctx context.Context, userDTO userDTO.CreateUserDTO) error
}

func NewCreateUserService(
	userRepository repository.IUserRepository,
) ICreateUserService {
	return &CreateUserService{
		UserRepository: userRepository,
	}
}
