package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type ListManagedHeadersResponse struct {
	Response
	Result ManagedHeaders `json:"result"`
}

type UpdateManagedHeadersParams struct {
	ZoneID string
	ManagedHeaders
	Filter func(ManagedHeader) bool
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
	Filter func(ManagedHeader) bool
}

func filterManagedHeaders(headers ManagedHeaders, filter func(ManagedHeader) bool) ManagedHeaders {
	if filter == nil {
		filter = func(ManagedHeader) bool {
			return true
		}
	}

	var managedRequestHeaders []ManagedHeader
	for i := range headers.ManagedRequestHeaders {
		if filter(headers.ManagedRequestHeaders[i]) {
			managedRequestHeaders = append(managedRequestHeaders, headers.ManagedRequestHeaders[i])
		}
	}
	var managedResponseHeaders []ManagedHeader
	for i := range headers.ManagedResponseHeaders {
		if filter(headers.ManagedResponseHeaders[i]) {
			managedResponseHeaders = append(managedResponseHeaders, headers.ManagedResponseHeaders[i])
		}
	}

	return ManagedHeaders{
		ManagedRequestHeaders:  managedRequestHeaders,
		ManagedResponseHeaders: managedResponseHeaders,
	}
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
		return ManagedHeaders{}, errors.Wrap(err, errUnmarshalError)
	}

	return filterManagedHeaders(result.Result, params.Filter), nil
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
		return ManagedHeaders{}, errors.Wrap(err, errUnmarshalError)
	}

	return filterManagedHeaders(result.Result, params.Filter), nil
}
