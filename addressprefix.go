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

// Add a new prefix under the account.
func (r *AddressPrefixService) New(ctx context.Context, accountID string, body AddressPrefixNewParams, opts ...option.RequestOption) (res *AddressPrefixNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all prefixes owned by the account.
func (r *AddressPrefixService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]AddressPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

// Modify the description for a prefix owned by the account.
func (r *AddressPrefixService) Edit(ctx context.Context, accountID string, prefixID string, body AddressPrefixEditParams, opts ...option.RequestOption) (res *AddressPrefixEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
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

type AddressPrefixNewResponse struct {
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
	JSON           addressPrefixNewResponseJSON `json:"-"`
}

// addressPrefixNewResponseJSON contains the JSON metadata for the struct
// [AddressPrefixNewResponse]
type addressPrefixNewResponseJSON struct {
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

func (r *AddressPrefixNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixListResponse struct {
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
	OnDemandLocked bool                          `json:"on_demand_locked"`
	JSON           addressPrefixListResponseJSON `json:"-"`
}

// addressPrefixListResponseJSON contains the JSON metadata for the struct
// [AddressPrefixListResponse]
type addressPrefixListResponseJSON struct {
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

func (r *AddressPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
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

type AddressPrefixEditResponse struct {
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
	OnDemandLocked bool                          `json:"on_demand_locked"`
	JSON           addressPrefixEditResponseJSON `json:"-"`
}

// addressPrefixEditResponseJSON contains the JSON metadata for the struct
// [AddressPrefixEditResponse]
type addressPrefixEditResponseJSON struct {
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

func (r *AddressPrefixEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type AddressPrefixNewParams struct {
	// Autonomous System Number (ASN) the prefix will be advertised under.
	Asn param.Field[int64] `json:"asn,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr,required"`
	// Identifier for the uploaded LOA document.
	LoaDocumentID param.Field[string] `json:"loa_document_id,required"`
}

func (r AddressPrefixNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixNewResponseEnvelope struct {
	Errors   []AddressPrefixNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixNewResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressPrefixNewResponseEnvelope]
type addressPrefixNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    addressPrefixNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressPrefixNewResponseEnvelopeErrors]
type addressPrefixNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    addressPrefixNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AddressPrefixNewResponseEnvelopeMessages]
type addressPrefixNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixNewResponseEnvelopeSuccess bool

const (
	AddressPrefixNewResponseEnvelopeSuccessTrue AddressPrefixNewResponseEnvelopeSuccess = true
)

type AddressPrefixListResponseEnvelope struct {
	Errors   []AddressPrefixListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressPrefixListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressPrefixListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressPrefixListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressPrefixListResponseEnvelopeJSON       `json:"-"`
}

// addressPrefixListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressPrefixListResponseEnvelope]
type addressPrefixListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixListResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    addressPrefixListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressPrefixListResponseEnvelopeErrors]
type addressPrefixListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixListResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    addressPrefixListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AddressPrefixListResponseEnvelopeMessages]
type addressPrefixListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixListResponseEnvelopeSuccess bool

const (
	AddressPrefixListResponseEnvelopeSuccessTrue AddressPrefixListResponseEnvelopeSuccess = true
)

type AddressPrefixListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       addressPrefixListResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressPrefixListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AddressPrefixListResponseEnvelopeResultInfo]
type addressPrefixListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type AddressPrefixEditParams struct {
	// Description of the prefix.
	Description param.Field[string] `json:"description,required"`
}

func (r AddressPrefixEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixEditResponseEnvelope struct {
	Errors   []AddressPrefixEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixEditResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixEditResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressPrefixEditResponseEnvelope]
type addressPrefixEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    addressPrefixEditResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressPrefixEditResponseEnvelopeErrors]
type addressPrefixEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    addressPrefixEditResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AddressPrefixEditResponseEnvelopeMessages]
type addressPrefixEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixEditResponseEnvelopeSuccess bool

const (
	AddressPrefixEditResponseEnvelopeSuccessTrue AddressPrefixEditResponseEnvelopeSuccess = true
)

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
