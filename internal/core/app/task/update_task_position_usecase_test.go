package taskcase_test

import (
	"context"
	"testing"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestChangeTaskPosition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tr := mock_repository.NewMockTaskRepository(ctrl)
	changeTaskPosition := taskcase.NewUpdateTaskPosition(tr)

	ctx := context.Background()

	t.Run("Should be able to change the task position", func(t *testing.T) {
		tasks := []taskdto.TaskPositionParams{
			{TaskId: 1, Position: 2},
			{TaskId: 2, Position: 4},
		}

		updatatePosition := taskdto.UpdatePositionDTO{Tasks: tasks}

		tr.EXPECT().
			UpdatePosition(ctx, gomock.Any()).
			Return(nil)

		err := changeTaskPosition.Execute(ctx, &updatatePosition)
		if err != nil {
			t.Errorf("Error was not expected, got: %v", err)
		}
	})

	t.Run("Should fail if the number of tasks is different than 2", func(t *testing.T) {
		tasks := []taskdto.TaskPositionParams{
			{TaskId: 1, Position: 2},
		}

		updatatePosition := taskdto.UpdatePositionDTO{Tasks: tasks}

		err := changeTaskPosition.Execute(ctx, &updatatePosition)
		if err == nil {
			t.Errorf("Error was expected, got: %v", err)
		}
	})

	t.Run("Should fail if position of the tasks are equals", func(t *testing.T) {
		tasks := []taskdto.TaskPositionParams{
			{TaskId: 1, Position: 2},
			{TaskId: 2, Position: 2},
		}

		updatatePosition := taskdto.UpdatePositionDTO{Tasks: tasks}

		err := changeTaskPosition.Execute(ctx, &updatatePosition)
		if err == nil {
			t.Errorf("Error was expected, got: %v", err)
		}
	})
}
