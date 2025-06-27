package user

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	"github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"go.uber.org/zap"
)

type CreateUserUseCase struct {
	UserRepository repository.IUserRepository
}

type ICreateUserUseCase interface {
	Execute(ctx context.Context, user user.IUserDomain) (user.IUserDomain, error)
}

func NewCreateUserUseCase(
	userRepository repository.IUserRepository,
) ICreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (us *CreateUserUseCase) Execute(
	ctx context.Context,
	userDomain user.IUserDomain,
) (user.IUserDomain, error) {
	logger.Info("Init CreateUser UseCase",
		zap.String("journey", "createUser"),
	)

	if isValidUser := userDomain.IsValid(); !isValidUser {
		logger.Error(
			"Error: User values are missing", nil,
			zap.String("journey", "createUser"),
		)

		return nil, user.ErrUserValuesMissing()
	}

	userAlreadyExists, _ := us.UserRepository.CheckUserExists(
		ctx, userDomain.GetEmail(),
	)

	if userAlreadyExists {
		logger.Error(
			"Error: User already exists", nil,
			zap.String("journey", "createUser"),
		)

		return nil, user.ErrUserAlreadyExists()
	}

	userCreated, err := us.UserRepository.Create(ctx, userDomain)

	if err != nil {
		logger.Error(
			"Error in repository", nil,
			zap.String("journey", "createUser"),
		)

		return nil, rest_err.NewInternalServerError(
			"Internal error saving user",
		)
	}

	logger.Info(
		"CreateUser executed successfully",
		zap.String("userId", userCreated.GetID()),
		zap.String("journey", "createUser"),
	)

	return userCreated, nil
}
