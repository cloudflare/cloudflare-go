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

// SettingPolishService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingPolishService] method
// instead.
type SettingPolishService struct {
	Options []option.RequestOption
}

// NewSettingPolishService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingPolishService(opts ...option.RequestOption) (r *SettingPolishService) {
	r = &SettingPolishService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/polish-solving-mobile-speed)
// for more information.
func (r *SettingPolishService) Update(ctx context.Context, zoneID string, body SettingPolishUpdateParams, opts ...option.RequestOption) (res *SettingPolishUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPolishUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/polish", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/polish-solving-mobile-speed)
// for more information.
func (r *SettingPolishService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SettingPolishGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPolishGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/polish", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type SettingPolishUpdateResponse struct {
	// ID of the zone setting.
	ID SettingPolishUpdateResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingPolishUpdateResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingPolishUpdateResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingPolishUpdateResponseJSON `json:"-"`
}

// settingPolishUpdateResponseJSON contains the JSON metadata for the struct
// [SettingPolishUpdateResponse]
type settingPolishUpdateResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingPolishUpdateResponseID string

const (
	SettingPolishUpdateResponseIDPolish SettingPolishUpdateResponseID = "polish"
)

// Current value of the zone setting.
type SettingPolishUpdateResponseValue string

const (
	SettingPolishUpdateResponseValueOff      SettingPolishUpdateResponseValue = "off"
	SettingPolishUpdateResponseValueLossless SettingPolishUpdateResponseValue = "lossless"
	SettingPolishUpdateResponseValueLossy    SettingPolishUpdateResponseValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingPolishUpdateResponseEditable bool

const (
	SettingPolishUpdateResponseEditableTrue  SettingPolishUpdateResponseEditable = true
	SettingPolishUpdateResponseEditableFalse SettingPolishUpdateResponseEditable = false
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type SettingPolishGetResponse struct {
	// ID of the zone setting.
	ID SettingPolishGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingPolishGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingPolishGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingPolishGetResponseJSON `json:"-"`
}

// settingPolishGetResponseJSON contains the JSON metadata for the struct
// [SettingPolishGetResponse]
type settingPolishGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type SettingPolishGetResponseID string

const (
	SettingPolishGetResponseIDPolish SettingPolishGetResponseID = "polish"
)

// Current value of the zone setting.
type SettingPolishGetResponseValue string

const (
	SettingPolishGetResponseValueOff      SettingPolishGetResponseValue = "off"
	SettingPolishGetResponseValueLossless SettingPolishGetResponseValue = "lossless"
	SettingPolishGetResponseValueLossy    SettingPolishGetResponseValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingPolishGetResponseEditable bool

const (
	SettingPolishGetResponseEditableTrue  SettingPolishGetResponseEditable = true
	SettingPolishGetResponseEditableFalse SettingPolishGetResponseEditable = false
)

type SettingPolishUpdateParams struct {
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Value param.Field[SettingPolishUpdateParamsValue] `json:"value,required"`
}

func (r SettingPolishUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type SettingPolishUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[SettingPolishUpdateParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingPolishUpdateParamsValueValue] `json:"value,required"`
}

func (r SettingPolishUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type SettingPolishUpdateParamsValueID string

const (
	SettingPolishUpdateParamsValueIDPolish SettingPolishUpdateParamsValueID = "polish"
)

// Current value of the zone setting.
type SettingPolishUpdateParamsValueValue string

const (
	SettingPolishUpdateParamsValueValueOff      SettingPolishUpdateParamsValueValue = "off"
	SettingPolishUpdateParamsValueValueLossless SettingPolishUpdateParamsValueValue = "lossless"
	SettingPolishUpdateParamsValueValueLossy    SettingPolishUpdateParamsValueValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingPolishUpdateParamsValueEditable bool

const (
	SettingPolishUpdateParamsValueEditableTrue  SettingPolishUpdateParamsValueEditable = true
	SettingPolishUpdateParamsValueEditableFalse SettingPolishUpdateParamsValueEditable = false
)

type SettingPolishUpdateResponseEnvelope struct {
	Errors   []SettingPolishUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingPolishUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result SettingPolishUpdateResponse             `json:"result"`
	JSON   settingPolishUpdateResponseEnvelopeJSON `json:"-"`
}

// settingPolishUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingPolishUpdateResponseEnvelope]
type settingPolishUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPolishUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingPolishUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// settingPolishUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingPolishUpdateResponseEnvelopeErrors]
type settingPolishUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPolishUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    settingPolishUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// settingPolishUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingPolishUpdateResponseEnvelopeMessages]
type settingPolishUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPolishGetResponseEnvelope struct {
	Errors   []SettingPolishGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingPolishGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result SettingPolishGetResponse             `json:"result"`
	JSON   settingPolishGetResponseEnvelopeJSON `json:"-"`
}

// settingPolishGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingPolishGetResponseEnvelope]
type settingPolishGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPolishGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingPolishGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingPolishGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingPolishGetResponseEnvelopeErrors]
type settingPolishGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingPolishGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingPolishGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingPolishGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingPolishGetResponseEnvelopeMessages]
type settingPolishGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
