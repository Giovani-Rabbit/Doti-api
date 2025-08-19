package authcase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	authpkg "github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
	"go.uber.org/zap"
)

type signIn struct {
	userRepository repository.UserRepository
}

type SignIn interface {
	Execute(
		ctx context.Context, userEntity userdomain.User,
	) (userdomain.User, string, error)
}

func NewLoginUseCase(
	userRepository repository.UserRepository,
) SignIn {
	return &signIn{
		userRepository: userRepository,
	}
}

func (su *signIn) Execute(
	ctx context.Context,
	userEntity userdomain.User,
) (userdomain.User, string, error) {
	logger.Info("Init Sign-In UseCase",
		zap.String("journey", "sign-in"),
	)

	if userEntity.GetEmail() == "" || userEntity.GetPassword() == "" {
		logger.Error(
			"Error: Email or Password is missing", nil,
			zap.String("journey", "sign-in"),
		)

		return nil, "", userdomain.ErrSignInValuesMissing()
	}

	if !userEntity.IsValidEmail() {
		logger.Error(
			"Error: Invalid e-mail format", nil,
			zap.String("journey", "sign-in"),
		)

		return nil, "", userdomain.ErrInvalidUserEmailFormat()
	}

	userEntity.EncryptPassword()
	user, err := su.userRepository.FindUserByEmailAndPassword(ctx, userEntity)

	if err != nil {
		logger.Error(
			"Error: Could not find user with those credentials", err,
			zap.String("journey", "sign-in"),
		)

		return nil, "", userdomain.ErrCouldNotFindUser()
	}

	token, err := authpkg.GenerateToken(user)

	if err != nil {
		logger.Error(
			"Error: Could not generate token", err,
			zap.String("journey", "sign-in"),
		)

		return nil, "", userdomain.ErrGeneratingToken()
	}

	logger.Info(
		"sign-in executed successfully",
		zap.String("token:", token),
		zap.String("journey", "sign-in"),
	)

	return user, token, nil
}
