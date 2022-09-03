package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var ErrMissingAccountLoadBalancerPoolID = errors.New("require missing account load balancer pool ID")

// AccountLoadBalancerPool represents a load balancer pool's properties.
type AccountLoadBalancerPool struct {
	ID                 string                                 `json:"id,omitempty"`
	CreatedOn          *time.Time                             `json:"created_on,omitempty"`
	ModifiedOn         *time.Time                             `json:"modified_on,omitempty"`
	Description        string                                 `json:"description"`
	Name               string                                 `json:"name"`
	Enabled            bool                                   `json:"enabled"`
	MinimumOrigins     int                                    `json:"minimum_origins,omitempty"`
	Monitor            string                                 `json:"monitor,omitempty"`
	Origins            []AccountLoadBalancerOrigin            `json:"origins"`
	NotificationEmail  string                                 `json:"notification_email,omitempty"`
	NotificationFilter *AccountLoadBalancerNotificationFilter `json:"notification_filter,omitempty"`
	// ToDO: Figure out how to handle these.
	//Latitude           *float32                               `json:"latitude,omitempty"`
	//Longitude          *float32                               `json:"longitude,omitempty"`
	LoadShedding   *AccountLoadBalancerPoolLoadShedding `json:"load_shedding,omitempty"`
	OriginSteering *AccountLoadBalancerOriginSteering   `json:"origin_steering,omitempty"`

	// CheckRegions defines the geographic region(s) from where to run health-checks from - e.g. "WNAM", "WEU", "SAF", "SAM".
	// Providing a null/empty value means "all regions", which may not be available to all plan types.
	CheckRegions []string `json:"check_regions"`
}

// AccountLoadBalancerOrigin represents a Load Balancer origin's properties.
type AccountLoadBalancerOrigin struct {
	Name    string              `json:"name"`
	Address string              `json:"address"`
	Enabled bool                `json:"enabled"`
	Weight  float64             `json:"weight"`
	Header  map[string][]string `json:"header"`
}

// AccountLoadBalancerOriginSteering controls origin selection for new sessions and traffic without session affinity.
type AccountLoadBalancerOriginSteering struct {
	// Policy defaults to "random" (weighted) when empty or unspecified.
	Policy string `json:"policy,omitempty"`
}

// AccountLoadBalancerPoolLoadShedding contains the settings for controlling load shedding.
type AccountLoadBalancerPoolLoadShedding struct {
	DefaultPercent float32 `json:"default_percent"`
	DefaultPolicy  string  `json:"default_policy"`
	SessionPercent float32 `json:"session_percent"`
	SessionPolicy  string  `json:"session_policy"`
}

type AccountLoadBalancerNotificationFilter struct {
	Origin *AccountLoadBalancerNotificationResourceType `json:"origin"`
	Pool   *AccountLoadBalancerNotificationResourceType `json:"pool"`
}

// AccountLoadBalancerPoolHealth represents the healthchecks from different PoPs for a pool.
type AccountLoadBalancerPoolHealth struct {
	ID        string                                      `json:"pool_id,omitempty"`
	PopHealth map[string]AccountLoadBalancerPoolPopHealth `json:"pop_health,omitempty"`
}

// AccountLoadBalancerPoolPopHealth represents the health of the pool for given PoP.
type AccountLoadBalancerPoolPopHealth struct {
	Healthy bool                                         `json:"healthy,omitempty"`
	Origins []map[string]AccountLoadBalancerOriginHealth `json:"origins,omitempty"`
}

// AccountLoadBalancerOriginHealth represents the health of the origin.
type AccountLoadBalancerOriginHealth struct {
	Healthy       bool     `json:"healthy,omitempty"`
	RTT           Duration `json:"rtt,omitempty"`
	FailureReason string   `json:"failure_reason,omitempty"`
	ResponseCode  int      `json:"response_code,omitempty"`
}

type AccountLoadBalancerNotificationResourceType struct {
	Disable bool        `json:"disable"`
	Healthy interface{} `json:"healthy"`
}

