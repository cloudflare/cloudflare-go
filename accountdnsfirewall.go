// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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
func (r *AccountDNSFirewallService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *DNSFirewallSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a DNS Firewall Cluster configuration.
func (r *AccountDNSFirewallService) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountDNSFirewallUpdateParams, opts ...option.RequestOption) (res *DNSFirewallSingleResponse, err error) {
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
func (r *AccountDNSFirewallService) DNSFirewallNewDNSFirewallCluster(ctx context.Context, accountIdentifier string, body AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParams, opts ...option.RequestOption) (res *DNSFirewallSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List configured DNS Firewall clusters for an account.
func (r *AccountDNSFirewallService) DNSFirewallListDNSFirewallClusters(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *DNSFirewallResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DNSFirewallResponseCollection struct {
	Errors     []DNSFirewallResponseCollectionError    `json:"errors"`
	Messages   []DNSFirewallResponseCollectionMessage  `json:"messages"`
	Result     []DNSFirewallResponseCollectionResult   `json:"result"`
	ResultInfo DNSFirewallResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DNSFirewallResponseCollectionSuccess `json:"success"`
	JSON    dnsFirewallResponseCollectionJSON    `json:"-"`
}

// dnsFirewallResponseCollectionJSON contains the JSON metadata for the struct
// [DNSFirewallResponseCollection]
type dnsFirewallResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallResponseCollectionError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    dnsFirewallResponseCollectionErrorJSON `json:"-"`
}

// dnsFirewallResponseCollectionErrorJSON contains the JSON metadata for the struct
// [DNSFirewallResponseCollectionError]
type dnsFirewallResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallResponseCollectionMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    dnsFirewallResponseCollectionMessageJSON `json:"-"`
}

// dnsFirewallResponseCollectionMessageJSON contains the JSON metadata for the
// struct [DNSFirewallResponseCollectionMessage]
type dnsFirewallResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallResponseCollectionResult struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                               `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []DNSFirewallResponseCollectionResultDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTtl float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTtl float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name      string                                        `json:"name,required"`
	OriginIPs []DNSFirewallResponseCollectionResultOriginIP `json:"origin_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation DNSFirewallResponseCollectionResultAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTtl float64 `json:"negative_cache_ttl,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the origin nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from origin nameservers (not
	// counting the initial attempt).
	Retries float64                                 `json:"retries"`
	JSON    dnsFirewallResponseCollectionResultJSON `json:"-"`
}

// dnsFirewallResponseCollectionResultJSON contains the JSON metadata for the
// struct [DNSFirewallResponseCollectionResult]
type dnsFirewallResponseCollectionResultJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTtl      apijson.Field
	MinimumCacheTtl      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	OriginIPs            apijson.Field
	AttackMitigation     apijson.Field
	NegativeCacheTtl     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DNSFirewallResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Origin DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type DNSFirewallResponseCollectionResultDNSFirewallIP interface {
	ImplementsDNSFirewallResponseCollectionResultDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallResponseCollectionResultDNSFirewallIP)(nil)).Elem(),
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

// Origin DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type DNSFirewallResponseCollectionResultOriginIP interface {
	ImplementsDNSFirewallResponseCollectionResultOriginIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallResponseCollectionResultOriginIP)(nil)).Elem(),
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
type DNSFirewallResponseCollectionResultAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the origin
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Only mitigate attacks when origin servers seem unhealthy.
	OnlyWhenOriginUnhealthy bool                                                    `json:"only_when_origin_unhealthy"`
	JSON                    dnsFirewallResponseCollectionResultAttackMitigationJSON `json:"-"`
}

// dnsFirewallResponseCollectionResultAttackMitigationJSON contains the JSON
// metadata for the struct [DNSFirewallResponseCollectionResultAttackMitigation]
type dnsFirewallResponseCollectionResultAttackMitigationJSON struct {
	Enabled                 apijson.Field
	OnlyWhenOriginUnhealthy apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *DNSFirewallResponseCollectionResultAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       dnsFirewallResponseCollectionResultInfoJSON `json:"-"`
}

// dnsFirewallResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [DNSFirewallResponseCollectionResultInfo]
type dnsFirewallResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallResponseCollectionSuccess bool

const (
	DNSFirewallResponseCollectionSuccessTrue DNSFirewallResponseCollectionSuccess = true
)

type DNSFirewallSingleResponse struct {
	Errors   []DNSFirewallSingleResponseError   `json:"errors"`
	Messages []DNSFirewallSingleResponseMessage `json:"messages"`
	Result   DNSFirewallSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DNSFirewallSingleResponseSuccess `json:"success"`
	JSON    dnsFirewallSingleResponseJSON    `json:"-"`
}

// dnsFirewallSingleResponseJSON contains the JSON metadata for the struct
// [DNSFirewallSingleResponse]
type dnsFirewallSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallSingleResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    dnsFirewallSingleResponseErrorJSON `json:"-"`
}

// dnsFirewallSingleResponseErrorJSON contains the JSON metadata for the struct
// [DNSFirewallSingleResponseError]
type dnsFirewallSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallSingleResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    dnsFirewallSingleResponseMessageJSON `json:"-"`
}

