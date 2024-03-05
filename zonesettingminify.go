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

// ZoneSettingMinifyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingMinifyService] method
// instead.
type ZoneSettingMinifyService struct {
	Options []option.RequestOption
}

// NewZoneSettingMinifyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingMinifyService(opts ...option.RequestOption) (r *ZoneSettingMinifyService) {
	r = &ZoneSettingMinifyService{}
	r.Options = opts
	return
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
func (r *ZoneSettingMinifyService) Edit(ctx context.Context, params ZoneSettingMinifyEditParams, opts ...option.RequestOption) (res *ZonesMinify, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingMinifyEditResponseEnvelope
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
func (r *ZoneSettingMinifyService) Get(ctx context.Context, query ZoneSettingMinifyGetParams, opts ...option.RequestOption) (res *ZonesMinify, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingMinifyGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/minify", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

func (r ZonesMinify) implementsZoneSettingEditResponse() {}

func (r ZonesMinify) implementsZoneSettingGetResponse() {}

// Zone setting identifier.
type ZonesMinifyID string

const (
	ZonesMinifyIDMinify ZonesMinifyID = "minify"
)

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

// Automatically minify all CSS files for your website.
type ZonesMinifyValueCss string

const (
	ZonesMinifyValueCssOn  ZonesMinifyValueCss = "on"
	ZonesMinifyValueCssOff ZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZonesMinifyValueHTML string

const (
	ZonesMinifyValueHTMLOn  ZonesMinifyValueHTML = "on"
	ZonesMinifyValueHTMLOff ZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZonesMinifyValueJs string

const (
	ZonesMinifyValueJsOn  ZonesMinifyValueJs = "on"
	ZonesMinifyValueJsOff ZonesMinifyValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesMinifyEditable bool

const (
	ZonesMinifyEditableTrue  ZonesMinifyEditable = true
	ZonesMinifyEditableFalse ZonesMinifyEditable = false
)

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

func (r ZonesMinifyParam) implementsZoneSettingEditParamsItem() {}

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

type ZoneSettingMinifyEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingMinifyEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingMinifyEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingMinifyEditParamsValue struct {
	// Automatically minify all CSS files for your website.
	Css param.Field[ZoneSettingMinifyEditParamsValueCss] `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML param.Field[ZoneSettingMinifyEditParamsValueHTML] `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js param.Field[ZoneSettingMinifyEditParamsValueJs] `json:"js"`
}

func (r ZoneSettingMinifyEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingMinifyEditParamsValueCss string

const (
	ZoneSettingMinifyEditParamsValueCssOn  ZoneSettingMinifyEditParamsValueCss = "on"
	ZoneSettingMinifyEditParamsValueCssOff ZoneSettingMinifyEditParamsValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingMinifyEditParamsValueHTML string

const (
	ZoneSettingMinifyEditParamsValueHTMLOn  ZoneSettingMinifyEditParamsValueHTML = "on"
	ZoneSettingMinifyEditParamsValueHTMLOff ZoneSettingMinifyEditParamsValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingMinifyEditParamsValueJs string

const (
	ZoneSettingMinifyEditParamsValueJsOn  ZoneSettingMinifyEditParamsValueJs = "on"
	ZoneSettingMinifyEditParamsValueJsOff ZoneSettingMinifyEditParamsValueJs = "off"
)

type ZoneSettingMinifyEditResponseEnvelope struct {
	Errors   []ZoneSettingMinifyEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingMinifyEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result ZonesMinify                               `json:"result"`
	JSON   zoneSettingMinifyEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingMinifyEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingMinifyEditResponseEnvelope]
type zoneSettingMinifyEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinifyEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingMinifyEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingMinifyEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingMinifyEditResponseEnvelopeErrors]
type zoneSettingMinifyEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinifyEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingMinifyEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingMinifyEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingMinifyEditResponseEnvelopeMessages]
type zoneSettingMinifyEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinifyGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingMinifyGetResponseEnvelope struct {
	Errors   []ZoneSettingMinifyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingMinifyGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result ZonesMinify                              `json:"result"`
	JSON   zoneSettingMinifyGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingMinifyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingMinifyGetResponseEnvelope]
type zoneSettingMinifyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinifyGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingMinifyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingMinifyGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingMinifyGetResponseEnvelopeErrors]
type zoneSettingMinifyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinifyGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingMinifyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingMinifyGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingMinifyGetResponseEnvelopeMessages]
type zoneSettingMinifyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
