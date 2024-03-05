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

// ZoneSettingIPGeolocationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingIPGeolocationService] method instead.
type ZoneSettingIPGeolocationService struct {
	Options []option.RequestOption
}

// NewZoneSettingIPGeolocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingIPGeolocationService(opts ...option.RequestOption) (r *ZoneSettingIPGeolocationService) {
	r = &ZoneSettingIPGeolocationService{}
	r.Options = opts
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
func (r *ZoneSettingIPGeolocationService) Edit(ctx context.Context, params ZoneSettingIPGeolocationEditParams, opts ...option.RequestOption) (res *ZonesIPGeolocation, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingIPGeolocationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ip_geolocation", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
func (r *ZoneSettingIPGeolocationService) Get(ctx context.Context, query ZoneSettingIPGeolocationGetParams, opts ...option.RequestOption) (res *ZonesIPGeolocation, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingIPGeolocationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ip_geolocation", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZonesIPGeolocation struct {
	// ID of the zone setting.
	ID ZonesIPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesIPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesIPGeolocationJSON `json:"-"`
}

// zonesIPGeolocationJSON contains the JSON metadata for the struct
// [ZonesIPGeolocation]
type zonesIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesIPGeolocation) implementsZoneSettingEditResponse() {}

func (r ZonesIPGeolocation) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesIPGeolocationID string

const (
	ZonesIPGeolocationIDIPGeolocation ZonesIPGeolocationID = "ip_geolocation"
)

// Current value of the zone setting.
type ZonesIPGeolocationValue string

const (
	ZonesIPGeolocationValueOn  ZonesIPGeolocationValue = "on"
	ZonesIPGeolocationValueOff ZonesIPGeolocationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesIPGeolocationEditable bool

const (
	ZonesIPGeolocationEditableTrue  ZonesIPGeolocationEditable = true
	ZonesIPGeolocationEditableFalse ZonesIPGeolocationEditable = false
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZonesIPGeolocationParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesIPGeolocationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesIPGeolocationValue] `json:"value,required"`
}

func (r ZonesIPGeolocationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesIPGeolocationParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingIPGeolocationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingIPGeolocationEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingIPGeolocationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingIPGeolocationEditParamsValue string

const (
	ZoneSettingIPGeolocationEditParamsValueOn  ZoneSettingIPGeolocationEditParamsValue = "on"
	ZoneSettingIPGeolocationEditParamsValueOff ZoneSettingIPGeolocationEditParamsValue = "off"
)

type ZoneSettingIPGeolocationEditResponseEnvelope struct {
	Errors   []ZoneSettingIPGeolocationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingIPGeolocationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result ZonesIPGeolocation                               `json:"result"`
	JSON   zoneSettingIPGeolocationEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingIPGeolocationEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingIPGeolocationEditResponseEnvelope]
type zoneSettingIPGeolocationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPGeolocationEditResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zoneSettingIPGeolocationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingIPGeolocationEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingIPGeolocationEditResponseEnvelopeErrors]
type zoneSettingIPGeolocationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPGeolocationEditResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingIPGeolocationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingIPGeolocationEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingIPGeolocationEditResponseEnvelopeMessages]
type zoneSettingIPGeolocationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPGeolocationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingIPGeolocationGetResponseEnvelope struct {
	Errors   []ZoneSettingIPGeolocationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingIPGeolocationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result ZonesIPGeolocation                              `json:"result"`
	JSON   zoneSettingIPGeolocationGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingIPGeolocationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingIPGeolocationGetResponseEnvelope]
type zoneSettingIPGeolocationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPGeolocationGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zoneSettingIPGeolocationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingIPGeolocationGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneSettingIPGeolocationGetResponseEnvelopeErrors]
type zoneSettingIPGeolocationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingIPGeolocationGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingIPGeolocationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingIPGeolocationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingIPGeolocationGetResponseEnvelopeMessages]
type zoneSettingIPGeolocationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingIPGeolocationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
