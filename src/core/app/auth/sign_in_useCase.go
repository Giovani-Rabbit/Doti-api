package authcase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	authdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/auth"
	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	authdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/auth/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	authpkg "github.com/Giovani-Coelho/Doti-API/src/pkg/auth"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"go.uber.org/zap"
)

type SignInUseCase struct {
	UserRepository repository.IUserRepository
}

type ISignInUseCase interface {
	Execute(
		ctx context.Context, userEntity userdomain.IUserDomain,
	) (userdomain.IUserDomain, *authdto.AuthTokenDTO, *rest_err.RestErr)
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
	userEntity userdomain.IUserDomain,
) (userdomain.IUserDomain, *authdto.AuthTokenDTO, *rest_err.RestErr) {
	logger.Info("Init Sign-In UseCase",
		zap.String("journey", "sign-in"),
	)

	if userEntity.GetEmail() == "" || userEntity.GetPassword() == "" {
		logger.Error(
			"Error: Email or Password is missing", nil,
			zap.String("journey", "sign-in"),
		)

		return nil, nil, userdomain.ErrSignInValuesMissing()
	}

	if !userEntity.IsValidEmail() {
		logger.Error(
			"Error: Invalid e-mail format", nil,
			zap.String("journey", "sign-in"),
		)

		return nil, nil, userdomain.ErrInvalidUserEmailFormat()
	}

	userEntity.EncryptPassword()
	user, err := su.UserRepository.FindUserByEmailAndPassword(ctx, userEntity)

	if err != nil {
		logger.Error(
			"Error: Could not find user with those credentials", err,
			zap.String("journey", "sign-in"),
		)

		return nil, nil, userdomain.ErrCouldNotFindUser()
	}

	token, err := authpkg.GenerateToken(user)

	if err != nil {
		logger.Error(
			"Error: Could not generate token", err,
			zap.String("journey", "sign-in"),
		)

		return nil, nil, authdomain.ErrGeneratingToken()
	}

	logger.Info(
		"sign-in executed successfully",
		zap.String("token:", token.AccessToken),
		zap.String("journey", "sign-in"),
	)

	return user, &token, nil
}
