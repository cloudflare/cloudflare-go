// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SiteLanService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSiteLanService] method instead.
type SiteLanService struct {
	Options []option.RequestOption
}

// NewSiteLanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteLanService(opts ...option.RequestOption) (r *SiteLanService) {
	r = &SiteLanService{}
	r.Options = opts
	return
}

// Creates a new LAN. If the site is in high availability mode, static_addressing
// is required along with secondary and virtual address.
func (r *SiteLanService) New(ctx context.Context, accountIdentifier string, siteIdentifier string, body SiteLanNewParams, opts ...option.RequestOption) (res *SiteLanNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteLanNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a specific LAN.
func (r *SiteLanService) Update(ctx context.Context, accountIdentifier string, siteIdentifier string, lanIdentifier string, body SiteLanUpdateParams, opts ...option.RequestOption) (res *SiteLanUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteLanUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", accountIdentifier, siteIdentifier, lanIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists LANs associated with an account and site.
func (r *SiteLanService) List(ctx context.Context, accountIdentifier string, siteIdentifier string, opts ...option.RequestOption) (res *SiteLanListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteLanListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a specific LAN.
func (r *SiteLanService) Delete(ctx context.Context, accountIdentifier string, siteIdentifier string, lanIdentifier string, opts ...option.RequestOption) (res *SiteLanDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteLanDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", accountIdentifier, siteIdentifier, lanIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific LAN.
func (r *SiteLanService) Get(ctx context.Context, accountIdentifier string, siteIdentifier string, lanIdentifier string, opts ...option.RequestOption) (res *SiteLanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteLanGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", accountIdentifier, siteIdentifier, lanIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SiteLanNewResponse struct {
	Lans []SiteLanNewResponseLan `json:"lans"`
	JSON siteLanNewResponseJSON  `json:"-"`
}

// siteLanNewResponseJSON contains the JSON metadata for the struct
// [SiteLanNewResponse]
type siteLanNewResponseJSON struct {
	Lans        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseJSON) RawJSON() string {
	return r.raw
}

type SiteLanNewResponseLan struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                                 `json:"ha_link"`
	Nat           SiteLanNewResponseLansNat            `json:"nat"`
	Physport      int64                                `json:"physport"`
	RoutedSubnets []SiteLanNewResponseLansRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing SiteLanNewResponseLansStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                     `json:"vlan_tag"`
	JSON    siteLanNewResponseLanJSON `json:"-"`
}

// siteLanNewResponseLanJSON contains the JSON metadata for the struct
// [SiteLanNewResponseLan]
type siteLanNewResponseLanJSON struct {
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

func (r *SiteLanNewResponseLan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseLanJSON) RawJSON() string {
	return r.raw
}

type SiteLanNewResponseLansNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                        `json:"static_prefix"`
	JSON         siteLanNewResponseLansNatJSON `json:"-"`
}

// siteLanNewResponseLansNatJSON contains the JSON metadata for the struct
// [SiteLanNewResponseLansNat]
type siteLanNewResponseLansNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLanNewResponseLansNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseLansNatJSON) RawJSON() string {
	return r.raw
}

type SiteLanNewResponseLansRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                                 `json:"prefix,required"`
	Nat    SiteLanNewResponseLansRoutedSubnetsNat `json:"nat"`
	JSON   siteLanNewResponseLansRoutedSubnetJSON `json:"-"`
}

// siteLanNewResponseLansRoutedSubnetJSON contains the JSON metadata for the struct
// [SiteLanNewResponseLansRoutedSubnet]
type siteLanNewResponseLansRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanNewResponseLansRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseLansRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type SiteLanNewResponseLansRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                     `json:"static_prefix"`
	JSON         siteLanNewResponseLansRoutedSubnetsNatJSON `json:"-"`
}

// siteLanNewResponseLansRoutedSubnetsNatJSON contains the JSON metadata for the
// struct [SiteLanNewResponseLansRoutedSubnetsNat]
type siteLanNewResponseLansRoutedSubnetsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLanNewResponseLansRoutedSubnetsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseLansRoutedSubnetsNatJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLanNewResponseLansStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                                           `json:"address,required"`
	DhcpRelay  SiteLanNewResponseLansStaticAddressingDhcpRelay  `json:"dhcp_relay"`
	DhcpServer SiteLanNewResponseLansStaticAddressingDhcpServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                                     `json:"virtual_address"`
	JSON           siteLanNewResponseLansStaticAddressingJSON `json:"-"`
}

