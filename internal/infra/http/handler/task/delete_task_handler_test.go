package taskhdl_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_taskcase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task/mocks"
	taskhdl "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/golang/mock/gomock"
)

func TestDeleteTaskHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock_taskcase.NewMockDelete(ctrl)
	handler := taskhdl.NewDeleteTaskHandler(usecase)

	t.Run("Should be able to delete task", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPatch, "/tasks/12", nil)

		usecase.EXPECT().
			Execute(gomock.Any(), gomock.Any()).
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

	t.Run("Should fail with an invalid taskId", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPatch, "/tasks/1abc2", nil)

		mux := http.NewServeMux()
		mux.HandleFunc("PATCH /tasks/{id}", handler.Execute)

		mux.ServeHTTP(rr, req)

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

		var res resp.RestErr
		if err := json.Unmarshal(body, &res); err != nil {
			t.Fatalf("failet to unmarshal body, err: %v", err)
		}

		if res.Status != resp.SttInvalidPathValue {
			t.Fatalf("expected status %v, got: %v",
				resp.SttInvalidPathValue,
				res.Status,
			)
		}
	})
}
