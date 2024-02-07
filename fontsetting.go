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

// FontSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFontSettingService] method
// instead.
type FontSettingService struct {
	Options []option.RequestOption
}

// NewFontSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFontSettingService(opts ...option.RequestOption) (r *FontSettingService) {
	r = &FontSettingService{}
	r.Options = opts
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *FontSettingService) Get(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *FontSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/fonts", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *FontSettingService) Update(ctx context.Context, zoneIdentifier string, body FontSettingUpdateParams, opts ...option.RequestOption) (res *FontSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/fonts", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type FontSettingGetResponse struct {
	Errors   []FontSettingGetResponseError   `json:"errors"`
	Messages []FontSettingGetResponseMessage `json:"messages"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result FontSettingGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                       `json:"success"`
	JSON    fontSettingGetResponseJSON `json:"-"`
}

// fontSettingGetResponseJSON contains the JSON metadata for the struct
// [FontSettingGetResponse]
type fontSettingGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FontSettingGetResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    fontSettingGetResponseErrorJSON `json:"-"`
}

// fontSettingGetResponseErrorJSON contains the JSON metadata for the struct
// [FontSettingGetResponseError]
type fontSettingGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FontSettingGetResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    fontSettingGetResponseMessageJSON `json:"-"`
}

// fontSettingGetResponseMessageJSON contains the JSON metadata for the struct
// [FontSettingGetResponseMessage]
type fontSettingGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type FontSettingGetResponseResult struct {
	// ID of the zone setting.
	ID FontSettingGetResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable FontSettingGetResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled or disabled.
	Value FontSettingGetResponseResultValue `json:"value"`
	JSON  fontSettingGetResponseResultJSON  `json:"-"`
}

// fontSettingGetResponseResultJSON contains the JSON metadata for the struct
// [FontSettingGetResponseResult]
type fontSettingGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type FontSettingGetResponseResultID string

const (
	FontSettingGetResponseResultIDFonts FontSettingGetResponseResultID = "fonts"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type FontSettingGetResponseResultEditable bool

const (
	FontSettingGetResponseResultEditableTrue  FontSettingGetResponseResultEditable = true
	FontSettingGetResponseResultEditableFalse FontSettingGetResponseResultEditable = false
)

// Whether the feature is enabled or disabled.
type FontSettingGetResponseResultValue string

const (
	FontSettingGetResponseResultValueOn  FontSettingGetResponseResultValue = "on"
	FontSettingGetResponseResultValueOff FontSettingGetResponseResultValue = "off"
)

type FontSettingUpdateResponse struct {
	Errors   []FontSettingUpdateResponseError   `json:"errors"`
	Messages []FontSettingUpdateResponseMessage `json:"messages"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result FontSettingUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                          `json:"success"`
	JSON    fontSettingUpdateResponseJSON `json:"-"`
}

// fontSettingUpdateResponseJSON contains the JSON metadata for the struct
// [FontSettingUpdateResponse]
type fontSettingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FontSettingUpdateResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    fontSettingUpdateResponseErrorJSON `json:"-"`
}

// fontSettingUpdateResponseErrorJSON contains the JSON metadata for the struct
// [FontSettingUpdateResponseError]
type fontSettingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FontSettingUpdateResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    fontSettingUpdateResponseMessageJSON `json:"-"`
}

// fontSettingUpdateResponseMessageJSON contains the JSON metadata for the struct
// [FontSettingUpdateResponseMessage]
type fontSettingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type FontSettingUpdateResponseResult struct {
	// ID of the zone setting.
	ID FontSettingUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable FontSettingUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled or disabled.
	Value FontSettingUpdateResponseResultValue `json:"value"`
	JSON  fontSettingUpdateResponseResultJSON  `json:"-"`
}

// fontSettingUpdateResponseResultJSON contains the JSON metadata for the struct
// [FontSettingUpdateResponseResult]
type fontSettingUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type FontSettingUpdateResponseResultID string

const (
	FontSettingUpdateResponseResultIDFonts FontSettingUpdateResponseResultID = "fonts"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type FontSettingUpdateResponseResultEditable bool

const (
	FontSettingUpdateResponseResultEditableTrue  FontSettingUpdateResponseResultEditable = true
	FontSettingUpdateResponseResultEditableFalse FontSettingUpdateResponseResultEditable = false
)

// Whether the feature is enabled or disabled.
type FontSettingUpdateResponseResultValue string

const (
	FontSettingUpdateResponseResultValueOn  FontSettingUpdateResponseResultValue = "on"
	FontSettingUpdateResponseResultValueOff FontSettingUpdateResponseResultValue = "off"
)

type FontSettingUpdateParams struct {
	// Whether the feature is enabled or disabled.
	Value param.Field[FontSettingUpdateParamsValue] `json:"value,required"`
}

func (r FontSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the feature is enabled or disabled.
type FontSettingUpdateParamsValue string

const (
	FontSettingUpdateParamsValueOn  FontSettingUpdateParamsValue = "on"
	FontSettingUpdateParamsValueOff FontSettingUpdateParamsValue = "off"
)