// siteLanNewResponseLansStaticAddressingJSON contains the JSON metadata for the
// struct [SiteLanNewResponseLansStaticAddressing]
type siteLanNewResponseLansStaticAddressingJSON struct {
	Address          apijson.Field
	DhcpRelay        apijson.Field
	DhcpServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteLanNewResponseLansStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseLansStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteLanNewResponseLansStaticAddressingDhcpRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                                            `json:"server_addresses"`
	JSON            siteLanNewResponseLansStaticAddressingDhcpRelayJSON `json:"-"`
}

// siteLanNewResponseLansStaticAddressingDhcpRelayJSON contains the JSON metadata
// for the struct [SiteLanNewResponseLansStaticAddressingDhcpRelay]
type siteLanNewResponseLansStaticAddressingDhcpRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SiteLanNewResponseLansStaticAddressingDhcpRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseLansStaticAddressingDhcpRelayJSON) RawJSON() string {
	return r.raw
}

type SiteLanNewResponseLansStaticAddressingDhcpServer struct {
	// A valid IPv4 address.
	DhcpPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DhcpPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                                    `json:"reservations"`
	JSON         siteLanNewResponseLansStaticAddressingDhcpServerJSON `json:"-"`
}

// siteLanNewResponseLansStaticAddressingDhcpServerJSON contains the JSON metadata
// for the struct [SiteLanNewResponseLansStaticAddressingDhcpServer]
type siteLanNewResponseLansStaticAddressingDhcpServerJSON struct {
	DhcpPoolEnd   apijson.Field
	DhcpPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SiteLanNewResponseLansStaticAddressingDhcpServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseLansStaticAddressingDhcpServerJSON) RawJSON() string {
	return r.raw
}

type SiteLanUpdateResponse struct {
	Lan  SiteLanUpdateResponseLan  `json:"lan"`
	JSON siteLanUpdateResponseJSON `json:"-"`
}

// siteLanUpdateResponseJSON contains the JSON metadata for the struct
// [SiteLanUpdateResponse]
type siteLanUpdateResponseJSON struct {
	Lan         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SiteLanUpdateResponseLan struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                                   `json:"ha_link"`
	Nat           SiteLanUpdateResponseLanNat            `json:"nat"`
	Physport      int64                                  `json:"physport"`
	RoutedSubnets []SiteLanUpdateResponseLanRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing SiteLanUpdateResponseLanStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                        `json:"vlan_tag"`
	JSON    siteLanUpdateResponseLanJSON `json:"-"`
}

// siteLanUpdateResponseLanJSON contains the JSON metadata for the struct
// [SiteLanUpdateResponseLan]
type siteLanUpdateResponseLanJSON struct {
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

func (r *SiteLanUpdateResponseLan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseLanJSON) RawJSON() string {
	return r.raw
}

type SiteLanUpdateResponseLanNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                          `json:"static_prefix"`
	JSON         siteLanUpdateResponseLanNatJSON `json:"-"`
}

// siteLanUpdateResponseLanNatJSON contains the JSON metadata for the struct
// [SiteLanUpdateResponseLanNat]
type siteLanUpdateResponseLanNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLanUpdateResponseLanNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseLanNatJSON) RawJSON() string {
	return r.raw
}

type SiteLanUpdateResponseLanRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                                   `json:"prefix,required"`
	Nat    SiteLanUpdateResponseLanRoutedSubnetsNat `json:"nat"`
	JSON   siteLanUpdateResponseLanRoutedSubnetJSON `json:"-"`
}

// siteLanUpdateResponseLanRoutedSubnetJSON contains the JSON metadata for the
// struct [SiteLanUpdateResponseLanRoutedSubnet]
type siteLanUpdateResponseLanRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanUpdateResponseLanRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseLanRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type SiteLanUpdateResponseLanRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                       `json:"static_prefix"`
	JSON         siteLanUpdateResponseLanRoutedSubnetsNatJSON `json:"-"`
}

