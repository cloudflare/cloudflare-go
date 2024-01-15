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

// AccountDNSFirewallService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDNSFirewallService] method
// instead.
type AccountDNSFirewallService struct {
	Options      []option.RequestOption
	DNSAnalytics *AccountDNSFirewallDNSAnalyticService
}

// NewAccountDNSFirewallService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDNSFirewallService(opts ...option.RequestOption) (r *AccountDNSFirewallService) {
	r = &AccountDNSFirewallService{}
	r.Options = opts
	r.DNSAnalytics = NewAccountDNSFirewallDNSAnalyticService(opts...)
	return
}

// Show a single configured DNS Firewall cluster for an account.
func (r *AccountDNSFirewallService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountDNSFirewallGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a DNS Firewall Cluster configuration.
func (r *AccountDNSFirewallService) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountDNSFirewallUpdateParams, opts ...option.RequestOption) (res *AccountDNSFirewallUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a configured DNS Firewall Cluster.
func (r *AccountDNSFirewallService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountDNSFirewallDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a configured DNS Firewall Cluster.
func (r *AccountDNSFirewallService) DNSFirewallNewDNSFirewallCluster(ctx context.Context, accountIdentifier string, body AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParams, opts ...option.RequestOption) (res *AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured DNS Firewall clusters for an account.
func (r *AccountDNSFirewallService) DNSFirewallListDNSFirewallClusters(ctx context.Context, accountIdentifier string, query AccountDNSFirewallDNSFirewallListDNSFirewallClustersParams, opts ...option.RequestOption) (res *shared.Page[AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall", accountIdentifier)
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

type AccountDNSFirewallGetResponse struct {
	Errors   []AccountDNSFirewallGetResponseError   `json:"errors"`
	Messages []AccountDNSFirewallGetResponseMessage `json:"messages"`
	Result   AccountDNSFirewallGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDNSFirewallGetResponseSuccess `json:"success"`
	JSON    accountDNSFirewallGetResponseJSON    `json:"-"`
}

// accountDNSFirewallGetResponseJSON contains the JSON metadata for the struct
// [AccountDNSFirewallGetResponse]
type accountDNSFirewallGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountDNSFirewallGetResponseErrorJSON `json:"-"`
}

// accountDNSFirewallGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountDNSFirewallGetResponseError]
type accountDNSFirewallGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountDNSFirewallGetResponseMessageJSON `json:"-"`
}

// accountDNSFirewallGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountDNSFirewallGetResponseMessage]
type accountDNSFirewallGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallGetResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                               `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []AccountDNSFirewallGetResponseResultDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTtl float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTtl float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                                          `json:"name,required"`
	UpstreamIPs []AccountDNSFirewallGetResponseResultUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation AccountDNSFirewallGetResponseResultAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTtl float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                                 `json:"retries"`
	JSON    accountDNSFirewallGetResponseResultJSON `json:"-"`
}

// accountDNSFirewallGetResponseResultJSON contains the JSON metadata for the
// struct [AccountDNSFirewallGetResponseResult]
type accountDNSFirewallGetResponseResultJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTtl      apijson.Field
	MinimumCacheTtl      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	NegativeCacheTtl     apijson.Field
	OriginIPs            apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountDNSFirewallGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AccountDNSFirewallGetResponseResultDNSFirewallIP interface {
	ImplementsAccountDNSFirewallGetResponseResultDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDNSFirewallGetResponseResultDNSFirewallIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Upstream DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AccountDNSFirewallGetResponseResultUpstreamIP interface {
	ImplementsAccountDNSFirewallGetResponseResultUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDNSFirewallGetResponseResultUpstreamIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Attack mitigation settings.
type AccountDNSFirewallGetResponseResultAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                                    `json:"only_when_upstream_unhealthy"`
	JSON                      accountDNSFirewallGetResponseResultAttackMitigationJSON `json:"-"`
}

// accountDNSFirewallGetResponseResultAttackMitigationJSON contains the JSON
// metadata for the struct [AccountDNSFirewallGetResponseResultAttackMitigation]
type accountDNSFirewallGetResponseResultAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AccountDNSFirewallGetResponseResultAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDNSFirewallGetResponseSuccess bool

const (
	AccountDNSFirewallGetResponseSuccessTrue AccountDNSFirewallGetResponseSuccess = true
)

type AccountDNSFirewallUpdateResponse struct {
	Errors   []AccountDNSFirewallUpdateResponseError   `json:"errors"`
	Messages []AccountDNSFirewallUpdateResponseMessage `json:"messages"`
	Result   AccountDNSFirewallUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDNSFirewallUpdateResponseSuccess `json:"success"`
	JSON    accountDNSFirewallUpdateResponseJSON    `json:"-"`
}

// accountDNSFirewallUpdateResponseJSON contains the JSON metadata for the struct
// [AccountDNSFirewallUpdateResponse]
type accountDNSFirewallUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountDNSFirewallUpdateResponseErrorJSON `json:"-"`
}

// accountDNSFirewallUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountDNSFirewallUpdateResponseError]
type accountDNSFirewallUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountDNSFirewallUpdateResponseMessageJSON `json:"-"`
}

// accountDNSFirewallUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountDNSFirewallUpdateResponseMessage]
type accountDNSFirewallUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallUpdateResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                                  `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []AccountDNSFirewallUpdateResponseResultDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTtl float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTtl float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                                             `json:"name,required"`
	UpstreamIPs []AccountDNSFirewallUpdateResponseResultUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation AccountDNSFirewallUpdateResponseResultAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTtl float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                                    `json:"retries"`
	JSON    accountDNSFirewallUpdateResponseResultJSON `json:"-"`
}

// accountDNSFirewallUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountDNSFirewallUpdateResponseResult]
type accountDNSFirewallUpdateResponseResultJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTtl      apijson.Field
	MinimumCacheTtl      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	NegativeCacheTtl     apijson.Field
	OriginIPs            apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountDNSFirewallUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AccountDNSFirewallUpdateResponseResultDNSFirewallIP interface {
	ImplementsAccountDNSFirewallUpdateResponseResultDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDNSFirewallUpdateResponseResultDNSFirewallIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Upstream DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AccountDNSFirewallUpdateResponseResultUpstreamIP interface {
	ImplementsAccountDNSFirewallUpdateResponseResultUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDNSFirewallUpdateResponseResultUpstreamIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Attack mitigation settings.
type AccountDNSFirewallUpdateResponseResultAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                                       `json:"only_when_upstream_unhealthy"`
	JSON                      accountDNSFirewallUpdateResponseResultAttackMitigationJSON `json:"-"`
}

// accountDNSFirewallUpdateResponseResultAttackMitigationJSON contains the JSON
// metadata for the struct [AccountDNSFirewallUpdateResponseResultAttackMitigation]
type accountDNSFirewallUpdateResponseResultAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AccountDNSFirewallUpdateResponseResultAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDNSFirewallUpdateResponseSuccess bool

const (
	AccountDNSFirewallUpdateResponseSuccessTrue AccountDNSFirewallUpdateResponseSuccess = true
)

type AccountDNSFirewallDeleteResponse struct {
	Errors   []AccountDNSFirewallDeleteResponseError   `json:"errors"`
	Messages []AccountDNSFirewallDeleteResponseMessage `json:"messages"`
	Result   AccountDNSFirewallDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDNSFirewallDeleteResponseSuccess `json:"success"`
	JSON    accountDNSFirewallDeleteResponseJSON    `json:"-"`
}

// accountDNSFirewallDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDNSFirewallDeleteResponse]
type accountDNSFirewallDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDeleteResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountDNSFirewallDeleteResponseErrorJSON `json:"-"`
}

// accountDNSFirewallDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountDNSFirewallDeleteResponseError]
type accountDNSFirewallDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDeleteResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountDNSFirewallDeleteResponseMessageJSON `json:"-"`
}

// accountDNSFirewallDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountDNSFirewallDeleteResponseMessage]
type accountDNSFirewallDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDeleteResponseResult struct {
	// Identifier
	ID   string                                     `json:"id"`
	JSON accountDNSFirewallDeleteResponseResultJSON `json:"-"`
}

// accountDNSFirewallDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountDNSFirewallDeleteResponseResult]
type accountDNSFirewallDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDNSFirewallDeleteResponseSuccess bool

const (
	AccountDNSFirewallDeleteResponseSuccessTrue AccountDNSFirewallDeleteResponseSuccess = true
)

type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponse struct {
	Errors   []AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseError   `json:"errors"`
	Messages []AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseMessage `json:"messages"`
	Result   AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseSuccess `json:"success"`
	JSON    accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseJSON    `json:"-"`
}

// accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseJSON contains the JSON
// metadata for the struct
// [AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponse]
type accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseErrorJSON `json:"-"`
}

// accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseError]
type accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseMessageJSON `json:"-"`
}

// accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseMessage]
type accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                                                            `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTtl float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTtl float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                                                                       `json:"name,required"`
	UpstreamIPs []AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTtl float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                                                              `json:"retries"`
	JSON    accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultJSON `json:"-"`
}

// accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultJSON contains
// the JSON metadata for the struct
// [AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResult]
type accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTtl      apijson.Field
	MinimumCacheTtl      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	NegativeCacheTtl     apijson.Field
	OriginIPs            apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultDNSFirewallIP interface {
	ImplementsAccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultDNSFirewallIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Upstream DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultUpstreamIP interface {
	ImplementsAccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultUpstreamIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Attack mitigation settings.
type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                                                                 `json:"only_when_upstream_unhealthy"`
	JSON                      accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultAttackMitigationJSON `json:"-"`
}

// accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultAttackMitigationJSON
// contains the JSON metadata for the struct
// [AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultAttackMitigation]
type accountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseResultAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseSuccess bool

const (
	AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseSuccessTrue AccountDNSFirewallDNSFirewallNewDNSFirewallClusterResponseSuccess = true
)

type AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                                                        `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTtl float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTtl float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name        string                                                                   `json:"name,required"`
	UpstreamIPs []AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseUpstreamIP `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTtl float64 `json:"negative_cache_ttl,nullable"`
	// Deprecated alias for "upstream_ips".
	OriginIPs interface{} `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries float64                                                          `json:"retries"`
	JSON    accountDNSFirewallDNSFirewallListDNSFirewallClustersResponseJSON `json:"-"`
}

