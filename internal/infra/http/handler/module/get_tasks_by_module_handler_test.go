package modulehandler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module/mocks"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	"github.com/golang/mock/gomock"
)

func TestGetTasksByModuleHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	getTasksByModuleUseCase := mock_modulecase.NewMockGetTasks(ctrl)
	getModulesByUserHandler := modulehandler.NewGetTasksByModuleHandler(getTasksByModuleUseCase)

	t.Run("Should be able to get tasks by module", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/modules/14/tasks", nil)
		rr := httptest.NewRecorder()

		taskQuantity := 10

		tasks := make([]taskdomain.Task, taskQuantity)
		for i := range taskQuantity {
			task := taskdomain.New(14, "New module", 1)
			tasks[i] = task
		}

		getTasksByModuleUseCase.EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Return(tasks, nil)

		getModulesByUserHandler.Execute(rr, req)

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		if rr.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d body: %s",
				http.StatusOK, rr.Code, body,
			)
		}

		var res []taskdto.TaskDTO
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("invalid return. Failed to unmarshal body: %v", err)
		}

		if len(res) != taskQuantity {
			t.Fatalf("expected 10 tasks, got %d", len(res))
		}
	})
}
