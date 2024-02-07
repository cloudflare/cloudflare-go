// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// LoadBalancerPoolService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerPoolService] method
// instead.
type LoadBalancerPoolService struct {
	Options    []option.RequestOption
	Health     *LoadBalancerPoolHealthService
	Previews   *LoadBalancerPoolPreviewService
	References *LoadBalancerPoolReferenceService
}

// NewLoadBalancerPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerPoolService(opts ...option.RequestOption) (r *LoadBalancerPoolService) {
	r = &LoadBalancerPoolService{}
	r.Options = opts
	r.Health = NewLoadBalancerPoolHealthService(opts...)
	r.Previews = NewLoadBalancerPoolPreviewService(opts...)
	r.References = NewLoadBalancerPoolReferenceService(opts...)
	return
}

// Fetch a single configured pool.
func (r *LoadBalancerPoolService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *LoadBalancerPoolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured pool.
func (r *LoadBalancerPoolService) Update(ctx context.Context, accountIdentifier string, identifier string, body LoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *LoadBalancerPoolUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured pool.
func (r *LoadBalancerPoolService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *LoadBalancerPoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new pool.
func (r *LoadBalancerPoolService) AccountLoadBalancerPoolsNewPool(ctx context.Context, accountIdentifier string, body LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams, opts ...option.RequestOption) (res *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured pools.
func (r *LoadBalancerPoolService) AccountLoadBalancerPoolsListPools(ctx context.Context, accountIdentifier string, query LoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams, opts ...option.RequestOption) (res *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Apply changes to a number of existing pools, overwriting the supplied
// properties. Pools are ordered by ascending `name`. Returns the list of affected
// pools. Supports the standard pagination query parameters, either
// `limit`/`offset` or `per_page`/`page`.
func (r *LoadBalancerPoolService) AccountLoadBalancerPoolsPatchPools(ctx context.Context, accountIdentifier string, body LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams, opts ...option.RequestOption) (res *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type LoadBalancerPoolGetResponse struct {
	Errors   []LoadBalancerPoolGetResponseError   `json:"errors"`
	Messages []LoadBalancerPoolGetResponseMessage `json:"messages"`
	Result   LoadBalancerPoolGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerPoolGetResponseSuccess `json:"success"`
	JSON    loadBalancerPoolGetResponseJSON    `json:"-"`
}

// loadBalancerPoolGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolGetResponse]
type loadBalancerPoolGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    loadBalancerPoolGetResponseErrorJSON `json:"-"`
}

// loadBalancerPoolGetResponseErrorJSON contains the JSON metadata for the struct
// [LoadBalancerPoolGetResponseError]
type loadBalancerPoolGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    loadBalancerPoolGetResponseMessageJSON `json:"-"`
}

// loadBalancerPoolGetResponseMessageJSON contains the JSON metadata for the struct
// [LoadBalancerPoolGetResponseMessage]
type loadBalancerPoolGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolGetResponseResult struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolGetResponseResultCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                      `json:"created_on" format:"date-time"`
	// A human-readable description of the pool.
	Description string `json:"description"`
	// This field shows up only if the pool is disabled. This field is set with the
	// time the pool was disabled at.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled bool `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude float64 `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding LoadBalancerPoolGetResponseResultLoadShedding `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude float64 `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins int64     `json:"minimum_origins"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor interface{} `json:"monitor"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name string `json:"name"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail string `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter LoadBalancerPoolGetResponseResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolGetResponseResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolGetResponseResultOrigin `json:"origins"`
	JSON    loadBalancerPoolGetResponseResultJSON     `json:"-"`
}

// loadBalancerPoolGetResponseResultJSON contains the JSON metadata for the struct
// [LoadBalancerPoolGetResponseResult]
type loadBalancerPoolGetResponseResultJSON struct {
	ID                 apijson.Field
	CheckRegions       apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	DisabledAt         apijson.Field
	Enabled            apijson.Field
	Latitude           apijson.Field
	LoadShedding       apijson.Field
	Longitude          apijson.Field
	MinimumOrigins     apijson.Field
	ModifiedOn         apijson.Field
	Monitor            apijson.Field
	Name               apijson.Field
	NotificationEmail  apijson.Field
	NotificationFilter apijson.Field
	OriginSteering     apijson.Field
	Origins            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolGetResponseResultCheckRegion string

const (
	LoadBalancerPoolGetResponseResultCheckRegionWnam       LoadBalancerPoolGetResponseResultCheckRegion = "WNAM"
	LoadBalancerPoolGetResponseResultCheckRegionEnam       LoadBalancerPoolGetResponseResultCheckRegion = "ENAM"
	LoadBalancerPoolGetResponseResultCheckRegionWeu        LoadBalancerPoolGetResponseResultCheckRegion = "WEU"
	LoadBalancerPoolGetResponseResultCheckRegionEeu        LoadBalancerPoolGetResponseResultCheckRegion = "EEU"
	LoadBalancerPoolGetResponseResultCheckRegionNsam       LoadBalancerPoolGetResponseResultCheckRegion = "NSAM"
	LoadBalancerPoolGetResponseResultCheckRegionSsam       LoadBalancerPoolGetResponseResultCheckRegion = "SSAM"
	LoadBalancerPoolGetResponseResultCheckRegionOc         LoadBalancerPoolGetResponseResultCheckRegion = "OC"
	LoadBalancerPoolGetResponseResultCheckRegionMe         LoadBalancerPoolGetResponseResultCheckRegion = "ME"
	LoadBalancerPoolGetResponseResultCheckRegionNaf        LoadBalancerPoolGetResponseResultCheckRegion = "NAF"
	LoadBalancerPoolGetResponseResultCheckRegionSaf        LoadBalancerPoolGetResponseResultCheckRegion = "SAF"
	LoadBalancerPoolGetResponseResultCheckRegionSas        LoadBalancerPoolGetResponseResultCheckRegion = "SAS"
	LoadBalancerPoolGetResponseResultCheckRegionSeas       LoadBalancerPoolGetResponseResultCheckRegion = "SEAS"
	LoadBalancerPoolGetResponseResultCheckRegionNeas       LoadBalancerPoolGetResponseResultCheckRegion = "NEAS"
	LoadBalancerPoolGetResponseResultCheckRegionAllRegions LoadBalancerPoolGetResponseResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolGetResponseResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolGetResponseResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolGetResponseResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolGetResponseResultLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolGetResponseResultLoadSheddingJSON contains the JSON metadata for
// the struct [LoadBalancerPoolGetResponseResultLoadShedding]
type loadBalancerPoolGetResponseResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolGetResponseResultLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolGetResponseResultLoadSheddingDefaultPolicyRandom LoadBalancerPoolGetResponseResultLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolGetResponseResultLoadSheddingDefaultPolicyHash   LoadBalancerPoolGetResponseResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolGetResponseResultLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolGetResponseResultLoadSheddingSessionPolicyHash LoadBalancerPoolGetResponseResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolGetResponseResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolGetResponseResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolGetResponseResultNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolGetResponseResultNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolGetResponseResultNotificationFilterJSON contains the JSON
// metadata for the struct [LoadBalancerPoolGetResponseResultNotificationFilter]
type loadBalancerPoolGetResponseResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolGetResponseResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                          `json:"healthy,nullable"`
	JSON    loadBalancerPoolGetResponseResultNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolGetResponseResultNotificationFilterOriginJSON contains the JSON
// metadata for the struct
// [LoadBalancerPoolGetResponseResultNotificationFilterOrigin]
type loadBalancerPoolGetResponseResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolGetResponseResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                        `json:"healthy,nullable"`
	JSON    loadBalancerPoolGetResponseResultNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolGetResponseResultNotificationFilterPoolJSON contains the JSON
// metadata for the struct
// [LoadBalancerPoolGetResponseResultNotificationFilterPool]
type loadBalancerPoolGetResponseResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolGetResponseResultOriginSteering struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy LoadBalancerPoolGetResponseResultOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolGetResponseResultOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolGetResponseResultOriginSteeringJSON contains the JSON metadata
// for the struct [LoadBalancerPoolGetResponseResultOriginSteering]
type loadBalancerPoolGetResponseResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of origin steering policy to use.
//
//   - `"random"`: Select an origin randomly.
//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
//     address.
//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
//     origin weights, as well as each origin's number of outstanding requests.
//     Origins with more pending requests are weighted proportionately less relative
//     to others.
//   - `"least_connections"`: Select an origin by taking into consideration origin
//     weights, as well as each origin's number of open connections. Origins with
//     more open connections are weighted proportionately less relative to others.
//     Supported for HTTP/1 and HTTP/2 connections.
type LoadBalancerPoolGetResponseResultOriginSteeringPolicy string

const (
	LoadBalancerPoolGetResponseResultOriginSteeringPolicyRandom                   LoadBalancerPoolGetResponseResultOriginSteeringPolicy = "random"
	LoadBalancerPoolGetResponseResultOriginSteeringPolicyHash                     LoadBalancerPoolGetResponseResultOriginSteeringPolicy = "hash"
	LoadBalancerPoolGetResponseResultOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolGetResponseResultOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolGetResponseResultOriginSteeringPolicyLeastConnections         LoadBalancerPoolGetResponseResultOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolGetResponseResultOrigin struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address string `json:"address"`
	// This field shows up only if the origin is disabled. This field is set with the
	// time the origin was disabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled bool `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header LoadBalancerPoolGetResponseResultOriginsHeader `json:"header"`
	// A human-identifiable name for the origin.
	Name string `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID string `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight float64                                     `json:"weight"`
	JSON   loadBalancerPoolGetResponseResultOriginJSON `json:"-"`
}

// loadBalancerPoolGetResponseResultOriginJSON contains the JSON metadata for the
// struct [LoadBalancerPoolGetResponseResultOrigin]
type loadBalancerPoolGetResponseResultOriginJSON struct {
	Address          apijson.Field
	DisabledAt       apijson.Field
	Enabled          apijson.Field
	Header           apijson.Field
	Name             apijson.Field
	VirtualNetworkID apijson.Field
	Weight           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolGetResponseResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                           `json:"Host"`
	JSON loadBalancerPoolGetResponseResultOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolGetResponseResultOriginsHeaderJSON contains the JSON metadata
// for the struct [LoadBalancerPoolGetResponseResultOriginsHeader]
type loadBalancerPoolGetResponseResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolGetResponseSuccess bool

const (
	LoadBalancerPoolGetResponseSuccessTrue LoadBalancerPoolGetResponseSuccess = true
)

type LoadBalancerPoolUpdateResponse struct {
	Errors   []LoadBalancerPoolUpdateResponseError   `json:"errors"`
	Messages []LoadBalancerPoolUpdateResponseMessage `json:"messages"`
	Result   LoadBalancerPoolUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerPoolUpdateResponseSuccess `json:"success"`
	JSON    loadBalancerPoolUpdateResponseJSON    `json:"-"`
}

// loadBalancerPoolUpdateResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolUpdateResponse]
type loadBalancerPoolUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    loadBalancerPoolUpdateResponseErrorJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseErrorJSON contains the JSON metadata for the
// struct [LoadBalancerPoolUpdateResponseError]
type loadBalancerPoolUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    loadBalancerPoolUpdateResponseMessageJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseMessageJSON contains the JSON metadata for the
// struct [LoadBalancerPoolUpdateResponseMessage]
type loadBalancerPoolUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolUpdateResponseResult struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolUpdateResponseResultCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                         `json:"created_on" format:"date-time"`
	// A human-readable description of the pool.
	Description string `json:"description"`
	// This field shows up only if the pool is disabled. This field is set with the
	// time the pool was disabled at.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled bool `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude float64 `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding LoadBalancerPoolUpdateResponseResultLoadShedding `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude float64 `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins int64     `json:"minimum_origins"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor interface{} `json:"monitor"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name string `json:"name"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail string `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter LoadBalancerPoolUpdateResponseResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolUpdateResponseResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolUpdateResponseResultOrigin `json:"origins"`
	JSON    loadBalancerPoolUpdateResponseResultJSON     `json:"-"`
}

// loadBalancerPoolUpdateResponseResultJSON contains the JSON metadata for the
// struct [LoadBalancerPoolUpdateResponseResult]
type loadBalancerPoolUpdateResponseResultJSON struct {
	ID                 apijson.Field
	CheckRegions       apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	DisabledAt         apijson.Field
	Enabled            apijson.Field
	Latitude           apijson.Field
	LoadShedding       apijson.Field
	Longitude          apijson.Field
	MinimumOrigins     apijson.Field
	ModifiedOn         apijson.Field
	Monitor            apijson.Field
	Name               apijson.Field
	NotificationEmail  apijson.Field
	NotificationFilter apijson.Field
	OriginSteering     apijson.Field
	Origins            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolUpdateResponseResultCheckRegion string

const (
	LoadBalancerPoolUpdateResponseResultCheckRegionWnam       LoadBalancerPoolUpdateResponseResultCheckRegion = "WNAM"
	LoadBalancerPoolUpdateResponseResultCheckRegionEnam       LoadBalancerPoolUpdateResponseResultCheckRegion = "ENAM"
	LoadBalancerPoolUpdateResponseResultCheckRegionWeu        LoadBalancerPoolUpdateResponseResultCheckRegion = "WEU"
	LoadBalancerPoolUpdateResponseResultCheckRegionEeu        LoadBalancerPoolUpdateResponseResultCheckRegion = "EEU"
	LoadBalancerPoolUpdateResponseResultCheckRegionNsam       LoadBalancerPoolUpdateResponseResultCheckRegion = "NSAM"
	LoadBalancerPoolUpdateResponseResultCheckRegionSsam       LoadBalancerPoolUpdateResponseResultCheckRegion = "SSAM"
	LoadBalancerPoolUpdateResponseResultCheckRegionOc         LoadBalancerPoolUpdateResponseResultCheckRegion = "OC"
	LoadBalancerPoolUpdateResponseResultCheckRegionMe         LoadBalancerPoolUpdateResponseResultCheckRegion = "ME"
	LoadBalancerPoolUpdateResponseResultCheckRegionNaf        LoadBalancerPoolUpdateResponseResultCheckRegion = "NAF"
	LoadBalancerPoolUpdateResponseResultCheckRegionSaf        LoadBalancerPoolUpdateResponseResultCheckRegion = "SAF"
	LoadBalancerPoolUpdateResponseResultCheckRegionSas        LoadBalancerPoolUpdateResponseResultCheckRegion = "SAS"
	LoadBalancerPoolUpdateResponseResultCheckRegionSeas       LoadBalancerPoolUpdateResponseResultCheckRegion = "SEAS"
	LoadBalancerPoolUpdateResponseResultCheckRegionNeas       LoadBalancerPoolUpdateResponseResultCheckRegion = "NEAS"
	LoadBalancerPoolUpdateResponseResultCheckRegionAllRegions LoadBalancerPoolUpdateResponseResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolUpdateResponseResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolUpdateResponseResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolUpdateResponseResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolUpdateResponseResultLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolUpdateResponseResultLoadSheddingJSON contains the JSON metadata
// for the struct [LoadBalancerPoolUpdateResponseResultLoadShedding]
type loadBalancerPoolUpdateResponseResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolUpdateResponseResultLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolUpdateResponseResultLoadSheddingDefaultPolicyRandom LoadBalancerPoolUpdateResponseResultLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolUpdateResponseResultLoadSheddingDefaultPolicyHash   LoadBalancerPoolUpdateResponseResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolUpdateResponseResultLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolUpdateResponseResultLoadSheddingSessionPolicyHash LoadBalancerPoolUpdateResponseResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolUpdateResponseResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolUpdateResponseResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolUpdateResponseResultNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolUpdateResponseResultNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseResultNotificationFilterJSON contains the JSON
// metadata for the struct [LoadBalancerPoolUpdateResponseResultNotificationFilter]
type loadBalancerPoolUpdateResponseResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolUpdateResponseResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                             `json:"healthy,nullable"`
	JSON    loadBalancerPoolUpdateResponseResultNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseResultNotificationFilterOriginJSON contains the
// JSON metadata for the struct
// [LoadBalancerPoolUpdateResponseResultNotificationFilterOrigin]
type loadBalancerPoolUpdateResponseResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolUpdateResponseResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                           `json:"healthy,nullable"`
	JSON    loadBalancerPoolUpdateResponseResultNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseResultNotificationFilterPoolJSON contains the JSON
// metadata for the struct
// [LoadBalancerPoolUpdateResponseResultNotificationFilterPool]
type loadBalancerPoolUpdateResponseResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolUpdateResponseResultOriginSteering struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy LoadBalancerPoolUpdateResponseResultOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolUpdateResponseResultOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolUpdateResponseResultOriginSteeringJSON contains the JSON
// metadata for the struct [LoadBalancerPoolUpdateResponseResultOriginSteering]
type loadBalancerPoolUpdateResponseResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of origin steering policy to use.
//
//   - `"random"`: Select an origin randomly.
//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
//     address.
//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
//     origin weights, as well as each origin's number of outstanding requests.
//     Origins with more pending requests are weighted proportionately less relative
//     to others.
//   - `"least_connections"`: Select an origin by taking into consideration origin
//     weights, as well as each origin's number of open connections. Origins with
//     more open connections are weighted proportionately less relative to others.
//     Supported for HTTP/1 and HTTP/2 connections.
type LoadBalancerPoolUpdateResponseResultOriginSteeringPolicy string

const (
	LoadBalancerPoolUpdateResponseResultOriginSteeringPolicyRandom                   LoadBalancerPoolUpdateResponseResultOriginSteeringPolicy = "random"
	LoadBalancerPoolUpdateResponseResultOriginSteeringPolicyHash                     LoadBalancerPoolUpdateResponseResultOriginSteeringPolicy = "hash"
	LoadBalancerPoolUpdateResponseResultOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolUpdateResponseResultOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolUpdateResponseResultOriginSteeringPolicyLeastConnections         LoadBalancerPoolUpdateResponseResultOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolUpdateResponseResultOrigin struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address string `json:"address"`
	// This field shows up only if the origin is disabled. This field is set with the
	// time the origin was disabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled bool `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header LoadBalancerPoolUpdateResponseResultOriginsHeader `json:"header"`
	// A human-identifiable name for the origin.
	Name string `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID string `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight float64                                        `json:"weight"`
	JSON   loadBalancerPoolUpdateResponseResultOriginJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseResultOriginJSON contains the JSON metadata for
// the struct [LoadBalancerPoolUpdateResponseResultOrigin]
type loadBalancerPoolUpdateResponseResultOriginJSON struct {
	Address          apijson.Field
	DisabledAt       apijson.Field
	Enabled          apijson.Field
	Header           apijson.Field
	Name             apijson.Field
	VirtualNetworkID apijson.Field
	Weight           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolUpdateResponseResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                              `json:"Host"`
	JSON loadBalancerPoolUpdateResponseResultOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseResultOriginsHeaderJSON contains the JSON metadata
// for the struct [LoadBalancerPoolUpdateResponseResultOriginsHeader]
type loadBalancerPoolUpdateResponseResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolUpdateResponseSuccess bool

const (
	LoadBalancerPoolUpdateResponseSuccessTrue LoadBalancerPoolUpdateResponseSuccess = true
)

type LoadBalancerPoolDeleteResponse struct {
	Errors   []LoadBalancerPoolDeleteResponseError   `json:"errors"`
	Messages []LoadBalancerPoolDeleteResponseMessage `json:"messages"`
	Result   LoadBalancerPoolDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerPoolDeleteResponseSuccess `json:"success"`
	JSON    loadBalancerPoolDeleteResponseJSON    `json:"-"`
}

// loadBalancerPoolDeleteResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolDeleteResponse]
type loadBalancerPoolDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    loadBalancerPoolDeleteResponseErrorJSON `json:"-"`
}

// loadBalancerPoolDeleteResponseErrorJSON contains the JSON metadata for the
// struct [LoadBalancerPoolDeleteResponseError]
type loadBalancerPoolDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    loadBalancerPoolDeleteResponseMessageJSON `json:"-"`
}

// loadBalancerPoolDeleteResponseMessageJSON contains the JSON metadata for the
// struct [LoadBalancerPoolDeleteResponseMessage]
type loadBalancerPoolDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolDeleteResponseResult struct {
	ID   string                                   `json:"id"`
	JSON loadBalancerPoolDeleteResponseResultJSON `json:"-"`
}

// loadBalancerPoolDeleteResponseResultJSON contains the JSON metadata for the
// struct [LoadBalancerPoolDeleteResponseResult]
type loadBalancerPoolDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolDeleteResponseSuccess bool

const (
	LoadBalancerPoolDeleteResponseSuccessTrue LoadBalancerPoolDeleteResponseSuccess = true
)

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse struct {
	Errors   []LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseError   `json:"errors"`
	Messages []LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseMessage `json:"messages"`
	Result   LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseSuccess `json:"success"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseJSON    `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseJSON contains the JSON
// metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseError struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseErrorJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseErrorJSON contains the
// JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseError]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseMessage struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseMessageJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseMessageJSON contains the
// JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseMessage]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResult struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                                  `json:"created_on" format:"date-time"`
	// A human-readable description of the pool.
	Description string `json:"description"`
	// This field shows up only if the pool is disabled. This field is set with the
	// time the pool was disabled at.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled bool `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude float64 `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadShedding `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude float64 `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins int64     `json:"minimum_origins"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor interface{} `json:"monitor"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name string `json:"name"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail string `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOrigin `json:"origins"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultJSON     `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultJSON contains the
// JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResult]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultJSON struct {
	ID                 apijson.Field
	CheckRegions       apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	DisabledAt         apijson.Field
	Enabled            apijson.Field
	Latitude           apijson.Field
	LoadShedding       apijson.Field
	Longitude          apijson.Field
	MinimumOrigins     apijson.Field
	ModifiedOn         apijson.Field
	Monitor            apijson.Field
	Name               apijson.Field
	NotificationEmail  apijson.Field
	NotificationFilter apijson.Field
	OriginSteering     apijson.Field
	Origins            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionWnam       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "WNAM"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionEnam       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "ENAM"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionWeu        LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "WEU"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionEeu        LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "EEU"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionNsam       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "NSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionSsam       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "SSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionOc         LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "OC"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionMe         LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "ME"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionNaf        LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "NAF"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionSaf        LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "SAF"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionSas        LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "SAS"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionSeas       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "SEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionNeas       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "NEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegionAllRegions LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadShedding]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingDefaultPolicyRandom LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingDefaultPolicyHash   LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingSessionPolicyHash LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilter]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                      `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                    `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterPool]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteering struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteering]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of origin steering policy to use.
//
//   - `"random"`: Select an origin randomly.
//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
//     address.
//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
//     origin weights, as well as each origin's number of outstanding requests.
//     Origins with more pending requests are weighted proportionately less relative
//     to others.
//   - `"least_connections"`: Select an origin by taking into consideration origin
//     weights, as well as each origin's number of open connections. Origins with
//     more open connections are weighted proportionately less relative to others.
//     Supported for HTTP/1 and HTTP/2 connections.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringPolicyRandom                   LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringPolicyHash                     LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringPolicy = "hash"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringPolicyLeastConnections         LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOrigin struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address string `json:"address"`
	// This field shows up only if the origin is disabled. This field is set with the
	// time the origin was disabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled bool `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginsHeader `json:"header"`
	// A human-identifiable name for the origin.
	Name string `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID string `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight float64                                                                 `json:"weight"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginJSON contains
// the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginJSON struct {
	Address          apijson.Field
	DisabledAt       apijson.Field
	Enabled          apijson.Field
	Header           apijson.Field
	Name             apijson.Field
	VirtualNetworkID apijson.Field
	Weight           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                       `json:"Host"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginsHeaderJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginsHeader]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseSuccess bool

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseSuccessTrue LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseSuccess = true
)

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse struct {
	Errors     []LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseError    `json:"errors"`
	Messages   []LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessage  `json:"messages"`
	Result     []LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResult   `json:"result"`
	ResultInfo LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseSuccess `json:"success"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseJSON    `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseJSON contains the JSON
// metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseErrorJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseErrorJSON contains the
// JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseError]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessageJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessageJSON contains
// the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessage]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResult struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                                    `json:"created_on" format:"date-time"`
	// A human-readable description of the pool.
	Description string `json:"description"`
	// This field shows up only if the pool is disabled. This field is set with the
	// time the pool was disabled at.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled bool `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude float64 `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadShedding `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude float64 `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins int64     `json:"minimum_origins"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor interface{} `json:"monitor"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name string `json:"name"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail string `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOrigin `json:"origins"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultJSON     `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultJSON contains the
// JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResult]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultJSON struct {
	ID                 apijson.Field
	CheckRegions       apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	DisabledAt         apijson.Field
	Enabled            apijson.Field
	Latitude           apijson.Field
	LoadShedding       apijson.Field
	Longitude          apijson.Field
	MinimumOrigins     apijson.Field
	ModifiedOn         apijson.Field
	Monitor            apijson.Field
	Name               apijson.Field
	NotificationEmail  apijson.Field
	NotificationFilter apijson.Field
	OriginSteering     apijson.Field
	Origins            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionWnam       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "WNAM"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionEnam       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "ENAM"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionWeu        LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "WEU"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionEeu        LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "EEU"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionNsam       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "NSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionSsam       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionOc         LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "OC"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionMe         LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "ME"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionNaf        LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "NAF"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionSaf        LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SAF"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionSas        LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SAS"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionSeas       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionNeas       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "NEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionAllRegions LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadShedding]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicyRandom LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicyHash   LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicyHash LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilter]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                        `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                      `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteering struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteering]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of origin steering policy to use.
