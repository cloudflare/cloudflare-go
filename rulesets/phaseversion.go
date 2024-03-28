// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PhaseVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPhaseVersionService] method
// instead.
type PhaseVersionService struct {
	Options []option.RequestOption
}

// NewPhaseVersionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPhaseVersionService(opts ...option.RequestOption) (r *PhaseVersionService) {
	r = &PhaseVersionService{}
	r.Options = opts
	return
}

// Fetches the versions of an account or zone entry point ruleset.
func (r *PhaseVersionService) List(ctx context.Context, rulesetPhase PhaseVersionListParamsRulesetPhase, query PhaseVersionListParams, opts ...option.RequestOption) (res *shared.SinglePage[PhaseVersionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches the versions of an account or zone entry point ruleset.
func (r *PhaseVersionService) ListAutoPaging(ctx context.Context, rulesetPhase PhaseVersionListParamsRulesetPhase, query PhaseVersionListParams, opts ...option.RequestOption) *shared.SinglePageAutoPager[PhaseVersionListResponse] {
	return shared.NewSinglePageAutoPager(r.List(ctx, rulesetPhase, query, opts...))
}

// Fetches a specific version of an account or zone entry point ruleset.
func (r *PhaseVersionService) Get(ctx context.Context, rulesetPhase PhaseVersionGetParamsRulesetPhase, rulesetVersion string, query PhaseVersionGetParams, opts ...option.RequestOption) (res *Ruleset, err error) {
	opts = append(r.Options[:], opts...)
	var env PhaseVersionGetResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type PhaseVersionListResponse struct {
	// The kind of the ruleset.
	Kind PhaseVersionListResponseKind `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase PhaseVersionListResponsePhase `json:"phase,required"`
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the ruleset.
	Version string                       `json:"version"`
	JSON    phaseVersionListResponseJSON `json:"-"`
}

// phaseVersionListResponseJSON contains the JSON metadata for the struct
// [PhaseVersionListResponse]
type phaseVersionListResponseJSON struct {
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

func (r *PhaseVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionListResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type PhaseVersionListResponseKind string

const (
	PhaseVersionListResponseKindManaged PhaseVersionListResponseKind = "managed"
	PhaseVersionListResponseKindCustom  PhaseVersionListResponseKind = "custom"
	PhaseVersionListResponseKindRoot    PhaseVersionListResponseKind = "root"
	PhaseVersionListResponseKindZone    PhaseVersionListResponseKind = "zone"
)

func (r PhaseVersionListResponseKind) IsKnown() bool {
	switch r {
	case PhaseVersionListResponseKindManaged, PhaseVersionListResponseKindCustom, PhaseVersionListResponseKindRoot, PhaseVersionListResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type PhaseVersionListResponsePhase string

const (
	PhaseVersionListResponsePhaseDDoSL4                         PhaseVersionListResponsePhase = "ddos_l4"
	PhaseVersionListResponsePhaseDDoSL7                         PhaseVersionListResponsePhase = "ddos_l7"
	PhaseVersionListResponsePhaseHTTPConfigSettings             PhaseVersionListResponsePhase = "http_config_settings"
	PhaseVersionListResponsePhaseHTTPCustomErrors               PhaseVersionListResponsePhase = "http_custom_errors"
	PhaseVersionListResponsePhaseHTTPLogCustomFields            PhaseVersionListResponsePhase = "http_log_custom_fields"
	PhaseVersionListResponsePhaseHTTPRatelimit                  PhaseVersionListResponsePhase = "http_ratelimit"
	PhaseVersionListResponsePhaseHTTPRequestCacheSettings       PhaseVersionListResponsePhase = "http_request_cache_settings"
	PhaseVersionListResponsePhaseHTTPRequestDynamicRedirect     PhaseVersionListResponsePhase = "http_request_dynamic_redirect"
	PhaseVersionListResponsePhaseHTTPRequestFirewallCustom      PhaseVersionListResponsePhase = "http_request_firewall_custom"
	PhaseVersionListResponsePhaseHTTPRequestFirewallManaged     PhaseVersionListResponsePhase = "http_request_firewall_managed"
	PhaseVersionListResponsePhaseHTTPRequestLateTransform       PhaseVersionListResponsePhase = "http_request_late_transform"
	PhaseVersionListResponsePhaseHTTPRequestOrigin              PhaseVersionListResponsePhase = "http_request_origin"
	PhaseVersionListResponsePhaseHTTPRequestRedirect            PhaseVersionListResponsePhase = "http_request_redirect"
	PhaseVersionListResponsePhaseHTTPRequestSanitize            PhaseVersionListResponsePhase = "http_request_sanitize"
	PhaseVersionListResponsePhaseHTTPRequestSbfm                PhaseVersionListResponsePhase = "http_request_sbfm"
	PhaseVersionListResponsePhaseHTTPRequestSelectConfiguration PhaseVersionListResponsePhase = "http_request_select_configuration"
	PhaseVersionListResponsePhaseHTTPRequestTransform           PhaseVersionListResponsePhase = "http_request_transform"
	PhaseVersionListResponsePhaseHTTPResponseCompression        PhaseVersionListResponsePhase = "http_response_compression"
	PhaseVersionListResponsePhaseHTTPResponseFirewallManaged    PhaseVersionListResponsePhase = "http_response_firewall_managed"
	PhaseVersionListResponsePhaseHTTPResponseHeadersTransform   PhaseVersionListResponsePhase = "http_response_headers_transform"
	PhaseVersionListResponsePhaseMagicTransit                   PhaseVersionListResponsePhase = "magic_transit"
	PhaseVersionListResponsePhaseMagicTransitIDsManaged         PhaseVersionListResponsePhase = "magic_transit_ids_managed"
	PhaseVersionListResponsePhaseMagicTransitManaged            PhaseVersionListResponsePhase = "magic_transit_managed"
)

func (r PhaseVersionListResponsePhase) IsKnown() bool {
	switch r {
	case PhaseVersionListResponsePhaseDDoSL4, PhaseVersionListResponsePhaseDDoSL7, PhaseVersionListResponsePhaseHTTPConfigSettings, PhaseVersionListResponsePhaseHTTPCustomErrors, PhaseVersionListResponsePhaseHTTPLogCustomFields, PhaseVersionListResponsePhaseHTTPRatelimit, PhaseVersionListResponsePhaseHTTPRequestCacheSettings, PhaseVersionListResponsePhaseHTTPRequestDynamicRedirect, PhaseVersionListResponsePhaseHTTPRequestFirewallCustom, PhaseVersionListResponsePhaseHTTPRequestFirewallManaged, PhaseVersionListResponsePhaseHTTPRequestLateTransform, PhaseVersionListResponsePhaseHTTPRequestOrigin, PhaseVersionListResponsePhaseHTTPRequestRedirect, PhaseVersionListResponsePhaseHTTPRequestSanitize, PhaseVersionListResponsePhaseHTTPRequestSbfm, PhaseVersionListResponsePhaseHTTPRequestSelectConfiguration, PhaseVersionListResponsePhaseHTTPRequestTransform, PhaseVersionListResponsePhaseHTTPResponseCompression, PhaseVersionListResponsePhaseHTTPResponseFirewallManaged, PhaseVersionListResponsePhaseHTTPResponseHeadersTransform, PhaseVersionListResponsePhaseMagicTransit, PhaseVersionListResponsePhaseMagicTransitIDsManaged, PhaseVersionListResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type PhaseVersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type PhaseVersionListParamsRulesetPhase string

const (
	PhaseVersionListParamsRulesetPhaseDDoSL4                         PhaseVersionListParamsRulesetPhase = "ddos_l4"
	PhaseVersionListParamsRulesetPhaseDDoSL7                         PhaseVersionListParamsRulesetPhase = "ddos_l7"
	PhaseVersionListParamsRulesetPhaseHTTPConfigSettings             PhaseVersionListParamsRulesetPhase = "http_config_settings"
	PhaseVersionListParamsRulesetPhaseHTTPCustomErrors               PhaseVersionListParamsRulesetPhase = "http_custom_errors"
	PhaseVersionListParamsRulesetPhaseHTTPLogCustomFields            PhaseVersionListParamsRulesetPhase = "http_log_custom_fields"
	PhaseVersionListParamsRulesetPhaseHTTPRatelimit                  PhaseVersionListParamsRulesetPhase = "http_ratelimit"
	PhaseVersionListParamsRulesetPhaseHTTPRequestCacheSettings       PhaseVersionListParamsRulesetPhase = "http_request_cache_settings"
	PhaseVersionListParamsRulesetPhaseHTTPRequestDynamicRedirect     PhaseVersionListParamsRulesetPhase = "http_request_dynamic_redirect"
	PhaseVersionListParamsRulesetPhaseHTTPRequestFirewallCustom      PhaseVersionListParamsRulesetPhase = "http_request_firewall_custom"
	PhaseVersionListParamsRulesetPhaseHTTPRequestFirewallManaged     PhaseVersionListParamsRulesetPhase = "http_request_firewall_managed"
	PhaseVersionListParamsRulesetPhaseHTTPRequestLateTransform       PhaseVersionListParamsRulesetPhase = "http_request_late_transform"
	PhaseVersionListParamsRulesetPhaseHTTPRequestOrigin              PhaseVersionListParamsRulesetPhase = "http_request_origin"
	PhaseVersionListParamsRulesetPhaseHTTPRequestRedirect            PhaseVersionListParamsRulesetPhase = "http_request_redirect"
	PhaseVersionListParamsRulesetPhaseHTTPRequestSanitize            PhaseVersionListParamsRulesetPhase = "http_request_sanitize"
	PhaseVersionListParamsRulesetPhaseHTTPRequestSbfm                PhaseVersionListParamsRulesetPhase = "http_request_sbfm"
	PhaseVersionListParamsRulesetPhaseHTTPRequestSelectConfiguration PhaseVersionListParamsRulesetPhase = "http_request_select_configuration"
	PhaseVersionListParamsRulesetPhaseHTTPRequestTransform           PhaseVersionListParamsRulesetPhase = "http_request_transform"
	PhaseVersionListParamsRulesetPhaseHTTPResponseCompression        PhaseVersionListParamsRulesetPhase = "http_response_compression"
	PhaseVersionListParamsRulesetPhaseHTTPResponseFirewallManaged    PhaseVersionListParamsRulesetPhase = "http_response_firewall_managed"
	PhaseVersionListParamsRulesetPhaseHTTPResponseHeadersTransform   PhaseVersionListParamsRulesetPhase = "http_response_headers_transform"
	PhaseVersionListParamsRulesetPhaseMagicTransit                   PhaseVersionListParamsRulesetPhase = "magic_transit"
	PhaseVersionListParamsRulesetPhaseMagicTransitIDsManaged         PhaseVersionListParamsRulesetPhase = "magic_transit_ids_managed"
	PhaseVersionListParamsRulesetPhaseMagicTransitManaged            PhaseVersionListParamsRulesetPhase = "magic_transit_managed"
)

func (r PhaseVersionListParamsRulesetPhase) IsKnown() bool {
	switch r {
	case PhaseVersionListParamsRulesetPhaseDDoSL4, PhaseVersionListParamsRulesetPhaseDDoSL7, PhaseVersionListParamsRulesetPhaseHTTPConfigSettings, PhaseVersionListParamsRulesetPhaseHTTPCustomErrors, PhaseVersionListParamsRulesetPhaseHTTPLogCustomFields, PhaseVersionListParamsRulesetPhaseHTTPRatelimit, PhaseVersionListParamsRulesetPhaseHTTPRequestCacheSettings, PhaseVersionListParamsRulesetPhaseHTTPRequestDynamicRedirect, PhaseVersionListParamsRulesetPhaseHTTPRequestFirewallCustom, PhaseVersionListParamsRulesetPhaseHTTPRequestFirewallManaged, PhaseVersionListParamsRulesetPhaseHTTPRequestLateTransform, PhaseVersionListParamsRulesetPhaseHTTPRequestOrigin, PhaseVersionListParamsRulesetPhaseHTTPRequestRedirect, PhaseVersionListParamsRulesetPhaseHTTPRequestSanitize, PhaseVersionListParamsRulesetPhaseHTTPRequestSbfm, PhaseVersionListParamsRulesetPhaseHTTPRequestSelectConfiguration, PhaseVersionListParamsRulesetPhaseHTTPRequestTransform, PhaseVersionListParamsRulesetPhaseHTTPResponseCompression, PhaseVersionListParamsRulesetPhaseHTTPResponseFirewallManaged, PhaseVersionListParamsRulesetPhaseHTTPResponseHeadersTransform, PhaseVersionListParamsRulesetPhaseMagicTransit, PhaseVersionListParamsRulesetPhaseMagicTransitIDsManaged, PhaseVersionListParamsRulesetPhaseMagicTransitManaged:
		return true
	}
	return false
}

type PhaseVersionGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type PhaseVersionGetParamsRulesetPhase string

const (
	PhaseVersionGetParamsRulesetPhaseDDoSL4                         PhaseVersionGetParamsRulesetPhase = "ddos_l4"
	PhaseVersionGetParamsRulesetPhaseDDoSL7                         PhaseVersionGetParamsRulesetPhase = "ddos_l7"
	PhaseVersionGetParamsRulesetPhaseHTTPConfigSettings             PhaseVersionGetParamsRulesetPhase = "http_config_settings"
	PhaseVersionGetParamsRulesetPhaseHTTPCustomErrors               PhaseVersionGetParamsRulesetPhase = "http_custom_errors"
	PhaseVersionGetParamsRulesetPhaseHTTPLogCustomFields            PhaseVersionGetParamsRulesetPhase = "http_log_custom_fields"
	PhaseVersionGetParamsRulesetPhaseHTTPRatelimit                  PhaseVersionGetParamsRulesetPhase = "http_ratelimit"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestCacheSettings       PhaseVersionGetParamsRulesetPhase = "http_request_cache_settings"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestDynamicRedirect     PhaseVersionGetParamsRulesetPhase = "http_request_dynamic_redirect"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallCustom      PhaseVersionGetParamsRulesetPhase = "http_request_firewall_custom"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallManaged     PhaseVersionGetParamsRulesetPhase = "http_request_firewall_managed"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestLateTransform       PhaseVersionGetParamsRulesetPhase = "http_request_late_transform"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestOrigin              PhaseVersionGetParamsRulesetPhase = "http_request_origin"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestRedirect            PhaseVersionGetParamsRulesetPhase = "http_request_redirect"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestSanitize            PhaseVersionGetParamsRulesetPhase = "http_request_sanitize"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestSbfm                PhaseVersionGetParamsRulesetPhase = "http_request_sbfm"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestSelectConfiguration PhaseVersionGetParamsRulesetPhase = "http_request_select_configuration"
	PhaseVersionGetParamsRulesetPhaseHTTPRequestTransform           PhaseVersionGetParamsRulesetPhase = "http_request_transform"
	PhaseVersionGetParamsRulesetPhaseHTTPResponseCompression        PhaseVersionGetParamsRulesetPhase = "http_response_compression"
	PhaseVersionGetParamsRulesetPhaseHTTPResponseFirewallManaged    PhaseVersionGetParamsRulesetPhase = "http_response_firewall_managed"
	PhaseVersionGetParamsRulesetPhaseHTTPResponseHeadersTransform   PhaseVersionGetParamsRulesetPhase = "http_response_headers_transform"
	PhaseVersionGetParamsRulesetPhaseMagicTransit                   PhaseVersionGetParamsRulesetPhase = "magic_transit"
	PhaseVersionGetParamsRulesetPhaseMagicTransitIDsManaged         PhaseVersionGetParamsRulesetPhase = "magic_transit_ids_managed"
	PhaseVersionGetParamsRulesetPhaseMagicTransitManaged            PhaseVersionGetParamsRulesetPhase = "magic_transit_managed"
)

func (r PhaseVersionGetParamsRulesetPhase) IsKnown() bool {
	switch r {
	case PhaseVersionGetParamsRulesetPhaseDDoSL4, PhaseVersionGetParamsRulesetPhaseDDoSL7, PhaseVersionGetParamsRulesetPhaseHTTPConfigSettings, PhaseVersionGetParamsRulesetPhaseHTTPCustomErrors, PhaseVersionGetParamsRulesetPhaseHTTPLogCustomFields, PhaseVersionGetParamsRulesetPhaseHTTPRatelimit, PhaseVersionGetParamsRulesetPhaseHTTPRequestCacheSettings, PhaseVersionGetParamsRulesetPhaseHTTPRequestDynamicRedirect, PhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallCustom, PhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallManaged, PhaseVersionGetParamsRulesetPhaseHTTPRequestLateTransform, PhaseVersionGetParamsRulesetPhaseHTTPRequestOrigin, PhaseVersionGetParamsRulesetPhaseHTTPRequestRedirect, PhaseVersionGetParamsRulesetPhaseHTTPRequestSanitize, PhaseVersionGetParamsRulesetPhaseHTTPRequestSbfm, PhaseVersionGetParamsRulesetPhaseHTTPRequestSelectConfiguration, PhaseVersionGetParamsRulesetPhaseHTTPRequestTransform, PhaseVersionGetParamsRulesetPhaseHTTPResponseCompression, PhaseVersionGetParamsRulesetPhaseHTTPResponseFirewallManaged, PhaseVersionGetParamsRulesetPhaseHTTPResponseHeadersTransform, PhaseVersionGetParamsRulesetPhaseMagicTransit, PhaseVersionGetParamsRulesetPhaseMagicTransitIDsManaged, PhaseVersionGetParamsRulesetPhaseMagicTransitManaged:
		return true
	}
	return false
}

// A response object.
type PhaseVersionGetResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseVersionGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result Ruleset `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    phaseVersionGetResponseEnvelopeJSON    `json:"-"`
}

// phaseVersionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PhaseVersionGetResponseEnvelope]
type phaseVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseVersionGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseVersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   phaseVersionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// phaseVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseEnvelopeErrors]
type phaseVersionGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseVersionGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    phaseVersionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// phaseVersionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseEnvelopeErrorsSource]
type phaseVersionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseVersionGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseVersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   phaseVersionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// phaseVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseEnvelopeMessages]
type phaseVersionGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseVersionGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    phaseVersionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// phaseVersionGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseEnvelopeMessagesSource]
type phaseVersionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PhaseVersionGetResponseEnvelopeSuccess bool

const (
	PhaseVersionGetResponseEnvelopeSuccessTrue PhaseVersionGetResponseEnvelopeSuccess = true
)

func (r PhaseVersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
