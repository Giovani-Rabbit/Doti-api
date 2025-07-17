package modulecase_test

import (
	"context"
	"testing"

	modulecase "github.com/Giovani-Coelho/Doti-API/src/core/app/module"
	moduledomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/module"
	mock "github.com/Giovani-Coelho/Doti-API/test/mocks/repository"
	"github.com/google/uuid"
)

func TestCreateModuleUseCase(t *testing.T) {
	module := &mock.MockModuleRepository{
		CreateFn: func(ctx context.Context, module moduledomain.IModuleDomain) (moduledomain.IModuleDomain, error) {
			_, err := uuid.Parse(module.GetUserId())

			if err != nil {
				return nil, err
			}

			return module, nil
		},
	}

	createModule := modulecase.NewCreateModuleUseCase(module)

	ctx := context.Background()

	newModule := moduledomain.NewCreateModule(
		"146681af-2cee-493a-a145-d23609ae056d",
		"Math",
		"PI",
	)

	t.Run("Should be able to create new module successfully", func(t *testing.T) {
		_, err := createModule.Execute(ctx, newModule)

		if err != nil {
			t.Fatalf("expected no error, but we got: %v", err)
		}
	})

	t.Run("should not be possible to create a new module with an invalid UUID", func(t *testing.T) {
		moduleWithInvalidUserID := moduledomain.NewCreateModule(
			"12345",
			"Math",
			"PI",
		)

		_, err := createModule.Execute(ctx, moduleWithInvalidUserID)

		if err == nil {
			t.Fatalf("expected error, but we got: %v", err)
		}
	})
}
