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

func TestRenameModuleUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	moduleRepo := mocks.NewMockModuleRepository(ctrl)
	renameCase := modulecase.NewRenameModuleUseCase(moduleRepo)

	ctx := context.Background()

	t.Run("Should be able to rename a module", func(t *testing.T) {
		moduleId := uuid.New().String()
		newName := "new Name"

		moduleRepo.EXPECT().
			UpdateModuleName(ctx, moduleId, newName)

		err := renameCase.Execute(ctx, moduleId, newName)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}
	})

	t.Run("Should not be able to rename if the module ID is invalid", func(t *testing.T) {
		newName := "new Name"

		err := renameCase.Execute(ctx, "1234", newName)

		if err == nil {
			t.Fatalf("an error was expected, but we got: %v", err)
		}

		sttErr := err.(*http.RestErr).Status

		if sttErr != moduledomain.SttInvalidModuleID {
			t.Fatalf("an error of the type INVALID_MODULE_ID was expected, but we got: %v", err)
		}
	})

	t.Run("Should not be able to rename if the new name is invalid", func(t *testing.T) {
		err := renameCase.Execute(ctx, uuid.New().String(), "  ")

		if err == nil {
			t.Fatalf("an error was expected, but we got: %v", err)
		}

		sttErr := err.(*http.RestErr).Status

		if sttErr != moduledomain.SttNewModuleNameIsEmpty {
			t.Fatalf("an error of the type NEW_MODULE_NAME_IS_EMPTY was expected, but we got: %v", err)
		}
	})
}
