// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// LoadBalancerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerService] method
// instead.
type LoadBalancerService struct {
	Options  []option.RequestOption
	Monitors *MonitorService
	Pools    *PoolService
	Previews *PreviewService
	Regions  *RegionService
	Searches *SearchService
}

// NewLoadBalancerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLoadBalancerService(opts ...option.RequestOption) (r *LoadBalancerService) {
	r = &LoadBalancerService{}
	r.Options = opts
	r.Monitors = NewMonitorService(opts...)
	r.Pools = NewPoolService(opts...)
	r.Previews = NewPreviewService(opts...)
	r.Regions = NewRegionService(opts...)
	r.Searches = NewSearchService(opts...)
	return
}

// Create a new load balancer.
func (r *LoadBalancerService) New(ctx context.Context, params LoadBalancerNewParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured load balancer.
func (r *LoadBalancerService) Update(ctx context.Context, loadBalancerID string, params LoadBalancerUpdateParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers/%s", params.ZoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured load balancers.
func (r *LoadBalancerService) List(ctx context.Context, query LoadBalancerListParams, opts ...option.RequestOption) (res *pagination.SinglePage[LoadBalancer], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/load_balancers", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List configured load balancers.
func (r *LoadBalancerService) ListAutoPaging(ctx context.Context, query LoadBalancerListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LoadBalancer] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured load balancer.
func (r *LoadBalancerService) Delete(ctx context.Context, loadBalancerID string, params LoadBalancerDeleteParams, opts ...option.RequestOption) (res *LoadBalancerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers/%s", params.ZoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing load balancer, overwriting the supplied properties.
func (r *LoadBalancerService) Edit(ctx context.Context, loadBalancerID string, params LoadBalancerEditParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers/%s", params.ZoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured load balancer.
func (r *LoadBalancerService) Get(ctx context.Context, loadBalancerID string, query LoadBalancerGetParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers/%s", query.ZoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancer struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerAdaptiveRouting `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools interface{} `json:"country_pools"`
	CreatedOn    time.Time   `json:"created_on" format:"date-time"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools []string `json:"default_pools"`
	// Object description.
	Description string `json:"description"`
	// Whether to enable (the default) this load balancer.
	Enabled bool `json:"enabled"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool interface{} `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy LoadBalancerLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                    `json:"modified_on" format:"date-time"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name string `json:"name"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools interface{} `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied bool `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering LoadBalancerRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancerRule `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"` or "" (default). The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity LoadBalancerSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerSessionAffinityAttributes `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTTL float64 `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy LoadBalancerSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64          `json:"ttl"`
	JSON loadBalancerJSON `json:"-"`
}

// loadBalancerJSON contains the JSON metadata for the struct [LoadBalancer]
type loadBalancerJSON struct {
	ID                        apijson.Field
	AdaptiveRouting           apijson.Field
	CountryPools              apijson.Field
	CreatedOn                 apijson.Field
	DefaultPools              apijson.Field
	Description               apijson.Field
	Enabled                   apijson.Field
	FallbackPool              apijson.Field
	LocationStrategy          apijson.Field
	ModifiedOn                apijson.Field
	Name                      apijson.Field
	PopPools                  apijson.Field
	Proxied                   apijson.Field
	RandomSteering            apijson.Field
	RegionPools               apijson.Field
	Rules                     apijson.Field
	SessionAffinity           apijson.Field
	SessionAffinityAttributes apijson.Field
	SessionAffinityTTL        apijson.Field
	SteeringPolicy            apijson.Field
	TTL                       apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LoadBalancer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerJSON) RawJSON() string {
	return r.raw
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                            `json:"failover_across_pools"`
	JSON                loadBalancerAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerAdaptiveRoutingJSON contains the JSON metadata for the struct
// [LoadBalancerAdaptiveRouting]
type loadBalancerAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerAdaptiveRoutingJSON) RawJSON() string {
	return r.raw
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerLocationStrategyJSON      `json:"-"`
}

// loadBalancerLocationStrategyJSON contains the JSON metadata for the struct
// [LoadBalancerLocationStrategy]
type loadBalancerLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerLocationStrategyJSON) RawJSON() string {
	return r.raw
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerLocationStrategyMode string

const (
	LoadBalancerLocationStrategyModePop        LoadBalancerLocationStrategyMode = "pop"
	LoadBalancerLocationStrategyModeResolverIP LoadBalancerLocationStrategyMode = "resolver_ip"
)

func (r LoadBalancerLocationStrategyMode) IsKnown() bool {
	switch r {
	case LoadBalancerLocationStrategyModePop, LoadBalancerLocationStrategyModeResolverIP:
		return true
	}
	return false
}

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerLocationStrategyPreferEcs string

const (
	LoadBalancerLocationStrategyPreferEcsAlways    LoadBalancerLocationStrategyPreferEcs = "always"
	LoadBalancerLocationStrategyPreferEcsNever     LoadBalancerLocationStrategyPreferEcs = "never"
	LoadBalancerLocationStrategyPreferEcsProximity LoadBalancerLocationStrategyPreferEcs = "proximity"
	LoadBalancerLocationStrategyPreferEcsGeo       LoadBalancerLocationStrategyPreferEcs = "geo"
)

func (r LoadBalancerLocationStrategyPreferEcs) IsKnown() bool {
	switch r {
	case LoadBalancerLocationStrategyPreferEcsAlways, LoadBalancerLocationStrategyPreferEcsNever, LoadBalancerLocationStrategyPreferEcsProximity, LoadBalancerLocationStrategyPreferEcsGeo:
		return true
	}
	return false
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                    `json:"pool_weights"`
	JSON        loadBalancerRandomSteeringJSON `json:"-"`
}

// loadBalancerRandomSteeringJSON contains the JSON metadata for the struct
// [LoadBalancerRandomSteering]
type loadBalancerRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerRandomSteeringJSON) RawJSON() string {
	return r.raw
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerRule struct {
	// The condition expressions to evaluate. If the condition evaluates to true, the
	// overrides or fixed_response in this rule will be applied. An empty condition is
	// always true. For more details on condition expressions, please see
	// https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.
	Condition string `json:"condition"`
	// Disable this specific rule. It will no longer be evaluated by this load
	// balancer.
	Disabled bool `json:"disabled"`
	// A collection of fields used to directly respond to the eyeball instead of
	// routing to a pool. If a fixed_response is supplied the rule will be marked as
	// terminates.
	FixedResponse LoadBalancerRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancerRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                 `json:"terminates"`
	JSON       loadBalancerRuleJSON `json:"-"`
}

// loadBalancerRuleJSON contains the JSON metadata for the struct
// [LoadBalancerRule]
type loadBalancerRuleJSON struct {
	Condition     apijson.Field
	Disabled      apijson.Field
	FixedResponse apijson.Field
	Name          apijson.Field
	Overrides     apijson.Field
	Priority      apijson.Field
	Terminates    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerRuleJSON) RawJSON() string {
	return r.raw
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                              `json:"status_code"`
	JSON       loadBalancerRulesFixedResponseJSON `json:"-"`
}

// loadBalancerRulesFixedResponseJSON contains the JSON metadata for the struct
// [LoadBalancerRulesFixedResponse]
type loadBalancerRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerRulesFixedResponseJSON) RawJSON() string {
	return r.raw
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools interface{} `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools []string `json:"default_pools"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool interface{} `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy LoadBalancerRulesOverridesLocationStrategy `json:"location_strategy"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools interface{} `json:"pop_pools"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering LoadBalancerRulesOverridesRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"` or "" (default). The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity LoadBalancerRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTTL float64 `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy LoadBalancerRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                        `json:"ttl"`
	JSON loadBalancerRulesOverridesJSON `json:"-"`
}

// loadBalancerRulesOverridesJSON contains the JSON metadata for the struct
// [LoadBalancerRulesOverrides]
type loadBalancerRulesOverridesJSON struct {
	AdaptiveRouting           apijson.Field
	CountryPools              apijson.Field
	DefaultPools              apijson.Field
	FallbackPool              apijson.Field
	LocationStrategy          apijson.Field
	PopPools                  apijson.Field
	RandomSteering            apijson.Field
	RegionPools               apijson.Field
	SessionAffinity           apijson.Field
	SessionAffinityAttributes apijson.Field
	SessionAffinityTTL        apijson.Field
	SteeringPolicy            apijson.Field
	TTL                       apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LoadBalancerRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerRulesOverridesJSON) RawJSON() string {
	return r.raw
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                          `json:"failover_across_pools"`
	JSON                loadBalancerRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerRulesOverridesAdaptiveRoutingJSON contains the JSON metadata for the
// struct [LoadBalancerRulesOverridesAdaptiveRouting]
type loadBalancerRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerRulesOverridesAdaptiveRoutingJSON) RawJSON() string {
	return r.raw
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerRulesOverridesLocationStrategyJSON      `json:"-"`
}

// loadBalancerRulesOverridesLocationStrategyJSON contains the JSON metadata for
// the struct [LoadBalancerRulesOverridesLocationStrategy]
type loadBalancerRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerRulesOverridesLocationStrategyJSON) RawJSON() string {
	return r.raw
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerRulesOverridesLocationStrategyMode string

const (
	LoadBalancerRulesOverridesLocationStrategyModePop        LoadBalancerRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerRulesOverridesLocationStrategyModeResolverIP LoadBalancerRulesOverridesLocationStrategyMode = "resolver_ip"
)

func (r LoadBalancerRulesOverridesLocationStrategyMode) IsKnown() bool {
	switch r {
	case LoadBalancerRulesOverridesLocationStrategyModePop, LoadBalancerRulesOverridesLocationStrategyModeResolverIP:
		return true
	}
	return false
}

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerRulesOverridesLocationStrategyPreferEcs = "geo"
)

func (r LoadBalancerRulesOverridesLocationStrategyPreferEcs) IsKnown() bool {
	switch r {
	case LoadBalancerRulesOverridesLocationStrategyPreferEcsAlways, LoadBalancerRulesOverridesLocationStrategyPreferEcsNever, LoadBalancerRulesOverridesLocationStrategyPreferEcsProximity, LoadBalancerRulesOverridesLocationStrategyPreferEcsGeo:
		return true
	}
	return false
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                  `json:"pool_weights"`
	JSON        loadBalancerRulesOverridesRandomSteeringJSON `json:"-"`
}

// loadBalancerRulesOverridesRandomSteeringJSON contains the JSON metadata for the
// struct [LoadBalancerRulesOverridesRandomSteering]
type loadBalancerRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerRulesOverridesRandomSteeringJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of session affinity the load balancer should use unless
// specified as `"none"` or "" (default). The supported types are:
//
//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
//     generated, encoding information of which origin the request will be forwarded
//     to. Subsequent requests, by the same client to the same load balancer, will be
//     sent to the origin server the cookie encodes, for the duration of the cookie
//     and as long as the origin server remains healthy. If the cookie has expired or
//     the origin server is unhealthy, then a new origin server is calculated and
//     used.
//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
//     selection is stable and based on the client's ip address.
//   - `"header"`: On the first request to a proxied load balancer, a session key
//     based on the configured HTTP headers (see
//     `session_affinity_attributes.headers`) is generated, encoding the request
//     headers used for storing in the load balancer session state which origin the
//     request will be forwarded to. Subsequent requests to the load balancer with
//     the same headers will be sent to the same origin server, for the duration of
//     the session and as long as the origin server remains healthy. If the session
//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
//     server is unhealthy, then a new origin server is calculated and used. See
//     `headers` in `session_affinity_attributes` for additional required
//     configuration.
type LoadBalancerRulesOverridesSessionAffinity string

const (
	LoadBalancerRulesOverridesSessionAffinityNone     LoadBalancerRulesOverridesSessionAffinity = "none"
	LoadBalancerRulesOverridesSessionAffinityCookie   LoadBalancerRulesOverridesSessionAffinity = "cookie"
	LoadBalancerRulesOverridesSessionAffinityIPCookie LoadBalancerRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerRulesOverridesSessionAffinityHeader   LoadBalancerRulesOverridesSessionAffinity = "header"
	LoadBalancerRulesOverridesSessionAffinityEmpty    LoadBalancerRulesOverridesSessionAffinity = "\"\""
)

func (r LoadBalancerRulesOverridesSessionAffinity) IsKnown() bool {
	switch r {
	case LoadBalancerRulesOverridesSessionAffinityNone, LoadBalancerRulesOverridesSessionAffinityCookie, LoadBalancerRulesOverridesSessionAffinityIPCookie, LoadBalancerRulesOverridesSessionAffinityHeader, LoadBalancerRulesOverridesSessionAffinityEmpty:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type LoadBalancerRulesOverridesSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration float64 `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers []string `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are:
	//
	//   - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	//     specified by the `headers` session affinity attribute, otherwise sessions
	//     aren't created.
	//   - `"false"`: Load balancing requests must contain _at least one_ of the HTTP
	//     headers specified by the `headers` session affinity attribute, otherwise
	//     sessions aren't created.
	RequireAllHeaders bool `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite LoadBalancerRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are:
	//
	//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
	//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
	//     originally pinned origin is available; note that this can potentially result
	//     in heavy origin flapping.
	//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
	//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
	//     currently not supported for session affinity by header.
	ZeroDowntimeFailover LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerRulesOverridesSessionAffinityAttributesJSON contains the JSON
// metadata for the struct [LoadBalancerRulesOverridesSessionAffinityAttributes]
type loadBalancerRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerRulesOverridesSessionAffinityAttributesJSON) RawJSON() string {
	return r.raw
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

func (r LoadBalancerRulesOverridesSessionAffinityAttributesSamesite) IsKnown() bool {
	switch r {
	case LoadBalancerRulesOverridesSessionAffinityAttributesSamesiteAuto, LoadBalancerRulesOverridesSessionAffinityAttributesSamesiteLax, LoadBalancerRulesOverridesSessionAffinityAttributesSamesiteNone, LoadBalancerRulesOverridesSessionAffinityAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerRulesOverridesSessionAffinityAttributesSecure = "Never"
)

func (r LoadBalancerRulesOverridesSessionAffinityAttributesSecure) IsKnown() bool {
	switch r {
	case LoadBalancerRulesOverridesSessionAffinityAttributesSecureAuto, LoadBalancerRulesOverridesSessionAffinityAttributesSecureAlways, LoadBalancerRulesOverridesSessionAffinityAttributesSecureNever:
		return true
	}
	return false
}

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. This feature is currently incompatible with Argo, Tiered
// Cache, and Bandwidth Alliance. The supported values are:
//
//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
//     originally pinned origin is available; note that this can potentially result
//     in heavy origin flapping.
//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
//     currently not supported for session affinity by header.
type LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

func (r LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailover) IsKnown() bool {
	switch r {
	case LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone, LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary, LoadBalancerRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky:
		return true
	}
	return false
}

// Steering Policy for this load balancer.
//
//   - `"off"`: Use `default_pools`.
//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
//     requests, the country for `country_pools` is determined by
//     `location_strategy`.
//   - `"random"`: Select a pool randomly.
//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
//     default_pools (requires pool health checks).
//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
//     pool using the Cloudflare PoP location for proxied requests or the location
//     determined by `location_strategy` for non-proxied requests.
//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of outstanding
//     requests. Pools with more pending requests are weighted proportionately less
//     relative to others.
//   - `"least_connections"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of open connections.
//     Pools with more open connections are weighted proportionately less relative to
//     others. Supported for HTTP/1 and HTTP/2 connections.
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerRulesOverridesSteeringPolicy string

const (
	LoadBalancerRulesOverridesSteeringPolicyOff                      LoadBalancerRulesOverridesSteeringPolicy = "off"
	LoadBalancerRulesOverridesSteeringPolicyGeo                      LoadBalancerRulesOverridesSteeringPolicy = "geo"
	LoadBalancerRulesOverridesSteeringPolicyRandom                   LoadBalancerRulesOverridesSteeringPolicy = "random"
	LoadBalancerRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerRulesOverridesSteeringPolicyProximity                LoadBalancerRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerRulesOverridesSteeringPolicyLeastConnections         LoadBalancerRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerRulesOverridesSteeringPolicyEmpty                    LoadBalancerRulesOverridesSteeringPolicy = "\"\""
)

func (r LoadBalancerRulesOverridesSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerRulesOverridesSteeringPolicyOff, LoadBalancerRulesOverridesSteeringPolicyGeo, LoadBalancerRulesOverridesSteeringPolicyRandom, LoadBalancerRulesOverridesSteeringPolicyDynamicLatency, LoadBalancerRulesOverridesSteeringPolicyProximity, LoadBalancerRulesOverridesSteeringPolicyLeastOutstandingRequests, LoadBalancerRulesOverridesSteeringPolicyLeastConnections, LoadBalancerRulesOverridesSteeringPolicyEmpty:
		return true
	}
	return false
}

// Specifies the type of session affinity the load balancer should use unless
// specified as `"none"` or "" (default). The supported types are:
//
//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
//     generated, encoding information of which origin the request will be forwarded
//     to. Subsequent requests, by the same client to the same load balancer, will be
//     sent to the origin server the cookie encodes, for the duration of the cookie
//     and as long as the origin server remains healthy. If the cookie has expired or
//     the origin server is unhealthy, then a new origin server is calculated and
//     used.
//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
//     selection is stable and based on the client's ip address.
//   - `"header"`: On the first request to a proxied load balancer, a session key
//     based on the configured HTTP headers (see
//     `session_affinity_attributes.headers`) is generated, encoding the request
//     headers used for storing in the load balancer session state which origin the
//     request will be forwarded to. Subsequent requests to the load balancer with
//     the same headers will be sent to the same origin server, for the duration of
//     the session and as long as the origin server remains healthy. If the session
//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
//     server is unhealthy, then a new origin server is calculated and used. See
//     `headers` in `session_affinity_attributes` for additional required
//     configuration.
type LoadBalancerSessionAffinity string

const (
	LoadBalancerSessionAffinityNone     LoadBalancerSessionAffinity = "none"
	LoadBalancerSessionAffinityCookie   LoadBalancerSessionAffinity = "cookie"
	LoadBalancerSessionAffinityIPCookie LoadBalancerSessionAffinity = "ip_cookie"
	LoadBalancerSessionAffinityHeader   LoadBalancerSessionAffinity = "header"
	LoadBalancerSessionAffinityEmpty    LoadBalancerSessionAffinity = "\"\""
)

func (r LoadBalancerSessionAffinity) IsKnown() bool {
	switch r {
	case LoadBalancerSessionAffinityNone, LoadBalancerSessionAffinityCookie, LoadBalancerSessionAffinityIPCookie, LoadBalancerSessionAffinityHeader, LoadBalancerSessionAffinityEmpty:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type LoadBalancerSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration float64 `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers []string `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are:
	//
	//   - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	//     specified by the `headers` session affinity attribute, otherwise sessions
	//     aren't created.
	//   - `"false"`: Load balancing requests must contain _at least one_ of the HTTP
	//     headers specified by the `headers` session affinity attribute, otherwise
	//     sessions aren't created.
	RequireAllHeaders bool `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite LoadBalancerSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerSessionAffinityAttributesSecure `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are:
	//
	//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
	//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
	//     originally pinned origin is available; note that this can potentially result
	//     in heavy origin flapping.
	//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
	//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
	//     currently not supported for session affinity by header.
	ZeroDowntimeFailover LoadBalancerSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerSessionAffinityAttributesJSON contains the JSON metadata for the
// struct [LoadBalancerSessionAffinityAttributes]
type loadBalancerSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerSessionAffinityAttributesJSON) RawJSON() string {
	return r.raw
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerSessionAffinityAttributesSamesite string

const (
	LoadBalancerSessionAffinityAttributesSamesiteAuto   LoadBalancerSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerSessionAffinityAttributesSamesiteLax    LoadBalancerSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerSessionAffinityAttributesSamesiteNone   LoadBalancerSessionAffinityAttributesSamesite = "None"
	LoadBalancerSessionAffinityAttributesSamesiteStrict LoadBalancerSessionAffinityAttributesSamesite = "Strict"
)

func (r LoadBalancerSessionAffinityAttributesSamesite) IsKnown() bool {
	switch r {
	case LoadBalancerSessionAffinityAttributesSamesiteAuto, LoadBalancerSessionAffinityAttributesSamesiteLax, LoadBalancerSessionAffinityAttributesSamesiteNone, LoadBalancerSessionAffinityAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerSessionAffinityAttributesSecure string

const (
	LoadBalancerSessionAffinityAttributesSecureAuto   LoadBalancerSessionAffinityAttributesSecure = "Auto"
	LoadBalancerSessionAffinityAttributesSecureAlways LoadBalancerSessionAffinityAttributesSecure = "Always"
	LoadBalancerSessionAffinityAttributesSecureNever  LoadBalancerSessionAffinityAttributesSecure = "Never"
)

func (r LoadBalancerSessionAffinityAttributesSecure) IsKnown() bool {
	switch r {
	case LoadBalancerSessionAffinityAttributesSecureAuto, LoadBalancerSessionAffinityAttributesSecureAlways, LoadBalancerSessionAffinityAttributesSecureNever:
		return true
	}
	return false
}

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. This feature is currently incompatible with Argo, Tiered
// Cache, and Bandwidth Alliance. The supported values are:
//
//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
//     originally pinned origin is available; note that this can potentially result
//     in heavy origin flapping.
//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
//     currently not supported for session affinity by header.
type LoadBalancerSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

func (r LoadBalancerSessionAffinityAttributesZeroDowntimeFailover) IsKnown() bool {
	switch r {
	case LoadBalancerSessionAffinityAttributesZeroDowntimeFailoverNone, LoadBalancerSessionAffinityAttributesZeroDowntimeFailoverTemporary, LoadBalancerSessionAffinityAttributesZeroDowntimeFailoverSticky:
		return true
	}
	return false
}

// Steering Policy for this load balancer.
//
//   - `"off"`: Use `default_pools`.
//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
//     requests, the country for `country_pools` is determined by
//     `location_strategy`.
//   - `"random"`: Select a pool randomly.
//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
//     default_pools (requires pool health checks).
//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
//     pool using the Cloudflare PoP location for proxied requests or the location
//     determined by `location_strategy` for non-proxied requests.
//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of outstanding
//     requests. Pools with more pending requests are weighted proportionately less
//     relative to others.
//   - `"least_connections"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of open connections.
//     Pools with more open connections are weighted proportionately less relative to
//     others. Supported for HTTP/1 and HTTP/2 connections.
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerSteeringPolicy string

const (
	LoadBalancerSteeringPolicyOff                      LoadBalancerSteeringPolicy = "off"
	LoadBalancerSteeringPolicyGeo                      LoadBalancerSteeringPolicy = "geo"
	LoadBalancerSteeringPolicyRandom                   LoadBalancerSteeringPolicy = "random"
	LoadBalancerSteeringPolicyDynamicLatency           LoadBalancerSteeringPolicy = "dynamic_latency"
	LoadBalancerSteeringPolicyProximity                LoadBalancerSteeringPolicy = "proximity"
	LoadBalancerSteeringPolicyLeastOutstandingRequests LoadBalancerSteeringPolicy = "least_outstanding_requests"
	LoadBalancerSteeringPolicyLeastConnections         LoadBalancerSteeringPolicy = "least_connections"
	LoadBalancerSteeringPolicyEmpty                    LoadBalancerSteeringPolicy = "\"\""
)

func (r LoadBalancerSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerSteeringPolicyOff, LoadBalancerSteeringPolicyGeo, LoadBalancerSteeringPolicyRandom, LoadBalancerSteeringPolicyDynamicLatency, LoadBalancerSteeringPolicyProximity, LoadBalancerSteeringPolicyLeastOutstandingRequests, LoadBalancerSteeringPolicyLeastConnections, LoadBalancerSteeringPolicyEmpty:
		return true
	}
	return false
}

type LoadBalancerDeleteResponse struct {
	ID   string                         `json:"id"`
	JSON loadBalancerDeleteResponseJSON `json:"-"`
}

// loadBalancerDeleteResponseJSON contains the JSON metadata for the struct
// [LoadBalancerDeleteResponse]
type loadBalancerDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerNewParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]string] `json:"default_pools,required"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[interface{}] `json:"fallback_pool,required"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name param.Field[string] `json:"name,required"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[LoadBalancerNewParamsAdaptiveRouting] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[interface{}] `json:"country_pools"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LoadBalancerNewParamsLocationStrategy] `json:"location_strategy"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools param.Field[interface{}] `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied param.Field[bool] `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[LoadBalancerNewParamsRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]LoadBalancerNewParamsRule] `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"` or "" (default). The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity param.Field[LoadBalancerNewParamsSessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[LoadBalancerNewParamsSessionAffinityAttributes] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTTL param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[LoadBalancerNewParamsSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r LoadBalancerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerNewParamsAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r LoadBalancerNewParamsAdaptiveRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerNewParamsLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[LoadBalancerNewParamsLocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[LoadBalancerNewParamsLocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r LoadBalancerNewParamsLocationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerNewParamsLocationStrategyMode string

const (
	LoadBalancerNewParamsLocationStrategyModePop        LoadBalancerNewParamsLocationStrategyMode = "pop"
	LoadBalancerNewParamsLocationStrategyModeResolverIP LoadBalancerNewParamsLocationStrategyMode = "resolver_ip"
)

func (r LoadBalancerNewParamsLocationStrategyMode) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsLocationStrategyModePop, LoadBalancerNewParamsLocationStrategyModeResolverIP:
		return true
	}
	return false
}

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerNewParamsLocationStrategyPreferEcs string

const (
	LoadBalancerNewParamsLocationStrategyPreferEcsAlways    LoadBalancerNewParamsLocationStrategyPreferEcs = "always"
	LoadBalancerNewParamsLocationStrategyPreferEcsNever     LoadBalancerNewParamsLocationStrategyPreferEcs = "never"
	LoadBalancerNewParamsLocationStrategyPreferEcsProximity LoadBalancerNewParamsLocationStrategyPreferEcs = "proximity"
	LoadBalancerNewParamsLocationStrategyPreferEcsGeo       LoadBalancerNewParamsLocationStrategyPreferEcs = "geo"
)

func (r LoadBalancerNewParamsLocationStrategyPreferEcs) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsLocationStrategyPreferEcsAlways, LoadBalancerNewParamsLocationStrategyPreferEcsNever, LoadBalancerNewParamsLocationStrategyPreferEcsProximity, LoadBalancerNewParamsLocationStrategyPreferEcsGeo:
		return true
	}
	return false
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerNewParamsRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[interface{}] `json:"pool_weights"`
}

func (r LoadBalancerNewParamsRandomSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerNewParamsRule struct {
	// The condition expressions to evaluate. If the condition evaluates to true, the
	// overrides or fixed_response in this rule will be applied. An empty condition is
	// always true. For more details on condition expressions, please see
	// https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.
	Condition param.Field[string] `json:"condition"`
	// Disable this specific rule. It will no longer be evaluated by this load
	// balancer.
	Disabled param.Field[bool] `json:"disabled"`
	// A collection of fields used to directly respond to the eyeball instead of
	// routing to a pool. If a fixed_response is supplied the rule will be marked as
	// terminates.
	FixedResponse param.Field[LoadBalancerNewParamsRulesFixedResponse] `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides param.Field[LoadBalancerNewParamsRulesOverrides] `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority param.Field[int64] `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates param.Field[bool] `json:"terminates"`
}

func (r LoadBalancerNewParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerNewParamsRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType param.Field[string] `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location param.Field[string] `json:"location"`
	// Text to include as the http body.
	MessageBody param.Field[string] `json:"message_body"`
	// The http status code to respond with.
	StatusCode param.Field[int64] `json:"status_code"`
}

func (r LoadBalancerNewParamsRulesFixedResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerNewParamsRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[LoadBalancerNewParamsRulesOverridesAdaptiveRouting] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[interface{}] `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]string] `json:"default_pools"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[interface{}] `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LoadBalancerNewParamsRulesOverridesLocationStrategy] `json:"location_strategy"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools param.Field[interface{}] `json:"pop_pools"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[LoadBalancerNewParamsRulesOverridesRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"` or "" (default). The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity param.Field[LoadBalancerNewParamsRulesOverridesSessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[LoadBalancerNewParamsRulesOverridesSessionAffinityAttributes] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTTL param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[LoadBalancerNewParamsRulesOverridesSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r LoadBalancerNewParamsRulesOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerNewParamsRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r LoadBalancerNewParamsRulesOverridesAdaptiveRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerNewParamsRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[LoadBalancerNewParamsRulesOverridesLocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r LoadBalancerNewParamsRulesOverridesLocationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerNewParamsRulesOverridesLocationStrategyMode string

const (
	LoadBalancerNewParamsRulesOverridesLocationStrategyModePop        LoadBalancerNewParamsRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP LoadBalancerNewParamsRulesOverridesLocationStrategyMode = "resolver_ip"
)

func (r LoadBalancerNewParamsRulesOverridesLocationStrategyMode) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsRulesOverridesLocationStrategyModePop, LoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP:
		return true
	}
	return false
}

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs = "geo"
)

func (r LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways, LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsNever, LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsProximity, LoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsGeo:
		return true
	}
	return false
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerNewParamsRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[interface{}] `json:"pool_weights"`
}

func (r LoadBalancerNewParamsRulesOverridesRandomSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of session affinity the load balancer should use unless
// specified as `"none"` or "" (default). The supported types are:
//
//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
//     generated, encoding information of which origin the request will be forwarded
//     to. Subsequent requests, by the same client to the same load balancer, will be
//     sent to the origin server the cookie encodes, for the duration of the cookie
//     and as long as the origin server remains healthy. If the cookie has expired or
//     the origin server is unhealthy, then a new origin server is calculated and
//     used.
//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
//     selection is stable and based on the client's ip address.
//   - `"header"`: On the first request to a proxied load balancer, a session key
//     based on the configured HTTP headers (see
//     `session_affinity_attributes.headers`) is generated, encoding the request
//     headers used for storing in the load balancer session state which origin the
//     request will be forwarded to. Subsequent requests to the load balancer with
//     the same headers will be sent to the same origin server, for the duration of
//     the session and as long as the origin server remains healthy. If the session
//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
//     server is unhealthy, then a new origin server is calculated and used. See
//     `headers` in `session_affinity_attributes` for additional required
//     configuration.
type LoadBalancerNewParamsRulesOverridesSessionAffinity string

const (
	LoadBalancerNewParamsRulesOverridesSessionAffinityNone     LoadBalancerNewParamsRulesOverridesSessionAffinity = "none"
	LoadBalancerNewParamsRulesOverridesSessionAffinityCookie   LoadBalancerNewParamsRulesOverridesSessionAffinity = "cookie"
	LoadBalancerNewParamsRulesOverridesSessionAffinityIPCookie LoadBalancerNewParamsRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerNewParamsRulesOverridesSessionAffinityHeader   LoadBalancerNewParamsRulesOverridesSessionAffinity = "header"
	LoadBalancerNewParamsRulesOverridesSessionAffinityEmpty    LoadBalancerNewParamsRulesOverridesSessionAffinity = "\"\""
)

func (r LoadBalancerNewParamsRulesOverridesSessionAffinity) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsRulesOverridesSessionAffinityNone, LoadBalancerNewParamsRulesOverridesSessionAffinityCookie, LoadBalancerNewParamsRulesOverridesSessionAffinityIPCookie, LoadBalancerNewParamsRulesOverridesSessionAffinityHeader, LoadBalancerNewParamsRulesOverridesSessionAffinityEmpty:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type LoadBalancerNewParamsRulesOverridesSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers param.Field[[]string] `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are:
	//
	//   - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	//     specified by the `headers` session affinity attribute, otherwise sessions
	//     aren't created.
	//   - `"false"`: Load balancing requests must contain _at least one_ of the HTTP
	//     headers specified by the `headers` session affinity attribute, otherwise
	//     sessions aren't created.
	RequireAllHeaders param.Field[bool] `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite param.Field[LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure] `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are:
	//
	//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
	//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
	//     originally pinned origin is available; note that this can potentially result
	//     in heavy origin flapping.
	//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
	//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
	//     currently not supported for session affinity by header.
	ZeroDowntimeFailover param.Field[LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r LoadBalancerNewParamsRulesOverridesSessionAffinityAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

func (r LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto, LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteLax, LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteNone, LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure = "Never"
)

func (r LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto, LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAlways, LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureNever:
		return true
	}
	return false
}

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. This feature is currently incompatible with Argo, Tiered
// Cache, and Bandwidth Alliance. The supported values are:
//
//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
//     originally pinned origin is available; note that this can potentially result
//     in heavy origin flapping.
//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
//     currently not supported for session affinity by header.
type LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

func (r LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone, LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary, LoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky:
		return true
	}
	return false
}

// Steering Policy for this load balancer.
//
//   - `"off"`: Use `default_pools`.
//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
//     requests, the country for `country_pools` is determined by
//     `location_strategy`.
//   - `"random"`: Select a pool randomly.
//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
//     default_pools (requires pool health checks).
//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
//     pool using the Cloudflare PoP location for proxied requests or the location
//     determined by `location_strategy` for non-proxied requests.
//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of outstanding
//     requests. Pools with more pending requests are weighted proportionately less
//     relative to others.
//   - `"least_connections"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of open connections.
//     Pools with more open connections are weighted proportionately less relative to
//     others. Supported for HTTP/1 and HTTP/2 connections.
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerNewParamsRulesOverridesSteeringPolicy string

const (
	LoadBalancerNewParamsRulesOverridesSteeringPolicyOff                      LoadBalancerNewParamsRulesOverridesSteeringPolicy = "off"
	LoadBalancerNewParamsRulesOverridesSteeringPolicyGeo                      LoadBalancerNewParamsRulesOverridesSteeringPolicy = "geo"
	LoadBalancerNewParamsRulesOverridesSteeringPolicyRandom                   LoadBalancerNewParamsRulesOverridesSteeringPolicy = "random"
	LoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerNewParamsRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerNewParamsRulesOverridesSteeringPolicyProximity                LoadBalancerNewParamsRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerNewParamsRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerNewParamsRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerNewParamsRulesOverridesSteeringPolicyLeastConnections         LoadBalancerNewParamsRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerNewParamsRulesOverridesSteeringPolicyEmpty                    LoadBalancerNewParamsRulesOverridesSteeringPolicy = "\"\""
)

func (r LoadBalancerNewParamsRulesOverridesSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsRulesOverridesSteeringPolicyOff, LoadBalancerNewParamsRulesOverridesSteeringPolicyGeo, LoadBalancerNewParamsRulesOverridesSteeringPolicyRandom, LoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency, LoadBalancerNewParamsRulesOverridesSteeringPolicyProximity, LoadBalancerNewParamsRulesOverridesSteeringPolicyLeastOutstandingRequests, LoadBalancerNewParamsRulesOverridesSteeringPolicyLeastConnections, LoadBalancerNewParamsRulesOverridesSteeringPolicyEmpty:
		return true
	}
	return false
}

// Specifies the type of session affinity the load balancer should use unless
// specified as `"none"` or "" (default). The supported types are:
//
//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
//     generated, encoding information of which origin the request will be forwarded
//     to. Subsequent requests, by the same client to the same load balancer, will be
//     sent to the origin server the cookie encodes, for the duration of the cookie
//     and as long as the origin server remains healthy. If the cookie has expired or
//     the origin server is unhealthy, then a new origin server is calculated and
//     used.
//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
//     selection is stable and based on the client's ip address.
//   - `"header"`: On the first request to a proxied load balancer, a session key
//     based on the configured HTTP headers (see
//     `session_affinity_attributes.headers`) is generated, encoding the request
//     headers used for storing in the load balancer session state which origin the
//     request will be forwarded to. Subsequent requests to the load balancer with
//     the same headers will be sent to the same origin server, for the duration of
//     the session and as long as the origin server remains healthy. If the session
//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
//     server is unhealthy, then a new origin server is calculated and used. See
//     `headers` in `session_affinity_attributes` for additional required
//     configuration.
type LoadBalancerNewParamsSessionAffinity string

const (
	LoadBalancerNewParamsSessionAffinityNone     LoadBalancerNewParamsSessionAffinity = "none"
	LoadBalancerNewParamsSessionAffinityCookie   LoadBalancerNewParamsSessionAffinity = "cookie"
	LoadBalancerNewParamsSessionAffinityIPCookie LoadBalancerNewParamsSessionAffinity = "ip_cookie"
	LoadBalancerNewParamsSessionAffinityHeader   LoadBalancerNewParamsSessionAffinity = "header"
	LoadBalancerNewParamsSessionAffinityEmpty    LoadBalancerNewParamsSessionAffinity = "\"\""
)

func (r LoadBalancerNewParamsSessionAffinity) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsSessionAffinityNone, LoadBalancerNewParamsSessionAffinityCookie, LoadBalancerNewParamsSessionAffinityIPCookie, LoadBalancerNewParamsSessionAffinityHeader, LoadBalancerNewParamsSessionAffinityEmpty:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type LoadBalancerNewParamsSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers param.Field[[]string] `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are:
	//
	//   - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	//     specified by the `headers` session affinity attribute, otherwise sessions
	//     aren't created.
	//   - `"false"`: Load balancing requests must contain _at least one_ of the HTTP
	//     headers specified by the `headers` session affinity attribute, otherwise
	//     sessions aren't created.
	RequireAllHeaders param.Field[bool] `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite param.Field[LoadBalancerNewParamsSessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[LoadBalancerNewParamsSessionAffinityAttributesSecure] `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are:
	//
	//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
	//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
	//     originally pinned origin is available; note that this can potentially result
	//     in heavy origin flapping.
	//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
	//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
	//     currently not supported for session affinity by header.
	ZeroDowntimeFailover param.Field[LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r LoadBalancerNewParamsSessionAffinityAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerNewParamsSessionAffinityAttributesSamesite string

const (
	LoadBalancerNewParamsSessionAffinityAttributesSamesiteAuto   LoadBalancerNewParamsSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerNewParamsSessionAffinityAttributesSamesiteLax    LoadBalancerNewParamsSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerNewParamsSessionAffinityAttributesSamesiteNone   LoadBalancerNewParamsSessionAffinityAttributesSamesite = "None"
	LoadBalancerNewParamsSessionAffinityAttributesSamesiteStrict LoadBalancerNewParamsSessionAffinityAttributesSamesite = "Strict"
)

func (r LoadBalancerNewParamsSessionAffinityAttributesSamesite) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsSessionAffinityAttributesSamesiteAuto, LoadBalancerNewParamsSessionAffinityAttributesSamesiteLax, LoadBalancerNewParamsSessionAffinityAttributesSamesiteNone, LoadBalancerNewParamsSessionAffinityAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerNewParamsSessionAffinityAttributesSecure string

const (
	LoadBalancerNewParamsSessionAffinityAttributesSecureAuto   LoadBalancerNewParamsSessionAffinityAttributesSecure = "Auto"
	LoadBalancerNewParamsSessionAffinityAttributesSecureAlways LoadBalancerNewParamsSessionAffinityAttributesSecure = "Always"
	LoadBalancerNewParamsSessionAffinityAttributesSecureNever  LoadBalancerNewParamsSessionAffinityAttributesSecure = "Never"
)

func (r LoadBalancerNewParamsSessionAffinityAttributesSecure) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsSessionAffinityAttributesSecureAuto, LoadBalancerNewParamsSessionAffinityAttributesSecureAlways, LoadBalancerNewParamsSessionAffinityAttributesSecureNever:
		return true
	}
	return false
}

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. This feature is currently incompatible with Argo, Tiered
// Cache, and Bandwidth Alliance. The supported values are:
//
//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
//     originally pinned origin is available; note that this can potentially result
//     in heavy origin flapping.
//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
//     currently not supported for session affinity by header.
type LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

func (r LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverNone, LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverTemporary, LoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverSticky:
		return true
	}
	return false
}

// Steering Policy for this load balancer.
//
//   - `"off"`: Use `default_pools`.
//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
//     requests, the country for `country_pools` is determined by
//     `location_strategy`.
//   - `"random"`: Select a pool randomly.
//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
//     default_pools (requires pool health checks).
//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
//     pool using the Cloudflare PoP location for proxied requests or the location
//     determined by `location_strategy` for non-proxied requests.
//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of outstanding
//     requests. Pools with more pending requests are weighted proportionately less
//     relative to others.
//   - `"least_connections"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of open connections.
//     Pools with more open connections are weighted proportionately less relative to
//     others. Supported for HTTP/1 and HTTP/2 connections.
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerNewParamsSteeringPolicy string

const (
	LoadBalancerNewParamsSteeringPolicyOff                      LoadBalancerNewParamsSteeringPolicy = "off"
	LoadBalancerNewParamsSteeringPolicyGeo                      LoadBalancerNewParamsSteeringPolicy = "geo"
	LoadBalancerNewParamsSteeringPolicyRandom                   LoadBalancerNewParamsSteeringPolicy = "random"
	LoadBalancerNewParamsSteeringPolicyDynamicLatency           LoadBalancerNewParamsSteeringPolicy = "dynamic_latency"
	LoadBalancerNewParamsSteeringPolicyProximity                LoadBalancerNewParamsSteeringPolicy = "proximity"
	LoadBalancerNewParamsSteeringPolicyLeastOutstandingRequests LoadBalancerNewParamsSteeringPolicy = "least_outstanding_requests"
	LoadBalancerNewParamsSteeringPolicyLeastConnections         LoadBalancerNewParamsSteeringPolicy = "least_connections"
	LoadBalancerNewParamsSteeringPolicyEmpty                    LoadBalancerNewParamsSteeringPolicy = "\"\""
)

func (r LoadBalancerNewParamsSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerNewParamsSteeringPolicyOff, LoadBalancerNewParamsSteeringPolicyGeo, LoadBalancerNewParamsSteeringPolicyRandom, LoadBalancerNewParamsSteeringPolicyDynamicLatency, LoadBalancerNewParamsSteeringPolicyProximity, LoadBalancerNewParamsSteeringPolicyLeastOutstandingRequests, LoadBalancerNewParamsSteeringPolicyLeastConnections, LoadBalancerNewParamsSteeringPolicyEmpty:
		return true
	}
	return false
}

type LoadBalancerNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   LoadBalancer                                              `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerNewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerNewResponseEnvelope]
type loadBalancerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerNewResponseEnvelopeSuccess bool

const (
	LoadBalancerNewResponseEnvelopeSuccessTrue LoadBalancerNewResponseEnvelopeSuccess = true
)

func (r LoadBalancerNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerUpdateParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]string] `json:"default_pools,required"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[interface{}] `json:"fallback_pool,required"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name param.Field[string] `json:"name,required"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[LoadBalancerUpdateParamsAdaptiveRouting] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[interface{}] `json:"country_pools"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) this load balancer.
	Enabled param.Field[bool] `json:"enabled"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LoadBalancerUpdateParamsLocationStrategy] `json:"location_strategy"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools param.Field[interface{}] `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied param.Field[bool] `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[LoadBalancerUpdateParamsRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]LoadBalancerUpdateParamsRule] `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"` or "" (default). The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity param.Field[LoadBalancerUpdateParamsSessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[LoadBalancerUpdateParamsSessionAffinityAttributes] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTTL param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[LoadBalancerUpdateParamsSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r LoadBalancerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerUpdateParamsAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r LoadBalancerUpdateParamsAdaptiveRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerUpdateParamsLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[LoadBalancerUpdateParamsLocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[LoadBalancerUpdateParamsLocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r LoadBalancerUpdateParamsLocationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerUpdateParamsLocationStrategyMode string

const (
	LoadBalancerUpdateParamsLocationStrategyModePop        LoadBalancerUpdateParamsLocationStrategyMode = "pop"
	LoadBalancerUpdateParamsLocationStrategyModeResolverIP LoadBalancerUpdateParamsLocationStrategyMode = "resolver_ip"
)

func (r LoadBalancerUpdateParamsLocationStrategyMode) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsLocationStrategyModePop, LoadBalancerUpdateParamsLocationStrategyModeResolverIP:
		return true
	}
	return false
}

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerUpdateParamsLocationStrategyPreferEcs string

