package modulecase_test

import (
	"context"
	"testing"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestGetModulesUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var uuiduserId = uuid.New().String()
	var modules []moduledomain.Module

	modulesQuantity := 6

	for range modulesQuantity {
		modules = append(
			modules,
			moduledomain.NewCreateModule(
				uuiduserId,
				"module",
				"icon",
			),
		)
	}

	moduleRepo := mock_repository.NewMockModuleRepository(ctrl)
	getModules := modulecase.NewGetModulesUseCase(moduleRepo)

	ctx := context.Background()

	t.Run("Should be able to list modules", func(t *testing.T) {
		moduleRepo.EXPECT().
			ListModulesByUserID(ctx, uuiduserId).
			Return(modules, nil)

		modulesList, err := getModules.Execute(ctx, uuiduserId)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}

		if len(modulesList) < modulesQuantity {
			t.Fatal("expected a list with 6 modules")
		}
	})
}
