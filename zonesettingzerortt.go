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

// ZoneSettingZeroRTTService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingZeroRTTService] method
// instead.
type ZoneSettingZeroRTTService struct {
	Options []option.RequestOption
}

// NewZoneSettingZeroRTTService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingZeroRTTService(opts ...option.RequestOption) (r *ZoneSettingZeroRTTService) {
	r = &ZoneSettingZeroRTTService{}
	r.Options = opts
	return
}

// Changes the 0-RTT session resumption setting.
func (r *ZoneSettingZeroRTTService) Edit(ctx context.Context, params ZoneSettingZeroRTTEditParams, opts ...option.RequestOption) (res *Zones0rtt, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingZeroRTTEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/0rtt", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets 0-RTT session resumption setting.
func (r *ZoneSettingZeroRTTService) Get(ctx context.Context, query ZoneSettingZeroRTTGetParams, opts ...option.RequestOption) (res *Zones0rtt, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingZeroRTTGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/0rtt", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

func (r Zones0rtt) implementsZoneSettingEditResponse() {}

func (r Zones0rtt) implementsZoneSettingGetResponse() {}

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

func (r Zones0rttParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingZeroRTTEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the 0-RTT setting.
	Value param.Field[ZoneSettingZeroRTTEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingZeroRTTEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the 0-RTT setting.
type ZoneSettingZeroRTTEditParamsValue string

const (
	ZoneSettingZeroRTTEditParamsValueOn  ZoneSettingZeroRTTEditParamsValue = "on"
	ZoneSettingZeroRTTEditParamsValueOff ZoneSettingZeroRTTEditParamsValue = "off"
)

type ZoneSettingZeroRTTEditResponseEnvelope struct {
	Errors   []ZoneSettingZeroRTTEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingZeroRTTEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result Zones0rtt                                  `json:"result"`
	JSON   zoneSettingZeroRTTEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingZeroRTTEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingZeroRTTEditResponseEnvelope]
type zoneSettingZeroRTTEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRTTEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZeroRTTEditResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingZeroRTTEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingZeroRTTEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingZeroRTTEditResponseEnvelopeErrors]
type zoneSettingZeroRTTEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRTTEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZeroRTTEditResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingZeroRTTEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingZeroRTTEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneSettingZeroRTTEditResponseEnvelopeMessages]
type zoneSettingZeroRTTEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRTTEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZeroRTTGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingZeroRTTGetResponseEnvelope struct {
	Errors   []ZoneSettingZeroRTTGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingZeroRTTGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result Zones0rtt                                 `json:"result"`
	JSON   zoneSettingZeroRTTGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingZeroRTTGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingZeroRTTGetResponseEnvelope]
type zoneSettingZeroRTTGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRTTGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZeroRTTGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingZeroRTTGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingZeroRTTGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingZeroRTTGetResponseEnvelopeErrors]
type zoneSettingZeroRTTGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRTTGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingZeroRTTGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingZeroRTTGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingZeroRTTGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingZeroRTTGetResponseEnvelopeMessages]
type zoneSettingZeroRTTGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRTTGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
