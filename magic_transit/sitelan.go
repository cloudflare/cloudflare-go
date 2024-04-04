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

type SiteLANNewResponse struct {
	LANs []SiteLANNewResponseLAN `json:"lans"`
	JSON siteLANNewResponseJSON  `json:"-"`
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

type SiteLANNewResponseLAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                                 `json:"ha_link"`
	Nat           SiteLANNewResponseLANsNat            `json:"nat"`
	Physport      int64                                `json:"physport"`
	RoutedSubnets []SiteLANNewResponseLANsRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing SiteLANNewResponseLANsStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                     `json:"vlan_tag"`
	JSON    siteLANNewResponseLANJSON `json:"-"`
}

// siteLANNewResponseLANJSON contains the JSON metadata for the struct
// [SiteLANNewResponseLAN]
type siteLANNewResponseLANJSON struct {
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

func (r *SiteLANNewResponseLAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANNewResponseLANJSON) RawJSON() string {
	return r.raw
}

type SiteLANNewResponseLANsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                        `json:"static_prefix"`
	JSON         siteLANNewResponseLANsNatJSON `json:"-"`
}

// siteLANNewResponseLANsNatJSON contains the JSON metadata for the struct
// [SiteLANNewResponseLANsNat]
type siteLANNewResponseLANsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLANNewResponseLANsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANNewResponseLANsNatJSON) RawJSON() string {
	return r.raw
}

type SiteLANNewResponseLANsRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                                 `json:"prefix,required"`
	Nat    SiteLANNewResponseLANsRoutedSubnetsNat `json:"nat"`
	JSON   siteLANNewResponseLANsRoutedSubnetJSON `json:"-"`
}

// siteLANNewResponseLANsRoutedSubnetJSON contains the JSON metadata for the struct
// [SiteLANNewResponseLANsRoutedSubnet]
type siteLANNewResponseLANsRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANNewResponseLANsRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANNewResponseLANsRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type SiteLANNewResponseLANsRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                     `json:"static_prefix"`
	JSON         siteLANNewResponseLANsRoutedSubnetsNatJSON `json:"-"`
}

// siteLANNewResponseLANsRoutedSubnetsNatJSON contains the JSON metadata for the
// struct [SiteLANNewResponseLANsRoutedSubnetsNat]
type siteLANNewResponseLANsRoutedSubnetsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLANNewResponseLANsRoutedSubnetsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANNewResponseLANsRoutedSubnetsNatJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLANNewResponseLANsStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                                           `json:"address,required"`
	DHCPRelay  SiteLANNewResponseLANsStaticAddressingDHCPRelay  `json:"dhcp_relay"`
	DHCPServer SiteLANNewResponseLANsStaticAddressingDHCPServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                                     `json:"virtual_address"`
	JSON           siteLANNewResponseLANsStaticAddressingJSON `json:"-"`
}

// siteLANNewResponseLANsStaticAddressingJSON contains the JSON metadata for the
// struct [SiteLANNewResponseLANsStaticAddressing]
type siteLANNewResponseLANsStaticAddressingJSON struct {
	Address          apijson.Field
	DHCPRelay        apijson.Field
	DHCPServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteLANNewResponseLANsStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANNewResponseLANsStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteLANNewResponseLANsStaticAddressingDHCPRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                                            `json:"server_addresses"`
	JSON            siteLANNewResponseLANsStaticAddressingDHCPRelayJSON `json:"-"`
}

// siteLANNewResponseLANsStaticAddressingDHCPRelayJSON contains the JSON metadata
// for the struct [SiteLANNewResponseLANsStaticAddressingDHCPRelay]
type siteLANNewResponseLANsStaticAddressingDHCPRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SiteLANNewResponseLANsStaticAddressingDHCPRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANNewResponseLANsStaticAddressingDHCPRelayJSON) RawJSON() string {
	return r.raw
}

