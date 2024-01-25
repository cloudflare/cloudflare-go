// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRulesetRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRulesetRuleService] method
// instead.
type AccountRulesetRuleService struct {
	Options []option.RequestOption
}

// NewAccountRulesetRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRulesetRuleService(opts ...option.RequestOption) (r *AccountRulesetRuleService) {
	r = &AccountRulesetRuleService{}
	r.Options = opts
	return
}

// Updates an existing rule in an account ruleset.
func (r *AccountRulesetRuleService) Update(ctx context.Context, accountID string, rulesetID string, ruleID string, body AccountRulesetRuleUpdateParams, opts ...option.RequestOption) (res *AccountRulesetRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/rules/%s", accountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an existing rule from an account ruleset.
func (r *AccountRulesetRuleService) Delete(ctx context.Context, accountID string, rulesetID string, ruleID string, opts ...option.RequestOption) (res *AccountRulesetRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/rules/%s", accountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new rule to an account ruleset. The rule will be added to the end of the
// existing list of rules in the ruleset by default.
func (r *AccountRulesetRuleService) AccountRulesetsNewAnAccountRulesetRule(ctx context.Context, accountID string, rulesetID string, body AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams, opts ...option.RequestOption) (res *AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/rules", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountRulesetRuleUpdateResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetRuleUpdateResponseMessage `json:"messages"`
	Result   AccountRulesetRuleUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetRuleUpdateResponseSuccess `json:"success"`
	JSON    accountRulesetRuleUpdateResponseJSON    `json:"-"`
}

// accountRulesetRuleUpdateResponseJSON contains the JSON metadata for the struct
// [AccountRulesetRuleUpdateResponse]
type accountRulesetRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetRuleUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetRuleUpdateResponseMessagesSource `json:"source"`
	JSON   accountRulesetRuleUpdateResponseMessageJSON    `json:"-"`
}

// accountRulesetRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountRulesetRuleUpdateResponseMessage]
type accountRulesetRuleUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetRuleUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                             `json:"pointer,required"`
	JSON    accountRulesetRuleUpdateResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetRuleUpdateResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountRulesetRuleUpdateResponseMessagesSource]
type accountRulesetRuleUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetRuleUpdateResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetRuleUpdateResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetRuleUpdateResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                     `json:"version"`
	JSON    accountRulesetRuleUpdateResponseResultJSON `json:"-"`
}

// accountRulesetRuleUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountRulesetRuleUpdateResponseResult]
type accountRulesetRuleUpdateResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetRuleUpdateResponseResultKind string

