package user_response

import (
	"time"

	userDomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
)

type UserCreatedResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserCreatedResponse(user userDomain.IUserDomain) UserCreatedResponse {
	return UserCreatedResponse{
		ID:        user.GetID(),
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreateAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}
}
