package moduledomain

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	SttCouldNotRenameModule = "COULD_NOT_RENAME_MODULE"
	SttCouldNotGetModules   = "COULD_NOT_GET_MODULES"
	SttInvalidModuleID      = "INVALID_MODULE_ID"
	SttNewModuleNameIsEmpty = "NEW_MODULE_NAME_IS_EMPTY"
)

func ErrInvalidModuleFields(err error) *http.RestErr {
	return http.NewBadRequestError(
		SttInvalidModuleID,
		err.Error(),
	)
}

func ErrInvalidModuleID() *http.RestErr {
	return http.NewBadRequestError(
		SttInvalidModuleID,
		"The module id is not a uuid",
	)
}

func ErrRenamingModule(err error) *http.RestErr {
	return http.NewRestError(
		SttCouldNotRenameModule,
		"Error renaming module",
		err,
	)
}

func ErrNewModuleNameIsEmpty() *http.RestErr {
	return http.NewBadRequestError(
		SttNewModuleNameIsEmpty,
		"The new module name is empty",
	)
}

func ErrGettingModule(err error) *http.RestErr {
	return http.NewRestError(
		SttCouldNotGetModules,
		"Internal error getting modules",
		err,
	)
}
