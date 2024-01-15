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

// ZoneRulesetRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRulesetRuleService] method
// instead.
type ZoneRulesetRuleService struct {
	Options []option.RequestOption
}

// NewZoneRulesetRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRulesetRuleService(opts ...option.RequestOption) (r *ZoneRulesetRuleService) {
	r = &ZoneRulesetRuleService{}
	r.Options = opts
	return
}

// Updates an existing rule in a zone ruleset.
func (r *ZoneRulesetRuleService) Update(ctx context.Context, zoneID string, rulesetID string, ruleID string, body ZoneRulesetRuleUpdateParams, opts ...option.RequestOption) (res *ZoneRulesetRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/rules/%s", zoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an existing rule from a zone ruleset.
func (r *ZoneRulesetRuleService) Delete(ctx context.Context, zoneID string, rulesetID string, ruleID string, opts ...option.RequestOption) (res *ZoneRulesetRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/rules/%s", zoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new rule to a zone ruleset. The rule will be added to the end of the
// existing list of rules in the ruleset by default.
func (r *ZoneRulesetRuleService) ZoneRulesetsNewAZoneRulesetRule(ctx context.Context, zoneID string, rulesetID string, body ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParams, opts ...option.RequestOption) (res *ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/rules", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneRulesetRuleUpdateResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetRuleUpdateResponseMessage `json:"messages"`
	Result   ZoneRulesetRuleUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetRuleUpdateResponseSuccess `json:"success"`
	JSON    zoneRulesetRuleUpdateResponseJSON    `json:"-"`
}

// zoneRulesetRuleUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetRuleUpdateResponse]
type zoneRulesetRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetRuleUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetRuleUpdateResponseMessagesSource `json:"source"`
	JSON   zoneRulesetRuleUpdateResponseMessageJSON    `json:"-"`
}

// zoneRulesetRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneRulesetRuleUpdateResponseMessage]
type zoneRulesetRuleUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetRuleUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    zoneRulesetRuleUpdateResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetRuleUpdateResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneRulesetRuleUpdateResponseMessagesSource]
type zoneRulesetRuleUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetRuleUpdateResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetRuleUpdateResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetRuleUpdateResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                  `json:"version"`
	JSON    zoneRulesetRuleUpdateResponseResultJSON `json:"-"`
}

// zoneRulesetRuleUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneRulesetRuleUpdateResponseResult]
type zoneRulesetRuleUpdateResponseResultJSON struct {
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

func (r *ZoneRulesetRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetRuleUpdateResponseResultKind string

const (
	ZoneRulesetRuleUpdateResponseResultKindManaged ZoneRulesetRuleUpdateResponseResultKind = "managed"
	ZoneRulesetRuleUpdateResponseResultKindCustom  ZoneRulesetRuleUpdateResponseResultKind = "custom"
	ZoneRulesetRuleUpdateResponseResultKindRoot    ZoneRulesetRuleUpdateResponseResultKind = "root"
	ZoneRulesetRuleUpdateResponseResultKindZone    ZoneRulesetRuleUpdateResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetRuleUpdateResponseResultPhase string

const (
	ZoneRulesetRuleUpdateResponseResultPhaseDdosL4                         ZoneRulesetRuleUpdateResponseResultPhase = "ddos_l4"
	ZoneRulesetRuleUpdateResponseResultPhaseDdosL7                         ZoneRulesetRuleUpdateResponseResultPhase = "ddos_l7"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPConfigSettings             ZoneRulesetRuleUpdateResponseResultPhase = "http_config_settings"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPCustomErrors               ZoneRulesetRuleUpdateResponseResultPhase = "http_custom_errors"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPLogCustomFields            ZoneRulesetRuleUpdateResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRatelimit                  ZoneRulesetRuleUpdateResponseResultPhase = "http_ratelimit"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetRuleUpdateResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetRuleUpdateResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetRuleUpdateResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetRuleUpdateResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetRuleUpdateResponseResultPhase = "http_request_late_transform"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestOrigin              ZoneRulesetRuleUpdateResponseResultPhase = "http_request_origin"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestRedirect            ZoneRulesetRuleUpdateResponseResultPhase = "http_request_redirect"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestSanitize            ZoneRulesetRuleUpdateResponseResultPhase = "http_request_sanitize"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestSbfm                ZoneRulesetRuleUpdateResponseResultPhase = "http_request_sbfm"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetRuleUpdateResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPRequestTransform           ZoneRulesetRuleUpdateResponseResultPhase = "http_request_transform"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPResponseCompression        ZoneRulesetRuleUpdateResponseResultPhase = "http_response_compression"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetRuleUpdateResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetRuleUpdateResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetRuleUpdateResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetRuleUpdateResponseResultPhaseMagicTransit                   ZoneRulesetRuleUpdateResponseResultPhase = "magic_transit"
	ZoneRulesetRuleUpdateResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetRuleUpdateResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetRuleUpdateResponseResultPhaseMagicTransitManaged            ZoneRulesetRuleUpdateResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetRuleUpdateResponseSuccess bool

const (
	ZoneRulesetRuleUpdateResponseSuccessTrue ZoneRulesetRuleUpdateResponseSuccess = true
)

type ZoneRulesetRuleDeleteResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetRuleDeleteResponseMessage `json:"messages"`
	Result   ZoneRulesetRuleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetRuleDeleteResponseSuccess `json:"success"`
	JSON    zoneRulesetRuleDeleteResponseJSON    `json:"-"`
}

// zoneRulesetRuleDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetRuleDeleteResponse]
type zoneRulesetRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetRuleDeleteResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetRuleDeleteResponseMessagesSource `json:"source"`
	JSON   zoneRulesetRuleDeleteResponseMessageJSON    `json:"-"`
}

// zoneRulesetRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneRulesetRuleDeleteResponseMessage]
type zoneRulesetRuleDeleteResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetRuleDeleteResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    zoneRulesetRuleDeleteResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetRuleDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneRulesetRuleDeleteResponseMessagesSource]
type zoneRulesetRuleDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetRuleDeleteResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetRuleDeleteResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetRuleDeleteResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                  `json:"version"`
	JSON    zoneRulesetRuleDeleteResponseResultJSON `json:"-"`
}

// zoneRulesetRuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneRulesetRuleDeleteResponseResult]
type zoneRulesetRuleDeleteResponseResultJSON struct {
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

func (r *ZoneRulesetRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetRuleDeleteResponseResultKind string

const (
	ZoneRulesetRuleDeleteResponseResultKindManaged ZoneRulesetRuleDeleteResponseResultKind = "managed"
	ZoneRulesetRuleDeleteResponseResultKindCustom  ZoneRulesetRuleDeleteResponseResultKind = "custom"
	ZoneRulesetRuleDeleteResponseResultKindRoot    ZoneRulesetRuleDeleteResponseResultKind = "root"
	ZoneRulesetRuleDeleteResponseResultKindZone    ZoneRulesetRuleDeleteResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetRuleDeleteResponseResultPhase string

const (
	ZoneRulesetRuleDeleteResponseResultPhaseDdosL4                         ZoneRulesetRuleDeleteResponseResultPhase = "ddos_l4"
	ZoneRulesetRuleDeleteResponseResultPhaseDdosL7                         ZoneRulesetRuleDeleteResponseResultPhase = "ddos_l7"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPConfigSettings             ZoneRulesetRuleDeleteResponseResultPhase = "http_config_settings"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPCustomErrors               ZoneRulesetRuleDeleteResponseResultPhase = "http_custom_errors"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPLogCustomFields            ZoneRulesetRuleDeleteResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRatelimit                  ZoneRulesetRuleDeleteResponseResultPhase = "http_ratelimit"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetRuleDeleteResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetRuleDeleteResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetRuleDeleteResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetRuleDeleteResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetRuleDeleteResponseResultPhase = "http_request_late_transform"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestOrigin              ZoneRulesetRuleDeleteResponseResultPhase = "http_request_origin"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestRedirect            ZoneRulesetRuleDeleteResponseResultPhase = "http_request_redirect"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestSanitize            ZoneRulesetRuleDeleteResponseResultPhase = "http_request_sanitize"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestSbfm                ZoneRulesetRuleDeleteResponseResultPhase = "http_request_sbfm"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetRuleDeleteResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPRequestTransform           ZoneRulesetRuleDeleteResponseResultPhase = "http_request_transform"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPResponseCompression        ZoneRulesetRuleDeleteResponseResultPhase = "http_response_compression"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetRuleDeleteResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetRuleDeleteResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetRuleDeleteResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetRuleDeleteResponseResultPhaseMagicTransit                   ZoneRulesetRuleDeleteResponseResultPhase = "magic_transit"
	ZoneRulesetRuleDeleteResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetRuleDeleteResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetRuleDeleteResponseResultPhaseMagicTransitManaged            ZoneRulesetRuleDeleteResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetRuleDeleteResponseSuccess bool

const (
	ZoneRulesetRuleDeleteResponseSuccessTrue ZoneRulesetRuleDeleteResponseSuccess = true
)

type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessage `json:"messages"`
	Result   ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseSuccess `json:"success"`
	JSON    zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseJSON    `json:"-"`
}

// zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseJSON contains the JSON
// metadata for the struct [ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponse]
type zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessagesSource `json:"source"`
	JSON   zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessageJSON    `json:"-"`
}

// zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessage]
type zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                   `json:"pointer,required"`
	JSON    zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessagesSource]
type zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                                           `json:"version"`
	JSON    zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultJSON `json:"-"`
}

// zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResult]
type zoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultJSON struct {
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

func (r *ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultKind string

const (
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultKindManaged ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultKind = "managed"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultKindCustom  ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultKind = "custom"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultKindRoot    ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultKind = "root"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultKindZone    ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase string

const (
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseDdosL4                         ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "ddos_l4"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseDdosL7                         ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "ddos_l7"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPConfigSettings             ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_config_settings"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPCustomErrors               ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_custom_errors"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPLogCustomFields            ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRatelimit                  ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_ratelimit"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_late_transform"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestOrigin              ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_origin"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestRedirect            ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_redirect"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestSanitize            ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_sanitize"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestSbfm                ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_sbfm"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPRequestTransform           ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_request_transform"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPResponseCompression        ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_response_compression"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseMagicTransit                   ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "magic_transit"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhaseMagicTransitManaged            ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseSuccess bool

const (
	ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseSuccessTrue ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleResponseSuccess = true
)

type ZoneRulesetRuleUpdateParams struct {
	Position param.Field[ZoneRulesetRuleUpdateParamsPosition] `json:"position"`
}

func (r ZoneRulesetRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [ZoneRulesetRuleUpdateParamsPositionObject],
// [ZoneRulesetRuleUpdateParamsPositionObject],
// [ZoneRulesetRuleUpdateParamsPositionObject].
type ZoneRulesetRuleUpdateParamsPosition interface {
	implementsZoneRulesetRuleUpdateParamsPosition()
}

type ZoneRulesetRuleUpdateParamsPositionObject struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r ZoneRulesetRuleUpdateParamsPositionObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleUpdateParamsPositionObject) implementsZoneRulesetRuleUpdateParamsPosition() {}

type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParams struct {
	Position param.Field[ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPosition] `json:"position"`
}

func (r ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPositionObject],
// [ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPositionObject],
// [ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPositionObject].
type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPosition interface {
	implementsZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPosition()
}

type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPositionObject struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPositionObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPositionObject) implementsZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsPosition() {
}
