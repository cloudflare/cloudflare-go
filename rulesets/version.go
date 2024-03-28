// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// VersionService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewVersionService] method instead.
type VersionService struct {
	Options []option.RequestOption
	ByTag   *VersionByTagService
}

// NewVersionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVersionService(opts ...option.RequestOption) (r *VersionService) {
	r = &VersionService{}
	r.Options = opts
	r.ByTag = NewVersionByTagService(opts...)
	return
}

// Fetches the versions of an account or zone ruleset.
func (r *VersionService) List(ctx context.Context, rulesetID string, query VersionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[VersionListResponse], err error) {
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
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions", accountOrZone, accountOrZoneID, rulesetID)
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

// Fetches the versions of an account or zone ruleset.
func (r *VersionService) ListAutoPaging(ctx context.Context, rulesetID string, query VersionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[VersionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, rulesetID, query, opts...))
}

// Deletes an existing version of an account or zone ruleset.
func (r *VersionService) Delete(ctx context.Context, rulesetID string, rulesetVersion string, body VersionDeleteParams, opts ...option.RequestOption) (err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches a specific version of an account or zone ruleset.
func (r *VersionService) Get(ctx context.Context, rulesetID string, rulesetVersion string, query VersionGetParams, opts ...option.RequestOption) (res *Ruleset, err error) {
	opts = append(r.Options[:], opts...)
	var env VersionGetResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type VersionListResponse struct {
	// The kind of the ruleset.
	Kind VersionListResponseKind `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase VersionListResponsePhase `json:"phase,required"`
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the ruleset.
	Version string                  `json:"version"`
	JSON    versionListResponseJSON `json:"-"`
}

// versionListResponseJSON contains the JSON metadata for the struct
// [VersionListResponse]
type versionListResponseJSON struct {
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

func (r *VersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type VersionListResponseKind string

const (
	VersionListResponseKindManaged VersionListResponseKind = "managed"
	VersionListResponseKindCustom  VersionListResponseKind = "custom"
	VersionListResponseKindRoot    VersionListResponseKind = "root"
	VersionListResponseKindZone    VersionListResponseKind = "zone"
)

func (r VersionListResponseKind) IsKnown() bool {
	switch r {
	case VersionListResponseKindManaged, VersionListResponseKindCustom, VersionListResponseKindRoot, VersionListResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type VersionListResponsePhase string

const (
	VersionListResponsePhaseDDoSL4                         VersionListResponsePhase = "ddos_l4"
	VersionListResponsePhaseDDoSL7                         VersionListResponsePhase = "ddos_l7"
	VersionListResponsePhaseHTTPConfigSettings             VersionListResponsePhase = "http_config_settings"
	VersionListResponsePhaseHTTPCustomErrors               VersionListResponsePhase = "http_custom_errors"
	VersionListResponsePhaseHTTPLogCustomFields            VersionListResponsePhase = "http_log_custom_fields"
	VersionListResponsePhaseHTTPRatelimit                  VersionListResponsePhase = "http_ratelimit"
	VersionListResponsePhaseHTTPRequestCacheSettings       VersionListResponsePhase = "http_request_cache_settings"
	VersionListResponsePhaseHTTPRequestDynamicRedirect     VersionListResponsePhase = "http_request_dynamic_redirect"
	VersionListResponsePhaseHTTPRequestFirewallCustom      VersionListResponsePhase = "http_request_firewall_custom"
	VersionListResponsePhaseHTTPRequestFirewallManaged     VersionListResponsePhase = "http_request_firewall_managed"
	VersionListResponsePhaseHTTPRequestLateTransform       VersionListResponsePhase = "http_request_late_transform"
	VersionListResponsePhaseHTTPRequestOrigin              VersionListResponsePhase = "http_request_origin"
	VersionListResponsePhaseHTTPRequestRedirect            VersionListResponsePhase = "http_request_redirect"
	VersionListResponsePhaseHTTPRequestSanitize            VersionListResponsePhase = "http_request_sanitize"
	VersionListResponsePhaseHTTPRequestSbfm                VersionListResponsePhase = "http_request_sbfm"
	VersionListResponsePhaseHTTPRequestSelectConfiguration VersionListResponsePhase = "http_request_select_configuration"
	VersionListResponsePhaseHTTPRequestTransform           VersionListResponsePhase = "http_request_transform"
	VersionListResponsePhaseHTTPResponseCompression        VersionListResponsePhase = "http_response_compression"
	VersionListResponsePhaseHTTPResponseFirewallManaged    VersionListResponsePhase = "http_response_firewall_managed"
	VersionListResponsePhaseHTTPResponseHeadersTransform   VersionListResponsePhase = "http_response_headers_transform"
	VersionListResponsePhaseMagicTransit                   VersionListResponsePhase = "magic_transit"
	VersionListResponsePhaseMagicTransitIDsManaged         VersionListResponsePhase = "magic_transit_ids_managed"
	VersionListResponsePhaseMagicTransitManaged            VersionListResponsePhase = "magic_transit_managed"
)

func (r VersionListResponsePhase) IsKnown() bool {
	switch r {
	case VersionListResponsePhaseDDoSL4, VersionListResponsePhaseDDoSL7, VersionListResponsePhaseHTTPConfigSettings, VersionListResponsePhaseHTTPCustomErrors, VersionListResponsePhaseHTTPLogCustomFields, VersionListResponsePhaseHTTPRatelimit, VersionListResponsePhaseHTTPRequestCacheSettings, VersionListResponsePhaseHTTPRequestDynamicRedirect, VersionListResponsePhaseHTTPRequestFirewallCustom, VersionListResponsePhaseHTTPRequestFirewallManaged, VersionListResponsePhaseHTTPRequestLateTransform, VersionListResponsePhaseHTTPRequestOrigin, VersionListResponsePhaseHTTPRequestRedirect, VersionListResponsePhaseHTTPRequestSanitize, VersionListResponsePhaseHTTPRequestSbfm, VersionListResponsePhaseHTTPRequestSelectConfiguration, VersionListResponsePhaseHTTPRequestTransform, VersionListResponsePhaseHTTPResponseCompression, VersionListResponsePhaseHTTPResponseFirewallManaged, VersionListResponsePhaseHTTPResponseHeadersTransform, VersionListResponsePhaseMagicTransit, VersionListResponsePhaseMagicTransitIDsManaged, VersionListResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type VersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type VersionDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type VersionGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type VersionGetResponseEnvelope struct {
	// A list of error messages.
	Errors []VersionGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []VersionGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result Ruleset `json:"result,required"`
	// Whether the API call was successful.
	Success VersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    versionGetResponseEnvelopeJSON    `json:"-"`
}

// versionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelope]
type versionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   versionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// versionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelopeErrors]
type versionGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    versionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// versionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [VersionGetResponseEnvelopeErrorsSource]
type versionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   versionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// versionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelopeMessages]
type versionGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    versionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// versionGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [VersionGetResponseEnvelopeMessagesSource]
type versionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VersionGetResponseEnvelopeSuccess bool

const (
	VersionGetResponseEnvelopeSuccessTrue VersionGetResponseEnvelopeSuccess = true
)

func (r VersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
