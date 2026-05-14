// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ddos_protection

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// AdvancedTCPProtectionSynProtectionRuleItemService contains methods and other
// services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionSynProtectionRuleItemService] method instead.
type AdvancedTCPProtectionSynProtectionRuleItemService struct {
	Options []option.RequestOption
}

// NewAdvancedTCPProtectionSynProtectionRuleItemService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAdvancedTCPProtectionSynProtectionRuleItemService(opts ...option.RequestOption) (r *AdvancedTCPProtectionSynProtectionRuleItemService) {
	r = &AdvancedTCPProtectionSynProtectionRuleItemService{}
	r.Options = opts
	return
}

// Delete a SYN Protection rule specified by the given UUID.
func (r *AdvancedTCPProtectionSynProtectionRuleItemService) Delete(ctx context.Context, ruleID string, body AdvancedTCPProtectionSynProtectionRuleItemDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionSynProtectionRuleItemDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/rules/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Update a SYN Protection rule specified by the given UUID.
func (r *AdvancedTCPProtectionSynProtectionRuleItemService) Edit(ctx context.Context, ruleID string, params AdvancedTCPProtectionSynProtectionRuleItemEditParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionSynProtectionRuleItemEditResponse, err error) {
	var env AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get a SYN Protection rule specified by the given UUID.
func (r *AdvancedTCPProtectionSynProtectionRuleItemService) Get(ctx context.Context, ruleID string, query AdvancedTCPProtectionSynProtectionRuleItemGetParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionSynProtectionRuleItemGetResponse, err error) {
	var env AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/syn_protection/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AdvancedTCPProtectionSynProtectionRuleItemDeleteResponse struct {
	Errors   []AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionSynProtectionRuleItemDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemDeleteResponseJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemDeleteResponse]
type advancedTCPProtectionSynProtectionRuleItemDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseError struct {
	Code             int64                                                                `json:"code" api:"required"`
	Message          string                                                               `json:"message" api:"required"`
	DocumentationURL string                                                               `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseError]
type advancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorsSource struct {
	Pointer string                                                                   `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorsSource]
type advancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseMessage struct {
	Code             int64                                                                  `json:"code" api:"required"`
	Message          string                                                                 `json:"message" api:"required"`
	DocumentationURL string                                                                 `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionRuleItemDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemDeleteResponseMessageJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseMessage]
type advancedTCPProtectionSynProtectionRuleItemDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseMessagesSource struct {
	Pointer string                                                                     `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionRuleItemDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemDeleteResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseMessagesSource]
type advancedTCPProtectionSynProtectionRuleItemDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseSuccessTrue AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionSynProtectionRuleItemDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionSynProtectionRuleItemEditResponse struct {
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
	Scope string                                                     `json:"scope" api:"required"`
	JSON  advancedTCPProtectionSynProtectionRuleItemEditResponseJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemEditResponseJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionSynProtectionRuleItemEditResponse]
type advancedTCPProtectionSynProtectionRuleItemEditResponseJSON struct {
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

func (r *AdvancedTCPProtectionSynProtectionRuleItemEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemEditResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemGetResponse struct {
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
	Scope string                                                    `json:"scope" api:"required"`
	JSON  advancedTCPProtectionSynProtectionRuleItemGetResponseJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemGetResponseJSON contains the JSON
// metadata for the struct [AdvancedTCPProtectionSynProtectionRuleItemGetResponse]
type advancedTCPProtectionSynProtectionRuleItemGetResponseJSON struct {
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

func (r *AdvancedTCPProtectionSynProtectionRuleItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionSynProtectionRuleItemEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The new burst sensitivity. Optional. Must be one of 'low', 'medium', 'high'.
	BurstSensitivity param.Field[string] `json:"burst_sensitivity"`
	// The new mitigation type. Optional. Must be one of 'challenge' or 'retransmit'.
	MitigationType param.Field[string] `json:"mitigation_type"`
	// The new mode for SYN Protection. Optional. Must be one of 'enabled', 'disabled',
	// 'monitoring'.
	Mode param.Field[string] `json:"mode"`
	// The new rate sensitivity. Optional. Must be one of 'low', 'medium', 'high'.
	RateSensitivity param.Field[string] `json:"rate_sensitivity"`
}

func (r AdvancedTCPProtectionSynProtectionRuleItemEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionSynProtectionRuleItemEditResponse                `json:"result"`
	JSON    advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelope]
type advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrors struct {
	Code             int64                                                                      `json:"code" api:"required"`
	Message          string                                                                     `json:"message" api:"required"`
	DocumentationURL string                                                                     `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrors]
type advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                                         `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsSource]
type advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessages struct {
	Code             int64                                                                        `json:"code" api:"required"`
	Message          string                                                                       `json:"message" api:"required"`
	DocumentationURL string                                                                       `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessages]
type advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                                           `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesSource]
type advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeSuccessTrue AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionSynProtectionRuleItemEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionSynProtectionRuleItemGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionSynProtectionRuleItemGetResponse                `json:"result"`
	JSON    advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelope]
type advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrors struct {
	Code             int64                                                                     `json:"code" api:"required"`
	Message          string                                                                    `json:"message" api:"required"`
	DocumentationURL string                                                                    `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrors]
type advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                                        `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsSource]
type advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessages struct {
	Code             int64                                                                       `json:"code" api:"required"`
	Message          string                                                                      `json:"message" api:"required"`
	DocumentationURL string                                                                      `json:"documentation_url"`
	Source           AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessages]
type advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                                          `json:"pointer"`
	JSON    advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesSource]
type advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeSuccessTrue AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionSynProtectionRuleItemGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
