package modulecase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	val "github.com/Giovani-Coelho/Doti-API/internal/pkg/validator"
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

	if !val.IsValidUUID(userId) {
		return nil, val.ErrInvalidUUID()
	}

	modules, err := gm.moduleRepo.ListModulesByUserID(ctx, userId)

	if err != nil {
		logger.Error(
			"Error in repository", err,
			zap.String("journey", "getModule"),
		)

		return nil, moduledomain.ErrGettingModule(err)
	}

	logger.Info("getModule executed successfully",
		zap.Int("ModulesLength", len(modules)),
		zap.String("journey", "getModule"),
	)

	return modules, nil
}