type AccountLoadBalancerPoolPreviewParameters struct {
	ID              string              `json:"-"`
	Port            int                 `json:"port"`
	Method          string              `json:"method"`
	Timeout         int                 `json:"timeout"`
	Path            string              `json:"path"`
	Retries         int                 `json:"retries"`
	FollowRedirects bool                `json:"follow_redirects"`
	ExpectedBody    string              `json:"expected_body"`
	ExpectedCodes   string              `json:"expected_codes"`
	Header          map[string][]string `json:"header"`
	AllowInsecure   bool                `json:"allow_insecure"`
	Type            string              `json:"type"`
	ProbeZone       string              `json:"probe_zone"`
}

type AccountLoadBalancerPoolPreview struct {
	PreviewID string            `json:"preview_id"`
	Pools     map[string]string `json:"pools"`
}

type AccountLoadBalancerPoolReferences struct {
	ResourceID    string `json:"resource_id"`
	ResourceName  string `json:"resource_name"`
	ResourceType  string `json:"resource_type"`
	ReferenceType string `json:"reference_type"`
}

// AccountLoadBalancerPoolResponse represents the response from the load balancer pool endpoints.
type AccountLoadBalancerPoolResponse struct {
	Result AccountLoadBalancerPool `json:"result"`
	Response
}

// ListAccountLoadBalancerPoolResponse represents the response from the load balancer pool endpoints.
type ListAccountLoadBalancerPoolResponse struct {
	Result []AccountLoadBalancerPool `json:"result"`
	Response
}

type GetAccountLoadBalancerPoolResponse struct {
	Result AccountLoadBalancerPool `json:"result"`
	Response
}

type GetAccountLoadBalancerPoolHealthResponse struct {
	Result AccountLoadBalancerPoolHealth `json:"result"`
	Response
}

type PreviewAccountLoadBalancerPoolResponse struct {
	Result AccountLoadBalancerPoolPreview `json:"result"`
	Response
}

type ListAccountLoadBalancerPoolReferencesResponse struct {
	Result []AccountLoadBalancerPoolReferences `json:"result"`
	Response
}

