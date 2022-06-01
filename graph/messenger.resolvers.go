package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"telegraph/model"
)

func (r *mutationResolver) CreateMessage(ctx context.Context, newMessage model.NewMessageInput) (*model.Message, error) {
	return r.client.Messenger.Create(newMessage.Convert())
}

func (r *mutationResolver) ReadMessage(ctx context.Context, messageID string, conversationID string) (*model.Message, error) {
	return r.client.Messenger.Read(messageID)
}

func (r *mutationResolver) DeleteMessage(ctx context.Context, messageID string, conversationID string) (*model.Message, error) {
	return r.client.Messenger.Delete(messageID)
}

func (r *queryResolver) GetMessage(ctx context.Context, messageID string, conversationID string) (*model.Message, error) {
	return r.client.Messenger.Get(messageID)
}

func (r *queryResolver) GetMessages(ctx context.Context, conversationID string) ([]*model.Message, error) {
	return r.client.Messenger.GetMany(conversationID)
}
