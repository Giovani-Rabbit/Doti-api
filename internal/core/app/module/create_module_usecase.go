package modulecase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"
	val "github.com/Giovani-Coelho/Doti-API/internal/pkg/validator"
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

	if !val.IsValidUUID(moduleEntity.GetUserId()) {
		logger.Error(
			"User id validation error", nil,
			zap.String("journey", "createModule"),
		)

		return nil, moduledomain.ErrInvalidUserId()
	}

	if !moduleEntity.IsValid() {
		logger.Error(
			"Module validation error", nil,
			zap.String("journey", "createModule"),
		)

		return nil, moduledomain.ErrInvalidModuleFields()
	}

	moduleCreated, err := mu.ModuleRepository.Create(ctx, moduleEntity)

	if err != nil {
		logger.Error(
			"Error in repository", err,
			zap.String("journey", "createModule"),
		)

		return nil, http.ErrInternal(
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
