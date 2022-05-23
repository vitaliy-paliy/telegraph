package directives

import (
	"context"
	"fmt"

	"telegraph/middleware"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	token := middleware.GetToken(ctx)
	if token == nil {
		return nil, &gqlerror.Error{Message: "Access Denied"}
	}
	fmt.Println(token)

	return next(ctx)
}
