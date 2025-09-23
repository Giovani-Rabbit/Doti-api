package taskcase

import (
	"context"
	"errors"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	"github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository"
	"go.uber.org/zap"
)

type UpdatePosition interface {
	Execute(ctx context.Context, params *taskdto.UpdatePositionDTO) error
}

type updatePosition struct {
	taskRepository repository.TaskRepository
}

func NewUpdateTaskPosition(
	tr repository.TaskRepository,
) UpdatePosition {
	return &updatePosition{
		taskRepository: tr,
	}
}

func (tp *updatePosition) Execute(
	ctx context.Context, params *taskdto.UpdatePositionDTO,
) error {
	logger.Info("Update task position",
		zap.String("journey", "updateTaskPosition"))

	if len(params.Tasks) != 2 {
		logger.Error("The tasks number must be equal to 2", nil,
			zap.Int("tasks_quantity", len(params.Tasks)),
			zap.String("journey", "updateTaskPosition"))

		return errors.New("expected two tasks to be updated")
	}

	if params.Tasks[0].Position == params.Tasks[1].Position {
		logger.Error("The task position cannot be equal", nil,
			zap.Int("position_0", int(params.Tasks[0].Position)),
			zap.Int("position_1", int(params.Tasks[1].Position)))

		return errors.New("the positions must be different")
	}

	err := tp.taskRepository.UpdatePosition(ctx, params.Tasks)
	if err != nil {
		logger.Error("error on update task position", err,
			zap.String("journey", "updateTaskPosition"))

		return err
	}

	logger.Info("task position updated succesfully",
		zap.String("journey", "updateTaskposition"))

	return nil
}
