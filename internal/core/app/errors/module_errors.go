package apperr

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	SttCouldNotRenameModule = "COULD_NOT_RENAME_MODULE"
	SttCouldNotGetModules   = "COULD_NOT_GET_MODULES"
)

func ErrRenamingModule(err error) *http.RestErr {
	return http.NewRestError(
		SttCouldNotRenameModule,
		"Error renaming module",
		err,
	)
}

func ErrGettingModule(err error) *http.RestErr {
	return http.NewRestError(
		SttCouldNotGetModules,
		"Internal error getting modules",
		err,
	)
}
