// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// UserLoadBalancerPoolService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserLoadBalancerPoolService]
// method instead.
type UserLoadBalancerPoolService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserLoadBalancerPoolService(opts ...option.RequestOption) (r *UserLoadBalancerPoolService) {
	r = &UserLoadBalancerPoolService{}
	r.Options = opts
	return
}

// Create a new pool.
func (r *UserLoadBalancerPoolService) New(ctx context.Context, body UserLoadBalancerPoolNewParams, opts ...option.RequestOption) (res *UserLoadBalancerPoolNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolNewResponseEnvelope
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured pool.
func (r *UserLoadBalancerPoolService) Update(ctx context.Context, poolID string, body UserLoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *UserLoadBalancerPoolUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolUpdateResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured pools.
func (r *UserLoadBalancerPoolService) List(ctx context.Context, query UserLoadBalancerPoolListParams, opts ...option.RequestOption) (res *[]UserLoadBalancerPoolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolListResponseEnvelope
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured pool.
func (r *UserLoadBalancerPoolService) Delete(ctx context.Context, poolID string, opts ...option.RequestOption) (res *UserLoadBalancerPoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolDeleteResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing pool, overwriting the supplied properties.
func (r *UserLoadBalancerPoolService) Edit(ctx context.Context, poolID string, body UserLoadBalancerPoolEditParams, opts ...option.RequestOption) (res *UserLoadBalancerPoolEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolEditResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured pool.
func (r *UserLoadBalancerPoolService) Get(ctx context.Context, poolID string, opts ...option.RequestOption) (res *UserLoadBalancerPoolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolGetResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch the latest pool health status for a single pool.
func (r *UserLoadBalancerPoolService) Health(ctx context.Context, poolID string, opts ...option.RequestOption) (res *UserLoadBalancerPoolHealthResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolHealthResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/health", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Preview pool health using provided monitor details. The returned preview_id can
// be used in the preview endpoint to retrieve the results.
func (r *UserLoadBalancerPoolService) Preview(ctx context.Context, poolID string, body UserLoadBalancerPoolPreviewParams, opts ...option.RequestOption) (res *UserLoadBalancerPoolPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolPreviewResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/preview", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the list of resources that reference the provided pool.
func (r *UserLoadBalancerPoolService) References(ctx context.Context, poolID string, opts ...option.RequestOption) (res *[]UserLoadBalancerPoolReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolReferencesResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/references", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserLoadBalancerPoolNewResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []UserLoadBalancerPoolNewResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                    `json:"created_on" format:"date-time"`
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
	LoadShedding UserLoadBalancerPoolNewResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter UserLoadBalancerPoolNewResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering UserLoadBalancerPoolNewResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []UserLoadBalancerPoolNewResponseOrigin `json:"origins"`
	JSON    userLoadBalancerPoolNewResponseJSON     `json:"-"`
}

// userLoadBalancerPoolNewResponseJSON contains the JSON metadata for the struct
// [UserLoadBalancerPoolNewResponse]
type userLoadBalancerPoolNewResponseJSON struct {
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

func (r *UserLoadBalancerPoolNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolNewResponseCheckRegion string

const (
	UserLoadBalancerPoolNewResponseCheckRegionWnam       UserLoadBalancerPoolNewResponseCheckRegion = "WNAM"
	UserLoadBalancerPoolNewResponseCheckRegionEnam       UserLoadBalancerPoolNewResponseCheckRegion = "ENAM"
	UserLoadBalancerPoolNewResponseCheckRegionWeu        UserLoadBalancerPoolNewResponseCheckRegion = "WEU"
	UserLoadBalancerPoolNewResponseCheckRegionEeu        UserLoadBalancerPoolNewResponseCheckRegion = "EEU"
	UserLoadBalancerPoolNewResponseCheckRegionNsam       UserLoadBalancerPoolNewResponseCheckRegion = "NSAM"
	UserLoadBalancerPoolNewResponseCheckRegionSsam       UserLoadBalancerPoolNewResponseCheckRegion = "SSAM"
	UserLoadBalancerPoolNewResponseCheckRegionOc         UserLoadBalancerPoolNewResponseCheckRegion = "OC"
	UserLoadBalancerPoolNewResponseCheckRegionMe         UserLoadBalancerPoolNewResponseCheckRegion = "ME"
	UserLoadBalancerPoolNewResponseCheckRegionNaf        UserLoadBalancerPoolNewResponseCheckRegion = "NAF"
	UserLoadBalancerPoolNewResponseCheckRegionSaf        UserLoadBalancerPoolNewResponseCheckRegion = "SAF"
	UserLoadBalancerPoolNewResponseCheckRegionSas        UserLoadBalancerPoolNewResponseCheckRegion = "SAS"
	UserLoadBalancerPoolNewResponseCheckRegionSeas       UserLoadBalancerPoolNewResponseCheckRegion = "SEAS"
	UserLoadBalancerPoolNewResponseCheckRegionNeas       UserLoadBalancerPoolNewResponseCheckRegion = "NEAS"
	UserLoadBalancerPoolNewResponseCheckRegionAllRegions UserLoadBalancerPoolNewResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolNewResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy UserLoadBalancerPoolNewResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy UserLoadBalancerPoolNewResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          userLoadBalancerPoolNewResponseLoadSheddingJSON          `json:"-"`
}

// userLoadBalancerPoolNewResponseLoadSheddingJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolNewResponseLoadShedding]
type userLoadBalancerPoolNewResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserLoadBalancerPoolNewResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolNewResponseLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolNewResponseLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolNewResponseLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolNewResponseLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolNewResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolNewResponseLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolNewResponseLoadSheddingSessionPolicyHash UserLoadBalancerPoolNewResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolNewResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin UserLoadBalancerPoolNewResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool UserLoadBalancerPoolNewResponseNotificationFilterPool `json:"pool,nullable"`
	JSON userLoadBalancerPoolNewResponseNotificationFilterJSON `json:"-"`
}

// userLoadBalancerPoolNewResponseNotificationFilterJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolNewResponseNotificationFilter]
type userLoadBalancerPoolNewResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolNewResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolNewResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                        `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolNewResponseNotificationFilterOriginJSON `json:"-"`
}

// userLoadBalancerPoolNewResponseNotificationFilterOriginJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerPoolNewResponseNotificationFilterOrigin]
type userLoadBalancerPoolNewResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolNewResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolNewResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                      `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolNewResponseNotificationFilterPoolJSON `json:"-"`
}

// userLoadBalancerPoolNewResponseNotificationFilterPoolJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolNewResponseNotificationFilterPool]
type userLoadBalancerPoolNewResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolNewResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolNewResponseOriginSteering struct {
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
	Policy UserLoadBalancerPoolNewResponseOriginSteeringPolicy `json:"policy"`
	JSON   userLoadBalancerPoolNewResponseOriginSteeringJSON   `json:"-"`
}

// userLoadBalancerPoolNewResponseOriginSteeringJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolNewResponseOriginSteering]
type userLoadBalancerPoolNewResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolNewResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type UserLoadBalancerPoolNewResponseOriginSteeringPolicy string

const (
	UserLoadBalancerPoolNewResponseOriginSteeringPolicyRandom                   UserLoadBalancerPoolNewResponseOriginSteeringPolicy = "random"
	UserLoadBalancerPoolNewResponseOriginSteeringPolicyHash                     UserLoadBalancerPoolNewResponseOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolNewResponseOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolNewResponseOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolNewResponseOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolNewResponseOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolNewResponseOrigin struct {
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
	Header UserLoadBalancerPoolNewResponseOriginsHeader `json:"header"`
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
	Weight float64                                   `json:"weight"`
	JSON   userLoadBalancerPoolNewResponseOriginJSON `json:"-"`
}

// userLoadBalancerPoolNewResponseOriginJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolNewResponseOrigin]
type userLoadBalancerPoolNewResponseOriginJSON struct {
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

func (r *UserLoadBalancerPoolNewResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolNewResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                         `json:"Host"`
	JSON userLoadBalancerPoolNewResponseOriginsHeaderJSON `json:"-"`
}

// userLoadBalancerPoolNewResponseOriginsHeaderJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolNewResponseOriginsHeader]
type userLoadBalancerPoolNewResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolNewResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolUpdateResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []UserLoadBalancerPoolUpdateResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                       `json:"created_on" format:"date-time"`
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
	LoadShedding UserLoadBalancerPoolUpdateResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter UserLoadBalancerPoolUpdateResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering UserLoadBalancerPoolUpdateResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []UserLoadBalancerPoolUpdateResponseOrigin `json:"origins"`
	JSON    userLoadBalancerPoolUpdateResponseJSON     `json:"-"`
}

// userLoadBalancerPoolUpdateResponseJSON contains the JSON metadata for the struct
// [UserLoadBalancerPoolUpdateResponse]
type userLoadBalancerPoolUpdateResponseJSON struct {
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

func (r *UserLoadBalancerPoolUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolUpdateResponseCheckRegion string

const (
	UserLoadBalancerPoolUpdateResponseCheckRegionWnam       UserLoadBalancerPoolUpdateResponseCheckRegion = "WNAM"
	UserLoadBalancerPoolUpdateResponseCheckRegionEnam       UserLoadBalancerPoolUpdateResponseCheckRegion = "ENAM"
	UserLoadBalancerPoolUpdateResponseCheckRegionWeu        UserLoadBalancerPoolUpdateResponseCheckRegion = "WEU"
	UserLoadBalancerPoolUpdateResponseCheckRegionEeu        UserLoadBalancerPoolUpdateResponseCheckRegion = "EEU"
	UserLoadBalancerPoolUpdateResponseCheckRegionNsam       UserLoadBalancerPoolUpdateResponseCheckRegion = "NSAM"
	UserLoadBalancerPoolUpdateResponseCheckRegionSsam       UserLoadBalancerPoolUpdateResponseCheckRegion = "SSAM"
	UserLoadBalancerPoolUpdateResponseCheckRegionOc         UserLoadBalancerPoolUpdateResponseCheckRegion = "OC"
	UserLoadBalancerPoolUpdateResponseCheckRegionMe         UserLoadBalancerPoolUpdateResponseCheckRegion = "ME"
	UserLoadBalancerPoolUpdateResponseCheckRegionNaf        UserLoadBalancerPoolUpdateResponseCheckRegion = "NAF"
	UserLoadBalancerPoolUpdateResponseCheckRegionSaf        UserLoadBalancerPoolUpdateResponseCheckRegion = "SAF"
	UserLoadBalancerPoolUpdateResponseCheckRegionSas        UserLoadBalancerPoolUpdateResponseCheckRegion = "SAS"
	UserLoadBalancerPoolUpdateResponseCheckRegionSeas       UserLoadBalancerPoolUpdateResponseCheckRegion = "SEAS"
	UserLoadBalancerPoolUpdateResponseCheckRegionNeas       UserLoadBalancerPoolUpdateResponseCheckRegion = "NEAS"
	UserLoadBalancerPoolUpdateResponseCheckRegionAllRegions UserLoadBalancerPoolUpdateResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolUpdateResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy UserLoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy UserLoadBalancerPoolUpdateResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          userLoadBalancerPoolUpdateResponseLoadSheddingJSON          `json:"-"`
}

// userLoadBalancerPoolUpdateResponseLoadSheddingJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolUpdateResponseLoadShedding]
type userLoadBalancerPoolUpdateResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserLoadBalancerPoolUpdateResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolUpdateResponseLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolUpdateResponseLoadSheddingSessionPolicyHash UserLoadBalancerPoolUpdateResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolUpdateResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin UserLoadBalancerPoolUpdateResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool UserLoadBalancerPoolUpdateResponseNotificationFilterPool `json:"pool,nullable"`
	JSON userLoadBalancerPoolUpdateResponseNotificationFilterJSON `json:"-"`
}

// userLoadBalancerPoolUpdateResponseNotificationFilterJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolUpdateResponseNotificationFilter]
type userLoadBalancerPoolUpdateResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolUpdateResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolUpdateResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                           `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolUpdateResponseNotificationFilterOriginJSON `json:"-"`
}

// userLoadBalancerPoolUpdateResponseNotificationFilterOriginJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerPoolUpdateResponseNotificationFilterOrigin]
type userLoadBalancerPoolUpdateResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolUpdateResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolUpdateResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                         `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolUpdateResponseNotificationFilterPoolJSON `json:"-"`
}

// userLoadBalancerPoolUpdateResponseNotificationFilterPoolJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerPoolUpdateResponseNotificationFilterPool]
type userLoadBalancerPoolUpdateResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolUpdateResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolUpdateResponseOriginSteering struct {
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
	Policy UserLoadBalancerPoolUpdateResponseOriginSteeringPolicy `json:"policy"`
	JSON   userLoadBalancerPoolUpdateResponseOriginSteeringJSON   `json:"-"`
}

// userLoadBalancerPoolUpdateResponseOriginSteeringJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolUpdateResponseOriginSteering]
type userLoadBalancerPoolUpdateResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolUpdateResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type UserLoadBalancerPoolUpdateResponseOriginSteeringPolicy string

const (
	UserLoadBalancerPoolUpdateResponseOriginSteeringPolicyRandom                   UserLoadBalancerPoolUpdateResponseOriginSteeringPolicy = "random"
	UserLoadBalancerPoolUpdateResponseOriginSteeringPolicyHash                     UserLoadBalancerPoolUpdateResponseOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolUpdateResponseOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolUpdateResponseOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolUpdateResponseOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolUpdateResponseOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolUpdateResponseOrigin struct {
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
	Header UserLoadBalancerPoolUpdateResponseOriginsHeader `json:"header"`
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
	Weight float64                                      `json:"weight"`
	JSON   userLoadBalancerPoolUpdateResponseOriginJSON `json:"-"`
}

// userLoadBalancerPoolUpdateResponseOriginJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolUpdateResponseOrigin]
type userLoadBalancerPoolUpdateResponseOriginJSON struct {
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

func (r *UserLoadBalancerPoolUpdateResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolUpdateResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                            `json:"Host"`
	JSON userLoadBalancerPoolUpdateResponseOriginsHeaderJSON `json:"-"`
}

// userLoadBalancerPoolUpdateResponseOriginsHeaderJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolUpdateResponseOriginsHeader]
type userLoadBalancerPoolUpdateResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolUpdateResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolListResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []UserLoadBalancerPoolListResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                     `json:"created_on" format:"date-time"`
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
	LoadShedding UserLoadBalancerPoolListResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter UserLoadBalancerPoolListResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering UserLoadBalancerPoolListResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []UserLoadBalancerPoolListResponseOrigin `json:"origins"`
	JSON    userLoadBalancerPoolListResponseJSON     `json:"-"`
}

// userLoadBalancerPoolListResponseJSON contains the JSON metadata for the struct
// [UserLoadBalancerPoolListResponse]
type userLoadBalancerPoolListResponseJSON struct {
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

func (r *UserLoadBalancerPoolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolListResponseCheckRegion string

const (
	UserLoadBalancerPoolListResponseCheckRegionWnam       UserLoadBalancerPoolListResponseCheckRegion = "WNAM"
	UserLoadBalancerPoolListResponseCheckRegionEnam       UserLoadBalancerPoolListResponseCheckRegion = "ENAM"
	UserLoadBalancerPoolListResponseCheckRegionWeu        UserLoadBalancerPoolListResponseCheckRegion = "WEU"
	UserLoadBalancerPoolListResponseCheckRegionEeu        UserLoadBalancerPoolListResponseCheckRegion = "EEU"
	UserLoadBalancerPoolListResponseCheckRegionNsam       UserLoadBalancerPoolListResponseCheckRegion = "NSAM"
	UserLoadBalancerPoolListResponseCheckRegionSsam       UserLoadBalancerPoolListResponseCheckRegion = "SSAM"
	UserLoadBalancerPoolListResponseCheckRegionOc         UserLoadBalancerPoolListResponseCheckRegion = "OC"
	UserLoadBalancerPoolListResponseCheckRegionMe         UserLoadBalancerPoolListResponseCheckRegion = "ME"
	UserLoadBalancerPoolListResponseCheckRegionNaf        UserLoadBalancerPoolListResponseCheckRegion = "NAF"
	UserLoadBalancerPoolListResponseCheckRegionSaf        UserLoadBalancerPoolListResponseCheckRegion = "SAF"
	UserLoadBalancerPoolListResponseCheckRegionSas        UserLoadBalancerPoolListResponseCheckRegion = "SAS"
	UserLoadBalancerPoolListResponseCheckRegionSeas       UserLoadBalancerPoolListResponseCheckRegion = "SEAS"
	UserLoadBalancerPoolListResponseCheckRegionNeas       UserLoadBalancerPoolListResponseCheckRegion = "NEAS"
	UserLoadBalancerPoolListResponseCheckRegionAllRegions UserLoadBalancerPoolListResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolListResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy UserLoadBalancerPoolListResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy UserLoadBalancerPoolListResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          userLoadBalancerPoolListResponseLoadSheddingJSON          `json:"-"`
}

// userLoadBalancerPoolListResponseLoadSheddingJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolListResponseLoadShedding]
type userLoadBalancerPoolListResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserLoadBalancerPoolListResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolListResponseLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolListResponseLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolListResponseLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolListResponseLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolListResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolListResponseLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolListResponseLoadSheddingSessionPolicyHash UserLoadBalancerPoolListResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolListResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin UserLoadBalancerPoolListResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool UserLoadBalancerPoolListResponseNotificationFilterPool `json:"pool,nullable"`
	JSON userLoadBalancerPoolListResponseNotificationFilterJSON `json:"-"`
}

// userLoadBalancerPoolListResponseNotificationFilterJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolListResponseNotificationFilter]
type userLoadBalancerPoolListResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolListResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolListResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                         `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolListResponseNotificationFilterOriginJSON `json:"-"`
}

// userLoadBalancerPoolListResponseNotificationFilterOriginJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerPoolListResponseNotificationFilterOrigin]
type userLoadBalancerPoolListResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolListResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolListResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                       `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolListResponseNotificationFilterPoolJSON `json:"-"`
}

// userLoadBalancerPoolListResponseNotificationFilterPoolJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolListResponseNotificationFilterPool]
type userLoadBalancerPoolListResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolListResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolListResponseOriginSteering struct {
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
	Policy UserLoadBalancerPoolListResponseOriginSteeringPolicy `json:"policy"`
	JSON   userLoadBalancerPoolListResponseOriginSteeringJSON   `json:"-"`
}

// userLoadBalancerPoolListResponseOriginSteeringJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolListResponseOriginSteering]
type userLoadBalancerPoolListResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolListResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type UserLoadBalancerPoolListResponseOriginSteeringPolicy string

const (
	UserLoadBalancerPoolListResponseOriginSteeringPolicyRandom                   UserLoadBalancerPoolListResponseOriginSteeringPolicy = "random"
	UserLoadBalancerPoolListResponseOriginSteeringPolicyHash                     UserLoadBalancerPoolListResponseOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolListResponseOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolListResponseOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolListResponseOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolListResponseOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolListResponseOrigin struct {
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
	Header UserLoadBalancerPoolListResponseOriginsHeader `json:"header"`
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
	Weight float64                                    `json:"weight"`
	JSON   userLoadBalancerPoolListResponseOriginJSON `json:"-"`
}

// userLoadBalancerPoolListResponseOriginJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolListResponseOrigin]
type userLoadBalancerPoolListResponseOriginJSON struct {
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

func (r *UserLoadBalancerPoolListResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolListResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                          `json:"Host"`
	JSON userLoadBalancerPoolListResponseOriginsHeaderJSON `json:"-"`
}

// userLoadBalancerPoolListResponseOriginsHeaderJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolListResponseOriginsHeader]
type userLoadBalancerPoolListResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolListResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolDeleteResponse struct {
	ID   string                                 `json:"id"`
	JSON userLoadBalancerPoolDeleteResponseJSON `json:"-"`
}

// userLoadBalancerPoolDeleteResponseJSON contains the JSON metadata for the struct
// [UserLoadBalancerPoolDeleteResponse]
type userLoadBalancerPoolDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolEditResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []UserLoadBalancerPoolEditResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                     `json:"created_on" format:"date-time"`
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
	LoadShedding UserLoadBalancerPoolEditResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter UserLoadBalancerPoolEditResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering UserLoadBalancerPoolEditResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []UserLoadBalancerPoolEditResponseOrigin `json:"origins"`
	JSON    userLoadBalancerPoolEditResponseJSON     `json:"-"`
}

// userLoadBalancerPoolEditResponseJSON contains the JSON metadata for the struct
// [UserLoadBalancerPoolEditResponse]
type userLoadBalancerPoolEditResponseJSON struct {
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

func (r *UserLoadBalancerPoolEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolEditResponseCheckRegion string

const (
	UserLoadBalancerPoolEditResponseCheckRegionWnam       UserLoadBalancerPoolEditResponseCheckRegion = "WNAM"
	UserLoadBalancerPoolEditResponseCheckRegionEnam       UserLoadBalancerPoolEditResponseCheckRegion = "ENAM"
	UserLoadBalancerPoolEditResponseCheckRegionWeu        UserLoadBalancerPoolEditResponseCheckRegion = "WEU"
	UserLoadBalancerPoolEditResponseCheckRegionEeu        UserLoadBalancerPoolEditResponseCheckRegion = "EEU"
	UserLoadBalancerPoolEditResponseCheckRegionNsam       UserLoadBalancerPoolEditResponseCheckRegion = "NSAM"
	UserLoadBalancerPoolEditResponseCheckRegionSsam       UserLoadBalancerPoolEditResponseCheckRegion = "SSAM"
	UserLoadBalancerPoolEditResponseCheckRegionOc         UserLoadBalancerPoolEditResponseCheckRegion = "OC"
	UserLoadBalancerPoolEditResponseCheckRegionMe         UserLoadBalancerPoolEditResponseCheckRegion = "ME"
	UserLoadBalancerPoolEditResponseCheckRegionNaf        UserLoadBalancerPoolEditResponseCheckRegion = "NAF"
	UserLoadBalancerPoolEditResponseCheckRegionSaf        UserLoadBalancerPoolEditResponseCheckRegion = "SAF"
	UserLoadBalancerPoolEditResponseCheckRegionSas        UserLoadBalancerPoolEditResponseCheckRegion = "SAS"
	UserLoadBalancerPoolEditResponseCheckRegionSeas       UserLoadBalancerPoolEditResponseCheckRegion = "SEAS"
	UserLoadBalancerPoolEditResponseCheckRegionNeas       UserLoadBalancerPoolEditResponseCheckRegion = "NEAS"
	UserLoadBalancerPoolEditResponseCheckRegionAllRegions UserLoadBalancerPoolEditResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolEditResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy UserLoadBalancerPoolEditResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy UserLoadBalancerPoolEditResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          userLoadBalancerPoolEditResponseLoadSheddingJSON          `json:"-"`
}

// userLoadBalancerPoolEditResponseLoadSheddingJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolEditResponseLoadShedding]
type userLoadBalancerPoolEditResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserLoadBalancerPoolEditResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolEditResponseLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolEditResponseLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolEditResponseLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolEditResponseLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolEditResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolEditResponseLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolEditResponseLoadSheddingSessionPolicyHash UserLoadBalancerPoolEditResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolEditResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin UserLoadBalancerPoolEditResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool UserLoadBalancerPoolEditResponseNotificationFilterPool `json:"pool,nullable"`
	JSON userLoadBalancerPoolEditResponseNotificationFilterJSON `json:"-"`
}

// userLoadBalancerPoolEditResponseNotificationFilterJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolEditResponseNotificationFilter]
type userLoadBalancerPoolEditResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolEditResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolEditResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                         `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolEditResponseNotificationFilterOriginJSON `json:"-"`
}

// userLoadBalancerPoolEditResponseNotificationFilterOriginJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerPoolEditResponseNotificationFilterOrigin]
type userLoadBalancerPoolEditResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolEditResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolEditResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                       `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolEditResponseNotificationFilterPoolJSON `json:"-"`
}

// userLoadBalancerPoolEditResponseNotificationFilterPoolJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolEditResponseNotificationFilterPool]
type userLoadBalancerPoolEditResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolEditResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolEditResponseOriginSteering struct {
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
	Policy UserLoadBalancerPoolEditResponseOriginSteeringPolicy `json:"policy"`
	JSON   userLoadBalancerPoolEditResponseOriginSteeringJSON   `json:"-"`
}

// userLoadBalancerPoolEditResponseOriginSteeringJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolEditResponseOriginSteering]
type userLoadBalancerPoolEditResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolEditResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type UserLoadBalancerPoolEditResponseOriginSteeringPolicy string

const (
	UserLoadBalancerPoolEditResponseOriginSteeringPolicyRandom                   UserLoadBalancerPoolEditResponseOriginSteeringPolicy = "random"
	UserLoadBalancerPoolEditResponseOriginSteeringPolicyHash                     UserLoadBalancerPoolEditResponseOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolEditResponseOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolEditResponseOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolEditResponseOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolEditResponseOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolEditResponseOrigin struct {
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
	Header UserLoadBalancerPoolEditResponseOriginsHeader `json:"header"`
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
	Weight float64                                    `json:"weight"`
	JSON   userLoadBalancerPoolEditResponseOriginJSON `json:"-"`
}

// userLoadBalancerPoolEditResponseOriginJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolEditResponseOrigin]
type userLoadBalancerPoolEditResponseOriginJSON struct {
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

func (r *UserLoadBalancerPoolEditResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolEditResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                          `json:"Host"`
	JSON userLoadBalancerPoolEditResponseOriginsHeaderJSON `json:"-"`
}

// userLoadBalancerPoolEditResponseOriginsHeaderJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolEditResponseOriginsHeader]
type userLoadBalancerPoolEditResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolEditResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolGetResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []UserLoadBalancerPoolGetResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                    `json:"created_on" format:"date-time"`
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
	LoadShedding UserLoadBalancerPoolGetResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter UserLoadBalancerPoolGetResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering UserLoadBalancerPoolGetResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []UserLoadBalancerPoolGetResponseOrigin `json:"origins"`
	JSON    userLoadBalancerPoolGetResponseJSON     `json:"-"`
}

// userLoadBalancerPoolGetResponseJSON contains the JSON metadata for the struct
// [UserLoadBalancerPoolGetResponse]
type userLoadBalancerPoolGetResponseJSON struct {
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

func (r *UserLoadBalancerPoolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolGetResponseCheckRegion string

const (
	UserLoadBalancerPoolGetResponseCheckRegionWnam       UserLoadBalancerPoolGetResponseCheckRegion = "WNAM"
	UserLoadBalancerPoolGetResponseCheckRegionEnam       UserLoadBalancerPoolGetResponseCheckRegion = "ENAM"
	UserLoadBalancerPoolGetResponseCheckRegionWeu        UserLoadBalancerPoolGetResponseCheckRegion = "WEU"
	UserLoadBalancerPoolGetResponseCheckRegionEeu        UserLoadBalancerPoolGetResponseCheckRegion = "EEU"
	UserLoadBalancerPoolGetResponseCheckRegionNsam       UserLoadBalancerPoolGetResponseCheckRegion = "NSAM"
	UserLoadBalancerPoolGetResponseCheckRegionSsam       UserLoadBalancerPoolGetResponseCheckRegion = "SSAM"
	UserLoadBalancerPoolGetResponseCheckRegionOc         UserLoadBalancerPoolGetResponseCheckRegion = "OC"
	UserLoadBalancerPoolGetResponseCheckRegionMe         UserLoadBalancerPoolGetResponseCheckRegion = "ME"
	UserLoadBalancerPoolGetResponseCheckRegionNaf        UserLoadBalancerPoolGetResponseCheckRegion = "NAF"
	UserLoadBalancerPoolGetResponseCheckRegionSaf        UserLoadBalancerPoolGetResponseCheckRegion = "SAF"
	UserLoadBalancerPoolGetResponseCheckRegionSas        UserLoadBalancerPoolGetResponseCheckRegion = "SAS"
	UserLoadBalancerPoolGetResponseCheckRegionSeas       UserLoadBalancerPoolGetResponseCheckRegion = "SEAS"
	UserLoadBalancerPoolGetResponseCheckRegionNeas       UserLoadBalancerPoolGetResponseCheckRegion = "NEAS"
	UserLoadBalancerPoolGetResponseCheckRegionAllRegions UserLoadBalancerPoolGetResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolGetResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy UserLoadBalancerPoolGetResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy UserLoadBalancerPoolGetResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          userLoadBalancerPoolGetResponseLoadSheddingJSON          `json:"-"`
}

// userLoadBalancerPoolGetResponseLoadSheddingJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolGetResponseLoadShedding]
type userLoadBalancerPoolGetResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserLoadBalancerPoolGetResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolGetResponseLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolGetResponseLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolGetResponseLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolGetResponseLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolGetResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolGetResponseLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolGetResponseLoadSheddingSessionPolicyHash UserLoadBalancerPoolGetResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolGetResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin UserLoadBalancerPoolGetResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool UserLoadBalancerPoolGetResponseNotificationFilterPool `json:"pool,nullable"`
	JSON userLoadBalancerPoolGetResponseNotificationFilterJSON `json:"-"`
}

// userLoadBalancerPoolGetResponseNotificationFilterJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolGetResponseNotificationFilter]
type userLoadBalancerPoolGetResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolGetResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolGetResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                        `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolGetResponseNotificationFilterOriginJSON `json:"-"`
}

// userLoadBalancerPoolGetResponseNotificationFilterOriginJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerPoolGetResponseNotificationFilterOrigin]
type userLoadBalancerPoolGetResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolGetResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolGetResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                      `json:"healthy,nullable"`
	JSON    userLoadBalancerPoolGetResponseNotificationFilterPoolJSON `json:"-"`
}

// userLoadBalancerPoolGetResponseNotificationFilterPoolJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolGetResponseNotificationFilterPool]
type userLoadBalancerPoolGetResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolGetResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolGetResponseOriginSteering struct {
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
	Policy UserLoadBalancerPoolGetResponseOriginSteeringPolicy `json:"policy"`
	JSON   userLoadBalancerPoolGetResponseOriginSteeringJSON   `json:"-"`
}

// userLoadBalancerPoolGetResponseOriginSteeringJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolGetResponseOriginSteering]
type userLoadBalancerPoolGetResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolGetResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type UserLoadBalancerPoolGetResponseOriginSteeringPolicy string

const (
	UserLoadBalancerPoolGetResponseOriginSteeringPolicyRandom                   UserLoadBalancerPoolGetResponseOriginSteeringPolicy = "random"
	UserLoadBalancerPoolGetResponseOriginSteeringPolicyHash                     UserLoadBalancerPoolGetResponseOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolGetResponseOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolGetResponseOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolGetResponseOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolGetResponseOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolGetResponseOrigin struct {
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
	Header UserLoadBalancerPoolGetResponseOriginsHeader `json:"header"`
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
	Weight float64                                   `json:"weight"`
	JSON   userLoadBalancerPoolGetResponseOriginJSON `json:"-"`
}

// userLoadBalancerPoolGetResponseOriginJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolGetResponseOrigin]
type userLoadBalancerPoolGetResponseOriginJSON struct {
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

func (r *UserLoadBalancerPoolGetResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolGetResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                         `json:"Host"`
	JSON userLoadBalancerPoolGetResponseOriginsHeaderJSON `json:"-"`
}

// userLoadBalancerPoolGetResponseOriginsHeaderJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolGetResponseOriginsHeader]
type userLoadBalancerPoolGetResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolGetResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of regions from which to run health checks. Null means every Cloudflare
// data center.
//
// Union satisfied by [UserLoadBalancerPoolHealthResponseUnknown] or
// [shared.UnionString].
type UserLoadBalancerPoolHealthResponse interface {
	ImplementsUserLoadBalancerPoolHealthResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserLoadBalancerPoolHealthResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserLoadBalancerPoolPreviewResponse struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string                       `json:"pools"`
	PreviewID string                                  `json:"preview_id"`
	JSON      userLoadBalancerPoolPreviewResponseJSON `json:"-"`
}

// userLoadBalancerPoolPreviewResponseJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolPreviewResponse]
type userLoadBalancerPoolPreviewResponseJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferencesResponse struct {
	ReferenceType UserLoadBalancerPoolReferencesResponseReferenceType `json:"reference_type"`
	ResourceID    string                                              `json:"resource_id"`
	ResourceName  string                                              `json:"resource_name"`
	ResourceType  string                                              `json:"resource_type"`
	JSON          userLoadBalancerPoolReferencesResponseJSON          `json:"-"`
}

// userLoadBalancerPoolReferencesResponseJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolReferencesResponse]
type userLoadBalancerPoolReferencesResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferencesResponseReferenceType string

const (
	UserLoadBalancerPoolReferencesResponseReferenceTypeStar     UserLoadBalancerPoolReferencesResponseReferenceType = "*"
	UserLoadBalancerPoolReferencesResponseReferenceTypeReferral UserLoadBalancerPoolReferencesResponseReferenceType = "referral"
	UserLoadBalancerPoolReferencesResponseReferenceTypeReferrer UserLoadBalancerPoolReferencesResponseReferenceType = "referrer"
)

type UserLoadBalancerPoolNewParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]UserLoadBalancerPoolNewParamsOrigin] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]UserLoadBalancerPoolNewParamsCheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[UserLoadBalancerPoolNewParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[UserLoadBalancerPoolNewParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[UserLoadBalancerPoolNewParamsOriginSteering] `json:"origin_steering"`
}

func (r UserLoadBalancerPoolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserLoadBalancerPoolNewParamsOrigin struct {
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
	Header param.Field[UserLoadBalancerPoolNewParamsOriginsHeader] `json:"header"`
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

func (r UserLoadBalancerPoolNewParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolNewParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r UserLoadBalancerPoolNewParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolNewParamsCheckRegion string

const (
	UserLoadBalancerPoolNewParamsCheckRegionWnam       UserLoadBalancerPoolNewParamsCheckRegion = "WNAM"
	UserLoadBalancerPoolNewParamsCheckRegionEnam       UserLoadBalancerPoolNewParamsCheckRegion = "ENAM"
	UserLoadBalancerPoolNewParamsCheckRegionWeu        UserLoadBalancerPoolNewParamsCheckRegion = "WEU"
	UserLoadBalancerPoolNewParamsCheckRegionEeu        UserLoadBalancerPoolNewParamsCheckRegion = "EEU"
	UserLoadBalancerPoolNewParamsCheckRegionNsam       UserLoadBalancerPoolNewParamsCheckRegion = "NSAM"
	UserLoadBalancerPoolNewParamsCheckRegionSsam       UserLoadBalancerPoolNewParamsCheckRegion = "SSAM"
	UserLoadBalancerPoolNewParamsCheckRegionOc         UserLoadBalancerPoolNewParamsCheckRegion = "OC"
	UserLoadBalancerPoolNewParamsCheckRegionMe         UserLoadBalancerPoolNewParamsCheckRegion = "ME"
	UserLoadBalancerPoolNewParamsCheckRegionNaf        UserLoadBalancerPoolNewParamsCheckRegion = "NAF"
	UserLoadBalancerPoolNewParamsCheckRegionSaf        UserLoadBalancerPoolNewParamsCheckRegion = "SAF"
	UserLoadBalancerPoolNewParamsCheckRegionSas        UserLoadBalancerPoolNewParamsCheckRegion = "SAS"
	UserLoadBalancerPoolNewParamsCheckRegionSeas       UserLoadBalancerPoolNewParamsCheckRegion = "SEAS"
	UserLoadBalancerPoolNewParamsCheckRegionNeas       UserLoadBalancerPoolNewParamsCheckRegion = "NEAS"
	UserLoadBalancerPoolNewParamsCheckRegionAllRegions UserLoadBalancerPoolNewParamsCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolNewParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[UserLoadBalancerPoolNewParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[UserLoadBalancerPoolNewParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r UserLoadBalancerPoolNewParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolNewParamsLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolNewParamsLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolNewParamsLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolNewParamsLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolNewParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolNewParamsLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolNewParamsLoadSheddingSessionPolicyHash UserLoadBalancerPoolNewParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolNewParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[UserLoadBalancerPoolNewParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[UserLoadBalancerPoolNewParamsNotificationFilterPool] `json:"pool"`
}

func (r UserLoadBalancerPoolNewParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolNewParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r UserLoadBalancerPoolNewParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolNewParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r UserLoadBalancerPoolNewParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolNewParamsOriginSteering struct {
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
	Policy param.Field[UserLoadBalancerPoolNewParamsOriginSteeringPolicy] `json:"policy"`
}

func (r UserLoadBalancerPoolNewParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type UserLoadBalancerPoolNewParamsOriginSteeringPolicy string

const (
	UserLoadBalancerPoolNewParamsOriginSteeringPolicyRandom                   UserLoadBalancerPoolNewParamsOriginSteeringPolicy = "random"
	UserLoadBalancerPoolNewParamsOriginSteeringPolicyHash                     UserLoadBalancerPoolNewParamsOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolNewParamsOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolNewParamsOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolNewParamsOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolNewParamsOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolNewResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolNewResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerPoolNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolNewResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolNewResponseEnvelope]
type userLoadBalancerPoolNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    userLoadBalancerPoolNewResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolNewResponseEnvelopeErrors]
type userLoadBalancerPoolNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    userLoadBalancerPoolNewResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolNewResponseEnvelopeMessages]
type userLoadBalancerPoolNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolNewResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolNewResponseEnvelopeSuccessTrue UserLoadBalancerPoolNewResponseEnvelopeSuccess = true
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

type UserLoadBalancerPoolUpdateResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerPoolUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolUpdateResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolUpdateResponseEnvelope]
type userLoadBalancerPoolUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    userLoadBalancerPoolUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolUpdateResponseEnvelopeErrors]
type userLoadBalancerPoolUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    userLoadBalancerPoolUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolUpdateResponseEnvelopeMessages]
type userLoadBalancerPoolUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolUpdateResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolUpdateResponseEnvelopeSuccessTrue UserLoadBalancerPoolUpdateResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolListParams struct {
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `query:"monitor"`
}

// URLQuery serializes [UserLoadBalancerPoolListParams]'s query parameters as
// `url.Values`.
func (r UserLoadBalancerPoolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserLoadBalancerPoolListResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolListResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserLoadBalancerPoolListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserLoadBalancerPoolListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancerPoolListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userLoadBalancerPoolListResponseEnvelopeJSON       `json:"-"`
}

// userLoadBalancerPoolListResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolListResponseEnvelope]
type userLoadBalancerPoolListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    userLoadBalancerPoolListResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolListResponseEnvelopeErrors]
type userLoadBalancerPoolListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    userLoadBalancerPoolListResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolListResponseEnvelopeMessages]
type userLoadBalancerPoolListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolListResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolListResponseEnvelopeSuccessTrue UserLoadBalancerPoolListResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       userLoadBalancerPoolListResponseEnvelopeResultInfoJSON `json:"-"`
}

// userLoadBalancerPoolListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolListResponseEnvelopeResultInfo]
type userLoadBalancerPoolListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolDeleteResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerPoolDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolDeleteResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolDeleteResponseEnvelope]
type userLoadBalancerPoolDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    userLoadBalancerPoolDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolDeleteResponseEnvelopeErrors]
type userLoadBalancerPoolDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    userLoadBalancerPoolDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolDeleteResponseEnvelopeMessages]
type userLoadBalancerPoolDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolDeleteResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolDeleteResponseEnvelopeSuccessTrue UserLoadBalancerPoolDeleteResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolEditParams struct {
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]UserLoadBalancerPoolEditParamsCheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[UserLoadBalancerPoolEditParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[UserLoadBalancerPoolEditParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[UserLoadBalancerPoolEditParamsOriginSteering] `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]UserLoadBalancerPoolEditParamsOrigin] `json:"origins"`
}

func (r UserLoadBalancerPoolEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type UserLoadBalancerPoolEditParamsCheckRegion string

const (
	UserLoadBalancerPoolEditParamsCheckRegionWnam       UserLoadBalancerPoolEditParamsCheckRegion = "WNAM"
	UserLoadBalancerPoolEditParamsCheckRegionEnam       UserLoadBalancerPoolEditParamsCheckRegion = "ENAM"
	UserLoadBalancerPoolEditParamsCheckRegionWeu        UserLoadBalancerPoolEditParamsCheckRegion = "WEU"
	UserLoadBalancerPoolEditParamsCheckRegionEeu        UserLoadBalancerPoolEditParamsCheckRegion = "EEU"
	UserLoadBalancerPoolEditParamsCheckRegionNsam       UserLoadBalancerPoolEditParamsCheckRegion = "NSAM"
	UserLoadBalancerPoolEditParamsCheckRegionSsam       UserLoadBalancerPoolEditParamsCheckRegion = "SSAM"
	UserLoadBalancerPoolEditParamsCheckRegionOc         UserLoadBalancerPoolEditParamsCheckRegion = "OC"
	UserLoadBalancerPoolEditParamsCheckRegionMe         UserLoadBalancerPoolEditParamsCheckRegion = "ME"
	UserLoadBalancerPoolEditParamsCheckRegionNaf        UserLoadBalancerPoolEditParamsCheckRegion = "NAF"
	UserLoadBalancerPoolEditParamsCheckRegionSaf        UserLoadBalancerPoolEditParamsCheckRegion = "SAF"
	UserLoadBalancerPoolEditParamsCheckRegionSas        UserLoadBalancerPoolEditParamsCheckRegion = "SAS"
	UserLoadBalancerPoolEditParamsCheckRegionSeas       UserLoadBalancerPoolEditParamsCheckRegion = "SEAS"
	UserLoadBalancerPoolEditParamsCheckRegionNeas       UserLoadBalancerPoolEditParamsCheckRegion = "NEAS"
	UserLoadBalancerPoolEditParamsCheckRegionAllRegions UserLoadBalancerPoolEditParamsCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type UserLoadBalancerPoolEditParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[UserLoadBalancerPoolEditParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[UserLoadBalancerPoolEditParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r UserLoadBalancerPoolEditParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type UserLoadBalancerPoolEditParamsLoadSheddingDefaultPolicy string

const (
	UserLoadBalancerPoolEditParamsLoadSheddingDefaultPolicyRandom UserLoadBalancerPoolEditParamsLoadSheddingDefaultPolicy = "random"
	UserLoadBalancerPoolEditParamsLoadSheddingDefaultPolicyHash   UserLoadBalancerPoolEditParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type UserLoadBalancerPoolEditParamsLoadSheddingSessionPolicy string

const (
	UserLoadBalancerPoolEditParamsLoadSheddingSessionPolicyHash UserLoadBalancerPoolEditParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type UserLoadBalancerPoolEditParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[UserLoadBalancerPoolEditParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[UserLoadBalancerPoolEditParamsNotificationFilterPool] `json:"pool"`
}

func (r UserLoadBalancerPoolEditParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolEditParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r UserLoadBalancerPoolEditParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type UserLoadBalancerPoolEditParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r UserLoadBalancerPoolEditParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type UserLoadBalancerPoolEditParamsOriginSteering struct {
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
	Policy param.Field[UserLoadBalancerPoolEditParamsOriginSteeringPolicy] `json:"policy"`
}

func (r UserLoadBalancerPoolEditParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type UserLoadBalancerPoolEditParamsOriginSteeringPolicy string

const (
	UserLoadBalancerPoolEditParamsOriginSteeringPolicyRandom                   UserLoadBalancerPoolEditParamsOriginSteeringPolicy = "random"
	UserLoadBalancerPoolEditParamsOriginSteeringPolicyHash                     UserLoadBalancerPoolEditParamsOriginSteeringPolicy = "hash"
	UserLoadBalancerPoolEditParamsOriginSteeringPolicyLeastOutstandingRequests UserLoadBalancerPoolEditParamsOriginSteeringPolicy = "least_outstanding_requests"
	UserLoadBalancerPoolEditParamsOriginSteeringPolicyLeastConnections         UserLoadBalancerPoolEditParamsOriginSteeringPolicy = "least_connections"
)

type UserLoadBalancerPoolEditParamsOrigin struct {
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
	Header param.Field[UserLoadBalancerPoolEditParamsOriginsHeader] `json:"header"`
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

func (r UserLoadBalancerPoolEditParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type UserLoadBalancerPoolEditParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r UserLoadBalancerPoolEditParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserLoadBalancerPoolEditResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolEditResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerPoolEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolEditResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolEditResponseEnvelope]
type userLoadBalancerPoolEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    userLoadBalancerPoolEditResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolEditResponseEnvelopeErrors]
type userLoadBalancerPoolEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    userLoadBalancerPoolEditResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolEditResponseEnvelopeMessages]
type userLoadBalancerPoolEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolEditResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolEditResponseEnvelopeSuccessTrue UserLoadBalancerPoolEditResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolGetResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolGetResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerPoolGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolGetResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserLoadBalancerPoolGetResponseEnvelope]
type userLoadBalancerPoolGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    userLoadBalancerPoolGetResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolGetResponseEnvelopeErrors]
type userLoadBalancerPoolGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    userLoadBalancerPoolGetResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolGetResponseEnvelopeMessages]
type userLoadBalancerPoolGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolGetResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolGetResponseEnvelopeSuccessTrue UserLoadBalancerPoolGetResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolHealthResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolHealthResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolHealthResponseEnvelopeMessages `json:"messages,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result UserLoadBalancerPoolHealthResponse `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolHealthResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolHealthResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolHealthResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolHealthResponseEnvelope]
type userLoadBalancerPoolHealthResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolHealthResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    userLoadBalancerPoolHealthResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolHealthResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolHealthResponseEnvelopeErrors]
type userLoadBalancerPoolHealthResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolHealthResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    userLoadBalancerPoolHealthResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolHealthResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolHealthResponseEnvelopeMessages]
type userLoadBalancerPoolHealthResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolHealthResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolHealthResponseEnvelopeSuccessTrue UserLoadBalancerPoolHealthResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolPreviewParams struct {
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[interface{}] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[UserLoadBalancerPoolPreviewParamsType] `json:"type"`
}

func (r UserLoadBalancerPoolPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type UserLoadBalancerPoolPreviewParamsType string

const (
	UserLoadBalancerPoolPreviewParamsTypeHTTP     UserLoadBalancerPoolPreviewParamsType = "http"
	UserLoadBalancerPoolPreviewParamsTypeHTTPS    UserLoadBalancerPoolPreviewParamsType = "https"
	UserLoadBalancerPoolPreviewParamsTypeTcp      UserLoadBalancerPoolPreviewParamsType = "tcp"
	UserLoadBalancerPoolPreviewParamsTypeUdpIcmp  UserLoadBalancerPoolPreviewParamsType = "udp_icmp"
	UserLoadBalancerPoolPreviewParamsTypeIcmpPing UserLoadBalancerPoolPreviewParamsType = "icmp_ping"
	UserLoadBalancerPoolPreviewParamsTypeSmtp     UserLoadBalancerPoolPreviewParamsType = "smtp"
)

type UserLoadBalancerPoolPreviewResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolPreviewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolPreviewResponseEnvelopeMessages `json:"messages,required"`
	Result   UserLoadBalancerPoolPreviewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolPreviewResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolPreviewResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolPreviewResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserLoadBalancerPoolPreviewResponseEnvelope]
type userLoadBalancerPoolPreviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolPreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolPreviewResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    userLoadBalancerPoolPreviewResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolPreviewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolPreviewResponseEnvelopeErrors]
type userLoadBalancerPoolPreviewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolPreviewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolPreviewResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    userLoadBalancerPoolPreviewResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolPreviewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolPreviewResponseEnvelopeMessages]
type userLoadBalancerPoolPreviewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolPreviewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolPreviewResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolPreviewResponseEnvelopeSuccessTrue UserLoadBalancerPoolPreviewResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolReferencesResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolReferencesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolReferencesResponseEnvelopeMessages `json:"messages,required"`
	// List of resources that reference a given pool.
	Result []UserLoadBalancerPoolReferencesResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserLoadBalancerPoolReferencesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancerPoolReferencesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userLoadBalancerPoolReferencesResponseEnvelopeJSON       `json:"-"`
}

// userLoadBalancerPoolReferencesResponseEnvelopeJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolReferencesResponseEnvelope]
type userLoadBalancerPoolReferencesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferencesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferencesResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    userLoadBalancerPoolReferencesResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolReferencesResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolReferencesResponseEnvelopeErrors]
type userLoadBalancerPoolReferencesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferencesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferencesResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    userLoadBalancerPoolReferencesResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolReferencesResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolReferencesResponseEnvelopeMessages]
type userLoadBalancerPoolReferencesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferencesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolReferencesResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolReferencesResponseEnvelopeSuccessTrue UserLoadBalancerPoolReferencesResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolReferencesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       userLoadBalancerPoolReferencesResponseEnvelopeResultInfoJSON `json:"-"`
}

// userLoadBalancerPoolReferencesResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerPoolReferencesResponseEnvelopeResultInfo]
type userLoadBalancerPoolReferencesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferencesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
