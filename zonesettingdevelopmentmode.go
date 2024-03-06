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

// ZoneSettingDevelopmentModeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingDevelopmentModeService] method instead.
type ZoneSettingDevelopmentModeService struct {
	Options []option.RequestOption
}

// NewZoneSettingDevelopmentModeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingDevelopmentModeService(opts ...option.RequestOption) (r *ZoneSettingDevelopmentModeService) {
	r = &ZoneSettingDevelopmentModeService{}
	r.Options = opts
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
func (r *ZoneSettingDevelopmentModeService) Edit(ctx context.Context, params ZoneSettingDevelopmentModeEditParams, opts ...option.RequestOption) (res *ZoneSettingDevelopmentModeEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingDevelopmentModeEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/development_mode", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
func (r *ZoneSettingDevelopmentModeService) Get(ctx context.Context, query ZoneSettingDevelopmentModeGetParams, opts ...option.RequestOption) (res *ZoneSettingDevelopmentModeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingDevelopmentModeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/development_mode", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingDevelopmentModeEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingDevelopmentModeEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingDevelopmentModeEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingDevelopmentModeEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                                    `json:"time_remaining"`
	JSON          zoneSettingDevelopmentModeEditResponseJSON `json:"-"`
}

// zoneSettingDevelopmentModeEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingDevelopmentModeEditResponse]
type zoneSettingDevelopmentModeEditResponseJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingDevelopmentModeEditResponseID string

const (
	ZoneSettingDevelopmentModeEditResponseIDDevelopmentMode ZoneSettingDevelopmentModeEditResponseID = "development_mode"
)

// Current value of the zone setting.
type ZoneSettingDevelopmentModeEditResponseValue string

const (
	ZoneSettingDevelopmentModeEditResponseValueOn  ZoneSettingDevelopmentModeEditResponseValue = "on"
	ZoneSettingDevelopmentModeEditResponseValueOff ZoneSettingDevelopmentModeEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingDevelopmentModeEditResponseEditable bool

const (
	ZoneSettingDevelopmentModeEditResponseEditableTrue  ZoneSettingDevelopmentModeEditResponseEditable = true
	ZoneSettingDevelopmentModeEditResponseEditableFalse ZoneSettingDevelopmentModeEditResponseEditable = false
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingDevelopmentModeGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingDevelopmentModeGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingDevelopmentModeGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingDevelopmentModeGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                                   `json:"time_remaining"`
	JSON          zoneSettingDevelopmentModeGetResponseJSON `json:"-"`
}

// zoneSettingDevelopmentModeGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingDevelopmentModeGetResponse]
type zoneSettingDevelopmentModeGetResponseJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingDevelopmentModeGetResponseID string

const (
	ZoneSettingDevelopmentModeGetResponseIDDevelopmentMode ZoneSettingDevelopmentModeGetResponseID = "development_mode"
)

// Current value of the zone setting.
type ZoneSettingDevelopmentModeGetResponseValue string

const (
	ZoneSettingDevelopmentModeGetResponseValueOn  ZoneSettingDevelopmentModeGetResponseValue = "on"
	ZoneSettingDevelopmentModeGetResponseValueOff ZoneSettingDevelopmentModeGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingDevelopmentModeGetResponseEditable bool

const (
	ZoneSettingDevelopmentModeGetResponseEditableTrue  ZoneSettingDevelopmentModeGetResponseEditable = true
	ZoneSettingDevelopmentModeGetResponseEditableFalse ZoneSettingDevelopmentModeGetResponseEditable = false
)

type ZoneSettingDevelopmentModeEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingDevelopmentModeEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingDevelopmentModeEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingDevelopmentModeEditParamsValue string

const (
	ZoneSettingDevelopmentModeEditParamsValueOn  ZoneSettingDevelopmentModeEditParamsValue = "on"
	ZoneSettingDevelopmentModeEditParamsValueOff ZoneSettingDevelopmentModeEditParamsValue = "off"
)

type ZoneSettingDevelopmentModeEditResponseEnvelope struct {
	Errors   []ZoneSettingDevelopmentModeEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingDevelopmentModeEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result ZoneSettingDevelopmentModeEditResponse             `json:"result"`
	JSON   zoneSettingDevelopmentModeEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingDevelopmentModeEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingDevelopmentModeEditResponseEnvelope]
type zoneSettingDevelopmentModeEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingDevelopmentModeEditResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingDevelopmentModeEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingDevelopmentModeEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingDevelopmentModeEditResponseEnvelopeErrors]
type zoneSettingDevelopmentModeEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingDevelopmentModeEditResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingDevelopmentModeEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingDevelopmentModeEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingDevelopmentModeEditResponseEnvelopeMessages]
type zoneSettingDevelopmentModeEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingDevelopmentModeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingDevelopmentModeGetResponseEnvelope struct {
	Errors   []ZoneSettingDevelopmentModeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingDevelopmentModeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result ZoneSettingDevelopmentModeGetResponse             `json:"result"`
	JSON   zoneSettingDevelopmentModeGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingDevelopmentModeGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingDevelopmentModeGetResponseEnvelope]
type zoneSettingDevelopmentModeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingDevelopmentModeGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingDevelopmentModeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingDevelopmentModeGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingDevelopmentModeGetResponseEnvelopeErrors]
type zoneSettingDevelopmentModeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingDevelopmentModeGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingDevelopmentModeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingDevelopmentModeGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingDevelopmentModeGetResponseEnvelopeMessages]
type zoneSettingDevelopmentModeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentModeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
