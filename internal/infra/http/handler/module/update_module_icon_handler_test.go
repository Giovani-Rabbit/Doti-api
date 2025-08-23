package modulehandler_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module/mocks"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TeseUpdateModuleIconHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	renameModuleUseCase := mock_modulecase.NewMockUpdateIcon(ctrl)
	renameModuleHandler := modulehandler.NewUpdateIconHandler(renameModuleUseCase)

	uuidstg := uuid.NewString()

	t.Run("Should be able to change the icon", func(t *testing.T) {
		requestBody := moduledto.UpdateIconDTO{Icon: "code"}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch,
			"/modules/"+uuidstg,
			bytes.NewReader(jsonBody),
		)

		renameModuleUseCase.EXPECT().
			Execute(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil)

		renameModuleHandler.Execute(rr, req)

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
			t.Fatalf("expected empty body, got %v", body)
		}
	})

	t.Run("Should fail if it is a invalid request body ", func(t *testing.T) {
		requestBody := moduledto.NewModuleNameDTO{Name: "giovani"}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodPatch,
			"/modules/"+uuidstg,
			bytes.NewReader(jsonBody),
		)

		renameModuleHandler.Execute(rr, req)

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

		var res *resp.RestErr
		err = json.Unmarshal(body, &res)
		if err != nil {
			t.Fatalf("invalid return. Failed to unmarshal body: %v", err)
		}

		if res.Status != resp.SttInvalidRequestBody {
			t.Fatalf("expected invalid request body, got %v", res.Status)
		}
	})
}
