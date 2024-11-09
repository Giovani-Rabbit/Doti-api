package userServices

import (
	"context"

	dtos "github.com/Giovani-Coelho/Doti-API/src/application/dtos/user"
)

func (us *CreateUserService) CreateUser(
	ctx context.Context,
	userDTO dtos.CreateUserDto,
) error {
	err := us.UserRepository.Create(ctx, userDTO)

	if err != nil {
		panic(err)
	}

	return nil
}