const (
	LoadBalancerUpdateParamsLocationStrategyPreferEcsAlways    LoadBalancerUpdateParamsLocationStrategyPreferEcs = "always"
	LoadBalancerUpdateParamsLocationStrategyPreferEcsNever     LoadBalancerUpdateParamsLocationStrategyPreferEcs = "never"
	LoadBalancerUpdateParamsLocationStrategyPreferEcsProximity LoadBalancerUpdateParamsLocationStrategyPreferEcs = "proximity"
	LoadBalancerUpdateParamsLocationStrategyPreferEcsGeo       LoadBalancerUpdateParamsLocationStrategyPreferEcs = "geo"
)

func (r LoadBalancerUpdateParamsLocationStrategyPreferEcs) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsLocationStrategyPreferEcsAlways, LoadBalancerUpdateParamsLocationStrategyPreferEcsNever, LoadBalancerUpdateParamsLocationStrategyPreferEcsProximity, LoadBalancerUpdateParamsLocationStrategyPreferEcsGeo:
		return true
	}
	return false
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerUpdateParamsRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[interface{}] `json:"pool_weights"`
}

func (r LoadBalancerUpdateParamsRandomSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerUpdateParamsRule struct {
	// The condition expressions to evaluate. If the condition evaluates to true, the
	// overrides or fixed_response in this rule will be applied. An empty condition is
	// always true. For more details on condition expressions, please see
	// https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.
	Condition param.Field[string] `json:"condition"`
	// Disable this specific rule. It will no longer be evaluated by this load
	// balancer.
	Disabled param.Field[bool] `json:"disabled"`
	// A collection of fields used to directly respond to the eyeball instead of
	// routing to a pool. If a fixed_response is supplied the rule will be marked as
	// terminates.
	FixedResponse param.Field[LoadBalancerUpdateParamsRulesFixedResponse] `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides param.Field[LoadBalancerUpdateParamsRulesOverrides] `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority param.Field[int64] `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates param.Field[bool] `json:"terminates"`
}

func (r LoadBalancerUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerUpdateParamsRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType param.Field[string] `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location param.Field[string] `json:"location"`
	// Text to include as the http body.
	MessageBody param.Field[string] `json:"message_body"`
	// The http status code to respond with.
	StatusCode param.Field[int64] `json:"status_code"`
}

func (r LoadBalancerUpdateParamsRulesFixedResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerUpdateParamsRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[LoadBalancerUpdateParamsRulesOverridesAdaptiveRouting] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[interface{}] `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]string] `json:"default_pools"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[interface{}] `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LoadBalancerUpdateParamsRulesOverridesLocationStrategy] `json:"location_strategy"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools param.Field[interface{}] `json:"pop_pools"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[LoadBalancerUpdateParamsRulesOverridesRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"` or "" (default). The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity param.Field[LoadBalancerUpdateParamsRulesOverridesSessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTTL param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[LoadBalancerUpdateParamsRulesOverridesSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r LoadBalancerUpdateParamsRulesOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerUpdateParamsRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r LoadBalancerUpdateParamsRulesOverridesAdaptiveRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerUpdateParamsRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[LoadBalancerUpdateParamsRulesOverridesLocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r LoadBalancerUpdateParamsRulesOverridesLocationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerUpdateParamsRulesOverridesLocationStrategyMode string

const (
	LoadBalancerUpdateParamsRulesOverridesLocationStrategyModePop        LoadBalancerUpdateParamsRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP LoadBalancerUpdateParamsRulesOverridesLocationStrategyMode = "resolver_ip"
)

func (r LoadBalancerUpdateParamsRulesOverridesLocationStrategyMode) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsRulesOverridesLocationStrategyModePop, LoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP:
		return true
	}
	return false
}

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs = "geo"
)

func (r LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways, LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsNever, LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsProximity, LoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsGeo:
		return true
	}
	return false
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerUpdateParamsRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[interface{}] `json:"pool_weights"`
}

func (r LoadBalancerUpdateParamsRulesOverridesRandomSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of session affinity the load balancer should use unless
// specified as `"none"` or "" (default). The supported types are:
//
//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
//     generated, encoding information of which origin the request will be forwarded
//     to. Subsequent requests, by the same client to the same load balancer, will be
//     sent to the origin server the cookie encodes, for the duration of the cookie
//     and as long as the origin server remains healthy. If the cookie has expired or
//     the origin server is unhealthy, then a new origin server is calculated and
//     used.
//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
//     selection is stable and based on the client's ip address.
//   - `"header"`: On the first request to a proxied load balancer, a session key
//     based on the configured HTTP headers (see
//     `session_affinity_attributes.headers`) is generated, encoding the request
//     headers used for storing in the load balancer session state which origin the
//     request will be forwarded to. Subsequent requests to the load balancer with
//     the same headers will be sent to the same origin server, for the duration of
//     the session and as long as the origin server remains healthy. If the session
//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
//     server is unhealthy, then a new origin server is calculated and used. See
//     `headers` in `session_affinity_attributes` for additional required
//     configuration.
type LoadBalancerUpdateParamsRulesOverridesSessionAffinity string

const (
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityNone     LoadBalancerUpdateParamsRulesOverridesSessionAffinity = "none"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie   LoadBalancerUpdateParamsRulesOverridesSessionAffinity = "cookie"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityIPCookie LoadBalancerUpdateParamsRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityHeader   LoadBalancerUpdateParamsRulesOverridesSessionAffinity = "header"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityEmpty    LoadBalancerUpdateParamsRulesOverridesSessionAffinity = "\"\""
)

func (r LoadBalancerUpdateParamsRulesOverridesSessionAffinity) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsRulesOverridesSessionAffinityNone, LoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie, LoadBalancerUpdateParamsRulesOverridesSessionAffinityIPCookie, LoadBalancerUpdateParamsRulesOverridesSessionAffinityHeader, LoadBalancerUpdateParamsRulesOverridesSessionAffinityEmpty:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers param.Field[[]string] `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are:
	//
	//   - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	//     specified by the `headers` session affinity attribute, otherwise sessions
	//     aren't created.
	//   - `"false"`: Load balancing requests must contain _at least one_ of the HTTP
	//     headers specified by the `headers` session affinity attribute, otherwise
	//     sessions aren't created.
	RequireAllHeaders param.Field[bool] `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite param.Field[LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure] `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are:
	//
	//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
	//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
	//     originally pinned origin is available; note that this can potentially result
	//     in heavy origin flapping.
	//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
	//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
	//     currently not supported for session affinity by header.
	ZeroDowntimeFailover param.Field[LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

func (r LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto, LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteLax, LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteNone, LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure = "Never"
)

func (r LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto, LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAlways, LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureNever:
		return true
	}
	return false
}

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. This feature is currently incompatible with Argo, Tiered
// Cache, and Bandwidth Alliance. The supported values are:
//
//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
//     originally pinned origin is available; note that this can potentially result
//     in heavy origin flapping.
//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
//     currently not supported for session affinity by header.
type LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

func (r LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone, LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary, LoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky:
		return true
	}
	return false
}

// Steering Policy for this load balancer.
//
//   - `"off"`: Use `default_pools`.
//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
//     requests, the country for `country_pools` is determined by
//     `location_strategy`.
//   - `"random"`: Select a pool randomly.
//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
//     default_pools (requires pool health checks).
//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
//     pool using the Cloudflare PoP location for proxied requests or the location
//     determined by `location_strategy` for non-proxied requests.
//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of outstanding
//     requests. Pools with more pending requests are weighted proportionately less
//     relative to others.
//   - `"least_connections"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of open connections.
//     Pools with more open connections are weighted proportionately less relative to
//     others. Supported for HTTP/1 and HTTP/2 connections.
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerUpdateParamsRulesOverridesSteeringPolicy string

const (
	LoadBalancerUpdateParamsRulesOverridesSteeringPolicyOff                      LoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "off"
	LoadBalancerUpdateParamsRulesOverridesSteeringPolicyGeo                      LoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "geo"
	LoadBalancerUpdateParamsRulesOverridesSteeringPolicyRandom                   LoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "random"
	LoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerUpdateParamsRulesOverridesSteeringPolicyProximity                LoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerUpdateParamsRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerUpdateParamsRulesOverridesSteeringPolicyLeastConnections         LoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerUpdateParamsRulesOverridesSteeringPolicyEmpty                    LoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "\"\""
)

func (r LoadBalancerUpdateParamsRulesOverridesSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsRulesOverridesSteeringPolicyOff, LoadBalancerUpdateParamsRulesOverridesSteeringPolicyGeo, LoadBalancerUpdateParamsRulesOverridesSteeringPolicyRandom, LoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency, LoadBalancerUpdateParamsRulesOverridesSteeringPolicyProximity, LoadBalancerUpdateParamsRulesOverridesSteeringPolicyLeastOutstandingRequests, LoadBalancerUpdateParamsRulesOverridesSteeringPolicyLeastConnections, LoadBalancerUpdateParamsRulesOverridesSteeringPolicyEmpty:
		return true
	}
	return false
}

// Specifies the type of session affinity the load balancer should use unless
// specified as `"none"` or "" (default). The supported types are:
//
//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
//     generated, encoding information of which origin the request will be forwarded
//     to. Subsequent requests, by the same client to the same load balancer, will be
//     sent to the origin server the cookie encodes, for the duration of the cookie
//     and as long as the origin server remains healthy. If the cookie has expired or
//     the origin server is unhealthy, then a new origin server is calculated and
//     used.
//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
//     selection is stable and based on the client's ip address.
//   - `"header"`: On the first request to a proxied load balancer, a session key
//     based on the configured HTTP headers (see
//     `session_affinity_attributes.headers`) is generated, encoding the request
//     headers used for storing in the load balancer session state which origin the
//     request will be forwarded to. Subsequent requests to the load balancer with
//     the same headers will be sent to the same origin server, for the duration of
//     the session and as long as the origin server remains healthy. If the session
//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
//     server is unhealthy, then a new origin server is calculated and used. See
//     `headers` in `session_affinity_attributes` for additional required
//     configuration.
type LoadBalancerUpdateParamsSessionAffinity string

const (
	LoadBalancerUpdateParamsSessionAffinityNone     LoadBalancerUpdateParamsSessionAffinity = "none"
	LoadBalancerUpdateParamsSessionAffinityCookie   LoadBalancerUpdateParamsSessionAffinity = "cookie"
	LoadBalancerUpdateParamsSessionAffinityIPCookie LoadBalancerUpdateParamsSessionAffinity = "ip_cookie"
	LoadBalancerUpdateParamsSessionAffinityHeader   LoadBalancerUpdateParamsSessionAffinity = "header"
	LoadBalancerUpdateParamsSessionAffinityEmpty    LoadBalancerUpdateParamsSessionAffinity = "\"\""
)

func (r LoadBalancerUpdateParamsSessionAffinity) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsSessionAffinityNone, LoadBalancerUpdateParamsSessionAffinityCookie, LoadBalancerUpdateParamsSessionAffinityIPCookie, LoadBalancerUpdateParamsSessionAffinityHeader, LoadBalancerUpdateParamsSessionAffinityEmpty:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type LoadBalancerUpdateParamsSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers param.Field[[]string] `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are:
	//
	//   - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	//     specified by the `headers` session affinity attribute, otherwise sessions
	//     aren't created.
	//   - `"false"`: Load balancing requests must contain _at least one_ of the HTTP
	//     headers specified by the `headers` session affinity attribute, otherwise
	//     sessions aren't created.
	RequireAllHeaders param.Field[bool] `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite param.Field[LoadBalancerUpdateParamsSessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[LoadBalancerUpdateParamsSessionAffinityAttributesSecure] `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are:
	//
	//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
	//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
	//     originally pinned origin is available; note that this can potentially result
	//     in heavy origin flapping.
	//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
	//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
	//     currently not supported for session affinity by header.
	ZeroDowntimeFailover param.Field[LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r LoadBalancerUpdateParamsSessionAffinityAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerUpdateParamsSessionAffinityAttributesSamesite string

const (
	LoadBalancerUpdateParamsSessionAffinityAttributesSamesiteAuto   LoadBalancerUpdateParamsSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerUpdateParamsSessionAffinityAttributesSamesiteLax    LoadBalancerUpdateParamsSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerUpdateParamsSessionAffinityAttributesSamesiteNone   LoadBalancerUpdateParamsSessionAffinityAttributesSamesite = "None"
	LoadBalancerUpdateParamsSessionAffinityAttributesSamesiteStrict LoadBalancerUpdateParamsSessionAffinityAttributesSamesite = "Strict"
)

func (r LoadBalancerUpdateParamsSessionAffinityAttributesSamesite) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsSessionAffinityAttributesSamesiteAuto, LoadBalancerUpdateParamsSessionAffinityAttributesSamesiteLax, LoadBalancerUpdateParamsSessionAffinityAttributesSamesiteNone, LoadBalancerUpdateParamsSessionAffinityAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerUpdateParamsSessionAffinityAttributesSecure string

const (
	LoadBalancerUpdateParamsSessionAffinityAttributesSecureAuto   LoadBalancerUpdateParamsSessionAffinityAttributesSecure = "Auto"
	LoadBalancerUpdateParamsSessionAffinityAttributesSecureAlways LoadBalancerUpdateParamsSessionAffinityAttributesSecure = "Always"
	LoadBalancerUpdateParamsSessionAffinityAttributesSecureNever  LoadBalancerUpdateParamsSessionAffinityAttributesSecure = "Never"
)

func (r LoadBalancerUpdateParamsSessionAffinityAttributesSecure) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsSessionAffinityAttributesSecureAuto, LoadBalancerUpdateParamsSessionAffinityAttributesSecureAlways, LoadBalancerUpdateParamsSessionAffinityAttributesSecureNever:
		return true
	}
	return false
}

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. This feature is currently incompatible with Argo, Tiered
// Cache, and Bandwidth Alliance. The supported values are:
//
//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
//     originally pinned origin is available; note that this can potentially result
//     in heavy origin flapping.
//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
//     currently not supported for session affinity by header.
type LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

func (r LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverNone, LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverTemporary, LoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverSticky:
		return true
	}
	return false
}

// Steering Policy for this load balancer.
//
//   - `"off"`: Use `default_pools`.
//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
//     requests, the country for `country_pools` is determined by
//     `location_strategy`.
//   - `"random"`: Select a pool randomly.
//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
//     default_pools (requires pool health checks).
//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
//     pool using the Cloudflare PoP location for proxied requests or the location
//     determined by `location_strategy` for non-proxied requests.
//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of outstanding
//     requests. Pools with more pending requests are weighted proportionately less
//     relative to others.
//   - `"least_connections"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of open connections.
//     Pools with more open connections are weighted proportionately less relative to
//     others. Supported for HTTP/1 and HTTP/2 connections.
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerUpdateParamsSteeringPolicy string

const (
	LoadBalancerUpdateParamsSteeringPolicyOff                      LoadBalancerUpdateParamsSteeringPolicy = "off"
	LoadBalancerUpdateParamsSteeringPolicyGeo                      LoadBalancerUpdateParamsSteeringPolicy = "geo"
	LoadBalancerUpdateParamsSteeringPolicyRandom                   LoadBalancerUpdateParamsSteeringPolicy = "random"
	LoadBalancerUpdateParamsSteeringPolicyDynamicLatency           LoadBalancerUpdateParamsSteeringPolicy = "dynamic_latency"
	LoadBalancerUpdateParamsSteeringPolicyProximity                LoadBalancerUpdateParamsSteeringPolicy = "proximity"
	LoadBalancerUpdateParamsSteeringPolicyLeastOutstandingRequests LoadBalancerUpdateParamsSteeringPolicy = "least_outstanding_requests"
	LoadBalancerUpdateParamsSteeringPolicyLeastConnections         LoadBalancerUpdateParamsSteeringPolicy = "least_connections"
	LoadBalancerUpdateParamsSteeringPolicyEmpty                    LoadBalancerUpdateParamsSteeringPolicy = "\"\""
)

func (r LoadBalancerUpdateParamsSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateParamsSteeringPolicyOff, LoadBalancerUpdateParamsSteeringPolicyGeo, LoadBalancerUpdateParamsSteeringPolicyRandom, LoadBalancerUpdateParamsSteeringPolicyDynamicLatency, LoadBalancerUpdateParamsSteeringPolicyProximity, LoadBalancerUpdateParamsSteeringPolicyLeastOutstandingRequests, LoadBalancerUpdateParamsSteeringPolicyLeastConnections, LoadBalancerUpdateParamsSteeringPolicyEmpty:
		return true
	}
	return false
}

type LoadBalancerUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   LoadBalancer                                              `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerUpdateResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerUpdateResponseEnvelope]
type loadBalancerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerUpdateResponseEnvelopeSuccess bool

const (
	LoadBalancerUpdateResponseEnvelopeSuccessTrue LoadBalancerUpdateResponseEnvelopeSuccess = true
)

func (r LoadBalancerUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerListParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LoadBalancerDeleteParams struct {
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r LoadBalancerDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LoadBalancerDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   LoadBalancerDeleteResponse                                `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerDeleteResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerDeleteResponseEnvelope]
type loadBalancerDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerDeleteResponseEnvelopeSuccess bool

const (
	LoadBalancerDeleteResponseEnvelopeSuccessTrue LoadBalancerDeleteResponseEnvelopeSuccess = true
)

func (r LoadBalancerDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerEditParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[LoadBalancerEditParamsAdaptiveRouting] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[interface{}] `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]string] `json:"default_pools"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) this load balancer.
	Enabled param.Field[bool] `json:"enabled"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[interface{}] `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LoadBalancerEditParamsLocationStrategy] `json:"location_strategy"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name param.Field[string] `json:"name"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools param.Field[interface{}] `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied param.Field[bool] `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[LoadBalancerEditParamsRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]LoadBalancerEditParamsRule] `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"` or "" (default). The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity param.Field[LoadBalancerEditParamsSessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[LoadBalancerEditParamsSessionAffinityAttributes] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTTL param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[LoadBalancerEditParamsSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r LoadBalancerEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerEditParamsAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r LoadBalancerEditParamsAdaptiveRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerEditParamsLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[LoadBalancerEditParamsLocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[LoadBalancerEditParamsLocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r LoadBalancerEditParamsLocationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerEditParamsLocationStrategyMode string

const (
	LoadBalancerEditParamsLocationStrategyModePop        LoadBalancerEditParamsLocationStrategyMode = "pop"
	LoadBalancerEditParamsLocationStrategyModeResolverIP LoadBalancerEditParamsLocationStrategyMode = "resolver_ip"
)

func (r LoadBalancerEditParamsLocationStrategyMode) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsLocationStrategyModePop, LoadBalancerEditParamsLocationStrategyModeResolverIP:
		return true
	}
	return false
}

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerEditParamsLocationStrategyPreferEcs string

const (
	LoadBalancerEditParamsLocationStrategyPreferEcsAlways    LoadBalancerEditParamsLocationStrategyPreferEcs = "always"
	LoadBalancerEditParamsLocationStrategyPreferEcsNever     LoadBalancerEditParamsLocationStrategyPreferEcs = "never"
	LoadBalancerEditParamsLocationStrategyPreferEcsProximity LoadBalancerEditParamsLocationStrategyPreferEcs = "proximity"
	LoadBalancerEditParamsLocationStrategyPreferEcsGeo       LoadBalancerEditParamsLocationStrategyPreferEcs = "geo"
)

func (r LoadBalancerEditParamsLocationStrategyPreferEcs) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsLocationStrategyPreferEcsAlways, LoadBalancerEditParamsLocationStrategyPreferEcsNever, LoadBalancerEditParamsLocationStrategyPreferEcsProximity, LoadBalancerEditParamsLocationStrategyPreferEcsGeo:
		return true
	}
	return false
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerEditParamsRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[interface{}] `json:"pool_weights"`
}

func (r LoadBalancerEditParamsRandomSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerEditParamsRule struct {
	// The condition expressions to evaluate. If the condition evaluates to true, the
	// overrides or fixed_response in this rule will be applied. An empty condition is
	// always true. For more details on condition expressions, please see
	// https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.
	Condition param.Field[string] `json:"condition"`
	// Disable this specific rule. It will no longer be evaluated by this load
	// balancer.
	Disabled param.Field[bool] `json:"disabled"`
	// A collection of fields used to directly respond to the eyeball instead of
	// routing to a pool. If a fixed_response is supplied the rule will be marked as
	// terminates.
	FixedResponse param.Field[LoadBalancerEditParamsRulesFixedResponse] `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides param.Field[LoadBalancerEditParamsRulesOverrides] `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority param.Field[int64] `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates param.Field[bool] `json:"terminates"`
}

func (r LoadBalancerEditParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerEditParamsRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType param.Field[string] `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location param.Field[string] `json:"location"`
	// Text to include as the http body.
	MessageBody param.Field[string] `json:"message_body"`
	// The http status code to respond with.
	StatusCode param.Field[int64] `json:"status_code"`
}

func (r LoadBalancerEditParamsRulesFixedResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerEditParamsRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[LoadBalancerEditParamsRulesOverridesAdaptiveRouting] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[interface{}] `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]string] `json:"default_pools"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[interface{}] `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LoadBalancerEditParamsRulesOverridesLocationStrategy] `json:"location_strategy"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools param.Field[interface{}] `json:"pop_pools"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[LoadBalancerEditParamsRulesOverridesRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"` or "" (default). The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity param.Field[LoadBalancerEditParamsRulesOverridesSessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[LoadBalancerEditParamsRulesOverridesSessionAffinityAttributes] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTTL param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[LoadBalancerEditParamsRulesOverridesSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r LoadBalancerEditParamsRulesOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerEditParamsRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r LoadBalancerEditParamsRulesOverridesAdaptiveRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerEditParamsRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[LoadBalancerEditParamsRulesOverridesLocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r LoadBalancerEditParamsRulesOverridesLocationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerEditParamsRulesOverridesLocationStrategyMode string

const (
	LoadBalancerEditParamsRulesOverridesLocationStrategyModePop        LoadBalancerEditParamsRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerEditParamsRulesOverridesLocationStrategyModeResolverIP LoadBalancerEditParamsRulesOverridesLocationStrategyMode = "resolver_ip"
)

func (r LoadBalancerEditParamsRulesOverridesLocationStrategyMode) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsRulesOverridesLocationStrategyModePop, LoadBalancerEditParamsRulesOverridesLocationStrategyModeResolverIP:
		return true
	}
	return false
}

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcs = "geo"
)

func (r LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcs) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsAlways, LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsNever, LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsProximity, LoadBalancerEditParamsRulesOverridesLocationStrategyPreferEcsGeo:
		return true
	}
	return false
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerEditParamsRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[interface{}] `json:"pool_weights"`
}

