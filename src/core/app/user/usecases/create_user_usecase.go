package user

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	userdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"go.uber.org/zap"
)

type CreateUserUseCase struct {
	UserRepository repository.IUserRepository
}

type ICreateUserUseCase interface {
	Execute(ctx context.Context, userDTO userdto.CreateUserDTO) (userDomain.IUserDomain, error)
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
	userDTO userdto.CreateUserDTO,
) (userDomain.IUserDomain, error) {
	logger.Info("Init CreateUser UseCase",
		zap.String("journey", "createUser"),
	)

	userAlreadyExists, _ := us.UserRepository.CheckUserExists(
		ctx, userDTO.Email,
	)

	if userAlreadyExists {
		logger.Error(
			"Error: User already exists", nil,
			zap.String("journey", "createUser"),
		)

		return nil, userDomain.ErrUserAlreadyExists()
	}

	createUserDomain := userDomain.NewCreateUserDomain(
		userDTO.Name,
		userDTO.Email,
		userDTO.Password,
	)

	user, err := us.UserRepository.Create(ctx, createUserDomain)

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
		zap.String("userId", user.GetID()),
		zap.String("journey", "createUser"),
	)

	return user, nil
}
