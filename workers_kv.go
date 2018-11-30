package cloudflare

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type StorageNamespaceRequest struct {
	Title string
}

type StorageNamespaceResponse struct {
	ID, Title string
}

func (a *API) CreateStorageNamespace(req StorageNamespaceRequest) (StorageNamespaceResponse, error) {
	uri := fmt.Sprintf("/account/%s/workers/namespaces/", a.OrganizationID)
	res, err := a.makeRequest("POST", uri, req)
	if err != nil {
		return StorageNamespaceResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := StorageNamespaceResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

func (a *API) DeleteStorageNamespace() {

}

func (a *API) UpdateStorageNamespace() {

}

func (a *API) ListStorageNamespaces() {

}

func (a *API) CreateStorageKV() {

}

func (a API) ReadStorageValue() {

}

func (a API) DeleteStorageKV() {
}

func (a API) ListStorageKeys() {

}
