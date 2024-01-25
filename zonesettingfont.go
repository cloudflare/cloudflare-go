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

// ZoneSettingFontService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingFontService] method
// instead.
type ZoneSettingFontService struct {
	Options []option.RequestOption
}

// NewZoneSettingFontService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingFontService(opts ...option.RequestOption) (r *ZoneSettingFontService) {
	r = &ZoneSettingFontService{}
	r.Options = opts
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *ZoneSettingFontService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingFontGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/fonts", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *ZoneSettingFontService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingFontUpdateParams, opts ...option.RequestOption) (res *ZoneSettingFontUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/fonts", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type CloudflareFonts struct {
	// ID of the zone setting.
	ID CloudflareFontsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable CloudflareFontsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled or disabled.
	Value CloudflareFontsValue `json:"value"`
	JSON  cloudflareFontsJSON  `json:"-"`
}

// cloudflareFontsJSON contains the JSON metadata for the struct [CloudflareFonts]
type cloudflareFontsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CloudflareFonts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CloudflareFontsID string

const (
	CloudflareFontsIDFonts CloudflareFontsID = "fonts"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type CloudflareFontsEditable bool

const (
	CloudflareFontsEditableTrue  CloudflareFontsEditable = true
	CloudflareFontsEditableFalse CloudflareFontsEditable = false
)

// Whether the feature is enabled or disabled.
type CloudflareFontsValue string

const (
	CloudflareFontsValueOn  CloudflareFontsValue = "on"
	CloudflareFontsValueOff CloudflareFontsValue = "off"
)

type ZoneSettingFontGetResponse struct {
	Errors   []ZoneSettingFontGetResponseError   `json:"errors"`
	Messages []ZoneSettingFontGetResponseMessage `json:"messages"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result CloudflareFonts `json:"result"`
	// Whether the API call was successful
	Success bool                           `json:"success"`
	JSON    zoneSettingFontGetResponseJSON `json:"-"`
}

// zoneSettingFontGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingFontGetResponse]
type zoneSettingFontGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingFontGetResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneSettingFontGetResponseErrorJSON `json:"-"`
}

// zoneSettingFontGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingFontGetResponseError]
type zoneSettingFontGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingFontGetResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSettingFontGetResponseMessageJSON `json:"-"`
}

// zoneSettingFontGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingFontGetResponseMessage]
type zoneSettingFontGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingFontUpdateResponse struct {
	Errors   []ZoneSettingFontUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingFontUpdateResponseMessage `json:"messages"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result CloudflareFonts `json:"result"`
	// Whether the API call was successful
	Success bool                              `json:"success"`
	JSON    zoneSettingFontUpdateResponseJSON `json:"-"`
}

// zoneSettingFontUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingFontUpdateResponse]
type zoneSettingFontUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingFontUpdateResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingFontUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingFontUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingFontUpdateResponseError]
type zoneSettingFontUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingFontUpdateResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingFontUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingFontUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingFontUpdateResponseMessage]
type zoneSettingFontUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingFontUpdateParams struct {
	// Whether the feature is enabled or disabled.
	Value param.Field[ZoneSettingFontUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingFontUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the feature is enabled or disabled.
type ZoneSettingFontUpdateParamsValue string

const (
	ZoneSettingFontUpdateParamsValueOn  ZoneSettingFontUpdateParamsValue = "on"
	ZoneSettingFontUpdateParamsValueOff ZoneSettingFontUpdateParamsValue = "off"
)
