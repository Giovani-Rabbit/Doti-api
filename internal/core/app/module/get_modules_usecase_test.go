package modulecase_test

import (
	"context"
	"testing"

	modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	mock_repository "github.com/Giovani-Coelho/Doti-API/test/mocks/repository"
	"github.com/google/uuid"
)

func TestGetModulesUseCase(t *testing.T) {

	var uuiduserId = uuid.New()
	var modules []moduledomain.Module

	modulesQuantity := 6

	for range modulesQuantity {
		modules = append(
			modules,
			moduledomain.NewCreateModule(
				uuiduserId.String(),
				"module",
				"icon",
			),
		)
	}

	moduleRepo := mock_repository.MockModuleRepository{
		ListModuleByUserIDFn: func(ctx context.Context, userId string) ([]moduledomain.Module, error) {
			return modules, nil
		},
	}

	getModules := modulecase.NewGetModulesUseCase(&moduleRepo)

	ctx := context.Background()

	t.Run("Should be able to list modules", func(t *testing.T) {
		modulesList, err := getModules.Execute(ctx, uuiduserId.String())

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}

		if len(modulesList) < modulesQuantity {
			t.Fatal("expected a list with 6 modules")
		}
	})
}
