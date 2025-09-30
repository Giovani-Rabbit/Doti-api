package taskcase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type UpdateCompletion interface {
	Execute(ctx context.Context, taskId int32, isComplete bool) error
}

type updateCompletion struct {
	taskRepo repository.TaskRepository
}

func NewUpdateCompletion(
	tr repository.TaskRepository,
) UpdateCompletion {
	return &updateCompletion{
		taskRepo: tr,
	}
}

func (uc *updateCompletion) Execute(
	ctx context.Context, taskId int32, isComplete bool,
) error {
	logger.Info("Init update task completion",
		zap.Int32("taskId", taskId),
		zap.Bool("isComplete", isComplete),
		zap.String("journey", "updateTaskCompletion"))

	exists, err := uc.taskRepo.CheckExists(ctx, taskId)
	if err != nil {
		logger.Info("Failed to check if task exists",
			zap.Int32("taskId", taskId),
			zap.String("journey", "updateTaskCompletion"))

		return taskdomain.ErrCheckingIfTaskExists(err)
	}

	if !exists {
		logger.Info("Task not found",
			zap.Int32("taskid", taskId),
			zap.String("journey", "updateTaskCompletion"))

		return taskdomain.ErrCouldNotFindTask()
	}

	err = uc.taskRepo.UpdateCompletion(ctx, taskId, isComplete)
	if err != nil {
		logger.Info("Failed to update task completion",
			zap.Int32("taskId", taskId),
			zap.Bool("isComplete", isComplete),
			zap.String("journey", "updateTaskCompletion"))

		return taskdomain.ErrCouldNotUpdateTask(err)
	}

	logger.Info("Task completion updated",
		zap.String("journey", "updateTaskCompletion"))

	return nil
}
