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

// DNSFirewallService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDNSFirewallService] method
// instead.
type DNSFirewallService struct {
	Options      []option.RequestOption
	DNSAnalytics *DNSFirewallDNSAnalyticService
}

// NewDNSFirewallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSFirewallService(opts ...option.RequestOption) (r *DNSFirewallService) {
	r = &DNSFirewallService{}
	r.Options = opts
	r.DNSAnalytics = NewDNSFirewallDNSAnalyticService(opts...)
	return
}

// Create a configured DNS Firewall Cluster.
func (r *DNSFirewallService) New(ctx context.Context, accountID string, body DNSFirewallNewParams, opts ...option.RequestOption) (res *DNSFirewallNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSFirewallNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured DNS Firewall clusters for an account.
func (r *DNSFirewallService) List(ctx context.Context, accountID string, query DNSFirewallListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[DNSFirewallListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall", accountID)
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

// List configured DNS Firewall clusters for an account.
func (r *DNSFirewallService) ListAutoPaging(ctx context.Context, accountID string, query DNSFirewallListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[DNSFirewallListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountID, query, opts...))
}

// Delete a configured DNS Firewall Cluster.
func (r *DNSFirewallService) Delete(ctx context.Context, accountID string, dnsFirewallID string, opts ...option.RequestOption) (res *DNSFirewallDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSFirewallDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", accountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a DNS Firewall Cluster configuration.
func (r *DNSFirewallService) Edit(ctx context.Context, accountID string, dnsFirewallID string, body DNSFirewallEditParams, opts ...option.RequestOption) (res *DNSFirewallEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSFirewallEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", accountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show a single configured DNS Firewall cluster for an account.
func (r *DNSFirewallService) Get(ctx context.Context, accountID string, dnsFirewallID string, opts ...option.RequestOption) (res *DNSFirewallGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSFirewallGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", accountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSFirewallNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                  `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []DNSFirewallNewResponseDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                             `json:"name,required"`
	UpstreamIPs []DNSFirewallNewResponseUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation DNSFirewallNewResponseAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                    `json:"retries"`
	JSON    dnsFirewallNewResponseJSON `json:"-"`
}

// dnsFirewallNewResponseJSON contains the JSON metadata for the struct
// [DNSFirewallNewResponse]
type dnsFirewallNewResponseJSON struct {
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

func (r *DNSFirewallNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type DNSFirewallNewResponseDNSFirewallIP interface {
	ImplementsDNSFirewallNewResponseDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallNewResponseDNSFirewallIP)(nil)).Elem(),
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
type DNSFirewallNewResponseUpstreamIP interface {
	ImplementsDNSFirewallNewResponseUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallNewResponseUpstreamIP)(nil)).Elem(),
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
type DNSFirewallNewResponseAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                       `json:"only_when_upstream_unhealthy"`
	JSON                      dnsFirewallNewResponseAttackMitigationJSON `json:"-"`
}

// dnsFirewallNewResponseAttackMitigationJSON contains the JSON metadata for the
// struct [DNSFirewallNewResponseAttackMitigation]
type dnsFirewallNewResponseAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *DNSFirewallNewResponseAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                   `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []DNSFirewallListResponseDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                              `json:"name,required"`
	UpstreamIPs []DNSFirewallListResponseUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation DNSFirewallListResponseAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                     `json:"retries"`
	JSON    dnsFirewallListResponseJSON `json:"-"`
}

// dnsFirewallListResponseJSON contains the JSON metadata for the struct
// [DNSFirewallListResponse]
type dnsFirewallListResponseJSON struct {
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

func (r *DNSFirewallListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type DNSFirewallListResponseDNSFirewallIP interface {
	ImplementsDNSFirewallListResponseDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallListResponseDNSFirewallIP)(nil)).Elem(),
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
type DNSFirewallListResponseUpstreamIP interface {
	ImplementsDNSFirewallListResponseUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallListResponseUpstreamIP)(nil)).Elem(),
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
type DNSFirewallListResponseAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                        `json:"only_when_upstream_unhealthy"`
	JSON                      dnsFirewallListResponseAttackMitigationJSON `json:"-"`
}

// dnsFirewallListResponseAttackMitigationJSON contains the JSON metadata for the
// struct [DNSFirewallListResponseAttackMitigation]
type dnsFirewallListResponseAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *DNSFirewallListResponseAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDeleteResponse struct {
	// Identifier
	ID   string                        `json:"id"`
	JSON dnsFirewallDeleteResponseJSON `json:"-"`
}

// dnsFirewallDeleteResponseJSON contains the JSON metadata for the struct
// [DNSFirewallDeleteResponse]
type dnsFirewallDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallEditResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                   `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []DNSFirewallEditResponseDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                              `json:"name,required"`
	UpstreamIPs []DNSFirewallEditResponseUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation DNSFirewallEditResponseAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                     `json:"retries"`
	JSON    dnsFirewallEditResponseJSON `json:"-"`
}

// dnsFirewallEditResponseJSON contains the JSON metadata for the struct
// [DNSFirewallEditResponse]
type dnsFirewallEditResponseJSON struct {
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

func (r *DNSFirewallEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type DNSFirewallEditResponseDNSFirewallIP interface {
	ImplementsDNSFirewallEditResponseDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallEditResponseDNSFirewallIP)(nil)).Elem(),
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
type DNSFirewallEditResponseUpstreamIP interface {
	ImplementsDNSFirewallEditResponseUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallEditResponseUpstreamIP)(nil)).Elem(),
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
type DNSFirewallEditResponseAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                        `json:"only_when_upstream_unhealthy"`
	JSON                      dnsFirewallEditResponseAttackMitigationJSON `json:"-"`
}

// dnsFirewallEditResponseAttackMitigationJSON contains the JSON metadata for the
// struct [DNSFirewallEditResponseAttackMitigation]
type dnsFirewallEditResponseAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *DNSFirewallEditResponseAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                  `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []DNSFirewallGetResponseDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                             `json:"name,required"`
	UpstreamIPs []DNSFirewallGetResponseUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation DNSFirewallGetResponseAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                    `json:"retries"`
	JSON    dnsFirewallGetResponseJSON `json:"-"`
}

// dnsFirewallGetResponseJSON contains the JSON metadata for the struct
// [DNSFirewallGetResponse]
type dnsFirewallGetResponseJSON struct {
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

func (r *DNSFirewallGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type DNSFirewallGetResponseDNSFirewallIP interface {
	ImplementsDNSFirewallGetResponseDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallGetResponseDNSFirewallIP)(nil)).Elem(),
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
type DNSFirewallGetResponseUpstreamIP interface {
	ImplementsDNSFirewallGetResponseUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallGetResponseUpstreamIP)(nil)).Elem(),
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
type DNSFirewallGetResponseAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                       `json:"only_when_upstream_unhealthy"`
	JSON                      dnsFirewallGetResponseAttackMitigationJSON `json:"-"`
}

// dnsFirewallGetResponseAttackMitigationJSON contains the JSON metadata for the
// struct [DNSFirewallGetResponseAttackMitigation]
type dnsFirewallGetResponseAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *DNSFirewallGetResponseAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallNewParams struct {
	// DNS Firewall Cluster Name.
	Name        param.Field[string]                           `json:"name,required"`
	UpstreamIPs param.Field[[]DNSFirewallNewParamsUpstreamIP] `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation param.Field[DNSFirewallNewParamsAttackMitigation] `json:"attack_mitigation"`
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

func (r DNSFirewallNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Upstream DNS Server IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type DNSFirewallNewParamsUpstreamIP interface {
	ImplementsDNSFirewallNewParamsUpstreamIP()
}

// Attack mitigation settings.
type DNSFirewallNewParamsAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled param.Field[bool] `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy param.Field[interface{}] `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy param.Field[bool] `json:"only_when_upstream_unhealthy"`
}

func (r DNSFirewallNewParamsAttackMitigation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSFirewallNewResponseEnvelope struct {
	Errors   []DNSFirewallNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSFirewallNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSFirewallNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsFirewallNewResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSFirewallNewResponseEnvelope]
type dnsFirewallNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dnsFirewallNewResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsFirewallNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSFirewallNewResponseEnvelopeErrors]
type dnsFirewallNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dnsFirewallNewResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsFirewallNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSFirewallNewResponseEnvelopeMessages]
type dnsFirewallNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallNewResponseEnvelopeSuccess bool

