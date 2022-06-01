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

func authorize(client *db.Client, params ...interface{}) error {
	ok, err := client.Enforcer.Enforce(params...)
	if err != nil {
		return &gqlerror.Error{Message: fmt.Sprintf("An error occurred while trying to authorize the user: %s", err)}
	}

	if !ok {
		return &gqlerror.Error{Message: "Not Authorized"}
	}

	return nil
}

func Messenger(client *db.Client) func(context.Context, interface{}, graphql.Resolver, model.Action) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, action model.Action) (interface{}, error) {
		var err error
		token := middleware.GetToken(ctx)
		rc := graphql.GetResolverContext(ctx)

		err = client.Enforcer.LoadPolicy()
		if err != nil {
			return nil, &gqlerror.Error{Message: fmt.Sprintf("An error occurred while trying to load enforcer policy: %s", err)}
		}

		switch action {
		case model.ActionCreate:
			newMessage := rc.Args["new_message"].(model.NewMessageInput)
			for _, userID := range []string{token.ID, newMessage.Sender, newMessage.Recipient} {
				if err = authorize(client, userID, newMessage.ConversationID, model.FriendshipPolicyEnum.FRIEND); err != nil {
					break
				}
			}
		case model.ActionGetOne, model.ActionGetMany, model.ActionUpdate, model.ActionDelete:
			conversationID := rc.Args["conversation_id"].(string)
			err = authorize(client, token.ID, conversationID, model.FriendshipPolicyEnum.FRIEND)
		default:
			err = fmt.Errorf("%s is not a valid Action", action.String())
		}

		if err != nil {
			return nil, err
		}

		return next(ctx)
	}
}
