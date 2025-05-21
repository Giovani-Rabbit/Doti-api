package userServices

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	userDTO "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user/dtos"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"go.uber.org/zap"
)

func (us *UserServices) CreateUser(
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
