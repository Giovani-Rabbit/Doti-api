package moduledomain

import "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"

const (
	SttCouldNotFindModuleByID = "COULD_NOT_FIND_MODULE_BY_ID"
	SttCouldNotGetModules     = "COULD_NOT_GET_MODULES"
	SttCouldNotPersistModule  = "COULD_NOT_PERSIST_MODULE"
	SttCouldNotRenameModule   = "COULD_NOT_RENAME_MODULE"
	SttDeletingModule         = "COULD_NOT_DELETE_MODULE"
	SttInvalidModuleFields    = "INVALID_MODULE_FIELDS"
	SttInvalidModuleID        = "INVALID_MODULE_ID"
	SttInvalidUserID          = "INVALID_USER_ID"
	SttNewModuleNameIsEmpty   = "NEW_MODULE_NAME_IS_EMPTY"
)

func ErrInvalidModuleFields() *http.RestErr {
	return http.NewBadRequestError(
		SttInvalidModuleFields,
		"Missing required fields",
	)
}

func ErrInvalidUserId() *http.RestErr {
	return http.NewBadRequestError(
		SttInvalidUserID,
		"The user id is invalid",
	)
}

func ErrInvalidModuleID() *http.RestErr {
	return http.NewBadRequestError(
		SttInvalidModuleID,
		"The module id is not a uuid",
	)
}

func ErrNewModuleNameIsEmpty() *http.RestErr {
	return http.NewBadRequestError(
		SttNewModuleNameIsEmpty,
		"The new module name is empty",
	)
}

func ErrCouldNotPersistModule(err error) *http.RestErr {
	return http.ErrInternal(
		"Error saving module",
		SttCouldNotPersistModule,
		err,
	)
}

func ErrCouldNotFindModuleByID() *http.RestErr {
	return http.NewCouldNotFind(
		"The module was not found",
		SttCouldNotFindModuleByID,
	)
}
