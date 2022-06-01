package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"telegraph/model"

	cmap "github.com/orcaman/concurrent-map"
)

func (r *mutationResolver) CreateMessage(ctx context.Context, newMessage model.NewMessageInput) (*model.Message, error) {
	message, err := r.client.Messenger.Create(newMessage.Convert())
	if err == nil {
		go r.UpdateMessageChannels(message)
	}

	return message, err
}

func (r *mutationResolver) ReadMessage(ctx context.Context, messageID string, conversationID string) (*model.Message, error) {
	message, err := r.client.Messenger.Read(messageID)
	if err == nil {
		go r.UpdateMessageChannels(message)
	}

	return message, err
}

func (r *mutationResolver) DeleteMessage(ctx context.Context, messageID string, conversationID string) (*model.Message, error) {
	message, err := r.client.Messenger.Delete(messageID)
	if err == nil {
		go r.UpdateMessageChannels(message)
	}

	return message, err
}

func (r *queryResolver) GetMessage(ctx context.Context, messageID string, conversationID string) (*model.Message, error) {
	return r.client.Messenger.Get(messageID)
}

func (r *queryResolver) GetMessages(ctx context.Context, conversationID string) ([]*model.Message, error) {
	return r.client.Messenger.GetMany(conversationID)
}

func (r *subscriptionResolver) MessagesSubscription(ctx context.Context, conversationID string, userID string) (<-chan *model.Message, error) {
	var channels cmap.ConcurrentMap
	ch := make(chan *model.Message, 1)

	if temp, ok := r.messageChannels.Get(conversationID); !ok {
		channels = cmap.New()
	} else {
		channels = temp.(cmap.ConcurrentMap)
	}

	channels.Set(userID, ch)
	r.messageChannels.Set(conversationID, channels)

	go func() {
		<-ctx.Done()
		close(ch)
		channels.Remove(userID)
		if channels.IsEmpty() {
			r.messageChannels.Remove(conversationID)
		}
	}()

	return ch, nil
}
