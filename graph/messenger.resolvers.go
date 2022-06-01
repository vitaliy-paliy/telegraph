package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"telegraph/model"
)

func (r *queryResolver) GetMessages(ctx context.Context, conversationID string) ([]*model.Message, error) {
	return []*model.Message{&model.Message{Text: "Success"}}, nil
}
