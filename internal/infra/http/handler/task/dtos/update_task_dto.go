package taskdto

type TaskPositionParams struct {
	TaskId   int32 `json:"task_id"`
	Position int32 `json:"position"`
}

type UpdatePositionDTO struct {
	Tasks []TaskPositionParams
}
