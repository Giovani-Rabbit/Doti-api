package user_request

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
