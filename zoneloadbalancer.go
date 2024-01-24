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
func (r *ZoneLoadBalancerService) New(ctx context.Context, identifier interface{}, body ZoneLoadBalancerNewParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/load_balancers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single configured load balancer.
func (r *ZoneLoadBalancerService) Get(ctx context.Context, identifier1 interface{}, identifier interface{}, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/load_balancers/%v", identifier1, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured load balancer.
func (r *ZoneLoadBalancerService) Update(ctx context.Context, identifier1 interface{}, identifier interface{}, body ZoneLoadBalancerUpdateParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/load_balancers/%v", identifier1, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List configured load balancers.
func (r *ZoneLoadBalancerService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *LoadBalancerCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/load_balancers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a configured load balancer.
func (r *ZoneLoadBalancerService) Delete(ctx context.Context, identifier1 interface{}, identifier interface{}, opts ...option.RequestOption) (res *IDResponseZtdTboH6, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/load_balancers/%v", identifier1, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type IDResponseZtdTboH6 struct {
	Errors   []IDResponseZtdTboH6Error   `json:"errors"`
	Messages []IDResponseZtdTboH6Message `json:"messages"`
	Result   IDResponseZtdTboH6Result    `json:"result"`
	// Whether the API call was successful
	Success IDResponseZtdTboH6Success `json:"success"`
	JSON    idResponseZtdTboH6JSON    `json:"-"`
}

// idResponseZtdTboH6JSON contains the JSON metadata for the struct
// [IDResponseZtdTboH6]
type idResponseZtdTboH6JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseZtdTboH6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponseZtdTboH6Error struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    idResponseZtdTboH6ErrorJSON `json:"-"`
}

// idResponseZtdTboH6ErrorJSON contains the JSON metadata for the struct
// [IDResponseZtdTboH6Error]
type idResponseZtdTboH6ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseZtdTboH6Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponseZtdTboH6Message struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    idResponseZtdTboH6MessageJSON `json:"-"`
}

// idResponseZtdTboH6MessageJSON contains the JSON metadata for the struct
// [IDResponseZtdTboH6Message]
type idResponseZtdTboH6MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseZtdTboH6Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponseZtdTboH6Result struct {
	ID   interface{}                  `json:"id"`
	JSON idResponseZtdTboH6ResultJSON `json:"-"`
}

// idResponseZtdTboH6ResultJSON contains the JSON metadata for the struct
// [IDResponseZtdTboH6Result]
type idResponseZtdTboH6ResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseZtdTboH6Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IDResponseZtdTboH6Success bool

const (
	IDResponseZtdTboH6SuccessTrue IDResponseZtdTboH6Success = true
)

type LoadBalancer struct {
	Errors   []LoadBalancerError   `json:"errors"`
	Messages []LoadBalancerMessage `json:"messages"`
	Result   LoadBalancerResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerSuccess `json:"success"`
	JSON    loadBalancerJSON    `json:"-"`
}

// loadBalancerJSON contains the JSON metadata for the struct [LoadBalancer]
type loadBalancerJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerError struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    loadBalancerErrorJSON `json:"-"`
}

// loadBalancerErrorJSON contains the JSON metadata for the struct
// [LoadBalancerError]
type loadBalancerErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMessage struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    loadBalancerMessageJSON `json:"-"`
}

