package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gql/graph/generated"
	"gql/graph/model"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	link := &model.Link{Name: input.Name, URL: input.URL}
	r.Db.Write(link)
	return link, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	links := *r.Db.ReadAll()
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
