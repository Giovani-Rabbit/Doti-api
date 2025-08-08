package userdomain

import (
	"net/mail"
	"strings"
	"time"
)

type userDomain struct {
	id        string
	email     string
	name      string
	password  string
	isAdmin   bool
	createdAt time.Time
	updatedAt time.Time
}

func (u *userDomain) GetID() string               { return u.id }
func (u *userDomain) GetName() string             { return u.name }
func (u *userDomain) GetEmail() string            { return u.email }
func (u *userDomain) GetPassword() string         { return u.password }
func (u *userDomain) GetIsAdmin() bool            { return u.isAdmin }
func (u *userDomain) GetCreateAt() time.Time      { return u.createdAt }
func (u *userDomain) GetUpdatedAt() time.Time     { return u.updatedAt }
func (u *userDomain) setPassword(password string) { u.password = password }

func (u *userDomain) IsValid() bool {
	return strings.TrimSpace(u.name) != "" &&
		strings.TrimSpace(u.email) != "" &&
		strings.TrimSpace(u.password) != ""
}

func (ud *userDomain) IsValidEmail() bool {
	_, err := mail.ParseAddress(ud.email)
	return err == nil
}