const (
	DNSFirewallNewResponseEnvelopeSuccessTrue DNSFirewallNewResponseEnvelopeSuccess = true
)

type DNSFirewallListParams struct {
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of clusters per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [DNSFirewallListParams]'s query parameters as `url.Values`.
func (r DNSFirewallListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DNSFirewallDeleteResponseEnvelope struct {
	Errors   []DNSFirewallDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSFirewallDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSFirewallDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsFirewallDeleteResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSFirewallDeleteResponseEnvelope]
type dnsFirewallDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    dnsFirewallDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsFirewallDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSFirewallDeleteResponseEnvelopeErrors]
type dnsFirewallDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    dnsFirewallDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsFirewallDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSFirewallDeleteResponseEnvelopeMessages]
type dnsFirewallDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallDeleteResponseEnvelopeSuccess bool

const (
	DNSFirewallDeleteResponseEnvelopeSuccessTrue DNSFirewallDeleteResponseEnvelopeSuccess = true
)

type DNSFirewallEditParams struct {
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests param.Field[bool]                                 `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       param.Field[[]DNSFirewallEditParamsDNSFirewallIP] `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback param.Field[bool] `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTTL param.Field[float64] `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTTL param.Field[float64] `json:"minimum_cache_ttl,required"`
	// DNS Firewall Cluster Name.
	Name        param.Field[string]                            `json:"name,required"`
	UpstreamIPs param.Field[[]DNSFirewallEditParamsUpstreamIP] `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation param.Field[DNSFirewallEditParamsAttackMitigation] `json:"attack_mitigation"`
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

func (r DNSFirewallEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type DNSFirewallEditParamsDNSFirewallIP interface {
	ImplementsDNSFirewallEditParamsDNSFirewallIP()
}

// Upstream DNS Server IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type DNSFirewallEditParamsUpstreamIP interface {
	ImplementsDNSFirewallEditParamsUpstreamIP()
}

// Attack mitigation settings.
type DNSFirewallEditParamsAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled param.Field[bool] `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy param.Field[interface{}] `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy param.Field[bool] `json:"only_when_upstream_unhealthy"`
}

func (r DNSFirewallEditParamsAttackMitigation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSFirewallEditResponseEnvelope struct {
	Errors   []DNSFirewallEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallEditResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSFirewallEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSFirewallEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsFirewallEditResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSFirewallEditResponseEnvelope]
type dnsFirewallEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallEditResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    dnsFirewallEditResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsFirewallEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSFirewallEditResponseEnvelopeErrors]
type dnsFirewallEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallEditResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    dnsFirewallEditResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsFirewallEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSFirewallEditResponseEnvelopeMessages]
type dnsFirewallEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallEditResponseEnvelopeSuccess bool

const (
	DNSFirewallEditResponseEnvelopeSuccessTrue DNSFirewallEditResponseEnvelopeSuccess = true
)

type DNSFirewallGetResponseEnvelope struct {
	Errors   []DNSFirewallGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSFirewallGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSFirewallGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsFirewallGetResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSFirewallGetResponseEnvelope]
type dnsFirewallGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dnsFirewallGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsFirewallGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSFirewallGetResponseEnvelopeErrors]
type dnsFirewallGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    dnsFirewallGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsFirewallGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSFirewallGetResponseEnvelopeMessages]
type dnsFirewallGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallGetResponseEnvelopeSuccess bool

const (
	DNSFirewallGetResponseEnvelopeSuccessTrue DNSFirewallGetResponseEnvelopeSuccess = true
)
