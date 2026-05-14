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

// AdvancedTCPProtectionTCPFlowProtectionRuleService contains methods and other
// services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionTCPFlowProtectionRuleService] method instead.
type AdvancedTCPProtectionTCPFlowProtectionRuleService struct {
	Options []option.RequestOption
	Items   *AdvancedTCPProtectionTCPFlowProtectionRuleItemService
}

// NewAdvancedTCPProtectionTCPFlowProtectionRuleService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAdvancedTCPProtectionTCPFlowProtectionRuleService(opts ...option.RequestOption) (r *AdvancedTCPProtectionTCPFlowProtectionRuleService) {
	r = &AdvancedTCPProtectionTCPFlowProtectionRuleService{}
	r.Options = opts
	r.Items = NewAdvancedTCPProtectionTCPFlowProtectionRuleItemService(opts...)
	return
}

// Create a TCP Flow Protection rule for an account.
func (r *AdvancedTCPProtectionTCPFlowProtectionRuleService) New(ctx context.Context, params AdvancedTCPProtectionTCPFlowProtectionRuleNewParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionTCPFlowProtectionRuleNewResponse, err error) {
	var env AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all TCP Flow Protection rules for an account.
func (r *AdvancedTCPProtectionTCPFlowProtectionRuleService) List(ctx context.Context, params AdvancedTCPProtectionTCPFlowProtectionRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AdvancedTCPProtectionTCPFlowProtectionRuleListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules", params.AccountID)
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

// List all TCP Flow Protection rules for an account.
func (r *AdvancedTCPProtectionTCPFlowProtectionRuleService) ListAutoPaging(ctx context.Context, params AdvancedTCPProtectionTCPFlowProtectionRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AdvancedTCPProtectionTCPFlowProtectionRuleListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete all TCP Flow Protection rules for an account.
func (r *AdvancedTCPProtectionTCPFlowProtectionRuleService) BulkDelete(ctx context.Context, body AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AdvancedTCPProtectionTCPFlowProtectionRuleNewResponse struct {
	// The unique ID of the TCP Flow Protection rule.
	ID string `json:"id" api:"required"`
	// The burst sensitivity. Must be one of 'low', 'medium', 'high'.
	BurstSensitivity string `json:"burst_sensitivity" api:"required"`
	// The creation timestamp of the TCP Flow Protection rule.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The mode for TCP Flow Protection. Must be one of 'enabled', 'disabled',
	// 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the TCP Flow Protection rule.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The name of the TCP Flow Protection rule. Value is relative to the 'scope'
	// setting. For 'global' scope, name should be 'global'. For either the 'region' or
	// 'datacenter' scope, name should be the actual name of the region or datacenter,
	// e.g., 'wnam' or 'lax'.
	Name string `json:"name" api:"required"`
	// The rate sensitivity. Must be one of 'low', 'medium', 'high'.
	RateSensitivity string `json:"rate_sensitivity" api:"required"`
	// The scope for the TCP Flow Protection rule. Must be one of 'global', 'region',
	// or 'datacenter'.
	Scope string                                                    `json:"scope" api:"required"`
	JSON  advancedTCPProtectionTCPFlowProtectionRuleNewResponseJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleNewResponseJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionTCPFlowProtectionRuleNewResponse]
type advancedTCPProtectionTCPFlowProtectionRuleNewResponseJSON struct {
	ID               apijson.Field
	BurstSensitivity apijson.Field
	CreatedOn        apijson.Field
	Mode             apijson.Field
	ModifiedOn       apijson.Field
	Name             apijson.Field
	RateSensitivity  apijson.Field
	Scope            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleListResponse struct {
	// The unique ID of the TCP Flow Protection rule.
	ID string `json:"id" api:"required"`
	// The burst sensitivity. Must be one of 'low', 'medium', 'high'.
	BurstSensitivity string `json:"burst_sensitivity" api:"required"`
	// The creation timestamp of the TCP Flow Protection rule.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The mode for TCP Flow Protection. Must be one of 'enabled', 'disabled',
	// 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the TCP Flow Protection rule.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The name of the TCP Flow Protection rule. Value is relative to the 'scope'
	// setting. For 'global' scope, name should be 'global'. For either the 'region' or
	// 'datacenter' scope, name should be the actual name of the region or datacenter,
	// e.g., 'wnam' or 'lax'.
	Name string `json:"name" api:"required"`
	// The rate sensitivity. Must be one of 'low', 'medium', 'high'.
	RateSensitivity string `json:"rate_sensitivity" api:"required"`
	// The scope for the TCP Flow Protection rule. Must be one of 'global', 'region',
	// or 'datacenter'.
	Scope string                                                     `json:"scope" api:"required"`
	JSON  advancedTCPProtectionTCPFlowProtectionRuleListResponseJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleListResponseJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionTCPFlowProtectionRuleListResponse]
type advancedTCPProtectionTCPFlowProtectionRuleListResponseJSON struct {
	ID               apijson.Field
	BurstSensitivity apijson.Field
	CreatedOn        apijson.Field
	Mode             apijson.Field
	ModifiedOn       apijson.Field
	Name             apijson.Field
	RateSensitivity  apijson.Field
	Scope            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponse struct {
	Errors   []AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponse]
type advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseError struct {
	Code             int64                                                                    `json:"code" api:"required"`
	Message          string                                                                   `json:"message" api:"required"`
	DocumentationURL string                                                                   `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseError]
type advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorsSource struct {
	Pointer string                                                                       `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorsSource]
type advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessage struct {
	Code             int64                                                                      `json:"code" api:"required"`
	Message          string                                                                     `json:"message" api:"required"`
	DocumentationURL string                                                                     `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessageJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessage]
type advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessagesSource struct {
	Pointer string                                                                         `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessagesSource]
type advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseSuccessTrue AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionTCPFlowProtectionRuleNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The burst sensitivity. Must be one of 'low', 'medium', 'high'.
	BurstSensitivity param.Field[string] `json:"burst_sensitivity" api:"required"`
	// The mode for the TCP Flow Protection. Must be one of 'enabled', 'disabled',
	// 'monitoring'.
	Mode param.Field[string] `json:"mode" api:"required"`
	// The name of the TCP Flow Protection rule. Value is relative to the 'scope'
	// setting. For 'global' scope, name should be 'global'. For either the 'region' or
	// 'datacenter' scope, name should be the actual name of the region or datacenter,
	// e.g., 'wnam' or 'lax'.
	Name param.Field[string] `json:"name" api:"required"`
	// The rate sensitivity. Must be one of 'low', 'medium', 'high'.
	RateSensitivity param.Field[string] `json:"rate_sensitivity" api:"required"`
	// The scope for the TCP Flow Protection rule.
	Scope param.Field[string] `json:"scope" api:"required"`
}

func (r AdvancedTCPProtectionTCPFlowProtectionRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionTCPFlowProtectionRuleNewResponse                `json:"result"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelope]
type advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrors struct {
	Code             int64                                                                     `json:"code" api:"required"`
	Message          string                                                                    `json:"message" api:"required"`
	DocumentationURL string                                                                    `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrors]
type advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                                        `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsSource]
type advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessages struct {
	Code             int64                                                                       `json:"code" api:"required"`
	Message          string                                                                      `json:"message" api:"required"`
	DocumentationURL string                                                                      `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessages]
type advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                                          `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesSource]
type advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeSuccessTrue AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionTCPFlowProtectionRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionTCPFlowProtectionRuleListParams struct {
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

// URLQuery serializes [AdvancedTCPProtectionTCPFlowProtectionRuleListParams]'s
// query parameters as `url.Values`.
func (r AdvancedTCPProtectionTCPFlowProtectionRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AdvancedTCPProtectionTCPFlowProtectionRuleBulkDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
