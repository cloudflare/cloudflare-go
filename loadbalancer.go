// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// LoadBalancerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerService] method
// instead.
type LoadBalancerService struct {
	Options  []option.RequestOption
	Monitors *LoadBalancerMonitorService
	Pools    *LoadBalancerPoolService
	Previews *LoadBalancerPreviewService
	Regions  *LoadBalancerRegionService
	Searches *LoadBalancerSearchService
}

// NewLoadBalancerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLoadBalancerService(opts ...option.RequestOption) (r *LoadBalancerService) {
	r = &LoadBalancerService{}
	r.Options = opts
	r.Monitors = NewLoadBalancerMonitorService(opts...)
	r.Pools = NewLoadBalancerPoolService(opts...)
	r.Previews = NewLoadBalancerPreviewService(opts...)
	r.Regions = NewLoadBalancerRegionService(opts...)
	r.Searches = NewLoadBalancerSearchService(opts...)
	return
}

// Create a new load balancer.
func (r *LoadBalancerService) New(ctx context.Context, identifier string, body LoadBalancerNewParams, opts ...option.RequestOption) (res *LoadBalancerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured load balancer.
func (r *LoadBalancerService) Get(ctx context.Context, identifier1 string, identifier string, opts ...option.RequestOption) (res *LoadBalancerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/load_balancers/%s", identifier1, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured load balancer.
func (r *LoadBalancerService) Update(ctx context.Context, identifier1 string, identifier string, body LoadBalancerUpdateParams, opts ...option.RequestOption) (res *LoadBalancerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/load_balancers/%s", identifier1, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List configured load balancers.
func (r *LoadBalancerService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *LoadBalancerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/load_balancers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a configured load balancer.
func (r *LoadBalancerService) Delete(ctx context.Context, identifier1 string, identifier string, opts ...option.RequestOption) (res *LoadBalancerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/load_balancers/%s", identifier1, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type LoadBalancerNewResponse struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerNewResponseAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerNewResponseLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                               `json:"modified_on" format:"date-time"`
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
	RandomSteering LoadBalancerNewResponseRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancerNewResponseRule `json:"rules"`
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
	SessionAffinity LoadBalancerNewResponseSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerNewResponseSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerNewResponseSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                     `json:"ttl"`
	JSON loadBalancerNewResponseJSON `json:"-"`
}

// loadBalancerNewResponseJSON contains the JSON metadata for the struct
// [LoadBalancerNewResponse]
type loadBalancerNewResponseJSON struct {
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

func (r *LoadBalancerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerNewResponseAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                       `json:"failover_across_pools"`
	JSON                loadBalancerNewResponseAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerNewResponseAdaptiveRoutingJSON contains the JSON metadata for the
// struct [LoadBalancerNewResponseAdaptiveRouting]
type loadBalancerNewResponseAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerNewResponseAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerNewResponseLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerNewResponseLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerNewResponseLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerNewResponseLocationStrategyJSON      `json:"-"`
}

// loadBalancerNewResponseLocationStrategyJSON contains the JSON metadata for the
// struct [LoadBalancerNewResponseLocationStrategy]
type loadBalancerNewResponseLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerNewResponseLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerNewResponseLocationStrategyMode string

const (
	LoadBalancerNewResponseLocationStrategyModePop        LoadBalancerNewResponseLocationStrategyMode = "pop"
	LoadBalancerNewResponseLocationStrategyModeResolverIP LoadBalancerNewResponseLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerNewResponseLocationStrategyPreferEcs string

const (
	LoadBalancerNewResponseLocationStrategyPreferEcsAlways    LoadBalancerNewResponseLocationStrategyPreferEcs = "always"
	LoadBalancerNewResponseLocationStrategyPreferEcsNever     LoadBalancerNewResponseLocationStrategyPreferEcs = "never"
	LoadBalancerNewResponseLocationStrategyPreferEcsProximity LoadBalancerNewResponseLocationStrategyPreferEcs = "proximity"
	LoadBalancerNewResponseLocationStrategyPreferEcsGeo       LoadBalancerNewResponseLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerNewResponseRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                               `json:"pool_weights"`
	JSON        loadBalancerNewResponseRandomSteeringJSON `json:"-"`
}

// loadBalancerNewResponseRandomSteeringJSON contains the JSON metadata for the
// struct [LoadBalancerNewResponseRandomSteering]
type loadBalancerNewResponseRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerNewResponseRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerNewResponseRule struct {
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
	FixedResponse LoadBalancerNewResponseRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancerNewResponseRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                            `json:"terminates"`
	JSON       loadBalancerNewResponseRuleJSON `json:"-"`
}

// loadBalancerNewResponseRuleJSON contains the JSON metadata for the struct
// [LoadBalancerNewResponseRule]
type loadBalancerNewResponseRuleJSON struct {
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

func (r *LoadBalancerNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerNewResponseRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                         `json:"status_code"`
	JSON       loadBalancerNewResponseRulesFixedResponseJSON `json:"-"`
}

// loadBalancerNewResponseRulesFixedResponseJSON contains the JSON metadata for the
// struct [LoadBalancerNewResponseRulesFixedResponse]
type loadBalancerNewResponseRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerNewResponseRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerNewResponseRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerNewResponseRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerNewResponseRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering LoadBalancerNewResponseRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity LoadBalancerNewResponseRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerNewResponseRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerNewResponseRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                   `json:"ttl"`
	JSON loadBalancerNewResponseRulesOverridesJSON `json:"-"`
}

// loadBalancerNewResponseRulesOverridesJSON contains the JSON metadata for the
// struct [LoadBalancerNewResponseRulesOverrides]
type loadBalancerNewResponseRulesOverridesJSON struct {
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

func (r *LoadBalancerNewResponseRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerNewResponseRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                     `json:"failover_across_pools"`
	JSON                loadBalancerNewResponseRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerNewResponseRulesOverridesAdaptiveRoutingJSON contains the JSON
// metadata for the struct [LoadBalancerNewResponseRulesOverridesAdaptiveRouting]
type loadBalancerNewResponseRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerNewResponseRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerNewResponseRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerNewResponseRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerNewResponseRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerNewResponseRulesOverridesLocationStrategyJSON      `json:"-"`
}

// loadBalancerNewResponseRulesOverridesLocationStrategyJSON contains the JSON
// metadata for the struct [LoadBalancerNewResponseRulesOverridesLocationStrategy]
type loadBalancerNewResponseRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerNewResponseRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerNewResponseRulesOverridesLocationStrategyMode string

const (
	LoadBalancerNewResponseRulesOverridesLocationStrategyModePop        LoadBalancerNewResponseRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerNewResponseRulesOverridesLocationStrategyModeResolverIP LoadBalancerNewResponseRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerNewResponseRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerNewResponseRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerNewResponseRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerNewResponseRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerNewResponseRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerNewResponseRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerNewResponseRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerNewResponseRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerNewResponseRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerNewResponseRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                             `json:"pool_weights"`
	JSON        loadBalancerNewResponseRulesOverridesRandomSteeringJSON `json:"-"`
}

// loadBalancerNewResponseRulesOverridesRandomSteeringJSON contains the JSON
// metadata for the struct [LoadBalancerNewResponseRulesOverridesRandomSteering]
type loadBalancerNewResponseRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerNewResponseRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type LoadBalancerNewResponseRulesOverridesSessionAffinity string

const (
	LoadBalancerNewResponseRulesOverridesSessionAffinityNone     LoadBalancerNewResponseRulesOverridesSessionAffinity = "none"
	LoadBalancerNewResponseRulesOverridesSessionAffinityCookie   LoadBalancerNewResponseRulesOverridesSessionAffinity = "cookie"
	LoadBalancerNewResponseRulesOverridesSessionAffinityIPCookie LoadBalancerNewResponseRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerNewResponseRulesOverridesSessionAffinityHeader   LoadBalancerNewResponseRulesOverridesSessionAffinity = "header"
	LoadBalancerNewResponseRulesOverridesSessionAffinityEmpty    LoadBalancerNewResponseRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerNewResponseRulesOverridesSessionAffinityAttributes struct {
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
	Samesite LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerNewResponseRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerNewResponseRulesOverridesSessionAffinityAttributesJSON contains the
// JSON metadata for the struct
// [LoadBalancerNewResponseRulesOverridesSessionAffinityAttributes]
type loadBalancerNewResponseRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerNewResponseRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesSecure = "Never"
)

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
type LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerNewResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

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
type LoadBalancerNewResponseRulesOverridesSteeringPolicy string

const (
	LoadBalancerNewResponseRulesOverridesSteeringPolicyOff                      LoadBalancerNewResponseRulesOverridesSteeringPolicy = "off"
	LoadBalancerNewResponseRulesOverridesSteeringPolicyGeo                      LoadBalancerNewResponseRulesOverridesSteeringPolicy = "geo"
	LoadBalancerNewResponseRulesOverridesSteeringPolicyRandom                   LoadBalancerNewResponseRulesOverridesSteeringPolicy = "random"
	LoadBalancerNewResponseRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerNewResponseRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerNewResponseRulesOverridesSteeringPolicyProximity                LoadBalancerNewResponseRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerNewResponseRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerNewResponseRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerNewResponseRulesOverridesSteeringPolicyLeastConnections         LoadBalancerNewResponseRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerNewResponseRulesOverridesSteeringPolicyEmpty                    LoadBalancerNewResponseRulesOverridesSteeringPolicy = "\"\""
)

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
type LoadBalancerNewResponseSessionAffinity string

const (
	LoadBalancerNewResponseSessionAffinityNone     LoadBalancerNewResponseSessionAffinity = "none"
	LoadBalancerNewResponseSessionAffinityCookie   LoadBalancerNewResponseSessionAffinity = "cookie"
	LoadBalancerNewResponseSessionAffinityIPCookie LoadBalancerNewResponseSessionAffinity = "ip_cookie"
	LoadBalancerNewResponseSessionAffinityHeader   LoadBalancerNewResponseSessionAffinity = "header"
	LoadBalancerNewResponseSessionAffinityEmpty    LoadBalancerNewResponseSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerNewResponseSessionAffinityAttributes struct {
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
	Samesite LoadBalancerNewResponseSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerNewResponseSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerNewResponseSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerNewResponseSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerNewResponseSessionAffinityAttributesJSON contains the JSON metadata
// for the struct [LoadBalancerNewResponseSessionAffinityAttributes]
type loadBalancerNewResponseSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerNewResponseSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerNewResponseSessionAffinityAttributesSamesite string

const (
	LoadBalancerNewResponseSessionAffinityAttributesSamesiteAuto   LoadBalancerNewResponseSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerNewResponseSessionAffinityAttributesSamesiteLax    LoadBalancerNewResponseSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerNewResponseSessionAffinityAttributesSamesiteNone   LoadBalancerNewResponseSessionAffinityAttributesSamesite = "None"
	LoadBalancerNewResponseSessionAffinityAttributesSamesiteStrict LoadBalancerNewResponseSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerNewResponseSessionAffinityAttributesSecure string

const (
	LoadBalancerNewResponseSessionAffinityAttributesSecureAuto   LoadBalancerNewResponseSessionAffinityAttributesSecure = "Auto"
	LoadBalancerNewResponseSessionAffinityAttributesSecureAlways LoadBalancerNewResponseSessionAffinityAttributesSecure = "Always"
	LoadBalancerNewResponseSessionAffinityAttributesSecureNever  LoadBalancerNewResponseSessionAffinityAttributesSecure = "Never"
)

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
type LoadBalancerNewResponseSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerNewResponseSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerNewResponseSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerNewResponseSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerNewResponseSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerNewResponseSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerNewResponseSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

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
type LoadBalancerNewResponseSteeringPolicy string

const (
	LoadBalancerNewResponseSteeringPolicyOff                      LoadBalancerNewResponseSteeringPolicy = "off"
	LoadBalancerNewResponseSteeringPolicyGeo                      LoadBalancerNewResponseSteeringPolicy = "geo"
	LoadBalancerNewResponseSteeringPolicyRandom                   LoadBalancerNewResponseSteeringPolicy = "random"
	LoadBalancerNewResponseSteeringPolicyDynamicLatency           LoadBalancerNewResponseSteeringPolicy = "dynamic_latency"
	LoadBalancerNewResponseSteeringPolicyProximity                LoadBalancerNewResponseSteeringPolicy = "proximity"
	LoadBalancerNewResponseSteeringPolicyLeastOutstandingRequests LoadBalancerNewResponseSteeringPolicy = "least_outstanding_requests"
	LoadBalancerNewResponseSteeringPolicyLeastConnections         LoadBalancerNewResponseSteeringPolicy = "least_connections"
	LoadBalancerNewResponseSteeringPolicyEmpty                    LoadBalancerNewResponseSteeringPolicy = "\"\""
)

type LoadBalancerGetResponse struct {
	Errors   []LoadBalancerGetResponseError   `json:"errors"`
	Messages []LoadBalancerGetResponseMessage `json:"messages"`
	Result   LoadBalancerGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerGetResponseSuccess `json:"success"`
	JSON    loadBalancerGetResponseJSON    `json:"-"`
}

// loadBalancerGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancerGetResponse]
type loadBalancerGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerGetResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    loadBalancerGetResponseErrorJSON `json:"-"`
}

// loadBalancerGetResponseErrorJSON contains the JSON metadata for the struct
// [LoadBalancerGetResponseError]
type loadBalancerGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerGetResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    loadBalancerGetResponseMessageJSON `json:"-"`
}

// loadBalancerGetResponseMessageJSON contains the JSON metadata for the struct
// [LoadBalancerGetResponseMessage]
type loadBalancerGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerGetResponseResult struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerGetResponseResultAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerGetResponseResultLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                                     `json:"modified_on" format:"date-time"`
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
	RandomSteering LoadBalancerGetResponseResultRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancerGetResponseResultRule `json:"rules"`
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
	SessionAffinity LoadBalancerGetResponseResultSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerGetResponseResultSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerGetResponseResultSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                           `json:"ttl"`
	JSON loadBalancerGetResponseResultJSON `json:"-"`
}

// loadBalancerGetResponseResultJSON contains the JSON metadata for the struct
// [LoadBalancerGetResponseResult]
type loadBalancerGetResponseResultJSON struct {
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

func (r *LoadBalancerGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerGetResponseResultAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                             `json:"failover_across_pools"`
	JSON                loadBalancerGetResponseResultAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerGetResponseResultAdaptiveRoutingJSON contains the JSON metadata for
// the struct [LoadBalancerGetResponseResultAdaptiveRouting]
type loadBalancerGetResponseResultAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerGetResponseResultAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerGetResponseResultLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerGetResponseResultLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerGetResponseResultLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerGetResponseResultLocationStrategyJSON      `json:"-"`
}

// loadBalancerGetResponseResultLocationStrategyJSON contains the JSON metadata for
// the struct [LoadBalancerGetResponseResultLocationStrategy]
type loadBalancerGetResponseResultLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseResultLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerGetResponseResultLocationStrategyMode string

const (
	LoadBalancerGetResponseResultLocationStrategyModePop        LoadBalancerGetResponseResultLocationStrategyMode = "pop"
	LoadBalancerGetResponseResultLocationStrategyModeResolverIP LoadBalancerGetResponseResultLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerGetResponseResultLocationStrategyPreferEcs string

const (
	LoadBalancerGetResponseResultLocationStrategyPreferEcsAlways    LoadBalancerGetResponseResultLocationStrategyPreferEcs = "always"
	LoadBalancerGetResponseResultLocationStrategyPreferEcsNever     LoadBalancerGetResponseResultLocationStrategyPreferEcs = "never"
	LoadBalancerGetResponseResultLocationStrategyPreferEcsProximity LoadBalancerGetResponseResultLocationStrategyPreferEcs = "proximity"
	LoadBalancerGetResponseResultLocationStrategyPreferEcsGeo       LoadBalancerGetResponseResultLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerGetResponseResultRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                     `json:"pool_weights"`
	JSON        loadBalancerGetResponseResultRandomSteeringJSON `json:"-"`
}

// loadBalancerGetResponseResultRandomSteeringJSON contains the JSON metadata for
// the struct [LoadBalancerGetResponseResultRandomSteering]
type loadBalancerGetResponseResultRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerGetResponseResultRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerGetResponseResultRule struct {
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
	FixedResponse LoadBalancerGetResponseResultRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancerGetResponseResultRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                                  `json:"terminates"`
	JSON       loadBalancerGetResponseResultRuleJSON `json:"-"`
}

// loadBalancerGetResponseResultRuleJSON contains the JSON metadata for the struct
// [LoadBalancerGetResponseResultRule]
type loadBalancerGetResponseResultRuleJSON struct {
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

func (r *LoadBalancerGetResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerGetResponseResultRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                               `json:"status_code"`
	JSON       loadBalancerGetResponseResultRulesFixedResponseJSON `json:"-"`
}

// loadBalancerGetResponseResultRulesFixedResponseJSON contains the JSON metadata
// for the struct [LoadBalancerGetResponseResultRulesFixedResponse]
type loadBalancerGetResponseResultRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseResultRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerGetResponseResultRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerGetResponseResultRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerGetResponseResultRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering LoadBalancerGetResponseResultRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity LoadBalancerGetResponseResultRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerGetResponseResultRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                         `json:"ttl"`
	JSON loadBalancerGetResponseResultRulesOverridesJSON `json:"-"`
}

// loadBalancerGetResponseResultRulesOverridesJSON contains the JSON metadata for
// the struct [LoadBalancerGetResponseResultRulesOverrides]
type loadBalancerGetResponseResultRulesOverridesJSON struct {
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

func (r *LoadBalancerGetResponseResultRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerGetResponseResultRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                           `json:"failover_across_pools"`
	JSON                loadBalancerGetResponseResultRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerGetResponseResultRulesOverridesAdaptiveRoutingJSON contains the JSON
// metadata for the struct
// [LoadBalancerGetResponseResultRulesOverridesAdaptiveRouting]
type loadBalancerGetResponseResultRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerGetResponseResultRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerGetResponseResultRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerGetResponseResultRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerGetResponseResultRulesOverridesLocationStrategyJSON      `json:"-"`
}

// loadBalancerGetResponseResultRulesOverridesLocationStrategyJSON contains the
// JSON metadata for the struct
// [LoadBalancerGetResponseResultRulesOverridesLocationStrategy]
type loadBalancerGetResponseResultRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseResultRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerGetResponseResultRulesOverridesLocationStrategyMode string

const (
	LoadBalancerGetResponseResultRulesOverridesLocationStrategyModePop        LoadBalancerGetResponseResultRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerGetResponseResultRulesOverridesLocationStrategyModeResolverIP LoadBalancerGetResponseResultRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerGetResponseResultRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                                   `json:"pool_weights"`
	JSON        loadBalancerGetResponseResultRulesOverridesRandomSteeringJSON `json:"-"`
}

// loadBalancerGetResponseResultRulesOverridesRandomSteeringJSON contains the JSON
// metadata for the struct
// [LoadBalancerGetResponseResultRulesOverridesRandomSteering]
type loadBalancerGetResponseResultRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerGetResponseResultRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type LoadBalancerGetResponseResultRulesOverridesSessionAffinity string

const (
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityNone     LoadBalancerGetResponseResultRulesOverridesSessionAffinity = "none"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityCookie   LoadBalancerGetResponseResultRulesOverridesSessionAffinity = "cookie"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityIPCookie LoadBalancerGetResponseResultRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityHeader   LoadBalancerGetResponseResultRulesOverridesSessionAffinity = "header"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityEmpty    LoadBalancerGetResponseResultRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributes struct {
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
	Samesite LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesJSON
// contains the JSON metadata for the struct
// [LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributes]
type loadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecure = "Never"
)

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
type LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

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
type LoadBalancerGetResponseResultRulesOverridesSteeringPolicy string

const (
	LoadBalancerGetResponseResultRulesOverridesSteeringPolicyOff                      LoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "off"
	LoadBalancerGetResponseResultRulesOverridesSteeringPolicyGeo                      LoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "geo"
	LoadBalancerGetResponseResultRulesOverridesSteeringPolicyRandom                   LoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "random"
	LoadBalancerGetResponseResultRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerGetResponseResultRulesOverridesSteeringPolicyProximity                LoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerGetResponseResultRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerGetResponseResultRulesOverridesSteeringPolicyLeastConnections         LoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerGetResponseResultRulesOverridesSteeringPolicyEmpty                    LoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "\"\""
)

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
type LoadBalancerGetResponseResultSessionAffinity string

const (
	LoadBalancerGetResponseResultSessionAffinityNone     LoadBalancerGetResponseResultSessionAffinity = "none"
	LoadBalancerGetResponseResultSessionAffinityCookie   LoadBalancerGetResponseResultSessionAffinity = "cookie"
	LoadBalancerGetResponseResultSessionAffinityIPCookie LoadBalancerGetResponseResultSessionAffinity = "ip_cookie"
	LoadBalancerGetResponseResultSessionAffinityHeader   LoadBalancerGetResponseResultSessionAffinity = "header"
	LoadBalancerGetResponseResultSessionAffinityEmpty    LoadBalancerGetResponseResultSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerGetResponseResultSessionAffinityAttributes struct {
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
	Samesite LoadBalancerGetResponseResultSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerGetResponseResultSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerGetResponseResultSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerGetResponseResultSessionAffinityAttributesJSON contains the JSON
// metadata for the struct [LoadBalancerGetResponseResultSessionAffinityAttributes]
type loadBalancerGetResponseResultSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerGetResponseResultSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerGetResponseResultSessionAffinityAttributesSamesite string

const (
	LoadBalancerGetResponseResultSessionAffinityAttributesSamesiteAuto   LoadBalancerGetResponseResultSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerGetResponseResultSessionAffinityAttributesSamesiteLax    LoadBalancerGetResponseResultSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerGetResponseResultSessionAffinityAttributesSamesiteNone   LoadBalancerGetResponseResultSessionAffinityAttributesSamesite = "None"
	LoadBalancerGetResponseResultSessionAffinityAttributesSamesiteStrict LoadBalancerGetResponseResultSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerGetResponseResultSessionAffinityAttributesSecure string

const (
	LoadBalancerGetResponseResultSessionAffinityAttributesSecureAuto   LoadBalancerGetResponseResultSessionAffinityAttributesSecure = "Auto"
	LoadBalancerGetResponseResultSessionAffinityAttributesSecureAlways LoadBalancerGetResponseResultSessionAffinityAttributesSecure = "Always"
	LoadBalancerGetResponseResultSessionAffinityAttributesSecureNever  LoadBalancerGetResponseResultSessionAffinityAttributesSecure = "Never"
)

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
type LoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

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
type LoadBalancerGetResponseResultSteeringPolicy string

const (
	LoadBalancerGetResponseResultSteeringPolicyOff                      LoadBalancerGetResponseResultSteeringPolicy = "off"
	LoadBalancerGetResponseResultSteeringPolicyGeo                      LoadBalancerGetResponseResultSteeringPolicy = "geo"
	LoadBalancerGetResponseResultSteeringPolicyRandom                   LoadBalancerGetResponseResultSteeringPolicy = "random"
	LoadBalancerGetResponseResultSteeringPolicyDynamicLatency           LoadBalancerGetResponseResultSteeringPolicy = "dynamic_latency"
	LoadBalancerGetResponseResultSteeringPolicyProximity                LoadBalancerGetResponseResultSteeringPolicy = "proximity"
	LoadBalancerGetResponseResultSteeringPolicyLeastOutstandingRequests LoadBalancerGetResponseResultSteeringPolicy = "least_outstanding_requests"
	LoadBalancerGetResponseResultSteeringPolicyLeastConnections         LoadBalancerGetResponseResultSteeringPolicy = "least_connections"
	LoadBalancerGetResponseResultSteeringPolicyEmpty                    LoadBalancerGetResponseResultSteeringPolicy = "\"\""
)

// Whether the API call was successful
type LoadBalancerGetResponseSuccess bool

const (
	LoadBalancerGetResponseSuccessTrue LoadBalancerGetResponseSuccess = true
)

type LoadBalancerUpdateResponse struct {
	Errors   []LoadBalancerUpdateResponseError   `json:"errors"`
	Messages []LoadBalancerUpdateResponseMessage `json:"messages"`
	Result   LoadBalancerUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerUpdateResponseSuccess `json:"success"`
	JSON    loadBalancerUpdateResponseJSON    `json:"-"`
}

// loadBalancerUpdateResponseJSON contains the JSON metadata for the struct
// [LoadBalancerUpdateResponse]
type loadBalancerUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerUpdateResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    loadBalancerUpdateResponseErrorJSON `json:"-"`
}

// loadBalancerUpdateResponseErrorJSON contains the JSON metadata for the struct
// [LoadBalancerUpdateResponseError]
type loadBalancerUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerUpdateResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    loadBalancerUpdateResponseMessageJSON `json:"-"`
}

// loadBalancerUpdateResponseMessageJSON contains the JSON metadata for the struct
// [LoadBalancerUpdateResponseMessage]
type loadBalancerUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerUpdateResponseResult struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerUpdateResponseResultAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerUpdateResponseResultLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                                        `json:"modified_on" format:"date-time"`
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
	RandomSteering LoadBalancerUpdateResponseResultRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancerUpdateResponseResultRule `json:"rules"`
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
	SessionAffinity LoadBalancerUpdateResponseResultSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerUpdateResponseResultSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerUpdateResponseResultSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                              `json:"ttl"`
	JSON loadBalancerUpdateResponseResultJSON `json:"-"`
}

// loadBalancerUpdateResponseResultJSON contains the JSON metadata for the struct
// [LoadBalancerUpdateResponseResult]
type loadBalancerUpdateResponseResultJSON struct {
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

func (r *LoadBalancerUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerUpdateResponseResultAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                `json:"failover_across_pools"`
	JSON                loadBalancerUpdateResponseResultAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerUpdateResponseResultAdaptiveRoutingJSON contains the JSON metadata
// for the struct [LoadBalancerUpdateResponseResultAdaptiveRouting]
type loadBalancerUpdateResponseResultAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseResultAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerUpdateResponseResultLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerUpdateResponseResultLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerUpdateResponseResultLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerUpdateResponseResultLocationStrategyJSON      `json:"-"`
}

// loadBalancerUpdateResponseResultLocationStrategyJSON contains the JSON metadata
// for the struct [LoadBalancerUpdateResponseResultLocationStrategy]
type loadBalancerUpdateResponseResultLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseResultLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerUpdateResponseResultLocationStrategyMode string

const (
	LoadBalancerUpdateResponseResultLocationStrategyModePop        LoadBalancerUpdateResponseResultLocationStrategyMode = "pop"
	LoadBalancerUpdateResponseResultLocationStrategyModeResolverIP LoadBalancerUpdateResponseResultLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerUpdateResponseResultLocationStrategyPreferEcs string

const (
	LoadBalancerUpdateResponseResultLocationStrategyPreferEcsAlways    LoadBalancerUpdateResponseResultLocationStrategyPreferEcs = "always"
	LoadBalancerUpdateResponseResultLocationStrategyPreferEcsNever     LoadBalancerUpdateResponseResultLocationStrategyPreferEcs = "never"
	LoadBalancerUpdateResponseResultLocationStrategyPreferEcsProximity LoadBalancerUpdateResponseResultLocationStrategyPreferEcs = "proximity"
	LoadBalancerUpdateResponseResultLocationStrategyPreferEcsGeo       LoadBalancerUpdateResponseResultLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerUpdateResponseResultRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                        `json:"pool_weights"`
	JSON        loadBalancerUpdateResponseResultRandomSteeringJSON `json:"-"`
}

// loadBalancerUpdateResponseResultRandomSteeringJSON contains the JSON metadata
// for the struct [LoadBalancerUpdateResponseResultRandomSteering]
type loadBalancerUpdateResponseResultRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseResultRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerUpdateResponseResultRule struct {
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
	FixedResponse LoadBalancerUpdateResponseResultRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancerUpdateResponseResultRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                                     `json:"terminates"`
	JSON       loadBalancerUpdateResponseResultRuleJSON `json:"-"`
}

// loadBalancerUpdateResponseResultRuleJSON contains the JSON metadata for the
// struct [LoadBalancerUpdateResponseResultRule]
type loadBalancerUpdateResponseResultRuleJSON struct {
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

func (r *LoadBalancerUpdateResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerUpdateResponseResultRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                                  `json:"status_code"`
	JSON       loadBalancerUpdateResponseResultRulesFixedResponseJSON `json:"-"`
}

// loadBalancerUpdateResponseResultRulesFixedResponseJSON contains the JSON
// metadata for the struct [LoadBalancerUpdateResponseResultRulesFixedResponse]
type loadBalancerUpdateResponseResultRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseResultRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerUpdateResponseResultRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerUpdateResponseResultRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerUpdateResponseResultRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering LoadBalancerUpdateResponseResultRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity LoadBalancerUpdateResponseResultRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                            `json:"ttl"`
	JSON loadBalancerUpdateResponseResultRulesOverridesJSON `json:"-"`
}

// loadBalancerUpdateResponseResultRulesOverridesJSON contains the JSON metadata
// for the struct [LoadBalancerUpdateResponseResultRulesOverrides]
type loadBalancerUpdateResponseResultRulesOverridesJSON struct {
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

func (r *LoadBalancerUpdateResponseResultRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerUpdateResponseResultRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                              `json:"failover_across_pools"`
	JSON                loadBalancerUpdateResponseResultRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerUpdateResponseResultRulesOverridesAdaptiveRoutingJSON contains the
// JSON metadata for the struct
// [LoadBalancerUpdateResponseResultRulesOverridesAdaptiveRouting]
type loadBalancerUpdateResponseResultRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseResultRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerUpdateResponseResultRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerUpdateResponseResultRulesOverridesLocationStrategyJSON      `json:"-"`
}

// loadBalancerUpdateResponseResultRulesOverridesLocationStrategyJSON contains the
// JSON metadata for the struct
// [LoadBalancerUpdateResponseResultRulesOverridesLocationStrategy]
type loadBalancerUpdateResponseResultRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseResultRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyMode string

const (
	LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyModePop        LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyModeResolverIP LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerUpdateResponseResultRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                                      `json:"pool_weights"`
	JSON        loadBalancerUpdateResponseResultRulesOverridesRandomSteeringJSON `json:"-"`
}

// loadBalancerUpdateResponseResultRulesOverridesRandomSteeringJSON contains the
// JSON metadata for the struct
// [LoadBalancerUpdateResponseResultRulesOverridesRandomSteering]
type loadBalancerUpdateResponseResultRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseResultRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type LoadBalancerUpdateResponseResultRulesOverridesSessionAffinity string

const (
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityNone     LoadBalancerUpdateResponseResultRulesOverridesSessionAffinity = "none"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityCookie   LoadBalancerUpdateResponseResultRulesOverridesSessionAffinity = "cookie"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityIPCookie LoadBalancerUpdateResponseResultRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityHeader   LoadBalancerUpdateResponseResultRulesOverridesSessionAffinity = "header"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityEmpty    LoadBalancerUpdateResponseResultRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributes struct {
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
	Samesite LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesJSON
// contains the JSON metadata for the struct
// [LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributes]
type loadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecure = "Never"
)

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
type LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

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
type LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy string

const (
	LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyOff                      LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "off"
	LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyGeo                      LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "geo"
	LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyRandom                   LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "random"
	LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyProximity                LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyLeastConnections         LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyEmpty                    LoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "\"\""
)

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
type LoadBalancerUpdateResponseResultSessionAffinity string

const (
	LoadBalancerUpdateResponseResultSessionAffinityNone     LoadBalancerUpdateResponseResultSessionAffinity = "none"
	LoadBalancerUpdateResponseResultSessionAffinityCookie   LoadBalancerUpdateResponseResultSessionAffinity = "cookie"
	LoadBalancerUpdateResponseResultSessionAffinityIPCookie LoadBalancerUpdateResponseResultSessionAffinity = "ip_cookie"
	LoadBalancerUpdateResponseResultSessionAffinityHeader   LoadBalancerUpdateResponseResultSessionAffinity = "header"
	LoadBalancerUpdateResponseResultSessionAffinityEmpty    LoadBalancerUpdateResponseResultSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerUpdateResponseResultSessionAffinityAttributes struct {
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
	Samesite LoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerUpdateResponseResultSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerUpdateResponseResultSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerUpdateResponseResultSessionAffinityAttributesJSON contains the JSON
// metadata for the struct
// [LoadBalancerUpdateResponseResultSessionAffinityAttributes]
type loadBalancerUpdateResponseResultSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseResultSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite string

const (
	LoadBalancerUpdateResponseResultSessionAffinityAttributesSamesiteAuto   LoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerUpdateResponseResultSessionAffinityAttributesSamesiteLax    LoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerUpdateResponseResultSessionAffinityAttributesSamesiteNone   LoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite = "None"
	LoadBalancerUpdateResponseResultSessionAffinityAttributesSamesiteStrict LoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerUpdateResponseResultSessionAffinityAttributesSecure string

const (
	LoadBalancerUpdateResponseResultSessionAffinityAttributesSecureAuto   LoadBalancerUpdateResponseResultSessionAffinityAttributesSecure = "Auto"
	LoadBalancerUpdateResponseResultSessionAffinityAttributesSecureAlways LoadBalancerUpdateResponseResultSessionAffinityAttributesSecure = "Always"
	LoadBalancerUpdateResponseResultSessionAffinityAttributesSecureNever  LoadBalancerUpdateResponseResultSessionAffinityAttributesSecure = "Never"
)

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
type LoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

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
type LoadBalancerUpdateResponseResultSteeringPolicy string

const (
	LoadBalancerUpdateResponseResultSteeringPolicyOff                      LoadBalancerUpdateResponseResultSteeringPolicy = "off"
	LoadBalancerUpdateResponseResultSteeringPolicyGeo                      LoadBalancerUpdateResponseResultSteeringPolicy = "geo"
	LoadBalancerUpdateResponseResultSteeringPolicyRandom                   LoadBalancerUpdateResponseResultSteeringPolicy = "random"
	LoadBalancerUpdateResponseResultSteeringPolicyDynamicLatency           LoadBalancerUpdateResponseResultSteeringPolicy = "dynamic_latency"
	LoadBalancerUpdateResponseResultSteeringPolicyProximity                LoadBalancerUpdateResponseResultSteeringPolicy = "proximity"
	LoadBalancerUpdateResponseResultSteeringPolicyLeastOutstandingRequests LoadBalancerUpdateResponseResultSteeringPolicy = "least_outstanding_requests"
	LoadBalancerUpdateResponseResultSteeringPolicyLeastConnections         LoadBalancerUpdateResponseResultSteeringPolicy = "least_connections"
	LoadBalancerUpdateResponseResultSteeringPolicyEmpty                    LoadBalancerUpdateResponseResultSteeringPolicy = "\"\""
)

// Whether the API call was successful
type LoadBalancerUpdateResponseSuccess bool

const (
	LoadBalancerUpdateResponseSuccessTrue LoadBalancerUpdateResponseSuccess = true
)

type LoadBalancerListResponse struct {
	Errors     []LoadBalancerListResponseError    `json:"errors"`
	Messages   []LoadBalancerListResponseMessage  `json:"messages"`
	Result     []LoadBalancerListResponseResult   `json:"result"`
	ResultInfo LoadBalancerListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerListResponseSuccess `json:"success"`
	JSON    loadBalancerListResponseJSON    `json:"-"`
}

// loadBalancerListResponseJSON contains the JSON metadata for the struct
// [LoadBalancerListResponse]
type loadBalancerListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    loadBalancerListResponseErrorJSON `json:"-"`
}

// loadBalancerListResponseErrorJSON contains the JSON metadata for the struct
// [LoadBalancerListResponseError]
type loadBalancerListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    loadBalancerListResponseMessageJSON `json:"-"`
}

// loadBalancerListResponseMessageJSON contains the JSON metadata for the struct
// [LoadBalancerListResponseMessage]
type loadBalancerListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListResponseResult struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerListResponseResultAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerListResponseResultLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                                      `json:"modified_on" format:"date-time"`
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
	RandomSteering LoadBalancerListResponseResultRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancerListResponseResultRule `json:"rules"`
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
	SessionAffinity LoadBalancerListResponseResultSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerListResponseResultSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerListResponseResultSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                            `json:"ttl"`
	JSON loadBalancerListResponseResultJSON `json:"-"`
}

// loadBalancerListResponseResultJSON contains the JSON metadata for the struct
// [LoadBalancerListResponseResult]
type loadBalancerListResponseResultJSON struct {
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

func (r *LoadBalancerListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerListResponseResultAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                              `json:"failover_across_pools"`
	JSON                loadBalancerListResponseResultAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerListResponseResultAdaptiveRoutingJSON contains the JSON metadata for
// the struct [LoadBalancerListResponseResultAdaptiveRouting]
type loadBalancerListResponseResultAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerListResponseResultAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerListResponseResultLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerListResponseResultLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerListResponseResultLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerListResponseResultLocationStrategyJSON      `json:"-"`
}

// loadBalancerListResponseResultLocationStrategyJSON contains the JSON metadata
// for the struct [LoadBalancerListResponseResultLocationStrategy]
type loadBalancerListResponseResultLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseResultLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerListResponseResultLocationStrategyMode string

const (
	LoadBalancerListResponseResultLocationStrategyModePop        LoadBalancerListResponseResultLocationStrategyMode = "pop"
	LoadBalancerListResponseResultLocationStrategyModeResolverIP LoadBalancerListResponseResultLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerListResponseResultLocationStrategyPreferEcs string

const (
	LoadBalancerListResponseResultLocationStrategyPreferEcsAlways    LoadBalancerListResponseResultLocationStrategyPreferEcs = "always"
	LoadBalancerListResponseResultLocationStrategyPreferEcsNever     LoadBalancerListResponseResultLocationStrategyPreferEcs = "never"
	LoadBalancerListResponseResultLocationStrategyPreferEcsProximity LoadBalancerListResponseResultLocationStrategyPreferEcs = "proximity"
	LoadBalancerListResponseResultLocationStrategyPreferEcsGeo       LoadBalancerListResponseResultLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerListResponseResultRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                      `json:"pool_weights"`
	JSON        loadBalancerListResponseResultRandomSteeringJSON `json:"-"`
}

// loadBalancerListResponseResultRandomSteeringJSON contains the JSON metadata for
// the struct [LoadBalancerListResponseResultRandomSteering]
type loadBalancerListResponseResultRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerListResponseResultRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerListResponseResultRule struct {
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
	FixedResponse LoadBalancerListResponseResultRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancerListResponseResultRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                                   `json:"terminates"`
	JSON       loadBalancerListResponseResultRuleJSON `json:"-"`
}

// loadBalancerListResponseResultRuleJSON contains the JSON metadata for the struct
// [LoadBalancerListResponseResultRule]
type loadBalancerListResponseResultRuleJSON struct {
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

func (r *LoadBalancerListResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerListResponseResultRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                                `json:"status_code"`
	JSON       loadBalancerListResponseResultRulesFixedResponseJSON `json:"-"`
}

// loadBalancerListResponseResultRulesFixedResponseJSON contains the JSON metadata
// for the struct [LoadBalancerListResponseResultRulesFixedResponse]
type loadBalancerListResponseResultRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseResultRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerListResponseResultRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerListResponseResultRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerListResponseResultRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering LoadBalancerListResponseResultRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity LoadBalancerListResponseResultRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerListResponseResultRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                          `json:"ttl"`
	JSON loadBalancerListResponseResultRulesOverridesJSON `json:"-"`
}

// loadBalancerListResponseResultRulesOverridesJSON contains the JSON metadata for
// the struct [LoadBalancerListResponseResultRulesOverrides]
type loadBalancerListResponseResultRulesOverridesJSON struct {
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

func (r *LoadBalancerListResponseResultRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerListResponseResultRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                            `json:"failover_across_pools"`
	JSON                loadBalancerListResponseResultRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerListResponseResultRulesOverridesAdaptiveRoutingJSON contains the
// JSON metadata for the struct
// [LoadBalancerListResponseResultRulesOverridesAdaptiveRouting]
type loadBalancerListResponseResultRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerListResponseResultRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerListResponseResultRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerListResponseResultRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerListResponseResultRulesOverridesLocationStrategyJSON      `json:"-"`
}

// loadBalancerListResponseResultRulesOverridesLocationStrategyJSON contains the
// JSON metadata for the struct
// [LoadBalancerListResponseResultRulesOverridesLocationStrategy]
type loadBalancerListResponseResultRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseResultRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerListResponseResultRulesOverridesLocationStrategyMode string

const (
	LoadBalancerListResponseResultRulesOverridesLocationStrategyModePop        LoadBalancerListResponseResultRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerListResponseResultRulesOverridesLocationStrategyModeResolverIP LoadBalancerListResponseResultRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerListResponseResultRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                                    `json:"pool_weights"`
	JSON        loadBalancerListResponseResultRulesOverridesRandomSteeringJSON `json:"-"`
}

// loadBalancerListResponseResultRulesOverridesRandomSteeringJSON contains the JSON
// metadata for the struct
// [LoadBalancerListResponseResultRulesOverridesRandomSteering]
type loadBalancerListResponseResultRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerListResponseResultRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type LoadBalancerListResponseResultRulesOverridesSessionAffinity string

const (
	LoadBalancerListResponseResultRulesOverridesSessionAffinityNone     LoadBalancerListResponseResultRulesOverridesSessionAffinity = "none"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityCookie   LoadBalancerListResponseResultRulesOverridesSessionAffinity = "cookie"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityIPCookie LoadBalancerListResponseResultRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityHeader   LoadBalancerListResponseResultRulesOverridesSessionAffinity = "header"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityEmpty    LoadBalancerListResponseResultRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributes struct {
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
	Samesite LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerListResponseResultRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerListResponseResultRulesOverridesSessionAffinityAttributesJSON
// contains the JSON metadata for the struct
// [LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributes]
type loadBalancerListResponseResultRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecure = "Never"
)

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
type LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

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
type LoadBalancerListResponseResultRulesOverridesSteeringPolicy string

const (
	LoadBalancerListResponseResultRulesOverridesSteeringPolicyOff                      LoadBalancerListResponseResultRulesOverridesSteeringPolicy = "off"
	LoadBalancerListResponseResultRulesOverridesSteeringPolicyGeo                      LoadBalancerListResponseResultRulesOverridesSteeringPolicy = "geo"
	LoadBalancerListResponseResultRulesOverridesSteeringPolicyRandom                   LoadBalancerListResponseResultRulesOverridesSteeringPolicy = "random"
	LoadBalancerListResponseResultRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerListResponseResultRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerListResponseResultRulesOverridesSteeringPolicyProximity                LoadBalancerListResponseResultRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerListResponseResultRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerListResponseResultRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerListResponseResultRulesOverridesSteeringPolicyLeastConnections         LoadBalancerListResponseResultRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerListResponseResultRulesOverridesSteeringPolicyEmpty                    LoadBalancerListResponseResultRulesOverridesSteeringPolicy = "\"\""
)

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
type LoadBalancerListResponseResultSessionAffinity string

const (
	LoadBalancerListResponseResultSessionAffinityNone     LoadBalancerListResponseResultSessionAffinity = "none"
	LoadBalancerListResponseResultSessionAffinityCookie   LoadBalancerListResponseResultSessionAffinity = "cookie"
	LoadBalancerListResponseResultSessionAffinityIPCookie LoadBalancerListResponseResultSessionAffinity = "ip_cookie"
	LoadBalancerListResponseResultSessionAffinityHeader   LoadBalancerListResponseResultSessionAffinity = "header"
	LoadBalancerListResponseResultSessionAffinityEmpty    LoadBalancerListResponseResultSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerListResponseResultSessionAffinityAttributes struct {
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
	Samesite LoadBalancerListResponseResultSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerListResponseResultSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerListResponseResultSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerListResponseResultSessionAffinityAttributesJSON contains the JSON
// metadata for the struct
// [LoadBalancerListResponseResultSessionAffinityAttributes]
type loadBalancerListResponseResultSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerListResponseResultSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerListResponseResultSessionAffinityAttributesSamesite string

const (
	LoadBalancerListResponseResultSessionAffinityAttributesSamesiteAuto   LoadBalancerListResponseResultSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerListResponseResultSessionAffinityAttributesSamesiteLax    LoadBalancerListResponseResultSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerListResponseResultSessionAffinityAttributesSamesiteNone   LoadBalancerListResponseResultSessionAffinityAttributesSamesite = "None"
	LoadBalancerListResponseResultSessionAffinityAttributesSamesiteStrict LoadBalancerListResponseResultSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerListResponseResultSessionAffinityAttributesSecure string

const (
	LoadBalancerListResponseResultSessionAffinityAttributesSecureAuto   LoadBalancerListResponseResultSessionAffinityAttributesSecure = "Auto"
	LoadBalancerListResponseResultSessionAffinityAttributesSecureAlways LoadBalancerListResponseResultSessionAffinityAttributesSecure = "Always"
	LoadBalancerListResponseResultSessionAffinityAttributesSecureNever  LoadBalancerListResponseResultSessionAffinityAttributesSecure = "Never"
)

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
type LoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

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
type LoadBalancerListResponseResultSteeringPolicy string

const (
	LoadBalancerListResponseResultSteeringPolicyOff                      LoadBalancerListResponseResultSteeringPolicy = "off"
	LoadBalancerListResponseResultSteeringPolicyGeo                      LoadBalancerListResponseResultSteeringPolicy = "geo"
	LoadBalancerListResponseResultSteeringPolicyRandom                   LoadBalancerListResponseResultSteeringPolicy = "random"
	LoadBalancerListResponseResultSteeringPolicyDynamicLatency           LoadBalancerListResponseResultSteeringPolicy = "dynamic_latency"
	LoadBalancerListResponseResultSteeringPolicyProximity                LoadBalancerListResponseResultSteeringPolicy = "proximity"
	LoadBalancerListResponseResultSteeringPolicyLeastOutstandingRequests LoadBalancerListResponseResultSteeringPolicy = "least_outstanding_requests"
	LoadBalancerListResponseResultSteeringPolicyLeastConnections         LoadBalancerListResponseResultSteeringPolicy = "least_connections"
	LoadBalancerListResponseResultSteeringPolicyEmpty                    LoadBalancerListResponseResultSteeringPolicy = "\"\""
)

type LoadBalancerListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       loadBalancerListResponseResultInfoJSON `json:"-"`
}

// loadBalancerListResponseResultInfoJSON contains the JSON metadata for the struct
// [LoadBalancerListResponseResultInfo]
type loadBalancerListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerListResponseSuccess bool

const (
	LoadBalancerListResponseSuccessTrue LoadBalancerListResponseSuccess = true
)

type LoadBalancerDeleteResponse struct {
	Errors   []LoadBalancerDeleteResponseError   `json:"errors"`
	Messages []LoadBalancerDeleteResponseMessage `json:"messages"`
	Result   LoadBalancerDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerDeleteResponseSuccess `json:"success"`
	JSON    loadBalancerDeleteResponseJSON    `json:"-"`
}

// loadBalancerDeleteResponseJSON contains the JSON metadata for the struct
// [LoadBalancerDeleteResponse]
type loadBalancerDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerDeleteResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    loadBalancerDeleteResponseErrorJSON `json:"-"`
}

// loadBalancerDeleteResponseErrorJSON contains the JSON metadata for the struct
// [LoadBalancerDeleteResponseError]
type loadBalancerDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerDeleteResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    loadBalancerDeleteResponseMessageJSON `json:"-"`
}

// loadBalancerDeleteResponseMessageJSON contains the JSON metadata for the struct
// [LoadBalancerDeleteResponseMessage]
type loadBalancerDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerDeleteResponseResult struct {
	ID   string                               `json:"id"`
	JSON loadBalancerDeleteResponseResultJSON `json:"-"`
}

// loadBalancerDeleteResponseResultJSON contains the JSON metadata for the struct
// [LoadBalancerDeleteResponseResult]
type loadBalancerDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerDeleteResponseSuccess bool

const (
	LoadBalancerDeleteResponseSuccessTrue LoadBalancerDeleteResponseSuccess = true
)

type LoadBalancerNewParams struct {
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

type LoadBalancerNewResponseEnvelope struct {
	Errors   []LoadBalancerNewResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerNewResponseEnvelopeMessages `json:"messages"`
	Result   LoadBalancerNewResponse                   `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerNewResponseEnvelopeSuccess `json:"success"`
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

type LoadBalancerNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    loadBalancerNewResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LoadBalancerNewResponseEnvelopeErrors]
type loadBalancerNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    loadBalancerNewResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LoadBalancerNewResponseEnvelopeMessages]
type loadBalancerNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerNewResponseEnvelopeSuccess bool

const (
	LoadBalancerNewResponseEnvelopeSuccessTrue LoadBalancerNewResponseEnvelopeSuccess = true
)

type LoadBalancerUpdateParams struct {
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
