// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *FirewallService) New(ctx context.Context, params FirewallNewParams, opts ...option.RequestOption) (res *DNSFirewallDNSFirewall, err error) {
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
func (r *FirewallService) List(ctx context.Context, params FirewallListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[DNSFirewallDNSFirewall], err error) {
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
func (r *FirewallService) ListAutoPaging(ctx context.Context, params FirewallListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[DNSFirewallDNSFirewall] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a configured DNS Firewall Cluster.
func (r *FirewallService) Delete(ctx context.Context, dnsFirewallID string, body FirewallDeleteParams, opts ...option.RequestOption) (res *FirewallDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", body.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a DNS Firewall Cluster configuration.
func (r *FirewallService) Edit(ctx context.Context, dnsFirewallID string, params FirewallEditParams, opts ...option.RequestOption) (res *DNSFirewallDNSFirewall, err error) {
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
func (r *FirewallService) Get(ctx context.Context, dnsFirewallID string, query FirewallGetParams, opts ...option.RequestOption) (res *DNSFirewallDNSFirewall, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", query.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSFirewallDNSFirewall struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                  `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []DNSFirewallDNSFirewallDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
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
	UpstreamIPs []DNSFirewallDNSFirewallUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation DNSFirewallDNSFirewallAttackMitigation `json:"attack_mitigation,nullable"`
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
	JSON    dnsFirewallDNSFirewallJSON `json:"-"`
}

// dnsFirewallDNSFirewallJSON contains the JSON metadata for the struct
// [DNSFirewallDNSFirewall]
type dnsFirewallDNSFirewallJSON struct {
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

func (r *DNSFirewallDNSFirewall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallDNSFirewallJSON) RawJSON() string {
	return r.raw
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type DNSFirewallDNSFirewallDNSFirewallIP interface {
	ImplementsDNSDNSFirewallDNSFirewallDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallDNSFirewallDNSFirewallIP)(nil)).Elem(),
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
type DNSFirewallDNSFirewallUpstreamIP interface {
	ImplementsDNSDNSFirewallDNSFirewallUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallDNSFirewallUpstreamIP)(nil)).Elem(),
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
type DNSFirewallDNSFirewallAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                       `json:"only_when_upstream_unhealthy"`
	JSON                      dnsFirewallDNSFirewallAttackMitigationJSON `json:"-"`
}

// dnsFirewallDNSFirewallAttackMitigationJSON contains the JSON metadata for the
// struct [DNSFirewallDNSFirewallAttackMitigation]
type dnsFirewallDNSFirewallAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *DNSFirewallDNSFirewallAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallDNSFirewallAttackMitigationJSON) RawJSON() string {
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
	Result   DNSFirewallDNSFirewall                `json:"result,required"`
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
	Result   DNSFirewallDNSFirewall                 `json:"result,required"`
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
	Result   DNSFirewallDNSFirewall                `json:"result,required"`
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
