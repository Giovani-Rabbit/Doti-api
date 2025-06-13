package userUseCase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	userDTO "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"go.uber.org/zap"
)

type CreateUserUseCase struct {
	UserRepository repository.IUserRepository
}

type ICreateUserUseCase interface {
	Execute(ctx context.Context, userDTO userDTO.CreateUserDTO) error
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
	userDTO userDTO.CreateUserDTO,
) error {
	logger.Info("Init CreateUser UseCase",
		zap.String("journey", "createUser"),
	)

	userAlreadyExists, _ := us.UserRepository.CheckUserExists(
		ctx, userDTO.Email,
	)

	if userAlreadyExists {
		return rest_err.NewBadRequestError(
			"User already exists",
		)
	}

	userDomain := userDomain.NewUserDomain(
		userDTO.Name,
		userDTO.Email,
		userDTO.Password,
	)

	err := us.UserRepository.Create(ctx, userDomain)

	if err != nil {
		return rest_err.NewInternalServerError(
			"Internal error when saving user",
		)
	}

	return nil
}
