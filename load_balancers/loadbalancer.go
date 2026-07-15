// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// LoadBalancerService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerService] method instead.
type LoadBalancerService struct {
	Options       []option.RequestOption
	Monitors      *MonitorService
	MonitorGroups *MonitorGroupService
	Pools         *PoolService
	Previews      *PreviewService
	Regions       *RegionService
	Searches      *SearchService
}

// NewLoadBalancerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLoadBalancerService(opts ...option.RequestOption) (r *LoadBalancerService) {
	r = &LoadBalancerService{}
	r.Options = opts
	r.Monitors = NewMonitorService(opts...)
	r.MonitorGroups = NewMonitorGroupService(opts...)
	r.Pools = NewPoolService(opts...)
	r.Previews = NewPreviewService(opts...)
	r.Regions = NewRegionService(opts...)
	r.Searches = NewSearchService(opts...)
	return
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type CheckRegion string

const (
	CheckRegionWnam       CheckRegion = "WNAM"
	CheckRegionEnam       CheckRegion = "ENAM"
	CheckRegionWeu        CheckRegion = "WEU"
	CheckRegionEeu        CheckRegion = "EEU"
	CheckRegionNsam       CheckRegion = "NSAM"
	CheckRegionSsam       CheckRegion = "SSAM"
	CheckRegionOc         CheckRegion = "OC"
	CheckRegionMe         CheckRegion = "ME"
	CheckRegionNaf        CheckRegion = "NAF"
	CheckRegionSaf        CheckRegion = "SAF"
	CheckRegionSas        CheckRegion = "SAS"
	CheckRegionSeas       CheckRegion = "SEAS"
	CheckRegionNeas       CheckRegion = "NEAS"
	CheckRegionAllRegions CheckRegion = "ALL_REGIONS"
)

func (r CheckRegion) IsKnown() bool {
	switch r {
	case CheckRegionWnam, CheckRegionEnam, CheckRegionWeu, CheckRegionEeu, CheckRegionNsam, CheckRegionSsam, CheckRegionOc, CheckRegionMe, CheckRegionNaf, CheckRegionSaf, CheckRegionSas, CheckRegionSeas, CheckRegionNeas, CheckRegionAllRegions:
		return true
	}
	return false
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type FilterOptions struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable" api:"nullable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool              `json:"healthy" api:"nullable"`
	JSON    filterOptionsJSON `json:"-"`
}

// filterOptionsJSON contains the JSON metadata for the struct [FilterOptions]
type filterOptionsJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterOptionsJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type FilterOptionsParam struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r FilterOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type Header struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []Host     `json:"Host"`
	JSON headerJSON `json:"-"`
}

// headerJSON contains the JSON metadata for the struct [Header]
type headerJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Header) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r headerJSON) RawJSON() string {
	return r.raw
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type HeaderParam struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]HostParam] `json:"Host"`
}

func (r HeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Host = string

type HostParam = string

// Configures load shedding policies and percentages for the pool.
type LoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadSheddingJSON          `json:"-"`
}

// loadSheddingJSON contains the JSON metadata for the struct [LoadShedding]
type loadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadSheddingJSON) RawJSON() string {
	return r.raw
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadSheddingDefaultPolicy string

const (
	LoadSheddingDefaultPolicyRandom LoadSheddingDefaultPolicy = "random"
	LoadSheddingDefaultPolicyHash   LoadSheddingDefaultPolicy = "hash"
)

func (r LoadSheddingDefaultPolicy) IsKnown() bool {
	switch r {
	case LoadSheddingDefaultPolicyRandom, LoadSheddingDefaultPolicyHash:
		return true
	}
	return false
}

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadSheddingSessionPolicy string

const (
	LoadSheddingSessionPolicyHash LoadSheddingSessionPolicy = "hash"
)

func (r LoadSheddingSessionPolicy) IsKnown() bool {
	switch r {
	case LoadSheddingSessionPolicyHash:
		return true
	}
	return false
}

// Configures load shedding policies and percentages for the pool.
type LoadSheddingParam struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[LoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[LoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r LoadSheddingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type NotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin FilterOptions `json:"origin" api:"nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool FilterOptions          `json:"pool" api:"nullable"`
	JSON notificationFilterJSON `json:"-"`
}

// notificationFilterJSON contains the JSON metadata for the struct
// [NotificationFilter]
type notificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationFilterJSON) RawJSON() string {
	return r.raw
}

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type NotificationFilterParam struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[FilterOptionsParam] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[FilterOptionsParam] `json:"pool"`
}

func (r NotificationFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Origin struct {
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
	// Whether to flatten CNAME records for this origin, resolving them to A/AAAA
	// records before returning to the client. When true (the default), the director
	// resolves CNAME addresses to their underlying A/AAAA records. When false, the
	// origin address is returned as a raw CNAME record without resolution. This
	// setting mirrors the DNS API record flatten_cname setting.
	FlattenCNAME bool `json:"flatten_cname"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header Header `json:"header"`
	// A human-identifiable name for the origin.
	Name string `json:"name"`
	// The port for upstream connections. A value of 0 means the default port for the
	// protocol will be used.
	Port int64 `json:"port"`
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
	Weight float64    `json:"weight"`
	JSON   originJSON `json:"-"`
}

// originJSON contains the JSON metadata for the struct [Origin]
type originJSON struct {
	Address          apijson.Field
	DisabledAt       apijson.Field
	Enabled          apijson.Field
	FlattenCNAME     apijson.Field
	Header           apijson.Field
	Name             apijson.Field
	Port             apijson.Field
	VirtualNetworkID apijson.Field
	Weight           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Origin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originJSON) RawJSON() string {
	return r.raw
}

type OriginParam struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address param.Field[string] `json:"address"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled param.Field[bool] `json:"enabled"`
	// Whether to flatten CNAME records for this origin, resolving them to A/AAAA
	// records before returning to the client. When true (the default), the director
	// resolves CNAME addresses to their underlying A/AAAA records. When false, the
	// origin address is returned as a raw CNAME record without resolution. This
	// setting mirrors the DNS API record flatten_cname setting.
	FlattenCNAME param.Field[bool] `json:"flatten_cname"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header param.Field[HeaderParam] `json:"header"`
	// A human-identifiable name for the origin.
	Name param.Field[string] `json:"name"`
	// The port for upstream connections. A value of 0 means the default port for the
	// protocol will be used.
	Port param.Field[int64] `json:"port"`
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

func (r OriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type OriginSteering struct {
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
	Policy OriginSteeringPolicy `json:"policy"`
	JSON   originSteeringJSON   `json:"-"`
}

// originSteeringJSON contains the JSON metadata for the struct [OriginSteering]
type originSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originSteeringJSON) RawJSON() string {
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
type OriginSteeringPolicy string

const (
	OriginSteeringPolicyRandom                   OriginSteeringPolicy = "random"
	OriginSteeringPolicyHash                     OriginSteeringPolicy = "hash"
	OriginSteeringPolicyLeastOutstandingRequests OriginSteeringPolicy = "least_outstanding_requests"
	OriginSteeringPolicyLeastConnections         OriginSteeringPolicy = "least_connections"
)

func (r OriginSteeringPolicy) IsKnown() bool {
	switch r {
	case OriginSteeringPolicyRandom, OriginSteeringPolicyHash, OriginSteeringPolicyLeastOutstandingRequests, OriginSteeringPolicyLeastConnections:
		return true
	}
	return false
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type OriginSteeringParam struct {
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
	Policy param.Field[OriginSteeringPolicy] `json:"policy"`
}

func (r OriginSteeringParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
