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
func (r *SettingMinifyService) Edit(ctx context.Context, params SettingMinifyEditParams, opts ...option.RequestOption) (res *ZoneSettingMinify, err error) {
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
func (r *SettingMinifyService) Get(ctx context.Context, query SettingMinifyGetParams, opts ...option.RequestOption) (res *ZoneSettingMinify, err error) {
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
type ZoneSettingMinify struct {
	// Zone setting identifier.
	ID ZoneSettingMinifyID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingMinifyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingMinifyJSON `json:"-"`
}

// zoneSettingMinifyJSON contains the JSON metadata for the struct
// [ZoneSettingMinify]
type zoneSettingMinifyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinifyJSON) RawJSON() string {
	return r.raw
}

// Zone setting identifier.
type ZoneSettingMinifyID string

const (
	ZoneSettingMinifyIDMinify ZoneSettingMinifyID = "minify"
)

func (r ZoneSettingMinifyID) IsKnown() bool {
	switch r {
	case ZoneSettingMinifyIDMinify:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css ZoneSettingMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML ZoneSettingMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   ZoneSettingMinifyValueJs   `json:"js"`
	JSON zoneSettingMinifyValueJSON `json:"-"`
}

// zoneSettingMinifyValueJSON contains the JSON metadata for the struct
// [ZoneSettingMinifyValue]
type zoneSettingMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingMinifyValueJSON) RawJSON() string {
	return r.raw
}

// Automatically minify all CSS files for your website.
type ZoneSettingMinifyValueCss string

const (
	ZoneSettingMinifyValueCssOn  ZoneSettingMinifyValueCss = "on"
	ZoneSettingMinifyValueCssOff ZoneSettingMinifyValueCss = "off"
)

func (r ZoneSettingMinifyValueCss) IsKnown() bool {
	switch r {
	case ZoneSettingMinifyValueCssOn, ZoneSettingMinifyValueCssOff:
		return true
	}
	return false
}

// Automatically minify all HTML files for your website.
type ZoneSettingMinifyValueHTML string

const (
	ZoneSettingMinifyValueHTMLOn  ZoneSettingMinifyValueHTML = "on"
	ZoneSettingMinifyValueHTMLOff ZoneSettingMinifyValueHTML = "off"
)

func (r ZoneSettingMinifyValueHTML) IsKnown() bool {
	switch r {
	case ZoneSettingMinifyValueHTMLOn, ZoneSettingMinifyValueHTMLOff:
		return true
	}
	return false
}

// Automatically minify all JavaScript files for your website.
type ZoneSettingMinifyValueJs string

const (
	ZoneSettingMinifyValueJsOn  ZoneSettingMinifyValueJs = "on"
	ZoneSettingMinifyValueJsOff ZoneSettingMinifyValueJs = "off"
)

func (r ZoneSettingMinifyValueJs) IsKnown() bool {
	switch r {
	case ZoneSettingMinifyValueJsOn, ZoneSettingMinifyValueJsOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMinifyEditable bool

const (
	ZoneSettingMinifyEditableTrue  ZoneSettingMinifyEditable = true
	ZoneSettingMinifyEditableFalse ZoneSettingMinifyEditable = false
)

func (r ZoneSettingMinifyEditable) IsKnown() bool {
	switch r {
	case ZoneSettingMinifyEditableTrue, ZoneSettingMinifyEditableFalse:
		return true
	}
	return false
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result ZoneSettingMinify                     `json:"result"`
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

type SettingMinifyGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingMinifyGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result ZoneSettingMinify                    `json:"result"`
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