func (r LoadBalancerEditParamsRulesOverridesRandomSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of session affinity the load balancer should use unless
// specified as `"none"` or "" (default). The supported types are:
//
//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
//     generated, encoding information of which origin the request will be forwarded
//     to. Subsequent requests, by the same client to the same load balancer, will be
//     sent to the origin server the cookie encodes, for the duration of the cookie
//     and as long as the origin server remains healthy. If the cookie has expired or
//     the origin server is unhealthy, then a new origin server is calculated and
//     used.
//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
//     selection is stable and based on the client's ip address.
//   - `"header"`: On the first request to a proxied load balancer, a session key
//     based on the configured HTTP headers (see
//     `session_affinity_attributes.headers`) is generated, encoding the request
//     headers used for storing in the load balancer session state which origin the
//     request will be forwarded to. Subsequent requests to the load balancer with
//     the same headers will be sent to the same origin server, for the duration of
//     the session and as long as the origin server remains healthy. If the session
//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
//     server is unhealthy, then a new origin server is calculated and used. See
//     `headers` in `session_affinity_attributes` for additional required
//     configuration.
type LoadBalancerEditParamsRulesOverridesSessionAffinity string

const (
	LoadBalancerEditParamsRulesOverridesSessionAffinityNone     LoadBalancerEditParamsRulesOverridesSessionAffinity = "none"
	LoadBalancerEditParamsRulesOverridesSessionAffinityCookie   LoadBalancerEditParamsRulesOverridesSessionAffinity = "cookie"
	LoadBalancerEditParamsRulesOverridesSessionAffinityIPCookie LoadBalancerEditParamsRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerEditParamsRulesOverridesSessionAffinityHeader   LoadBalancerEditParamsRulesOverridesSessionAffinity = "header"
	LoadBalancerEditParamsRulesOverridesSessionAffinityEmpty    LoadBalancerEditParamsRulesOverridesSessionAffinity = "\"\""
)

func (r LoadBalancerEditParamsRulesOverridesSessionAffinity) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsRulesOverridesSessionAffinityNone, LoadBalancerEditParamsRulesOverridesSessionAffinityCookie, LoadBalancerEditParamsRulesOverridesSessionAffinityIPCookie, LoadBalancerEditParamsRulesOverridesSessionAffinityHeader, LoadBalancerEditParamsRulesOverridesSessionAffinityEmpty:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type LoadBalancerEditParamsRulesOverridesSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers param.Field[[]string] `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are:
	//
	//   - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	//     specified by the `headers` session affinity attribute, otherwise sessions
	//     aren't created.
	//   - `"false"`: Load balancing requests must contain _at least one_ of the HTTP
	//     headers specified by the `headers` session affinity attribute, otherwise
	//     sessions aren't created.
	RequireAllHeaders param.Field[bool] `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite param.Field[LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecure] `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are:
	//
	//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
	//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
	//     originally pinned origin is available; note that this can potentially result
	//     in heavy origin flapping.
	//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
	//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
	//     currently not supported for session affinity by header.
	ZeroDowntimeFailover param.Field[LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r LoadBalancerEditParamsRulesOverridesSessionAffinityAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

func (r LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesite) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteAuto, LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteLax, LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteNone, LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecure = "Never"
)

func (r LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecure) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureAuto, LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureAlways, LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesSecureNever:
		return true
	}
	return false
}

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. This feature is currently incompatible with Argo, Tiered
// Cache, and Bandwidth Alliance. The supported values are:
//
//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
//     originally pinned origin is available; note that this can potentially result
//     in heavy origin flapping.
//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
//     currently not supported for session affinity by header.
type LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

func (r LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone, LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary, LoadBalancerEditParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky:
		return true
	}
	return false
}

// Steering Policy for this load balancer.
//
//   - `"off"`: Use `default_pools`.
//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
//     requests, the country for `country_pools` is determined by
//     `location_strategy`.
//   - `"random"`: Select a pool randomly.
//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
//     default_pools (requires pool health checks).
//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
//     pool using the Cloudflare PoP location for proxied requests or the location
//     determined by `location_strategy` for non-proxied requests.
//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of outstanding
//     requests. Pools with more pending requests are weighted proportionately less
//     relative to others.
//   - `"least_connections"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of open connections.
//     Pools with more open connections are weighted proportionately less relative to
//     others. Supported for HTTP/1 and HTTP/2 connections.
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerEditParamsRulesOverridesSteeringPolicy string

const (
	LoadBalancerEditParamsRulesOverridesSteeringPolicyOff                      LoadBalancerEditParamsRulesOverridesSteeringPolicy = "off"
	LoadBalancerEditParamsRulesOverridesSteeringPolicyGeo                      LoadBalancerEditParamsRulesOverridesSteeringPolicy = "geo"
	LoadBalancerEditParamsRulesOverridesSteeringPolicyRandom                   LoadBalancerEditParamsRulesOverridesSteeringPolicy = "random"
	LoadBalancerEditParamsRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerEditParamsRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerEditParamsRulesOverridesSteeringPolicyProximity                LoadBalancerEditParamsRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerEditParamsRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerEditParamsRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerEditParamsRulesOverridesSteeringPolicyLeastConnections         LoadBalancerEditParamsRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerEditParamsRulesOverridesSteeringPolicyEmpty                    LoadBalancerEditParamsRulesOverridesSteeringPolicy = "\"\""
)

func (r LoadBalancerEditParamsRulesOverridesSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsRulesOverridesSteeringPolicyOff, LoadBalancerEditParamsRulesOverridesSteeringPolicyGeo, LoadBalancerEditParamsRulesOverridesSteeringPolicyRandom, LoadBalancerEditParamsRulesOverridesSteeringPolicyDynamicLatency, LoadBalancerEditParamsRulesOverridesSteeringPolicyProximity, LoadBalancerEditParamsRulesOverridesSteeringPolicyLeastOutstandingRequests, LoadBalancerEditParamsRulesOverridesSteeringPolicyLeastConnections, LoadBalancerEditParamsRulesOverridesSteeringPolicyEmpty:
		return true
	}
	return false
}

