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

// AdvancedTCPProtectionSynProtectionRuleService contains methods and other
// services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionSynProtectionRuleService] method instead.
type AdvancedTCPProtectionSynProtectionRuleService struct {
	Options []option.RequestOption
	Items   *AdvancedTCPProtectionSynProtectionRuleItemService
}

// NewAdvancedTCPProtectionSynProtectionRuleService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAdvancedTCPProtectionSynProtectionRuleService(opts ...option.RequestOption) (r *AdvancedTCPProtectionSynProtectionRuleService) {
	r = &AdvancedTCPProtectionSynProtectionRuleService{}
	r.Options = opts
	r.Items = NewAdvancedTCPProtectionSynProtectionRuleItemService(opts...)
	return
}

// Create a SYN Protection rule for an account.
func (r *AdvancedTCPProtectionSynProtectionRuleService) New(ctx context.Context, params AdvancedTCPProtectionSynProtectionRuleNewParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionSynProtectionRuleNewResponse, err error) {
	var env AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all SYN Protection rules for an account.
func (r *AdvancedTCPProtectionSynProtectionRuleService) List(ctx context.Context, params AdvancedTCPProtectionSynProtectionRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AdvancedTCPProtectionSynProtectionRuleListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/rules", params.AccountID)
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

// List all SYN Protection rules for an account.
func (r *AdvancedTCPProtectionSynProtectionRuleService) ListAutoPaging(ctx context.Context, params AdvancedTCPProtectionSynProtectionRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AdvancedTCPProtectionSynProtectionRuleListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete all SYN Protection rules for an account.
func (r *AdvancedTCPProtectionSynProtectionRuleService) BulkDelete(ctx context.Context, body AdvancedTCPProtectionSynProtectionRuleBulkDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/rules", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type AdvancedTCPProtectionSynProtectionRuleNewResponse struct {
	// The unique ID of the SYN Protection rule.
	ID string `json:"id" api:"required"`
	// The burst sensitivity. Must be one of 'low', 'medium', 'high'.
	BurstSensitivity string `json:"burst_sensitivity" api:"required"`
	// The creation timestamp of the SYN Protection rule.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The type of mitigation for SYN Protection. Must be one of 'challenge' or
	// 'retransmit'.
	MitigationType string `json:"mitigation_type" api:"required"`
	// The mode for SYN Protection. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the SYN Protection rule.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The name of the SYN Protection rule. Value is relative to the 'scope' setting.
	// For 'global' scope, name should be 'global'. For either the 'region' or
	// 'datacenter' scope, name should be the actual name of the region or datacenter,
	// e.g., 'wnam' or 'lax'.
	Name string `json:"name" api:"required"`
	// The rate sensitivity. Must be one of 'low', 'medium', 'high'.
	RateSensitivity string `json:"rate_sensitivity" api:"required"`
	// The scope for the SYN Protection rule. Must be one of 'global', 'region', or
	// 'datacenter'.
	Scope string                                                `json:"scope" api:"required"`
	JSON  advancedTCPProtectionSynProtectionRuleNewResponseJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleNewResponseJSON contains the JSON metadata
// for the struct [AdvancedTCPProtectionSynProtectionRuleNewResponse]
type advancedTCPProtectionSynProtectionRuleNewResponseJSON struct {
	ID               apijson.Field
	BurstSensitivity apijson.Field
	CreatedOn        apijson.Field
	MitigationType   apijson.Field
	Mode             apijson.Field
	ModifiedOn       apijson.Field
	Name             apijson.Field
	RateSensitivity  apijson.Field
	Scope            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleListResponse struct {
	// The unique ID of the SYN Protection rule.
	ID string `json:"id" api:"required"`
	// The burst sensitivity. Must be one of 'low', 'medium', 'high'.
	BurstSensitivity string `json:"burst_sensitivity" api:"required"`
	// The creation timestamp of the SYN Protection rule.
	CreatedOn time.Time `json:"created_on" api:"required" format:"date-time"`
	// The type of mitigation for SYN Protection. Must be one of 'challenge' or
	// 'retransmit'.
	MitigationType string `json:"mitigation_type" api:"required"`
	// The mode for SYN Protection. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode string `json:"mode" api:"required"`
	// The last modification timestamp of the SYN Protection rule.
	ModifiedOn time.Time `json:"modified_on" api:"required" format:"date-time"`
	// The name of the SYN Protection rule. Value is relative to the 'scope' setting.
	// For 'global' scope, name should be 'global'. For either the 'region' or
	// 'datacenter' scope, name should be the actual name of the region or datacenter,
	// e.g., 'wnam' or 'lax'.
	Name string `json:"name" api:"required"`
	// The rate sensitivity. Must be one of 'low', 'medium', 'high'.
	RateSensitivity string `json:"rate_sensitivity" api:"required"`
	// The scope for the SYN Protection rule. Must be one of 'global', 'region', or
	// 'datacenter'.
	Scope string                                                 `json:"scope" api:"required"`
	JSON  advancedTCPProtectionSynProtectionRuleListResponseJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleListResponseJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionSynProtectionRuleListResponse]
type advancedTCPProtectionSynProtectionRuleListResponseJSON struct {
	ID               apijson.Field
	BurstSensitivity apijson.Field
	CreatedOn        apijson.Field
	MitigationType   apijson.Field
	Mode             apijson.Field
	ModifiedOn       apijson.Field
	Name             apijson.Field
	RateSensitivity  apijson.Field
	Scope            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponse struct {
	Errors   []AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionSynProtectionRuleBulkDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleBulkDeleteResponseJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponse]
type advancedTCPProtectionSynProtectionRuleBulkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseError struct {
	Code             int64                                                                `json:"code" api:"required"`
	Message          string                                                               `json:"message" api:"required"`
	DocumentationURL string                                                               `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseError]
type advancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorsSource struct {
	Pointer string                                                                   `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorsSource]
type advancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleBulkDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessage struct {
	Code             int64                                                                  `json:"code" api:"required"`
	Message          string                                                                 `json:"message" api:"required"`
	DocumentationURL string                                                                 `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessageJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessage]
type advancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessagesSource struct {
	Pointer string                                                                     `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessagesSource]
type advancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleBulkDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseSuccessTrue AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionSynProtectionRuleBulkDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionSynProtectionRuleNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The burst sensitivity. Must be one of 'low', 'medium', 'high'.
	BurstSensitivity param.Field[string] `json:"burst_sensitivity" api:"required"`
	// The mode for SYN Protection. Must be one of 'enabled', 'disabled', 'monitoring'.
	Mode param.Field[string] `json:"mode" api:"required"`
	// The name of the SYN Protection rule. Value is relative to the 'scope' setting.
	// For 'global' scope, name should be 'global'. For either the 'region' or
	// 'datacenter' scope, name should be the actual name of the region or datacenter,
	// e.g., 'wnam' or 'lax'.
	Name param.Field[string] `json:"name" api:"required"`
	// The rate sensitivity. Must be one of 'low', 'medium', 'high'.
	RateSensitivity param.Field[string] `json:"rate_sensitivity" api:"required"`
	// The scope for the SYN Protection rule. Must be one of 'global', 'region', or
	// 'datacenter'.
	Scope param.Field[string] `json:"scope" api:"required"`
	// The type of mitigation. Must be one of 'challenge' or 'retransmit'. Optional.
	// Defaults to 'challenge'.
	MitigationType param.Field[string] `json:"mitigation_type"`
}

func (r AdvancedTCPProtectionSynProtectionRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionSynProtectionRuleNewResponse                `json:"result"`
	JSON    advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelope]
type advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrors struct {
	Code             int64                                                                 `json:"code" api:"required"`
	Message          string                                                                `json:"message" api:"required"`
	DocumentationURL string                                                                `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrors]
type advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                                    `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsSource]
type advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessages struct {
	Code             int64                                                                   `json:"code" api:"required"`
	Message          string                                                                  `json:"message" api:"required"`
	DocumentationURL string                                                                  `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessages]
type advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                                      `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesSource]
type advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeSuccessTrue AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionSynProtectionRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionSynProtectionRuleListParams struct {
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

// URLQuery serializes [AdvancedTCPProtectionSynProtectionRuleListParams]'s query
// parameters as `url.Values`.
func (r AdvancedTCPProtectionSynProtectionRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AdvancedTCPProtectionSynProtectionRuleBulkDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