//
//   - `"random"`: Select an origin randomly.
//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
//     address.
//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
//     origin weights, as well as each origin's number of outstanding requests.
//     Origins with more pending requests are weighted proportionately less relative
//     to others.
//   - `"least_connections"`: Select an origin by taking into consideration origin
//     weights, as well as each origin's number of open connections. Origins with
//     more open connections are weighted proportionately less relative to others.
//     Supported for HTTP/1 and HTTP/2 connections.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyRandom                   LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyHash                     LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "hash"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyLeastConnections         LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOrigin struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address string `json:"address"`
	// This field shows up only if the origin is disabled. This field is set with the
	// time the origin was disabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled bool `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeader `json:"header"`
	// A human-identifiable name for the origin.
	Name string `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID string `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight float64                                                                   `json:"weight"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginJSON struct {
	Address          apijson.Field
	DisabledAt       apijson.Field
	Enabled          apijson.Field
	Header           apijson.Field
	Name             apijson.Field
	VirtualNetworkID apijson.Field
	Weight           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                         `json:"Host"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeaderJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeader]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfoJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfoJSON contains
// the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfo]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseSuccess bool

const (
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseSuccessTrue LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseSuccess = true
)

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse struct {
	Errors     []LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseError    `json:"errors"`
	Messages   []LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessage  `json:"messages"`
	Result     []LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResult   `json:"result"`
	ResultInfo LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseSuccess `json:"success"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseJSON    `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseJSON contains the JSON
// metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseErrorJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseErrorJSON contains the
// JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseError]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessageJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessageJSON contains
// the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessage]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResult struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                                     `json:"created_on" format:"date-time"`
	// A human-readable description of the pool.
	Description string `json:"description"`
	// This field shows up only if the pool is disabled. This field is set with the
	// time the pool was disabled at.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled bool `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude float64 `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadShedding `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude float64 `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins int64     `json:"minimum_origins"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor interface{} `json:"monitor"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name string `json:"name"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail string `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOrigin `json:"origins"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultJSON     `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultJSON contains
// the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResult]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultJSON struct {
	ID                 apijson.Field
	CheckRegions       apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	DisabledAt         apijson.Field
	Enabled            apijson.Field
	Latitude           apijson.Field
	LoadShedding       apijson.Field
	Longitude          apijson.Field
	MinimumOrigins     apijson.Field
	ModifiedOn         apijson.Field
	Monitor            apijson.Field
	Name               apijson.Field
	NotificationEmail  apijson.Field
	NotificationFilter apijson.Field
	OriginSteering     apijson.Field
	Origins            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionWnam       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "WNAM"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionEnam       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "ENAM"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionWeu        LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "WEU"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionEeu        LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "EEU"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionNsam       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "NSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSsam       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionOc         LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "OC"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionMe         LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "ME"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionNaf        LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "NAF"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSaf        LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SAF"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSas        LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SAS"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSeas       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionNeas       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "NEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionAllRegions LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadShedding]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicyRandom LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicyHash   LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicyHash LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                         `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                       `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteering struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteering]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of origin steering policy to use.
