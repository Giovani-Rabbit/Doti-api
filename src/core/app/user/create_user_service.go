package userServices

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	userDTO "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/repository"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"go.uber.org/zap"
)

type CreateUserService struct {
	UserRepository repository.IUserRepository
}

type ICreateUserService interface {
	Execute(ctx context.Context, userDTO userDTO.CreateUserDTO) error
}

func NewCreateUserService(
	userRepository repository.IUserRepository,
) ICreateUserService {
	return &CreateUserService{
		UserRepository: userRepository,
	}
}

func (us *CreateUserService) Execute(
	ctx context.Context,
	userDTO userDTO.CreateUserDTO,
) error {
	logger.Info("Init CreateUser service",
		zap.String("journey", "createUser"),
	)

	userAlreadyExists, _ := us.UserRepository.CheckUserExists(ctx, userDTO.Email)

	if userAlreadyExists {
		return rest_err.NewBadRequestError(
			"User already exists",
		)
	}

	err := us.UserRepository.Create(ctx, userDTO)

	if err != nil {
		return rest_err.NewInternalServerError(
			"Internal error when saving user",
		)
	}

	return nil
}
