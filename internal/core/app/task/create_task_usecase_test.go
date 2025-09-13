package taskcase_test

import (
	"context"
	"testing"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
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
		task := taskdomain.New(1, "task", 2)

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

	t.Run("Should fail if the fields are invalid", func(t *testing.T) {
		task := taskdomain.New(1, "", 0)

		_, err := createTask.Execute(ctx, task)
		if err == nil {
			t.Errorf("Error was expected, got: %v", err)
		}

		respErr := resp.AsRestErr(err)

		if respErr.Status != taskdomain.SttInvalidFields {
			t.Errorf("Expected %v error, got: %v",
				taskdomain.SttInvalidFields,
				respErr.Status,
			)
		}
	})

	t.Run("Should fail if module does not exists", func(t *testing.T) {
		task := taskdomain.New(1, "task", 2)

		moduleRepository.EXPECT().
			CheckExistsById(gomock.Any(), task.ModuleId()).
			Return(false, nil)

		_, err := createTask.Execute(ctx, task)
		if err == nil {
			t.Errorf("Expected error, got: %v", err)
		}

		respErr := resp.AsRestErr(err)

		if respErr.Status != moduledomain.SttCouldNotFindModuleByID {
			t.Errorf("Expected status %v, got: %v",
				moduledomain.SttCouldNotFindModuleByID,
				respErr.Status,
			)
		}
	})
}
