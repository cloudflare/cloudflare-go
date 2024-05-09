// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SiteWANService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteWANService] method instead.
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
func (r *SiteWANService) New(ctx context.Context, siteID string, params SiteWANNewParams, opts ...option.RequestOption) (res *[]WAN, err error) {
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
func (r *SiteWANService) Update(ctx context.Context, siteID string, wanID string, params SiteWANUpdateParams, opts ...option.RequestOption) (res *WAN, err error) {
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
func (r *SiteWANService) List(ctx context.Context, siteID string, query SiteWANListParams, opts ...option.RequestOption) (res *pagination.SinglePage[WAN], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans", query.AccountID, siteID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists WANs associated with an account and site.
func (r *SiteWANService) ListAutoPaging(ctx context.Context, siteID string, query SiteWANListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[WAN] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, siteID, query, opts...))
}

// Remove a specific WAN.
func (r *SiteWANService) Delete(ctx context.Context, siteID string, wanID string, body SiteWANDeleteParams, opts ...option.RequestOption) (res *WAN, err error) {
	opts = append(r.Options[:], opts...)
	var env SiteWANDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", body.AccountID, siteID, wanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific WAN.
func (r *SiteWANService) Get(ctx context.Context, siteID string, wanID string, query SiteWANGetParams, opts ...option.RequestOption) (res *WAN, err error) {
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

type WAN struct {
	// Identifier
	ID       string `json:"id"`
	Name     string `json:"name"`
	Physport int64  `json:"physport"`
	// Priority of WAN for traffic loadbalancing.
	Priority int64 `json:"priority"`
	// Identifier
	SiteID string `json:"site_id"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing WANStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64   `json:"vlan_tag"`
	JSON    wanJSON `json:"-"`
}

// wanJSON contains the JSON metadata for the struct [WAN]
type wanJSON struct {
	ID               apijson.Field
	Name             apijson.Field
	Physport         apijson.Field
	Priority         apijson.Field
	SiteID           apijson.Field
	StaticAddressing apijson.Field
	VlanTag          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WAN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wanJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type WANStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                  `json:"secondary_address"`
	JSON             wanStaticAddressingJSON `json:"-"`
}

// wanStaticAddressingJSON contains the JSON metadata for the struct
// [WANStaticAddressing]
type wanStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WANStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wanStaticAddressingJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type WANStaticAddressingParam struct {
	// A valid CIDR notation representing an IP range.
	Address param.Field[string] `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress param.Field[string] `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
}

func (r WANStaticAddressingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWANNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Physport  param.Field[int64]  `json:"physport,required"`
	// VLAN port number.
	VlanTag  param.Field[int64]  `json:"vlan_tag,required"`
	Name     param.Field[string] `json:"name"`
	Priority param.Field[int64]  `json:"priority"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing param.Field[WANStaticAddressingParam] `json:"static_addressing"`
}

func (r SiteWANNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWANNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []WAN                 `json:"result,required"`
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
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `json:"name"`
	Physport  param.Field[int64]  `json:"physport"`
	Priority  param.Field[int64]  `json:"priority"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing param.Field[WANStaticAddressingParam] `json:"static_addressing"`
	// VLAN port number.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r SiteWANUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteWANUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   WAN                   `json:"result,required"`
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

type SiteWANDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteWANDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   WAN                   `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   WAN                   `json:"result,required"`
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
