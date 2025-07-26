package authdto

type SignInDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponseDTO struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	ExpiresIn   int64  `json:"expires_in"`
	AccessToken string `json:"accessToken"`
}
