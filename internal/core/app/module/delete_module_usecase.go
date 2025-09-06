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

type Delete interface {
	Execute(ctx context.Context, id string) error
}

type delete struct {
	moduleRepository repository.ModuleRepository
}

func NewDeleteModuleUseCase(
	moduleRepo repository.ModuleRepository,
) Delete {
	return &delete{
		moduleRepository: moduleRepo,
	}
}

func (dm *delete) Execute(ctx context.Context, id string) error {
	logger.Info("Init delete module usecase",
		zap.String("journey", "deleteModule"),
		zap.String("moduleId", id))

	if strings.TrimSpace(id) == "" {
		logger.Error("module id is empty", nil,
			zap.String("journey", "deleteModule"),
			zap.String("moduleId", id))

		return moduledomain.ErrInvalidModuleID()
	}

	intId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		logger.Error("Invalid, id is not a int", err,
			zap.String("journey", "deleteModule"),
			zap.String("moduleId", id))

		return moduledomain.ErrInvalidModuleID()
	}

	int32Id := int32(intId)

	moduleExists, err := dm.moduleRepository.CheckExistsById(ctx, int32Id)
	if err != nil {
		logger.Error("Error finding a module", err,
			zap.String("journey", "deleteModule"),
			zap.Int32("moduleId", int32Id))

		return moduledomain.ErrCouldNotPersistModule(err)
	}

	if !moduleExists {
		logger.Info("Module not find",
			zap.String("journey", "deleteModule"))

		return moduledomain.ErrCouldNotFindModuleByID()
	}

	err = dm.moduleRepository.DeleteModule(ctx, int32Id)
	if err != nil {
		logger.Error("Error deleting a module", err,
			zap.String("journey", "deleteModule"),
			zap.Bool("moduleExists", moduleExists),
			zap.Int32("moduleId", int32Id))

		return moduledomain.ErrCouldNotPersistModule(err)
	}

	logger.Info("Module deleted successfully",
		zap.String("journey", "createModule"),
		zap.Int32("moduleId", int32Id))

	return nil
}
