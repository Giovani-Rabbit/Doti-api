package taskdomain

import (
	"strings"
	"time"
)

type taskDomain struct {
	id          int32
	moduleId    int32
	name        string
	isCompleted bool
	position    int32
	createdAt   time.Time
	updatedAt   time.Time
}

func (t *taskDomain) ID() int32            { return t.id }
func (t *taskDomain) ModuleId() int32      { return t.moduleId }
func (t *taskDomain) Name() string         { return t.name }
func (t *taskDomain) IsCompleted() bool    { return t.isCompleted }
func (t *taskDomain) Position() int32      { return t.position }
func (t *taskDomain) CreatedAt() time.Time { return t.createdAt }
func (t *taskDomain) UpdatedAt() time.Time { return t.updatedAt }

func (t *taskDomain) IsValidToCreate() bool {
	return t.moduleId >= 0 &&
		strings.TrimSpace(t.name) != "" &&
		t.position >= 0
}
