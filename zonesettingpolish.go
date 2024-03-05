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
func (r *ZoneSettingPolishService) Edit(ctx context.Context, params ZoneSettingPolishEditParams, opts ...option.RequestOption) (res *ZonesPolish, err error) {
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
func (r *ZoneSettingPolishService) Get(ctx context.Context, query ZoneSettingPolishGetParams, opts ...option.RequestOption) (res *ZonesPolish, err error) {
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

func (r ZonesPolish) implementsZoneSettingEditResponse() {}

func (r ZonesPolish) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesPolishID string

const (
	ZonesPolishIDPolish ZonesPolishID = "polish"
)

// Current value of the zone setting.
type ZonesPolishValue string

const (
	ZonesPolishValueOff      ZonesPolishValue = "off"
	ZonesPolishValueLossless ZonesPolishValue = "lossless"
	ZonesPolishValueLossy    ZonesPolishValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesPolishEditable bool

const (
	ZonesPolishEditableTrue  ZonesPolishEditable = true
	ZonesPolishEditableFalse ZonesPolishEditable = false
)

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

func (r ZonesPolishParam) implementsZoneSettingEditParamsItem() {}

type ZoneSettingPolishEditParams struct {
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

func (r ZoneSettingPolishEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	Result ZonesPolish                               `json:"result"`
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
	Result ZonesPolish                              `json:"result"`
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
