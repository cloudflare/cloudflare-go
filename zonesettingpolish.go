// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZoneSettingPolishService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingPolishService] method
// instead.
type ZoneSettingPolishService struct {
	Options []option.RequestOption
}

// NewZoneSettingPolishService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingPolishService(opts ...option.RequestOption) (r *ZoneSettingPolishService) {
	r = &ZoneSettingPolishService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/polish-solving-mobile-speed)
// for more information.
func (r *ZoneSettingPolishService) Edit(ctx context.Context, params ZoneSettingPolishEditParams, opts ...option.RequestOption) (res *ZoneSettingPolishEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingPolishEditResponseEnvelope
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
func (r *ZoneSettingPolishService) Get(ctx context.Context, query ZoneSettingPolishGetParams, opts ...option.RequestOption) (res *ZoneSettingPolishGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingPolishGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/polish", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
type ZoneSettingPolishEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingPolishEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingPolishEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPolishEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingPolishEditResponseJSON `json:"-"`
}

// zoneSettingPolishEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingPolishEditResponse]
type zoneSettingPolishEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPolishEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingPolishEditResponseID string

const (
	ZoneSettingPolishEditResponseIDPolish ZoneSettingPolishEditResponseID = "polish"
)

// Current value of the zone setting.
type ZoneSettingPolishEditResponseValue string

const (
	ZoneSettingPolishEditResponseValueOff      ZoneSettingPolishEditResponseValue = "off"
	ZoneSettingPolishEditResponseValueLossless ZoneSettingPolishEditResponseValue = "lossless"
	ZoneSettingPolishEditResponseValueLossy    ZoneSettingPolishEditResponseValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPolishEditResponseEditable bool

const (
	ZoneSettingPolishEditResponseEditableTrue  ZoneSettingPolishEditResponseEditable = true
	ZoneSettingPolishEditResponseEditableFalse ZoneSettingPolishEditResponseEditable = false
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingPolishGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingPolishGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingPolishGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPolishGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingPolishGetResponseJSON `json:"-"`
}

// zoneSettingPolishGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingPolishGetResponse]
type zoneSettingPolishGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPolishGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingPolishGetResponseID string

const (
	ZoneSettingPolishGetResponseIDPolish ZoneSettingPolishGetResponseID = "polish"
)

// Current value of the zone setting.
type ZoneSettingPolishGetResponseValue string

const (
	ZoneSettingPolishGetResponseValueOff      ZoneSettingPolishGetResponseValue = "off"
	ZoneSettingPolishGetResponseValueLossless ZoneSettingPolishGetResponseValue = "lossless"
	ZoneSettingPolishGetResponseValueLossy    ZoneSettingPolishGetResponseValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPolishGetResponseEditable bool

const (
	ZoneSettingPolishGetResponseEditableTrue  ZoneSettingPolishGetResponseEditable = true
	ZoneSettingPolishGetResponseEditableFalse ZoneSettingPolishGetResponseEditable = false
)

type ZoneSettingPolishEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Value param.Field[ZoneSettingPolishEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingPolishEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingPolishEditParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingPolishEditParamsValueID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingPolishEditParamsValueValue] `json:"value,required"`
}

func (r ZoneSettingPolishEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingPolishEditParamsValueID string

const (
	ZoneSettingPolishEditParamsValueIDPolish ZoneSettingPolishEditParamsValueID = "polish"
)

// Current value of the zone setting.
type ZoneSettingPolishEditParamsValueValue string

const (
	ZoneSettingPolishEditParamsValueValueOff      ZoneSettingPolishEditParamsValueValue = "off"
	ZoneSettingPolishEditParamsValueValueLossless ZoneSettingPolishEditParamsValueValue = "lossless"
	ZoneSettingPolishEditParamsValueValueLossy    ZoneSettingPolishEditParamsValueValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPolishEditParamsValueEditable bool

const (
	ZoneSettingPolishEditParamsValueEditableTrue  ZoneSettingPolishEditParamsValueEditable = true
	ZoneSettingPolishEditParamsValueEditableFalse ZoneSettingPolishEditParamsValueEditable = false
)

type ZoneSettingPolishEditResponseEnvelope struct {
	Errors   []ZoneSettingPolishEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingPolishEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result ZoneSettingPolishEditResponse             `json:"result"`
	JSON   zoneSettingPolishEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingPolishEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingPolishEditResponseEnvelope]
type zoneSettingPolishEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPolishEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingPolishEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingPolishEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingPolishEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingPolishEditResponseEnvelopeErrors]
type zoneSettingPolishEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPolishEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingPolishEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingPolishEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingPolishEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingPolishEditResponseEnvelopeMessages]
type zoneSettingPolishEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPolishEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingPolishGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingPolishGetResponseEnvelope struct {
	Errors   []ZoneSettingPolishGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingPolishGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result ZoneSettingPolishGetResponse             `json:"result"`
	JSON   zoneSettingPolishGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingPolishGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneSettingPolishGetResponseEnvelope]
type zoneSettingPolishGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPolishGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingPolishGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneSettingPolishGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingPolishGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneSettingPolishGetResponseEnvelopeErrors]
type zoneSettingPolishGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPolishGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingPolishGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingPolishGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingPolishGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneSettingPolishGetResponseEnvelopeMessages]
type zoneSettingPolishGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingPolishGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
