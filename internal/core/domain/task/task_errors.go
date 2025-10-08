package taskdomain

import (
	"fmt"

	resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"
)

const (
	SttCouldNotPersist        = "COULD_NOT_PERSIST_TASK"
	SttCouldNotListTasks      = "COULD_NOT_LIST_TASKS"
	SttCouldNotVerifyPosition = "COULD_NOT_VERIFY_POSITION"
	SttCouldNotUpdateTask     = "COULD_NOT_UPDATE_TASK"
	SttCouldNotFindTask       = "COULD_NOT_FIND_TASK"
	SttCouldNotFindOwner      = "COULD_NOT_FIND_OWNER"
	SttCouldNotDeleteTask     = "COULD_NOT_DELETE_TASK"
	SttInvalidTaskOwner       = "INVALID_TASK_OWNER"
	SttInternalRepositoryErr  = "INTERNAL_REPOSITORY_ERROR"
	SttInvalidFields          = "INVALID_TASK_FIELDS"
	SttInvalidTaskName        = "INVALID_TASK_NAME"
	SttRepeatedPosition       = "REPEATED_TASK_POSITION"
	SttUnavailablePosition    = "UNAVAILABLE_POSITION"
)

func ErrInvalidFields() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidFields,
		"The fields required to create a task are invalid",
	)
}

func ErrInvalidTaskOwner() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidTaskOwner,
		"This task does not belong to the logged in user",
	)
}

func ErrInvalidTaskName() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidTaskName,
		"Invalid task name, please try again.",
	)
}

func ErrCouldNotFindTask() *resp.RestErr {
	return resp.NewBadRequestError(
		SttCouldNotFindTask,
		"Could not find task by id",
	)
}

func ErrCouldNotFindOwner(err error) *resp.RestErr {
	return resp.NewBadRequestError(
		SttCouldNotFindOwner,
		fmt.Sprintf("Failed to get task owner. %v", err.Error()),
	)
}

func ErrCouldNotToCreate() *resp.RestErr {
	return resp.NewBadRequestError(
		SttCouldNotPersist,
		"Failed to persist a new task",
	)
}

func ErrCouldNotListTasks() *resp.RestErr {
	return resp.NewBadRequestError(
		SttCouldNotListTasks,
		"Failed to get task list from database",
	)
}

func ErrCouldNotDeleteTask() *resp.RestErr {
	return resp.NewBadRequestError(
		SttCouldNotDeleteTask,
		"Failed to delete a task",
	)
}

func ErrCouldNotVerifyPosition() *resp.RestErr {
	return resp.NewBadRequestError(
		SttCouldNotVerifyPosition,
		"Failed to check if task position is available",
	)
}

func ErrCheckingIfTaskExists(err error) *resp.RestErr {
	return resp.NewErrInternal(
		"Failed to check if task exists",
		SttInternalRepositoryErr,
		err,
	)
}

func ErrCouldNotUpdateTask(err error) *resp.RestErr {
	return resp.NewErrInternal(
		"Failed to update task",
		SttCouldNotUpdateTask,
		err,
	)
}

func ErrRepeatedPosition() *resp.RestErr {
	return resp.NewBadRequestError(
		SttRepeatedPosition,
		"The task position is repeated, make sure that all positions are different",
	)
}

func ErrUnavailableTaskPosition() *resp.RestErr {
	return resp.NewBadRequestError(
		SttUnavailablePosition,
		"The task position is already in use",
	)
}
