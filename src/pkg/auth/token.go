package auth

import (
	"os"
	"strings"
	"time"

	"github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	"github.com/Giovani-Coelho/Doti-API/src/infra/persistence/db/sqlc"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var (
	JWT_TOKEN_KEY = "JWT_TOKEN_KEY"
)

func GenerateToken(user user.IUserDomain) (string, error) {
	secret := os.Getenv(JWT_TOKEN_KEY)

	claims := jwt.MapClaims{
		"id":    user.GetID(),
		"email": user.GetEmail(),
		"name":  user.GetName(),
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

func VerifyToken(tokenValue string) (*sqlc.User, error) {
	secret := os.Getenv(JWT_TOKEN_KEY)

	token, err := jwt.Parse(
		RemoveBearerPrefix(tokenValue),
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(secret), nil
			}

			return nil, rest_err.NewBadRequestError("INVALID_TOKEN", "invalid token")
		},
	)

	if err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedRequestError("invalid token")
	}

	return &sqlc.User{
		ID:    claims["id"].(uuid.UUID),
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix("Bearer ", token)
	}

	return token
}
