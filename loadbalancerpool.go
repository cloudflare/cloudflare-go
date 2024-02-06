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
	var env LoadBalancerPoolGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured pool.
func (r *LoadBalancerPoolService) Update(ctx context.Context, accountIdentifier string, identifier string, body LoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *LoadBalancerPoolUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured pool.
func (r *LoadBalancerPoolService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *LoadBalancerPoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a new pool.
func (r *LoadBalancerPoolService) AccountLoadBalancerPoolsNewPool(ctx context.Context, accountIdentifier string, body LoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams, opts ...option.RequestOption) (res *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured pools.
func (r *LoadBalancerPoolService) AccountLoadBalancerPoolsListPools(ctx context.Context, accountIdentifier string, query LoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams, opts ...option.RequestOption) (res *[]LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
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
func (r *LoadBalancerPoolService) AccountLoadBalancerPoolsPatchPools(ctx context.Context, accountIdentifier string, body LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams, opts ...option.RequestOption) (res *[]LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion `json:"check_regions,nullable"`
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
	LoadShedding LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOrigin `json:"origins"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseJSON     `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseJSON contains the JSON
// metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseJSON struct {
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

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionWnam       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "WNAM"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionEnam       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "ENAM"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionWeu        LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "WEU"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionEeu        LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "EEU"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionNsam       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "NSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionSsam       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "SSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionOc         LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "OC"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionMe         LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "ME"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionNaf        LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "NAF"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionSaf        LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "SAF"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionSas        LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "SAS"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionSeas       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "SEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionNeas       LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "NEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegionAllRegions LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingJSON contains
// the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadShedding]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicyRandom LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicyHash   LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingSessionPolicyHash LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilter]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                              `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterPool]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteering struct {
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
	Policy LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteering]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringPolicyRandom                   LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringPolicyHash                     LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy = "hash"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringPolicyLeastConnections         LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOrigin struct {
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
	Header LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginsHeader `json:"header"`
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
	JSON   loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginJSON contains the
// JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginJSON struct {
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

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                 `json:"Host"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginsHeaderJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginsHeader]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                              `json:"created_on" format:"date-time"`
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
	LoadShedding LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOrigin `json:"origins"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseJSON     `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseJSON contains the JSON
// metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseJSON struct {
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

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionWnam       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "WNAM"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionEnam       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "ENAM"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionWeu        LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "WEU"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionEeu        LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "EEU"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionNsam       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "NSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionSsam       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "SSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionOc         LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "OC"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionMe         LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "ME"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionNaf        LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "NAF"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionSaf        LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "SAF"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionSas        LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "SAS"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionSeas       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "SEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionNeas       LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "NEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegionAllRegions LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadShedding]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicyRandom LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicyHash   LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingSessionPolicyHash LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilter]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                  `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterPool]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteering struct {
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
	Policy LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteering]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringPolicyRandom                   LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringPolicyHash                     LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy = "hash"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringPolicyLeastConnections         LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOrigin struct {
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
	Header LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginsHeader `json:"header"`
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
	Weight float64                                                             `json:"weight"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginJSON contains the
// JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginJSON struct {
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

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                   `json:"Host"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginsHeaderJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginsHeader]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                                                               `json:"created_on" format:"date-time"`
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
	LoadShedding LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadShedding `json:"load_shedding"`
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
	NotificationFilter LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOrigin `json:"origins"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseJSON     `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseJSON contains the JSON
// metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseJSON struct {
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

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionWnam       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "WNAM"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionEnam       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "ENAM"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionWeu        LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "WEU"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionEeu        LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "EEU"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionNsam       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "NSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionSsam       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "SSAM"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionOc         LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "OC"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionMe         LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "ME"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionNaf        LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "NAF"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionSaf        LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "SAF"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionSas        LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "SAS"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionSeas       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "SEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionNeas       LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "NEAS"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegionAllRegions LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingJSON          `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadShedding]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicyRandom LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicyHash   LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingSessionPolicyHash LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterPool `json:"pool,nullable"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilter]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                   `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterOriginJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                                                 `json:"healthy,nullable"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterPoolJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterPoolJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterPool]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteering struct {
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
	Policy LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy `json:"policy"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringJSON   `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteering]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteering) UnmarshalJSON(data []byte) (err error) {
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
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy string

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicyRandom                   LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy = "random"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicyHash                     LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy = "hash"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicyLeastConnections         LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginSteeringPolicy = "least_connections"
)

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOrigin struct {
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
	Header LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginsHeader `json:"header"`
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
	Weight float64                                                              `json:"weight"`
	JSON   loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginJSON contains
// the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOrigin]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginJSON struct {
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

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                                                    `json:"Host"`
	JSON loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginsHeaderJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginsHeaderJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginsHeader]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolGetResponseEnvelope struct {
	Errors   []LoadBalancerPoolGetResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerPoolGetResponseEnvelopeMessages `json:"messages"`
	Result   LoadBalancerPoolGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerPoolGetResponseEnvelopeSuccess `json:"success"`
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

type LoadBalancerPoolUpdateResponseEnvelope struct {
	Errors   []LoadBalancerPoolUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerPoolUpdateResponseEnvelopeMessages `json:"messages"`
	Result   LoadBalancerPoolUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerPoolUpdateResponseEnvelopeSuccess `json:"success"`
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

type LoadBalancerPoolDeleteResponseEnvelope struct {
	Errors   []LoadBalancerPoolDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerPoolDeleteResponseEnvelopeMessages `json:"messages"`
	Result   LoadBalancerPoolDeleteResponse                   `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerPoolDeleteResponseEnvelopeSuccess `json:"success"`
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

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelope struct {
	Errors   []LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeMessages `json:"messages"`
	Result   LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponse                   `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeSuccess `json:"success"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelope]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeErrors]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeMessages]
type loadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeSuccessTrue LoadBalancerPoolAccountLoadBalancerPoolsNewPoolResponseEnvelopeSuccess = true
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

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelope struct {
	Errors     []LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeErrors   `json:"errors"`
	Messages   []LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeMessages `json:"messages"`
	Result     []LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponse                 `json:"result"`
	ResultInfo LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeSuccess `json:"success"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelope]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeErrors]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeMessages]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeResultInfo]
type loadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeSuccessTrue LoadBalancerPoolAccountLoadBalancerPoolsListPoolsResponseEnvelopeSuccess = true
)

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

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelope struct {
	Errors     []LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeErrors   `json:"errors"`
	Messages   []LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeMessages `json:"messages"`
	Result     []LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponse                 `json:"result"`
	ResultInfo LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeSuccess `json:"success"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelope]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeErrors struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeErrors]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeMessages struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeMessages]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfo]
type loadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeSuccessTrue LoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsResponseEnvelopeSuccess = true
)
