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

// AdvancedTCPProtectionPrefixService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionPrefixService] method instead.
type AdvancedTCPProtectionPrefixService struct {
	Options []option.RequestOption
	Items   *AdvancedTCPProtectionPrefixItemService
}

// NewAdvancedTCPProtectionPrefixService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAdvancedTCPProtectionPrefixService(opts ...option.RequestOption) (r *AdvancedTCPProtectionPrefixService) {
	r = &AdvancedTCPProtectionPrefixService{}
	r.Options = opts
	r.Items = NewAdvancedTCPProtectionPrefixItemService(opts...)
	return
}

// Create a prefix for an account.
func (r *AdvancedTCPProtectionPrefixService) New(ctx context.Context, params AdvancedTCPProtectionPrefixNewParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionPrefixNewResponse, err error) {
	var env AdvancedTCPProtectionPrefixNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/prefixes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all prefixes for an account.
func (r *AdvancedTCPProtectionPrefixService) List(ctx context.Context, params AdvancedTCPProtectionPrefixListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AdvancedTCPProtectionPrefixListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/prefixes", params.AccountID)
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

// List all prefixes for an account.
func (r *AdvancedTCPProtectionPrefixService) ListAutoPaging(ctx context.Context, params AdvancedTCPProtectionPrefixListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AdvancedTCPProtectionPrefixListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Create multiple prefixes for an account.
func (r *AdvancedTCPProtectionPrefixService) BulkNew(ctx context.Context, params AdvancedTCPProtectionPrefixBulkNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[AdvancedTCPProtectionPrefixBulkNewResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/prefixes/bulk", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Create multiple prefixes for an account.
func (r *AdvancedTCPProtectionPrefixService) BulkNewAutoPaging(ctx context.Context, params AdvancedTCPProtectionPrefixBulkNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AdvancedTCPProtectionPrefixBulkNewResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkNew(ctx, params, opts...))
}

// Delete all prefixes for an account.
func (r *AdvancedTCPProtectionPrefixService) BulkDelete(ctx context.Context, body AdvancedTCPProtectionPrefixBulkDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionPrefixBulkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/prefixes", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AdvancedTCPProtectionPrefixNewResponse struct {
	// The unique ID of the prefix.
	ID string `json:"id" api:"required"`
	// A comment describing the prefix.
	Comment string `json:"comment" api:"required"`
	// The creation timestamp of the prefix.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Whether to exclude the prefix from protection.
	Excluded bool `json:"excluded" api:"required"`
	// The last modification timestamp of the prefix.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The prefix in CIDR format.
	Prefix string                                     `json:"prefix" api:"required"`
	JSON   advancedTCPProtectionPrefixNewResponseJSON `json:"-"`
}

// advancedTCPProtectionPrefixNewResponseJSON contains the JSON metadata for the
// struct [AdvancedTCPProtectionPrefixNewResponse]
type advancedTCPProtectionPrefixNewResponseJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Excluded    apijson.Field
	ModifiedOn  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixNewResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixListResponse struct {
	// The unique ID of the prefix.
	ID string `json:"id" api:"required"`
	// A comment describing the prefix.
	Comment string `json:"comment" api:"required"`
	// The creation timestamp of the prefix.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Whether to exclude the prefix from protection.
	Excluded bool `json:"excluded" api:"required"`
	// The last modification timestamp of the prefix.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The prefix in CIDR format.
	Prefix string                                      `json:"prefix" api:"required"`
	JSON   advancedTCPProtectionPrefixListResponseJSON `json:"-"`
}

// advancedTCPProtectionPrefixListResponseJSON contains the JSON metadata for the
// struct [AdvancedTCPProtectionPrefixListResponse]
type advancedTCPProtectionPrefixListResponseJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Excluded    apijson.Field
	ModifiedOn  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixListResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixBulkNewResponse struct {
	// The unique ID of the prefix.
	ID string `json:"id" api:"required"`
	// A comment describing the prefix.
	Comment string `json:"comment" api:"required"`
	// The creation timestamp of the prefix.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Whether to exclude the prefix from protection.
	Excluded bool `json:"excluded" api:"required"`
	// The last modification timestamp of the prefix.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The prefix in CIDR format.
	Prefix string                                         `json:"prefix" api:"required"`
	JSON   advancedTCPProtectionPrefixBulkNewResponseJSON `json:"-"`
}

// advancedTCPProtectionPrefixBulkNewResponseJSON contains the JSON metadata for
// the struct [AdvancedTCPProtectionPrefixBulkNewResponse]
type advancedTCPProtectionPrefixBulkNewResponseJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Excluded    apijson.Field
	ModifiedOn  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixBulkDeleteResponse struct {
	Errors   []AdvancedTCPProtectionPrefixBulkDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionPrefixBulkDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionPrefixBulkDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionPrefixBulkDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionPrefixBulkDeleteResponseJSON contains the JSON metadata for
// the struct [AdvancedTCPProtectionPrefixBulkDeleteResponse]
type advancedTCPProtectionPrefixBulkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixBulkDeleteResponseError struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AdvancedTCPProtectionPrefixBulkDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionPrefixBulkDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionPrefixBulkDeleteResponseErrorJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionPrefixBulkDeleteResponseError]
type advancedTCPProtectionPrefixBulkDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixBulkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixBulkDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixBulkDeleteResponseErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    advancedTCPProtectionPrefixBulkDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionPrefixBulkDeleteResponseErrorsSourceJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionPrefixBulkDeleteResponseErrorsSource]
type advancedTCPProtectionPrefixBulkDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixBulkDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixBulkDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixBulkDeleteResponseMessage struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AdvancedTCPProtectionPrefixBulkDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionPrefixBulkDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionPrefixBulkDeleteResponseMessageJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionPrefixBulkDeleteResponseMessage]
type advancedTCPProtectionPrefixBulkDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixBulkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixBulkDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixBulkDeleteResponseMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    advancedTCPProtectionPrefixBulkDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionPrefixBulkDeleteResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionPrefixBulkDeleteResponseMessagesSource]
type advancedTCPProtectionPrefixBulkDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixBulkDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixBulkDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionPrefixBulkDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionPrefixBulkDeleteResponseSuccessTrue AdvancedTCPProtectionPrefixBulkDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionPrefixBulkDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionPrefixBulkDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionPrefixNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A comment describing the prefix.
	Comment param.Field[string] `json:"comment" api:"required"`
	// Whether to exclude the prefix from protection.
	Excluded param.Field[bool] `json:"excluded" api:"required"`
	// The prefix to add in CIDR format.
	Prefix param.Field[string] `json:"prefix" api:"required"`
}

func (r AdvancedTCPProtectionPrefixNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionPrefixNewResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionPrefixNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionPrefixNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionPrefixNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionPrefixNewResponse                `json:"result"`
	JSON    advancedTCPProtectionPrefixNewResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionPrefixNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [AdvancedTCPProtectionPrefixNewResponseEnvelope]
type advancedTCPProtectionPrefixNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixNewResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AdvancedTCPProtectionPrefixNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionPrefixNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionPrefixNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionPrefixNewResponseEnvelopeErrors]
type advancedTCPProtectionPrefixNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    advancedTCPProtectionPrefixNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionPrefixNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionPrefixNewResponseEnvelopeErrorsSource]
type advancedTCPProtectionPrefixNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixNewResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AdvancedTCPProtectionPrefixNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionPrefixNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionPrefixNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionPrefixNewResponseEnvelopeMessages]
type advancedTCPProtectionPrefixNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionPrefixNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    advancedTCPProtectionPrefixNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionPrefixNewResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionPrefixNewResponseEnvelopeMessagesSource]
type advancedTCPProtectionPrefixNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionPrefixNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionPrefixNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionPrefixNewResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionPrefixNewResponseEnvelopeSuccessTrue AdvancedTCPProtectionPrefixNewResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionPrefixNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionPrefixNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionPrefixListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The direction of ordering (ASC or DESC). Defaults to 'ASC'.
	Direction param.Field[string] `query:"direction"`
	// The field to order by. Defaults to 'prefix'.
	Order param.Field[string] `query:"order"`
	// The page number for pagination. Defaults to 1.
	Page param.Field[int64] `query:"page"`
	// The number of items per page. Must be between 10 and 1000. Defaults to 25.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [AdvancedTCPProtectionPrefixListParams]'s query parameters
// as `url.Values`.
func (r AdvancedTCPProtectionPrefixListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AdvancedTCPProtectionPrefixBulkNewParams struct {
	// Identifier.
	AccountID param.Field[string]                            `path:"account_id" api:"required"`
	Body      []AdvancedTCPProtectionPrefixBulkNewParamsBody `json:"body" api:"required"`
}

func (r AdvancedTCPProtectionPrefixBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AdvancedTCPProtectionPrefixBulkNewParamsBody struct {
	// A comment describing the prefix.
	Comment param.Field[string] `json:"comment" api:"required"`
	// Whether to exclude the prefix from protection.
	Excluded param.Field[bool] `json:"excluded" api:"required"`
	// The prefix to add in CIDR format.
	Prefix param.Field[string] `json:"prefix" api:"required"`
}

func (r AdvancedTCPProtectionPrefixBulkNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionPrefixBulkDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