//
//   - `"random"`: Select an origin randomly.
//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
//     address.
//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
//     origin weights, as well as each origin's number of outstanding requests.
//     Origins with more pending requests are weighted proportionately less relative
//     to others.
//   - `"least_connections"`: Select an origin by taking into consideration origin
//     weights, as well as each origin's number of open connections. Origins with
//     more open connections are weighted proportionately less relative to others.
//     Supported for HTTP/1 and HTTP/2 connections.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyRandom                   LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyHash                     LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "hash"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyLeastConnections         LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOrigin struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address string `json:"address"`
	// This field shows up only if the origin is disabled. This field is set with the
	// time the origin was disabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled bool `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader `json:"header"`
	// A human-identifiable name for the origin.
	Name string `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID string `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight float64                                                                    `json:"weight"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginJSON struct {
	Address          apijson.Field
	DisabledAt       apijson.Field
	Enabled          apijson.Field
	Header           apijson.Field
	Name             apijson.Field
	VirtualNetworkID apijson.Field
	Weight           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                          `json:"Host"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeaderJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                  `json:"total_count"`
	JSON       loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfoJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfo]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseSuccess bool

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseSuccessTrue LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseSuccess = true
)

type LoadBalancerPoolUpdateParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]LoadBalancerPoolUpdateParamsOrigin] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]LoadBalancerPoolUpdateParamsCheckRegion] `json:"check_regions"`
	// A human-readable description of the pool.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled param.Field[bool] `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude param.Field[float64] `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding param.Field[LoadBalancerPoolUpdateParamsLoadShedding] `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude param.Field[float64] `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins param.Field[int64] `json:"minimum_origins"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `json:"monitor"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail param.Field[string] `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter param.Field[LoadBalancerPoolUpdateParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[LoadBalancerPoolUpdateParamsOriginSteering] `json:"origin_steering"`
}

