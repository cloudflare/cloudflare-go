// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAddressAddressMapService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAddressAddressMapService] method instead.
type AccountAddressAddressMapService struct {
	Options  []option.RequestOption
	Accounts *AccountAddressAddressMapAccountService
	IPs      *AccountAddressAddressMapIPService
	Zones    *AccountAddressAddressMapZoneService
}

// NewAccountAddressAddressMapService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressAddressMapService(opts ...option.RequestOption) (r *AccountAddressAddressMapService) {
	r = &AccountAddressAddressMapService{}
	r.Options = opts
	r.Accounts = NewAccountAddressAddressMapAccountService(opts...)
	r.IPs = NewAccountAddressAddressMapIPService(opts...)
	r.Zones = NewAccountAddressAddressMapZoneService(opts...)
	return
}

// Create a new address map under the account.
func (r *AccountAddressAddressMapService) New(ctx context.Context, accountIdentifier string, body AccountAddressAddressMapNewParams, opts ...option.RequestOption) (res *AccountAddressAddressMapNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Show a particular address map owned by the account.
func (r *AccountAddressAddressMapService) Get(ctx context.Context, accountIdentifier string, addressMapIdentifier string, opts ...option.RequestOption) (res *AccountAddressAddressMapGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", accountIdentifier, addressMapIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify properties of an address map owned by the account.
func (r *AccountAddressAddressMapService) Update(ctx context.Context, accountIdentifier string, addressMapIdentifier string, body AccountAddressAddressMapUpdateParams, opts ...option.RequestOption) (res *AccountAddressAddressMapUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", accountIdentifier, addressMapIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all address maps owned by the account.
func (r *AccountAddressAddressMapService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountAddressAddressMapListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a particular address map owned by the account. An Address Map must be
// disabled before it can be deleted.
func (r *AccountAddressAddressMapService) Delete(ctx context.Context, accountIdentifier string, addressMapIdentifier string, opts ...option.RequestOption) (res *AccountAddressAddressMapDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", accountIdentifier, addressMapIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountAddressAddressMapNewResponse struct {
	Errors   []AccountAddressAddressMapNewResponseError   `json:"errors"`
	Messages []AccountAddressAddressMapNewResponseMessage `json:"messages"`
	Result   AccountAddressAddressMapNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressAddressMapNewResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapNewResponseJSON    `json:"-"`
}

// accountAddressAddressMapNewResponseJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapNewResponse]
type accountAddressAddressMapNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapNewResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountAddressAddressMapNewResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapNewResponseError]
type accountAddressAddressMapNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapNewResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountAddressAddressMapNewResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapNewResponseMessageJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapNewResponseMessage]
type accountAddressAddressMapNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapNewResponseResult struct {
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
	IPs []AccountAddressAddressMapNewResponseResultIP `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []AccountAddressAddressMapNewResponseResultMembership `json:"memberships"`
	ModifiedAt  time.Time                                             `json:"modified_at" format:"date-time"`
	JSON        accountAddressAddressMapNewResponseResultJSON         `json:"-"`
}

// accountAddressAddressMapNewResponseResultJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapNewResponseResult]
type accountAddressAddressMapNewResponseResultJSON struct {
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

func (r *AccountAddressAddressMapNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapNewResponseResultIP struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An IPv4 or IPv6 address.
	IP   string                                          `json:"ip"`
	JSON accountAddressAddressMapNewResponseResultIPJSON `json:"-"`
}

// accountAddressAddressMapNewResponseResultIPJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapNewResponseResultIP]
type accountAddressAddressMapNewResponseResultIPJSON struct {
	CreatedAt   apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapNewResponseResultIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapNewResponseResultMembership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifier
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind AccountAddressAddressMapNewResponseResultMembershipsKind `json:"kind"`
	JSON accountAddressAddressMapNewResponseResultMembershipJSON  `json:"-"`
}

// accountAddressAddressMapNewResponseResultMembershipJSON contains the JSON
// metadata for the struct [AccountAddressAddressMapNewResponseResultMembership]
type accountAddressAddressMapNewResponseResultMembershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapNewResponseResultMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the membership.
type AccountAddressAddressMapNewResponseResultMembershipsKind string

const (
	AccountAddressAddressMapNewResponseResultMembershipsKindZone    AccountAddressAddressMapNewResponseResultMembershipsKind = "zone"
	AccountAddressAddressMapNewResponseResultMembershipsKindAccount AccountAddressAddressMapNewResponseResultMembershipsKind = "account"
)

// Whether the API call was successful
type AccountAddressAddressMapNewResponseSuccess bool

const (
	AccountAddressAddressMapNewResponseSuccessTrue AccountAddressAddressMapNewResponseSuccess = true
)

type AccountAddressAddressMapGetResponse struct {
	Errors   []AccountAddressAddressMapGetResponseError   `json:"errors"`
	Messages []AccountAddressAddressMapGetResponseMessage `json:"messages"`
	Result   AccountAddressAddressMapGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressAddressMapGetResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapGetResponseJSON    `json:"-"`
}

// accountAddressAddressMapGetResponseJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapGetResponse]
type accountAddressAddressMapGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapGetResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountAddressAddressMapGetResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapGetResponseError]
type accountAddressAddressMapGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapGetResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountAddressAddressMapGetResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapGetResponseMessage]
type accountAddressAddressMapGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapGetResponseResult struct {
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
	IPs []AccountAddressAddressMapGetResponseResultIP `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []AccountAddressAddressMapGetResponseResultMembership `json:"memberships"`
	ModifiedAt  time.Time                                             `json:"modified_at" format:"date-time"`
	JSON        accountAddressAddressMapGetResponseResultJSON         `json:"-"`
}

// accountAddressAddressMapGetResponseResultJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapGetResponseResult]
type accountAddressAddressMapGetResponseResultJSON struct {
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

func (r *AccountAddressAddressMapGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapGetResponseResultIP struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An IPv4 or IPv6 address.
	IP   string                                          `json:"ip"`
	JSON accountAddressAddressMapGetResponseResultIPJSON `json:"-"`
}

// accountAddressAddressMapGetResponseResultIPJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapGetResponseResultIP]
type accountAddressAddressMapGetResponseResultIPJSON struct {
	CreatedAt   apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapGetResponseResultIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapGetResponseResultMembership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Identifier
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind AccountAddressAddressMapGetResponseResultMembershipsKind `json:"kind"`
	JSON accountAddressAddressMapGetResponseResultMembershipJSON  `json:"-"`
}

// accountAddressAddressMapGetResponseResultMembershipJSON contains the JSON
// metadata for the struct [AccountAddressAddressMapGetResponseResultMembership]
type accountAddressAddressMapGetResponseResultMembershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapGetResponseResultMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the membership.
type AccountAddressAddressMapGetResponseResultMembershipsKind string

const (
	AccountAddressAddressMapGetResponseResultMembershipsKindZone    AccountAddressAddressMapGetResponseResultMembershipsKind = "zone"
	AccountAddressAddressMapGetResponseResultMembershipsKindAccount AccountAddressAddressMapGetResponseResultMembershipsKind = "account"
)

// Whether the API call was successful
type AccountAddressAddressMapGetResponseSuccess bool

const (
	AccountAddressAddressMapGetResponseSuccessTrue AccountAddressAddressMapGetResponseSuccess = true
)

type AccountAddressAddressMapUpdateResponse struct {
	Errors   []AccountAddressAddressMapUpdateResponseError   `json:"errors"`
	Messages []AccountAddressAddressMapUpdateResponseMessage `json:"messages"`
	Result   AccountAddressAddressMapUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressAddressMapUpdateResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapUpdateResponseJSON    `json:"-"`
}

// accountAddressAddressMapUpdateResponseJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapUpdateResponse]
type accountAddressAddressMapUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountAddressAddressMapUpdateResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapUpdateResponseError]
type accountAddressAddressMapUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountAddressAddressMapUpdateResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapUpdateResponseMessage]
type accountAddressAddressMapUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapUpdateResponseResult struct {
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
	Enabled    bool                                             `json:"enabled,nullable"`
	ModifiedAt time.Time                                        `json:"modified_at" format:"date-time"`
	JSON       accountAddressAddressMapUpdateResponseResultJSON `json:"-"`
}

// accountAddressAddressMapUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapUpdateResponseResult]
type accountAddressAddressMapUpdateResponseResultJSON struct {
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

func (r *AccountAddressAddressMapUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressAddressMapUpdateResponseSuccess bool

const (
	AccountAddressAddressMapUpdateResponseSuccessTrue AccountAddressAddressMapUpdateResponseSuccess = true
)

type AccountAddressAddressMapListResponse struct {
	Errors     []AccountAddressAddressMapListResponseError    `json:"errors"`
	Messages   []AccountAddressAddressMapListResponseMessage  `json:"messages"`
	Result     []AccountAddressAddressMapListResponseResult   `json:"result"`
	ResultInfo AccountAddressAddressMapListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressAddressMapListResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapListResponseJSON    `json:"-"`
}

// accountAddressAddressMapListResponseJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapListResponse]
type accountAddressAddressMapListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapListResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountAddressAddressMapListResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapListResponseErrorJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapListResponseError]
type accountAddressAddressMapListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapListResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountAddressAddressMapListResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapListResponseMessageJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapListResponseMessage]
type accountAddressAddressMapListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapListResponseResult struct {
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
	Enabled    bool                                           `json:"enabled,nullable"`
	ModifiedAt time.Time                                      `json:"modified_at" format:"date-time"`
	JSON       accountAddressAddressMapListResponseResultJSON `json:"-"`
}

// accountAddressAddressMapListResponseResultJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapListResponseResult]
type accountAddressAddressMapListResponseResultJSON struct {
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

func (r *AccountAddressAddressMapListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       accountAddressAddressMapListResponseResultInfoJSON `json:"-"`
}

// accountAddressAddressMapListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountAddressAddressMapListResponseResultInfo]
type accountAddressAddressMapListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressAddressMapListResponseSuccess bool

const (
	AccountAddressAddressMapListResponseSuccessTrue AccountAddressAddressMapListResponseSuccess = true
)

type AccountAddressAddressMapDeleteResponse struct {
	Errors     []AccountAddressAddressMapDeleteResponseError    `json:"errors"`
	Messages   []AccountAddressAddressMapDeleteResponseMessage  `json:"messages"`
	Result     []interface{}                                    `json:"result,nullable"`
	ResultInfo AccountAddressAddressMapDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressAddressMapDeleteResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapDeleteResponseJSON    `json:"-"`
}

// accountAddressAddressMapDeleteResponseJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapDeleteResponse]
type accountAddressAddressMapDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapDeleteResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountAddressAddressMapDeleteResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapDeleteResponseError]
type accountAddressAddressMapDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapDeleteResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountAddressAddressMapDeleteResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapDeleteResponseMessage]
type accountAddressAddressMapDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       accountAddressAddressMapDeleteResponseResultInfoJSON `json:"-"`
}

// accountAddressAddressMapDeleteResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountAddressAddressMapDeleteResponseResultInfo]
type accountAddressAddressMapDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressAddressMapDeleteResponseSuccess bool

const (
	AccountAddressAddressMapDeleteResponseSuccessTrue AccountAddressAddressMapDeleteResponseSuccess = true
)

type AccountAddressAddressMapNewParams struct {
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description param.Field[string] `json:"description"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountAddressAddressMapNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAddressAddressMapUpdateParams struct {
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

func (r AccountAddressAddressMapUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
