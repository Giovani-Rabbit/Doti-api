package modulecase

import (
	"context"
	"strings"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	val "github.com/Giovani-Coelho/Doti-API/internal/pkg/validator"
	"go.uber.org/zap"
)

type updateIcon struct {
	moduleRepository repository.ModuleRepository
}

type UpdateIcon interface {
	Execute(ctx context.Context, moduleId string, icon string) error
}

func NewUpdateModuleIconUseCase(
	moduleRepo repository.ModuleRepository,
) UpdateIcon {
	return &updateIcon{
		moduleRepository: moduleRepo,
	}
}

func (ui *updateIcon) Execute(
	ctx context.Context, moduleId string, icon string,
) error {
	logger.Info("Init change module icon",
		zap.String("journey", "changeIcon"),
	)

	if !val.IsValidUUID(moduleId) {
		logger.Error("Invalid module id", nil,
			zap.String("journey", "changeIcon"))

		return moduledomain.ErrInvalidModuleID()
	}

	if strings.TrimSpace(icon) == "" {
		logger.Error("Error, the new icon is empty", nil,
			zap.String("journey", "changeIcon"),
		)

		return moduledomain.ErrNewModuleIconIsEmpty()
	}

	icon = strings.TrimSpace(icon)

	err := ui.moduleRepository.UpdateIcon(ctx, moduleId, icon)

	if err != nil {
		logger.Error("Error on update module icon", err,
			zap.String("journey", "changeIcon"),
		)

		return moduledomain.ErrCouldNotPersistModule(err)
	}

	logger.Info("Icon updated successfully",
		zap.String("journey", "changeIcon"),
	)

	return nil
}
