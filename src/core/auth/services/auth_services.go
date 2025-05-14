package authServices

import (
	"context"

	authDTO "github.com/Giovani-Coelho/Doti-API/src/core/auth/dto"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
)

type AuthServices struct {
	UserRepository repository.IUserRepository
}

type IAuthServices interface {
	LoginUser(ctx context.Context, userDTO authDTO.SignInDTO) (sqlc.User, string, error)
}

func NewAuthServices(
	userRepository repository.IUserRepository,
) IAuthServices {
	return &AuthServices{
		UserRepository: userRepository,
	}
}
