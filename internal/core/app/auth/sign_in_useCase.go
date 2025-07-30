package authcase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	authdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/auth"
	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	authpkg "github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"
	"go.uber.org/zap"
)

type SignInUseCase struct {
	UserRepository repository.IUserRepository
}

type ISignInUseCase interface {
	Execute(
		ctx context.Context, userEntity userdomain.User,
	) (userdomain.User, string, *http.RestErr)
}

func NewLoginUseCase(
	userRepository repository.IUserRepository,
) ISignInUseCase {
	return &SignInUseCase{
		UserRepository: userRepository,
	}
}

func (su *SignInUseCase) Execute(
	ctx context.Context,
	userEntity userdomain.User,
) (userdomain.User, string, *http.RestErr) {
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
	user, err := su.UserRepository.FindUserByEmailAndPassword(ctx, userEntity)

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

		return nil, "", authdomain.ErrGeneratingToken()
	}

	logger.Info(
		"sign-in executed successfully",
		zap.String("token:", token),
		zap.String("journey", "sign-in"),
	)

	return user, token, nil
}
