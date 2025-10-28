package detailshdl_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_detailscase "github.com/Giovani-Coelho/Doti-API/internal/core/app/task_details/mocks"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	detailshdl "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task_details"
	detailsdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task_details/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/golang/mock/gomock"
)

func TestTaskDetailsUpdateDescriptionHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock_detailscase.NewMockUpdateDescription(ctrl)
	handler := detailshdl.NewUpdateDescriptionHandler(usecase)

	t.Run("Should be able update the description", func(t *testing.T) {
		requestBody := detailsdto.UpdateDescriptionRequest{
			Description: "blablabla",
		}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch, "/tasks/123/details", bytes.NewReader(jsonBody),
		)

		usecase.EXPECT().
			Execute(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil)

		mux := http.NewServeMux()
		mux.HandleFunc("PATCH /tasks/{id}/details", handler.Execute)

		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Fatalf(
				"expected code %d, got %d",
				http.StatusCreated, rr.Code,
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

	t.Run("Should be able update the description", func(t *testing.T) {
		requestBody := taskdto.UpdateTaskNameHttpBody{
			TaskName: "taskname",
		}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch, "/tasks/123/details", bytes.NewReader(jsonBody),
		)

		mux := http.NewServeMux()
		mux.HandleFunc("PATCH /tasks/{id}/details", handler.Execute)

		mux.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf(
				"expected code %d, got %d",
				http.StatusCreated, rr.Code,
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
