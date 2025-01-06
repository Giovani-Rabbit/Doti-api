package auth

import (
	"os"
	"time"

	"github.com/Giovani-Coelho/Doti-API/src/infra/database/db/sqlc"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_TOKEN_KEY = "JWT_TOKEN_KEY"
)

func GenerateToken(
	user sqlc.User,
) (string, error) {
	secret := os.Getenv(JWT_TOKEN_KEY)

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerError(
			"Error trying to generate jwt token",
		)
	}

	return tokenString, nil
}
