package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrMissingAccountLoadBalancerMonitorIdentifier = errors.New("require missing account load balancer identifier not found")

type AccountLoadBalancerMonitor struct {
	ID              string              `json:"id,omitempty"`
	CreatedOn       *time.Time          `json:"created_on,omitempty"`
	ModifiedOn      *time.Time          `json:"modified_on,omitempty"`
	Type            string              `json:"type"`
	Description     string              `json:"description"`
	Method          string              `json:"method"`
	Path            string              `json:"path"`
	Header          map[string][]string `json:"header"`
	Port            int                 `json:"port"`
	Timeout         int                 `json:"timeout"`
	Retries         int                 `json:"retries"`
	Interval        int                 `json:"interval"`
	ExpectedBody    string              `json:"expected_body"`
	ExpectedCodes   string              `json:"expected_codes"`
	FollowRedirects bool                `json:"follow_redirects"`
	AllowInsecure   bool                `json:"allow_insecure"`
	ConsecutiveUp   int                 `json:"consecutive_up,omitempty"`
	ConsecutiveDown int                 `json:"consecutive_down,omitempty"`
	ProbeZone       string              `json:"probe_zone"`
}

type PreviewAccountLoadBalancerMonitor struct {
	PreviewID string            `json:"preview_id"`
	Pools     map[string]string `json:"pools"`
}

// TODO: Check API Response
type PreviewAccountLoadBalancerMonitorResult struct {
}

type ListAccountLoadBalancerMonitorResponse struct {
	Result []AccountLoadBalancerMonitor `json:"result"`
	Response
}

type CreateAccountLoadBalancerMonitorResponse struct {
	Result AccountLoadBalancerMonitor `json:"result"`
	Response
}
type GetAccountLoadBalancerMonitorResponse struct {
	Result AccountLoadBalancerMonitor `json:"result"`
	Response
}

type UpdateAccountLoadBalancerMonitorResponse struct {
	Result AccountLoadBalancerMonitor `json:"result"`
	Response
}
type PatchAccountLoadBalancerMonitorResponse struct {
	Result AccountLoadBalancerMonitor `json:"result"`
	Response
}

type PreviewAccountLoadBalancerMonitorParameters struct {
	Identifier string `json:"-"`
	AccountLoadBalancerMonitor
}

type PreviewAccountLoadBalancerMonitorResponse struct {
	Result PreviewAccountLoadBalancerMonitor `json:"result"`
	Response
}

type PreviewAccountLoadBalancerMonitorResultResponse struct {
	Result PreviewAccountLoadBalancerMonitorResult `json:"result"`
	Response
}