type SiteLANNewResponseLANsStaticAddressingDHCPServer struct {
	// A valid IPv4 address.
	DHCPPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                                    `json:"reservations"`
	JSON         siteLANNewResponseLANsStaticAddressingDHCPServerJSON `json:"-"`
}

// siteLANNewResponseLANsStaticAddressingDHCPServerJSON contains the JSON metadata
// for the struct [SiteLANNewResponseLANsStaticAddressingDHCPServer]
type siteLANNewResponseLANsStaticAddressingDHCPServerJSON struct {
	DHCPPoolEnd   apijson.Field
	DHCPPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SiteLANNewResponseLANsStaticAddressingDHCPServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANNewResponseLANsStaticAddressingDHCPServerJSON) RawJSON() string {
	return r.raw
}

type SiteLANUpdateResponse struct {
	LAN  SiteLANUpdateResponseLAN  `json:"lan"`
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

type SiteLANUpdateResponseLAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                                   `json:"ha_link"`
	Nat           SiteLANUpdateResponseLANNat            `json:"nat"`
	Physport      int64                                  `json:"physport"`
	RoutedSubnets []SiteLANUpdateResponseLANRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing SiteLANUpdateResponseLANStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                        `json:"vlan_tag"`
	JSON    siteLANUpdateResponseLANJSON `json:"-"`
}

// siteLANUpdateResponseLANJSON contains the JSON metadata for the struct
// [SiteLANUpdateResponseLAN]
type siteLANUpdateResponseLANJSON struct {
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

func (r *SiteLANUpdateResponseLAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANUpdateResponseLANJSON) RawJSON() string {
	return r.raw
}

type SiteLANUpdateResponseLANNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                          `json:"static_prefix"`
	JSON         siteLANUpdateResponseLANNatJSON `json:"-"`
}

// siteLANUpdateResponseLANNatJSON contains the JSON metadata for the struct
// [SiteLANUpdateResponseLANNat]
type siteLANUpdateResponseLANNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLANUpdateResponseLANNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANUpdateResponseLANNatJSON) RawJSON() string {
	return r.raw
}

type SiteLANUpdateResponseLANRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                                   `json:"prefix,required"`
	Nat    SiteLANUpdateResponseLANRoutedSubnetsNat `json:"nat"`
	JSON   siteLANUpdateResponseLANRoutedSubnetJSON `json:"-"`
}

// siteLANUpdateResponseLANRoutedSubnetJSON contains the JSON metadata for the
// struct [SiteLANUpdateResponseLANRoutedSubnet]
type siteLANUpdateResponseLANRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANUpdateResponseLANRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANUpdateResponseLANRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type SiteLANUpdateResponseLANRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                       `json:"static_prefix"`
	JSON         siteLANUpdateResponseLANRoutedSubnetsNatJSON `json:"-"`
}

// siteLANUpdateResponseLANRoutedSubnetsNatJSON contains the JSON metadata for the
// struct [SiteLANUpdateResponseLANRoutedSubnetsNat]
type siteLANUpdateResponseLANRoutedSubnetsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLANUpdateResponseLANRoutedSubnetsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANUpdateResponseLANRoutedSubnetsNatJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLANUpdateResponseLANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                                             `json:"address,required"`
	DHCPRelay  SiteLANUpdateResponseLANStaticAddressingDHCPRelay  `json:"dhcp_relay"`
	DHCPServer SiteLANUpdateResponseLANStaticAddressingDHCPServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                                       `json:"virtual_address"`
	JSON           siteLANUpdateResponseLANStaticAddressingJSON `json:"-"`
}

// siteLANUpdateResponseLANStaticAddressingJSON contains the JSON metadata for the
// struct [SiteLANUpdateResponseLANStaticAddressing]
type siteLANUpdateResponseLANStaticAddressingJSON struct {
	Address          apijson.Field
	DHCPRelay        apijson.Field
	DHCPServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteLANUpdateResponseLANStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANUpdateResponseLANStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteLANUpdateResponseLANStaticAddressingDHCPRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                                              `json:"server_addresses"`
	JSON            siteLANUpdateResponseLANStaticAddressingDHCPRelayJSON `json:"-"`
}

