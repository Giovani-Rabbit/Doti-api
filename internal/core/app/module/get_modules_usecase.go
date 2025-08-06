package modulecase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"
	"go.uber.org/zap"
)

type GetModulesUseCase interface {
	Execute(ctx context.Context, userId string) ([]moduledomain.Module, error)
}

type getModulesUseCase struct {
	moduleRepo repository.ModuleRepository
}

func NewGetModulesUseCase(
	moduleRepo repository.ModuleRepository,
) GetModulesUseCase {
	return &getModulesUseCase{
		moduleRepo: moduleRepo,
	}
}

func (gm *getModulesUseCase) Execute(
	ctx context.Context, userId string,
) ([]moduledomain.Module, error) {
	logger.Info("Init get module",
		zap.String("journey", "getModule"),
	)

	modules, err := gm.moduleRepo.ListModulesByUserID(ctx, userId)

	if err != nil {
		logger.Error(
			"Error in repository", err,
			zap.String("journey", "getModule"),
		)

		return nil, http.NewInternalServerError(
			"Internal error getting module",
		)
	}

	logger.Info("getModule executed successfully",
		zap.Int("ModulesLength", len(modules)),
		zap.String("journey", "getModule"),
	)

	return modules, nil
}
