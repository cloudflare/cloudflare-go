// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package precursor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// PrecursorService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrecursorService] method instead.
type PrecursorService struct {
	Options []option.RequestOption
}

// NewPrecursorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrecursorService(opts ...option.RequestOption) (r *PrecursorService) {
	r = &PrecursorService{}
	r.Options = opts
	return
}

// Updates the Precursor configuration for a zone.
//
// `mode` controls whether Precursor detection is off, monitoring, or armed.
// `show_interstitial` controls the zone-level response when armed. `rules` is the
// ordered list of overrides for matching requests.
//
// This is a partial update: only the fields present in the request body are
// changed.
//
//   - Sending an empty array (`[]`) clears all rules.
//   - At least one of `mode`, `show_interstitial`, or `rules` must be present; an
//     empty body (`{}`) is rejected with `400`.
//   - Rule `id` is read-only (assigned by Cloudflare) and ignored on input.
//   - Rule `show_interstitial` is required and controls the response for requests
//     that match that rule.
//   - Rule `expression` is limited to 4000 characters. The limit applies to each
//     rule individually, not to the combined size of all rules.
func (r *PrecursorService) Update(ctx context.Context, params PrecursorUpdateParams, opts ...option.RequestOption) (res *PrecursorConfig, err error) {
	var env PrecursorUpdateResponseEnvelope
	if params.CloudflareVersion.Present {
		opts = append(opts, option.WithHeader("Cloudflare-Version", fmt.Sprintf("%v", params.CloudflareVersion)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Cloudflare-Version", "1.0")}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/precursor", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve a zone's Precursor configuration: the detection `mode`, whether to
// `show_interstitial`, and the ordered list of `rules`.
func (r *PrecursorService) Get(ctx context.Context, params PrecursorGetParams, opts ...option.RequestOption) (res *PrecursorConfig, err error) {
	var env PrecursorGetResponseEnvelope
	if params.CloudflareVersion.Present {
		opts = append(opts, option.WithHeader("Cloudflare-Version", fmt.Sprintf("%v", params.CloudflareVersion)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Cloudflare-Version", "1.0")}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/precursor", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type PrecursorConfig struct {
	// Controls Precursor detection for the zone. `off` disables detection, `monitor`
	// records detections without affecting bot scores, and `arm` allows detections to
	// affect bot scores.
	Mode PrecursorConfigMode `json:"mode"`
	// The ordered list of Precursor rules for the zone.
	Rules []PrecursorConfigRule `json:"rules"`
	// Whether Precursor shows an interstitial when it takes action on a request.
	ShowInterstitial bool                `json:"show_interstitial"`
	JSON             precursorConfigJSON `json:"-"`
}

// precursorConfigJSON contains the JSON metadata for the struct [PrecursorConfig]
type precursorConfigJSON struct {
	Mode             apijson.Field
	Rules            apijson.Field
	ShowInterstitial apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrecursorConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorConfigJSON) RawJSON() string {
	return r.raw
}

// Controls Precursor detection for the zone. `off` disables detection, `monitor`
// records detections without affecting bot scores, and `arm` allows detections to
// affect bot scores.
type PrecursorConfigMode string

const (
	PrecursorConfigModeOff     PrecursorConfigMode = "off"
	PrecursorConfigModeMonitor PrecursorConfigMode = "monitor"
	PrecursorConfigModeArm     PrecursorConfigMode = "arm"
)

func (r PrecursorConfigMode) IsKnown() bool {
	switch r {
	case PrecursorConfigModeOff, PrecursorConfigModeMonitor, PrecursorConfigModeArm:
		return true
	}
	return false
}

type PrecursorConfigRule struct {
	// The filter expression that determines which requests the rule matches.
	Expression string `json:"expression" api:"required"`
	// Whether Precursor shows an interstitial when it takes action on a request.
	ShowInterstitial bool `json:"show_interstitial" api:"required"`
	// The read-only identifier that Cloudflare assigns to the rule.
	ID string `json:"id"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule is active.
	Enabled bool                    `json:"enabled"`
	JSON    precursorConfigRuleJSON `json:"-"`
}

// precursorConfigRuleJSON contains the JSON metadata for the struct
// [PrecursorConfigRule]
type precursorConfigRuleJSON struct {
	Expression       apijson.Field
	ShowInterstitial apijson.Field
	ID               apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrecursorConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorConfigRuleJSON) RawJSON() string {
	return r.raw
}

type PrecursorConfigParam struct {
	// Controls Precursor detection for the zone. `off` disables detection, `monitor`
	// records detections without affecting bot scores, and `arm` allows detections to
	// affect bot scores.
	Mode param.Field[PrecursorConfigMode] `json:"mode"`
	// The ordered list of Precursor rules for the zone.
	Rules param.Field[[]PrecursorConfigRuleParam] `json:"rules"`
	// Whether Precursor shows an interstitial when it takes action on a request.
	ShowInterstitial param.Field[bool] `json:"show_interstitial"`
}

func (r PrecursorConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrecursorConfigRuleParam struct {
	// The filter expression that determines which requests the rule matches.
	Expression param.Field[string] `json:"expression" api:"required"`
	// Whether Precursor shows an interstitial when it takes action on a request.
	ShowInterstitial param.Field[bool] `json:"show_interstitial" api:"required"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule is active.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r PrecursorConfigRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrecursorUpdateParams struct {
	// Identifier.
	ZoneID            param.Field[string]                                 `path:"zone_id" api:"required"`
	PrecursorConfig   PrecursorConfigParam                                `json:"precursor_config" api:"required"`
	CloudflareVersion param.Field[PrecursorUpdateParamsCloudflareVersion] `header:"Cloudflare-Version" api:"required"`
}

func (r PrecursorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrecursorConfig)
}

type PrecursorUpdateParamsCloudflareVersion string

const (
	PrecursorUpdateParamsCloudflareVersion1_0 PrecursorUpdateParamsCloudflareVersion = "1.0"
)

func (r PrecursorUpdateParamsCloudflareVersion) IsKnown() bool {
	switch r {
	case PrecursorUpdateParamsCloudflareVersion1_0:
		return true
	}
	return false
}

type PrecursorUpdateResponseEnvelope struct {
	Errors   []PrecursorUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []PrecursorUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success PrecursorUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  PrecursorConfig                        `json:"result"`
	JSON    precursorUpdateResponseEnvelopeJSON    `json:"-"`
}

// precursorUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrecursorUpdateResponseEnvelope]
type precursorUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrecursorUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrecursorUpdateResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           PrecursorUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             precursorUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// precursorUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrecursorUpdateResponseEnvelopeErrors]
type precursorUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrecursorUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrecursorUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    precursorUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// precursorUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [PrecursorUpdateResponseEnvelopeErrorsSource]
type precursorUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrecursorUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type PrecursorUpdateResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           PrecursorUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             precursorUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// precursorUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PrecursorUpdateResponseEnvelopeMessages]
type precursorUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrecursorUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PrecursorUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    precursorUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// precursorUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [PrecursorUpdateResponseEnvelopeMessagesSource]
type precursorUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrecursorUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PrecursorUpdateResponseEnvelopeSuccess bool

const (
	PrecursorUpdateResponseEnvelopeSuccessTrue PrecursorUpdateResponseEnvelopeSuccess = true
)

func (r PrecursorUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrecursorUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrecursorGetParams struct {
	// Identifier.
	ZoneID            param.Field[string]                              `path:"zone_id" api:"required"`
	CloudflareVersion param.Field[PrecursorGetParamsCloudflareVersion] `header:"Cloudflare-Version" api:"required"`
}

type PrecursorGetParamsCloudflareVersion string

const (
	PrecursorGetParamsCloudflareVersion1_0 PrecursorGetParamsCloudflareVersion = "1.0"
)

func (r PrecursorGetParamsCloudflareVersion) IsKnown() bool {
	switch r {
	case PrecursorGetParamsCloudflareVersion1_0:
		return true
	}
	return false
}

type PrecursorGetResponseEnvelope struct {
	Errors   []PrecursorGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []PrecursorGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success PrecursorGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  PrecursorConfig                     `json:"result"`
	JSON    precursorGetResponseEnvelopeJSON    `json:"-"`
}

// precursorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrecursorGetResponseEnvelope]
type precursorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrecursorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrecursorGetResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           PrecursorGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             precursorGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// precursorGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PrecursorGetResponseEnvelopeErrors]
type precursorGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrecursorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrecursorGetResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    precursorGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// precursorGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [PrecursorGetResponseEnvelopeErrorsSource]
type precursorGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrecursorGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type PrecursorGetResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           PrecursorGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             precursorGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// precursorGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PrecursorGetResponseEnvelopeMessages]
type precursorGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrecursorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PrecursorGetResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    precursorGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// precursorGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [PrecursorGetResponseEnvelopeMessagesSource]
type precursorGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrecursorGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PrecursorGetResponseEnvelopeSuccess bool

const (
	PrecursorGetResponseEnvelopeSuccessTrue PrecursorGetResponseEnvelopeSuccess = true
)

func (r PrecursorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrecursorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
