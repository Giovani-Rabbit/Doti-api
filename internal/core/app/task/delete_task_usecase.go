package taskcase

import (
	"context"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type Delete interface {
	Execute(ctx context.Context, taskId int32) error
}

type delete struct {
	taskRepository repository.TaskRepository
}

func NewDeleteTaskUseCase(
	taskRepo repository.TaskRepository,
) Delete {
	return &delete{
		taskRepository: taskRepo,
	}
}

func (dt *delete) Execute(
	ctx context.Context, taskId int32,
) error {
	logger.Info("Init delete task",
		zap.Int32("taskId", taskId),
		zap.String("journey", "deleteTask"))

	err := dt.taskRepository.Delete(ctx, taskId)
	if err != nil {
		logger.Error("Failed to delete task", err,
			zap.String("journey", "deleteTask"))

		return taskdomain.ErrCouldNotDeleteTask()
	}

	logger.Info("Task deleted",
		zap.String("journey", "deleteTask"))

	return nil
}
