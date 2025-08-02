package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	authdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/auth"
	userdomain "github.com/Giovani-Coelho/Doti-API/internal/core/domain/user"
	httperr "github.com/Giovani-Coelho/Doti-API/internal/pkg/handlers/http"
	"github.com/golang-jwt/jwt/v5"
)

const JWT_TOKEN_KEY = "JWT_TOKEN_KEY"

var secretKey = os.Getenv(JWT_TOKEN_KEY)

func GenerateToken(user userdomain.User) (string, error) {
	claims := authdomain.AuthClaims{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenValue string) (*authdomain.AuthClaims, *httperr.RestErr) {
	claims := &authdomain.AuthClaims{}

	token, err := jwt.ParseWithClaims(tokenValue, claims,
		func(t *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil || !token.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, httperr.NewUnauthorizedRequestError("Expired Token")
		}

		return nil, httperr.NewUnauthorizedRequestError("Unauthorized")
	}

	return claims, nil
}

func GetAuthenticatedUser(r *http.Request) (*authdomain.AuthClaims, error) {
	cookieToken, err := GetTokenFromCookie(r)

	if err == nil {
		return cookieToken, nil
	}

	headerToken, err := GetTokenFromHeader(r)

	if err == nil {
		return headerToken, nil
	}

	return nil, errors.New("no token found")
}

func GetTokenFromHeader(r *http.Request) (*authdomain.AuthClaims, error) {
	authHeader := r.Header.Get("Authorization")

	user, err := VerifyToken(authHeader)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetTokenFromCookie(r *http.Request) (*authdomain.AuthClaims, error) {
	cookie, err := r.Cookie("access-token")

	if err != nil {
		return nil, err
	}

	user, invalid := VerifyToken(cookie.Value)

	if invalid != nil {
		return nil, err
	}

	return user, err
}
