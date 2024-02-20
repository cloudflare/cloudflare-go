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
func (r *SettingZeroRttService) Update(ctx context.Context, zoneID string, body SettingZeroRttUpdateParams, opts ...option.RequestOption) (res *SettingZeroRttUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingZeroRttUpdateResponseEnvelope
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
type SettingZeroRttUpdateResponse struct {
	// ID of the zone setting.
	ID SettingZeroRttUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingZeroRttUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingZeroRttUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingZeroRttUpdateResponseJSON `json:"-"`
}

// settingZeroRttUpdateResponseJSON contains the JSON metadata for the struct
// [SettingZeroRttUpdateResponse]
type settingZeroRttUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingZeroRttUpdateResponseID string

const (
	SettingZeroRttUpdateResponseID0rtt SettingZeroRttUpdateResponseID = "0rtt"
)

// Current value of the zone setting.
type SettingZeroRttUpdateResponseValue string

const (
	SettingZeroRttUpdateResponseValueOn  SettingZeroRttUpdateResponseValue = "on"
	SettingZeroRttUpdateResponseValueOff SettingZeroRttUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingZeroRttUpdateResponseEditable bool

const (
	SettingZeroRttUpdateResponseEditableTrue  SettingZeroRttUpdateResponseEditable = true
	SettingZeroRttUpdateResponseEditableFalse SettingZeroRttUpdateResponseEditable = false
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

type SettingZeroRttUpdateParams struct {
	// Value of the 0-RTT setting.
	Value param.Field[SettingZeroRttUpdateParamsValue] `json:"value,required"`
}

func (r SettingZeroRttUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the 0-RTT setting.
type SettingZeroRttUpdateParamsValue string

const (
	SettingZeroRttUpdateParamsValueOn  SettingZeroRttUpdateParamsValue = "on"
	SettingZeroRttUpdateParamsValueOff SettingZeroRttUpdateParamsValue = "off"
)

type SettingZeroRttUpdateResponseEnvelope struct {
	Errors   []SettingZeroRttUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingZeroRttUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result SettingZeroRttUpdateResponse             `json:"result"`
	JSON   settingZeroRttUpdateResponseEnvelopeJSON `json:"-"`
}

// settingZeroRttUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingZeroRttUpdateResponseEnvelope]
type settingZeroRttUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingZeroRttUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    settingZeroRttUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingZeroRttUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingZeroRttUpdateResponseEnvelopeErrors]
type settingZeroRttUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingZeroRttUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingZeroRttUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingZeroRttUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingZeroRttUpdateResponseEnvelopeMessages]
type settingZeroRttUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRttUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
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
