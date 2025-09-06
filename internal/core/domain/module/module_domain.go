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

func (md *moduleDomain) GetID() int32 {
	return md.id

}
func (md *moduleDomain) GetUserId() string {
	return md.userId
}

func (md *moduleDomain) GetName() string {
	return md.name
}

func (md *moduleDomain) GetIsOpen() bool {
	return md.isOpen
}

func (md *moduleDomain) GetIcon() string {
	return md.icon
}

func (md *moduleDomain) GetCreateAt() time.Time {
	return md.createdAt
}

func (md *moduleDomain) GetUpdatedAt() time.Time {
	return md.updatedAt
}

func (md *moduleDomain) IsValid() bool {
	return strings.TrimSpace(md.icon) != "" && strings.TrimSpace(md.name) != ""
}
