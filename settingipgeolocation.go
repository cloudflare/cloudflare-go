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
func (r *SettingIPGeolocationService) Update(ctx context.Context, zoneID string, body SettingIPGeolocationUpdateParams, opts ...option.RequestOption) (res *SettingIPGeolocationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPGeolocationUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ip_geolocation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
func (r *SettingIPGeolocationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingIPGeolocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPGeolocationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ip_geolocation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingIPGeolocationUpdateResponse struct {
	// ID of the zone setting.
	ID SettingIPGeolocationUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingIPGeolocationUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingIPGeolocationUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingIPGeolocationUpdateResponseJSON `json:"-"`
}

// settingIPGeolocationUpdateResponseJSON contains the JSON metadata for the struct
// [SettingIPGeolocationUpdateResponse]
type settingIPGeolocationUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingIPGeolocationUpdateResponseID string

const (
	SettingIPGeolocationUpdateResponseIDIPGeolocation SettingIPGeolocationUpdateResponseID = "ip_geolocation"
)

// Current value of the zone setting.
type SettingIPGeolocationUpdateResponseValue string

const (
	SettingIPGeolocationUpdateResponseValueOn  SettingIPGeolocationUpdateResponseValue = "on"
	SettingIPGeolocationUpdateResponseValueOff SettingIPGeolocationUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingIPGeolocationUpdateResponseEditable bool

const (
	SettingIPGeolocationUpdateResponseEditableTrue  SettingIPGeolocationUpdateResponseEditable = true
	SettingIPGeolocationUpdateResponseEditableFalse SettingIPGeolocationUpdateResponseEditable = false
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

type SettingIPGeolocationUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingIPGeolocationUpdateParamsValue] `json:"value,required"`
}

func (r SettingIPGeolocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingIPGeolocationUpdateParamsValue string

const (
	SettingIPGeolocationUpdateParamsValueOn  SettingIPGeolocationUpdateParamsValue = "on"
	SettingIPGeolocationUpdateParamsValueOff SettingIPGeolocationUpdateParamsValue = "off"
)

type SettingIPGeolocationUpdateResponseEnvelope struct {
	Errors   []SettingIPGeolocationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingIPGeolocationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result SettingIPGeolocationUpdateResponse             `json:"result"`
	JSON   settingIPGeolocationUpdateResponseEnvelopeJSON `json:"-"`
}

// settingIPGeolocationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingIPGeolocationUpdateResponseEnvelope]
type settingIPGeolocationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPGeolocationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingIPGeolocationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingIPGeolocationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingIPGeolocationUpdateResponseEnvelopeErrors]
type settingIPGeolocationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingIPGeolocationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingIPGeolocationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingIPGeolocationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingIPGeolocationUpdateResponseEnvelopeMessages]
type settingIPGeolocationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
