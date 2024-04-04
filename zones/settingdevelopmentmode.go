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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *SettingDevelopmentModeService) Edit(ctx context.Context, params SettingDevelopmentModeEditParams, opts ...option.RequestOption) (res *ZoneSettingDevelopmentMode, err error) {
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
func (r *SettingDevelopmentModeService) Get(ctx context.Context, query SettingDevelopmentModeGetParams, opts ...option.RequestOption) (res *ZoneSettingDevelopmentMode, err error) {
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
type ZoneSettingDevelopmentMode struct {
	// ID of the zone setting.
	ID ZoneSettingDevelopmentModeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingDevelopmentModeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                        `json:"time_remaining"`
	JSON          zoneSettingDevelopmentModeJSON `json:"-"`
}

// zoneSettingDevelopmentModeJSON contains the JSON metadata for the struct
// [ZoneSettingDevelopmentMode]
type zoneSettingDevelopmentModeJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingDevelopmentModeJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingDevelopmentMode) implementsZonesSettingEditResponse() {}

func (r ZoneSettingDevelopmentMode) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingDevelopmentModeID string

const (
	ZoneSettingDevelopmentModeIDDevelopmentMode ZoneSettingDevelopmentModeID = "development_mode"
)

func (r ZoneSettingDevelopmentModeID) IsKnown() bool {
	switch r {
	case ZoneSettingDevelopmentModeIDDevelopmentMode:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingDevelopmentModeValue string

const (
	ZoneSettingDevelopmentModeValueOn  ZoneSettingDevelopmentModeValue = "on"
	ZoneSettingDevelopmentModeValueOff ZoneSettingDevelopmentModeValue = "off"
)

func (r ZoneSettingDevelopmentModeValue) IsKnown() bool {
	switch r {
	case ZoneSettingDevelopmentModeValueOn, ZoneSettingDevelopmentModeValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingDevelopmentModeEditable bool

const (
	ZoneSettingDevelopmentModeEditableTrue  ZoneSettingDevelopmentModeEditable = true
	ZoneSettingDevelopmentModeEditableFalse ZoneSettingDevelopmentModeEditable = false
)

func (r ZoneSettingDevelopmentModeEditable) IsKnown() bool {
	switch r {
	case ZoneSettingDevelopmentModeEditableTrue, ZoneSettingDevelopmentModeEditableFalse:
		return true
	}
	return false
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingDevelopmentModeParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingDevelopmentModeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingDevelopmentModeValue] `json:"value,required"`
}

func (r ZoneSettingDevelopmentModeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingDevelopmentModeParam) implementsZonesSettingEditParamsItem() {}

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

func (r SettingDevelopmentModeEditParamsValue) IsKnown() bool {
	switch r {
	case SettingDevelopmentModeEditParamsValueOn, SettingDevelopmentModeEditParamsValueOff:
		return true
	}
	return false
}

type SettingDevelopmentModeEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result ZoneSettingDevelopmentMode                     `json:"result"`
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

type SettingDevelopmentModeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingDevelopmentModeGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result ZoneSettingDevelopmentMode                    `json:"result"`
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
