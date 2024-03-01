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

// AddressingPrefixService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressingPrefixService] method
// instead.
type AddressingPrefixService struct {
	Options     []option.RequestOption
	BGP         *AddressingPrefixBGPService
	Delegations *AddressingPrefixDelegationService
}

// NewAddressingPrefixService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressingPrefixService(opts ...option.RequestOption) (r *AddressingPrefixService) {
	r = &AddressingPrefixService{}
	r.Options = opts
	r.BGP = NewAddressingPrefixBGPService(opts...)
	r.Delegations = NewAddressingPrefixDelegationService(opts...)
	return
}

// Add a new prefix under the account.
func (r *AddressingPrefixService) New(ctx context.Context, params AddressingPrefixNewParams, opts ...option.RequestOption) (res *AddressingPrefixNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all prefixes owned by the account.
func (r *AddressingPrefixService) List(ctx context.Context, query AddressingPrefixListParams, opts ...option.RequestOption) (res *[]AddressingPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an unapproved prefix owned by the account.
func (r *AddressingPrefixService) Delete(ctx context.Context, prefixID string, body AddressingPrefixDeleteParams, opts ...option.RequestOption) (res *AddressingPrefixDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", body.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify the description for a prefix owned by the account.
func (r *AddressingPrefixService) Edit(ctx context.Context, prefixID string, params AddressingPrefixEditParams, opts ...option.RequestOption) (res *AddressingPrefixEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a particular prefix owned by the account.
func (r *AddressingPrefixService) Get(ctx context.Context, prefixID string, query AddressingPrefixGetParams, opts ...option.RequestOption) (res *AddressingPrefixGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingPrefixNewResponse struct {
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
	LOADocumentID string    `json:"loa_document_id,nullable"`
	ModifiedAt    time.Time `json:"modified_at" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                            `json:"on_demand_locked"`
	JSON           addressingPrefixNewResponseJSON `json:"-"`
}

// addressingPrefixNewResponseJSON contains the JSON metadata for the struct
// [AddressingPrefixNewResponse]
type addressingPrefixNewResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	ASN                  apijson.Field
	Cidr                 apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	LOADocumentID        apijson.Field
	ModifiedAt           apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixListResponse struct {
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
	LOADocumentID string    `json:"loa_document_id,nullable"`
	ModifiedAt    time.Time `json:"modified_at" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                             `json:"on_demand_locked"`
	JSON           addressingPrefixListResponseJSON `json:"-"`
}

// addressingPrefixListResponseJSON contains the JSON metadata for the struct
// [AddressingPrefixListResponse]
type addressingPrefixListResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	ASN                  apijson.Field
	Cidr                 apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	LOADocumentID        apijson.Field
	ModifiedAt           apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AddressingPrefixDeleteResponseUnknown],
// [AddressingPrefixDeleteResponseArray] or [shared.UnionString].
type AddressingPrefixDeleteResponse interface {
	ImplementsAddressingPrefixDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingPrefixDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingPrefixDeleteResponseArray []interface{}

func (r AddressingPrefixDeleteResponseArray) ImplementsAddressingPrefixDeleteResponse() {}

type AddressingPrefixEditResponse struct {
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
	LOADocumentID string    `json:"loa_document_id,nullable"`
	ModifiedAt    time.Time `json:"modified_at" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                             `json:"on_demand_locked"`
	JSON           addressingPrefixEditResponseJSON `json:"-"`
}

// addressingPrefixEditResponseJSON contains the JSON metadata for the struct
// [AddressingPrefixEditResponse]
type addressingPrefixEditResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	ASN                  apijson.Field
	Cidr                 apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	LOADocumentID        apijson.Field
	ModifiedAt           apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixGetResponse struct {
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
	LOADocumentID string    `json:"loa_document_id,nullable"`
	ModifiedAt    time.Time `json:"modified_at" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                            `json:"on_demand_locked"`
	JSON           addressingPrefixGetResponseJSON `json:"-"`
}

// addressingPrefixGetResponseJSON contains the JSON metadata for the struct
// [AddressingPrefixGetResponse]
type addressingPrefixGetResponseJSON struct {
	ID                   apijson.Field
	AccountID            apijson.Field
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	Approved             apijson.Field
	ASN                  apijson.Field
	Cidr                 apijson.Field
	CreatedAt            apijson.Field
	Description          apijson.Field
	LOADocumentID        apijson.Field
	ModifiedAt           apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AddressingPrefixGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN param.Field[int64] `json:"asn,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr,required"`
	// Identifier for the uploaded LOA document.
	LOADocumentID param.Field[string] `json:"loa_document_id,required"`
}

func (r AddressingPrefixNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingPrefixNewResponseEnvelope struct {
	Errors   []AddressingPrefixNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixNewResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressingPrefixNewResponseEnvelope]
type addressingPrefixNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    addressingPrefixNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressingPrefixNewResponseEnvelopeErrors]
type addressingPrefixNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    addressingPrefixNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressingPrefixNewResponseEnvelopeMessages]
type addressingPrefixNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixNewResponseEnvelopeSuccess bool

const (
	AddressingPrefixNewResponseEnvelopeSuccessTrue AddressingPrefixNewResponseEnvelopeSuccess = true
)

type AddressingPrefixListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixListResponseEnvelope struct {
	Errors   []AddressingPrefixListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressingPrefixListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingPrefixListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingPrefixListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingPrefixListResponseEnvelopeJSON       `json:"-"`
}

// addressingPrefixListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressingPrefixListResponseEnvelope]
type addressingPrefixListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    addressingPrefixListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressingPrefixListResponseEnvelopeErrors]
type addressingPrefixListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    addressingPrefixListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressingPrefixListResponseEnvelopeMessages]
type addressingPrefixListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixListResponseEnvelopeSuccess bool

const (
	AddressingPrefixListResponseEnvelopeSuccessTrue AddressingPrefixListResponseEnvelopeSuccess = true
)

type AddressingPrefixListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       addressingPrefixListResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingPrefixListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AddressingPrefixListResponseEnvelopeResultInfo]
type addressingPrefixListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixDeleteResponseEnvelope struct {
	Errors   []AddressingPrefixDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingPrefixDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingPrefixDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingPrefixDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressingPrefixDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressingPrefixDeleteResponseEnvelope]
type addressingPrefixDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    addressingPrefixDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressingPrefixDeleteResponseEnvelopeErrors]
type addressingPrefixDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    addressingPrefixDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressingPrefixDeleteResponseEnvelopeMessages]
type addressingPrefixDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixDeleteResponseEnvelopeSuccess bool

const (
	AddressingPrefixDeleteResponseEnvelopeSuccessTrue AddressingPrefixDeleteResponseEnvelopeSuccess = true
)

type AddressingPrefixDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       addressingPrefixDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingPrefixDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AddressingPrefixDeleteResponseEnvelopeResultInfo]
type addressingPrefixDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Description of the prefix.
	Description param.Field[string] `json:"description,required"`
}

func (r AddressingPrefixEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingPrefixEditResponseEnvelope struct {
	Errors   []AddressingPrefixEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixEditResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixEditResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressingPrefixEditResponseEnvelope]
type addressingPrefixEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixEditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    addressingPrefixEditResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressingPrefixEditResponseEnvelopeErrors]
type addressingPrefixEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixEditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    addressingPrefixEditResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressingPrefixEditResponseEnvelopeMessages]
type addressingPrefixEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixEditResponseEnvelopeSuccess bool

const (
	AddressingPrefixEditResponseEnvelopeSuccessTrue AddressingPrefixEditResponseEnvelopeSuccess = true
)

type AddressingPrefixGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixGetResponseEnvelope struct {
	Errors   []AddressingPrefixGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixGetResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressingPrefixGetResponseEnvelope]
type addressingPrefixGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    addressingPrefixGetResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressingPrefixGetResponseEnvelopeErrors]
type addressingPrefixGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingPrefixGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    addressingPrefixGetResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressingPrefixGetResponseEnvelopeMessages]
type addressingPrefixGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingPrefixGetResponseEnvelopeSuccess bool

const (
	AddressingPrefixGetResponseEnvelopeSuccessTrue AddressingPrefixGetResponseEnvelopeSuccess = true
)