// siteLANUpdateResponseLANStaticAddressingDHCPRelayJSON contains the JSON metadata
// for the struct [SiteLANUpdateResponseLANStaticAddressingDHCPRelay]
type siteLANUpdateResponseLANStaticAddressingDHCPRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SiteLANUpdateResponseLANStaticAddressingDHCPRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANUpdateResponseLANStaticAddressingDHCPRelayJSON) RawJSON() string {
	return r.raw
}

type SiteLANUpdateResponseLANStaticAddressingDHCPServer struct {
	// A valid IPv4 address.
	DHCPPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                                      `json:"reservations"`
	JSON         siteLANUpdateResponseLANStaticAddressingDHCPServerJSON `json:"-"`
}

// siteLANUpdateResponseLANStaticAddressingDHCPServerJSON contains the JSON
// metadata for the struct [SiteLANUpdateResponseLANStaticAddressingDHCPServer]
type siteLANUpdateResponseLANStaticAddressingDHCPServerJSON struct {
	DHCPPoolEnd   apijson.Field
	DHCPPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SiteLANUpdateResponseLANStaticAddressingDHCPServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANUpdateResponseLANStaticAddressingDHCPServerJSON) RawJSON() string {
	return r.raw
}

type SiteLANListResponse struct {
	LANs []SiteLANListResponseLAN `json:"lans"`
	JSON siteLANListResponseJSON  `json:"-"`
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

type SiteLANListResponseLAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                                  `json:"ha_link"`
	Nat           SiteLANListResponseLANsNat            `json:"nat"`
	Physport      int64                                 `json:"physport"`
	RoutedSubnets []SiteLANListResponseLANsRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing SiteLANListResponseLANsStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                      `json:"vlan_tag"`
	JSON    siteLANListResponseLANJSON `json:"-"`
}

// siteLANListResponseLANJSON contains the JSON metadata for the struct
// [SiteLANListResponseLAN]
type siteLANListResponseLANJSON struct {
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

func (r *SiteLANListResponseLAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANListResponseLANJSON) RawJSON() string {
	return r.raw
}

type SiteLANListResponseLANsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                         `json:"static_prefix"`
	JSON         siteLANListResponseLANsNatJSON `json:"-"`
}

// siteLANListResponseLANsNatJSON contains the JSON metadata for the struct
// [SiteLANListResponseLANsNat]
type siteLANListResponseLANsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLANListResponseLANsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANListResponseLANsNatJSON) RawJSON() string {
	return r.raw
}

type SiteLANListResponseLANsRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                                  `json:"prefix,required"`
	Nat    SiteLANListResponseLANsRoutedSubnetsNat `json:"nat"`
	JSON   siteLANListResponseLANsRoutedSubnetJSON `json:"-"`
}

// siteLANListResponseLANsRoutedSubnetJSON contains the JSON metadata for the
// struct [SiteLANListResponseLANsRoutedSubnet]
type siteLANListResponseLANsRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANListResponseLANsRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANListResponseLANsRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type SiteLANListResponseLANsRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                      `json:"static_prefix"`
	JSON         siteLANListResponseLANsRoutedSubnetsNatJSON `json:"-"`
}

// siteLANListResponseLANsRoutedSubnetsNatJSON contains the JSON metadata for the
// struct [SiteLANListResponseLANsRoutedSubnetsNat]
type siteLANListResponseLANsRoutedSubnetsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLANListResponseLANsRoutedSubnetsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANListResponseLANsRoutedSubnetsNatJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLANListResponseLANsStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                                            `json:"address,required"`
	DHCPRelay  SiteLANListResponseLANsStaticAddressingDHCPRelay  `json:"dhcp_relay"`
	DHCPServer SiteLANListResponseLANsStaticAddressingDHCPServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                                      `json:"virtual_address"`
	JSON           siteLANListResponseLANsStaticAddressingJSON `json:"-"`
}

