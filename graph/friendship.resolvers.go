package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"telegraph/model"
)

func (r *mutationResolver) CreateFriendship(ctx context.Context, newFriendship model.NewFriendshipInput) (*model.Friendship, error) {
	return r.client.Friendship.Create(newFriendship.Convert())
}

func (r *mutationResolver) AcceptFriendship(ctx context.Context, friendshipID string) (*model.Friendship, error) {
	return r.client.Friendship.Accept(friendshipID)
}

func (r *mutationResolver) DeleteFriendship(ctx context.Context, friendshipID string) (*model.Friendship, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetFriendship(ctx context.Context, friendshipID string) (*model.Friendship, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetFriendships(ctx context.Context, userID string) ([]*model.Friendship, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPendingFriendships(ctx context.Context, userID string) ([]*model.Friendship, error) {
	panic(fmt.Errorf("not implemented"))
}
