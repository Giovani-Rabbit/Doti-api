package modulecase

import (
	"context"
	"strings"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	val "github.com/Giovani-Coelho/Doti-API/internal/pkg/validator"
	"go.uber.org/zap"
)

type rename struct {
	ModuleRepo repository.ModuleRepository
}

type Rename interface {
	Execute(ctx context.Context, id string, name string) error
}

func NewRenameModuleUseCase(
	moduleRepo repository.ModuleRepository,
) Rename {
	return &rename{
		ModuleRepo: moduleRepo,
	}
}

func (rm *rename) Execute(
	ctx context.Context,
	moduleId string, name string,
) error {
	logger.Info("Init rename module",
		zap.String("journey", "renameModule"),
	)

	if !val.IsValidUUID(moduleId) {
		logger.Error("Invalid module id", nil,
			zap.String("journey", "renameModule"))

		return moduledomain.ErrInvalidModuleID()
	}

	if strings.TrimSpace(name) == "" {
		logger.Error("Error, the new name is empty", nil,
			zap.String("journey", "renameModule"),
		)

		return moduledomain.ErrNewModuleNameIsEmpty()
	}

	name = strings.TrimSpace(name)

	err := rm.ModuleRepo.UpdateModuleName(ctx, moduleId, name)

	if err != nil {
		logger.Error("Error on change module name", err,
			zap.String("journey", "renameModule"),
		)

		return moduledomain.ErrCouldNotPersistModule(err)
	}

	logger.Info("Module renamed successfully",
		zap.String("journey", "renameModule"),
	)

	return nil
}
