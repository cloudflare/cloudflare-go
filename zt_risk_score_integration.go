package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var (
	ErrMissingRiskScoreIntegrationID = errors.New("missing required Risk Score Integration UUID")
)

// Represents a Risk Score Sharing Integration.
// If enabled, Cloudflare will share changes in user risk score with the specified vendor.
type RiskScoreIntegration struct {
	ID              string     `json:"id,omitempty"`
	IntegrationType string     `json:"integration_type,omitempty"`
	TenantUrl       string     `json:"tenant_url,omitempty"`
	ReferenceID     string     `json:"reference_id,omitempty"`
	AccountTag      string     `json:"account_tag,omitempty"`
	WellKnownUrl    string     `json:"well_known_url,omitempty"`
	Active          *bool      `json:"active,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
}

// The properties required to create a new risk score integration.
type RiskScoreIntegrationCreateRequest struct {
	IntegrationType string `json:"integration_type,omitempty"`
	TenantUrl       string `json:"tenant_url,omitempty"`
	ReferenceID     string `json:"reference_id,omitempty"`
}

// The properties required to update a risk score integration.
type RiskScoreIntegrationUpdateRequest struct {
	IntegrationType string `json:"integration_type,omitempty"`
	TenantUrl       string `json:"tenant_url,omitempty"`
	ReferenceID     string `json:"reference_id,omitempty"`
	Active          *bool  `json:"active,omitempty"`
}

// A list of risk score integrations as returned from the API.
type RiskScoreIntegrationListResponse struct {
	Result []RiskScoreIntegration `json:"result"`
	Response
}

// A risk score integration as returned from the API.
type RiskScoreIntegrationResponse struct {
	Result RiskScoreIntegration `json:"result"`
	Response
}

type ListRiskScoreIntegrationParams struct{}

// GetRiskScoreIntegration returns a single Risk Score Integration by its ID.
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/risk_scoring/subresources/integrations/methods/get/
func (api *API) GetRiskScoreIntegration(ctx context.Context, rc *ResourceContainer, integrationID string) (RiskScoreIntegration, error) {
	if rc.Identifier == "" {
		return RiskScoreIntegration{}, ErrMissingResourceIdentifier
	}

	if integrationID == "" {
		return RiskScoreIntegration{}, ErrMissingRiskScoreIntegrationID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/zt_risk_scoring/integrations/%s", rc.Level, rc.Identifier, integrationID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return RiskScoreIntegration{}, err
	}

	var integrationResponse RiskScoreIntegrationResponse
	err = json.Unmarshal(res, &integrationResponse)
	if err != nil {
		return RiskScoreIntegration{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return integrationResponse.Result, nil
}

// ListRiskScoreIntegrations returns all Risk Score Integrations for an account.
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/risk_scoring/subresources/integrations/methods/list/
func (api *API) ListRiskScoreIntegrations(ctx context.Context, rc *ResourceContainer, params ListRiskScoreIntegrationParams) ([]RiskScoreIntegration, error) {
	if rc.Identifier == "" {
		return []RiskScoreIntegration{}, ErrMissingResourceIdentifier
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/zt_risk_scoring/integrations", rc.Level, rc.Identifier), nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []RiskScoreIntegration{}, err
	}

	var integrationsResponse RiskScoreIntegrationListResponse
	err = json.Unmarshal(res, &integrationsResponse)
	if err != nil {
		return []RiskScoreIntegration{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return integrationsResponse.Result, nil
}

// CreateRiskScoreIntegration creates a new Risk Score Integration.
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/risk_scoring/subresources/integrations/methods/create/
func (api *API) CreateRiskScoreIntegration(ctx context.Context, rc *ResourceContainer, params RiskScoreIntegrationCreateRequest) (RiskScoreIntegration, error) {
	if rc.Identifier == "" {
		return RiskScoreIntegration{}, ErrMissingResourceIdentifier
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/zt_risk_scoring/integrations", rc.Level, rc.Identifier), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return RiskScoreIntegration{}, err
	}

	var integrationResponse RiskScoreIntegrationResponse
	err = json.Unmarshal(res, &integrationResponse)
	if err != nil {
		return RiskScoreIntegration{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return integrationResponse.Result, nil
}

// UpdateRiskScoreIntegration updates a Risk Score Integration.
//
// Can be used to disable or active the integration using the "active" boolean property.
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/risk_scoring/subresources/integrations/methods/update/
func (api *API) UpdateRiskScoreIntegration(ctx context.Context, rc *ResourceContainer, integrationID string, params RiskScoreIntegrationUpdateRequest) (RiskScoreIntegration, error) {
	if rc.Identifier == "" {
		return RiskScoreIntegration{}, ErrMissingResourceIdentifier
	}

	if integrationID == "" {
		return RiskScoreIntegration{}, ErrMissingRiskScoreIntegrationID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/zt_risk_scoring/integrations/%s", rc.Level, rc.Identifier, integrationID), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return RiskScoreIntegration{}, err
	}

	var integrationResponse RiskScoreIntegrationResponse
	err = json.Unmarshal(res, &integrationResponse)
	if err != nil {
		return RiskScoreIntegration{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return integrationResponse.Result, nil
}

// DeleteRiskScoreIntegration deletes a Risk Score Integration.
//
// API reference: https://developers.cloudflare.com/api/resources/zero_trust/subresources/risk_scoring/subresources/integrations/methods/delete/
func (api *API) DeleteRiskScoreIntegration(ctx context.Context, rc *ResourceContainer, integrationID string) error {
	if rc.Identifier == "" {
		return ErrMissingResourceIdentifier
	}

	if integrationID == "" {
		return ErrMissingRiskScoreIntegrationID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/zt_risk_scoring/integrations/%s", rc.Level, rc.Identifier, integrationID), nil)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	return err
}
