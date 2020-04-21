package netmgrgql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nlepage/go-netmgr"
	"github.com/nlepage/go-netmgr/gql/generated"
	"github.com/nlepage/go-netmgr/gql/model"
)

func (r *mutationResolver) NetworkManager(ctx context.Context, input model.NetworkManagerInput) (netmgr.NetworkManager, error) {
	if input.WirelessEnabled != nil {
		if err := netmgr.SetWirelessEnabled(*input.WirelessEnabled); err != nil {
			return nil, err
		}
	}
	if input.WwanEnabled != nil {
		if err := netmgr.SetWwanEnabled(*input.WwanEnabled); err != nil {
			return nil, err
		}
	}
	nm, err := netmgr.System()
	if err != nil {
		return nil, err
	}
	return nm, nil
}

func (r *queryResolver) NetworkManager(ctx context.Context) (netmgr.NetworkManager, error) {
	nm, err := netmgr.System()
	if err != nil {
		return nil, err
	}
	return nm, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
