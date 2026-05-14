// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ddos_protection

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AdvancedTCPProtectionSynProtectionFilterService contains methods and other
// services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionSynProtectionFilterService] method instead.
type AdvancedTCPProtectionSynProtectionFilterService struct {
	Options []option.RequestOption
	Items   *AdvancedTCPProtectionSynProtectionFilterItemService
}

// NewAdvancedTCPProtectionSynProtectionFilterService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAdvancedTCPProtectionSynProtectionFilterService(opts ...option.RequestOption) (r *AdvancedTCPProtectionSynProtectionFilterService) {
	r = &AdvancedTCPProtectionSynProtectionFilterService{}
	r.Options = opts
	r.Items = NewAdvancedTCPProtectionSynProtectionFilterItemService(opts...)
	return
}

// Create a SYN Protection filter for an account.
func (r *AdvancedTCPProtectionSynProtectionFilterService) New(ctx context.Context, params AdvancedTCPProtectionSynProtectionFilterNewParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionSynProtectionFilterNewResponse, err error) {
	var env AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/filters", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all SYN Protection filters for an account.
func (r *AdvancedTCPProtectionSynProtectionFilterService) List(ctx context.Context, params AdvancedTCPProtectionSynProtectionFilterListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AdvancedTCPProtectionSynProtectionFilterListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/filters", params.AccountID)
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

// List all SYN Protection filters for an account.
func (r *AdvancedTCPProtectionSynProtectionFilterService) ListAutoPaging(ctx context.Context, params AdvancedTCPProtectionSynProtectionFilterListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AdvancedTCPProtectionSynProtectionFilterListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete all SYN Protection filters for an account.
func (r *AdvancedTCPProtectionSynProtectionFilterService) BulkDelete(ctx context.Context, body AdvancedTCPProtectionSynProtectionFilterBulkDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/filters", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AdvancedTCPProtectionSynProtectionFilterNewResponse struct {
	// The unique ID of the expression filter.
	ID string `json:"id" api:"required"`
	// The creation timestamp of the expression filter.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The filter expression.
	Expression string `json:"expression" api:"required"`
	// The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the expression filter.
	ModifiedOn time.Time                                               `json:"modified_on" api:"required" format:"date-time"`
	JSON       advancedTCPProtectionSynProtectionFilterNewResponseJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterNewResponseJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionSynProtectionFilterNewResponse]
type advancedTCPProtectionSynProtectionFilterNewResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Expression  apijson.Field
	Mode        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterNewResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterListResponse struct {
	// The unique ID of the expression filter.
	ID string `json:"id" api:"required"`
	// The creation timestamp of the expression filter.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The filter expression.
	Expression string `json:"expression" api:"required"`
	// The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the expression filter.
	ModifiedOn time.Time                                                `json:"modified_on" api:"required" format:"date-time"`
	JSON       advancedTCPProtectionSynProtectionFilterListResponseJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterListResponseJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionSynProtectionFilterListResponse]
type advancedTCPProtectionSynProtectionFilterListResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Expression  apijson.Field
	Mode        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterListResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponse struct {
	Errors   []AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionSynProtectionFilterBulkDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterBulkDeleteResponseJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponse]
type advancedTCPProtectionSynProtectionFilterBulkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseError struct {
	Code             int64                                                                  `json:"code" api:"required"`
	Message          string                                                                 `json:"message" api:"required"`
	DocumentationURL string                                                                 `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseError]
type advancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorsSource struct {
	Pointer string                                                                     `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorsSource]
type advancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterBulkDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessage struct {
	Code             int64                                                                    `json:"code" api:"required"`
	Message          string                                                                   `json:"message" api:"required"`
	DocumentationURL string                                                                   `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessageJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessage]
type advancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessagesSource struct {
	Pointer string                                                                       `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessagesSource]
type advancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterBulkDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseSuccessTrue AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionSynProtectionFilterBulkDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionSynProtectionFilterNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The filter expression.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode param.Field[string] `json:"mode" api:"required"`
}

func (r AdvancedTCPProtectionSynProtectionFilterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionSynProtectionFilterNewResponse                `json:"result"`
	JSON    advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelope]
type advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrors struct {
	Code             int64                                                                   `json:"code" api:"required"`
	Message          string                                                                  `json:"message" api:"required"`
	DocumentationURL string                                                                  `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrors]
type advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                                      `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsSource]
type advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessages struct {
	Code             int64                                                                     `json:"code" api:"required"`
	Message          string                                                                    `json:"message" api:"required"`
	DocumentationURL string                                                                    `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessages]
type advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                                        `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesSource]
type advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionFilterNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeSuccessTrue AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionSynProtectionFilterNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionSynProtectionFilterListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The direction of ordering (ASC or DESC). Defaults to 'ASC'.
	Direction param.Field[string] `query:"direction"`
	// The mode of the filters to get. Optional. Valid values: 'enabled', 'disabled',
	// 'monitoring'.
	Mode param.Field[string] `query:"mode"`
	// The field to order by. Defaults to 'prefix'.
	Order param.Field[string] `query:"order"`
	// The page number for pagination. Defaults to 1.
	Page param.Field[int64] `query:"page"`
	// The number of items per page. Must be between 10 and 1000. Defaults to 25.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [AdvancedTCPProtectionSynProtectionFilterListParams]'s query
// parameters as `url.Values`.
func (r AdvancedTCPProtectionSynProtectionFilterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AdvancedTCPProtectionSynProtectionFilterBulkDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
