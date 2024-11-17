package user_test

import (
	"context"
	"testing"

	dtos "github.com/Giovani-Coelho/Doti-API/src/application/dtos/user"
	userServices "github.com/Giovani-Coelho/Doti-API/src/application/services/user/createUser"
	mocks "github.com/Giovani-Coelho/Doti-API/test/mocks/user"
)

func TestUserRepository_CreateAndCheck(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{}
	userService := userServices.NewCreateUserService(mockRepo)

	t.Run("Create new user successfully", func(t *testing.T) {
		userDTO := dtos.CreateUserDto{
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
			ctx context.Context, userDTO dtos.CreateUserDto,
		) error {
			return nil
		}

		err := userService.CreateUser(context.Background(), userDTO)

		if err != nil {
			t.Fatalf("expected no error, but got: %v", err)
		}
	})

	t.Run("User already exists", func(t *testing.T) {
		userDTO := dtos.CreateUserDto{
			Name:     "Existing User",
			Email:    "existinguser@example.com",
			Password: "password123",
		}

		mockRepo.MockCheckUserExists = func(
			ctx context.Context, email string,
		) (bool, error) {
			return true, nil
		}

		err := userService.CreateUser(context.Background(), userDTO)

		if err != nil {
			t.Fatalf("expected no error, but got: %v", err)
		}
	})
}
