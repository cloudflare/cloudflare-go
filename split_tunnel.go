package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// SplitTunnelIncludeGetResponse represents the response from the get split
// tunnel endpoints.
type SplitTunnelGetResponse struct {
	Response
	Result []SplitTunnel `json:"result"`
}

// SplitTunnel represents the individual tunnel struct.
type SplitTunnel struct {
	Address     string `json:"address,omitempty"`
	Host        string `json:"host,omitempty"`
	Description string `json:"description,omitempty"`
}


// SplitTunnelInclude returns all split tunnel include rules within an account.
//
// API reference: https://api.cloudflare.com/#device-policy-get-split-tunnel-include-list
func (api *API) SplitTunnelInclude(ctx context.Context, accountID string) ([]SplitTunnel, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/include", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []SplitTunnel{}, err
	}

	var splitTunnelIncludeGetResponse	SplitTunnelGetResponse
	err = json.Unmarshal(res, &splitTunnelIncludeGetResponse)
	if err != nil {
		return []SplitTunnel{}, errors.Wrap(err, errUnmarshalError)
	}

	return splitTunnelIncludeGetResponse.Result, nil
}

// UpdateSplitTunnelInclude updates the existing split tunnel include policy.
//
// API reference: https://api.cloudflare.com/#device-policy-set-split-tunnel-include-list
func (api *API) UpdateSplitTunnelInclude(ctx context.Context, accountID string, tunnels []SplitTunnel) ([]SplitTunnel, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/include", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, tunnels)
	if err != nil {
		return []SplitTunnel{}, err
	}

	var splitTunnelIncludeGetResponse	SplitTunnelGetResponse
	err = json.Unmarshal(res, &splitTunnelIncludeGetResponse)
	if err != nil {
		return []SplitTunnel{}, errors.Wrap(err, errUnmarshalError)
	}

	return splitTunnelIncludeGetResponse.Result, nil
}

// SplitTunnelExclude returns all split tunnel exclude rules within an account.
//
// API reference: https://api.cloudflare.com/#device-policy-get-split-tunnel-exclude-list
func (api *API) SplitTunnelExclude(ctx context.Context, accountID string) ([]SplitTunnel, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/exclude", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []SplitTunnel{}, err
	}

	var splitTunnelExcludeGetResponse	SplitTunnelGetResponse
	err = json.Unmarshal(res, &splitTunnelExcludeGetResponse)
	if err != nil {
		return []SplitTunnel{}, errors.Wrap(err, errUnmarshalError)
	}

	return splitTunnelExcludeGetResponse.Result, nil
}

// UpdateSplitTunnelExclude updates the existing split tunnel exclude policy.
//
// API reference: https://api.cloudflare.com/#device-policy-set-split-tunnel-exclude-list
func (api *API) UpdateSplitTunnelExclude(ctx context.Context, accountID string, tunnels []SplitTunnel) ([]SplitTunnel, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/exclude", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, tunnels)
	if err != nil {
		return []SplitTunnel{}, err
	}

	var splitTunnelExcludeGetResponse	SplitTunnelGetResponse
	err = json.Unmarshal(res, &splitTunnelExcludeGetResponse)
	if err != nil {
		return []SplitTunnel{}, errors.Wrap(err, errUnmarshalError)
	}

	return splitTunnelExcludeGetResponse.Result, nil
}