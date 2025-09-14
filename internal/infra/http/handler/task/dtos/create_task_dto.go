package taskdto

type CreateTaskDTO struct {
	ModuleId int    `json:"module_id"`
	TaskName string `json:"task_name"`
	Position int    `json:"position"`
}
