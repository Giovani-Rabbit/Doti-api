package detailsdomain

import resp "github.com/Giovani-Coelho/Doti-API/internal/infra/http/responder"

const (
	SttInvalidDescription = "INVALID_DESCRIPTION_DETAILS"
	SttCouldNotUpdateDescription
)

func ErrInvalidDescription() *resp.RestErr {
	return resp.NewBadRequestError(
		SttInvalidDescription,
		"Invalid description",
	)
}

func ErrUpdatingDescription(err error) *resp.RestErr {
	return resp.NewErrInternal(
		"Unexpected error updating description",
		SttCouldNotUpdateDescription,
		err,
	)
}
