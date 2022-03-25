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
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

// tunnelRouteListResponse is the API response for listing tunnel routes.
type tunnelRouteListResponse struct {
	Response
	Result []TunnelRoute `json:"result"`
}

type tunnelRouteForIpResponse struct {
	Response
	Result TunnelRoute `json:"result"`
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
func (api *API) TunnelRouteForIp(ctx context.Context, ip string) (TunnelRoute, error) {
	uri := fmt.Sprintf("/%s/%s/teamnet/routes/ip/%s", AccountRouteRoot, api.AccountID, ip)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TunnelRoute{}, err
	}

	var routeResponse tunnelRouteForIpResponse
	err = json.Unmarshal(res, &routeResponse)
	if err != nil {
		return TunnelRoute{}, errors.Wrap(err, errUnmarshalError)
	}

	return routeResponse.Result, nil
}

// CreateTunnelRoute add a new route to the account's routing table for the given tunnel
//
// See: https://api.cloudflare.com/#tunnel-route-create-route
func (api *API) CreateTunnelRoute(ctx context.Context, tunnelId string, ipNetwork string) (TunnelRoute, error) {
	return api.createTunnelRoute(ctx, tunnelId, ipNetwork, "")
}

// CreateTunnelRouteWithComment add a new route to the account's routing table for the given tunnel
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

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return TunnelRoute{}, err
	}

	var routeResponse tunnelRouteForIpResponse
	err = json.Unmarshal(res, &routeResponse)
	if err != nil {
		return TunnelRoute{}, errors.Wrap(err, errUnmarshalError)
	}

	return routeResponse.Result, nil
}
