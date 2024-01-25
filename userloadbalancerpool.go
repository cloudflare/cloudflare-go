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

// UserLoadBalancerPoolService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserLoadBalancerPoolService]
// method instead.
type UserLoadBalancerPoolService struct {
	Options    []option.RequestOption
	Health     *UserLoadBalancerPoolHealthService
	Previews   *UserLoadBalancerPoolPreviewService
	References *UserLoadBalancerPoolReferenceService
}

// NewUserLoadBalancerPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserLoadBalancerPoolService(opts ...option.RequestOption) (r *UserLoadBalancerPoolService) {
	r = &UserLoadBalancerPoolService{}
	r.Options = opts
	r.Health = NewUserLoadBalancerPoolHealthService(opts...)
	r.Previews = NewUserLoadBalancerPoolPreviewService(opts...)
	r.References = NewUserLoadBalancerPoolReferenceService(opts...)
	return
}

// Fetch a single configured pool.
func (r *UserLoadBalancerPoolService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured pool.
func (r *UserLoadBalancerPoolService) Update(ctx context.Context, identifier string, body UserLoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured pool.
func (r *UserLoadBalancerPoolService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserLoadBalancerPoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new pool.
func (r *UserLoadBalancerPoolService) LoadBalancerPoolsNewPool(ctx context.Context, body UserLoadBalancerPoolLoadBalancerPoolsNewPoolParams, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured pools.
func (r *UserLoadBalancerPoolService) LoadBalancerPoolsListPools(ctx context.Context, query UserLoadBalancerPoolLoadBalancerPoolsListPoolsParams, opts ...option.RequestOption) (res *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Apply changes to a number of existing pools, overwriting the supplied
// properties. Pools are ordered by ascending `name`. Returns the list of affected
// pools. Supports the standard pagination query parameters, either
// `limit`/`offset` or `per_page`/`page`.
func (r *UserLoadBalancerPoolService) LoadBalancerPoolsPatchPools(ctx context.Context, body UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParams, opts ...option.RequestOption) (res *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Apply changes to an existing pool, overwriting the supplied properties.
func (r *UserLoadBalancerPoolService) Patch(ctx context.Context, identifier string, body UserLoadBalancerPoolPatchParams, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type UserLoadBalancerPoolDeleteResponse struct {
	Errors   []UserLoadBalancerPoolDeleteResponseError   `json:"errors"`
	Messages []UserLoadBalancerPoolDeleteResponseMessage `json:"messages"`
	Result   UserLoadBalancerPoolDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolDeleteResponseSuccess `json:"success"`
	JSON    userLoadBalancerPoolDeleteResponseJSON    `json:"-"`
}

// userLoadBalancerPoolDeleteResponseJSON contains the JSON metadata for the struct
// [UserLoadBalancerPoolDeleteResponse]
type userLoadBalancerPoolDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    userLoadBalancerPoolDeleteResponseErrorJSON `json:"-"`
}

// userLoadBalancerPoolDeleteResponseErrorJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolDeleteResponseError]
type userLoadBalancerPoolDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    userLoadBalancerPoolDeleteResponseMessageJSON `json:"-"`
}

// userLoadBalancerPoolDeleteResponseMessageJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolDeleteResponseMessage]
type userLoadBalancerPoolDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolDeleteResponseResult struct {
	ID   string                                       `json:"id"`
	JSON userLoadBalancerPoolDeleteResponseResultJSON `json:"-"`
}

// userLoadBalancerPoolDeleteResponseResultJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolDeleteResponseResult]
type userLoadBalancerPoolDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolDeleteResponseSuccess bool

const (
	UserLoadBalancerPoolDeleteResponseSuccessTrue UserLoadBalancerPoolDeleteResponseSuccess = true
)

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponse struct {
	Errors     []UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseError    `json:"errors"`
	Messages   []UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseMessage  `json:"messages"`
	Result     []UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResult   `json:"result"`
	ResultInfo UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseSuccess `json:"success"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseJSON    `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponse]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseErrorJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseErrorJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseError]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseMessageJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseMessageJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseMessage]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResult struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                                 `json:"created_on" format:"date-time"`
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
	LoadShedding UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadShedding `json:"load_shedding"`
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
	NotificationFilter UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOrigin `json:"origins"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultJSON     `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResult]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultJSON struct {
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

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion string

const (
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionWnam       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "WNAM"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionEnam       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "ENAM"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionWeu        UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "WEU"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionEeu        UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "EEU"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionNsam       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "NSAM"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionSsam       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SSAM"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionOc         UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "OC"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionMe         UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "ME"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionNaf        UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "NAF"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionSaf        UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SAF"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionSas        UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SAS"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionSeas       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "SEAS"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionNeas       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "NEAS"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegionAllRegions UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingJSON          `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadShedding]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicyHash UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool `json:"pool,nullable"`
	JSON userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilter]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                     `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterOriginJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                   `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterPoolJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteering struct {
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
	Policy UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy `json:"policy"`
	JSON   userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringJSON   `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteering]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyRandom                   UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyHash                     UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOrigin struct {
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
	Header UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginsHeader `json:"header"`
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
	Weight float64                                                                `json:"weight"`
	JSON   userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOrigin]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginJSON struct {
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

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                      `json:"Host"`
	JSON userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginsHeaderJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginsHeaderJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginsHeader]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                              `json:"total_count"`
	JSON       userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultInfoJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultInfoJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultInfo]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseSuccess bool

const (
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseSuccessTrue UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseSuccess = true
)

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponse struct {
	Errors     []UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseError    `json:"errors"`
	Messages   []UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseMessage  `json:"messages"`
	Result     []UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResult   `json:"result"`
	ResultInfo UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseSuccess `json:"success"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseJSON    `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponse]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseError struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseErrorJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseErrorJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseError]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseMessage struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseMessageJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseMessageJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseMessage]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResult struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion `json:"check_regions,nullable"`
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
	LoadShedding UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadShedding `json:"load_shedding"`
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
	NotificationFilter UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOrigin `json:"origins"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultJSON     `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResult]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultJSON struct {
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

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion string

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionWnam       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "WNAM"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionEnam       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "ENAM"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionWeu        UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "WEU"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionEeu        UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "EEU"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionNsam       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "NSAM"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSsam       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SSAM"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionOc         UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "OC"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionMe         UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "ME"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionNaf        UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "NAF"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSaf        UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SAF"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSas        UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SAS"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionSeas       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "SEAS"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionNeas       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "NEAS"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegionAllRegions UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingJSON          `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadShedding]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicyHash UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool `json:"pool,nullable"`
	JSON userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                      `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOriginJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                    `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPoolJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteering struct {
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
	Policy UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy `json:"policy"`
	JSON   userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringJSON   `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteering]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyRandom                   UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyHash                     UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOrigin struct {
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
	Header UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader `json:"header"`
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
	JSON   userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOrigin]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginJSON struct {
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

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                       `json:"Host"`
	JSON userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginsHeaderJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginsHeaderJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                               `json:"total_count"`
	JSON       userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultInfoJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultInfoJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultInfo]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseSuccess bool

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseSuccessTrue UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseSuccess = true
)

type UserLoadBalancerPoolUpdateParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]UserLoadBalancerPoolUpdateParamsOrigin] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]UserLoadBalancerPoolUpdateParamsCheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[UserLoadBalancerPoolUpdateParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[UserLoadBalancerPoolUpdateParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[UserLoadBalancerPoolUpdateParamsOriginSteering] `json:"origin_steering"`
}

func (r UserLoadBalancerPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserLoadBalancerPoolUpdateParamsOrigin struct {
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
	Header param.Field[UserLoadBalancerPoolUpdateParamsOriginsHeader] `json:"header"`
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

func (r UserLoadBalancerPoolUpdateParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolUpdateParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r UserLoadBalancerPoolUpdateParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolUpdateParamsCheckRegion string

const (
	UserLoadBalancerPoolUpdateParamsCheckRegionWnam       UserLoadBalancerPoolUpdateParamsCheckRegion = "WNAM"
	UserLoadBalancerPoolUpdateParamsCheckRegionEnam       UserLoadBalancerPoolUpdateParamsCheckRegion = "ENAM"
	UserLoadBalancerPoolUpdateParamsCheckRegionWeu        UserLoadBalancerPoolUpdateParamsCheckRegion = "WEU"
	UserLoadBalancerPoolUpdateParamsCheckRegionEeu        UserLoadBalancerPoolUpdateParamsCheckRegion = "EEU"
	UserLoadBalancerPoolUpdateParamsCheckRegionNsam       UserLoadBalancerPoolUpdateParamsCheckRegion = "NSAM"
	UserLoadBalancerPoolUpdateParamsCheckRegionSsam       UserLoadBalancerPoolUpdateParamsCheckRegion = "SSAM"
	UserLoadBalancerPoolUpdateParamsCheckRegionOc         UserLoadBalancerPoolUpdateParamsCheckRegion = "OC"
	UserLoadBalancerPoolUpdateParamsCheckRegionMe         UserLoadBalancerPoolUpdateParamsCheckRegion = "ME"
	UserLoadBalancerPoolUpdateParamsCheckRegionNaf        UserLoadBalancerPoolUpdateParamsCheckRegion = "NAF"
	UserLoadBalancerPoolUpdateParamsCheckRegionSaf        UserLoadBalancerPoolUpdateParamsCheckRegion = "SAF"
	UserLoadBalancerPoolUpdateParamsCheckRegionSas        UserLoadBalancerPoolUpdateParamsCheckRegion = "SAS"
	UserLoadBalancerPoolUpdateParamsCheckRegionSeas       UserLoadBalancerPoolUpdateParamsCheckRegion = "SEAS"
	UserLoadBalancerPoolUpdateParamsCheckRegionNeas       UserLoadBalancerPoolUpdateParamsCheckRegion = "NEAS"
	UserLoadBalancerPoolUpdateParamsCheckRegionAllRegions UserLoadBalancerPoolUpdateParamsCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolUpdateParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[UserLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[UserLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r UserLoadBalancerPoolUpdateParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicyHash UserLoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolUpdateParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[UserLoadBalancerPoolUpdateParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[UserLoadBalancerPoolUpdateParamsNotificationFilterPool] `json:"pool"`
}

func (r UserLoadBalancerPoolUpdateParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolUpdateParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r UserLoadBalancerPoolUpdateParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolUpdateParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r UserLoadBalancerPoolUpdateParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolUpdateParamsOriginSteering struct {
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
	Policy param.Field[UserLoadBalancerPoolUpdateParamsOriginSteeringPolicy] `json:"policy"`
}

func (r UserLoadBalancerPoolUpdateParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type UserLoadBalancerPoolUpdateParamsOriginSteeringPolicy string

const (
	UserLoadBalancerPoolUpdateParamsOriginSteeringPolicyRandom                   UserLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "random"
	UserLoadBalancerPoolUpdateParamsOriginSteeringPolicyHash                     UserLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolUpdateParamsOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolUpdateParamsOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOrigin] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteering] `json:"origin_steering"`
}

func (r UserLoadBalancerPoolLoadBalancerPoolsNewPoolParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOrigin struct {
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
	Header param.Field[UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginsHeader] `json:"header"`
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

func (r UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion string

const (
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionWnam       UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "WNAM"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionEnam       UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "ENAM"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionWeu        UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "WEU"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionEeu        UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "EEU"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionNsam       UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "NSAM"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionSsam       UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "SSAM"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionOc         UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "OC"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionMe         UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "ME"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionNaf        UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "NAF"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionSaf        UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "SAF"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionSas        UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "SAS"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionSeas       UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "SEAS"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionNeas       UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "NEAS"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegionAllRegions UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicyHash UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilterPool] `json:"pool"`
}

func (r UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteering struct {
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
	Policy param.Field[UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy] `json:"policy"`
}

func (r UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyRandom                   UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyHash                     UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsParams struct {
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `query:"monitor"`
}

// URLQuery serializes [UserLoadBalancerPoolLoadBalancerPoolsListPoolsParams]'s
// query parameters as `url.Values`.
func (r UserLoadBalancerPoolLoadBalancerPoolsListPoolsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParams struct {
	// The email address to send health status notifications to. This field is now
	// deprecated in favor of Cloudflare Notifications for Load Balancing, so only
	// resetting this field with an empty string `""` is accepted.
	NotificationEmail param.Field[UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParamsNotificationEmail] `json:"notification_email"`
}

func (r UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The email address to send health status notifications to. This field is now
// deprecated in favor of Cloudflare Notifications for Load Balancing, so only
// resetting this field with an empty string `""` is accepted.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParamsNotificationEmail string

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParamsNotificationEmailEmpty UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParamsNotificationEmail = "\"\""
)

type UserLoadBalancerPoolPatchParams struct {
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]UserLoadBalancerPoolPatchParamsCheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[UserLoadBalancerPoolPatchParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[UserLoadBalancerPoolPatchParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[UserLoadBalancerPoolPatchParamsOriginSteering] `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]UserLoadBalancerPoolPatchParamsOrigin] `json:"origins"`
}

func (r UserLoadBalancerPoolPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolPatchParamsCheckRegion string

const (
	UserLoadBalancerPoolPatchParamsCheckRegionWnam       UserLoadBalancerPoolPatchParamsCheckRegion = "WNAM"
	UserLoadBalancerPoolPatchParamsCheckRegionEnam       UserLoadBalancerPoolPatchParamsCheckRegion = "ENAM"
	UserLoadBalancerPoolPatchParamsCheckRegionWeu        UserLoadBalancerPoolPatchParamsCheckRegion = "WEU"
	UserLoadBalancerPoolPatchParamsCheckRegionEeu        UserLoadBalancerPoolPatchParamsCheckRegion = "EEU"
	UserLoadBalancerPoolPatchParamsCheckRegionNsam       UserLoadBalancerPoolPatchParamsCheckRegion = "NSAM"
	UserLoadBalancerPoolPatchParamsCheckRegionSsam       UserLoadBalancerPoolPatchParamsCheckRegion = "SSAM"
	UserLoadBalancerPoolPatchParamsCheckRegionOc         UserLoadBalancerPoolPatchParamsCheckRegion = "OC"
	UserLoadBalancerPoolPatchParamsCheckRegionMe         UserLoadBalancerPoolPatchParamsCheckRegion = "ME"
	UserLoadBalancerPoolPatchParamsCheckRegionNaf        UserLoadBalancerPoolPatchParamsCheckRegion = "NAF"
	UserLoadBalancerPoolPatchParamsCheckRegionSaf        UserLoadBalancerPoolPatchParamsCheckRegion = "SAF"
	UserLoadBalancerPoolPatchParamsCheckRegionSas        UserLoadBalancerPoolPatchParamsCheckRegion = "SAS"
	UserLoadBalancerPoolPatchParamsCheckRegionSeas       UserLoadBalancerPoolPatchParamsCheckRegion = "SEAS"
	UserLoadBalancerPoolPatchParamsCheckRegionNeas       UserLoadBalancerPoolPatchParamsCheckRegion = "NEAS"
	UserLoadBalancerPoolPatchParamsCheckRegionAllRegions UserLoadBalancerPoolPatchParamsCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolPatchParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[UserLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[UserLoadBalancerPoolPatchParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r UserLoadBalancerPoolPatchParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolPatchParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolPatchParamsLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolPatchParamsLoadSheddingSessionPolicyHash UserLoadBalancerPoolPatchParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolPatchParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[UserLoadBalancerPoolPatchParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[UserLoadBalancerPoolPatchParamsNotificationFilterPool] `json:"pool"`
}

func (r UserLoadBalancerPoolPatchParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolPatchParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r UserLoadBalancerPoolPatchParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolPatchParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r UserLoadBalancerPoolPatchParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolPatchParamsOriginSteering struct {
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
	Policy param.Field[UserLoadBalancerPoolPatchParamsOriginSteeringPolicy] `json:"policy"`
}

func (r UserLoadBalancerPoolPatchParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type UserLoadBalancerPoolPatchParamsOriginSteeringPolicy string

const (
	UserLoadBalancerPoolPatchParamsOriginSteeringPolicyRandom                   UserLoadBalancerPoolPatchParamsOriginSteeringPolicy = "random"
	UserLoadBalancerPoolPatchParamsOriginSteeringPolicyHash                     UserLoadBalancerPoolPatchParamsOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolPatchParamsOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolPatchParamsOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolPatchParamsOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolPatchParamsOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolPatchParamsOrigin struct {
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
	Header param.Field[UserLoadBalancerPoolPatchParamsOriginsHeader] `json:"header"`
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

func (r UserLoadBalancerPoolPatchParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolPatchParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r UserLoadBalancerPoolPatchParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
