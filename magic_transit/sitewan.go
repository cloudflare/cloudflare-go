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

// SiteWanService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSiteWanService] method instead.
type SiteWanService struct {
	Options []option.RequestOption
}

// NewSiteWanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteWanService(opts ...option.RequestOption) (r *SiteWanService) {
	r = &SiteWanService{}
	r.Options = opts
	return
}

// Creates a new WAN.
func (r *SiteWanService) New(ctx context.Context, accountIdentifier string, siteIdentifier string, body SiteWanNewParams, opts ...option.RequestOption) (res *SiteWanNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWanNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a specific WAN.
func (r *SiteWanService) Update(ctx context.Context, accountIdentifier string, siteIdentifier string, wanIdentifier string, body SiteWanUpdateParams, opts ...option.RequestOption) (res *SiteWanUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWanUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", accountIdentifier, siteIdentifier, wanIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists WANs associated with an account and site.
func (r *SiteWanService) List(ctx context.Context, accountIdentifier string, siteIdentifier string, opts ...option.RequestOption) (res *SiteWanListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWanListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans", accountIdentifier, siteIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a specific WAN.
func (r *SiteWanService) Delete(ctx context.Context, accountIdentifier string, siteIdentifier string, wanIdentifier string, opts ...option.RequestOption) (res *SiteWanDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWanDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", accountIdentifier, siteIdentifier, wanIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific WAN.
func (r *SiteWanService) Get(ctx context.Context, accountIdentifier string, siteIdentifier string, wanIdentifier string, opts ...option.RequestOption) (res *SiteWanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWanGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", accountIdentifier, siteIdentifier, wanIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SiteWanNewResponse struct {
	Wans []SiteWanNewResponseWan `json:"wans"`
	JSON siteWanNewResponseJSON  `json:"-"`
}

// siteWanNewResponseJSON contains the JSON metadata for the struct
// [SiteWanNewResponse]
type siteWanNewResponseJSON struct {
	Wans        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanNewResponseJSON) RawJSON() string {
	return r.raw
}

type SiteWanNewResponseWan struct {
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
	StaticAddressing SiteWanNewResponseWansStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                     `json:"vlan_tag"`
	JSON    siteWanNewResponseWanJSON `json:"-"`
}

// siteWanNewResponseWanJSON contains the JSON metadata for the struct
// [SiteWanNewResponseWan]
type siteWanNewResponseWanJSON struct {
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

func (r *SiteWanNewResponseWan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanNewResponseWanJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWanNewResponseWansStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                                     `json:"secondary_address"`
	JSON             siteWanNewResponseWansStaticAddressingJSON `json:"-"`
}

// siteWanNewResponseWansStaticAddressingJSON contains the JSON metadata for the
// struct [SiteWanNewResponseWansStaticAddressing]
type siteWanNewResponseWansStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWanNewResponseWansStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanNewResponseWansStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteWanUpdateResponse struct {
	Wan  SiteWanUpdateResponseWan  `json:"wan"`
	JSON siteWanUpdateResponseJSON `json:"-"`
}

// siteWanUpdateResponseJSON contains the JSON metadata for the struct
// [SiteWanUpdateResponse]
type siteWanUpdateResponseJSON struct {
	Wan         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SiteWanUpdateResponseWan struct {
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
	StaticAddressing SiteWanUpdateResponseWanStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                        `json:"vlan_tag"`
	JSON    siteWanUpdateResponseWanJSON `json:"-"`
}

// siteWanUpdateResponseWanJSON contains the JSON metadata for the struct
// [SiteWanUpdateResponseWan]
type siteWanUpdateResponseWanJSON struct {
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

func (r *SiteWanUpdateResponseWan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanUpdateResponseWanJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWanUpdateResponseWanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                                       `json:"secondary_address"`
	JSON             siteWanUpdateResponseWanStaticAddressingJSON `json:"-"`
}

// siteWanUpdateResponseWanStaticAddressingJSON contains the JSON metadata for the
// struct [SiteWanUpdateResponseWanStaticAddressing]
type siteWanUpdateResponseWanStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWanUpdateResponseWanStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanUpdateResponseWanStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteWanListResponse struct {
	Wans []SiteWanListResponseWan `json:"wans"`
	JSON siteWanListResponseJSON  `json:"-"`
}

// siteWanListResponseJSON contains the JSON metadata for the struct
// [SiteWanListResponse]
type siteWanListResponseJSON struct {
	Wans        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanListResponseJSON) RawJSON() string {
	return r.raw
}

type SiteWanListResponseWan struct {
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
	StaticAddressing SiteWanListResponseWansStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                      `json:"vlan_tag"`
	JSON    siteWanListResponseWanJSON `json:"-"`
}

// siteWanListResponseWanJSON contains the JSON metadata for the struct
// [SiteWanListResponseWan]
type siteWanListResponseWanJSON struct {
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

func (r *SiteWanListResponseWan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanListResponseWanJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWanListResponseWansStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                                      `json:"secondary_address"`
	JSON             siteWanListResponseWansStaticAddressingJSON `json:"-"`
}

// siteWanListResponseWansStaticAddressingJSON contains the JSON metadata for the
// struct [SiteWanListResponseWansStaticAddressing]
type siteWanListResponseWansStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWanListResponseWansStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanListResponseWansStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteWanDeleteResponse struct {
	Deleted    bool                            `json:"deleted"`
	DeletedWan SiteWanDeleteResponseDeletedWan `json:"deleted_wan"`
	JSON       siteWanDeleteResponseJSON       `json:"-"`
}

// siteWanDeleteResponseJSON contains the JSON metadata for the struct
// [SiteWanDeleteResponse]
type siteWanDeleteResponseJSON struct {
	Deleted     apijson.Field
	DeletedWan  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SiteWanDeleteResponseDeletedWan struct {
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
	StaticAddressing SiteWanDeleteResponseDeletedWanStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                               `json:"vlan_tag"`
	JSON    siteWanDeleteResponseDeletedWanJSON `json:"-"`
}

// siteWanDeleteResponseDeletedWanJSON contains the JSON metadata for the struct
// [SiteWanDeleteResponseDeletedWan]
type siteWanDeleteResponseDeletedWanJSON struct {
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

func (r *SiteWanDeleteResponseDeletedWan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanDeleteResponseDeletedWanJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWanDeleteResponseDeletedWanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                                              `json:"secondary_address"`
	JSON             siteWanDeleteResponseDeletedWanStaticAddressingJSON `json:"-"`
}

// siteWanDeleteResponseDeletedWanStaticAddressingJSON contains the JSON metadata
// for the struct [SiteWanDeleteResponseDeletedWanStaticAddressing]
type siteWanDeleteResponseDeletedWanStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWanDeleteResponseDeletedWanStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanDeleteResponseDeletedWanStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteWanGetResponse struct {
	Wan  SiteWanGetResponseWan  `json:"wan"`
	JSON siteWanGetResponseJSON `json:"-"`
}

// siteWanGetResponseJSON contains the JSON metadata for the struct
// [SiteWanGetResponse]
type siteWanGetResponseJSON struct {
	Wan         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanGetResponseJSON) RawJSON() string {
	return r.raw
}

type SiteWanGetResponseWan struct {
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
	StaticAddressing SiteWanGetResponseWanStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64                     `json:"vlan_tag"`
	JSON    siteWanGetResponseWanJSON `json:"-"`
}

// siteWanGetResponseWanJSON contains the JSON metadata for the struct
// [SiteWanGetResponseWan]
type siteWanGetResponseWanJSON struct {
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

func (r *SiteWanGetResponseWan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanGetResponseWanJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWanGetResponseWanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                                    `json:"secondary_address"`
	JSON             siteWanGetResponseWanStaticAddressingJSON `json:"-"`
}

// siteWanGetResponseWanStaticAddressingJSON contains the JSON metadata for the
// struct [SiteWanGetResponseWanStaticAddressing]
type siteWanGetResponseWanStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SiteWanGetResponseWanStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanGetResponseWanStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type SiteWanNewParams struct {
	Wan param.Field[SiteWanNewParamsWan] `json:"wan"`
}

func (r SiteWanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWanNewParamsWan struct {
	Physport param.Field[int64] `json:"physport,required"`
	// VLAN port number.
	VlanTag     param.Field[int64]  `json:"vlan_tag,required"`
	Description param.Field[string] `json:"description"`
	Priority    param.Field[int64]  `json:"priority"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing param.Field[SiteWanNewParamsWanStaticAddressing] `json:"static_addressing"`
}

func (r SiteWanNewParamsWan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWanNewParamsWanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address param.Field[string] `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress param.Field[string] `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
}

func (r SiteWanNewParamsWanStaticAddressing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWanNewResponseEnvelope struct {
	Errors   []SiteWanNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteWanNewResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteWanNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteWanNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteWanNewResponseEnvelopeJSON    `json:"-"`
}

// siteWanNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteWanNewResponseEnvelope]
type siteWanNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteWanNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    siteWanNewResponseEnvelopeErrorsJSON `json:"-"`
}

// siteWanNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteWanNewResponseEnvelopeErrors]
type siteWanNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteWanNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    siteWanNewResponseEnvelopeMessagesJSON `json:"-"`
}

// siteWanNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SiteWanNewResponseEnvelopeMessages]
type siteWanNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteWanNewResponseEnvelopeSuccess bool

const (
	SiteWanNewResponseEnvelopeSuccessTrue SiteWanNewResponseEnvelopeSuccess = true
)

type SiteWanUpdateParams struct {
	Wan param.Field[SiteWanUpdateParamsWan] `json:"wan"`
}

func (r SiteWanUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWanUpdateParamsWan struct {
	Description param.Field[string] `json:"description"`
	Physport    param.Field[int64]  `json:"physport"`
	Priority    param.Field[int64]  `json:"priority"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing param.Field[SiteWanUpdateParamsWanStaticAddressing] `json:"static_addressing"`
	// VLAN port number.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r SiteWanUpdateParamsWan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type SiteWanUpdateParamsWanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address param.Field[string] `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress param.Field[string] `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
}

func (r SiteWanUpdateParamsWanStaticAddressing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWanUpdateResponseEnvelope struct {
	Errors   []SiteWanUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteWanUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteWanUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteWanUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteWanUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteWanUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteWanUpdateResponseEnvelope]
type siteWanUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteWanUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteWanUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// siteWanUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SiteWanUpdateResponseEnvelopeErrors]
type siteWanUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteWanUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    siteWanUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// siteWanUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteWanUpdateResponseEnvelopeMessages]
type siteWanUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteWanUpdateResponseEnvelopeSuccess bool

const (
	SiteWanUpdateResponseEnvelopeSuccessTrue SiteWanUpdateResponseEnvelopeSuccess = true
)

type SiteWanListResponseEnvelope struct {
	Errors   []SiteWanListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteWanListResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteWanListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteWanListResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteWanListResponseEnvelopeJSON    `json:"-"`
}

// siteWanListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteWanListResponseEnvelope]
type siteWanListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteWanListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    siteWanListResponseEnvelopeErrorsJSON `json:"-"`
}

// siteWanListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteWanListResponseEnvelopeErrors]
type siteWanListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteWanListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteWanListResponseEnvelopeMessagesJSON `json:"-"`
}

// siteWanListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteWanListResponseEnvelopeMessages]
type siteWanListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteWanListResponseEnvelopeSuccess bool

const (
	SiteWanListResponseEnvelopeSuccessTrue SiteWanListResponseEnvelopeSuccess = true
)

type SiteWanDeleteResponseEnvelope struct {
	Errors   []SiteWanDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteWanDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteWanDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteWanDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteWanDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteWanDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteWanDeleteResponseEnvelope]
type siteWanDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteWanDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    siteWanDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// siteWanDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SiteWanDeleteResponseEnvelopeErrors]
type siteWanDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteWanDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    siteWanDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// siteWanDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SiteWanDeleteResponseEnvelopeMessages]
type siteWanDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteWanDeleteResponseEnvelopeSuccess bool

const (
	SiteWanDeleteResponseEnvelopeSuccessTrue SiteWanDeleteResponseEnvelopeSuccess = true
)

type SiteWanGetResponseEnvelope struct {
	Errors   []SiteWanGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SiteWanGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SiteWanGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SiteWanGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteWanGetResponseEnvelopeJSON    `json:"-"`
}

// siteWanGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteWanGetResponseEnvelope]
type siteWanGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SiteWanGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    siteWanGetResponseEnvelopeErrorsJSON `json:"-"`
}

// siteWanGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SiteWanGetResponseEnvelopeErrors]
type siteWanGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SiteWanGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    siteWanGetResponseEnvelopeMessagesJSON `json:"-"`
}

// siteWanGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SiteWanGetResponseEnvelopeMessages]
type siteWanGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteWanGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteWanGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteWanGetResponseEnvelopeSuccess bool

const (
	SiteWanGetResponseEnvelopeSuccessTrue SiteWanGetResponseEnvelopeSuccess = true
)
