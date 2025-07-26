package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	authdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/auth"
	userdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/user"
	authdto "github.com/Giovani-Coelho/Doti-API/src/infra/http/handler/auth/dtos"
	rest_err "github.com/Giovani-Coelho/Doti-API/src/pkg/handlers/http"
	"github.com/golang-jwt/jwt/v5"
)

const JWT_TOKEN_KEY = "JWT_TOKEN_KEY"

var secretKey = os.Getenv(JWT_TOKEN_KEY)

func GenerateToken(user userdomain.IUserDomain) (authdto.AuthTokenDTO, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := authdomain.AuthClaims{
		ID:    user.GetID(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))

	if err != nil {
		return authdto.AuthTokenDTO{},
			rest_err.NewUnauthorizedRequestError("Error generating token")
	}

	return authdto.AuthTokenDTO{
		AccessToken: token,
		ExpiresIn:   int64(time.Until(expirationTime).Seconds()),
	}, nil
}

func VerifyToken(tokenValue string) (*authdomain.AuthClaims, *rest_err.RestErr) {
	claims := &authdomain.AuthClaims{}

	token, err := jwt.ParseWithClaims(tokenValue, claims,
		func(t *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil || !token.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, rest_err.NewUnauthorizedRequestError("Expired Token")
		}

		return nil, rest_err.NewUnauthorizedRequestError("Unauthorized")
	}

	return claims, nil
}

func GetAuthenticatedUser(r *http.Request) (*authdomain.AuthClaims, error) {
	authHeader := r.Header.Get("Authorization")

	user, err := VerifyToken(authHeader)

	if err != nil {
		return nil, err
	}

	return user, nil
}
