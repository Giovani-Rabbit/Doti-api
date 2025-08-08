package modulecase_test

import (
	"context"
	"testing"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestCreateModuleUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	moduleRepo := mocks.NewMockModuleRepository(ctrl)
	ctx := context.Background()

	userID := uuid.New().String()

	module := moduledomain.NewCreateModule(
		userID,
		"new module",
		"icon",
	)

	createModulecase := modulecase.NewCreateModuleUseCase(moduleRepo)

	t.Run("Should be able to create new module successfully", func(t *testing.T) {
		moduleRepo.EXPECT().
			Create(ctx, module).
			Return(module, nil)

		_, err := createModulecase.Execute(ctx, module)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}
	})

	t.Run("should not be possible to create a new module with an invalid UUID", func(t *testing.T) {
		moduleWithInvalidUserID := moduledomain.NewCreateModule(
			"12345",
			"new module",
			"icon",
		)

		_, err := createModulecase.Execute(ctx, moduleWithInvalidUserID)

		if err == nil {
			t.Fatalf("expected error, but we got: %v", err)
		}

		sttErr := err.(*http.RestErr).Status

		if sttErr != moduledomain.SttInvalidUserID {
			t.Fatalf("an error of the type INVALID_USER_ID was expected, but we got: %v", sttErr)
		}
	})

	t.Run("Should not be able to create new module with empty fields", func(t *testing.T) {
		emptyModule := moduledomain.NewCreateModule(userID, "", "")

		_, err := createModulecase.Execute(ctx, emptyModule)

		if err == nil {
			t.Fatalf("an error was expected, but we got: %v", err)
		}

		sttErr := err.(*http.RestErr).Status

		if sttErr != moduledomain.SttInvalidModuleFields {
			t.Fatalf("an error of the type INVALID_MODULE_FIELDS was expected, but we got: %v", sttErr)
		}
	})
}
