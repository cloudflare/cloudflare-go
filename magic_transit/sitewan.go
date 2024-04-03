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

// SiteWANService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSiteWANService] method instead.
type SiteWANService struct {
	Options []option.RequestOption
}

// NewSiteWANService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteWANService(opts ...option.RequestOption) (r *SiteWANService) {
	r = &SiteWANService{}
	r.Options = opts
	return
}

// Creates a new WAN.
func (r *SiteWANService) New(ctx context.Context, siteID string, params SiteWANNewParams, opts ...option.RequestOption) (res *SiteWANNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWANNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a specific WAN.
func (r *SiteWANService) Update(ctx context.Context, siteID string, wanID string, params SiteWANUpdateParams, opts ...option.RequestOption) (res *SiteWANUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWANUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", params.AccountID, siteID, wanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists WANs associated with an account and site.
func (r *SiteWANService) List(ctx context.Context, siteID string, query SiteWANListParams, opts ...option.RequestOption) (res *SiteWANListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWANListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans", query.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a specific WAN.
func (r *SiteWANService) Delete(ctx context.Context, siteID string, wanID string, params SiteWANDeleteParams, opts ...option.RequestOption) (res *SiteWANDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWANDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", params.AccountID, siteID, wanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific WAN.
func (r *SiteWANService) Get(ctx context.Context, siteID string, wanID string, query SiteWANGetParams, opts ...option.RequestOption) (res *SiteWANGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWANGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", query.AccountID, siteID, wanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SiteWANNewResponse struct {
	WANs []SiteWANNewResponseWAN `json:"wans"`
	JSON siteWANNewResponseJSON  `json:"-"`
}

// siteWANNewResponseJSON contains the JSON metadata for the struct
// [SiteWANNewResponse]
type siteWANNewResponseJSON struct {
	WANs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANNewResponseJSON) RawJSON() string {
	return r.raw
}

type SiteWANNewResponseWAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	Physport    int64  `json:"physport"`
	// Priority of WAN for traffic loadbalancing.
	Priority int64 `json:"priority"`
	// Identifier
	SiteID string `json:"site_id"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing SiteWANNewResponseWANsStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                     `json:"vlan_tag"`
	JSON    siteWANNewResponseWANJSON `json:"-"`
}

// siteWANNewResponseWANJSON contains the JSON metadata for the struct
// [SiteWANNewResponseWAN]
type siteWANNewResponseWANJSON struct {
	ID               apijson.Field
	Description      apijson.Field
	Physport         apijson.Field
	Priority         apijson.Field
	SiteID           apijson.Field
	StaticAddressing apijson.Field
	VlanTag          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWANNewResponseWAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANNewResponseWANJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWANNewResponseWANsStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                                     `json:"secondary_address"`
	JSON             siteWANNewResponseWANsStaticAddressingJSON `json:"-"`
}

// siteWANNewResponseWANsStaticAddressingJSON contains the JSON metadata for the
// struct [SiteWANNewResponseWANsStaticAddressing]
type siteWANNewResponseWANsStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWANNewResponseWANsStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANNewResponseWANsStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteWANUpdateResponse struct {
	WAN  SiteWANUpdateResponseWAN  `json:"wan"`
	JSON siteWANUpdateResponseJSON `json:"-"`
}

// siteWANUpdateResponseJSON contains the JSON metadata for the struct
// [SiteWANUpdateResponse]
type siteWANUpdateResponseJSON struct {
	WAN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SiteWANUpdateResponseWAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	Physport    int64  `json:"physport"`
	// Priority of WAN for traffic loadbalancing.
	Priority int64 `json:"priority"`
	// Identifier
	SiteID string `json:"site_id"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing SiteWANUpdateResponseWANStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                        `json:"vlan_tag"`
	JSON    siteWANUpdateResponseWANJSON `json:"-"`
}

// siteWANUpdateResponseWANJSON contains the JSON metadata for the struct
// [SiteWANUpdateResponseWAN]
type siteWANUpdateResponseWANJSON struct {
	ID               apijson.Field
	Description      apijson.Field
	Physport         apijson.Field
	Priority         apijson.Field
	SiteID           apijson.Field
	StaticAddressing apijson.Field
	VlanTag          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWANUpdateResponseWAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANUpdateResponseWANJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWANUpdateResponseWANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                                       `json:"secondary_address"`
	JSON             siteWANUpdateResponseWANStaticAddressingJSON `json:"-"`
}

// siteWANUpdateResponseWANStaticAddressingJSON contains the JSON metadata for the
// struct [SiteWANUpdateResponseWANStaticAddressing]
type siteWANUpdateResponseWANStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWANUpdateResponseWANStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANUpdateResponseWANStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteWANListResponse struct {
	WANs []SiteWANListResponseWAN `json:"wans"`
	JSON siteWANListResponseJSON  `json:"-"`
}

// siteWANListResponseJSON contains the JSON metadata for the struct
// [SiteWANListResponse]
type siteWANListResponseJSON struct {
	WANs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANListResponseJSON) RawJSON() string {
	return r.raw
}

type SiteWANListResponseWAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	Physport    int64  `json:"physport"`
	// Priority of WAN for traffic loadbalancing.
	Priority int64 `json:"priority"`
	// Identifier
	SiteID string `json:"site_id"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing SiteWANListResponseWANsStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                      `json:"vlan_tag"`
	JSON    siteWANListResponseWANJSON `json:"-"`
}

// siteWANListResponseWANJSON contains the JSON metadata for the struct
// [SiteWANListResponseWAN]
type siteWANListResponseWANJSON struct {
	ID               apijson.Field
	Description      apijson.Field
	Physport         apijson.Field
	Priority         apijson.Field
	SiteID           apijson.Field
	StaticAddressing apijson.Field
	VlanTag          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWANListResponseWAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANListResponseWANJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWANListResponseWANsStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                                      `json:"secondary_address"`
	JSON             siteWANListResponseWANsStaticAddressingJSON `json:"-"`
}

// siteWANListResponseWANsStaticAddressingJSON contains the JSON metadata for the
// struct [SiteWANListResponseWANsStaticAddressing]
type siteWANListResponseWANsStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWANListResponseWANsStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANListResponseWANsStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteWANDeleteResponse struct {
	Deleted    bool                            `json:"deleted"`
	DeletedWAN SiteWANDeleteResponseDeletedWAN `json:"deleted_wan"`
	JSON       siteWANDeleteResponseJSON       `json:"-"`
}

// siteWANDeleteResponseJSON contains the JSON metadata for the struct
// [SiteWANDeleteResponse]
type siteWANDeleteResponseJSON struct {
	Deleted     apijson.Field
	DeletedWAN  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SiteWANDeleteResponseDeletedWAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	Physport    int64  `json:"physport"`
	// Priority of WAN for traffic loadbalancing.
	Priority int64 `json:"priority"`
	// Identifier
	SiteID string `json:"site_id"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing SiteWANDeleteResponseDeletedWANStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                               `json:"vlan_tag"`
	JSON    siteWANDeleteResponseDeletedWANJSON `json:"-"`
}

// siteWANDeleteResponseDeletedWANJSON contains the JSON metadata for the struct
// [SiteWANDeleteResponseDeletedWAN]
type siteWANDeleteResponseDeletedWANJSON struct {
	ID               apijson.Field
	Description      apijson.Field
	Physport         apijson.Field
	Priority         apijson.Field
	SiteID           apijson.Field
	StaticAddressing apijson.Field
	VlanTag          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWANDeleteResponseDeletedWAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANDeleteResponseDeletedWANJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWANDeleteResponseDeletedWANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                                              `json:"secondary_address"`
	JSON             siteWANDeleteResponseDeletedWANStaticAddressingJSON `json:"-"`
}

// siteWANDeleteResponseDeletedWANStaticAddressingJSON contains the JSON metadata
// for the struct [SiteWANDeleteResponseDeletedWANStaticAddressing]
type siteWANDeleteResponseDeletedWANStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWANDeleteResponseDeletedWANStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANDeleteResponseDeletedWANStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteWANGetResponse struct {
	WAN  SiteWANGetResponseWAN  `json:"wan"`
	JSON siteWANGetResponseJSON `json:"-"`
}

// siteWANGetResponseJSON contains the JSON metadata for the struct
// [SiteWANGetResponse]
type siteWANGetResponseJSON struct {
	WAN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANGetResponseJSON) RawJSON() string {
	return r.raw
}

type SiteWANGetResponseWAN struct {
	// Identifier
	ID          string `json:"id"`
	Description string `json:"description"`
	Physport    int64  `json:"physport"`
	// Priority of WAN for traffic loadbalancing.
	Priority int64 `json:"priority"`
	// Identifier
	SiteID string `json:"site_id"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing SiteWANGetResponseWANStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                     `json:"vlan_tag"`
	JSON    siteWANGetResponseWANJSON `json:"-"`
}

// siteWANGetResponseWANJSON contains the JSON metadata for the struct
// [SiteWANGetResponseWAN]
type siteWANGetResponseWANJSON struct {
	ID               apijson.Field
	Description      apijson.Field
	Physport         apijson.Field
	Priority         apijson.Field
	SiteID           apijson.Field
	StaticAddressing apijson.Field
	VlanTag          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWANGetResponseWAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANGetResponseWANJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWANGetResponseWANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                                    `json:"secondary_address"`
	JSON             siteWANGetResponseWANStaticAddressingJSON `json:"-"`
}

// siteWANGetResponseWANStaticAddressingJSON contains the JSON metadata for the
// struct [SiteWANGetResponseWANStaticAddressing]
type siteWANGetResponseWANStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWANGetResponseWANStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANGetResponseWANStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteWANNewParams struct {
	// Identifier
	AccountID param.Field[string]              `path:"account_id,required"`
	WAN       param.Field[SiteWANNewParamsWAN] `json:"wan"`
}

func (r SiteWANNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWANNewParamsWAN struct {
	Physport param.Field[int64] `json:"physport,required"`
	// VLAN port number.
	VlanTag     param.Field[int64]  `json:"vlan_tag,required"`
	Description param.Field[string] `json:"description"`
	Priority    param.Field[int64]  `json:"priority"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing param.Field[SiteWANNewParamsWANStaticAddressing] `json:"static_addressing"`
}

func (r SiteWANNewParamsWAN) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWANNewParamsWANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address param.Field[string] `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress param.Field[string] `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
}

func (r SiteWANNewParamsWANStaticAddressing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWANNewResponseEnvelope struct {
	Errors   []SiteWANNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteWANNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteWANNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteWANNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteWANNewResponseEnvelopeJSON    `json:"-"`
}

// siteWANNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteWANNewResponseEnvelope]
type siteWANNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteWANNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    siteWANNewResponseEnvelopeErrorsJSON `json:"-"`
}

// siteWANNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteWANNewResponseEnvelopeErrors]
type siteWANNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteWANNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    siteWANNewResponseEnvelopeMessagesJSON `json:"-"`
}

// siteWANNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SiteWANNewResponseEnvelopeMessages]
type siteWANNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteWANNewResponseEnvelopeSuccess bool

const (
	SiteWANNewResponseEnvelopeSuccessTrue SiteWANNewResponseEnvelopeSuccess = true
)

func (r SiteWANNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteWANNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteWANUpdateParams struct {
	// Identifier
	AccountID param.Field[string]                 `path:"account_id,required"`
	WAN       param.Field[SiteWANUpdateParamsWAN] `json:"wan"`
}

func (r SiteWANUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWANUpdateParamsWAN struct {
	Description param.Field[string] `json:"description"`
	Physport    param.Field[int64]  `json:"physport"`
	Priority    param.Field[int64]  `json:"priority"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing param.Field[SiteWANUpdateParamsWANStaticAddressing] `json:"static_addressing"`
	// VLAN port number.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r SiteWANUpdateParamsWAN) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWANUpdateParamsWANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address param.Field[string] `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress param.Field[string] `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
}

func (r SiteWANUpdateParamsWANStaticAddressing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWANUpdateResponseEnvelope struct {
	Errors   []SiteWANUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteWANUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteWANUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteWANUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteWANUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteWANUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteWANUpdateResponseEnvelope]
type siteWANUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteWANUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteWANUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// siteWANUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SiteWANUpdateResponseEnvelopeErrors]
type siteWANUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteWANUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    siteWANUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// siteWANUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteWANUpdateResponseEnvelopeMessages]
type siteWANUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteWANUpdateResponseEnvelopeSuccess bool

const (
	SiteWANUpdateResponseEnvelopeSuccessTrue SiteWANUpdateResponseEnvelopeSuccess = true
)

func (r SiteWANUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteWANUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteWANListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteWANListResponseEnvelope struct {
	Errors   []SiteWANListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteWANListResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteWANListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteWANListResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteWANListResponseEnvelopeJSON    `json:"-"`
}

// siteWANListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteWANListResponseEnvelope]
type siteWANListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteWANListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    siteWANListResponseEnvelopeErrorsJSON `json:"-"`
}

// siteWANListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteWANListResponseEnvelopeErrors]
type siteWANListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteWANListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteWANListResponseEnvelopeMessagesJSON `json:"-"`
}

// siteWANListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteWANListResponseEnvelopeMessages]
type siteWANListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteWANListResponseEnvelopeSuccess bool

const (
	SiteWANListResponseEnvelopeSuccessTrue SiteWANListResponseEnvelopeSuccess = true
)

func (r SiteWANListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteWANListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteWANDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r SiteWANDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SiteWANDeleteResponseEnvelope struct {
	Errors   []SiteWANDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteWANDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteWANDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteWANDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteWANDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteWANDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteWANDeleteResponseEnvelope]
type siteWANDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteWANDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteWANDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// siteWANDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SiteWANDeleteResponseEnvelopeErrors]
type siteWANDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteWANDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    siteWANDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// siteWANDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteWANDeleteResponseEnvelopeMessages]
type siteWANDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteWANDeleteResponseEnvelopeSuccess bool

const (
	SiteWANDeleteResponseEnvelopeSuccessTrue SiteWANDeleteResponseEnvelopeSuccess = true
)

func (r SiteWANDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteWANDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteWANGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteWANGetResponseEnvelope struct {
	Errors   []SiteWANGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteWANGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteWANGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteWANGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteWANGetResponseEnvelopeJSON    `json:"-"`
}

// siteWANGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteWANGetResponseEnvelope]
type siteWANGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteWANGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    siteWANGetResponseEnvelopeErrorsJSON `json:"-"`
}

// siteWANGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteWANGetResponseEnvelopeErrors]
type siteWANGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteWANGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    siteWANGetResponseEnvelopeMessagesJSON `json:"-"`
}

// siteWANGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SiteWANGetResponseEnvelopeMessages]
type siteWANGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWANGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWANGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteWANGetResponseEnvelopeSuccess bool

const (
	SiteWANGetResponseEnvelopeSuccessTrue SiteWANGetResponseEnvelopeSuccess = true
)

func (r SiteWANGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteWANGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
