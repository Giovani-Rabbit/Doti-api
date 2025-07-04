package userdomain

import "time"

type IUserDomain interface {
	EncryptPassword() error

	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetIsAdmin() bool
	GetCreateAt() time.Time
	GetUpdatedAt() time.Time

	setPassword(password string) // private

	IsValid() bool
	IsValidEmail() bool
}

func NewUserDomain(
	id string,
	name string,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) IUserDomain {
	return &userDomain{
		id:        id,
		name:      name,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func NewCreateUserDomain(
	name string, email string, password string,
) IUserDomain {
	return &userDomain{
		name:     name,
		email:    email,
		password: password,
	}
}

func NewSignInUserDomain(
	email string, password string,
) IUserDomain {
	return &userDomain{
		email:    email,
		password: password,
	}
}
