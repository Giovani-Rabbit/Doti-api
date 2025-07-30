package mock_repository

import (
	"context"
	"errors"

	module "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
)

type MockModuleRepository struct {
	Modules []sqlc.Module

	CreateFn func(ctx context.Context, module module.Module) (module.Module, error)
}

func (m *MockModuleRepository) Create(
	ctx context.Context,
	module module.Module,
) (module.Module, error) {
	if m.CreateFn == nil {
		return nil, errors.New("createFn not implemented")
	}

	return m.CreateFn(ctx, module)
}
