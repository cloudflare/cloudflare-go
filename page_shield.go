package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// PageShield represents the page shield object minus any timestamps
type PageShield struct {
	Enabled                        bool `json:"enabled"`
	UseCloudflareReportingEndpoint bool `json:"use_cloudflare_reporting_endpoint"`
	UseConnectionURLPath           bool `json:"use_connection_url_path"`
}

// PageShieldSettings represents the page shield settings for a zone
type PageShieldSettings struct {
	PageShield
	UpdatedAt string `json:"updated_at"`
}

// PageShieldSettingsResponse represents the response from the page shield settings endpoint
type PageShieldSettingsResponse struct {
	PageShield PageShieldSettings `json:"result"`
	Response
}

// GetPageShieldSettings returns the page shield settings for a zone
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-get-page-shield-settings
func (api *API) GetPageShieldSettings(ctx context.Context, rc *ResourceContainer) (*PageShieldSettingsResponse, error) {
	uri := fmt.Sprintf("/zones/%s/page_shield", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var psResponse PageShieldSettingsResponse
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &psResponse, nil
}

// UpdatePageShieldSettings updates the page shield settings for a zone
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-update-page-shield-settings
func (api *API) UpdatePageShieldSettings(ctx context.Context, rc *ResourceContainer, params PageShield) (*PageShieldSettingsResponse, error) {
	uri := fmt.Sprintf("/zones/%s/page_shield", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return nil, err
	}

	var psResponse PageShieldSettingsResponse
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &psResponse, nil
}
