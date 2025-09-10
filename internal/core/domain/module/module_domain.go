package moduledomain

import (
	"strings"
	"time"
)

type moduleDomain struct {
	id        int32
	userId    string
	name      string
	isOpen    bool
	icon      string
	createdAt time.Time
	updatedAt time.Time
}

func (md *moduleDomain) ID() int32            { return md.id }
func (md *moduleDomain) UserId() string       { return md.userId }
func (md *moduleDomain) Name() string         { return md.name }
func (md *moduleDomain) IsOpen() bool         { return md.isOpen }
func (md *moduleDomain) Icon() string         { return md.icon }
func (md *moduleDomain) CreateAt() time.Time  { return md.createdAt }
func (md *moduleDomain) UpdatedAt() time.Time { return md.updatedAt }

func (md *moduleDomain) IsValid() bool {
	return strings.TrimSpace(md.icon) != "" &&
		strings.TrimSpace(md.name) != "" &&
		strings.TrimSpace(md.userId) != ""
}
