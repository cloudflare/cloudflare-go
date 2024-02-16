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
	References *LoadBalancerPoolReferenceService
}

// NewLoadBalancerPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerPoolService(opts ...option.RequestOption) (r *LoadBalancerPoolService) {
	r = &LoadBalancerPoolService{}
	r.Options = opts
	r.Health = NewLoadBalancerPoolHealthService(opts...)
	r.References = NewLoadBalancerPoolReferenceService(opts...)
	return
}

// Create a new pool.
func (r *LoadBalancerPoolService) New(ctx context.Context, accountID string, body LoadBalancerPoolNewParams, opts ...option.RequestOption) (res *LoadBalancerPoolNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing pool, overwriting the supplied properties.
func (r *LoadBalancerPoolService) Update(ctx context.Context, accountID string, poolID string, body LoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *LoadBalancerPoolUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured pools.
func (r *LoadBalancerPoolService) List(ctx context.Context, accountID string, query LoadBalancerPoolListParams, opts ...option.RequestOption) (res *[]LoadBalancerPoolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured pool.
func (r *LoadBalancerPoolService) Delete(ctx context.Context, accountID string, poolID string, opts ...option.RequestOption) (res *LoadBalancerPoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured pool.
func (r *LoadBalancerPoolService) Get(ctx context.Context, accountID string, poolID string, opts ...option.RequestOption) (res *LoadBalancerPoolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerPoolNewResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolNewResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                `json:"created_on" format:"date-time"`
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
	LoadShedding LoadBalancerPoolNewResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter LoadBalancerPoolNewResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolNewResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolNewResponseOrigin `json:"origins"`
	JSON    loadBalancerPoolNewResponseJSON     `json:"-"`
}

// loadBalancerPoolNewResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolNewResponse]
type loadBalancerPoolNewResponseJSON struct {
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

func (r *LoadBalancerPoolNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolNewResponseCheckRegion string

const (
	LoadBalancerPoolNewResponseCheckRegionWnam       LoadBalancerPoolNewResponseCheckRegion = "WNAM"
	LoadBalancerPoolNewResponseCheckRegionEnam       LoadBalancerPoolNewResponseCheckRegion = "ENAM"
	LoadBalancerPoolNewResponseCheckRegionWeu        LoadBalancerPoolNewResponseCheckRegion = "WEU"
	LoadBalancerPoolNewResponseCheckRegionEeu        LoadBalancerPoolNewResponseCheckRegion = "EEU"
	LoadBalancerPoolNewResponseCheckRegionNsam       LoadBalancerPoolNewResponseCheckRegion = "NSAM"
	LoadBalancerPoolNewResponseCheckRegionSsam       LoadBalancerPoolNewResponseCheckRegion = "SSAM"
	LoadBalancerPoolNewResponseCheckRegionOc         LoadBalancerPoolNewResponseCheckRegion = "OC"
	LoadBalancerPoolNewResponseCheckRegionMe         LoadBalancerPoolNewResponseCheckRegion = "ME"
	LoadBalancerPoolNewResponseCheckRegionNaf        LoadBalancerPoolNewResponseCheckRegion = "NAF"
	LoadBalancerPoolNewResponseCheckRegionSaf        LoadBalancerPoolNewResponseCheckRegion = "SAF"
	LoadBalancerPoolNewResponseCheckRegionSas        LoadBalancerPoolNewResponseCheckRegion = "SAS"
	LoadBalancerPoolNewResponseCheckRegionSeas       LoadBalancerPoolNewResponseCheckRegion = "SEAS"
	LoadBalancerPoolNewResponseCheckRegionNeas       LoadBalancerPoolNewResponseCheckRegion = "NEAS"
	LoadBalancerPoolNewResponseCheckRegionAllRegions LoadBalancerPoolNewResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolNewResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolNewResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolNewResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolNewResponseLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolNewResponseLoadSheddingJSON contains the JSON metadata for the
// struct [LoadBalancerPoolNewResponseLoadShedding]
type loadBalancerPoolNewResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolNewResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolNewResponseLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolNewResponseLoadSheddingDefaultPolicyRandom LoadBalancerPoolNewResponseLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolNewResponseLoadSheddingDefaultPolicyHash   LoadBalancerPoolNewResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolNewResponseLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolNewResponseLoadSheddingSessionPolicyHash LoadBalancerPoolNewResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolNewResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolNewResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolNewResponseNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolNewResponseNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolNewResponseNotificationFilterJSON contains the JSON metadata for
// the struct [LoadBalancerPoolNewResponseNotificationFilter]
type loadBalancerPoolNewResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolNewResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolNewResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                    `json:"healthy,nullable"`
	JSON    loadBalancerPoolNewResponseNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolNewResponseNotificationFilterOriginJSON contains the JSON
// metadata for the struct [LoadBalancerPoolNewResponseNotificationFilterOrigin]
type loadBalancerPoolNewResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolNewResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolNewResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                  `json:"healthy,nullable"`
	JSON    loadBalancerPoolNewResponseNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolNewResponseNotificationFilterPoolJSON contains the JSON metadata
// for the struct [LoadBalancerPoolNewResponseNotificationFilterPool]
type loadBalancerPoolNewResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolNewResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolNewResponseOriginSteering struct {
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
	Policy LoadBalancerPoolNewResponseOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolNewResponseOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolNewResponseOriginSteeringJSON contains the JSON metadata for the
// struct [LoadBalancerPoolNewResponseOriginSteering]
type loadBalancerPoolNewResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolNewResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type LoadBalancerPoolNewResponseOriginSteeringPolicy string

const (
	LoadBalancerPoolNewResponseOriginSteeringPolicyRandom                   LoadBalancerPoolNewResponseOriginSteeringPolicy = "random"
	LoadBalancerPoolNewResponseOriginSteeringPolicyHash                     LoadBalancerPoolNewResponseOriginSteeringPolicy = "hash"
	LoadBalancerPoolNewResponseOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolNewResponseOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolNewResponseOriginSteeringPolicyLeastConnections         LoadBalancerPoolNewResponseOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolNewResponseOrigin struct {
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
	Header LoadBalancerPoolNewResponseOriginsHeader `json:"header"`
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
	Weight float64                               `json:"weight"`
	JSON   loadBalancerPoolNewResponseOriginJSON `json:"-"`
}

// loadBalancerPoolNewResponseOriginJSON contains the JSON metadata for the struct
// [LoadBalancerPoolNewResponseOrigin]
type loadBalancerPoolNewResponseOriginJSON struct {
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

func (r *LoadBalancerPoolNewResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolNewResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                     `json:"Host"`
	JSON loadBalancerPoolNewResponseOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolNewResponseOriginsHeaderJSON contains the JSON metadata for the
// struct [LoadBalancerPoolNewResponseOriginsHeader]
type loadBalancerPoolNewResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolNewResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolUpdateResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolUpdateResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                   `json:"created_on" format:"date-time"`
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
	LoadShedding LoadBalancerPoolUpdateResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter LoadBalancerPoolUpdateResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolUpdateResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolUpdateResponseOrigin `json:"origins"`
	JSON    loadBalancerPoolUpdateResponseJSON     `json:"-"`
}

// loadBalancerPoolUpdateResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolUpdateResponse]
type loadBalancerPoolUpdateResponseJSON struct {
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

func (r *LoadBalancerPoolUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolUpdateResponseCheckRegion string

const (
	LoadBalancerPoolUpdateResponseCheckRegionWnam       LoadBalancerPoolUpdateResponseCheckRegion = "WNAM"
	LoadBalancerPoolUpdateResponseCheckRegionEnam       LoadBalancerPoolUpdateResponseCheckRegion = "ENAM"
	LoadBalancerPoolUpdateResponseCheckRegionWeu        LoadBalancerPoolUpdateResponseCheckRegion = "WEU"
	LoadBalancerPoolUpdateResponseCheckRegionEeu        LoadBalancerPoolUpdateResponseCheckRegion = "EEU"
	LoadBalancerPoolUpdateResponseCheckRegionNsam       LoadBalancerPoolUpdateResponseCheckRegion = "NSAM"
	LoadBalancerPoolUpdateResponseCheckRegionSsam       LoadBalancerPoolUpdateResponseCheckRegion = "SSAM"
	LoadBalancerPoolUpdateResponseCheckRegionOc         LoadBalancerPoolUpdateResponseCheckRegion = "OC"
	LoadBalancerPoolUpdateResponseCheckRegionMe         LoadBalancerPoolUpdateResponseCheckRegion = "ME"
	LoadBalancerPoolUpdateResponseCheckRegionNaf        LoadBalancerPoolUpdateResponseCheckRegion = "NAF"
	LoadBalancerPoolUpdateResponseCheckRegionSaf        LoadBalancerPoolUpdateResponseCheckRegion = "SAF"
	LoadBalancerPoolUpdateResponseCheckRegionSas        LoadBalancerPoolUpdateResponseCheckRegion = "SAS"
	LoadBalancerPoolUpdateResponseCheckRegionSeas       LoadBalancerPoolUpdateResponseCheckRegion = "SEAS"
	LoadBalancerPoolUpdateResponseCheckRegionNeas       LoadBalancerPoolUpdateResponseCheckRegion = "NEAS"
	LoadBalancerPoolUpdateResponseCheckRegionAllRegions LoadBalancerPoolUpdateResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolUpdateResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolUpdateResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolUpdateResponseLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolUpdateResponseLoadSheddingJSON contains the JSON metadata for
// the struct [LoadBalancerPoolUpdateResponseLoadShedding]
type loadBalancerPoolUpdateResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicyRandom LoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicyHash   LoadBalancerPoolUpdateResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolUpdateResponseLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolUpdateResponseLoadSheddingSessionPolicyHash LoadBalancerPoolUpdateResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolUpdateResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolUpdateResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolUpdateResponseNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolUpdateResponseNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseNotificationFilterJSON contains the JSON metadata
// for the struct [LoadBalancerPoolUpdateResponseNotificationFilter]
type loadBalancerPoolUpdateResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolUpdateResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                       `json:"healthy,nullable"`
	JSON    loadBalancerPoolUpdateResponseNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseNotificationFilterOriginJSON contains the JSON
// metadata for the struct [LoadBalancerPoolUpdateResponseNotificationFilterOrigin]
type loadBalancerPoolUpdateResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolUpdateResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                     `json:"healthy,nullable"`
	JSON    loadBalancerPoolUpdateResponseNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseNotificationFilterPoolJSON contains the JSON
// metadata for the struct [LoadBalancerPoolUpdateResponseNotificationFilterPool]
type loadBalancerPoolUpdateResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolUpdateResponseOriginSteering struct {
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
	Policy LoadBalancerPoolUpdateResponseOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolUpdateResponseOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolUpdateResponseOriginSteeringJSON contains the JSON metadata for
// the struct [LoadBalancerPoolUpdateResponseOriginSteering]
type loadBalancerPoolUpdateResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type LoadBalancerPoolUpdateResponseOriginSteeringPolicy string

const (
	LoadBalancerPoolUpdateResponseOriginSteeringPolicyRandom                   LoadBalancerPoolUpdateResponseOriginSteeringPolicy = "random"
	LoadBalancerPoolUpdateResponseOriginSteeringPolicyHash                     LoadBalancerPoolUpdateResponseOriginSteeringPolicy = "hash"
	LoadBalancerPoolUpdateResponseOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolUpdateResponseOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolUpdateResponseOriginSteeringPolicyLeastConnections         LoadBalancerPoolUpdateResponseOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolUpdateResponseOrigin struct {
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
	Header LoadBalancerPoolUpdateResponseOriginsHeader `json:"header"`
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
	Weight float64                                  `json:"weight"`
	JSON   loadBalancerPoolUpdateResponseOriginJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseOriginJSON contains the JSON metadata for the
// struct [LoadBalancerPoolUpdateResponseOrigin]
type loadBalancerPoolUpdateResponseOriginJSON struct {
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

func (r *LoadBalancerPoolUpdateResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolUpdateResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                        `json:"Host"`
	JSON loadBalancerPoolUpdateResponseOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseOriginsHeaderJSON contains the JSON metadata for
// the struct [LoadBalancerPoolUpdateResponseOriginsHeader]
type loadBalancerPoolUpdateResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolListResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolListResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                 `json:"created_on" format:"date-time"`
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
	LoadShedding LoadBalancerPoolListResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter LoadBalancerPoolListResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolListResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolListResponseOrigin `json:"origins"`
	JSON    loadBalancerPoolListResponseJSON     `json:"-"`
}

// loadBalancerPoolListResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolListResponse]
type loadBalancerPoolListResponseJSON struct {
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

func (r *LoadBalancerPoolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolListResponseCheckRegion string

const (
	LoadBalancerPoolListResponseCheckRegionWnam       LoadBalancerPoolListResponseCheckRegion = "WNAM"
	LoadBalancerPoolListResponseCheckRegionEnam       LoadBalancerPoolListResponseCheckRegion = "ENAM"
	LoadBalancerPoolListResponseCheckRegionWeu        LoadBalancerPoolListResponseCheckRegion = "WEU"
	LoadBalancerPoolListResponseCheckRegionEeu        LoadBalancerPoolListResponseCheckRegion = "EEU"
	LoadBalancerPoolListResponseCheckRegionNsam       LoadBalancerPoolListResponseCheckRegion = "NSAM"
	LoadBalancerPoolListResponseCheckRegionSsam       LoadBalancerPoolListResponseCheckRegion = "SSAM"
	LoadBalancerPoolListResponseCheckRegionOc         LoadBalancerPoolListResponseCheckRegion = "OC"
	LoadBalancerPoolListResponseCheckRegionMe         LoadBalancerPoolListResponseCheckRegion = "ME"
	LoadBalancerPoolListResponseCheckRegionNaf        LoadBalancerPoolListResponseCheckRegion = "NAF"
	LoadBalancerPoolListResponseCheckRegionSaf        LoadBalancerPoolListResponseCheckRegion = "SAF"
	LoadBalancerPoolListResponseCheckRegionSas        LoadBalancerPoolListResponseCheckRegion = "SAS"
	LoadBalancerPoolListResponseCheckRegionSeas       LoadBalancerPoolListResponseCheckRegion = "SEAS"
	LoadBalancerPoolListResponseCheckRegionNeas       LoadBalancerPoolListResponseCheckRegion = "NEAS"
	LoadBalancerPoolListResponseCheckRegionAllRegions LoadBalancerPoolListResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolListResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolListResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolListResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolListResponseLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolListResponseLoadSheddingJSON contains the JSON metadata for the
// struct [LoadBalancerPoolListResponseLoadShedding]
type loadBalancerPoolListResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolListResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolListResponseLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolListResponseLoadSheddingDefaultPolicyRandom LoadBalancerPoolListResponseLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolListResponseLoadSheddingDefaultPolicyHash   LoadBalancerPoolListResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolListResponseLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolListResponseLoadSheddingSessionPolicyHash LoadBalancerPoolListResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolListResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolListResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolListResponseNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolListResponseNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolListResponseNotificationFilterJSON contains the JSON metadata
// for the struct [LoadBalancerPoolListResponseNotificationFilter]
type loadBalancerPoolListResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolListResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolListResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                     `json:"healthy,nullable"`
	JSON    loadBalancerPoolListResponseNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolListResponseNotificationFilterOriginJSON contains the JSON
// metadata for the struct [LoadBalancerPoolListResponseNotificationFilterOrigin]
type loadBalancerPoolListResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolListResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolListResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                   `json:"healthy,nullable"`
	JSON    loadBalancerPoolListResponseNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolListResponseNotificationFilterPoolJSON contains the JSON
// metadata for the struct [LoadBalancerPoolListResponseNotificationFilterPool]
type loadBalancerPoolListResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolListResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolListResponseOriginSteering struct {
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
	Policy LoadBalancerPoolListResponseOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolListResponseOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolListResponseOriginSteeringJSON contains the JSON metadata for
// the struct [LoadBalancerPoolListResponseOriginSteering]
type loadBalancerPoolListResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolListResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type LoadBalancerPoolListResponseOriginSteeringPolicy string

const (
	LoadBalancerPoolListResponseOriginSteeringPolicyRandom                   LoadBalancerPoolListResponseOriginSteeringPolicy = "random"
	LoadBalancerPoolListResponseOriginSteeringPolicyHash                     LoadBalancerPoolListResponseOriginSteeringPolicy = "hash"
	LoadBalancerPoolListResponseOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolListResponseOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolListResponseOriginSteeringPolicyLeastConnections         LoadBalancerPoolListResponseOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolListResponseOrigin struct {
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
	Header LoadBalancerPoolListResponseOriginsHeader `json:"header"`
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
	Weight float64                                `json:"weight"`
	JSON   loadBalancerPoolListResponseOriginJSON `json:"-"`
}

// loadBalancerPoolListResponseOriginJSON contains the JSON metadata for the struct
// [LoadBalancerPoolListResponseOrigin]
type loadBalancerPoolListResponseOriginJSON struct {
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

func (r *LoadBalancerPoolListResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolListResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                      `json:"Host"`
	JSON loadBalancerPoolListResponseOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolListResponseOriginsHeaderJSON contains the JSON metadata for the
// struct [LoadBalancerPoolListResponseOriginsHeader]
type loadBalancerPoolListResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolListResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolDeleteResponse struct {
	ID   string                             `json:"id"`
	JSON loadBalancerPoolDeleteResponseJSON `json:"-"`
}

// loadBalancerPoolDeleteResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolDeleteResponse]
type loadBalancerPoolDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolGetResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolGetResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                `json:"created_on" format:"date-time"`
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
	LoadShedding LoadBalancerPoolGetResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter LoadBalancerPoolGetResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolGetResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolGetResponseOrigin `json:"origins"`
	JSON    loadBalancerPoolGetResponseJSON     `json:"-"`
}

// loadBalancerPoolGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolGetResponse]
type loadBalancerPoolGetResponseJSON struct {
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

func (r *LoadBalancerPoolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolGetResponseCheckRegion string

const (
	LoadBalancerPoolGetResponseCheckRegionWnam       LoadBalancerPoolGetResponseCheckRegion = "WNAM"
	LoadBalancerPoolGetResponseCheckRegionEnam       LoadBalancerPoolGetResponseCheckRegion = "ENAM"
	LoadBalancerPoolGetResponseCheckRegionWeu        LoadBalancerPoolGetResponseCheckRegion = "WEU"
	LoadBalancerPoolGetResponseCheckRegionEeu        LoadBalancerPoolGetResponseCheckRegion = "EEU"
	LoadBalancerPoolGetResponseCheckRegionNsam       LoadBalancerPoolGetResponseCheckRegion = "NSAM"
	LoadBalancerPoolGetResponseCheckRegionSsam       LoadBalancerPoolGetResponseCheckRegion = "SSAM"
	LoadBalancerPoolGetResponseCheckRegionOc         LoadBalancerPoolGetResponseCheckRegion = "OC"
	LoadBalancerPoolGetResponseCheckRegionMe         LoadBalancerPoolGetResponseCheckRegion = "ME"
	LoadBalancerPoolGetResponseCheckRegionNaf        LoadBalancerPoolGetResponseCheckRegion = "NAF"
	LoadBalancerPoolGetResponseCheckRegionSaf        LoadBalancerPoolGetResponseCheckRegion = "SAF"
	LoadBalancerPoolGetResponseCheckRegionSas        LoadBalancerPoolGetResponseCheckRegion = "SAS"
	LoadBalancerPoolGetResponseCheckRegionSeas       LoadBalancerPoolGetResponseCheckRegion = "SEAS"
	LoadBalancerPoolGetResponseCheckRegionNeas       LoadBalancerPoolGetResponseCheckRegion = "NEAS"
	LoadBalancerPoolGetResponseCheckRegionAllRegions LoadBalancerPoolGetResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolGetResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolGetResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolGetResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolGetResponseLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolGetResponseLoadSheddingJSON contains the JSON metadata for the
// struct [LoadBalancerPoolGetResponseLoadShedding]
type loadBalancerPoolGetResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolGetResponseLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolGetResponseLoadSheddingDefaultPolicyRandom LoadBalancerPoolGetResponseLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolGetResponseLoadSheddingDefaultPolicyHash   LoadBalancerPoolGetResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolGetResponseLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolGetResponseLoadSheddingSessionPolicyHash LoadBalancerPoolGetResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolGetResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolGetResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolGetResponseNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolGetResponseNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolGetResponseNotificationFilterJSON contains the JSON metadata for
// the struct [LoadBalancerPoolGetResponseNotificationFilter]
type loadBalancerPoolGetResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolGetResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                    `json:"healthy,nullable"`
	JSON    loadBalancerPoolGetResponseNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolGetResponseNotificationFilterOriginJSON contains the JSON
// metadata for the struct [LoadBalancerPoolGetResponseNotificationFilterOrigin]
type loadBalancerPoolGetResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolGetResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                  `json:"healthy,nullable"`
	JSON    loadBalancerPoolGetResponseNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolGetResponseNotificationFilterPoolJSON contains the JSON metadata
// for the struct [LoadBalancerPoolGetResponseNotificationFilterPool]
type loadBalancerPoolGetResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolGetResponseOriginSteering struct {
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
	Policy LoadBalancerPoolGetResponseOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolGetResponseOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolGetResponseOriginSteeringJSON contains the JSON metadata for the
// struct [LoadBalancerPoolGetResponseOriginSteering]
type loadBalancerPoolGetResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type LoadBalancerPoolGetResponseOriginSteeringPolicy string

const (
	LoadBalancerPoolGetResponseOriginSteeringPolicyRandom                   LoadBalancerPoolGetResponseOriginSteeringPolicy = "random"
	LoadBalancerPoolGetResponseOriginSteeringPolicyHash                     LoadBalancerPoolGetResponseOriginSteeringPolicy = "hash"
	LoadBalancerPoolGetResponseOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolGetResponseOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolGetResponseOriginSteeringPolicyLeastConnections         LoadBalancerPoolGetResponseOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolGetResponseOrigin struct {
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
	Header LoadBalancerPoolGetResponseOriginsHeader `json:"header"`
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
	Weight float64                               `json:"weight"`
	JSON   loadBalancerPoolGetResponseOriginJSON `json:"-"`
}

// loadBalancerPoolGetResponseOriginJSON contains the JSON metadata for the struct
// [LoadBalancerPoolGetResponseOrigin]
type loadBalancerPoolGetResponseOriginJSON struct {
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

func (r *LoadBalancerPoolGetResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolGetResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                     `json:"Host"`
	JSON loadBalancerPoolGetResponseOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolGetResponseOriginsHeaderJSON contains the JSON metadata for the
// struct [LoadBalancerPoolGetResponseOriginsHeader]
type loadBalancerPoolGetResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolNewParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]LoadBalancerPoolNewParamsOrigin] `json:"origins,required"`
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
	LoadShedding param.Field[LoadBalancerPoolNewParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[LoadBalancerPoolNewParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[LoadBalancerPoolNewParamsOriginSteering] `json:"origin_steering"`
}

func (r LoadBalancerPoolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerPoolNewParamsOrigin struct {
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
	Header param.Field[LoadBalancerPoolNewParamsOriginsHeader] `json:"header"`
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

func (r LoadBalancerPoolNewParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolNewParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r LoadBalancerPoolNewParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolNewParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[LoadBalancerPoolNewParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[LoadBalancerPoolNewParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r LoadBalancerPoolNewParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolNewParamsLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolNewParamsLoadSheddingDefaultPolicyRandom LoadBalancerPoolNewParamsLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolNewParamsLoadSheddingDefaultPolicyHash   LoadBalancerPoolNewParamsLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolNewParamsLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolNewParamsLoadSheddingSessionPolicyHash LoadBalancerPoolNewParamsLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolNewParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[LoadBalancerPoolNewParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[LoadBalancerPoolNewParamsNotificationFilterPool] `json:"pool"`
}

func (r LoadBalancerPoolNewParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolNewParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r LoadBalancerPoolNewParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolNewParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r LoadBalancerPoolNewParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolNewParamsOriginSteering struct {
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
	Policy param.Field[LoadBalancerPoolNewParamsOriginSteeringPolicy] `json:"policy"`
}

func (r LoadBalancerPoolNewParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type LoadBalancerPoolNewParamsOriginSteeringPolicy string

const (
	LoadBalancerPoolNewParamsOriginSteeringPolicyRandom                   LoadBalancerPoolNewParamsOriginSteeringPolicy = "random"
	LoadBalancerPoolNewParamsOriginSteeringPolicyHash                     LoadBalancerPoolNewParamsOriginSteeringPolicy = "hash"
	LoadBalancerPoolNewParamsOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolNewParamsOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolNewParamsOriginSteeringPolicyLeastConnections         LoadBalancerPoolNewParamsOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolNewResponseEnvelope struct {
	Errors   []LoadBalancerPoolNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPoolNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerPoolNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPoolNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPoolNewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPoolNewResponseEnvelope]
type loadBalancerPoolNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    loadBalancerPoolNewResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LoadBalancerPoolNewResponseEnvelopeErrors]
type loadBalancerPoolNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    loadBalancerPoolNewResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LoadBalancerPoolNewResponseEnvelopeMessages]
type loadBalancerPoolNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolNewResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolNewResponseEnvelopeSuccessTrue LoadBalancerPoolNewResponseEnvelopeSuccess = true
)

type LoadBalancerPoolUpdateParams struct {
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
	NotificationFilter param.Field[LoadBalancerPoolUpdateParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[LoadBalancerPoolUpdateParamsOriginSteering] `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]LoadBalancerPoolUpdateParamsOrigin] `json:"origins"`
}

func (r LoadBalancerPoolUpdateParams) MarshalJSON() (data []byte, err error) {
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

type LoadBalancerPoolUpdateResponseEnvelope struct {
	Errors   []LoadBalancerPoolUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPoolUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerPoolUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPoolUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPoolUpdateResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPoolUpdateResponseEnvelope]
type loadBalancerPoolUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    loadBalancerPoolUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerPoolUpdateResponseEnvelopeErrors]
type loadBalancerPoolUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    loadBalancerPoolUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerPoolUpdateResponseEnvelopeMessages]
type loadBalancerPoolUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolUpdateResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolUpdateResponseEnvelopeSuccessTrue LoadBalancerPoolUpdateResponseEnvelopeSuccess = true
)

type LoadBalancerPoolListParams struct {
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `query:"monitor"`
}

// URLQuery serializes [LoadBalancerPoolListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerPoolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LoadBalancerPoolListResponseEnvelope struct {
	Errors   []LoadBalancerPoolListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPoolListResponseEnvelopeMessages `json:"messages,required"`
	Result   []LoadBalancerPoolListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerPoolListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerPoolListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerPoolListResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerPoolListResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPoolListResponseEnvelope]
type loadBalancerPoolListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    loadBalancerPoolListResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerPoolListResponseEnvelopeErrors]
type loadBalancerPoolListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    loadBalancerPoolListResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LoadBalancerPoolListResponseEnvelopeMessages]
type loadBalancerPoolListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolListResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolListResponseEnvelopeSuccessTrue LoadBalancerPoolListResponseEnvelopeSuccess = true
)

type LoadBalancerPoolListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       loadBalancerPoolListResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerPoolListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [LoadBalancerPoolListResponseEnvelopeResultInfo]
type loadBalancerPoolListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolDeleteResponseEnvelope struct {
	Errors   []LoadBalancerPoolDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPoolDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerPoolDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPoolDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPoolDeleteResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPoolDeleteResponseEnvelope]
type loadBalancerPoolDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    loadBalancerPoolDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerPoolDeleteResponseEnvelopeErrors]
type loadBalancerPoolDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    loadBalancerPoolDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerPoolDeleteResponseEnvelopeMessages]
type loadBalancerPoolDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolDeleteResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolDeleteResponseEnvelopeSuccessTrue LoadBalancerPoolDeleteResponseEnvelopeSuccess = true
)

type LoadBalancerPoolGetResponseEnvelope struct {
	Errors   []LoadBalancerPoolGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPoolGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerPoolGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPoolGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPoolGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPoolGetResponseEnvelope]
type loadBalancerPoolGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    loadBalancerPoolGetResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LoadBalancerPoolGetResponseEnvelopeErrors]
type loadBalancerPoolGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    loadBalancerPoolGetResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LoadBalancerPoolGetResponseEnvelopeMessages]
type loadBalancerPoolGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolGetResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolGetResponseEnvelopeSuccessTrue LoadBalancerPoolGetResponseEnvelopeSuccess = true
)
