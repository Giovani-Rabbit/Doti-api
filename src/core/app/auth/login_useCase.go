package auth

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	authDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/auth"
	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	userdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	authpkg "github.com/Giovani-Coelho/Doti-API/src/pkg/auth"
	"go.uber.org/zap"
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
	logger.Info("Init Login UseCase",
		zap.String("journey", "login"),
	)

	user, err := lu.UserRepository.FindUserByEmailAndPassword(ctx, userDTO)

	if err != nil {
		logger.Error(
			"Error: Could not find user with those credentials", err,
			zap.String("journey", "login"),
		)

		return nil, "", userDomain.ErrCouldNotFindUser()
	}

	token, err := authpkg.GenerateToken(user)

	if err != nil {
		logger.Error(
			"Error: Could not generate token", err,
			zap.String("journey", "login"),
		)

		return nil, "", authDomain.ErrGeneratingToken()
	}

	logger.Info(
		"Login executed successfully",
		zap.String("token:", token),
		zap.String("journey", "login"),
	)

	return user, token, nil
}
