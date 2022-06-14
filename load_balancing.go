package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// LoadBalancerPool represents a load balancer pool's properties.
type LoadBalancerPool struct {
	ID                string                      `json:"id,omitempty"`
	CreatedOn         *time.Time                  `json:"created_on,omitempty"`
	ModifiedOn        *time.Time                  `json:"modified_on,omitempty"`
	Description       string                      `json:"description"`
	Name              string                      `json:"name"`
	Enabled           bool                        `json:"enabled"`
	MinimumOrigins    int                         `json:"minimum_origins,omitempty"`
	Monitor           string                      `json:"monitor,omitempty"`
	Origins           []LoadBalancerOrigin        `json:"origins"`
	NotificationEmail string                      `json:"notification_email,omitempty"`
	Latitude          *float32                    `json:"latitude,omitempty"`
	Longitude         *float32                    `json:"longitude,omitempty"`
	LoadShedding      *LoadBalancerLoadShedding   `json:"load_shedding,omitempty"`
	OriginSteering    *LoadBalancerOriginSteering `json:"origin_steering,omitempty"`

	// CheckRegions defines the geographic region(s) from where to run health-checks from - e.g. "WNAM", "WEU", "SAF", "SAM".
	// Providing a null/empty value means "all regions", which may not be available to all plan types.
	CheckRegions []string `json:"check_regions"`
}

// LoadBalancerOrigin represents a Load Balancer origin's properties.
type LoadBalancerOrigin struct {
	Name    string              `json:"name"`
	Address string              `json:"address"`
	Enabled bool                `json:"enabled"`
	Weight  float64             `json:"weight"`
	Header  map[string][]string `json:"header"`
}

// LoadBalancerOriginSteering controls origin selection for new sessions and traffic without session affinity.
type LoadBalancerOriginSteering struct {
	// Policy defaults to "random" (weighted) when empty or unspecified.
	Policy string `json:"policy,omitempty"`
}

// LoadBalancerMonitor represents a load balancer monitor's properties.
type LoadBalancerMonitor struct {
	ID              string              `json:"id,omitempty"`
	CreatedOn       *time.Time          `json:"created_on,omitempty"`
	ModifiedOn      *time.Time          `json:"modified_on,omitempty"`
	Type            string              `json:"type"`
	Description     string              `json:"description"`
	Method          string              `json:"method"`
	Path            string              `json:"path"`
	Header          map[string][]string `json:"header"`
	Timeout         int                 `json:"timeout"`
	Retries         int                 `json:"retries"`
	Interval        int                 `json:"interval"`
	Port            uint16              `json:"port,omitempty"`
	ExpectedBody    string              `json:"expected_body"`
	ExpectedCodes   string              `json:"expected_codes"`
	FollowRedirects bool                `json:"follow_redirects"`
	AllowInsecure   bool                `json:"allow_insecure"`
	ProbeZone       string              `json:"probe_zone"`
}

// LoadBalancer represents a load balancer's properties.
type LoadBalancer struct {
	ID                        string                     `json:"id,omitempty"`
	CreatedOn                 *time.Time                 `json:"created_on,omitempty"`
	ModifiedOn                *time.Time                 `json:"modified_on,omitempty"`
	Description               string                     `json:"description"`
	Name                      string                     `json:"name"`
	TTL                       int                        `json:"ttl,omitempty"`
	FallbackPool              string                     `json:"fallback_pool"`
	DefaultPools              []string                   `json:"default_pools"`
	RegionPools               map[string][]string        `json:"region_pools"`
	PopPools                  map[string][]string        `json:"pop_pools"`
	CountryPools              map[string][]string        `json:"country_pools"`
	Proxied                   bool                       `json:"proxied"`
	Enabled                   *bool                      `json:"enabled,omitempty"`
	Persistence               string                     `json:"session_affinity,omitempty"`
	PersistenceTTL            int                        `json:"session_affinity_ttl,omitempty"`
	SessionAffinityAttributes *SessionAffinityAttributes `json:"session_affinity_attributes,omitempty"`
	Rules                     []*LoadBalancerRule        `json:"rules,omitempty"`
	RandomSteering            *RandomSteering            `json:"random_steering,omitempty"`

	// SteeringPolicy controls pool selection logic.
	// "off" select pools in DefaultPools order
	// "geo" select pools based on RegionPools/PopPools/CountryPools
	// "dynamic_latency" select pools based on RTT (requires health checks)
	// "random" selects pools in a random order
	// "proximity" select pools based on 'distance' from request
	// "" maps to "geo" if RegionPools or PopPools or CountryPools have entries otherwise "off"
	SteeringPolicy string `json:"steering_policy,omitempty"`
}

