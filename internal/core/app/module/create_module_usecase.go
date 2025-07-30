package modulecase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"
	"go.uber.org/zap"
)

type createModuleUseCase struct {
	ModuleRepository repository.IModuleRepository
}

type CreateModuleUseCase interface {
	Execute(ctx context.Context, moduleEntity moduledomain.Module) (moduledomain.Module, error)
}

func NewCreateModuleUseCase(
	moduleRepo repository.IModuleRepository,
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

	moduleCreated, err := mu.ModuleRepository.Create(ctx, moduleEntity)

	if err != nil {
		logger.Error(
			"Error in repository", err,
			zap.String("journey", "createModule"),
		)

		return nil, http.NewInternalServerError(
			"Internal error saving module",
		)
	}

	logger.Info(
		"CreateModule executed successfully",
		zap.String("moduleID", moduleCreated.GetID()),
		zap.String("journey", "createModule"),
	)

	return moduleCreated, nil
}
