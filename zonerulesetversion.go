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

// ZoneRulesetVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRulesetVersionService] method
// instead.
type ZoneRulesetVersionService struct {
	Options []option.RequestOption
}

// NewZoneRulesetVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneRulesetVersionService(opts ...option.RequestOption) (r *ZoneRulesetVersionService) {
	r = &ZoneRulesetVersionService{}
	r.Options = opts
	return
}

// Fetches a specific version of a zone ruleset.
func (r *ZoneRulesetVersionService) Get(ctx context.Context, zoneID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (res *ZoneRulesetVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/versions/%s", zoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing version of a zone ruleset.
func (r *ZoneRulesetVersionService) Delete(ctx context.Context, zoneID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/versions/%s", zoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the versions of a zone ruleset.
func (r *ZoneRulesetVersionService) ZoneRulesetsListAZoneRulesetSVersions(ctx context.Context, zoneID string, rulesetID string, opts ...option.RequestOption) (res *ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/versions", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneRulesetVersionGetResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetVersionGetResponseMessage `json:"messages"`
	Result   ZoneRulesetVersionGetResponseResult    `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetVersionGetResponseSuccess `json:"success"`
	JSON    zoneRulesetVersionGetResponseJSON    `json:"-"`
}

// zoneRulesetVersionGetResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetVersionGetResponse]
type zoneRulesetVersionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetVersionGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetVersionGetResponseMessagesSource `json:"source"`
	JSON   zoneRulesetVersionGetResponseMessageJSON    `json:"-"`
}

// zoneRulesetVersionGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneRulesetVersionGetResponseMessage]
type zoneRulesetVersionGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetVersionGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    zoneRulesetVersionGetResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetVersionGetResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneRulesetVersionGetResponseMessagesSource]
type zoneRulesetVersionGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetVersionGetResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetVersionGetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetVersionGetResponseResultPhase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                                  `json:"version"`
	JSON    zoneRulesetVersionGetResponseResultJSON `json:"-"`
}

// zoneRulesetVersionGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneRulesetVersionGetResponseResult]
type zoneRulesetVersionGetResponseResultJSON struct {
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

func (r *ZoneRulesetVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetVersionGetResponseResultKind string

const (
	ZoneRulesetVersionGetResponseResultKindManaged ZoneRulesetVersionGetResponseResultKind = "managed"
	ZoneRulesetVersionGetResponseResultKindCustom  ZoneRulesetVersionGetResponseResultKind = "custom"
	ZoneRulesetVersionGetResponseResultKindRoot    ZoneRulesetVersionGetResponseResultKind = "root"
	ZoneRulesetVersionGetResponseResultKindZone    ZoneRulesetVersionGetResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetVersionGetResponseResultPhase string

const (
	ZoneRulesetVersionGetResponseResultPhaseDdosL4                         ZoneRulesetVersionGetResponseResultPhase = "ddos_l4"
	ZoneRulesetVersionGetResponseResultPhaseDdosL7                         ZoneRulesetVersionGetResponseResultPhase = "ddos_l7"
	ZoneRulesetVersionGetResponseResultPhaseHTTPConfigSettings             ZoneRulesetVersionGetResponseResultPhase = "http_config_settings"
	ZoneRulesetVersionGetResponseResultPhaseHTTPCustomErrors               ZoneRulesetVersionGetResponseResultPhase = "http_custom_errors"
	ZoneRulesetVersionGetResponseResultPhaseHTTPLogCustomFields            ZoneRulesetVersionGetResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRatelimit                  ZoneRulesetVersionGetResponseResultPhase = "http_ratelimit"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetVersionGetResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetVersionGetResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetVersionGetResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetVersionGetResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetVersionGetResponseResultPhase = "http_request_late_transform"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestOrigin              ZoneRulesetVersionGetResponseResultPhase = "http_request_origin"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestRedirect            ZoneRulesetVersionGetResponseResultPhase = "http_request_redirect"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestSanitize            ZoneRulesetVersionGetResponseResultPhase = "http_request_sanitize"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestSbfm                ZoneRulesetVersionGetResponseResultPhase = "http_request_sbfm"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetVersionGetResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetVersionGetResponseResultPhaseHTTPRequestTransform           ZoneRulesetVersionGetResponseResultPhase = "http_request_transform"
	ZoneRulesetVersionGetResponseResultPhaseHTTPResponseCompression        ZoneRulesetVersionGetResponseResultPhase = "http_response_compression"
	ZoneRulesetVersionGetResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetVersionGetResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetVersionGetResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetVersionGetResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetVersionGetResponseResultPhaseMagicTransit                   ZoneRulesetVersionGetResponseResultPhase = "magic_transit"
	ZoneRulesetVersionGetResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetVersionGetResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetVersionGetResponseResultPhaseMagicTransitManaged            ZoneRulesetVersionGetResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetVersionGetResponseSuccess bool

const (
	ZoneRulesetVersionGetResponseSuccessTrue ZoneRulesetVersionGetResponseSuccess = true
)

type ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponse struct {
	Errors interface{} `json:"errors"`
	// A list of warning messages.
	Messages []ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessage `json:"messages"`
	// A list of rulesets. The returned information will not include the rules in each
	// ruleset.
	Result []ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResult `json:"result"`
	// Whether the API call was successful.
	Success ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseSuccess `json:"success"`
	JSON    zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseJSON    `json:"-"`
}

// zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseJSON contains the
// JSON metadata for the struct
// [ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponse]
type zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessagesSource `json:"source"`
	JSON   zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessageJSON    `json:"-"`
}

// zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessage]
type zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                            `json:"pointer,required"`
	JSON    zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessagesSource]
type zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResult struct {
	ID interface{} `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase `json:"phase"`
	// The version of the ruleset.
	Version string                                                                    `json:"version"`
	JSON    zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultJSON `json:"-"`
}

// zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResult]
type zoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultJSON struct {
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

func (r *ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultKind string

const (
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultKindManaged ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultKind = "managed"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultKindCustom  ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultKind = "custom"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultKindRoot    ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultKind = "root"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultKindZone    ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultKind = "zone"
)

// The phase of the ruleset.
type ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase string

const (
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseDdosL4                         ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "ddos_l4"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseDdosL7                         ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "ddos_l7"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPConfigSettings             ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_config_settings"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPCustomErrors               ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_custom_errors"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPLogCustomFields            ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_log_custom_fields"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRatelimit                  ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_ratelimit"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestCacheSettings       ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_cache_settings"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestDynamicRedirect     ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_dynamic_redirect"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestFirewallCustom      ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_firewall_custom"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestFirewallManaged     ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_firewall_managed"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestLateTransform       ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_late_transform"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestOrigin              ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_origin"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestRedirect            ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_redirect"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestSanitize            ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_sanitize"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestSbfm                ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_sbfm"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestSelectConfiguration ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_select_configuration"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPRequestTransform           ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_request_transform"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPResponseCompression        ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_response_compression"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPResponseFirewallManaged    ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_response_firewall_managed"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseHTTPResponseHeadersTransform   ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "http_response_headers_transform"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseMagicTransit                   ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "magic_transit"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseMagicTransitIDsManaged         ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "magic_transit_ids_managed"
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhaseMagicTransitManaged            ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseResultPhase = "magic_transit_managed"
)

// Whether the API call was successful.
type ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseSuccess bool

const (
	ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseSuccessTrue ZoneRulesetVersionZoneRulesetsListAZoneRulesetSVersionsResponseSuccess = true
)
