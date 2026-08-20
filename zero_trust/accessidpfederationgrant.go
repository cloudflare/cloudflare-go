// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// AccessIdPFederationGrantService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessIdPFederationGrantService] method instead.
type AccessIdPFederationGrantService struct {
	Options []option.RequestOption
}

// NewAccessIdPFederationGrantService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessIdPFederationGrantService(opts ...option.RequestOption) (r *AccessIdPFederationGrantService) {
	r = &AccessIdPFederationGrantService{}
	r.Options = opts
	return
}

// Creates an IdP federation grant for the specified identity provider, making it
// available for federation to other accounts in the same Cloudflare organization.
//
// The account must belong to a Cloudflare organization. One-time pin and
// Cloudflare-managed identity providers cannot be federated. An account can
// federate at most five identity providers at a time.
func (r *AccessIdPFederationGrantService) New(ctx context.Context, params AccessIdPFederationGrantNewParams, opts ...option.RequestOption) (res *IdPFederationGrant, err error) {
	var env AccessIdPFederationGrantNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/idp_federation_grants", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists the IdP federation grants owned by the account.
func (r *AccessIdPFederationGrantService) List(ctx context.Context, query AccessIdPFederationGrantListParams, opts ...option.RequestOption) (res *[]IdPFederationGrant, err error) {
	var env AccessIdPFederationGrantListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/idp_federation_grants", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Deletes an IdP federation grant. The identity provider remains in the account,
// but it is no longer available for federation to other accounts.
func (r *AccessIdPFederationGrantService) Delete(ctx context.Context, grantID string, body AccessIdPFederationGrantDeleteParams, opts ...option.RequestOption) (res *AccessIdPFederationGrantDeleteResponse, err error) {
	var env AccessIdPFederationGrantDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if grantID == "" {
		err = errors.New("missing required grant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/idp_federation_grants/%s", body.AccountID, grantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves a single IdP federation grant by its UID.
func (r *AccessIdPFederationGrantService) Get(ctx context.Context, grantID string, query AccessIdPFederationGrantGetParams, opts ...option.RequestOption) (res *IdPFederationGrant, err error) {
	var env AccessIdPFederationGrantGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if grantID == "" {
		err = errors.New("missing required grant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/idp_federation_grants/%s", query.AccountID, grantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type IdPFederationGrant struct {
	// UID of the IdP federation grant.
	ID string `json:"id" api:"required"`
	// UID of the identity provider being federated.
	IdPID string                 `json:"idp_id" api:"required" format:"uuid"`
	JSON  IdPFederationGrantJSON `json:"-"`
}

// IdPFederationGrantJSON contains the JSON metadata for the struct
// [IdPFederationGrant]
type IdPFederationGrantJSON struct {
	ID          apijson.Field
	IdPID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdPFederationGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdPFederationGrantJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantDeleteResponse struct {
	// UID of the deleted IdP federation grant.
	ID   string                                     `json:"id"`
	JSON accessIdPFederationGrantDeleteResponseJSON `json:"-"`
}

// accessIdPFederationGrantDeleteResponseJSON contains the JSON metadata for the
// struct [AccessIdPFederationGrantDeleteResponse]
type accessIdPFederationGrantDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// UID of the identity provider to federate. Must be an existing identity provider
	// in this account. One-time pin and Cloudflare-managed identity providers cannot
	// be federated.
	IdPID param.Field[string] `json:"idp_id" api:"required" format:"uuid"`
}

func (r AccessIdPFederationGrantNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessIdPFederationGrantNewResponseEnvelope struct {
	Errors   []AccessIdPFederationGrantNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessIdPFederationGrantNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessIdPFederationGrantNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  IdPFederationGrant                                 `json:"result"`
	JSON    accessIdPFederationGrantNewResponseEnvelopeJSON    `json:"-"`
}

// accessIdPFederationGrantNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessIdPFederationGrantNewResponseEnvelope]
type accessIdPFederationGrantNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantNewResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           AccessIdPFederationGrantNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessIdPFederationGrantNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessIdPFederationGrantNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessIdPFederationGrantNewResponseEnvelopeErrors]
type accessIdPFederationGrantNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessIdPFederationGrantNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    accessIdPFederationGrantNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessIdPFederationGrantNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessIdPFederationGrantNewResponseEnvelopeErrorsSource]
type accessIdPFederationGrantNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantNewResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccessIdPFederationGrantNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessIdPFederationGrantNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessIdPFederationGrantNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessIdPFederationGrantNewResponseEnvelopeMessages]
type accessIdPFederationGrantNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessIdPFederationGrantNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accessIdPFederationGrantNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessIdPFederationGrantNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessIdPFederationGrantNewResponseEnvelopeMessagesSource]
type accessIdPFederationGrantNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// AccessIdPFederationGrantNewResponseEnvelopeSuccess indicates whether the API call was successful.
type AccessIdPFederationGrantNewResponseEnvelopeSuccess bool

const (
	AccessIdPFederationGrantNewResponseEnvelopeSuccessTrue AccessIdPFederationGrantNewResponseEnvelopeSuccess = true
)

func (r AccessIdPFederationGrantNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessIdPFederationGrantNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessIdPFederationGrantListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessIdPFederationGrantListResponseEnvelope struct {
	Errors   []AccessIdPFederationGrantListResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessIdPFederationGrantListResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success    AccessIdPFederationGrantListResponseEnvelopeSuccess    `json:"success" api:"required"`
	Result     []IdPFederationGrant                                   `json:"result"`
	ResultInfo AccessIdPFederationGrantListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessIdPFederationGrantListResponseEnvelopeJSON       `json:"-"`
}

// accessIdPFederationGrantListResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessIdPFederationGrantListResponseEnvelope]
type accessIdPFederationGrantListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantListResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccessIdPFederationGrantListResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessIdPFederationGrantListResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessIdPFederationGrantListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessIdPFederationGrantListResponseEnvelopeErrors]
type accessIdPFederationGrantListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessIdPFederationGrantListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantListResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accessIdPFederationGrantListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessIdPFederationGrantListResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessIdPFederationGrantListResponseEnvelopeErrorsSource]
type accessIdPFederationGrantListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantListResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AccessIdPFederationGrantListResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessIdPFederationGrantListResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessIdPFederationGrantListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessIdPFederationGrantListResponseEnvelopeMessages]
type accessIdPFederationGrantListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessIdPFederationGrantListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantListResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    accessIdPFederationGrantListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessIdPFederationGrantListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessIdPFederationGrantListResponseEnvelopeMessagesSource]
type accessIdPFederationGrantListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// AccessIdPFederationGrantListResponseEnvelopeSuccess indicates whether the API call was successful.
type AccessIdPFederationGrantListResponseEnvelopeSuccess bool

const (
	AccessIdPFederationGrantListResponseEnvelopeSuccessTrue AccessIdPFederationGrantListResponseEnvelopeSuccess = true
)

func (r AccessIdPFederationGrantListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessIdPFederationGrantListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessIdPFederationGrantListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64 `json:"total_count"`
	// The number of total pages in the entire result set.
	TotalPages float64                                                    `json:"total_pages"`
	JSON       accessIdPFederationGrantListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessIdPFederationGrantListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AccessIdPFederationGrantListResponseEnvelopeResultInfo]
type accessIdPFederationGrantListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessIdPFederationGrantDeleteResponseEnvelope struct {
	Errors   []AccessIdPFederationGrantDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessIdPFederationGrantDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessIdPFederationGrantDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessIdPFederationGrantDeleteResponse                `json:"result"`
	JSON    accessIdPFederationGrantDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessIdPFederationGrantDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessIdPFederationGrantDeleteResponseEnvelope]
type accessIdPFederationGrantDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantDeleteResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AccessIdPFederationGrantDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessIdPFederationGrantDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessIdPFederationGrantDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessIdPFederationGrantDeleteResponseEnvelopeErrors]
type accessIdPFederationGrantDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessIdPFederationGrantDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    accessIdPFederationGrantDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessIdPFederationGrantDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessIdPFederationGrantDeleteResponseEnvelopeErrorsSource]
type accessIdPFederationGrantDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantDeleteResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AccessIdPFederationGrantDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessIdPFederationGrantDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessIdPFederationGrantDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessIdPFederationGrantDeleteResponseEnvelopeMessages]
type accessIdPFederationGrantDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessIdPFederationGrantDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    accessIdPFederationGrantDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessIdPFederationGrantDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessIdPFederationGrantDeleteResponseEnvelopeMessagesSource]
type accessIdPFederationGrantDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// AccessIdPFederationGrantDeleteResponseEnvelopeSuccess indicates whether the API call was successful.
type AccessIdPFederationGrantDeleteResponseEnvelopeSuccess bool

