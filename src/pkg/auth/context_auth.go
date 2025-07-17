package auth

import (
	"context"

	authdomain "github.com/Giovani-Coelho/Doti-API/src/core/domain/auth"
)

type contextKey int

const authenticatedUserKey contextKey = 0

func GetAuthenticatedUserFromContext(ctx context.Context) (*authdomain.AuthClaims, error) {
	user, ok := ctx.Value(authenticatedUserKey).(*authdomain.AuthClaims)

	if !ok || user == nil {
		return nil, authdomain.ErrGetUserFromContext()
	}

	return user, nil
}

func SetUserInContext(
	ctx context.Context, user *authdomain.AuthClaims,
) context.Context {
	return context.WithValue(ctx, authenticatedUserKey, user)
}
