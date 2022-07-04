package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"errors"
)

var ErrMissingVnetName = errors.New("required missing virtual network name")

// TunnelVirtualNetwork is segregation of Tunnel IP Routes via Virtualized
// Networks to handle overlapping private IPs in your origins.
type TunnelVirtualNetwork struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	IsDefaultNetwork bool       `json:"is_default_network"`
	Comment          string     `json:"comment"`
	CreatedAt        *time.Time `json:"created_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}

type TunnelVirtualNetworksListParams struct {
	AccountID string `url:"-"`
	ID        string `url:"id,omitempty"`
	Name      string `url:"name,omitempty"`
	IsDefault *bool  `url:"is_default,omitempty"`
	IsDeleted *bool  `url:"is_deleted,omitempty"`

	PaginationOptions
}

type TunnelVirtualNetworkCreateParams struct {
	AccountID string `json:"-"`
	Name      string `json:"name"`
	Comment   string `json:"comment"`
	IsDefault bool   `json:"is_default"`
}

type TunnelVirtualNetworkUpdateParams struct {
	AccountID        string `json:"-"`
	VnetID           string `json:"-"`
	Name             string `json:"name,omitempty"`
	Comment          string `json:"comment,omitempty"`
	IsDefaultNetwork *bool  `json:"is_default_network,omitempty"`
}

type TunnelVirtualNetworkDeleteParams struct {
	AccountID string `json:"-"`
	VnetID    string `json:"-"`
}

// tunnelRouteListResponse is the API response for listing tunnel virtual
// networks.
type tunnelVirtualNetworkListResponse struct {
	Response
	Result []TunnelVirtualNetwork `json:"result"`
}

type tunnelVirtualNetworkResponse struct {
	Response
	Result TunnelVirtualNetwork `json:"result"`
}

// ListTunnelVirtualNetworks lists all defined virtual networks for tunnels in
// the account.
//
// API reference: https://api.cloudflare.com/#tunnel-virtual-network-list-virtual-networks
func (api *API) ListTunnelVirtualNetworks(ctx context.Context, params TunnelVirtualNetworksListParams) ([]TunnelVirtualNetwork, error) {
	if params.AccountID == "" {
		return []TunnelVirtualNetwork{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/teamnet/virtual_networks", AccountRouteRoot, params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, params)
	if err != nil {
		return []TunnelVirtualNetwork{}, err
	}

	var resp tunnelVirtualNetworkListResponse
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return []TunnelVirtualNetwork{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return resp.Result, nil
}

// CreateTunnelVirtualNetwork adds a new virtual network to the account.
//
// API reference: https://api.cloudflare.com/#tunnel-virtual-network-create-virtual-network
func (api *API) CreateTunnelVirtualNetwork(ctx context.Context, params TunnelVirtualNetworkCreateParams) (TunnelVirtualNetwork, error) {
	if params.AccountID == "" {
		return TunnelVirtualNetwork{}, ErrMissingAccountID
	}

	if params.Name == "" {
		return TunnelVirtualNetwork{}, ErrMissingVnetName
	}

	uri := fmt.Sprintf("/%s/%s/teamnet/virtual_networks", AccountRouteRoot, params.AccountID)

	responseBody, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return TunnelVirtualNetwork{}, err
	}

	var resp tunnelVirtualNetworkResponse
	err = json.Unmarshal(responseBody, &resp)
	if err != nil {
		return TunnelVirtualNetwork{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return resp.Result, nil
}

// DeleteTunnelVirtualNetwork deletes an existing virtual network from the
// account.
//
// API reference: https://api.cloudflare.com/#tunnel-virtual-network-delete-virtual-network
func (api *API) DeleteTunnelVirtualNetwork(ctx context.Context, params TunnelVirtualNetworkDeleteParams) error {
	if params.AccountID == "" {
		return ErrMissingAccountID
	}

	uri := fmt.Sprintf("/%s/%s/teamnet/virtual_networks/%s", AccountRouteRoot, params.AccountID, params.VnetID)

	responseBody, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	var resp tunnelVirtualNetworkResponse
	err = json.Unmarshal(responseBody, &resp)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

// UpdateTunnelRoute updates an existing virtual network in the account.
//
// API reference: https://api.cloudflare.com/#tunnel-virtual-network-update-virtual-network
func (api *API) UpdateTunnelVirtualNetwork(ctx context.Context, params TunnelVirtualNetworkUpdateParams) (TunnelVirtualNetwork, error) {
	if params.AccountID == "" {
		return TunnelVirtualNetwork{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/%s/%s/teamnet/virtual_networks/%s", AccountRouteRoot, params.AccountID, params.VnetID)

	responseBody, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return TunnelVirtualNetwork{}, err
	}

	var resp tunnelVirtualNetworkResponse
	err = json.Unmarshal(responseBody, &resp)
	if err != nil {
		return TunnelVirtualNetwork{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return resp.Result, nil
}
