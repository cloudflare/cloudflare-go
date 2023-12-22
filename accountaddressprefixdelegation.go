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

// AccountAddressPrefixDelegationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAddressPrefixDelegationService] method instead.
type AccountAddressPrefixDelegationService struct {
	Options []option.RequestOption
}

// NewAccountAddressPrefixDelegationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressPrefixDelegationService(opts ...option.RequestOption) (r *AccountAddressPrefixDelegationService) {
	r = &AccountAddressPrefixDelegationService{}
	r.Options = opts
	return
}

// Delete an account delegation for a given IP prefix.
func (r *AccountAddressPrefixDelegationService) Delete(ctx context.Context, accountIdentifier string, prefixIdentifier string, delegationIdentifier string, opts ...option.RequestOption) (res *IDResponseJbqQi2td, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations/%s", accountIdentifier, prefixIdentifier, delegationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new account delegation for a given IP prefix.
func (r *AccountAddressPrefixDelegationService) IPAddressManagementPrefixDelegationNewPrefixDelegation(ctx context.Context, accountIdentifier string, prefixIdentifier string, body AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationParams, opts ...option.RequestOption) (res *SchemasSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all delegations for a given account IP prefix.
func (r *AccountAddressPrefixDelegationService) IPAddressManagementPrefixDelegationListPrefixDelegations(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *SchemasResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IDResponseJbqQi2td struct {
	Errors   []IDResponseJbqQi2tdError   `json:"errors"`
	Messages []IDResponseJbqQi2tdMessage `json:"messages"`
	Result   IDResponseJbqQi2tdResult    `json:"result"`
	// Whether the API call was successful
	Success IDResponseJbqQi2tdSuccess `json:"success"`
	JSON    idResponseJbqQi2tdJSON    `json:"-"`
}

// idResponseJbqQi2tdJSON contains the JSON metadata for the struct
// [IDResponseJbqQi2td]
type idResponseJbqQi2tdJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseJbqQi2td) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponseJbqQi2tdError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    idResponseJbqQi2tdErrorJSON `json:"-"`
}

// idResponseJbqQi2tdErrorJSON contains the JSON metadata for the struct
// [IDResponseJbqQi2tdError]
type idResponseJbqQi2tdErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseJbqQi2tdError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponseJbqQi2tdMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    idResponseJbqQi2tdMessageJSON `json:"-"`
}

// idResponseJbqQi2tdMessageJSON contains the JSON metadata for the struct
// [IDResponseJbqQi2tdMessage]
type idResponseJbqQi2tdMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseJbqQi2tdMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IDResponseJbqQi2tdResult struct {
	// Delegation identifier tag.
	ID   string                       `json:"id"`
	JSON idResponseJbqQi2tdResultJSON `json:"-"`
}

// idResponseJbqQi2tdResultJSON contains the JSON metadata for the struct
// [IDResponseJbqQi2tdResult]
type idResponseJbqQi2tdResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseJbqQi2tdResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IDResponseJbqQi2tdSuccess bool

const (
	IDResponseJbqQi2tdSuccessTrue IDResponseJbqQi2tdSuccess = true
)

type SchemasResponseCollection struct {
	Errors     []SchemasResponseCollectionError    `json:"errors"`
	Messages   []SchemasResponseCollectionMessage  `json:"messages"`
	Result     []SchemasResponseCollectionResult   `json:"result"`
	ResultInfo SchemasResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasResponseCollectionSuccess `json:"success"`
	JSON    schemasResponseCollectionJSON    `json:"-"`
}

// schemasResponseCollectionJSON contains the JSON metadata for the struct
// [SchemasResponseCollection]
type schemasResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    schemasResponseCollectionErrorJSON `json:"-"`
}

// schemasResponseCollectionErrorJSON contains the JSON metadata for the struct
// [SchemasResponseCollectionError]
type schemasResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    schemasResponseCollectionMessageJSON `json:"-"`
}

// schemasResponseCollectionMessageJSON contains the JSON metadata for the struct
// [SchemasResponseCollectionMessage]
type schemasResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionResult struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                              `json:"parent_prefix_id"`
	JSON           schemasResponseCollectionResultJSON `json:"-"`
}

// schemasResponseCollectionResultJSON contains the JSON metadata for the struct
// [SchemasResponseCollectionResult]
type schemasResponseCollectionResultJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       schemasResponseCollectionResultInfoJSON `json:"-"`
}

// schemasResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [SchemasResponseCollectionResultInfo]
type schemasResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasResponseCollectionSuccess bool

const (
	SchemasResponseCollectionSuccessTrue SchemasResponseCollectionSuccess = true
)

type SchemasSingleResponse struct {
	Errors   []SchemasSingleResponseError   `json:"errors"`
	Messages []SchemasSingleResponseMessage `json:"messages"`
	Result   SchemasSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success SchemasSingleResponseSuccess `json:"success"`
	JSON    schemasSingleResponseJSON    `json:"-"`
}

// schemasSingleResponseJSON contains the JSON metadata for the struct
// [SchemasSingleResponse]
type schemasSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    schemasSingleResponseErrorJSON `json:"-"`
}

// schemasSingleResponseErrorJSON contains the JSON metadata for the struct
// [SchemasSingleResponseError]
type schemasSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    schemasSingleResponseMessageJSON `json:"-"`
}

// schemasSingleResponseMessageJSON contains the JSON metadata for the struct
// [SchemasSingleResponseMessage]
type schemasSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasSingleResponseResult struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                          `json:"parent_prefix_id"`
	JSON           schemasSingleResponseResultJSON `json:"-"`
}

// schemasSingleResponseResultJSON contains the JSON metadata for the struct
// [SchemasSingleResponseResult]
type schemasSingleResponseResultJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasSingleResponseSuccess bool

const (
	SchemasSingleResponseSuccessTrue SchemasSingleResponseSuccess = true
)

type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationParams struct {
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr,required"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID param.Field[string] `json:"delegated_account_id,required"`
}

func (r AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
