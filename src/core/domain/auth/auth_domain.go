package authdomain

import "github.com/golang-jwt/jwt/v5"

type AuthClaims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
