package entity

import "github.com/gofrs/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
}
