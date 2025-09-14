package taskcase_test

import (
	"context"
	"errors"
	"testing"

	taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task"
	moduledomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/module"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	mock_repository "github.com/Giovani-Coelho/Doti-API/internal/infra/persistence/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestGetTasksByModuleId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskRepository := mock_repository.NewMockTaskRepository(ctrl)
	moduleRepository := mock_repository.NewMockModuleRepository(ctrl)
	getTaskList := taskcase.NewGetTasksByModuleId(taskRepository, moduleRepository)

	ctx := context.Background()

	var tasks []taskdomain.Task
	taskQuantity := 6

	for i := range taskQuantity {
		tasks = append(tasks, taskdomain.New(i, "task", i))
	}

	t.Run("Should be able to get tasks", func(t *testing.T) {
		moduleId := int32(321)

		moduleRepository.EXPECT().
			CheckExistsById(gomock.Any(), moduleId).
			Return(true, nil)

		taskRepository.EXPECT().
			ListByModuleId(gomock.Any(), moduleId).
			Return(tasks, nil)

		taskList, err := getTaskList.Execute(ctx, moduleId)
		if err != nil {
			t.Fatalf("No error was expected, got: %v", err)
		}

		if len(taskList) != taskQuantity {
			t.Fatalf("Invalid task quantity, expected %v, got: %v",
				taskQuantity, len(taskList),
			)
		}
	})

	t.Run("Should fail if the module does not exists", func(t *testing.T) {
		moduleId := int32(321)

		moduleRepository.EXPECT().
			CheckExistsById(gomock.Any(), moduleId).
			Return(false, nil)

		taskList, err := getTaskList.Execute(ctx, moduleId)
		if err == nil {
			t.Fatalf("An error was expected, got: %v", err)
		}

		if len(taskList) == taskQuantity {
			t.Fatalf("Expected a empty task list, got: %v", taskList)
		}

		res := resp.AsRestErr(err)

		if res.Status != moduledomain.SttCouldNotFindModuleByID {
			t.Fatalf("Expected status %v, got: %v",
				moduledomain.SttCouldNotFindModuleByID, res.Status,
			)
		}
	})

	t.Run("Should fail to check if module exists", func(t *testing.T) {
		moduleId := int32(321)

		moduleRepository.EXPECT().
			CheckExistsById(gomock.Any(), moduleId).
			Return(false, errors.New("repository error"))

		taskList, err := getTaskList.Execute(ctx, moduleId)
		if err == nil {
			t.Fatalf("An error was expected, got: %v", err)
		}

		if len(taskList) == taskQuantity {
			t.Fatalf("Expected a empty task list, got: %v", taskList)
		}

		res := resp.AsRestErr(err)

		if res.Status != moduledomain.SttCouldNotVerifyIfModuleExists {
			t.Fatalf("Expected status %v, got: %v",
				moduledomain.SttCouldNotVerifyIfModuleExists, res.Status,
			)
		}
	})

	t.Run("Should fail to get the task list", func(t *testing.T) {
		moduleId := int32(321)

		moduleRepository.EXPECT().
			CheckExistsById(gomock.Any(), moduleId).
			Return(true, nil)

		taskRepository.EXPECT().
			ListByModuleId(gomock.Any(), moduleId).
			Return(nil, errors.New("repository error"))

		taskList, err := getTaskList.Execute(ctx, moduleId)
		if err == nil {
			t.Fatalf("An error was expected, got: %v", err)
		}

		if len(taskList) == taskQuantity {
			t.Fatalf("Expected a empty task list, got: %v", taskList)
		}

		res := resp.AsRestErr(err)

		if res.Status != taskdomain.SttCouldNotListTasks {
			t.Fatalf("Expected status %v, got: %v",
				taskdomain.SttCouldNotListTasks, res.Status,
			)
		}
	})
}
