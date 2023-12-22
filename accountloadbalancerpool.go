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
func (r *AccountLoadBalancerPoolService) Get(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *PoolSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured pool.
func (r *AccountLoadBalancerPoolService) Update(ctx context.Context, accountIdentifier string, identifier interface{}, body AccountLoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *PoolSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured pool.
func (r *AccountLoadBalancerPoolService) Delete(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *SchemasIDResponseLwyhetfd, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new pool.
func (r *AccountLoadBalancerPoolService) AccountLoadBalancerPoolsNewPool(ctx context.Context, accountIdentifier string, body AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParams, opts ...option.RequestOption) (res *PoolSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured pools.
func (r *AccountLoadBalancerPoolService) AccountLoadBalancerPoolsListPools(ctx context.Context, accountIdentifier string, query AccountLoadBalancerPoolAccountLoadBalancerPoolsListPoolsParams, opts ...option.RequestOption) (res *PoolResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Apply changes to a number of existing pools, overwriting the supplied
// properties. Pools are ordered by ascending `name`. Returns the list of affected
// pools. Supports the standard pagination query parameters, either
// `limit`/`offset` or `per_page`/`page`.
func (r *AccountLoadBalancerPoolService) AccountLoadBalancerPoolsPatchPools(ctx context.Context, accountIdentifier string, body AccountLoadBalancerPoolAccountLoadBalancerPoolsPatchPoolsParams, opts ...option.RequestOption) (res *PoolResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type PoolResponseCollection struct {
	Errors     []PoolResponseCollectionError    `json:"errors"`
	Messages   []PoolResponseCollectionMessage  `json:"messages"`
	Result     []PoolResponseCollectionResult   `json:"result"`
	ResultInfo PoolResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success PoolResponseCollectionSuccess `json:"success"`
	JSON    poolResponseCollectionJSON    `json:"-"`
}

// poolResponseCollectionJSON contains the JSON metadata for the struct
// [PoolResponseCollection]
type poolResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoolResponseCollectionError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    poolResponseCollectionErrorJSON `json:"-"`
}

// poolResponseCollectionErrorJSON contains the JSON metadata for the struct
// [PoolResponseCollectionError]
type poolResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoolResponseCollectionMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    poolResponseCollectionMessageJSON `json:"-"`
}

// poolResponseCollectionMessageJSON contains the JSON metadata for the struct
// [PoolResponseCollectionMessage]
type poolResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoolResponseCollectionResult struct {
	ID interface{} `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []PoolResponseCollectionResultCheckRegion `json:"check_regions,nullable"`
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
	LoadShedding PoolResponseCollectionResultLoadShedding `json:"load_shedding"`
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
	// The email address to send health status notifications to. This can be an
	// individual mailbox or a mailing list. Multiple emails can be supplied as a comma
	// delimited list.
	NotificationEmail string `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter PoolResponseCollectionResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering PoolResponseCollectionResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []PoolResponseCollectionResultOrigin `json:"origins"`
	JSON    poolResponseCollectionResultJSON     `json:"-"`
}

// poolResponseCollectionResultJSON contains the JSON metadata for the struct
// [PoolResponseCollectionResult]
type poolResponseCollectionResultJSON struct {
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

func (r *PoolResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolResponseCollectionResultCheckRegion string

const (
	PoolResponseCollectionResultCheckRegionWnam       PoolResponseCollectionResultCheckRegion = "WNAM"
	PoolResponseCollectionResultCheckRegionEnam       PoolResponseCollectionResultCheckRegion = "ENAM"
	PoolResponseCollectionResultCheckRegionWeu        PoolResponseCollectionResultCheckRegion = "WEU"
	PoolResponseCollectionResultCheckRegionEeu        PoolResponseCollectionResultCheckRegion = "EEU"
	PoolResponseCollectionResultCheckRegionNsam       PoolResponseCollectionResultCheckRegion = "NSAM"
	PoolResponseCollectionResultCheckRegionSsam       PoolResponseCollectionResultCheckRegion = "SSAM"
	PoolResponseCollectionResultCheckRegionOc         PoolResponseCollectionResultCheckRegion = "OC"
	PoolResponseCollectionResultCheckRegionMe         PoolResponseCollectionResultCheckRegion = "ME"
	PoolResponseCollectionResultCheckRegionNaf        PoolResponseCollectionResultCheckRegion = "NAF"
	PoolResponseCollectionResultCheckRegionSaf        PoolResponseCollectionResultCheckRegion = "SAF"
	PoolResponseCollectionResultCheckRegionSas        PoolResponseCollectionResultCheckRegion = "SAS"
	PoolResponseCollectionResultCheckRegionSeas       PoolResponseCollectionResultCheckRegion = "SEAS"
	PoolResponseCollectionResultCheckRegionNeas       PoolResponseCollectionResultCheckRegion = "NEAS"
	PoolResponseCollectionResultCheckRegionAllRegions PoolResponseCollectionResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type PoolResponseCollectionResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy PoolResponseCollectionResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy PoolResponseCollectionResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          poolResponseCollectionResultLoadSheddingJSON          `json:"-"`
}

// poolResponseCollectionResultLoadSheddingJSON contains the JSON metadata for the
// struct [PoolResponseCollectionResultLoadShedding]
type poolResponseCollectionResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PoolResponseCollectionResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolResponseCollectionResultLoadSheddingDefaultPolicy string

const (
	PoolResponseCollectionResultLoadSheddingDefaultPolicyRandom PoolResponseCollectionResultLoadSheddingDefaultPolicy = "random"
	PoolResponseCollectionResultLoadSheddingDefaultPolicyHash   PoolResponseCollectionResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolResponseCollectionResultLoadSheddingSessionPolicy string

const (
	PoolResponseCollectionResultLoadSheddingSessionPolicyHash PoolResponseCollectionResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolResponseCollectionResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin PoolResponseCollectionResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool PoolResponseCollectionResultNotificationFilterPool `json:"pool,nullable"`
	JSON poolResponseCollectionResultNotificationFilterJSON `json:"-"`
}

// poolResponseCollectionResultNotificationFilterJSON contains the JSON metadata
// for the struct [PoolResponseCollectionResultNotificationFilter]
type poolResponseCollectionResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResponseCollectionResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolResponseCollectionResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                     `json:"healthy,nullable"`
	JSON    poolResponseCollectionResultNotificationFilterOriginJSON `json:"-"`
}

// poolResponseCollectionResultNotificationFilterOriginJSON contains the JSON
// metadata for the struct [PoolResponseCollectionResultNotificationFilterOrigin]
type poolResponseCollectionResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResponseCollectionResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolResponseCollectionResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                   `json:"healthy,nullable"`
	JSON    poolResponseCollectionResultNotificationFilterPoolJSON `json:"-"`
}

// poolResponseCollectionResultNotificationFilterPoolJSON contains the JSON
// metadata for the struct [PoolResponseCollectionResultNotificationFilterPool]
type poolResponseCollectionResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResponseCollectionResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolResponseCollectionResultOriginSteering struct {
	// The type of origin steering policy to use, either "random" or "hash" (based on
	// CF-Connecting-IP).
	Policy PoolResponseCollectionResultOriginSteeringPolicy `json:"policy"`
	JSON   poolResponseCollectionResultOriginSteeringJSON   `json:"-"`
}

// poolResponseCollectionResultOriginSteeringJSON contains the JSON metadata for
// the struct [PoolResponseCollectionResultOriginSteering]
type poolResponseCollectionResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResponseCollectionResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of origin steering policy to use, either "random" or "hash" (based on
// CF-Connecting-IP).
type PoolResponseCollectionResultOriginSteeringPolicy string

const (
	PoolResponseCollectionResultOriginSteeringPolicyRandom PoolResponseCollectionResultOriginSteeringPolicy = "random"
	PoolResponseCollectionResultOriginSteeringPolicyHash   PoolResponseCollectionResultOriginSteeringPolicy = "hash"
)

type PoolResponseCollectionResultOrigin struct {
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
	Header PoolResponseCollectionResultOriginsHeader `json:"header"`
	// A human-identifiable name for the origin.
	Name string `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID string `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	Weight float64                                `json:"weight"`
	JSON   poolResponseCollectionResultOriginJSON `json:"-"`
}

// poolResponseCollectionResultOriginJSON contains the JSON metadata for the struct
// [PoolResponseCollectionResultOrigin]
type poolResponseCollectionResultOriginJSON struct {
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

func (r *PoolResponseCollectionResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolResponseCollectionResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                      `json:"Host"`
	JSON poolResponseCollectionResultOriginsHeaderJSON `json:"-"`
}

// poolResponseCollectionResultOriginsHeaderJSON contains the JSON metadata for the
// struct [PoolResponseCollectionResultOriginsHeader]
type poolResponseCollectionResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResponseCollectionResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoolResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       poolResponseCollectionResultInfoJSON `json:"-"`
}

// poolResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [PoolResponseCollectionResultInfo]
type poolResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PoolResponseCollectionSuccess bool

const (
	PoolResponseCollectionSuccessTrue PoolResponseCollectionSuccess = true
)

type PoolSingleResponse struct {
	Errors   []PoolSingleResponseError   `json:"errors"`
	Messages []PoolSingleResponseMessage `json:"messages"`
	Result   PoolSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success PoolSingleResponseSuccess `json:"success"`
	JSON    poolSingleResponseJSON    `json:"-"`
}

// poolSingleResponseJSON contains the JSON metadata for the struct
// [PoolSingleResponse]
type poolSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoolSingleResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    poolSingleResponseErrorJSON `json:"-"`
}

// poolSingleResponseErrorJSON contains the JSON metadata for the struct
// [PoolSingleResponseError]
type poolSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoolSingleResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    poolSingleResponseMessageJSON `json:"-"`
}

// poolSingleResponseMessageJSON contains the JSON metadata for the struct
// [PoolSingleResponseMessage]
type poolSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoolSingleResponseResult struct {
	ID interface{} `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []PoolSingleResponseResultCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time                             `json:"created_on" format:"date-time"`
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
	LoadShedding PoolSingleResponseResultLoadShedding `json:"load_shedding"`
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
	// The email address to send health status notifications to. This can be an
	// individual mailbox or a mailing list. Multiple emails can be supplied as a comma
	// delimited list.
	NotificationEmail string `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter PoolSingleResponseResultNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering PoolSingleResponseResultOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []PoolSingleResponseResultOrigin `json:"origins"`
	JSON    poolSingleResponseResultJSON     `json:"-"`
}

// poolSingleResponseResultJSON contains the JSON metadata for the struct
// [PoolSingleResponseResult]
type poolSingleResponseResultJSON struct {
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

func (r *PoolSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolSingleResponseResultCheckRegion string

const (
	PoolSingleResponseResultCheckRegionWnam       PoolSingleResponseResultCheckRegion = "WNAM"
	PoolSingleResponseResultCheckRegionEnam       PoolSingleResponseResultCheckRegion = "ENAM"
	PoolSingleResponseResultCheckRegionWeu        PoolSingleResponseResultCheckRegion = "WEU"
	PoolSingleResponseResultCheckRegionEeu        PoolSingleResponseResultCheckRegion = "EEU"
	PoolSingleResponseResultCheckRegionNsam       PoolSingleResponseResultCheckRegion = "NSAM"
	PoolSingleResponseResultCheckRegionSsam       PoolSingleResponseResultCheckRegion = "SSAM"
	PoolSingleResponseResultCheckRegionOc         PoolSingleResponseResultCheckRegion = "OC"
	PoolSingleResponseResultCheckRegionMe         PoolSingleResponseResultCheckRegion = "ME"
	PoolSingleResponseResultCheckRegionNaf        PoolSingleResponseResultCheckRegion = "NAF"
	PoolSingleResponseResultCheckRegionSaf        PoolSingleResponseResultCheckRegion = "SAF"
	PoolSingleResponseResultCheckRegionSas        PoolSingleResponseResultCheckRegion = "SAS"
	PoolSingleResponseResultCheckRegionSeas       PoolSingleResponseResultCheckRegion = "SEAS"
	PoolSingleResponseResultCheckRegionNeas       PoolSingleResponseResultCheckRegion = "NEAS"
	PoolSingleResponseResultCheckRegionAllRegions PoolSingleResponseResultCheckRegion = "ALL_REGIONS"
)

// Configures load shedding policies and percentages for the pool.
type PoolSingleResponseResultLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy PoolSingleResponseResultLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy PoolSingleResponseResultLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          poolSingleResponseResultLoadSheddingJSON          `json:"-"`
}

// poolSingleResponseResultLoadSheddingJSON contains the JSON metadata for the
// struct [PoolSingleResponseResultLoadShedding]
type poolSingleResponseResultLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PoolSingleResponseResultLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolSingleResponseResultLoadSheddingDefaultPolicy string

const (
	PoolSingleResponseResultLoadSheddingDefaultPolicyRandom PoolSingleResponseResultLoadSheddingDefaultPolicy = "random"
	PoolSingleResponseResultLoadSheddingDefaultPolicyHash   PoolSingleResponseResultLoadSheddingDefaultPolicy = "hash"
)

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolSingleResponseResultLoadSheddingSessionPolicy string

const (
	PoolSingleResponseResultLoadSheddingSessionPolicyHash PoolSingleResponseResultLoadSheddingSessionPolicy = "hash"
)

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolSingleResponseResultNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin PoolSingleResponseResultNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool PoolSingleResponseResultNotificationFilterPool `json:"pool,nullable"`
	JSON poolSingleResponseResultNotificationFilterJSON `json:"-"`
}

// poolSingleResponseResultNotificationFilterJSON contains the JSON metadata for
// the struct [PoolSingleResponseResultNotificationFilter]
type poolSingleResponseResultNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolSingleResponseResultNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolSingleResponseResultNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                                 `json:"healthy,nullable"`
	JSON    poolSingleResponseResultNotificationFilterOriginJSON `json:"-"`
}

// poolSingleResponseResultNotificationFilterOriginJSON contains the JSON metadata
// for the struct [PoolSingleResponseResultNotificationFilterOrigin]
type poolSingleResponseResultNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolSingleResponseResultNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolSingleResponseResultNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                                               `json:"healthy,nullable"`
	JSON    poolSingleResponseResultNotificationFilterPoolJSON `json:"-"`
}

// poolSingleResponseResultNotificationFilterPoolJSON contains the JSON metadata
// for the struct [PoolSingleResponseResultNotificationFilterPool]
type poolSingleResponseResultNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolSingleResponseResultNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolSingleResponseResultOriginSteering struct {
	// The type of origin steering policy to use, either "random" or "hash" (based on
	// CF-Connecting-IP).
	Policy PoolSingleResponseResultOriginSteeringPolicy `json:"policy"`
	JSON   poolSingleResponseResultOriginSteeringJSON   `json:"-"`
}

// poolSingleResponseResultOriginSteeringJSON contains the JSON metadata for the
// struct [PoolSingleResponseResultOriginSteering]
type poolSingleResponseResultOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolSingleResponseResultOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of origin steering policy to use, either "random" or "hash" (based on
// CF-Connecting-IP).
type PoolSingleResponseResultOriginSteeringPolicy string

const (
	PoolSingleResponseResultOriginSteeringPolicyRandom PoolSingleResponseResultOriginSteeringPolicy = "random"
	PoolSingleResponseResultOriginSteeringPolicyHash   PoolSingleResponseResultOriginSteeringPolicy = "hash"
)

type PoolSingleResponseResultOrigin struct {
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
	Header PoolSingleResponseResultOriginsHeader `json:"header"`
	// A human-identifiable name for the origin.
	Name string `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID string `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	Weight float64                            `json:"weight"`
	JSON   poolSingleResponseResultOriginJSON `json:"-"`
}

// poolSingleResponseResultOriginJSON contains the JSON metadata for the struct
// [PoolSingleResponseResultOrigin]
type poolSingleResponseResultOriginJSON struct {
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

func (r *PoolSingleResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolSingleResponseResultOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string                                  `json:"Host"`
	JSON poolSingleResponseResultOriginsHeaderJSON `json:"-"`
}

// poolSingleResponseResultOriginsHeaderJSON contains the JSON metadata for the
// struct [PoolSingleResponseResultOriginsHeader]
type poolSingleResponseResultOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolSingleResponseResultOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PoolSingleResponseSuccess bool

const (
	PoolSingleResponseSuccessTrue PoolSingleResponseSuccess = true
)

type SchemasIDResponseLwyhetfd struct {
	Errors   []SchemasIDResponseLwyhetfdError   `json:"errors"`
	Messages []SchemasIDResponseLwyhetfdMessage `json:"messages"`
	Result   SchemasIDResponseLwyhetfdResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasIDResponseLwyhetfdSuccess `json:"success"`
	JSON    schemasIDResponseLwyhetfdJSON    `json:"-"`
}

// schemasIDResponseLwyhetfdJSON contains the JSON metadata for the struct
// [SchemasIDResponseLwyhetfd]
type schemasIDResponseLwyhetfdJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseLwyhetfd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasIDResponseLwyhetfdError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    schemasIDResponseLwyhetfdErrorJSON `json:"-"`
}

// schemasIDResponseLwyhetfdErrorJSON contains the JSON metadata for the struct
// [SchemasIDResponseLwyhetfdError]
type schemasIDResponseLwyhetfdErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseLwyhetfdError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasIDResponseLwyhetfdMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    schemasIDResponseLwyhetfdMessageJSON `json:"-"`
}

// schemasIDResponseLwyhetfdMessageJSON contains the JSON metadata for the struct
// [SchemasIDResponseLwyhetfdMessage]
type schemasIDResponseLwyhetfdMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseLwyhetfdMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasIDResponseLwyhetfdResult struct {
	ID   interface{}                         `json:"id"`
	JSON schemasIDResponseLwyhetfdResultJSON `json:"-"`
}

// schemasIDResponseLwyhetfdResultJSON contains the JSON metadata for the struct
// [SchemasIDResponseLwyhetfdResult]
type schemasIDResponseLwyhetfdResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseLwyhetfdResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasIDResponseLwyhetfdSuccess bool

const (
	SchemasIDResponseLwyhetfdSuccessTrue SchemasIDResponseLwyhetfdSuccess = true
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
	// The email address to send health status notifications to. This can be an
	// individual mailbox or a mailing list. Multiple emails can be supplied as a comma
	// delimited list.
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
	// The type of origin steering policy to use, either "random" or "hash" (based on
	// CF-Connecting-IP).
	Policy param.Field[AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicy] `json:"policy"`
}

func (r AccountLoadBalancerPoolUpdateParamsOriginSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of origin steering policy to use, either "random" or "hash" (based on
// CF-Connecting-IP).
type AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicy string

const (
	AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicyRandom AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "random"
	AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicyHash   AccountLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "hash"
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
	// The email address to send health status notifications to. This can be an
	// individual mailbox or a mailing list. Multiple emails can be supplied as a comma
	// delimited list.
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
	// The type of origin steering policy to use, either "random" or "hash" (based on
	// CF-Connecting-IP).
	Policy param.Field[AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy] `json:"policy"`
}

func (r AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of origin steering policy to use, either "random" or "hash" (based on
// CF-Connecting-IP).
type AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy string

const (
	AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyRandom AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "random"
	AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyHash   AccountLoadBalancerPoolAccountLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "hash"
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
