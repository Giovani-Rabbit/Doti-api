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
			return userdomain.New(
				"1", "giovani",
				"newuser@example.com",
				"password123",
				time.Now(), time.Now(),
			), nil
		},
	}

	signInCase := authcase.NewLoginUseCase(mockRepo)
	ctx := context.Background()

	t.Run("Should be able to sign-in successfully.", func(t *testing.T) {
		credentials := userdomain.NewSignInUser(
			"newuserexample@adawd.com",
			"password123",
		)

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

	t.Run("Should not be able to sign-in without credentials.", func(t *testing.T) {
		credentials := userdomain.NewSignInUser("", "")
		_, _, err := signInCase.Execute(ctx, credentials)

		if err == nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}

		if err.Status != userdomain.SttUserValuesMissing {
			t.Fatalf("Expected error: USER_VALUES_MISSING, got %s", err.Status)
		}
	})

	t.Run("Should not be able to sign-in with invalid e-mail.", func(t *testing.T) {
		credentials := userdomain.NewSignInUser("giovaniemail.com", "123")
		_, _, err := signInCase.Execute(ctx, credentials)

		if err == nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}

		if err.Status != userdomain.SttInvalidUserEmailFormat {
			t.Fatalf("Expected error: INVALID_USER_EMAIL_FORMAT, got %s", err.Status)
		}
	})
}
