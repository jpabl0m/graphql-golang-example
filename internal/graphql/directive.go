package graphql

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/stobita/graphql-golang-example/internal/usecase"
)

func RequireAuthDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	userID := ctx.Value(usecase.UserIDContextKey)
	if userID == nil {
		return nil, errors.New("Access denied")
	}
	return next(ctx)
}
