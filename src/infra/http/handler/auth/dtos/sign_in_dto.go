package authdto

import "github.com/gofrs/uuid"

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponseDTO struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}
