// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// FirewallService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFirewallService] method instead.
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
	var env FirewallNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured DNS Firewall clusters for an account.
func (r *FirewallService) List(ctx context.Context, params FirewallListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[FirewallListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
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
func (r *FirewallService) ListAutoPaging(ctx context.Context, params FirewallListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[FirewallListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a configured DNS Firewall Cluster.
func (r *FirewallService) Delete(ctx context.Context, dnsFirewallID string, body FirewallDeleteParams, opts ...option.RequestOption) (res *FirewallDeleteResponse, err error) {
	var env FirewallDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", body.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify a DNS Firewall Cluster configuration.
func (r *FirewallService) Edit(ctx context.Context, dnsFirewallID string, params FirewallEditParams, opts ...option.RequestOption) (res *FirewallEditResponse, err error) {
	var env FirewallEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
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
	var env FirewallGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", query.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Attack mitigation settings.
type AttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                 `json:"only_when_upstream_unhealthy"`
	JSON                      attackMitigationJSON `json:"-"`
}

// attackMitigationJSON contains the JSON metadata for the struct
// [AttackMitigation]
type attackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackMitigationJSON) RawJSON() string {
	return r.raw
}

// Attack mitigation settings.
type AttackMitigationParam struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled param.Field[bool] `json:"enabled"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy param.Field[bool] `json:"only_when_upstream_unhealthy"`
}

func (r AttackMitigationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallIPs = string

type UpstreamIPs = string

type UpstreamIPsParam = string

type FirewallNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool          `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallIPs `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	ECSFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS cache TTL. This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS cache TTL. This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name string `json:"name,required"`
	// Negative DNS cache TTL. This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,required,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,required,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries     float64       `json:"retries,required"`
	UpstreamIPs []UpstreamIPs `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation AttackMitigation        `json:"attack_mitigation,nullable"`
	JSON             firewallNewResponseJSON `json:"-"`
}

// firewallNewResponseJSON contains the JSON metadata for the struct
// [FirewallNewResponse]
type firewallNewResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	ECSFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	NegativeCacheTTL     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FirewallNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallNewResponseJSON) RawJSON() string {
	return r.raw
}

type FirewallListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool          `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallIPs `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	ECSFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS cache TTL. This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS cache TTL. This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name string `json:"name,required"`
	// Negative DNS cache TTL. This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,required,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,required,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries     float64       `json:"retries,required"`
	UpstreamIPs []UpstreamIPs `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation AttackMitigation         `json:"attack_mitigation,nullable"`
	JSON             firewallListResponseJSON `json:"-"`
}

// firewallListResponseJSON contains the JSON metadata for the struct
// [FirewallListResponse]
type firewallListResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	ECSFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	NegativeCacheTTL     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FirewallListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallListResponseJSON) RawJSON() string {
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
	DeprecateAnyRequests bool          `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallIPs `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	ECSFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS cache TTL. This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS cache TTL. This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name string `json:"name,required"`
	// Negative DNS cache TTL. This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,required,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,required,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries     float64       `json:"retries,required"`
	UpstreamIPs []UpstreamIPs `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation AttackMitigation         `json:"attack_mitigation,nullable"`
	JSON             firewallEditResponseJSON `json:"-"`
}

// firewallEditResponseJSON contains the JSON metadata for the struct
// [FirewallEditResponse]
type firewallEditResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	ECSFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	NegativeCacheTTL     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FirewallEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallEditResponseJSON) RawJSON() string {
	return r.raw
}

type FirewallGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool          `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallIPs `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	ECSFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS cache TTL. This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS cache TTL. This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name string `json:"name,required"`
	// Negative DNS cache TTL. This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,required,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,required,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries     float64       `json:"retries,required"`
	UpstreamIPs []UpstreamIPs `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation AttackMitigation        `json:"attack_mitigation,nullable"`
	JSON             firewallGetResponseJSON `json:"-"`
}

// firewallGetResponseJSON contains the JSON metadata for the struct
// [FirewallGetResponse]
type firewallGetResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	ECSFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	NegativeCacheTTL     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *FirewallGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallGetResponseJSON) RawJSON() string {
	return r.raw
}

type FirewallNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// DNS Firewall Cluster Name.
	Name        param.Field[string]             `json:"name,required"`
	UpstreamIPs param.Field[[]UpstreamIPsParam] `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation param.Field[AttackMitigationParam] `json:"attack_mitigation"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests param.Field[bool] `json:"deprecate_any_requests"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	ECSFallback param.Field[bool] `json:"ecs_fallback"`
	// Maximum DNS cache TTL. This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL param.Field[float64] `json:"maximum_cache_ttl"`
	// Minimum DNS cache TTL. This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL param.Field[float64] `json:"minimum_cache_ttl"`
	// Negative DNS cache TTL. This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL param.Field[float64] `json:"negative_cache_ttl"`
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

type FirewallNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success FirewallNewResponseEnvelopeSuccess `json:"success,required"`
	Result  FirewallNewResponse                `json:"result"`
	JSON    firewallNewResponseEnvelopeJSON    `json:"-"`
}

// firewallNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallNewResponseEnvelope]
type firewallNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallNewResponseEnvelopeSuccess bool

const (
	FirewallNewResponseEnvelopeSuccessTrue FirewallNewResponseEnvelopeSuccess = true
)

func (r FirewallNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type FirewallDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type FirewallDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success FirewallDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  FirewallDeleteResponse                `json:"result"`
	JSON    firewallDeleteResponseEnvelopeJSON    `json:"-"`
}

// firewallDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallDeleteResponseEnvelope]
type firewallDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallDeleteResponseEnvelopeSuccess bool

const (
	FirewallDeleteResponseEnvelopeSuccessTrue FirewallDeleteResponseEnvelopeSuccess = true
)

func (r FirewallDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FirewallEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Attack mitigation settings.
	AttackMitigation param.Field[AttackMitigationParam] `json:"attack_mitigation"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests param.Field[bool] `json:"deprecate_any_requests"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	ECSFallback param.Field[bool] `json:"ecs_fallback"`
	// Maximum DNS cache TTL. This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL param.Field[float64] `json:"maximum_cache_ttl"`
	// Minimum DNS cache TTL. This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL param.Field[float64] `json:"minimum_cache_ttl"`
	// DNS Firewall Cluster Name.
	Name param.Field[string] `json:"name"`
	// Negative DNS cache TTL. This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL param.Field[float64] `json:"negative_cache_ttl"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit param.Field[float64] `json:"ratelimit"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries     param.Field[float64]            `json:"retries"`
	UpstreamIPs param.Field[[]UpstreamIPsParam] `json:"upstream_ips" format:"ipv4"`
}

func (r FirewallEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success FirewallEditResponseEnvelopeSuccess `json:"success,required"`
	Result  FirewallEditResponse                `json:"result"`
	JSON    firewallEditResponseEnvelopeJSON    `json:"-"`
}

// firewallEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallEditResponseEnvelope]
type firewallEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallEditResponseEnvelopeSuccess bool

const (
	FirewallEditResponseEnvelopeSuccessTrue FirewallEditResponseEnvelopeSuccess = true
)

func (r FirewallEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FirewallGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type FirewallGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success FirewallGetResponseEnvelopeSuccess `json:"success,required"`
	Result  FirewallGetResponse                `json:"result"`
	JSON    firewallGetResponseEnvelopeJSON    `json:"-"`
}

// firewallGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallGetResponseEnvelope]
type firewallGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallGetResponseEnvelopeSuccess bool

const (
	FirewallGetResponseEnvelopeSuccessTrue FirewallGetResponseEnvelopeSuccess = true
)

func (r FirewallGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FirewallGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
