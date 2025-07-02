package user

import "time"

type IUserDomain interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetIsAdmin() bool
	GetCreateAt() time.Time
	GetUpdatedAt() time.Time

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
		password: encryptPassword(password),
	}
}

func NewSignInUserDomain(
	email string, password string,
) IUserDomain {
	return &userDomain{
		email:    email,
		password: encryptPassword(password),
	}
}
