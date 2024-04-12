// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AddressMapService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAddressMapService] method instead.
type AddressMapService struct {
	Options  []option.RequestOption
	Accounts *AddressMapAccountService
	IPs      *AddressMapIPService
	Zones    *AddressMapZoneService
}

// NewAddressMapService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressMapService(opts ...option.RequestOption) (r *AddressMapService) {
	r = &AddressMapService{}
	r.Options = opts
	r.Accounts = NewAddressMapAccountService(opts...)
	r.IPs = NewAddressMapIPService(opts...)
	r.Zones = NewAddressMapZoneService(opts...)
	return
}

// Create a new address map under the account.
func (r *AddressMapService) New(ctx context.Context, params AddressMapNewParams, opts ...option.RequestOption) (res *AddressMapNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all address maps owned by the account.
func (r *AddressMapService) List(ctx context.Context, query AddressMapListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AddressMap], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", query.AccountID)
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

// List all address maps owned by the account.
func (r *AddressMapService) ListAutoPaging(ctx context.Context, query AddressMapListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AddressMap] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a particular address map owned by the account. An Address Map must be
// disabled before it can be deleted.
func (r *AddressMapService) Delete(ctx context.Context, addressMapID string, params AddressMapDeleteParams, opts ...option.RequestOption) (res *AddressMapDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", params.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify properties of an address map owned by the account.
func (r *AddressMapService) Edit(ctx context.Context, addressMapID string, params AddressMapEditParams, opts ...option.RequestOption) (res *AddressMap, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", params.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show a particular address map owned by the account.
func (r *AddressMapService) Get(ctx context.Context, addressMapID string, query AddressMapGetParams, opts ...option.RequestOption) (res *AddressMapGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", query.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressMap struct {
	// Identifier
	ID string `json:"id"`
	// If set to false, then the Address Map cannot be deleted via API. This is true
	// for Cloudflare-managed maps.
	CanDelete bool `json:"can_delete"`
	// If set to false, then the IPs on the Address Map cannot be modified via the API.
	// This is true for Cloudflare-managed maps.
	CanModifyIPs bool      `json:"can_modify_ips"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// If you have legacy TLS clients which do not send the TLS server name indicator,
	// then you can specify one default SNI on the map. If Cloudflare receives a TLS
	// handshake from a client without an SNI, it will respond with the default SNI on
	// those IPs. The default SNI can be any valid zone or subdomain owned by the
	// account.
	DefaultSni string `json:"default_sni,nullable"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description string `json:"description,nullable"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled    bool           `json:"enabled,nullable"`
	ModifiedAt time.Time      `json:"modified_at" format:"date-time"`
	JSON       addressMapJSON `json:"-"`
}

// addressMapJSON contains the JSON metadata for the struct [AddressMap]
type addressMapJSON struct {
	ID           apijson.Field
	CanDelete    apijson.Field
	CanModifyIPs apijson.Field
	CreatedAt    apijson.Field
	DefaultSni   apijson.Field
	Description  apijson.Field
	Enabled      apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressMap) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapJSON) RawJSON() string {
	return r.raw
}

type AddressMapNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// If set to false, then the Address Map cannot be deleted via API. This is true
	// for Cloudflare-managed maps.
	CanDelete bool `json:"can_delete"`
	// If set to false, then the IPs on the Address Map cannot be modified via the API.
	// This is true for Cloudflare-managed maps.
	CanModifyIPs bool      `json:"can_modify_ips"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// If you have legacy TLS clients which do not send the TLS server name indicator,
	// then you can specify one default SNI on the map. If Cloudflare receives a TLS
	// handshake from a client without an SNI, it will respond with the default SNI on
	// those IPs. The default SNI can be any valid zone or subdomain owned by the
	// account.
	DefaultSni string `json:"default_sni,nullable"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description string `json:"description,nullable"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled bool `json:"enabled,nullable"`
	// The set of IPs on the Address Map.
	IPs []AddressMapNewResponseIP `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []AddressMapNewResponseMembership `json:"memberships"`
	ModifiedAt  time.Time                         `json:"modified_at" format:"date-time"`
	JSON        addressMapNewResponseJSON         `json:"-"`
}

// addressMapNewResponseJSON contains the JSON metadata for the struct
// [AddressMapNewResponse]
type addressMapNewResponseJSON struct {
	ID           apijson.Field
	CanDelete    apijson.Field
	CanModifyIPs apijson.Field
	CreatedAt    apijson.Field
	DefaultSni   apijson.Field
	Description  apijson.Field
	Enabled      apijson.Field
	IPs          apijson.Field
	Memberships  apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressMapNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseJSON) RawJSON() string {
	return r.raw
}

type AddressMapNewResponseIP struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An IPv4 or IPv6 address.
	IP   string                      `json:"ip"`
	JSON addressMapNewResponseIPJSON `json:"-"`
}

// addressMapNewResponseIPJSON contains the JSON metadata for the struct
// [AddressMapNewResponseIP]
type addressMapNewResponseIPJSON struct {
	CreatedAt   apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapNewResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseIPJSON) RawJSON() string {
	return r.raw
}

type AddressMapNewResponseMembership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifier
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind AddressMapNewResponseMembershipsKind `json:"kind"`
	JSON addressMapNewResponseMembershipJSON  `json:"-"`
}

// addressMapNewResponseMembershipJSON contains the JSON metadata for the struct
// [AddressMapNewResponseMembership]
type addressMapNewResponseMembershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapNewResponseMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseMembershipJSON) RawJSON() string {
	return r.raw
}

// The type of the membership.
type AddressMapNewResponseMembershipsKind string

const (
	AddressMapNewResponseMembershipsKindZone    AddressMapNewResponseMembershipsKind = "zone"
	AddressMapNewResponseMembershipsKindAccount AddressMapNewResponseMembershipsKind = "account"
)

func (r AddressMapNewResponseMembershipsKind) IsKnown() bool {
	switch r {
	case AddressMapNewResponseMembershipsKindZone, AddressMapNewResponseMembershipsKindAccount:
		return true
	}
	return false
}

// Union satisfied by [addressing.AddressMapDeleteResponseUnknown],
// [addressing.AddressMapDeleteResponseArray] or [shared.UnionString].
type AddressMapDeleteResponseUnion interface {
	ImplementsAddressingAddressMapDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressMapDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressMapDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressMapDeleteResponseArray []interface{}

func (r AddressMapDeleteResponseArray) ImplementsAddressingAddressMapDeleteResponseUnion() {}

type AddressMapGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// If set to false, then the Address Map cannot be deleted via API. This is true
	// for Cloudflare-managed maps.
	CanDelete bool `json:"can_delete"`
	// If set to false, then the IPs on the Address Map cannot be modified via the API.
	// This is true for Cloudflare-managed maps.
	CanModifyIPs bool      `json:"can_modify_ips"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// If you have legacy TLS clients which do not send the TLS server name indicator,
	// then you can specify one default SNI on the map. If Cloudflare receives a TLS
	// handshake from a client without an SNI, it will respond with the default SNI on
	// those IPs. The default SNI can be any valid zone or subdomain owned by the
	// account.
	DefaultSni string `json:"default_sni,nullable"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description string `json:"description,nullable"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled bool `json:"enabled,nullable"`
	// The set of IPs on the Address Map.
	IPs []AddressMapGetResponseIP `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []AddressMapGetResponseMembership `json:"memberships"`
	ModifiedAt  time.Time                         `json:"modified_at" format:"date-time"`
	JSON        addressMapGetResponseJSON         `json:"-"`
}

// addressMapGetResponseJSON contains the JSON metadata for the struct
// [AddressMapGetResponse]
type addressMapGetResponseJSON struct {
	ID           apijson.Field
	CanDelete    apijson.Field
	CanModifyIPs apijson.Field
	CreatedAt    apijson.Field
	DefaultSni   apijson.Field
	Description  apijson.Field
	Enabled      apijson.Field
	IPs          apijson.Field
	Memberships  apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressMapGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseJSON) RawJSON() string {
	return r.raw
}

type AddressMapGetResponseIP struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An IPv4 or IPv6 address.
	IP   string                      `json:"ip"`
	JSON addressMapGetResponseIPJSON `json:"-"`
}

// addressMapGetResponseIPJSON contains the JSON metadata for the struct
// [AddressMapGetResponseIP]
type addressMapGetResponseIPJSON struct {
	CreatedAt   apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapGetResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseIPJSON) RawJSON() string {
	return r.raw
}

type AddressMapGetResponseMembership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifier
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind AddressMapGetResponseMembershipsKind `json:"kind"`
	JSON addressMapGetResponseMembershipJSON  `json:"-"`
}

// addressMapGetResponseMembershipJSON contains the JSON metadata for the struct
// [AddressMapGetResponseMembership]
type addressMapGetResponseMembershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapGetResponseMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseMembershipJSON) RawJSON() string {
	return r.raw
}

// The type of the membership.
type AddressMapGetResponseMembershipsKind string

const (
	AddressMapGetResponseMembershipsKindZone    AddressMapGetResponseMembershipsKind = "zone"
	AddressMapGetResponseMembershipsKindAccount AddressMapGetResponseMembershipsKind = "account"
)

func (r AddressMapGetResponseMembershipsKind) IsKnown() bool {
	switch r {
	case AddressMapGetResponseMembershipsKindZone, AddressMapGetResponseMembershipsKindAccount:
		return true
	}
	return false
}

type AddressMapNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description param.Field[string] `json:"description"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AddressMapNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressMapNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AddressMapNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success AddressMapNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressMapNewResponseEnvelopeJSON    `json:"-"`
}

// addressMapNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapNewResponseEnvelope]
type addressMapNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapNewResponseEnvelopeSuccess bool

const (
	AddressMapNewResponseEnvelopeSuccessTrue AddressMapNewResponseEnvelopeSuccess = true
)

func (r AddressMapNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressMapDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r AddressMapDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AddressMapDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   AddressMapDeleteResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressMapDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressMapDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressMapDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapDeleteResponseEnvelope]
type addressMapDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapDeleteResponseEnvelopeSuccess bool

const (
	AddressMapDeleteResponseEnvelopeSuccessTrue AddressMapDeleteResponseEnvelopeSuccess = true
)

func (r AddressMapDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       addressMapDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressMapDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AddressMapDeleteResponseEnvelopeResultInfo]
type addressMapDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// If you have legacy TLS clients which do not send the TLS server name indicator,
	// then you can specify one default SNI on the map. If Cloudflare receives a TLS
	// handshake from a client without an SNI, it will respond with the default SNI on
	// those IPs. The default SNI can be any valid zone or subdomain owned by the
	// account.
	DefaultSni param.Field[string] `json:"default_sni"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description param.Field[string] `json:"description"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AddressMapEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressMapEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AddressMap            `json:"result,required"`
	// Whether the API call was successful
	Success AddressMapEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressMapEditResponseEnvelopeJSON    `json:"-"`
}

// addressMapEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapEditResponseEnvelope]
type addressMapEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapEditResponseEnvelopeSuccess bool

const (
	AddressMapEditResponseEnvelopeSuccessTrue AddressMapEditResponseEnvelopeSuccess = true
)

func (r AddressMapEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressMapGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AddressMapGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success AddressMapGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressMapGetResponseEnvelopeJSON    `json:"-"`
}

// addressMapGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapGetResponseEnvelope]
type addressMapGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapGetResponseEnvelopeSuccess bool

const (
	AddressMapGetResponseEnvelopeSuccessTrue AddressMapGetResponseEnvelopeSuccess = true
)

func (r AddressMapGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