// LoadBalancerLoadShedding contains the settings for controlling load shedding.
type LoadBalancerLoadShedding struct {
	DefaultPercent float32 `json:"default_percent,omitempty"`
	DefaultPolicy  string  `json:"default_policy,omitempty"`
	SessionPercent float32 `json:"session_percent,omitempty"`
	SessionPolicy  string  `json:"session_policy,omitempty"`
}

// LoadBalancerRule represents a single rule entry for a Load Balancer. Each rules
// is run one after the other in priority order. Disabled rules are skipped.
type LoadBalancerRule struct {
	Overrides LoadBalancerRuleOverrides `json:"overrides"`

	// Name is required but is only used for human readability
	Name string `json:"name"`

	Condition string `json:"condition"`

	// Priority controls the order of rule execution the lowest value will be invoked first
	Priority int `json:"priority"`

	// FixedResponse if set and the condition is true we will not run
	// routing logic but rather directly respond with the provided fields.
	// FixedResponse implies terminates.
	FixedResponse *LoadBalancerFixedResponseData `json:"fixed_response,omitempty"`

	Disabled bool `json:"disabled"`

	// Terminates flag this rule as 'terminating'. No further rules will
	// be executed after this one.
	Terminates bool `json:"terminates,omitempty"`
}

// LoadBalancerFixedResponseData contains all the data needed to generate
// a fixed response from a Load Balancer. This behavior can be enabled via Rules.
type LoadBalancerFixedResponseData struct {
	// MessageBody data to write into the http body
	MessageBody string `json:"message_body,omitempty"`
	// StatusCode the http status code to response with
	StatusCode int `json:"status_code,omitempty"`
	// ContentType value of the http 'content-type' header
	ContentType string `json:"content_type,omitempty"`
	// Location value of the http 'location' header
	Location string `json:"location,omitempty"`
}

// LoadBalancerRuleOverrides are the set of field overridable by the rules system.
type LoadBalancerRuleOverrides struct {
	// session affinity
	Persistence    string `json:"session_affinity,omitempty"`
	PersistenceTTL *uint  `json:"session_affinity_ttl,omitempty"`

	SessionAffinityAttrs *LoadBalancerRuleOverridesSessionAffinityAttrs `json:"session_affinity_attributes,omitempty"`

	TTL uint `json:"ttl,omitempty"`

	SteeringPolicy string `json:"steering_policy,omitempty"`
	FallbackPool   string `json:"fallback_pool,omitempty"`

	DefaultPools []string            `json:"default_pools,omitempty"`
	PoPPools     map[string][]string `json:"pop_pools,omitempty"`
	RegionPools  map[string][]string `json:"region_pools,omitempty"`
	CountryPools map[string][]string `json:"country_pools,omitempty"`

	RandomSteering *RandomSteering `json:"random_steering,omitempty"`
}

// RandomSteering represents fields used to set pool weights on a load balancer
// with "random" steering policy.
type RandomSteering struct {
	DefaultWeight float64            `json:"default_weight,omitempty"`
	PoolWeights   map[string]float64 `json:"pool_weights,omitempty"`
}

// LoadBalancerRuleOverridesSessionAffinityAttrs mimics SessionAffinityAttributes without the
// DrainDuration field as that field can not be overwritten via rules.
type LoadBalancerRuleOverridesSessionAffinityAttrs struct {
	SameSite             string `json:"samesite,omitempty"`
	Secure               string `json:"secure,omitempty"`
	ZeroDowntimeFailover string `json:"zero_downtime_failover,omitempty"`
}

// SessionAffinityAttributes represents the fields used to set attributes in a load balancer session affinity cookie.
type SessionAffinityAttributes struct {
	SameSite             string `json:"samesite,omitempty"`
	Secure               string `json:"secure,omitempty"`
	DrainDuration        int    `json:"drain_duration,omitempty"`
	ZeroDowntimeFailover string `json:"zero_downtime_failover,omitempty"`
}

