package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"telegraph/graph/generated"
)

func (r *mutationResolver) Welcome(ctx context.Context) (string, error) {
	return "Welcome to Telegraph.", nil
}

func (r *queryResolver) Welcome(ctx context.Context) (string, error) {
	return "Welcome to Telegraph.", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
