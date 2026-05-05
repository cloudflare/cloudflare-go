// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// ResourceLibraryCategoryService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceLibraryCategoryService] method instead.
type ResourceLibraryCategoryService struct {
	Options []option.RequestOption
}

// NewResourceLibraryCategoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewResourceLibraryCategoryService(opts ...option.RequestOption) (r *ResourceLibraryCategoryService) {
	r = &ResourceLibraryCategoryService{}
	r.Options = opts
	return
}

// List application categories.
func (r *ResourceLibraryCategoryService) List(ctx context.Context, params ResourceLibraryCategoryListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ResourceLibraryCategoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/resource-library/categories", params.AccountID)
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

// List application categories.
func (r *ResourceLibraryCategoryService) ListAutoPaging(ctx context.Context, params ResourceLibraryCategoryListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ResourceLibraryCategoryListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Get application category by ID.
func (r *ResourceLibraryCategoryService) Get(ctx context.Context, id string, query ResourceLibraryCategoryGetParams, opts ...option.RequestOption) (res *ResourceLibraryCategoryGetResponse, err error) {
	var env ResourceLibraryCategoryGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/resource-library/categories/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ResourceLibraryCategoryListResponse struct {
	// Returns the category ID.
	ID string `json:"id" api:"required"`
	// Returns the category creation time.
	CreatedAt string `json:"created_at" api:"required"`
	// Returns the category description.
	Description string `json:"description" api:"required"`
	// Returns the category name.
	Name string                                  `json:"name" api:"required"`
	JSON resourceLibraryCategoryListResponseJSON `json:"-"`
}

// resourceLibraryCategoryListResponseJSON contains the JSON metadata for the
// struct [ResourceLibraryCategoryListResponse]
type resourceLibraryCategoryListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryCategoryListResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryCategoryGetResponse struct {
	// Returns the category ID.
	ID string `json:"id" api:"required"`
	// Returns the category creation time.
	CreatedAt string `json:"created_at" api:"required"`
	// Returns the category description.
	Description string `json:"description" api:"required"`
	// Returns the category name.
	Name string                                 `json:"name" api:"required"`
	JSON resourceLibraryCategoryGetResponseJSON `json:"-"`
}

// resourceLibraryCategoryGetResponseJSON contains the JSON metadata for the struct
// [ResourceLibraryCategoryGetResponse]
type resourceLibraryCategoryGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryCategoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryCategoryGetResponseJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryCategoryListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Limit of number of results to return.
	Limit param.Field[int64] `query:"limit"`
	// Offset of results to return.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [ResourceLibraryCategoryListParams]'s query parameters as
// `url.Values`.
func (r ResourceLibraryCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ResourceLibraryCategoryGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ResourceLibraryCategoryGetResponseEnvelope struct {
	Errors   []ResourceLibraryCategoryGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ResourceLibraryCategoryGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Indicates whether the API call was successful.
	Success ResourceLibraryCategoryGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  ResourceLibraryCategoryGetResponse                `json:"result"`
	JSON    resourceLibraryCategoryGetResponseEnvelopeJSON    `json:"-"`
}

// resourceLibraryCategoryGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ResourceLibraryCategoryGetResponseEnvelope]
type resourceLibraryCategoryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryCategoryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryCategoryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryCategoryGetResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           ResourceLibraryCategoryGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             resourceLibraryCategoryGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// resourceLibraryCategoryGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ResourceLibraryCategoryGetResponseEnvelopeErrors]
type resourceLibraryCategoryGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceLibraryCategoryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryCategoryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryCategoryGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    resourceLibraryCategoryGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// resourceLibraryCategoryGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ResourceLibraryCategoryGetResponseEnvelopeErrorsSource]
type resourceLibraryCategoryGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryCategoryGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryCategoryGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryCategoryGetResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           ResourceLibraryCategoryGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             resourceLibraryCategoryGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// resourceLibraryCategoryGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ResourceLibraryCategoryGetResponseEnvelopeMessages]
type resourceLibraryCategoryGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResourceLibraryCategoryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryCategoryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ResourceLibraryCategoryGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    resourceLibraryCategoryGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// resourceLibraryCategoryGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [ResourceLibraryCategoryGetResponseEnvelopeMessagesSource]
type resourceLibraryCategoryGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResourceLibraryCategoryGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resourceLibraryCategoryGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful.
type ResourceLibraryCategoryGetResponseEnvelopeSuccess bool

const (
	ResourceLibraryCategoryGetResponseEnvelopeSuccessTrue ResourceLibraryCategoryGetResponseEnvelopeSuccess = true
)

func (r ResourceLibraryCategoryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ResourceLibraryCategoryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
