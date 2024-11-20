package userServices

import (
	"context"
	"fmt"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/application/user/dtos"
)

func (us *CreateUserService) CreateUser(
	ctx context.Context,
	userDTO userDTO.CreateUserDTO,
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
