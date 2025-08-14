package modulecase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	val "github.com/Giovani-Coelho/Doti-API/internal/pkg/validator"
	"go.uber.org/zap"
)

type DeleteModuleUseCase interface {
	Execute(ctx context.Context, id string) error
}

type deleteModuleUseCase struct {
	moduleRepository repository.ModuleRepository
}

func NewDeleteModuleUseCase(
	moduleRepo repository.ModuleRepository,
) DeleteModuleUseCase {
	return &deleteModuleUseCase{
		moduleRepository: moduleRepo,
	}
}

func (dm *deleteModuleUseCase) Execute(ctx context.Context, id string) error {
	logger.Info("Init delete module usecase",
		zap.String("journey", "deleteModule"),
	)

	if !val.IsValidUUID(id) {
		logger.Error("Invalid, id is not a uuid", nil,
			zap.String("journey", "deleteModule"),
		)

		return moduledomain.ErrInvalidModuleID()
	}

	moduleExists, err := dm.moduleRepository.CheckExistsById(ctx, id)

	if err != nil {
		logger.Error("Error finding a module", err,
			zap.String("journey", "deleteModule"),
		)

		return moduledomain.ErrCouldNotPersistModule(err)
	}

	if !moduleExists {
		logger.Info("Module not find",
			zap.String("journey", "deleteModule"),
		)

		return moduledomain.ErrCouldNotFindModuleByID()
	}

	err = dm.moduleRepository.DeleteModule(ctx, id)

	if err != nil {
		logger.Error("Error deleting a module", err,
			zap.String("journey", "deleteModule"),
		)

		return moduledomain.ErrCouldNotPersistModule(err)
	}

	logger.Info(
		"Module deleted successfully",
		zap.String("journey", "createModule"),
	)

	return nil
}
