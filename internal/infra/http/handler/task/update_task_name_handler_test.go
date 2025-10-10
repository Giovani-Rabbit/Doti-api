package taskhdl_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task/mocks"
	taskhdl "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/golang/mock/gomock"
)

func TestUpdateTaskName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock_taskcase.NewMockUpdateName(ctrl)
	handler := taskhdl.NewUpdateTaskNameHandler(usecase)

	t.Run("Should be able to update task name", func(t *testing.T) {
		requestBody := taskdto.UpdateTaskNameHttpBody{TaskName: "task 3"}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		usecase.EXPECT().
			Execute(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil)

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch, "/tasks/12/rename", bytes.NewReader(jsonBody),
		)

		mux := http.NewServeMux()
		mux.HandleFunc("PATCH /tasks/{id}/rename", handler.Execute)

		mux.ServeHTTP(rr, req)

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

	t.Run("Should fail if task id is invalid", func(t *testing.T) {
		requestBody := taskdto.UpdateTaskNameHttpBody{TaskName: "task 3"}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch, "/tasks/abc123/rename", bytes.NewReader(jsonBody),
		)

		mux := http.NewServeMux()
		mux.HandleFunc("PATCH /tasks/{id}/rename", handler.Execute)

		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf(
				"expected code %d, got %d",
				http.StatusNoContent, rr.Code,
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

		if response.Status != resp.SttInvalidPathValue {
			t.Fatalf("expected status %v, got %v",
				resp.SttInvalidPathValue,
				response.Status,
			)
		}
	})

	t.Run("Should fail if task body is invalid", func(t *testing.T) {
		requestBody := taskdto.UpdateCompletionDTO{IsComplete: false}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch, "/tasks/123/rename", bytes.NewReader(jsonBody),
		)

		mux := http.NewServeMux()
		mux.HandleFunc("PATCH /tasks/{id}/rename", handler.Execute)

		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf(
				"expected code %d, got %d",
				http.StatusNoContent, rr.Code,
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
				resp.SttInvalidPathValue,
				response.Status,
			)
		}
	})
}
