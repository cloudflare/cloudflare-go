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

// AccountRulesetVersionByTagService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountRulesetVersionByTagService] method instead.
type AccountRulesetVersionByTagService struct {
	Options []option.RequestOption
}

// NewAccountRulesetVersionByTagService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountRulesetVersionByTagService(opts ...option.RequestOption) (r *AccountRulesetVersionByTagService) {
	r = &AccountRulesetVersionByTagService{}
	r.Options = opts
	return
}

// Fetches the rules of a managed account ruleset version for a given tag.
func (r *AccountRulesetVersionByTagService) Get(ctx context.Context, accountID string, rulesetID string, rulesetVersion string, ruleTag string, opts ...option.RequestOption) (res *AccountRulesetVersionByTagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s/by_tag/%s", accountID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountRulesetVersionByTagGetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []AccountRulesetVersionByTagGetResponseMessage `json:"messages"`
	Result   AccountRulesetVersionByTagGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success AccountRulesetVersionByTagGetResponseSuccess `json:"success"`
	JSON    accountRulesetVersionByTagGetResponseJSON    `json:"-"`
}

// accountRulesetVersionByTagGetResponseJSON contains the JSON metadata for the
// struct [AccountRulesetVersionByTagGetResponse]
type accountRulesetVersionByTagGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionByTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type AccountRulesetVersionByTagGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetVersionByTagGetResponseMessagesSource `json:"source"`
	JSON   accountRulesetVersionByTagGetResponseMessageJSON    `json:"-"`
}

// accountRulesetVersionByTagGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountRulesetVersionByTagGetResponseMessage]
type accountRulesetVersionByTagGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionByTagGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type AccountRulesetVersionByTagGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                  `json:"pointer,required"`
	JSON    accountRulesetVersionByTagGetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetVersionByTagGetResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AccountRulesetVersionByTagGetResponseMessagesSource]
type accountRulesetVersionByTagGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetVersionByTagGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountRulesetVersionByTagGetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind AccountRulesetVersionByTagGetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase AccountRulesetVersionByTagGetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                          `json:"version"`
	JSON    accountRulesetVersionByTagGetResponseResultJSON `json:"-"`
}

// accountRulesetVersionByTagGetResponseResultJSON contains the JSON metadata for
// the struct [AccountRulesetVersionByTagGetResponseResult]
type accountRulesetVersionByTagGetResponseResultJSON struct {
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

func (r *AccountRulesetVersionByTagGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type AccountRulesetVersionByTagGetResponseResultKind string

const (
	AccountRulesetVersionByTagGetResponseResultKindManaged AccountRulesetVersionByTagGetResponseResultKind = "managed"
	AccountRulesetVersionByTagGetResponseResultKindCustom  AccountRulesetVersionByTagGetResponseResultKind = "custom"
	AccountRulesetVersionByTagGetResponseResultKindRoot    AccountRulesetVersionByTagGetResponseResultKind = "root"
	AccountRulesetVersionByTagGetResponseResultKindZone    AccountRulesetVersionByTagGetResponseResultKind = "zone"
)

// The phase of the ruleset.
type AccountRulesetVersionByTagGetResponseResultPhase string

const (
	AccountRulesetVersionByTagGetResponseResultPhaseDdosL4                         AccountRulesetVersionByTagGetResponseResultPhase = "ddos_l4"
	AccountRulesetVersionByTagGetResponseResultPhaseDdosL7                         AccountRulesetVersionByTagGetResponseResultPhase = "ddos_l7"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPConfigSettings             AccountRulesetVersionByTagGetResponseResultPhase = "http_config_settings"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPCustomErrors               AccountRulesetVersionByTagGetResponseResultPhase = "http_custom_errors"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPLogCustomFields            AccountRulesetVersionByTagGetResponseResultPhase = "http_log_custom_fields"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRatelimit                  AccountRulesetVersionByTagGetResponseResultPhase = "http_ratelimit"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestCacheSettings       AccountRulesetVersionByTagGetResponseResultPhase = "http_request_cache_settings"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestDynamicRedirect     AccountRulesetVersionByTagGetResponseResultPhase = "http_request_dynamic_redirect"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestFirewallCustom      AccountRulesetVersionByTagGetResponseResultPhase = "http_request_firewall_custom"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestFirewallManaged     AccountRulesetVersionByTagGetResponseResultPhase = "http_request_firewall_managed"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestLateTransform       AccountRulesetVersionByTagGetResponseResultPhase = "http_request_late_transform"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestOrigin              AccountRulesetVersionByTagGetResponseResultPhase = "http_request_origin"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestRedirect            AccountRulesetVersionByTagGetResponseResultPhase = "http_request_redirect"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestSanitize            AccountRulesetVersionByTagGetResponseResultPhase = "http_request_sanitize"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestSbfm                AccountRulesetVersionByTagGetResponseResultPhase = "http_request_sbfm"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestSelectConfiguration AccountRulesetVersionByTagGetResponseResultPhase = "http_request_select_configuration"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPRequestTransform           AccountRulesetVersionByTagGetResponseResultPhase = "http_request_transform"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPResponseCompression        AccountRulesetVersionByTagGetResponseResultPhase = "http_response_compression"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPResponseFirewallManaged    AccountRulesetVersionByTagGetResponseResultPhase = "http_response_firewall_managed"
	AccountRulesetVersionByTagGetResponseResultPhaseHTTPResponseHeadersTransform   AccountRulesetVersionByTagGetResponseResultPhase = "http_response_headers_transform"
	AccountRulesetVersionByTagGetResponseResultPhaseMagicTransit                   AccountRulesetVersionByTagGetResponseResultPhase = "magic_transit"
	AccountRulesetVersionByTagGetResponseResultPhaseMagicTransitIDsManaged         AccountRulesetVersionByTagGetResponseResultPhase = "magic_transit_ids_managed"
	AccountRulesetVersionByTagGetResponseResultPhaseMagicTransitManaged            AccountRulesetVersionByTagGetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type AccountRulesetVersionByTagGetResponseSuccess bool

const (
	AccountRulesetVersionByTagGetResponseSuccessTrue AccountRulesetVersionByTagGetResponseSuccess = true
)
