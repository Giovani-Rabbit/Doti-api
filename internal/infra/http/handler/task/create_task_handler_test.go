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
	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	taskhdl "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task"
	taskdto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/task/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/golang/mock/gomock"
)

func TestCreateTaskHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createTaskCase := mock_taskcase.NewMockCreate(ctrl)
	createTaskHandler := taskhdl.NewCreateTaskHandler(createTaskCase)

	t.Run("Should be able to create task", func(t *testing.T) {
		requestBody := taskdto.CreateTaskDTO{
			ModuleId: 1,
			TaskName: "New Task",
			Position: 0,
		}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPost, "/task", bytes.NewReader(jsonBody),
		)

		taskEntity := taskdomain.New(
			requestBody.ModuleId,
			requestBody.TaskName,
			requestBody.Position,
		)

		createTaskCase.EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Return(taskEntity, nil)

		createTaskHandler.Execute(rr, req)

		if rr.Code != http.StatusCreated {
			t.Fatalf(
				"expected code %d, got %d",
				http.StatusNoContent, rr.Code,
			)
		}

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		var response taskdto.TaskDTO
		if err := json.Unmarshal(body, &response); err != nil {
			t.Fatalf("failed to unmarshal body: %v", err)
		}

		if response.ModuleId != int32(requestBody.ModuleId) {
			t.Fatalf("The moduleId is different")
		}
	})

	t.Run("Should fail with a invalid body", func(t *testing.T) {
		requestBody := moduledto.UpdateIconDTO{Icon: "123"}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPost, "/task", bytes.NewReader(jsonBody),
		)

		createTaskHandler.Execute(rr, req)

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
				resp.SttInvalidRequestBody,
				response.Status,
			)
		}
	})
}
