// File generated from our OpenAPI spec by Stainless.

package load_balancers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// PoolService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPoolService] method instead.
type PoolService struct {
	Options    []option.RequestOption
	Health     *PoolHealthService
	References *PoolReferenceService
}

// NewPoolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPoolService(opts ...option.RequestOption) (r *PoolService) {
	r = &PoolService{}
	r.Options = opts
	r.Health = NewPoolHealthService(opts...)
	r.References = NewPoolReferenceService(opts...)
	return
}

// Create a new pool.
func (r *PoolService) New(ctx context.Context, params PoolNewParams, opts ...option.RequestOption) (res *PoolNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured pool.
func (r *PoolService) Update(ctx context.Context, poolID string, params PoolUpdateParams, opts ...option.RequestOption) (res *PoolUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", params.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured pools.
func (r *PoolService) List(ctx context.Context, params PoolListParams, opts ...option.RequestOption) (res *[]PoolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured pool.
func (r *PoolService) Delete(ctx context.Context, poolID string, body PoolDeleteParams, opts ...option.RequestOption) (res *PoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", body.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing pool, overwriting the supplied properties.
func (r *PoolService) Edit(ctx context.Context, poolID string, params PoolEditParams, opts ...option.RequestOption) (res *PoolEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", params.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured pool.
func (r *PoolService) Get(ctx context.Context, poolID string, query PoolGetParams, opts ...option.RequestOption) (res *PoolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", query.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PoolNewResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []PoolNewResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                    `json:"created_on" format:"date-time"`
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
	LoadShedding PoolNewResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter PoolNewResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering PoolNewResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []PoolNewResponseOrigin `json:"origins"`
	JSON    poolNewResponseJSON     `json:"-"`
}

// poolNewResponseJSON contains the JSON metadata for the struct [PoolNewResponse]
type poolNewResponseJSON struct {
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

func (r *PoolNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolNewResponseCheckRegion string

const (
	PoolNewResponseCheckRegionWnam       PoolNewResponseCheckRegion = "WNAM"
	PoolNewResponseCheckRegionEnam       PoolNewResponseCheckRegion = "ENAM"
	PoolNewResponseCheckRegionWeu        PoolNewResponseCheckRegion = "WEU"
	PoolNewResponseCheckRegionEeu        PoolNewResponseCheckRegion = "EEU"
	PoolNewResponseCheckRegionNsam       PoolNewResponseCheckRegion = "NSAM"
	PoolNewResponseCheckRegionSsam       PoolNewResponseCheckRegion = "SSAM"
	PoolNewResponseCheckRegionOc         PoolNewResponseCheckRegion = "OC"
	PoolNewResponseCheckRegionMe         PoolNewResponseCheckRegion = "ME"
	PoolNewResponseCheckRegionNaf        PoolNewResponseCheckRegion = "NAF"
	PoolNewResponseCheckRegionSaf        PoolNewResponseCheckRegion = "SAF"
	PoolNewResponseCheckRegionSas        PoolNewResponseCheckRegion = "SAS"
	PoolNewResponseCheckRegionSeas       PoolNewResponseCheckRegion = "SEAS"
	PoolNewResponseCheckRegionNeas       PoolNewResponseCheckRegion = "NEAS"
	PoolNewResponseCheckRegionAllRegions PoolNewResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type PoolNewResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy PoolNewResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy PoolNewResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          poolNewResponseLoadSheddingJSON          `json:"-"`
}

// poolNewResponseLoadSheddingJSON contains the JSON metadata for the struct
// [PoolNewResponseLoadShedding]
type poolNewResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PoolNewResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseLoadSheddingJSON) RawJSON() string {
	return r.raw
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolNewResponseLoadSheddingDefaultPolicy string

const (
	PoolNewResponseLoadSheddingDefaultPolicyRandom PoolNewResponseLoadSheddingDefaultPolicy = "random"
	PoolNewResponseLoadSheddingDefaultPolicyHash   PoolNewResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolNewResponseLoadSheddingSessionPolicy string

const (
	PoolNewResponseLoadSheddingSessionPolicyHash PoolNewResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolNewResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin PoolNewResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool PoolNewResponseNotificationFilterPool `json:"pool,nullable"`
	JSON poolNewResponseNotificationFilterJSON `json:"-"`
}

// poolNewResponseNotificationFilterJSON contains the JSON metadata for the struct
// [PoolNewResponseNotificationFilter]
type poolNewResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNewResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseNotificationFilterJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolNewResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                        `json:"healthy,nullable"`
	JSON    poolNewResponseNotificationFilterOriginJSON `json:"-"`
}

// poolNewResponseNotificationFilterOriginJSON contains the JSON metadata for the
// struct [PoolNewResponseNotificationFilterOrigin]
type poolNewResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNewResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseNotificationFilterOriginJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolNewResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                      `json:"healthy,nullable"`
	JSON    poolNewResponseNotificationFilterPoolJSON `json:"-"`
}

// poolNewResponseNotificationFilterPoolJSON contains the JSON metadata for the
// struct [PoolNewResponseNotificationFilterPool]
type poolNewResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNewResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseNotificationFilterPoolJSON) RawJSON() string {
	return r.raw
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolNewResponseOriginSteering struct {
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
	Policy PoolNewResponseOriginSteeringPolicy `json:"policy"`
	JSON   poolNewResponseOriginSteeringJSON   `json:"-"`
}

// poolNewResponseOriginSteeringJSON contains the JSON metadata for the struct
// [PoolNewResponseOriginSteering]
type poolNewResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNewResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseOriginSteeringJSON) RawJSON() string {
	return r.raw
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
type PoolNewResponseOriginSteeringPolicy string

const (
	PoolNewResponseOriginSteeringPolicyRandom                   PoolNewResponseOriginSteeringPolicy = "random"
	PoolNewResponseOriginSteeringPolicyHash                     PoolNewResponseOriginSteeringPolicy = "hash"
	PoolNewResponseOriginSteeringPolicyLeastOutstandingRequests PoolNewResponseOriginSteeringPolicy = "least_outstanding_requests"
	PoolNewResponseOriginSteeringPolicyLeastConnections         PoolNewResponseOriginSteeringPolicy = "least_connections"
)

type PoolNewResponseOrigin struct {
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
	Header PoolNewResponseOriginsHeader `json:"header"`
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
	Weight float64                   `json:"weight"`
	JSON   poolNewResponseOriginJSON `json:"-"`
}

// poolNewResponseOriginJSON contains the JSON metadata for the struct
// [PoolNewResponseOrigin]
type poolNewResponseOriginJSON struct {
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

func (r *PoolNewResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseOriginJSON) RawJSON() string {
	return r.raw
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolNewResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                         `json:"Host"`
	JSON poolNewResponseOriginsHeaderJSON `json:"-"`
}

// poolNewResponseOriginsHeaderJSON contains the JSON metadata for the struct
// [PoolNewResponseOriginsHeader]
type poolNewResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNewResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseOriginsHeaderJSON) RawJSON() string {
	return r.raw
}

type PoolUpdateResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []PoolUpdateResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                       `json:"created_on" format:"date-time"`
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
	LoadShedding PoolUpdateResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter PoolUpdateResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering PoolUpdateResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []PoolUpdateResponseOrigin `json:"origins"`
	JSON    poolUpdateResponseJSON     `json:"-"`
}

// poolUpdateResponseJSON contains the JSON metadata for the struct
// [PoolUpdateResponse]
type poolUpdateResponseJSON struct {
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

func (r *PoolUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolUpdateResponseCheckRegion string

const (
	PoolUpdateResponseCheckRegionWnam       PoolUpdateResponseCheckRegion = "WNAM"
	PoolUpdateResponseCheckRegionEnam       PoolUpdateResponseCheckRegion = "ENAM"
	PoolUpdateResponseCheckRegionWeu        PoolUpdateResponseCheckRegion = "WEU"
	PoolUpdateResponseCheckRegionEeu        PoolUpdateResponseCheckRegion = "EEU"
	PoolUpdateResponseCheckRegionNsam       PoolUpdateResponseCheckRegion = "NSAM"
	PoolUpdateResponseCheckRegionSsam       PoolUpdateResponseCheckRegion = "SSAM"
	PoolUpdateResponseCheckRegionOc         PoolUpdateResponseCheckRegion = "OC"
	PoolUpdateResponseCheckRegionMe         PoolUpdateResponseCheckRegion = "ME"
	PoolUpdateResponseCheckRegionNaf        PoolUpdateResponseCheckRegion = "NAF"
	PoolUpdateResponseCheckRegionSaf        PoolUpdateResponseCheckRegion = "SAF"
	PoolUpdateResponseCheckRegionSas        PoolUpdateResponseCheckRegion = "SAS"
	PoolUpdateResponseCheckRegionSeas       PoolUpdateResponseCheckRegion = "SEAS"
	PoolUpdateResponseCheckRegionNeas       PoolUpdateResponseCheckRegion = "NEAS"
	PoolUpdateResponseCheckRegionAllRegions PoolUpdateResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type PoolUpdateResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy PoolUpdateResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy PoolUpdateResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          poolUpdateResponseLoadSheddingJSON          `json:"-"`
}

// poolUpdateResponseLoadSheddingJSON contains the JSON metadata for the struct
// [PoolUpdateResponseLoadShedding]
type poolUpdateResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PoolUpdateResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseLoadSheddingJSON) RawJSON() string {
	return r.raw
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolUpdateResponseLoadSheddingDefaultPolicy string

const (
	PoolUpdateResponseLoadSheddingDefaultPolicyRandom PoolUpdateResponseLoadSheddingDefaultPolicy = "random"
	PoolUpdateResponseLoadSheddingDefaultPolicyHash   PoolUpdateResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolUpdateResponseLoadSheddingSessionPolicy string

const (
	PoolUpdateResponseLoadSheddingSessionPolicyHash PoolUpdateResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolUpdateResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin PoolUpdateResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool PoolUpdateResponseNotificationFilterPool `json:"pool,nullable"`
	JSON poolUpdateResponseNotificationFilterJSON `json:"-"`
}

// poolUpdateResponseNotificationFilterJSON contains the JSON metadata for the
// struct [PoolUpdateResponseNotificationFilter]
type poolUpdateResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolUpdateResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseNotificationFilterJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolUpdateResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                           `json:"healthy,nullable"`
	JSON    poolUpdateResponseNotificationFilterOriginJSON `json:"-"`
}

// poolUpdateResponseNotificationFilterOriginJSON contains the JSON metadata for
// the struct [PoolUpdateResponseNotificationFilterOrigin]
type poolUpdateResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolUpdateResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseNotificationFilterOriginJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolUpdateResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                         `json:"healthy,nullable"`
	JSON    poolUpdateResponseNotificationFilterPoolJSON `json:"-"`
}

// poolUpdateResponseNotificationFilterPoolJSON contains the JSON metadata for the
// struct [PoolUpdateResponseNotificationFilterPool]
type poolUpdateResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolUpdateResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseNotificationFilterPoolJSON) RawJSON() string {
	return r.raw
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolUpdateResponseOriginSteering struct {
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
	Policy PoolUpdateResponseOriginSteeringPolicy `json:"policy"`
	JSON   poolUpdateResponseOriginSteeringJSON   `json:"-"`
}

// poolUpdateResponseOriginSteeringJSON contains the JSON metadata for the struct
// [PoolUpdateResponseOriginSteering]
type poolUpdateResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolUpdateResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseOriginSteeringJSON) RawJSON() string {
	return r.raw
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
type PoolUpdateResponseOriginSteeringPolicy string

const (
	PoolUpdateResponseOriginSteeringPolicyRandom                   PoolUpdateResponseOriginSteeringPolicy = "random"
	PoolUpdateResponseOriginSteeringPolicyHash                     PoolUpdateResponseOriginSteeringPolicy = "hash"
	PoolUpdateResponseOriginSteeringPolicyLeastOutstandingRequests PoolUpdateResponseOriginSteeringPolicy = "least_outstanding_requests"
	PoolUpdateResponseOriginSteeringPolicyLeastConnections         PoolUpdateResponseOriginSteeringPolicy = "least_connections"
)

type PoolUpdateResponseOrigin struct {
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
	Header PoolUpdateResponseOriginsHeader `json:"header"`
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
	Weight float64                      `json:"weight"`
	JSON   poolUpdateResponseOriginJSON `json:"-"`
}

// poolUpdateResponseOriginJSON contains the JSON metadata for the struct
// [PoolUpdateResponseOrigin]
type poolUpdateResponseOriginJSON struct {
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

func (r *PoolUpdateResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseOriginJSON) RawJSON() string {
	return r.raw
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolUpdateResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                            `json:"Host"`
	JSON poolUpdateResponseOriginsHeaderJSON `json:"-"`
}

// poolUpdateResponseOriginsHeaderJSON contains the JSON metadata for the struct
// [PoolUpdateResponseOriginsHeader]
type poolUpdateResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolUpdateResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseOriginsHeaderJSON) RawJSON() string {
	return r.raw
}

type PoolListResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []PoolListResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                     `json:"created_on" format:"date-time"`
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
	LoadShedding PoolListResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter PoolListResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering PoolListResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []PoolListResponseOrigin `json:"origins"`
	JSON    poolListResponseJSON     `json:"-"`
}

// poolListResponseJSON contains the JSON metadata for the struct
// [PoolListResponse]
type poolListResponseJSON struct {
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

func (r *PoolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolListResponseCheckRegion string

const (
	PoolListResponseCheckRegionWnam       PoolListResponseCheckRegion = "WNAM"
	PoolListResponseCheckRegionEnam       PoolListResponseCheckRegion = "ENAM"
	PoolListResponseCheckRegionWeu        PoolListResponseCheckRegion = "WEU"
	PoolListResponseCheckRegionEeu        PoolListResponseCheckRegion = "EEU"
	PoolListResponseCheckRegionNsam       PoolListResponseCheckRegion = "NSAM"
	PoolListResponseCheckRegionSsam       PoolListResponseCheckRegion = "SSAM"
	PoolListResponseCheckRegionOc         PoolListResponseCheckRegion = "OC"
	PoolListResponseCheckRegionMe         PoolListResponseCheckRegion = "ME"
	PoolListResponseCheckRegionNaf        PoolListResponseCheckRegion = "NAF"
	PoolListResponseCheckRegionSaf        PoolListResponseCheckRegion = "SAF"
	PoolListResponseCheckRegionSas        PoolListResponseCheckRegion = "SAS"
	PoolListResponseCheckRegionSeas       PoolListResponseCheckRegion = "SEAS"
	PoolListResponseCheckRegionNeas       PoolListResponseCheckRegion = "NEAS"
	PoolListResponseCheckRegionAllRegions PoolListResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type PoolListResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy PoolListResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy PoolListResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          poolListResponseLoadSheddingJSON          `json:"-"`
}

// poolListResponseLoadSheddingJSON contains the JSON metadata for the struct
// [PoolListResponseLoadShedding]
type poolListResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PoolListResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseLoadSheddingJSON) RawJSON() string {
	return r.raw
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolListResponseLoadSheddingDefaultPolicy string

const (
	PoolListResponseLoadSheddingDefaultPolicyRandom PoolListResponseLoadSheddingDefaultPolicy = "random"
	PoolListResponseLoadSheddingDefaultPolicyHash   PoolListResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolListResponseLoadSheddingSessionPolicy string

const (
	PoolListResponseLoadSheddingSessionPolicyHash PoolListResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolListResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin PoolListResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool PoolListResponseNotificationFilterPool `json:"pool,nullable"`
	JSON poolListResponseNotificationFilterJSON `json:"-"`
}

// poolListResponseNotificationFilterJSON contains the JSON metadata for the struct
// [PoolListResponseNotificationFilter]
type poolListResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolListResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseNotificationFilterJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolListResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                         `json:"healthy,nullable"`
	JSON    poolListResponseNotificationFilterOriginJSON `json:"-"`
}

// poolListResponseNotificationFilterOriginJSON contains the JSON metadata for the
// struct [PoolListResponseNotificationFilterOrigin]
type poolListResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolListResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseNotificationFilterOriginJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolListResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                       `json:"healthy,nullable"`
	JSON    poolListResponseNotificationFilterPoolJSON `json:"-"`
}

// poolListResponseNotificationFilterPoolJSON contains the JSON metadata for the
// struct [PoolListResponseNotificationFilterPool]
type poolListResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolListResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseNotificationFilterPoolJSON) RawJSON() string {
	return r.raw
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolListResponseOriginSteering struct {
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
	Policy PoolListResponseOriginSteeringPolicy `json:"policy"`
	JSON   poolListResponseOriginSteeringJSON   `json:"-"`
}

// poolListResponseOriginSteeringJSON contains the JSON metadata for the struct
// [PoolListResponseOriginSteering]
type poolListResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolListResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseOriginSteeringJSON) RawJSON() string {
	return r.raw
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
type PoolListResponseOriginSteeringPolicy string

const (
	PoolListResponseOriginSteeringPolicyRandom                   PoolListResponseOriginSteeringPolicy = "random"
	PoolListResponseOriginSteeringPolicyHash                     PoolListResponseOriginSteeringPolicy = "hash"
	PoolListResponseOriginSteeringPolicyLeastOutstandingRequests PoolListResponseOriginSteeringPolicy = "least_outstanding_requests"
	PoolListResponseOriginSteeringPolicyLeastConnections         PoolListResponseOriginSteeringPolicy = "least_connections"
)

type PoolListResponseOrigin struct {
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
	Header PoolListResponseOriginsHeader `json:"header"`
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
	Weight float64                    `json:"weight"`
	JSON   poolListResponseOriginJSON `json:"-"`
}

// poolListResponseOriginJSON contains the JSON metadata for the struct
// [PoolListResponseOrigin]
type poolListResponseOriginJSON struct {
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

func (r *PoolListResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseOriginJSON) RawJSON() string {
	return r.raw
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolListResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                          `json:"Host"`
	JSON poolListResponseOriginsHeaderJSON `json:"-"`
}

// poolListResponseOriginsHeaderJSON contains the JSON metadata for the struct
// [PoolListResponseOriginsHeader]
type poolListResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolListResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseOriginsHeaderJSON) RawJSON() string {
	return r.raw
}

type PoolDeleteResponse struct {
	ID   string                 `json:"id"`
	JSON poolDeleteResponseJSON `json:"-"`
}

// poolDeleteResponseJSON contains the JSON metadata for the struct
// [PoolDeleteResponse]
type poolDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PoolEditResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []PoolEditResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                     `json:"created_on" format:"date-time"`
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
	LoadShedding PoolEditResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter PoolEditResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering PoolEditResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []PoolEditResponseOrigin `json:"origins"`
	JSON    poolEditResponseJSON     `json:"-"`
}

// poolEditResponseJSON contains the JSON metadata for the struct
// [PoolEditResponse]
type poolEditResponseJSON struct {
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

func (r *PoolEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolEditResponseCheckRegion string

const (
	PoolEditResponseCheckRegionWnam       PoolEditResponseCheckRegion = "WNAM"
	PoolEditResponseCheckRegionEnam       PoolEditResponseCheckRegion = "ENAM"
	PoolEditResponseCheckRegionWeu        PoolEditResponseCheckRegion = "WEU"
	PoolEditResponseCheckRegionEeu        PoolEditResponseCheckRegion = "EEU"
	PoolEditResponseCheckRegionNsam       PoolEditResponseCheckRegion = "NSAM"
	PoolEditResponseCheckRegionSsam       PoolEditResponseCheckRegion = "SSAM"
	PoolEditResponseCheckRegionOc         PoolEditResponseCheckRegion = "OC"
	PoolEditResponseCheckRegionMe         PoolEditResponseCheckRegion = "ME"
	PoolEditResponseCheckRegionNaf        PoolEditResponseCheckRegion = "NAF"
	PoolEditResponseCheckRegionSaf        PoolEditResponseCheckRegion = "SAF"
	PoolEditResponseCheckRegionSas        PoolEditResponseCheckRegion = "SAS"
	PoolEditResponseCheckRegionSeas       PoolEditResponseCheckRegion = "SEAS"
	PoolEditResponseCheckRegionNeas       PoolEditResponseCheckRegion = "NEAS"
	PoolEditResponseCheckRegionAllRegions PoolEditResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type PoolEditResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy PoolEditResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy PoolEditResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          poolEditResponseLoadSheddingJSON          `json:"-"`
}

// poolEditResponseLoadSheddingJSON contains the JSON metadata for the struct
// [PoolEditResponseLoadShedding]
type poolEditResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PoolEditResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseLoadSheddingJSON) RawJSON() string {
	return r.raw
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolEditResponseLoadSheddingDefaultPolicy string

const (
	PoolEditResponseLoadSheddingDefaultPolicyRandom PoolEditResponseLoadSheddingDefaultPolicy = "random"
	PoolEditResponseLoadSheddingDefaultPolicyHash   PoolEditResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolEditResponseLoadSheddingSessionPolicy string

const (
	PoolEditResponseLoadSheddingSessionPolicyHash PoolEditResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolEditResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin PoolEditResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool PoolEditResponseNotificationFilterPool `json:"pool,nullable"`
	JSON poolEditResponseNotificationFilterJSON `json:"-"`
}

// poolEditResponseNotificationFilterJSON contains the JSON metadata for the struct
// [PoolEditResponseNotificationFilter]
type poolEditResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolEditResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseNotificationFilterJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolEditResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                         `json:"healthy,nullable"`
	JSON    poolEditResponseNotificationFilterOriginJSON `json:"-"`
}

// poolEditResponseNotificationFilterOriginJSON contains the JSON metadata for the
// struct [PoolEditResponseNotificationFilterOrigin]
type poolEditResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolEditResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseNotificationFilterOriginJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolEditResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                       `json:"healthy,nullable"`
	JSON    poolEditResponseNotificationFilterPoolJSON `json:"-"`
}

// poolEditResponseNotificationFilterPoolJSON contains the JSON metadata for the
// struct [PoolEditResponseNotificationFilterPool]
type poolEditResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolEditResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseNotificationFilterPoolJSON) RawJSON() string {
	return r.raw
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolEditResponseOriginSteering struct {
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
	Policy PoolEditResponseOriginSteeringPolicy `json:"policy"`
	JSON   poolEditResponseOriginSteeringJSON   `json:"-"`
}

// poolEditResponseOriginSteeringJSON contains the JSON metadata for the struct
// [PoolEditResponseOriginSteering]
type poolEditResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolEditResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseOriginSteeringJSON) RawJSON() string {
	return r.raw
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
type PoolEditResponseOriginSteeringPolicy string

const (
	PoolEditResponseOriginSteeringPolicyRandom                   PoolEditResponseOriginSteeringPolicy = "random"
	PoolEditResponseOriginSteeringPolicyHash                     PoolEditResponseOriginSteeringPolicy = "hash"
	PoolEditResponseOriginSteeringPolicyLeastOutstandingRequests PoolEditResponseOriginSteeringPolicy = "least_outstanding_requests"
	PoolEditResponseOriginSteeringPolicyLeastConnections         PoolEditResponseOriginSteeringPolicy = "least_connections"
)

type PoolEditResponseOrigin struct {
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
	Header PoolEditResponseOriginsHeader `json:"header"`
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
	Weight float64                    `json:"weight"`
	JSON   poolEditResponseOriginJSON `json:"-"`
}

// poolEditResponseOriginJSON contains the JSON metadata for the struct
// [PoolEditResponseOrigin]
type poolEditResponseOriginJSON struct {
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

func (r *PoolEditResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseOriginJSON) RawJSON() string {
	return r.raw
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolEditResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                          `json:"Host"`
	JSON poolEditResponseOriginsHeaderJSON `json:"-"`
}

// poolEditResponseOriginsHeaderJSON contains the JSON metadata for the struct
// [PoolEditResponseOriginsHeader]
type poolEditResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolEditResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseOriginsHeaderJSON) RawJSON() string {
	return r.raw
}

type PoolGetResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []PoolGetResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                    `json:"created_on" format:"date-time"`
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
	LoadShedding PoolGetResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter PoolGetResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering PoolGetResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []PoolGetResponseOrigin `json:"origins"`
	JSON    poolGetResponseJSON     `json:"-"`
}

// poolGetResponseJSON contains the JSON metadata for the struct [PoolGetResponse]
type poolGetResponseJSON struct {
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

func (r *PoolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolGetResponseCheckRegion string

const (
	PoolGetResponseCheckRegionWnam       PoolGetResponseCheckRegion = "WNAM"
	PoolGetResponseCheckRegionEnam       PoolGetResponseCheckRegion = "ENAM"
	PoolGetResponseCheckRegionWeu        PoolGetResponseCheckRegion = "WEU"
	PoolGetResponseCheckRegionEeu        PoolGetResponseCheckRegion = "EEU"
	PoolGetResponseCheckRegionNsam       PoolGetResponseCheckRegion = "NSAM"
	PoolGetResponseCheckRegionSsam       PoolGetResponseCheckRegion = "SSAM"
	PoolGetResponseCheckRegionOc         PoolGetResponseCheckRegion = "OC"
	PoolGetResponseCheckRegionMe         PoolGetResponseCheckRegion = "ME"
	PoolGetResponseCheckRegionNaf        PoolGetResponseCheckRegion = "NAF"
	PoolGetResponseCheckRegionSaf        PoolGetResponseCheckRegion = "SAF"
	PoolGetResponseCheckRegionSas        PoolGetResponseCheckRegion = "SAS"
	PoolGetResponseCheckRegionSeas       PoolGetResponseCheckRegion = "SEAS"
	PoolGetResponseCheckRegionNeas       PoolGetResponseCheckRegion = "NEAS"
	PoolGetResponseCheckRegionAllRegions PoolGetResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type PoolGetResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy PoolGetResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy PoolGetResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          poolGetResponseLoadSheddingJSON          `json:"-"`
}

// poolGetResponseLoadSheddingJSON contains the JSON metadata for the struct
// [PoolGetResponseLoadShedding]
type poolGetResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PoolGetResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseLoadSheddingJSON) RawJSON() string {
	return r.raw
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolGetResponseLoadSheddingDefaultPolicy string

const (
	PoolGetResponseLoadSheddingDefaultPolicyRandom PoolGetResponseLoadSheddingDefaultPolicy = "random"
	PoolGetResponseLoadSheddingDefaultPolicyHash   PoolGetResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolGetResponseLoadSheddingSessionPolicy string

const (
	PoolGetResponseLoadSheddingSessionPolicyHash PoolGetResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolGetResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin PoolGetResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool PoolGetResponseNotificationFilterPool `json:"pool,nullable"`
	JSON poolGetResponseNotificationFilterJSON `json:"-"`
}

// poolGetResponseNotificationFilterJSON contains the JSON metadata for the struct
// [PoolGetResponseNotificationFilter]
type poolGetResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolGetResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseNotificationFilterJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolGetResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                        `json:"healthy,nullable"`
	JSON    poolGetResponseNotificationFilterOriginJSON `json:"-"`
}

// poolGetResponseNotificationFilterOriginJSON contains the JSON metadata for the
// struct [PoolGetResponseNotificationFilterOrigin]
type poolGetResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolGetResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseNotificationFilterOriginJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolGetResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                      `json:"healthy,nullable"`
	JSON    poolGetResponseNotificationFilterPoolJSON `json:"-"`
}

// poolGetResponseNotificationFilterPoolJSON contains the JSON metadata for the
// struct [PoolGetResponseNotificationFilterPool]
type poolGetResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolGetResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseNotificationFilterPoolJSON) RawJSON() string {
	return r.raw
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolGetResponseOriginSteering struct {
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
	Policy PoolGetResponseOriginSteeringPolicy `json:"policy"`
	JSON   poolGetResponseOriginSteeringJSON   `json:"-"`
}

// poolGetResponseOriginSteeringJSON contains the JSON metadata for the struct
// [PoolGetResponseOriginSteering]
type poolGetResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolGetResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseOriginSteeringJSON) RawJSON() string {
	return r.raw
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
type PoolGetResponseOriginSteeringPolicy string

const (
	PoolGetResponseOriginSteeringPolicyRandom                   PoolGetResponseOriginSteeringPolicy = "random"
	PoolGetResponseOriginSteeringPolicyHash                     PoolGetResponseOriginSteeringPolicy = "hash"
	PoolGetResponseOriginSteeringPolicyLeastOutstandingRequests PoolGetResponseOriginSteeringPolicy = "least_outstanding_requests"
	PoolGetResponseOriginSteeringPolicyLeastConnections         PoolGetResponseOriginSteeringPolicy = "least_connections"
)

type PoolGetResponseOrigin struct {
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
	Header PoolGetResponseOriginsHeader `json:"header"`
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
	Weight float64                   `json:"weight"`
	JSON   poolGetResponseOriginJSON `json:"-"`
}

// poolGetResponseOriginJSON contains the JSON metadata for the struct
// [PoolGetResponseOrigin]
type poolGetResponseOriginJSON struct {
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

func (r *PoolGetResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseOriginJSON) RawJSON() string {
	return r.raw
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolGetResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                         `json:"Host"`
	JSON poolGetResponseOriginsHeaderJSON `json:"-"`
}

// poolGetResponseOriginsHeaderJSON contains the JSON metadata for the struct
// [PoolGetResponseOriginsHeader]
type poolGetResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolGetResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseOriginsHeaderJSON) RawJSON() string {
	return r.raw
}

type PoolNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]PoolNewParamsOrigin] `json:"origins,required"`
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
	LoadShedding param.Field[PoolNewParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[PoolNewParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[PoolNewParamsOriginSteering] `json:"origin_steering"`
}

func (r PoolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PoolNewParamsOrigin struct {
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
	Header param.Field[PoolNewParamsOriginsHeader] `json:"header"`
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

func (r PoolNewParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolNewParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r PoolNewParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures load shedding policies and percentages for the pool.
type PoolNewParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[PoolNewParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[PoolNewParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r PoolNewParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolNewParamsLoadSheddingDefaultPolicy string

const (
	PoolNewParamsLoadSheddingDefaultPolicyRandom PoolNewParamsLoadSheddingDefaultPolicy = "random"
	PoolNewParamsLoadSheddingDefaultPolicyHash   PoolNewParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolNewParamsLoadSheddingSessionPolicy string

const (
	PoolNewParamsLoadSheddingSessionPolicyHash PoolNewParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolNewParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[PoolNewParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[PoolNewParamsNotificationFilterPool] `json:"pool"`
}

func (r PoolNewParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolNewParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r PoolNewParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolNewParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r PoolNewParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolNewParamsOriginSteering struct {
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
	Policy param.Field[PoolNewParamsOriginSteeringPolicy] `json:"policy"`
}

func (r PoolNewParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type PoolNewParamsOriginSteeringPolicy string

const (
	PoolNewParamsOriginSteeringPolicyRandom                   PoolNewParamsOriginSteeringPolicy = "random"
	PoolNewParamsOriginSteeringPolicyHash                     PoolNewParamsOriginSteeringPolicy = "hash"
	PoolNewParamsOriginSteeringPolicyLeastOutstandingRequests PoolNewParamsOriginSteeringPolicy = "least_outstanding_requests"
	PoolNewParamsOriginSteeringPolicyLeastConnections         PoolNewParamsOriginSteeringPolicy = "least_connections"
)

type PoolNewResponseEnvelope struct {
	Errors   []PoolNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PoolNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PoolNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PoolNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolNewResponseEnvelopeJSON    `json:"-"`
}

// poolNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolNewResponseEnvelope]
type poolNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PoolNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    poolNewResponseEnvelopeErrorsJSON `json:"-"`
}

// poolNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PoolNewResponseEnvelopeErrors]
type poolNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PoolNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    poolNewResponseEnvelopeMessagesJSON `json:"-"`
}

// poolNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PoolNewResponseEnvelopeMessages]
type poolNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolNewResponseEnvelopeSuccess bool

const (
	PoolNewResponseEnvelopeSuccessTrue PoolNewResponseEnvelopeSuccess = true
)

type PoolUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]PoolUpdateParamsOrigin] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]PoolUpdateParamsCheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[PoolUpdateParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[PoolUpdateParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[PoolUpdateParamsOriginSteering] `json:"origin_steering"`
}

func (r PoolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PoolUpdateParamsOrigin struct {
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
	Header param.Field[PoolUpdateParamsOriginsHeader] `json:"header"`
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

func (r PoolUpdateParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolUpdateParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r PoolUpdateParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolUpdateParamsCheckRegion string

const (
	PoolUpdateParamsCheckRegionWnam       PoolUpdateParamsCheckRegion = "WNAM"
	PoolUpdateParamsCheckRegionEnam       PoolUpdateParamsCheckRegion = "ENAM"
	PoolUpdateParamsCheckRegionWeu        PoolUpdateParamsCheckRegion = "WEU"
	PoolUpdateParamsCheckRegionEeu        PoolUpdateParamsCheckRegion = "EEU"
	PoolUpdateParamsCheckRegionNsam       PoolUpdateParamsCheckRegion = "NSAM"
	PoolUpdateParamsCheckRegionSsam       PoolUpdateParamsCheckRegion = "SSAM"
	PoolUpdateParamsCheckRegionOc         PoolUpdateParamsCheckRegion = "OC"
	PoolUpdateParamsCheckRegionMe         PoolUpdateParamsCheckRegion = "ME"
	PoolUpdateParamsCheckRegionNaf        PoolUpdateParamsCheckRegion = "NAF"
	PoolUpdateParamsCheckRegionSaf        PoolUpdateParamsCheckRegion = "SAF"
	PoolUpdateParamsCheckRegionSas        PoolUpdateParamsCheckRegion = "SAS"
	PoolUpdateParamsCheckRegionSeas       PoolUpdateParamsCheckRegion = "SEAS"
	PoolUpdateParamsCheckRegionNeas       PoolUpdateParamsCheckRegion = "NEAS"
	PoolUpdateParamsCheckRegionAllRegions PoolUpdateParamsCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type PoolUpdateParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[PoolUpdateParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[PoolUpdateParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r PoolUpdateParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolUpdateParamsLoadSheddingDefaultPolicy string

const (
	PoolUpdateParamsLoadSheddingDefaultPolicyRandom PoolUpdateParamsLoadSheddingDefaultPolicy = "random"
	PoolUpdateParamsLoadSheddingDefaultPolicyHash   PoolUpdateParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolUpdateParamsLoadSheddingSessionPolicy string

const (
	PoolUpdateParamsLoadSheddingSessionPolicyHash PoolUpdateParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolUpdateParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[PoolUpdateParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[PoolUpdateParamsNotificationFilterPool] `json:"pool"`
}

func (r PoolUpdateParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolUpdateParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r PoolUpdateParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolUpdateParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r PoolUpdateParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolUpdateParamsOriginSteering struct {
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
	Policy param.Field[PoolUpdateParamsOriginSteeringPolicy] `json:"policy"`
}

func (r PoolUpdateParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type PoolUpdateParamsOriginSteeringPolicy string

const (
	PoolUpdateParamsOriginSteeringPolicyRandom                   PoolUpdateParamsOriginSteeringPolicy = "random"
	PoolUpdateParamsOriginSteeringPolicyHash                     PoolUpdateParamsOriginSteeringPolicy = "hash"
	PoolUpdateParamsOriginSteeringPolicyLeastOutstandingRequests PoolUpdateParamsOriginSteeringPolicy = "least_outstanding_requests"
	PoolUpdateParamsOriginSteeringPolicyLeastConnections         PoolUpdateParamsOriginSteeringPolicy = "least_connections"
)

type PoolUpdateResponseEnvelope struct {
	Errors   []PoolUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PoolUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   PoolUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PoolUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolUpdateResponseEnvelopeJSON    `json:"-"`
}

// poolUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolUpdateResponseEnvelope]
type poolUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PoolUpdateResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    poolUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// poolUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PoolUpdateResponseEnvelopeErrors]
type poolUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PoolUpdateResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    poolUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// poolUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PoolUpdateResponseEnvelopeMessages]
type poolUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolUpdateResponseEnvelopeSuccess bool

const (
	PoolUpdateResponseEnvelopeSuccessTrue PoolUpdateResponseEnvelopeSuccess = true
)

type PoolListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `query:"monitor"`
}

// URLQuery serializes [PoolListParams]'s query parameters as `url.Values`.
func (r PoolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PoolListResponseEnvelope struct {
	Errors   []PoolListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PoolListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PoolListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PoolListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PoolListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       poolListResponseEnvelopeJSON       `json:"-"`
}

// poolListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolListResponseEnvelope]
type poolListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PoolListResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    poolListResponseEnvelopeErrorsJSON `json:"-"`
}

// poolListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PoolListResponseEnvelopeErrors]
type poolListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PoolListResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    poolListResponseEnvelopeMessagesJSON `json:"-"`
}

// poolListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PoolListResponseEnvelopeMessages]
type poolListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolListResponseEnvelopeSuccess bool

const (
	PoolListResponseEnvelopeSuccessTrue PoolListResponseEnvelopeSuccess = true
)

type PoolListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       poolListResponseEnvelopeResultInfoJSON `json:"-"`
}

// poolListResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [PoolListResponseEnvelopeResultInfo]
type poolListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type PoolDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PoolDeleteResponseEnvelope struct {
	Errors   []PoolDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PoolDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   PoolDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PoolDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolDeleteResponseEnvelopeJSON    `json:"-"`
}

// poolDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolDeleteResponseEnvelope]
type poolDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PoolDeleteResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    poolDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// poolDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PoolDeleteResponseEnvelopeErrors]
type poolDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PoolDeleteResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    poolDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// poolDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PoolDeleteResponseEnvelopeMessages]
type poolDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolDeleteResponseEnvelopeSuccess bool

const (
	PoolDeleteResponseEnvelopeSuccessTrue PoolDeleteResponseEnvelopeSuccess = true
)

type PoolEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]PoolEditParamsCheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[PoolEditParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[PoolEditParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[PoolEditParamsOriginSteering] `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]PoolEditParamsOrigin] `json:"origins"`
}

func (r PoolEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolEditParamsCheckRegion string

const (
	PoolEditParamsCheckRegionWnam       PoolEditParamsCheckRegion = "WNAM"
	PoolEditParamsCheckRegionEnam       PoolEditParamsCheckRegion = "ENAM"
	PoolEditParamsCheckRegionWeu        PoolEditParamsCheckRegion = "WEU"
	PoolEditParamsCheckRegionEeu        PoolEditParamsCheckRegion = "EEU"
	PoolEditParamsCheckRegionNsam       PoolEditParamsCheckRegion = "NSAM"
	PoolEditParamsCheckRegionSsam       PoolEditParamsCheckRegion = "SSAM"
	PoolEditParamsCheckRegionOc         PoolEditParamsCheckRegion = "OC"
	PoolEditParamsCheckRegionMe         PoolEditParamsCheckRegion = "ME"
	PoolEditParamsCheckRegionNaf        PoolEditParamsCheckRegion = "NAF"
	PoolEditParamsCheckRegionSaf        PoolEditParamsCheckRegion = "SAF"
	PoolEditParamsCheckRegionSas        PoolEditParamsCheckRegion = "SAS"
	PoolEditParamsCheckRegionSeas       PoolEditParamsCheckRegion = "SEAS"
	PoolEditParamsCheckRegionNeas       PoolEditParamsCheckRegion = "NEAS"
	PoolEditParamsCheckRegionAllRegions PoolEditParamsCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type PoolEditParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[PoolEditParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[PoolEditParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r PoolEditParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolEditParamsLoadSheddingDefaultPolicy string

const (
	PoolEditParamsLoadSheddingDefaultPolicyRandom PoolEditParamsLoadSheddingDefaultPolicy = "random"
	PoolEditParamsLoadSheddingDefaultPolicyHash   PoolEditParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolEditParamsLoadSheddingSessionPolicy string

const (
	PoolEditParamsLoadSheddingSessionPolicyHash PoolEditParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolEditParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[PoolEditParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[PoolEditParamsNotificationFilterPool] `json:"pool"`
}

func (r PoolEditParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolEditParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r PoolEditParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolEditParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r PoolEditParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolEditParamsOriginSteering struct {
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
	Policy param.Field[PoolEditParamsOriginSteeringPolicy] `json:"policy"`
}

func (r PoolEditParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type PoolEditParamsOriginSteeringPolicy string

const (
	PoolEditParamsOriginSteeringPolicyRandom                   PoolEditParamsOriginSteeringPolicy = "random"
	PoolEditParamsOriginSteeringPolicyHash                     PoolEditParamsOriginSteeringPolicy = "hash"
	PoolEditParamsOriginSteeringPolicyLeastOutstandingRequests PoolEditParamsOriginSteeringPolicy = "least_outstanding_requests"
	PoolEditParamsOriginSteeringPolicyLeastConnections         PoolEditParamsOriginSteeringPolicy = "least_connections"
)

type PoolEditParamsOrigin struct {
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
	Header param.Field[PoolEditParamsOriginsHeader] `json:"header"`
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

func (r PoolEditParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolEditParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r PoolEditParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PoolEditResponseEnvelope struct {
	Errors   []PoolEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PoolEditResponseEnvelopeMessages `json:"messages,required"`
	Result   PoolEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PoolEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolEditResponseEnvelopeJSON    `json:"-"`
}

// poolEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolEditResponseEnvelope]
type poolEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PoolEditResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    poolEditResponseEnvelopeErrorsJSON `json:"-"`
}

// poolEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PoolEditResponseEnvelopeErrors]
type poolEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PoolEditResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    poolEditResponseEnvelopeMessagesJSON `json:"-"`
}

// poolEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PoolEditResponseEnvelopeMessages]
type poolEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolEditResponseEnvelopeSuccess bool

const (
	PoolEditResponseEnvelopeSuccessTrue PoolEditResponseEnvelopeSuccess = true
)

type PoolGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PoolGetResponseEnvelope struct {
	Errors   []PoolGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PoolGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PoolGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success PoolGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolGetResponseEnvelopeJSON    `json:"-"`
}

// poolGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolGetResponseEnvelope]
type poolGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PoolGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    poolGetResponseEnvelopeErrorsJSON `json:"-"`
}

// poolGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PoolGetResponseEnvelopeErrors]
type poolGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PoolGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    poolGetResponseEnvelopeMessagesJSON `json:"-"`
}

// poolGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PoolGetResponseEnvelopeMessages]
type poolGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolGetResponseEnvelopeSuccess bool

const (
	PoolGetResponseEnvelopeSuccessTrue PoolGetResponseEnvelopeSuccess = true
)