// siteLANListResponseLANsStaticAddressingJSON contains the JSON metadata for the
// struct [SiteLANListResponseLANsStaticAddressing]
type siteLANListResponseLANsStaticAddressingJSON struct {
	Address          apijson.Field
	DHCPRelay        apijson.Field
	DHCPServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteLANListResponseLANsStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANListResponseLANsStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteLANListResponseLANsStaticAddressingDHCPRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                                             `json:"server_addresses"`
	JSON            siteLANListResponseLANsStaticAddressingDHCPRelayJSON `json:"-"`
}

// siteLANListResponseLANsStaticAddressingDHCPRelayJSON contains the JSON metadata
// for the struct [SiteLANListResponseLANsStaticAddressingDHCPRelay]
type siteLANListResponseLANsStaticAddressingDHCPRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SiteLANListResponseLANsStaticAddressingDHCPRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANListResponseLANsStaticAddressingDHCPRelayJSON) RawJSON() string {
	return r.raw
}

type SiteLANListResponseLANsStaticAddressingDHCPServer struct {
	// A valid IPv4 address.
	DHCPPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                                     `json:"reservations"`
	JSON         siteLANListResponseLANsStaticAddressingDHCPServerJSON `json:"-"`
}

// siteLANListResponseLANsStaticAddressingDHCPServerJSON contains the JSON metadata
// for the struct [SiteLANListResponseLANsStaticAddressingDHCPServer]
type siteLANListResponseLANsStaticAddressingDHCPServerJSON struct {
	DHCPPoolEnd   apijson.Field
	DHCPPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SiteLANListResponseLANsStaticAddressingDHCPServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANListResponseLANsStaticAddressingDHCPServerJSON) RawJSON() string {
	return r.raw
}

type SiteLANDeleteResponse struct {
	Deleted    bool                            `json:"deleted"`
	DeletedLAN SiteLANDeleteResponseDeletedLAN `json:"deleted_lan"`
	JSON       siteLANDeleteResponseJSON       `json:"-"`
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

type SiteLANDeleteResponseDeletedLAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                                          `json:"ha_link"`
	Nat           SiteLANDeleteResponseDeletedLANNat            `json:"nat"`
	Physport      int64                                         `json:"physport"`
	RoutedSubnets []SiteLANDeleteResponseDeletedLANRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing SiteLANDeleteResponseDeletedLANStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                               `json:"vlan_tag"`
	JSON    siteLANDeleteResponseDeletedLANJSON `json:"-"`
}

// siteLANDeleteResponseDeletedLANJSON contains the JSON metadata for the struct
// [SiteLANDeleteResponseDeletedLAN]
type siteLANDeleteResponseDeletedLANJSON struct {
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

func (r *SiteLANDeleteResponseDeletedLAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANDeleteResponseDeletedLANJSON) RawJSON() string {
	return r.raw
}

type SiteLANDeleteResponseDeletedLANNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                 `json:"static_prefix"`
	JSON         siteLANDeleteResponseDeletedLANNatJSON `json:"-"`
}

// siteLANDeleteResponseDeletedLANNatJSON contains the JSON metadata for the struct
// [SiteLANDeleteResponseDeletedLANNat]
type siteLANDeleteResponseDeletedLANNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLANDeleteResponseDeletedLANNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANDeleteResponseDeletedLANNatJSON) RawJSON() string {
	return r.raw
}

type SiteLANDeleteResponseDeletedLANRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                                          `json:"prefix,required"`
	Nat    SiteLANDeleteResponseDeletedLANRoutedSubnetsNat `json:"nat"`
	JSON   siteLANDeleteResponseDeletedLANRoutedSubnetJSON `json:"-"`
}

