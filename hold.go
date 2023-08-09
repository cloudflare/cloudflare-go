package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// Retrieve whether the zone is subject to a zone hold, and metadata about the hold.
type ZoneHold struct {
	Hold              bool    `json:"hold"`
	IncludeSubdomains *bool   `json:"include_subdomains,omitempty"`
	HoldAfter         *string `json:"hold_after,omitempty"`
}

// ZoneHoldResponse represents a response from the Zone Hold endpoint.
type ZoneHoldResponse struct {
	Result ZoneHold `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// ZoneHoldDeleteResponse represents a response from the Delete Zone Hold
// endpoint.
type ZoneHoldDeleteResponse struct {
	Result ZoneHold `json:"result"`
}

// ZoneHoldCreateParams represents params for the Create Zone Hold
// endpoint.
type ZoneHoldCreateParams struct {
	ID                string `json:"id"`
	Hold              bool   `json:"hold"`
	IncludeSubdomains *bool  `json:"include_subdomains,omitempty"`
}

// ZoneHoldDeleteParams represents params for the Delete Zone Hold
// endpoint.
type ZoneHoldDeleteParams struct {
	ID        string  `json:"id"`
	Hold      bool    `json:"hold"`
	HoldAfter *string `json:"hold_after,omitempty"`
}

// Enforce a zone hold on the zone, blocking the creation and activation of zones with this zone's hostname.
//
// API reference: https://developers.cloudflare.com/api/operations/zones-0-hold-post
func (api *API) CreateZoneHold(ctx context.Context, zoneID string, params ZoneHoldCreateParams) (ZoneHold, error) {
	uri := fmt.Sprintf("/zones/%s/hold", zoneID)
	if params.IncludeSubdomains != nil {
		uri = fmt.Sprintf("%s?include_subdomains=%t", uri, *params.IncludeSubdomains)
	}
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return ZoneHold{}, err
	}

	response := &ZoneHoldResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneHold{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	if !response.Success {
		return ZoneHold{}, fmt.Errorf(response.Errors[0].Message)
	}

	return response.Result, nil
}

// Stop enforcement of a zone hold on the zone, permanently or temporarily, allowing the creation and activation of zones with this zone's hostname.
//
// API reference:https://developers.cloudflare.com/api/operations/zones-0-hold-delete
func (api *API) DeleteZoneHold(ctx context.Context, zoneID string, params ZoneHoldDeleteParams) (ZoneHold, error) {
	uri := fmt.Sprintf("/zones/%s/hold", zoneID)
	if params.HoldAfter != nil {
		uri = fmt.Sprintf("%s?hold_after=%s", uri, *params.HoldAfter)
	}
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return ZoneHold{}, err
	}

	response := &ZoneHoldDeleteResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneHold{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// Retrieve whether the zone is subject to a zone hold, and metadata about the hold.
//
// API reference: https://developers.cloudflare.com/api/operations/zones-0-hold-get
func (api *API) ZoneHold(ctx context.Context, zoneID string) (ZoneHold, error) {
	uri := fmt.Sprintf("/zones/%s/hold", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZoneHold{}, err
	}

	response := &ZoneHoldResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneHold{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}
