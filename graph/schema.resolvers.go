package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/brandaogabriel7/studio-sol-back-end-test/graph/model"
)

// Verify is the resolver for the verify field.
func (r *queryResolver) Verify(ctx context.Context, password string, rules []*model.Rule) (*model.Verify, error) {
	verifyResponse := r.PasswordValidationService.Validate(password, rules)
	return &verifyResponse, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
