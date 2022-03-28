package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// TunnelRoute is the full record for a route.
type TunnelRoute struct {
	Network    string     `json:"network"`
	TunnelId   string     `json:"tunnel_id"`
	TunnelName string     `json:"tunnel_name"`
	Comment    string     `json:"comment"`
	CreatedAt  *time.Time `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

type TunnelRoutesListParams struct {
	PaginationOptions
	TunnelId        string     `json:"tunnel_id,omitempty"`
	Comment         string     `json:"comment,omitempty"`
	IsDeleted       *bool      `json:"is_deleted,omitempty"`
	NetworkSubset   string     `json:"network_subset,omitempty"`
	NetworkSuperset string     `json:"network_superset,omitempty"`
	ExistedAt       *time.Time `json:"existed_at,omitempty"`
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

// TunnelRoutes lists all defined routes for tunnels in the account.
//
// See: https://api.cloudflare.com/#tunnel-route-list-tunnel-routes
func (api *API) TunnelRoutes(ctx context.Context, params TunnelRoutesListParams) ([]TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes", AccountRouteRoot, api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, params)

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

// TunnelRouteForIp finds the Tunnel Route that encompasses the given IP
//
// See: https://api.cloudflare.com/#tunnel-route-get-tunnel-route-by-ip
func (api *API) TunnelRouteForIP(ctx context.Context, ip string) (TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes/ip/%s", AccountRouteRoot, api.AccountID, ip)

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

// CreateTunnelRoute add a new route to the account's routing table for the given tunnel.
//
// See: https://api.cloudflare.com/#tunnel-route-create-route
func (api *API) CreateTunnelRoute(ctx context.Context, tunnelId string, ipNetwork string, comment string) (TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, api.AccountID, url.PathEscape(ipNetwork))

	params := make(map[string]string)
	params["tunnel_id"] = tunnelId
	if comment != "" {
		params["comment"] = comment
	}

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

// DeleteTunnelRoute delete an existing route from the account's routing table
//
// See: https://api.cloudflare.com/#tunnel-route-delete-route
func (api *API) DeleteTunnelRoute(ctx context.Context, ipNetwork string) (TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, api.AccountID, url.PathEscape(ipNetwork))

	responseBody, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
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

// UpdateTunnelRoute update an existing route in the account's routing table for the given tunnel
//
// See: https://api.cloudflare.com/#tunnel-route-update-route
func (api *API) UpdateTunnelRoute(ctx context.Context, tunnelId string, currentNetwork, newNetwork, comment string) (TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, api.AccountID, url.PathEscape(currentNetwork))

	params := map[string]string{
		"tunnel_id":   tunnelId,
		"new_network": newNetwork,
	}
	if comment != "" {
		params["comment"] = comment
	}

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
