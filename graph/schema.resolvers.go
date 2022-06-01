package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"telegraph/graph/generated"
)

func (r *mutationResolver) Welcome(ctx context.Context) (string, error) {
	panic(fmt.Errorf("Welcome!"))
}

func (r *queryResolver) Welcome(ctx context.Context) (string, error) {
	panic(fmt.Errorf("Welcome!"))
}

func (r *subscriptionResolver) Welcome(ctx context.Context) (<-chan string, error) {
	panic(fmt.Errorf("Welcome!"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
