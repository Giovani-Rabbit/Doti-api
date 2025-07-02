package usercase_test

import (
	"context"
	"testing"

	usercase "github.com/Giovani-Coelho/Doti-API/src/core/app/user/usecases"
	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	mock_repository "github.com/Giovani-Coelho/Doti-API/test/mocks/repository"
)

func TestCreateUserUseCase(t *testing.T) {
	mockRepo := &mock_repository.MockUserRepository{
		CreateFn: func(ctx context.Context, user userdomain.IUserDomain) (userdomain.IUserDomain, error) {
			return user, nil
		},
		CheckUserExistsFn: func(ctx context.Context, email string) (bool, error) {
			return false, nil
		},
	}

	createUser := usercase.NewCreateUserUseCase(mockRepo)

	ctx := context.Background()

	user := userdomain.NewCreateUserDomain(
		"New User",
		"newuser@example.com",
		"password123",
	)

	t.Run("Create new user successfully", func(t *testing.T) {
		_, err := createUser.Execute(ctx, user)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}
	})

	t.Run("User already exists", func(t *testing.T) {
		mockRepo.CheckUserExistsFn = func(
			ctx context.Context, email string,
		) (bool, error) {
			return true, nil
		}

		_, err := createUser.Execute(ctx, user)

		if err == nil {
			t.Fatalf("expected: the user already exists. But we got: %v", err)
		}
	})
}
