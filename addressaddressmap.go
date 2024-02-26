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

// AddressAddressMapService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressAddressMapService] method
// instead.
type AddressAddressMapService struct {
	Options  []option.RequestOption
	Accounts *AddressAddressMapAccountService
	IPs      *AddressAddressMapIPService
	Zones    *AddressAddressMapZoneService
}

// NewAddressAddressMapService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressAddressMapService(opts ...option.RequestOption) (r *AddressAddressMapService) {
	r = &AddressAddressMapService{}
	r.Options = opts
	r.Accounts = NewAddressAddressMapAccountService(opts...)
	r.IPs = NewAddressAddressMapIPService(opts...)
	r.Zones = NewAddressAddressMapZoneService(opts...)
	return
}

// Create a new address map under the account.
func (r *AddressAddressMapService) New(ctx context.Context, params AddressAddressMapNewParams, opts ...option.RequestOption) (res *AddressAddressMapNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all address maps owned by the account.
func (r *AddressAddressMapService) List(ctx context.Context, query AddressAddressMapListParams, opts ...option.RequestOption) (res *[]AddressAddressMapListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapListResponseEnvelope
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
func (r *AddressAddressMapService) Delete(ctx context.Context, addressMapID string, body AddressAddressMapDeleteParams, opts ...option.RequestOption) (res *AddressAddressMapDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", body.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify properties of an address map owned by the account.
func (r *AddressAddressMapService) Edit(ctx context.Context, addressMapID string, params AddressAddressMapEditParams, opts ...option.RequestOption) (res *AddressAddressMapEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", params.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show a particular address map owned by the account.
func (r *AddressAddressMapService) Get(ctx context.Context, addressMapID string, query AddressAddressMapGetParams, opts ...option.RequestOption) (res *AddressAddressMapGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", query.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressAddressMapNewResponse struct {
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
	IPs []AddressAddressMapNewResponseIP `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []AddressAddressMapNewResponseMembership `json:"memberships"`
	ModifiedAt  time.Time                                `json:"modified_at" format:"date-time"`
	JSON        addressAddressMapNewResponseJSON         `json:"-"`
}

// addressAddressMapNewResponseJSON contains the JSON metadata for the struct
// [AddressAddressMapNewResponse]
type addressAddressMapNewResponseJSON struct {
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

func (r *AddressAddressMapNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapNewResponseIP struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An IPv4 or IPv6 address.
	IP   string                             `json:"ip"`
	JSON addressAddressMapNewResponseIPJSON `json:"-"`
}

// addressAddressMapNewResponseIPJSON contains the JSON metadata for the struct
// [AddressAddressMapNewResponseIP]
type addressAddressMapNewResponseIPJSON struct {
	CreatedAt   apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapNewResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapNewResponseMembership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifier
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind AddressAddressMapNewResponseMembershipsKind `json:"kind"`
	JSON addressAddressMapNewResponseMembershipJSON  `json:"-"`
}

// addressAddressMapNewResponseMembershipJSON contains the JSON metadata for the
// struct [AddressAddressMapNewResponseMembership]
type addressAddressMapNewResponseMembershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapNewResponseMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the membership.
type AddressAddressMapNewResponseMembershipsKind string

const (
	AddressAddressMapNewResponseMembershipsKindZone    AddressAddressMapNewResponseMembershipsKind = "zone"
	AddressAddressMapNewResponseMembershipsKindAccount AddressAddressMapNewResponseMembershipsKind = "account"
)

type AddressAddressMapListResponse struct {
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
	Enabled    bool                              `json:"enabled,nullable"`
	ModifiedAt time.Time                         `json:"modified_at" format:"date-time"`
	JSON       addressAddressMapListResponseJSON `json:"-"`
}

// addressAddressMapListResponseJSON contains the JSON metadata for the struct
// [AddressAddressMapListResponse]
type addressAddressMapListResponseJSON struct {
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

func (r *AddressAddressMapListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AddressAddressMapDeleteResponseUnknown],
// [AddressAddressMapDeleteResponseArray] or [shared.UnionString].
type AddressAddressMapDeleteResponse interface {
	ImplementsAddressAddressMapDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressAddressMapDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressAddressMapDeleteResponseArray []interface{}

func (r AddressAddressMapDeleteResponseArray) ImplementsAddressAddressMapDeleteResponse() {}

type AddressAddressMapEditResponse struct {
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
	Enabled    bool                              `json:"enabled,nullable"`
	ModifiedAt time.Time                         `json:"modified_at" format:"date-time"`
	JSON       addressAddressMapEditResponseJSON `json:"-"`
}

// addressAddressMapEditResponseJSON contains the JSON metadata for the struct
// [AddressAddressMapEditResponse]
type addressAddressMapEditResponseJSON struct {
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

func (r *AddressAddressMapEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapGetResponse struct {
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
	IPs []AddressAddressMapGetResponseIP `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []AddressAddressMapGetResponseMembership `json:"memberships"`
	ModifiedAt  time.Time                                `json:"modified_at" format:"date-time"`
	JSON        addressAddressMapGetResponseJSON         `json:"-"`
}

// addressAddressMapGetResponseJSON contains the JSON metadata for the struct
// [AddressAddressMapGetResponse]
type addressAddressMapGetResponseJSON struct {
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

func (r *AddressAddressMapGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapGetResponseIP struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An IPv4 or IPv6 address.
	IP   string                             `json:"ip"`
	JSON addressAddressMapGetResponseIPJSON `json:"-"`
}

// addressAddressMapGetResponseIPJSON contains the JSON metadata for the struct
// [AddressAddressMapGetResponseIP]
type addressAddressMapGetResponseIPJSON struct {
	CreatedAt   apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapGetResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapGetResponseMembership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifier
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind AddressAddressMapGetResponseMembershipsKind `json:"kind"`
	JSON addressAddressMapGetResponseMembershipJSON  `json:"-"`
}

// addressAddressMapGetResponseMembershipJSON contains the JSON metadata for the
// struct [AddressAddressMapGetResponseMembership]
type addressAddressMapGetResponseMembershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapGetResponseMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the membership.
type AddressAddressMapGetResponseMembershipsKind string

const (
	AddressAddressMapGetResponseMembershipsKindZone    AddressAddressMapGetResponseMembershipsKind = "zone"
	AddressAddressMapGetResponseMembershipsKindAccount AddressAddressMapGetResponseMembershipsKind = "account"
)

type AddressAddressMapNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description param.Field[string] `json:"description"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AddressAddressMapNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressAddressMapNewResponseEnvelope struct {
	Errors   []AddressAddressMapNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressAddressMapNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressAddressMapNewResponseEnvelopeJSON    `json:"-"`
}

// addressAddressMapNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressAddressMapNewResponseEnvelope]
type addressAddressMapNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    addressAddressMapNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressAddressMapNewResponseEnvelopeErrors]
type addressAddressMapNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    addressAddressMapNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressAddressMapNewResponseEnvelopeMessages]
type addressAddressMapNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapNewResponseEnvelopeSuccess bool

const (
	AddressAddressMapNewResponseEnvelopeSuccessTrue AddressAddressMapNewResponseEnvelopeSuccess = true
)

type AddressAddressMapListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressAddressMapListResponseEnvelope struct {
	Errors   []AddressAddressMapListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressAddressMapListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressAddressMapListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressAddressMapListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressAddressMapListResponseEnvelopeJSON       `json:"-"`
}

// addressAddressMapListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressAddressMapListResponseEnvelope]
type addressAddressMapListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    addressAddressMapListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressAddressMapListResponseEnvelopeErrors]
type addressAddressMapListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    addressAddressMapListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressAddressMapListResponseEnvelopeMessages]
type addressAddressMapListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapListResponseEnvelopeSuccess bool

const (
	AddressAddressMapListResponseEnvelopeSuccessTrue AddressAddressMapListResponseEnvelopeSuccess = true
)

type AddressAddressMapListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       addressAddressMapListResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressAddressMapListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AddressAddressMapListResponseEnvelopeResultInfo]
type addressAddressMapListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressAddressMapDeleteResponseEnvelope struct {
	Errors   []AddressAddressMapDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressAddressMapDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressAddressMapDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressAddressMapDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressAddressMapDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressAddressMapDeleteResponseEnvelope]
type addressAddressMapDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    addressAddressMapDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressAddressMapDeleteResponseEnvelopeErrors]
type addressAddressMapDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    addressAddressMapDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressAddressMapDeleteResponseEnvelopeMessages]
type addressAddressMapDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapDeleteResponseEnvelopeSuccess bool

const (
	AddressAddressMapDeleteResponseEnvelopeSuccessTrue AddressAddressMapDeleteResponseEnvelopeSuccess = true
)

type AddressAddressMapDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       addressAddressMapDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressAddressMapDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AddressAddressMapDeleteResponseEnvelopeResultInfo]
type addressAddressMapDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapEditParams struct {
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

func (r AddressAddressMapEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressAddressMapEditResponseEnvelope struct {
	Errors   []AddressAddressMapEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapEditResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressAddressMapEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressAddressMapEditResponseEnvelopeJSON    `json:"-"`
}

// addressAddressMapEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressAddressMapEditResponseEnvelope]
type addressAddressMapEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    addressAddressMapEditResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressAddressMapEditResponseEnvelopeErrors]
type addressAddressMapEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    addressAddressMapEditResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressAddressMapEditResponseEnvelopeMessages]
type addressAddressMapEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapEditResponseEnvelopeSuccess bool

const (
	AddressAddressMapEditResponseEnvelopeSuccessTrue AddressAddressMapEditResponseEnvelopeSuccess = true
)

type AddressAddressMapGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressAddressMapGetResponseEnvelope struct {
	Errors   []AddressAddressMapGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressAddressMapGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressAddressMapGetResponseEnvelopeJSON    `json:"-"`
}

// addressAddressMapGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressAddressMapGetResponseEnvelope]
type addressAddressMapGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    addressAddressMapGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressAddressMapGetResponseEnvelopeErrors]
type addressAddressMapGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    addressAddressMapGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressAddressMapGetResponseEnvelopeMessages]
type addressAddressMapGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapGetResponseEnvelopeSuccess bool

const (
	AddressAddressMapGetResponseEnvelopeSuccessTrue AddressAddressMapGetResponseEnvelopeSuccess = true
)