// dnsFirewallSingleResponseMessageJSON contains the JSON metadata for the struct
// [DNSFirewallSingleResponseMessage]
type dnsFirewallSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallSingleResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// Deprecate the response to ANY requests.
	DeprecateAnyRequests bool                                           `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []DNSFirewallSingleResponseResultDNSFirewallIP `json:"dns_firewall_ips,required" format:"ipv4"`
	// Forward client IP (resolver) subnet if no EDNS Client Subnet is sent.
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS Cache TTL.
	MaximumCacheTtl float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS Cache TTL.
	MinimumCacheTtl float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall Cluster Name.
	Name      string                                    `json:"name,required"`
	OriginIPs []DNSFirewallSingleResponseResultOriginIP `json:"origin_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation DNSFirewallSingleResponseResultAttackMitigation `json:"attack_mitigation,nullable"`
	// Negative DNS Cache TTL.
	NegativeCacheTtl float64 `json:"negative_cache_ttl,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the origin nameservers configured on the cluster).
	Ratelimit float64 `json:"ratelimit,nullable"`
	// Number of retries for fetching DNS responses from origin nameservers (not
	// counting the initial attempt).
	Retries float64                             `json:"retries"`
	JSON    dnsFirewallSingleResponseResultJSON `json:"-"`
}

// dnsFirewallSingleResponseResultJSON contains the JSON metadata for the struct
// [DNSFirewallSingleResponseResult]
type dnsFirewallSingleResponseResultJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTtl      apijson.Field
	MinimumCacheTtl      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	OriginIPs            apijson.Field
	AttackMitigation     apijson.Field
	NegativeCacheTtl     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DNSFirewallSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Origin DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type DNSFirewallSingleResponseResultDNSFirewallIP interface {
	ImplementsDNSFirewallSingleResponseResultDNSFirewallIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallSingleResponseResultDNSFirewallIP)(nil)).Elem(),
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

// Origin DNS Server IPv4 Address.
//
// Union satisfied by [shared.UnionString] or [shared.UnionString].
type DNSFirewallSingleResponseResultOriginIP interface {
	ImplementsDNSFirewallSingleResponseResultOriginIP()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSFirewallSingleResponseResultOriginIP)(nil)).Elem(),
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
type DNSFirewallSingleResponseResultAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the origin
	// DNS servers protected.
	Enabled bool `json:"enabled"`
	// Only mitigate attacks when origin servers seem unhealthy.
	OnlyWhenOriginUnhealthy bool                                                `json:"only_when_origin_unhealthy"`
	JSON                    dnsFirewallSingleResponseResultAttackMitigationJSON `json:"-"`
}

// dnsFirewallSingleResponseResultAttackMitigationJSON contains the JSON metadata
// for the struct [DNSFirewallSingleResponseResultAttackMitigation]
type dnsFirewallSingleResponseResultAttackMitigationJSON struct {
	Enabled                 apijson.Field
	OnlyWhenOriginUnhealthy apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *DNSFirewallSingleResponseResultAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallSingleResponseSuccess bool

const (
	DNSFirewallSingleResponseSuccessTrue DNSFirewallSingleResponseSuccess = true
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
	Name      param.Field[string]                                   `json:"name,required"`
	OriginIPs param.Field[[]AccountDNSFirewallUpdateParamsOriginIP] `json:"origin_ips,required" format:"ipv4"`
	// Attack mitigation settings.
	AttackMitigation param.Field[AccountDNSFirewallUpdateParamsAttackMitigation] `json:"attack_mitigation"`
	// Negative DNS Cache TTL.
	NegativeCacheTtl param.Field[float64] `json:"negative_cache_ttl"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the origin nameservers configured on the cluster).
	Ratelimit param.Field[float64] `json:"ratelimit"`
	// Number of retries for fetching DNS responses from origin nameservers (not
	// counting the initial attempt).
	Retries param.Field[float64] `json:"retries"`
}

func (r AccountDNSFirewallUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Origin DNS Server IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccountDNSFirewallUpdateParamsDNSFirewallIP interface {
	ImplementsAccountDNSFirewallUpdateParamsDNSFirewallIP()
}

// Origin DNS Server IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccountDNSFirewallUpdateParamsOriginIP interface {
	ImplementsAccountDNSFirewallUpdateParamsOriginIP()
}

// Attack mitigation settings.
type AccountDNSFirewallUpdateParamsAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the origin
	// DNS servers protected.
	Enabled param.Field[bool] `json:"enabled"`
	// Only mitigate attacks when origin servers seem unhealthy.
	OnlyWhenOriginUnhealthy param.Field[bool] `json:"only_when_origin_unhealthy"`
}

func (r AccountDNSFirewallUpdateParamsAttackMitigation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParams struct {
	// DNS Firewall Cluster Name.
	Name      param.Field[string]                                                             `json:"name,required"`
	OriginIPs param.Field[[]AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsOriginIP] `json:"origin_ips,required" format:"ipv4"`
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
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the origin nameservers configured on the cluster).
	Ratelimit param.Field[float64] `json:"ratelimit"`
	// Number of retries for fetching DNS responses from origin nameservers (not
	// counting the initial attempt).
	Retries param.Field[float64] `json:"retries"`
}

func (r AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Origin DNS Server IPv4 Address.
//
// Satisfied by [shared.UnionString], [shared.UnionString].
type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsOriginIP interface {
	ImplementsAccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsOriginIP()
}

// Attack mitigation settings.
type AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsAttackMitigation struct {
	// When enabled, random-prefix attacks are automatically mitigated and the origin
	// DNS servers protected.
	Enabled param.Field[bool] `json:"enabled"`
	// Only mitigate attacks when origin servers seem unhealthy.
	OnlyWhenOriginUnhealthy param.Field[bool] `json:"only_when_origin_unhealthy"`
}

func (r AccountDNSFirewallDNSFirewallNewDNSFirewallClusterParamsAttackMitigation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