// siteLANDeleteResponseDeletedLANRoutedSubnetJSON contains the JSON metadata for
// the struct [SiteLANDeleteResponseDeletedLANRoutedSubnet]
type siteLANDeleteResponseDeletedLANRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANDeleteResponseDeletedLANRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANDeleteResponseDeletedLANRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type SiteLANDeleteResponseDeletedLANRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                              `json:"static_prefix"`
	JSON         siteLANDeleteResponseDeletedLANRoutedSubnetsNatJSON `json:"-"`
}

// siteLANDeleteResponseDeletedLANRoutedSubnetsNatJSON contains the JSON metadata
// for the struct [SiteLANDeleteResponseDeletedLANRoutedSubnetsNat]
type siteLANDeleteResponseDeletedLANRoutedSubnetsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLANDeleteResponseDeletedLANRoutedSubnetsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANDeleteResponseDeletedLANRoutedSubnetsNatJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLANDeleteResponseDeletedLANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                                                    `json:"address,required"`
	DHCPRelay  SiteLANDeleteResponseDeletedLANStaticAddressingDHCPRelay  `json:"dhcp_relay"`
	DHCPServer SiteLANDeleteResponseDeletedLANStaticAddressingDHCPServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                                              `json:"virtual_address"`
	JSON           siteLANDeleteResponseDeletedLANStaticAddressingJSON `json:"-"`
}

// siteLANDeleteResponseDeletedLANStaticAddressingJSON contains the JSON metadata
// for the struct [SiteLANDeleteResponseDeletedLANStaticAddressing]
type siteLANDeleteResponseDeletedLANStaticAddressingJSON struct {
	Address          apijson.Field
	DHCPRelay        apijson.Field
	DHCPServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteLANDeleteResponseDeletedLANStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANDeleteResponseDeletedLANStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteLANDeleteResponseDeletedLANStaticAddressingDHCPRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                                                     `json:"server_addresses"`
	JSON            siteLANDeleteResponseDeletedLANStaticAddressingDHCPRelayJSON `json:"-"`
}

// siteLANDeleteResponseDeletedLANStaticAddressingDHCPRelayJSON contains the JSON
// metadata for the struct
// [SiteLANDeleteResponseDeletedLANStaticAddressingDHCPRelay]
type siteLANDeleteResponseDeletedLANStaticAddressingDHCPRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SiteLANDeleteResponseDeletedLANStaticAddressingDHCPRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANDeleteResponseDeletedLANStaticAddressingDHCPRelayJSON) RawJSON() string {
	return r.raw
}

type SiteLANDeleteResponseDeletedLANStaticAddressingDHCPServer struct {
	// A valid IPv4 address.
	DHCPPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                                             `json:"reservations"`
	JSON         siteLANDeleteResponseDeletedLANStaticAddressingDHCPServerJSON `json:"-"`
}

// siteLANDeleteResponseDeletedLANStaticAddressingDHCPServerJSON contains the JSON
// metadata for the struct
// [SiteLANDeleteResponseDeletedLANStaticAddressingDHCPServer]
type siteLANDeleteResponseDeletedLANStaticAddressingDHCPServerJSON struct {
	DHCPPoolEnd   apijson.Field
	DHCPPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SiteLANDeleteResponseDeletedLANStaticAddressingDHCPServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANDeleteResponseDeletedLANStaticAddressingDHCPServerJSON) RawJSON() string {
	return r.raw
}

type SiteLANGetResponse struct {
	LAN  SiteLANGetResponseLAN  `json:"lan"`
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

type SiteLANGetResponseLAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                                `json:"ha_link"`
	Nat           SiteLANGetResponseLANNat            `json:"nat"`
	Physport      int64                               `json:"physport"`
	RoutedSubnets []SiteLANGetResponseLANRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing SiteLANGetResponseLANStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                     `json:"vlan_tag"`
	JSON    siteLANGetResponseLANJSON `json:"-"`
}