// Specifies the type of session affinity the load balancer should use unless
// specified as `"none"` or "" (default). The supported types are:
//
//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
//     generated, encoding information of which origin the request will be forwarded
//     to. Subsequent requests, by the same client to the same load balancer, will be
//     sent to the origin server the cookie encodes, for the duration of the cookie
//     and as long as the origin server remains healthy. If the cookie has expired or
//     the origin server is unhealthy, then a new origin server is calculated and
//     used.
//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
//     selection is stable and based on the client's ip address.
//   - `"header"`: On the first request to a proxied load balancer, a session key
//     based on the configured HTTP headers (see
//     `session_affinity_attributes.headers`) is generated, encoding the request
//     headers used for storing in the load balancer session state which origin the
//     request will be forwarded to. Subsequent requests to the load balancer with
//     the same headers will be sent to the same origin server, for the duration of
//     the session and as long as the origin server remains healthy. If the session
//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
//     server is unhealthy, then a new origin server is calculated and used. See
//     `headers` in `session_affinity_attributes` for additional required
//     configuration.
type LoadBalancerEditParamsSessionAffinity string

const (
	LoadBalancerEditParamsSessionAffinityNone     LoadBalancerEditParamsSessionAffinity = "none"
	LoadBalancerEditParamsSessionAffinityCookie   LoadBalancerEditParamsSessionAffinity = "cookie"
	LoadBalancerEditParamsSessionAffinityIPCookie LoadBalancerEditParamsSessionAffinity = "ip_cookie"
	LoadBalancerEditParamsSessionAffinityHeader   LoadBalancerEditParamsSessionAffinity = "header"
	LoadBalancerEditParamsSessionAffinityEmpty    LoadBalancerEditParamsSessionAffinity = "\"\""
)

