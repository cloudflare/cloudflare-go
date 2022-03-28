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
	Network    string    `json:"network"`
	TunnelId   string    `json:"tunnel_id"`
	TunnelName string    `json:"tunnel_name"`
	Comment    string    `json:"comment"`
	CreatedAt  *time.Time `json:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
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

func extractTunnelRouteResponse(responseBody []byte, err error) (TunnelRoute, error) {
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

// TunnelRoutes lists all defined routes for tunnels in the account.
//
// See: https://api.cloudflare.com/#tunnel-route-list-tunnel-routes
func (api *API) TunnelRoutes(ctx context.Context) ([]TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes", AccountRouteRoot, api.AccountID)
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

// TunnelRouteForIp finds the Tunnel Route that encompasses the given IP
//
// See: https://api.cloudflare.com/#tunnel-route-get-tunnel-route-by-ip
func (api *API) TunnelRouteForIP(ctx context.Context, ip string) (TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes/ip/%s", AccountRouteRoot, api.AccountID, ip)
	return extractTunnelRouteResponse(api.makeRequestContext(ctx, http.MethodGet, uri, nil))
}

// CreateTunnelRoute add a new route to the account's routing table for the given tunnel.
//
// See: https://api.cloudflare.com/#tunnel-route-create-route
func (api *API) CreateTunnelRoute(ctx context.Context, tunnelId string, ipNetwork string) (TunnelRoute, error) {
	return api.createTunnelRoute(ctx, tunnelId, ipNetwork, "")
}

// CreateTunnelRouteWithComment add a new route to the account's routing table for the given tunnel.
//
// See: https://api.cloudflare.com/#tunnel-route-create-route
func (api *API) CreateTunnelRouteWithComment(ctx context.Context, tunnelId string, ipNetwork string, comment string) (TunnelRoute, error) {
	return api.createTunnelRoute(ctx, tunnelId, ipNetwork, comment)
}

func (api *API) createTunnelRoute(ctx context.Context, tunnelId string, ipNetwork string, comment string) (TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, api.AccountID, url.PathEscape(ipNetwork))

	params := make(map[string]string)
	params["tunnel_id"] = tunnelId
	if comment != "" {
		params["comment"] = comment
	}

	return extractTunnelRouteResponse(api.makeRequestContext(ctx, http.MethodPost, uri, params))
}

// DeleteTunnelRoute delete an existing route from the account's routing table
//
// See: https://api.cloudflare.com/#tunnel-route-delete-route
func (api *API) DeleteTunnelRoute(ctx context.Context, ipNetwork string) (TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, api.AccountID, url.PathEscape(ipNetwork))
	return extractTunnelRouteResponse(api.makeRequestContext(ctx, http.MethodDelete, uri, nil))
}

// UpdateTunnelRoute update an existing route in the account's routing table for the given tunnel
//
// See: https://api.cloudflare.com/#tunnel-route-update-route
func (api *API) UpdateTunnelRoute(ctx context.Context, tunnelId string, currentNetwork string, newNetwork string) (TunnelRoute, error) {
	return api.updateTunnelRoute(ctx, tunnelId, currentNetwork, newNetwork, "")
}

// UpdateTunnelRouteWithComment update an existing route in the account's routing table for the given tunnel
//
// See: https://api.cloudflare.com/#tunnel-route-update-route
func (api *API) UpdateTunnelRouteWithComment(ctx context.Context, tunnelId string, currentNetwork string, newNetwork string, comment string) (TunnelRoute, error) {
	return api.updateTunnelRoute(ctx, tunnelId, currentNetwork, newNetwork, comment)
}

func (api *API) updateTunnelRoute(ctx context.Context, tunnelId string, currentNetwork string, newNetwork string, comment string) (TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, api.AccountID, url.PathEscape(currentNetwork))

	params := map[string]string{
		"tunnel_id":   tunnelId,
		"new_network": newNetwork,
	}
	if comment != "" {
		params["comment"] = comment
	}

	return extractTunnelRouteResponse(api.makeRequestContext(ctx, http.MethodPatch, uri, params))
}
