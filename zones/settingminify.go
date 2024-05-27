// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// SettingMinifyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingMinifyService] method instead.
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
func (r *SettingMinifyService) Edit(ctx context.Context, params SettingMinifyEditParams, opts ...option.RequestOption) (res *Minify, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinifyEditResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
func (r *SettingMinifyService) Get(ctx context.Context, query SettingMinifyGetParams, opts ...option.RequestOption) (res *Minify, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinifyGetResponseEnvelope
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
type Minify struct {
	// Zone setting identifier.
	ID MinifyID `json:"id,required"`
	// Current value of the zone setting.
	Value MinifyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time  `json:"modified_on,nullable" format:"date-time"`
	JSON       minifyJSON `json:"-"`
}

// minifyJSON contains the JSON metadata for the struct [Minify]
type minifyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Minify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minifyJSON) RawJSON() string {
	return r.raw
}

// Zone setting identifier.
type MinifyID string

const (
	MinifyIDMinify MinifyID = "minify"
)

func (r MinifyID) IsKnown() bool {
	switch r {
	case MinifyIDMinify:
		return true
	}
	return false
}

// Current value of the zone setting.
type MinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css MinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML MinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	JS   MinifyValueJS   `json:"js"`
	JSON minifyValueJSON `json:"-"`
}

// minifyValueJSON contains the JSON metadata for the struct [MinifyValue]
type minifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	JS          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minifyValueJSON) RawJSON() string {
	return r.raw
}

// Automatically minify all CSS files for your website.
type MinifyValueCss string

const (
	MinifyValueCssOn  MinifyValueCss = "on"
	MinifyValueCssOff MinifyValueCss = "off"
)

func (r MinifyValueCss) IsKnown() bool {
	switch r {
	case MinifyValueCssOn, MinifyValueCssOff:
		return true
	}
	return false
}

// Automatically minify all HTML files for your website.
type MinifyValueHTML string

const (
	MinifyValueHTMLOn  MinifyValueHTML = "on"
	MinifyValueHTMLOff MinifyValueHTML = "off"
)

func (r MinifyValueHTML) IsKnown() bool {
	switch r {
	case MinifyValueHTMLOn, MinifyValueHTMLOff:
		return true
	}
	return false
}

// Automatically minify all JavaScript files for your website.
type MinifyValueJS string

const (
	MinifyValueJSOn  MinifyValueJS = "on"
	MinifyValueJSOff MinifyValueJS = "off"
)

func (r MinifyValueJS) IsKnown() bool {
	switch r {
	case MinifyValueJSOn, MinifyValueJSOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type MinifyEditable bool

const (
	MinifyEditableTrue  MinifyEditable = true
	MinifyEditableFalse MinifyEditable = false
)

func (r MinifyEditable) IsKnown() bool {
	switch r {
	case MinifyEditableTrue, MinifyEditableFalse:
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
	JS param.Field[SettingMinifyEditParamsValueJS] `json:"js"`
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
type SettingMinifyEditParamsValueJS string

const (
	SettingMinifyEditParamsValueJSOn  SettingMinifyEditParamsValueJS = "on"
	SettingMinifyEditParamsValueJSOff SettingMinifyEditParamsValueJS = "off"
)

func (r SettingMinifyEditParamsValueJS) IsKnown() bool {
	switch r {
	case SettingMinifyEditParamsValueJSOn, SettingMinifyEditParamsValueJSOff:
		return true
	}
	return false
}

type SettingMinifyEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result Minify                                `json:"result"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result Minify                               `json:"result"`
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
