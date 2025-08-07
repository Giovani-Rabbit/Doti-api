package usercase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"
	"go.uber.org/zap"
)

type createUserUseCase struct {
	UserRepository repository.UserRepository
}

type CreateUserUseCase interface {
	Execute(ctx context.Context, userEntity userdomain.User) (userdomain.User, *http.RestErr)
}

func NewCreateUserUseCase(
	userRepository repository.UserRepository,
) CreateUserUseCase {
	return &createUserUseCase{
		UserRepository: userRepository,
	}
}

func (us *createUserUseCase) Execute(
	ctx context.Context,
	userEntity userdomain.User,
) (userdomain.User, *http.RestErr) {
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
			"Error in repository", err,
			zap.String("journey", "createUser"),
		)

		return nil, http.NewInternalServerError(
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
