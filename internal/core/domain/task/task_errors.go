package taskdomain

import resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"

const (
	SttInvalidFields          = "INVALID_TASK_FIELDS"
	SttCouldNotPersist        = "COULD_NOT_PERSIST_TASK"
	SttCouldNotListTasks      = "COULD_NOT_LIST_TASKS"
	SttCouldNotVerifyPosition = "COULD_NOT_VERIFY_POSITION"
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

func ErrUnavailableTaskPosition() *resp.RestErr {
	return resp.NewBadRequestError(
		SttUnavailablePosition,
		"The task position is already in use",
	)
}
