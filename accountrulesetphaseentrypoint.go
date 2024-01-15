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

// AccountRulesetPhaseEntrypointService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountRulesetPhaseEntrypointService] method instead.
type AccountRulesetPhaseEntrypointService struct {
	Options  []option.RequestOption
	Versions *AccountRulesetPhaseEntrypointVersionService
}

// NewAccountRulesetPhaseEntrypointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountRulesetPhaseEntrypointService(opts ...option.RequestOption) (r *AccountRulesetPhaseEntrypointService) {
	r = &AccountRulesetPhaseEntrypointService{}
	r.Options = opts
	r.Versions = NewAccountRulesetPhaseEntrypointVersionService(opts...)
	return
}

// Fetches the latest version of the account entry point ruleset for a given phase.
func (r *AccountRulesetPhaseEntrypointService) AccountRulesetsGetAnAccountEntryPointRuleset(ctx context.Context, accountID string, rulesetPhase AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase, opts ...option.RequestOption) (res *AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%v/entrypoint", accountID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an account entry point ruleset, creating a new version.
func (r *AccountRulesetPhaseEntrypointService) AccountRulesetsUpdateAnAccountEntryPointRuleset(ctx context.Context, accountID string, rulesetPhase AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase, body AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParams, opts ...option.RequestOption) (res *AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%v/entrypoint", accountID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessage `json:"messages"`
	Result   AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseSuccess `json:"success"`
	JSON    accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponse]
type accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessagesSource `json:"source"`
	JSON   accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessageJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessage]
type accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                                              `json:"pointer,required"`
	JSON    accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessagesSource]
type accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                                                                      `json:"version"`
	JSON    accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultJSON `json:"-"`
}

// accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResult]
type accountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultJSON struct {
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

func (r *AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultKind string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultKindManaged AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultKind = "managed"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultKindCustom  AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultKind = "custom"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultKindRoot    AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultKind = "root"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultKindZone    AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseDdosL4                         AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "ddos_l4"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseDdosL7                         AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "ddos_l7"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPConfigSettings             AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_config_settings"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPCustomErrors               AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_custom_errors"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPLogCustomFields            AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_log_custom_fields"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRatelimit                  AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_ratelimit"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_cache_settings"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestLateTransform       AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_late_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestOrigin              AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_origin"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestRedirect            AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_redirect"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestSanitize            AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_sanitize"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestSbfm                AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_sbfm"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_select_configuration"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestTransform           AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_request_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPResponseCompression        AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_response_compression"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "http_response_headers_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseMagicTransit                   AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "magic_transit"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseMagicTransitIDsManaged         AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhaseMagicTransitManaged            AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseSuccess bool

const (
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseSuccessTrue AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetResponseSuccess = true
)

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessage `json:"messages"`
	Result   AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseSuccess `json:"success"`
	JSON    accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponse]
type accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessagesSource `json:"source"`
	JSON   accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessageJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessage]
type accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                                                 `json:"pointer,required"`
	JSON    accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessagesSource]
type accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                                                                         `json:"version"`
	JSON    accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultJSON `json:"-"`
}

// accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResult]
type accountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultJSON struct {
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

func (r *AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultKind string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultKindManaged AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultKind = "managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultKindCustom  AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultKind = "custom"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultKindRoot    AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultKind = "root"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultKindZone    AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseDdosL4                         AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "ddos_l4"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseDdosL7                         AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "ddos_l7"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPConfigSettings             AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_config_settings"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPCustomErrors               AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_custom_errors"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPLogCustomFields            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_log_custom_fields"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRatelimit                  AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_ratelimit"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_cache_settings"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestLateTransform       AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_late_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestOrigin              AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_origin"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestRedirect            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_redirect"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestSanitize            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_sanitize"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestSbfm                AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_sbfm"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_select_configuration"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPRequestTransform           AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_request_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPResponseCompression        AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_response_compression"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "http_response_headers_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseMagicTransit                   AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "magic_transit"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseMagicTransitIDsManaged         AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhaseMagicTransitManaged            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseSuccess bool

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseSuccessTrue AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetResponseSuccess = true
)

// The phase of the ruleset.
type AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseDdosL4                         AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "ddos_l4"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseDdosL7                         AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "ddos_l7"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPConfigSettings             AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_config_settings"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPCustomErrors               AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_custom_errors"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPLogCustomFields            AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_log_custom_fields"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRatelimit                  AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_ratelimit"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestCacheSettings       AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_cache_settings"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestDynamicRedirect     AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_dynamic_redirect"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestFirewallCustom      AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_firewall_custom"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestFirewallManaged     AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_firewall_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestLateTransform       AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_late_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestOrigin              AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_origin"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestRedirect            AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_redirect"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestSanitize            AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_sanitize"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestSbfm                AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_sbfm"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestSelectConfiguration AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_select_configuration"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestTransform           AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPResponseCompression        AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_response_compression"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPResponseFirewallManaged    AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_response_firewall_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseHTTPResponseHeadersTransform   AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "http_response_headers_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseMagicTransit                   AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "magic_transit"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseMagicTransitIDsManaged         AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "magic_transit_ids_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhaseMagicTransitManaged            AccountRulesetPhaseEntrypointAccountRulesetsGetAnAccountEntryPointRulesetParamsRulesetPhase = "magic_transit_managed"
)

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParams struct {
	ID param.Field[interface{}] `json:"id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase] `json:"phase"`
	// The list of rules in the ruleset.
	Rules param.Field[[]AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule] `json:"rules"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The phase of the ruleset.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseDdosL4                         AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "ddos_l4"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseDdosL7                         AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "ddos_l7"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPConfigSettings             AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_config_settings"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPCustomErrors               AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_custom_errors"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPLogCustomFields            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_log_custom_fields"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRatelimit                  AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_ratelimit"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestCacheSettings       AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_cache_settings"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestDynamicRedirect     AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_dynamic_redirect"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestFirewallCustom      AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_firewall_custom"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestFirewallManaged     AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_firewall_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestLateTransform       AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_late_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestOrigin              AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_origin"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestRedirect            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_redirect"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestSanitize            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_sanitize"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestSbfm                AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_sbfm"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestSelectConfiguration AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_select_configuration"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPRequestTransform           AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_request_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPResponseCompression        AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_response_compression"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPResponseFirewallManaged    AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_response_firewall_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseHTTPResponseHeadersTransform   AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "http_response_headers_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseMagicTransit                   AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "magic_transit"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseMagicTransitIDsManaged         AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "magic_transit_ids_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhaseMagicTransitManaged            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesetPhase = "magic_transit_managed"
)

