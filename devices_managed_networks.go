package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	TlsSockAddr string `json:"tls_sockaddr,omitempty"`
	Sha256      string `json:"sha256,omitempty"`
}

type DeviceManagedNetworks struct {
	NetworkID string  `json:"network_id,omitempty"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Config    *Config `json:"config"`
}

type DeviceManagedNetworksResponse struct {
	Response
	Result DeviceManagedNetworks `json:"result"`
}

type DeviceManagedNetworksListResponse struct {
	Response
	Result []DeviceManagedNetworks `json:"result"`
}

// ListTeamsDevice returns all devices for a given account.
//
// API reference : https://api.cloudflare.com/#device-managed-networks-list-device-managed-networks
func (api *API) ListManagedNetworks(ctx context.Context, accountID string) ([]DeviceManagedNetworks, error) {
	uri := fmt.Sprintf("/%s/%s/devices/networks", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []DeviceManagedNetworks{}, err
	}

	var response DeviceManagedNetworksListResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return []DeviceManagedNetworks{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// CreateDeviceManagedNetwork creates a new Device Managed Network.
//
// API reference: https://api.cloudflare.com/#device-managed-networks-create-device-managed-network
func (api *API) CreateDeviceManagedNetwork(ctx context.Context, accountID string, req DeviceManagedNetworks) (DeviceManagedNetworks, error) {
	uri := fmt.Sprintf("/%s/%s/devices/networks", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, req)
	if err != nil {
		return DeviceManagedNetworks{}, err
	}

	var deviceManagedNetworksResponse DeviceManagedNetworksResponse
	if err := json.Unmarshal(res, &deviceManagedNetworksResponse); err != nil {
		return DeviceManagedNetworks{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return deviceManagedNetworksResponse.Result, err
}

// UpdateDeviceManagedNetwork Update a Device Managed Network.
//
// API reference: https://api.cloudflare.com/#device-managed-networks-update-device-managed-network
func (api *API) UpdateDeviceManagedNetwork(ctx context.Context, accountID string, networkID string, req DeviceManagedNetworks) (DeviceManagedNetworks, error) {
	uri := fmt.Sprintf("/%s/%s/devices/networks/%s", AccountRouteRoot, accountID, networkID)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, req)
	if err != nil {
		return DeviceManagedNetworks{}, err
	}

	var deviceManagedNetworksResponse DeviceManagedNetworksResponse

	if err := json.Unmarshal(res, &deviceManagedNetworksResponse); err != nil {
		return DeviceManagedNetworks{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return deviceManagedNetworksResponse.Result, err
}

// GetDeviceManagedNetwork gets a single Managed Network.
//
// API reference: https://api.cloudflare.com/#device-managed-networks-device-managed-network-details
func (api *API) GetDeviceManagedNetwork(ctx context.Context, accountID string, networkID string) (DeviceManagedNetworks, error) {
	uri := fmt.Sprintf("/%s/%s/devices/networks/%s", AccountRouteRoot, accountID, networkID)

	deviceManagedNetworksResponse := DeviceManagedNetworksResponse{}
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DeviceManagedNetworks{}, err
	}

	if err := json.Unmarshal(res, &deviceManagedNetworksResponse); err != nil {
		return DeviceManagedNetworks{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return deviceManagedNetworksResponse.Result, err
}

// DeleteManagedNetworks deletes a Device Managed Network. Returns the remaining device managed networks for the account.
//
// API reference: https://api.cloudflare.com/#device-managed-networks-delete-device-managed-network
func (api *API) DeleteManagedNetworks(ctx context.Context, accountID, networkID string) ([]DeviceManagedNetworks, error) {
	uri := fmt.Sprintf("/%s/%s/devices/networks/%s", AccountRouteRoot, accountID, networkID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return []DeviceManagedNetworks{}, err
	}

	var response DeviceManagedNetworksListResponse
	if err := json.Unmarshal(res, &response); err != nil {
		return []DeviceManagedNetworks{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, err
}
