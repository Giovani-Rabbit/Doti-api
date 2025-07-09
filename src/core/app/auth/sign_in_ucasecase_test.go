package authcase_test

import (
	"context"
	"testing"
	"time"

	authcase "github.com/Giovani-Coelho/Doti-API/src/core/app/auth"
	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/src/pkg/auth"
	mock_repository "github.com/Giovani-Coelho/Doti-API/test/mocks/repository"
)

func TestSignInUseCase(t *testing.T) {
	mockRepo := &mock_repository.MockUserRepository{
		FindUserByEmailAndPasswordFn: func(
			ctx context.Context, args userdomain.IUserDomain,
		) (userdomain.IUserDomain, error) {
			return userdomain.NewUserDomain(
				"1", "giovani",
				"newuser@example.com",
				"password123",
				time.Now(), time.Now(),
			), nil
		},
	}

	signInCase := authcase.NewLoginUseCase(mockRepo)
	ctx := context.Background()

	credentials := userdomain.NewSignInUserDomain(
		"newuserexample@adawd.com",
		"password123",
	)

	t.Run("Should be able to log in successfully.", func(t *testing.T) {
		loggedUser, token, err := signInCase.Execute(ctx, credentials)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}

		userToken, err := auth.VerifyToken(token)

		if err != nil {
			t.Fatalf("Invalid Token: %v", err)
		}

		if loggedUser.GetEmail() != userToken.GetEmail() {
			t.Fatalf("Token Validation: Invalid e-mail: %s", userToken.GetEmail())
		}

		if loggedUser.GetName() != userToken.GetName() {
			t.Fatalf("Token Validation: Invalid username: %s", userToken.GetName())
		}
	})
}
