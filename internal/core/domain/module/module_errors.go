package moduledomain

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

func ErrRenamingModule() *http.RestErr {
	return http.NewInternalServerError(
		"Error renaming module",
	)
}
