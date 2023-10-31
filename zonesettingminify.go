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
func (r *ZoneSettingMinifyService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingMinifyUpdateParams, opts ...option.RequestOption) (res *ZoneSettingMinifyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/minify", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
func (r *ZoneSettingMinifyService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingMinifyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/minify", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingMinifyUpdateResponse struct {
	Errors   []ZoneSettingMinifyUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingMinifyUpdateResponseMessage `json:"messages"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result ZoneSettingMinifyUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingMinifyUpdateResponseJSON
}

// zoneSettingMinifyUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingMinifyUpdateResponse]
type zoneSettingMinifyUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinifyUpdateResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingMinifyUpdateResponseErrorJSON
}

// zoneSettingMinifyUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingMinifyUpdateResponseError]
type zoneSettingMinifyUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinifyUpdateResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingMinifyUpdateResponseMessageJSON
}

// zoneSettingMinifyUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingMinifyUpdateResponseMessage]
type zoneSettingMinifyUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingMinifyUpdateResponseResult struct {
	// Zone setting identifier.
	ID ZoneSettingMinifyUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMinifyUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingMinifyUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingMinifyUpdateResponseResultJSON
}

// zoneSettingMinifyUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingMinifyUpdateResponseResult]
type zoneSettingMinifyUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type ZoneSettingMinifyUpdateResponseResultID string

const (
	ZoneSettingMinifyUpdateResponseResultIDMinify ZoneSettingMinifyUpdateResponseResultID = "minify"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMinifyUpdateResponseResultEditable bool

const (
	ZoneSettingMinifyUpdateResponseResultEditableTrue  ZoneSettingMinifyUpdateResponseResultEditable = true
	ZoneSettingMinifyUpdateResponseResultEditableFalse ZoneSettingMinifyUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingMinifyUpdateResponseResultValue struct {
	// Automatically minify all CSS files for your website.
	Css ZoneSettingMinifyUpdateResponseResultValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML ZoneSettingMinifyUpdateResponseResultValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   ZoneSettingMinifyUpdateResponseResultValueJs `json:"js"`
	JSON zoneSettingMinifyUpdateResponseResultValueJSON
}

// zoneSettingMinifyUpdateResponseResultValueJSON contains the JSON metadata for
// the struct [ZoneSettingMinifyUpdateResponseResultValue]
type zoneSettingMinifyUpdateResponseResultValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyUpdateResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingMinifyUpdateResponseResultValueCss string

const (
	ZoneSettingMinifyUpdateResponseResultValueCssOn  ZoneSettingMinifyUpdateResponseResultValueCss = "on"
	ZoneSettingMinifyUpdateResponseResultValueCssOff ZoneSettingMinifyUpdateResponseResultValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingMinifyUpdateResponseResultValueHTML string

const (
	ZoneSettingMinifyUpdateResponseResultValueHTMLOn  ZoneSettingMinifyUpdateResponseResultValueHTML = "on"
	ZoneSettingMinifyUpdateResponseResultValueHTMLOff ZoneSettingMinifyUpdateResponseResultValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingMinifyUpdateResponseResultValueJs string

const (
	ZoneSettingMinifyUpdateResponseResultValueJsOn  ZoneSettingMinifyUpdateResponseResultValueJs = "on"
	ZoneSettingMinifyUpdateResponseResultValueJsOff ZoneSettingMinifyUpdateResponseResultValueJs = "off"
)

type ZoneSettingMinifyListResponse struct {
	Errors   []ZoneSettingMinifyListResponseError   `json:"errors"`
	Messages []ZoneSettingMinifyListResponseMessage `json:"messages"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result ZoneSettingMinifyListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool `json:"success"`
	JSON    zoneSettingMinifyListResponseJSON
}

// zoneSettingMinifyListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingMinifyListResponse]
type zoneSettingMinifyListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinifyListResponseError struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingMinifyListResponseErrorJSON
}

// zoneSettingMinifyListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingMinifyListResponseError]
type zoneSettingMinifyListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingMinifyListResponseMessage struct {
	Code    int64  `json:"code,required"`
	Message string `json:"message,required"`
	JSON    zoneSettingMinifyListResponseMessageJSON
}

// zoneSettingMinifyListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingMinifyListResponseMessage]
type zoneSettingMinifyListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingMinifyListResponseResult struct {
	// Zone setting identifier.
	ID ZoneSettingMinifyListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingMinifyListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingMinifyListResponseResultValue `json:"value"`
	JSON  zoneSettingMinifyListResponseResultJSON
}

// zoneSettingMinifyListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingMinifyListResponseResult]
type zoneSettingMinifyListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Zone setting identifier.
type ZoneSettingMinifyListResponseResultID string

const (
	ZoneSettingMinifyListResponseResultIDMinify ZoneSettingMinifyListResponseResultID = "minify"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingMinifyListResponseResultEditable bool

const (
	ZoneSettingMinifyListResponseResultEditableTrue  ZoneSettingMinifyListResponseResultEditable = true
	ZoneSettingMinifyListResponseResultEditableFalse ZoneSettingMinifyListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingMinifyListResponseResultValue struct {
	// Automatically minify all CSS files for your website.
	Css ZoneSettingMinifyListResponseResultValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML ZoneSettingMinifyListResponseResultValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   ZoneSettingMinifyListResponseResultValueJs `json:"js"`
	JSON zoneSettingMinifyListResponseResultValueJSON
}

// zoneSettingMinifyListResponseResultValueJSON contains the JSON metadata for the
// struct [ZoneSettingMinifyListResponseResultValue]
type zoneSettingMinifyListResponseResultValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingMinifyListResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingMinifyListResponseResultValueCss string

const (
	ZoneSettingMinifyListResponseResultValueCssOn  ZoneSettingMinifyListResponseResultValueCss = "on"
	ZoneSettingMinifyListResponseResultValueCssOff ZoneSettingMinifyListResponseResultValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingMinifyListResponseResultValueHTML string

const (
	ZoneSettingMinifyListResponseResultValueHTMLOn  ZoneSettingMinifyListResponseResultValueHTML = "on"
	ZoneSettingMinifyListResponseResultValueHTMLOff ZoneSettingMinifyListResponseResultValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingMinifyListResponseResultValueJs string

const (
	ZoneSettingMinifyListResponseResultValueJsOn  ZoneSettingMinifyListResponseResultValueJs = "on"
	ZoneSettingMinifyListResponseResultValueJsOff ZoneSettingMinifyListResponseResultValueJs = "off"
)

type ZoneSettingMinifyUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingMinifyUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingMinifyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingMinifyUpdateParamsValue struct {
	// Automatically minify all CSS files for your website.
	Css param.Field[ZoneSettingMinifyUpdateParamsValueCss] `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML param.Field[ZoneSettingMinifyUpdateParamsValueHTML] `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js param.Field[ZoneSettingMinifyUpdateParamsValueJs] `json:"js"`
}

func (r ZoneSettingMinifyUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingMinifyUpdateParamsValueCss string

const (
	ZoneSettingMinifyUpdateParamsValueCssOn  ZoneSettingMinifyUpdateParamsValueCss = "on"
	ZoneSettingMinifyUpdateParamsValueCssOff ZoneSettingMinifyUpdateParamsValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingMinifyUpdateParamsValueHTML string

const (
	ZoneSettingMinifyUpdateParamsValueHTMLOn  ZoneSettingMinifyUpdateParamsValueHTML = "on"
	ZoneSettingMinifyUpdateParamsValueHTMLOff ZoneSettingMinifyUpdateParamsValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingMinifyUpdateParamsValueJs string

const (
	ZoneSettingMinifyUpdateParamsValueJsOn  ZoneSettingMinifyUpdateParamsValueJs = "on"
	ZoneSettingMinifyUpdateParamsValueJsOff ZoneSettingMinifyUpdateParamsValueJs = "off"
)
