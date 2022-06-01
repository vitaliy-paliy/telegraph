package directives

import (
	"context"
	"fmt"

	"telegraph/db"
	"telegraph/middleware"
	"telegraph/model"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Friendship(client *db.Client) func(context.Context, interface{}, graphql.Resolver, model.Action) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, action model.Action) (interface{}, error) {
		var err error
		token := middleware.GetToken(ctx)
		rc := graphql.GetResolverContext(ctx)

		switch action {
		case model.ActionCreate:
			newFriendship := rc.Args["new_friendship"].(model.NewFriendshipInput)
			if token.ID != newFriendship.Sender {
				err = &gqlerror.Error{Message: "Not Authorized"}
			}
		case model.ActionGetOne, model.ActionUpdate, model.ActionDelete:
			friendshipID := rc.Args["friendship_id"]
			err = client.Enforcer.LoadPolicy()
			if err != nil {
				err = &gqlerror.Error{Message: fmt.Sprintf("An error occurred while trying to load enforcer policy: %s", err)}
				break
			}

			for _, policy := range []string{model.FriendshipPolicyEnum.FRIEND, model.FriendshipPolicyEnum.SENDER, model.FriendshipPolicyEnum.RECIPIENT} {
				ok, err := client.Enforcer.Enforce(token.ID, friendshipID, policy)
				if err != nil {
					err = &gqlerror.Error{Message: fmt.Sprintf("An error occurred while trying to authorize the user: %s", err)}
					break
				}

				if ok {
					return next(ctx)
				}
			}
			if err == nil {
				err = &gqlerror.Error{Message: "Not Authorized"}
			}
		case model.ActionGetMany:
			userID := rc.Args["user_id"]
			if userID != token.ID {
				err = &gqlerror.Error{Message: "Not Authorized"}
			}
		default:
			err = fmt.Errorf("%s is not a valid Action", action.String())
		}

		if err != nil {
			return nil, err
		}

		return next(ctx)
	}
}
