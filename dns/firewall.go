// File generated from our OpenAPI spec by Stainless.

package dns

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// FirewallService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewFirewallService] method instead.
type FirewallService struct {
	Options   []option.RequestOption
	Analytics *FirewallAnalyticsService
}

// NewFirewallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallService(opts ...option.RequestOption) (r *FirewallService) {
	r = &FirewallService{}
	r.Options = opts
	r.Analytics = NewFirewallAnalyticsService(opts...)
	return
}

// Create a configured DNS Firewall Cluster.
func (r *FirewallService) New(ctx context.Context, params FirewallNewParams, opts ...option.RequestOption) (res *FirewallNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured DNS Firewall clusters for an account.
func (r *FirewallService) List(ctx context.Context, params FirewallListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FirewallListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall", params.AccountID)
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

// List configured DNS Firewall clusters for an account.
func (r *FirewallService) ListAutoPaging(ctx context.Context, params FirewallListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FirewallListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a configured DNS Firewall Cluster.
func (r *FirewallService) Delete(ctx context.Context, dnsFirewallID string, body FirewallDeleteParams, opts ...option.RequestOption) (res *FirewallDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", body.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a DNS Firewall Cluster configuration.
func (r *FirewallService) Edit(ctx context.Context, dnsFirewallID string, params FirewallEditParams, opts ...option.RequestOption) (res *FirewallEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", params.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show a single configured DNS Firewall cluster for an account.
func (r *FirewallService) Get(ctx context.Context, dnsFirewallID string, query FirewallGetParams, opts ...option.RequestOption) (res *FirewallGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", query.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                               `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallNewResponseDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                          `json:"name,required"`
	UpstreamIPs []FirewallNewResponseUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation FirewallNewResponseAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                 `json:"retries"`
	JSON    firewallNewResponseJSON `json:"-"`
}

// firewallNewResponseJSON contains the JSON metadata for the struct
// [FirewallNewResponse]
type firewallNewResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	NegativeCacheTTL     apijson.Field
	OriginIPs            apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FirewallNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallNewResponseJSON) RawJSON() string {
	return r.raw
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type FirewallNewResponseDNSFirewallIP interface {
	ImplementsDNSFirewallNewResponseDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallNewResponseDNSFirewallIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Upstream DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type FirewallNewResponseUpstreamIP interface {
	ImplementsDNSFirewallNewResponseUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallNewResponseUpstreamIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Attack mitigation settings.
type FirewallNewResponseAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                    `json:"only_when_upstream_unhealthy"`
	JSON                      firewallNewResponseAttackMitigationJSON `json:"-"`
}

// firewallNewResponseAttackMitigationJSON contains the JSON metadata for the
// struct [FirewallNewResponseAttackMitigation]
type firewallNewResponseAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *FirewallNewResponseAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallNewResponseAttackMitigationJSON) RawJSON() string {
	return r.raw
}

type FirewallListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallListResponseDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                           `json:"name,required"`
	UpstreamIPs []FirewallListResponseUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation FirewallListResponseAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                  `json:"retries"`
	JSON    firewallListResponseJSON `json:"-"`
}

// firewallListResponseJSON contains the JSON metadata for the struct
// [FirewallListResponse]
type firewallListResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	NegativeCacheTTL     apijson.Field
	OriginIPs            apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FirewallListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallListResponseJSON) RawJSON() string {
	return r.raw
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type FirewallListResponseDNSFirewallIP interface {
	ImplementsDNSFirewallListResponseDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallListResponseDNSFirewallIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Upstream DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type FirewallListResponseUpstreamIP interface {
	ImplementsDNSFirewallListResponseUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallListResponseUpstreamIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Attack mitigation settings.
type FirewallListResponseAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                     `json:"only_when_upstream_unhealthy"`
	JSON                      firewallListResponseAttackMitigationJSON `json:"-"`
}

// firewallListResponseAttackMitigationJSON contains the JSON metadata for the
// struct [FirewallListResponseAttackMitigation]
type firewallListResponseAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *FirewallListResponseAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallListResponseAttackMitigationJSON) RawJSON() string {
	return r.raw
}

type FirewallDeleteResponse struct {
	// Identifier
	ID   string                     `json:"id"`
	JSON firewallDeleteResponseJSON `json:"-"`
}

// firewallDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallDeleteResponse]
type firewallDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type FirewallEditResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallEditResponseDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                           `json:"name,required"`
	UpstreamIPs []FirewallEditResponseUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation FirewallEditResponseAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                  `json:"retries"`
	JSON    firewallEditResponseJSON `json:"-"`
}

// firewallEditResponseJSON contains the JSON metadata for the struct
// [FirewallEditResponse]
type firewallEditResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	NegativeCacheTTL     apijson.Field
	OriginIPs            apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FirewallEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallEditResponseJSON) RawJSON() string {
	return r.raw
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type FirewallEditResponseDNSFirewallIP interface {
	ImplementsDNSFirewallEditResponseDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallEditResponseDNSFirewallIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Upstream DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type FirewallEditResponseUpstreamIP interface {
	ImplementsDNSFirewallEditResponseUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallEditResponseUpstreamIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Attack mitigation settings.
type FirewallEditResponseAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                     `json:"only_when_upstream_unhealthy"`
	JSON                      firewallEditResponseAttackMitigationJSON `json:"-"`
}

// firewallEditResponseAttackMitigationJSON contains the JSON metadata for the
// struct [FirewallEditResponseAttackMitigation]
type firewallEditResponseAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *FirewallEditResponseAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallEditResponseAttackMitigationJSON) RawJSON() string {
	return r.raw
}

type FirewallGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                               `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallGetResponseDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                          `json:"name,required"`
	UpstreamIPs []FirewallGetResponseUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation FirewallGetResponseAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                 `json:"retries"`
	JSON    firewallGetResponseJSON `json:"-"`
}

// firewallGetResponseJSON contains the JSON metadata for the struct
// [FirewallGetResponse]
type firewallGetResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	NegativeCacheTTL     apijson.Field
	OriginIPs            apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FirewallGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallGetResponseJSON) RawJSON() string {
	return r.raw
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type FirewallGetResponseDNSFirewallIP interface {
	ImplementsDNSFirewallGetResponseDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallGetResponseDNSFirewallIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Upstream DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type FirewallGetResponseUpstreamIP interface {
	ImplementsDNSFirewallGetResponseUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallGetResponseUpstreamIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Attack mitigation settings.
type FirewallGetResponseAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                    `json:"only_when_upstream_unhealthy"`
	JSON                      firewallGetResponseAttackMitigationJSON `json:"-"`
}

// firewallGetResponseAttackMitigationJSON contains the JSON metadata for the
// struct [FirewallGetResponseAttackMitigation]
type firewallGetResponseAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *FirewallGetResponseAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallGetResponseAttackMitigationJSON) RawJSON() string {
	return r.raw
}

type FirewallNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// DNS Firewall Cluster Name.
	Name        param.Field[string]                        `json:"name,required"`
	UpstreamIPs param.Field[[]FirewallNewParamsUpstreamIP] `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation param.Field[FirewallNewParamsAttackMitigation] `json:"attack_mitigation"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests param.Field[bool] `json:"deprecate_any_requests"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback param.Field[bool] `json:"ecs_fallback"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL param.Field[float64] `json:"maximum_cache_ttl"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL param.Field[float64] `json:"minimum_cache_ttl"`
	// Negative DNS Cache TTL.
	NegativeCacheTTL param.Field[float64] `json:"negative_cache_ttl"`
	// Deprecated alias for "upstream_ips".
	OriginIPs param.Field[interface{}] `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit param.Field[float64] `json:"ratelimit"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries param.Field[float64] `json:"retries"`
}

func (r FirewallNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Upstream DNS Server IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type FirewallNewParamsUpstreamIP interface {
	ImplementsDNSFirewallNewParamsUpstreamIP()
}

// Attack mitigation settings.
type FirewallNewParamsAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled param.Field[bool] `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy param.Field[interface{}] `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy param.Field[bool] `json:"only_when_upstream_unhealthy"`
}

func (r FirewallNewParamsAttackMitigation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallNewResponseEnvelope struct {
	Errors   []FirewallNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallNewResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallNewResponseEnvelopeJSON    `json:"-"`
}

// firewallNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallNewResponseEnvelope]
type firewallNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    firewallNewResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FirewallNewResponseEnvelopeErrors]
type firewallNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallNewResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    firewallNewResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallNewResponseEnvelopeMessages]
type firewallNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallNewResponseEnvelopeSuccess bool

const (
	FirewallNewResponseEnvelopeSuccessTrue FirewallNewResponseEnvelopeSuccess = true
)

type FirewallListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of clusters per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallListParams]'s query parameters as `url.Values`.
func (r FirewallListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type FirewallDeleteResponseEnvelope struct {
	Errors   []FirewallDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallDeleteResponseEnvelopeJSON    `json:"-"`
}

// firewallDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallDeleteResponseEnvelope]
type firewallDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallDeleteResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    firewallDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallDeleteResponseEnvelopeErrors]
type firewallDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallDeleteResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    firewallDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallDeleteResponseEnvelopeMessages]
type firewallDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallDeleteResponseEnvelopeSuccess bool

