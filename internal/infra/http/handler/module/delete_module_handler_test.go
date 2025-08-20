package modulehandler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_modulecase "github.com/Giovani-Coelho/Doti-API/internal/core/app/module/mocks"
	modulehandler "github.com/Giovani-Coelho/Doti-API/internal/infra/http/handler/module"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestDeleteModuleHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	deleteModuleUseCase := mock_modulecase.NewMockDelete(ctrl)
	deleteModuleHandler := modulehandler.NewDeleteHandler(deleteModuleUseCase)

	t.Run("Should be able to delete a module", func(*testing.T) {
		uuid := uuid.New()

		req := httptest.NewRequest(http.MethodDelete, "/modules/"+uuid.String(), nil)
		rr := httptest.NewRecorder()

		deleteModuleUseCase.EXPECT().
			Execute(gomock.Any(), gomock.Any()).
			Return(nil)

		deleteModuleHandler.Execute(rr, req)

		body, err := io.ReadAll(rr.Body)
		if err != nil {
			t.Fatalf("failed to get body: %v", err)
		}

		if rr.Code != http.StatusNoContent {
			t.Fatalf("expected code %v, got %v", http.StatusNoContent, rr.Code)
		}

		if len(body) != 0 {
			t.Fatalf("expected empty body, got %v", body)
		}
	})
}