// LoadBalancerOriginHealth represents the health of the origin.
type LoadBalancerOriginHealth struct {
	Healthy       bool     `json:"healthy,omitempty"`
	RTT           Duration `json:"rtt,omitempty"`
	FailureReason string   `json:"failure_reason,omitempty"`
	ResponseCode  int      `json:"response_code,omitempty"`
}

// LoadBalancerPoolPopHealth represents the health of the pool for given PoP.
type LoadBalancerPoolPopHealth struct {
	Healthy bool                                  `json:"healthy,omitempty"`
	Origins []map[string]LoadBalancerOriginHealth `json:"origins,omitempty"`
}

// LoadBalancerPoolHealth represents the healthchecks from different PoPs for a pool.
type LoadBalancerPoolHealth struct {
	ID        string                               `json:"pool_id,omitempty"`
	PopHealth map[string]LoadBalancerPoolPopHealth `json:"pop_health,omitempty"`
}

// loadBalancerPoolResponse represents the response from the load balancer pool endpoints.
type loadBalancerPoolResponse struct {
	Response
	Result LoadBalancerPool `json:"result"`
}

// loadBalancerPoolListResponse represents the response from the List Pools endpoint.
type loadBalancerPoolListResponse struct {
	Response
	Result     []LoadBalancerPool `json:"result"`
	ResultInfo ResultInfo         `json:"result_info"`
}

// loadBalancerMonitorResponse represents the response from the load balancer monitor endpoints.
type loadBalancerMonitorResponse struct {
	Response
	Result LoadBalancerMonitor `json:"result"`
}

// loadBalancerMonitorListResponse represents the response from the List Monitors endpoint.
type loadBalancerMonitorListResponse struct {
	Response
	Result     []LoadBalancerMonitor `json:"result"`
	ResultInfo ResultInfo            `json:"result_info"`
}

// loadBalancerResponse represents the response from the load balancer endpoints.
type loadBalancerResponse struct {
	Response
	Result LoadBalancer `json:"result"`
}

// loadBalancerListResponse represents the response from the List Load Balancers endpoint.
type loadBalancerListResponse struct {
	Response
	Result     []LoadBalancer `json:"result"`
	ResultInfo ResultInfo     `json:"result_info"`
}

// loadBalancerPoolHealthResponse represents the response from the Pool Health Details endpoint.
type loadBalancerPoolHealthResponse struct {
	Response
	Result LoadBalancerPoolHealth `json:"result"`
}

type CreateLoadBalancerPoolParams struct {
	AccountID string
	ZoneID    string

	LoadBalancerPool LoadBalancerPool
}

type ListLoadBalancerPoolParams struct {
	AccountID string
	ZoneID    string

	PaginationOptions
}

type LoadBalancerPoolParams struct {
	AccountID string
	ZoneID    string
	PoolID    string
}

type DeleteLoadBalancerPoolParams struct {
	AccountID string
	ZoneID    string
	PoolID    string
}

type UpdateLoadBalacerPoolParams struct {
	AccountID string
	ZoneID    string

	LoadBalancer LoadBalancerPool
}

type CreateLoadBalancerMonitorParams struct {
	AccountID string
	ZoneID    string

	LoadBalancerMonitor LoadBalancerMonitor
}

type ListLoadBalancerMonitorParams struct {
	AccountID string
	ZoneID    string

	PaginationOptions
}

type LoadBalancerMonitorParams struct {
	AccountID string
	ZoneID    string
	MonitorID string
}

type DeletLoadBalancerMonitorParams struct {
	AccountID string
	ZoneID    string
	MonitorID string
}

type UpdateLoadBalancerMonitorParams struct {
	AccountID string
	ZoneID    string

	LoadBalancerMonitor LoadBalancerMonitor
}

type CreateLoadBalancerParams struct {
	AccountID string
	ZoneID    string

	LoadBalancer LoadBalancer
}

type ListLoadBalancerParams struct {
	AccountID string
	ZoneID    string

	PaginationOptions
}

type LoadBalancerParams struct {
	AccountID      string
	ZoneID         string
	LoadBalancerID string
}

