package directives

import (
	"context"
	"fmt"

	"telegraph/db"
	"telegraph/middleware"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Authorization(client *db.Client) func(context.Context, interface{}, graphql.Resolver) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		token := middleware.GetToken(ctx)
		rc := graphql.GetResolverContext(ctx)
		conversationID := rc.Args["conversation_id"]

		if token == nil {
			return nil, &gqlerror.Error{Message: "Access Denied"}
		}

		err := client.Enforcer.LoadPolicy()
		if err != nil {
			return nil, &gqlerror.Error{Message: fmt.Sprintf("An error occurred while trying to load enforcer policy: %s", err)}
		}

		ok, err := client.Enforcer.Enforce(fmt.Sprint(token.ID), fmt.Sprint(conversationID), "read")
		if err != nil {
			return nil, &gqlerror.Error{Message: fmt.Sprintf("An error occurred while trying to authorize the user: %s", err)}
		}

		if !ok {
			return nil, &gqlerror.Error{Message: "Access Denied"}
		}

		return next(ctx)
	}
}
