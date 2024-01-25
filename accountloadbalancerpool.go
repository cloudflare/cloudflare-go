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

// AccountLoadBalancerPoolService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLoadBalancerPoolService] method instead.
type AccountLoadBalancerPoolService struct {
	Options    []option.RequestOption
	Health     *AccountLoadBalancerPoolHealthService
	Previews   *AccountLoadBalancerPoolPreviewService
	References *AccountLoadBalancerPoolReferenceService
}

// NewAccountLoadBalancerPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerPoolService(opts ...option.RequestOption) (r *AccountLoadBalancerPoolService) {
	r = &AccountLoadBalancerPoolService{}
	r.Options = opts
	r.Health = NewAccountLoadBalancerPoolHealthService(opts...)
	r.Previews = NewAccountLoadBalancerPoolPreviewService(opts...)
	r.References = NewAccountLoadBalancerPoolReferenceService(opts...)
	return
}

// Fetch a single configured pool.
func (r *AccountLoadBalancerPoolService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured pool.
func (r *AccountLoadBalancerPoolService) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountLoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured pool.
func (r *AccountLoadBalancerPoolService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountLoadBalancerPoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new pool.
func (r *AccountLoadBalancerPoolService) AccountLoadBalancerPoolsNewPool(ctx context.Context, accountIdentifier string, body AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured pools.
func (r *AccountLoadBalancerPoolService) AccountLoadBalancerPoolsListPools(ctx context.Context, accountIdentifier string, query AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams, opts ...option.RequestOption) (res *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Apply changes to a number of existing pools, overwriting the supplied
// properties. Pools are ordered by ascending `name`. Returns the list of affected
// pools. Supports the standard pagination query parameters, either
// `limit`/`offset` or `per_page`/`page`.
func (r *AccountLoadBalancerPoolService) AccountLoadBalancerPoolsPatchPools(ctx context.Context, accountIdentifier string, body AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams, opts ...option.RequestOption) (res *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Apply changes to an existing pool, overwriting the supplied properties.
func (r *AccountLoadBalancerPoolService) Patch(ctx context.Context, accountIdentifier string, identifier string, body AccountLoadBalancerPoolPatchParams, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type Pool struct {
	Errors   []PoolError   `json:"errors"`
	Messages []PoolMessage `json:"messages"`
	Result   PoolResult    `json:"result"`
	// Whether the API call was successful
	Success PoolSuccess `json:"success"`
	JSON    poolJSON    `json:"-"`
}

// poolJSON contains the JSON metadata for the struct [Pool]
type poolJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoolError struct {
	Code    int64         `json:"code,required"`
	Message string        `json:"message,required"`
	JSON    poolErrorJSON `json:"-"`
}

// poolErrorJSON contains the JSON metadata for the struct [PoolError]
type poolErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoolMessage struct {
	Code    int64           `json:"code,required"`
	Message string          `json:"message,required"`
	JSON    poolMessageJSON `json:"-"`
}

// poolMessageJSON contains the JSON metadata for the struct [PoolMessage]
type poolMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoolResult struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []PoolResultCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time               `json:"created_on" format:"date-time"`
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
	LoadShedding PoolResultLoadShedding `json:"load_shedding"`
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
	NotificationFilter PoolResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering PoolResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []PoolResultOrigin `json:"origins"`
	JSON    poolResultJSON     `json:"-"`
}

// poolResultJSON contains the JSON metadata for the struct [PoolResult]
type poolResultJSON struct {
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

func (r *PoolResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolResultCheckRegion string

const (
	PoolResultCheckRegionWnam       PoolResultCheckRegion = "WNAM"
	PoolResultCheckRegionEnam       PoolResultCheckRegion = "ENAM"
	PoolResultCheckRegionWeu        PoolResultCheckRegion = "WEU"
	PoolResultCheckRegionEeu        PoolResultCheckRegion = "EEU"
	PoolResultCheckRegionNsam       PoolResultCheckRegion = "NSAM"
	PoolResultCheckRegionSsam       PoolResultCheckRegion = "SSAM"
	PoolResultCheckRegionOc         PoolResultCheckRegion = "OC"
	PoolResultCheckRegionMe         PoolResultCheckRegion = "ME"
	PoolResultCheckRegionNaf        PoolResultCheckRegion = "NAF"
	PoolResultCheckRegionSaf        PoolResultCheckRegion = "SAF"
	PoolResultCheckRegionSas        PoolResultCheckRegion = "SAS"
	PoolResultCheckRegionSeas       PoolResultCheckRegion = "SEAS"
	PoolResultCheckRegionNeas       PoolResultCheckRegion = "NEAS"
	PoolResultCheckRegionAllRegions PoolResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type PoolResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy PoolResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy PoolResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          poolResultLoadSheddingJSON          `json:"-"`
}

// poolResultLoadSheddingJSON contains the JSON metadata for the struct
// [PoolResultLoadShedding]
type poolResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PoolResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolResultLoadSheddingDefaultPolicy string

const (
	PoolResultLoadSheddingDefaultPolicyRandom PoolResultLoadSheddingDefaultPolicy = "random"
	PoolResultLoadSheddingDefaultPolicyHash   PoolResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolResultLoadSheddingSessionPolicy string

const (
	PoolResultLoadSheddingSessionPolicyHash PoolResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin PoolResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool PoolResultNotificationFilterPool `json:"pool,nullable"`
	JSON poolResultNotificationFilterJSON `json:"-"`
}

// poolResultNotificationFilterJSON contains the JSON metadata for the struct
// [PoolResultNotificationFilter]
type poolResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                   `json:"healthy,nullable"`
	JSON    poolResultNotificationFilterOriginJSON `json:"-"`
}

// poolResultNotificationFilterOriginJSON contains the JSON metadata for the struct
// [PoolResultNotificationFilterOrigin]
type poolResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                 `json:"healthy,nullable"`
	JSON    poolResultNotificationFilterPoolJSON `json:"-"`
}

// poolResultNotificationFilterPoolJSON contains the JSON metadata for the struct
// [PoolResultNotificationFilterPool]
type poolResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolResultOriginSteering struct {
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
	Policy PoolResultOriginSteeringPolicy `json:"policy"`
	JSON   poolResultOriginSteeringJSON   `json:"-"`
}

// poolResultOriginSteeringJSON contains the JSON metadata for the struct
// [PoolResultOriginSteering]
type poolResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type PoolResultOriginSteeringPolicy string

const (
	PoolResultOriginSteeringPolicyRandom                   PoolResultOriginSteeringPolicy = "random"
	PoolResultOriginSteeringPolicyHash                     PoolResultOriginSteeringPolicy = "hash"
	PoolResultOriginSteeringPolicyLeastOutstandingRequests PoolResultOriginSteeringPolicy = "least_outstanding_requests"
	PoolResultOriginSteeringPolicyLeastConnections         PoolResultOriginSteeringPolicy = "least_connections"
)

type PoolResultOrigin struct {
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
	Header PoolResultOriginsHeader `json:"header"`
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
	Weight float64              `json:"weight"`
	JSON   poolResultOriginJSON `json:"-"`
}

// poolResultOriginJSON contains the JSON metadata for the struct
// [PoolResultOrigin]
type poolResultOriginJSON struct {
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

func (r *PoolResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                    `json:"Host"`
	JSON poolResultOriginsHeaderJSON `json:"-"`
}

// poolResultOriginsHeaderJSON contains the JSON metadata for the struct
// [PoolResultOriginsHeader]
type poolResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PoolSuccess bool

const (
	PoolSuccessTrue PoolSuccess = true
)

type AccountLoadBalancerPoolDeleteResponse struct {
	Errors   []AccountLoadBalancerPoolDeleteResponseError   `json:"errors"`
	Messages []AccountLoadBalancerPoolDeleteResponseMessage `json:"messages"`
	Result   AccountLoadBalancerPoolDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountLoadBalancerPoolDeleteResponseSuccess `json:"success"`
	JSON    accountLoadBalancerPoolDeleteResponseJSON    `json:"-"`
}

// accountLoadBalancerPoolDeleteResponseJSON contains the JSON metadata for the
// struct [AccountLoadBalancerPoolDeleteResponse]
type accountLoadBalancerPoolDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountLoadBalancerPoolDeleteResponseErrorJSON `json:"-"`
}

// accountLoadBalancerPoolDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountLoadBalancerPoolDeleteResponseError]
type accountLoadBalancerPoolDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountLoadBalancerPoolDeleteResponseMessageJSON `json:"-"`
}

// accountLoadBalancerPoolDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountLoadBalancerPoolDeleteResponseMessage]
type accountLoadBalancerPoolDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolDeleteResponseResult struct {
	ID   string                                          `json:"id"`
	JSON accountLoadBalancerPoolDeleteResponseResultJSON `json:"-"`
}

// accountLoadBalancerPoolDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountLoadBalancerPoolDeleteResponseResult]
type accountLoadBalancerPoolDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerPoolDeleteResponseSuccess bool

const (
	AccountLoadBalancerPoolDeleteResponseSuccessTrue AccountLoadBalancerPoolDeleteResponseSuccess = true
)

type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse struct {
	Errors     []AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseError    `json:"errors"`
	Messages   []AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessage  `json:"messages"`
	Result     []AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResult   `json:"result"`
	ResultInfo AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseSuccess `json:"success"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseJSON    `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseJSON contains
// the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseErrorJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseError]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessageJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessage]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResult struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                                           `json:"created_on" format:"date-time"`
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
	LoadShedding AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadShedding `json:"load_shedding"`
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
	NotificationFilter AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOrigin `json:"origins"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultJSON     `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResult]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultJSON struct {
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

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionWnam       AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "WNAM"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionEnam       AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "ENAM"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionWeu        AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "WEU"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionEeu        AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "EEU"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionNsam       AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "NSAM"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionSsam       AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SSAM"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionOc         AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "OC"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionMe         AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "ME"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionNaf        AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "NAF"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionSaf        AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SAF"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionSas        AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SAS"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionSeas       AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SEAS"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionNeas       AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "NEAS"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegionAllRegions AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingJSON          `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadShedding]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicyRandom AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy = "random"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicyHash   AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicy string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicyHash AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool `json:"pool,nullable"`
	JSON accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilter]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                               `json:"healthy,nullable"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOriginJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                             `json:"healthy,nullable"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPoolJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteering struct {
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
	Policy AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy `json:"policy"`
	JSON   accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringJSON   `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteering]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyRandom                   AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "random"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyHash                     AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "hash"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyLeastOutstandingRequests AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "least_outstanding_requests"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyLeastConnections         AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "least_connections"
)

type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOrigin struct {
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
	Header AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeader `json:"header"`
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
	Weight float64                                                                          `json:"weight"`
	JSON   accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOrigin]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginJSON struct {
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

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                                `json:"Host"`
	JSON accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeaderJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeaderJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeader]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                        `json:"total_count"`
	JSON       accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfoJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfo]
type accountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseSuccess bool

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseSuccessTrue AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseSuccess = true
)

type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse struct {
	Errors     []AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseError    `json:"errors"`
	Messages   []AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessage  `json:"messages"`
	Result     []AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResult   `json:"result"`
	ResultInfo AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseSuccess `json:"success"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseJSON    `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseJSON contains
// the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseErrorJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseError]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessageJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessage]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResult struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                                            `json:"created_on" format:"date-time"`
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
	LoadShedding AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadShedding `json:"load_shedding"`
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
	NotificationFilter AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOrigin `json:"origins"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultJSON     `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResult]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultJSON struct {
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

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionWnam       AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "WNAM"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionEnam       AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "ENAM"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionWeu        AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "WEU"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionEeu        AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "EEU"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionNsam       AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "NSAM"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSsam       AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SSAM"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionOc         AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "OC"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionMe         AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "ME"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionNaf        AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "NAF"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSaf        AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SAF"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSas        AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SAS"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSeas       AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SEAS"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionNeas       AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "NEAS"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegionAllRegions AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingJSON          `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadShedding]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicyRandom AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy = "random"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicyHash   AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicy string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicyHash AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool `json:"pool,nullable"`
	JSON accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                                `json:"healthy,nullable"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOriginJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                              `json:"healthy,nullable"`
	JSON    accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPoolJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteering struct {
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
	Policy AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy `json:"policy"`
	JSON   accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringJSON   `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteering]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyRandom                   AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "random"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyHash                     AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "hash"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyLeastOutstandingRequests AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "least_outstanding_requests"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyLeastConnections         AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "least_connections"
)

type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOrigin struct {
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
	Header AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader `json:"header"`
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
	Weight float64                                                                           `json:"weight"`
	JSON   accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOrigin]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginJSON struct {
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

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                                 `json:"Host"`
	JSON accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeaderJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeaderJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfoJSON `json:"-"`
}

// accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfo]
type accountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseSuccess bool

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseSuccessTrue AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseSuccess = true
)

type AccountLoadBalancerPoolUpdateParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]AccountLoadBalancerPoolUpdateParamsOrigin] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]AccountLoadBalancerPoolUpdateParamsCheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[AccountLoadBalancerPoolUpdateParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[AccountLoadBalancerPoolUpdateParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[AccountLoadBalancerPoolUpdateParamsOriginSteering] `json:"origin_steering"`
}

func (r AccountLoadBalancerPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountLoadBalancerPoolUpdateParamsOrigin struct {
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
	Header param.Field[AccountLoadBalancerPoolUpdateParamsOriginsHeader] `json:"header"`
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

func (r AccountLoadBalancerPoolUpdateParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type AccountLoadBalancerPoolUpdateParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r AccountLoadBalancerPoolUpdateParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type AccountLoadBalancerPoolUpdateParamsCheckRegion string

const (
	AccountLoadBalancerPoolUpdateParamsCheckRegionWnam       AccountLoadBalancerPoolUpdateParamsCheckRegion = "WNAM"
	AccountLoadBalancerPoolUpdateParamsCheckRegionEnam       AccountLoadBalancerPoolUpdateParamsCheckRegion = "ENAM"
	AccountLoadBalancerPoolUpdateParamsCheckRegionWeu        AccountLoadBalancerPoolUpdateParamsCheckRegion = "WEU"
	AccountLoadBalancerPoolUpdateParamsCheckRegionEeu        AccountLoadBalancerPoolUpdateParamsCheckRegion = "EEU"
	AccountLoadBalancerPoolUpdateParamsCheckRegionNsam       AccountLoadBalancerPoolUpdateParamsCheckRegion = "NSAM"
	AccountLoadBalancerPoolUpdateParamsCheckRegionSsam       AccountLoadBalancerPoolUpdateParamsCheckRegion = "SSAM"
	AccountLoadBalancerPoolUpdateParamsCheckRegionOc         AccountLoadBalancerPoolUpdateParamsCheckRegion = "OC"
	AccountLoadBalancerPoolUpdateParamsCheckRegionMe         AccountLoadBalancerPoolUpdateParamsCheckRegion = "ME"
	AccountLoadBalancerPoolUpdateParamsCheckRegionNaf        AccountLoadBalancerPoolUpdateParamsCheckRegion = "NAF"
	AccountLoadBalancerPoolUpdateParamsCheckRegionSaf        AccountLoadBalancerPoolUpdateParamsCheckRegion = "SAF"
	AccountLoadBalancerPoolUpdateParamsCheckRegionSas        AccountLoadBalancerPoolUpdateParamsCheckRegion = "SAS"
	AccountLoadBalancerPoolUpdateParamsCheckRegionSeas       AccountLoadBalancerPoolUpdateParamsCheckRegion = "SEAS"
	AccountLoadBalancerPoolUpdateParamsCheckRegionNeas       AccountLoadBalancerPoolUpdateParamsCheckRegion = "NEAS"
	AccountLoadBalancerPoolUpdateParamsCheckRegionAllRegions AccountLoadBalancerPoolUpdateParamsCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type AccountLoadBalancerPoolUpdateParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[AccountLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[AccountLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r AccountLoadBalancerPoolUpdateParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type AccountLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy string

const (
	AccountLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyRandom AccountLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy = "random"
	AccountLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyHash   AccountLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type AccountLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy string

const (
	AccountLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicyHash AccountLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type AccountLoadBalancerPoolUpdateParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[AccountLoadBalancerPoolUpdateParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[AccountLoadBalancerPoolUpdateParamsNotificationFilterPool] `json:"pool"`
}

func (r AccountLoadBalancerPoolUpdateParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type AccountLoadBalancerPoolUpdateParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r AccountLoadBalancerPoolUpdateParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type AccountLoadBalancerPoolUpdateParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r AccountLoadBalancerPoolUpdateParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type AccountLoadBalancerPoolUpdateParamsOriginSteering struct {
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
	Policy param.Field[AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicy] `json:"policy"`
}

func (r AccountLoadBalancerPoolUpdateParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicy string

const (
	AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicyRandom                   AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "random"
	AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicyHash                     AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "hash"
	AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicyLeastOutstandingRequests AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "least_outstanding_requests"
	AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicyLeastConnections         AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "least_connections"
)

type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOrigin] `json:"origins,required"`
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
	LoadShedding param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteering] `json:"origin_steering"`
}

func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOrigin struct {
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
	Header param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginsHeader] `json:"header"`
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

func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures load shedding policies and percentages for the pool.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicyRandom AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy = "random"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicyHash   AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicy string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicyHash AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterPool] `json:"pool"`
}

func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteering struct {
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
	Policy param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy] `json:"policy"`
}

func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyRandom                   AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "random"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyHash                     AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "hash"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyLeastOutstandingRequests AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "least_outstanding_requests"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyLeastConnections         AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "least_connections"
)

type AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams struct {
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `query:"monitor"`
}

// URLQuery serializes
// [AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams]'s query
// parameters as `url.Values`.
func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams struct {
	// The email address to send health status notifications to. This field is now
	// deprecated in favor of Cloudflare Notifications for Load Balancing, so only
	// resetting this field with an empty string `""` is accepted.
	NotificationEmail param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParamsNotificationEmail] `json:"notification_email"`
}

func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The email address to send health status notifications to. This field is now
// deprecated in favor of Cloudflare Notifications for Load Balancing, so only
// resetting this field with an empty string `""` is accepted.
type AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParamsNotificationEmail string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParamsNotificationEmailEmpty AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParamsNotificationEmail = "\"\""
)

type AccountLoadBalancerPoolPatchParams struct {
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]AccountLoadBalancerPoolPatchParamsCheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[AccountLoadBalancerPoolPatchParamsLoadShedding] `json:"load_shedding"`
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
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail param.Field[string] `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter param.Field[AccountLoadBalancerPoolPatchParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[AccountLoadBalancerPoolPatchParamsOriginSteering] `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]AccountLoadBalancerPoolPatchParamsOrigin] `json:"origins"`
}

func (r AccountLoadBalancerPoolPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type AccountLoadBalancerPoolPatchParamsCheckRegion string

const (
	AccountLoadBalancerPoolPatchParamsCheckRegionWnam       AccountLoadBalancerPoolPatchParamsCheckRegion = "WNAM"
	AccountLoadBalancerPoolPatchParamsCheckRegionEnam       AccountLoadBalancerPoolPatchParamsCheckRegion = "ENAM"
	AccountLoadBalancerPoolPatchParamsCheckRegionWeu        AccountLoadBalancerPoolPatchParamsCheckRegion = "WEU"
	AccountLoadBalancerPoolPatchParamsCheckRegionEeu        AccountLoadBalancerPoolPatchParamsCheckRegion = "EEU"
	AccountLoadBalancerPoolPatchParamsCheckRegionNsam       AccountLoadBalancerPoolPatchParamsCheckRegion = "NSAM"
	AccountLoadBalancerPoolPatchParamsCheckRegionSsam       AccountLoadBalancerPoolPatchParamsCheckRegion = "SSAM"
	AccountLoadBalancerPoolPatchParamsCheckRegionOc         AccountLoadBalancerPoolPatchParamsCheckRegion = "OC"
	AccountLoadBalancerPoolPatchParamsCheckRegionMe         AccountLoadBalancerPoolPatchParamsCheckRegion = "ME"
	AccountLoadBalancerPoolPatchParamsCheckRegionNaf        AccountLoadBalancerPoolPatchParamsCheckRegion = "NAF"
	AccountLoadBalancerPoolPatchParamsCheckRegionSaf        AccountLoadBalancerPoolPatchParamsCheckRegion = "SAF"
	AccountLoadBalancerPoolPatchParamsCheckRegionSas        AccountLoadBalancerPoolPatchParamsCheckRegion = "SAS"
	AccountLoadBalancerPoolPatchParamsCheckRegionSeas       AccountLoadBalancerPoolPatchParamsCheckRegion = "SEAS"
	AccountLoadBalancerPoolPatchParamsCheckRegionNeas       AccountLoadBalancerPoolPatchParamsCheckRegion = "NEAS"
	AccountLoadBalancerPoolPatchParamsCheckRegionAllRegions AccountLoadBalancerPoolPatchParamsCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type AccountLoadBalancerPoolPatchParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[AccountLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[AccountLoadBalancerPoolPatchParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r AccountLoadBalancerPoolPatchParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type AccountLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicy string

const (
	AccountLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicyRandom AccountLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicy = "random"
	AccountLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicyHash   AccountLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type AccountLoadBalancerPoolPatchParamsLoadSheddingSessionPolicy string

const (
	AccountLoadBalancerPoolPatchParamsLoadSheddingSessionPolicyHash AccountLoadBalancerPoolPatchParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type AccountLoadBalancerPoolPatchParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[AccountLoadBalancerPoolPatchParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[AccountLoadBalancerPoolPatchParamsNotificationFilterPool] `json:"pool"`
}

func (r AccountLoadBalancerPoolPatchParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type AccountLoadBalancerPoolPatchParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r AccountLoadBalancerPoolPatchParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type AccountLoadBalancerPoolPatchParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r AccountLoadBalancerPoolPatchParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type AccountLoadBalancerPoolPatchParamsOriginSteering struct {
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
	Policy param.Field[AccountLoadBalancerPoolPatchParamsOriginSteeringPolicy] `json:"policy"`
}

func (r AccountLoadBalancerPoolPatchParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type AccountLoadBalancerPoolPatchParamsOriginSteeringPolicy string

const (
	AccountLoadBalancerPoolPatchParamsOriginSteeringPolicyRandom                   AccountLoadBalancerPoolPatchParamsOriginSteeringPolicy = "random"
	AccountLoadBalancerPoolPatchParamsOriginSteeringPolicyHash                     AccountLoadBalancerPoolPatchParamsOriginSteeringPolicy = "hash"
	AccountLoadBalancerPoolPatchParamsOriginSteeringPolicyLeastOutstandingRequests AccountLoadBalancerPoolPatchParamsOriginSteeringPolicy = "least_outstanding_requests"
	AccountLoadBalancerPoolPatchParamsOriginSteeringPolicyLeastConnections         AccountLoadBalancerPoolPatchParamsOriginSteeringPolicy = "least_connections"
)

type AccountLoadBalancerPoolPatchParamsOrigin struct {
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
	Header param.Field[AccountLoadBalancerPoolPatchParamsOriginsHeader] `json:"header"`
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

func (r AccountLoadBalancerPoolPatchParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type AccountLoadBalancerPoolPatchParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r AccountLoadBalancerPoolPatchParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