// siteLanUpdateResponseLanRoutedSubnetsNatJSON contains the JSON metadata for the
// struct [SiteLanUpdateResponseLanRoutedSubnetsNat]
type siteLanUpdateResponseLanRoutedSubnetsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLanUpdateResponseLanRoutedSubnetsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseLanRoutedSubnetsNatJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLanUpdateResponseLanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                                             `json:"address,required"`
	DhcpRelay  SiteLanUpdateResponseLanStaticAddressingDhcpRelay  `json:"dhcp_relay"`
	DhcpServer SiteLanUpdateResponseLanStaticAddressingDhcpServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                                       `json:"virtual_address"`
	JSON           siteLanUpdateResponseLanStaticAddressingJSON `json:"-"`
}

// siteLanUpdateResponseLanStaticAddressingJSON contains the JSON metadata for the
// struct [SiteLanUpdateResponseLanStaticAddressing]
type siteLanUpdateResponseLanStaticAddressingJSON struct {
	Address          apijson.Field
	DhcpRelay        apijson.Field
	DhcpServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteLanUpdateResponseLanStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseLanStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteLanUpdateResponseLanStaticAddressingDhcpRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                                              `json:"server_addresses"`
	JSON            siteLanUpdateResponseLanStaticAddressingDhcpRelayJSON `json:"-"`
}

// siteLanUpdateResponseLanStaticAddressingDhcpRelayJSON contains the JSON metadata
// for the struct [SiteLanUpdateResponseLanStaticAddressingDhcpRelay]
type siteLanUpdateResponseLanStaticAddressingDhcpRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SiteLanUpdateResponseLanStaticAddressingDhcpRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseLanStaticAddressingDhcpRelayJSON) RawJSON() string {
	return r.raw
}

type SiteLanUpdateResponseLanStaticAddressingDhcpServer struct {
	// A valid IPv4 address.
	DhcpPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DhcpPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                                      `json:"reservations"`
	JSON         siteLanUpdateResponseLanStaticAddressingDhcpServerJSON `json:"-"`
}

// siteLanUpdateResponseLanStaticAddressingDhcpServerJSON contains the JSON
// metadata for the struct [SiteLanUpdateResponseLanStaticAddressingDhcpServer]
type siteLanUpdateResponseLanStaticAddressingDhcpServerJSON struct {
	DhcpPoolEnd   apijson.Field
	DhcpPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SiteLanUpdateResponseLanStaticAddressingDhcpServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseLanStaticAddressingDhcpServerJSON) RawJSON() string {
	return r.raw
}

type SiteLanListResponse struct {
	Lans []SiteLanListResponseLan `json:"lans"`
	JSON siteLanListResponseJSON  `json:"-"`
}

// siteLanListResponseJSON contains the JSON metadata for the struct
// [SiteLanListResponse]
type siteLanListResponseJSON struct {
	Lans        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseJSON) RawJSON() string {
	return r.raw
}

type SiteLanListResponseLan struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                                  `json:"ha_link"`
	Nat           SiteLanListResponseLansNat            `json:"nat"`
	Physport      int64                                 `json:"physport"`
	RoutedSubnets []SiteLanListResponseLansRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing SiteLanListResponseLansStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                      `json:"vlan_tag"`
	JSON    siteLanListResponseLanJSON `json:"-"`
}

// siteLanListResponseLanJSON contains the JSON metadata for the struct
// [SiteLanListResponseLan]
type siteLanListResponseLanJSON struct {
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

func (r *SiteLanListResponseLan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseLanJSON) RawJSON() string {
	return r.raw
}

type SiteLanListResponseLansNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                         `json:"static_prefix"`
	JSON         siteLanListResponseLansNatJSON `json:"-"`
}

// siteLanListResponseLansNatJSON contains the JSON metadata for the struct
// [SiteLanListResponseLansNat]
type siteLanListResponseLansNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLanListResponseLansNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseLansNatJSON) RawJSON() string {
	return r.raw
}

type SiteLanListResponseLansRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                                  `json:"prefix,required"`
	Nat    SiteLanListResponseLansRoutedSubnetsNat `json:"nat"`
	JSON   siteLanListResponseLansRoutedSubnetJSON `json:"-"`
}

