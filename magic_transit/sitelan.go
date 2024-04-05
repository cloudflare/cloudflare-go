// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SiteLANService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSiteLANService] method instead.
type SiteLANService struct {
	Options []option.RequestOption
}

// NewSiteLANService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteLANService(opts ...option.RequestOption) (r *SiteLANService) {
	r = &SiteLANService{}
	r.Options = opts
	return
}

// Creates a new LAN. If the site is in high availability mode, static_addressing
// is required along with secondary and virtual address.
func (r *SiteLANService) New(ctx context.Context, siteID string, params SiteLANNewParams, opts ...option.RequestOption) (res *SiteLANNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteLANNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a specific LAN.
func (r *SiteLANService) Update(ctx context.Context, siteID string, lanID string, params SiteLANUpdateParams, opts ...option.RequestOption) (res *SiteLANUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteLANUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", params.AccountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists LANs associated with an account and site.
func (r *SiteLANService) List(ctx context.Context, siteID string, query SiteLANListParams, opts ...option.RequestOption) (res *SiteLANListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteLANListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans", query.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a specific LAN.
func (r *SiteLANService) Delete(ctx context.Context, siteID string, lanID string, params SiteLANDeleteParams, opts ...option.RequestOption) (res *SiteLANDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteLANDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", params.AccountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific LAN.
func (r *SiteLANService) Get(ctx context.Context, siteID string, lanID string, query SiteLANGetParams, opts ...option.RequestOption) (res *SiteLANGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteLANGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", query.AccountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DHCPRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string      `json:"server_addresses"`
	JSON            dhcpRelayJSON `json:"-"`
}

// dhcpRelayJSON contains the JSON metadata for the struct [DHCPRelay]
type dhcpRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DHCPRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dhcpRelayJSON) RawJSON() string {
	return r.raw
}

type DHCPRelayParam struct {
	// List of DHCP server IPs.
	ServerAddresses param.Field[[]string] `json:"server_addresses"`
}

func (r DHCPRelayParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DHCPServer struct {
	// A valid IPv4 address.
	DHCPPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string `json:"reservations"`
	JSON         dhcpServerJSON    `json:"-"`
}

// dhcpServerJSON contains the JSON metadata for the struct [DHCPServer]
type dhcpServerJSON struct {
	DHCPPoolEnd   apijson.Field
	DHCPPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DHCPServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dhcpServerJSON) RawJSON() string {
	return r.raw
}

type DHCPServerParam struct {
	// A valid IPv4 address.
	DHCPPoolEnd param.Field[string] `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart param.Field[string] `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer param.Field[string] `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations param.Field[map[string]string] `json:"reservations"`
}

func (r DHCPServerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool           `json:"ha_link"`
	Nat           Nat            `json:"nat"`
	Physport      int64          `json:"physport"`
	RoutedSubnets []RoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing StaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64   `json:"vlan_tag"`
	JSON    lanJSON `json:"-"`
}

// lanJSON contains the JSON metadata for the struct [LAN]
type lanJSON struct {
	ID               apijson.Field
	Description      apijson.Field
	HaLink           apijson.Field
	Nat              apijson.Field
	Physport         apijson.Field
	RoutedSubnets    apijson.Field
	SiteID           apijson.Field
	StaticAddressing apijson.Field
	VlanTag          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lanJSON) RawJSON() string {
	return r.raw
}

type Nat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string  `json:"static_prefix"`
	JSON         natJSON `json:"-"`
}

// natJSON contains the JSON metadata for the struct [Nat]
type natJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *Nat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r natJSON) RawJSON() string {
	return r.raw
}

type NatParam struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r NatParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string           `json:"prefix,required"`
	Nat    Nat              `json:"nat"`
	JSON   routedSubnetJSON `json:"-"`
}

// routedSubnetJSON contains the JSON metadata for the struct [RoutedSubnet]
type routedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routedSubnetJSON) RawJSON() string {
	return r.raw
}

type RoutedSubnetParam struct {
	// A valid IPv4 address.
	NextHop param.Field[string] `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix param.Field[string]   `json:"prefix,required"`
	Nat    param.Field[NatParam] `json:"nat"`
}

func (r RoutedSubnetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type StaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string     `json:"address,required"`
	DHCPRelay  DHCPRelay  `json:"dhcp_relay"`
	DHCPServer DHCPServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string               `json:"virtual_address"`
	JSON           staticAddressingJSON `json:"-"`
}

// staticAddressingJSON contains the JSON metadata for the struct
// [StaticAddressing]
type staticAddressingJSON struct {
	Address          apijson.Field
	DHCPRelay        apijson.Field
	DHCPServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *StaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r staticAddressingJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type StaticAddressingParam struct {
	// A valid CIDR notation representing an IP range.
	Address    param.Field[string]          `json:"address,required"`
	DHCPRelay  param.Field[DHCPRelayParam]  `json:"dhcp_relay"`
	DHCPServer param.Field[DHCPServerParam] `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress param.Field[string] `json:"virtual_address"`
}

func (r StaticAddressingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANNewResponse struct {
	LANs []LAN                  `json:"lans"`
	JSON siteLANNewResponseJSON `json:"-"`
}

// siteLANNewResponseJSON contains the JSON metadata for the struct
// [SiteLANNewResponse]
type siteLANNewResponseJSON struct {
	LANs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANNewResponseJSON) RawJSON() string {
	return r.raw
}

type SiteLANUpdateResponse struct {
	LAN  LAN                       `json:"lan"`
	JSON siteLANUpdateResponseJSON `json:"-"`
}

// siteLANUpdateResponseJSON contains the JSON metadata for the struct
// [SiteLANUpdateResponse]
type siteLANUpdateResponseJSON struct {
	LAN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SiteLANListResponse struct {
	LANs []LAN                   `json:"lans"`
	JSON siteLANListResponseJSON `json:"-"`
}

// siteLANListResponseJSON contains the JSON metadata for the struct
// [SiteLANListResponse]
type siteLANListResponseJSON struct {
	LANs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANListResponseJSON) RawJSON() string {
	return r.raw
}

type SiteLANDeleteResponse struct {
	Deleted    bool                      `json:"deleted"`
	DeletedLAN LAN                       `json:"deleted_lan"`
	JSON       siteLANDeleteResponseJSON `json:"-"`
}

// siteLANDeleteResponseJSON contains the JSON metadata for the struct
// [SiteLANDeleteResponse]
type siteLANDeleteResponseJSON struct {
	Deleted     apijson.Field
	DeletedLAN  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SiteLANGetResponse struct {
	LAN  LAN                    `json:"lan"`
	JSON siteLANGetResponseJSON `json:"-"`
}

// siteLANGetResponseJSON contains the JSON metadata for the struct
// [SiteLANGetResponse]
type siteLANGetResponseJSON struct {
	LAN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANGetResponseJSON) RawJSON() string {
	return r.raw
}

type SiteLANNewParams struct {
	// Identifier
	AccountID param.Field[string]              `path:"account_id,required"`
	LAN       param.Field[SiteLANNewParamsLAN] `json:"lan"`
}

func (r SiteLANNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANNewParamsLAN struct {
	Physport param.Field[int64] `json:"physport,required"`
	// VLAN port number.
	VlanTag     param.Field[int64]  `json:"vlan_tag,required"`
	Description param.Field[string] `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        param.Field[bool]                `json:"ha_link"`
	Nat           param.Field[NatParam]            `json:"nat"`
	RoutedSubnets param.Field[[]RoutedSubnetParam] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[StaticAddressingParam] `json:"static_addressing"`
}

func (r SiteLANNewParamsLAN) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   SiteLANNewResponse                                        `json:"result,required"`
	// Whether the API call was successful
	Success SiteLANNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLANNewResponseEnvelopeJSON    `json:"-"`
}

// siteLANNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLANNewResponseEnvelope]
type siteLANNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLANNewResponseEnvelopeSuccess bool

const (
	SiteLANNewResponseEnvelopeSuccessTrue SiteLANNewResponseEnvelopeSuccess = true
)

func (r SiteLANNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteLANNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteLANUpdateParams struct {
	// Identifier
	AccountID param.Field[string]                 `path:"account_id,required"`
	LAN       param.Field[SiteLANUpdateParamsLAN] `json:"lan"`
}

func (r SiteLANUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANUpdateParamsLAN struct {
	Description   param.Field[string]              `json:"description"`
	Nat           param.Field[NatParam]            `json:"nat"`
	Physport      param.Field[int64]               `json:"physport"`
	RoutedSubnets param.Field[[]RoutedSubnetParam] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[StaticAddressingParam] `json:"static_addressing"`
	// VLAN port number.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r SiteLANUpdateParamsLAN) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   SiteLANUpdateResponse                                     `json:"result,required"`
	// Whether the API call was successful
	Success SiteLANUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLANUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteLANUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLANUpdateResponseEnvelope]
type siteLANUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLANUpdateResponseEnvelopeSuccess bool

const (
	SiteLANUpdateResponseEnvelopeSuccessTrue SiteLANUpdateResponseEnvelopeSuccess = true
)

func (r SiteLANUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteLANUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteLANListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteLANListResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   SiteLANListResponse                                       `json:"result,required"`
	// Whether the API call was successful
	Success SiteLANListResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLANListResponseEnvelopeJSON    `json:"-"`
}

// siteLANListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLANListResponseEnvelope]
type siteLANListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLANListResponseEnvelopeSuccess bool

const (
	SiteLANListResponseEnvelopeSuccessTrue SiteLANListResponseEnvelopeSuccess = true
)

func (r SiteLANListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteLANListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteLANDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r SiteLANDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SiteLANDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   SiteLANDeleteResponse                                     `json:"result,required"`
	// Whether the API call was successful
	Success SiteLANDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLANDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteLANDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLANDeleteResponseEnvelope]
type siteLANDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLANDeleteResponseEnvelopeSuccess bool

const (
	SiteLANDeleteResponseEnvelopeSuccessTrue SiteLANDeleteResponseEnvelopeSuccess = true
)

func (r SiteLANDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteLANDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteLANGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteLANGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   SiteLANGetResponse                                        `json:"result,required"`
	// Whether the API call was successful
	Success SiteLANGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLANGetResponseEnvelopeJSON    `json:"-"`
}

// siteLANGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLANGetResponseEnvelope]
type siteLANGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLANGetResponseEnvelopeSuccess bool

const (
	SiteLANGetResponseEnvelopeSuccessTrue SiteLANGetResponseEnvelopeSuccess = true
)

func (r SiteLANGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteLANGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