// loadBalancerMessageJSON contains the JSON metadata for the struct
// [LoadBalancerMessage]
type loadBalancerMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerResult struct {
	ID interface{} `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerResultAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerResultLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                          `json:"modified_on" format:"date-time"`
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
	// Configures pool weights for random steering. When steering_policy is 'random', a
	// random pool is selected with probability proportional to these pool weights.
	RandomSteering LoadBalancerResultRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancerResultRule `json:"rules"`
	// Configures cookie attributes for session affinity cookie.
	SessionAffinityAttributes LoadBalancerResultSessionAffinityAttributes `json:"session_affinity_attributes"`
	// Time, in seconds, until this load balancer's session affinity cookie expires
	// after being created. This parameter is ignored unless a supported session
	// affinity policy is set. The current default of 23 hours will be used unless
	// session_affinity_ttl is explicitly set. The accepted range of values is between
	// [1800, 604800]. Once the expiry time has been reached, subsequent requests may
	// get sent to a different origin server.
	SessionAffinityTtl float64 `json:"session_affinity_ttl"`
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
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy LoadBalancerResultSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl  float64                `json:"ttl"`
	JSON loadBalancerResultJSON `json:"-"`
}

// loadBalancerResultJSON contains the JSON metadata for the struct
// [LoadBalancerResult]
type loadBalancerResultJSON struct {
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
	SessionAffinityAttributes apijson.Field
	SessionAffinityTtl        apijson.Field
	SteeringPolicy            apijson.Field
	Ttl                       apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LoadBalancerResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerResultAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                  `json:"failover_across_pools"`
	JSON                loadBalancerResultAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerResultAdaptiveRoutingJSON contains the JSON metadata for the struct
// [LoadBalancerResultAdaptiveRouting]
type loadBalancerResultAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerResultAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerResultLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerResultLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerResultLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerResultLocationStrategyJSON      `json:"-"`
}

// loadBalancerResultLocationStrategyJSON contains the JSON metadata for the struct
// [LoadBalancerResultLocationStrategy]
type loadBalancerResultLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerResultLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerResultLocationStrategyMode string

const (
	LoadBalancerResultLocationStrategyModePop        LoadBalancerResultLocationStrategyMode = "pop"
	LoadBalancerResultLocationStrategyModeResolverIP LoadBalancerResultLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerResultLocationStrategyPreferEcs string

const (
	LoadBalancerResultLocationStrategyPreferEcsAlways    LoadBalancerResultLocationStrategyPreferEcs = "always"
	LoadBalancerResultLocationStrategyPreferEcsNever     LoadBalancerResultLocationStrategyPreferEcs = "never"
	LoadBalancerResultLocationStrategyPreferEcsProximity LoadBalancerResultLocationStrategyPreferEcs = "proximity"
	LoadBalancerResultLocationStrategyPreferEcsGeo       LoadBalancerResultLocationStrategyPreferEcs = "geo"
)

// Configures pool weights for random steering. When steering_policy is 'random', a
// random pool is selected with probability proportional to these pool weights.
type LoadBalancerResultRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                          `json:"pool_weights"`
	JSON        loadBalancerResultRandomSteeringJSON `json:"-"`
}

// loadBalancerResultRandomSteeringJSON contains the JSON metadata for the struct
// [LoadBalancerResultRandomSteering]
type loadBalancerResultRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerResultRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerResultRule struct {
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
	FixedResponse LoadBalancerResultRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancerResultRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                       `json:"terminates"`
	JSON       loadBalancerResultRuleJSON `json:"-"`
}

// loadBalancerResultRuleJSON contains the JSON metadata for the struct
// [LoadBalancerResultRule]
type loadBalancerResultRuleJSON struct {
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

func (r *LoadBalancerResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerResultRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                    `json:"status_code"`
	JSON       loadBalancerResultRulesFixedResponseJSON `json:"-"`
}

// loadBalancerResultRulesFixedResponseJSON contains the JSON metadata for the
// struct [LoadBalancerResultRulesFixedResponse]
type loadBalancerResultRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerResultRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerResultRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerResultRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerResultRulesOverridesLocationStrategy `json:"location_strategy"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools interface{} `json:"pop_pools"`
	// Configures pool weights for random steering. When steering_policy is 'random', a
	// random pool is selected with probability proportional to these pool weights.
	RandomSteering LoadBalancerResultRulesOverridesRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// Configures cookie attributes for session affinity cookie.
	SessionAffinityAttributes LoadBalancerResultRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
	// Time, in seconds, until this load balancer's session affinity cookie expires
	// after being created. This parameter is ignored unless a supported session
	// affinity policy is set. The current default of 23 hours will be used unless
	// session_affinity_ttl is explicitly set. The accepted range of values is between
	// [1800, 604800]. Once the expiry time has been reached, subsequent requests may
	// get sent to a different origin server.
	SessionAffinityTtl float64 `json:"session_affinity_ttl"`
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
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy LoadBalancerResultRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl  float64                              `json:"ttl"`
	JSON loadBalancerResultRulesOverridesJSON `json:"-"`
}

// loadBalancerResultRulesOverridesJSON contains the JSON metadata for the struct
// [LoadBalancerResultRulesOverrides]
type loadBalancerResultRulesOverridesJSON struct {
	AdaptiveRouting           apijson.Field
	CountryPools              apijson.Field
	DefaultPools              apijson.Field
	FallbackPool              apijson.Field
	LocationStrategy          apijson.Field
	PopPools                  apijson.Field
	RandomSteering            apijson.Field
	RegionPools               apijson.Field
	SessionAffinityAttributes apijson.Field
	SessionAffinityTtl        apijson.Field
	SteeringPolicy            apijson.Field
	Ttl                       apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LoadBalancerResultRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerResultRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                `json:"failover_across_pools"`
	JSON                loadBalancerResultRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerResultRulesOverridesAdaptiveRoutingJSON contains the JSON metadata
// for the struct [LoadBalancerResultRulesOverridesAdaptiveRouting]
type loadBalancerResultRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerResultRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerResultRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerResultRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerResultRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerResultRulesOverridesLocationStrategyJSON      `json:"-"`
}

// loadBalancerResultRulesOverridesLocationStrategyJSON contains the JSON metadata
// for the struct [LoadBalancerResultRulesOverridesLocationStrategy]
type loadBalancerResultRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerResultRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerResultRulesOverridesLocationStrategyMode string

const (
	LoadBalancerResultRulesOverridesLocationStrategyModePop        LoadBalancerResultRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerResultRulesOverridesLocationStrategyModeResolverIP LoadBalancerResultRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerResultRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerResultRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerResultRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerResultRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerResultRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerResultRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerResultRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerResultRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerResultRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights for random steering. When steering_policy is 'random', a
// random pool is selected with probability proportional to these pool weights.
type LoadBalancerResultRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                        `json:"pool_weights"`
	JSON        loadBalancerResultRulesOverridesRandomSteeringJSON `json:"-"`
}

