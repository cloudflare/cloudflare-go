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

// ZoneRulesetPhaseEntrypointVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneRulesetPhaseEntrypointVersionService] method instead.
type ZoneRulesetPhaseEntrypointVersionService struct {
	Options []option.RequestOption
}

// NewZoneRulesetPhaseEntrypointVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneRulesetPhaseEntrypointVersionService(opts ...option.RequestOption) (r *ZoneRulesetPhaseEntrypointVersionService) {
	r = &ZoneRulesetPhaseEntrypointVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of a zone entry point ruleset.
func (r *ZoneRulesetPhaseEntrypointVersionService) Get(ctx context.Context, zoneID string, rulesetPhase ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase, rulesetVersion string, opts ...option.RequestOption) (res *ZoneRulesetPhaseEntrypointVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/%v/entrypoint/versions/%s", zoneID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the versions of a zone entry point ruleset.
func (r *ZoneRulesetPhaseEntrypointVersionService) ZoneRulesetsListAZoneEntryPointRulesetSVersions(ctx context.Context, zoneID string, rulesetPhase ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase, opts ...option.RequestOption) (res *ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/%v/entrypoint/versions", zoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneRulesetPhaseEntrypointVersionGetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetPhaseEntrypointVersionGetResponseMessage `json:"messages"`
	Result   ZoneRulesetPhaseEntrypointVersionGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetPhaseEntrypointVersionGetResponseSuccess `json:"success"`
	JSON    zoneRulesetPhaseEntrypointVersionGetResponseJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionGetResponseJSON contains the JSON metadata for
// the struct [ZoneRulesetPhaseEntrypointVersionGetResponse]
type zoneRulesetPhaseEntrypointVersionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetPhaseEntrypointVersionGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetPhaseEntrypointVersionGetResponseMessagesSource `json:"source"`
	JSON   zoneRulesetPhaseEntrypointVersionGetResponseMessageJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionGetResponseMessageJSON contains the JSON
// metadata for the struct [ZoneRulesetPhaseEntrypointVersionGetResponseMessage]
type zoneRulesetPhaseEntrypointVersionGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetPhaseEntrypointVersionGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                         `json:"pointer,required"`
	JSON    zoneRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON contains the JSON
// metadata for the struct
// [ZoneRulesetPhaseEntrypointVersionGetResponseMessagesSource]
type zoneRulesetPhaseEntrypointVersionGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetPhaseEntrypointVersionGetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetPhaseEntrypointVersionGetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                                 `json:"version"`
	JSON    zoneRulesetPhaseEntrypointVersionGetResponseResultJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionGetResponseResultJSON contains the JSON
// metadata for the struct [ZoneRulesetPhaseEntrypointVersionGetResponseResult]
type zoneRulesetPhaseEntrypointVersionGetResponseResultJSON struct {
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

func (r *ZoneRulesetPhaseEntrypointVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetPhaseEntrypointVersionGetResponseResultKind string

const (
	ZoneRulesetPhaseEntrypointVersionGetResponseResultKindManaged ZoneRulesetPhaseEntrypointVersionGetResponseResultKind = "managed"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultKindCustom  ZoneRulesetPhaseEntrypointVersionGetResponseResultKind = "custom"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultKindRoot    ZoneRulesetPhaseEntrypointVersionGetResponseResultKind = "root"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultKindZone    ZoneRulesetPhaseEntrypointVersionGetResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase string

const (
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseDdosL4                         ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "ddos_l4"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseDdosL7                         ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "ddos_l7"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPConfigSettings             ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_config_settings"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPCustomErrors               ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_custom_errors"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPLogCustomFields            ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRatelimit                  ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_ratelimit"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestOrigin              ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_origin"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestRedirect            ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_redirect"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestSanitize            ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_sanitize"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestSbfm                ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_sbfm"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPRequestTransform           ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPResponseCompression        ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_response_compression"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseMagicTransit                   ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "magic_transit"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetPhaseEntrypointVersionGetResponseResultPhaseMagicTransitManaged            ZoneRulesetPhaseEntrypointVersionGetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetPhaseEntrypointVersionGetResponseSuccess bool

const (
	ZoneRulesetPhaseEntrypointVersionGetResponseSuccessTrue ZoneRulesetPhaseEntrypointVersionGetResponseSuccess = true
)

type ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessage `json:"messages"`
	// A list of rulesets. The returned information will not include the rules in each
	// ruleset.
	Result []ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResult `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseSuccess `json:"success"`
	JSON    zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponse]
type zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessagesSource `json:"source"`
	JSON   zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessageJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessage]
type zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                                                     `json:"pointer,required"`
	JSON    zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessagesSource]
type zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase `json:"phase"`
	// The version of the ruleset.
	Version string                                                                                             `json:"version"`
	JSON    zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResult]
type zoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultJSON struct {
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

func (r *ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultKind string

const (
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultKindManaged ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultKind = "managed"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultKindCustom  ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultKind = "custom"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultKindRoot    ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultKind = "root"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultKindZone    ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase string

const (
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseDdosL4                         ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "ddos_l4"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseDdosL7                         ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "ddos_l7"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPConfigSettings             ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_config_settings"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPCustomErrors               ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_custom_errors"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPLogCustomFields            ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRatelimit                  ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_ratelimit"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestOrigin              ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_origin"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestRedirect            ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_redirect"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestSanitize            ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_sanitize"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestSbfm                ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_sbfm"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPRequestTransform           ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPResponseCompression        ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_response_compression"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseMagicTransit                   ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "magic_transit"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhaseMagicTransitManaged            ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseSuccess bool

const (
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseSuccessTrue ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsResponseSuccess = true
)

// The phase of the ruleset.
type ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase string

const (
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseDdosL4                         ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "ddos_l4"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseDdosL7                         ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "ddos_l7"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPConfigSettings             ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_config_settings"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPCustomErrors               ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_custom_errors"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPLogCustomFields            ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_log_custom_fields"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRatelimit                  ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_ratelimit"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestCacheSettings       ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_cache_settings"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestDynamicRedirect     ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_dynamic_redirect"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestFirewallCustom      ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_firewall_custom"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestFirewallManaged     ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_firewall_managed"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestLateTransform       ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestOrigin              ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_origin"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestRedirect            ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_redirect"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestSanitize            ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_sanitize"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestSbfm                ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_sbfm"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestSelectConfiguration ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_select_configuration"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPRequestTransform           ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPResponseCompression        ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_response_compression"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPResponseFirewallManaged    ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_response_firewall_managed"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseHTTPResponseHeadersTransform   ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "http_response_headers_transform"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseMagicTransit                   ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "magic_transit"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseMagicTransitIDsManaged         ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "magic_transit_ids_managed"
	ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhaseMagicTransitManaged            ZoneRulesetPhaseEntrypointVersionGetParamsRulesetPhase = "magic_transit_managed"
)

// The phase of the ruleset.
type ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase string

const (
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseDdosL4                         ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "ddos_l4"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseDdosL7                         ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "ddos_l7"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPConfigSettings             ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_config_settings"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPCustomErrors               ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_custom_errors"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPLogCustomFields            ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_log_custom_fields"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRatelimit                  ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_ratelimit"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestCacheSettings       ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_cache_settings"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestDynamicRedirect     ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_dynamic_redirect"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestFirewallCustom      ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_firewall_custom"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestFirewallManaged     ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_firewall_managed"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestLateTransform       ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestOrigin              ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_origin"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestRedirect            ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_redirect"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestSanitize            ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_sanitize"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestSbfm                ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_sbfm"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestSelectConfiguration ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_select_configuration"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPRequestTransform           ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPResponseCompression        ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_response_compression"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPResponseFirewallManaged    ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_response_firewall_managed"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseHTTPResponseHeadersTransform   ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "http_response_headers_transform"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseMagicTransit                   ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "magic_transit"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseMagicTransitIDsManaged         ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "magic_transit_ids_managed"
	ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhaseMagicTransitManaged            ZoneRulesetPhaseEntrypointVersionZoneRulesetsListAZoneEntryPointRulesetSVersionsParamsRulesetPhase = "magic_transit_managed"
)
