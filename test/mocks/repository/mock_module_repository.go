package mock_repository

import (
	"context"
	"errors"

	module "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/db/sqlc"
)

type MockModuleRepository struct {
	Modules []sqlc.Module

	CreateFn             func(ctx context.Context, module module.Module) (module.Module, error)
	ListModuleByUserIDFn func(ctx context.Context, userId string) ([]module.Module, error)
	UpdateModuleNameFn   func(ctx context.Context, id string, name string) error
}

func (m *MockModuleRepository) ListModulesByUserID(
	ctx context.Context,
	userId string,
) ([]module.Module, error) {
	if m.ListModuleByUserIDFn == nil {
		return nil, errors.New("ListModuleByUserIDFn not implemented")
	}

	return m.ListModuleByUserIDFn(ctx, userId)
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

func (m *MockModuleRepository) UpdateModuleName(
	ctx context.Context,
	id string, name string,
) error {
	if m.UpdateModuleNameFn == nil {
		return errors.New("UpdateModuleNameFn not implemented")
	}

	return m.UpdateModuleNameFn(ctx, id, name)
}
