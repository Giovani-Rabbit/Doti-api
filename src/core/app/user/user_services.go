package userServices

import (
	"context"

	// userDTO "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
)

type UserServices struct {
	UserRepository repository.IUserRepository
}

type IUserServices interface {
	findUserByEmailAndPassword(ctx context.Context, email string, password string) (sqlc.User, error)
}

func NewUserServices(
	userRepository repository.IUserRepository,
) IUserServices {
	return &UserServices{
		UserRepository: userRepository,
	}
}
