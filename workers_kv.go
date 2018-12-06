package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// StorageNamespaceRequest
type StorageNamespaceRequest struct {
	Title string
}

// StorageNamespaceResponse
type StorageNamespaceResponse struct {
	Response
	ID, Title string
}

// StorageNamespace
type StorageNamespace struct {
	ID, Title string
}

// ListStorageNamespacesResponse
type ListStorageNamespacesResponse struct {
	Response
	Result     []StorageNamespace `json:"result"`
	ResultInfo `json:"result_info"`
}

func (api *API) CreateStorageNamespace(ctx context.Context, req *StorageNamespaceRequest) (StorageNamespaceResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/namespaces", api.OrganizationID)
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

func (api *API) ListStorageNamespaces(ctx context.Context) (ListStorageNamespacesResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/namespaces", api.OrganizationID)
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

func (api *API) DeleteStorageNamespace(ctx context.Context, namespace string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/namespaces/%s", api.OrganizationID, namespace)
	fmt.Println(uri)
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

func (api *API) UpdateStorageNamespace() {

}

func (api *API) CreateStorageKV() {

}

func (api API) ReadStorageValue() {

}

func (api API) DeleteStorageKV() {
}

func (api API) ListStorageKeys() {

}