// accountDNSFirewallDNSFirewallListDNSFirewallClustersResponseJSON contains the
// JSON metadata for the struct
// [AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponse]
type accountDNSFirewallDNSFirewallListDNSFirewallClustersResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTtl      apijson.Field
	MinimumCacheTtl      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	NegativeCacheTtl     apijson.Field
	OriginIPs            apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseDNSFirewallIP interface {
	ImplementsAccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseDNSFirewallIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Upstream DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseUpstreamIP interface {
	ImplementsAccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseUpstreamIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseUpstreamIP)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Attack mitigation settings.
type AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy interface{} `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy bool                                                                             `json:"only_when_upstream_unhealthy"`
	JSON                      accountDNSFirewallDNSFirewallListDNSFirewallClustersResponseAttackMitigationJSON `json:"-"`
}

// accountDNSFirewallDNSFirewallListDNSFirewallClustersResponseAttackMitigationJSON
// contains the JSON metadata for the struct
// [AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseAttackMitigation]
type accountDNSFirewallDNSFirewallListDNSFirewallClustersResponseAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenOriginUnhealthy   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSFirewallListDNSFirewallClustersResponseAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallUpdateParams struct {
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests param.Field[bool]                                          `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       param.Field[[]AccountDNSFirewallUpdateParamsDNSFirewallIP] `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback param.Field[bool] `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTtl param.Field[float64] `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTtl param.Field[float64] `json:"minimum_cache_ttl,required"`
	// DNS Firewall Cluster Name.
	Name        param.Field[string]                                     `json:"name,required"`
	UpstreamIPs param.Field[[]AccountDNSFirewallUpdateParamsUpstreamIP] `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation param.Field[AccountDNSFirewallUpdateParamsAttackMitigation] `json:"attack_mitigation"`
	// Negative DNS Cache TTL.
	NegativeCacheTtl param.Field[float64] `json:"negative_cache_ttl"`
	// Deprecated alias for "upstream_ips".
	OriginIPs param.Field[interface{}] `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit param.Field[float64] `json:"ratelimit"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries param.Field[float64] `json:"retries"`
}

func (r AccountDNSFirewallUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloudflare-assigned DNS IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccountDNSFirewallUpdateParamsDNSFirewallIP interface {
	ImplementsAccountDNSFirewallUpdateParamsDNSFirewallIP()
}

// Upstream DNS Server IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccountDNSFirewallUpdateParamsUpstreamIP interface {
	ImplementsAccountDNSFirewallUpdateParamsUpstreamIP()
}

// Attack mitigation settings.
type AccountDNSFirewallUpdateParamsAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled param.Field[bool] `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy param.Field[interface{}] `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy param.Field[bool] `json:"only_when_upstream_unhealthy"`
}

func (r AccountDNSFirewallUpdateParamsAttackMitigation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParams struct {
	// DNS Firewall Cluster Name.
	Name        param.Field[string]                                                               `json:"name,required"`
	UpstreamIPs param.Field[[]AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsUpstreamIP] `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation param.Field[AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsAttackMitigation] `json:"attack_mitigation"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests param.Field[bool] `json:"deprecate_any_requests"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback param.Field[bool] `json:"ecs_fallback"`
	// Maximum DNS Cache TTL.
	MaximumCacheTtl param.Field[float64] `json:"maximum_cache_ttl"`
	// Minimum DNS Cache TTL.
	MinimumCacheTtl param.Field[float64] `json:"minimum_cache_ttl"`
	// Negative DNS Cache TTL.
	NegativeCacheTtl param.Field[float64] `json:"negative_cache_ttl"`
	// Deprecated alias for "upstream_ips".
	OriginIPs param.Field[interface{}] `json:"origin_ips"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster).
	Ratelimit param.Field[float64] `json:"ratelimit"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt).
	Retries param.Field[float64] `json:"retries"`
}

func (r AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Upstream DNS Server IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsUpstreamIP interface {
	ImplementsAccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsUpstreamIP()
}

// Attack mitigation settings.
type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the upstream
	// DNS servers protected.
	Enabled param.Field[bool] `json:"enabled"`
	// Deprecated alias for "only_when_upstream_unhealthy".
	OnlyWhenOriginUnhealthy param.Field[interface{}] `json:"only_when_origin_unhealthy"`
	// Only mitigate attacks when upstream servers seem unhealthy.
	OnlyWhenUpstreamUnhealthy param.Field[bool] `json:"only_when_upstream_unhealthy"`
}

func (r AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsAttackMitigation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDNSFirewallDNSFirewallListDNSFirewallClustersParams struct {
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of clusters per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [AccountDNSFirewallDNSFirewallListDNSFirewallClustersParams]'s query parameters
// as `url.Values`.
func (r AccountDNSFirewallDNSFirewallListDNSFirewallClustersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
