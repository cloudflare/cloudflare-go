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

// SettingFontSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingFontSettingService] method
// instead.
type SettingFontSettingService struct {
	Options []option.RequestOption
}

// NewSettingFontSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingFontSettingService(opts ...option.RequestOption) (r *SettingFontSettingService) {
	r = &SettingFontSettingService{}
	r.Options = opts
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *SettingFontSettingService) Edit(ctx context.Context, params SettingFontSettingEditParams, opts ...option.RequestOption) (res *SettingFontSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingFontSettingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/fonts", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *SettingFontSettingService) Get(ctx context.Context, query SettingFontSettingGetParams, opts ...option.RequestOption) (res *SettingFontSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingFontSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/fonts", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type SettingFontSettingEditResponse struct {
	// ID of the zone setting.
	ID SettingFontSettingEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingFontSettingEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingFontSettingEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingFontSettingEditResponseJSON `json:"-"`
}

// settingFontSettingEditResponseJSON contains the JSON metadata for the struct
// [SettingFontSettingEditResponse]
type settingFontSettingEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingFontSettingEditResponseID string

const (
	SettingFontSettingEditResponseIDFonts SettingFontSettingEditResponseID = "fonts"
)

// Current value of the zone setting.
type SettingFontSettingEditResponseValue string

const (
	SettingFontSettingEditResponseValueOn  SettingFontSettingEditResponseValue = "on"
	SettingFontSettingEditResponseValueOff SettingFontSettingEditResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingFontSettingEditResponseEditable bool

const (
	SettingFontSettingEditResponseEditableTrue  SettingFontSettingEditResponseEditable = true
	SettingFontSettingEditResponseEditableFalse SettingFontSettingEditResponseEditable = false
)

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type SettingFontSettingGetResponse struct {
	// ID of the zone setting.
	ID SettingFontSettingGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingFontSettingGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingFontSettingGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingFontSettingGetResponseJSON `json:"-"`
}

// settingFontSettingGetResponseJSON contains the JSON metadata for the struct
// [SettingFontSettingGetResponse]
type settingFontSettingGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingFontSettingGetResponseID string

const (
	SettingFontSettingGetResponseIDFonts SettingFontSettingGetResponseID = "fonts"
)

// Current value of the zone setting.
type SettingFontSettingGetResponseValue string

const (
	SettingFontSettingGetResponseValueOn  SettingFontSettingGetResponseValue = "on"
	SettingFontSettingGetResponseValueOff SettingFontSettingGetResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingFontSettingGetResponseEditable bool

const (
	SettingFontSettingGetResponseEditableTrue  SettingFontSettingGetResponseEditable = true
	SettingFontSettingGetResponseEditableFalse SettingFontSettingGetResponseEditable = false
)

type SettingFontSettingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Whether the feature is enabled or disabled.
	Value param.Field[SettingFontSettingEditParamsValue] `json:"value,required"`
}

func (r SettingFontSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the feature is enabled or disabled.
type SettingFontSettingEditParamsValue string

const (
	SettingFontSettingEditParamsValueOn  SettingFontSettingEditParamsValue = "on"
	SettingFontSettingEditParamsValueOff SettingFontSettingEditParamsValue = "off"
)

type SettingFontSettingEditResponseEnvelope struct {
	Errors   []SettingFontSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingFontSettingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result SettingFontSettingEditResponse             `json:"result"`
	JSON   settingFontSettingEditResponseEnvelopeJSON `json:"-"`
}

// settingFontSettingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingFontSettingEditResponseEnvelope]
type settingFontSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingFontSettingEditResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    settingFontSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingFontSettingEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingFontSettingEditResponseEnvelopeErrors]
type settingFontSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingFontSettingEditResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingFontSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingFontSettingEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingFontSettingEditResponseEnvelopeMessages]
type settingFontSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingFontSettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingFontSettingGetResponseEnvelope struct {
	Errors   []SettingFontSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingFontSettingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result SettingFontSettingGetResponse             `json:"result"`
	JSON   settingFontSettingGetResponseEnvelopeJSON `json:"-"`
}

// settingFontSettingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingFontSettingGetResponseEnvelope]
type settingFontSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingFontSettingGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingFontSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingFontSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingFontSettingGetResponseEnvelopeErrors]
type settingFontSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingFontSettingGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    settingFontSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingFontSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingFontSettingGetResponseEnvelopeMessages]
type settingFontSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
