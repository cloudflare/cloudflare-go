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
)

// SettingH2PrioritizationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSettingH2PrioritizationService] method instead.
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
func (r *SettingH2PrioritizationService) Edit(ctx context.Context, params SettingH2PrioritizationEditParams, opts ...option.RequestOption) (res *ZoneSettingH2Prioritization, err error) {
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
func (r *SettingH2PrioritizationService) Get(ctx context.Context, query SettingH2PrioritizationGetParams, opts ...option.RequestOption) (res *ZoneSettingH2Prioritization, err error) {
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
type ZoneSettingH2Prioritization struct {
	// ID of the zone setting.
	ID ZoneSettingH2PrioritizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingH2PrioritizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingH2PrioritizationJSON `json:"-"`
}

// zoneSettingH2PrioritizationJSON contains the JSON metadata for the struct
// [ZoneSettingH2Prioritization]
type zoneSettingH2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingH2PrioritizationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingH2Prioritization) implementsZonesSettingEditResponse() {}

func (r ZoneSettingH2Prioritization) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingH2PrioritizationID string

const (
	ZoneSettingH2PrioritizationIDH2Prioritization ZoneSettingH2PrioritizationID = "h2_prioritization"
)

func (r ZoneSettingH2PrioritizationID) IsKnown() bool {
	switch r {
	case ZoneSettingH2PrioritizationIDH2Prioritization:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingH2PrioritizationValue string

const (
	ZoneSettingH2PrioritizationValueOn     ZoneSettingH2PrioritizationValue = "on"
	ZoneSettingH2PrioritizationValueOff    ZoneSettingH2PrioritizationValue = "off"
	ZoneSettingH2PrioritizationValueCustom ZoneSettingH2PrioritizationValue = "custom"
)

func (r ZoneSettingH2PrioritizationValue) IsKnown() bool {
	switch r {
	case ZoneSettingH2PrioritizationValueOn, ZoneSettingH2PrioritizationValueOff, ZoneSettingH2PrioritizationValueCustom:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingH2PrioritizationEditable bool

const (
	ZoneSettingH2PrioritizationEditableTrue  ZoneSettingH2PrioritizationEditable = true
	ZoneSettingH2PrioritizationEditableFalse ZoneSettingH2PrioritizationEditable = false
)

func (r ZoneSettingH2PrioritizationEditable) IsKnown() bool {
	switch r {
	case ZoneSettingH2PrioritizationEditableTrue, ZoneSettingH2PrioritizationEditableFalse:
		return true
	}
	return false
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingH2PrioritizationParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingH2PrioritizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingH2PrioritizationValue] `json:"value,required"`
}

func (r ZoneSettingH2PrioritizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingH2PrioritizationParam) implementsZonesSettingEditParamsItem() {}

type SettingH2PrioritizationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Value param.Field[ZoneSettingH2PrioritizationParam] `json:"value,required"`
}

func (r SettingH2PrioritizationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingH2PrioritizationEditResponseEnvelope struct {
	Errors   []SettingH2PrioritizationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingH2PrioritizationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result ZoneSettingH2Prioritization                     `json:"result"`
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

type SettingH2PrioritizationEditResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingH2PrioritizationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingH2PrioritizationEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingH2PrioritizationEditResponseEnvelopeErrors]
type settingH2PrioritizationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingH2PrioritizationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingH2PrioritizationEditResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    settingH2PrioritizationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingH2PrioritizationEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingH2PrioritizationEditResponseEnvelopeMessages]
type settingH2PrioritizationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingH2PrioritizationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingH2PrioritizationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingH2PrioritizationGetResponseEnvelope struct {
	Errors   []SettingH2PrioritizationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingH2PrioritizationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result ZoneSettingH2Prioritization                    `json:"result"`
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

type SettingH2PrioritizationGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingH2PrioritizationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingH2PrioritizationGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingH2PrioritizationGetResponseEnvelopeErrors]
type settingH2PrioritizationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingH2PrioritizationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingH2PrioritizationGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingH2PrioritizationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingH2PrioritizationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingH2PrioritizationGetResponseEnvelopeMessages]
type settingH2PrioritizationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingH2PrioritizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingH2PrioritizationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
