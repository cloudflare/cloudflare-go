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
func (r *ZoneSettingPolishService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingPolishUpdateParams, opts ...option.RequestOption) (res *ZoneSettingPolishUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/polish", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/polish-solving-mobile-speed)
// for more information.
func (r *ZoneSettingPolishService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingPolishListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/polish", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingPolishUpdateResponse struct {
	Errors   []ZoneSettingPolishUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingPolishUpdateResponseMessage `json:"messages"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result ZoneSettingPolishUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                `json:"success"`
	JSON    zoneSettingPolishUpdateResponseJSON `json:"-"`
}

// zoneSettingPolishUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingPolishUpdateResponse]
type zoneSettingPolishUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPolishUpdateResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingPolishUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingPolishUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingPolishUpdateResponseError]
type zoneSettingPolishUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPolishUpdateResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSettingPolishUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingPolishUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingPolishUpdateResponseMessage]
type zoneSettingPolishUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingPolishUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingPolishUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPolishUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingPolishUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingPolishUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingPolishUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingPolishUpdateResponseResult]
type zoneSettingPolishUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingPolishUpdateResponseResultID string

const (
	ZoneSettingPolishUpdateResponseResultIDPolish ZoneSettingPolishUpdateResponseResultID = "polish"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPolishUpdateResponseResultEditable bool

const (
	ZoneSettingPolishUpdateResponseResultEditableTrue  ZoneSettingPolishUpdateResponseResultEditable = true
	ZoneSettingPolishUpdateResponseResultEditableFalse ZoneSettingPolishUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingPolishUpdateResponseResultValue string

const (
	ZoneSettingPolishUpdateResponseResultValueOff      ZoneSettingPolishUpdateResponseResultValue = "off"
	ZoneSettingPolishUpdateResponseResultValueLossless ZoneSettingPolishUpdateResponseResultValue = "lossless"
	ZoneSettingPolishUpdateResponseResultValueLossy    ZoneSettingPolishUpdateResponseResultValue = "lossy"
)

type ZoneSettingPolishListResponse struct {
	Errors   []ZoneSettingPolishListResponseError   `json:"errors"`
	Messages []ZoneSettingPolishListResponseMessage `json:"messages"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result ZoneSettingPolishListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                              `json:"success"`
	JSON    zoneSettingPolishListResponseJSON `json:"-"`
}

// zoneSettingPolishListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingPolishListResponse]
type zoneSettingPolishListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPolishListResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingPolishListResponseErrorJSON `json:"-"`
}

// zoneSettingPolishListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingPolishListResponseError]
type zoneSettingPolishListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingPolishListResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingPolishListResponseMessageJSON `json:"-"`
}

// zoneSettingPolishListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingPolishListResponseMessage]
type zoneSettingPolishListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingPolishListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingPolishListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingPolishListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingPolishListResponseResultValue `json:"value"`
	JSON  zoneSettingPolishListResponseResultJSON  `json:"-"`
}

// zoneSettingPolishListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingPolishListResponseResult]
type zoneSettingPolishListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingPolishListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingPolishListResponseResultID string

const (
	ZoneSettingPolishListResponseResultIDPolish ZoneSettingPolishListResponseResultID = "polish"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPolishListResponseResultEditable bool

const (
	ZoneSettingPolishListResponseResultEditableTrue  ZoneSettingPolishListResponseResultEditable = true
	ZoneSettingPolishListResponseResultEditableFalse ZoneSettingPolishListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingPolishListResponseResultValue string

const (
	ZoneSettingPolishListResponseResultValueOff      ZoneSettingPolishListResponseResultValue = "off"
	ZoneSettingPolishListResponseResultValueLossless ZoneSettingPolishListResponseResultValue = "lossless"
	ZoneSettingPolishListResponseResultValueLossy    ZoneSettingPolishListResponseResultValue = "lossy"
)

type ZoneSettingPolishUpdateParams struct {
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Value param.Field[ZoneSettingPolishUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingPolishUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingPolishUpdateParamsValue struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingPolishUpdateParamsValueID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingPolishUpdateParamsValueValue] `json:"value"`
}

func (r ZoneSettingPolishUpdateParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the zone setting.
type ZoneSettingPolishUpdateParamsValueID string

const (
	ZoneSettingPolishUpdateParamsValueIDPolish ZoneSettingPolishUpdateParamsValueID = "polish"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingPolishUpdateParamsValueEditable bool

const (
	ZoneSettingPolishUpdateParamsValueEditableTrue  ZoneSettingPolishUpdateParamsValueEditable = true
	ZoneSettingPolishUpdateParamsValueEditableFalse ZoneSettingPolishUpdateParamsValueEditable = false
)

// Value of the zone setting.
type ZoneSettingPolishUpdateParamsValueValue string

const (
	ZoneSettingPolishUpdateParamsValueValueOff      ZoneSettingPolishUpdateParamsValueValue = "off"
	ZoneSettingPolishUpdateParamsValueValueLossless ZoneSettingPolishUpdateParamsValueValue = "lossless"
	ZoneSettingPolishUpdateParamsValueValueLossy    ZoneSettingPolishUpdateParamsValueValue = "lossy"
)
