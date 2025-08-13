package modulecase_test

import (
	"context"
	"errors"
	"testing"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestDeleteModuleUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	moduleRepo := mock_repository.NewMockModuleRepository(ctrl)
	deleteModuleCase := modulecase.NewDeleteModuleUseCase(moduleRepo)

	ctx := context.Background()

	t.Run("Should delete module successfully", func(t *testing.T) {
		moduleID := uuid.New().String()

		moduleRepo.EXPECT().DeleteModule(gomock.Any(), moduleID)

		err := deleteModuleCase.Execute(ctx, moduleID)
		if err != nil {
			t.Fatalf("Expected nil error, got %v", err)
		}
	})

	t.Run("Should return error if id is invalid", func(t *testing.T) {
		moduleID := "12345"

		err := deleteModuleCase.Execute(ctx, moduleID)
		if err == nil {
			t.Fatal("Expected an error for invalid module id, but got nil")
		}

		restErr, ok := err.(*http.RestErr)
		if !ok {
			t.Fatalf("Expected error of type RestErr, but got: %T", err)
		}

		if restErr.Status != moduledomain.SttInvalidModuleID {
			t.Fatalf(
				"An status error of INVALID_MODULE_ID was expected, but we got: %s",
				restErr.Status,
			)
		}
	})

	t.Run("Should return error if repository fails", func(t *testing.T) {
		moduleID := uuid.New().String()

		moduleRepo.EXPECT().
			DeleteModule(gomock.Any(), moduleID).
			Return(errors.New("db error"))

		err := deleteModuleCase.Execute(ctx, moduleID)

		if err == nil {
			t.Fatal("Expected error from repository, got nil")
		}

		restErr, ok := err.(*http.RestErr)
		if !ok {
			t.Fatalf("Expected *http.RestErr, got %T", err)
		}

		if restErr.Status != moduledomain.SttCouldNotPersistModule {
			t.Fatalf("Expected status DELETING_MODULE, got %s", restErr.Status)
		}
	})
}