// siteLanListResponseLansRoutedSubnetJSON contains the JSON metadata for the
// struct [SiteLanListResponseLansRoutedSubnet]
type siteLanListResponseLansRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanListResponseLansRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseLansRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type SiteLanListResponseLansRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                      `json:"static_prefix"`
	JSON         siteLanListResponseLansRoutedSubnetsNatJSON `json:"-"`
}

// siteLanListResponseLansRoutedSubnetsNatJSON contains the JSON metadata for the
// struct [SiteLanListResponseLansRoutedSubnetsNat]
type siteLanListResponseLansRoutedSubnetsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLanListResponseLansRoutedSubnetsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseLansRoutedSubnetsNatJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLanListResponseLansStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                                            `json:"address,required"`
	DhcpRelay  SiteLanListResponseLansStaticAddressingDhcpRelay  `json:"dhcp_relay"`
	DhcpServer SiteLanListResponseLansStaticAddressingDhcpServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                                      `json:"virtual_address"`
	JSON           siteLanListResponseLansStaticAddressingJSON `json:"-"`
}

// siteLanListResponseLansStaticAddressingJSON contains the JSON metadata for the
// struct [SiteLanListResponseLansStaticAddressing]
type siteLanListResponseLansStaticAddressingJSON struct {
	Address          apijson.Field
	DhcpRelay        apijson.Field
	DhcpServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteLanListResponseLansStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseLansStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteLanListResponseLansStaticAddressingDhcpRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                                             `json:"server_addresses"`
	JSON            siteLanListResponseLansStaticAddressingDhcpRelayJSON `json:"-"`
}

// siteLanListResponseLansStaticAddressingDhcpRelayJSON contains the JSON metadata
// for the struct [SiteLanListResponseLansStaticAddressingDhcpRelay]
type siteLanListResponseLansStaticAddressingDhcpRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SiteLanListResponseLansStaticAddressingDhcpRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseLansStaticAddressingDhcpRelayJSON) RawJSON() string {
	return r.raw
}

type SiteLanListResponseLansStaticAddressingDhcpServer struct {
	// A valid IPv4 address.
	DhcpPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DhcpPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                                     `json:"reservations"`
	JSON         siteLanListResponseLansStaticAddressingDhcpServerJSON `json:"-"`
}

// siteLanListResponseLansStaticAddressingDhcpServerJSON contains the JSON metadata
// for the struct [SiteLanListResponseLansStaticAddressingDhcpServer]
type siteLanListResponseLansStaticAddressingDhcpServerJSON struct {
	DhcpPoolEnd   apijson.Field
	DhcpPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SiteLanListResponseLansStaticAddressingDhcpServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseLansStaticAddressingDhcpServerJSON) RawJSON() string {
	return r.raw
}

type SiteLanDeleteResponse struct {
	Deleted    bool                            `json:"deleted"`
	DeletedLan SiteLanDeleteResponseDeletedLan `json:"deleted_lan"`
	JSON       siteLanDeleteResponseJSON       `json:"-"`
}

// siteLanDeleteResponseJSON contains the JSON metadata for the struct
// [SiteLanDeleteResponse]
type siteLanDeleteResponseJSON struct {
	Deleted     apijson.Field
	DeletedLan  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SiteLanDeleteResponseDeletedLan struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                                          `json:"ha_link"`
	Nat           SiteLanDeleteResponseDeletedLanNat            `json:"nat"`
	Physport      int64                                         `json:"physport"`
	RoutedSubnets []SiteLanDeleteResponseDeletedLanRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing SiteLanDeleteResponseDeletedLanStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                               `json:"vlan_tag"`
	JSON    siteLanDeleteResponseDeletedLanJSON `json:"-"`
}

// siteLanDeleteResponseDeletedLanJSON contains the JSON metadata for the struct
// [SiteLanDeleteResponseDeletedLan]
type siteLanDeleteResponseDeletedLanJSON struct {
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

func (r *SiteLanDeleteResponseDeletedLan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseDeletedLanJSON) RawJSON() string {
	return r.raw
}

type SiteLanDeleteResponseDeletedLanNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                 `json:"static_prefix"`
	JSON         siteLanDeleteResponseDeletedLanNatJSON `json:"-"`
}

