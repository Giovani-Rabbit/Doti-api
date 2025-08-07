package modulecase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	apperr "github.com/Giovani-Coelho/Doti-API/internal/core/domain/errors"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type createModuleUseCase struct {
	ModuleRepository repository.ModuleRepository
}

type CreateModuleUseCase interface {
	Execute(ctx context.Context, moduleEntity moduledomain.Module) (moduledomain.Module, error)
}

func NewCreateModuleUseCase(
	moduleRepo repository.ModuleRepository,
) CreateModuleUseCase {
	return &createModuleUseCase{
		ModuleRepository: moduleRepo,
	}
}

func (mu *createModuleUseCase) Execute(
	ctx context.Context,
	moduleEntity moduledomain.Module,
) (moduledomain.Module, error) {
	logger.Info("Init create module",
		zap.String("journey", "createModule"),
	)

	if err := moduleEntity.IsValid(); err != nil {
		logger.Error(
			"Validation domain error", nil,
			zap.String("journey", "createModule"),
		)

		return nil, apperr.ErrValidationDomain(err)
	}

	moduleCreated, err := mu.ModuleRepository.Create(ctx, moduleEntity)

	if err != nil {
		logger.Error(
			"Error in repository", err,
			zap.String("journey", "createModule"),
		)

		return nil, apperr.ErrInternal(
			"Error saving module", err,
		)
	}

	logger.Info(
		"CreateModule executed successfully",
		zap.String("moduleID", moduleCreated.GetID()),
		zap.String("journey", "createModule"),
	)

	return moduleCreated, nil
}
