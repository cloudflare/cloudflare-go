package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrMissingNetwork      = errors.New("missing required network parameter")
	ErrInvalidNetworkValue = errors.New("invalid IP parameter. Cannot use CIDR ranges for this endpoint.")
)

// TunnelRoute is the full record for a route.
type TunnelRoute struct {
	Network    string     `json:"network"`
	TunnelID   string     `json:"tunnel_id"`
	TunnelName string     `json:"tunnel_name"`
	Comment    string     `json:"comment"`
	CreatedAt  *time.Time `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type TunnelRoutesListParams struct {
	AccountID       string     `url:"-"`
	TunnelID        string     `url:"tunnel_id,omitempty"`
	Comment         string     `url:"comment,omitempty"`
	IsDeleted       *bool      `url:"is_deleted,omitempty"`
	NetworkSubset   string     `url:"network_subset,omitempty"`
	NetworkSuperset string     `url:"network_superset,omitempty"`
	ExistedAt       *time.Time `url:"existed_at,omitempty"`
	PaginationOptions
}

type TunnelRoutesCreateParams struct {
	AccountID string `json:"-"`
	Network   string `json:"-"`
	TunnelID  string `json:"tunnel_id"`
	Comment   string `json:"comment,omitempty"`
}

type TunnelRoutesUpdateParams struct {
	AccountID string `json:"-"`
	Network   string `json:"network"`
	TunnelID  string `json:"tunnel_id"`
	Comment   string `json:"comment,omitempty"`
}

type TunnelRoutesForIPParams struct {
	AccountID string `json:"-"`
	Network   string `json:"-"`
}

type TunnelRoutesDeleteParams struct {
	AccountID string `json:"-"`
	Network   string `json:"-"`
}

// tunnelRouteListResponse is the API response for listing tunnel routes.
type tunnelRouteListResponse struct {
	Response
	Result []TunnelRoute `json:"result"`
}

type tunnelRouteResponse struct {
	Response
	Result TunnelRoute `json:"result"`
}

// ListTunnelRoutes lists all defined routes for tunnels in the account.
//
// See: https://api.cloudflare.com/#tunnel-route-list-tunnel-routes
func (api *API) ListTunnelRoutes(ctx context.Context, params TunnelRoutesListParams) ([]TunnelRoute, error) {
	if params.AccountID == "" {
		return []TunnelRoute{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/teamnet/routes", AccountRouteRoot, params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []TunnelRoute{}, err
	}

	var resp tunnelRouteListResponse
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return []TunnelRoute{}, errors.Wrap(err, errUnmarshalError)
	}

	return resp.Result, nil
}

// GetTunnelRouteForIP finds the Tunnel Route that encompasses the given IP.
//
// See: https://api.cloudflare.com/#tunnel-route-get-tunnel-route-by-ip
func (api *API) GetTunnelRouteForIP(ctx context.Context, params TunnelRoutesForIPParams) (TunnelRoute, error) {
	if params.AccountID == "" {
		return TunnelRoute{}, ErrMissingAccountID
	}

	if params.Network == "" {
		return TunnelRoute{}, ErrMissingNetwork
	}

	if strings.Contains(params.Network, "/") {
		return TunnelRoute{}, ErrInvalidNetworkValue
	}

	uri := fmt.Sprintf("/%s/%s/teamnet/routes/ip/%s", AccountRouteRoot, params.AccountID, params.Network)

	responseBody, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TunnelRoute{}, err
	}

	var routeResponse tunnelRouteResponse
	err = json.Unmarshal(responseBody, &routeResponse)
	if err != nil {
		return TunnelRoute{}, errors.Wrap(err, errUnmarshalError)
	}

	return routeResponse.Result, nil
}

// CreateTunnelRoute add a new route to the account routing table for the given
// tunnel.
//
// See: https://api.cloudflare.com/#tunnel-route-create-route
func (api *API) CreateTunnelRoute(ctx context.Context, params TunnelRoutesCreateParams) (TunnelRoute, error) {
	if params.AccountID == "" {
		return TunnelRoute{}, ErrMissingAccountID
	}

	if params.Network == "" {
		return TunnelRoute{}, ErrMissingNetwork
	}

	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, params.AccountID, url.PathEscape(params.Network))

	responseBody, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return TunnelRoute{}, err
	}

	var routeResponse tunnelRouteResponse
	err = json.Unmarshal(responseBody, &routeResponse)
	if err != nil {
		return TunnelRoute{}, errors.Wrap(err, errUnmarshalError)
	}

	return routeResponse.Result, nil
}

// DeleteTunnelRoute delete an existing route from the account routing table.
//
// See: https://api.cloudflare.com/#tunnel-route-delete-route
func (api *API) DeleteTunnelRoute(ctx context.Context, params TunnelRoutesDeleteParams) error {
	if params.AccountID == "" {
		return ErrMissingAccountID
	}

	if params.Network == "" {
		return ErrMissingNetwork
	}

	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, params.AccountID, url.PathEscape(params.Network))

	responseBody, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	var routeResponse tunnelRouteResponse
	err = json.Unmarshal(responseBody, &routeResponse)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}

	return nil
}

// UpdateTunnelRoute updates an existing route in the account routing table for
// the given tunnel.
//
// See: https://api.cloudflare.com/#tunnel-route-update-route
func (api *API) UpdateTunnelRoute(ctx context.Context, params TunnelRoutesUpdateParams) (TunnelRoute, error) {
	if params.AccountID == "" {
		return TunnelRoute{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, params.AccountID, url.PathEscape(params.Network))

	responseBody, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return TunnelRoute{}, err
	}

	var routeResponse tunnelRouteResponse
	err = json.Unmarshal(responseBody, &routeResponse)
	if err != nil {
		return TunnelRoute{}, errors.Wrap(err, errUnmarshalError)
	}

	return routeResponse.Result, nil
}
