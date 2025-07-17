package mock_repository

import (
	"context"
	"errors"

	module "github.com/Giovani-Coelho/Doti-API/src/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
)

type MockModuleRepository struct {
	Modules []sqlc.Module

	CreateFn func(ctx context.Context, module module.IModuleDomain) (module.IModuleDomain, error)
}

func (m *MockModuleRepository) Create(
	ctx context.Context,
	module module.IModuleDomain,
) (module.IModuleDomain, error) {
	if m.CreateFn == nil {
		return nil, errors.New("createFn not implemented")
	}

	return m.CreateFn(ctx, module)
}
