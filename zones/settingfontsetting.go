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
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingFontSettingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingFontSettingService] method instead.
type SettingFontSettingService struct {
	Options []option.RequestOption
}

// NewSettingFontSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingFontSettingService(opts ...option.RequestOption) (r *SettingFontSettingService) {
	r = &SettingFontSettingService{}
	r.Options = opts
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *SettingFontSettingService) Edit(ctx context.Context, params SettingFontSettingEditParams, opts ...option.RequestOption) (res *FontSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingFontSettingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/fonts", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *SettingFontSettingService) Get(ctx context.Context, query SettingFontSettingGetParams, opts ...option.RequestOption) (res *FontSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingFontSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/fonts", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type FontSettings struct {
	// ID of the zone setting.
	ID FontSettingsID `json:"id,required"`
	// Current value of the zone setting.
	Value FontSettingsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable FontSettingsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on,nullable" format:"date-time"`
	JSON       fontSettingsJSON `json:"-"`
}

// fontSettingsJSON contains the JSON metadata for the struct [FontSettings]
type fontSettingsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fontSettingsJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type FontSettingsID string

const (
	FontSettingsIDFonts FontSettingsID = "fonts"
)

func (r FontSettingsID) IsKnown() bool {
	switch r {
	case FontSettingsIDFonts:
		return true
	}
	return false
}

// Current value of the zone setting.
type FontSettingsValue string

const (
	FontSettingsValueOn  FontSettingsValue = "on"
	FontSettingsValueOff FontSettingsValue = "off"
)

func (r FontSettingsValue) IsKnown() bool {
	switch r {
	case FontSettingsValueOn, FontSettingsValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type FontSettingsEditable bool

const (
	FontSettingsEditableTrue  FontSettingsEditable = true
	FontSettingsEditableFalse FontSettingsEditable = false
)

func (r FontSettingsEditable) IsKnown() bool {
	switch r {
	case FontSettingsEditableTrue, FontSettingsEditableFalse:
		return true
	}
	return false
}

type SettingFontSettingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Whether the feature is enabled or disabled.
	Value param.Field[SettingFontSettingEditParamsValue] `json:"value,required"`
}

func (r SettingFontSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the feature is enabled or disabled.
type SettingFontSettingEditParamsValue string

const (
	SettingFontSettingEditParamsValueOn  SettingFontSettingEditParamsValue = "on"
	SettingFontSettingEditParamsValueOff SettingFontSettingEditParamsValue = "off"
)

func (r SettingFontSettingEditParamsValue) IsKnown() bool {
	switch r {
	case SettingFontSettingEditParamsValueOn, SettingFontSettingEditParamsValueOff:
		return true
	}
	return false
}

type SettingFontSettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result FontSettings                               `json:"result"`
	JSON   settingFontSettingEditResponseEnvelopeJSON `json:"-"`
}

// settingFontSettingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingFontSettingEditResponseEnvelope]
type settingFontSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingFontSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingFontSettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingFontSettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result FontSettings                              `json:"result"`
	JSON   settingFontSettingGetResponseEnvelopeJSON `json:"-"`
}

// settingFontSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingFontSettingGetResponseEnvelope]
type settingFontSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingFontSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
