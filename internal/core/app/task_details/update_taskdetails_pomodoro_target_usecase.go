package detailscase

import (
	"context"
	"errors"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	detailsdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task_details"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type UpdatePomodoroTarget interface {
	Execute(ctx context.Context, taskId int32, target int) error
}

type updatePomodoroTarget struct {
	taskDetailsRepo repository.TaskDetailsRepository
}

func NewTaskDetailsUpdatePomodoroTarget(
	taskDetailsRepo repository.TaskDetailsRepository,
) UpdatePomodoroTarget {
	return &updatePomodoroTarget{
		taskDetailsRepo: taskDetailsRepo,
	}
}

func (pt *updatePomodoroTarget) Execute(
	ctx context.Context, taskId int32, target int,
) error {
	logger.Info("Init update pomodoro target",
		zap.Int32("taskId", taskId),
		zap.Int("pomodoroTarget", target),
		zap.String("journey", "updatePomodoroTarget"))

	err := pt.taskDetailsRepo.UpdatePomodoroTarget(ctx, taskId, target)
	if err != nil {
		var restErr *resp.RestErr
		if errors.As(err, &restErr) {
			logger.Error("Could not find task", err,
				zap.Int32("taskId", taskId),
				zap.String("journey", "updatePomodoroTarget"))

			return err
		}

		logger.Error("Error updating pomodoro target", err,
			zap.Int32("taskId", taskId),
			zap.String("journey", "updatePomodoroTarget"))

		return detailsdomain.ErrUpdatingDescription(err)
	}

	logger.Info("pomodoro target updated",
		zap.String("journey", "updatePomodoroTarget"))

	return nil
}
