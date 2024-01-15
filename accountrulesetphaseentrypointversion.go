// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountRulesetPhaseEntrypointVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountRulesetPhaseEntrypointVersionService] method instead.
type AccountRulesetPhaseEntrypointVersionService struct {
	Options []option.RequestOption
}

// NewAccountRulesetPhaseEntrypointVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountRulesetPhaseEntrypointVersionService(opts ...option.RequestOption) (r *AccountRulesetPhaseEntrypointVersionService) {
	r = &AccountRulesetPhaseEntrypointVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of an account entry point ruleset.
func (r *AccountRulesetPhaseEntrypointVersionService) Get(ctx context.Context, accountID string, rulesetPhase AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase, rulesetVersion string, opts ...option.RequestOption) (res *AccountRulesetPhaseEntrypointVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%v/entrypoint/versions/%s", accountID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the versions of an account entry point ruleset.
func (r *AccountRulesetPhaseEntrypointVersionService) AccountRulesetsListAnAccountEntryPointRulesetSVersions(ctx context.Context, accountID string, rulesetPhase AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase, opts ...option.RequestOption) (res *AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%v/entrypoint/versions", accountID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountRulesetPhaseEntrypointVersionGetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetPhaseEntrypointVersionGetResponseMessage `json:"messages"`
	Result   AccountRulesetPhaseEntrypointVersionGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetPhaseEntrypointVersionGetResponseSuccess `json:"success"`
	JSON    accountRulesetPhaseEntrypointVersionGetResponseJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointVersionGetResponseJSON contains the JSON metadata
// for the struct [AccountRulesetPhaseEntrypointVersionGetResponse]
type accountRulesetPhaseEntrypointVersionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetPhaseEntrypointVersionGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetPhaseEntrypointVersionGetResponseMessagesSource `json:"source"`
	JSON   accountRulesetPhaseEntrypointVersionGetResponseMessageJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointVersionGetResponseMessageJSON contains the JSON
// metadata for the struct [AccountRulesetPhaseEntrypointVersionGetResponseMessage]
type accountRulesetPhaseEntrypointVersionGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetPhaseEntrypointVersionGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                            `json:"pointer,required"`
	JSON    accountRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccountRulesetPhaseEntrypointVersionGetResponseMessagesSource]
type accountRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetPhaseEntrypointVersionGetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetPhaseEntrypointVersionGetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetPhaseEntrypointVersionGetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                                    `json:"version"`
	JSON    accountRulesetPhaseEntrypointVersionGetResponseResultJSON `json:"-"`
}

