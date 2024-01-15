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

// AccountAddressPrefixService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAddressPrefixService]
// method instead.
type AccountAddressPrefixService struct {
	Options     []option.RequestOption
	Bgps        *AccountAddressPrefixBgpService
	Delegations *AccountAddressPrefixDelegationService
}

// NewAccountAddressPrefixService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAddressPrefixService(opts ...option.RequestOption) (r *AccountAddressPrefixService) {
	r = &AccountAddressPrefixService{}
	r.Options = opts
	r.Bgps = NewAccountAddressPrefixBgpService(opts...)
	r.Delegations = NewAccountAddressPrefixDelegationService(opts...)
	return
}

// List a particular prefix owned by the account.
func (r *AccountAddressPrefixService) Get(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *AccountAddressPrefixGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify the description for a prefix owned by the account.
func (r *AccountAddressPrefixService) Update(ctx context.Context, accountIdentifier string, prefixIdentifier string, body AccountAddressPrefixUpdateParams, opts ...option.RequestOption) (res *AccountAddressPrefixUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete an unapproved prefix owned by the account.
func (r *AccountAddressPrefixService) Delete(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *AccountAddressPrefixDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a new prefix under the account.
func (r *AccountAddressPrefixService) IPAddressManagementPrefixesAddPrefix(ctx context.Context, accountIdentifier string, body AccountAddressPrefixIPAddressManagementPrefixesAddPrefixParams, opts ...option.RequestOption) (res *AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all prefixes owned by the account.
func (r *AccountAddressPrefixService) IPAddressManagementPrefixesListPrefixes(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAddressPrefixGetResponse struct {
	Errors   []AccountAddressPrefixGetResponseError   `json:"errors"`
	Messages []AccountAddressPrefixGetResponseMessage `json:"messages"`
	Result   AccountAddressPrefixGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressPrefixGetResponseSuccess `json:"success"`
	JSON    accountAddressPrefixGetResponseJSON    `json:"-"`
}

// accountAddressPrefixGetResponseJSON contains the JSON metadata for the struct
// [AccountAddressPrefixGetResponse]
type accountAddressPrefixGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixGetResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountAddressPrefixGetResponseErrorJSON `json:"-"`
}

// accountAddressPrefixGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountAddressPrefixGetResponseError]
type accountAddressPrefixGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixGetResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountAddressPrefixGetResponseMessageJSON `json:"-"`
}

// accountAddressPrefixGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountAddressPrefixGetResponseMessage]
type accountAddressPrefixGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixGetResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Identifier
	AccountID string `json:"account_id"`
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Approval state of the prefix (P = pending, V = active).
	Approved string `json:"approved"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN int64 `json:"asn,nullable"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Description of the prefix.
	Description string `json:"description"`
	// Identifier for the uploaded LOA document.
	LoaDocumentID string    `json:"loa_document_id,nullable"`
	ModifiedAt    time.Time `json:"modified_at" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                                      `json:"on_demand_locked"`
	JSON           accountAddressPrefixGetResponseResultJSON `json:"-"`
}

// accountAddressPrefixGetResponseResultJSON contains the JSON metadata for the
// struct [AccountAddressPrefixGetResponseResult]
type accountAddressPrefixGetResponseResultJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	ASN                  apijson.Field
	Cidr                 apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	LoaDocumentID        apijson.Field
	ModifiedAt           apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAddressPrefixGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressPrefixGetResponseSuccess bool

const (
	AccountAddressPrefixGetResponseSuccessTrue AccountAddressPrefixGetResponseSuccess = true
)

type AccountAddressPrefixUpdateResponse struct {
	Errors   []AccountAddressPrefixUpdateResponseError   `json:"errors"`
	Messages []AccountAddressPrefixUpdateResponseMessage `json:"messages"`
	Result   AccountAddressPrefixUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressPrefixUpdateResponseSuccess `json:"success"`
	JSON    accountAddressPrefixUpdateResponseJSON    `json:"-"`
}

// accountAddressPrefixUpdateResponseJSON contains the JSON metadata for the struct
// [AccountAddressPrefixUpdateResponse]
type accountAddressPrefixUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixUpdateResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountAddressPrefixUpdateResponseErrorJSON `json:"-"`
}

// accountAddressPrefixUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountAddressPrefixUpdateResponseError]
type accountAddressPrefixUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixUpdateResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountAddressPrefixUpdateResponseMessageJSON `json:"-"`
}

// accountAddressPrefixUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountAddressPrefixUpdateResponseMessage]
type accountAddressPrefixUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixUpdateResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Identifier
	AccountID string `json:"account_id"`
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Approval state of the prefix (P = pending, V = active).
	Approved string `json:"approved"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN int64 `json:"asn,nullable"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Description of the prefix.
	Description string `json:"description"`
	// Identifier for the uploaded LOA document.
	LoaDocumentID string    `json:"loa_document_id,nullable"`
	ModifiedAt    time.Time `json:"modified_at" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                                         `json:"on_demand_locked"`
	JSON           accountAddressPrefixUpdateResponseResultJSON `json:"-"`
}

// accountAddressPrefixUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountAddressPrefixUpdateResponseResult]
type accountAddressPrefixUpdateResponseResultJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	ASN                  apijson.Field
	Cidr                 apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	LoaDocumentID        apijson.Field
	ModifiedAt           apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAddressPrefixUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressPrefixUpdateResponseSuccess bool

const (
	AccountAddressPrefixUpdateResponseSuccessTrue AccountAddressPrefixUpdateResponseSuccess = true
)

type AccountAddressPrefixDeleteResponse struct {
	Errors     []AccountAddressPrefixDeleteResponseError    `json:"errors"`
	Messages   []AccountAddressPrefixDeleteResponseMessage  `json:"messages"`
	Result     []interface{}                                `json:"result,nullable"`
	ResultInfo AccountAddressPrefixDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressPrefixDeleteResponseSuccess `json:"success"`
	JSON    accountAddressPrefixDeleteResponseJSON    `json:"-"`
}

// accountAddressPrefixDeleteResponseJSON contains the JSON metadata for the struct
// [AccountAddressPrefixDeleteResponse]
type accountAddressPrefixDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountAddressPrefixDeleteResponseErrorJSON `json:"-"`
}

// accountAddressPrefixDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountAddressPrefixDeleteResponseError]
type accountAddressPrefixDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountAddressPrefixDeleteResponseMessageJSON `json:"-"`
}

// accountAddressPrefixDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountAddressPrefixDeleteResponseMessage]
type accountAddressPrefixDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       accountAddressPrefixDeleteResponseResultInfoJSON `json:"-"`
}

// accountAddressPrefixDeleteResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountAddressPrefixDeleteResponseResultInfo]
type accountAddressPrefixDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressPrefixDeleteResponseSuccess bool

const (
	AccountAddressPrefixDeleteResponseSuccessTrue AccountAddressPrefixDeleteResponseSuccess = true
)

type AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponse struct {
	Errors   []AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseError   `json:"errors"`
	Messages []AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseMessage `json:"messages"`
	Result   AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseSuccess `json:"success"`
	JSON    accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseJSON    `json:"-"`
}

// accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseJSON contains
// the JSON metadata for the struct
// [AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponse]
type accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseErrorJSON `json:"-"`
}

// accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseError]
type accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseMessageJSON `json:"-"`
}

// accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseMessage]
type accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Identifier
	AccountID string `json:"account_id"`
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Approval state of the prefix (P = pending, V = active).
	Approved string `json:"approved"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN int64 `json:"asn,nullable"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Description of the prefix.
	Description string `json:"description"`
	// Identifier for the uploaded LOA document.
	LoaDocumentID string    `json:"loa_document_id,nullable"`
	ModifiedAt    time.Time `json:"modified_at" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                                                                       `json:"on_demand_locked"`
	JSON           accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseResultJSON `json:"-"`
}

// accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseResult]
type accountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseResultJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	ASN                  apijson.Field
	Cidr                 apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	LoaDocumentID        apijson.Field
	ModifiedAt           apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseSuccess bool

const (
	AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseSuccessTrue AccountAddressPrefixIPAddressManagementPrefixesAddPrefixResponseSuccess = true
)

type AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponse struct {
	Errors     []AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseError    `json:"errors"`
	Messages   []AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseMessage  `json:"messages"`
	Result     []AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResult   `json:"result"`
	ResultInfo AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseSuccess `json:"success"`
	JSON    accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseJSON    `json:"-"`
}

// accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseJSON contains
// the JSON metadata for the struct
// [AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponse]
type accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseError struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseErrorJSON `json:"-"`
}

// accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseError]
type accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseMessage struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseMessageJSON `json:"-"`
}

// accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseMessage]
type accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Identifier
	AccountID string `json:"account_id"`
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Approval state of the prefix (P = pending, V = active).
	Approved string `json:"approved"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN int64 `json:"asn,nullable"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Description of the prefix.
	Description string `json:"description"`
	// Identifier for the uploaded LOA document.
	LoaDocumentID string    `json:"loa_document_id,nullable"`
	ModifiedAt    time.Time `json:"modified_at" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                                                                          `json:"on_demand_locked"`
	JSON           accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResultJSON `json:"-"`
}

// accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResult]
type accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResultJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	ASN                  apijson.Field
	Cidr                 apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	LoaDocumentID        apijson.Field
	ModifiedAt           apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                           `json:"total_count"`
	JSON       accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResultInfoJSON `json:"-"`
}

// accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResultInfo]
type accountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseSuccess bool

const (
	AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseSuccessTrue AccountAddressPrefixIPAddressManagementPrefixesListPrefixesResponseSuccess = true
)

type AccountAddressPrefixUpdateParams struct {
	// Description of the prefix.
	Description param.Field[string] `json:"description,required"`
}

func (r AccountAddressPrefixUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAddressPrefixIPAddressManagementPrefixesAddPrefixParams struct {
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN param.Field[int64] `json:"asn,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr,required"`
	// Identifier for the uploaded LOA document.
	LoaDocumentID param.Field[string] `json:"loa_document_id,required"`
}

func (r AccountAddressPrefixIPAddressManagementPrefixesAddPrefixParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
