package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"telegraph/db"
)

type Resolver struct{
	client *db.Client
}

func NewResolver(client *db.Client) *Resolver {
	return &Resolver {
		client: client,
	}
}
