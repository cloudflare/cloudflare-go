// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

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
func (r *UserLoadBalancerPoolService) Get(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *PoolSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured pool.
func (r *UserLoadBalancerPoolService) Update(ctx context.Context, identifier interface{}, body UserLoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *PoolSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a configured pool.
func (r *UserLoadBalancerPoolService) Delete(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *SchemasIDResponseLwyhetfd, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new pool.
func (r *UserLoadBalancerPoolService) LoadBalancerPoolsNewPool(ctx context.Context, body UserLoadBalancerPoolLoadBalancerPoolsNewPoolParams, opts ...option.RequestOption) (res *PoolSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured pools.
func (r *UserLoadBalancerPoolService) LoadBalancerPoolsListPools(ctx context.Context, query UserLoadBalancerPoolLoadBalancerPoolsListPoolsParams, opts ...option.RequestOption) (res *PoolResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Apply changes to a number of existing pools, overwriting the supplied
// properties. Pools are ordered by ascending `name`. Returns the list of affected
// pools. Supports the standard pagination query parameters, either
// `limit`/`offset` or `per_page`/`page`.
func (r *UserLoadBalancerPoolService) LoadBalancerPoolsPatchPools(ctx context.Context, body UserLoadBalancerPoolLoadBalancerPoolsPatchPoolsParams, opts ...option.RequestOption) (res *PoolResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancers/pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

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
	// The email address to send health status notifications to. This can be an
	// individual mailbox or a mailing list. Multiple emails can be supplied as a comma
	// delimited list.
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
	// The type of origin steering policy to use, either "random" or "hash" (based on
	// CF-Connecting-IP).
	Policy param.Field[UserLoadBalancerPoolUpdateParamsOriginSteeringPolicy] `json:"policy"`
}

func (r UserLoadBalancerPoolUpdateParamsOriginSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of origin steering policy to use, either "random" or "hash" (based on
// CF-Connecting-IP).
type UserLoadBalancerPoolUpdateParamsOriginSteeringPolicy string

const (
	UserLoadBalancerPoolUpdateParamsOriginSteeringPolicyRandom UserLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "random"
	UserLoadBalancerPoolUpdateParamsOriginSteeringPolicyHash   UserLoadBalancerPoolUpdateParamsOriginSteeringPolicy = "hash"
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
	// The email address to send health status notifications to. This can be an
	// individual mailbox or a mailing list. Multiple emails can be supplied as a comma
	// delimited list.
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
	// The type of origin steering policy to use, either "random" or "hash" (based on
	// CF-Connecting-IP).
	Policy param.Field[UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy] `json:"policy"`
}

func (r UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of origin steering policy to use, either "random" or "hash" (based on
// CF-Connecting-IP).
type UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy string

const (
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyRandom UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "random"
	UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicyHash   UserLoadBalancerPoolLoadBalancerPoolsNewPoolParamsOriginSteeringPolicy = "hash"
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
