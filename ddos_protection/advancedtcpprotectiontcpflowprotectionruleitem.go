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

// AdvancedTCPProtectionTCPFlowProtectionRuleItemService contains methods and other
// services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdvancedTCPProtectionTCPFlowProtectionRuleItemService] method instead.
type AdvancedTCPProtectionTCPFlowProtectionRuleItemService struct {
	Options []option.RequestOption
}

// NewAdvancedTCPProtectionTCPFlowProtectionRuleItemService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAdvancedTCPProtectionTCPFlowProtectionRuleItemService(opts ...option.RequestOption) (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemService) {
	r = &AdvancedTCPProtectionTCPFlowProtectionRuleItemService{}
	r.Options = opts
	return
}

// Delete a TCP Flow Protection rule specified by the given UUID.
func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemService) Delete(ctx context.Context, ruleID string, body AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Update a TCP Flow Protection rule specified by the given UUID.
func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemService) Edit(ctx context.Context, ruleID string, params AdvancedTCPProtectionTCPFlowProtectionRuleItemEditParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponse, err error) {
	var env AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get a TCP Flow Protection rule specified by the given UUID.
func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemService) Get(ctx context.Context, ruleID string, query AdvancedTCPProtectionTCPFlowProtectionRuleItemGetParams, opts ...option.RequestOption) (res *AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponse, err error) {
	var env AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/advanced_tcp_protection/configs/tcp_flow_protection/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponse struct {
	Errors   []AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseError   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseSuccess `json:"success" api:"required"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseJSON contains the
// JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponse]
type advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseError struct {
	Code             int64                                                                    `json:"code" api:"required"`
	Message          string                                                                   `json:"message" api:"required"`
	DocumentationURL string                                                                   `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorsSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseError]
type advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorsSource struct {
	Pointer string                                                                       `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorsSource]
type advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessage struct {
	Code             int64                                                                      `json:"code" api:"required"`
	Message          string                                                                     `json:"message" api:"required"`
	DocumentationURL string                                                                     `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessagesSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessageJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessageJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessage]
type advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessagesSource struct {
	Pointer string                                                                         `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessagesSource]
type advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseSuccess bool

const (
	AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseSuccessTrue AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseSuccess = true
)

func (r AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponse struct {
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
	Scope string                                                         `json:"scope" api:"required"`
	JSON  advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponse]
type advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseJSON struct {
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

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponse struct {
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
	Scope string                                                        `json:"scope" api:"required"`
	JSON  advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseJSON contains the JSON
// metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponse]
type advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseJSON struct {
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

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The new burst sensitivity. Optional. Must be one of 'low', 'medium', 'high'.
	BurstSensitivity param.Field[string] `json:"burst_sensitivity"`
	// The new mode for TCP Flow Protection. Optional. Must be one of 'enabled',
	// 'disabled', 'monitoring'.
	Mode param.Field[string] `json:"mode"`
	// The new rate sensitivity. Optional. Must be one of 'low', 'medium', 'high'.
	RateSensitivity param.Field[string] `json:"rate_sensitivity"`
}

func (r AdvancedTCPProtectionTCPFlowProtectionRuleItemEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponse                `json:"result"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelope]
type advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrors struct {
	Code             int64                                                                          `json:"code" api:"required"`
	Message          string                                                                         `json:"message" api:"required"`
	DocumentationURL string                                                                         `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrors]
type advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                                             `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsSource]
type advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessages struct {
	Code             int64                                                                            `json:"code" api:"required"`
	Message          string                                                                           `json:"message" api:"required"`
	DocumentationURL string                                                                           `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessages]
type advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                                               `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesSource]
type advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeSuccessTrue AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionTCPFlowProtectionRuleItemEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelope struct {
	Errors   []AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponse                `json:"result"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeJSON    `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelope]
type advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrors struct {
	Code             int64                                                                         `json:"code" api:"required"`
	Message          string                                                                        `json:"message" api:"required"`
	DocumentationURL string                                                                        `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrors]
type advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                                            `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsSource]
type advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessages struct {
	Code             int64                                                                           `json:"code" api:"required"`
	Message          string                                                                          `json:"message" api:"required"`
	DocumentationURL string                                                                          `json:"documentation_url"`
	Source           AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessages]
type advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                                              `json:"pointer"`
	JSON    advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesSource]
type advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeSuccess bool

const (
	AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeSuccessTrue AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeSuccess = true
)

func (r AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AdvancedTCPProtectionTCPFlowProtectionRuleItemGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