// loadBalancerResultRulesOverridesRandomSteeringJSON contains the JSON metadata
// for the struct [LoadBalancerResultRulesOverridesRandomSteering]
type loadBalancerResultRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerResultRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures cookie attributes for session affinity cookie.
type LoadBalancerResultRulesOverridesSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration float64 `json:"drain_duration"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite LoadBalancerResultRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerResultRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. Value "none" means no failover takes place for sessions
	// pinned to the origin (default). Value "temporary" means traffic will be sent to
	// another other healthy origin until the originally pinned origin is available;
	// note that this can potentially result in heavy origin flapping. Value "sticky"
	// means the session affinity cookie is updated and subsequent requests are sent to
	// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
	// and Bandwidth Alliance.
	ZeroDowntimeFailover LoadBalancerResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerResultRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerResultRulesOverridesSessionAffinityAttributesJSON contains the JSON
// metadata for the struct
// [LoadBalancerResultRulesOverridesSessionAffinityAttributes]
type loadBalancerResultRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerResultRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerResultRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerResultRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerResultRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerResultRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerResultRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerResultRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerResultRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerResultRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerResultRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerResultRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerResultRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerResultRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerResultRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerResultRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerResultRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerResultRulesOverridesSessionAffinityAttributesSecure = "Never"
)

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. Value "none" means no failover takes place for sessions
// pinned to the origin (default). Value "temporary" means traffic will be sent to
// another other healthy origin until the originally pinned origin is available;
// note that this can potentially result in heavy origin flapping. Value "sticky"
// means the session affinity cookie is updated and subsequent requests are sent to
// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
// and Bandwidth Alliance.
type LoadBalancerResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerResultRulesOverridesSteeringPolicy string

const (
	LoadBalancerResultRulesOverridesSteeringPolicyOff            LoadBalancerResultRulesOverridesSteeringPolicy = "off"
	LoadBalancerResultRulesOverridesSteeringPolicyGeo            LoadBalancerResultRulesOverridesSteeringPolicy = "geo"
	LoadBalancerResultRulesOverridesSteeringPolicyRandom         LoadBalancerResultRulesOverridesSteeringPolicy = "random"
	LoadBalancerResultRulesOverridesSteeringPolicyDynamicLatency LoadBalancerResultRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerResultRulesOverridesSteeringPolicyProximity      LoadBalancerResultRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerResultRulesOverridesSteeringPolicyEmpty          LoadBalancerResultRulesOverridesSteeringPolicy = "\"\""
)

// Configures cookie attributes for session affinity cookie.
type LoadBalancerResultSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration float64 `json:"drain_duration"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite LoadBalancerResultSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerResultSessionAffinityAttributesSecure `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. Value "none" means no failover takes place for sessions
	// pinned to the origin (default). Value "temporary" means traffic will be sent to
	// another other healthy origin until the originally pinned origin is available;
	// note that this can potentially result in heavy origin flapping. Value "sticky"
	// means the session affinity cookie is updated and subsequent requests are sent to
	// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
	// and Bandwidth Alliance.
	ZeroDowntimeFailover LoadBalancerResultSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerResultSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerResultSessionAffinityAttributesJSON contains the JSON metadata for
// the struct [LoadBalancerResultSessionAffinityAttributes]
type loadBalancerResultSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerResultSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerResultSessionAffinityAttributesSamesite string

const (
	LoadBalancerResultSessionAffinityAttributesSamesiteAuto   LoadBalancerResultSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerResultSessionAffinityAttributesSamesiteLax    LoadBalancerResultSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerResultSessionAffinityAttributesSamesiteNone   LoadBalancerResultSessionAffinityAttributesSamesite = "None"
	LoadBalancerResultSessionAffinityAttributesSamesiteStrict LoadBalancerResultSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerResultSessionAffinityAttributesSecure string

const (
	LoadBalancerResultSessionAffinityAttributesSecureAuto   LoadBalancerResultSessionAffinityAttributesSecure = "Auto"
	LoadBalancerResultSessionAffinityAttributesSecureAlways LoadBalancerResultSessionAffinityAttributesSecure = "Always"
	LoadBalancerResultSessionAffinityAttributesSecureNever  LoadBalancerResultSessionAffinityAttributesSecure = "Never"
)

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. Value "none" means no failover takes place for sessions
// pinned to the origin (default). Value "temporary" means traffic will be sent to
// another other healthy origin until the originally pinned origin is available;
// note that this can potentially result in heavy origin flapping. Value "sticky"
// means the session affinity cookie is updated and subsequent requests are sent to
// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
// and Bandwidth Alliance.
type LoadBalancerResultSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerResultSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerResultSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerResultSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerResultSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerResultSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerResultSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerResultSteeringPolicy string

const (
	LoadBalancerResultSteeringPolicyOff            LoadBalancerResultSteeringPolicy = "off"
	LoadBalancerResultSteeringPolicyGeo            LoadBalancerResultSteeringPolicy = "geo"
	LoadBalancerResultSteeringPolicyRandom         LoadBalancerResultSteeringPolicy = "random"
	LoadBalancerResultSteeringPolicyDynamicLatency LoadBalancerResultSteeringPolicy = "dynamic_latency"
	LoadBalancerResultSteeringPolicyProximity      LoadBalancerResultSteeringPolicy = "proximity"
	LoadBalancerResultSteeringPolicyEmpty          LoadBalancerResultSteeringPolicy = "\"\""
)

// Whether the API call was successful
type LoadBalancerSuccess bool

const (
	LoadBalancerSuccessTrue LoadBalancerSuccess = true
)

type LoadBalancerCollection struct {
	Errors     []LoadBalancerCollectionError    `json:"errors"`
	Messages   []LoadBalancerCollectionMessage  `json:"messages"`
	Result     []LoadBalancerCollectionResult   `json:"result"`
	ResultInfo LoadBalancerCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerCollectionSuccess `json:"success"`
	JSON    loadBalancerCollectionJSON    `json:"-"`
}

// loadBalancerCollectionJSON contains the JSON metadata for the struct
// [LoadBalancerCollection]
type loadBalancerCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerCollectionError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    loadBalancerCollectionErrorJSON `json:"-"`
}

// loadBalancerCollectionErrorJSON contains the JSON metadata for the struct
// [LoadBalancerCollectionError]
type loadBalancerCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerCollectionMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    loadBalancerCollectionMessageJSON `json:"-"`
}

// loadBalancerCollectionMessageJSON contains the JSON metadata for the struct
// [LoadBalancerCollectionMessage]
type loadBalancerCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerCollectionResult struct {
	ID interface{} `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerCollectionResultAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerCollectionResultLocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time                                    `json:"modified_on" format:"date-time"`
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
	// Configures pool weights for random steering. When steering_policy is 'random', a
	// random pool is selected with probability proportional to these pool weights.
	RandomSteering LoadBalancerCollectionResultRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancerCollectionResultRule `json:"rules"`
	// Configures cookie attributes for session affinity cookie.
	SessionAffinityAttributes LoadBalancerCollectionResultSessionAffinityAttributes `json:"session_affinity_attributes"`
	// Time, in seconds, until this load balancer's session affinity cookie expires
	// after being created. This parameter is ignored unless a supported session
	// affinity policy is set. The current default of 23 hours will be used unless
	// session_affinity_ttl is explicitly set. The accepted range of values is between
	// [1800, 604800]. Once the expiry time has been reached, subsequent requests may
	// get sent to a different origin server.
	SessionAffinityTtl float64 `json:"session_affinity_ttl"`
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
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy LoadBalancerCollectionResultSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl  float64                          `json:"ttl"`
	JSON loadBalancerCollectionResultJSON `json:"-"`
}

