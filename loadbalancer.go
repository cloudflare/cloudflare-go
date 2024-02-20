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
func (r *LoadBalancerService) New(ctx context.Context, zoneID string, body LoadBalancerNewParams, opts ...option.RequestOption) (res *LoadBalancerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured load balancers.
func (r *LoadBalancerService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]LoadBalancerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerListResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured load balancer.
func (r *LoadBalancerService) Delete(ctx context.Context, zoneID string, loadBalancerID string, opts ...option.RequestOption) (res *LoadBalancerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers/%s", zoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing load balancer, overwriting the supplied properties.
func (r *LoadBalancerService) Edit(ctx context.Context, zoneID string, loadBalancerID string, body LoadBalancerEditParams, opts ...option.RequestOption) (res *LoadBalancerEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers/%s", zoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured load balancer.
func (r *LoadBalancerService) Get(ctx context.Context, zoneID string, loadBalancerID string, opts ...option.RequestOption) (res *LoadBalancerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/load_balancers/%s", zoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

type LoadBalancerListResponse struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerListResponseAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerListResponseLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                                `json:"modified_on" format:"date-time"`
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
	RandomSteering LoadBalancerListResponseRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancerListResponseRule `json:"rules"`
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
	SessionAffinity LoadBalancerListResponseSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerListResponseSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerListResponseSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                      `json:"ttl"`
	JSON loadBalancerListResponseJSON `json:"-"`
}

// loadBalancerListResponseJSON contains the JSON metadata for the struct
// [LoadBalancerListResponse]
type loadBalancerListResponseJSON struct {
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

func (r *LoadBalancerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerListResponseAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                        `json:"failover_across_pools"`
	JSON                loadBalancerListResponseAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerListResponseAdaptiveRoutingJSON contains the JSON metadata for the
// struct [LoadBalancerListResponseAdaptiveRouting]
type loadBalancerListResponseAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerListResponseAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerListResponseLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerListResponseLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerListResponseLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerListResponseLocationStrategyJSON      `json:"-"`
}

// loadBalancerListResponseLocationStrategyJSON contains the JSON metadata for the
// struct [LoadBalancerListResponseLocationStrategy]
type loadBalancerListResponseLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerListResponseLocationStrategyMode string

const (
	LoadBalancerListResponseLocationStrategyModePop        LoadBalancerListResponseLocationStrategyMode = "pop"
	LoadBalancerListResponseLocationStrategyModeResolverIP LoadBalancerListResponseLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerListResponseLocationStrategyPreferEcs string

const (
	LoadBalancerListResponseLocationStrategyPreferEcsAlways    LoadBalancerListResponseLocationStrategyPreferEcs = "always"
	LoadBalancerListResponseLocationStrategyPreferEcsNever     LoadBalancerListResponseLocationStrategyPreferEcs = "never"
	LoadBalancerListResponseLocationStrategyPreferEcsProximity LoadBalancerListResponseLocationStrategyPreferEcs = "proximity"
	LoadBalancerListResponseLocationStrategyPreferEcsGeo       LoadBalancerListResponseLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerListResponseRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                `json:"pool_weights"`
	JSON        loadBalancerListResponseRandomSteeringJSON `json:"-"`
}

// loadBalancerListResponseRandomSteeringJSON contains the JSON metadata for the
// struct [LoadBalancerListResponseRandomSteering]
type loadBalancerListResponseRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerListResponseRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerListResponseRule struct {
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
	FixedResponse LoadBalancerListResponseRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancerListResponseRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                             `json:"terminates"`
	JSON       loadBalancerListResponseRuleJSON `json:"-"`
}

// loadBalancerListResponseRuleJSON contains the JSON metadata for the struct
// [LoadBalancerListResponseRule]
type loadBalancerListResponseRuleJSON struct {
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

func (r *LoadBalancerListResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerListResponseRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                          `json:"status_code"`
	JSON       loadBalancerListResponseRulesFixedResponseJSON `json:"-"`
}

// loadBalancerListResponseRulesFixedResponseJSON contains the JSON metadata for
// the struct [LoadBalancerListResponseRulesFixedResponse]
type loadBalancerListResponseRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerListResponseRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerListResponseRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerListResponseRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering LoadBalancerListResponseRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity LoadBalancerListResponseRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerListResponseRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerListResponseRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                    `json:"ttl"`
	JSON loadBalancerListResponseRulesOverridesJSON `json:"-"`
}

// loadBalancerListResponseRulesOverridesJSON contains the JSON metadata for the
// struct [LoadBalancerListResponseRulesOverrides]
type loadBalancerListResponseRulesOverridesJSON struct {
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

func (r *LoadBalancerListResponseRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerListResponseRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                      `json:"failover_across_pools"`
	JSON                loadBalancerListResponseRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerListResponseRulesOverridesAdaptiveRoutingJSON contains the JSON
// metadata for the struct [LoadBalancerListResponseRulesOverridesAdaptiveRouting]
type loadBalancerListResponseRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerListResponseRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerListResponseRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerListResponseRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerListResponseRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerListResponseRulesOverridesLocationStrategyJSON      `json:"-"`
}

// loadBalancerListResponseRulesOverridesLocationStrategyJSON contains the JSON
// metadata for the struct [LoadBalancerListResponseRulesOverridesLocationStrategy]
type loadBalancerListResponseRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerListResponseRulesOverridesLocationStrategyMode string

const (
	LoadBalancerListResponseRulesOverridesLocationStrategyModePop        LoadBalancerListResponseRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerListResponseRulesOverridesLocationStrategyModeResolverIP LoadBalancerListResponseRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerListResponseRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerListResponseRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerListResponseRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerListResponseRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerListResponseRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerListResponseRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerListResponseRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerListResponseRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerListResponseRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerListResponseRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                              `json:"pool_weights"`
	JSON        loadBalancerListResponseRulesOverridesRandomSteeringJSON `json:"-"`
}

// loadBalancerListResponseRulesOverridesRandomSteeringJSON contains the JSON
// metadata for the struct [LoadBalancerListResponseRulesOverridesRandomSteering]
type loadBalancerListResponseRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerListResponseRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
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
type LoadBalancerListResponseRulesOverridesSessionAffinity string

const (
	LoadBalancerListResponseRulesOverridesSessionAffinityNone     LoadBalancerListResponseRulesOverridesSessionAffinity = "none"
	LoadBalancerListResponseRulesOverridesSessionAffinityCookie   LoadBalancerListResponseRulesOverridesSessionAffinity = "cookie"
	LoadBalancerListResponseRulesOverridesSessionAffinityIPCookie LoadBalancerListResponseRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerListResponseRulesOverridesSessionAffinityHeader   LoadBalancerListResponseRulesOverridesSessionAffinity = "header"
	LoadBalancerListResponseRulesOverridesSessionAffinityEmpty    LoadBalancerListResponseRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerListResponseRulesOverridesSessionAffinityAttributes struct {
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
	Samesite LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerListResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerListResponseRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerListResponseRulesOverridesSessionAffinityAttributesJSON contains the
// JSON metadata for the struct
// [LoadBalancerListResponseRulesOverridesSessionAffinityAttributes]
type loadBalancerListResponseRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerListResponseRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerListResponseRulesOverridesSessionAffinityAttributesSecure = "Never"
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
type LoadBalancerListResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerListResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerListResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerListResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerListResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerListResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerListResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type LoadBalancerListResponseRulesOverridesSteeringPolicy string

const (
	LoadBalancerListResponseRulesOverridesSteeringPolicyOff                      LoadBalancerListResponseRulesOverridesSteeringPolicy = "off"
	LoadBalancerListResponseRulesOverridesSteeringPolicyGeo                      LoadBalancerListResponseRulesOverridesSteeringPolicy = "geo"
	LoadBalancerListResponseRulesOverridesSteeringPolicyRandom                   LoadBalancerListResponseRulesOverridesSteeringPolicy = "random"
	LoadBalancerListResponseRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerListResponseRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerListResponseRulesOverridesSteeringPolicyProximity                LoadBalancerListResponseRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerListResponseRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerListResponseRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerListResponseRulesOverridesSteeringPolicyLeastConnections         LoadBalancerListResponseRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerListResponseRulesOverridesSteeringPolicyEmpty                    LoadBalancerListResponseRulesOverridesSteeringPolicy = "\"\""
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
type LoadBalancerListResponseSessionAffinity string

const (
	LoadBalancerListResponseSessionAffinityNone     LoadBalancerListResponseSessionAffinity = "none"
	LoadBalancerListResponseSessionAffinityCookie   LoadBalancerListResponseSessionAffinity = "cookie"
	LoadBalancerListResponseSessionAffinityIPCookie LoadBalancerListResponseSessionAffinity = "ip_cookie"
	LoadBalancerListResponseSessionAffinityHeader   LoadBalancerListResponseSessionAffinity = "header"
	LoadBalancerListResponseSessionAffinityEmpty    LoadBalancerListResponseSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerListResponseSessionAffinityAttributes struct {
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
	Samesite LoadBalancerListResponseSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerListResponseSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerListResponseSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerListResponseSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerListResponseSessionAffinityAttributesJSON contains the JSON metadata
// for the struct [LoadBalancerListResponseSessionAffinityAttributes]
type loadBalancerListResponseSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerListResponseSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerListResponseSessionAffinityAttributesSamesite string

const (
	LoadBalancerListResponseSessionAffinityAttributesSamesiteAuto   LoadBalancerListResponseSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerListResponseSessionAffinityAttributesSamesiteLax    LoadBalancerListResponseSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerListResponseSessionAffinityAttributesSamesiteNone   LoadBalancerListResponseSessionAffinityAttributesSamesite = "None"
	LoadBalancerListResponseSessionAffinityAttributesSamesiteStrict LoadBalancerListResponseSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerListResponseSessionAffinityAttributesSecure string

const (
	LoadBalancerListResponseSessionAffinityAttributesSecureAuto   LoadBalancerListResponseSessionAffinityAttributesSecure = "Auto"
	LoadBalancerListResponseSessionAffinityAttributesSecureAlways LoadBalancerListResponseSessionAffinityAttributesSecure = "Always"
	LoadBalancerListResponseSessionAffinityAttributesSecureNever  LoadBalancerListResponseSessionAffinityAttributesSecure = "Never"
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
type LoadBalancerListResponseSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerListResponseSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerListResponseSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerListResponseSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerListResponseSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerListResponseSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerListResponseSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type LoadBalancerListResponseSteeringPolicy string

const (
	LoadBalancerListResponseSteeringPolicyOff                      LoadBalancerListResponseSteeringPolicy = "off"
	LoadBalancerListResponseSteeringPolicyGeo                      LoadBalancerListResponseSteeringPolicy = "geo"
	LoadBalancerListResponseSteeringPolicyRandom                   LoadBalancerListResponseSteeringPolicy = "random"
	LoadBalancerListResponseSteeringPolicyDynamicLatency           LoadBalancerListResponseSteeringPolicy = "dynamic_latency"
	LoadBalancerListResponseSteeringPolicyProximity                LoadBalancerListResponseSteeringPolicy = "proximity"
	LoadBalancerListResponseSteeringPolicyLeastOutstandingRequests LoadBalancerListResponseSteeringPolicy = "least_outstanding_requests"
	LoadBalancerListResponseSteeringPolicyLeastConnections         LoadBalancerListResponseSteeringPolicy = "least_connections"
	LoadBalancerListResponseSteeringPolicyEmpty                    LoadBalancerListResponseSteeringPolicy = "\"\""
)

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

type LoadBalancerEditResponse struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerEditResponseAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerEditResponseLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                                `json:"modified_on" format:"date-time"`
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
	RandomSteering LoadBalancerEditResponseRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancerEditResponseRule `json:"rules"`
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
	SessionAffinity LoadBalancerEditResponseSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerEditResponseSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerEditResponseSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                      `json:"ttl"`
	JSON loadBalancerEditResponseJSON `json:"-"`
}

// loadBalancerEditResponseJSON contains the JSON metadata for the struct
// [LoadBalancerEditResponse]
type loadBalancerEditResponseJSON struct {
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

func (r *LoadBalancerEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerEditResponseAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                        `json:"failover_across_pools"`
	JSON                loadBalancerEditResponseAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerEditResponseAdaptiveRoutingJSON contains the JSON metadata for the
// struct [LoadBalancerEditResponseAdaptiveRouting]
type loadBalancerEditResponseAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerEditResponseAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerEditResponseLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerEditResponseLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerEditResponseLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerEditResponseLocationStrategyJSON      `json:"-"`
}

// loadBalancerEditResponseLocationStrategyJSON contains the JSON metadata for the
// struct [LoadBalancerEditResponseLocationStrategy]
type loadBalancerEditResponseLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerEditResponseLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerEditResponseLocationStrategyMode string

const (
	LoadBalancerEditResponseLocationStrategyModePop        LoadBalancerEditResponseLocationStrategyMode = "pop"
	LoadBalancerEditResponseLocationStrategyModeResolverIP LoadBalancerEditResponseLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerEditResponseLocationStrategyPreferEcs string

const (
	LoadBalancerEditResponseLocationStrategyPreferEcsAlways    LoadBalancerEditResponseLocationStrategyPreferEcs = "always"
	LoadBalancerEditResponseLocationStrategyPreferEcsNever     LoadBalancerEditResponseLocationStrategyPreferEcs = "never"
	LoadBalancerEditResponseLocationStrategyPreferEcsProximity LoadBalancerEditResponseLocationStrategyPreferEcs = "proximity"
	LoadBalancerEditResponseLocationStrategyPreferEcsGeo       LoadBalancerEditResponseLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerEditResponseRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                `json:"pool_weights"`
	JSON        loadBalancerEditResponseRandomSteeringJSON `json:"-"`
}

// loadBalancerEditResponseRandomSteeringJSON contains the JSON metadata for the
// struct [LoadBalancerEditResponseRandomSteering]
type loadBalancerEditResponseRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerEditResponseRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerEditResponseRule struct {
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
	FixedResponse LoadBalancerEditResponseRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancerEditResponseRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                             `json:"terminates"`
	JSON       loadBalancerEditResponseRuleJSON `json:"-"`
}

// loadBalancerEditResponseRuleJSON contains the JSON metadata for the struct
// [LoadBalancerEditResponseRule]
type loadBalancerEditResponseRuleJSON struct {
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

func (r *LoadBalancerEditResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerEditResponseRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                          `json:"status_code"`
	JSON       loadBalancerEditResponseRulesFixedResponseJSON `json:"-"`
}

// loadBalancerEditResponseRulesFixedResponseJSON contains the JSON metadata for
// the struct [LoadBalancerEditResponseRulesFixedResponse]
type loadBalancerEditResponseRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerEditResponseRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerEditResponseRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerEditResponseRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerEditResponseRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering LoadBalancerEditResponseRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity LoadBalancerEditResponseRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerEditResponseRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerEditResponseRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                    `json:"ttl"`
	JSON loadBalancerEditResponseRulesOverridesJSON `json:"-"`
}

// loadBalancerEditResponseRulesOverridesJSON contains the JSON metadata for the
// struct [LoadBalancerEditResponseRulesOverrides]
type loadBalancerEditResponseRulesOverridesJSON struct {
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

func (r *LoadBalancerEditResponseRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerEditResponseRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                      `json:"failover_across_pools"`
	JSON                loadBalancerEditResponseRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerEditResponseRulesOverridesAdaptiveRoutingJSON contains the JSON
// metadata for the struct [LoadBalancerEditResponseRulesOverridesAdaptiveRouting]
type loadBalancerEditResponseRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerEditResponseRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerEditResponseRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerEditResponseRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerEditResponseRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerEditResponseRulesOverridesLocationStrategyJSON      `json:"-"`
}

// loadBalancerEditResponseRulesOverridesLocationStrategyJSON contains the JSON
// metadata for the struct [LoadBalancerEditResponseRulesOverridesLocationStrategy]
type loadBalancerEditResponseRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerEditResponseRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerEditResponseRulesOverridesLocationStrategyMode string

const (
	LoadBalancerEditResponseRulesOverridesLocationStrategyModePop        LoadBalancerEditResponseRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerEditResponseRulesOverridesLocationStrategyModeResolverIP LoadBalancerEditResponseRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerEditResponseRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerEditResponseRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerEditResponseRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerEditResponseRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerEditResponseRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerEditResponseRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerEditResponseRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerEditResponseRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerEditResponseRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerEditResponseRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                              `json:"pool_weights"`
	JSON        loadBalancerEditResponseRulesOverridesRandomSteeringJSON `json:"-"`
}

// loadBalancerEditResponseRulesOverridesRandomSteeringJSON contains the JSON
// metadata for the struct [LoadBalancerEditResponseRulesOverridesRandomSteering]
type loadBalancerEditResponseRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerEditResponseRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
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
type LoadBalancerEditResponseRulesOverridesSessionAffinity string

const (
	LoadBalancerEditResponseRulesOverridesSessionAffinityNone     LoadBalancerEditResponseRulesOverridesSessionAffinity = "none"
	LoadBalancerEditResponseRulesOverridesSessionAffinityCookie   LoadBalancerEditResponseRulesOverridesSessionAffinity = "cookie"
	LoadBalancerEditResponseRulesOverridesSessionAffinityIPCookie LoadBalancerEditResponseRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerEditResponseRulesOverridesSessionAffinityHeader   LoadBalancerEditResponseRulesOverridesSessionAffinity = "header"
	LoadBalancerEditResponseRulesOverridesSessionAffinityEmpty    LoadBalancerEditResponseRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerEditResponseRulesOverridesSessionAffinityAttributes struct {
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
	Samesite LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerEditResponseRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerEditResponseRulesOverridesSessionAffinityAttributesJSON contains the
// JSON metadata for the struct
// [LoadBalancerEditResponseRulesOverridesSessionAffinityAttributes]
type loadBalancerEditResponseRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerEditResponseRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesSecure = "Never"
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
type LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerEditResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type LoadBalancerEditResponseRulesOverridesSteeringPolicy string

const (
	LoadBalancerEditResponseRulesOverridesSteeringPolicyOff                      LoadBalancerEditResponseRulesOverridesSteeringPolicy = "off"
	LoadBalancerEditResponseRulesOverridesSteeringPolicyGeo                      LoadBalancerEditResponseRulesOverridesSteeringPolicy = "geo"
	LoadBalancerEditResponseRulesOverridesSteeringPolicyRandom                   LoadBalancerEditResponseRulesOverridesSteeringPolicy = "random"
	LoadBalancerEditResponseRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerEditResponseRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerEditResponseRulesOverridesSteeringPolicyProximity                LoadBalancerEditResponseRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerEditResponseRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerEditResponseRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerEditResponseRulesOverridesSteeringPolicyLeastConnections         LoadBalancerEditResponseRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerEditResponseRulesOverridesSteeringPolicyEmpty                    LoadBalancerEditResponseRulesOverridesSteeringPolicy = "\"\""
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
type LoadBalancerEditResponseSessionAffinity string

const (
	LoadBalancerEditResponseSessionAffinityNone     LoadBalancerEditResponseSessionAffinity = "none"
	LoadBalancerEditResponseSessionAffinityCookie   LoadBalancerEditResponseSessionAffinity = "cookie"
	LoadBalancerEditResponseSessionAffinityIPCookie LoadBalancerEditResponseSessionAffinity = "ip_cookie"
	LoadBalancerEditResponseSessionAffinityHeader   LoadBalancerEditResponseSessionAffinity = "header"
	LoadBalancerEditResponseSessionAffinityEmpty    LoadBalancerEditResponseSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerEditResponseSessionAffinityAttributes struct {
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
	Samesite LoadBalancerEditResponseSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerEditResponseSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerEditResponseSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerEditResponseSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerEditResponseSessionAffinityAttributesJSON contains the JSON metadata
// for the struct [LoadBalancerEditResponseSessionAffinityAttributes]
type loadBalancerEditResponseSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerEditResponseSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerEditResponseSessionAffinityAttributesSamesite string

const (
	LoadBalancerEditResponseSessionAffinityAttributesSamesiteAuto   LoadBalancerEditResponseSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerEditResponseSessionAffinityAttributesSamesiteLax    LoadBalancerEditResponseSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerEditResponseSessionAffinityAttributesSamesiteNone   LoadBalancerEditResponseSessionAffinityAttributesSamesite = "None"
	LoadBalancerEditResponseSessionAffinityAttributesSamesiteStrict LoadBalancerEditResponseSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerEditResponseSessionAffinityAttributesSecure string

const (
	LoadBalancerEditResponseSessionAffinityAttributesSecureAuto   LoadBalancerEditResponseSessionAffinityAttributesSecure = "Auto"
	LoadBalancerEditResponseSessionAffinityAttributesSecureAlways LoadBalancerEditResponseSessionAffinityAttributesSecure = "Always"
	LoadBalancerEditResponseSessionAffinityAttributesSecureNever  LoadBalancerEditResponseSessionAffinityAttributesSecure = "Never"
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
type LoadBalancerEditResponseSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerEditResponseSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerEditResponseSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerEditResponseSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerEditResponseSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerEditResponseSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerEditResponseSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type LoadBalancerEditResponseSteeringPolicy string

const (
	LoadBalancerEditResponseSteeringPolicyOff                      LoadBalancerEditResponseSteeringPolicy = "off"
	LoadBalancerEditResponseSteeringPolicyGeo                      LoadBalancerEditResponseSteeringPolicy = "geo"
	LoadBalancerEditResponseSteeringPolicyRandom                   LoadBalancerEditResponseSteeringPolicy = "random"
	LoadBalancerEditResponseSteeringPolicyDynamicLatency           LoadBalancerEditResponseSteeringPolicy = "dynamic_latency"
	LoadBalancerEditResponseSteeringPolicyProximity                LoadBalancerEditResponseSteeringPolicy = "proximity"
	LoadBalancerEditResponseSteeringPolicyLeastOutstandingRequests LoadBalancerEditResponseSteeringPolicy = "least_outstanding_requests"
	LoadBalancerEditResponseSteeringPolicyLeastConnections         LoadBalancerEditResponseSteeringPolicy = "least_connections"
	LoadBalancerEditResponseSteeringPolicyEmpty                    LoadBalancerEditResponseSteeringPolicy = "\"\""
)

type LoadBalancerGetResponse struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerGetResponseAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerGetResponseLocationStrategy `json:"location_strategy"`
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
	RandomSteering LoadBalancerGetResponseRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancerGetResponseRule `json:"rules"`
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
	SessionAffinity LoadBalancerGetResponseSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerGetResponseSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerGetResponseSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                     `json:"ttl"`
	JSON loadBalancerGetResponseJSON `json:"-"`
}

// loadBalancerGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancerGetResponse]
type loadBalancerGetResponseJSON struct {
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

func (r *LoadBalancerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerGetResponseAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                       `json:"failover_across_pools"`
	JSON                loadBalancerGetResponseAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerGetResponseAdaptiveRoutingJSON contains the JSON metadata for the
// struct [LoadBalancerGetResponseAdaptiveRouting]
type loadBalancerGetResponseAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerGetResponseAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerGetResponseLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerGetResponseLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerGetResponseLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerGetResponseLocationStrategyJSON      `json:"-"`
}

// loadBalancerGetResponseLocationStrategyJSON contains the JSON metadata for the
// struct [LoadBalancerGetResponseLocationStrategy]
type loadBalancerGetResponseLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerGetResponseLocationStrategyMode string

const (
	LoadBalancerGetResponseLocationStrategyModePop        LoadBalancerGetResponseLocationStrategyMode = "pop"
	LoadBalancerGetResponseLocationStrategyModeResolverIP LoadBalancerGetResponseLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerGetResponseLocationStrategyPreferEcs string

const (
	LoadBalancerGetResponseLocationStrategyPreferEcsAlways    LoadBalancerGetResponseLocationStrategyPreferEcs = "always"
	LoadBalancerGetResponseLocationStrategyPreferEcsNever     LoadBalancerGetResponseLocationStrategyPreferEcs = "never"
	LoadBalancerGetResponseLocationStrategyPreferEcsProximity LoadBalancerGetResponseLocationStrategyPreferEcs = "proximity"
	LoadBalancerGetResponseLocationStrategyPreferEcsGeo       LoadBalancerGetResponseLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerGetResponseRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                               `json:"pool_weights"`
	JSON        loadBalancerGetResponseRandomSteeringJSON `json:"-"`
}

// loadBalancerGetResponseRandomSteeringJSON contains the JSON metadata for the
// struct [LoadBalancerGetResponseRandomSteering]
type loadBalancerGetResponseRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerGetResponseRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerGetResponseRule struct {
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
	FixedResponse LoadBalancerGetResponseRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancerGetResponseRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                            `json:"terminates"`
	JSON       loadBalancerGetResponseRuleJSON `json:"-"`
}

// loadBalancerGetResponseRuleJSON contains the JSON metadata for the struct
// [LoadBalancerGetResponseRule]
type loadBalancerGetResponseRuleJSON struct {
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

func (r *LoadBalancerGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerGetResponseRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                         `json:"status_code"`
	JSON       loadBalancerGetResponseRulesFixedResponseJSON `json:"-"`
}

// loadBalancerGetResponseRulesFixedResponseJSON contains the JSON metadata for the
// struct [LoadBalancerGetResponseRulesFixedResponse]
type loadBalancerGetResponseRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerGetResponseRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerGetResponseRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerGetResponseRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering LoadBalancerGetResponseRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity LoadBalancerGetResponseRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes LoadBalancerGetResponseRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy LoadBalancerGetResponseRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                   `json:"ttl"`
	JSON loadBalancerGetResponseRulesOverridesJSON `json:"-"`
}

// loadBalancerGetResponseRulesOverridesJSON contains the JSON metadata for the
// struct [LoadBalancerGetResponseRulesOverrides]
type loadBalancerGetResponseRulesOverridesJSON struct {
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

func (r *LoadBalancerGetResponseRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerGetResponseRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                     `json:"failover_across_pools"`
	JSON                loadBalancerGetResponseRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerGetResponseRulesOverridesAdaptiveRoutingJSON contains the JSON
// metadata for the struct [LoadBalancerGetResponseRulesOverridesAdaptiveRouting]
type loadBalancerGetResponseRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerGetResponseRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerGetResponseRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerGetResponseRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerGetResponseRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerGetResponseRulesOverridesLocationStrategyJSON      `json:"-"`
}

// loadBalancerGetResponseRulesOverridesLocationStrategyJSON contains the JSON
// metadata for the struct [LoadBalancerGetResponseRulesOverridesLocationStrategy]
type loadBalancerGetResponseRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerGetResponseRulesOverridesLocationStrategyMode string

const (
	LoadBalancerGetResponseRulesOverridesLocationStrategyModePop        LoadBalancerGetResponseRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerGetResponseRulesOverridesLocationStrategyModeResolverIP LoadBalancerGetResponseRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerGetResponseRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerGetResponseRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerGetResponseRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerGetResponseRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerGetResponseRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerGetResponseRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerGetResponseRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerGetResponseRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerGetResponseRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type LoadBalancerGetResponseRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                             `json:"pool_weights"`
	JSON        loadBalancerGetResponseRulesOverridesRandomSteeringJSON `json:"-"`
}

// loadBalancerGetResponseRulesOverridesRandomSteeringJSON contains the JSON
// metadata for the struct [LoadBalancerGetResponseRulesOverridesRandomSteering]
type loadBalancerGetResponseRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerGetResponseRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
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
type LoadBalancerGetResponseRulesOverridesSessionAffinity string

const (
	LoadBalancerGetResponseRulesOverridesSessionAffinityNone     LoadBalancerGetResponseRulesOverridesSessionAffinity = "none"
	LoadBalancerGetResponseRulesOverridesSessionAffinityCookie   LoadBalancerGetResponseRulesOverridesSessionAffinity = "cookie"
	LoadBalancerGetResponseRulesOverridesSessionAffinityIPCookie LoadBalancerGetResponseRulesOverridesSessionAffinity = "ip_cookie"
	LoadBalancerGetResponseRulesOverridesSessionAffinityHeader   LoadBalancerGetResponseRulesOverridesSessionAffinity = "header"
	LoadBalancerGetResponseRulesOverridesSessionAffinityEmpty    LoadBalancerGetResponseRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerGetResponseRulesOverridesSessionAffinityAttributes struct {
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
	Samesite LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerGetResponseRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerGetResponseRulesOverridesSessionAffinityAttributesJSON contains the
// JSON metadata for the struct
// [LoadBalancerGetResponseRulesOverridesSessionAffinityAttributes]
type loadBalancerGetResponseRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerGetResponseRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesSecure = "Never"
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
type LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerGetResponseRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type LoadBalancerGetResponseRulesOverridesSteeringPolicy string

const (
	LoadBalancerGetResponseRulesOverridesSteeringPolicyOff                      LoadBalancerGetResponseRulesOverridesSteeringPolicy = "off"
	LoadBalancerGetResponseRulesOverridesSteeringPolicyGeo                      LoadBalancerGetResponseRulesOverridesSteeringPolicy = "geo"
	LoadBalancerGetResponseRulesOverridesSteeringPolicyRandom                   LoadBalancerGetResponseRulesOverridesSteeringPolicy = "random"
	LoadBalancerGetResponseRulesOverridesSteeringPolicyDynamicLatency           LoadBalancerGetResponseRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerGetResponseRulesOverridesSteeringPolicyProximity                LoadBalancerGetResponseRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerGetResponseRulesOverridesSteeringPolicyLeastOutstandingRequests LoadBalancerGetResponseRulesOverridesSteeringPolicy = "least_outstanding_requests"
	LoadBalancerGetResponseRulesOverridesSteeringPolicyLeastConnections         LoadBalancerGetResponseRulesOverridesSteeringPolicy = "least_connections"
	LoadBalancerGetResponseRulesOverridesSteeringPolicyEmpty                    LoadBalancerGetResponseRulesOverridesSteeringPolicy = "\"\""
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
type LoadBalancerGetResponseSessionAffinity string

const (
	LoadBalancerGetResponseSessionAffinityNone     LoadBalancerGetResponseSessionAffinity = "none"
	LoadBalancerGetResponseSessionAffinityCookie   LoadBalancerGetResponseSessionAffinity = "cookie"
	LoadBalancerGetResponseSessionAffinityIPCookie LoadBalancerGetResponseSessionAffinity = "ip_cookie"
	LoadBalancerGetResponseSessionAffinityHeader   LoadBalancerGetResponseSessionAffinity = "header"
	LoadBalancerGetResponseSessionAffinityEmpty    LoadBalancerGetResponseSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type LoadBalancerGetResponseSessionAffinityAttributes struct {
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
	Samesite LoadBalancerGetResponseSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerGetResponseSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover LoadBalancerGetResponseSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerGetResponseSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerGetResponseSessionAffinityAttributesJSON contains the JSON metadata
// for the struct [LoadBalancerGetResponseSessionAffinityAttributes]
type loadBalancerGetResponseSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerGetResponseSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerGetResponseSessionAffinityAttributesSamesite string

const (
	LoadBalancerGetResponseSessionAffinityAttributesSamesiteAuto   LoadBalancerGetResponseSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerGetResponseSessionAffinityAttributesSamesiteLax    LoadBalancerGetResponseSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerGetResponseSessionAffinityAttributesSamesiteNone   LoadBalancerGetResponseSessionAffinityAttributesSamesite = "None"
	LoadBalancerGetResponseSessionAffinityAttributesSamesiteStrict LoadBalancerGetResponseSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerGetResponseSessionAffinityAttributesSecure string

const (
	LoadBalancerGetResponseSessionAffinityAttributesSecureAuto   LoadBalancerGetResponseSessionAffinityAttributesSecure = "Auto"
	LoadBalancerGetResponseSessionAffinityAttributesSecureAlways LoadBalancerGetResponseSessionAffinityAttributesSecure = "Always"
	LoadBalancerGetResponseSessionAffinityAttributesSecureNever  LoadBalancerGetResponseSessionAffinityAttributesSecure = "Never"
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
type LoadBalancerGetResponseSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerGetResponseSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerGetResponseSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerGetResponseSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerGetResponseSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerGetResponseSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerGetResponseSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type LoadBalancerGetResponseSteeringPolicy string

const (
	LoadBalancerGetResponseSteeringPolicyOff                      LoadBalancerGetResponseSteeringPolicy = "off"
	LoadBalancerGetResponseSteeringPolicyGeo                      LoadBalancerGetResponseSteeringPolicy = "geo"
	LoadBalancerGetResponseSteeringPolicyRandom                   LoadBalancerGetResponseSteeringPolicy = "random"
	LoadBalancerGetResponseSteeringPolicyDynamicLatency           LoadBalancerGetResponseSteeringPolicy = "dynamic_latency"
	LoadBalancerGetResponseSteeringPolicyProximity                LoadBalancerGetResponseSteeringPolicy = "proximity"
	LoadBalancerGetResponseSteeringPolicyLeastOutstandingRequests LoadBalancerGetResponseSteeringPolicy = "least_outstanding_requests"
	LoadBalancerGetResponseSteeringPolicyLeastConnections         LoadBalancerGetResponseSteeringPolicy = "least_connections"
	LoadBalancerGetResponseSteeringPolicyEmpty                    LoadBalancerGetResponseSteeringPolicy = "\"\""
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
	Errors   []LoadBalancerNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerNewResponse                   `json:"result,required"`
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

type LoadBalancerListResponseEnvelope struct {
	Errors   []LoadBalancerListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerListResponseEnvelopeMessages `json:"messages,required"`
	Result   []LoadBalancerListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerListResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerListResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerListResponseEnvelope]
type loadBalancerListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    loadBalancerListResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LoadBalancerListResponseEnvelopeErrors]
type loadBalancerListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    loadBalancerListResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LoadBalancerListResponseEnvelopeMessages]
type loadBalancerListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerListResponseEnvelopeSuccess bool

const (
	LoadBalancerListResponseEnvelopeSuccessTrue LoadBalancerListResponseEnvelopeSuccess = true
)

type LoadBalancerListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       loadBalancerListResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [LoadBalancerListResponseEnvelopeResultInfo]
type loadBalancerListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerDeleteResponseEnvelope struct {
	Errors   []LoadBalancerDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerDeleteResponse                   `json:"result,required"`
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

type LoadBalancerDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    loadBalancerDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LoadBalancerDeleteResponseEnvelopeErrors]
type loadBalancerDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    loadBalancerDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LoadBalancerDeleteResponseEnvelopeMessages]
type loadBalancerDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerDeleteResponseEnvelopeSuccess bool

const (
	LoadBalancerDeleteResponseEnvelopeSuccessTrue LoadBalancerDeleteResponseEnvelopeSuccess = true
)

type LoadBalancerEditParams struct {
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

type LoadBalancerEditResponseEnvelope struct {
	Errors   []LoadBalancerEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerEditResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerEditResponse                   `json:"result,required"`
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

type LoadBalancerEditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    loadBalancerEditResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LoadBalancerEditResponseEnvelopeErrors]
type loadBalancerEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerEditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    loadBalancerEditResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LoadBalancerEditResponseEnvelopeMessages]
type loadBalancerEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerEditResponseEnvelopeSuccess bool

const (
	LoadBalancerEditResponseEnvelopeSuccessTrue LoadBalancerEditResponseEnvelopeSuccess = true
)

type LoadBalancerGetResponseEnvelope struct {
	Errors   []LoadBalancerGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerGetResponse                   `json:"result,required"`
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

type LoadBalancerGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    loadBalancerGetResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LoadBalancerGetResponseEnvelopeErrors]
type loadBalancerGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    loadBalancerGetResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LoadBalancerGetResponseEnvelopeMessages]
type loadBalancerGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerGetResponseEnvelopeSuccess bool

const (
	LoadBalancerGetResponseEnvelopeSuccessTrue LoadBalancerGetResponseEnvelopeSuccess = true
)
