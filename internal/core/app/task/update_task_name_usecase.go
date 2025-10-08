package taskcase

import (
	"context"
	"errors"
	"strings"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type UpdateName interface {
	Execute(ctx context.Context, taskId int32, name string) error
}

type updateName struct {
	taskRepository repository.TaskRepository
}

func NewUpdateTaskName(
	taskRepository repository.TaskRepository,
) UpdateName {
	return &updateName{
		taskRepository: taskRepository,
	}
}

func (un *updateName) Execute(
	ctx context.Context, taskId int32, name string,
) error {
	logger.Info("update task name",
		zap.Int32("taskId", taskId),
		zap.String("taskName", name),
		zap.String("journey", "updateTaskName"))

	if strings.TrimSpace(name) == "" || len(name) > 100 {
		logger.Error("invalid task name", nil,
			zap.String("taskName", name),
			zap.String("journey", "updateTaskName"))

		return taskdomain.ErrInvalidTaskName()
	}

	err := un.taskRepository.UpdateName(ctx, taskId, name)
	if err != nil {
		var restErr *resp.RestErr
		if errors.As(err, &restErr) {
			if restErr.Status == taskdomain.SttCouldNotFindTask {
				logger.Error("task not found", err,
					zap.Int32("taskId", taskId),
					zap.String("taskName", name),
					zap.String("journey", "updateTaskName"))

				return err
			}
		}

		logger.Error("error updating task name", err,
			zap.Int32("taskId", taskId),
			zap.String("taskName", name),
			zap.String("journey", "updateTaskName"))

		return taskdomain.ErrCouldNotUpdateTask(err)
	}

	logger.Info("task name updated",
		zap.String("journey", "updateTaskName"))

	return nil
}
