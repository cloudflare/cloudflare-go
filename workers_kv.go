package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// WorkersKVNamespace contains the unique identifier and title of a kv namespace.
type WorkersKVNamespace struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// WorkersKVNamespaceRequest provides parameters for creating and updating kv namespaces.
type WorkersKVNamespaceRequest struct {
	Title string `json:"title"`
}

// WorkersKVNamespaceResponse is the response received when creating or updating kv namespaces.
type WorkersKVNamespaceResponse struct {
	Response
	Result WorkersKVNamespace `json:"result"`
}

// WorkersKVNamespaceListResponse contains a slice of kv namespaces associated with an
// account, pagination information, and an embedded response struct.
type WorkersKVNamespaceListResponse struct {
	Response
	Result     []WorkersKVNamespace `json:"result"`
	ResultInfo `json:"result_info"`
}

// ListWorkersKVNamespacesResponse is deprecated, use WorkersKVNamespaceListResponse instead.
type ListWorkersKVNamespacesResponse = WorkersKVNamespaceListResponse

// WorkersKVKey contains the name and expiration of a kv key.
type WorkersKVKey struct {
	Name       string `json:"name"`
	Expiration uint   `json:"expiration,omitempty"`
}

// WorkersKVPair contains all key, value, and expiration of a kv pair.
type WorkersKVPair struct {
	Name          string `json:"key"`
	Value         string `json:"value"`
	Expiration    uint   `json:"expiration,omitempty"`
	ExpirationTTL uint   `json:"expiration_ttl,omitempty"`
}

// WorkersKVKeyListResponse is the response received when listing kv keys in a namespace.
type WorkersKVKeyListResponse struct {
	Response
	Result     []WorkersKVKey         `json:"result"`
	ResultInfo WorkersKVKeyPagination `json:"result_info"`
}

// WorkersKVKeyPagination contains metadata about listing kv keys in a namespace.
type WorkersKVKeyPagination struct {
	Count  uint   `json:"count"`
	Cursor string `json:"cursor"`
}

// CreateWorkersKVNamespace creates a namespace under the given title.
// A 400 is returned if the account already owns a namespace with this title.
// A namespace must be explicitly deleted to be replaced.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-create-a-namespace
func (api *API) CreateWorkersKVNamespace(ctx context.Context, req *WorkersKVNamespaceRequest) (WorkersKVNamespaceResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces", api.OrganizationID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, req)
	if err != nil {
		return WorkersKVNamespaceResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := WorkersKVNamespaceResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ListWorkersKVNamespaces lists kv namespaces.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-list-namespaces
func (api *API) ListWorkersKVNamespaces(ctx context.Context) (WorkersKVNamespaceListResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces", api.OrganizationID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkersKVNamespaceListResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := WorkersKVNamespaceListResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// DeleteWorkersKVNamespace deletes the namespace corresponding to the given ID.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-remove-a-namespace
func (api *API) DeleteWorkersKVNamespace(ctx context.Context, namespaceID string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.OrganizationID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// UpdateWorkersKVNamespace modifies a namespace's title.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-rename-a-namespace
func (api *API) UpdateWorkersKVNamespace(ctx context.Context, namespaceID string, req *WorkersKVNamespaceRequest) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.OrganizationID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, req)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// WriteWorkersKVPair writes a value identified by a key.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-key-value-pair
func (api *API) WriteWorkersKVPair(ctx context.Context, namespaceID string, pair *WorkersKVPair) (Response, error) {
	params := url.Values{}
	expires := pair.Expiration
	if expires > 0 {
		params.Set("expiration", strconv.Itoa(int(expires)))
	}
	expires = pair.ExpirationTTL
	if expires > 0 {
		params.Set("expiration_ttl", strconv.Itoa(int(expires)))
	}

	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s?"+params.Encode(), api.OrganizationID, namespaceID, url.PathEscape(pair.Name))
	res, err := api.makeRequestWithAuthTypeAndHeaders(
		ctx, http.MethodPut, uri, pair.Value, api.authType, http.Header{"Content-Type": []string{"application/octet-stream"}})
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// WriteWorkersKV writes a value identified by a key.
//
// Deprecated: Use WriteWorkersKVPair instead.
func (api API) WriteWorkersKV(ctx context.Context, namespaceID, key string, value []byte) (Response, error) {
	return api.WriteWorkersKVPair(ctx, namespaceID, &WorkersKVPair{Name: key, Value: string(value)})
}

// WriteWorkersKVPairBulk writes multiple key-value pairs.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-key-value-pair
func (api *API) WriteWorkersKVPairBulk(ctx context.Context, namespaceID string, pairs []*WorkersKVPair) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/bulk", api.OrganizationID, namespaceID)
	res, err := api.makeRequestWithAuthTypeAndHeaders(
		ctx, http.MethodPut, uri, pairs, api.authType, http.Header{"Content-Type": []string{"application/json"}})
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ReadWorkersKVPair returns the kv pair associated with the given key in the given namespace.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-read-key-value-pair
func (api *API) ReadWorkersKVPair(ctx context.Context, namespaceID, key string) (WorkersKVPair, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.OrganizationID, namespaceID, url.PathEscape(key))
	res, headers, err := api.makeRequestWithAuthTypeAndHeadersTuple(ctx, http.MethodGet, uri, nil, api.authType, http.Header{"Content-Type": []string{"application/octet-stream"}})
	if err != nil {
		return WorkersKVPair{}, errors.Wrap(err, errMakeRequestError)
	}

	expiration, _ := strconv.ParseUint(headers.Get("Expiration"), 10, 64)
	return WorkersKVPair{
		Name:       key,
		Value:      string(res),
		Expiration: uint(expiration),
	}, nil
}

// ReadWorkersKV returns the kv pair associated with the given key in the given namespace.
//
// Deprecated: Use ReadWorkersKVPair instead.
func (api API) ReadWorkersKV(ctx context.Context, namespaceID, key string) ([]byte, error) {
	pair, err := api.ReadWorkersKVPair(ctx, namespaceID, key)
	return []byte(pair.Value), err
}

// DeleteWorkersKVPair deletes the kv pair associated with the given key in the given namespace.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-delete-key-value-pair
func (api *API) DeleteWorkersKVPair(ctx context.Context, namespaceID, key string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.OrganizationID, namespaceID, url.PathEscape(key))
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}
	return result, err
}

// DeleteWorkersKV deletes the kv pair associated with the given key in the given namespace.
//
// Deprecated: Use DeleteWorkersKVPair instead.
func (api API) DeleteWorkersKV(ctx context.Context, namespaceID, key string) (Response, error) {
	return api.DeleteWorkersKVPair(ctx, namespaceID, key)
}

// ListWorkersKVKeys lists the keys in the given namespace.
//
// API Reference: https://api.cloudflare.com/#workers-kv-namespace-list-a-namespace-s-keys
func (api *API) ListWorkersKVKeys(ctx context.Context, namespaceID string, request *WorkersKVKeyPagination) (WorkersKVKeyListResponse, error) {
	params := url.Values{}
	limit := request.Count
	if limit > 0 {
		params.Set("limit", strconv.Itoa(int(limit)))
	}
	cursor := request.Cursor
	if cursor != "" {
		params.Set("cursor", cursor)
	}

	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys?"+params.Encode(), api.OrganizationID, namespaceID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkersKVKeyListResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := WorkersKVKeyListResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}
	return result, err
}
