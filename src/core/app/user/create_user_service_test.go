package userServices_test

import (
	"context"
	"testing"

	userServices "github.com/Giovani-Coelho/Doti-API/src/core/app/user"
	userDTO "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user/dtos"
	mocks "github.com/Giovani-Coelho/Doti-API/test/mocks/user"
)

func TestUserRepository_CreateAndCheck(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{}
	createUser := userServices.NewCreateUserService(mockRepo)

	t.Run("Create new user successfully", func(t *testing.T) {
		user := userDTO.CreateUserDTO{
			Name:     "New User",
			Email:    "newuser@example.com",
			Password: "password123",
		}

		mockRepo.MockCheckUserExists = func(
			ctx context.Context, email string,
		) (bool, error) {
			return false, nil
		}

		mockRepo.MockCreate = func(
			ctx context.Context, user userDTO.CreateUserDTO,
		) error {
			return nil
		}

		err := createUser.Execute(context.Background(), user)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}
	})

	t.Run("User already exists", func(t *testing.T) {
		userDTO := userDTO.CreateUserDTO{
			Name:     "Existing User",
			Email:    "existinguser@example.com",
			Password: "password123",
		}

		mockRepo.MockCheckUserExists = func(
			ctx context.Context, email string,
		) (bool, error) {
			return true, nil
		}

		err := createUser.Execute(context.Background(), userDTO)

		if err == nil {
			t.Fatalf("expected: the user already exists. But we got: %v", err)
		}
	})
}
