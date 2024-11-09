package userServices

import (
	"context"
	"fmt"

	dtos "github.com/Giovani-Coelho/Doti-API/src/application/dtos/user"
)

func (us *CreateUserService) CreateUser(
	ctx context.Context,
	userDTO dtos.CreateUserDto,
) error {
	userAlreadyExists, err := us.UserRepository.CheckUserExists(ctx, userDTO.Email)

	if err != nil {
		panic(err)
	}

	if userAlreadyExists {
		fmt.Println("Usuario ja existe")
		return nil
	}

	err = us.UserRepository.Create(ctx, userDTO)

	if err != nil {
		panic(err)
	}

	return nil
}
