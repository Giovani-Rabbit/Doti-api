package modulecase_test

import (
	"context"
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

	t.Run("Should be able to delete a module by id", func(t *testing.T) {
		moduleID := uuid.New().String()

		moduleRepo.EXPECT().DeleteModule(ctx, moduleID)
		err := deleteModuleCase.Execute(ctx, moduleID)

		if err != nil {
			t.Fatalf("The error was not expected, but we got: %s", err)
		}
	})

	t.Run("Should not be able to delete a module with an invalid id", func(t *testing.T) {
		moduleID := "12345"
		err := deleteModuleCase.Execute(ctx, moduleID)

		if err == nil {
			t.Fatalf("An invalid module id error was expected, but we got %s", err)
		}

		sttErr := err.(*http.RestErr).Status

		if sttErr != moduledomain.SttInvalidModuleID {
			t.Fatalf("An status error of INVALID_MODULE_ID was expected, but we got: %s", sttErr)
		}
	})
}
