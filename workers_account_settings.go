package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type WorkersAccountSettings struct {
	DefaultUsageModel string `json:"default_usage_model,omitempty"`
	GreenCompute      bool   `json:"green_compute,omitempty"`
}

type CreateWorkersAccountSettingsParameters struct {
	AccountID         string `json:"-"`
	DefaultUsageModel string `json:"default_usage_model,omitempty"`
	GreenCompute      bool   `json:"green_compute,omitempty"`
}
type CreateWorkersAccountSettingsResponse struct {
	Response
	Result WorkersAccountSettings
}

type GetWorkersAccountSettingsParameters struct {
	AccountID string
}

type GetWorkersAccountSettingsResponse struct {
	Response
	Result WorkersAccountSettings
}

// CreateWorkersAccountSettings Starts a tail that receives logs and exception from a Worker
//
// API reference: https://api.cloudflare.com/#worker-account-settings-create-worker-account-settings
func (api *API) CreateWorkersAccountSettings(ctx context.Context, params CreateWorkersAccountSettingsParameters) (WorkersAccountSettings, error) {
	if params.AccountID == "" {
		return WorkersAccountSettings{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/account-settings", params.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return WorkersAccountSettings{}, err
	}
	var workersAccountSettingsResponse CreateWorkersAccountSettingsResponse
	if err := json.Unmarshal(res, &workersAccountSettingsResponse); err != nil {
		return WorkersAccountSettings{}, err
	}
	return workersAccountSettingsResponse.Result, nil
}

// GetWorkersAccountSettings Starts a tail that receives logs and exception from a Worker
//
// API reference: https://api.cloudflare.com/#worker-account-settings-fetch-worker-account-settings
func (api *API) GetWorkersAccountSettings(ctx context.Context, params GetWorkersAccountSettingsParameters) (WorkersAccountSettings, error) {
	if params.AccountID == "" {
		return WorkersAccountSettings{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/account-settings", params.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, params)
	if err != nil {
		return WorkersAccountSettings{}, err
	}
	var workersAccountSettingsResponse CreateWorkersAccountSettingsResponse
	if err := json.Unmarshal(res, &workersAccountSettingsResponse); err != nil {
		return WorkersAccountSettings{}, err
	}
	return workersAccountSettingsResponse.Result, nil
}
