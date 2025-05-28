package userUseCase_test

import (
	"context"
	"testing"

	userUseCase "github.com/Giovani-Coelho/Doti-API/src/core/app/user/usecases"
	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	userDTO "github.com/Giovani-Coelho/Doti-API/src/infra/http/controller/user/dtos"
	mocks "github.com/Giovani-Coelho/Doti-API/test/mocks/user"
)

func TestUserRepository_CreateAndCheck(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{}
	createUser := userUseCase.NewCreateUserUseCase(mockRepo)

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
			ctx context.Context,
			user userDomain.IUserDomain,
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
