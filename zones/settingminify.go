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

// SettingMinifyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingMinifyService] method
// instead.
type SettingMinifyService struct {
	Options []option.RequestOption
}

// NewSettingMinifyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingMinifyService(opts ...option.RequestOption) (r *SettingMinifyService) {
	r = &SettingMinifyService{}
	r.Options = opts
	return
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
func (r *SettingMinifyService) Edit(ctx context.Context, params SettingMinifyEditParams, opts ...option.RequestOption) (res *ZonesMinify, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinifyEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/minify", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
func (r *SettingMinifyService) Get(ctx context.Context, query SettingMinifyGetParams, opts ...option.RequestOption) (res *ZonesMinify, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinifyGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/minify", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZonesMinify struct {
	// Zone setting identifier.
	ID ZonesMinifyID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesMinifyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time       `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesMinifyJSON `json:"-"`
}

// zonesMinifyJSON contains the JSON metadata for the struct [ZonesMinify]
type zonesMinifyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesMinifyJSON) RawJSON() string {
	return r.raw
}

func (r ZonesMinify) implementsZonesSettingEditResponse() {}

func (r ZonesMinify) implementsZonesSettingGetResponse() {}

// Zone setting identifier.
type ZonesMinifyID string

const (
	ZonesMinifyIDMinify ZonesMinifyID = "minify"
)

func (r ZonesMinifyID) IsKnown() bool {
	switch r {
	case ZonesMinifyIDMinify:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css ZonesMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML ZonesMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   ZonesMinifyValueJs   `json:"js"`
	JSON zonesMinifyValueJSON `json:"-"`
}

// zonesMinifyValueJSON contains the JSON metadata for the struct
// [ZonesMinifyValue]
type zonesMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesMinifyValueJSON) RawJSON() string {
	return r.raw
}

// Automatically minify all CSS files for your website.
type ZonesMinifyValueCss string

const (
	ZonesMinifyValueCssOn  ZonesMinifyValueCss = "on"
	ZonesMinifyValueCssOff ZonesMinifyValueCss = "off"
)

func (r ZonesMinifyValueCss) IsKnown() bool {
	switch r {
	case ZonesMinifyValueCssOn, ZonesMinifyValueCssOff:
		return true
	}
	return false
}

// Automatically minify all HTML files for your website.
type ZonesMinifyValueHTML string

const (
	ZonesMinifyValueHTMLOn  ZonesMinifyValueHTML = "on"
	ZonesMinifyValueHTMLOff ZonesMinifyValueHTML = "off"
)

func (r ZonesMinifyValueHTML) IsKnown() bool {
	switch r {
	case ZonesMinifyValueHTMLOn, ZonesMinifyValueHTMLOff:
		return true
	}
	return false
}

// Automatically minify all JavaScript files for your website.
type ZonesMinifyValueJs string

const (
	ZonesMinifyValueJsOn  ZonesMinifyValueJs = "on"
	ZonesMinifyValueJsOff ZonesMinifyValueJs = "off"
)

func (r ZonesMinifyValueJs) IsKnown() bool {
	switch r {
	case ZonesMinifyValueJsOn, ZonesMinifyValueJsOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesMinifyEditable bool

const (
	ZonesMinifyEditableTrue  ZonesMinifyEditable = true
	ZonesMinifyEditableFalse ZonesMinifyEditable = false
)

func (r ZonesMinifyEditable) IsKnown() bool {
	switch r {
	case ZonesMinifyEditableTrue, ZonesMinifyEditableFalse:
		return true
	}
	return false
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZonesMinifyParam struct {
	// Zone setting identifier.
	ID param.Field[ZonesMinifyID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesMinifyValueParam] `json:"value,required"`
}

func (r ZonesMinifyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesMinifyParam) implementsZonesSettingEditParamsItem() {}

// Current value of the zone setting.
type ZonesMinifyValueParam struct {
	// Automatically minify all CSS files for your website.
	Css param.Field[ZonesMinifyValueCss] `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML param.Field[ZonesMinifyValueHTML] `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js param.Field[ZonesMinifyValueJs] `json:"js"`
}

func (r ZonesMinifyValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingMinifyEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingMinifyEditParamsValue] `json:"value,required"`
}

func (r SettingMinifyEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMinifyEditParamsValue struct {
	// Automatically minify all CSS files for your website.
	Css param.Field[SettingMinifyEditParamsValueCss] `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML param.Field[SettingMinifyEditParamsValueHTML] `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js param.Field[SettingMinifyEditParamsValueJs] `json:"js"`
}

func (r SettingMinifyEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Automatically minify all CSS files for your website.
type SettingMinifyEditParamsValueCss string

const (
	SettingMinifyEditParamsValueCssOn  SettingMinifyEditParamsValueCss = "on"
	SettingMinifyEditParamsValueCssOff SettingMinifyEditParamsValueCss = "off"
)

func (r SettingMinifyEditParamsValueCss) IsKnown() bool {
	switch r {
	case SettingMinifyEditParamsValueCssOn, SettingMinifyEditParamsValueCssOff:
		return true
	}
	return false
}

// Automatically minify all HTML files for your website.
type SettingMinifyEditParamsValueHTML string

const (
	SettingMinifyEditParamsValueHTMLOn  SettingMinifyEditParamsValueHTML = "on"
	SettingMinifyEditParamsValueHTMLOff SettingMinifyEditParamsValueHTML = "off"
)

func (r SettingMinifyEditParamsValueHTML) IsKnown() bool {
	switch r {
	case SettingMinifyEditParamsValueHTMLOn, SettingMinifyEditParamsValueHTMLOff:
		return true
	}
	return false
}

// Automatically minify all JavaScript files for your website.
type SettingMinifyEditParamsValueJs string

const (
	SettingMinifyEditParamsValueJsOn  SettingMinifyEditParamsValueJs = "on"
	SettingMinifyEditParamsValueJsOff SettingMinifyEditParamsValueJs = "off"
)

func (r SettingMinifyEditParamsValueJs) IsKnown() bool {
	switch r {
	case SettingMinifyEditParamsValueJsOn, SettingMinifyEditParamsValueJsOff:
		return true
	}
	return false
}

type SettingMinifyEditResponseEnvelope struct {
	Errors   []SettingMinifyEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMinifyEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result ZonesMinify                           `json:"result"`
	JSON   settingMinifyEditResponseEnvelopeJSON `json:"-"`
}

// settingMinifyEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingMinifyEditResponseEnvelope]
type settingMinifyEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinifyEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingMinifyEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingMinifyEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMinifyEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingMinifyEditResponseEnvelopeErrors]
type settingMinifyEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinifyEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingMinifyEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingMinifyEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMinifyEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingMinifyEditResponseEnvelopeMessages]
type settingMinifyEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinifyEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingMinifyGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMinifyGetResponseEnvelope struct {
	Errors   []SettingMinifyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMinifyGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result ZonesMinify                          `json:"result"`
	JSON   settingMinifyGetResponseEnvelopeJSON `json:"-"`
}

// settingMinifyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingMinifyGetResponseEnvelope]
type settingMinifyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinifyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingMinifyGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingMinifyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMinifyGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingMinifyGetResponseEnvelopeErrors]
type settingMinifyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinifyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingMinifyGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingMinifyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMinifyGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingMinifyGetResponseEnvelopeMessages]
type settingMinifyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingMinifyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
