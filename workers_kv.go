package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"

	"github.com/pkg/errors"
)

// StorageNamespaceRequest provides parameters for creating and updating storage namespaces
type StorageNamespaceRequest struct {
	Title string
}

// StorageNamespaceResponse is the response received when creating storage namespaces
type StorageNamespaceResponse struct {
	Response
	Result StorageNamespace `json:"result"`
}

// StorageNamespace contains the unique identifier and title of a storage namespace
type StorageNamespace struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// ListStorageNamespacesResponse contains a slice of storage namespaces associated with an
// account, pagination information, and an embedded response struct
type ListStorageNamespacesResponse struct {
	Response
	Result     []StorageNamespace `json:"result"`
	ResultInfo `json:"result_info"`
}

// StorageKey is a key name used to identify a storage value
type StorageKey struct {
	Name string `json:"name"`
}

// ListStorageKeysResponse contains a slice of keys belonging to a storage namespace,
// pagination information, and an embedded response struct
type ListStorageKeysResponse struct {
	Response
	Result     []StorageKey `json:"result"`
	ResultInfo `json:"result_info"`
}

// CreateStorageNamespace creates a namespace under the given title.
// A 400 is returned if the account already owns a namespace with this title.
// A namespace must be explicitly deleted to be replaced.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-create-a-namespace
func (api *API) CreateStorageNamespace(ctx context.Context, req *StorageNamespaceRequest) (StorageNamespaceResponse, error) {
	uri := path.Join("/accounts", api.OrganizationID, "storage/kv/namespaces")
	res, err := api.makeRequestContext(ctx, "POST", uri, req)
	if err != nil {
		return StorageNamespaceResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := StorageNamespaceResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ListStorageNamespaces lists storage namespaces
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-list-namespaces
func (api *API) ListStorageNamespaces(ctx context.Context) (ListStorageNamespacesResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces", api.OrganizationID)
	res, err := api.makeRequestContext(ctx, "GET", uri, nil)
	if err != nil {
		return ListStorageNamespacesResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := ListStorageNamespacesResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// DeleteStorageNamespace deletes the namespace corresponding to the given ID
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-remove-a-namespace
func (api *API) DeleteStorageNamespace(ctx context.Context, namespace string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.OrganizationID, namespace)
	res, err := api.makeRequestContext(ctx, "DELETE", uri, nil)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// UpdateStorageNamespace modifies a namespace's title
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-rename-a-namespace
func (api *API) UpdateStorageNamespace(ctx context.Context, namespace string, req *StorageNamespaceRequest) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s", api.OrganizationID, namespace)
	res, err := api.makeRequestContext(ctx, "PUT", uri, req)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// CreateStorageKV writes a value identified by a key. Body should be the value to be stored.
// Existing values will be overwritten.
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-write-key-value-pair
func (api *API) CreateStorageKV(ctx context.Context, namespace, key string, value io.Reader) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.OrganizationID, namespace, key)
	res, err := api.makeRequestWithAuthTypeAndHeaders(
		ctx, "PUT", uri, value, api.authType, http.Header{"Content-Type": []string{"binary/octet-stream"}},
	)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ReadStorageKV returns the value associated with the given key in the given namespace
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-read-key-value-pair
func (api API) ReadStorageKV(ctx context.Context, namespace, key string) ([]byte, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.OrganizationID, namespace, key)
	res, err := api.makeRequestContext(ctx, "GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	return res, nil
}

// DeleteStorageKV deletes a key and value for a provided storage namespace
//
// API reference: https://api.cloudflare.com/#workers-kv-namespace-delete-key-value-pair
func (api API) DeleteStorageKV(ctx context.Context, namespace, key string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/values/%s", api.OrganizationID, namespace, key)
	res, err := api.makeRequestContext(ctx, "DELETE", uri, nil)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}
	return result, err
}

// ListStorageKeys lists a namespace's keys
//
// API Reference: https://api.cloudflare.com/#workers-kv-namespace-list-a-namespace-s-keys
func (api API) ListStorageKeys(ctx context.Context, namespace string) (ListStorageKeysResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/storage/kv/namespaces/%s/keys", api.OrganizationID, namespace)
	res, err := api.makeRequestContext(ctx, "GET", uri, nil)
	if err != nil {
		return ListStorageKeysResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := ListStorageKeysResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}
	return result, err
}
