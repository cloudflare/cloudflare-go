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

// AddressingAddressMapService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressingAddressMapService]
// method instead.
type AddressingAddressMapService struct {
	Options  []option.RequestOption
	Accounts *AddressingAddressMapAccountService
	IPs      *AddressingAddressMapIPService
	Zones    *AddressingAddressMapZoneService
}

// NewAddressingAddressMapService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressingAddressMapService(opts ...option.RequestOption) (r *AddressingAddressMapService) {
	r = &AddressingAddressMapService{}
	r.Options = opts
	r.Accounts = NewAddressingAddressMapAccountService(opts...)
	r.IPs = NewAddressingAddressMapIPService(opts...)
	r.Zones = NewAddressingAddressMapZoneService(opts...)
	return
}

// Create a new address map under the account.
func (r *AddressingAddressMapService) New(ctx context.Context, params AddressingAddressMapNewParams, opts ...option.RequestOption) (res *AddressingAddressMapNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all address maps owned by the account.
func (r *AddressingAddressMapService) List(ctx context.Context, query AddressingAddressMapListParams, opts ...option.RequestOption) (res *[]AddressingAddressMapListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a particular address map owned by the account. An Address Map must be
// disabled before it can be deleted.
func (r *AddressingAddressMapService) Delete(ctx context.Context, addressMapID string, body AddressingAddressMapDeleteParams, opts ...option.RequestOption) (res *AddressingAddressMapDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", body.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify properties of an address map owned by the account.
func (r *AddressingAddressMapService) Edit(ctx context.Context, addressMapID string, params AddressingAddressMapEditParams, opts ...option.RequestOption) (res *AddressingAddressMapEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", params.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show a particular address map owned by the account.
func (r *AddressingAddressMapService) Get(ctx context.Context, addressMapID string, query AddressingAddressMapGetParams, opts ...option.RequestOption) (res *AddressingAddressMapGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", query.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingAddressMapNewResponse struct {
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
	IPs []AddressingAddressMapNewResponseIP `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []AddressingAddressMapNewResponseMembership `json:"memberships"`
	ModifiedAt  time.Time                                   `json:"modified_at" format:"date-time"`
	JSON        addressingAddressMapNewResponseJSON         `json:"-"`
}

// addressingAddressMapNewResponseJSON contains the JSON metadata for the struct
// [AddressingAddressMapNewResponse]
type addressingAddressMapNewResponseJSON struct {
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

func (r *AddressingAddressMapNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapNewResponseJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapNewResponseIP struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An IPv4 or IPv6 address.
	IP   string                                `json:"ip"`
	JSON addressingAddressMapNewResponseIPJSON `json:"-"`
}

// addressingAddressMapNewResponseIPJSON contains the JSON metadata for the struct
// [AddressingAddressMapNewResponseIP]
type addressingAddressMapNewResponseIPJSON struct {
	CreatedAt   apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapNewResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapNewResponseIPJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapNewResponseMembership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifier
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind AddressingAddressMapNewResponseMembershipsKind `json:"kind"`
	JSON addressingAddressMapNewResponseMembershipJSON  `json:"-"`
}

// addressingAddressMapNewResponseMembershipJSON contains the JSON metadata for the
// struct [AddressingAddressMapNewResponseMembership]
type addressingAddressMapNewResponseMembershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapNewResponseMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapNewResponseMembershipJSON) RawJSON() string {
	return r.raw
}

// The type of the membership.
type AddressingAddressMapNewResponseMembershipsKind string

const (
	AddressingAddressMapNewResponseMembershipsKindZone    AddressingAddressMapNewResponseMembershipsKind = "zone"
	AddressingAddressMapNewResponseMembershipsKindAccount AddressingAddressMapNewResponseMembershipsKind = "account"
)

type AddressingAddressMapListResponse struct {
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
	Enabled    bool                                 `json:"enabled,nullable"`
	ModifiedAt time.Time                            `json:"modified_at" format:"date-time"`
	JSON       addressingAddressMapListResponseJSON `json:"-"`
}

// addressingAddressMapListResponseJSON contains the JSON metadata for the struct
// [AddressingAddressMapListResponse]
type addressingAddressMapListResponseJSON struct {
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

func (r *AddressingAddressMapListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapListResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [AddressingAddressMapDeleteResponseUnknown],
// [AddressingAddressMapDeleteResponseArray] or [shared.UnionString].
type AddressingAddressMapDeleteResponse interface {
	ImplementsAddressingAddressMapDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingAddressMapDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingAddressMapDeleteResponseArray []interface{}

func (r AddressingAddressMapDeleteResponseArray) ImplementsAddressingAddressMapDeleteResponse() {}

type AddressingAddressMapEditResponse struct {
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
	Enabled    bool                                 `json:"enabled,nullable"`
	ModifiedAt time.Time                            `json:"modified_at" format:"date-time"`
	JSON       addressingAddressMapEditResponseJSON `json:"-"`
}

// addressingAddressMapEditResponseJSON contains the JSON metadata for the struct
// [AddressingAddressMapEditResponse]
type addressingAddressMapEditResponseJSON struct {
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

func (r *AddressingAddressMapEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapEditResponseJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapGetResponse struct {
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
	IPs []AddressingAddressMapGetResponseIP `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []AddressingAddressMapGetResponseMembership `json:"memberships"`
	ModifiedAt  time.Time                                   `json:"modified_at" format:"date-time"`
	JSON        addressingAddressMapGetResponseJSON         `json:"-"`
}

// addressingAddressMapGetResponseJSON contains the JSON metadata for the struct
// [AddressingAddressMapGetResponse]
type addressingAddressMapGetResponseJSON struct {
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

func (r *AddressingAddressMapGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapGetResponseJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapGetResponseIP struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An IPv4 or IPv6 address.
	IP   string                                `json:"ip"`
	JSON addressingAddressMapGetResponseIPJSON `json:"-"`
}

// addressingAddressMapGetResponseIPJSON contains the JSON metadata for the struct
// [AddressingAddressMapGetResponseIP]
type addressingAddressMapGetResponseIPJSON struct {
	CreatedAt   apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapGetResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapGetResponseIPJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapGetResponseMembership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifier
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind AddressingAddressMapGetResponseMembershipsKind `json:"kind"`
	JSON addressingAddressMapGetResponseMembershipJSON  `json:"-"`
}

// addressingAddressMapGetResponseMembershipJSON contains the JSON metadata for the
// struct [AddressingAddressMapGetResponseMembership]
type addressingAddressMapGetResponseMembershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapGetResponseMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapGetResponseMembershipJSON) RawJSON() string {
	return r.raw
}

// The type of the membership.
type AddressingAddressMapGetResponseMembershipsKind string

const (
	AddressingAddressMapGetResponseMembershipsKindZone    AddressingAddressMapGetResponseMembershipsKind = "zone"
	AddressingAddressMapGetResponseMembershipsKindAccount AddressingAddressMapGetResponseMembershipsKind = "account"
)

type AddressingAddressMapNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description param.Field[string] `json:"description"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AddressingAddressMapNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingAddressMapNewResponseEnvelope struct {
	Errors   []AddressingAddressMapNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingAddressMapNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingAddressMapNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingAddressMapNewResponseEnvelopeJSON    `json:"-"`
}

// addressingAddressMapNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressingAddressMapNewResponseEnvelope]
type addressingAddressMapNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    addressingAddressMapNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressingAddressMapNewResponseEnvelopeErrors]
type addressingAddressMapNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    addressingAddressMapNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressingAddressMapNewResponseEnvelopeMessages]
type addressingAddressMapNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapNewResponseEnvelopeSuccess bool

const (
	AddressingAddressMapNewResponseEnvelopeSuccessTrue AddressingAddressMapNewResponseEnvelopeSuccess = true
)

type AddressingAddressMapListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingAddressMapListResponseEnvelope struct {
	Errors   []AddressingAddressMapListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressingAddressMapListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingAddressMapListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingAddressMapListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingAddressMapListResponseEnvelopeJSON       `json:"-"`
}

// addressingAddressMapListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressingAddressMapListResponseEnvelope]
type addressingAddressMapListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    addressingAddressMapListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressingAddressMapListResponseEnvelopeErrors]
type addressingAddressMapListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressingAddressMapListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressingAddressMapListResponseEnvelopeMessages]
type addressingAddressMapListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapListResponseEnvelopeSuccess bool

const (
	AddressingAddressMapListResponseEnvelopeSuccessTrue AddressingAddressMapListResponseEnvelopeSuccess = true
)

type AddressingAddressMapListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       addressingAddressMapListResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingAddressMapListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressingAddressMapListResponseEnvelopeResultInfo]
type addressingAddressMapListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingAddressMapDeleteResponseEnvelope struct {
	Errors   []AddressingAddressMapDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingAddressMapDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingAddressMapDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingAddressMapDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingAddressMapDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressingAddressMapDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingAddressMapDeleteResponseEnvelope]
type addressingAddressMapDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressingAddressMapDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressingAddressMapDeleteResponseEnvelopeErrors]
type addressingAddressMapDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressingAddressMapDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingAddressMapDeleteResponseEnvelopeMessages]
type addressingAddressMapDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapDeleteResponseEnvelopeSuccess bool

const (
	AddressingAddressMapDeleteResponseEnvelopeSuccessTrue AddressingAddressMapDeleteResponseEnvelopeSuccess = true
)

type AddressingAddressMapDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       addressingAddressMapDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingAddressMapDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressingAddressMapDeleteResponseEnvelopeResultInfo]
type addressingAddressMapDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapEditParams struct {
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

func (r AddressingAddressMapEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingAddressMapEditResponseEnvelope struct {
	Errors   []AddressingAddressMapEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapEditResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingAddressMapEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingAddressMapEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingAddressMapEditResponseEnvelopeJSON    `json:"-"`
}

// addressingAddressMapEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressingAddressMapEditResponseEnvelope]
type addressingAddressMapEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    addressingAddressMapEditResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressingAddressMapEditResponseEnvelopeErrors]
type addressingAddressMapEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressingAddressMapEditResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressingAddressMapEditResponseEnvelopeMessages]
type addressingAddressMapEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapEditResponseEnvelopeSuccess bool

const (
	AddressingAddressMapEditResponseEnvelopeSuccessTrue AddressingAddressMapEditResponseEnvelopeSuccess = true
)

type AddressingAddressMapGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingAddressMapGetResponseEnvelope struct {
	Errors   []AddressingAddressMapGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingAddressMapGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingAddressMapGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingAddressMapGetResponseEnvelopeJSON    `json:"-"`
}

// addressingAddressMapGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressingAddressMapGetResponseEnvelope]
type addressingAddressMapGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    addressingAddressMapGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressingAddressMapGetResponseEnvelopeErrors]
type addressingAddressMapGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    addressingAddressMapGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressingAddressMapGetResponseEnvelopeMessages]
type addressingAddressMapGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapGetResponseEnvelopeSuccess bool

const (
	AddressingAddressMapGetResponseEnvelopeSuccessTrue AddressingAddressMapGetResponseEnvelopeSuccess = true
)
