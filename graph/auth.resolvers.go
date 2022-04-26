package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"telegraph/model"
)

func (r *mutationResolver) SignUp(ctx context.Context, newUser model.NewUserInput) (*model.User, error) {
	return r.client.Auth.SignUp(newUser.Convert())
}

func (r *queryResolver) SignIn(ctx context.Context, phoneNumber string) (*model.User, error) {
	return r.client.Auth.SignIn(phoneNumber)
}

func (r *queryResolver) SendOtp(ctx context.Context, phoneNumber string) (string, error) {
	return r.client.Auth.SendOTP(phoneNumber)
}
