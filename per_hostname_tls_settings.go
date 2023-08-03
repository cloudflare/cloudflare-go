package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// HostnameTLSSetting represents the metadata for a user-created tls setting.
type HostnameTLSSetting struct {
	Hostname  string      `json:"hostname"`
	Value     interface{} `json:"value"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// HostnameTLSSettingResponse represents the response from the PUT and DELETE endpoints for per-hostname tls settings.
type HostnameTLSSettingResponse struct {
	Response
	Result HostnameTLSSetting `json:"result"`
}

// HostnameTLSSettingsResponse represents the response from the retrieval endpoint for per-hostname tls settings.
type HostnameTLSSettingsResponse struct {
	Response
	Result     []HostnameTLSSetting `json:"result"`
	ResultInfo `json:"result_info"`
}

// EditHostnameTLSSettingParams represents the data related to the per-hostname tls setting being edited.
type EditHostnameTLSSettingParams struct {
	Value interface{} `json:"value"`
}

type ListHostnameTLSSettingsParams struct {
	PaginationOptions
	Limit    int      `url:"limit,omitempty"`
	Offset   int      `url:"offset,omitempty"`
	Hostname []string `url:"hostname,omitempty"`
}

var (
	ErrMissingHostnameTLSSettingName = errors.New("tls setting name required but missing")
)

// ListHostnameTLSSettings returns a list of all user-created tls setting values for the specified setting and hostnames.
//
// API reference: https://developers.cloudflare.com/api/operations/per-hostname-tls-settings-list
func (api *API) ListHostnameTLSSettings(ctx context.Context, rc *ResourceContainer, setting string, params ListHostnameTLSSettingsParams) ([]HostnameTLSSetting, ResultInfo, error) {
	if rc.Identifier == "" {
		return []HostnameTLSSetting{}, ResultInfo{}, ErrMissingZoneID
	}
	if setting == "" {
		return []HostnameTLSSetting{}, ResultInfo{}, ErrMissingHostnameTLSSettingName
	}

	uri := fmt.Sprintf("/zones/%s/hostnames/settings/%s", rc.Identifier, setting)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, params)
	if err != nil {
		return []HostnameTLSSetting{}, ResultInfo{}, err
	}
	var r HostnameTLSSettingsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []HostnameTLSSetting{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, r.ResultInfo, err
}

// EditHostnameTLSSetting will update the per-hostname tls setting for the specified hostname.
//
// API reference: https://developers.cloudflare.com/api/operations/per-hostname-tls-settings-put
func (api *API) EditHostnameTLSSetting(ctx context.Context, rc *ResourceContainer, setting, hostname string, params EditHostnameTLSSettingParams) (HostnameTLSSetting, error) {
	if rc.Identifier == "" {
		return HostnameTLSSetting{}, ErrMissingZoneID
	}
	if setting == "" {
		return HostnameTLSSetting{}, ErrMissingHostnameTLSSettingName
	}

	uri := fmt.Sprintf("/zones/%s/hostnames/settings/%s/%s", rc.Identifier, setting, hostname)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return HostnameTLSSetting{}, err
	}
	var r HostnameTLSSettingResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return HostnameTLSSetting{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteHostnameTLSSetting will delete the specified per-hostname tls setting.
//
// API reference: https://developers.cloudflare.com/api/operations/per-hostname-tls-settings-delete
func (api *API) DeleteHostnameTLSSetting(ctx context.Context, rc *ResourceContainer, setting, hostname string) (HostnameTLSSetting, error) {
	if rc.Identifier == "" {
		return HostnameTLSSetting{}, ErrMissingZoneID
	}
	if setting == "" {
		return HostnameTLSSetting{}, ErrMissingHostnameTLSSettingName
	}

	uri := fmt.Sprintf("/zones/%s/hostnames/settings/%s/%s", rc.Identifier, setting, hostname)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return HostnameTLSSetting{}, err
	}
	var r HostnameTLSSettingResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return HostnameTLSSetting{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