// accountRulesetPhaseEntrypointVersionGetResponseResultJSON contains the JSON
// metadata for the struct [AccountRulesetPhaseEntrypointVersionGetResponseResult]
type accountRulesetPhaseEntrypointVersionGetResponseResultJSON struct {
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

func (r *AccountRulesetPhaseEntrypointVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetPhaseEntrypointVersionGetResponseResultKind string

const (
	AccountRulesetPhaseEntrypointVersionGetResponseResultKindManaged AccountRulesetPhaseEntrypointVersionGetResponseResultKind = "managed"
	AccountRulesetPhaseEntrypointVersionGetResponseResultKindCustom  AccountRulesetPhaseEntrypointVersionGetResponseResultKind = "custom"
	AccountRulesetPhaseEntrypointVersionGetResponseResultKindRoot    AccountRulesetPhaseEntrypointVersionGetResponseResultKind = "root"
	AccountRulesetPhaseEntrypointVersionGetResponseResultKindZone    AccountRulesetPhaseEntrypointVersionGetResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetPhaseEntrypointVersionGetResponseResultPhase string

const (
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseDdosL4                         AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "ddos_l4"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseDdosL7                         AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "ddos_l7"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPConfigSettings             AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_config_settings"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPCustomErrors               AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_custom_errors"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPLogCustomFields            AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_log_custom_fields"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRatelimit                  AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_ratelimit"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_cache_settings"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestLateTransform       AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_late_transform"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestOrigin              AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_origin"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestRedirect            AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_redirect"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestSanitize            AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_sanitize"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestSbfm                AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_sbfm"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_select_configuration"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestTransform           AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_transform"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPResponseCompression        AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_response_compression"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_response_headers_transform"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseMagicTransit                   AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "magic_transit"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseMagicTransitIDsManaged         AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetPhaseEntrypointVersionGetResponseResultPhaseMagicTransitManaged            AccountRulesetPhaseEntrypointVersionGetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetPhaseEntrypointVersionGetResponseSuccess bool

const (
	AccountRulesetPhaseEntrypointVersionGetResponseSuccessTrue AccountRulesetPhaseEntrypointVersionGetResponseSuccess = true
)

type AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessage `json:"messages"`
	// A list of rulesets. The returned information will not include the rules in each
	// ruleset.
	Result []AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResult `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseSuccess `json:"success"`
	JSON    accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponse]
type accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessagesSource `json:"source"`
	JSON   accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessageJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessage]
type accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                                                               `json:"pointer,required"`
	JSON    accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessagesSource]
type accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase `json:"phase"`
	// The version of the ruleset.
	Version string                                                                                                       `json:"version"`
	JSON    accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultJSON `json:"-"`
}

// accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResult]
type accountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultKind string

const (
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultKindManaged AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultKind = "managed"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultKindCustom  AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultKind = "custom"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultKindRoot    AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultKind = "root"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultKindZone    AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase string

const (
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseDdosL4                         AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "ddos_l4"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseDdosL7                         AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "ddos_l7"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPConfigSettings             AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_config_settings"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPCustomErrors               AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_custom_errors"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPLogCustomFields            AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_log_custom_fields"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRatelimit                  AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_ratelimit"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_cache_settings"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestLateTransform       AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_late_transform"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestOrigin              AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_origin"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestRedirect            AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_redirect"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestSanitize            AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_sanitize"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestSbfm                AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_sbfm"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_select_configuration"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestTransform           AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_request_transform"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPResponseCompression        AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_response_compression"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "http_response_headers_transform"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseMagicTransit                   AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "magic_transit"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseMagicTransitIDsManaged         AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhaseMagicTransitManaged            AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseSuccess bool

const (
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseSuccessTrue AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsResponseSuccess = true
)

// The phase of the ruleset.
type AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase string

const (
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseDdosL4                         AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "ddos_l4"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseDdosL7                         AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "ddos_l7"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPConfigSettings             AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_config_settings"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPCustomErrors               AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_custom_errors"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPLogCustomFields            AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_log_custom_fields"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRatelimit                  AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_ratelimit"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestCacheSettings       AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_cache_settings"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestDynamicRedirect     AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_dynamic_redirect"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestFirewallCustom      AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_firewall_custom"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestFirewallManaged     AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_firewall_managed"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestLateTransform       AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_late_transform"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestOrigin              AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_origin"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestRedirect            AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_redirect"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestSanitize            AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_sanitize"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestSbfm                AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_sbfm"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestSelectConfiguration AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_select_configuration"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestTransform           AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_transform"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPResponseCompression        AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_response_compression"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPResponseFirewallManaged    AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_response_firewall_managed"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPResponseHeadersTransform   AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_response_headers_transform"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseMagicTransit                   AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "magic_transit"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseMagicTransitIDsManaged         AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "magic_transit_ids_managed"
	AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhaseMagicTransitManaged            AccountRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "magic_transit_managed"
)

// The phase of the ruleset.
type AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase string

const (
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseDdosL4                         AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "ddos_l4"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseDdosL7                         AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "ddos_l7"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPConfigSettings             AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_config_settings"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPCustomErrors               AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_custom_errors"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPLogCustomFields            AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_log_custom_fields"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRatelimit                  AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_ratelimit"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestCacheSettings       AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_cache_settings"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestDynamicRedirect     AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_dynamic_redirect"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestFirewallCustom      AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_firewall_custom"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestFirewallManaged     AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_firewall_managed"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestLateTransform       AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_late_transform"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestOrigin              AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_origin"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestRedirect            AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_redirect"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestSanitize            AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_sanitize"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestSbfm                AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_sbfm"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestSelectConfiguration AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_select_configuration"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestTransform           AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_transform"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPResponseCompression        AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_response_compression"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPResponseFirewallManaged    AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_response_firewall_managed"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseHTTPResponseHeadersTransform   AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "http_response_headers_transform"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseMagicTransit                   AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "magic_transit"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseMagicTransitIDsManaged         AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "magic_transit_ids_managed"
	AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhaseMagicTransitManaged            AccountRulesetPhaseEntrypointVersionAccountRulesetsListAnAccountEntryPointRulesetSVersionsParamsRulesetPhase = "magic_transit_managed"
)
