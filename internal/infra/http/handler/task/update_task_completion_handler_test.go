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

func TestUpdateTaskCompletion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock_taskcase.NewMockUpdateCompletion(ctrl)
	handler := taskhdl.NewUpdateCompletionHandler(usecase)

	t.Run("Should be able to update the task completion", func(t *testing.T) {
		taskCompletion := taskdto.UpdateCompletionDTO{
			IsComplete: true,
		}

		jsonBody, err := json.Marshal(taskCompletion)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch, "/tasks/12", bytes.NewReader(jsonBody),
		)

		usecase.EXPECT().
			Execute(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil)

		mux := http.NewServeMux()
		mux.HandleFunc("PATCH /tasks/{id}", handler.Execute)

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

	t.Run("Should fail with a invalid param", func(t *testing.T) {
		taskCompletion := taskdto.UpdateCompletionDTO{
			IsComplete: true,
		}

		jsonBody, err := json.Marshal(taskCompletion)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch, "/tasks/abc", bytes.NewReader(jsonBody),
		)

		mux := http.NewServeMux()
		mux.HandleFunc("PATCH /tasks/{id}", handler.Execute)

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

		var res *resp.RestErr
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("invalid return. Failed to unmarshal body: %v", err)
		}

		if res.Status != resp.SttInvalidPathValue {
			t.Fatalf("expected status %v, got: %v",
				resp.SttInvalidPathValue,
				res.Status,
			)
		}
	})

	t.Run("Should fail with a invalid body", func(t *testing.T) {
		taskCompletion := taskdto.UpdatePositionDTO{
			MovedTasks: nil,
		}

		jsonBody, err := json.Marshal(taskCompletion)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch, "/tasks/123", bytes.NewReader(jsonBody),
		)

		mux := http.NewServeMux()
		mux.HandleFunc("PATCH /tasks/{id}", handler.Execute)

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

		var res *resp.RestErr
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("invalid return. Failed to unmarshal body: %v", err)
		}

		if res.Status != resp.SttInvalidRequestBody {
			t.Fatalf("expected status %v, got: %v",
				resp.SttInvalidRequestBody,
				res.Status,
			)
		}
	})
}
