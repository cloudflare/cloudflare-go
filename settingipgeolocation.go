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

// SettingIPGeolocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingIPGeolocationService]
// method instead.
type SettingIPGeolocationService struct {
	Options []option.RequestOption
}

// NewSettingIPGeolocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingIPGeolocationService(opts ...option.RequestOption) (r *SettingIPGeolocationService) {
	r = &SettingIPGeolocationService{}
	r.Options = opts
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
func (r *SettingIPGeolocationService) Edit(ctx context.Context, params SettingIPGeolocationEditParams, opts ...option.RequestOption) (res *SettingIPGeolocationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPGeolocationEditResponseEnvelope
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
func (r *SettingIPGeolocationService) Get(ctx context.Context, query SettingIPGeolocationGetParams, opts ...option.RequestOption) (res *SettingIPGeolocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPGeolocationGetResponseEnvelope
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
type SettingIPGeolocationEditResponse struct {
	// ID of the zone setting.
	ID SettingIPGeolocationEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingIPGeolocationEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingIPGeolocationEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingIPGeolocationEditResponseJSON `json:"-"`
}

// settingIPGeolocationEditResponseJSON contains the JSON metadata for the struct
// [SettingIPGeolocationEditResponse]
type settingIPGeolocationEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingIPGeolocationEditResponseID string

const (
	SettingIPGeolocationEditResponseIDIPGeolocation SettingIPGeolocationEditResponseID = "ip_geolocation"
)

// Current value of the zone setting.
type SettingIPGeolocationEditResponseValue string

const (
	SettingIPGeolocationEditResponseValueOn  SettingIPGeolocationEditResponseValue = "on"
	SettingIPGeolocationEditResponseValueOff SettingIPGeolocationEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingIPGeolocationEditResponseEditable bool

const (
	SettingIPGeolocationEditResponseEditableTrue  SettingIPGeolocationEditResponseEditable = true
	SettingIPGeolocationEditResponseEditableFalse SettingIPGeolocationEditResponseEditable = false
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingIPGeolocationGetResponse struct {
	// ID of the zone setting.
	ID SettingIPGeolocationGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingIPGeolocationGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingIPGeolocationGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingIPGeolocationGetResponseJSON `json:"-"`
}

// settingIPGeolocationGetResponseJSON contains the JSON metadata for the struct
// [SettingIPGeolocationGetResponse]
type settingIPGeolocationGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingIPGeolocationGetResponseID string

const (
	SettingIPGeolocationGetResponseIDIPGeolocation SettingIPGeolocationGetResponseID = "ip_geolocation"
)

// Current value of the zone setting.
type SettingIPGeolocationGetResponseValue string

const (
	SettingIPGeolocationGetResponseValueOn  SettingIPGeolocationGetResponseValue = "on"
	SettingIPGeolocationGetResponseValueOff SettingIPGeolocationGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingIPGeolocationGetResponseEditable bool

const (
	SettingIPGeolocationGetResponseEditableTrue  SettingIPGeolocationGetResponseEditable = true
	SettingIPGeolocationGetResponseEditableFalse SettingIPGeolocationGetResponseEditable = false
)

type SettingIPGeolocationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingIPGeolocationEditParamsValue] `json:"value,required"`
}

func (r SettingIPGeolocationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingIPGeolocationEditParamsValue string

const (
	SettingIPGeolocationEditParamsValueOn  SettingIPGeolocationEditParamsValue = "on"
	SettingIPGeolocationEditParamsValueOff SettingIPGeolocationEditParamsValue = "off"
)

type SettingIPGeolocationEditResponseEnvelope struct {
	Errors   []SettingIPGeolocationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingIPGeolocationEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result SettingIPGeolocationEditResponse             `json:"result"`
	JSON   settingIPGeolocationEditResponseEnvelopeJSON `json:"-"`
}

// settingIPGeolocationEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingIPGeolocationEditResponseEnvelope]
type settingIPGeolocationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPGeolocationEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingIPGeolocationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingIPGeolocationEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingIPGeolocationEditResponseEnvelopeErrors]
type settingIPGeolocationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPGeolocationEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingIPGeolocationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingIPGeolocationEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingIPGeolocationEditResponseEnvelopeMessages]
type settingIPGeolocationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPGeolocationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingIPGeolocationGetResponseEnvelope struct {
	Errors   []SettingIPGeolocationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingIPGeolocationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result SettingIPGeolocationGetResponse             `json:"result"`
	JSON   settingIPGeolocationGetResponseEnvelopeJSON `json:"-"`
}

// settingIPGeolocationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingIPGeolocationGetResponseEnvelope]
type settingIPGeolocationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPGeolocationGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingIPGeolocationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingIPGeolocationGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingIPGeolocationGetResponseEnvelopeErrors]
type settingIPGeolocationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPGeolocationGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingIPGeolocationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingIPGeolocationGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingIPGeolocationGetResponseEnvelopeMessages]
type settingIPGeolocationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
