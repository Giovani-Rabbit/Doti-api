package modulehandler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module/mocks"
	authdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/auth"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
	moduledto "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module/dtos"
	"github.com/Giovani-Coelho/Doti-API/internal/pkg/auth"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestRenameModuleHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	renameModuleUseCase := mock_modulecase.NewMockRename(ctrl)
	renameModuleHandler := modulehandler.NewRenameHandler(renameModuleUseCase)

	authUser := &authdomain.AuthClaims{
		ID:    uuid.NewString(),
		Name:  "Giovani",
		Email: "giovani@example.com",
	}

	t.Run("Should be able to rename the module", func(t *testing.T) {
		requestBody := moduledto.NewModuleNameDTO{Name: "anything"}

		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}

		uuidstg := uuid.NewString()
		req := httptest.NewRequest(
			http.MethodPatch,
			"/modules/"+uuidstg,
			bytes.NewReader(jsonBody),
		)

		ctx := context.WithValue(req.Context(), auth.AuthenticatedUserKey, authUser)

		req = req.WithContext(ctx)
		rr := httptest.NewRecorder()

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
}
