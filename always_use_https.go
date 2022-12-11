package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AlwaysUseHTTPS struct {
	ID         string    `json:"id,omitempty"`
	Value      string    `json:"value,omitempty"`
	Editable   bool      `json:"editable,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}

type AlwaysUseHTTPSResponse struct {
	Response
	Result AlwaysUseHTTPS `json:"result"`
}

// GetAlwaysUseHTTPS Get Always Use HTTPS Settings for a Zone.
//
// API Reference: https://api.cloudflare.com/#zone-settings-get-always-use-https-setting
func (api *API) GetAlwaysUseHTTPS(ctx context.Context, zoneID string) (AlwaysUseHTTPS, error) {
	uri := fmt.Sprintf("/zones/%s/settings/always_use_https", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AlwaysUseHTTPS{}, err
	}

	var r AlwaysUseHTTPSResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return AlwaysUseHTTPS{}, err
	}
	return r.Result, nil
}

// SetAlwaysUseHTTPS Set Always Use HTTPS Settings for a Zone.
//
// API Reference: https://api.cloudflare.com/#zone-settings-change-always-use-https-setting
func (api *API) SetAlwaysUseHTTPS(ctx context.Context, zoneID string, enabled bool) (AlwaysUseHTTPS, error) {
	params := AlwaysUseHTTPS{Value: "off"}
	if enabled {
		params.Value = "on"
	}

	uri := fmt.Sprintf("/zones/%s/settings/always_use_https", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return AlwaysUseHTTPS{}, err
	}

	var r AlwaysUseHTTPSResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return AlwaysUseHTTPS{}, err
	}
	return r.Result, nil
}
