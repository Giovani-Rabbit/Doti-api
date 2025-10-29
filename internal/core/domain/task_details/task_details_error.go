package detailsdomain

import resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"

const (
	SttCouldNotUpdateDescription    = "ERROR_UPDATING_DESCRIPTION"
	SttCouldNotUpdatePomodoroTarget = "ERROR_UPDATING_POMODORO_TARGET"
)

func ErrUpdatingDescription(err error) *resp.RestErr {
	return resp.NewErrInternal(
		"Unexpected error updating description",
		SttCouldNotUpdateDescription,
		err,
	)
}

func ErrUpdatingPomodoroTarget(err error) *resp.RestErr {
	return resp.NewErrInternal(
		"Unexpected error updating pomodoro target",
		SttCouldNotUpdatePomodoroTarget,
		err,
	)
}
