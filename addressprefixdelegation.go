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

// Create a new account delegation for a given IP prefix.
func (r *AddressPrefixDelegationService) New(ctx context.Context, accountID string, prefixID string, body AddressPrefixDelegationNewParams, opts ...option.RequestOption) (res *AddressPrefixDelegationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixDelegationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all delegations for a given account IP prefix.
func (r *AddressPrefixDelegationService) List(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *[]AddressPrefixDelegationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixDelegationListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an account delegation for a given IP prefix.
func (r *AddressPrefixDelegationService) Delete(ctx context.Context, accountID string, prefixID string, delegationID string, opts ...option.RequestOption) (res *AddressPrefixDelegationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressPrefixDelegationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations/%s", accountID, prefixID, delegationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressPrefixDelegationNewResponse struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                                 `json:"parent_prefix_id"`
	JSON           addressPrefixDelegationNewResponseJSON `json:"-"`
}

// addressPrefixDelegationNewResponseJSON contains the JSON metadata for the struct
// [AddressPrefixDelegationNewResponse]
type addressPrefixDelegationNewResponseJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AddressPrefixDelegationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationListResponse struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                                  `json:"parent_prefix_id"`
	JSON           addressPrefixDelegationListResponseJSON `json:"-"`
}

// addressPrefixDelegationListResponseJSON contains the JSON metadata for the
// struct [AddressPrefixDelegationListResponse]
type addressPrefixDelegationListResponseJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AddressPrefixDelegationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type AddressPrefixDelegationNewParams struct {
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr,required"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID param.Field[string] `json:"delegated_account_id,required"`
}

func (r AddressPrefixDelegationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressPrefixDelegationNewResponseEnvelope struct {
	Errors   []AddressPrefixDelegationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixDelegationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressPrefixDelegationNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressPrefixDelegationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressPrefixDelegationNewResponseEnvelopeJSON    `json:"-"`
}

// addressPrefixDelegationNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressPrefixDelegationNewResponseEnvelope]
type addressPrefixDelegationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressPrefixDelegationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixDelegationNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressPrefixDelegationNewResponseEnvelopeErrors]
type addressPrefixDelegationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressPrefixDelegationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixDelegationNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressPrefixDelegationNewResponseEnvelopeMessages]
type addressPrefixDelegationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixDelegationNewResponseEnvelopeSuccess bool

const (
	AddressPrefixDelegationNewResponseEnvelopeSuccessTrue AddressPrefixDelegationNewResponseEnvelopeSuccess = true
)

type AddressPrefixDelegationListResponseEnvelope struct {
	Errors   []AddressPrefixDelegationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressPrefixDelegationListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressPrefixDelegationListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressPrefixDelegationListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressPrefixDelegationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressPrefixDelegationListResponseEnvelopeJSON       `json:"-"`
}

// addressPrefixDelegationListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressPrefixDelegationListResponseEnvelope]
type addressPrefixDelegationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    addressPrefixDelegationListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressPrefixDelegationListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressPrefixDelegationListResponseEnvelopeErrors]
type addressPrefixDelegationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressPrefixDelegationListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressPrefixDelegationListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressPrefixDelegationListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressPrefixDelegationListResponseEnvelopeMessages]
type addressPrefixDelegationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressPrefixDelegationListResponseEnvelopeSuccess bool

const (
	AddressPrefixDelegationListResponseEnvelopeSuccessTrue AddressPrefixDelegationListResponseEnvelopeSuccess = true
)

type AddressPrefixDelegationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       addressPrefixDelegationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressPrefixDelegationListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressPrefixDelegationListResponseEnvelopeResultInfo]
type addressPrefixDelegationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressPrefixDelegationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
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
