package authcase_test

import (
	"context"
	"testing"

	authcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/auth"
	apperr "github.com/Giovani-Coelho/Doti-API/internal/core/app/errors"
	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
	"github.com/golang/mock/gomock"
)

func TestSignInUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	signInCase := authcase.NewLoginUseCase(mockRepo)
	ctx := context.Background()

	credentials := userdomain.NewSignInUser(
		"newuserexample@adawd.com",
		"password123",
	)

	mockRepo.EXPECT().
		FindUserByEmailAndPassword(ctx, credentials).
		Return(credentials, nil)

	t.Run("Should be able to sign-in successfully.", func(t *testing.T) {

		loggedUser, token, err := signInCase.Execute(ctx, credentials)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}

		userToken, err := auth.VerifyToken(token)

		if err != nil {
			t.Fatalf("Invalid Token: %v", err)
		}

		if loggedUser.GetEmail() != userToken.Email {
			t.Fatalf("Token Validation: Invalid e-mail: %s", userToken.Email)
		}

		if loggedUser.GetName() != userToken.Name {
			t.Fatalf("Token Validation: Invalid username: %s", userToken.Name)
		}
	})

	t.Run("Should not be able to sign-in without credentials.", func(t *testing.T) {
		credentials := userdomain.NewSignInUser("", "")
		_, _, err := signInCase.Execute(ctx, credentials)

		if err == nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}

		if err.Status != apperr.SttUserValuesMissing {
			t.Fatalf("Expected error: USER_VALUES_MISSING, got %s", err.Status)
		}
	})

	t.Run("Should not be able to sign-in with invalid e-mail.", func(t *testing.T) {
		credentials := userdomain.NewSignInUser("giovaniemail.com", "123")
		_, _, err := signInCase.Execute(ctx, credentials)

		if err == nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}

		if err.Status != apperr.SttInvalidUserEmailFormat {
			t.Fatalf("Expected error: INVALID_USER_EMAIL_FORMAT, got %s", err.Status)
		}
	})
}
