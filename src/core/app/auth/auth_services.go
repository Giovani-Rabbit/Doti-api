package authServices

import (
	"context"

	user_request "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
)

type AuthServices struct {
	UserRepository repository.IUserRepository
}

type IAuthServices interface {
	LoginUser(ctx context.Context, userDTO user_request.SignInDTO) (sqlc.User, string, error)
}

func NewAuthServices(
	userRepository repository.IUserRepository,
) IAuthServices {
	return &AuthServices{
		UserRepository: userRepository,
	}
}
