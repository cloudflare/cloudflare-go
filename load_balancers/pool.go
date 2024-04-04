// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/user"
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
func (r *PoolService) New(ctx context.Context, params PoolNewParams, opts ...option.RequestOption) (res *user.Pool, err error) {
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
func (r *PoolService) Update(ctx context.Context, poolID string, params PoolUpdateParams, opts ...option.RequestOption) (res *user.Pool, err error) {
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
func (r *PoolService) List(ctx context.Context, params PoolListParams, opts ...option.RequestOption) (res *pagination.SinglePage[user.Pool], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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
func (r *PoolService) ListAutoPaging(ctx context.Context, params PoolListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[user.Pool] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete a configured pool.
func (r *PoolService) Delete(ctx context.Context, poolID string, params PoolDeleteParams, opts ...option.RequestOption) (res *PoolDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", params.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing pool, overwriting the supplied properties.
func (r *PoolService) Edit(ctx context.Context, poolID string, params PoolEditParams, opts ...option.RequestOption) (res *user.Pool, err error) {
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
func (r *PoolService) Get(ctx context.Context, poolID string, query PoolGetParams, opts ...option.RequestOption) (res *user.Pool, err error) {
	opts = append(r.Options[:], opts...)
	var env PoolGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", query.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

func (r PoolNewParamsLoadSheddingDefaultPolicy) IsKnown() bool {
	switch r {
	case PoolNewParamsLoadSheddingDefaultPolicyRandom, PoolNewParamsLoadSheddingDefaultPolicyHash:
		return true
	}
	return false
}

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolNewParamsLoadSheddingSessionPolicy string

const (
	PoolNewParamsLoadSheddingSessionPolicyHash PoolNewParamsLoadSheddingSessionPolicy = "hash"
)

func (r PoolNewParamsLoadSheddingSessionPolicy) IsKnown() bool {
	switch r {
	case PoolNewParamsLoadSheddingSessionPolicyHash:
		return true
	}
	return false
}

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

func (r PoolNewParamsOriginSteeringPolicy) IsKnown() bool {
	switch r {
	case PoolNewParamsOriginSteeringPolicyRandom, PoolNewParamsOriginSteeringPolicyHash, PoolNewParamsOriginSteeringPolicyLeastOutstandingRequests, PoolNewParamsOriginSteeringPolicyLeastConnections:
		return true
	}
	return false
}

type PoolNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   user.Pool             `json:"result,required"`
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

// Whether the API call was successful
type PoolNewResponseEnvelopeSuccess bool

const (
	PoolNewResponseEnvelopeSuccessTrue PoolNewResponseEnvelopeSuccess = true
)

func (r PoolNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r PoolUpdateParamsCheckRegion) IsKnown() bool {
	switch r {
	case PoolUpdateParamsCheckRegionWnam, PoolUpdateParamsCheckRegionEnam, PoolUpdateParamsCheckRegionWeu, PoolUpdateParamsCheckRegionEeu, PoolUpdateParamsCheckRegionNsam, PoolUpdateParamsCheckRegionSsam, PoolUpdateParamsCheckRegionOc, PoolUpdateParamsCheckRegionMe, PoolUpdateParamsCheckRegionNaf, PoolUpdateParamsCheckRegionSaf, PoolUpdateParamsCheckRegionSas, PoolUpdateParamsCheckRegionSeas, PoolUpdateParamsCheckRegionNeas, PoolUpdateParamsCheckRegionAllRegions:
		return true
	}
	return false
}

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

func (r PoolUpdateParamsLoadSheddingDefaultPolicy) IsKnown() bool {
	switch r {
	case PoolUpdateParamsLoadSheddingDefaultPolicyRandom, PoolUpdateParamsLoadSheddingDefaultPolicyHash:
		return true
	}
	return false
}

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolUpdateParamsLoadSheddingSessionPolicy string

const (
	PoolUpdateParamsLoadSheddingSessionPolicyHash PoolUpdateParamsLoadSheddingSessionPolicy = "hash"
)

func (r PoolUpdateParamsLoadSheddingSessionPolicy) IsKnown() bool {
	switch r {
	case PoolUpdateParamsLoadSheddingSessionPolicyHash:
		return true
	}
	return false
}

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

func (r PoolUpdateParamsOriginSteeringPolicy) IsKnown() bool {
	switch r {
	case PoolUpdateParamsOriginSteeringPolicyRandom, PoolUpdateParamsOriginSteeringPolicyHash, PoolUpdateParamsOriginSteeringPolicyLeastOutstandingRequests, PoolUpdateParamsOriginSteeringPolicyLeastConnections:
		return true
	}
	return false
}

type PoolUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   user.Pool             `json:"result,required"`
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

// Whether the API call was successful
type PoolUpdateResponseEnvelopeSuccess bool

const (
	PoolUpdateResponseEnvelopeSuccessTrue PoolUpdateResponseEnvelopeSuccess = true
)

func (r PoolUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PoolDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r PoolDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PoolDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   PoolDeleteResponse    `json:"result,required"`
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

// Whether the API call was successful
type PoolDeleteResponseEnvelopeSuccess bool

const (
	PoolDeleteResponseEnvelopeSuccessTrue PoolDeleteResponseEnvelopeSuccess = true
)

func (r PoolDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r PoolEditParamsCheckRegion) IsKnown() bool {
	switch r {
	case PoolEditParamsCheckRegionWnam, PoolEditParamsCheckRegionEnam, PoolEditParamsCheckRegionWeu, PoolEditParamsCheckRegionEeu, PoolEditParamsCheckRegionNsam, PoolEditParamsCheckRegionSsam, PoolEditParamsCheckRegionOc, PoolEditParamsCheckRegionMe, PoolEditParamsCheckRegionNaf, PoolEditParamsCheckRegionSaf, PoolEditParamsCheckRegionSas, PoolEditParamsCheckRegionSeas, PoolEditParamsCheckRegionNeas, PoolEditParamsCheckRegionAllRegions:
		return true
	}
	return false
}

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

func (r PoolEditParamsLoadSheddingDefaultPolicy) IsKnown() bool {
	switch r {
	case PoolEditParamsLoadSheddingDefaultPolicyRandom, PoolEditParamsLoadSheddingDefaultPolicyHash:
		return true
	}
	return false
}

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type PoolEditParamsLoadSheddingSessionPolicy string

const (
	PoolEditParamsLoadSheddingSessionPolicyHash PoolEditParamsLoadSheddingSessionPolicy = "hash"
)

func (r PoolEditParamsLoadSheddingSessionPolicy) IsKnown() bool {
	switch r {
	case PoolEditParamsLoadSheddingSessionPolicyHash:
		return true
	}
	return false
}

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

func (r PoolEditParamsOriginSteeringPolicy) IsKnown() bool {
	switch r {
	case PoolEditParamsOriginSteeringPolicyRandom, PoolEditParamsOriginSteeringPolicyHash, PoolEditParamsOriginSteeringPolicyLeastOutstandingRequests, PoolEditParamsOriginSteeringPolicyLeastConnections:
		return true
	}
	return false
}

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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   user.Pool             `json:"result,required"`
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

// Whether the API call was successful
type PoolEditResponseEnvelopeSuccess bool

const (
	PoolEditResponseEnvelopeSuccessTrue PoolEditResponseEnvelopeSuccess = true
)

func (r PoolEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PoolGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PoolGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   user.Pool             `json:"result,required"`
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

// Whether the API call was successful
type PoolGetResponseEnvelopeSuccess bool

const (
	PoolGetResponseEnvelopeSuccessTrue PoolGetResponseEnvelopeSuccess = true
)

func (r PoolGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
