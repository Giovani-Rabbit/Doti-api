package modulecase_test

import (
	"context"
	"testing"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
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

	moduleRepo.EXPECT().
		Create(ctx, module).
		Return(module, nil)

	createModulecase := modulecase.NewCreateModuleUseCase(moduleRepo)

	t.Run("Should be able to create new module successfully", func(t *testing.T) {
		_, err := createModulecase.Execute(ctx, module)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}
	})

	t.Run("Should not be able to create new module with empty fields", func(t *testing.T) {
		emptyModule := moduledomain.NewCreateModule(userID, "", "")

		_, err := createModulecase.Execute(ctx, emptyModule)

		if err == nil {
			t.Fatalf("an error was expected, but we got: %v", err)
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
	})
}
