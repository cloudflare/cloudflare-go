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

// AccountRulesetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRulesetService] method
// instead.
type AccountRulesetService struct {
	Options  []option.RequestOption
	Phases   *AccountRulesetPhaseService
	Rules    *AccountRulesetRuleService
	Versions *AccountRulesetVersionService
}

// NewAccountRulesetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRulesetService(opts ...option.RequestOption) (r *AccountRulesetService) {
	r = &AccountRulesetService{}
	r.Options = opts
	r.Phases = NewAccountRulesetPhaseService(opts...)
	r.Rules = NewAccountRulesetRuleService(opts...)
	r.Versions = NewAccountRulesetVersionService(opts...)
	return
}

// Fetches the latest version of an account ruleset.
func (r *AccountRulesetService) Get(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (res *AccountRulesetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an account ruleset, creating a new version.
func (r *AccountRulesetService) Update(ctx context.Context, accountID string, rulesetID string, body AccountRulesetUpdateParams, opts ...option.RequestOption) (res *AccountRulesetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes all versions of an existing account ruleset.
func (r *AccountRulesetService) Delete(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Creates a ruleset at the account level.
func (r *AccountRulesetService) AccountRulesetsNewAnAccountRuleset(ctx context.Context, accountID string, body AccountRulesetAccountRulesetsNewAnAccountRulesetParams, opts ...option.RequestOption) (res *AccountRulesetAccountRulesetsNewAnAccountRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches all rulesets at the account level.
func (r *AccountRulesetService) AccountRulesetsListAccountRulesets(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountRulesetAccountRulesetsListAccountRulesetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountRulesetGetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetGetResponseMessage `json:"messages"`
	Result   AccountRulesetGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetGetResponseSuccess `json:"success"`
	JSON    accountRulesetGetResponseJSON    `json:"-"`
}

// accountRulesetGetResponseJSON contains the JSON metadata for the struct
// [AccountRulesetGetResponse]
type accountRulesetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetGetResponseMessagesSource `json:"source"`
	JSON   accountRulesetGetResponseMessageJSON    `json:"-"`
}

// accountRulesetGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountRulesetGetResponseMessage]
type accountRulesetGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                      `json:"pointer,required"`
	JSON    accountRulesetGetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetGetResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountRulesetGetResponseMessagesSource]
type accountRulesetGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetGetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetGetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetGetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                              `json:"version"`
	JSON    accountRulesetGetResponseResultJSON `json:"-"`
}

// accountRulesetGetResponseResultJSON contains the JSON metadata for the struct
// [AccountRulesetGetResponseResult]
type accountRulesetGetResponseResultJSON struct {
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

func (r *AccountRulesetGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetGetResponseResultKind string

const (
	AccountRulesetGetResponseResultKindManaged AccountRulesetGetResponseResultKind = "managed"
	AccountRulesetGetResponseResultKindCustom  AccountRulesetGetResponseResultKind = "custom"
	AccountRulesetGetResponseResultKindRoot    AccountRulesetGetResponseResultKind = "root"
	AccountRulesetGetResponseResultKindZone    AccountRulesetGetResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetGetResponseResultPhase string

const (
	AccountRulesetGetResponseResultPhaseDdosL4                         AccountRulesetGetResponseResultPhase = "ddos_l4"
	AccountRulesetGetResponseResultPhaseDdosL7                         AccountRulesetGetResponseResultPhase = "ddos_l7"
	AccountRulesetGetResponseResultPhaseHTTPConfigSettings             AccountRulesetGetResponseResultPhase = "http_config_settings"
	AccountRulesetGetResponseResultPhaseHTTPCustomErrors               AccountRulesetGetResponseResultPhase = "http_custom_errors"
	AccountRulesetGetResponseResultPhaseHTTPLogCustomFields            AccountRulesetGetResponseResultPhase = "http_log_custom_fields"
	AccountRulesetGetResponseResultPhaseHTTPRatelimit                  AccountRulesetGetResponseResultPhase = "http_ratelimit"
	AccountRulesetGetResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetGetResponseResultPhase = "http_request_cache_settings"
	AccountRulesetGetResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetGetResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetGetResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetGetResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetGetResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetGetResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetGetResponseResultPhaseHTTPRequestLateTransform       AccountRulesetGetResponseResultPhase = "http_request_late_transform"
	AccountRulesetGetResponseResultPhaseHTTPRequestOrigin              AccountRulesetGetResponseResultPhase = "http_request_origin"
	AccountRulesetGetResponseResultPhaseHTTPRequestRedirect            AccountRulesetGetResponseResultPhase = "http_request_redirect"
	AccountRulesetGetResponseResultPhaseHTTPRequestSanitize            AccountRulesetGetResponseResultPhase = "http_request_sanitize"
	AccountRulesetGetResponseResultPhaseHTTPRequestSbfm                AccountRulesetGetResponseResultPhase = "http_request_sbfm"
	AccountRulesetGetResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetGetResponseResultPhase = "http_request_select_configuration"
	AccountRulesetGetResponseResultPhaseHTTPRequestTransform           AccountRulesetGetResponseResultPhase = "http_request_transform"
	AccountRulesetGetResponseResultPhaseHTTPResponseCompression        AccountRulesetGetResponseResultPhase = "http_response_compression"
	AccountRulesetGetResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetGetResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetGetResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetGetResponseResultPhase = "http_response_headers_transform"
	AccountRulesetGetResponseResultPhaseMagicTransit                   AccountRulesetGetResponseResultPhase = "magic_transit"
	AccountRulesetGetResponseResultPhaseMagicTransitIDsManaged         AccountRulesetGetResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetGetResponseResultPhaseMagicTransitManaged            AccountRulesetGetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetGetResponseSuccess bool

const (
	AccountRulesetGetResponseSuccessTrue AccountRulesetGetResponseSuccess = true
)

type AccountRulesetUpdateResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetUpdateResponseMessage `json:"messages"`
	Result   AccountRulesetUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetUpdateResponseSuccess `json:"success"`
	JSON    accountRulesetUpdateResponseJSON    `json:"-"`
}

// accountRulesetUpdateResponseJSON contains the JSON metadata for the struct
// [AccountRulesetUpdateResponse]
type accountRulesetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetUpdateResponseMessagesSource `json:"source"`
	JSON   accountRulesetUpdateResponseMessageJSON    `json:"-"`
}

// accountRulesetUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountRulesetUpdateResponseMessage]
type accountRulesetUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                         `json:"pointer,required"`
	JSON    accountRulesetUpdateResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetUpdateResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountRulesetUpdateResponseMessagesSource]
type accountRulesetUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetUpdateResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetUpdateResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetUpdateResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                 `json:"version"`
	JSON    accountRulesetUpdateResponseResultJSON `json:"-"`
}

// accountRulesetUpdateResponseResultJSON contains the JSON metadata for the struct
// [AccountRulesetUpdateResponseResult]
type accountRulesetUpdateResponseResultJSON struct {
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

func (r *AccountRulesetUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetUpdateResponseResultKind string

const (
	AccountRulesetUpdateResponseResultKindManaged AccountRulesetUpdateResponseResultKind = "managed"
	AccountRulesetUpdateResponseResultKindCustom  AccountRulesetUpdateResponseResultKind = "custom"
	AccountRulesetUpdateResponseResultKindRoot    AccountRulesetUpdateResponseResultKind = "root"
	AccountRulesetUpdateResponseResultKindZone    AccountRulesetUpdateResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetUpdateResponseResultPhase string

const (
	AccountRulesetUpdateResponseResultPhaseDdosL4                         AccountRulesetUpdateResponseResultPhase = "ddos_l4"
	AccountRulesetUpdateResponseResultPhaseDdosL7                         AccountRulesetUpdateResponseResultPhase = "ddos_l7"
	AccountRulesetUpdateResponseResultPhaseHTTPConfigSettings             AccountRulesetUpdateResponseResultPhase = "http_config_settings"
	AccountRulesetUpdateResponseResultPhaseHTTPCustomErrors               AccountRulesetUpdateResponseResultPhase = "http_custom_errors"
	AccountRulesetUpdateResponseResultPhaseHTTPLogCustomFields            AccountRulesetUpdateResponseResultPhase = "http_log_custom_fields"
	AccountRulesetUpdateResponseResultPhaseHTTPRatelimit                  AccountRulesetUpdateResponseResultPhase = "http_ratelimit"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetUpdateResponseResultPhase = "http_request_cache_settings"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetUpdateResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetUpdateResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetUpdateResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestLateTransform       AccountRulesetUpdateResponseResultPhase = "http_request_late_transform"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestOrigin              AccountRulesetUpdateResponseResultPhase = "http_request_origin"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestRedirect            AccountRulesetUpdateResponseResultPhase = "http_request_redirect"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestSanitize            AccountRulesetUpdateResponseResultPhase = "http_request_sanitize"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestSbfm                AccountRulesetUpdateResponseResultPhase = "http_request_sbfm"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetUpdateResponseResultPhase = "http_request_select_configuration"
	AccountRulesetUpdateResponseResultPhaseHTTPRequestTransform           AccountRulesetUpdateResponseResultPhase = "http_request_transform"
	AccountRulesetUpdateResponseResultPhaseHTTPResponseCompression        AccountRulesetUpdateResponseResultPhase = "http_response_compression"
	AccountRulesetUpdateResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetUpdateResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetUpdateResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetUpdateResponseResultPhase = "http_response_headers_transform"
	AccountRulesetUpdateResponseResultPhaseMagicTransit                   AccountRulesetUpdateResponseResultPhase = "magic_transit"
	AccountRulesetUpdateResponseResultPhaseMagicTransitIDsManaged         AccountRulesetUpdateResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetUpdateResponseResultPhaseMagicTransitManaged            AccountRulesetUpdateResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetUpdateResponseSuccess bool

const (
	AccountRulesetUpdateResponseSuccessTrue AccountRulesetUpdateResponseSuccess = true
)

type AccountRulesetAccountRulesetsNewAnAccountRulesetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetAccountRulesetsNewAnAccountRulesetResponseMessage `json:"messages"`
	Result   AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetAccountRulesetsNewAnAccountRulesetResponseSuccess `json:"success"`
	JSON    accountRulesetAccountRulesetsNewAnAccountRulesetResponseJSON    `json:"-"`
}

// accountRulesetAccountRulesetsNewAnAccountRulesetResponseJSON contains the JSON
// metadata for the struct
// [AccountRulesetAccountRulesetsNewAnAccountRulesetResponse]
type accountRulesetAccountRulesetsNewAnAccountRulesetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetAccountRulesetsNewAnAccountRulesetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetAccountRulesetsNewAnAccountRulesetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetAccountRulesetsNewAnAccountRulesetResponseMessagesSource `json:"source"`
	JSON   accountRulesetAccountRulesetsNewAnAccountRulesetResponseMessageJSON    `json:"-"`
}

// accountRulesetAccountRulesetsNewAnAccountRulesetResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountRulesetAccountRulesetsNewAnAccountRulesetResponseMessage]
type accountRulesetAccountRulesetsNewAnAccountRulesetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetAccountRulesetsNewAnAccountRulesetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetAccountRulesetsNewAnAccountRulesetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                     `json:"pointer,required"`
	JSON    accountRulesetAccountRulesetsNewAnAccountRulesetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetAccountRulesetsNewAnAccountRulesetResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AccountRulesetAccountRulesetsNewAnAccountRulesetResponseMessagesSource]
type accountRulesetAccountRulesetsNewAnAccountRulesetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetAccountRulesetsNewAnAccountRulesetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                                             `json:"version"`
	JSON    accountRulesetAccountRulesetsNewAnAccountRulesetResponseResultJSON `json:"-"`
}

// accountRulesetAccountRulesetsNewAnAccountRulesetResponseResultJSON contains the
// JSON metadata for the struct
// [AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResult]
type accountRulesetAccountRulesetsNewAnAccountRulesetResponseResultJSON struct {
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

func (r *AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultKind string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultKindManaged AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultKind = "managed"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultKindCustom  AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultKind = "custom"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultKindRoot    AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultKind = "root"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultKindZone    AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseDdosL4                         AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "ddos_l4"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseDdosL7                         AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "ddos_l7"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPConfigSettings             AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_config_settings"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPCustomErrors               AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_custom_errors"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPLogCustomFields            AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_log_custom_fields"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRatelimit                  AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_ratelimit"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_cache_settings"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestLateTransform       AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_late_transform"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestOrigin              AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_origin"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestRedirect            AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_redirect"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestSanitize            AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_sanitize"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestSbfm                AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_sbfm"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_select_configuration"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPRequestTransform           AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_request_transform"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPResponseCompression        AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_response_compression"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "http_response_headers_transform"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseMagicTransit                   AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "magic_transit"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseMagicTransitIDsManaged         AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhaseMagicTransitManaged            AccountRulesetAccountRulesetsNewAnAccountRulesetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetAccountRulesetsNewAnAccountRulesetResponseSuccess bool

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetResponseSuccessTrue AccountRulesetAccountRulesetsNewAnAccountRulesetResponseSuccess = true
)

type AccountRulesetAccountRulesetsListAccountRulesetsResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetAccountRulesetsListAccountRulesetsResponseMessage `json:"messages"`
	// A list of rulesets. The returned information will not include the rules in each
	// ruleset.
	Result []AccountRulesetAccountRulesetsListAccountRulesetsResponseResult `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetAccountRulesetsListAccountRulesetsResponseSuccess `json:"success"`
	JSON    accountRulesetAccountRulesetsListAccountRulesetsResponseJSON    `json:"-"`
}

// accountRulesetAccountRulesetsListAccountRulesetsResponseJSON contains the JSON
// metadata for the struct
// [AccountRulesetAccountRulesetsListAccountRulesetsResponse]
type accountRulesetAccountRulesetsListAccountRulesetsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetAccountRulesetsListAccountRulesetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetAccountRulesetsListAccountRulesetsResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetAccountRulesetsListAccountRulesetsResponseMessagesSource `json:"source"`
	JSON   accountRulesetAccountRulesetsListAccountRulesetsResponseMessageJSON    `json:"-"`
}

// accountRulesetAccountRulesetsListAccountRulesetsResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountRulesetAccountRulesetsListAccountRulesetsResponseMessage]
type accountRulesetAccountRulesetsListAccountRulesetsResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetAccountRulesetsListAccountRulesetsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetAccountRulesetsListAccountRulesetsResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                     `json:"pointer,required"`
	JSON    accountRulesetAccountRulesetsListAccountRulesetsResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetAccountRulesetsListAccountRulesetsResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AccountRulesetAccountRulesetsListAccountRulesetsResponseMessagesSource]
type accountRulesetAccountRulesetsListAccountRulesetsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetAccountRulesetsListAccountRulesetsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetAccountRulesetsListAccountRulesetsResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetAccountRulesetsListAccountRulesetsResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase `json:"phase"`
	// The version of the ruleset.
	Version string                                                             `json:"version"`
	JSON    accountRulesetAccountRulesetsListAccountRulesetsResponseResultJSON `json:"-"`
}

// accountRulesetAccountRulesetsListAccountRulesetsResponseResultJSON contains the
// JSON metadata for the struct
// [AccountRulesetAccountRulesetsListAccountRulesetsResponseResult]
type accountRulesetAccountRulesetsListAccountRulesetsResponseResultJSON struct {
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

func (r *AccountRulesetAccountRulesetsListAccountRulesetsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetAccountRulesetsListAccountRulesetsResponseResultKind string

const (
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultKindManaged AccountRulesetAccountRulesetsListAccountRulesetsResponseResultKind = "managed"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultKindCustom  AccountRulesetAccountRulesetsListAccountRulesetsResponseResultKind = "custom"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultKindRoot    AccountRulesetAccountRulesetsListAccountRulesetsResponseResultKind = "root"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultKindZone    AccountRulesetAccountRulesetsListAccountRulesetsResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase string

const (
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseDdosL4                         AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "ddos_l4"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseDdosL7                         AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "ddos_l7"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPConfigSettings             AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_config_settings"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPCustomErrors               AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_custom_errors"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPLogCustomFields            AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_log_custom_fields"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRatelimit                  AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_ratelimit"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_cache_settings"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestLateTransform       AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_late_transform"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestOrigin              AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_origin"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestRedirect            AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_redirect"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestSanitize            AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_sanitize"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestSbfm                AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_sbfm"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_select_configuration"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPRequestTransform           AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_request_transform"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPResponseCompression        AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_response_compression"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "http_response_headers_transform"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseMagicTransit                   AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "magic_transit"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseMagicTransitIDsManaged         AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhaseMagicTransitManaged            AccountRulesetAccountRulesetsListAccountRulesetsResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetAccountRulesetsListAccountRulesetsResponseSuccess bool

const (
	AccountRulesetAccountRulesetsListAccountRulesetsResponseSuccessTrue AccountRulesetAccountRulesetsListAccountRulesetsResponseSuccess = true
)

type AccountRulesetUpdateParams struct {
	ID param.Field[interface{}] `json:"id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[AccountRulesetUpdateParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[AccountRulesetUpdateParamsPhase] `json:"phase"`
	// The list of rules in the ruleset.
	Rules param.Field[[]AccountRulesetUpdateParamsRule] `json:"rules"`
}

func (r AccountRulesetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type AccountRulesetUpdateParamsKind string

const (
	AccountRulesetUpdateParamsKindManaged AccountRulesetUpdateParamsKind = "managed"
	AccountRulesetUpdateParamsKindCustom  AccountRulesetUpdateParamsKind = "custom"
	AccountRulesetUpdateParamsKindRoot    AccountRulesetUpdateParamsKind = "root"
	AccountRulesetUpdateParamsKindZone    AccountRulesetUpdateParamsKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetUpdateParamsPhase string

const (
	AccountRulesetUpdateParamsPhaseDdosL4                         AccountRulesetUpdateParamsPhase = "ddos_l4"
	AccountRulesetUpdateParamsPhaseDdosL7                         AccountRulesetUpdateParamsPhase = "ddos_l7"
	AccountRulesetUpdateParamsPhaseHTTPConfigSettings             AccountRulesetUpdateParamsPhase = "http_config_settings"
	AccountRulesetUpdateParamsPhaseHTTPCustomErrors               AccountRulesetUpdateParamsPhase = "http_custom_errors"
	AccountRulesetUpdateParamsPhaseHTTPLogCustomFields            AccountRulesetUpdateParamsPhase = "http_log_custom_fields"
	AccountRulesetUpdateParamsPhaseHTTPRatelimit                  AccountRulesetUpdateParamsPhase = "http_ratelimit"
	AccountRulesetUpdateParamsPhaseHTTPRequestCacheSettings       AccountRulesetUpdateParamsPhase = "http_request_cache_settings"
	AccountRulesetUpdateParamsPhaseHTTPRequestDynamicRedirect     AccountRulesetUpdateParamsPhase = "http_request_dynamic_redirect"
	AccountRulesetUpdateParamsPhaseHTTPRequestFirewallCustom      AccountRulesetUpdateParamsPhase = "http_request_firewall_custom"
	AccountRulesetUpdateParamsPhaseHTTPRequestFirewallManaged     AccountRulesetUpdateParamsPhase = "http_request_firewall_managed"
	AccountRulesetUpdateParamsPhaseHTTPRequestLateTransform       AccountRulesetUpdateParamsPhase = "http_request_late_transform"
	AccountRulesetUpdateParamsPhaseHTTPRequestOrigin              AccountRulesetUpdateParamsPhase = "http_request_origin"
	AccountRulesetUpdateParamsPhaseHTTPRequestRedirect            AccountRulesetUpdateParamsPhase = "http_request_redirect"
	AccountRulesetUpdateParamsPhaseHTTPRequestSanitize            AccountRulesetUpdateParamsPhase = "http_request_sanitize"
	AccountRulesetUpdateParamsPhaseHTTPRequestSbfm                AccountRulesetUpdateParamsPhase = "http_request_sbfm"
	AccountRulesetUpdateParamsPhaseHTTPRequestSelectConfiguration AccountRulesetUpdateParamsPhase = "http_request_select_configuration"
	AccountRulesetUpdateParamsPhaseHTTPRequestTransform           AccountRulesetUpdateParamsPhase = "http_request_transform"
	AccountRulesetUpdateParamsPhaseHTTPResponseCompression        AccountRulesetUpdateParamsPhase = "http_response_compression"
	AccountRulesetUpdateParamsPhaseHTTPResponseFirewallManaged    AccountRulesetUpdateParamsPhase = "http_response_firewall_managed"
	AccountRulesetUpdateParamsPhaseHTTPResponseHeadersTransform   AccountRulesetUpdateParamsPhase = "http_response_headers_transform"
	AccountRulesetUpdateParamsPhaseMagicTransit                   AccountRulesetUpdateParamsPhase = "magic_transit"
	AccountRulesetUpdateParamsPhaseMagicTransitIDsManaged         AccountRulesetUpdateParamsPhase = "magic_transit_ids_managed"
	AccountRulesetUpdateParamsPhaseMagicTransitManaged            AccountRulesetUpdateParamsPhase = "magic_transit_managed"
)

// Satisfied by [AccountRulesetUpdateParamsRulesOexZd8xKBlockRule],
// [AccountRulesetUpdateParamsRulesOexZd8xKExecuteRule],
// [AccountRulesetUpdateParamsRulesOexZd8xKLogRule],
// [AccountRulesetUpdateParamsRulesOexZd8xKSkipRule].
type AccountRulesetUpdateParamsRule interface {
	implementsAccountRulesetUpdateParamsRule()
}

type AccountRulesetUpdateParamsRulesOexZd8xKBlockRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                           `json:"id"`
	Action           param.Field[AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                      `json:"description"`
	Enabled          param.Field[interface{}]                                                      `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKBlockRule) implementsAccountRulesetUpdateParamsRule() {
}

type AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleAction string

const (
	AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionBlock AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleAction = "block"
)

type AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParametersResponse] `json:"response"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetUpdateParamsRulesOexZd8xKExecuteRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                             `json:"id"`
	Action           param.Field[AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                        `json:"description"`
	Enabled          param.Field[interface{}]                                                        `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKExecuteRule) implementsAccountRulesetUpdateParamsRule() {
}

type AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleAction string

const (
	AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionExecute AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleAction = "execute"
)

type AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParameters struct {
	ID param.Field[interface{}] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverrides struct {
	Action param.Field[interface{}] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	Enabled    param.Field[interface{}]                                                                           `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules            param.Field[[]AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule] `json:"rules"`
	SensitivityLevel param.Field[interface{}]                                                                       `json:"sensitivity_level"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory struct {
	Category         param.Field[interface{}] `json:"category,required"`
	Action           param.Field[interface{}] `json:"action"`
	Enabled          param.Field[interface{}] `json:"enabled"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule-level override
type AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule struct {
	ID      param.Field[interface{}] `json:"id,required"`
	Action  param.Field[interface{}] `json:"action"`
	Enabled param.Field[interface{}] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold   param.Field[int64]       `json:"score_threshold"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetUpdateParamsRulesOexZd8xKLogRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                               `json:"id"`
	Action           param.Field[AccountRulesetUpdateParamsRulesOexZd8xKLogRuleAction] `json:"action"`
	ActionParameters param.Field[interface{}]                                          `json:"action_parameters"`
	Description      param.Field[interface{}]                                          `json:"description"`
	Enabled          param.Field[interface{}]                                          `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetUpdateParamsRulesOexZd8xKLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKLogRule) implementsAccountRulesetUpdateParamsRule() {}

type AccountRulesetUpdateParamsRulesOexZd8xKLogRuleAction string

const (
	AccountRulesetUpdateParamsRulesOexZd8xKLogRuleActionLog AccountRulesetUpdateParamsRulesOexZd8xKLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type AccountRulesetUpdateParamsRulesOexZd8xKLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetUpdateParamsRulesOexZd8xKSkipRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                          `json:"id"`
	Action           param.Field[AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                     `json:"description"`
	Enabled          param.Field[interface{}]                                                     `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKSkipRule) implementsAccountRulesetUpdateParamsRule() {}

type AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleAction string

const (
	AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionSkip AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleAction = "skip"
)

type AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]interface{}] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[interface{}] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]interface{}] `json:"rulesets"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name of a legacy security product to skip the execution of.
type AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct string

const (
	AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductBic           AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "bic"
	AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductHot           AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "hot"
	AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductRateLimit     AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "rateLimit"
	AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductSecurityLevel AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "securityLevel"
	AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductUaBlock       AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "uaBlock"
	AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductWaf           AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "waf"
	AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductZoneLockdown  AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersRuleset string

const (
	AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersRulesetCurrent AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetUpdateParamsRulesOexZd8xKSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetParams struct {
	ID param.Field[interface{}] `json:"id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase] `json:"phase"`
	// The list of rules in the ruleset.
	Rules param.Field[[]AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule] `json:"rules"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKindManaged AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind = "managed"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKindCustom  AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind = "custom"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKindRoot    AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind = "root"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKindZone    AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseDdosL4                         AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "ddos_l4"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseDdosL7                         AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "ddos_l7"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPConfigSettings             AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_config_settings"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPCustomErrors               AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_custom_errors"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPLogCustomFields            AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_log_custom_fields"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRatelimit                  AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_ratelimit"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestCacheSettings       AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_cache_settings"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestDynamicRedirect     AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_dynamic_redirect"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestFirewallCustom      AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_firewall_custom"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestFirewallManaged     AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_firewall_managed"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestLateTransform       AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_late_transform"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestOrigin              AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_origin"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestRedirect            AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_redirect"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestSanitize            AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_sanitize"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestSbfm                AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_sbfm"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestSelectConfiguration AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_select_configuration"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPRequestTransform           AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_request_transform"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPResponseCompression        AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_response_compression"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPResponseFirewallManaged    AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_response_firewall_managed"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseHTTPResponseHeadersTransform   AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "http_response_headers_transform"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseMagicTransit                   AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "magic_transit"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseMagicTransitIDsManaged         AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "magic_transit_ids_managed"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhaseMagicTransitManaged            AccountRulesetAccountRulesetsNewAnAccountRulesetParamsPhase = "magic_transit_managed"
)

// Satisfied by
// [AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRule],
// [AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRule],
// [AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRule],
// [AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRule].
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule interface {
	implementsAccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule()
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                       `json:"id"`
	Action           param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                  `json:"description"`
	Enabled          param.Field[interface{}]                                                                                  `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRule) implementsAccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule() {
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleAction string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionBlock AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleAction = "block"
)

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse] `json:"response"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                         `json:"id"`
	Action           param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                    `json:"description"`
	Enabled          param.Field[interface{}]                                                                                    `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRule) implementsAccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule() {
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleAction string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionExecute AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleAction = "execute"
)

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParameters struct {
	ID param.Field[interface{}] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverrides struct {
	Action param.Field[interface{}] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	Enabled    param.Field[interface{}]                                                                                                       `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules            param.Field[[]AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule] `json:"rules"`
	SensitivityLevel param.Field[interface{}]                                                                                                   `json:"sensitivity_level"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory struct {
	Category         param.Field[interface{}] `json:"category,required"`
	Action           param.Field[interface{}] `json:"action"`
	Enabled          param.Field[interface{}] `json:"enabled"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule-level override
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule struct {
	ID      param.Field[interface{}] `json:"id,required"`
	Action  param.Field[interface{}] `json:"action"`
	Enabled param.Field[interface{}] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold   param.Field[int64]       `json:"score_threshold"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                           `json:"id"`
	Action           param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRuleAction] `json:"action"`
	ActionParameters param.Field[interface{}]                                                                      `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                      `json:"description"`
	Enabled          param.Field[interface{}]                                                                      `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRule) implementsAccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule() {
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRuleAction string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRuleActionLog AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                      `json:"id"`
	Action           param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                 `json:"description"`
	Enabled          param.Field[interface{}]                                                                                 `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRule) implementsAccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule() {
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleAction string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionSkip AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleAction = "skip"
)

type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]interface{}] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[interface{}] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]interface{}] `json:"rulesets"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name of a legacy security product to skip the execution of.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductBic           AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "bic"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductHot           AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "hot"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductRateLimit     AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "rateLimit"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductSecurityLevel AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "securityLevel"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductUaBlock       AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "uaBlock"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductWaf           AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "waf"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductZoneLockdown  AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersRuleset string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersRulesetCurrent AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesOexZd8xKSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
