package netmgrgql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nlepage/go-netmgr"
	"github.com/nlepage/go-netmgr/gql/generated"
	"github.com/nlepage/go-netmgr/gql/model"
)

func (r *mutationResolver) NetworkManager(ctx context.Context, input model.NetworkManagerInput) (*model.NetworkManager, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) NetworkManager(ctx context.Context) (*model.NetworkManager, error) {
	nm, err := netmgr.System()
	if err != nil {
		return nil, err
	}
	return &model.NetworkManager{nm}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
