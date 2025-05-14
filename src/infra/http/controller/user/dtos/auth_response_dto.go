package userDTO

import "github.com/gofrs/uuid"

type AuthResponseDTO struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}
