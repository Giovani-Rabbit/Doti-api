package auth

import (
	"os"
	"strings"
	"time"

	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"github.com/golang-jwt/jwt"
)

const (
	JWT_TOKEN_KEY = "JWT_TOKEN_KEY"
)

func GenerateToken(user userdomain.IUserDomain) (string, error) {
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

func VerifyToken(tokenValue string) (userdomain.IUserDomain, *rest_err.RestErr) {
	secret := os.Getenv(JWT_TOKEN_KEY)

	token, err := jwt.Parse(
		removeBearerPrefix(tokenValue),
		func(t *jwt.Token) (any, error) {
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

	return userdomain.New(
		claims["id"].(string),
		claims["name"].(string),
		claims["email"].(string),
		"",
		time.Now(),
		time.Now(),
	), nil
}

func removeBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix("Bearer ", token)
	}

	return token
}
