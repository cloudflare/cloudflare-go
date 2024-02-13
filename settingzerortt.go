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

// SettingZeroRttService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingZeroRttService] method
// instead.
type SettingZeroRttService struct {
	Options []option.RequestOption
}

// NewSettingZeroRttService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingZeroRttService(opts ...option.RequestOption) (r *SettingZeroRttService) {
	r = &SettingZeroRttService{}
	r.Options = opts
	return
}

// Gets 0-RTT session resumption setting.
func (r *SettingZeroRttService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingZeroRttGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingZeroRttGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/0rtt", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Changes the 0-RTT session resumption setting.
func (r *SettingZeroRttService) ZoneSettingsChange0RttSessionResumptionSetting(ctx context.Context, zoneID string, body SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParams, opts ...option.RequestOption) (res *SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/0rtt", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// 0-RTT session resumption enabled for this zone.
type SettingZeroRttGetResponse struct {
	// ID of the zone setting.
	ID SettingZeroRttGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingZeroRttGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingZeroRttGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingZeroRttGetResponseJSON `json:"-"`
}

// settingZeroRttGetResponseJSON contains the JSON metadata for the struct
// [SettingZeroRttGetResponse]
type settingZeroRttGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingZeroRttGetResponseID string

const (
	SettingZeroRttGetResponseID0rtt SettingZeroRttGetResponseID = "0rtt"
)

// Current value of the zone setting.
type SettingZeroRttGetResponseValue string

const (
	SettingZeroRttGetResponseValueOn  SettingZeroRttGetResponseValue = "on"
	SettingZeroRttGetResponseValueOff SettingZeroRttGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingZeroRttGetResponseEditable bool

const (
	SettingZeroRttGetResponseEditableTrue  SettingZeroRttGetResponseEditable = true
	SettingZeroRttGetResponseEditableFalse SettingZeroRttGetResponseEditable = false
)

// 0-RTT session resumption enabled for this zone.
type SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponse struct {
	// ID of the zone setting.
	ID SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseJSON `json:"-"`
}

// settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseJSON
// contains the JSON metadata for the struct
// [SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponse]
type settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseID string

const (
	SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseID0rtt SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseID = "0rtt"
)

// Current value of the zone setting.
type SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseValue string

const (
	SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseValueOn  SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseValue = "on"
	SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseValueOff SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEditable bool

const (
	SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEditableTrue  SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEditable = true
	SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEditableFalse SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEditable = false
)

type SettingZeroRttGetResponseEnvelope struct {
	Errors   []SettingZeroRttGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingZeroRttGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result SettingZeroRttGetResponse             `json:"result"`
	JSON   settingZeroRttGetResponseEnvelopeJSON `json:"-"`
}

// settingZeroRttGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZeroRttGetResponseEnvelope]
type settingZeroRttGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingZeroRttGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingZeroRttGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingZeroRttGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingZeroRttGetResponseEnvelopeErrors]
type settingZeroRttGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingZeroRttGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingZeroRttGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingZeroRttGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingZeroRttGetResponseEnvelopeMessages]
type settingZeroRttGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParams struct {
	// Value of the 0-RTT setting.
	Value param.Field[SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValue] `json:"value,required"`
}

func (r SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the 0-RTT setting.
type SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValue string

const (
	SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValueOn  SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValue = "on"
	SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValueOff SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingParamsValue = "off"
)

type SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelope struct {
	Errors   []SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponse             `json:"result"`
	JSON   settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeJSON `json:"-"`
}

// settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelope]
type settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeErrors struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeErrorsJSON `json:"-"`
}

// settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeErrors]
type settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeMessages struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeMessagesJSON `json:"-"`
}

// settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeMessages]
type settingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttZoneSettingsChange0RttSessionResumptionSettingResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
