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

type ManagedHeadersOptions struct {
	ZoneID string
}

func (api *API) ListZoneManagedHeaders(ctx context.Context, opts ManagedHeadersOptions) (ManagedHeaders, error) {
	uri := fmt.Sprintf("/zones/%s/managed_headers", opts.ZoneID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ManagedHeaders{}, err
	}

	result := ListManagedHeadersResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ManagedHeaders{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

func (api *API) UpdateZoneManagedHeaders(ctx context.Context, headers ManagedHeaders, opts ManagedHeadersOptions) (ManagedHeaders, error) {
	uri := fmt.Sprintf("/zones/%s/managed_headers", opts.ZoneID)

	request, err := json.Marshal(headers)
	if err != nil {
		return ManagedHeaders{}, errors.Wrap(err, errMakeRequestError)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, request)
	if err != nil {
		return ManagedHeaders{}, err
	}

	result := ListManagedHeadersResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ManagedHeaders{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}
