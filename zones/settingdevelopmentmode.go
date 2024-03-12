// File generated from our OpenAPI spec by Stainless.

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

// SettingDevelopmentModeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingDevelopmentModeService]
// method instead.
type SettingDevelopmentModeService struct {
	Options []option.RequestOption
}

// NewSettingDevelopmentModeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingDevelopmentModeService(opts ...option.RequestOption) (r *SettingDevelopmentModeService) {
	r = &SettingDevelopmentModeService{}
	r.Options = opts
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
func (r *SettingDevelopmentModeService) Edit(ctx context.Context, params SettingDevelopmentModeEditParams, opts ...option.RequestOption) (res *ZonesDevelopmentMode, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingDevelopmentModeEditResponseEnvelope
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
func (r *SettingDevelopmentModeService) Get(ctx context.Context, query SettingDevelopmentModeGetParams, opts ...option.RequestOption) (res *ZonesDevelopmentMode, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingDevelopmentModeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/development_mode", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
type ZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID ZonesDevelopmentModeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesDevelopmentModeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                  `json:"time_remaining"`
	JSON          zonesDevelopmentModeJSON `json:"-"`
}

// zonesDevelopmentModeJSON contains the JSON metadata for the struct
// [ZonesDevelopmentMode]
type zonesDevelopmentModeJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZonesDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesDevelopmentModeJSON) RawJSON() string {
	return r.raw
}

func (r ZonesDevelopmentMode) implementsZonesSettingEditResponse() {}

func (r ZonesDevelopmentMode) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesDevelopmentModeID string

const (
	ZonesDevelopmentModeIDDevelopmentMode ZonesDevelopmentModeID = "development_mode"
)

// Current value of the zone setting.
type ZonesDevelopmentModeValue string

const (
	ZonesDevelopmentModeValueOn  ZonesDevelopmentModeValue = "on"
	ZonesDevelopmentModeValueOff ZonesDevelopmentModeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesDevelopmentModeEditable bool

const (
	ZonesDevelopmentModeEditableTrue  ZonesDevelopmentModeEditable = true
	ZonesDevelopmentModeEditableFalse ZonesDevelopmentModeEditable = false
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZonesDevelopmentModeParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesDevelopmentModeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesDevelopmentModeValue] `json:"value,required"`
}

func (r ZonesDevelopmentModeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesDevelopmentModeParam) implementsZonesSettingEditParamsItem() {}

type SettingDevelopmentModeEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingDevelopmentModeEditParamsValue] `json:"value,required"`
}

func (r SettingDevelopmentModeEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingDevelopmentModeEditParamsValue string

const (
	SettingDevelopmentModeEditParamsValueOn  SettingDevelopmentModeEditParamsValue = "on"
	SettingDevelopmentModeEditParamsValueOff SettingDevelopmentModeEditParamsValue = "off"
)

type SettingDevelopmentModeEditResponseEnvelope struct {
	Errors   []SettingDevelopmentModeEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingDevelopmentModeEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result ZonesDevelopmentMode                           `json:"result"`
	JSON   settingDevelopmentModeEditResponseEnvelopeJSON `json:"-"`
}

// settingDevelopmentModeEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingDevelopmentModeEditResponseEnvelope]
type settingDevelopmentModeEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDevelopmentModeEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDevelopmentModeEditResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingDevelopmentModeEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingDevelopmentModeEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingDevelopmentModeEditResponseEnvelopeErrors]
type settingDevelopmentModeEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDevelopmentModeEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingDevelopmentModeEditResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingDevelopmentModeEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingDevelopmentModeEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingDevelopmentModeEditResponseEnvelopeMessages]
type settingDevelopmentModeEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDevelopmentModeEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingDevelopmentModeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingDevelopmentModeGetResponseEnvelope struct {
	Errors   []SettingDevelopmentModeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingDevelopmentModeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result ZonesDevelopmentMode                          `json:"result"`
	JSON   settingDevelopmentModeGetResponseEnvelopeJSON `json:"-"`
}

// settingDevelopmentModeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingDevelopmentModeGetResponseEnvelope]
type settingDevelopmentModeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDevelopmentModeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingDevelopmentModeGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingDevelopmentModeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingDevelopmentModeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingDevelopmentModeGetResponseEnvelopeErrors]
type settingDevelopmentModeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDevelopmentModeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingDevelopmentModeGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingDevelopmentModeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingDevelopmentModeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingDevelopmentModeGetResponseEnvelopeMessages]
type settingDevelopmentModeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingDevelopmentModeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingDevelopmentModeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