// loadBalancerCollectionResultJSON contains the JSON metadata for the struct
// [LoadBalancerCollectionResult]
type loadBalancerCollectionResultJSON struct {
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
	SessionAffinityAttributes apijson.Field
	SessionAffinityTtl        apijson.Field
	SteeringPolicy            apijson.Field
	Ttl                       apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LoadBalancerCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerCollectionResultAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                            `json:"failover_across_pools"`
	JSON                loadBalancerCollectionResultAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerCollectionResultAdaptiveRoutingJSON contains the JSON metadata for
// the struct [LoadBalancerCollectionResultAdaptiveRouting]
type loadBalancerCollectionResultAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerCollectionResultLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerCollectionResultLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerCollectionResultLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerCollectionResultLocationStrategyJSON      `json:"-"`
}

// loadBalancerCollectionResultLocationStrategyJSON contains the JSON metadata for
// the struct [LoadBalancerCollectionResultLocationStrategy]
type loadBalancerCollectionResultLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerCollectionResultLocationStrategyMode string

const (
	LoadBalancerCollectionResultLocationStrategyModePop        LoadBalancerCollectionResultLocationStrategyMode = "pop"
	LoadBalancerCollectionResultLocationStrategyModeResolverIP LoadBalancerCollectionResultLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerCollectionResultLocationStrategyPreferEcs string

const (
	LoadBalancerCollectionResultLocationStrategyPreferEcsAlways    LoadBalancerCollectionResultLocationStrategyPreferEcs = "always"
	LoadBalancerCollectionResultLocationStrategyPreferEcsNever     LoadBalancerCollectionResultLocationStrategyPreferEcs = "never"
	LoadBalancerCollectionResultLocationStrategyPreferEcsProximity LoadBalancerCollectionResultLocationStrategyPreferEcs = "proximity"
	LoadBalancerCollectionResultLocationStrategyPreferEcsGeo       LoadBalancerCollectionResultLocationStrategyPreferEcs = "geo"
)

// Configures pool weights for random steering. When steering_policy is 'random', a
// random pool is selected with probability proportional to these pool weights.
type LoadBalancerCollectionResultRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                    `json:"pool_weights"`
	JSON        loadBalancerCollectionResultRandomSteeringJSON `json:"-"`
}

// loadBalancerCollectionResultRandomSteeringJSON contains the JSON metadata for
// the struct [LoadBalancerCollectionResultRandomSteering]
type loadBalancerCollectionResultRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancerCollectionResultRule struct {
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
	FixedResponse LoadBalancerCollectionResultRulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancerCollectionResultRulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                                 `json:"terminates"`
	JSON       loadBalancerCollectionResultRuleJSON `json:"-"`
}

