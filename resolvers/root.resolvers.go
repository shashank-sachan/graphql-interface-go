package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"graphql-interface-go/data"
	"graphql-interface-go/generated"
	"graphql-interface-go/models"
)

// GetData is the resolver for the GetData field.
func (r *queryResolver) GetData(ctx context.Context) (*models.Root, error) {
	root := models.Root{}
	if err := json.Unmarshal([]byte(data.RootData), &root); err != nil {
		return nil, err
	}

	return &root, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