const (
	FirewallDeleteResponseEnvelopeSuccessTrue FirewallDeleteResponseEnvelopeSuccess = true
)

type FirewallEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests param.Field[bool]                              `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       param.Field[[]FirewallEditParamsDNSFirewallIP] `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback param.Field[bool] `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL param.Field[float64] `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL param.Field[float64] `json:"minimum_cache_ttl,required"`
	// DNS Firewall Cluster Name.
	Name        param.Field[string]                         `json:"name,required"`
	UpstreamIPs param.Field[[]FirewallEditParamsUpstreamIP] `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation param.Field[FirewallEditParamsAttackMitigation] `json:"attack_mitigation"`
	// Negative DNS Cache TTL.
	NegativeCacheTTL param.Field[float64] `json:"negative_cache_ttl"`
	// Deprecated alias for "upstream_ips".
	OriginIPs param.Field[interface{}] `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit param.Field[float64] `json:"ratelimit"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries param.Field[float64] `json:"retries"`
}

func (r FirewallEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type FirewallEditParamsDNSFirewallIP interface {
	ImplementsDNSFirewallEditParamsDNSFirewallIP()
}

// Upstream DNS Server IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type FirewallEditParamsUpstreamIP interface {
	ImplementsDNSFirewallEditParamsUpstreamIP()
}

// Attack mitigation settings.
type FirewallEditParamsAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled param.Field[bool] `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy param.Field[interface{}] `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy param.Field[bool] `json:"only_when_upstream_unhealthy"`
}

func (r FirewallEditParamsAttackMitigation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallEditResponseEnvelope struct {
	Errors   []FirewallEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallEditResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallEditResponseEnvelopeJSON    `json:"-"`
}

// firewallEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallEditResponseEnvelope]
type firewallEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallEditResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    firewallEditResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FirewallEditResponseEnvelopeErrors]
type firewallEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallEditResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    firewallEditResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallEditResponseEnvelopeMessages]
type firewallEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallEditResponseEnvelopeSuccess bool

const (
	FirewallEditResponseEnvelopeSuccessTrue FirewallEditResponseEnvelopeSuccess = true
)

type FirewallGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type FirewallGetResponseEnvelope struct {
	Errors   []FirewallGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallGetResponseEnvelopeJSON    `json:"-"`
}

// firewallGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallGetResponseEnvelope]
type firewallGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    firewallGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FirewallGetResponseEnvelopeErrors]
type firewallGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    firewallGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallGetResponseEnvelopeMessages]
type firewallGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallGetResponseEnvelopeSuccess bool

const (
	FirewallGetResponseEnvelopeSuccessTrue FirewallGetResponseEnvelopeSuccess = true
)
