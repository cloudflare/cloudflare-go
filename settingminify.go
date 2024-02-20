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
func (r *SettingMinifyService) Edit(ctx context.Context, zoneID string, body SettingMinifyEditParams, opts ...option.RequestOption) (res *SettingMinifyEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinifyEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/minify", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
func (r *SettingMinifyService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingMinifyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinifyGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/minify", zoneID)
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
type SettingMinifyEditResponse struct {
	// Zone setting identifier.
	ID SettingMinifyEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMinifyEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMinifyEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMinifyEditResponseJSON `json:"-"`
}

// settingMinifyEditResponseJSON contains the JSON metadata for the struct
// [SettingMinifyEditResponse]
type settingMinifyEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type SettingMinifyEditResponseID string

const (
	SettingMinifyEditResponseIDMinify SettingMinifyEditResponseID = "minify"
)

// Current value of the zone setting.
type SettingMinifyEditResponseValue struct {
	// Automatically minify all CSS files for your website.
	Css SettingMinifyEditResponseValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML SettingMinifyEditResponseValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   SettingMinifyEditResponseValueJs   `json:"js"`
	JSON settingMinifyEditResponseValueJSON `json:"-"`
}

// settingMinifyEditResponseValueJSON contains the JSON metadata for the struct
// [SettingMinifyEditResponseValue]
type settingMinifyEditResponseValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type SettingMinifyEditResponseValueCss string

const (
	SettingMinifyEditResponseValueCssOn  SettingMinifyEditResponseValueCss = "on"
	SettingMinifyEditResponseValueCssOff SettingMinifyEditResponseValueCss = "off"
)

// Automatically minify all HTML files for your website.
type SettingMinifyEditResponseValueHTML string

const (
	SettingMinifyEditResponseValueHTMLOn  SettingMinifyEditResponseValueHTML = "on"
	SettingMinifyEditResponseValueHTMLOff SettingMinifyEditResponseValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type SettingMinifyEditResponseValueJs string

const (
	SettingMinifyEditResponseValueJsOn  SettingMinifyEditResponseValueJs = "on"
	SettingMinifyEditResponseValueJsOff SettingMinifyEditResponseValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMinifyEditResponseEditable bool

const (
	SettingMinifyEditResponseEditableTrue  SettingMinifyEditResponseEditable = true
	SettingMinifyEditResponseEditableFalse SettingMinifyEditResponseEditable = false
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type SettingMinifyGetResponse struct {
	// Zone setting identifier.
	ID SettingMinifyGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMinifyGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMinifyGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMinifyGetResponseJSON `json:"-"`
}

// settingMinifyGetResponseJSON contains the JSON metadata for the struct
// [SettingMinifyGetResponse]
type settingMinifyGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type SettingMinifyGetResponseID string

const (
	SettingMinifyGetResponseIDMinify SettingMinifyGetResponseID = "minify"
)

// Current value of the zone setting.
type SettingMinifyGetResponseValue struct {
	// Automatically minify all CSS files for your website.
	Css SettingMinifyGetResponseValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML SettingMinifyGetResponseValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   SettingMinifyGetResponseValueJs   `json:"js"`
	JSON settingMinifyGetResponseValueJSON `json:"-"`
}

// settingMinifyGetResponseValueJSON contains the JSON metadata for the struct
// [SettingMinifyGetResponseValue]
type settingMinifyGetResponseValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type SettingMinifyGetResponseValueCss string

const (
	SettingMinifyGetResponseValueCssOn  SettingMinifyGetResponseValueCss = "on"
	SettingMinifyGetResponseValueCssOff SettingMinifyGetResponseValueCss = "off"
)

// Automatically minify all HTML files for your website.
type SettingMinifyGetResponseValueHTML string

const (
	SettingMinifyGetResponseValueHTMLOn  SettingMinifyGetResponseValueHTML = "on"
	SettingMinifyGetResponseValueHTMLOff SettingMinifyGetResponseValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type SettingMinifyGetResponseValueJs string

const (
	SettingMinifyGetResponseValueJsOn  SettingMinifyGetResponseValueJs = "on"
	SettingMinifyGetResponseValueJsOff SettingMinifyGetResponseValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMinifyGetResponseEditable bool

const (
	SettingMinifyGetResponseEditableTrue  SettingMinifyGetResponseEditable = true
	SettingMinifyGetResponseEditableFalse SettingMinifyGetResponseEditable = false
)

type SettingMinifyEditParams struct {
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

// Automatically minify all HTML files for your website.
type SettingMinifyEditParamsValueHTML string

const (
	SettingMinifyEditParamsValueHTMLOn  SettingMinifyEditParamsValueHTML = "on"
	SettingMinifyEditParamsValueHTMLOff SettingMinifyEditParamsValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type SettingMinifyEditParamsValueJs string

const (
	SettingMinifyEditParamsValueJsOn  SettingMinifyEditParamsValueJs = "on"
	SettingMinifyEditParamsValueJsOff SettingMinifyEditParamsValueJs = "off"
)

type SettingMinifyEditResponseEnvelope struct {
	Errors   []SettingMinifyEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMinifyEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result SettingMinifyEditResponse             `json:"result"`
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

type SettingMinifyGetResponseEnvelope struct {
	Errors   []SettingMinifyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMinifyGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result SettingMinifyGetResponse             `json:"result"`
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
