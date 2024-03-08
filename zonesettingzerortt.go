// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ZoneSettingZeroRTTService) Edit(ctx context.Context, params ZoneSettingZeroRTTEditParams, opts ...option.RequestOption) (res *ZoneSettingZeroRTTEditResponse, err error) {
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
func (r *ZoneSettingZeroRTTService) Get(ctx context.Context, query ZoneSettingZeroRTTGetParams, opts ...option.RequestOption) (res *ZoneSettingZeroRTTGetResponse, err error) {
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
type ZoneSettingZeroRTTEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingZeroRTTEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingZeroRTTEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingZeroRTTEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingZeroRTTEditResponseJSON `json:"-"`
}

// zoneSettingZeroRTTEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingZeroRTTEditResponse]
type zoneSettingZeroRTTEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRTTEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingZeroRTTEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingZeroRTTEditResponseID string

const (
	ZoneSettingZeroRTTEditResponseID0rtt ZoneSettingZeroRTTEditResponseID = "0rtt"
)

// Current value of the zone setting.
type ZoneSettingZeroRTTEditResponseValue string

const (
	ZoneSettingZeroRTTEditResponseValueOn  ZoneSettingZeroRTTEditResponseValue = "on"
	ZoneSettingZeroRTTEditResponseValueOff ZoneSettingZeroRTTEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingZeroRTTEditResponseEditable bool

const (
	ZoneSettingZeroRTTEditResponseEditableTrue  ZoneSettingZeroRTTEditResponseEditable = true
	ZoneSettingZeroRTTEditResponseEditableFalse ZoneSettingZeroRTTEditResponseEditable = false
)

// 0-RTT session resumption enabled for this zone.
type ZoneSettingZeroRTTGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingZeroRTTGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingZeroRTTGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingZeroRTTGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingZeroRTTGetResponseJSON `json:"-"`
}

// zoneSettingZeroRTTGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingZeroRTTGetResponse]
type zoneSettingZeroRTTGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingZeroRTTGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingZeroRTTGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingZeroRTTGetResponseID string

const (
	ZoneSettingZeroRTTGetResponseID0rtt ZoneSettingZeroRTTGetResponseID = "0rtt"
)

// Current value of the zone setting.
type ZoneSettingZeroRTTGetResponseValue string

const (
	ZoneSettingZeroRTTGetResponseValueOn  ZoneSettingZeroRTTGetResponseValue = "on"
	ZoneSettingZeroRTTGetResponseValueOff ZoneSettingZeroRTTGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingZeroRTTGetResponseEditable bool

const (
	ZoneSettingZeroRTTGetResponseEditableTrue  ZoneSettingZeroRTTGetResponseEditable = true
	ZoneSettingZeroRTTGetResponseEditableFalse ZoneSettingZeroRTTGetResponseEditable = false
)

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
	Result ZoneSettingZeroRTTEditResponse             `json:"result"`
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

func (r zoneSettingZeroRTTEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingZeroRTTEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingZeroRTTEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result ZoneSettingZeroRTTGetResponse             `json:"result"`
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

func (r zoneSettingZeroRTTGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingZeroRTTGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingZeroRTTGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