// siteLanDeleteResponseDeletedLanNatJSON contains the JSON metadata for the struct
// [SiteLanDeleteResponseDeletedLanNat]
type siteLanDeleteResponseDeletedLanNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLanDeleteResponseDeletedLanNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseDeletedLanNatJSON) RawJSON() string {
	return r.raw
}

type SiteLanDeleteResponseDeletedLanRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                                          `json:"prefix,required"`
	Nat    SiteLanDeleteResponseDeletedLanRoutedSubnetsNat `json:"nat"`
	JSON   siteLanDeleteResponseDeletedLanRoutedSubnetJSON `json:"-"`
}

// siteLanDeleteResponseDeletedLanRoutedSubnetJSON contains the JSON metadata for
// the struct [SiteLanDeleteResponseDeletedLanRoutedSubnet]
type siteLanDeleteResponseDeletedLanRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanDeleteResponseDeletedLanRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseDeletedLanRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type SiteLanDeleteResponseDeletedLanRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                              `json:"static_prefix"`
	JSON         siteLanDeleteResponseDeletedLanRoutedSubnetsNatJSON `json:"-"`
}

// siteLanDeleteResponseDeletedLanRoutedSubnetsNatJSON contains the JSON metadata
// for the struct [SiteLanDeleteResponseDeletedLanRoutedSubnetsNat]
type siteLanDeleteResponseDeletedLanRoutedSubnetsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLanDeleteResponseDeletedLanRoutedSubnetsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseDeletedLanRoutedSubnetsNatJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLanDeleteResponseDeletedLanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                                                    `json:"address,required"`
	DhcpRelay  SiteLanDeleteResponseDeletedLanStaticAddressingDhcpRelay  `json:"dhcp_relay"`
	DhcpServer SiteLanDeleteResponseDeletedLanStaticAddressingDhcpServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                                              `json:"virtual_address"`
	JSON           siteLanDeleteResponseDeletedLanStaticAddressingJSON `json:"-"`
}

// siteLanDeleteResponseDeletedLanStaticAddressingJSON contains the JSON metadata
// for the struct [SiteLanDeleteResponseDeletedLanStaticAddressing]
type siteLanDeleteResponseDeletedLanStaticAddressingJSON struct {
	Address          apijson.Field
	DhcpRelay        apijson.Field
	DhcpServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteLanDeleteResponseDeletedLanStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseDeletedLanStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteLanDeleteResponseDeletedLanStaticAddressingDhcpRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                                                     `json:"server_addresses"`
	JSON            siteLanDeleteResponseDeletedLanStaticAddressingDhcpRelayJSON `json:"-"`
}

// siteLanDeleteResponseDeletedLanStaticAddressingDhcpRelayJSON contains the JSON
// metadata for the struct
// [SiteLanDeleteResponseDeletedLanStaticAddressingDhcpRelay]
type siteLanDeleteResponseDeletedLanStaticAddressingDhcpRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SiteLanDeleteResponseDeletedLanStaticAddressingDhcpRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseDeletedLanStaticAddressingDhcpRelayJSON) RawJSON() string {
	return r.raw
}

type SiteLanDeleteResponseDeletedLanStaticAddressingDhcpServer struct {
	// A valid IPv4 address.
	DhcpPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DhcpPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                                             `json:"reservations"`
	JSON         siteLanDeleteResponseDeletedLanStaticAddressingDhcpServerJSON `json:"-"`
}

// siteLanDeleteResponseDeletedLanStaticAddressingDhcpServerJSON contains the JSON
// metadata for the struct
// [SiteLanDeleteResponseDeletedLanStaticAddressingDhcpServer]
type siteLanDeleteResponseDeletedLanStaticAddressingDhcpServerJSON struct {
	DhcpPoolEnd   apijson.Field
	DhcpPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SiteLanDeleteResponseDeletedLanStaticAddressingDhcpServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseDeletedLanStaticAddressingDhcpServerJSON) RawJSON() string {
	return r.raw
}

type SiteLanGetResponse struct {
	Lan  SiteLanGetResponseLan  `json:"lan"`
	JSON siteLanGetResponseJSON `json:"-"`
}

