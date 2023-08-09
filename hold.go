package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// ZoneHold represents a Zone Hold rule. A rule only permits access to
// the provided URL pattern(s) from the given IP address(es) or subnet(s).
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

// ZoneHoldDeleteResponse represents a response from the List Zone Hold
// endpoint.
type ZoneHoldDeleteResponse struct {
	Result ZoneHold `json:"result"`
}

// ZoneHoldCreateParams contains required and optional params
// for creating a zone Hold.
type ZoneHoldCreateParams struct {
	ID                string `json:"id"`
	Hold              bool   `json:"hold"`
	IncludeSubdomains *bool  `json:"include_subdomains,omitempty"`
}

// ZoneHoldDeleteParams contains required and optional params
// for updating a zone Hold.
type ZoneHoldDeleteParams struct {
	ID        string  `json:"id"`
	Hold      bool    `json:"hold"`
	HoldAfter *string `json:"hold_after,omitempty"`
}

// CreateZoneHold creates a Zone ZoneHold rule for the given zone ID.
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

// DeleteZoneHold deletes a Zone ZoneHold rule (based on the ID) for the given zone ID.
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

// ZoneHold retrieves a Zone ZoneHold rule (based on the ID) for the given zone ID.
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
