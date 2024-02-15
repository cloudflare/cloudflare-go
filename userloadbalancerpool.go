// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
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

// Create a new pool.
func (r *UserLoadBalancerPoolService) LoadBalancerPoolsNewPool(ctx context.Context, body UserLoadBalancerPoolLoadBalancerPoolsNewPoolParams, opts ...option.RequestOption) (res *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelope
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured pools.
func (r *UserLoadBalancerPoolService) LoadBalancerPoolsListPools(ctx context.Context, query UserLoadBalancerPoolLoadBalancerPoolsListPoolsParams, opts ...option.RequestOption) (res *[]UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelope
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to a number of existing pools, overwriting the supplied
// properties. Pools are ordered by ascending `name`. Returns the list of affected
// pools. Supports the standard pagination query parameters, either
// `limit`/`offset` or `per_page`/`page`.
func (r *UserLoadBalancerPoolService) LoadBalancerPoolsPatchPools(ctx context.Context, body UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParams, opts ...option.RequestOption) (res *[]UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelope
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                         `json:"created_on" format:"date-time"`
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
	LoadShedding UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOrigin `json:"origins"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseJSON     `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponse]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseJSON struct {
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

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion string

const (
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionWnam       UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "WNAM"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionEnam       UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "ENAM"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionWeu        UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "WEU"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionEeu        UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "EEU"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionNsam       UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "NSAM"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionSsam       UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "SSAM"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionOc         UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "OC"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionMe         UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "ME"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionNaf        UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "NAF"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionSaf        UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "SAF"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionSas        UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "SAS"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionSeas       UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "SEAS"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionNeas       UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "NEAS"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegionAllRegions UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingJSON          `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadShedding]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingSessionPolicyHash UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterPool `json:"pool,nullable"`
	JSON userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilter]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                             `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterOriginJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterOrigin]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                           `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterPoolJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterPool]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteering struct {
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
	Policy UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy `json:"policy"`
	JSON   userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringJSON   `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteering]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringPolicyRandom                   UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringPolicyHash                     UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOrigin struct {
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
	Header UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginsHeader `json:"header"`
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
	Weight float64                                                        `json:"weight"`
	JSON   userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOrigin]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginJSON struct {
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

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                              `json:"Host"`
	JSON userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginsHeaderJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginsHeaderJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginsHeader]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                           `json:"created_on" format:"date-time"`
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
	LoadShedding UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOrigin `json:"origins"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseJSON     `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponse]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseJSON struct {
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

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion string

const (
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionWnam       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "WNAM"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionEnam       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "ENAM"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionWeu        UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "WEU"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionEeu        UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "EEU"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionNsam       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "NSAM"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionSsam       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "SSAM"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionOc         UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "OC"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionMe         UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "ME"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionNaf        UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "NAF"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionSaf        UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "SAF"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionSas        UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "SAS"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionSeas       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "SEAS"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionNeas       UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "NEAS"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegionAllRegions UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingJSON          `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadShedding]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingSessionPolicyHash UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterPool `json:"pool,nullable"`
	JSON userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilter]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                               `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterOriginJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterOrigin]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                             `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterPoolJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterPool]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteering struct {
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
	Policy UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy `json:"policy"`
	JSON   userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringJSON   `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteering]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringPolicyRandom                   UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringPolicyHash                     UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOrigin struct {
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
	Header UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginsHeader `json:"header"`
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
	Weight float64                                                          `json:"weight"`
	JSON   userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOrigin]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginJSON struct {
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

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                `json:"Host"`
	JSON userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginsHeaderJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginsHeaderJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginsHeader]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                            `json:"created_on" format:"date-time"`
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
	LoadShedding UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOrigin `json:"origins"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseJSON     `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponse]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseJSON struct {
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

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion string

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionWnam       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "WNAM"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionEnam       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "ENAM"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionWeu        UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "WEU"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionEeu        UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "EEU"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionNsam       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "NSAM"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionSsam       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "SSAM"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionOc         UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "OC"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionMe         UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "ME"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionNaf        UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "NAF"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionSaf        UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "SAF"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionSas        UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "SAS"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionSeas       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "SEAS"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionNeas       UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "NEAS"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegionAllRegions UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingJSON          `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadShedding]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingSessionPolicyHash UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterPool `json:"pool,nullable"`
	JSON userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilter]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterOriginJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterOrigin]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                              `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterPoolJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterPool]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteering struct {
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
	Policy UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy `json:"policy"`
	JSON   userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringJSON   `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteering]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicyRandom                   UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicyHash                     UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOrigin struct {
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
	Header UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginsHeader `json:"header"`
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
	Weight float64                                                           `json:"weight"`
	JSON   userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOrigin]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginJSON struct {
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

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                 `json:"Host"`
	JSON userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginsHeaderJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginsHeaderJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginsHeader]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelope]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeErrors]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeMessages]
type userLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeSuccessTrue UserLoadBalancerPoolLoadBalancerPoolsNewPoolResponseEnvelopeSuccess = true
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

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeJSON       `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelope]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeErrors struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeErrors]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeMessages struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeMessages]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeSuccessTrue UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                      `json:"total_count"`
	JSON       userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeResultInfoJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeResultInfo]
type userLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsListPoolsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeJSON       `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelope]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeErrors]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeMessages]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeSuccessTrue UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                       `json:"total_count"`
	JSON       userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfoJSON `json:"-"`
}

// userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfo]
type userLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
