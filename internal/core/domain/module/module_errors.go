package moduledomain

import resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"

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
	SttNewIconNameIsEmpty     = "NEW_ICON_NAME_IS_EMPTY"
)

func ErrInvalidModuleFields() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidModuleFields,
		"Missing required fields",
	)
}

func ErrInvalidUserId() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidUserID,
		"The user id is invalid",
	)
}

func ErrInvalidModuleID() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidModuleID,
		"The module id is not a uuid",
	)
}

func ErrNewModuleNameIsEmpty() *resp.RestErr {
	return resp.NewBadRequestError(
		SttNewModuleNameIsEmpty,
		"The new module name is empty",
	)
}

func ErrNewModuleIconIsEmpty() *resp.RestErr {
	return resp.NewBadRequestError(
		SttNewIconNameIsEmpty,
		"The new module icon is empty",
	)
}

func ErrCouldNotPersistModule(err error) *resp.RestErr {
	return resp.NewErrInternal(
		"Error saving module",
		SttCouldNotPersistModule,
		err,
	)
}

func ErrCouldNotFindModuleByID() *resp.RestErr {
	return resp.NewCouldNotFind(
		"The module was not found",
		SttCouldNotFindModuleByID,
	)
}
