package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"telegraph/db"
	"telegraph/model"

	"github.com/orcaman/concurrent-map"
)

type Resolver struct {
	client          *db.Client
	messageChannels cmap.ConcurrentMap
}

func NewResolver(client *db.Client) *Resolver {
	return &Resolver{
		client:          client,
		messageChannels: cmap.New(),
	}
}

func (r *Resolver) UpdateMessageChannels(message *model.Message) {
	if temp, ok := r.messageChannels.Get(message.ConversationID); ok {
		ms := temp.(cmap.ConcurrentMap)
		for _, ch := range ms.Items() {
			ch.(chan *model.Message) <- message
		}
	}
}
