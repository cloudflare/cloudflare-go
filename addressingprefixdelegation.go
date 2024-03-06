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

// AddressingPrefixDelegationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressingPrefixDelegationService] method instead.
type AddressingPrefixDelegationService struct {
	Options []option.RequestOption
}

// NewAddressingPrefixDelegationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAddressingPrefixDelegationService(opts ...option.RequestOption) (r *AddressingPrefixDelegationService) {
	r = &AddressingPrefixDelegationService{}
	r.Options = opts
	return
}

// Create a new account delegation for a given IP prefix.
func (r *AddressingPrefixDelegationService) New(ctx context.Context, prefixID string, params AddressingPrefixDelegationNewParams, opts ...option.RequestOption) (res *AddressingPrefixDelegationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixDelegationNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all delegations for a given account IP prefix.
func (r *AddressingPrefixDelegationService) List(ctx context.Context, prefixID string, query AddressingPrefixDelegationListParams, opts ...option.RequestOption) (res *[]AddressingPrefixDelegationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixDelegationListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an account delegation for a given IP prefix.
func (r *AddressingPrefixDelegationService) Delete(ctx context.Context, prefixID string, delegationID string, body AddressingPrefixDelegationDeleteParams, opts ...option.RequestOption) (res *AddressingPrefixDelegationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingPrefixDelegationDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/delegations/%s", body.AccountID, prefixID, delegationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingPrefixDelegationNewResponse struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                                    `json:"parent_prefix_id"`
	JSON           addressingPrefixDelegationNewResponseJSON `json:"-"`
}

// addressingPrefixDelegationNewResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixDelegationNewResponse]
type addressingPrefixDelegationNewResponseJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AddressingPrefixDelegationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationNewResponseJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixDelegationListResponse struct {
	// Delegation identifier tag.
	ID string `json:"id"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr      string    `json:"cidr"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID string    `json:"delegated_account_id"`
	ModifiedAt         time.Time `json:"modified_at" format:"date-time"`
	// Identifier
	ParentPrefixID string                                     `json:"parent_prefix_id"`
	JSON           addressingPrefixDelegationListResponseJSON `json:"-"`
}

// addressingPrefixDelegationListResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixDelegationListResponse]
type addressingPrefixDelegationListResponseJSON struct {
	ID                 apijson.Field
	Cidr               apijson.Field
	CreatedAt          apijson.Field
	DelegatedAccountID apijson.Field
	ModifiedAt         apijson.Field
	ParentPrefixID     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AddressingPrefixDelegationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationListResponseJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixDelegationDeleteResponse struct {
	// Delegation identifier tag.
	ID   string                                       `json:"id"`
	JSON addressingPrefixDelegationDeleteResponseJSON `json:"-"`
}

// addressingPrefixDelegationDeleteResponseJSON contains the JSON metadata for the
// struct [AddressingPrefixDelegationDeleteResponse]
type addressingPrefixDelegationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixDelegationNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr,required"`
	// Account identifier for the account to which prefix is being delegated.
	DelegatedAccountID param.Field[string] `json:"delegated_account_id,required"`
}

func (r AddressingPrefixDelegationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingPrefixDelegationNewResponseEnvelope struct {
	Errors   []AddressingPrefixDelegationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixDelegationNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixDelegationNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixDelegationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixDelegationNewResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixDelegationNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingPrefixDelegationNewResponseEnvelope]
type addressingPrefixDelegationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixDelegationNewResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressingPrefixDelegationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixDelegationNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixDelegationNewResponseEnvelopeErrors]
type addressingPrefixDelegationNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixDelegationNewResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    addressingPrefixDelegationNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixDelegationNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixDelegationNewResponseEnvelopeMessages]
type addressingPrefixDelegationNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingPrefixDelegationNewResponseEnvelopeSuccess bool

const (
	AddressingPrefixDelegationNewResponseEnvelopeSuccessTrue AddressingPrefixDelegationNewResponseEnvelopeSuccess = true
)

type AddressingPrefixDelegationListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixDelegationListResponseEnvelope struct {
	Errors   []AddressingPrefixDelegationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixDelegationListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AddressingPrefixDelegationListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingPrefixDelegationListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingPrefixDelegationListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingPrefixDelegationListResponseEnvelopeJSON       `json:"-"`
}

// addressingPrefixDelegationListResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressingPrefixDelegationListResponseEnvelope]
type addressingPrefixDelegationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixDelegationListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressingPrefixDelegationListResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixDelegationListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixDelegationListResponseEnvelopeErrors]
type addressingPrefixDelegationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixDelegationListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    addressingPrefixDelegationListResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixDelegationListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingPrefixDelegationListResponseEnvelopeMessages]
type addressingPrefixDelegationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingPrefixDelegationListResponseEnvelopeSuccess bool

const (
	AddressingPrefixDelegationListResponseEnvelopeSuccessTrue AddressingPrefixDelegationListResponseEnvelopeSuccess = true
)

type AddressingPrefixDelegationListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       addressingPrefixDelegationListResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingPrefixDelegationListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AddressingPrefixDelegationListResponseEnvelopeResultInfo]
type addressingPrefixDelegationListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixDelegationDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingPrefixDelegationDeleteResponseEnvelope struct {
	Errors   []AddressingPrefixDelegationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingPrefixDelegationDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingPrefixDelegationDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingPrefixDelegationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingPrefixDelegationDeleteResponseEnvelopeJSON    `json:"-"`
}

// addressingPrefixDelegationDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressingPrefixDelegationDeleteResponseEnvelope]
type addressingPrefixDelegationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixDelegationDeleteResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    addressingPrefixDelegationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingPrefixDelegationDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingPrefixDelegationDeleteResponseEnvelopeErrors]
type addressingPrefixDelegationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingPrefixDelegationDeleteResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    addressingPrefixDelegationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingPrefixDelegationDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AddressingPrefixDelegationDeleteResponseEnvelopeMessages]
type addressingPrefixDelegationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingPrefixDelegationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingPrefixDelegationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingPrefixDelegationDeleteResponseEnvelopeSuccess bool

const (
	AddressingPrefixDelegationDeleteResponseEnvelopeSuccessTrue AddressingPrefixDelegationDeleteResponseEnvelopeSuccess = true
)
