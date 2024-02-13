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

// AddressPrefixDelegationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressPrefixDelegationService] method instead.
type AddressPrefixDelegationService struct {
	Options []option.RequestOption
}

// NewAddressPrefixDelegationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressPrefixDelegationService(opts ...option.RequestOption) (r *AddressPrefixDelegationService) {
	r = &AddressPrefixDelegationService{}
	r.Options = opts
	return
}

// Delete an account delegation for a given IP prefix.
func (r *AddressPrefixDelegationService) Delete(ctx context.Context, accountIdentifier string, prefixIdentifier string, delegationIdentifier string, opts ...option.RequestOption) (res *AddressPrefixDelegationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixDelegationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations/%s", accountIdentifier, prefixIdentifier, delegationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a new account delegation for a given IP prefix.
func (r *AddressPrefixDelegationService) IPAddressManagementPrefixDelegationNewPrefixDelegation(ctx context.Context, accountIdentifier string, prefixIdentifier string, body AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationParams, opts ...option.RequestOption) (res *AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all delegations for a given account IP prefix.
func (r *AddressPrefixDelegationService) IPAddressManagementPrefixDelegationListPrefixDelegations(ctx context.Context, accountIdentifier string, prefixIdentifier string, opts ...option.RequestOption) (res *[]AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", accountIdentifier, prefixIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressPrefixDelegationDeleteResponse struct {
	// Delegation identifier tag.
	ID   string                                    `json:"id"`
	JSON addressPrefixDelegationDeleteResponseJSON `json:"-"`
}

// addressPrefixDelegationDeleteResponseJSON contains the JSON metadata for the
// struct [AddressPrefixDelegationDeleteResponse]
type addressPrefixDelegationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponse struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                                                                                    `json:"parent_prefix_id"`
	JSON           addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseJSON `json:"-"`
}

// addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseJSON
// contains the JSON metadata for the struct
// [AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponse]
type addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponse struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                                                                                      `json:"parent_prefix_id"`
	JSON           addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseJSON `json:"-"`
}

// addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseJSON
// contains the JSON metadata for the struct
// [AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponse]
type addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationDeleteResponseEnvelope struct {
	Errors   []AddressPrefixDelegationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixDelegationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixDelegationDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixDelegationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixDelegationDeleteResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixDelegationDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressPrefixDelegationDeleteResponseEnvelope]
type addressPrefixDelegationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressPrefixDelegationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixDelegationDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressPrefixDelegationDeleteResponseEnvelopeErrors]
type addressPrefixDelegationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressPrefixDelegationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixDelegationDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressPrefixDelegationDeleteResponseEnvelopeMessages]
type addressPrefixDelegationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixDelegationDeleteResponseEnvelopeSuccess bool

const (
	AddressPrefixDelegationDeleteResponseEnvelopeSuccessTrue AddressPrefixDelegationDeleteResponseEnvelopeSuccess = true
)

type AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationParams struct {
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr,required"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID param.Field[string] `json:"delegated_account_id,required"`
}

func (r AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelope struct {
	Errors   []AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelope]
type addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeErrors struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeErrors]
type addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeMessages struct {
	Code    int64                                                                                                     `json:"code,required"`
	Message string                                                                                                    `json:"message,required"`
	JSON    addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeMessages]
type addressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeSuccess bool

const (
	AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeSuccessTrue AddressPrefixDelegationIPAddressManagementPrefixDelegationNewPrefixDelegationResponseEnvelopeSuccess = true
)

type AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelope struct {
	Errors   []AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeJSON       `json:"-"`
}

// addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelope]
type addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeErrors struct {
	Code    int64                                                                                                     `json:"code,required"`
	Message string                                                                                                    `json:"message,required"`
	JSON    addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeErrors]
type addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeMessages struct {
	Code    int64                                                                                                       `json:"code,required"`
	Message string                                                                                                      `json:"message,required"`
	JSON    addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeMessages]
type addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeSuccess bool

const (
	AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeSuccessTrue AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeSuccess = true
)

type AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                       `json:"total_count"`
	JSON       addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeResultInfo]
type addressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationIPAddressManagementPrefixDelegationListPrefixDelegationsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
