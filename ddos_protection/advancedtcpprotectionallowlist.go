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

// AdvancedTCPProtectionAllowlistService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionAllowlistService] method instead.
type AdvancedTCPProtectionAllowlistService struct {
	Options []option.RequestOption
	Items   *AdvancedTCPProtectionAllowlistItemService
}

// NewAdvancedTCPProtectionAllowlistService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAdvancedTCPProtectionAllowlistService(opts ...option.RequestOption) (r *AdvancedTCPProtectionAllowlistService) {
	r = &AdvancedTCPProtectionAllowlistService{}
	r.Options = opts
	r.Items = NewAdvancedTCPProtectionAllowlistItemService(opts...)
	return
}

// Create an allowlist prefix for an account.
func (r *AdvancedTCPProtectionAllowlistService) New(ctx context.Context, params AdvancedTCPProtectionAllowlistNewParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionAllowlistNewResponse, err error) {
	var env AdvancedTCPProtectionAllowlistNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/allowlist", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all allowlist prefixes for an account.
func (r *AdvancedTCPProtectionAllowlistService) List(ctx context.Context, params AdvancedTCPProtectionAllowlistListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AdvancedTCPProtectionAllowlistListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/allowlist", params.AccountID)
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

// List all allowlist prefixes for an account.
func (r *AdvancedTCPProtectionAllowlistService) ListAutoPaging(ctx context.Context, params AdvancedTCPProtectionAllowlistListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AdvancedTCPProtectionAllowlistListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete all allowlist prefixes for an account.
func (r *AdvancedTCPProtectionAllowlistService) BulkDelete(ctx context.Context, body AdvancedTCPProtectionAllowlistBulkDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionAllowlistBulkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/allowlist", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AdvancedTCPProtectionAllowlistNewResponse struct {
	// The unique ID of the allowlist prefix.
	ID string `json:"id" api:"required"`
	// An optional comment describing the allowlist prefix.
	Comment string `json:"comment" api:"required"`
	// The creation timestamp of the allowlist prefix.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Whether to enable the allowlist prefix into effect. Defaults to false.
	Enabled bool `json:"enabled" api:"required"`
	// The last modification timestamp of the allowlist prefix.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The allowlist prefix in CIDR format.
	Prefix string                                        `json:"prefix" api:"required"`
	JSON   advancedTCPProtectionAllowlistNewResponseJSON `json:"-"`
}

// advancedTCPProtectionAllowlistNewResponseJSON contains the JSON metadata for the
// struct [AdvancedTCPProtectionAllowlistNewResponse]
type advancedTCPProtectionAllowlistNewResponseJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Enabled     apijson.Field
	ModifiedOn  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistNewResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistListResponse struct {
	// The unique ID of the allowlist prefix.
	ID string `json:"id" api:"required"`
	// An optional comment describing the allowlist prefix.
	Comment string `json:"comment" api:"required"`
	// The creation timestamp of the allowlist prefix.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// Whether to enable the allowlist prefix into effect. Defaults to false.
	Enabled bool `json:"enabled" api:"required"`
	// The last modification timestamp of the allowlist prefix.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The allowlist prefix in CIDR format.
	Prefix string                                         `json:"prefix" api:"required"`
	JSON   advancedTCPProtectionAllowlistListResponseJSON `json:"-"`
}

// advancedTCPProtectionAllowlistListResponseJSON contains the JSON metadata for
// the struct [AdvancedTCPProtectionAllowlistListResponse]
type advancedTCPProtectionAllowlistListResponseJSON struct {
	ID          apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Enabled     apijson.Field
	ModifiedOn  apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistListResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistBulkDeleteResponse struct {
	Errors   []AdvancedTCPProtectionAllowlistBulkDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionAllowlistBulkDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionAllowlistBulkDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionAllowlistBulkDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionAllowlistBulkDeleteResponseJSON contains the JSON metadata
// for the struct [AdvancedTCPProtectionAllowlistBulkDeleteResponse]
type advancedTCPProtectionAllowlistBulkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistBulkDeleteResponseError struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AdvancedTCPProtectionAllowlistBulkDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionAllowlistBulkDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionAllowlistBulkDeleteResponseErrorJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionAllowlistBulkDeleteResponseError]
type advancedTCPProtectionAllowlistBulkDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistBulkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistBulkDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistBulkDeleteResponseErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    advancedTCPProtectionAllowlistBulkDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionAllowlistBulkDeleteResponseErrorsSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistBulkDeleteResponseErrorsSource]
type advancedTCPProtectionAllowlistBulkDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistBulkDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistBulkDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistBulkDeleteResponseMessage struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           AdvancedTCPProtectionAllowlistBulkDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionAllowlistBulkDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionAllowlistBulkDeleteResponseMessageJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionAllowlistBulkDeleteResponseMessage]
type advancedTCPProtectionAllowlistBulkDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistBulkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistBulkDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistBulkDeleteResponseMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    advancedTCPProtectionAllowlistBulkDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionAllowlistBulkDeleteResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistBulkDeleteResponseMessagesSource]
type advancedTCPProtectionAllowlistBulkDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistBulkDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistBulkDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionAllowlistBulkDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionAllowlistBulkDeleteResponseSuccessTrue AdvancedTCPProtectionAllowlistBulkDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionAllowlistBulkDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionAllowlistBulkDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionAllowlistNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// An comment describing the allowlist prefix.
	Comment param.Field[string] `json:"comment" api:"required"`
	// Whether to enable the allowlist prefix into effect.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// The allowlist prefix to add in CIDR format.
	Prefix param.Field[string] `json:"prefix" api:"required"`
}

func (r AdvancedTCPProtectionAllowlistNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionAllowlistNewResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionAllowlistNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionAllowlistNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionAllowlistNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionAllowlistNewResponse                `json:"result"`
	JSON    advancedTCPProtectionAllowlistNewResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionAllowlistNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [AdvancedTCPProtectionAllowlistNewResponseEnvelope]
type advancedTCPProtectionAllowlistNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistNewResponseEnvelopeErrors struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           AdvancedTCPProtectionAllowlistNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionAllowlistNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionAllowlistNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionAllowlistNewResponseEnvelopeErrors]
type advancedTCPProtectionAllowlistNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    advancedTCPProtectionAllowlistNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionAllowlistNewResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistNewResponseEnvelopeErrorsSource]
type advancedTCPProtectionAllowlistNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistNewResponseEnvelopeMessages struct {
	Code             int64                                                           `json:"code" api:"required"`
	Message          string                                                          `json:"message" api:"required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           AdvancedTCPProtectionAllowlistNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionAllowlistNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionAllowlistNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionAllowlistNewResponseEnvelopeMessages]
type advancedTCPProtectionAllowlistNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionAllowlistNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    advancedTCPProtectionAllowlistNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionAllowlistNewResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionAllowlistNewResponseEnvelopeMessagesSource]
type advancedTCPProtectionAllowlistNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionAllowlistNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionAllowlistNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionAllowlistNewResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionAllowlistNewResponseEnvelopeSuccessTrue AdvancedTCPProtectionAllowlistNewResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionAllowlistNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionAllowlistNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionAllowlistListParams struct {
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

// URLQuery serializes [AdvancedTCPProtectionAllowlistListParams]'s query
// parameters as `url.Values`.
func (r AdvancedTCPProtectionAllowlistListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AdvancedTCPProtectionAllowlistBulkDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
