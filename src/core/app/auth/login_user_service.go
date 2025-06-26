package auth

import (
	"context"

	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	userdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	"github.com/Giovani-Coelho/Doti-API/src/pkg/auth"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

type LoginUseCase struct {
	UserRepository repository.IUserRepository
}

type ILoginUseCase interface {
	Execute(ctx context.Context, userDTO userdto.SignInDTO) (userDomain.IUserDomain, string, error)
}

func NewLoginUseCase(
	userRepository repository.IUserRepository,
) ILoginUseCase {
	return &LoginUseCase{
		UserRepository: userRepository,
	}
}

func (lu *LoginUseCase) Execute(
	ctx context.Context,
	userDTO userdto.SignInDTO,
) (userDomain.IUserDomain, string, error) {
	user, err := lu.UserRepository.FindUserByEmailAndPassword(ctx, userDTO)

	if err != nil {
		return nil, "", rest_err.NewBadRequestError(
			"UNKNOWN",
			"Invalid login credentials",
		)
	}

	token, err := auth.GenerateToken(user)

	if err != nil {
		return nil, "", rest_err.NewBadRequestError(
			"UNKNOWN",
			"Error on generate Token",
		)
	}

	return user, token, nil
}