// ListAccountLoadBalancerMonitors List configured monitors for an account.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-monitors-list-monitors
func (api *API) ListAccountLoadBalancerMonitors(ctx context.Context, rc *ResourceContainer) ([]AccountLoadBalancerMonitor, error) {
	if rc.Identifier == "" {
		return []AccountLoadBalancerMonitor{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/monitors", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccountLoadBalancerMonitor{}, err
	}

	var result ListAccountLoadBalancerMonitorResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return []AccountLoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// CreateAccountLoadBalancerMonitor Create a configured monitor.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-monitors-create-monitor
func (api *API) CreateAccountLoadBalancerMonitor(ctx context.Context, rc *ResourceContainer, params AccountLoadBalancerMonitor) (AccountLoadBalancerMonitor, error) {
	if rc.Identifier == "" {
		return AccountLoadBalancerMonitor{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/monitors", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return AccountLoadBalancerMonitor{}, err
	}

	var result CreateAccountLoadBalancerMonitorResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return AccountLoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// GetAccountLoadBalancerMonitor List a single configured monitor for an account.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-monitors-monitor-details
func (api *API) GetAccountLoadBalancerMonitor(ctx context.Context, rc *ResourceContainer, identifier string) (AccountLoadBalancerMonitor, error) {
	if rc.Identifier == "" {
		return AccountLoadBalancerMonitor{}, ErrMissingAccountID
	}

	if identifier == "" {
		return AccountLoadBalancerMonitor{}, ErrMissingAccountLoadBalancerMonitorIdentifier
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", rc.Identifier, identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccountLoadBalancerMonitor{}, err
	}

	var result GetAccountLoadBalancerMonitorResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return AccountLoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// UpdateAccountLoadBalancerMonitor Modify a configured monitor.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-monitors-update-monitor
func (api *API) UpdateAccountLoadBalancerMonitor(ctx context.Context, rc *ResourceContainer, params AccountLoadBalancerMonitor) (AccountLoadBalancerMonitor, error) {
	if rc.Identifier == "" {
		return AccountLoadBalancerMonitor{}, ErrMissingAccountID
	}
	if params.ID == "" {
		return AccountLoadBalancerMonitor{}, ErrMissingAccountLoadBalancerMonitorIdentifier
	}
	uri := fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return AccountLoadBalancerMonitor{}, err
	}

	var result GetAccountLoadBalancerMonitorResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return AccountLoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// PatchAccountLoadBalancerMonitor Apply changes to an existing monitor, overwriting the supplied properties.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-monitors-patch-monitor
func (api *API) PatchAccountLoadBalancerMonitor(ctx context.Context, rc *ResourceContainer, params AccountLoadBalancerMonitor) (AccountLoadBalancerMonitor, error) {
	if rc.Identifier == "" {
		return AccountLoadBalancerMonitor{}, ErrMissingAccountID
	}
	if params.ID == "" {
		return AccountLoadBalancerMonitor{}, ErrMissingAccountLoadBalancerMonitorIdentifier
	}
	uri := fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return AccountLoadBalancerMonitor{}, err
	}

	var result GetAccountLoadBalancerMonitorResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return AccountLoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// DeleteAccountLoadBalancerMonitor Delete a configured monitor.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-monitors-delete-monitor
func (api *API) DeleteAccountLoadBalancerMonitor(ctx context.Context, rc *ResourceContainer, identifier string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}
	if identifier == "" {
		return ErrMissingAccountLoadBalancerMonitorIdentifier
	}
	uri := fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", rc.Identifier, identifier)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	return err
}

// PreviewAccountLoadBalancerMonitor Preview pools using the specified monitor with provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-monitors-preview-monitor
func (api *API) PreviewAccountLoadBalancerMonitor(ctx context.Context, rc *ResourceContainer, params PreviewAccountLoadBalancerMonitorParameters) (PreviewAccountLoadBalancerMonitor, error) {
	if rc.Identifier == "" {
		return PreviewAccountLoadBalancerMonitor{}, ErrMissingAccountID
	}
	if params.Identifier == "" {
		return PreviewAccountLoadBalancerMonitor{}, ErrMissingAccountLoadBalancerMonitorIdentifier
	}
	uri := fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s/preview", rc.Identifier, params.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PreviewAccountLoadBalancerMonitor{}, err
	}

	var result PreviewAccountLoadBalancerMonitorResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return PreviewAccountLoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// PreviewAccountLoadBalancerResult Get the result of a previous preview operation using the provided preview_id.
// TODO: Verify API RESPONSE
// API reference: https://api.cloudflare.com/#account-load-balancer-monitors-preview-monitor
func (api *API) PreviewAccountLoadBalancerResult(ctx context.Context, rc *ResourceContainer, previewID string) (PreviewAccountLoadBalancerMonitorResult, error) {
	if rc.Identifier == "" {
		return PreviewAccountLoadBalancerMonitorResult{}, ErrMissingAccountID
	}
	if previewID == "" {
		return PreviewAccountLoadBalancerMonitorResult{}, errors.New("require previewID is missing")
	}
	uri := fmt.Sprintf("/accounts/%s/load_balancers/monitors/preview/%s", rc.Identifier, previewID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PreviewAccountLoadBalancerMonitorResult{}, err
	}

	var result PreviewAccountLoadBalancerMonitorResultResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return PreviewAccountLoadBalancerMonitorResult{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}
