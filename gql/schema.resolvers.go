package netmgrgql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/nlepage/go-netmgr"
	"github.com/nlepage/go-netmgr/gql/generated"
	"github.com/nlepage/go-netmgr/gql/model"
)

func (r *deviceResolver) ID(ctx context.Context, obj netmgr.Device) (string, error) {
	path := string(obj.Path())
	i := strings.LastIndex(path, "/")
	if i == -1 {
		return "", fmt.Errorf("Path has no slashes: %#v", path)
	}
	return path[i+1:], nil
}

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
	if input.ConnectivityCheckEnabled != nil {
		if err := netmgr.SetConnectivityCheckEnabled(*input.ConnectivityCheckEnabled); err != nil {
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

// Device returns generated.DeviceResolver implementation.
func (r *Resolver) Device() generated.DeviceResolver { return &deviceResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type deviceResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
