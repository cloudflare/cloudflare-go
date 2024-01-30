// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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
	Errors   []ZoneLoadBalancerNewResponseError   `json:"errors,required"`
	Messages []ZoneLoadBalancerNewResponseMessage `json:"messages,required"`
	Result   ZoneLoadBalancerNewResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success ZoneLoadBalancerNewResponseSuccess `json:"success,required"`
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

// Union satisfied by [ZoneLoadBalancerNewResponseResultUnknown],
// [ZoneLoadBalancerNewResponseResultArray] or [shared.UnionString].
type ZoneLoadBalancerNewResponseResult interface {
	ImplementsZoneLoadBalancerNewResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneLoadBalancerNewResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneLoadBalancerNewResponseResultArray []interface{}

func (r ZoneLoadBalancerNewResponseResultArray) ImplementsZoneLoadBalancerNewResponseResult() {}

// Whether the API call was successful
type ZoneLoadBalancerNewResponseSuccess bool

const (
	ZoneLoadBalancerNewResponseSuccessTrue ZoneLoadBalancerNewResponseSuccess = true
)

type ZoneLoadBalancerGetResponse struct {
	Errors   []ZoneLoadBalancerGetResponseError   `json:"errors,required"`
	Messages []ZoneLoadBalancerGetResponseMessage `json:"messages,required"`
	Result   ZoneLoadBalancerGetResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success ZoneLoadBalancerGetResponseSuccess `json:"success,required"`
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

// Union satisfied by [ZoneLoadBalancerGetResponseResultUnknown],
// [ZoneLoadBalancerGetResponseResultArray] or [shared.UnionString].
type ZoneLoadBalancerGetResponseResult interface {
	ImplementsZoneLoadBalancerGetResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneLoadBalancerGetResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneLoadBalancerGetResponseResultArray []interface{}

func (r ZoneLoadBalancerGetResponseResultArray) ImplementsZoneLoadBalancerGetResponseResult() {}

// Whether the API call was successful
type ZoneLoadBalancerGetResponseSuccess bool

const (
	ZoneLoadBalancerGetResponseSuccessTrue ZoneLoadBalancerGetResponseSuccess = true
)

type ZoneLoadBalancerUpdateResponse struct {
	Errors   []ZoneLoadBalancerUpdateResponseError   `json:"errors,required"`
	Messages []ZoneLoadBalancerUpdateResponseMessage `json:"messages,required"`
	Result   ZoneLoadBalancerUpdateResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success ZoneLoadBalancerUpdateResponseSuccess `json:"success,required"`
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

// Union satisfied by [ZoneLoadBalancerUpdateResponseResultUnknown],
// [ZoneLoadBalancerUpdateResponseResultArray] or [shared.UnionString].
type ZoneLoadBalancerUpdateResponseResult interface {
	ImplementsZoneLoadBalancerUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneLoadBalancerUpdateResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneLoadBalancerUpdateResponseResultArray []interface{}

func (r ZoneLoadBalancerUpdateResponseResultArray) ImplementsZoneLoadBalancerUpdateResponseResult() {}

// Whether the API call was successful
type ZoneLoadBalancerUpdateResponseSuccess bool

const (
	ZoneLoadBalancerUpdateResponseSuccessTrue ZoneLoadBalancerUpdateResponseSuccess = true
)

type ZoneLoadBalancerListResponse struct {
	Errors   []ZoneLoadBalancerListResponseError   `json:"errors,required"`
	Messages []ZoneLoadBalancerListResponseMessage `json:"messages,required"`
	Result   ZoneLoadBalancerListResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success    ZoneLoadBalancerListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneLoadBalancerListResponseResultInfo `json:"result_info"`
	JSON       zoneLoadBalancerListResponseJSON       `json:"-"`
}

// zoneLoadBalancerListResponseJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerListResponse]
type zoneLoadBalancerListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
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

// Union satisfied by [ZoneLoadBalancerListResponseResultUnknown],
// [ZoneLoadBalancerListResponseResultArray] or [shared.UnionString].
type ZoneLoadBalancerListResponseResult interface {
	ImplementsZoneLoadBalancerListResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneLoadBalancerListResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneLoadBalancerListResponseResultArray []interface{}

func (r ZoneLoadBalancerListResponseResultArray) ImplementsZoneLoadBalancerListResponseResult() {}

// Whether the API call was successful
type ZoneLoadBalancerListResponseSuccess bool

const (
	ZoneLoadBalancerListResponseSuccessTrue ZoneLoadBalancerListResponseSuccess = true
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

type ZoneLoadBalancerDeleteResponse struct {
	Errors   []ZoneLoadBalancerDeleteResponseError   `json:"errors,required"`
	Messages []ZoneLoadBalancerDeleteResponseMessage `json:"messages,required"`
	Result   ZoneLoadBalancerDeleteResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success ZoneLoadBalancerDeleteResponseSuccess `json:"success,required"`
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

// Union satisfied by [ZoneLoadBalancerDeleteResponseResultUnknown],
// [ZoneLoadBalancerDeleteResponseResultArray] or [shared.UnionString].
type ZoneLoadBalancerDeleteResponseResult interface {
	ImplementsZoneLoadBalancerDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneLoadBalancerDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneLoadBalancerDeleteResponseResultArray []interface{}

func (r ZoneLoadBalancerDeleteResponseResultArray) ImplementsZoneLoadBalancerDeleteResponseResult() {}

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