// siteLanGetResponseJSON contains the JSON metadata for the struct
// [SiteLanGetResponse]
type siteLanGetResponseJSON struct {
	Lan         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseJSON) RawJSON() string {
	return r.raw
}

type SiteLanGetResponseLan struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                                `json:"ha_link"`
	Nat           SiteLanGetResponseLanNat            `json:"nat"`
	Physport      int64                               `json:"physport"`
	RoutedSubnets []SiteLanGetResponseLanRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing SiteLanGetResponseLanStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                     `json:"vlan_tag"`
	JSON    siteLanGetResponseLanJSON `json:"-"`
}

// siteLanGetResponseLanJSON contains the JSON metadata for the struct
// [SiteLanGetResponseLan]
type siteLanGetResponseLanJSON struct {
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

func (r *SiteLanGetResponseLan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseLanJSON) RawJSON() string {
	return r.raw
}

type SiteLanGetResponseLanNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                       `json:"static_prefix"`
	JSON         siteLanGetResponseLanNatJSON `json:"-"`
}

// siteLanGetResponseLanNatJSON contains the JSON metadata for the struct
// [SiteLanGetResponseLanNat]
type siteLanGetResponseLanNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLanGetResponseLanNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseLanNatJSON) RawJSON() string {
	return r.raw
}

type SiteLanGetResponseLanRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                                `json:"prefix,required"`
	Nat    SiteLanGetResponseLanRoutedSubnetsNat `json:"nat"`
	JSON   siteLanGetResponseLanRoutedSubnetJSON `json:"-"`
}

// siteLanGetResponseLanRoutedSubnetJSON contains the JSON metadata for the struct
// [SiteLanGetResponseLanRoutedSubnet]
type siteLanGetResponseLanRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanGetResponseLanRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseLanRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type SiteLanGetResponseLanRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string                                    `json:"static_prefix"`
	JSON         siteLanGetResponseLanRoutedSubnetsNatJSON `json:"-"`
}

// siteLanGetResponseLanRoutedSubnetsNatJSON contains the JSON metadata for the
// struct [SiteLanGetResponseLanRoutedSubnetsNat]
type siteLanGetResponseLanRoutedSubnetsNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SiteLanGetResponseLanRoutedSubnetsNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseLanRoutedSubnetsNatJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLanGetResponseLanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                                          `json:"address,required"`
	DhcpRelay  SiteLanGetResponseLanStaticAddressingDhcpRelay  `json:"dhcp_relay"`
	DhcpServer SiteLanGetResponseLanStaticAddressingDhcpServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                                    `json:"virtual_address"`
	JSON           siteLanGetResponseLanStaticAddressingJSON `json:"-"`
}

// siteLanGetResponseLanStaticAddressingJSON contains the JSON metadata for the
// struct [SiteLanGetResponseLanStaticAddressing]
type siteLanGetResponseLanStaticAddressingJSON struct {
	Address          apijson.Field
	DhcpRelay        apijson.Field
	DhcpServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteLanGetResponseLanStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseLanStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteLanGetResponseLanStaticAddressingDhcpRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                                           `json:"server_addresses"`
	JSON            siteLanGetResponseLanStaticAddressingDhcpRelayJSON `json:"-"`
}

// siteLanGetResponseLanStaticAddressingDhcpRelayJSON contains the JSON metadata
// for the struct [SiteLanGetResponseLanStaticAddressingDhcpRelay]
type siteLanGetResponseLanStaticAddressingDhcpRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SiteLanGetResponseLanStaticAddressingDhcpRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseLanStaticAddressingDhcpRelayJSON) RawJSON() string {
	return r.raw
}

