package detailscase

import (
	"context"
	"errors"
	"strings"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	detailsdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task_details"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type UpdateDescription interface {
	Execute(ctx context.Context, taskId int32, description string) error
}

type updateDescription struct {
	taskDetailsRepo repository.TaskDetailsRepository
}

func NewTaskDetailsUseCase(
	taskDetailsRepo repository.TaskDetailsRepository,
) UpdateDescription {
	return &updateDescription{
		taskDetailsRepo: taskDetailsRepo,
	}
}

func (td *updateDescription) Execute(
	ctx context.Context, taskId int32, description string,
) error {
	logger.Info("Init update description",
		zap.Int32("taskId", taskId),
		zap.String("description", description),
		zap.String("journey", "updateDescDetails"))

	if strings.TrimSpace(description) == "" {
		logger.Error("invalid description", nil,
			zap.String("description", description),
			zap.String("journey", "updateDescDetails"))

		return detailsdomain.ErrInvalidDescription()
	}

	err := td.taskDetailsRepo.UpdateDescription(ctx, taskId, description)
	if err != nil {
		var restErr *resp.RestErr
		if errors.As(err, &restErr) {
			logger.Error("Could not find task", err,
				zap.Int32("taskId", taskId),
				zap.String("journey", "updateDescDetails"))

			return err
		}

		logger.Error("Error updating description", err,
			zap.Int32("taskId", taskId),
			zap.String("journey", "updateDescDetails"))

		return detailsdomain.ErrUpdatingDescription(err)
	}

	logger.Info("Description updated",
		zap.String("journey", "updateDescDetails"))

	return nil
}
