package taskdomain

import resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"

const (
	SttInvalidFields   = "INVALID_TASK_FIELDS"
	SttCouldNotPersist = "COULD_NOT_PERSIST_TASK"
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