func (r LoadBalancerPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerPoolUpdateParamsOrigin struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address param.Field[string] `json:"address"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled param.Field[bool] `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header param.Field[LoadBalancerPoolUpdateParamsOriginsHeader] `json:"header"`
	// A human-identifiable name for the origin.
	Name param.Field[string] `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight param.Field[float64] `json:"weight"`
}

func (r LoadBalancerPoolUpdateParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolUpdateParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r LoadBalancerPoolUpdateParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolUpdateParamsCheckRegion string

const (
	LoadBalancerPoolUpdateParamsCheckRegionWnam       LoadBalancerPoolUpdateParamsCheckRegion = "WNAM"
	LoadBalancerPoolUpdateParamsCheckRegionEnam       LoadBalancerPoolUpdateParamsCheckRegion = "ENAM"
	LoadBalancerPoolUpdateParamsCheckRegionWeu        LoadBalancerPoolUpdateParamsCheckRegion = "WEU"
	LoadBalancerPoolUpdateParamsCheckRegionEeu        LoadBalancerPoolUpdateParamsCheckRegion = "EEU"
	LoadBalancerPoolUpdateParamsCheckRegionNsam       LoadBalancerPoolUpdateParamsCheckRegion = "NSAM"
	LoadBalancerPoolUpdateParamsCheckRegionSsam       LoadBalancerPoolUpdateParamsCheckRegion = "SSAM"
	LoadBalancerPoolUpdateParamsCheckRegionOc         LoadBalancerPoolUpdateParamsCheckRegion = "OC"
	LoadBalancerPoolUpdateParamsCheckRegionMe         LoadBalancerPoolUpdateParamsCheckRegion = "ME"
	LoadBalancerPoolUpdateParamsCheckRegionNaf        LoadBalancerPoolUpdateParamsCheckRegion = "NAF"
	LoadBalancerPoolUpdateParamsCheckRegionSaf        LoadBalancerPoolUpdateParamsCheckRegion = "SAF"
	LoadBalancerPoolUpdateParamsCheckRegionSas        LoadBalancerPoolUpdateParamsCheckRegion = "SAS"
	LoadBalancerPoolUpdateParamsCheckRegionSeas       LoadBalancerPoolUpdateParamsCheckRegion = "SEAS"
	LoadBalancerPoolUpdateParamsCheckRegionNeas       LoadBalancerPoolUpdateParamsCheckRegion = "NEAS"
	LoadBalancerPoolUpdateParamsCheckRegionAllRegions LoadBalancerPoolUpdateParamsCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolUpdateParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[LoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[LoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r LoadBalancerPoolUpdateParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyRandom LoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyHash   LoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolUpdateParamsLoadSheddingSessionPolicyHash LoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolUpdateParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[LoadBalancerPoolUpdateParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[LoadBalancerPoolUpdateParamsNotificationFilterPool] `json:"pool"`
}

func (r LoadBalancerPoolUpdateParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolUpdateParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r LoadBalancerPoolUpdateParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolUpdateParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r LoadBalancerPoolUpdateParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolUpdateParamsOriginSteering struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy param.Field[LoadBalancerPoolUpdateParamsOriginSteeringPolicy] `json:"policy"`
}

func (r LoadBalancerPoolUpdateParamsOriginSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of origin steering policy to use.
//
//   - `"random"`: Select an origin randomly.
//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
//     address.
//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
//     origin weights, as well as each origin's number of outstanding requests.
//     Origins with more pending requests are weighted proportionately less relative
//     to others.
//   - `"least_connections"`: Select an origin by taking into consideration origin
//     weights, as well as each origin's number of open connections. Origins with
//     more open connections are weighted proportionately less relative to others.
//     Supported for HTTP/1 and HTTP/2 connections.
type LoadBalancerPoolUpdateParamsOriginSteeringPolicy string

const (
	LoadBalancerPoolUpdateParamsOriginSteeringPolicyRandom                   LoadBalancerPoolUpdateParamsOriginSteeringPolicy = "random"
	LoadBalancerPoolUpdateParamsOriginSteeringPolicyHash                     LoadBalancerPoolUpdateParamsOriginSteeringPolicy = "hash"
	LoadBalancerPoolUpdateParamsOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolUpdateParamsOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolUpdateParamsOriginSteeringPolicyLeastConnections         LoadBalancerPoolUpdateParamsOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOrigin] `json:"origins,required"`
	// A human-readable description of the pool.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled param.Field[bool] `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude param.Field[float64] `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding param.Field[LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadShedding] `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude param.Field[float64] `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins param.Field[int64] `json:"minimum_origins"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `json:"monitor"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail param.Field[string] `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter param.Field[LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteering] `json:"origin_steering"`
}

