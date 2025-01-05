package mocks

import (
	"context"

	userDTO "github.com/Giovani-Coelho/Doti-API/src/application/user/dtos"
	"github.com/Giovani-Coelho/Doti-API/src/infra/database/db/sqlc"
)

type MockUserRepository struct {
	MockCreate                     func(ctx context.Context, userDto userDTO.CreateUserDTO) error
	MockCheckUserExists            func(ctx context.Context, email string) (bool, error)
	MockFindUserByEmail            func(ctx context.Context, email string) (sqlc.User, error)
	MockFindUserByEmailAndPassword func(
		ctx context.Context,
		args sqlc.FindUserByEmailAndPasswordParams,
	) (sqlc.User, error)
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

func (m *MockUserRepository) FindUserByEmail(
	ctx context.Context,
	email string,
) (sqlc.User, error) {
	if m.MockFindUserByEmail != nil {
		return m.MockFindUserByEmail(ctx, email)
	}

	return sqlc.User{}, nil
}

func (m *MockUserRepository) FindUserByEmailAndPassword(
	ctx context.Context,
	args sqlc.FindUserByEmailAndPasswordParams,
) (sqlc.User, error) {
	if m.MockFindUserByEmailAndPassword != nil {
		return m.MockFindUserByEmailAndPassword(ctx, args)
	}

	return sqlc.User{}, nil
}
