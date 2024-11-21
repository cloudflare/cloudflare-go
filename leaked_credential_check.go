package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type LeakedCredentialCheckGetStatusParams struct{}

type LeakedCredentialCheckStatus struct {
	Enabled *bool `json:"enabled"`
}

type LeakCredentialCheckStatusResponse struct {
	Response
	Result LeakedCredentialCheckStatus `json:"result"`
}

type LeakCredentialCheckSetStatusParams struct {
	Enabled *bool `json:"enabled"`
}

type LeakedCredentialCheckListDetectionsParams struct{}

type LeakedCredentialCheckDetectionEntry struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LeakedCredentialCheckListDetectionsResponse struct {
	Response
	Result []LeakedCredentialCheckDetectionEntry `json:"result"`
}

type LeakedCredentialCheckCreateDetectionParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LeakedCredentialCheckCreateDetectionResponse struct {
	Response
	Result LeakedCredentialCheckDetectionEntry `json:"result"`
}

type LeakedCredentialCheckDeleteDetectionParams struct {
	DetectionID string
}

type LeakedCredentialCheckDeleteDetectionResponse struct {
	Response
	Result []struct{} `json:"result"`
}

type LeakedCredentialCheckUpdateDetectionParams struct {
	LeakedCredentialCheckDetectionEntry
}
type LeakedCredentialCheckUpdateDetectionResponse struct {
	Response
	Result LeakedCredentialCheckDetectionEntry
}

// LeakCredentialCheckGetStatus returns whether Leaked credential check is enabled or not. It is false by default.
//
// API reference: https://developers.cloudflare.com/api/operations/waf-product-api-leaked-credentials-get-status
func (api *API) LeakedCredentialCheckGetStatus(ctx context.Context, rc *ResourceContainer, params LeakedCredentialCheckGetStatusParams) (LeakedCredentialCheckStatus, error) {
	if rc.Identifier == "" {
		return LeakedCredentialCheckStatus{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/leaked-credential-checks", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LeakedCredentialCheckStatus{}, err
	}
	result := LeakCredentialCheckStatusResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return LeakedCredentialCheckStatus{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// LeakedCredentialCheckSetStatus enable or disable the Leak Credential Check. Returns the status.
//
// API reference: https://developers.cloudflare.com/api/operations/waf-product-api-leaked-credentials-set-status
func (api *API) LeakedCredentialCheckSetStatus(ctx context.Context, rc *ResourceContainer, params LeakCredentialCheckSetStatusParams) (LeakedCredentialCheckStatus, error) {
	if rc.Identifier == "" {
		return LeakedCredentialCheckStatus{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/leaked-credential-checks", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return LeakedCredentialCheckStatus{}, err
	}
	result := LeakCredentialCheckStatusResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return LeakedCredentialCheckStatus{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// LeakedCredentialCheckListDetections lists user-defined detection patterns for Leaked Credential Checks.
//
// API reference: https://developers.cloudflare.com/api/operations/waf-product-api-leaked-credentials-list-detections
func (api *API) LeakedCredentialCheckListDetections(ctx context.Context, rc *ResourceContainer, params LeakedCredentialCheckListDetectionsParams) ([]LeakedCredentialCheckDetectionEntry, error) {
	if rc.Identifier == "" {
		return []LeakedCredentialCheckDetectionEntry{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/leaked-credential-checks/detections", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, params)
	if err != nil {
		return []LeakedCredentialCheckDetectionEntry{}, err
	}
	result := LeakedCredentialCheckListDetectionsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []LeakedCredentialCheckDetectionEntry{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// LeakedCredentialCheckCreateDetection creates user-defined detection pattern for Leaked Credential Checks
//
// API reference: https://developers.cloudflare.com/api/operations/waf-product-api-leaked-credentials-create-detection
func (api *API) LeakedCredentialCheckCreateDetection(ctx context.Context, rc *ResourceContainer, params LeakedCredentialCheckCreateDetectionParams) (LeakedCredentialCheckDetectionEntry, error) {
	if rc.Identifier == "" {
		return LeakedCredentialCheckDetectionEntry{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/leaked-credential-checks/detections", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return LeakedCredentialCheckDetectionEntry{}, err
	}
	result := LeakedCredentialCheckCreateDetectionResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return LeakedCredentialCheckDetectionEntry{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// LeakedCredentialCheckDeleteDetection removes user-defined detection pattern for Leaked Credential Checks
//
// API reference: https://developers.cloudflare.com/api/operations/waf-product-api-leaked-credentials-delete-detection
func (api *API) LeakedCredentialCheckDeleteDetection(ctx context.Context, rc *ResourceContainer, params LeakedCredentialCheckDeleteDetectionParams) (LeakedCredentialCheckDeleteDetectionResponse, error) {
	if rc.Identifier == "" {
		return LeakedCredentialCheckDeleteDetectionResponse{}, ErrMissingZoneID
	}
	if params.DetectionID == "" {
		return LeakedCredentialCheckDeleteDetectionResponse{}, ErrMissingDetectionID
	}

	uri := fmt.Sprintf("/zones/%s/leaked-credential-checks/detections/%s", rc.Identifier, params.DetectionID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return LeakedCredentialCheckDeleteDetectionResponse{}, err
	}
	result := LeakedCredentialCheckDeleteDetectionResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return LeakedCredentialCheckDeleteDetectionResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result, nil
}

// LeakedCredentialCheckUpdateDetection updates user-defined detection pattern for Leaked Credential Checks. Returns updated detection.
//
// API reference: https://developers.cloudflare.com/api/operations/waf-product-api-leaked-credentials-update-detection
func (api *API) LeakedCredentialCheckUpdateDetection(ctx context.Context, rc *ResourceContainer, params LeakedCredentialCheckUpdateDetectionParams) (LeakedCredentialCheckDetectionEntry, error) {
	if rc.Identifier == "" {
		return LeakedCredentialCheckDetectionEntry{}, ErrMissingZoneID
	}
	if params.ID == "" {
		return LeakedCredentialCheckDetectionEntry{}, ErrMissingDetectionID
	}

	uri := fmt.Sprintf("/zones/%s/leaked-credential-checks/detections/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return LeakedCredentialCheckDetectionEntry{}, err
	}
	result := LeakedCredentialCheckUpdateDetectionResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return LeakedCredentialCheckDetectionEntry{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}
