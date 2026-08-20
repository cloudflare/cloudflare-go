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
// `default_mode` sets the zone-level enforcement mode. `enforcement_rules` is the
// ordered list of rules that override enforcement for matching requests.
//
// This is a partial update: only the fields present in the request body are
// changed.
//
//   - Sending an empty array (`[]`) clears all enforcement rules.
//   - At least one of `default_mode` or `enforcement_rules` must be present; an
//     empty body (`{}`) is rejected with `400`.
//   - Rule `id` is read-only (assigned by Cloudflare) and ignored on input.
//   - Rule `mode` must be `min-friction` or `max-security` (`off` is not a valid
//     rule mode; use `default_mode` to disable enforcement).
func (r *PrecursorService) Update(ctx context.Context, params PrecursorUpdateParams, opts ...option.RequestOption) (res *PrecursorConfig, err error) {
	var env PrecursorUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
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

// Retrieve a zone's Precursor configuration: the zone-level `default_mode` and the
// ordered list of `enforcement_rules`.
func (r *PrecursorService) Get(ctx context.Context, query PrecursorGetParams, opts ...option.RequestOption) (res *PrecursorConfig, err error) {
	var env PrecursorGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/precursor", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type EnforcementRule struct {
	// The filter expression that determines which requests the rule matches.
	Expression string `json:"expression" api:"required"`
	// The override mode Precursor applies to requests matching an enforcement rule.
	// Unlike `default_mode`, this cannot be `off`.
	Mode EnforcementRuleMode `json:"mode" api:"required"`
	// The read-only identifier that Cloudflare assigns to the rule.
	ID string `json:"id"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule is active.
	Enabled bool                `json:"enabled"`
	JSON    enforcementRuleJSON `json:"-"`
}

// enforcementRuleJSON contains the JSON metadata for the struct [EnforcementRule]
type enforcementRuleJSON struct {
	Expression  apijson.Field
	Mode        apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnforcementRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enforcementRuleJSON) RawJSON() string {
	return r.raw
}

// The override mode Precursor applies to requests matching an enforcement rule.
// Unlike `default_mode`, this cannot be `off`.
type EnforcementRuleMode string

const (
	EnforcementRuleModeMinFriction EnforcementRuleMode = "min-friction"
	EnforcementRuleModeMaxSecurity EnforcementRuleMode = "max-security"
)

func (r EnforcementRuleMode) IsKnown() bool {
	switch r {
	case EnforcementRuleModeMinFriction, EnforcementRuleModeMaxSecurity:
		return true
	}
	return false
}

type EnforcementRuleParam struct {
	// The filter expression that determines which requests the rule matches.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The override mode Precursor applies to requests matching an enforcement rule.
	// Unlike `default_mode`, this cannot be `off`.
	Mode param.Field[EnforcementRuleMode] `json:"mode" api:"required"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule is active.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r EnforcementRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrecursorConfig struct {
	// The zone-level Precursor enforcement mode applied to requests that do not match
	// a more specific enforcement rule.
	DefaultMode PrecursorConfigDefaultMode `json:"default_mode"`
	// The ordered list of enforcement rules for the zone.
	EnforcementRules []EnforcementRule   `json:"enforcement_rules"`
	JSON             precursorConfigJSON `json:"-"`
}

// precursorConfigJSON contains the JSON metadata for the struct [PrecursorConfig]
type precursorConfigJSON struct {
	DefaultMode      apijson.Field
	EnforcementRules apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrecursorConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r precursorConfigJSON) RawJSON() string {
	return r.raw
}

// The zone-level Precursor enforcement mode applied to requests that do not match
// a more specific enforcement rule.
type PrecursorConfigDefaultMode string

const (
	PrecursorConfigDefaultModeOff         PrecursorConfigDefaultMode = "off"
	PrecursorConfigDefaultModeMinFriction PrecursorConfigDefaultMode = "min-friction"
	PrecursorConfigDefaultModeMaxSecurity PrecursorConfigDefaultMode = "max-security"
)

func (r PrecursorConfigDefaultMode) IsKnown() bool {
	switch r {
	case PrecursorConfigDefaultModeOff, PrecursorConfigDefaultModeMinFriction, PrecursorConfigDefaultModeMaxSecurity:
		return true
	}
	return false
}

type PrecursorConfigParam struct {
	// The zone-level Precursor enforcement mode applied to requests that do not match
	// a more specific enforcement rule.
	DefaultMode param.Field[PrecursorConfigDefaultMode] `json:"default_mode"`
	// The ordered list of enforcement rules for the zone.
	EnforcementRules param.Field[[]EnforcementRuleParam] `json:"enforcement_rules"`
}

func (r PrecursorConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrecursorUpdateParams struct {
	// Identifier.
	ZoneID          param.Field[string]  `path:"zone_id" api:"required"`
	PrecursorConfig PrecursorConfigParam `json:"precursor_config" api:"required"`
}

func (r PrecursorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PrecursorConfig)
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
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
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
