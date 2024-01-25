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

// ZoneRulesetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRulesetService] method
// instead.
type ZoneRulesetService struct {
	Options  []option.RequestOption
	Phases   *ZoneRulesetPhaseService
	Rules    *ZoneRulesetRuleService
	Versions *ZoneRulesetVersionService
}

// NewZoneRulesetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRulesetService(opts ...option.RequestOption) (r *ZoneRulesetService) {
	r = &ZoneRulesetService{}
	r.Options = opts
	r.Phases = NewZoneRulesetPhaseService(opts...)
	r.Rules = NewZoneRulesetRuleService(opts...)
	r.Versions = NewZoneRulesetVersionService(opts...)
	return
}

// Fetches the latest version of a zone ruleset.
func (r *ZoneRulesetService) Get(ctx context.Context, zoneID string, rulesetID string, opts ...option.RequestOption) (res *ZoneRulesetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a zone ruleset, creating a new version.
func (r *ZoneRulesetService) Update(ctx context.Context, zoneID string, rulesetID string, body ZoneRulesetUpdateParams, opts ...option.RequestOption) (res *ZoneRulesetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes all versions of an existing zone ruleset.
func (r *ZoneRulesetService) Delete(ctx context.Context, zoneID string, rulesetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Creates a ruleset at the zone level.
func (r *ZoneRulesetService) ZoneRulesetsNewAZoneRuleset(ctx context.Context, zoneID string, body ZoneRulesetZoneRulesetsNewAZoneRulesetParams, opts ...option.RequestOption) (res *ZoneRulesetZoneRulesetsNewAZoneRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches all rulesets at the zone level.
func (r *ZoneRulesetService) ZoneRulesetsListZoneRulesets(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneRulesetZoneRulesetsListZoneRulesetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneRulesetGetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetGetResponseMessage `json:"messages"`
	Result   ZoneRulesetGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetGetResponseSuccess `json:"success"`
	JSON    zoneRulesetGetResponseJSON    `json:"-"`
}

// zoneRulesetGetResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetGetResponse]
type zoneRulesetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetGetResponseMessagesSource `json:"source"`
	JSON   zoneRulesetGetResponseMessageJSON    `json:"-"`
}

// zoneRulesetGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRulesetGetResponseMessage]
type zoneRulesetGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    zoneRulesetGetResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetGetResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ZoneRulesetGetResponseMessagesSource]
type zoneRulesetGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetGetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetGetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetGetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                           `json:"version"`
	JSON    zoneRulesetGetResponseResultJSON `json:"-"`
}

// zoneRulesetGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneRulesetGetResponseResult]
type zoneRulesetGetResponseResultJSON struct {
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

func (r *ZoneRulesetGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetGetResponseResultKind string

const (
	ZoneRulesetGetResponseResultKindManaged ZoneRulesetGetResponseResultKind = "managed"
	ZoneRulesetGetResponseResultKindCustom  ZoneRulesetGetResponseResultKind = "custom"
	ZoneRulesetGetResponseResultKindRoot    ZoneRulesetGetResponseResultKind = "root"
	ZoneRulesetGetResponseResultKindZone    ZoneRulesetGetResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetGetResponseResultPhase string

const (
	ZoneRulesetGetResponseResultPhaseDdosL4                         ZoneRulesetGetResponseResultPhase = "ddos_l4"
	ZoneRulesetGetResponseResultPhaseDdosL7                         ZoneRulesetGetResponseResultPhase = "ddos_l7"
	ZoneRulesetGetResponseResultPhaseHTTPConfigSettings             ZoneRulesetGetResponseResultPhase = "http_config_settings"
	ZoneRulesetGetResponseResultPhaseHTTPCustomErrors               ZoneRulesetGetResponseResultPhase = "http_custom_errors"
	ZoneRulesetGetResponseResultPhaseHTTPLogCustomFields            ZoneRulesetGetResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetGetResponseResultPhaseHTTPRatelimit                  ZoneRulesetGetResponseResultPhase = "http_ratelimit"
	ZoneRulesetGetResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetGetResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetGetResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetGetResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetGetResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetGetResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetGetResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetGetResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetGetResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetGetResponseResultPhase = "http_request_late_transform"
	ZoneRulesetGetResponseResultPhaseHTTPRequestOrigin              ZoneRulesetGetResponseResultPhase = "http_request_origin"
	ZoneRulesetGetResponseResultPhaseHTTPRequestRedirect            ZoneRulesetGetResponseResultPhase = "http_request_redirect"
	ZoneRulesetGetResponseResultPhaseHTTPRequestSanitize            ZoneRulesetGetResponseResultPhase = "http_request_sanitize"
	ZoneRulesetGetResponseResultPhaseHTTPRequestSbfm                ZoneRulesetGetResponseResultPhase = "http_request_sbfm"
	ZoneRulesetGetResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetGetResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetGetResponseResultPhaseHTTPRequestTransform           ZoneRulesetGetResponseResultPhase = "http_request_transform"
	ZoneRulesetGetResponseResultPhaseHTTPResponseCompression        ZoneRulesetGetResponseResultPhase = "http_response_compression"
	ZoneRulesetGetResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetGetResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetGetResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetGetResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetGetResponseResultPhaseMagicTransit                   ZoneRulesetGetResponseResultPhase = "magic_transit"
	ZoneRulesetGetResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetGetResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetGetResponseResultPhaseMagicTransitManaged            ZoneRulesetGetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetGetResponseSuccess bool

const (
	ZoneRulesetGetResponseSuccessTrue ZoneRulesetGetResponseSuccess = true
)

type ZoneRulesetUpdateResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetUpdateResponseMessage `json:"messages"`
	Result   ZoneRulesetUpdateResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetUpdateResponseSuccess `json:"success"`
	JSON    zoneRulesetUpdateResponseJSON    `json:"-"`
}

// zoneRulesetUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetUpdateResponse]
type zoneRulesetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetUpdateResponseMessagesSource `json:"source"`
	JSON   zoneRulesetUpdateResponseMessageJSON    `json:"-"`
}

// zoneRulesetUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRulesetUpdateResponseMessage]
type zoneRulesetUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                      `json:"pointer,required"`
	JSON    zoneRulesetUpdateResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetUpdateResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ZoneRulesetUpdateResponseMessagesSource]
type zoneRulesetUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetUpdateResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetUpdateResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetUpdateResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                              `json:"version"`
	JSON    zoneRulesetUpdateResponseResultJSON `json:"-"`
}

// zoneRulesetUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneRulesetUpdateResponseResult]
type zoneRulesetUpdateResponseResultJSON struct {
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

func (r *ZoneRulesetUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetUpdateResponseResultKind string

const (
	ZoneRulesetUpdateResponseResultKindManaged ZoneRulesetUpdateResponseResultKind = "managed"
	ZoneRulesetUpdateResponseResultKindCustom  ZoneRulesetUpdateResponseResultKind = "custom"
	ZoneRulesetUpdateResponseResultKindRoot    ZoneRulesetUpdateResponseResultKind = "root"
	ZoneRulesetUpdateResponseResultKindZone    ZoneRulesetUpdateResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetUpdateResponseResultPhase string

const (
	ZoneRulesetUpdateResponseResultPhaseDdosL4                         ZoneRulesetUpdateResponseResultPhase = "ddos_l4"
	ZoneRulesetUpdateResponseResultPhaseDdosL7                         ZoneRulesetUpdateResponseResultPhase = "ddos_l7"
	ZoneRulesetUpdateResponseResultPhaseHTTPConfigSettings             ZoneRulesetUpdateResponseResultPhase = "http_config_settings"
	ZoneRulesetUpdateResponseResultPhaseHTTPCustomErrors               ZoneRulesetUpdateResponseResultPhase = "http_custom_errors"
	ZoneRulesetUpdateResponseResultPhaseHTTPLogCustomFields            ZoneRulesetUpdateResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetUpdateResponseResultPhaseHTTPRatelimit                  ZoneRulesetUpdateResponseResultPhase = "http_ratelimit"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetUpdateResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetUpdateResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetUpdateResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetUpdateResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetUpdateResponseResultPhase = "http_request_late_transform"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestOrigin              ZoneRulesetUpdateResponseResultPhase = "http_request_origin"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestRedirect            ZoneRulesetUpdateResponseResultPhase = "http_request_redirect"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestSanitize            ZoneRulesetUpdateResponseResultPhase = "http_request_sanitize"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestSbfm                ZoneRulesetUpdateResponseResultPhase = "http_request_sbfm"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetUpdateResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetUpdateResponseResultPhaseHTTPRequestTransform           ZoneRulesetUpdateResponseResultPhase = "http_request_transform"
	ZoneRulesetUpdateResponseResultPhaseHTTPResponseCompression        ZoneRulesetUpdateResponseResultPhase = "http_response_compression"
	ZoneRulesetUpdateResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetUpdateResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetUpdateResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetUpdateResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetUpdateResponseResultPhaseMagicTransit                   ZoneRulesetUpdateResponseResultPhase = "magic_transit"
	ZoneRulesetUpdateResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetUpdateResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetUpdateResponseResultPhaseMagicTransitManaged            ZoneRulesetUpdateResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetUpdateResponseSuccess bool

const (
	ZoneRulesetUpdateResponseSuccessTrue ZoneRulesetUpdateResponseSuccess = true
)

type ZoneRulesetZoneRulesetsNewAZoneRulesetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetZoneRulesetsNewAZoneRulesetResponseMessage `json:"messages"`
	Result   ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetZoneRulesetsNewAZoneRulesetResponseSuccess `json:"success"`
	JSON    zoneRulesetZoneRulesetsNewAZoneRulesetResponseJSON    `json:"-"`
}

// zoneRulesetZoneRulesetsNewAZoneRulesetResponseJSON contains the JSON metadata
// for the struct [ZoneRulesetZoneRulesetsNewAZoneRulesetResponse]
type zoneRulesetZoneRulesetsNewAZoneRulesetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetZoneRulesetsNewAZoneRulesetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetZoneRulesetsNewAZoneRulesetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetZoneRulesetsNewAZoneRulesetResponseMessagesSource `json:"source"`
	JSON   zoneRulesetZoneRulesetsNewAZoneRulesetResponseMessageJSON    `json:"-"`
}

// zoneRulesetZoneRulesetsNewAZoneRulesetResponseMessageJSON contains the JSON
// metadata for the struct [ZoneRulesetZoneRulesetsNewAZoneRulesetResponseMessage]
type zoneRulesetZoneRulesetsNewAZoneRulesetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetZoneRulesetsNewAZoneRulesetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetZoneRulesetsNewAZoneRulesetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                           `json:"pointer,required"`
	JSON    zoneRulesetZoneRulesetsNewAZoneRulesetResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetZoneRulesetsNewAZoneRulesetResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [ZoneRulesetZoneRulesetsNewAZoneRulesetResponseMessagesSource]
type zoneRulesetZoneRulesetsNewAZoneRulesetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetZoneRulesetsNewAZoneRulesetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                                   `json:"version"`
	JSON    zoneRulesetZoneRulesetsNewAZoneRulesetResponseResultJSON `json:"-"`
}

// zoneRulesetZoneRulesetsNewAZoneRulesetResponseResultJSON contains the JSON
// metadata for the struct [ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResult]
type zoneRulesetZoneRulesetsNewAZoneRulesetResponseResultJSON struct {
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

func (r *ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultKind string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultKindManaged ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultKind = "managed"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultKindCustom  ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultKind = "custom"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultKindRoot    ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultKind = "root"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultKindZone    ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseDdosL4                         ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "ddos_l4"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseDdosL7                         ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "ddos_l7"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPConfigSettings             ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_config_settings"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPCustomErrors               ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_custom_errors"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPLogCustomFields            ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRatelimit                  ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_ratelimit"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_late_transform"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestOrigin              ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_origin"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestRedirect            ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_redirect"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestSanitize            ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_sanitize"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestSbfm                ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_sbfm"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPRequestTransform           ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_request_transform"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPResponseCompression        ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_response_compression"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseMagicTransit                   ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "magic_transit"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhaseMagicTransitManaged            ZoneRulesetZoneRulesetsNewAZoneRulesetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetZoneRulesetsNewAZoneRulesetResponseSuccess bool

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetResponseSuccessTrue ZoneRulesetZoneRulesetsNewAZoneRulesetResponseSuccess = true
)

type ZoneRulesetZoneRulesetsListZoneRulesetsResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetZoneRulesetsListZoneRulesetsResponseMessage `json:"messages"`
	// A list of rulesets. The returned information will not include the rules in each
	// ruleset.
	Result []ZoneRulesetZoneRulesetsListZoneRulesetsResponseResult `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetZoneRulesetsListZoneRulesetsResponseSuccess `json:"success"`
	JSON    zoneRulesetZoneRulesetsListZoneRulesetsResponseJSON    `json:"-"`
}

// zoneRulesetZoneRulesetsListZoneRulesetsResponseJSON contains the JSON metadata
// for the struct [ZoneRulesetZoneRulesetsListZoneRulesetsResponse]
type zoneRulesetZoneRulesetsListZoneRulesetsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetZoneRulesetsListZoneRulesetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetZoneRulesetsListZoneRulesetsResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetZoneRulesetsListZoneRulesetsResponseMessagesSource `json:"source"`
	JSON   zoneRulesetZoneRulesetsListZoneRulesetsResponseMessageJSON    `json:"-"`
}

// zoneRulesetZoneRulesetsListZoneRulesetsResponseMessageJSON contains the JSON
// metadata for the struct [ZoneRulesetZoneRulesetsListZoneRulesetsResponseMessage]
type zoneRulesetZoneRulesetsListZoneRulesetsResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetZoneRulesetsListZoneRulesetsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetZoneRulesetsListZoneRulesetsResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                            `json:"pointer,required"`
	JSON    zoneRulesetZoneRulesetsListZoneRulesetsResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetZoneRulesetsListZoneRulesetsResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [ZoneRulesetZoneRulesetsListZoneRulesetsResponseMessagesSource]
type zoneRulesetZoneRulesetsListZoneRulesetsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetZoneRulesetsListZoneRulesetsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetZoneRulesetsListZoneRulesetsResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase `json:"phase"`
	// The version of the ruleset.
	Version string                                                    `json:"version"`
	JSON    zoneRulesetZoneRulesetsListZoneRulesetsResponseResultJSON `json:"-"`
}

// zoneRulesetZoneRulesetsListZoneRulesetsResponseResultJSON contains the JSON
// metadata for the struct [ZoneRulesetZoneRulesetsListZoneRulesetsResponseResult]
type zoneRulesetZoneRulesetsListZoneRulesetsResponseResultJSON struct {
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

func (r *ZoneRulesetZoneRulesetsListZoneRulesetsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultKind string

const (
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultKindManaged ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultKind = "managed"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultKindCustom  ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultKind = "custom"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultKindRoot    ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultKind = "root"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultKindZone    ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase string

const (
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseDdosL4                         ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "ddos_l4"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseDdosL7                         ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "ddos_l7"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPConfigSettings             ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_config_settings"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPCustomErrors               ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_custom_errors"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPLogCustomFields            ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRatelimit                  ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_ratelimit"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_late_transform"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestOrigin              ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_origin"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestRedirect            ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_redirect"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestSanitize            ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_sanitize"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestSbfm                ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_sbfm"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPRequestTransform           ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_request_transform"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPResponseCompression        ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_response_compression"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseMagicTransit                   ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "magic_transit"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhaseMagicTransitManaged            ZoneRulesetZoneRulesetsListZoneRulesetsResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetZoneRulesetsListZoneRulesetsResponseSuccess bool

const (
	ZoneRulesetZoneRulesetsListZoneRulesetsResponseSuccessTrue ZoneRulesetZoneRulesetsListZoneRulesetsResponseSuccess = true
)

type ZoneRulesetUpdateParams struct {
	ID param.Field[interface{}] `json:"id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[ZoneRulesetUpdateParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[ZoneRulesetUpdateParamsPhase] `json:"phase"`
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetUpdateParamsRule] `json:"rules"`
}

func (r ZoneRulesetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type ZoneRulesetUpdateParamsKind string

const (
	ZoneRulesetUpdateParamsKindManaged ZoneRulesetUpdateParamsKind = "managed"
	ZoneRulesetUpdateParamsKindCustom  ZoneRulesetUpdateParamsKind = "custom"
	ZoneRulesetUpdateParamsKindRoot    ZoneRulesetUpdateParamsKind = "root"
	ZoneRulesetUpdateParamsKindZone    ZoneRulesetUpdateParamsKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetUpdateParamsPhase string

const (
	ZoneRulesetUpdateParamsPhaseDdosL4                         ZoneRulesetUpdateParamsPhase = "ddos_l4"
	ZoneRulesetUpdateParamsPhaseDdosL7                         ZoneRulesetUpdateParamsPhase = "ddos_l7"
	ZoneRulesetUpdateParamsPhaseHTTPConfigSettings             ZoneRulesetUpdateParamsPhase = "http_config_settings"
	ZoneRulesetUpdateParamsPhaseHTTPCustomErrors               ZoneRulesetUpdateParamsPhase = "http_custom_errors"
	ZoneRulesetUpdateParamsPhaseHTTPLogCustomFields            ZoneRulesetUpdateParamsPhase = "http_log_custom_fields"
	ZoneRulesetUpdateParamsPhaseHTTPRatelimit                  ZoneRulesetUpdateParamsPhase = "http_ratelimit"
	ZoneRulesetUpdateParamsPhaseHTTPRequestCacheSettings       ZoneRulesetUpdateParamsPhase = "http_request_cache_settings"
	ZoneRulesetUpdateParamsPhaseHTTPRequestDynamicRedirect     ZoneRulesetUpdateParamsPhase = "http_request_dynamic_redirect"
	ZoneRulesetUpdateParamsPhaseHTTPRequestFirewallCustom      ZoneRulesetUpdateParamsPhase = "http_request_firewall_custom"
	ZoneRulesetUpdateParamsPhaseHTTPRequestFirewallManaged     ZoneRulesetUpdateParamsPhase = "http_request_firewall_managed"
	ZoneRulesetUpdateParamsPhaseHTTPRequestLateTransform       ZoneRulesetUpdateParamsPhase = "http_request_late_transform"
	ZoneRulesetUpdateParamsPhaseHTTPRequestOrigin              ZoneRulesetUpdateParamsPhase = "http_request_origin"
	ZoneRulesetUpdateParamsPhaseHTTPRequestRedirect            ZoneRulesetUpdateParamsPhase = "http_request_redirect"
	ZoneRulesetUpdateParamsPhaseHTTPRequestSanitize            ZoneRulesetUpdateParamsPhase = "http_request_sanitize"
	ZoneRulesetUpdateParamsPhaseHTTPRequestSbfm                ZoneRulesetUpdateParamsPhase = "http_request_sbfm"
	ZoneRulesetUpdateParamsPhaseHTTPRequestSelectConfiguration ZoneRulesetUpdateParamsPhase = "http_request_select_configuration"
	ZoneRulesetUpdateParamsPhaseHTTPRequestTransform           ZoneRulesetUpdateParamsPhase = "http_request_transform"
	ZoneRulesetUpdateParamsPhaseHTTPResponseCompression        ZoneRulesetUpdateParamsPhase = "http_response_compression"
	ZoneRulesetUpdateParamsPhaseHTTPResponseFirewallManaged    ZoneRulesetUpdateParamsPhase = "http_response_firewall_managed"
	ZoneRulesetUpdateParamsPhaseHTTPResponseHeadersTransform   ZoneRulesetUpdateParamsPhase = "http_response_headers_transform"
	ZoneRulesetUpdateParamsPhaseMagicTransit                   ZoneRulesetUpdateParamsPhase = "magic_transit"
	ZoneRulesetUpdateParamsPhaseMagicTransitIDsManaged         ZoneRulesetUpdateParamsPhase = "magic_transit_ids_managed"
	ZoneRulesetUpdateParamsPhaseMagicTransitManaged            ZoneRulesetUpdateParamsPhase = "magic_transit_managed"
)

// Satisfied by [ZoneRulesetUpdateParamsRulesOexZd8xKBlockRule],
// [ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRule],
// [ZoneRulesetUpdateParamsRulesOexZd8xKLogRule],
// [ZoneRulesetUpdateParamsRulesOexZd8xKSkipRule].
type ZoneRulesetUpdateParamsRule interface {
	implementsZoneRulesetUpdateParamsRule()
}

type ZoneRulesetUpdateParamsRulesOexZd8xKBlockRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                        `json:"id"`
	Action           param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                   `json:"description"`
	Enabled          param.Field[interface{}]                                                   `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKBlockRule) implementsZoneRulesetUpdateParamsRule() {}

type ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleAction string

const (
	ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleActionBlock ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleAction = "block"
)

type ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParametersResponse] `json:"response"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                          `json:"id"`
	Action           param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                     `json:"description"`
	Enabled          param.Field[interface{}]                                                     `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRule) implementsZoneRulesetUpdateParamsRule() {}

type ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleAction string

const (
	ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionExecute ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleAction = "execute"
)

type ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParameters struct {
	ID param.Field[interface{}] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverrides struct {
	Action param.Field[interface{}] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	Enabled    param.Field[interface{}]                                                                        `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules            param.Field[[]ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule] `json:"rules"`
	SensitivityLevel param.Field[interface{}]                                                                    `json:"sensitivity_level"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory struct {
	Category         param.Field[interface{}] `json:"category,required"`
	Action           param.Field[interface{}] `json:"action"`
	Enabled          param.Field[interface{}] `json:"enabled"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule-level override
type ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule struct {
	ID      param.Field[interface{}] `json:"id,required"`
	Action  param.Field[interface{}] `json:"action"`
	Enabled param.Field[interface{}] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold   param.Field[int64]       `json:"score_threshold"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetUpdateParamsRulesOexZd8xKLogRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                            `json:"id"`
	Action           param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKLogRuleAction] `json:"action"`
	ActionParameters param.Field[interface{}]                                       `json:"action_parameters"`
	Description      param.Field[interface{}]                                       `json:"description"`
	Enabled          param.Field[interface{}]                                       `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKLogRule) implementsZoneRulesetUpdateParamsRule() {}

type ZoneRulesetUpdateParamsRulesOexZd8xKLogRuleAction string

const (
	ZoneRulesetUpdateParamsRulesOexZd8xKLogRuleActionLog ZoneRulesetUpdateParamsRulesOexZd8xKLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type ZoneRulesetUpdateParamsRulesOexZd8xKLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetUpdateParamsRulesOexZd8xKSkipRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                       `json:"id"`
	Action           param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                  `json:"description"`
	Enabled          param.Field[interface{}]                                                  `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKSkipRule) implementsZoneRulesetUpdateParamsRule() {}

type ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleAction string

const (
	ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionSkip ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleAction = "skip"
)

type ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]interface{}] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[interface{}] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]interface{}] `json:"rulesets"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name of a legacy security product to skip the execution of.
type ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct string

const (
	ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductBic           ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "bic"
	ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductHot           ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "hot"
	ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductRateLimit     ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "rateLimit"
	ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductSecurityLevel ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "securityLevel"
	ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductUaBlock       ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "uaBlock"
	ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductWaf           ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "waf"
	ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProductZoneLockdown  ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersRuleset string

const (
	ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersRulesetCurrent ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetUpdateParamsRulesOexZd8xKSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetParams struct {
	ID param.Field[interface{}] `json:"id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase] `json:"phase"`
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule] `json:"rules"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKindManaged ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind = "managed"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKindCustom  ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind = "custom"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKindRoot    ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind = "root"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKindZone    ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseDdosL4                         ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "ddos_l4"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseDdosL7                         ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "ddos_l7"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPConfigSettings             ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_config_settings"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPCustomErrors               ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_custom_errors"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPLogCustomFields            ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_log_custom_fields"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRatelimit                  ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_ratelimit"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestCacheSettings       ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_cache_settings"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestDynamicRedirect     ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_dynamic_redirect"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestFirewallCustom      ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_firewall_custom"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestFirewallManaged     ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_firewall_managed"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestLateTransform       ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_late_transform"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestOrigin              ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_origin"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestRedirect            ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_redirect"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestSanitize            ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_sanitize"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestSbfm                ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_sbfm"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestSelectConfiguration ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_select_configuration"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPRequestTransform           ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_request_transform"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPResponseCompression        ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_response_compression"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPResponseFirewallManaged    ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_response_firewall_managed"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseHTTPResponseHeadersTransform   ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "http_response_headers_transform"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseMagicTransit                   ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "magic_transit"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseMagicTransitIDsManaged         ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "magic_transit_ids_managed"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhaseMagicTransitManaged            ZoneRulesetZoneRulesetsNewAZoneRulesetParamsPhase = "magic_transit_managed"
)

// Satisfied by
// [ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRule],
// [ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRule],
// [ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRule],
// [ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRule].
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule interface {
	implementsZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule()
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                             `json:"id"`
	Action           param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                        `json:"description"`
	Enabled          param.Field[interface{}]                                                                        `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRule) implementsZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule() {
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleAction string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleActionBlock ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleAction = "block"
)

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse] `json:"response"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                               `json:"id"`
	Action           param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                          `json:"description"`
	Enabled          param.Field[interface{}]                                                                          `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRule) implementsZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule() {
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleAction string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionExecute ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleAction = "execute"
)

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParameters struct {
	ID param.Field[interface{}] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverrides struct {
	Action param.Field[interface{}] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	Enabled    param.Field[interface{}]                                                                                             `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules            param.Field[[]ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule] `json:"rules"`
	SensitivityLevel param.Field[interface{}]                                                                                         `json:"sensitivity_level"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory struct {
	Category         param.Field[interface{}] `json:"category,required"`
	Action           param.Field[interface{}] `json:"action"`
	Enabled          param.Field[interface{}] `json:"enabled"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule-level override
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule struct {
	ID      param.Field[interface{}] `json:"id,required"`
	Action  param.Field[interface{}] `json:"action"`
	Enabled param.Field[interface{}] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold   param.Field[int64]       `json:"score_threshold"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                 `json:"id"`
	Action           param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRuleAction] `json:"action"`
	ActionParameters param.Field[interface{}]                                                            `json:"action_parameters"`
	Description      param.Field[interface{}]                                                            `json:"description"`
	Enabled          param.Field[interface{}]                                                            `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRule) implementsZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule() {
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRuleAction string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRuleActionLog ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                            `json:"id"`
	Action           param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                       `json:"description"`
	Enabled          param.Field[interface{}]                                                                       `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRule) implementsZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule() {
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleAction string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionSkip ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleAction = "skip"
)

type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]interface{}] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[interface{}] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]interface{}] `json:"rulesets"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name of a legacy security product to skip the execution of.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductBic           ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "bic"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductHot           ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "hot"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductRateLimit     ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "rateLimit"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductSecurityLevel ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "securityLevel"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductUaBlock       ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "uaBlock"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductWaf           ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "waf"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProductZoneLockdown  ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersRuleset string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersRulesetCurrent ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesOexZd8xKSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
