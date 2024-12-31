package userServices

import (
	"context"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/application/user/dtos"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
)

func (us *UserServices) CreateUser(
	ctx context.Context,
	userDTO userDTO.CreateUserDTO,
) error {
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