// siteLANGetResponseLANJSON contains the JSON metadata for the struct
// [SiteLANGetResponseLAN]
type siteLANGetResponseLANJSON struct {
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

func (r *SiteLANGetResponseLAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANGetResponseLANJSON) RawJSON() string {
	return r.raw
}

type SiteLANGetResponseLANNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                       `json:"static_prefix"`
	JSON         siteLANGetResponseLANNatJSON `json:"-"`
}

// siteLANGetResponseLANNatJSON contains the JSON metadata for the struct
// [SiteLANGetResponseLANNat]
type siteLANGetResponseLANNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLANGetResponseLANNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANGetResponseLANNatJSON) RawJSON() string {
	return r.raw
}

type SiteLANGetResponseLANRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                                `json:"prefix,required"`
	Nat    SiteLANGetResponseLANRoutedSubnetsNat `json:"nat"`
	JSON   siteLANGetResponseLANRoutedSubnetJSON `json:"-"`
}

// siteLANGetResponseLANRoutedSubnetJSON contains the JSON metadata for the struct
// [SiteLANGetResponseLANRoutedSubnet]
type siteLANGetResponseLANRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLANGetResponseLANRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANGetResponseLANRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type SiteLANGetResponseLANRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                    `json:"static_prefix"`
	JSON         siteLANGetResponseLANRoutedSubnetsNatJSON `json:"-"`
}

// siteLANGetResponseLANRoutedSubnetsNatJSON contains the JSON metadata for the
// struct [SiteLANGetResponseLANRoutedSubnetsNat]
type siteLANGetResponseLANRoutedSubnetsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLANGetResponseLANRoutedSubnetsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANGetResponseLANRoutedSubnetsNatJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLANGetResponseLANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                                          `json:"address,required"`
	DHCPRelay  SiteLANGetResponseLANStaticAddressingDHCPRelay  `json:"dhcp_relay"`
	DHCPServer SiteLANGetResponseLANStaticAddressingDHCPServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                                    `json:"virtual_address"`
	JSON           siteLANGetResponseLANStaticAddressingJSON `json:"-"`
}

// siteLANGetResponseLANStaticAddressingJSON contains the JSON metadata for the
// struct [SiteLANGetResponseLANStaticAddressing]
type siteLANGetResponseLANStaticAddressingJSON struct {
	Address          apijson.Field
	DHCPRelay        apijson.Field
	DHCPServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteLANGetResponseLANStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANGetResponseLANStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteLANGetResponseLANStaticAddressingDHCPRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                                           `json:"server_addresses"`
	JSON            siteLANGetResponseLANStaticAddressingDHCPRelayJSON `json:"-"`
}

// siteLANGetResponseLANStaticAddressingDHCPRelayJSON contains the JSON metadata
// for the struct [SiteLANGetResponseLANStaticAddressingDHCPRelay]
type siteLANGetResponseLANStaticAddressingDHCPRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SiteLANGetResponseLANStaticAddressingDHCPRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANGetResponseLANStaticAddressingDHCPRelayJSON) RawJSON() string {
	return r.raw
}

type SiteLANGetResponseLANStaticAddressingDHCPServer struct {
	// A valid IPv4 address.
	DHCPPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                                   `json:"reservations"`
	JSON         siteLANGetResponseLANStaticAddressingDHCPServerJSON `json:"-"`
}

// siteLANGetResponseLANStaticAddressingDHCPServerJSON contains the JSON metadata
// for the struct [SiteLANGetResponseLANStaticAddressingDHCPServer]
type siteLANGetResponseLANStaticAddressingDHCPServerJSON struct {
	DHCPPoolEnd   apijson.Field
	DHCPPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SiteLANGetResponseLANStaticAddressingDHCPServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLANGetResponseLANStaticAddressingDHCPServerJSON) RawJSON() string {
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
	HaLink        param.Field[bool]                              `json:"ha_link"`
	Nat           param.Field[SiteLANNewParamsLANNat]            `json:"nat"`
	RoutedSubnets param.Field[[]SiteLANNewParamsLANRoutedSubnet] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[SiteLANNewParamsLANStaticAddressing] `json:"static_addressing"`
}

