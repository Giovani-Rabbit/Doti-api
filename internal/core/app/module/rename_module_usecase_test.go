package modulecase_test

import (
	"context"
	"testing"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestRenameModuleUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	moduleRepo := mock_repository.NewMockModuleRepository(ctrl)
	renameCase := modulecase.NewRenameModuleUseCase(moduleRepo)

	ctx := context.Background()

	t.Run("Should be able to rename a module", func(t *testing.T) {
		moduleId := "1234"
		newName := "new Name"

		moduleRepo.EXPECT().
			UpdateModuleName(ctx, gomock.Any(), gomock.Any()).Return(nil)

		err := renameCase.Execute(ctx, moduleId, newName)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}
	})

	t.Run("Should not be able to rename if the module ID is invalid", func(t *testing.T) {
		newName := "new Name"

		err := renameCase.Execute(ctx, "abc123", newName)

		if err == nil {
			t.Fatalf("an error was expected, but we got: %v", err)
		}

		sttErr := err.(*resp.RestErr).Status

		if sttErr != moduledomain.SttInvalidModuleID {
			t.Fatalf(
				"an error of the type INVALID_MODULE_ID was expected, but we got: %v",
				err,
			)
		}
	})

	t.Run("Should not be able to rename if the new name is invalid", func(t *testing.T) {
		err := renameCase.Execute(ctx, "1234", "  ")

		if err == nil {
			t.Fatalf("an error was expected, but we got: %v", err)
		}

		sttErr := err.(*resp.RestErr).Status

		if sttErr != moduledomain.SttNewModuleNameIsEmpty {
			t.Fatalf(
				"an error of the type %v was expected, but we got: %v",
				moduledomain.SttNewModuleNameIsEmpty, sttErr,
			)
		}
	})
}
