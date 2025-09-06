package modulecase

import (
	"context"
	"strconv"
	"strings"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
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
		zap.String("moduleId", moduleId),
		zap.String("newName", name))

	if moduleId == "" {
		logger.Error("Module id is empty", nil,
			zap.String("journey", "renameModule"))

		return moduledomain.ErrInvalidModuleID()
	}

	intId, err := strconv.ParseInt(moduleId, 10, 32)
	if err != nil {
		logger.Error("Invalid module id", err,
			zap.String("journey", "renameModule"),
			zap.String("moduleId", moduleId))

		return moduledomain.ErrInvalidModuleID()
	}

	moduleIdInt := int32(intId)

	if strings.TrimSpace(name) == "" {
		logger.Error("Error, the new name is empty", nil,
			zap.String("journey", "renameModule"),
			zap.Int32("moduleId", moduleIdInt))

		return moduledomain.ErrNewModuleNameIsEmpty()
	}

	name = strings.TrimSpace(name)

	err = rm.ModuleRepo.UpdateModuleName(ctx, moduleIdInt, name)
	if err != nil {
		logger.Error("Error on change module name", err,
			zap.String("journey", "renameModule"),
			zap.Int32("moduleId", moduleIdInt),
			zap.String("newName", name))

		return moduledomain.ErrCouldNotPersistModule(err)
	}

	logger.Info("Module renamed successfully",
		zap.String("journey", "renameModule"),
		zap.Int32("moduleId", moduleIdInt),
		zap.String("newName", name))

	return nil
}
