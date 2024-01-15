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
func (r *AccountAddressPrefixDelegationService) Delete(ctx context.Context, accountIdentifier string, prefixIdentifier string, delegationIdentifier string, opts ...option.RequestOption) (res *AccountAddressPrefixDelegationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations/%s", accountIdentifier, prefixIdentifier, delegationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new account delegation for a given IP prefix.
func (r *AccountAddressPrefixDelegationService) IPAddressManagementPrefixDelegationNewPrefixDelegation(ctx context.Context, accountIdentifier string, prefixIdentifier string, body AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationParams, opts ...option.RequestOption) (res *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all delegations for a given account IP prefix.
func (r *AccountAddressPrefixDelegationService) IPAddressManagementPrefixDelegationListPrefixDelegations(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAddressPrefixDelegationDeleteResponse struct {
	Errors   []AccountAddressPrefixDelegationDeleteResponseError   `json:"errors"`
	Messages []AccountAddressPrefixDelegationDeleteResponseMessage `json:"messages"`
	Result   AccountAddressPrefixDelegationDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressPrefixDelegationDeleteResponseSuccess `json:"success"`
	JSON    accountAddressPrefixDelegationDeleteResponseJSON    `json:"-"`
}

// accountAddressPrefixDelegationDeleteResponseJSON contains the JSON metadata for
// the struct [AccountAddressPrefixDelegationDeleteResponse]
type accountAddressPrefixDelegationDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDelegationDeleteResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountAddressPrefixDelegationDeleteResponseErrorJSON `json:"-"`
}

// accountAddressPrefixDelegationDeleteResponseErrorJSON contains the JSON metadata
// for the struct [AccountAddressPrefixDelegationDeleteResponseError]
type accountAddressPrefixDelegationDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDelegationDeleteResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountAddressPrefixDelegationDeleteResponseMessageJSON `json:"-"`
}

// accountAddressPrefixDelegationDeleteResponseMessageJSON contains the JSON
// metadata for the struct [AccountAddressPrefixDelegationDeleteResponseMessage]
type accountAddressPrefixDelegationDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDelegationDeleteResponseResult struct {
	// Delegation identifier tag.
	ID   string                                                 `json:"id"`
	JSON accountAddressPrefixDelegationDeleteResponseResultJSON `json:"-"`
}

// accountAddressPrefixDelegationDeleteResponseResultJSON contains the JSON
// metadata for the struct [AccountAddressPrefixDelegationDeleteResponseResult]
type accountAddressPrefixDelegationDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressPrefixDelegationDeleteResponseSuccess bool

const (
	AccountAddressPrefixDelegationDeleteResponseSuccessTrue AccountAddressPrefixDelegationDeleteResponseSuccess = true
)

type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponse struct {
	Errors   []AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseError   `json:"errors"`
	Messages []AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseMessage `json:"messages"`
	Result   AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseSuccess `json:"success"`
	JSON    accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseJSON    `json:"-"`
}

// accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponse]
type accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseError struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseErrorJSON `json:"-"`
}

// accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseError]
type accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseMessage struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseMessageJSON `json:"-"`
}

// accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseMessage]
type accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseResult struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                                                                                                 `json:"parent_prefix_id"`
	JSON           accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseResultJSON `json:"-"`
}

// accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseResult]
type accountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseResultJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseSuccess bool

const (
	AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseSuccessTrue AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseSuccess = true
)

type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponse struct {
	Errors     []AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseError    `json:"errors"`
	Messages   []AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseMessage  `json:"messages"`
	Result     []AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResult   `json:"result"`
	ResultInfo AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseSuccess `json:"success"`
	JSON    accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseJSON    `json:"-"`
}

// accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponse]
type accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseError struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseErrorJSON `json:"-"`
}

// accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseError]
type accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseMessage struct {
	Code    int64                                                                                                     `json:"code,required"`
	Message string                                                                                                    `json:"message,required"`
	JSON    accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseMessageJSON `json:"-"`
}

// accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseMessage]
type accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResult struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                                                                                                   `json:"parent_prefix_id"`
	JSON           accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResultJSON `json:"-"`
}

// accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResult]
type accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResultJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                      `json:"total_count"`
	JSON       accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResultInfoJSON `json:"-"`
}

// accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResultInfo]
type accountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseSuccess bool

const (
	AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseSuccessTrue AccountAddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseSuccess = true
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
