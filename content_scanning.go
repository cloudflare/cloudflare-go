package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type ContentScanningStatusResult struct {
	Value    string `json:"value"`
	Modified string `json:"modified"`
}

type ContentScanningEnableParams struct{}

type ContentScanningEnableResponse struct {
	Response
	Result *ContentScanningStatusResult `json:"result"` // nil
}

type ContentScanningDisableParams struct{}

type ContentScanningDisableResponse struct {
	Response
	Result *ContentScanningStatusResult `json:"result"` // nil
}

type ContentScanningStatusParams struct{}

type ContentScanningStatusResponse struct {
	Response
	Result *ContentScanningStatusResult `json:"result"` // object or nil
}

type ContentScanningCustomExpression struct {
	ID      string `json:"id"`
	Payload string `json:"payload"`
}

type ContentScanningListCustomExpressionsParams struct{}

type ContentScanningListCustomExpressionsResponse struct {
	Response
	Result []ContentScanningCustomExpression `json:"result"`
}

type ContentScanningCustomPayload struct {
	Payload string `json:"payload"`
}

type ContentScanningAddCustomExpressionsParams struct {
	Payloads []ContentScanningCustomPayload
}

type ContentScanningAddCustomExpressionsResponse struct {
	Response
	Result []ContentScanningCustomExpression `json:"result"`
}

type ContentScanningDeleteCustomExpressionsParams struct {
	ID string
}

type ContentScanningDeleteCustomExpressionsResponse struct {
	Response
	Result []ContentScanningCustomExpression `json:"result"`
}

// ContentScanningEnable enables Content Scanning.
//
// API reference: https://developers.cloudflare.com/api/resources/content_scanning/methods/enable/
func (api *API) ContentScanningEnable(ctx context.Context, rc *ResourceContainer, params ContentScanningEnableParams) (ContentScanningEnableResponse, error) {
	if rc.Identifier == "" {
		return ContentScanningEnableResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/content-upload-scan/enable", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return ContentScanningEnableResponse{}, err
	}
	result := ContentScanningEnableResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ContentScanningEnableResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result, nil
}

// ContentScanningDisable disables Content Scanning.
//
// API reference: https://developers.cloudflare.com/api/resources/content_scanning/methods/disable/
func (api *API) ContentScanningDisable(ctx context.Context, rc *ResourceContainer, params ContentScanningDisableParams) (ContentScanningDisableResponse, error) {
	if rc.Identifier == "" {
		return ContentScanningDisableResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/content-upload-scan/disable", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return ContentScanningDisableResponse{}, err
	}
	result := ContentScanningDisableResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ContentScanningDisableResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result, nil
}

// ContentScanningStatus reads the status of Content Scanning.
//
// API reference: https://developers.cloudflare.com/api/resources/content_scanning/subresources/settings/methods/get/
func (api *API) ContentScanningStatus(ctx context.Context, rc *ResourceContainer, params ContentScanningStatusParams) (ContentScanningStatusResponse, error) {
	if rc.Identifier == "" {
		return ContentScanningStatusResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/content-upload-scan/settings", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ContentScanningStatusResponse{}, err
	}
	result := ContentScanningStatusResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ContentScanningStatusResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result, nil
}

// ContentScanningListCustomExpressions retrieves the list of existing custom scan expressions for Content Scanning
//
// API reference: https://developers.cloudflare.com/api/resources/content_scanning/subresources/payloads/methods/list/
func (api *API) ContentScanningListCustomExpressions(ctx context.Context, rc *ResourceContainer, params ContentScanningListCustomExpressionsParams) ([]ContentScanningCustomExpression, error) {
	if rc.Identifier == "" {
		return []ContentScanningCustomExpression{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/content-upload-scan/payloads", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []ContentScanningCustomExpression{}, err
	}
	result := ContentScanningListCustomExpressionsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []ContentScanningCustomExpression{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// ContentScanningAddCustomExpressions add new custom scan expressions for Content Scanning
//
// API reference: https://developers.cloudflare.com/api/resources/content_scanning/subresources/payloads/methods/list/
func (api *API) ContentScanningAddCustomExpressions(ctx context.Context, rc *ResourceContainer, params ContentScanningAddCustomExpressionsParams) ([]ContentScanningCustomExpression, error) {
	if rc.Identifier == "" {
		return []ContentScanningCustomExpression{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/content-upload-scan/payloads", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.Payloads)
	if err != nil {
		return []ContentScanningCustomExpression{}, err
	}
	result := ContentScanningAddCustomExpressionsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []ContentScanningCustomExpression{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// ContentScanningDeleteCustomExpressions delete custom scan expressions for Content Scanning
//
// API reference: https://developers.cloudflare.com/api/resources/content_scanning/subresources/payloads/methods/delete/
func (api *API) ContentScanningDeleteCustomExpression(ctx context.Context, rc *ResourceContainer, params ContentScanningDeleteCustomExpressionsParams) ([]ContentScanningCustomExpression, error) {
	if rc.Identifier == "" {
		return []ContentScanningCustomExpression{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/content-upload-scan/payloads/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return []ContentScanningCustomExpression{}, err
	}
	result := ContentScanningDeleteCustomExpressionsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []ContentScanningCustomExpression{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}
