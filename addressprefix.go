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

// AddressPrefixService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressPrefixService] method
// instead.
type AddressPrefixService struct {
	Options     []option.RequestOption
	BGPs        *AddressPrefixBGPService
	Delegations *AddressPrefixDelegationService
}

// NewAddressPrefixService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressPrefixService(opts ...option.RequestOption) (r *AddressPrefixService) {
	r = &AddressPrefixService{}
	r.Options = opts
	r.BGPs = NewAddressPrefixBGPService(opts...)
	r.Delegations = NewAddressPrefixDelegationService(opts...)
	return
}

// Modify the description for a prefix owned by the account.
func (r *AddressPrefixService) Update(ctx context.Context, accountID string, prefixID string, body AddressPrefixUpdateParams, opts ...option.RequestOption) (res *AddressPrefixUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an unapproved prefix owned by the account.
func (r *AddressPrefixService) Delete(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *AddressPrefixDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a particular prefix owned by the account.
func (r *AddressPrefixService) Get(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *AddressPrefixGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add a new prefix under the account.
func (r *AddressPrefixService) IPAddressManagementPrefixesAddPrefix(ctx context.Context, accountID string, body AddressPrefixIPAddressManagementPrefixesAddPrefixParams, opts ...option.RequestOption) (res *AddressPrefixIPAddressManagementPrefixesAddPrefixResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all prefixes owned by the account.
func (r *AddressPrefixService) IPAddressManagementPrefixesListPrefixes(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]AddressPrefixIPAddressManagementPrefixesListPrefixesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressPrefixUpdateResponse struct {
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
	Asn int64 `json:"asn,nullable"`
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
	OnDemandLocked bool                            `json:"on_demand_locked"`
	JSON           addressPrefixUpdateResponseJSON `json:"-"`
}

// addressPrefixUpdateResponseJSON contains the JSON metadata for the struct
// [AddressPrefixUpdateResponse]
type addressPrefixUpdateResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	Asn                  apijson.Field
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

func (r *AddressPrefixUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AddressPrefixDeleteResponseUnknown],
// [AddressPrefixDeleteResponseArray] or [shared.UnionString].
type AddressPrefixDeleteResponse interface {
	ImplementsAddressPrefixDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressPrefixDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressPrefixDeleteResponseArray []interface{}

func (r AddressPrefixDeleteResponseArray) ImplementsAddressPrefixDeleteResponse() {}

type AddressPrefixGetResponse struct {
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
	Asn int64 `json:"asn,nullable"`
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
	OnDemandLocked bool                         `json:"on_demand_locked"`
	JSON           addressPrefixGetResponseJSON `json:"-"`
}

// addressPrefixGetResponseJSON contains the JSON metadata for the struct
// [AddressPrefixGetResponse]
type addressPrefixGetResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	Asn                  apijson.Field
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

func (r *AddressPrefixGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixIPAddressManagementPrefixesAddPrefixResponse struct {
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
	Asn int64 `json:"asn,nullable"`
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
	OnDemandLocked bool                                                          `json:"on_demand_locked"`
	JSON           addressPrefixIPAddressManagementPrefixesAddPrefixResponseJSON `json:"-"`
}

// addressPrefixIPAddressManagementPrefixesAddPrefixResponseJSON contains the JSON
// metadata for the struct
// [AddressPrefixIPAddressManagementPrefixesAddPrefixResponse]
type addressPrefixIPAddressManagementPrefixesAddPrefixResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	Asn                  apijson.Field
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

func (r *AddressPrefixIPAddressManagementPrefixesAddPrefixResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixIPAddressManagementPrefixesListPrefixesResponse struct {
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
	Asn int64 `json:"asn,nullable"`
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
	OnDemandLocked bool                                                             `json:"on_demand_locked"`
	JSON           addressPrefixIPAddressManagementPrefixesListPrefixesResponseJSON `json:"-"`
}

// addressPrefixIPAddressManagementPrefixesListPrefixesResponseJSON contains the
// JSON metadata for the struct
// [AddressPrefixIPAddressManagementPrefixesListPrefixesResponse]
type addressPrefixIPAddressManagementPrefixesListPrefixesResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	Asn                  apijson.Field
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

func (r *AddressPrefixIPAddressManagementPrefixesListPrefixesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixUpdateParams struct {
	// Description of the prefix.
	Description param.Field[string] `json:"description,required"`
}

func (r AddressPrefixUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixUpdateResponseEnvelope struct {
	Errors   []AddressPrefixUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixUpdateResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressPrefixUpdateResponseEnvelope]
type addressPrefixUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    addressPrefixUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressPrefixUpdateResponseEnvelopeErrors]
type addressPrefixUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    addressPrefixUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressPrefixUpdateResponseEnvelopeMessages]
type addressPrefixUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixUpdateResponseEnvelopeSuccess bool

const (
	AddressPrefixUpdateResponseEnvelopeSuccessTrue AddressPrefixUpdateResponseEnvelopeSuccess = true
)

type AddressPrefixDeleteResponseEnvelope struct {
	Errors   []AddressPrefixDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressPrefixDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressPrefixDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressPrefixDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressPrefixDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressPrefixDeleteResponseEnvelope]
type addressPrefixDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    addressPrefixDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressPrefixDeleteResponseEnvelopeErrors]
type addressPrefixDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    addressPrefixDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressPrefixDeleteResponseEnvelopeMessages]
type addressPrefixDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixDeleteResponseEnvelopeSuccess bool

const (
	AddressPrefixDeleteResponseEnvelopeSuccessTrue AddressPrefixDeleteResponseEnvelopeSuccess = true
)

type AddressPrefixDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       addressPrefixDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressPrefixDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AddressPrefixDeleteResponseEnvelopeResultInfo]
type addressPrefixDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixGetResponseEnvelope struct {
	Errors   []AddressPrefixGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixGetResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressPrefixGetResponseEnvelope]
type addressPrefixGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    addressPrefixGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressPrefixGetResponseEnvelopeErrors]
type addressPrefixGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    addressPrefixGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AddressPrefixGetResponseEnvelopeMessages]
type addressPrefixGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixGetResponseEnvelopeSuccess bool

const (
	AddressPrefixGetResponseEnvelopeSuccessTrue AddressPrefixGetResponseEnvelopeSuccess = true
)

type AddressPrefixIPAddressManagementPrefixesAddPrefixParams struct {
	// Autonomous System Number (ASN) the prefix will be advertised under.
	Asn param.Field[int64] `json:"asn,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr,required"`
	// Identifier for the uploaded LOA document.
	LoaDocumentID param.Field[string] `json:"loa_document_id,required"`
}

func (r AddressPrefixIPAddressManagementPrefixesAddPrefixParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelope struct {
	Errors   []AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixIPAddressManagementPrefixesAddPrefixResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelope]
type addressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    addressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeErrors]
type addressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    addressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeMessages]
type addressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeSuccess bool

const (
	AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeSuccessTrue AddressPrefixIPAddressManagementPrefixesAddPrefixResponseEnvelopeSuccess = true
)

type AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelope struct {
	Errors   []AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressPrefixIPAddressManagementPrefixesListPrefixesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeJSON       `json:"-"`
}

// addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelope]
type addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeErrors struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeErrors]
type addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeMessages struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeMessages]
type addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeSuccess bool

const (
	AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeSuccessTrue AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeSuccess = true
)

type AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                            `json:"total_count"`
	JSON       addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeResultInfo]
type addressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixIPAddressManagementPrefixesListPrefixesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
