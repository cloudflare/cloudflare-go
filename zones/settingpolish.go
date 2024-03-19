// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *SettingPolishService) Edit(ctx context.Context, params SettingPolishEditParams, opts ...option.RequestOption) (res *ZonesPolish, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPolishEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/polish", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/polish-solving-mobile-speed)
// for more information.
func (r *SettingPolishService) Get(ctx context.Context, query SettingPolishGetParams, opts ...option.RequestOption) (res *ZonesPolish, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingPolishGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/polish", query.ZoneID)
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
type ZonesPolish struct {
	// ID of the zone setting.
	ID ZonesPolishID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesPolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time       `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesPolishJSON `json:"-"`
}

// zonesPolishJSON contains the JSON metadata for the struct [ZonesPolish]
type zonesPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesPolishJSON) RawJSON() string {
	return r.raw
}

func (r ZonesPolish) implementsZonesSettingEditResponse() {}

func (r ZonesPolish) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesPolishID string

const (
	ZonesPolishIDPolish ZonesPolishID = "polish"
)

func (r ZonesPolishID) IsKnown() bool {
	switch r {
	case ZonesPolishIDPolish:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZonesPolishValue string

const (
	ZonesPolishValueOff      ZonesPolishValue = "off"
	ZonesPolishValueLossless ZonesPolishValue = "lossless"
	ZonesPolishValueLossy    ZonesPolishValue = "lossy"
)

func (r ZonesPolishValue) IsKnown() bool {
	switch r {
	case ZonesPolishValueOff, ZonesPolishValueLossless, ZonesPolishValueLossy:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesPolishEditable bool

const (
	ZonesPolishEditableTrue  ZonesPolishEditable = true
	ZonesPolishEditableFalse ZonesPolishEditable = false
)

func (r ZonesPolishEditable) IsKnown() bool {
	switch r {
	case ZonesPolishEditableTrue, ZonesPolishEditableFalse:
		return true
	}
	return false
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZonesPolishParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesPolishID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesPolishValue] `json:"value,required"`
}

func (r ZonesPolishParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesPolishParam) implementsZonesSettingEditParamsItem() {}

type SettingPolishEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Value param.Field[ZonesPolishParam] `json:"value,required"`
}

func (r SettingPolishEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingPolishEditResponseEnvelope struct {
	Errors   []SettingPolishEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingPolishEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result ZonesPolish                           `json:"result"`
	JSON   settingPolishEditResponseEnvelopeJSON `json:"-"`
}

// settingPolishEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingPolishEditResponseEnvelope]
type settingPolishEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPolishEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingPolishEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingPolishEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingPolishEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingPolishEditResponseEnvelopeErrors]
type settingPolishEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPolishEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingPolishEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    settingPolishEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingPolishEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingPolishEditResponseEnvelopeMessages]
type settingPolishEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingPolishEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingPolishEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingPolishGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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
	Result ZonesPolish                          `json:"result"`
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

func (r settingPolishGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r settingPolishGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r settingPolishGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
