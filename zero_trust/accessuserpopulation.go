// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AccessUserPopulationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessUserPopulationService] method instead.
type AccessUserPopulationService struct {
	Options []option.RequestOption
	Users   *AccessUserPopulationUserService
}

// NewAccessUserPopulationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessUserPopulationService(opts ...option.RequestOption) (r *AccessUserPopulationService) {
	r = &AccessUserPopulationService{}
	r.Options = opts
	r.Users = NewAccessUserPopulationUserService(opts...)
	return
}

// Creates a user population in an account. Requires permission to update the
// account's Access organization.
func (r *AccessUserPopulationService) New(ctx context.Context, params AccessUserPopulationNewParams, opts ...option.RequestOption) (res *AccessUserPopulationNewResponse, err error) {
	var env AccessUserPopulationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a user population's name.
func (r *AccessUserPopulationService) Update(ctx context.Context, userPopulationID string, params AccessUserPopulationUpdateParams, opts ...option.RequestOption) (res *AccessUserPopulationUpdateResponse, err error) {
	var env AccessUserPopulationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userPopulationID == "" {
		err = errors.New("missing required user_population_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations/%s", params.AccountID, userPopulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists the user populations in an account.
func (r *AccessUserPopulationService) List(ctx context.Context, params AccessUserPopulationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessUserPopulationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists the user populations in an account.
func (r *AccessUserPopulationService) ListAutoPaging(ctx context.Context, params AccessUserPopulationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessUserPopulationListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a user population that is not attached to an application.
func (r *AccessUserPopulationService) Delete(ctx context.Context, userPopulationID string, body AccessUserPopulationDeleteParams, opts ...option.RequestOption) (res *AccessUserPopulationDeleteResponse, err error) {
	var env AccessUserPopulationDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userPopulationID == "" {
		err = errors.New("missing required user_population_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations/%s", body.AccountID, userPopulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches a user population.
func (r *AccessUserPopulationService) Get(ctx context.Context, userPopulationID string, query AccessUserPopulationGetParams, opts ...option.RequestOption) (res *AccessUserPopulationGetResponse, err error) {
	var env AccessUserPopulationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userPopulationID == "" {
		err = errors.New("missing required user_population_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/user_populations/%s", query.AccountID, userPopulationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AccessUserPopulationNewResponse struct {
	// UUID.
	ID string `json:"id" api:"required"`
	// The name of the user population. Leading and trailing whitespace is trimmed
	// before validation. After trimming, the name must contain at least one
	// non-whitespace character and must not exceed 128 UTF-8 bytes.
	Name string                              `json:"name" api:"required"`
	JSON accessUserPopulationNewResponseJSON `json:"-"`
}

// accessUserPopulationNewResponseJSON contains the JSON metadata for the struct
// [AccessUserPopulationNewResponse]
type accessUserPopulationNewResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUpdateResponse struct {
	// UUID.
	ID string `json:"id" api:"required"`
	// The name of the user population. Leading and trailing whitespace is trimmed
	// before validation. After trimming, the name must contain at least one
	// non-whitespace character and must not exceed 128 UTF-8 bytes.
	Name string                                 `json:"name" api:"required"`
	JSON accessUserPopulationUpdateResponseJSON `json:"-"`
}

// accessUserPopulationUpdateResponseJSON contains the JSON metadata for the struct
// [AccessUserPopulationUpdateResponse]
type accessUserPopulationUpdateResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationListResponse struct {
	// UUID.
	ID string `json:"id" api:"required"`
	// The name of the user population. Leading and trailing whitespace is trimmed
	// before validation. After trimming, the name must contain at least one
	// non-whitespace character and must not exceed 128 UTF-8 bytes.
	Name string                               `json:"name" api:"required"`
	JSON accessUserPopulationListResponseJSON `json:"-"`
}

// accessUserPopulationListResponseJSON contains the JSON metadata for the struct
// [AccessUserPopulationListResponse]
type accessUserPopulationListResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationDeleteResponse struct {
	// UUID.
	ID   string                                 `json:"id"`
	JSON accessUserPopulationDeleteResponseJSON `json:"-"`
}

// accessUserPopulationDeleteResponseJSON contains the JSON metadata for the struct
// [AccessUserPopulationDeleteResponse]
type accessUserPopulationDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationGetResponse struct {
	// UUID.
	ID string `json:"id" api:"required"`
	// The name of the user population. Leading and trailing whitespace is trimmed
	// before validation. After trimming, the name must contain at least one
	// non-whitespace character and must not exceed 128 UTF-8 bytes.
	Name string                              `json:"name" api:"required"`
	JSON accessUserPopulationGetResponseJSON `json:"-"`
}

// accessUserPopulationGetResponseJSON contains the JSON metadata for the struct
// [AccessUserPopulationGetResponse]
type accessUserPopulationGetResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The name of the user population. Leading and trailing whitespace is trimmed
	// before validation. After trimming, the name must contain at least one
	// non-whitespace character and must not exceed 128 UTF-8 bytes.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r AccessUserPopulationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationNewResponseEnvelope struct {
	Errors   []AccessUserPopulationNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserPopulationNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserPopulationNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessUserPopulationNewResponse                `json:"result"`
	JSON    accessUserPopulationNewResponseEnvelopeJSON    `json:"-"`
}

// accessUserPopulationNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessUserPopulationNewResponseEnvelope]
type accessUserPopulationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationNewResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AccessUserPopulationNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserPopulationNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserPopulationNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessUserPopulationNewResponseEnvelopeErrors]
type accessUserPopulationNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    accessUserPopulationNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserPopulationNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessUserPopulationNewResponseEnvelopeErrorsSource]
type accessUserPopulationNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationNewResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AccessUserPopulationNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserPopulationNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserPopulationNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessUserPopulationNewResponseEnvelopeMessages]
type accessUserPopulationNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    accessUserPopulationNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserPopulationNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessUserPopulationNewResponseEnvelopeMessagesSource]
type accessUserPopulationNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserPopulationNewResponseEnvelopeSuccess bool

const (
	AccessUserPopulationNewResponseEnvelopeSuccessTrue AccessUserPopulationNewResponseEnvelopeSuccess = true
)

func (r AccessUserPopulationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserPopulationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessUserPopulationUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The name of the user population. Leading and trailing whitespace is trimmed
	// before validation. After trimming, the name must contain at least one
	// non-whitespace character and must not exceed 128 UTF-8 bytes.
	Name param.Field[string] `json:"name" api:"required"`
}

func (r AccessUserPopulationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessUserPopulationUpdateResponseEnvelope struct {
	Errors   []AccessUserPopulationUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserPopulationUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserPopulationUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessUserPopulationUpdateResponse                `json:"result"`
	JSON    accessUserPopulationUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessUserPopulationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessUserPopulationUpdateResponseEnvelope]
type accessUserPopulationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUpdateResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessUserPopulationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserPopulationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserPopulationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessUserPopulationUpdateResponseEnvelopeErrors]
type accessUserPopulationUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessUserPopulationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserPopulationUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessUserPopulationUpdateResponseEnvelopeErrorsSource]
type accessUserPopulationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUpdateResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccessUserPopulationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserPopulationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserPopulationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserPopulationUpdateResponseEnvelopeMessages]
type accessUserPopulationUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accessUserPopulationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserPopulationUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessUserPopulationUpdateResponseEnvelopeMessagesSource]
type accessUserPopulationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserPopulationUpdateResponseEnvelopeSuccess bool

const (
	AccessUserPopulationUpdateResponseEnvelopeSuccessTrue AccessUserPopulationUpdateResponseEnvelopeSuccess = true
)

func (r AccessUserPopulationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserPopulationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessUserPopulationListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Page number of results.
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [AccessUserPopulationListParams]'s query parameters as
// `url.Values`.
func (r AccessUserPopulationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessUserPopulationDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessUserPopulationDeleteResponseEnvelope struct {
	Errors   []AccessUserPopulationDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserPopulationDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserPopulationDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessUserPopulationDeleteResponse                `json:"result"`
	JSON    accessUserPopulationDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessUserPopulationDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessUserPopulationDeleteResponseEnvelope]
type accessUserPopulationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationDeleteResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessUserPopulationDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserPopulationDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserPopulationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessUserPopulationDeleteResponseEnvelopeErrors]
type accessUserPopulationDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessUserPopulationDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserPopulationDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessUserPopulationDeleteResponseEnvelopeErrorsSource]
type accessUserPopulationDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationDeleteResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccessUserPopulationDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserPopulationDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserPopulationDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserPopulationDeleteResponseEnvelopeMessages]
type accessUserPopulationDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accessUserPopulationDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserPopulationDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessUserPopulationDeleteResponseEnvelopeMessagesSource]
type accessUserPopulationDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserPopulationDeleteResponseEnvelopeSuccess bool

const (
	AccessUserPopulationDeleteResponseEnvelopeSuccessTrue AccessUserPopulationDeleteResponseEnvelopeSuccess = true
)

func (r AccessUserPopulationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserPopulationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessUserPopulationGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessUserPopulationGetResponseEnvelope struct {
	Errors   []AccessUserPopulationGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessUserPopulationGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessUserPopulationGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessUserPopulationGetResponse                `json:"result"`
	JSON    accessUserPopulationGetResponseEnvelopeJSON    `json:"-"`
}

// accessUserPopulationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessUserPopulationGetResponseEnvelope]
type accessUserPopulationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AccessUserPopulationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserPopulationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserPopulationGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessUserPopulationGetResponseEnvelopeErrors]
type accessUserPopulationGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    accessUserPopulationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserPopulationGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessUserPopulationGetResponseEnvelopeErrorsSource]
type accessUserPopulationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AccessUserPopulationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserPopulationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserPopulationGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessUserPopulationGetResponseEnvelopeMessages]
type accessUserPopulationGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserPopulationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserPopulationGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    accessUserPopulationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserPopulationGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessUserPopulationGetResponseEnvelopeMessagesSource]
type accessUserPopulationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserPopulationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserPopulationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserPopulationGetResponseEnvelopeSuccess bool

const (
	AccessUserPopulationGetResponseEnvelopeSuccessTrue AccessUserPopulationGetResponseEnvelopeSuccess = true
)

func (r AccessUserPopulationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserPopulationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
