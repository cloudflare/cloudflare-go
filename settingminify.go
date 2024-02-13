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
func (r *SettingMinifyService) Update(ctx context.Context, zoneID string, body SettingMinifyUpdateParams, opts ...option.RequestOption) (res *SettingMinifyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingMinifyUpdateResponseEnvelope
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
type SettingMinifyUpdateResponse struct {
	// Zone setting identifier.
	ID SettingMinifyUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingMinifyUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingMinifyUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingMinifyUpdateResponseJSON `json:"-"`
}

// settingMinifyUpdateResponseJSON contains the JSON metadata for the struct
// [SettingMinifyUpdateResponse]
type settingMinifyUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type SettingMinifyUpdateResponseID string

const (
	SettingMinifyUpdateResponseIDMinify SettingMinifyUpdateResponseID = "minify"
)

// Current value of the zone setting.
type SettingMinifyUpdateResponseValue struct {
	// Automatically minify all CSS files for your website.
	Css SettingMinifyUpdateResponseValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML SettingMinifyUpdateResponseValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   SettingMinifyUpdateResponseValueJs   `json:"js"`
	JSON settingMinifyUpdateResponseValueJSON `json:"-"`
}

// settingMinifyUpdateResponseValueJSON contains the JSON metadata for the struct
// [SettingMinifyUpdateResponseValue]
type settingMinifyUpdateResponseValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyUpdateResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type SettingMinifyUpdateResponseValueCss string

const (
	SettingMinifyUpdateResponseValueCssOn  SettingMinifyUpdateResponseValueCss = "on"
	SettingMinifyUpdateResponseValueCssOff SettingMinifyUpdateResponseValueCss = "off"
)

// Automatically minify all HTML files for your website.
type SettingMinifyUpdateResponseValueHTML string

const (
	SettingMinifyUpdateResponseValueHTMLOn  SettingMinifyUpdateResponseValueHTML = "on"
	SettingMinifyUpdateResponseValueHTMLOff SettingMinifyUpdateResponseValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type SettingMinifyUpdateResponseValueJs string

const (
	SettingMinifyUpdateResponseValueJsOn  SettingMinifyUpdateResponseValueJs = "on"
	SettingMinifyUpdateResponseValueJsOff SettingMinifyUpdateResponseValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingMinifyUpdateResponseEditable bool

const (
	SettingMinifyUpdateResponseEditableTrue  SettingMinifyUpdateResponseEditable = true
	SettingMinifyUpdateResponseEditableFalse SettingMinifyUpdateResponseEditable = false
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

type SettingMinifyUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[SettingMinifyUpdateParamsValue] `json:"value,required"`
}

func (r SettingMinifyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingMinifyUpdateParamsValue struct {
	// Automatically minify all CSS files for your website.
	Css param.Field[SettingMinifyUpdateParamsValueCss] `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML param.Field[SettingMinifyUpdateParamsValueHTML] `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js param.Field[SettingMinifyUpdateParamsValueJs] `json:"js"`
}

func (r SettingMinifyUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Automatically minify all CSS files for your website.
type SettingMinifyUpdateParamsValueCss string

const (
	SettingMinifyUpdateParamsValueCssOn  SettingMinifyUpdateParamsValueCss = "on"
	SettingMinifyUpdateParamsValueCssOff SettingMinifyUpdateParamsValueCss = "off"
)

// Automatically minify all HTML files for your website.
type SettingMinifyUpdateParamsValueHTML string

const (
	SettingMinifyUpdateParamsValueHTMLOn  SettingMinifyUpdateParamsValueHTML = "on"
	SettingMinifyUpdateParamsValueHTMLOff SettingMinifyUpdateParamsValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type SettingMinifyUpdateParamsValueJs string

const (
	SettingMinifyUpdateParamsValueJsOn  SettingMinifyUpdateParamsValueJs = "on"
	SettingMinifyUpdateParamsValueJsOff SettingMinifyUpdateParamsValueJs = "off"
)

type SettingMinifyUpdateResponseEnvelope struct {
	Errors   []SettingMinifyUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingMinifyUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result SettingMinifyUpdateResponse             `json:"result"`
	JSON   settingMinifyUpdateResponseEnvelopeJSON `json:"-"`
}

// settingMinifyUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingMinifyUpdateResponseEnvelope]
type settingMinifyUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMinifyUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingMinifyUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingMinifyUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingMinifyUpdateResponseEnvelopeErrors]
type settingMinifyUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingMinifyUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingMinifyUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingMinifyUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingMinifyUpdateResponseEnvelopeMessages]
type settingMinifyUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingMinifyUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
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
