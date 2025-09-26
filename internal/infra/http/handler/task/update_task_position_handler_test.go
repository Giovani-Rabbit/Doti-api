package taskhdl_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task/mocks"
	taskdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/task"
	taskhdl "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/golang/mock/gomock"
)

func TestUpdateTaskPosition(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock_taskcase.NewMockUpdatePosition(ctrl)
	handler := taskhdl.NewUpdatePositionHandler(usecase)

	t.Run("Should be able to move the task position", func(t *testing.T) {
		movedTasks := []taskdomain.TaskPositionParams{
			{Id: 1, Position: 4},
			{Id: 2, Position: 6},
		}
		requestBody := taskdto.UpdatePositionDTO{MovedTasks: movedTasks}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch, "/tasks", bytes.NewReader(jsonBody),
		)

		usecase.EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Return(nil)

		handler.Execute(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Fatalf(
				"expected code %d, got %d",
				http.StatusNoContent, rr.Code,
			)
		}

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		if len(body) != 0 {
			t.Fatalf("expected a empty body, got: %v", body)
		}
	})

	t.Run("Should fail if the body is invalid", func(t *testing.T) {
		requestBody := taskdto.CreateTaskDTO{}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch, "/tasks", bytes.NewReader(jsonBody),
		)

		handler.Execute(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf(
				"expected code %d, got %d",
				http.StatusBadRequest, rr.Code,
			)
		}

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		var response resp.RestErr
		if err := json.Unmarshal(body, &response); err != nil {
			t.Fatalf("failed to unmarshal body: %v", err)
		}

		if response.Status != resp.SttInvalidRequestBody {
			t.Fatalf("expected status %v, got %v",
				resp.SttInvalidRequestBody,
				response.Status,
			)
		}
	})
}
