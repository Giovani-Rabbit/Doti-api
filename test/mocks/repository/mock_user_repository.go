package mock_repository

import (
	"context"
	"errors"

	user "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
)

type MockUserRepository struct {
	Users []sqlc.User

	CreateFn                     func(ctx context.Context, user user.User) (user.User, error)
	CheckUserExistsFn            func(ctx context.Context, email string) (bool, error)
	FindUserByEmailFn            func(ctx context.Context, email string) (user.User, error)
	FindUserByEmailAndPasswordFn func(ctx context.Context, args user.User) (user.User, error)
}

func (m *MockUserRepository) Create(
	ctx context.Context, user user.User,
) (user.User, error) {
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
) (user.User, error) {
	if m.FindUserByEmailFn != nil {
		return m.FindUserByEmailFn(ctx, email)
	}

	return nil, errors.New("FindUserByEmailFn not implemented")
}

func (m *MockUserRepository) FindUserByEmailAndPassword(
	ctx context.Context, args user.User,
) (user.User, error) {
	if m.FindUserByEmailAndPasswordFn != nil {
		return m.FindUserByEmailAndPasswordFn(ctx, args)
	}

	return nil, errors.New("FindUserByEmailAndPasswordFn not implemented")
}
