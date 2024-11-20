package mocks

import (
	"context"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/application/user/dtos"
)

type MockUserRepository struct {
	MockCreate          func(ctx context.Context, userDto userDTO.CreateUserDTO) error
	MockCheckUserExists func(ctx context.Context, email string) (bool, error)
}

func (m *MockUserRepository) Create(
	ctx context.Context,
	userDto userDTO.CreateUserDTO,
) error {
	if m.MockCreate != nil {
		return m.MockCreate(ctx, userDto)
	}

	return nil
}

func (m *MockUserRepository) CheckUserExists(
	ctx context.Context,
	email string,
) (bool, error) {
	if m.MockCheckUserExists != nil {
		return m.MockCheckUserExists(ctx, email)
	}

	return false, nil
}