type SiteLanGetResponseLanStaticAddressingDhcpServer struct {
	// A valid IPv4 address.
	DhcpPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DhcpPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer string `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                                   `json:"reservations"`
	JSON         siteLanGetResponseLanStaticAddressingDhcpServerJSON `json:"-"`
}

// siteLanGetResponseLanStaticAddressingDhcpServerJSON contains the JSON metadata
// for the struct [SiteLanGetResponseLanStaticAddressingDhcpServer]
type siteLanGetResponseLanStaticAddressingDhcpServerJSON struct {
	DhcpPoolEnd   apijson.Field
	DhcpPoolStart apijson.Field
	DNSServer     apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SiteLanGetResponseLanStaticAddressingDhcpServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseLanStaticAddressingDhcpServerJSON) RawJSON() string {
	return r.raw
}

type SiteLanNewParams struct {
	Lan param.Field[SiteLanNewParamsLan] `json:"lan"`
}

func (r SiteLanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanNewParamsLan struct {
	Physport param.Field[int64] `json:"physport,required"`
	// VLAN port number.
	VlanTag     param.Field[int64]  `json:"vlan_tag,required"`
	Description param.Field[string] `json:"description"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        param.Field[bool]                              `json:"ha_link"`
	Nat           param.Field[SiteLanNewParamsLanNat]            `json:"nat"`
	RoutedSubnets param.Field[[]SiteLanNewParamsLanRoutedSubnet] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[SiteLanNewParamsLanStaticAddressing] `json:"static_addressing"`
}

func (r SiteLanNewParamsLan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanNewParamsLanNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r SiteLanNewParamsLanNat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanNewParamsLanRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop param.Field[string] `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix param.Field[string]                              `json:"prefix,required"`
	Nat    param.Field[SiteLanNewParamsLanRoutedSubnetsNat] `json:"nat"`
}

func (r SiteLanNewParamsLanRoutedSubnet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanNewParamsLanRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r SiteLanNewParamsLanRoutedSubnetsNat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLanNewParamsLanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    param.Field[string]                                        `json:"address,required"`
	DhcpRelay  param.Field[SiteLanNewParamsLanStaticAddressingDhcpRelay]  `json:"dhcp_relay"`
	DhcpServer param.Field[SiteLanNewParamsLanStaticAddressingDhcpServer] `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress param.Field[string] `json:"virtual_address"`
}

