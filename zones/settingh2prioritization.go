// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingH2PrioritizationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingH2PrioritizationService] method instead.
type SettingH2PrioritizationService struct {
	Options []option.RequestOption
}

// NewSettingH2PrioritizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingH2PrioritizationService(opts ...option.RequestOption) (r *SettingH2PrioritizationService) {
	r = &SettingH2PrioritizationService{}
	r.Options = opts
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *SettingH2PrioritizationService) Edit(ctx context.Context, params SettingH2PrioritizationEditParams, opts ...option.RequestOption) (res *H2Prioritization, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingH2PrioritizationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *SettingH2PrioritizationService) Get(ctx context.Context, query SettingH2PrioritizationGetParams, opts ...option.RequestOption) (res *H2Prioritization, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingH2PrioritizationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type H2Prioritization struct {
	// ID of the zone setting.
	ID H2PrioritizationID `json:"id,required"`
	// Current value of the zone setting.
	Value H2PrioritizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable H2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time            `json:"modified_on,nullable" format:"date-time"`
	JSON       h2PrioritizationJSON `json:"-"`
}

// h2PrioritizationJSON contains the JSON metadata for the struct
// [H2Prioritization]
type h2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *H2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r h2PrioritizationJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type H2PrioritizationID string

const (
	H2PrioritizationIDH2Prioritization H2PrioritizationID = "h2_prioritization"
)

func (r H2PrioritizationID) IsKnown() bool {
	switch r {
	case H2PrioritizationIDH2Prioritization:
		return true
	}
	return false
}

// Current value of the zone setting.
type H2PrioritizationValue string

const (
	H2PrioritizationValueOn     H2PrioritizationValue = "on"
	H2PrioritizationValueOff    H2PrioritizationValue = "off"
	H2PrioritizationValueCustom H2PrioritizationValue = "custom"
)

func (r H2PrioritizationValue) IsKnown() bool {
	switch r {
	case H2PrioritizationValueOn, H2PrioritizationValueOff, H2PrioritizationValueCustom:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type H2PrioritizationEditable bool

const (
	H2PrioritizationEditableTrue  H2PrioritizationEditable = true
	H2PrioritizationEditableFalse H2PrioritizationEditable = false
)

func (r H2PrioritizationEditable) IsKnown() bool {
	switch r {
	case H2PrioritizationEditableTrue, H2PrioritizationEditableFalse:
		return true
	}
	return false
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type H2PrioritizationParam struct {
	// ID of the zone setting.
	ID param.Field[H2PrioritizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[H2PrioritizationValue] `json:"value,required"`
}

func (r H2PrioritizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingH2PrioritizationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Value param.Field[H2PrioritizationParam] `json:"value,required"`
}

func (r SettingH2PrioritizationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingH2PrioritizationEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result H2Prioritization                                `json:"result"`
	JSON   settingH2PrioritizationEditResponseEnvelopeJSON `json:"-"`
}

// settingH2PrioritizationEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingH2PrioritizationEditResponseEnvelope]
type settingH2PrioritizationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingH2PrioritizationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingH2PrioritizationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingH2PrioritizationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result H2Prioritization                               `json:"result"`
	JSON   settingH2PrioritizationGetResponseEnvelopeJSON `json:"-"`
}

// settingH2PrioritizationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingH2PrioritizationGetResponseEnvelope]
type settingH2PrioritizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingH2PrioritizationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
