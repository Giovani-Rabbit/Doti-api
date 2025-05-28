package mocks

import (
	"context"

	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
)

type MockUserRepository struct {
	MockCreate func(
		ctx context.Context,
		userDomain userDomain.IUserDomain,
	) error

	MockCheckUserExists func(
		ctx context.Context, email string,
	) (bool, error)

	MockFindUserByEmail func(
		ctx context.Context, email string,
	) (sqlc.User, error)

	MockFindUserByEmailAndPassword func(
		ctx context.Context,
		args sqlc.FindUserByEmailAndPasswordParams,
	) (sqlc.User, error)
}

func (m *MockUserRepository) Create(
	ctx context.Context,
	userDomain userDomain.IUserDomain,
) error {
	if m.MockCreate != nil {
		return m.MockCreate(ctx, userDomain)
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
