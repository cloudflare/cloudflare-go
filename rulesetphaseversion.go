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

// RulesetPhaseVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRulesetPhaseVersionService]
// method instead.
type RulesetPhaseVersionService struct {
	Options []option.RequestOption
}

// NewRulesetPhaseVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRulesetPhaseVersionService(opts ...option.RequestOption) (r *RulesetPhaseVersionService) {
	r = &RulesetPhaseVersionService{}
	r.Options = opts
	return
}

// Fetches the versions of an account or zone entry point ruleset.
func (r *RulesetPhaseVersionService) List(ctx context.Context, rulesetPhase RulesetPhaseVersionListParamsRulesetPhase, query RulesetPhaseVersionListParams, opts ...option.RequestOption) (res *RulesetsRulesetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetPhaseVersionListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint/versions", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a specific version of an account or zone entry point ruleset.
func (r *RulesetPhaseVersionService) Get(ctx context.Context, rulesetPhase RulesetPhaseVersionGetParamsRulesetPhase, rulesetVersion string, query RulesetPhaseVersionGetParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetPhaseVersionGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint/versions/%s", accountOrZone, accountOrZoneID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type RulesetPhaseVersionListResponse struct {
	// The kind of the ruleset.
	Kind RulesetPhaseVersionListResponseKind `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetPhaseVersionListResponsePhase `json:"phase,required"`
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the ruleset.
	Version string                              `json:"version"`
	JSON    rulesetPhaseVersionListResponseJSON `json:"-"`
}

// rulesetPhaseVersionListResponseJSON contains the JSON metadata for the struct
// [RulesetPhaseVersionListResponse]
type rulesetPhaseVersionListResponseJSON struct {
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	ID          apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetPhaseVersionListResponseKind string

const (
	RulesetPhaseVersionListResponseKindManaged RulesetPhaseVersionListResponseKind = "managed"
	RulesetPhaseVersionListResponseKindCustom  RulesetPhaseVersionListResponseKind = "custom"
	RulesetPhaseVersionListResponseKindRoot    RulesetPhaseVersionListResponseKind = "root"
	RulesetPhaseVersionListResponseKindZone    RulesetPhaseVersionListResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetPhaseVersionListResponsePhase string

const (
	RulesetPhaseVersionListResponsePhaseDDOSL4                         RulesetPhaseVersionListResponsePhase = "ddos_l4"
	RulesetPhaseVersionListResponsePhaseDDOSL7                         RulesetPhaseVersionListResponsePhase = "ddos_l7"
	RulesetPhaseVersionListResponsePhaseHTTPConfigSettings             RulesetPhaseVersionListResponsePhase = "http_config_settings"
	RulesetPhaseVersionListResponsePhaseHTTPCustomErrors               RulesetPhaseVersionListResponsePhase = "http_custom_errors"
	RulesetPhaseVersionListResponsePhaseHTTPLogCustomFields            RulesetPhaseVersionListResponsePhase = "http_log_custom_fields"
	RulesetPhaseVersionListResponsePhaseHTTPRatelimit                  RulesetPhaseVersionListResponsePhase = "http_ratelimit"
	RulesetPhaseVersionListResponsePhaseHTTPRequestCacheSettings       RulesetPhaseVersionListResponsePhase = "http_request_cache_settings"
	RulesetPhaseVersionListResponsePhaseHTTPRequestDynamicRedirect     RulesetPhaseVersionListResponsePhase = "http_request_dynamic_redirect"
	RulesetPhaseVersionListResponsePhaseHTTPRequestFirewallCustom      RulesetPhaseVersionListResponsePhase = "http_request_firewall_custom"
	RulesetPhaseVersionListResponsePhaseHTTPRequestFirewallManaged     RulesetPhaseVersionListResponsePhase = "http_request_firewall_managed"
	RulesetPhaseVersionListResponsePhaseHTTPRequestLateTransform       RulesetPhaseVersionListResponsePhase = "http_request_late_transform"
	RulesetPhaseVersionListResponsePhaseHTTPRequestOrigin              RulesetPhaseVersionListResponsePhase = "http_request_origin"
	RulesetPhaseVersionListResponsePhaseHTTPRequestRedirect            RulesetPhaseVersionListResponsePhase = "http_request_redirect"
	RulesetPhaseVersionListResponsePhaseHTTPRequestSanitize            RulesetPhaseVersionListResponsePhase = "http_request_sanitize"
	RulesetPhaseVersionListResponsePhaseHTTPRequestSbfm                RulesetPhaseVersionListResponsePhase = "http_request_sbfm"
	RulesetPhaseVersionListResponsePhaseHTTPRequestSelectConfiguration RulesetPhaseVersionListResponsePhase = "http_request_select_configuration"
	RulesetPhaseVersionListResponsePhaseHTTPRequestTransform           RulesetPhaseVersionListResponsePhase = "http_request_transform"
	RulesetPhaseVersionListResponsePhaseHTTPResponseCompression        RulesetPhaseVersionListResponsePhase = "http_response_compression"
	RulesetPhaseVersionListResponsePhaseHTTPResponseFirewallManaged    RulesetPhaseVersionListResponsePhase = "http_response_firewall_managed"
	RulesetPhaseVersionListResponsePhaseHTTPResponseHeadersTransform   RulesetPhaseVersionListResponsePhase = "http_response_headers_transform"
	RulesetPhaseVersionListResponsePhaseMagicTransit                   RulesetPhaseVersionListResponsePhase = "magic_transit"
	RulesetPhaseVersionListResponsePhaseMagicTransitIDsManaged         RulesetPhaseVersionListResponsePhase = "magic_transit_ids_managed"
	RulesetPhaseVersionListResponsePhaseMagicTransitManaged            RulesetPhaseVersionListResponsePhase = "magic_transit_managed"
)

type RulesetPhaseVersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type RulesetPhaseVersionListParamsRulesetPhase string

const (
	RulesetPhaseVersionListParamsRulesetPhaseDDOSL4                         RulesetPhaseVersionListParamsRulesetPhase = "ddos_l4"
	RulesetPhaseVersionListParamsRulesetPhaseDDOSL7                         RulesetPhaseVersionListParamsRulesetPhase = "ddos_l7"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPConfigSettings             RulesetPhaseVersionListParamsRulesetPhase = "http_config_settings"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPCustomErrors               RulesetPhaseVersionListParamsRulesetPhase = "http_custom_errors"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPLogCustomFields            RulesetPhaseVersionListParamsRulesetPhase = "http_log_custom_fields"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRatelimit                  RulesetPhaseVersionListParamsRulesetPhase = "http_ratelimit"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestCacheSettings       RulesetPhaseVersionListParamsRulesetPhase = "http_request_cache_settings"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestDynamicRedirect     RulesetPhaseVersionListParamsRulesetPhase = "http_request_dynamic_redirect"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestFirewallCustom      RulesetPhaseVersionListParamsRulesetPhase = "http_request_firewall_custom"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestFirewallManaged     RulesetPhaseVersionListParamsRulesetPhase = "http_request_firewall_managed"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestLateTransform       RulesetPhaseVersionListParamsRulesetPhase = "http_request_late_transform"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestOrigin              RulesetPhaseVersionListParamsRulesetPhase = "http_request_origin"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestRedirect            RulesetPhaseVersionListParamsRulesetPhase = "http_request_redirect"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestSanitize            RulesetPhaseVersionListParamsRulesetPhase = "http_request_sanitize"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestSbfm                RulesetPhaseVersionListParamsRulesetPhase = "http_request_sbfm"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestSelectConfiguration RulesetPhaseVersionListParamsRulesetPhase = "http_request_select_configuration"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestTransform           RulesetPhaseVersionListParamsRulesetPhase = "http_request_transform"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPResponseCompression        RulesetPhaseVersionListParamsRulesetPhase = "http_response_compression"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPResponseFirewallManaged    RulesetPhaseVersionListParamsRulesetPhase = "http_response_firewall_managed"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPResponseHeadersTransform   RulesetPhaseVersionListParamsRulesetPhase = "http_response_headers_transform"
	RulesetPhaseVersionListParamsRulesetPhaseMagicTransit                   RulesetPhaseVersionListParamsRulesetPhase = "magic_transit"
	RulesetPhaseVersionListParamsRulesetPhaseMagicTransitIDsManaged         RulesetPhaseVersionListParamsRulesetPhase = "magic_transit_ids_managed"
	RulesetPhaseVersionListParamsRulesetPhaseMagicTransitManaged            RulesetPhaseVersionListParamsRulesetPhase = "magic_transit_managed"
)

// A response object.
type RulesetPhaseVersionListResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetPhaseVersionListResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetPhaseVersionListResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetsResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetPhaseVersionListResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetPhaseVersionListResponseEnvelopeJSON    `json:"-"`
}

// rulesetPhaseVersionListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RulesetPhaseVersionListResponseEnvelope]
type rulesetPhaseVersionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseVersionListResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseVersionListResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetPhaseVersionListResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetPhaseVersionListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RulesetPhaseVersionListResponseEnvelopeErrors]
type rulesetPhaseVersionListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseVersionListResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                  `json:"pointer,required"`
	JSON    rulesetPhaseVersionListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetPhaseVersionListResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionListResponseEnvelopeErrorsSource]
type rulesetPhaseVersionListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseVersionListResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseVersionListResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetPhaseVersionListResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetPhaseVersionListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RulesetPhaseVersionListResponseEnvelopeMessages]
type rulesetPhaseVersionListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseVersionListResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                    `json:"pointer,required"`
	JSON    rulesetPhaseVersionListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetPhaseVersionListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionListResponseEnvelopeMessagesSource]
type rulesetPhaseVersionListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetPhaseVersionListResponseEnvelopeSuccess bool

const (
	RulesetPhaseVersionListResponseEnvelopeSuccessTrue RulesetPhaseVersionListResponseEnvelopeSuccess = true
)

type RulesetPhaseVersionGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type RulesetPhaseVersionGetParamsRulesetPhase string

const (
	RulesetPhaseVersionGetParamsRulesetPhaseDDOSL4                         RulesetPhaseVersionGetParamsRulesetPhase = "ddos_l4"
	RulesetPhaseVersionGetParamsRulesetPhaseDDOSL7                         RulesetPhaseVersionGetParamsRulesetPhase = "ddos_l7"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPConfigSettings             RulesetPhaseVersionGetParamsRulesetPhase = "http_config_settings"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPCustomErrors               RulesetPhaseVersionGetParamsRulesetPhase = "http_custom_errors"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPLogCustomFields            RulesetPhaseVersionGetParamsRulesetPhase = "http_log_custom_fields"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRatelimit                  RulesetPhaseVersionGetParamsRulesetPhase = "http_ratelimit"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestCacheSettings       RulesetPhaseVersionGetParamsRulesetPhase = "http_request_cache_settings"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestDynamicRedirect     RulesetPhaseVersionGetParamsRulesetPhase = "http_request_dynamic_redirect"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallCustom      RulesetPhaseVersionGetParamsRulesetPhase = "http_request_firewall_custom"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallManaged     RulesetPhaseVersionGetParamsRulesetPhase = "http_request_firewall_managed"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestLateTransform       RulesetPhaseVersionGetParamsRulesetPhase = "http_request_late_transform"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestOrigin              RulesetPhaseVersionGetParamsRulesetPhase = "http_request_origin"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestRedirect            RulesetPhaseVersionGetParamsRulesetPhase = "http_request_redirect"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestSanitize            RulesetPhaseVersionGetParamsRulesetPhase = "http_request_sanitize"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestSbfm                RulesetPhaseVersionGetParamsRulesetPhase = "http_request_sbfm"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestSelectConfiguration RulesetPhaseVersionGetParamsRulesetPhase = "http_request_select_configuration"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestTransform           RulesetPhaseVersionGetParamsRulesetPhase = "http_request_transform"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPResponseCompression        RulesetPhaseVersionGetParamsRulesetPhase = "http_response_compression"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPResponseFirewallManaged    RulesetPhaseVersionGetParamsRulesetPhase = "http_response_firewall_managed"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPResponseHeadersTransform   RulesetPhaseVersionGetParamsRulesetPhase = "http_response_headers_transform"
	RulesetPhaseVersionGetParamsRulesetPhaseMagicTransit                   RulesetPhaseVersionGetParamsRulesetPhase = "magic_transit"
	RulesetPhaseVersionGetParamsRulesetPhaseMagicTransitIDsManaged         RulesetPhaseVersionGetParamsRulesetPhase = "magic_transit_ids_managed"
	RulesetPhaseVersionGetParamsRulesetPhaseMagicTransitManaged            RulesetPhaseVersionGetParamsRulesetPhase = "magic_transit_managed"
)

// A response object.
type RulesetPhaseVersionGetResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetPhaseVersionGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetPhaseVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetPhaseVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetPhaseVersionGetResponseEnvelopeJSON    `json:"-"`
}

// rulesetPhaseVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RulesetPhaseVersionGetResponseEnvelope]
type rulesetPhaseVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseVersionGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseVersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetPhaseVersionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetPhaseVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RulesetPhaseVersionGetResponseEnvelopeErrors]
type rulesetPhaseVersionGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseVersionGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                 `json:"pointer,required"`
	JSON    rulesetPhaseVersionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionGetResponseEnvelopeErrorsSource]
type rulesetPhaseVersionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseVersionGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseVersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetPhaseVersionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetPhaseVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RulesetPhaseVersionGetResponseEnvelopeMessages]
type rulesetPhaseVersionGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseVersionGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                   `json:"pointer,required"`
	JSON    rulesetPhaseVersionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionGetResponseEnvelopeMessagesSource]
type rulesetPhaseVersionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetPhaseVersionGetResponseEnvelopeSuccess bool

const (
	RulesetPhaseVersionGetResponseEnvelopeSuccessTrue RulesetPhaseVersionGetResponseEnvelopeSuccess = true
)