const (
	AccountRulesetRuleUpdateResponseResultKindManaged AccountRulesetRuleUpdateResponseResultKind = "managed"
	AccountRulesetRuleUpdateResponseResultKindCustom  AccountRulesetRuleUpdateResponseResultKind = "custom"
	AccountRulesetRuleUpdateResponseResultKindRoot    AccountRulesetRuleUpdateResponseResultKind = "root"
	AccountRulesetRuleUpdateResponseResultKindZone    AccountRulesetRuleUpdateResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetRuleUpdateResponseResultPhase string

const (
	AccountRulesetRuleUpdateResponseResultPhaseDdosL4                         AccountRulesetRuleUpdateResponseResultPhase = "ddos_l4"
	AccountRulesetRuleUpdateResponseResultPhaseDdosL7                         AccountRulesetRuleUpdateResponseResultPhase = "ddos_l7"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPConfigSettings             AccountRulesetRuleUpdateResponseResultPhase = "http_config_settings"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPCustomErrors               AccountRulesetRuleUpdateResponseResultPhase = "http_custom_errors"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPLogCustomFields            AccountRulesetRuleUpdateResponseResultPhase = "http_log_custom_fields"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRatelimit                  AccountRulesetRuleUpdateResponseResultPhase = "http_ratelimit"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetRuleUpdateResponseResultPhase = "http_request_cache_settings"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetRuleUpdateResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetRuleUpdateResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetRuleUpdateResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestLateTransform       AccountRulesetRuleUpdateResponseResultPhase = "http_request_late_transform"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestOrigin              AccountRulesetRuleUpdateResponseResultPhase = "http_request_origin"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestRedirect            AccountRulesetRuleUpdateResponseResultPhase = "http_request_redirect"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestSanitize            AccountRulesetRuleUpdateResponseResultPhase = "http_request_sanitize"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestSbfm                AccountRulesetRuleUpdateResponseResultPhase = "http_request_sbfm"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetRuleUpdateResponseResultPhase = "http_request_select_configuration"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPRequestTransform           AccountRulesetRuleUpdateResponseResultPhase = "http_request_transform"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPResponseCompression        AccountRulesetRuleUpdateResponseResultPhase = "http_response_compression"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetRuleUpdateResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetRuleUpdateResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetRuleUpdateResponseResultPhase = "http_response_headers_transform"
	AccountRulesetRuleUpdateResponseResultPhaseMagicTransit                   AccountRulesetRuleUpdateResponseResultPhase = "magic_transit"
	AccountRulesetRuleUpdateResponseResultPhaseMagicTransitIDsManaged         AccountRulesetRuleUpdateResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetRuleUpdateResponseResultPhaseMagicTransitManaged            AccountRulesetRuleUpdateResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetRuleUpdateResponseSuccess bool

const (
	AccountRulesetRuleUpdateResponseSuccessTrue AccountRulesetRuleUpdateResponseSuccess = true
)

type AccountRulesetRuleDeleteResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetRuleDeleteResponseMessage `json:"messages"`
	Result   AccountRulesetRuleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetRuleDeleteResponseSuccess `json:"success"`
	JSON    accountRulesetRuleDeleteResponseJSON    `json:"-"`
}

// accountRulesetRuleDeleteResponseJSON contains the JSON metadata for the struct
// [AccountRulesetRuleDeleteResponse]
type accountRulesetRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetRuleDeleteResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetRuleDeleteResponseMessagesSource `json:"source"`
	JSON   accountRulesetRuleDeleteResponseMessageJSON    `json:"-"`
}

// accountRulesetRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountRulesetRuleDeleteResponseMessage]
type accountRulesetRuleDeleteResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetRuleDeleteResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                             `json:"pointer,required"`
	JSON    accountRulesetRuleDeleteResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetRuleDeleteResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountRulesetRuleDeleteResponseMessagesSource]
type accountRulesetRuleDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetRuleDeleteResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetRuleDeleteResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetRuleDeleteResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                     `json:"version"`
	JSON    accountRulesetRuleDeleteResponseResultJSON `json:"-"`
}

// accountRulesetRuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountRulesetRuleDeleteResponseResult]
type accountRulesetRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetRuleDeleteResponseResultKind string

const (
	AccountRulesetRuleDeleteResponseResultKindManaged AccountRulesetRuleDeleteResponseResultKind = "managed"
	AccountRulesetRuleDeleteResponseResultKindCustom  AccountRulesetRuleDeleteResponseResultKind = "custom"
	AccountRulesetRuleDeleteResponseResultKindRoot    AccountRulesetRuleDeleteResponseResultKind = "root"
	AccountRulesetRuleDeleteResponseResultKindZone    AccountRulesetRuleDeleteResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetRuleDeleteResponseResultPhase string

const (
	AccountRulesetRuleDeleteResponseResultPhaseDdosL4                         AccountRulesetRuleDeleteResponseResultPhase = "ddos_l4"
	AccountRulesetRuleDeleteResponseResultPhaseDdosL7                         AccountRulesetRuleDeleteResponseResultPhase = "ddos_l7"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPConfigSettings             AccountRulesetRuleDeleteResponseResultPhase = "http_config_settings"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPCustomErrors               AccountRulesetRuleDeleteResponseResultPhase = "http_custom_errors"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPLogCustomFields            AccountRulesetRuleDeleteResponseResultPhase = "http_log_custom_fields"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRatelimit                  AccountRulesetRuleDeleteResponseResultPhase = "http_ratelimit"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetRuleDeleteResponseResultPhase = "http_request_cache_settings"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetRuleDeleteResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetRuleDeleteResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetRuleDeleteResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestLateTransform       AccountRulesetRuleDeleteResponseResultPhase = "http_request_late_transform"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestOrigin              AccountRulesetRuleDeleteResponseResultPhase = "http_request_origin"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestRedirect            AccountRulesetRuleDeleteResponseResultPhase = "http_request_redirect"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestSanitize            AccountRulesetRuleDeleteResponseResultPhase = "http_request_sanitize"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestSbfm                AccountRulesetRuleDeleteResponseResultPhase = "http_request_sbfm"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetRuleDeleteResponseResultPhase = "http_request_select_configuration"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPRequestTransform           AccountRulesetRuleDeleteResponseResultPhase = "http_request_transform"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPResponseCompression        AccountRulesetRuleDeleteResponseResultPhase = "http_response_compression"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetRuleDeleteResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetRuleDeleteResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetRuleDeleteResponseResultPhase = "http_response_headers_transform"
	AccountRulesetRuleDeleteResponseResultPhaseMagicTransit                   AccountRulesetRuleDeleteResponseResultPhase = "magic_transit"
	AccountRulesetRuleDeleteResponseResultPhaseMagicTransitIDsManaged         AccountRulesetRuleDeleteResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetRuleDeleteResponseResultPhaseMagicTransitManaged            AccountRulesetRuleDeleteResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetRuleDeleteResponseSuccess bool

const (
	AccountRulesetRuleDeleteResponseSuccessTrue AccountRulesetRuleDeleteResponseSuccess = true
)

type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessage `json:"messages"`
	Result   AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseSuccess `json:"success"`
	JSON    accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseJSON    `json:"-"`
}

// accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseJSON contains
// the JSON metadata for the struct
// [AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponse]
type accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessagesSource `json:"source"`
	JSON   accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessageJSON    `json:"-"`
}

// accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessage]
type accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                             `json:"pointer,required"`
	JSON    accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessagesSource]
type accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                                                     `json:"version"`
	JSON    accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultJSON `json:"-"`
}

// accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultJSON
// contains the JSON metadata for the struct
// [AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResult]
type accountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultKind string

const (
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultKindManaged AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultKind = "managed"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultKindCustom  AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultKind = "custom"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultKindRoot    AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultKind = "root"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultKindZone    AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase string

const (
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseDdosL4                         AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "ddos_l4"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseDdosL7                         AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "ddos_l7"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPConfigSettings             AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_config_settings"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPCustomErrors               AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_custom_errors"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPLogCustomFields            AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_log_custom_fields"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRatelimit                  AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_ratelimit"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_cache_settings"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestLateTransform       AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_late_transform"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestOrigin              AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_origin"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestRedirect            AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_redirect"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestSanitize            AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_sanitize"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestSbfm                AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_sbfm"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_select_configuration"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPRequestTransform           AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_request_transform"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPResponseCompression        AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_response_compression"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "http_response_headers_transform"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseMagicTransit                   AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "magic_transit"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseMagicTransitIDsManaged         AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhaseMagicTransitManaged            AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseSuccess bool

const (
	AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseSuccessTrue AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleResponseSuccess = true
)

type AccountRulesetRuleUpdateParams struct {
	Position param.Field[AccountRulesetRuleUpdateParamsPosition] `json:"position"`
}

func (r AccountRulesetRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [AccountRulesetRuleUpdateParamsPositionObject],
// [AccountRulesetRuleUpdateParamsPositionObject],
// [AccountRulesetRuleUpdateParamsPositionObject].
type AccountRulesetRuleUpdateParamsPosition interface {
	implementsAccountRulesetRuleUpdateParamsPosition()
}

type AccountRulesetRuleUpdateParamsPositionObject struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r AccountRulesetRuleUpdateParamsPositionObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleUpdateParamsPositionObject) implementsAccountRulesetRuleUpdateParamsPosition() {
}

type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams struct {
	Position param.Field[AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsPosition] `json:"position"`
}

func (r AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsPositionObject],
// [AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsPositionObject],
// [AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsPositionObject].
type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsPosition interface {
	implementsAccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsPosition()
}

type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsPositionObject struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsPositionObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsPositionObject) implementsAccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsPosition() {
}
