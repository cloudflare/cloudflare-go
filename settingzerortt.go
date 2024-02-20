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

// Changes the 0-RTT session resumption setting.
func (r *SettingZeroRttService) Edit(ctx context.Context, zoneID string, body SettingZeroRttEditParams, opts ...option.RequestOption) (res *SettingZeroRttEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingZeroRttEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/0rtt", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// 0-RTT session resumption enabled for this zone.
type SettingZeroRttEditResponse struct {
	// ID of the zone setting.
	ID SettingZeroRttEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingZeroRttEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingZeroRttEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingZeroRttEditResponseJSON `json:"-"`
}

// settingZeroRttEditResponseJSON contains the JSON metadata for the struct
// [SettingZeroRttEditResponse]
type settingZeroRttEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingZeroRttEditResponseID string

const (
	SettingZeroRttEditResponseID0rtt SettingZeroRttEditResponseID = "0rtt"
)

// Current value of the zone setting.
type SettingZeroRttEditResponseValue string

const (
	SettingZeroRttEditResponseValueOn  SettingZeroRttEditResponseValue = "on"
	SettingZeroRttEditResponseValueOff SettingZeroRttEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingZeroRttEditResponseEditable bool

const (
	SettingZeroRttEditResponseEditableTrue  SettingZeroRttEditResponseEditable = true
	SettingZeroRttEditResponseEditableFalse SettingZeroRttEditResponseEditable = false
)

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

type SettingZeroRttEditParams struct {
	// Value of the 0-RTT setting.
	Value param.Field[SettingZeroRttEditParamsValue] `json:"value,required"`
}

func (r SettingZeroRttEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the 0-RTT setting.
type SettingZeroRttEditParamsValue string

const (
	SettingZeroRttEditParamsValueOn  SettingZeroRttEditParamsValue = "on"
	SettingZeroRttEditParamsValueOff SettingZeroRttEditParamsValue = "off"
)

type SettingZeroRttEditResponseEnvelope struct {
	Errors   []SettingZeroRttEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingZeroRttEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result SettingZeroRttEditResponse             `json:"result"`
	JSON   settingZeroRttEditResponseEnvelopeJSON `json:"-"`
}

// settingZeroRttEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZeroRttEditResponseEnvelope]
type settingZeroRttEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingZeroRttEditResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingZeroRttEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingZeroRttEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingZeroRttEditResponseEnvelopeErrors]
type settingZeroRttEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingZeroRttEditResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    settingZeroRttEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingZeroRttEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingZeroRttEditResponseEnvelopeMessages]
type settingZeroRttEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