type DeleteLoadBalancerParams struct {
	AccountID      string
	ZoneID         string
	LoadBalancerID string
}

type UpdateLoadBalancerParams struct {
	AccountID string
	ZoneID    string

	LoadBalancer LoadBalancer
}

type LoadBalancerPoolHealthParams struct {
	AccountID string
	ZoneID    string
	PoolID    string
}

var (
	ErrMissingPoolID         = errors.New("missing required pool ID")
	ErrMissingMonitorID      = errors.New("missing required monitor ID")
	ErrMissingLoadBalancerID = errors.New("missing required load balancer ID")
)

// CreateLoadBalancerPool creates a new load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-create-pool
func (api *API) CreateLoadBalancerPool(ctx context.Context, params CreateLoadBalancerPoolParams) (LoadBalancerPool, error) {
	uri := "/user/load_balancers/pools"

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/pools", params.ZoneID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools", params.AccountID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.LoadBalancerPool)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancerPools lists load balancer pools connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-list-pools
func (api *API) ListLoadBalancerPools(ctx context.Context, params ListLoadBalancerPoolParams) ([]LoadBalancerPool, error) {
	uri := "/user/load_balancers/pools"

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/pools", params.ZoneID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools", params.AccountID)
	}

	uri = buildURI(uri, params.PaginationOptions)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerPoolListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerPool returns the details for a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-details
func (api *API) LoadBalancerPool(ctx context.Context, params LoadBalancerPoolParams) (LoadBalancerPool, error) {
	if params.PoolID == "" {
		return LoadBalancerPool{}, ErrMissingPoolID
	}

	uri := fmt.Sprintf("/user/load_balancers/pools/%s", params.PoolID)

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/pools/%s", params.ZoneID, params.PoolID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", params.AccountID, params.PoolID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancerPool disables and deletes a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-delete-pool
func (api *API) DeleteLoadBalancerPool(ctx context.Context, params DeleteLoadBalancerPoolParams) error {
	if params.PoolID == "" {
		return ErrMissingPoolID
	}

	uri := fmt.Sprintf("/user/load_balancers/pools/%s", params.PoolID)

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/pools/%s", params.ZoneID, params.PoolID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", params.AccountID, params.PoolID)
	}

	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}

	return nil
}

// UpdateLoadBalancerPool modifies a configured load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-update-pool
func (api *API) UpdateLoadBalancerPool(ctx context.Context, params UpdateLoadBalacerPoolParams) (LoadBalancerPool, error) {
	if params.LoadBalancer.ID == "" {
		return LoadBalancerPool{}, ErrMissingPoolID
	}

	uri := fmt.Sprintf("/user/load_balancers/pools/%s", params.LoadBalancer.ID)

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/pools/%s", params.ZoneID, params.LoadBalancer.ID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", params.AccountID, params.LoadBalancer.ID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.LoadBalancer)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CreateLoadBalancerMonitor creates a new load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-create-monitor
func (api *API) CreateLoadBalancerMonitor(ctx context.Context, params CreateLoadBalancerMonitorParams) (LoadBalancerMonitor, error) {
	uri := "/user/load_balancers/monitors"

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/monitors", params.ZoneID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/monitors", params.AccountID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.LoadBalancerMonitor)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancerMonitors lists load balancer monitors connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-list-monitors
func (api *API) ListLoadBalancerMonitors(ctx context.Context, params ListLoadBalancerMonitorParams) ([]LoadBalancerMonitor, error) {
	uri := "/user/load_balancers/monitors"

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/monitors", params.ZoneID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/monitors", params.AccountID)
	}

	uri = buildURI(uri, params.PaginationOptions)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerMonitorListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerMonitor returns the details for a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-monitor-details
