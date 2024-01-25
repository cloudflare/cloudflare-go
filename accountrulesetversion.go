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

// AccountRulesetVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRulesetVersionService]
// method instead.
type AccountRulesetVersionService struct {
	Options []option.RequestOption
	ByTags  *AccountRulesetVersionByTagService
}

// NewAccountRulesetVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRulesetVersionService(opts ...option.RequestOption) (r *AccountRulesetVersionService) {
	r = &AccountRulesetVersionService{}
	r.Options = opts
	r.ByTags = NewAccountRulesetVersionByTagService(opts...)
	return
}

// Fetches a specific version of an account ruleset.
func (r *AccountRulesetVersionService) Get(ctx context.Context, accountID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (res *AccountRulesetVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s", accountID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing version of an account ruleset.
func (r *AccountRulesetVersionService) Delete(ctx context.Context, accountID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s", accountID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the versions of an account ruleset.
func (r *AccountRulesetVersionService) AccountRulesetsListAnAccountRulesetSVersions(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (res *AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountRulesetVersionGetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetVersionGetResponseMessage `json:"messages"`
	Result   AccountRulesetVersionGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetVersionGetResponseSuccess `json:"success"`
	JSON    accountRulesetVersionGetResponseJSON    `json:"-"`
}

// accountRulesetVersionGetResponseJSON contains the JSON metadata for the struct
// [AccountRulesetVersionGetResponse]
type accountRulesetVersionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetVersionGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetVersionGetResponseMessagesSource `json:"source"`
	JSON   accountRulesetVersionGetResponseMessageJSON    `json:"-"`
}

// accountRulesetVersionGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountRulesetVersionGetResponseMessage]
type accountRulesetVersionGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetVersionGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                             `json:"pointer,required"`
	JSON    accountRulesetVersionGetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetVersionGetResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountRulesetVersionGetResponseMessagesSource]
type accountRulesetVersionGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetVersionGetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetVersionGetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetVersionGetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                     `json:"version"`
	JSON    accountRulesetVersionGetResponseResultJSON `json:"-"`
}

// accountRulesetVersionGetResponseResultJSON contains the JSON metadata for the
// struct [AccountRulesetVersionGetResponseResult]
type accountRulesetVersionGetResponseResultJSON struct {
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

func (r *AccountRulesetVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetVersionGetResponseResultKind string

const (
	AccountRulesetVersionGetResponseResultKindManaged AccountRulesetVersionGetResponseResultKind = "managed"
	AccountRulesetVersionGetResponseResultKindCustom  AccountRulesetVersionGetResponseResultKind = "custom"
	AccountRulesetVersionGetResponseResultKindRoot    AccountRulesetVersionGetResponseResultKind = "root"
	AccountRulesetVersionGetResponseResultKindZone    AccountRulesetVersionGetResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetVersionGetResponseResultPhase string

const (
	AccountRulesetVersionGetResponseResultPhaseDdosL4                         AccountRulesetVersionGetResponseResultPhase = "ddos_l4"
	AccountRulesetVersionGetResponseResultPhaseDdosL7                         AccountRulesetVersionGetResponseResultPhase = "ddos_l7"
	AccountRulesetVersionGetResponseResultPhaseHTTPConfigSettings             AccountRulesetVersionGetResponseResultPhase = "http_config_settings"
	AccountRulesetVersionGetResponseResultPhaseHTTPCustomErrors               AccountRulesetVersionGetResponseResultPhase = "http_custom_errors"
	AccountRulesetVersionGetResponseResultPhaseHTTPLogCustomFields            AccountRulesetVersionGetResponseResultPhase = "http_log_custom_fields"
	AccountRulesetVersionGetResponseResultPhaseHTTPRatelimit                  AccountRulesetVersionGetResponseResultPhase = "http_ratelimit"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetVersionGetResponseResultPhase = "http_request_cache_settings"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetVersionGetResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetVersionGetResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetVersionGetResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestLateTransform       AccountRulesetVersionGetResponseResultPhase = "http_request_late_transform"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestOrigin              AccountRulesetVersionGetResponseResultPhase = "http_request_origin"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestRedirect            AccountRulesetVersionGetResponseResultPhase = "http_request_redirect"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestSanitize            AccountRulesetVersionGetResponseResultPhase = "http_request_sanitize"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestSbfm                AccountRulesetVersionGetResponseResultPhase = "http_request_sbfm"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetVersionGetResponseResultPhase = "http_request_select_configuration"
	AccountRulesetVersionGetResponseResultPhaseHTTPRequestTransform           AccountRulesetVersionGetResponseResultPhase = "http_request_transform"
	AccountRulesetVersionGetResponseResultPhaseHTTPResponseCompression        AccountRulesetVersionGetResponseResultPhase = "http_response_compression"
	AccountRulesetVersionGetResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetVersionGetResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetVersionGetResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetVersionGetResponseResultPhase = "http_response_headers_transform"
	AccountRulesetVersionGetResponseResultPhaseMagicTransit                   AccountRulesetVersionGetResponseResultPhase = "magic_transit"
	AccountRulesetVersionGetResponseResultPhaseMagicTransitIDsManaged         AccountRulesetVersionGetResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetVersionGetResponseResultPhaseMagicTransitManaged            AccountRulesetVersionGetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetVersionGetResponseSuccess bool

const (
	AccountRulesetVersionGetResponseSuccessTrue AccountRulesetVersionGetResponseSuccess = true
)

type AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessage `json:"messages"`
	// A list of rulesets. The returned information will not include the rules in each
	// ruleset.
	Result []AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResult `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseSuccess `json:"success"`
	JSON    accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseJSON    `json:"-"`
}

// accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseJSON
// contains the JSON metadata for the struct
// [AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponse]
type accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessagesSource `json:"source"`
	JSON   accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessageJSON    `json:"-"`
}

// accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessage]
type accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                                      `json:"pointer,required"`
	JSON    accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessagesSource]
type accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase `json:"phase"`
	// The version of the ruleset.
	Version string                                                                              `json:"version"`
	JSON    accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultJSON `json:"-"`
}

// accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResult]
type accountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultJSON struct {
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

func (r *AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultKind string

const (
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultKindManaged AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultKind = "managed"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultKindCustom  AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultKind = "custom"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultKindRoot    AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultKind = "root"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultKindZone    AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase string

const (
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseDdosL4                         AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "ddos_l4"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseDdosL7                         AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "ddos_l7"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPConfigSettings             AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_config_settings"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPCustomErrors               AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_custom_errors"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPLogCustomFields            AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_log_custom_fields"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRatelimit                  AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_ratelimit"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_cache_settings"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestLateTransform       AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_late_transform"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestOrigin              AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_origin"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestRedirect            AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_redirect"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestSanitize            AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_sanitize"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestSbfm                AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_sbfm"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_select_configuration"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPRequestTransform           AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_request_transform"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPResponseCompression        AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_response_compression"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "http_response_headers_transform"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseMagicTransit                   AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "magic_transit"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseMagicTransitIDsManaged         AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhaseMagicTransitManaged            AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseSuccess bool

const (
	AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseSuccessTrue AccountRulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseSuccess = true
)
