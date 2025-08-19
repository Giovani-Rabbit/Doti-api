package resp

const (
	SttInvalidUUID = "INVALID_UUID"
)

func NewInvalidUUID() error {
	return NewBadRequestError(
		SttInvalidUUID,
		"Invalid ID format",
	)
}
