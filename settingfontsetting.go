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
func (r *SettingFontSettingService) Update(ctx context.Context, zoneID string, body SettingFontSettingUpdateParams, opts ...option.RequestOption) (res *SettingFontSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingFontSettingUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/fonts", zoneID)
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
func (r *SettingFontSettingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingFontSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingFontSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/fonts", zoneID)
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
type SettingFontSettingUpdateResponse struct {
	// ID of the zone setting.
	ID SettingFontSettingUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingFontSettingUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingFontSettingUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingFontSettingUpdateResponseJSON `json:"-"`
}

// settingFontSettingUpdateResponseJSON contains the JSON metadata for the struct
// [SettingFontSettingUpdateResponse]
type settingFontSettingUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingFontSettingUpdateResponseID string

const (
	SettingFontSettingUpdateResponseIDFonts SettingFontSettingUpdateResponseID = "fonts"
)

// Current value of the zone setting.
type SettingFontSettingUpdateResponseValue string

const (
	SettingFontSettingUpdateResponseValueOn  SettingFontSettingUpdateResponseValue = "on"
	SettingFontSettingUpdateResponseValueOff SettingFontSettingUpdateResponseValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingFontSettingUpdateResponseEditable bool

const (
	SettingFontSettingUpdateResponseEditableTrue  SettingFontSettingUpdateResponseEditable = true
	SettingFontSettingUpdateResponseEditableFalse SettingFontSettingUpdateResponseEditable = false
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

type SettingFontSettingUpdateParams struct {
	// Whether the feature is enabled or disabled.
	Value param.Field[SettingFontSettingUpdateParamsValue] `json:"value,required"`
}

func (r SettingFontSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the feature is enabled or disabled.
type SettingFontSettingUpdateParamsValue string

const (
	SettingFontSettingUpdateParamsValueOn  SettingFontSettingUpdateParamsValue = "on"
	SettingFontSettingUpdateParamsValueOff SettingFontSettingUpdateParamsValue = "off"
)

type SettingFontSettingUpdateResponseEnvelope struct {
	Errors   []SettingFontSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingFontSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result SettingFontSettingUpdateResponse             `json:"result"`
	JSON   settingFontSettingUpdateResponseEnvelopeJSON `json:"-"`
}

// settingFontSettingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingFontSettingUpdateResponseEnvelope]
type settingFontSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingFontSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    settingFontSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingFontSettingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingFontSettingUpdateResponseEnvelopeErrors]
type settingFontSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingFontSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingFontSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingFontSettingUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingFontSettingUpdateResponseEnvelopeMessages]
type settingFontSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingFontSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
