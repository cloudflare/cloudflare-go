package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ListManagedHeadersResponse struct {
	Response
	Result ManagedHeaders `json:"result"`
}

type UpdateManagedHeadersParams struct {
	ZoneID string
	ManagedHeaders
}

type ManagedHeaders struct {
	ManagedRequestHeaders  []ManagedHeader `json:"managed_request_headers"`
	ManagedResponseHeaders []ManagedHeader `json:"managed_response_headers"`
}

type ManagedHeader struct {
	ID            string   `json:"id"`
	Enabled       bool     `json:"enabled"`
	HasCoflict    bool     `json:"has_conflict,omitempty"`
	ConflictsWith []string `json:"conflicts_with,omitempty"`
}

type ListManagedHeadersParams struct {
	ZoneID string
}

func (api *API) ListZoneManagedHeaders(ctx context.Context, params ListManagedHeadersParams) (ManagedHeaders, error) {
	if params.ZoneID == "" {
		return ManagedHeaders{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/managed_headers", params.ZoneID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ManagedHeaders{}, err
	}

	result := ListManagedHeadersResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ManagedHeaders{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

func (api *API) UpdateZoneManagedHeaders(ctx context.Context, params UpdateManagedHeadersParams) (ManagedHeaders, error) {
	if params.ZoneID == "" {
		return ManagedHeaders{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/managed_headers", params.ZoneID)

	payload, err := json.Marshal(params.ManagedHeaders)
	if err != nil {
		return ManagedHeaders{}, err
	}

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, payload)
	if err != nil {
		return ManagedHeaders{}, err
	}

	result := ListManagedHeadersResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ManagedHeaders{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}
