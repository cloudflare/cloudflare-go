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

// ZoneLoadBalancerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneLoadBalancerService] method
// instead.
type ZoneLoadBalancerService struct {
	Options []option.RequestOption
}

// NewZoneLoadBalancerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneLoadBalancerService(opts ...option.RequestOption) (r *ZoneLoadBalancerService) {
	r = &ZoneLoadBalancerService{}
	r.Options = opts
	return
}

// Create a new load balancer.
func (r *ZoneLoadBalancerService) New(ctx context.Context, identifier string, body ZoneLoadBalancerNewParams, opts ...option.RequestOption) (res *ZoneLoadBalancerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/load_balancers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single configured load balancer.
func (r *ZoneLoadBalancerService) Get(ctx context.Context, identifier1 string, identifier string, opts ...option.RequestOption) (res *ZoneLoadBalancerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/load_balancers/%s", identifier1, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured load balancer.
func (r *ZoneLoadBalancerService) Update(ctx context.Context, identifier1 string, identifier string, body ZoneLoadBalancerUpdateParams, opts ...option.RequestOption) (res *ZoneLoadBalancerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/load_balancers/%s", identifier1, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List configured load balancers.
func (r *ZoneLoadBalancerService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneLoadBalancerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/load_balancers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a configured load balancer.
func (r *ZoneLoadBalancerService) Delete(ctx context.Context, identifier1 string, identifier string, opts ...option.RequestOption) (res *ZoneLoadBalancerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/load_balancers/%s", identifier1, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ZoneLoadBalancerNewResponse struct {
	Errors   []ZoneLoadBalancerNewResponseError   `json:"errors"`
	Messages []ZoneLoadBalancerNewResponseMessage `json:"messages"`
	Result   ZoneLoadBalancerNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneLoadBalancerNewResponseSuccess `json:"success"`
	JSON    zoneLoadBalancerNewResponseJSON    `json:"-"`
}

// zoneLoadBalancerNewResponseJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerNewResponse]
type zoneLoadBalancerNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerNewResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneLoadBalancerNewResponseErrorJSON `json:"-"`
}

// zoneLoadBalancerNewResponseErrorJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerNewResponseError]
type zoneLoadBalancerNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerNewResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneLoadBalancerNewResponseMessageJSON `json:"-"`
}

// zoneLoadBalancerNewResponseMessageJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerNewResponseMessage]
type zoneLoadBalancerNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerNewResponseResult struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting ZoneLoadBalancerNewResponseResultAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy ZoneLoadBalancerNewResponseResultLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                                         `json:"modified_on" format:"date-time"`
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
	RandomSteering ZoneLoadBalancerNewResponseResultRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []ZoneLoadBalancerNewResponseResultRule `json:"rules"`
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
	SessionAffinity ZoneLoadBalancerNewResponseResultSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes ZoneLoadBalancerNewResponseResultSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy ZoneLoadBalancerNewResponseResultSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                               `json:"ttl"`
	JSON zoneLoadBalancerNewResponseResultJSON `json:"-"`
}

// zoneLoadBalancerNewResponseResultJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerNewResponseResult]
type zoneLoadBalancerNewResponseResultJSON struct {
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

func (r *ZoneLoadBalancerNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerNewResponseResultAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                 `json:"failover_across_pools"`
	JSON                zoneLoadBalancerNewResponseResultAdaptiveRoutingJSON `json:"-"`
}

// zoneLoadBalancerNewResponseResultAdaptiveRoutingJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerNewResponseResultAdaptiveRouting]
type zoneLoadBalancerNewResponseResultAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseResultAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerNewResponseResultLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode ZoneLoadBalancerNewResponseResultLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs ZoneLoadBalancerNewResponseResultLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      zoneLoadBalancerNewResponseResultLocationStrategyJSON      `json:"-"`
}

// zoneLoadBalancerNewResponseResultLocationStrategyJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerNewResponseResultLocationStrategy]
type zoneLoadBalancerNewResponseResultLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseResultLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerNewResponseResultLocationStrategyMode string

const (
	ZoneLoadBalancerNewResponseResultLocationStrategyModePop        ZoneLoadBalancerNewResponseResultLocationStrategyMode = "pop"
	ZoneLoadBalancerNewResponseResultLocationStrategyModeResolverIP ZoneLoadBalancerNewResponseResultLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerNewResponseResultLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerNewResponseResultLocationStrategyPreferEcsAlways    ZoneLoadBalancerNewResponseResultLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerNewResponseResultLocationStrategyPreferEcsNever     ZoneLoadBalancerNewResponseResultLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerNewResponseResultLocationStrategyPreferEcsProximity ZoneLoadBalancerNewResponseResultLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerNewResponseResultLocationStrategyPreferEcsGeo       ZoneLoadBalancerNewResponseResultLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerNewResponseResultRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                         `json:"pool_weights"`
	JSON        zoneLoadBalancerNewResponseResultRandomSteeringJSON `json:"-"`
}

// zoneLoadBalancerNewResponseResultRandomSteeringJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerNewResponseResultRandomSteering]
type zoneLoadBalancerNewResponseResultRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseResultRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type ZoneLoadBalancerNewResponseResultRule struct {
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
	FixedResponse ZoneLoadBalancerNewResponseResultRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides ZoneLoadBalancerNewResponseResultRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                                      `json:"terminates"`
	JSON       zoneLoadBalancerNewResponseResultRuleJSON `json:"-"`
}

// zoneLoadBalancerNewResponseResultRuleJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerNewResponseResultRule]
type zoneLoadBalancerNewResponseResultRuleJSON struct {
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

func (r *ZoneLoadBalancerNewResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type ZoneLoadBalancerNewResponseResultRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                                   `json:"status_code"`
	JSON       zoneLoadBalancerNewResponseResultRulesFixedResponseJSON `json:"-"`
}

// zoneLoadBalancerNewResponseResultRulesFixedResponseJSON contains the JSON
// metadata for the struct [ZoneLoadBalancerNewResponseResultRulesFixedResponse]
type zoneLoadBalancerNewResponseResultRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseResultRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type ZoneLoadBalancerNewResponseResultRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting ZoneLoadBalancerNewResponseResultRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering ZoneLoadBalancerNewResponseResultRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                             `json:"ttl"`
	JSON zoneLoadBalancerNewResponseResultRulesOverridesJSON `json:"-"`
}

// zoneLoadBalancerNewResponseResultRulesOverridesJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerNewResponseResultRulesOverrides]
type zoneLoadBalancerNewResponseResultRulesOverridesJSON struct {
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

func (r *ZoneLoadBalancerNewResponseResultRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerNewResponseResultRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                               `json:"failover_across_pools"`
	JSON                zoneLoadBalancerNewResponseResultRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// zoneLoadBalancerNewResponseResultRulesOverridesAdaptiveRoutingJSON contains the
// JSON metadata for the struct
// [ZoneLoadBalancerNewResponseResultRulesOverridesAdaptiveRouting]
type zoneLoadBalancerNewResponseResultRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseResultRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      zoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyJSON      `json:"-"`
}

// zoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyJSON contains the
// JSON metadata for the struct
// [ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategy]
type zoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyMode string

const (
	ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyModePop        ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyMode = "pop"
	ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyModeResolverIP ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyPreferEcsAlways    ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyPreferEcsNever     ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyPreferEcsProximity ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyPreferEcsGeo       ZoneLoadBalancerNewResponseResultRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerNewResponseResultRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                                       `json:"pool_weights"`
	JSON        zoneLoadBalancerNewResponseResultRulesOverridesRandomSteeringJSON `json:"-"`
}

// zoneLoadBalancerNewResponseResultRulesOverridesRandomSteeringJSON contains the
// JSON metadata for the struct
// [ZoneLoadBalancerNewResponseResultRulesOverridesRandomSteering]
type zoneLoadBalancerNewResponseResultRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseResultRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
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
type ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinity string

const (
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityNone     ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinity = "none"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityCookie   ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinity = "cookie"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityIPCookie ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinity = "ip_cookie"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityHeader   ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinity = "header"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityEmpty    ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributes struct {
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
	Samesite ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 zoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// zoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesJSON
// contains the JSON metadata for the struct
// [ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributes]
type zoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSecureAuto   ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSecureAlways ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSecureNever  ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerNewResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicy string

const (
	ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicyOff                      ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicy = "off"
	ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicyGeo                      ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicy = "geo"
	ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicyRandom                   ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicy = "random"
	ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicyDynamicLatency           ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicyProximity                ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicy = "proximity"
	ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicyLeastConnections         ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicy = "least_connections"
	ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicyEmpty                    ZoneLoadBalancerNewResponseResultRulesOverridesSteeringPolicy = "\"\""
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
type ZoneLoadBalancerNewResponseResultSessionAffinity string

const (
	ZoneLoadBalancerNewResponseResultSessionAffinityNone     ZoneLoadBalancerNewResponseResultSessionAffinity = "none"
	ZoneLoadBalancerNewResponseResultSessionAffinityCookie   ZoneLoadBalancerNewResponseResultSessionAffinity = "cookie"
	ZoneLoadBalancerNewResponseResultSessionAffinityIPCookie ZoneLoadBalancerNewResponseResultSessionAffinity = "ip_cookie"
	ZoneLoadBalancerNewResponseResultSessionAffinityHeader   ZoneLoadBalancerNewResponseResultSessionAffinity = "header"
	ZoneLoadBalancerNewResponseResultSessionAffinityEmpty    ZoneLoadBalancerNewResponseResultSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerNewResponseResultSessionAffinityAttributes struct {
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
	Samesite ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover ZoneLoadBalancerNewResponseResultSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 zoneLoadBalancerNewResponseResultSessionAffinityAttributesJSON                 `json:"-"`
}

// zoneLoadBalancerNewResponseResultSessionAffinityAttributesJSON contains the JSON
// metadata for the struct
// [ZoneLoadBalancerNewResponseResultSessionAffinityAttributes]
type zoneLoadBalancerNewResponseResultSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneLoadBalancerNewResponseResultSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSecureAuto   ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSecureAlways ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSecureNever  ZoneLoadBalancerNewResponseResultSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerNewResponseResultSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerNewResponseResultSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerNewResponseResultSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerNewResponseResultSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerNewResponseResultSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerNewResponseResultSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerNewResponseResultSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerNewResponseResultSteeringPolicy string

const (
	ZoneLoadBalancerNewResponseResultSteeringPolicyOff                      ZoneLoadBalancerNewResponseResultSteeringPolicy = "off"
	ZoneLoadBalancerNewResponseResultSteeringPolicyGeo                      ZoneLoadBalancerNewResponseResultSteeringPolicy = "geo"
	ZoneLoadBalancerNewResponseResultSteeringPolicyRandom                   ZoneLoadBalancerNewResponseResultSteeringPolicy = "random"
	ZoneLoadBalancerNewResponseResultSteeringPolicyDynamicLatency           ZoneLoadBalancerNewResponseResultSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerNewResponseResultSteeringPolicyProximity                ZoneLoadBalancerNewResponseResultSteeringPolicy = "proximity"
	ZoneLoadBalancerNewResponseResultSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerNewResponseResultSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerNewResponseResultSteeringPolicyLeastConnections         ZoneLoadBalancerNewResponseResultSteeringPolicy = "least_connections"
	ZoneLoadBalancerNewResponseResultSteeringPolicyEmpty                    ZoneLoadBalancerNewResponseResultSteeringPolicy = "\"\""
)

// Whether the API call was successful
type ZoneLoadBalancerNewResponseSuccess bool

const (
	ZoneLoadBalancerNewResponseSuccessTrue ZoneLoadBalancerNewResponseSuccess = true
)

type ZoneLoadBalancerGetResponse struct {
	Errors   []ZoneLoadBalancerGetResponseError   `json:"errors"`
	Messages []ZoneLoadBalancerGetResponseMessage `json:"messages"`
	Result   ZoneLoadBalancerGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneLoadBalancerGetResponseSuccess `json:"success"`
	JSON    zoneLoadBalancerGetResponseJSON    `json:"-"`
}

// zoneLoadBalancerGetResponseJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerGetResponse]
type zoneLoadBalancerGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneLoadBalancerGetResponseErrorJSON `json:"-"`
}

// zoneLoadBalancerGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerGetResponseError]
type zoneLoadBalancerGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneLoadBalancerGetResponseMessageJSON `json:"-"`
}

// zoneLoadBalancerGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerGetResponseMessage]
type zoneLoadBalancerGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerGetResponseResult struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting ZoneLoadBalancerGetResponseResultAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy ZoneLoadBalancerGetResponseResultLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                                         `json:"modified_on" format:"date-time"`
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
	RandomSteering ZoneLoadBalancerGetResponseResultRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []ZoneLoadBalancerGetResponseResultRule `json:"rules"`
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
	SessionAffinity ZoneLoadBalancerGetResponseResultSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes ZoneLoadBalancerGetResponseResultSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy ZoneLoadBalancerGetResponseResultSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                               `json:"ttl"`
	JSON zoneLoadBalancerGetResponseResultJSON `json:"-"`
}

// zoneLoadBalancerGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerGetResponseResult]
type zoneLoadBalancerGetResponseResultJSON struct {
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

func (r *ZoneLoadBalancerGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerGetResponseResultAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                 `json:"failover_across_pools"`
	JSON                zoneLoadBalancerGetResponseResultAdaptiveRoutingJSON `json:"-"`
}

// zoneLoadBalancerGetResponseResultAdaptiveRoutingJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerGetResponseResultAdaptiveRouting]
type zoneLoadBalancerGetResponseResultAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseResultAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerGetResponseResultLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode ZoneLoadBalancerGetResponseResultLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs ZoneLoadBalancerGetResponseResultLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      zoneLoadBalancerGetResponseResultLocationStrategyJSON      `json:"-"`
}

// zoneLoadBalancerGetResponseResultLocationStrategyJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerGetResponseResultLocationStrategy]
type zoneLoadBalancerGetResponseResultLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseResultLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerGetResponseResultLocationStrategyMode string

const (
	ZoneLoadBalancerGetResponseResultLocationStrategyModePop        ZoneLoadBalancerGetResponseResultLocationStrategyMode = "pop"
	ZoneLoadBalancerGetResponseResultLocationStrategyModeResolverIP ZoneLoadBalancerGetResponseResultLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerGetResponseResultLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerGetResponseResultLocationStrategyPreferEcsAlways    ZoneLoadBalancerGetResponseResultLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerGetResponseResultLocationStrategyPreferEcsNever     ZoneLoadBalancerGetResponseResultLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerGetResponseResultLocationStrategyPreferEcsProximity ZoneLoadBalancerGetResponseResultLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerGetResponseResultLocationStrategyPreferEcsGeo       ZoneLoadBalancerGetResponseResultLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerGetResponseResultRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                         `json:"pool_weights"`
	JSON        zoneLoadBalancerGetResponseResultRandomSteeringJSON `json:"-"`
}

// zoneLoadBalancerGetResponseResultRandomSteeringJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerGetResponseResultRandomSteering]
type zoneLoadBalancerGetResponseResultRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseResultRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type ZoneLoadBalancerGetResponseResultRule struct {
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
	FixedResponse ZoneLoadBalancerGetResponseResultRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides ZoneLoadBalancerGetResponseResultRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                                      `json:"terminates"`
	JSON       zoneLoadBalancerGetResponseResultRuleJSON `json:"-"`
}

// zoneLoadBalancerGetResponseResultRuleJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerGetResponseResultRule]
type zoneLoadBalancerGetResponseResultRuleJSON struct {
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

func (r *ZoneLoadBalancerGetResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type ZoneLoadBalancerGetResponseResultRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                                   `json:"status_code"`
	JSON       zoneLoadBalancerGetResponseResultRulesFixedResponseJSON `json:"-"`
}

// zoneLoadBalancerGetResponseResultRulesFixedResponseJSON contains the JSON
// metadata for the struct [ZoneLoadBalancerGetResponseResultRulesFixedResponse]
type zoneLoadBalancerGetResponseResultRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseResultRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type ZoneLoadBalancerGetResponseResultRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting ZoneLoadBalancerGetResponseResultRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering ZoneLoadBalancerGetResponseResultRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                             `json:"ttl"`
	JSON zoneLoadBalancerGetResponseResultRulesOverridesJSON `json:"-"`
}

// zoneLoadBalancerGetResponseResultRulesOverridesJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerGetResponseResultRulesOverrides]
type zoneLoadBalancerGetResponseResultRulesOverridesJSON struct {
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

func (r *ZoneLoadBalancerGetResponseResultRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerGetResponseResultRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                               `json:"failover_across_pools"`
	JSON                zoneLoadBalancerGetResponseResultRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// zoneLoadBalancerGetResponseResultRulesOverridesAdaptiveRoutingJSON contains the
// JSON metadata for the struct
// [ZoneLoadBalancerGetResponseResultRulesOverridesAdaptiveRouting]
type zoneLoadBalancerGetResponseResultRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseResultRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      zoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyJSON      `json:"-"`
}

// zoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyJSON contains the
// JSON metadata for the struct
// [ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategy]
type zoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyMode string

const (
	ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyModePop        ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyMode = "pop"
	ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyModeResolverIP ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcsAlways    ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcsNever     ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcsProximity ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcsGeo       ZoneLoadBalancerGetResponseResultRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerGetResponseResultRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                                       `json:"pool_weights"`
	JSON        zoneLoadBalancerGetResponseResultRulesOverridesRandomSteeringJSON `json:"-"`
}

// zoneLoadBalancerGetResponseResultRulesOverridesRandomSteeringJSON contains the
// JSON metadata for the struct
// [ZoneLoadBalancerGetResponseResultRulesOverridesRandomSteering]
type zoneLoadBalancerGetResponseResultRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseResultRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
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
type ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinity string

const (
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityNone     ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinity = "none"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityCookie   ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinity = "cookie"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityIPCookie ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinity = "ip_cookie"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityHeader   ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinity = "header"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityEmpty    ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributes struct {
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
	Samesite ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 zoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// zoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesJSON
// contains the JSON metadata for the struct
// [ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributes]
type zoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecureAuto   ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecureAlways ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecureNever  ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerGetResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicy string

const (
	ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicyOff                      ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "off"
	ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicyGeo                      ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "geo"
	ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicyRandom                   ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "random"
	ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicyDynamicLatency           ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicyProximity                ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "proximity"
	ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicyLeastConnections         ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "least_connections"
	ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicyEmpty                    ZoneLoadBalancerGetResponseResultRulesOverridesSteeringPolicy = "\"\""
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
type ZoneLoadBalancerGetResponseResultSessionAffinity string

const (
	ZoneLoadBalancerGetResponseResultSessionAffinityNone     ZoneLoadBalancerGetResponseResultSessionAffinity = "none"
	ZoneLoadBalancerGetResponseResultSessionAffinityCookie   ZoneLoadBalancerGetResponseResultSessionAffinity = "cookie"
	ZoneLoadBalancerGetResponseResultSessionAffinityIPCookie ZoneLoadBalancerGetResponseResultSessionAffinity = "ip_cookie"
	ZoneLoadBalancerGetResponseResultSessionAffinityHeader   ZoneLoadBalancerGetResponseResultSessionAffinity = "header"
	ZoneLoadBalancerGetResponseResultSessionAffinityEmpty    ZoneLoadBalancerGetResponseResultSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerGetResponseResultSessionAffinityAttributes struct {
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
	Samesite ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover ZoneLoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 zoneLoadBalancerGetResponseResultSessionAffinityAttributesJSON                 `json:"-"`
}

// zoneLoadBalancerGetResponseResultSessionAffinityAttributesJSON contains the JSON
// metadata for the struct
// [ZoneLoadBalancerGetResponseResultSessionAffinityAttributes]
type zoneLoadBalancerGetResponseResultSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneLoadBalancerGetResponseResultSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSecureAuto   ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSecureAlways ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSecureNever  ZoneLoadBalancerGetResponseResultSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerGetResponseResultSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerGetResponseResultSteeringPolicy string

const (
	ZoneLoadBalancerGetResponseResultSteeringPolicyOff                      ZoneLoadBalancerGetResponseResultSteeringPolicy = "off"
	ZoneLoadBalancerGetResponseResultSteeringPolicyGeo                      ZoneLoadBalancerGetResponseResultSteeringPolicy = "geo"
	ZoneLoadBalancerGetResponseResultSteeringPolicyRandom                   ZoneLoadBalancerGetResponseResultSteeringPolicy = "random"
	ZoneLoadBalancerGetResponseResultSteeringPolicyDynamicLatency           ZoneLoadBalancerGetResponseResultSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerGetResponseResultSteeringPolicyProximity                ZoneLoadBalancerGetResponseResultSteeringPolicy = "proximity"
	ZoneLoadBalancerGetResponseResultSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerGetResponseResultSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerGetResponseResultSteeringPolicyLeastConnections         ZoneLoadBalancerGetResponseResultSteeringPolicy = "least_connections"
	ZoneLoadBalancerGetResponseResultSteeringPolicyEmpty                    ZoneLoadBalancerGetResponseResultSteeringPolicy = "\"\""
)

// Whether the API call was successful
type ZoneLoadBalancerGetResponseSuccess bool

const (
	ZoneLoadBalancerGetResponseSuccessTrue ZoneLoadBalancerGetResponseSuccess = true
)

type ZoneLoadBalancerUpdateResponse struct {
	Errors   []ZoneLoadBalancerUpdateResponseError   `json:"errors"`
	Messages []ZoneLoadBalancerUpdateResponseMessage `json:"messages"`
	Result   ZoneLoadBalancerUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneLoadBalancerUpdateResponseSuccess `json:"success"`
	JSON    zoneLoadBalancerUpdateResponseJSON    `json:"-"`
}

// zoneLoadBalancerUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerUpdateResponse]
type zoneLoadBalancerUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneLoadBalancerUpdateResponseErrorJSON `json:"-"`
}

// zoneLoadBalancerUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerUpdateResponseError]
type zoneLoadBalancerUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneLoadBalancerUpdateResponseMessageJSON `json:"-"`
}

// zoneLoadBalancerUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerUpdateResponseMessage]
type zoneLoadBalancerUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerUpdateResponseResult struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting ZoneLoadBalancerUpdateResponseResultAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy ZoneLoadBalancerUpdateResponseResultLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                                            `json:"modified_on" format:"date-time"`
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
	RandomSteering ZoneLoadBalancerUpdateResponseResultRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []ZoneLoadBalancerUpdateResponseResultRule `json:"rules"`
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
	SessionAffinity ZoneLoadBalancerUpdateResponseResultSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy ZoneLoadBalancerUpdateResponseResultSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                  `json:"ttl"`
	JSON zoneLoadBalancerUpdateResponseResultJSON `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerUpdateResponseResult]
type zoneLoadBalancerUpdateResponseResultJSON struct {
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

func (r *ZoneLoadBalancerUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerUpdateResponseResultAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                    `json:"failover_across_pools"`
	JSON                zoneLoadBalancerUpdateResponseResultAdaptiveRoutingJSON `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultAdaptiveRoutingJSON contains the JSON
// metadata for the struct [ZoneLoadBalancerUpdateResponseResultAdaptiveRouting]
type zoneLoadBalancerUpdateResponseResultAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseResultAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerUpdateResponseResultLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode ZoneLoadBalancerUpdateResponseResultLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs ZoneLoadBalancerUpdateResponseResultLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      zoneLoadBalancerUpdateResponseResultLocationStrategyJSON      `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultLocationStrategyJSON contains the JSON
// metadata for the struct [ZoneLoadBalancerUpdateResponseResultLocationStrategy]
type zoneLoadBalancerUpdateResponseResultLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseResultLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerUpdateResponseResultLocationStrategyMode string

const (
	ZoneLoadBalancerUpdateResponseResultLocationStrategyModePop        ZoneLoadBalancerUpdateResponseResultLocationStrategyMode = "pop"
	ZoneLoadBalancerUpdateResponseResultLocationStrategyModeResolverIP ZoneLoadBalancerUpdateResponseResultLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerUpdateResponseResultLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerUpdateResponseResultLocationStrategyPreferEcsAlways    ZoneLoadBalancerUpdateResponseResultLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerUpdateResponseResultLocationStrategyPreferEcsNever     ZoneLoadBalancerUpdateResponseResultLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerUpdateResponseResultLocationStrategyPreferEcsProximity ZoneLoadBalancerUpdateResponseResultLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerUpdateResponseResultLocationStrategyPreferEcsGeo       ZoneLoadBalancerUpdateResponseResultLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerUpdateResponseResultRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                            `json:"pool_weights"`
	JSON        zoneLoadBalancerUpdateResponseResultRandomSteeringJSON `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultRandomSteeringJSON contains the JSON
// metadata for the struct [ZoneLoadBalancerUpdateResponseResultRandomSteering]
type zoneLoadBalancerUpdateResponseResultRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseResultRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type ZoneLoadBalancerUpdateResponseResultRule struct {
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
	FixedResponse ZoneLoadBalancerUpdateResponseResultRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides ZoneLoadBalancerUpdateResponseResultRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                                         `json:"terminates"`
	JSON       zoneLoadBalancerUpdateResponseResultRuleJSON `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultRuleJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerUpdateResponseResultRule]
type zoneLoadBalancerUpdateResponseResultRuleJSON struct {
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

func (r *ZoneLoadBalancerUpdateResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type ZoneLoadBalancerUpdateResponseResultRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                                      `json:"status_code"`
	JSON       zoneLoadBalancerUpdateResponseResultRulesFixedResponseJSON `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultRulesFixedResponseJSON contains the JSON
// metadata for the struct [ZoneLoadBalancerUpdateResponseResultRulesFixedResponse]
type zoneLoadBalancerUpdateResponseResultRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseResultRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type ZoneLoadBalancerUpdateResponseResultRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting ZoneLoadBalancerUpdateResponseResultRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering ZoneLoadBalancerUpdateResponseResultRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                                `json:"ttl"`
	JSON zoneLoadBalancerUpdateResponseResultRulesOverridesJSON `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultRulesOverridesJSON contains the JSON
// metadata for the struct [ZoneLoadBalancerUpdateResponseResultRulesOverrides]
type zoneLoadBalancerUpdateResponseResultRulesOverridesJSON struct {
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

func (r *ZoneLoadBalancerUpdateResponseResultRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerUpdateResponseResultRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                                  `json:"failover_across_pools"`
	JSON                zoneLoadBalancerUpdateResponseResultRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultRulesOverridesAdaptiveRoutingJSON contains
// the JSON metadata for the struct
// [ZoneLoadBalancerUpdateResponseResultRulesOverridesAdaptiveRouting]
type zoneLoadBalancerUpdateResponseResultRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseResultRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      zoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyJSON      `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyJSON contains
// the JSON metadata for the struct
// [ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategy]
type zoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyMode string

const (
	ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyModePop        ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyMode = "pop"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyModeResolverIP ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcsAlways    ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcsNever     ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcsProximity ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcsGeo       ZoneLoadBalancerUpdateResponseResultRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerUpdateResponseResultRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                                          `json:"pool_weights"`
	JSON        zoneLoadBalancerUpdateResponseResultRulesOverridesRandomSteeringJSON `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultRulesOverridesRandomSteeringJSON contains
// the JSON metadata for the struct
// [ZoneLoadBalancerUpdateResponseResultRulesOverridesRandomSteering]
type zoneLoadBalancerUpdateResponseResultRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseResultRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
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
type ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinity string

const (
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityNone     ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinity = "none"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityCookie   ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinity = "cookie"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityIPCookie ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinity = "ip_cookie"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityHeader   ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinity = "header"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityEmpty    ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributes struct {
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
	Samesite ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 zoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesJSON
// contains the JSON metadata for the struct
// [ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributes]
type zoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecureAuto   ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecureAlways ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecureNever  ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerUpdateResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy string

const (
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyOff                      ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "off"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyGeo                      ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "geo"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyRandom                   ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "random"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyDynamicLatency           ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyProximity                ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "proximity"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyLeastConnections         ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "least_connections"
	ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicyEmpty                    ZoneLoadBalancerUpdateResponseResultRulesOverridesSteeringPolicy = "\"\""
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
type ZoneLoadBalancerUpdateResponseResultSessionAffinity string

const (
	ZoneLoadBalancerUpdateResponseResultSessionAffinityNone     ZoneLoadBalancerUpdateResponseResultSessionAffinity = "none"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityCookie   ZoneLoadBalancerUpdateResponseResultSessionAffinity = "cookie"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityIPCookie ZoneLoadBalancerUpdateResponseResultSessionAffinity = "ip_cookie"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityHeader   ZoneLoadBalancerUpdateResponseResultSessionAffinity = "header"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityEmpty    ZoneLoadBalancerUpdateResponseResultSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributes struct {
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
	Samesite ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 zoneLoadBalancerUpdateResponseResultSessionAffinityAttributesJSON                 `json:"-"`
}

// zoneLoadBalancerUpdateResponseResultSessionAffinityAttributesJSON contains the
// JSON metadata for the struct
// [ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributes]
type zoneLoadBalancerUpdateResponseResultSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSecureAuto   ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSecureAlways ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSecureNever  ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerUpdateResponseResultSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerUpdateResponseResultSteeringPolicy string

const (
	ZoneLoadBalancerUpdateResponseResultSteeringPolicyOff                      ZoneLoadBalancerUpdateResponseResultSteeringPolicy = "off"
	ZoneLoadBalancerUpdateResponseResultSteeringPolicyGeo                      ZoneLoadBalancerUpdateResponseResultSteeringPolicy = "geo"
	ZoneLoadBalancerUpdateResponseResultSteeringPolicyRandom                   ZoneLoadBalancerUpdateResponseResultSteeringPolicy = "random"
	ZoneLoadBalancerUpdateResponseResultSteeringPolicyDynamicLatency           ZoneLoadBalancerUpdateResponseResultSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerUpdateResponseResultSteeringPolicyProximity                ZoneLoadBalancerUpdateResponseResultSteeringPolicy = "proximity"
	ZoneLoadBalancerUpdateResponseResultSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerUpdateResponseResultSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerUpdateResponseResultSteeringPolicyLeastConnections         ZoneLoadBalancerUpdateResponseResultSteeringPolicy = "least_connections"
	ZoneLoadBalancerUpdateResponseResultSteeringPolicyEmpty                    ZoneLoadBalancerUpdateResponseResultSteeringPolicy = "\"\""
)

// Whether the API call was successful
type ZoneLoadBalancerUpdateResponseSuccess bool

const (
	ZoneLoadBalancerUpdateResponseSuccessTrue ZoneLoadBalancerUpdateResponseSuccess = true
)

type ZoneLoadBalancerListResponse struct {
	Errors     []ZoneLoadBalancerListResponseError    `json:"errors"`
	Messages   []ZoneLoadBalancerListResponseMessage  `json:"messages"`
	Result     []ZoneLoadBalancerListResponseResult   `json:"result"`
	ResultInfo ZoneLoadBalancerListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneLoadBalancerListResponseSuccess `json:"success"`
	JSON    zoneLoadBalancerListResponseJSON    `json:"-"`
}

// zoneLoadBalancerListResponseJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerListResponse]
type zoneLoadBalancerListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneLoadBalancerListResponseErrorJSON `json:"-"`
}

// zoneLoadBalancerListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerListResponseError]
type zoneLoadBalancerListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneLoadBalancerListResponseMessageJSON `json:"-"`
}

// zoneLoadBalancerListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerListResponseMessage]
type zoneLoadBalancerListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerListResponseResult struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting ZoneLoadBalancerListResponseResultAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy ZoneLoadBalancerListResponseResultLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                                          `json:"modified_on" format:"date-time"`
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
	RandomSteering ZoneLoadBalancerListResponseResultRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []ZoneLoadBalancerListResponseResultRule `json:"rules"`
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
	SessionAffinity ZoneLoadBalancerListResponseResultSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes ZoneLoadBalancerListResponseResultSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy ZoneLoadBalancerListResponseResultSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                `json:"ttl"`
	JSON zoneLoadBalancerListResponseResultJSON `json:"-"`
}

// zoneLoadBalancerListResponseResultJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerListResponseResult]
type zoneLoadBalancerListResponseResultJSON struct {
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

func (r *ZoneLoadBalancerListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerListResponseResultAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                  `json:"failover_across_pools"`
	JSON                zoneLoadBalancerListResponseResultAdaptiveRoutingJSON `json:"-"`
}

// zoneLoadBalancerListResponseResultAdaptiveRoutingJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerListResponseResultAdaptiveRouting]
type zoneLoadBalancerListResponseResultAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseResultAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerListResponseResultLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode ZoneLoadBalancerListResponseResultLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs ZoneLoadBalancerListResponseResultLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      zoneLoadBalancerListResponseResultLocationStrategyJSON      `json:"-"`
}

// zoneLoadBalancerListResponseResultLocationStrategyJSON contains the JSON
// metadata for the struct [ZoneLoadBalancerListResponseResultLocationStrategy]
type zoneLoadBalancerListResponseResultLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseResultLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerListResponseResultLocationStrategyMode string

const (
	ZoneLoadBalancerListResponseResultLocationStrategyModePop        ZoneLoadBalancerListResponseResultLocationStrategyMode = "pop"
	ZoneLoadBalancerListResponseResultLocationStrategyModeResolverIP ZoneLoadBalancerListResponseResultLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerListResponseResultLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerListResponseResultLocationStrategyPreferEcsAlways    ZoneLoadBalancerListResponseResultLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerListResponseResultLocationStrategyPreferEcsNever     ZoneLoadBalancerListResponseResultLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerListResponseResultLocationStrategyPreferEcsProximity ZoneLoadBalancerListResponseResultLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerListResponseResultLocationStrategyPreferEcsGeo       ZoneLoadBalancerListResponseResultLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerListResponseResultRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                          `json:"pool_weights"`
	JSON        zoneLoadBalancerListResponseResultRandomSteeringJSON `json:"-"`
}

// zoneLoadBalancerListResponseResultRandomSteeringJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerListResponseResultRandomSteering]
type zoneLoadBalancerListResponseResultRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseResultRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type ZoneLoadBalancerListResponseResultRule struct {
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
	FixedResponse ZoneLoadBalancerListResponseResultRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides ZoneLoadBalancerListResponseResultRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                                       `json:"terminates"`
	JSON       zoneLoadBalancerListResponseResultRuleJSON `json:"-"`
}

// zoneLoadBalancerListResponseResultRuleJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerListResponseResultRule]
type zoneLoadBalancerListResponseResultRuleJSON struct {
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

func (r *ZoneLoadBalancerListResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type ZoneLoadBalancerListResponseResultRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                                    `json:"status_code"`
	JSON       zoneLoadBalancerListResponseResultRulesFixedResponseJSON `json:"-"`
}

// zoneLoadBalancerListResponseResultRulesFixedResponseJSON contains the JSON
// metadata for the struct [ZoneLoadBalancerListResponseResultRulesFixedResponse]
type zoneLoadBalancerListResponseResultRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseResultRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type ZoneLoadBalancerListResponseResultRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting ZoneLoadBalancerListResponseResultRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategy `json:"location_strategy"`
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
	RandomSteering ZoneLoadBalancerListResponseResultRulesOverridesRandomSteering `json:"random_steering"`
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
	SessionAffinity ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
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
	SteeringPolicy ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64                                              `json:"ttl"`
	JSON zoneLoadBalancerListResponseResultRulesOverridesJSON `json:"-"`
}

// zoneLoadBalancerListResponseResultRulesOverridesJSON contains the JSON metadata
// for the struct [ZoneLoadBalancerListResponseResultRulesOverrides]
type zoneLoadBalancerListResponseResultRulesOverridesJSON struct {
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

func (r *ZoneLoadBalancerListResponseResultRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerListResponseResultRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                                `json:"failover_across_pools"`
	JSON                zoneLoadBalancerListResponseResultRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// zoneLoadBalancerListResponseResultRulesOverridesAdaptiveRoutingJSON contains the
// JSON metadata for the struct
// [ZoneLoadBalancerListResponseResultRulesOverridesAdaptiveRouting]
type zoneLoadBalancerListResponseResultRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseResultRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      zoneLoadBalancerListResponseResultRulesOverridesLocationStrategyJSON      `json:"-"`
}

// zoneLoadBalancerListResponseResultRulesOverridesLocationStrategyJSON contains
// the JSON metadata for the struct
// [ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategy]
type zoneLoadBalancerListResponseResultRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyMode string

const (
	ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyModePop        ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyMode = "pop"
	ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyModeResolverIP ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcsAlways    ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcsNever     ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcsProximity ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcsGeo       ZoneLoadBalancerListResponseResultRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerListResponseResultRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                                        `json:"pool_weights"`
	JSON        zoneLoadBalancerListResponseResultRulesOverridesRandomSteeringJSON `json:"-"`
}

// zoneLoadBalancerListResponseResultRulesOverridesRandomSteeringJSON contains the
// JSON metadata for the struct
// [ZoneLoadBalancerListResponseResultRulesOverridesRandomSteering]
type zoneLoadBalancerListResponseResultRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseResultRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
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
type ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinity string

const (
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityNone     ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinity = "none"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityCookie   ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinity = "cookie"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityIPCookie ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinity = "ip_cookie"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityHeader   ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinity = "header"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityEmpty    ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributes struct {
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
	Samesite ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 zoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// zoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesJSON
// contains the JSON metadata for the struct
// [ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributes]
type zoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecureAuto   ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecureAlways ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecureNever  ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerListResponseResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicy string

const (
	ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicyOff                      ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicy = "off"
	ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicyGeo                      ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicy = "geo"
	ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicyRandom                   ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicy = "random"
	ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicyDynamicLatency           ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicyProximity                ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicy = "proximity"
	ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicyLeastConnections         ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicy = "least_connections"
	ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicyEmpty                    ZoneLoadBalancerListResponseResultRulesOverridesSteeringPolicy = "\"\""
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
type ZoneLoadBalancerListResponseResultSessionAffinity string

const (
	ZoneLoadBalancerListResponseResultSessionAffinityNone     ZoneLoadBalancerListResponseResultSessionAffinity = "none"
	ZoneLoadBalancerListResponseResultSessionAffinityCookie   ZoneLoadBalancerListResponseResultSessionAffinity = "cookie"
	ZoneLoadBalancerListResponseResultSessionAffinityIPCookie ZoneLoadBalancerListResponseResultSessionAffinity = "ip_cookie"
	ZoneLoadBalancerListResponseResultSessionAffinityHeader   ZoneLoadBalancerListResponseResultSessionAffinity = "header"
	ZoneLoadBalancerListResponseResultSessionAffinityEmpty    ZoneLoadBalancerListResponseResultSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerListResponseResultSessionAffinityAttributes struct {
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
	Samesite ZoneLoadBalancerListResponseResultSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure ZoneLoadBalancerListResponseResultSessionAffinityAttributesSecure `json:"secure"`
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
	ZeroDowntimeFailover ZoneLoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 zoneLoadBalancerListResponseResultSessionAffinityAttributesJSON                 `json:"-"`
}

// zoneLoadBalancerListResponseResultSessionAffinityAttributesJSON contains the
// JSON metadata for the struct
// [ZoneLoadBalancerListResponseResultSessionAffinityAttributes]
type zoneLoadBalancerListResponseResultSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseResultSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerListResponseResultSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerListResponseResultSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerListResponseResultSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerListResponseResultSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerListResponseResultSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerListResponseResultSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerListResponseResultSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerListResponseResultSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerListResponseResultSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerListResponseResultSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerListResponseResultSessionAffinityAttributesSecureAuto   ZoneLoadBalancerListResponseResultSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerListResponseResultSessionAffinityAttributesSecureAlways ZoneLoadBalancerListResponseResultSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerListResponseResultSessionAffinityAttributesSecureNever  ZoneLoadBalancerListResponseResultSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerListResponseResultSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerListResponseResultSteeringPolicy string

const (
	ZoneLoadBalancerListResponseResultSteeringPolicyOff                      ZoneLoadBalancerListResponseResultSteeringPolicy = "off"
	ZoneLoadBalancerListResponseResultSteeringPolicyGeo                      ZoneLoadBalancerListResponseResultSteeringPolicy = "geo"
	ZoneLoadBalancerListResponseResultSteeringPolicyRandom                   ZoneLoadBalancerListResponseResultSteeringPolicy = "random"
	ZoneLoadBalancerListResponseResultSteeringPolicyDynamicLatency           ZoneLoadBalancerListResponseResultSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerListResponseResultSteeringPolicyProximity                ZoneLoadBalancerListResponseResultSteeringPolicy = "proximity"
	ZoneLoadBalancerListResponseResultSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerListResponseResultSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerListResponseResultSteeringPolicyLeastConnections         ZoneLoadBalancerListResponseResultSteeringPolicy = "least_connections"
	ZoneLoadBalancerListResponseResultSteeringPolicyEmpty                    ZoneLoadBalancerListResponseResultSteeringPolicy = "\"\""
)

type ZoneLoadBalancerListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       zoneLoadBalancerListResponseResultInfoJSON `json:"-"`
}

// zoneLoadBalancerListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerListResponseResultInfo]
type zoneLoadBalancerListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLoadBalancerListResponseSuccess bool

const (
	ZoneLoadBalancerListResponseSuccessTrue ZoneLoadBalancerListResponseSuccess = true
)

type ZoneLoadBalancerDeleteResponse struct {
	Errors   []ZoneLoadBalancerDeleteResponseError   `json:"errors"`
	Messages []ZoneLoadBalancerDeleteResponseMessage `json:"messages"`
	Result   ZoneLoadBalancerDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneLoadBalancerDeleteResponseSuccess `json:"success"`
	JSON    zoneLoadBalancerDeleteResponseJSON    `json:"-"`
}

// zoneLoadBalancerDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerDeleteResponse]
type zoneLoadBalancerDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneLoadBalancerDeleteResponseErrorJSON `json:"-"`
}

// zoneLoadBalancerDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerDeleteResponseError]
type zoneLoadBalancerDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneLoadBalancerDeleteResponseMessageJSON `json:"-"`
}

// zoneLoadBalancerDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerDeleteResponseMessage]
type zoneLoadBalancerDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneLoadBalancerDeleteResponseResult struct {
	ID   string                                   `json:"id"`
	JSON zoneLoadBalancerDeleteResponseResultJSON `json:"-"`
}

// zoneLoadBalancerDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerDeleteResponseResult]
type zoneLoadBalancerDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneLoadBalancerDeleteResponseSuccess bool

const (
	ZoneLoadBalancerDeleteResponseSuccessTrue ZoneLoadBalancerDeleteResponseSuccess = true
)

type ZoneLoadBalancerNewParams struct {
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
	AdaptiveRouting param.Field[ZoneLoadBalancerNewParamsAdaptiveRouting] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[interface{}] `json:"country_pools"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[ZoneLoadBalancerNewParamsLocationStrategy] `json:"location_strategy"`
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
	RandomSteering param.Field[ZoneLoadBalancerNewParamsRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]ZoneLoadBalancerNewParamsRule] `json:"rules"`
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
	SessionAffinity param.Field[ZoneLoadBalancerNewParamsSessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[ZoneLoadBalancerNewParamsSessionAffinityAttributes] `json:"session_affinity_attributes"`
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
	SteeringPolicy param.Field[ZoneLoadBalancerNewParamsSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r ZoneLoadBalancerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerNewParamsAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r ZoneLoadBalancerNewParamsAdaptiveRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerNewParamsLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[ZoneLoadBalancerNewParamsLocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[ZoneLoadBalancerNewParamsLocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r ZoneLoadBalancerNewParamsLocationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerNewParamsLocationStrategyMode string

const (
	ZoneLoadBalancerNewParamsLocationStrategyModePop        ZoneLoadBalancerNewParamsLocationStrategyMode = "pop"
	ZoneLoadBalancerNewParamsLocationStrategyModeResolverIP ZoneLoadBalancerNewParamsLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerNewParamsLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerNewParamsLocationStrategyPreferEcsAlways    ZoneLoadBalancerNewParamsLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerNewParamsLocationStrategyPreferEcsNever     ZoneLoadBalancerNewParamsLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerNewParamsLocationStrategyPreferEcsProximity ZoneLoadBalancerNewParamsLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerNewParamsLocationStrategyPreferEcsGeo       ZoneLoadBalancerNewParamsLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerNewParamsRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[interface{}] `json:"pool_weights"`
}

func (r ZoneLoadBalancerNewParamsRandomSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type ZoneLoadBalancerNewParamsRule struct {
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
	FixedResponse param.Field[ZoneLoadBalancerNewParamsRulesFixedResponse] `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides param.Field[ZoneLoadBalancerNewParamsRulesOverrides] `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority param.Field[int64] `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates param.Field[bool] `json:"terminates"`
}

func (r ZoneLoadBalancerNewParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type ZoneLoadBalancerNewParamsRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType param.Field[string] `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location param.Field[string] `json:"location"`
	// Text to include as the http body.
	MessageBody param.Field[string] `json:"message_body"`
	// The http status code to respond with.
	StatusCode param.Field[int64] `json:"status_code"`
}

func (r ZoneLoadBalancerNewParamsRulesFixedResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type ZoneLoadBalancerNewParamsRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[ZoneLoadBalancerNewParamsRulesOverridesAdaptiveRouting] `json:"adaptive_routing"`
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
	LocationStrategy param.Field[ZoneLoadBalancerNewParamsRulesOverridesLocationStrategy] `json:"location_strategy"`
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
	RandomSteering param.Field[ZoneLoadBalancerNewParamsRulesOverridesRandomSteering] `json:"random_steering"`
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
	SessionAffinity param.Field[ZoneLoadBalancerNewParamsRulesOverridesSessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributes] `json:"session_affinity_attributes"`
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
	SteeringPolicy param.Field[ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r ZoneLoadBalancerNewParamsRulesOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerNewParamsRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r ZoneLoadBalancerNewParamsRulesOverridesAdaptiveRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerNewParamsRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r ZoneLoadBalancerNewParamsRulesOverridesLocationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyMode string

const (
	ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyModePop        ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyMode = "pop"
	ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyModeResolverIP ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsAlways    ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsNever     ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsProximity ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcsGeo       ZoneLoadBalancerNewParamsRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerNewParamsRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[interface{}] `json:"pool_weights"`
}

func (r ZoneLoadBalancerNewParamsRulesOverridesRandomSteering) MarshalJSON() (data []byte, err error) {
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
type ZoneLoadBalancerNewParamsRulesOverridesSessionAffinity string

const (
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityNone     ZoneLoadBalancerNewParamsRulesOverridesSessionAffinity = "none"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityCookie   ZoneLoadBalancerNewParamsRulesOverridesSessionAffinity = "cookie"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityIPCookie ZoneLoadBalancerNewParamsRulesOverridesSessionAffinity = "ip_cookie"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityHeader   ZoneLoadBalancerNewParamsRulesOverridesSessionAffinity = "header"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityEmpty    ZoneLoadBalancerNewParamsRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributes struct {
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
	Samesite param.Field[ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure] `json:"secure"`
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
	ZeroDowntimeFailover param.Field[ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAuto   ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureAlways ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecureNever  ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy string

const (
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyOff                      ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "off"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyGeo                      ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "geo"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyRandom                   ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "random"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency           ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyProximity                ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "proximity"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyLeastConnections         ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "least_connections"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyEmpty                    ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "\"\""
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
type ZoneLoadBalancerNewParamsSessionAffinity string

const (
	ZoneLoadBalancerNewParamsSessionAffinityNone     ZoneLoadBalancerNewParamsSessionAffinity = "none"
	ZoneLoadBalancerNewParamsSessionAffinityCookie   ZoneLoadBalancerNewParamsSessionAffinity = "cookie"
	ZoneLoadBalancerNewParamsSessionAffinityIPCookie ZoneLoadBalancerNewParamsSessionAffinity = "ip_cookie"
	ZoneLoadBalancerNewParamsSessionAffinityHeader   ZoneLoadBalancerNewParamsSessionAffinity = "header"
	ZoneLoadBalancerNewParamsSessionAffinityEmpty    ZoneLoadBalancerNewParamsSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerNewParamsSessionAffinityAttributes struct {
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
	Samesite param.Field[ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[ZoneLoadBalancerNewParamsSessionAffinityAttributesSecure] `json:"secure"`
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
	ZeroDowntimeFailover param.Field[ZoneLoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r ZoneLoadBalancerNewParamsSessionAffinityAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerNewParamsSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerNewParamsSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerNewParamsSessionAffinityAttributesSecureAuto   ZoneLoadBalancerNewParamsSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerNewParamsSessionAffinityAttributesSecureAlways ZoneLoadBalancerNewParamsSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerNewParamsSessionAffinityAttributesSecureNever  ZoneLoadBalancerNewParamsSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerNewParamsSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerNewParamsSteeringPolicy string

const (
	ZoneLoadBalancerNewParamsSteeringPolicyOff                      ZoneLoadBalancerNewParamsSteeringPolicy = "off"
	ZoneLoadBalancerNewParamsSteeringPolicyGeo                      ZoneLoadBalancerNewParamsSteeringPolicy = "geo"
	ZoneLoadBalancerNewParamsSteeringPolicyRandom                   ZoneLoadBalancerNewParamsSteeringPolicy = "random"
	ZoneLoadBalancerNewParamsSteeringPolicyDynamicLatency           ZoneLoadBalancerNewParamsSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerNewParamsSteeringPolicyProximity                ZoneLoadBalancerNewParamsSteeringPolicy = "proximity"
	ZoneLoadBalancerNewParamsSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerNewParamsSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerNewParamsSteeringPolicyLeastConnections         ZoneLoadBalancerNewParamsSteeringPolicy = "least_connections"
	ZoneLoadBalancerNewParamsSteeringPolicyEmpty                    ZoneLoadBalancerNewParamsSteeringPolicy = "\"\""
)

type ZoneLoadBalancerUpdateParams struct {
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
	AdaptiveRouting param.Field[ZoneLoadBalancerUpdateParamsAdaptiveRouting] `json:"adaptive_routing"`
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
	LocationStrategy param.Field[ZoneLoadBalancerUpdateParamsLocationStrategy] `json:"location_strategy"`
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
	RandomSteering param.Field[ZoneLoadBalancerUpdateParamsRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]ZoneLoadBalancerUpdateParamsRule] `json:"rules"`
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
	SessionAffinity param.Field[ZoneLoadBalancerUpdateParamsSessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[ZoneLoadBalancerUpdateParamsSessionAffinityAttributes] `json:"session_affinity_attributes"`
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
	SteeringPolicy param.Field[ZoneLoadBalancerUpdateParamsSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r ZoneLoadBalancerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerUpdateParamsAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r ZoneLoadBalancerUpdateParamsAdaptiveRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerUpdateParamsLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[ZoneLoadBalancerUpdateParamsLocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r ZoneLoadBalancerUpdateParamsLocationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerUpdateParamsLocationStrategyMode string

const (
	ZoneLoadBalancerUpdateParamsLocationStrategyModePop        ZoneLoadBalancerUpdateParamsLocationStrategyMode = "pop"
	ZoneLoadBalancerUpdateParamsLocationStrategyModeResolverIP ZoneLoadBalancerUpdateParamsLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcsAlways    ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcsNever     ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcsProximity ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcsGeo       ZoneLoadBalancerUpdateParamsLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerUpdateParamsRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[interface{}] `json:"pool_weights"`
}

func (r ZoneLoadBalancerUpdateParamsRandomSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type ZoneLoadBalancerUpdateParamsRule struct {
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
	FixedResponse param.Field[ZoneLoadBalancerUpdateParamsRulesFixedResponse] `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides param.Field[ZoneLoadBalancerUpdateParamsRulesOverrides] `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority param.Field[int64] `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates param.Field[bool] `json:"terminates"`
}

func (r ZoneLoadBalancerUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type ZoneLoadBalancerUpdateParamsRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType param.Field[string] `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location param.Field[string] `json:"location"`
	// Text to include as the http body.
	MessageBody param.Field[string] `json:"message_body"`
	// The http status code to respond with.
	StatusCode param.Field[int64] `json:"status_code"`
}

func (r ZoneLoadBalancerUpdateParamsRulesFixedResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type ZoneLoadBalancerUpdateParamsRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesAdaptiveRouting] `json:"adaptive_routing"`
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
	LocationStrategy param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategy] `json:"location_strategy"`
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
	RandomSteering param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesRandomSteering] `json:"random_steering"`
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
	SessionAffinity param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes] `json:"session_affinity_attributes"`
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
	SteeringPolicy param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r ZoneLoadBalancerUpdateParamsRulesOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type ZoneLoadBalancerUpdateParamsRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r ZoneLoadBalancerUpdateParamsRulesOverridesAdaptiveRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyMode string

const (
	ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyModePop        ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyMode = "pop"
	ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyModeResolverIP ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs string

const (
	ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsAlways    ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs = "always"
	ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsNever     ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs = "never"
	ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsProximity ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs = "proximity"
	ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcsGeo       ZoneLoadBalancerUpdateParamsRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type ZoneLoadBalancerUpdateParamsRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[interface{}] `json:"pool_weights"`
}

func (r ZoneLoadBalancerUpdateParamsRulesOverridesRandomSteering) MarshalJSON() (data []byte, err error) {
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
type ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinity string

const (
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityNone     ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinity = "none"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityCookie   ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinity = "cookie"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityIPCookie ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinity = "ip_cookie"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityHeader   ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinity = "header"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityEmpty    ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes struct {
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
	Samesite param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure] `json:"secure"`
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
	ZeroDowntimeFailover param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAuto   ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureAlways ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecureNever  ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy string

const (
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyOff                      ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "off"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyGeo                      ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "geo"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyRandom                   ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "random"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency           ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyProximity                ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "proximity"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyLeastConnections         ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "least_connections"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyEmpty                    ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "\"\""
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
type ZoneLoadBalancerUpdateParamsSessionAffinity string

const (
	ZoneLoadBalancerUpdateParamsSessionAffinityNone     ZoneLoadBalancerUpdateParamsSessionAffinity = "none"
	ZoneLoadBalancerUpdateParamsSessionAffinityCookie   ZoneLoadBalancerUpdateParamsSessionAffinity = "cookie"
	ZoneLoadBalancerUpdateParamsSessionAffinityIPCookie ZoneLoadBalancerUpdateParamsSessionAffinity = "ip_cookie"
	ZoneLoadBalancerUpdateParamsSessionAffinityHeader   ZoneLoadBalancerUpdateParamsSessionAffinity = "header"
	ZoneLoadBalancerUpdateParamsSessionAffinityEmpty    ZoneLoadBalancerUpdateParamsSessionAffinity = "\"\""
)

// Configures attributes for session affinity.
type ZoneLoadBalancerUpdateParamsSessionAffinityAttributes struct {
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
	Samesite param.Field[ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSecure] `json:"secure"`
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
	ZeroDowntimeFailover param.Field[ZoneLoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r ZoneLoadBalancerUpdateParamsSessionAffinityAttributes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesite string

const (
	ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesiteAuto   ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesite = "Auto"
	ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesiteLax    ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesite = "Lax"
	ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesiteNone   ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesite = "None"
	ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesiteStrict ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSecure string

const (
	ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSecureAuto   ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSecure = "Auto"
	ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSecureAlways ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSecure = "Always"
	ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSecureNever  ZoneLoadBalancerUpdateParamsSessionAffinityAttributesSecure = "Never"
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
type ZoneLoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover string

const (
	ZoneLoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverNone      ZoneLoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover = "none"
	ZoneLoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverTemporary ZoneLoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	ZoneLoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailoverSticky    ZoneLoadBalancerUpdateParamsSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
type ZoneLoadBalancerUpdateParamsSteeringPolicy string

const (
	ZoneLoadBalancerUpdateParamsSteeringPolicyOff                      ZoneLoadBalancerUpdateParamsSteeringPolicy = "off"
	ZoneLoadBalancerUpdateParamsSteeringPolicyGeo                      ZoneLoadBalancerUpdateParamsSteeringPolicy = "geo"
	ZoneLoadBalancerUpdateParamsSteeringPolicyRandom                   ZoneLoadBalancerUpdateParamsSteeringPolicy = "random"
	ZoneLoadBalancerUpdateParamsSteeringPolicyDynamicLatency           ZoneLoadBalancerUpdateParamsSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerUpdateParamsSteeringPolicyProximity                ZoneLoadBalancerUpdateParamsSteeringPolicy = "proximity"
	ZoneLoadBalancerUpdateParamsSteeringPolicyLeastOutstandingRequests ZoneLoadBalancerUpdateParamsSteeringPolicy = "least_outstanding_requests"
	ZoneLoadBalancerUpdateParamsSteeringPolicyLeastConnections         ZoneLoadBalancerUpdateParamsSteeringPolicy = "least_connections"
	ZoneLoadBalancerUpdateParamsSteeringPolicyEmpty                    ZoneLoadBalancerUpdateParamsSteeringPolicy = "\"\""
)
