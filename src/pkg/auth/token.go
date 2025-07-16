package auth

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Giovani-Coelho/Doti-API/config/logger"
	authdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/auth"
	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"github.com/golang-jwt/jwt"
)

const (
	JWT_TOKEN_KEY = "JWT_TOKEN_KEY"
)

var secretKey = os.Getenv(JWT_TOKEN_KEY)

func GenerateToken(user userdomain.IUserDomain) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.GetID(),
		"email": user.GetEmail(),
		"name":  user.GetName(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", rest_err.NewInternalServerError(
			"Error trying to generate jwt token",
		)
	}

	return tokenString, nil
}

func VerifyToken(tokenValue string) (*authdomain.AuthClaims, *rest_err.RestErr) {
	token, err := jwt.Parse(
		removeBearerPrefix(tokenValue),
		func(t *jwt.Token) (any, error) {
			// validate the method used by jwt
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Error("unexpected signing method", nil)
				return nil, rest_err.NewBadRequestError(
					"INVALID_JWT_METHOD", "Unexpected signing jwt method",
				)
			}

			return []byte(secretKey), nil
		},
	)

	if err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("Unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
	}

	if err := claims.Valid(); err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("Token expired")
	}

	return &authdomain.AuthClaims{
		ID:    claims["id"].(string),
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
	}, nil
}

func GetAuthenticatedUser(r *http.Request) (*authdomain.AuthClaims, error) {
	authHeader := r.Header.Get("Authorization")

	user, err := VerifyToken(authHeader)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func removeBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix("Bearer ", token)
	}

	return token
}