func (r SiteLANNewParamsLAN) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANNewParamsLANNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r SiteLANNewParamsLANNat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANNewParamsLANRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop param.Field[string] `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix param.Field[string]                              `json:"prefix,required"`
	Nat    param.Field[SiteLANNewParamsLANRoutedSubnetsNat] `json:"nat"`
}

func (r SiteLANNewParamsLANRoutedSubnet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANNewParamsLANRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r SiteLANNewParamsLANRoutedSubnetsNat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLANNewParamsLANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    param.Field[string]                                        `json:"address,required"`
	DHCPRelay  param.Field[SiteLANNewParamsLANStaticAddressingDHCPRelay]  `json:"dhcp_relay"`
	DHCPServer param.Field[SiteLANNewParamsLANStaticAddressingDHCPServer] `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress param.Field[string] `json:"virtual_address"`
}

func (r SiteLANNewParamsLANStaticAddressing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANNewParamsLANStaticAddressingDHCPRelay struct {
	// List of DHCP server IPs.
	ServerAddresses param.Field[[]string] `json:"server_addresses"`
}

func (r SiteLANNewParamsLANStaticAddressingDHCPRelay) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANNewParamsLANStaticAddressingDHCPServer struct {
	// A valid IPv4 address.
	DHCPPoolEnd param.Field[string] `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart param.Field[string] `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer param.Field[string] `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations param.Field[map[string]string] `json:"reservations"`
}

func (r SiteLANNewParamsLANStaticAddressingDHCPServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteLANNewResponse    `json:"result,required"`
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
	Description   param.Field[string]                               `json:"description"`
	Nat           param.Field[SiteLANUpdateParamsLANNat]            `json:"nat"`
	Physport      param.Field[int64]                                `json:"physport"`
	RoutedSubnets param.Field[[]SiteLANUpdateParamsLANRoutedSubnet] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[SiteLANUpdateParamsLANStaticAddressing] `json:"static_addressing"`
	// VLAN port number.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r SiteLANUpdateParamsLAN) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANUpdateParamsLANNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r SiteLANUpdateParamsLANNat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANUpdateParamsLANRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop param.Field[string] `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix param.Field[string]                                 `json:"prefix,required"`
	Nat    param.Field[SiteLANUpdateParamsLANRoutedSubnetsNat] `json:"nat"`
}

func (r SiteLANUpdateParamsLANRoutedSubnet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANUpdateParamsLANRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r SiteLANUpdateParamsLANRoutedSubnetsNat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLANUpdateParamsLANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    param.Field[string]                                           `json:"address,required"`
	DHCPRelay  param.Field[SiteLANUpdateParamsLANStaticAddressingDHCPRelay]  `json:"dhcp_relay"`
	DHCPServer param.Field[SiteLANUpdateParamsLANStaticAddressingDHCPServer] `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress param.Field[string] `json:"virtual_address"`
}

func (r SiteLANUpdateParamsLANStaticAddressing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANUpdateParamsLANStaticAddressingDHCPRelay struct {
	// List of DHCP server IPs.
	ServerAddresses param.Field[[]string] `json:"server_addresses"`
}

func (r SiteLANUpdateParamsLANStaticAddressingDHCPRelay) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANUpdateParamsLANStaticAddressingDHCPServer struct {
	// A valid IPv4 address.
	DHCPPoolEnd param.Field[string] `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DHCPPoolStart param.Field[string] `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer param.Field[string] `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations param.Field[map[string]string] `json:"reservations"`
}

func (r SiteLANUpdateParamsLANStaticAddressingDHCPServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLANUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteLANUpdateResponse `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteLANListResponse   `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteLANDeleteResponse `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SiteLANGetResponse    `json:"result,required"`
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