func (r LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOrigin struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address param.Field[string] `json:"address"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled param.Field[bool] `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header param.Field[LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginsHeader] `json:"header"`
	// A human-identifiable name for the origin.
	Name param.Field[string] `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight param.Field[float64] `json:"weight"`
}

func (r LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicyRandom LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicyHash   LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicyHash LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterPool] `json:"pool"`
}

func (r LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteering struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy param.Field[LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy] `json:"policy"`
}

func (r LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of origin steering policy to use.
//
//   - `"random"`: Select an origin randomly.
//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
//     address.
//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
//     origin weights, as well as each origin's number of outstanding requests.
//     Origins with more pending requests are weighted proportionately less relative
//     to others.
//   - `"least_connections"`: Select an origin by taking into consideration origin
//     weights, as well as each origin's number of open connections. Origins with
//     more open connections are weighted proportionately less relative to others.
//     Supported for HTTP/1 and HTTP/2 connections.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyRandom                   LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyHash                     LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "hash"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyLeastConnections         LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams struct {
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `query:"monitor"`
}

// URLQuery serializes [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams]'s
// query parameters as `url.Values`.
func (r LoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams struct {
	// The email address to send health status notifications to. This field is now
	// deprecated in favor of Cloudflare Notifications for Load Balancing, so only
	// resetting this field with an empty string `""` is accepted.
	NotificationEmail param.Field[LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParamsNotificationEmail] `json:"notification_email"`
}

func (r LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The email address to send health status notifications to. This field is now
// deprecated in favor of Cloudflare Notifications for Load Balancing, so only
// resetting this field with an empty string `""` is accepted.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParamsNotificationEmail string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParamsNotificationEmailEmpty LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParamsNotificationEmail = "\"\""
)
