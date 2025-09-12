package taskcase_test

import (
	"context"
	"testing"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestCreateTaskUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepository := mock_repository.NewMockTaskRepository(ctrl)
	moduleRepository := mock_repository.NewMockModuleRepository(ctrl)
	createTask := taskcase.NewCreateTaskUseCase(taskRepository, moduleRepository)

	ctx := context.Background()

	t.Run("Should be able to create a task", func(t *testing.T) {
		task := taskdomain.New(int32(1), "task", 2)

		moduleRepository.EXPECT().
			CheckExistsById(gomock.Any(), task.ModuleId()).
			Return(true, nil)

		taskRepository.EXPECT().
			Create(gomock.Any(), task).
			Return(task, nil)

		_, err := createTask.Execute(ctx, task)
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	})
}