// The kind of the ruleset.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsKind string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsKindManaged AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsKind = "managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsKindCustom  AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsKind = "custom"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsKindRoot    AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsKind = "root"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsKindZone    AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseDdosL4                         AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "ddos_l4"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseDdosL7                         AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "ddos_l7"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPConfigSettings             AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_config_settings"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPCustomErrors               AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_custom_errors"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPLogCustomFields            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_log_custom_fields"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRatelimit                  AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_ratelimit"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestCacheSettings       AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_cache_settings"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestDynamicRedirect     AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_dynamic_redirect"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestFirewallCustom      AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_firewall_custom"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestFirewallManaged     AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_firewall_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestLateTransform       AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_late_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestOrigin              AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_origin"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestRedirect            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_redirect"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestSanitize            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_sanitize"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestSbfm                AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_sbfm"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestSelectConfiguration AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_select_configuration"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPRequestTransform           AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_request_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPResponseCompression        AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_response_compression"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPResponseFirewallManaged    AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_response_firewall_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseHTTPResponseHeadersTransform   AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "http_response_headers_transform"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseMagicTransit                   AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "magic_transit"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseMagicTransitIDsManaged         AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "magic_transit_ids_managed"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhaseMagicTransitManaged            AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsPhase = "magic_transit_managed"
)

// Satisfied by
// [AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRule],
// [AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRule],
// [AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRule],
// [AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRule].
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule interface {
	implementsAccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule()
}

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                                                   `json:"id"`
	Action           param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                                              `json:"description"`
	Enabled          param.Field[interface{}]                                                                                                              `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRule) implementsAccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule() {
}

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleAction string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleActionBlock AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleAction = "block"
)

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse] `json:"response"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                                                     `json:"id"`
	Action           param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                                                `json:"description"`
	Enabled          param.Field[interface{}]                                                                                                                `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRule) implementsAccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule() {
}

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleAction string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionExecute AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleAction = "execute"
)

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParameters struct {
	ID param.Field[interface{}] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverrides struct {
	Action param.Field[interface{}] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	Enabled    param.Field[interface{}]                                                                                                                                   `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules            param.Field[[]AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule] `json:"rules"`
	SensitivityLevel param.Field[interface{}]                                                                                                                               `json:"sensitivity_level"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory struct {
	Category         param.Field[interface{}] `json:"category,required"`
	Action           param.Field[interface{}] `json:"action"`
	Enabled          param.Field[interface{}] `json:"enabled"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule-level override
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule struct {
	ID      param.Field[interface{}] `json:"id,required"`
	Action  param.Field[interface{}] `json:"action"`
	Enabled param.Field[interface{}] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold   param.Field[int64]       `json:"score_threshold"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                                       `json:"id"`
	Action           param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRuleAction] `json:"action"`
	ActionParameters param.Field[interface{}]                                                                                                  `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                                  `json:"description"`
	Enabled          param.Field[interface{}]                                                                                                  `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRule) implementsAccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule() {
}

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRuleAction string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRuleActionLog AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                                                  `json:"id"`
	Action           param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                                             `json:"description"`
	Enabled          param.Field[interface{}]                                                                                                             `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRule) implementsAccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule() {
}

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleAction string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionSkip AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleAction = "skip"
)

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]interface{}] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[interface{}] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]interface{}] `json:"rulesets"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name of a legacy security product to skip the execution of.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductBic           AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "bic"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductHot           AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "hot"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductRateLimit     AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "rateLimit"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductSecurityLevel AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "securityLevel"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductUaBlock       AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "uaBlock"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductWaf           AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "waf"
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductZoneLockdown  AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersRuleset string

const (
	AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersRulesetCurrent AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesOexZd8xKSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