func (r SiteLanNewParamsLanStaticAddressing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanNewParamsLanStaticAddressingDhcpRelay struct {
	// List of DHCP server IPs.
	ServerAddresses param.Field[[]string] `json:"server_addresses"`
}

func (r SiteLanNewParamsLanStaticAddressingDhcpRelay) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanNewParamsLanStaticAddressingDhcpServer struct {
	// A valid IPv4 address.
	DhcpPoolEnd param.Field[string] `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DhcpPoolStart param.Field[string] `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer param.Field[string] `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations param.Field[map[string]string] `json:"reservations"`
}

func (r SiteLanNewParamsLanStaticAddressingDhcpServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanNewResponseEnvelope struct {
	Errors   []SiteLanNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteLanNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteLanNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteLanNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLanNewResponseEnvelopeJSON    `json:"-"`
}

// siteLanNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLanNewResponseEnvelope]
type siteLanNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteLanNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    siteLanNewResponseEnvelopeErrorsJSON `json:"-"`
}

// siteLanNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteLanNewResponseEnvelopeErrors]
type siteLanNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteLanNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    siteLanNewResponseEnvelopeMessagesJSON `json:"-"`
}

// siteLanNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SiteLanNewResponseEnvelopeMessages]
type siteLanNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLanNewResponseEnvelopeSuccess bool

const (
	SiteLanNewResponseEnvelopeSuccessTrue SiteLanNewResponseEnvelopeSuccess = true
)

type SiteLanUpdateParams struct {
	Lan param.Field[SiteLanUpdateParamsLan] `json:"lan"`
}

func (r SiteLanUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanUpdateParamsLan struct {
	Description   param.Field[string]                               `json:"description"`
	Nat           param.Field[SiteLanUpdateParamsLanNat]            `json:"nat"`
	Physport      param.Field[int64]                                `json:"physport"`
	RoutedSubnets param.Field[[]SiteLanUpdateParamsLanRoutedSubnet] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[SiteLanUpdateParamsLanStaticAddressing] `json:"static_addressing"`
	// VLAN port number.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r SiteLanUpdateParamsLan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanUpdateParamsLanNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r SiteLanUpdateParamsLanNat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanUpdateParamsLanRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop param.Field[string] `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix param.Field[string]                                 `json:"prefix,required"`
	Nat    param.Field[SiteLanUpdateParamsLanRoutedSubnetsNat] `json:"nat"`
}

func (r SiteLanUpdateParamsLanRoutedSubnet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanUpdateParamsLanRoutedSubnetsNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r SiteLanUpdateParamsLanRoutedSubnetsNat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type SiteLanUpdateParamsLanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    param.Field[string]                                           `json:"address,required"`
	DhcpRelay  param.Field[SiteLanUpdateParamsLanStaticAddressingDhcpRelay]  `json:"dhcp_relay"`
	DhcpServer param.Field[SiteLanUpdateParamsLanStaticAddressingDhcpServer] `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress param.Field[string] `json:"virtual_address"`
}

func (r SiteLanUpdateParamsLanStaticAddressing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanUpdateParamsLanStaticAddressingDhcpRelay struct {
	// List of DHCP server IPs.
	ServerAddresses param.Field[[]string] `json:"server_addresses"`
}

func (r SiteLanUpdateParamsLanStaticAddressingDhcpRelay) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanUpdateParamsLanStaticAddressingDhcpServer struct {
	// A valid IPv4 address.
	DhcpPoolEnd param.Field[string] `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DhcpPoolStart param.Field[string] `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer param.Field[string] `json:"dns_server"`
	// Mapping of MAC addresses to IP addresses
	Reservations param.Field[map[string]string] `json:"reservations"`
}

func (r SiteLanUpdateParamsLanStaticAddressingDhcpServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteLanUpdateResponseEnvelope struct {
	Errors   []SiteLanUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteLanUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteLanUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteLanUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLanUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteLanUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLanUpdateResponseEnvelope]
type siteLanUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteLanUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteLanUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// siteLanUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SiteLanUpdateResponseEnvelopeErrors]
type siteLanUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteLanUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    siteLanUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// siteLanUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteLanUpdateResponseEnvelopeMessages]
type siteLanUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLanUpdateResponseEnvelopeSuccess bool

const (
	SiteLanUpdateResponseEnvelopeSuccessTrue SiteLanUpdateResponseEnvelopeSuccess = true
)

type SiteLanListResponseEnvelope struct {
	Errors   []SiteLanListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteLanListResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteLanListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteLanListResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLanListResponseEnvelopeJSON    `json:"-"`
}

// siteLanListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLanListResponseEnvelope]
type siteLanListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteLanListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    siteLanListResponseEnvelopeErrorsJSON `json:"-"`
}

// siteLanListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteLanListResponseEnvelopeErrors]
type siteLanListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteLanListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteLanListResponseEnvelopeMessagesJSON `json:"-"`
}

// siteLanListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteLanListResponseEnvelopeMessages]
type siteLanListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLanListResponseEnvelopeSuccess bool

const (
	SiteLanListResponseEnvelopeSuccessTrue SiteLanListResponseEnvelopeSuccess = true
)

type SiteLanDeleteResponseEnvelope struct {
	Errors   []SiteLanDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteLanDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteLanDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteLanDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLanDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteLanDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLanDeleteResponseEnvelope]
type siteLanDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteLanDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteLanDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// siteLanDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SiteLanDeleteResponseEnvelopeErrors]
type siteLanDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteLanDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    siteLanDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// siteLanDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteLanDeleteResponseEnvelopeMessages]
type siteLanDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLanDeleteResponseEnvelopeSuccess bool

const (
	SiteLanDeleteResponseEnvelopeSuccessTrue SiteLanDeleteResponseEnvelopeSuccess = true
)

type SiteLanGetResponseEnvelope struct {
	Errors   []SiteLanGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteLanGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteLanGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteLanGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteLanGetResponseEnvelopeJSON    `json:"-"`
}

// siteLanGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteLanGetResponseEnvelope]
type siteLanGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteLanGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    siteLanGetResponseEnvelopeErrorsJSON `json:"-"`
}

// siteLanGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteLanGetResponseEnvelopeErrors]
type siteLanGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteLanGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    siteLanGetResponseEnvelopeMessagesJSON `json:"-"`
}

// siteLanGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SiteLanGetResponseEnvelopeMessages]
type siteLanGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteLanGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteLanGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteLanGetResponseEnvelopeSuccess bool

const (
	SiteLanGetResponseEnvelopeSuccessTrue SiteLanGetResponseEnvelopeSuccess = true
)