// loadBalancerCollectionResultRuleJSON contains the JSON metadata for the struct
// [LoadBalancerCollectionResultRule]
type loadBalancerCollectionResultRuleJSON struct {
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

func (r *LoadBalancerCollectionResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancerCollectionResultRulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                                              `json:"status_code"`
	JSON       loadBalancerCollectionResultRulesFixedResponseJSON `json:"-"`
}

// loadBalancerCollectionResultRulesFixedResponseJSON contains the JSON metadata
// for the struct [LoadBalancerCollectionResultRulesFixedResponse]
type loadBalancerCollectionResultRulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultRulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancerCollectionResultRulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting LoadBalancerCollectionResultRulesOverridesAdaptiveRouting `json:"adaptive_routing"`
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
	LocationStrategy LoadBalancerCollectionResultRulesOverridesLocationStrategy `json:"location_strategy"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools interface{} `json:"pop_pools"`
	// Configures pool weights for random steering. When steering_policy is 'random', a
	// random pool is selected with probability proportional to these pool weights.
	RandomSteering LoadBalancerCollectionResultRulesOverridesRandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools interface{} `json:"region_pools"`
	// Configures cookie attributes for session affinity cookie.
	SessionAffinityAttributes LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributes `json:"session_affinity_attributes"`
	// Time, in seconds, until this load balancer's session affinity cookie expires
	// after being created. This parameter is ignored unless a supported session
	// affinity policy is set. The current default of 23 hours will be used unless
	// session_affinity_ttl is explicitly set. The accepted range of values is between
	// [1800, 604800]. Once the expiry time has been reached, subsequent requests may
	// get sent to a different origin server.
	SessionAffinityTtl float64 `json:"session_affinity_ttl"`
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
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy LoadBalancerCollectionResultRulesOverridesSteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl  float64                                        `json:"ttl"`
	JSON loadBalancerCollectionResultRulesOverridesJSON `json:"-"`
}

// loadBalancerCollectionResultRulesOverridesJSON contains the JSON metadata for
// the struct [LoadBalancerCollectionResultRulesOverrides]
type loadBalancerCollectionResultRulesOverridesJSON struct {
	AdaptiveRouting           apijson.Field
	CountryPools              apijson.Field
	DefaultPools              apijson.Field
	FallbackPool              apijson.Field
	LocationStrategy          apijson.Field
	PopPools                  apijson.Field
	RandomSteering            apijson.Field
	RegionPools               apijson.Field
	SessionAffinityAttributes apijson.Field
	SessionAffinityTtl        apijson.Field
	SteeringPolicy            apijson.Field
	Ttl                       apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultRulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type LoadBalancerCollectionResultRulesOverridesAdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                                                          `json:"failover_across_pools"`
	JSON                loadBalancerCollectionResultRulesOverridesAdaptiveRoutingJSON `json:"-"`
}

// loadBalancerCollectionResultRulesOverridesAdaptiveRoutingJSON contains the JSON
// metadata for the struct
// [LoadBalancerCollectionResultRulesOverridesAdaptiveRouting]
type loadBalancerCollectionResultRulesOverridesAdaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultRulesOverridesAdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LoadBalancerCollectionResultRulesOverridesLocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LoadBalancerCollectionResultRulesOverridesLocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LoadBalancerCollectionResultRulesOverridesLocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      loadBalancerCollectionResultRulesOverridesLocationStrategyJSON      `json:"-"`
}

// loadBalancerCollectionResultRulesOverridesLocationStrategyJSON contains the JSON
// metadata for the struct
// [LoadBalancerCollectionResultRulesOverridesLocationStrategy]
type loadBalancerCollectionResultRulesOverridesLocationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultRulesOverridesLocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LoadBalancerCollectionResultRulesOverridesLocationStrategyMode string

const (
	LoadBalancerCollectionResultRulesOverridesLocationStrategyModePop        LoadBalancerCollectionResultRulesOverridesLocationStrategyMode = "pop"
	LoadBalancerCollectionResultRulesOverridesLocationStrategyModeResolverIP LoadBalancerCollectionResultRulesOverridesLocationStrategyMode = "resolver_ip"
)

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LoadBalancerCollectionResultRulesOverridesLocationStrategyPreferEcs string

const (
	LoadBalancerCollectionResultRulesOverridesLocationStrategyPreferEcsAlways    LoadBalancerCollectionResultRulesOverridesLocationStrategyPreferEcs = "always"
	LoadBalancerCollectionResultRulesOverridesLocationStrategyPreferEcsNever     LoadBalancerCollectionResultRulesOverridesLocationStrategyPreferEcs = "never"
	LoadBalancerCollectionResultRulesOverridesLocationStrategyPreferEcsProximity LoadBalancerCollectionResultRulesOverridesLocationStrategyPreferEcs = "proximity"
	LoadBalancerCollectionResultRulesOverridesLocationStrategyPreferEcsGeo       LoadBalancerCollectionResultRulesOverridesLocationStrategyPreferEcs = "geo"
)

// Configures pool weights for random steering. When steering_policy is 'random', a
// random pool is selected with probability proportional to these pool weights.
type LoadBalancerCollectionResultRulesOverridesRandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights interface{}                                                  `json:"pool_weights"`
	JSON        loadBalancerCollectionResultRulesOverridesRandomSteeringJSON `json:"-"`
}

// loadBalancerCollectionResultRulesOverridesRandomSteeringJSON contains the JSON
// metadata for the struct
// [LoadBalancerCollectionResultRulesOverridesRandomSteering]
type loadBalancerCollectionResultRulesOverridesRandomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultRulesOverridesRandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures cookie attributes for session affinity cookie.
type LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration float64 `json:"drain_duration"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSecure `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. Value "none" means no failover takes place for sessions
	// pinned to the origin (default). Value "temporary" means traffic will be sent to
	// another other healthy origin until the originally pinned origin is available;
	// note that this can potentially result in heavy origin flapping. Value "sticky"
	// means the session affinity cookie is updated and subsequent requests are sent to
	// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
	// and Bandwidth Alliance.
	ZeroDowntimeFailover LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerCollectionResultRulesOverridesSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerCollectionResultRulesOverridesSessionAffinityAttributesJSON contains
// the JSON metadata for the struct
// [LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributes]
type loadBalancerCollectionResultRulesOverridesSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSamesite string

const (
	LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSamesiteAuto   LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSamesiteLax    LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSamesiteNone   LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSamesite = "None"
	LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSamesiteStrict LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSecure string

const (
	LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSecureAuto   LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSecure = "Auto"
	LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSecureAlways LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSecure = "Always"
	LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSecureNever  LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesSecure = "Never"
)

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. Value "none" means no failover takes place for sessions
// pinned to the origin (default). Value "temporary" means traffic will be sent to
// another other healthy origin until the originally pinned origin is available;
// note that this can potentially result in heavy origin flapping. Value "sticky"
// means the session affinity cookie is updated and subsequent requests are sent to
// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
// and Bandwidth Alliance.
type LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerCollectionResultRulesOverridesSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerCollectionResultRulesOverridesSteeringPolicy string

const (
	LoadBalancerCollectionResultRulesOverridesSteeringPolicyOff            LoadBalancerCollectionResultRulesOverridesSteeringPolicy = "off"
	LoadBalancerCollectionResultRulesOverridesSteeringPolicyGeo            LoadBalancerCollectionResultRulesOverridesSteeringPolicy = "geo"
	LoadBalancerCollectionResultRulesOverridesSteeringPolicyRandom         LoadBalancerCollectionResultRulesOverridesSteeringPolicy = "random"
	LoadBalancerCollectionResultRulesOverridesSteeringPolicyDynamicLatency LoadBalancerCollectionResultRulesOverridesSteeringPolicy = "dynamic_latency"
	LoadBalancerCollectionResultRulesOverridesSteeringPolicyProximity      LoadBalancerCollectionResultRulesOverridesSteeringPolicy = "proximity"
	LoadBalancerCollectionResultRulesOverridesSteeringPolicyEmpty          LoadBalancerCollectionResultRulesOverridesSteeringPolicy = "\"\""
)

// Configures cookie attributes for session affinity cookie.
type LoadBalancerCollectionResultSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration float64 `json:"drain_duration"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite LoadBalancerCollectionResultSessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure LoadBalancerCollectionResultSessionAffinityAttributesSecure `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. Value "none" means no failover takes place for sessions
	// pinned to the origin (default). Value "temporary" means traffic will be sent to
	// another other healthy origin until the originally pinned origin is available;
	// note that this can potentially result in heavy origin flapping. Value "sticky"
	// means the session affinity cookie is updated and subsequent requests are sent to
	// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
	// and Bandwidth Alliance.
	ZeroDowntimeFailover LoadBalancerCollectionResultSessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 loadBalancerCollectionResultSessionAffinityAttributesJSON                 `json:"-"`
}

// loadBalancerCollectionResultSessionAffinityAttributesJSON contains the JSON
// metadata for the struct [LoadBalancerCollectionResultSessionAffinityAttributes]
type loadBalancerCollectionResultSessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultSessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type LoadBalancerCollectionResultSessionAffinityAttributesSamesite string

const (
	LoadBalancerCollectionResultSessionAffinityAttributesSamesiteAuto   LoadBalancerCollectionResultSessionAffinityAttributesSamesite = "Auto"
	LoadBalancerCollectionResultSessionAffinityAttributesSamesiteLax    LoadBalancerCollectionResultSessionAffinityAttributesSamesite = "Lax"
	LoadBalancerCollectionResultSessionAffinityAttributesSamesiteNone   LoadBalancerCollectionResultSessionAffinityAttributesSamesite = "None"
	LoadBalancerCollectionResultSessionAffinityAttributesSamesiteStrict LoadBalancerCollectionResultSessionAffinityAttributesSamesite = "Strict"
)

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type LoadBalancerCollectionResultSessionAffinityAttributesSecure string

const (
	LoadBalancerCollectionResultSessionAffinityAttributesSecureAuto   LoadBalancerCollectionResultSessionAffinityAttributesSecure = "Auto"
	LoadBalancerCollectionResultSessionAffinityAttributesSecureAlways LoadBalancerCollectionResultSessionAffinityAttributesSecure = "Always"
	LoadBalancerCollectionResultSessionAffinityAttributesSecureNever  LoadBalancerCollectionResultSessionAffinityAttributesSecure = "Never"
)

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. Value "none" means no failover takes place for sessions
// pinned to the origin (default). Value "temporary" means traffic will be sent to
// another other healthy origin until the originally pinned origin is available;
// note that this can potentially result in heavy origin flapping. Value "sticky"
// means the session affinity cookie is updated and subsequent requests are sent to
// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
// and Bandwidth Alliance.
type LoadBalancerCollectionResultSessionAffinityAttributesZeroDowntimeFailover string

const (
	LoadBalancerCollectionResultSessionAffinityAttributesZeroDowntimeFailoverNone      LoadBalancerCollectionResultSessionAffinityAttributesZeroDowntimeFailover = "none"
	LoadBalancerCollectionResultSessionAffinityAttributesZeroDowntimeFailoverTemporary LoadBalancerCollectionResultSessionAffinityAttributesZeroDowntimeFailover = "temporary"
	LoadBalancerCollectionResultSessionAffinityAttributesZeroDowntimeFailoverSticky    LoadBalancerCollectionResultSessionAffinityAttributesZeroDowntimeFailover = "sticky"
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
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type LoadBalancerCollectionResultSteeringPolicy string

const (
	LoadBalancerCollectionResultSteeringPolicyOff            LoadBalancerCollectionResultSteeringPolicy = "off"
	LoadBalancerCollectionResultSteeringPolicyGeo            LoadBalancerCollectionResultSteeringPolicy = "geo"
	LoadBalancerCollectionResultSteeringPolicyRandom         LoadBalancerCollectionResultSteeringPolicy = "random"
	LoadBalancerCollectionResultSteeringPolicyDynamicLatency LoadBalancerCollectionResultSteeringPolicy = "dynamic_latency"
	LoadBalancerCollectionResultSteeringPolicyProximity      LoadBalancerCollectionResultSteeringPolicy = "proximity"
	LoadBalancerCollectionResultSteeringPolicyEmpty          LoadBalancerCollectionResultSteeringPolicy = "\"\""
)

type LoadBalancerCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       loadBalancerCollectionResultInfoJSON `json:"-"`
}

// loadBalancerCollectionResultInfoJSON contains the JSON metadata for the struct
// [LoadBalancerCollectionResultInfo]
type loadBalancerCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerCollectionSuccess bool

const (
	LoadBalancerCollectionSuccessTrue LoadBalancerCollectionSuccess = true
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
	// Configures pool weights for random steering. When steering_policy is 'random', a
	// random pool is selected with probability proportional to these pool weights.
	RandomSteering param.Field[ZoneLoadBalancerNewParamsRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]ZoneLoadBalancerNewParamsRule] `json:"rules"`
	// Configures cookie attributes for session affinity cookie.
	SessionAffinityAttributes param.Field[ZoneLoadBalancerNewParamsSessionAffinityAttributes] `json:"session_affinity_attributes"`
	// Time, in seconds, until this load balancer's session affinity cookie expires
	// after being created. This parameter is ignored unless a supported session
	// affinity policy is set. The current default of 23 hours will be used unless
	// session_affinity_ttl is explicitly set. The accepted range of values is between
	// [1800, 604800]. Once the expiry time has been reached, subsequent requests may
	// get sent to a different origin server.
	SessionAffinityTtl param.Field[float64] `json:"session_affinity_ttl"`
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
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[ZoneLoadBalancerNewParamsSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl param.Field[float64] `json:"ttl"`
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

// Configures pool weights for random steering. When steering_policy is 'random', a
// random pool is selected with probability proportional to these pool weights.
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
	// Configures pool weights for random steering. When steering_policy is 'random', a
	// random pool is selected with probability proportional to these pool weights.
	RandomSteering param.Field[ZoneLoadBalancerNewParamsRulesOverridesRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// Configures cookie attributes for session affinity cookie.
	SessionAffinityAttributes param.Field[ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributes] `json:"session_affinity_attributes"`
	// Time, in seconds, until this load balancer's session affinity cookie expires
	// after being created. This parameter is ignored unless a supported session
	// affinity policy is set. The current default of 23 hours will be used unless
	// session_affinity_ttl is explicitly set. The accepted range of values is between
	// [1800, 604800]. Once the expiry time has been reached, subsequent requests may
	// get sent to a different origin server.
	SessionAffinityTtl param.Field[float64] `json:"session_affinity_ttl"`
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
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl param.Field[float64] `json:"ttl"`
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

// Configures pool weights for random steering. When steering_policy is 'random', a
// random pool is selected with probability proportional to these pool weights.
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

// Configures cookie attributes for session affinity cookie.
type ZoneLoadBalancerNewParamsRulesOverridesSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
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
	// affinity is enabled. Value "none" means no failover takes place for sessions
	// pinned to the origin (default). Value "temporary" means traffic will be sent to
	// another other healthy origin until the originally pinned origin is available;
	// note that this can potentially result in heavy origin flapping. Value "sticky"
	// means the session affinity cookie is updated and subsequent requests are sent to
	// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
	// and Bandwidth Alliance.
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
// affinity is enabled. Value "none" means no failover takes place for sessions
// pinned to the origin (default). Value "temporary" means traffic will be sent to
// another other healthy origin until the originally pinned origin is available;
// note that this can potentially result in heavy origin flapping. Value "sticky"
// means the session affinity cookie is updated and subsequent requests are sent to
// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
// and Bandwidth Alliance.
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
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy string

const (
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyOff            ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "off"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyGeo            ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "geo"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyRandom         ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "random"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyDynamicLatency ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyProximity      ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "proximity"
	ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicyEmpty          ZoneLoadBalancerNewParamsRulesOverridesSteeringPolicy = "\"\""
)

// Configures cookie attributes for session affinity cookie.
type ZoneLoadBalancerNewParamsSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
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
	// affinity is enabled. Value "none" means no failover takes place for sessions
	// pinned to the origin (default). Value "temporary" means traffic will be sent to
	// another other healthy origin until the originally pinned origin is available;
	// note that this can potentially result in heavy origin flapping. Value "sticky"
	// means the session affinity cookie is updated and subsequent requests are sent to
	// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
	// and Bandwidth Alliance.
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
// affinity is enabled. Value "none" means no failover takes place for sessions
// pinned to the origin (default). Value "temporary" means traffic will be sent to
// another other healthy origin until the originally pinned origin is available;
// note that this can potentially result in heavy origin flapping. Value "sticky"
// means the session affinity cookie is updated and subsequent requests are sent to
// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
// and Bandwidth Alliance.
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
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type ZoneLoadBalancerNewParamsSteeringPolicy string

const (
	ZoneLoadBalancerNewParamsSteeringPolicyOff            ZoneLoadBalancerNewParamsSteeringPolicy = "off"
	ZoneLoadBalancerNewParamsSteeringPolicyGeo            ZoneLoadBalancerNewParamsSteeringPolicy = "geo"
	ZoneLoadBalancerNewParamsSteeringPolicyRandom         ZoneLoadBalancerNewParamsSteeringPolicy = "random"
	ZoneLoadBalancerNewParamsSteeringPolicyDynamicLatency ZoneLoadBalancerNewParamsSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerNewParamsSteeringPolicyProximity      ZoneLoadBalancerNewParamsSteeringPolicy = "proximity"
	ZoneLoadBalancerNewParamsSteeringPolicyEmpty          ZoneLoadBalancerNewParamsSteeringPolicy = "\"\""
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
	// Configures pool weights for random steering. When steering_policy is 'random', a
	// random pool is selected with probability proportional to these pool weights.
	RandomSteering param.Field[ZoneLoadBalancerUpdateParamsRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]ZoneLoadBalancerUpdateParamsRule] `json:"rules"`
	// Configures cookie attributes for session affinity cookie.
	SessionAffinityAttributes param.Field[ZoneLoadBalancerUpdateParamsSessionAffinityAttributes] `json:"session_affinity_attributes"`
	// Time, in seconds, until this load balancer's session affinity cookie expires
	// after being created. This parameter is ignored unless a supported session
	// affinity policy is set. The current default of 23 hours will be used unless
	// session_affinity_ttl is explicitly set. The accepted range of values is between
	// [1800, 604800]. Once the expiry time has been reached, subsequent requests may
	// get sent to a different origin server.
	SessionAffinityTtl param.Field[float64] `json:"session_affinity_ttl"`
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
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[ZoneLoadBalancerUpdateParamsSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl param.Field[float64] `json:"ttl"`
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

// Configures pool weights for random steering. When steering_policy is 'random', a
// random pool is selected with probability proportional to these pool weights.
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
	// Configures pool weights for random steering. When steering_policy is 'random', a
	// random pool is selected with probability proportional to these pool weights.
	RandomSteering param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesRandomSteering] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[interface{}] `json:"region_pools"`
	// Configures cookie attributes for session affinity cookie.
	SessionAffinityAttributes param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes] `json:"session_affinity_attributes"`
	// Time, in seconds, until this load balancer's session affinity cookie expires
	// after being created. This parameter is ignored unless a supported session
	// affinity policy is set. The current default of 23 hours will be used unless
	// session_affinity_ttl is explicitly set. The accepted range of values is between
	// [1800, 604800]. Once the expiry time has been reached, subsequent requests may
	// get sent to a different origin server.
	SessionAffinityTtl param.Field[float64] `json:"session_affinity_ttl"`
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
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl param.Field[float64] `json:"ttl"`
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

// Configures pool weights for random steering. When steering_policy is 'random', a
// random pool is selected with probability proportional to these pool weights.
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

// Configures cookie attributes for session affinity cookie.
type ZoneLoadBalancerUpdateParamsRulesOverridesSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
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
	// affinity is enabled. Value "none" means no failover takes place for sessions
	// pinned to the origin (default). Value "temporary" means traffic will be sent to
	// another other healthy origin until the originally pinned origin is available;
	// note that this can potentially result in heavy origin flapping. Value "sticky"
	// means the session affinity cookie is updated and subsequent requests are sent to
	// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
	// and Bandwidth Alliance.
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
// affinity is enabled. Value "none" means no failover takes place for sessions
// pinned to the origin (default). Value "temporary" means traffic will be sent to
// another other healthy origin until the originally pinned origin is available;
// note that this can potentially result in heavy origin flapping. Value "sticky"
// means the session affinity cookie is updated and subsequent requests are sent to
// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
// and Bandwidth Alliance.
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
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy string

const (
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyOff            ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "off"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyGeo            ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "geo"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyRandom         ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "random"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyDynamicLatency ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyProximity      ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "proximity"
	ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicyEmpty          ZoneLoadBalancerUpdateParamsRulesOverridesSteeringPolicy = "\"\""
)

// Configures cookie attributes for session affinity cookie.
type ZoneLoadBalancerUpdateParamsSessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
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
	// affinity is enabled. Value "none" means no failover takes place for sessions
	// pinned to the origin (default). Value "temporary" means traffic will be sent to
	// another other healthy origin until the originally pinned origin is available;
	// note that this can potentially result in heavy origin flapping. Value "sticky"
	// means the session affinity cookie is updated and subsequent requests are sent to
	// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
	// and Bandwidth Alliance.
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
// affinity is enabled. Value "none" means no failover takes place for sessions
// pinned to the origin (default). Value "temporary" means traffic will be sent to
// another other healthy origin until the originally pinned origin is available;
// note that this can potentially result in heavy origin flapping. Value "sticky"
// means the session affinity cookie is updated and subsequent requests are sent to
// the new origin. This feature is currently incompatible with Argo, Tiered Cache,
// and Bandwidth Alliance.
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
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type ZoneLoadBalancerUpdateParamsSteeringPolicy string

const (
	ZoneLoadBalancerUpdateParamsSteeringPolicyOff            ZoneLoadBalancerUpdateParamsSteeringPolicy = "off"
	ZoneLoadBalancerUpdateParamsSteeringPolicyGeo            ZoneLoadBalancerUpdateParamsSteeringPolicy = "geo"
	ZoneLoadBalancerUpdateParamsSteeringPolicyRandom         ZoneLoadBalancerUpdateParamsSteeringPolicy = "random"
	ZoneLoadBalancerUpdateParamsSteeringPolicyDynamicLatency ZoneLoadBalancerUpdateParamsSteeringPolicy = "dynamic_latency"
	ZoneLoadBalancerUpdateParamsSteeringPolicyProximity      ZoneLoadBalancerUpdateParamsSteeringPolicy = "proximity"
	ZoneLoadBalancerUpdateParamsSteeringPolicyEmpty          ZoneLoadBalancerUpdateParamsSteeringPolicy = "\"\""
)