// CreateAccountLoadBalancerPool creates a new account load balancer pool.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-pools-create-pool
func (api *API) CreateAccountLoadBalancerPool(ctx context.Context, rc *ResourceContainer, pool AccountLoadBalancerPool) (AccountLoadBalancerPool, error) {
	if rc.Identifier == "" {
		return AccountLoadBalancerPool{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/pools", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, pool)
	if err != nil {
		return AccountLoadBalancerPool{}, err
	}
	var r AccountLoadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AccountLoadBalancerPool{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListAccountLoadBalancerPools lists load balancer pools connected to an account.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-pools-list-pools
func (api *API) ListAccountLoadBalancerPools(ctx context.Context, rc *ResourceContainer, monitor string) ([]AccountLoadBalancerPool, error) {
	if rc.Identifier == "" {
		return []AccountLoadBalancerPool{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/load_balancers/pools", rc.Identifier), monitor)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var r ListAccountLoadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// GetAccountLoadBalancerPool Fetch a single configured pool.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-pools-pool-details
func (api *API) GetAccountLoadBalancerPool(ctx context.Context, rc *ResourceContainer, poolID string) (AccountLoadBalancerPool, error) {
	if rc.Identifier == "" {
		return AccountLoadBalancerPool{}, ErrMissingAccountID
	}
	if poolID == "" {
		return AccountLoadBalancerPool{}, ErrMissingAccountLoadBalancerPoolID
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", rc.Identifier, poolID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccountLoadBalancerPool{}, err
	}
	var r GetAccountLoadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AccountLoadBalancerPool{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// GetAccountLoadBalancerPoolHeath Fetch the latest pool health status for a single pool.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-pools-pool-health-details
func (api *API) GetAccountLoadBalancerPoolHeath(ctx context.Context, rc *ResourceContainer, poolID string) (AccountLoadBalancerPoolHealth, error) {
	if rc.Identifier == "" {
		return AccountLoadBalancerPoolHealth{}, ErrMissingAccountID
	}
	if poolID == "" {
		return AccountLoadBalancerPoolHealth{}, ErrMissingAccountLoadBalancerPoolID
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/pools/%s/health", rc.Identifier, poolID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccountLoadBalancerPoolHealth{}, err
	}
	var r GetAccountLoadBalancerPoolHealthResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AccountLoadBalancerPoolHealth{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateAccountLoadBalancerPool modifies a configured load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-update-pool
func (api *API) UpdateAccountLoadBalancerPool(ctx context.Context, rc *ResourceContainer, pool AccountLoadBalancerPool) (AccountLoadBalancerPool, error) {
	if rc.Identifier == "" {
		return AccountLoadBalancerPool{}, ErrMissingAccountID
	}
	if pool.ID == "" {
		return AccountLoadBalancerPool{}, ErrMissingAccountLoadBalancerPoolID
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", rc.Identifier, pool.ID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, pool)
	if err != nil {
		return AccountLoadBalancerPool{}, err
	}

	var r AccountLoadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AccountLoadBalancerPool{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// PatchAccountLoadBalancerPool Apply changes to an existing pool, overwriting the supplied properties.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-pools-patch-pool
func (api *API) PatchAccountLoadBalancerPool(ctx context.Context, rc *ResourceContainer, params AccountLoadBalancerPool) (AccountLoadBalancerPool, error) {
	if rc.Identifier == "" {
		return AccountLoadBalancerPool{}, ErrMissingAccountID
	}

	if params.ID == "" {
		return AccountLoadBalancerPool{}, ErrMissingAccountLoadBalancerPoolID
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return AccountLoadBalancerPool{}, err
	}

	var result AccountLoadBalancerPoolResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return AccountLoadBalancerPool{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// DeleteAccountLoadBalancerPool Delete a configured pool.
//
// API reference: https://api.cloudflare.com/#account-load-balancer-pools-delete-pool
func (api *API) DeleteAccountLoadBalancerPool(ctx context.Context, rc *ResourceContainer, poolID string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}
	if poolID == "" {
		return ErrMissingAccountLoadBalancerPoolID
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", rc.Identifier, poolID)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	return err
}

// PreviewAccountLoadBalancerPool Preview pool health using provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results..
//
// API reference: https://api.cloudflare.com/#account-load-balancer-pools-preview-pool
func (api *API) PreviewAccountLoadBalancerPool(ctx context.Context, rc *ResourceContainer, params AccountLoadBalancerPoolPreviewParameters) (AccountLoadBalancerPoolPreview, error) {
	if rc.Identifier == "" {
		return AccountLoadBalancerPoolPreview{}, ErrMissingAccountID
	}
	if params.ID == "" {
		return AccountLoadBalancerPoolPreview{}, ErrMissingAccountLoadBalancerPoolID
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/pools/%s/preview", rc.Identifier, params.ID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return AccountLoadBalancerPoolPreview{}, err
	}
	var r PreviewAccountLoadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AccountLoadBalancerPoolPreview{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListAccountLoadBalancerPoolReferences Preview pool health using provided monitor details. The returned preview_id can be used in the preview endpoint to retrieve the results..
//
// API reference: https://api.cloudflare.com/#account-load-balancer-pools-list-pool-references
func (api *API) ListAccountLoadBalancerPoolReferences(ctx context.Context, rc *ResourceContainer, poolID string) ([]AccountLoadBalancerPoolReferences, error) {
	if rc.Identifier == "" {
		return []AccountLoadBalancerPoolReferences{}, ErrMissingAccountID
	}
	if poolID == "" {
		return []AccountLoadBalancerPoolReferences{}, ErrMissingAccountLoadBalancerPoolID
	}

	uri := fmt.Sprintf("/accounts/%s/load_balancers/pools/%s/references", rc.Identifier, poolID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return []AccountLoadBalancerPoolReferences{}, err
	}
	var r ListAccountLoadBalancerPoolReferencesResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []AccountLoadBalancerPoolReferences{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
