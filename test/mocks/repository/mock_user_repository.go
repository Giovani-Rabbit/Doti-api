package mock_repository

import (
	"context"
	"errors"

	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
)

type MockUserRepository struct {
	Users []sqlc.User

	CreateFn                     func(ctx context.Context, user userDomain.IUserDomain) (userDomain.IUserDomain, error)
	CheckUserExistsFn            func(ctx context.Context, email string) (bool, error)
	FindUserByEmailFn            func(ctx context.Context, email string) (sqlc.User, error)
	FindUserByEmailAndPasswordFn func(ctx context.Context, args sqlc.FindUserByEmailAndPasswordParams) (sqlc.User, error)
}

func (m *MockUserRepository) Create(
	ctx context.Context, user userDomain.IUserDomain,
) (userDomain.IUserDomain, error) {
	if m.CreateFn != nil {
		return m.CreateFn(ctx, user)
	}

	return nil, errors.New("CreateFn not implemented")
}

func (m *MockUserRepository) CheckUserExists(
	ctx context.Context, email string,
) (bool, error) {
	if m.CheckUserExistsFn != nil {
		return m.CheckUserExistsFn(ctx, email)
	}

	return false, errors.New("CheckUserExistsFn not implemented")
}

func (m *MockUserRepository) FindUserByEmail(
	ctx context.Context, email string,
) (sqlc.User, error) {
	if m.FindUserByEmailFn != nil {
		return m.FindUserByEmailFn(ctx, email)
	}

	return sqlc.User{}, errors.New("FindUserByEmailFn not implemented")
}

func (m *MockUserRepository) FindUserByEmailAndPassword(
	ctx context.Context, args sqlc.FindUserByEmailAndPasswordParams,
) (sqlc.User, error) {
	if m.FindUserByEmailAndPasswordFn != nil {
		return m.FindUserByEmailAndPasswordFn(ctx, args)
	}

	return sqlc.User{}, errors.New("FindUserByEmailAndPasswordFn not implemented")
}
