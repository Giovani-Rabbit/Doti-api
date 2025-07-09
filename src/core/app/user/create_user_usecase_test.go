package usercase_test

import (
	"context"
	"testing"

	usercase "github.com/Giovani-Coelho/Doti-API/src/core/app/user"
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

	t.Run("Should be able to create new user successfully", func(t *testing.T) {
		_, err := createUser.Execute(ctx, user)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}
	})

	t.Run("Should not be able to create an already existing user", func(t *testing.T) {
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

	t.Run("Should not be able to use a invalid password", func(t *testing.T) {
		userInvalidPassword := userdomain.NewCreateUserDomain(
			"giovani",
			"giovani@emai.com",
			"12",
		)

		// password must contain both letter and numbers
		// passsword must be at least 4 characters
		_, err := createUser.Execute(ctx, userInvalidPassword)

		if err == nil {
			t.Fatalf("An Error was expected. But we got nil")
		}

		if err.Status != userdomain.SttInvalidPassword {
			t.Fatalf("Expected invalid password error")
		}
	})

	t.Run("Should not be able to use a invalid e-mail", func(t *testing.T) {
		userInvalidPassword := userdomain.NewCreateUserDomain(
			"giovani",
			"giovaniemai.com",
			"abc123",
		)

		// password must contain both letter and numbers
		// passsword must be at least 4 characters
		_, err := createUser.Execute(ctx, userInvalidPassword)

		if err == nil {
			t.Fatalf("An Error was expected. But we got nil")
		}

		if err.Status != userdomain.SttInvalidUserEmailFormat {
			t.Fatalf("Expected invalid email format error, gor: %s", err)
		}
	})
}
