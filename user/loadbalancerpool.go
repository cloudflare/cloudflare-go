// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// LoadBalancerPoolService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerPoolService] method
// instead.
type LoadBalancerPoolService struct {
	Options []option.RequestOption
}

// NewLoadBalancerPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerPoolService(opts ...option.RequestOption) (r *LoadBalancerPoolService) {
	r = &LoadBalancerPoolService{}
	r.Options = opts
	return
}

// Create a new pool.
func (r *LoadBalancerPoolService) New(ctx context.Context, body LoadBalancerPoolNewParams, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolNewResponseEnvelope
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a configured pool.
func (r *LoadBalancerPoolService) Update(ctx context.Context, poolID string, body LoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolUpdateResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured pools.
func (r *LoadBalancerPoolService) List(ctx context.Context, query LoadBalancerPoolListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Pool], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/load_balancers/pools"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List configured pools.
func (r *LoadBalancerPoolService) ListAutoPaging(ctx context.Context, query LoadBalancerPoolListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Pool] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured pool.
func (r *LoadBalancerPoolService) Delete(ctx context.Context, poolID string, body LoadBalancerPoolDeleteParams, opts ...option.RequestOption) (res *LoadBalancerPoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolDeleteResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing pool, overwriting the supplied properties.
func (r *LoadBalancerPoolService) Edit(ctx context.Context, poolID string, body LoadBalancerPoolEditParams, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolEditResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured pool.
func (r *LoadBalancerPoolService) Get(ctx context.Context, poolID string, opts ...option.RequestOption) (res *Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolGetResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch the latest pool health status for a single pool.
func (r *LoadBalancerPoolService) Health(ctx context.Context, poolID string, opts ...option.RequestOption) (res *LoadBalancerPoolHealthResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolHealthResponseEnvelope
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
func (r *LoadBalancerPoolService) Preview(ctx context.Context, poolID string, body LoadBalancerPoolPreviewParams, opts ...option.RequestOption) (res *LoadBalancerPoolPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolPreviewResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/preview", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the list of resources that reference the provided pool.
func (r *LoadBalancerPoolService) References(ctx context.Context, poolID string, opts ...option.RequestOption) (res *[]LoadBalancerPoolReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolReferencesResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/references", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Pool struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []PoolCheckRegion `json:"check_regions,nullable"`
	CreatedOn    time.Time         `json:"created_on" format:"date-time"`
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
	LoadShedding PoolLoadShedding `json:"load_shedding"`
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
	NotificationFilter PoolNotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering PoolOriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []PoolOrigin `json:"origins"`
	JSON    poolJSON     `json:"-"`
}

// poolJSON contains the JSON metadata for the struct [Pool]
type poolJSON struct {
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

func (r *Pool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolJSON) RawJSON() string {
	return r.raw
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type PoolCheckRegion string

const (
	PoolCheckRegionWnam       PoolCheckRegion = "WNAM"
	PoolCheckRegionEnam       PoolCheckRegion = "ENAM"
	PoolCheckRegionWeu        PoolCheckRegion = "WEU"
	PoolCheckRegionEeu        PoolCheckRegion = "EEU"
	PoolCheckRegionNsam       PoolCheckRegion = "NSAM"
	PoolCheckRegionSsam       PoolCheckRegion = "SSAM"
	PoolCheckRegionOc         PoolCheckRegion = "OC"
	PoolCheckRegionMe         PoolCheckRegion = "ME"
	PoolCheckRegionNaf        PoolCheckRegion = "NAF"
	PoolCheckRegionSaf        PoolCheckRegion = "SAF"
	PoolCheckRegionSas        PoolCheckRegion = "SAS"
	PoolCheckRegionSeas       PoolCheckRegion = "SEAS"
	PoolCheckRegionNeas       PoolCheckRegion = "NEAS"
	PoolCheckRegionAllRegions PoolCheckRegion = "ALL_REGIONS"
)

func (r PoolCheckRegion) IsKnown() bool {
	switch r {
	case PoolCheckRegionWnam, PoolCheckRegionEnam, PoolCheckRegionWeu, PoolCheckRegionEeu, PoolCheckRegionNsam, PoolCheckRegionSsam, PoolCheckRegionOc, PoolCheckRegionMe, PoolCheckRegionNaf, PoolCheckRegionSaf, PoolCheckRegionSas, PoolCheckRegionSeas, PoolCheckRegionNeas, PoolCheckRegionAllRegions:
		return true
	}
	return false
}

// Configures load shedding policies and percentages for the pool.
type PoolLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy PoolLoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy PoolLoadSheddingSessionPolicy `json:"session_policy"`
	JSON          poolLoadSheddingJSON          `json:"-"`
}

// poolLoadSheddingJSON contains the JSON metadata for the struct
// [PoolLoadShedding]
type poolLoadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PoolLoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolLoadSheddingJSON) RawJSON() string {
	return r.raw
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type PoolLoadSheddingDefaultPolicy string

const (
	PoolLoadSheddingDefaultPolicyRandom PoolLoadSheddingDefaultPolicy = "random"
	PoolLoadSheddingDefaultPolicyHash   PoolLoadSheddingDefaultPolicy = "hash"
)

func (r PoolLoadSheddingDefaultPolicy) IsKnown() bool {
	switch r {
	case PoolLoadSheddingDefaultPolicyRandom, PoolLoadSheddingDefaultPolicyHash:
		return true
	}
	return false
}

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolLoadSheddingSessionPolicy string

const (
	PoolLoadSheddingSessionPolicyHash PoolLoadSheddingSessionPolicy = "hash"
)

func (r PoolLoadSheddingSessionPolicy) IsKnown() bool {
	switch r {
	case PoolLoadSheddingSessionPolicyHash:
		return true
	}
	return false
}

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type PoolNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin PoolNotificationFilterOrigin `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool PoolNotificationFilterPool `json:"pool,nullable"`
	JSON poolNotificationFilterJSON `json:"-"`
}

// poolNotificationFilterJSON contains the JSON metadata for the struct
// [PoolNotificationFilter]
type poolNotificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNotificationFilterJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                             `json:"healthy,nullable"`
	JSON    poolNotificationFilterOriginJSON `json:"-"`
}

// poolNotificationFilterOriginJSON contains the JSON metadata for the struct
// [PoolNotificationFilterOrigin]
type poolNotificationFilterOriginJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNotificationFilterOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNotificationFilterOriginJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type PoolNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool                           `json:"healthy,nullable"`
	JSON    poolNotificationFilterPoolJSON `json:"-"`
}

// poolNotificationFilterPoolJSON contains the JSON metadata for the struct
// [PoolNotificationFilterPool]
type poolNotificationFilterPoolJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolNotificationFilterPool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolNotificationFilterPoolJSON) RawJSON() string {
	return r.raw
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type PoolOriginSteering struct {
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
	Policy PoolOriginSteeringPolicy `json:"policy"`
	JSON   poolOriginSteeringJSON   `json:"-"`
}

// poolOriginSteeringJSON contains the JSON metadata for the struct
// [PoolOriginSteering]
type poolOriginSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolOriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolOriginSteeringJSON) RawJSON() string {
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
type PoolOriginSteeringPolicy string

const (
	PoolOriginSteeringPolicyRandom                   PoolOriginSteeringPolicy = "random"
	PoolOriginSteeringPolicyHash                     PoolOriginSteeringPolicy = "hash"
	PoolOriginSteeringPolicyLeastOutstandingRequests PoolOriginSteeringPolicy = "least_outstanding_requests"
	PoolOriginSteeringPolicyLeastConnections         PoolOriginSteeringPolicy = "least_connections"
)

func (r PoolOriginSteeringPolicy) IsKnown() bool {
	switch r {
	case PoolOriginSteeringPolicyRandom, PoolOriginSteeringPolicyHash, PoolOriginSteeringPolicyLeastOutstandingRequests, PoolOriginSteeringPolicyLeastConnections:
		return true
	}
	return false
}

type PoolOrigin struct {
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
	Header PoolOriginsHeader `json:"header"`
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
	Weight float64        `json:"weight"`
	JSON   poolOriginJSON `json:"-"`
}

// poolOriginJSON contains the JSON metadata for the struct [PoolOrigin]
type poolOriginJSON struct {
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

func (r *PoolOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolOriginJSON) RawJSON() string {
	return r.raw
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type PoolOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string              `json:"Host"`
	JSON poolOriginsHeaderJSON `json:"-"`
}

// poolOriginsHeaderJSON contains the JSON metadata for the struct
// [PoolOriginsHeader]
type poolOriginsHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolOriginsHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolOriginsHeaderJSON) RawJSON() string {
	return r.raw
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

func (r loadBalancerPoolDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A list of regions from which to run health checks. Null means every Cloudflare
// data center.
//
// Union satisfied by [user.LoadBalancerPoolHealthResponseUnknown] or
// [shared.UnionString].
type LoadBalancerPoolHealthResponseUnion interface {
	ImplementsUserLoadBalancerPoolHealthResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LoadBalancerPoolHealthResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type LoadBalancerPoolPreviewResponse struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string                   `json:"pools"`
	PreviewID string                              `json:"preview_id"`
	JSON      loadBalancerPoolPreviewResponseJSON `json:"-"`
}

// loadBalancerPoolPreviewResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolPreviewResponse]
type loadBalancerPoolPreviewResponseJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolPreviewResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerPoolReferencesResponse struct {
	ReferenceType shared.UnnamedSchemaRef146             `json:"reference_type"`
	ResourceID    string                                 `json:"resource_id"`
	ResourceName  string                                 `json:"resource_name"`
	ResourceType  string                                 `json:"resource_type"`
	JSON          loadBalancerPoolReferencesResponseJSON `json:"-"`
}

// loadBalancerPoolReferencesResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPoolReferencesResponse]
type loadBalancerPoolReferencesResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerPoolReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolReferencesResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerPoolNewParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]LoadBalancerPoolNewParamsOrigin] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]LoadBalancerPoolNewParamsCheckRegion] `json:"check_regions"`
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

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolNewParamsCheckRegion string

const (
	LoadBalancerPoolNewParamsCheckRegionWnam       LoadBalancerPoolNewParamsCheckRegion = "WNAM"
	LoadBalancerPoolNewParamsCheckRegionEnam       LoadBalancerPoolNewParamsCheckRegion = "ENAM"
	LoadBalancerPoolNewParamsCheckRegionWeu        LoadBalancerPoolNewParamsCheckRegion = "WEU"
	LoadBalancerPoolNewParamsCheckRegionEeu        LoadBalancerPoolNewParamsCheckRegion = "EEU"
	LoadBalancerPoolNewParamsCheckRegionNsam       LoadBalancerPoolNewParamsCheckRegion = "NSAM"
	LoadBalancerPoolNewParamsCheckRegionSsam       LoadBalancerPoolNewParamsCheckRegion = "SSAM"
	LoadBalancerPoolNewParamsCheckRegionOc         LoadBalancerPoolNewParamsCheckRegion = "OC"
	LoadBalancerPoolNewParamsCheckRegionMe         LoadBalancerPoolNewParamsCheckRegion = "ME"
	LoadBalancerPoolNewParamsCheckRegionNaf        LoadBalancerPoolNewParamsCheckRegion = "NAF"
	LoadBalancerPoolNewParamsCheckRegionSaf        LoadBalancerPoolNewParamsCheckRegion = "SAF"
	LoadBalancerPoolNewParamsCheckRegionSas        LoadBalancerPoolNewParamsCheckRegion = "SAS"
	LoadBalancerPoolNewParamsCheckRegionSeas       LoadBalancerPoolNewParamsCheckRegion = "SEAS"
	LoadBalancerPoolNewParamsCheckRegionNeas       LoadBalancerPoolNewParamsCheckRegion = "NEAS"
	LoadBalancerPoolNewParamsCheckRegionAllRegions LoadBalancerPoolNewParamsCheckRegion = "ALL_REGIONS"
)

func (r LoadBalancerPoolNewParamsCheckRegion) IsKnown() bool {
	switch r {
	case LoadBalancerPoolNewParamsCheckRegionWnam, LoadBalancerPoolNewParamsCheckRegionEnam, LoadBalancerPoolNewParamsCheckRegionWeu, LoadBalancerPoolNewParamsCheckRegionEeu, LoadBalancerPoolNewParamsCheckRegionNsam, LoadBalancerPoolNewParamsCheckRegionSsam, LoadBalancerPoolNewParamsCheckRegionOc, LoadBalancerPoolNewParamsCheckRegionMe, LoadBalancerPoolNewParamsCheckRegionNaf, LoadBalancerPoolNewParamsCheckRegionSaf, LoadBalancerPoolNewParamsCheckRegionSas, LoadBalancerPoolNewParamsCheckRegionSeas, LoadBalancerPoolNewParamsCheckRegionNeas, LoadBalancerPoolNewParamsCheckRegionAllRegions:
		return true
	}
	return false
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

func (r LoadBalancerPoolNewParamsLoadSheddingDefaultPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerPoolNewParamsLoadSheddingDefaultPolicyRandom, LoadBalancerPoolNewParamsLoadSheddingDefaultPolicyHash:
		return true
	}
	return false
}

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolNewParamsLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolNewParamsLoadSheddingSessionPolicyHash LoadBalancerPoolNewParamsLoadSheddingSessionPolicy = "hash"
)

func (r LoadBalancerPoolNewParamsLoadSheddingSessionPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerPoolNewParamsLoadSheddingSessionPolicyHash:
		return true
	}
	return false
}

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

func (r LoadBalancerPoolNewParamsOriginSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerPoolNewParamsOriginSteeringPolicyRandom, LoadBalancerPoolNewParamsOriginSteeringPolicyHash, LoadBalancerPoolNewParamsOriginSteeringPolicyLeastOutstandingRequests, LoadBalancerPoolNewParamsOriginSteeringPolicyLeastConnections:
		return true
	}
	return false
}

type LoadBalancerPoolNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Pool                  `json:"result,required"`
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

func (r loadBalancerPoolNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolNewResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolNewResponseEnvelopeSuccessTrue LoadBalancerPoolNewResponseEnvelopeSuccess = true
)

func (r LoadBalancerPoolNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerPoolNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r LoadBalancerPoolUpdateParamsCheckRegion) IsKnown() bool {
	switch r {
	case LoadBalancerPoolUpdateParamsCheckRegionWnam, LoadBalancerPoolUpdateParamsCheckRegionEnam, LoadBalancerPoolUpdateParamsCheckRegionWeu, LoadBalancerPoolUpdateParamsCheckRegionEeu, LoadBalancerPoolUpdateParamsCheckRegionNsam, LoadBalancerPoolUpdateParamsCheckRegionSsam, LoadBalancerPoolUpdateParamsCheckRegionOc, LoadBalancerPoolUpdateParamsCheckRegionMe, LoadBalancerPoolUpdateParamsCheckRegionNaf, LoadBalancerPoolUpdateParamsCheckRegionSaf, LoadBalancerPoolUpdateParamsCheckRegionSas, LoadBalancerPoolUpdateParamsCheckRegionSeas, LoadBalancerPoolUpdateParamsCheckRegionNeas, LoadBalancerPoolUpdateParamsCheckRegionAllRegions:
		return true
	}
	return false
}

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

func (r LoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyRandom, LoadBalancerPoolUpdateParamsLoadSheddingDefaultPolicyHash:
		return true
	}
	return false
}

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolUpdateParamsLoadSheddingSessionPolicyHash LoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy = "hash"
)

func (r LoadBalancerPoolUpdateParamsLoadSheddingSessionPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerPoolUpdateParamsLoadSheddingSessionPolicyHash:
		return true
	}
	return false
}

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

func (r LoadBalancerPoolUpdateParamsOriginSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerPoolUpdateParamsOriginSteeringPolicyRandom, LoadBalancerPoolUpdateParamsOriginSteeringPolicyHash, LoadBalancerPoolUpdateParamsOriginSteeringPolicyLeastOutstandingRequests, LoadBalancerPoolUpdateParamsOriginSteeringPolicyLeastConnections:
		return true
	}
	return false
}

type LoadBalancerPoolUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Pool                  `json:"result,required"`
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

func (r loadBalancerPoolUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolUpdateResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolUpdateResponseEnvelopeSuccessTrue LoadBalancerPoolUpdateResponseEnvelopeSuccess = true
)

func (r LoadBalancerPoolUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerPoolUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerPoolListParams struct {
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[interface{}] `query:"monitor"`
}

// URLQuery serializes [LoadBalancerPoolListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerPoolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LoadBalancerPoolDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r LoadBalancerPoolDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LoadBalancerPoolDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   LoadBalancerPoolDeleteResponse `json:"result,required"`
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

func (r loadBalancerPoolDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolDeleteResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolDeleteResponseEnvelopeSuccessTrue LoadBalancerPoolDeleteResponseEnvelopeSuccess = true
)

func (r LoadBalancerPoolDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerPoolDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerPoolEditParams struct {
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]LoadBalancerPoolEditParamsCheckRegion] `json:"check_regions"`
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
	LoadShedding param.Field[LoadBalancerPoolEditParamsLoadShedding] `json:"load_shedding"`
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
	NotificationFilter param.Field[LoadBalancerPoolEditParamsNotificationFilter] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[LoadBalancerPoolEditParamsOriginSteering] `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]LoadBalancerPoolEditParamsOrigin] `json:"origins"`
}

func (r LoadBalancerPoolEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type LoadBalancerPoolEditParamsCheckRegion string

const (
	LoadBalancerPoolEditParamsCheckRegionWnam       LoadBalancerPoolEditParamsCheckRegion = "WNAM"
	LoadBalancerPoolEditParamsCheckRegionEnam       LoadBalancerPoolEditParamsCheckRegion = "ENAM"
	LoadBalancerPoolEditParamsCheckRegionWeu        LoadBalancerPoolEditParamsCheckRegion = "WEU"
	LoadBalancerPoolEditParamsCheckRegionEeu        LoadBalancerPoolEditParamsCheckRegion = "EEU"
	LoadBalancerPoolEditParamsCheckRegionNsam       LoadBalancerPoolEditParamsCheckRegion = "NSAM"
	LoadBalancerPoolEditParamsCheckRegionSsam       LoadBalancerPoolEditParamsCheckRegion = "SSAM"
	LoadBalancerPoolEditParamsCheckRegionOc         LoadBalancerPoolEditParamsCheckRegion = "OC"
	LoadBalancerPoolEditParamsCheckRegionMe         LoadBalancerPoolEditParamsCheckRegion = "ME"
	LoadBalancerPoolEditParamsCheckRegionNaf        LoadBalancerPoolEditParamsCheckRegion = "NAF"
	LoadBalancerPoolEditParamsCheckRegionSaf        LoadBalancerPoolEditParamsCheckRegion = "SAF"
	LoadBalancerPoolEditParamsCheckRegionSas        LoadBalancerPoolEditParamsCheckRegion = "SAS"
	LoadBalancerPoolEditParamsCheckRegionSeas       LoadBalancerPoolEditParamsCheckRegion = "SEAS"
	LoadBalancerPoolEditParamsCheckRegionNeas       LoadBalancerPoolEditParamsCheckRegion = "NEAS"
	LoadBalancerPoolEditParamsCheckRegionAllRegions LoadBalancerPoolEditParamsCheckRegion = "ALL_REGIONS"
)

func (r LoadBalancerPoolEditParamsCheckRegion) IsKnown() bool {
	switch r {
	case LoadBalancerPoolEditParamsCheckRegionWnam, LoadBalancerPoolEditParamsCheckRegionEnam, LoadBalancerPoolEditParamsCheckRegionWeu, LoadBalancerPoolEditParamsCheckRegionEeu, LoadBalancerPoolEditParamsCheckRegionNsam, LoadBalancerPoolEditParamsCheckRegionSsam, LoadBalancerPoolEditParamsCheckRegionOc, LoadBalancerPoolEditParamsCheckRegionMe, LoadBalancerPoolEditParamsCheckRegionNaf, LoadBalancerPoolEditParamsCheckRegionSaf, LoadBalancerPoolEditParamsCheckRegionSas, LoadBalancerPoolEditParamsCheckRegionSeas, LoadBalancerPoolEditParamsCheckRegionNeas, LoadBalancerPoolEditParamsCheckRegionAllRegions:
		return true
	}
	return false
}

// Configures load shedding policies and percentages for the pool.
type LoadBalancerPoolEditParamsLoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[LoadBalancerPoolEditParamsLoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[LoadBalancerPoolEditParamsLoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r LoadBalancerPoolEditParamsLoadShedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadBalancerPoolEditParamsLoadSheddingDefaultPolicy string

const (
	LoadBalancerPoolEditParamsLoadSheddingDefaultPolicyRandom LoadBalancerPoolEditParamsLoadSheddingDefaultPolicy = "random"
	LoadBalancerPoolEditParamsLoadSheddingDefaultPolicyHash   LoadBalancerPoolEditParamsLoadSheddingDefaultPolicy = "hash"
)

func (r LoadBalancerPoolEditParamsLoadSheddingDefaultPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerPoolEditParamsLoadSheddingDefaultPolicyRandom, LoadBalancerPoolEditParamsLoadSheddingDefaultPolicyHash:
		return true
	}
	return false
}

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadBalancerPoolEditParamsLoadSheddingSessionPolicy string

const (
	LoadBalancerPoolEditParamsLoadSheddingSessionPolicyHash LoadBalancerPoolEditParamsLoadSheddingSessionPolicy = "hash"
)

func (r LoadBalancerPoolEditParamsLoadSheddingSessionPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerPoolEditParamsLoadSheddingSessionPolicyHash:
		return true
	}
	return false
}

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type LoadBalancerPoolEditParamsNotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[LoadBalancerPoolEditParamsNotificationFilterOrigin] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[LoadBalancerPoolEditParamsNotificationFilterPool] `json:"pool"`
}

func (r LoadBalancerPoolEditParamsNotificationFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolEditParamsNotificationFilterOrigin struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r LoadBalancerPoolEditParamsNotificationFilterOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type LoadBalancerPoolEditParamsNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r LoadBalancerPoolEditParamsNotificationFilterPool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type LoadBalancerPoolEditParamsOriginSteering struct {
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
	Policy param.Field[LoadBalancerPoolEditParamsOriginSteeringPolicy] `json:"policy"`
}

func (r LoadBalancerPoolEditParamsOriginSteering) MarshalJSON() (data []byte, err error) {
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
type LoadBalancerPoolEditParamsOriginSteeringPolicy string

const (
	LoadBalancerPoolEditParamsOriginSteeringPolicyRandom                   LoadBalancerPoolEditParamsOriginSteeringPolicy = "random"
	LoadBalancerPoolEditParamsOriginSteeringPolicyHash                     LoadBalancerPoolEditParamsOriginSteeringPolicy = "hash"
	LoadBalancerPoolEditParamsOriginSteeringPolicyLeastOutstandingRequests LoadBalancerPoolEditParamsOriginSteeringPolicy = "least_outstanding_requests"
	LoadBalancerPoolEditParamsOriginSteeringPolicyLeastConnections         LoadBalancerPoolEditParamsOriginSteeringPolicy = "least_connections"
)

func (r LoadBalancerPoolEditParamsOriginSteeringPolicy) IsKnown() bool {
	switch r {
	case LoadBalancerPoolEditParamsOriginSteeringPolicyRandom, LoadBalancerPoolEditParamsOriginSteeringPolicyHash, LoadBalancerPoolEditParamsOriginSteeringPolicyLeastOutstandingRequests, LoadBalancerPoolEditParamsOriginSteeringPolicyLeastConnections:
		return true
	}
	return false
}

type LoadBalancerPoolEditParamsOrigin struct {
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
	Header param.Field[LoadBalancerPoolEditParamsOriginsHeader] `json:"header"`
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

func (r LoadBalancerPoolEditParamsOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type LoadBalancerPoolEditParamsOriginsHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r LoadBalancerPoolEditParamsOriginsHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerPoolEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Pool                  `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPoolEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPoolEditResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPoolEditResponseEnvelope]
type loadBalancerPoolEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolEditResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolEditResponseEnvelopeSuccessTrue LoadBalancerPoolEditResponseEnvelopeSuccess = true
)

func (r LoadBalancerPoolEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerPoolEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerPoolGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Pool                  `json:"result,required"`
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

func (r loadBalancerPoolGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolGetResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolGetResponseEnvelopeSuccessTrue LoadBalancerPoolGetResponseEnvelopeSuccess = true
)

func (r LoadBalancerPoolGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerPoolGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerPoolHealthResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result LoadBalancerPoolHealthResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPoolHealthResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPoolHealthResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolHealthResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPoolHealthResponseEnvelope]
type loadBalancerPoolHealthResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolHealthResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolHealthResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolHealthResponseEnvelopeSuccessTrue LoadBalancerPoolHealthResponseEnvelopeSuccess = true
)

func (r LoadBalancerPoolHealthResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerPoolHealthResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerPoolPreviewParams struct {
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
	Type param.Field[LoadBalancerPoolPreviewParamsType] `json:"type"`
}

func (r LoadBalancerPoolPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type LoadBalancerPoolPreviewParamsType string

const (
	LoadBalancerPoolPreviewParamsTypeHTTP     LoadBalancerPoolPreviewParamsType = "http"
	LoadBalancerPoolPreviewParamsTypeHTTPS    LoadBalancerPoolPreviewParamsType = "https"
	LoadBalancerPoolPreviewParamsTypeTcp      LoadBalancerPoolPreviewParamsType = "tcp"
	LoadBalancerPoolPreviewParamsTypeUdpIcmp  LoadBalancerPoolPreviewParamsType = "udp_icmp"
	LoadBalancerPoolPreviewParamsTypeIcmpPing LoadBalancerPoolPreviewParamsType = "icmp_ping"
	LoadBalancerPoolPreviewParamsTypeSmtp     LoadBalancerPoolPreviewParamsType = "smtp"
)

func (r LoadBalancerPoolPreviewParamsType) IsKnown() bool {
	switch r {
	case LoadBalancerPoolPreviewParamsTypeHTTP, LoadBalancerPoolPreviewParamsTypeHTTPS, LoadBalancerPoolPreviewParamsTypeTcp, LoadBalancerPoolPreviewParamsTypeUdpIcmp, LoadBalancerPoolPreviewParamsTypeIcmpPing, LoadBalancerPoolPreviewParamsTypeSmtp:
		return true
	}
	return false
}

type LoadBalancerPoolPreviewResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   LoadBalancerPoolPreviewResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPoolPreviewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPoolPreviewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolPreviewResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPoolPreviewResponseEnvelope]
type loadBalancerPoolPreviewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolPreviewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolPreviewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolPreviewResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolPreviewResponseEnvelopeSuccessTrue LoadBalancerPoolPreviewResponseEnvelopeSuccess = true
)

func (r LoadBalancerPoolPreviewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerPoolPreviewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerPoolReferencesResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// List of resources that reference a given pool.
	Result []LoadBalancerPoolReferencesResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerPoolReferencesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerPoolReferencesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerPoolReferencesResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerPoolReferencesResponseEnvelopeJSON contains the JSON metadata for
// the struct [LoadBalancerPoolReferencesResponseEnvelope]
type loadBalancerPoolReferencesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferencesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolReferencesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolReferencesResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolReferencesResponseEnvelopeSuccessTrue LoadBalancerPoolReferencesResponseEnvelopeSuccess = true
)

func (r LoadBalancerPoolReferencesResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerPoolReferencesResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerPoolReferencesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       loadBalancerPoolReferencesResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerPoolReferencesResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [LoadBalancerPoolReferencesResponseEnvelopeResultInfo]
type loadBalancerPoolReferencesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferencesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolReferencesResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
