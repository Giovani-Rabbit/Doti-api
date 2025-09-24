package taskdomain

import resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"

const (
	SttInvalidFields          = "INVALID_TASK_FIELDS"
	SttCouldNotPersist        = "COULD_NOT_PERSIST_TASK"
	SttCouldNotListTasks      = "COULD_NOT_LIST_TASKS"
	SttCouldNotVerifyPosition = "COULD_NOT_VERIFY_POSITION"
	SttCouldNotUpdateTask     = "COULD_NOT_UPDATE_TASK"
	SttInvalidTaskQuantity    = "INVALID_TASK_QUANTITY"
	SttInvalidPosition        = "INVALID_POSITION"
	SttUnavailablePosition    = "UNAVAILABLE_POSITION"
)

func ErrInvalidFields() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidFields,
		"The fields required to create a task are invalid",
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

func ErrCouldNotVerifyPosition() *resp.RestErr {
	return resp.NewBadRequestError(
		SttCouldNotVerifyPosition,
		"Failed to check if task position is available",
	)
}

func ErrCouldNotUpdateTask(err error) *resp.RestErr {
	return resp.NewErrInternal(
		"Failed to update task",
		SttCouldNotUpdateTask,
		err,
	)
}

func ErrInvalidPosition() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidPosition,
		"Task positions should be different",
	)
}

func ErrInvalidTaskQuantity() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidTaskQuantity,
		"The number of tasks should be 2",
	)
}

func ErrUnavailableTaskPosition() *resp.RestErr {
	return resp.NewBadRequestError(
		SttUnavailablePosition,
		"The task position is already in use",
	)
}