func (api *API) LoadBalancerMonitor(ctx context.Context, params LoadBalancerMonitorParams) (LoadBalancerMonitor, error) {
	if params.MonitorID == "" {
		return LoadBalancerMonitor{}, ErrMissingMonitorID
	}

	uri := fmt.Sprintf("/user/load_balancers/monitors/%s", params.MonitorID)

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/monitors/%s", params.ZoneID, params.MonitorID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", params.AccountID, params.MonitorID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancerMonitor disables and deletes a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-delete-monitor
func (api *API) DeleteLoadBalancerMonitor(ctx context.Context, params DeletLoadBalancerMonitorParams) error {
	if params.MonitorID == "" {
		return ErrMissingMonitorID
	}

	uri := fmt.Sprintf("/user/load_balancers/monitors/%s", params.MonitorID)

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/monitors/%s", params.ZoneID, params.MonitorID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", params.AccountID, params.MonitorID)
	}

	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// UpdateLoadBalancerMonitor modifies a configured load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-update-monitor
func (api *API) UpdateLoadBalancerMonitor(ctx context.Context, params UpdateLoadBalancerMonitorParams) (LoadBalancerMonitor, error) {
	if params.LoadBalancerMonitor.ID == "" {
		return LoadBalancerMonitor{}, ErrMissingMonitorID
	}

	uri := fmt.Sprintf("/user/load_balancers/monitors/%s", params.LoadBalancerMonitor.ID)

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/monitors/%s", params.ZoneID, params.LoadBalancerMonitor.ID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", params.AccountID, params.LoadBalancerMonitor.ID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.LoadBalancerMonitor)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CreateLoadBalancer creates a new load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-create-load-balancer
func (api *API) CreateLoadBalancer(ctx context.Context, params CreateLoadBalancerParams) (LoadBalancer, error) {
	uri := "/user/load_balancers"

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers", params.ZoneID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers", params.AccountID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.LoadBalancer)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancers lists load balancers configured on a zone.
//
// API reference: https://api.cloudflare.com/#load-balancers-list-load-balancers
func (api *API) ListLoadBalancers(ctx context.Context, params ListLoadBalancerParams) ([]LoadBalancer, error) {
	uri := "/user/load_balancers"

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers", params.ZoneID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers", params.AccountID)
	}

	uri = buildURI(uri, params.PaginationOptions)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancer returns the details for a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-load-balancer-details
func (api *API) LoadBalancer(ctx context.Context, params LoadBalancerParams) (LoadBalancer, error) {
	if params.LoadBalancerID == "" {
		return LoadBalancer{}, ErrMissingLoadBalancerID
	}

	uri := fmt.Sprintf("/user/load_balancers/%s", params.LoadBalancerID)

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/%s", params.ZoneID, params.LoadBalancerID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/%s", params.AccountID, params.LoadBalancerID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancer disables and deletes a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-delete-load-balancer
func (api *API) DeleteLoadBalancer(ctx context.Context, params DeleteLoadBalancerParams) error {
	if params.LoadBalancerID == "" {
		return ErrMissingLoadBalancerID
	}

	uri := fmt.Sprintf("/user/load_balancers/%s", params.LoadBalancerID)

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/%s", params.ZoneID, params.LoadBalancerID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/%s", params.AccountID, params.LoadBalancerID)
	}

	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// UpdateLoadBalancer modifies a configured load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-update-load-balancer
func (api *API) UpdateLoadBalancer(ctx context.Context, params UpdateLoadBalancerParams) (LoadBalancer, error) {
	if params.LoadBalancer.ID == "" {
		return LoadBalancer{}, ErrMissingLoadBalancerID
	}

	uri := fmt.Sprintf("/user/load_balancers/%s", params.LoadBalancer.ID)

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/%s", params.ZoneID, params.LoadBalancer.ID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/%s", params.AccountID, params.LoadBalancer.ID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.LoadBalancer)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerPoolHealth fetches the latest healtcheck details for a single pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-health-details
func (api *API) LoadBalancerPoolHealth(ctx context.Context, params LoadBalancerPoolHealthParams) (LoadBalancerPoolHealth, error) {
	if params.PoolID == "" {
		return LoadBalancerPoolHealth{}, ErrMissingPoolID
	}

	uri := fmt.Sprintf("/user/load_balancers/pools/%s/health", params.PoolID)

	if params.ZoneID != "" {
		uri = fmt.Sprintf("/zones/%s/load_balancers/pools/%s/health", params.ZoneID, params.PoolID)
	}

	if params.AccountID != "" {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools/%s/health", params.AccountID, params.PoolID)
	}
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerPoolHealth{}, err
	}
	var r loadBalancerPoolHealthResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPoolHealth{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}
