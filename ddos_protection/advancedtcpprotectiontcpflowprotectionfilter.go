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

// AdvancedTCPProtectionTCPFlowProtectionFilterService contains methods and other
// services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionTCPFlowProtectionFilterService] method instead.
type AdvancedTCPProtectionTCPFlowProtectionFilterService struct {
	Options []option.RequestOption
	Items   *AdvancedTCPProtectionTCPFlowProtectionFilterItemService
}

// NewAdvancedTCPProtectionTCPFlowProtectionFilterService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAdvancedTCPProtectionTCPFlowProtectionFilterService(opts ...option.RequestOption) (r *AdvancedTCPProtectionTCPFlowProtectionFilterService) {
	r = &AdvancedTCPProtectionTCPFlowProtectionFilterService{}
	r.Options = opts
	r.Items = NewAdvancedTCPProtectionTCPFlowProtectionFilterItemService(opts...)
	return
}

// Create a TCP Flow Protection filter for an account.
func (r *AdvancedTCPProtectionTCPFlowProtectionFilterService) New(ctx context.Context, params AdvancedTCPProtectionTCPFlowProtectionFilterNewParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionTCPFlowProtectionFilterNewResponse, err error) {
	var env AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all TCP Flow Protection filters for an account.
func (r *AdvancedTCPProtectionTCPFlowProtectionFilterService) List(ctx context.Context, params AdvancedTCPProtectionTCPFlowProtectionFilterListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AdvancedTCPProtectionTCPFlowProtectionFilterListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters", params.AccountID)
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

// List all TCP Flow Protection filters for an account.
func (r *AdvancedTCPProtectionTCPFlowProtectionFilterService) ListAutoPaging(ctx context.Context, params AdvancedTCPProtectionTCPFlowProtectionFilterListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AdvancedTCPProtectionTCPFlowProtectionFilterListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete all TCP Flow Protection filters for an account.
func (r *AdvancedTCPProtectionTCPFlowProtectionFilterService) BulkDelete(ctx context.Context, body AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/filters", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AdvancedTCPProtectionTCPFlowProtectionFilterNewResponse struct {
	// The unique ID of the expression filter.
	ID string `json:"id" api:"required"`
	// The creation timestamp of the expression filter.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The filter expression.
	Expression string `json:"expression" api:"required"`
	// The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the expression filter.
	ModifiedOn time.Time                                                   `json:"modified_on" api:"required" format:"date-time"`
	JSON       advancedTCPProtectionTCPFlowProtectionFilterNewResponseJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterNewResponseJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterNewResponse]
type advancedTCPProtectionTCPFlowProtectionFilterNewResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Expression  apijson.Field
	Mode        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterNewResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterListResponse struct {
	// The unique ID of the expression filter.
	ID string `json:"id" api:"required"`
	// The creation timestamp of the expression filter.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The filter expression.
	Expression string `json:"expression" api:"required"`
	// The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the expression filter.
	ModifiedOn time.Time                                                    `json:"modified_on" api:"required" format:"date-time"`
	JSON       advancedTCPProtectionTCPFlowProtectionFilterListResponseJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterListResponseJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterListResponse]
type advancedTCPProtectionTCPFlowProtectionFilterListResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Expression  apijson.Field
	Mode        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterListResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponse struct {
	Errors   []AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponse]
type advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseError struct {
	Code             int64                                                                      `json:"code" api:"required"`
	Message          string                                                                     `json:"message" api:"required"`
	DocumentationURL string                                                                     `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseError]
type advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorsSource struct {
	Pointer string                                                                         `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorsSource]
type advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessage struct {
	Code             int64                                                                        `json:"code" api:"required"`
	Message          string                                                                       `json:"message" api:"required"`
	DocumentationURL string                                                                       `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessageJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessage]
type advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessagesSource struct {
	Pointer string                                                                           `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessagesSource]
type advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseSuccessTrue AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionTCPFlowProtectionFilterNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The filter expression.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The filter's mode. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode param.Field[string] `json:"mode" api:"required"`
}

func (r AdvancedTCPProtectionTCPFlowProtectionFilterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionTCPFlowProtectionFilterNewResponse                `json:"result"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelope]
type advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrors struct {
	Code             int64                                                                       `json:"code" api:"required"`
	Message          string                                                                      `json:"message" api:"required"`
	DocumentationURL string                                                                      `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrors]
type advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                                          `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsSource]
type advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessages struct {
	Code             int64                                                                         `json:"code" api:"required"`
	Message          string                                                                        `json:"message" api:"required"`
	DocumentationURL string                                                                        `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessages]
type advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                                            `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesSource]
type advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeSuccessTrue AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionTCPFlowProtectionFilterNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionTCPFlowProtectionFilterListParams struct {
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

// URLQuery serializes [AdvancedTCPProtectionTCPFlowProtectionFilterListParams]'s
// query parameters as `url.Values`.
func (r AdvancedTCPProtectionTCPFlowProtectionFilterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AdvancedTCPProtectionTCPFlowProtectionFilterBulkDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
