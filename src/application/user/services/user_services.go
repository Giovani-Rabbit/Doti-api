package userServices

import (
	"context"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/application/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/repository"
)

type UserServices struct {
	UserRepository repository.IUserRepository
}

type IUserServices interface {
	CreateUser(ctx context.Context, userDTO userDTO.CreateUserDTO) error
	FindUserByEmail(ctx context.Context, email string) (sqlc.User, error)
}

func NewUserServices(
	userRepository repository.IUserRepository,
) IUserServices {
	return &UserServices{
		UserRepository: userRepository,
	}
}
