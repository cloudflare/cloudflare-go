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
func (r *AccountAddressPrefixService) Get(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *SingleResponseZcMisXxG, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify the description for a prefix owned by the account.
func (r *AccountAddressPrefixService) Update(ctx context.Context, accountIdentifier string, prefixIdentifier string, body AccountAddressPrefixUpdateParams, opts ...option.RequestOption) (res *SingleResponseZcMisXxG, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete an unapproved prefix owned by the account.
func (r *AccountAddressPrefixService) Delete(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *APIResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a new prefix under the account.
func (r *AccountAddressPrefixService) IPAddressManagementPrefixesAddPrefix(ctx context.Context, accountIdentifier string, body AccountAddressPrefixIPAddressManagementPrefixesAddPrefixParams, opts ...option.RequestOption) (res *SingleResponseZcMisXxG, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all prefixes owned by the account.
func (r *AccountAddressPrefixService) IPAddressManagementPrefixesListPrefixes(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *ResponseCollectionC81Zx96F, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResponseCollectionC81Zx96F struct {
	Errors     []ResponseCollectionC81Zx96FError    `json:"errors"`
	Messages   []ResponseCollectionC81Zx96FMessage  `json:"messages"`
	Result     []ResponseCollectionC81Zx96FResult   `json:"result"`
	ResultInfo ResponseCollectionC81Zx96FResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollectionC81Zx96FSuccess `json:"success"`
	JSON    responseCollectionC81Zx96FJSON    `json:"-"`
}

// responseCollectionC81Zx96FJSON contains the JSON metadata for the struct
// [ResponseCollectionC81Zx96F]
type responseCollectionC81Zx96FJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionC81Zx96F) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionC81Zx96FError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollectionC81Zx96FErrorJSON `json:"-"`
}

// responseCollectionC81Zx96FErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionC81Zx96FError]
type responseCollectionC81Zx96FErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionC81Zx96FError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionC81Zx96FMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollectionC81Zx96FMessageJSON `json:"-"`
}

// responseCollectionC81Zx96FMessageJSON contains the JSON metadata for the struct
// [ResponseCollectionC81Zx96FMessage]
type responseCollectionC81Zx96FMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionC81Zx96FMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionC81Zx96FResult struct {
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
	OnDemandLocked bool                                 `json:"on_demand_locked"`
	JSON           responseCollectionC81Zx96FResultJSON `json:"-"`
}

// responseCollectionC81Zx96FResultJSON contains the JSON metadata for the struct
// [ResponseCollectionC81Zx96FResult]
type responseCollectionC81Zx96FResultJSON struct {
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

func (r *ResponseCollectionC81Zx96FResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionC81Zx96FResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollectionC81Zx96FResultInfoJSON `json:"-"`
}

// responseCollectionC81Zx96FResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionC81Zx96FResultInfo]
type responseCollectionC81Zx96FResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionC81Zx96FResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollectionC81Zx96FSuccess bool

const (
	ResponseCollectionC81Zx96FSuccessTrue ResponseCollectionC81Zx96FSuccess = true
)

type SingleResponseZcMisXxG struct {
	Errors   []SingleResponseZcMisXxGError   `json:"errors"`
	Messages []SingleResponseZcMisXxGMessage `json:"messages"`
	Result   SingleResponseZcMisXxGResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseZcMisXxGSuccess `json:"success"`
	JSON    singleResponseZcMisXxGJSON    `json:"-"`
}

// singleResponseZcMisXxGJSON contains the JSON metadata for the struct
// [SingleResponseZcMisXxG]
type singleResponseZcMisXxGJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseZcMisXxG) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseZcMisXxGError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleResponseZcMisXxGErrorJSON `json:"-"`
}

// singleResponseZcMisXxGErrorJSON contains the JSON metadata for the struct
// [SingleResponseZcMisXxGError]
type singleResponseZcMisXxGErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseZcMisXxGError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseZcMisXxGMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleResponseZcMisXxGMessageJSON `json:"-"`
}

// singleResponseZcMisXxGMessageJSON contains the JSON metadata for the struct
// [SingleResponseZcMisXxGMessage]
type singleResponseZcMisXxGMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseZcMisXxGMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseZcMisXxGResult struct {
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
	OnDemandLocked bool                             `json:"on_demand_locked"`
	JSON           singleResponseZcMisXxGResultJSON `json:"-"`
}

// singleResponseZcMisXxGResultJSON contains the JSON metadata for the struct
// [SingleResponseZcMisXxGResult]
type singleResponseZcMisXxGResultJSON struct {
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

func (r *SingleResponseZcMisXxGResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleResponseZcMisXxGSuccess bool

const (
	SingleResponseZcMisXxGSuccessTrue SingleResponseZcMisXxGSuccess = true
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