const (
	AccessIdPFederationGrantDeleteResponseEnvelopeSuccessTrue AccessIdPFederationGrantDeleteResponseEnvelopeSuccess = true
)

func (r AccessIdPFederationGrantDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessIdPFederationGrantDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessIdPFederationGrantGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessIdPFederationGrantGetResponseEnvelope struct {
	Errors   []AccessIdPFederationGrantGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessIdPFederationGrantGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessIdPFederationGrantGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  IdPFederationGrant                                 `json:"result"`
	JSON    accessIdPFederationGrantGetResponseEnvelopeJSON    `json:"-"`
}

// accessIdPFederationGrantGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessIdPFederationGrantGetResponseEnvelope]
type accessIdPFederationGrantGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantGetResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           AccessIdPFederationGrantGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessIdPFederationGrantGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessIdPFederationGrantGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessIdPFederationGrantGetResponseEnvelopeErrors]
type accessIdPFederationGrantGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessIdPFederationGrantGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    accessIdPFederationGrantGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessIdPFederationGrantGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessIdPFederationGrantGetResponseEnvelopeErrorsSource]
type accessIdPFederationGrantGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantGetResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccessIdPFederationGrantGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessIdPFederationGrantGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessIdPFederationGrantGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessIdPFederationGrantGetResponseEnvelopeMessages]
type accessIdPFederationGrantGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessIdPFederationGrantGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessIdPFederationGrantGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accessIdPFederationGrantGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessIdPFederationGrantGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessIdPFederationGrantGetResponseEnvelopeMessagesSource]
type accessIdPFederationGrantGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessIdPFederationGrantGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdPFederationGrantGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// AccessIdPFederationGrantGetResponseEnvelopeSuccess indicates whether the API call was successful.
type AccessIdPFederationGrantGetResponseEnvelopeSuccess bool

const (
	AccessIdPFederationGrantGetResponseEnvelopeSuccessTrue AccessIdPFederationGrantGetResponseEnvelopeSuccess = true
)

func (r AccessIdPFederationGrantGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessIdPFederationGrantGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
