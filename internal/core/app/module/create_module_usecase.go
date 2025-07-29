package modulecase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	rest_err "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"
	"go.uber.org/zap"
)

type CreateModuleUseCase struct {
	ModuleRepository repository.IModuleRepository
}

type ICreateModuleUseCase interface {
	Execute(ctx context.Context, moduleEntity moduledomain.IModuleDomain) (moduledomain.IModuleDomain, error)
}

func NewCreateModuleUseCase(
	moduleRepo repository.IModuleRepository,
) ICreateModuleUseCase {
	return &CreateModuleUseCase{
		ModuleRepository: moduleRepo,
	}
}

func (mu *CreateModuleUseCase) Execute(
	ctx context.Context,
	moduleEntity moduledomain.IModuleDomain,
) (moduledomain.IModuleDomain, error) {
	logger.Info("Init create module",
		zap.String("journey", "createModule"),
	)

	moduleCreated, err := mu.ModuleRepository.Create(ctx, moduleEntity)

	if err != nil {
		logger.Error(
			"Error in repository", err,
			zap.String("journey", "createModule"),
		)

		return nil, rest_err.NewInternalServerError(
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
