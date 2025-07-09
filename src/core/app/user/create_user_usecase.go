package usercase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"go.uber.org/zap"
)

type CreateUserUseCase struct {
	UserRepository repository.IUserRepository
}

type ICreateUserUseCase interface {
	Execute(ctx context.Context, userEntity userdomain.IUserDomain) (userdomain.IUserDomain, *rest_err.RestErr)
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
	userEntity userdomain.IUserDomain,
) (userdomain.IUserDomain, *rest_err.RestErr) {
	logger.Info("Init CreateUser UseCase",
		zap.String("journey", "createUser"),
	)

	if !userEntity.IsValid() {
		logger.Error(
			"Error: User values are missing", nil,
			zap.String("journey", "createUser"),
		)

		return nil, userdomain.ErrUserValuesMissing()
	}

	if err := userEntity.ValidatePassword(); err != nil {
		logger.Error(
			"Error: Invalid Password", nil,
			zap.String("journey", "createUser"),
		)

		return nil, userdomain.ErrInvalidPassword(err)
	}

	userEntity.EncryptPassword()

	if !userEntity.IsValidEmail() {
		logger.Error(
			"Error: Invalid user email format", nil,
			zap.String("journey", "createUser"),
		)

		return nil, userdomain.ErrInvalidUserEmailFormat()
	}

	userAlreadyExists, _ := us.UserRepository.CheckUserExists(
		ctx, userEntity.GetEmail(),
	)

	if userAlreadyExists {
		logger.Error(
			"Error: User already exists", nil,
			zap.String("journey", "createUser"),
		)

		return nil, userdomain.ErrUserAlreadyExists()
	}

	userCreated, err := us.UserRepository.Create(ctx, userEntity)

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
