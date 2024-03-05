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

// RulesetVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRulesetVersionService] method
// instead.
type RulesetVersionService struct {
	Options []option.RequestOption
	ByTags  *RulesetVersionByTagService
}

// NewRulesetVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRulesetVersionService(opts ...option.RequestOption) (r *RulesetVersionService) {
	r = &RulesetVersionService{}
	r.Options = opts
	r.ByTags = NewRulesetVersionByTagService(opts...)
	return
}

// Fetches the versions of an account or zone ruleset.
func (r *RulesetVersionService) List(ctx context.Context, rulesetID string, query RulesetVersionListParams, opts ...option.RequestOption) (res *RulesetsRulesetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetVersionListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing version of an account or zone ruleset.
func (r *RulesetVersionService) Delete(ctx context.Context, rulesetID string, rulesetVersion string, body RulesetVersionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Fetches a specific version of an account or zone ruleset.
func (r *RulesetVersionService) Get(ctx context.Context, rulesetID string, rulesetVersion string, query RulesetVersionGetParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetVersionGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type RulesetVersionListResponse struct {
	// The kind of the ruleset.
	Kind RulesetVersionListResponseKind `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetVersionListResponsePhase `json:"phase,required"`
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the ruleset.
	Version string                         `json:"version"`
	JSON    rulesetVersionListResponseJSON `json:"-"`
}

// rulesetVersionListResponseJSON contains the JSON metadata for the struct
// [RulesetVersionListResponse]
type rulesetVersionListResponseJSON struct {
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

func (r *RulesetVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetVersionListResponseKind string

const (
	RulesetVersionListResponseKindManaged RulesetVersionListResponseKind = "managed"
	RulesetVersionListResponseKindCustom  RulesetVersionListResponseKind = "custom"
	RulesetVersionListResponseKindRoot    RulesetVersionListResponseKind = "root"
	RulesetVersionListResponseKindZone    RulesetVersionListResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetVersionListResponsePhase string

const (
	RulesetVersionListResponsePhaseDDOSL4                         RulesetVersionListResponsePhase = "ddos_l4"
	RulesetVersionListResponsePhaseDDOSL7                         RulesetVersionListResponsePhase = "ddos_l7"
	RulesetVersionListResponsePhaseHTTPConfigSettings             RulesetVersionListResponsePhase = "http_config_settings"
	RulesetVersionListResponsePhaseHTTPCustomErrors               RulesetVersionListResponsePhase = "http_custom_errors"
	RulesetVersionListResponsePhaseHTTPLogCustomFields            RulesetVersionListResponsePhase = "http_log_custom_fields"
	RulesetVersionListResponsePhaseHTTPRatelimit                  RulesetVersionListResponsePhase = "http_ratelimit"
	RulesetVersionListResponsePhaseHTTPRequestCacheSettings       RulesetVersionListResponsePhase = "http_request_cache_settings"
	RulesetVersionListResponsePhaseHTTPRequestDynamicRedirect     RulesetVersionListResponsePhase = "http_request_dynamic_redirect"
	RulesetVersionListResponsePhaseHTTPRequestFirewallCustom      RulesetVersionListResponsePhase = "http_request_firewall_custom"
	RulesetVersionListResponsePhaseHTTPRequestFirewallManaged     RulesetVersionListResponsePhase = "http_request_firewall_managed"
	RulesetVersionListResponsePhaseHTTPRequestLateTransform       RulesetVersionListResponsePhase = "http_request_late_transform"
	RulesetVersionListResponsePhaseHTTPRequestOrigin              RulesetVersionListResponsePhase = "http_request_origin"
	RulesetVersionListResponsePhaseHTTPRequestRedirect            RulesetVersionListResponsePhase = "http_request_redirect"
	RulesetVersionListResponsePhaseHTTPRequestSanitize            RulesetVersionListResponsePhase = "http_request_sanitize"
	RulesetVersionListResponsePhaseHTTPRequestSbfm                RulesetVersionListResponsePhase = "http_request_sbfm"
	RulesetVersionListResponsePhaseHTTPRequestSelectConfiguration RulesetVersionListResponsePhase = "http_request_select_configuration"
	RulesetVersionListResponsePhaseHTTPRequestTransform           RulesetVersionListResponsePhase = "http_request_transform"
	RulesetVersionListResponsePhaseHTTPResponseCompression        RulesetVersionListResponsePhase = "http_response_compression"
	RulesetVersionListResponsePhaseHTTPResponseFirewallManaged    RulesetVersionListResponsePhase = "http_response_firewall_managed"
	RulesetVersionListResponsePhaseHTTPResponseHeadersTransform   RulesetVersionListResponsePhase = "http_response_headers_transform"
	RulesetVersionListResponsePhaseMagicTransit                   RulesetVersionListResponsePhase = "magic_transit"
	RulesetVersionListResponsePhaseMagicTransitIDsManaged         RulesetVersionListResponsePhase = "magic_transit_ids_managed"
	RulesetVersionListResponsePhaseMagicTransitManaged            RulesetVersionListResponsePhase = "magic_transit_managed"
)

type RulesetVersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type RulesetVersionListResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetVersionListResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetVersionListResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetsResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetVersionListResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetVersionListResponseEnvelopeJSON    `json:"-"`
}

// rulesetVersionListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetVersionListResponseEnvelope]
type rulesetVersionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetVersionListResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionListResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetVersionListResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetVersionListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetVersionListResponseEnvelopeErrors]
type rulesetVersionListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetVersionListResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                             `json:"pointer,required"`
	JSON    rulesetVersionListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetVersionListResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RulesetVersionListResponseEnvelopeErrorsSource]
type rulesetVersionListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetVersionListResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionListResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetVersionListResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetVersionListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RulesetVersionListResponseEnvelopeMessages]
type rulesetVersionListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetVersionListResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                               `json:"pointer,required"`
	JSON    rulesetVersionListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetVersionListResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RulesetVersionListResponseEnvelopeMessagesSource]
type rulesetVersionListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetVersionListResponseEnvelopeSuccess bool

const (
	RulesetVersionListResponseEnvelopeSuccessTrue RulesetVersionListResponseEnvelopeSuccess = true
)

type RulesetVersionDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type RulesetVersionGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type RulesetVersionGetResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetVersionGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetVersionGetResponseEnvelopeJSON    `json:"-"`
}

// rulesetVersionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetVersionGetResponseEnvelope]
type rulesetVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetVersionGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetVersionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetVersionGetResponseEnvelopeErrors]
type rulesetVersionGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetVersionGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    rulesetVersionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetVersionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetVersionGetResponseEnvelopeErrorsSource]
type rulesetVersionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetVersionGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetVersionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetVersionGetResponseEnvelopeMessages]
type rulesetVersionGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetVersionGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                              `json:"pointer,required"`
	JSON    rulesetVersionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetVersionGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RulesetVersionGetResponseEnvelopeMessagesSource]
type rulesetVersionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetVersionGetResponseEnvelopeSuccess bool

const (
	RulesetVersionGetResponseEnvelopeSuccessTrue RulesetVersionGetResponseEnvelopeSuccess = true
)