func (r LoadBalancerEditParamsSessionAffinity) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsSessionAffinityNone, LoadBalancerEditParamsSessionAffinityCookie, LoadBalancerEditParamsSessionAffinityIPCookie, LoadBalancerEditParamsSessionAffinityHeader, LoadBalancerEditParamsSessionAffinityEmpty:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type LoadBalancerEditParamsSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers param.Field[[]string] `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are:
	//
	//   - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	//     specified by the `headers` session affinity attribute, otherwise sessions
	//     aren't created.
	//   - `"false"`: Load balancing requests must contain _at least one_ of the HTTP
	//     headers specified by the `headers` session affinity attribute, otherwise
	//     sessions aren't created.
	RequireAllHeaders param.Field[bool] `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite param.Field[LoadBalancerEditParamsSessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[LoadBalancerEditParamsSessionAffinityAttributesSecure] `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are:
	//
	//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
	//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
	//     originally pinned origin is available; note that this can potentially result
	//     in heavy origin flapping.
	//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
	//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
	//     currently not supported for session affinity by header.
	ZeroDowntimeFailover param.Field[LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r LoadBalancerEditParamsSessionAffinityAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerEditParamsSessionAffinityAttributesSamesite string

const (
	LoadBalancerEditParamsSessionAffinityAttributesSamesiteAuto   LoadBalancerEditParamsSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerEditParamsSessionAffinityAttributesSamesiteLax    LoadBalancerEditParamsSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerEditParamsSessionAffinityAttributesSamesiteNone   LoadBalancerEditParamsSessionAffinityAttributesSamesite = "None"
	LoadBalancerEditParamsSessionAffinityAttributesSamesiteStrict LoadBalancerEditParamsSessionAffinityAttributesSamesite = "Strict"
)

func (r LoadBalancerEditParamsSessionAffinityAttributesSamesite) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsSessionAffinityAttributesSamesiteAuto, LoadBalancerEditParamsSessionAffinityAttributesSamesiteLax, LoadBalancerEditParamsSessionAffinityAttributesSamesiteNone, LoadBalancerEditParamsSessionAffinityAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerEditParamsSessionAffinityAttributesSecure string

const (
	LoadBalancerEditParamsSessionAffinityAttributesSecureAuto   LoadBalancerEditParamsSessionAffinityAttributesSecure = "Auto"
	LoadBalancerEditParamsSessionAffinityAttributesSecureAlways LoadBalancerEditParamsSessionAffinityAttributesSecure = "Always"
	LoadBalancerEditParamsSessionAffinityAttributesSecureNever  LoadBalancerEditParamsSessionAffinityAttributesSecure = "Never"
)

func (r LoadBalancerEditParamsSessionAffinityAttributesSecure) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsSessionAffinityAttributesSecureAuto, LoadBalancerEditParamsSessionAffinityAttributesSecureAlways, LoadBalancerEditParamsSessionAffinityAttributesSecureNever:
		return true
	}
	return false
}

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. This feature is currently incompatible with Argo, Tiered
// Cache, and Bandwidth Alliance. The supported values are:
//
//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
//     originally pinned origin is available; note that this can potentially result
//     in heavy origin flapping.
//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
//     currently not supported for session affinity by header.
type LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

func (r LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailover) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailoverNone, LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailoverTemporary, LoadBalancerEditParamsSessionAffinityAttributesZeroDowntimeFailoverSticky:
		return true
	}
	return false
}

// Steering Policy for this load balancer.
//
//   - `"off"`: Use `default_pools`.
//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
//     requests, the country for `country_pools` is determined by
//     `location_strategy`.
//   - `"random"`: Select a pool randomly.
//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
//     default_pools (requires pool health checks).
//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
//     pool using the Cloudflare PoP location for proxied requests or the location
//     determined by `location_strategy` for non-proxied requests.
//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of outstanding
//     requests. Pools with more pending requests are weighted proportionately less
//     relative to others.
//   - `"least_connections"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of open connections.
//     Pools with more open connections are weighted proportionately less relative to
//     others. Supported for HTTP/1 and HTTP/2 connections.
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerEditParamsSteeringPolicy string

const (
	LoadBalancerEditParamsSteeringPolicyOff                      LoadBalancerEditParamsSteeringPolicy = "off"
	LoadBalancerEditParamsSteeringPolicyGeo                      LoadBalancerEditParamsSteeringPolicy = "geo"
	LoadBalancerEditParamsSteeringPolicyRandom                   LoadBalancerEditParamsSteeringPolicy = "random"
	LoadBalancerEditParamsSteeringPolicyDynamicLatency           LoadBalancerEditParamsSteeringPolicy = "dynamic_latency"
	LoadBalancerEditParamsSteeringPolicyProximity                LoadBalancerEditParamsSteeringPolicy = "proximity"
	LoadBalancerEditParamsSteeringPolicyLeastOutstandingRequests LoadBalancerEditParamsSteeringPolicy = "least_outstanding_requests"
	LoadBalancerEditParamsSteeringPolicyLeastConnections         LoadBalancerEditParamsSteeringPolicy = "least_connections"
	LoadBalancerEditParamsSteeringPolicyEmpty                    LoadBalancerEditParamsSteeringPolicy = "\"\""
)

func (r LoadBalancerEditParamsSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerEditParamsSteeringPolicyOff, LoadBalancerEditParamsSteeringPolicyGeo, LoadBalancerEditParamsSteeringPolicyRandom, LoadBalancerEditParamsSteeringPolicyDynamicLatency, LoadBalancerEditParamsSteeringPolicyProximity, LoadBalancerEditParamsSteeringPolicyLeastOutstandingRequests, LoadBalancerEditParamsSteeringPolicyLeastConnections, LoadBalancerEditParamsSteeringPolicyEmpty:
		return true
	}
	return false
}

type LoadBalancerEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   LoadBalancer                                              `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerEditResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerEditResponseEnvelope]
type loadBalancerEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerEditResponseEnvelopeSuccess bool

const (
	LoadBalancerEditResponseEnvelopeSuccessTrue LoadBalancerEditResponseEnvelopeSuccess = true
)

func (r LoadBalancerEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerGetParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LoadBalancerGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   LoadBalancer                                              `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerGetResponseEnvelope]
type loadBalancerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerGetResponseEnvelopeSuccess bool

const (
	LoadBalancerGetResponseEnvelopeSuccessTrue LoadBalancerGetResponseEnvelopeSuccess = true
)

func (r LoadBalancerGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
