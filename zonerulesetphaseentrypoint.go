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

// ZoneRulesetPhaseEntrypointService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneRulesetPhaseEntrypointService] method instead.
type ZoneRulesetPhaseEntrypointService struct {
	Options  []option.RequestOption
	Versions *ZoneRulesetPhaseEntrypointVersionService
}

// NewZoneRulesetPhaseEntrypointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneRulesetPhaseEntrypointService(opts ...option.RequestOption) (r *ZoneRulesetPhaseEntrypointService) {
	r = &ZoneRulesetPhaseEntrypointService{}
	r.Options = opts
	r.Versions = NewZoneRulesetPhaseEntrypointVersionService(opts...)
	return
}

// Fetches the latest version of the zone entry point ruleset for a given phase.
func (r *ZoneRulesetPhaseEntrypointService) TransformRulesListTransformRules(ctx context.Context, zoneID string, rulesetPhase ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase, opts ...option.RequestOption) (res *ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/%v/entrypoint", zoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a zone entry point ruleset, creating a new version.
func (r *ZoneRulesetPhaseEntrypointService) TransformRulesUpdateTransformRules(ctx context.Context, zoneID string, rulesetPhase ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase, body ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParams, opts ...option.RequestOption) (res *ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/%v/entrypoint", zoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessage `json:"messages"`
	Result   ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseSuccess `json:"success"`
	JSON    zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseJSON contains
// the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponse]
type zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessagesSource `json:"source"`
	JSON   zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessageJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessage]
type zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                               `json:"pointer,required"`
	JSON    zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessagesSource]
type zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                                                       `json:"version"`
	JSON    zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResult]
type zoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultJSON struct {
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

func (r *ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultKind string

const (
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultKindManaged ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultKind = "managed"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultKindCustom  ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultKind = "custom"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultKindRoot    ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultKind = "root"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultKindZone    ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase string

const (
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseDdosL4                         ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "ddos_l4"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseDdosL7                         ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "ddos_l7"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPConfigSettings             ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_config_settings"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPCustomErrors               ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_custom_errors"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPLogCustomFields            ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRatelimit                  ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_ratelimit"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestOrigin              ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_origin"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestRedirect            ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_redirect"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestSanitize            ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_sanitize"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestSbfm                ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_sbfm"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPRequestTransform           ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPResponseCompression        ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_response_compression"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseMagicTransit                   ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "magic_transit"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhaseMagicTransitManaged            ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseSuccess bool

const (
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseSuccessTrue ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesResponseSuccess = true
)

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessage `json:"messages"`
	Result   ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseSuccess `json:"success"`
	JSON    zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponse]
type zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessagesSource `json:"source"`
	JSON   zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessageJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessage]
type zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                                 `json:"pointer,required"`
	JSON    zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessagesSource]
type zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                                                         `json:"version"`
	JSON    zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResult]
type zoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultJSON struct {
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

func (r *ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultKind string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultKindManaged ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultKind = "managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultKindCustom  ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultKind = "custom"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultKindRoot    ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultKind = "root"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultKindZone    ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseDdosL4                         ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "ddos_l4"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseDdosL7                         ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "ddos_l7"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPConfigSettings             ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_config_settings"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPCustomErrors               ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_custom_errors"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPLogCustomFields            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRatelimit                  ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_ratelimit"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestOrigin              ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_origin"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestRedirect            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_redirect"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestSanitize            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_sanitize"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestSbfm                ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_sbfm"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPRequestTransform           ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPResponseCompression        ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_response_compression"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseMagicTransit                   ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "magic_transit"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhaseMagicTransitManaged            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseSuccess bool

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseSuccessTrue ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesResponseSuccess = true
)

// The phase of the ruleset.
type ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase string

const (
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseDdosL4                         ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "ddos_l4"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseDdosL7                         ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "ddos_l7"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPConfigSettings             ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_config_settings"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPCustomErrors               ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_custom_errors"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPLogCustomFields            ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_log_custom_fields"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRatelimit                  ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_ratelimit"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestCacheSettings       ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_cache_settings"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestDynamicRedirect     ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_dynamic_redirect"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestFirewallCustom      ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_firewall_custom"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestFirewallManaged     ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_firewall_managed"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestLateTransform       ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestOrigin              ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_origin"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestRedirect            ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_redirect"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestSanitize            ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_sanitize"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestSbfm                ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_sbfm"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestSelectConfiguration ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_select_configuration"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPRequestTransform           ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPResponseCompression        ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_response_compression"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPResponseFirewallManaged    ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_response_firewall_managed"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseHTTPResponseHeadersTransform   ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "http_response_headers_transform"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseMagicTransit                   ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "magic_transit"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseMagicTransitIDsManaged         ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "magic_transit_ids_managed"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhaseMagicTransitManaged            ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsRulesetPhase = "magic_transit_managed"
)

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParams struct {
	ID param.Field[interface{}] `json:"id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase] `json:"phase"`
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule] `json:"rules"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The phase of the ruleset.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseDdosL4                         ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "ddos_l4"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseDdosL7                         ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "ddos_l7"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPConfigSettings             ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_config_settings"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPCustomErrors               ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_custom_errors"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPLogCustomFields            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_log_custom_fields"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRatelimit                  ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_ratelimit"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestCacheSettings       ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_cache_settings"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestDynamicRedirect     ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_dynamic_redirect"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestFirewallCustom      ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_firewall_custom"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestFirewallManaged     ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_firewall_managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestLateTransform       ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestOrigin              ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_origin"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestRedirect            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_redirect"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestSanitize            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_sanitize"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestSbfm                ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_sbfm"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestSelectConfiguration ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_select_configuration"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPRequestTransform           ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPResponseCompression        ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_response_compression"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPResponseFirewallManaged    ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_response_firewall_managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseHTTPResponseHeadersTransform   ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "http_response_headers_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseMagicTransit                   ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "magic_transit"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseMagicTransitIDsManaged         ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "magic_transit_ids_managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhaseMagicTransitManaged            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesetPhase = "magic_transit_managed"
)

// The kind of the ruleset.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKind string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKindManaged ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKind = "managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKindCustom  ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKind = "custom"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKindRoot    ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKind = "root"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKindZone    ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseDdosL4                         ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "ddos_l4"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseDdosL7                         ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "ddos_l7"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPConfigSettings             ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_config_settings"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPCustomErrors               ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_custom_errors"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPLogCustomFields            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_log_custom_fields"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRatelimit                  ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_ratelimit"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestCacheSettings       ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_cache_settings"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestDynamicRedirect     ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_dynamic_redirect"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestFirewallCustom      ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_firewall_custom"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestFirewallManaged     ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_firewall_managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestLateTransform       ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestOrigin              ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_origin"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestRedirect            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_redirect"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestSanitize            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_sanitize"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestSbfm                ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_sbfm"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestSelectConfiguration ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_select_configuration"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestTransform           ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPResponseCompression        ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_response_compression"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPResponseFirewallManaged    ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_response_firewall_managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPResponseHeadersTransform   ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_response_headers_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseMagicTransit                   ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "magic_transit"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseMagicTransitIDsManaged         ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "magic_transit_ids_managed"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseMagicTransitManaged            ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "magic_transit_managed"
)

// Satisfied by
// [ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRule],
// [ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRule],
// [ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRule],
// [ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRule].
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule interface {
	implementsZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule()
}

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                                   `json:"id"`
	Action           param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                              `json:"description"`
	Enabled          param.Field[interface{}]                                                                                              `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRule) implementsZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule() {
}

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleAction string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionBlock ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleAction = "block"
)

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParametersResponse] `json:"response"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKBlockRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                                     `json:"id"`
	Action           param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                                `json:"description"`
	Enabled          param.Field[interface{}]                                                                                                `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRule) implementsZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule() {
}

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleAction string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionExecute ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleAction = "execute"
)

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParameters struct {
	ID param.Field[interface{}] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersOverrides struct {
	Action param.Field[interface{}] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	Enabled    param.Field[interface{}]                                                                                                                   `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules            param.Field[[]ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule] `json:"rules"`
	SensitivityLevel param.Field[interface{}]                                                                                                               `json:"sensitivity_level"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory struct {
	Category         param.Field[interface{}] `json:"category,required"`
	Action           param.Field[interface{}] `json:"action"`
	Enabled          param.Field[interface{}] `json:"enabled"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule-level override
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule struct {
	ID      param.Field[interface{}] `json:"id,required"`
	Action  param.Field[interface{}] `json:"action"`
	Enabled param.Field[interface{}] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold   param.Field[int64]       `json:"score_threshold"`
	SensitivityLevel param.Field[interface{}] `json:"sensitivity_level"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKExecuteRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                       `json:"id"`
	Action           param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRuleAction] `json:"action"`
	ActionParameters param.Field[interface{}]                                                                                  `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                  `json:"description"`
	Enabled          param.Field[interface{}]                                                                                  `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRule) implementsZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule() {
}

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRuleAction string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRuleActionLog ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKLogRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRule struct {
	// The unique ID of the rule.
	ID               param.Field[string]                                                                                                  `json:"id"`
	Action           param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParameters] `json:"action_parameters"`
	Description      param.Field[interface{}]                                                                                             `json:"description"`
	Enabled          param.Field[interface{}]                                                                                             `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRule) implementsZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule() {
}

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleAction string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionSkip ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleAction = "skip"
)

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]interface{}] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[interface{}] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]interface{}] `json:"rulesets"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name of a legacy security product to skip the execution of.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProduct string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProductBic           ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProduct = "bic"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProductHot           ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProduct = "hot"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProductRateLimit     ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProduct = "rateLimit"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProductSecurityLevel ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProduct = "securityLevel"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProductUaBlock       ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProduct = "uaBlock"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProductWaf           ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProduct = "waf"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProductZoneLockdown  ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersRuleset string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersRulesetCurrent ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesOexZd8xKSkipRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
