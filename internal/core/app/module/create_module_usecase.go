package modulecase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	val "github.com/Giovani-Coelho/Doti-API/internal/pkg/validator"
	"go.uber.org/zap"
)

type create struct {
	ModuleRepository repository.ModuleRepository
}

type Create interface {
	Execute(ctx context.Context, moduleEntity moduledomain.Module) (moduledomain.Module, error)
}

func NewCreateModuleUseCase(
	moduleRepo repository.ModuleRepository,
) Create {
	return &create{
		ModuleRepository: moduleRepo,
	}
}

func (mu *create) Execute(
	ctx context.Context,
	moduleEntity moduledomain.Module,
) (moduledomain.Module, error) {
	logger.Info("Init create module",
		zap.String("journey", "createModule"),
		zap.String("userId", moduleEntity.GetUserId()),
		zap.String("moduleName", moduleEntity.GetName()))

	if !val.IsValidUUID(moduleEntity.GetUserId()) {
		logger.Error("User id validation error", nil,
			zap.String("journey", "createModule"),
			zap.String("userId", moduleEntity.GetUserId()))

		return nil, moduledomain.ErrInvalidUserId()
	}

	if !moduleEntity.IsValid() {
		logger.Error("Module validation error", nil,
			zap.String("journey", "createModule"),
			zap.String("moduleName", moduleEntity.GetName()),
			zap.String("Icon", moduleEntity.GetIcon()))

		return nil, moduledomain.ErrInvalidModuleFields()
	}

	moduleCreated, err := mu.ModuleRepository.Create(ctx, moduleEntity)
	if err != nil {
		logger.Error("Error in repository", err,
			zap.String("journey", "createModule"),
			zap.String("userId", moduleEntity.GetUserId()),
			zap.String("moduleName", moduleEntity.GetName()),
			zap.String("Icon", moduleEntity.GetIcon()))

		return nil, moduledomain.ErrCouldNotPersistModule(err)
	}

	logger.Info("CreateModule executed successfully",
		zap.Int32("moduleID", moduleCreated.GetID()),
		zap.String("journey", "createModule"))

	return moduleCreated, nil
}
