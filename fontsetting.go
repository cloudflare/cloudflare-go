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
	var env FontSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/fonts", zoneIdentifier)
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
func (r *FontSettingService) Update(ctx context.Context, zoneIdentifier string, body FontSettingUpdateParams, opts ...option.RequestOption) (res *FontSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FontSettingUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/fonts", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type FontSettingGetResponse struct {
	// ID of the zone setting.
	ID FontSettingGetResponseID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable FontSettingGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled or disabled.
	Value FontSettingGetResponseValue `json:"value"`
	JSON  fontSettingGetResponseJSON  `json:"-"`
}

// fontSettingGetResponseJSON contains the JSON metadata for the struct
// [FontSettingGetResponse]
type fontSettingGetResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type FontSettingGetResponseID string

const (
	FontSettingGetResponseIDFonts FontSettingGetResponseID = "fonts"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type FontSettingGetResponseEditable bool

const (
	FontSettingGetResponseEditableTrue  FontSettingGetResponseEditable = true
	FontSettingGetResponseEditableFalse FontSettingGetResponseEditable = false
)

// Whether the feature is enabled or disabled.
type FontSettingGetResponseValue string

const (
	FontSettingGetResponseValueOn  FontSettingGetResponseValue = "on"
	FontSettingGetResponseValueOff FontSettingGetResponseValue = "off"
)

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type FontSettingUpdateResponse struct {
	// ID of the zone setting.
	ID FontSettingUpdateResponseID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable FontSettingUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled or disabled.
	Value FontSettingUpdateResponseValue `json:"value"`
	JSON  fontSettingUpdateResponseJSON  `json:"-"`
}

// fontSettingUpdateResponseJSON contains the JSON metadata for the struct
// [FontSettingUpdateResponse]
type fontSettingUpdateResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type FontSettingUpdateResponseID string

const (
	FontSettingUpdateResponseIDFonts FontSettingUpdateResponseID = "fonts"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type FontSettingUpdateResponseEditable bool

const (
	FontSettingUpdateResponseEditableTrue  FontSettingUpdateResponseEditable = true
	FontSettingUpdateResponseEditableFalse FontSettingUpdateResponseEditable = false
)

// Whether the feature is enabled or disabled.
type FontSettingUpdateResponseValue string

const (
	FontSettingUpdateResponseValueOn  FontSettingUpdateResponseValue = "on"
	FontSettingUpdateResponseValueOff FontSettingUpdateResponseValue = "off"
)

type FontSettingGetResponseEnvelope struct {
	Errors   []FontSettingGetResponseEnvelopeErrors   `json:"errors"`
	Messages []FontSettingGetResponseEnvelopeMessages `json:"messages"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result FontSettingGetResponse `json:"result"`
	// Whether the API call was successful
	Success bool                               `json:"success"`
	JSON    fontSettingGetResponseEnvelopeJSON `json:"-"`
}

// fontSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FontSettingGetResponseEnvelope]
type fontSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FontSettingGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    fontSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// fontSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FontSettingGetResponseEnvelopeErrors]
type fontSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FontSettingGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    fontSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// fontSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FontSettingGetResponseEnvelopeMessages]
type fontSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type FontSettingUpdateResponseEnvelope struct {
	Errors   []FontSettingUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []FontSettingUpdateResponseEnvelopeMessages `json:"messages"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result FontSettingUpdateResponse `json:"result"`
	// Whether the API call was successful
	Success bool                                  `json:"success"`
	JSON    fontSettingUpdateResponseEnvelopeJSON `json:"-"`
}

// fontSettingUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [FontSettingUpdateResponseEnvelope]
type fontSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FontSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    fontSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// fontSettingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FontSettingUpdateResponseEnvelopeErrors]
type fontSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FontSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    fontSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// fontSettingUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FontSettingUpdateResponseEnvelopeMessages]
type fontSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FontSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
