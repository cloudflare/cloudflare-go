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

// SettingZeroRTTService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingZeroRTTService] method
// instead.
type SettingZeroRTTService struct {
	Options []option.RequestOption
}

// NewSettingZeroRTTService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingZeroRTTService(opts ...option.RequestOption) (r *SettingZeroRTTService) {
	r = &SettingZeroRTTService{}
	r.Options = opts
	return
}

// Changes the 0-RTT session resumption setting.
func (r *SettingZeroRTTService) Edit(ctx context.Context, params SettingZeroRTTEditParams, opts ...option.RequestOption) (res *Zones0rtt, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingZeroRTTEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/0rtt", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets 0-RTT session resumption setting.
func (r *SettingZeroRTTService) Get(ctx context.Context, query SettingZeroRTTGetParams, opts ...option.RequestOption) (res *Zones0rtt, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingZeroRTTGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/0rtt", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// 0-RTT session resumption enabled for this zone.
type Zones0rtt struct {
	// ID of the zone setting.
	ID Zones0rttID `json:"id,required"`
	// Current value of the zone setting.
	Value Zones0rttValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable Zones0rttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time     `json:"modified_on,nullable" format:"date-time"`
	JSON       zones0rttJSON `json:"-"`
}

// zones0rttJSON contains the JSON metadata for the struct [Zones0rtt]
type zones0rttJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Zones0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zones0rttJSON) RawJSON() string {
	return r.raw
}

func (r Zones0rtt) implementsZonesSettingEditResponse() {}

func (r Zones0rtt) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type Zones0rttID string

const (
	Zones0rttID0rtt Zones0rttID = "0rtt"
)

// Current value of the zone setting.
type Zones0rttValue string

const (
	Zones0rttValueOn  Zones0rttValue = "on"
	Zones0rttValueOff Zones0rttValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type Zones0rttEditable bool

const (
	Zones0rttEditableTrue  Zones0rttEditable = true
	Zones0rttEditableFalse Zones0rttEditable = false
)

// 0-RTT session resumption enabled for this zone.
type Zones0rttParam struct {
	// ID of the zone setting.
	ID param.Field[Zones0rttID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[Zones0rttValue] `json:"value,required"`
}

func (r Zones0rttParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r Zones0rttParam) implementsZonesSettingEditParamsItem() {}

type SettingZeroRTTEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the 0-RTT setting.
	Value param.Field[SettingZeroRTTEditParamsValue] `json:"value,required"`
}

func (r SettingZeroRTTEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the 0-RTT setting.
type SettingZeroRTTEditParamsValue string

const (
	SettingZeroRTTEditParamsValueOn  SettingZeroRTTEditParamsValue = "on"
	SettingZeroRTTEditParamsValueOff SettingZeroRTTEditParamsValue = "off"
)

type SettingZeroRTTEditResponseEnvelope struct {
	Errors   []SettingZeroRTTEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingZeroRTTEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result Zones0rtt                              `json:"result"`
	JSON   settingZeroRTTEditResponseEnvelopeJSON `json:"-"`
}

// settingZeroRTTEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZeroRTTEditResponseEnvelope]
type settingZeroRTTEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRTTEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZeroRTTEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingZeroRTTEditResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingZeroRTTEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingZeroRTTEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingZeroRTTEditResponseEnvelopeErrors]
type settingZeroRTTEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRTTEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZeroRTTEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingZeroRTTEditResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    settingZeroRTTEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingZeroRTTEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingZeroRTTEditResponseEnvelopeMessages]
type settingZeroRTTEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRTTEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZeroRTTEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingZeroRTTGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingZeroRTTGetResponseEnvelope struct {
	Errors   []SettingZeroRTTGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingZeroRTTGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result Zones0rtt                             `json:"result"`
	JSON   settingZeroRTTGetResponseEnvelopeJSON `json:"-"`
}

// settingZeroRTTGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZeroRTTGetResponseEnvelope]
type settingZeroRTTGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRTTGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZeroRTTGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingZeroRTTGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingZeroRTTGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingZeroRTTGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingZeroRTTGetResponseEnvelopeErrors]
type settingZeroRTTGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRTTGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZeroRTTGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingZeroRTTGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingZeroRTTGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingZeroRTTGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingZeroRTTGetResponseEnvelopeMessages]
type settingZeroRTTGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZeroRTTGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZeroRTTGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
