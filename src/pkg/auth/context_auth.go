package auth

import (
	"context"
	"errors"

	authdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/auth"
)

type contextKey int

const authenticatedUserKey contextKey = 0

func GetAuthenticatedUserFromContext(
	ctx context.Context,
) (*authdomain.AuthClaims, bool) {
	user, ok := ctx.Value(authenticatedUserKey).(*authdomain.AuthClaims)
	return user, ok
}

func GetUserFromContext(ctx context.Context) (*authdomain.AuthClaims, error) {
	user, ok := ctx.Value(authenticatedUserKey).(*authdomain.AuthClaims)

	if !ok || user == nil {
		return nil, errors.New("user not found in context")
	}

	return user, nil
}

func SetUserInContext(
	ctx context.Context, user *authdomain.AuthClaims,
) context.Context {
	return context.WithValue(ctx, authenticatedUserKey, user)
}
