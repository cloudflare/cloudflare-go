package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// KeylessSSL represents Keyless SSL configuration.
type KeylessSSL struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Host        string    `json:"host"`
	Port        int       `json:"port"`
	Status      string    `json:"status"`
	Enabled     bool      `json:"enabled"`
	Permissions []string  `json:"permissions"`
	CreatedOn   time.Time `json:"created_on"`
	ModifiedOn  time.Time `json:"modified_on"`
}

// KeylessSSLCreateRequest represents the request format made for creating KeylessSSL.
type KeylessSSLCreateRequest struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Certificate  string `json:"certificate"`
	Name         string `json:"name,omitempty"`
	BundleMethod string `json:"bundle_method,omitempty"`
}

// KeylessSSLDetailResponse is the API response, containing a single Keyless SSL.
type KeylessSSLDetailResponse struct {
	Response
	Result KeylessSSL `json:"result"`
}

// KeylessSSLListResponse represents the response from the Keyless SSL list endpoint.
type KeylessSSLListResponse struct {
	Response
	Result []KeylessSSL `json:"result"`
}

// KeylessSSLUpdateRequest represents the request for updating KeylessSSL.
type KeylessSSLUpdateRequest struct {
	Host    string `json:"host,omitempty"`
	Name    string `json:"name,omitempty"`
	Port    int    `json:"port,omitempty"`
	Enabled *bool  `json:"enabled,omitempty"`
}

// CreateKeyless creates a new Keyless SSL configuration for the zone.
//
// API reference: https://api.cloudflare.com/#keyless-ssl-for-a-zone-create-keyless-ssl-configuration
func (api *API) CreateKeyless(ctx context.Context, zoneID string, keylessSSL KeylessSSLCreateRequest) (KeylessSSL, error) {
	uri := fmt.Sprintf("/zones/%s/keyless_certificates", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, keylessSSL)
	if err != nil {
		return KeylessSSL{}, err
	}

	var keylessSSLDetailResponse KeylessSSLDetailResponse
	err = json.Unmarshal(res, &keylessSSLDetailResponse)
	if err != nil {
		return KeylessSSL{}, errors.Wrap(err, errUnmarshalError)
	}

	return keylessSSLDetailResponse.Result, nil
}

// ListKeyless lists Keyless SSL configurations for a zone.
//
// API reference: https://api.cloudflare.com/#keyless-ssl-for-a-zone-list-keyless-ssl-configurations
func (api *API) ListKeyless(ctx context.Context, zoneID string) ([]KeylessSSL, error) {
	uri := fmt.Sprintf("/zones/%s/keyless_certificates", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var keylessSSLListResponse KeylessSSLListResponse
	err = json.Unmarshal(res, &keylessSSLListResponse)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return keylessSSLListResponse.Result, nil
}

// Keyless provides the configuration for a given Keyless SSL identifier.
//
// API reference: https://api.cloudflare.com/#keyless-ssl-for-a-zone-keyless-ssl-details
func (api *API) Keyless(ctx context.Context, zoneID, keylessSSLID string) (KeylessSSL, error) {
	uri := fmt.Sprintf("/zones/%s/keyless_certificates/%s", zoneID, keylessSSLID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return KeylessSSL{}, err
	}

	var keylessResponse KeylessSSLDetailResponse
	err = json.Unmarshal(res, &keylessResponse)
	if err != nil {
		return KeylessSSL{}, errors.Wrap(err, errUnmarshalError)
	}

	return keylessResponse.Result, nil
}

// UpdateKeyless updates an existing Keyless SSL configuration.
//
// API reference: https://api.cloudflare.com/#keyless-ssl-for-a-zone-edit-keyless-ssl-configuration
func (api *API) UpdateKeyless(ctx context.Context, zoneID, kelessSSLID string, keylessSSL KeylessSSLUpdateRequest) (KeylessSSL, error) {
	uri := fmt.Sprintf("/zones/%s/keyless_certificates/%s", zoneID, kelessSSLID)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, keylessSSL)
	if err != nil {
		return KeylessSSL{}, err
	}

	var keylessSSLDetailResponse KeylessSSLDetailResponse
	err = json.Unmarshal(res, &keylessSSLDetailResponse)
	if err != nil {
		return KeylessSSL{}, errors.Wrap(err, errUnmarshalError)
	}

	return keylessSSLDetailResponse.Result, nil
}

// DeleteKeyless deletes an existing Keyless SSL configuration.
//
// API reference: https://api.cloudflare.com/#keyless-ssl-for-a-zone-delete-keyless-ssl-configuration
func (api *API) DeleteKeyless(ctx context.Context, zoneID, keylessSSLID string) error {
	uri := fmt.Sprintf("/zones/%s/keyless_certificates/%s", zoneID, keylessSSLID)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}
