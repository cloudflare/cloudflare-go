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

// ZoneSettingWebpService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingWebpService] method
// instead.
type ZoneSettingWebpService struct {
	Options []option.RequestOption
}

// NewZoneSettingWebpService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingWebpService(opts ...option.RequestOption) (r *ZoneSettingWebpService) {
	r = &ZoneSettingWebpService{}
	r.Options = opts
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *ZoneSettingWebpService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingWebpUpdateParams, opts ...option.RequestOption) (res *ZoneSettingWebpUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/webp", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *ZoneSettingWebpService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingWebpListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/webp", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingWebpUpdateResponse struct {
	Errors   []ZoneSettingWebpUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingWebpUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                `json:"success,required"`
	Result  ZoneSettingWebpUpdateResponseResult `json:"result"`
	JSON    zoneSettingWebpUpdateResponseJSON   `json:"-"`
}

// zoneSettingWebpUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWebpUpdateResponse]
type zoneSettingWebpUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebpUpdateResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingWebpUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingWebpUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingWebpUpdateResponseError]
type zoneSettingWebpUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebpUpdateResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingWebpUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingWebpUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingWebpUpdateResponseMessage]
type zoneSettingWebpUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebpUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingWebpUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingWebpUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWebpUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingWebpUpdateResponseResultJSON `json:"-"`
}

// zoneSettingWebpUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingWebpUpdateResponseResult]
type zoneSettingWebpUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingWebpUpdateResponseResultID string

const (
	ZoneSettingWebpUpdateResponseResultIDWebp ZoneSettingWebpUpdateResponseResultID = "webp"
)

// Value of the zone setting.
type ZoneSettingWebpUpdateResponseResultValue string

const (
	ZoneSettingWebpUpdateResponseResultValueOff ZoneSettingWebpUpdateResponseResultValue = "off"
	ZoneSettingWebpUpdateResponseResultValueOn  ZoneSettingWebpUpdateResponseResultValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWebpUpdateResponseResultEditable bool

const (
	ZoneSettingWebpUpdateResponseResultEditableTrue  ZoneSettingWebpUpdateResponseResultEditable = true
	ZoneSettingWebpUpdateResponseResultEditableFalse ZoneSettingWebpUpdateResponseResultEditable = false
)

type ZoneSettingWebpListResponse struct {
	Errors   []ZoneSettingWebpListResponseError   `json:"errors,required"`
	Messages []ZoneSettingWebpListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                              `json:"success,required"`
	Result  ZoneSettingWebpListResponseResult `json:"result"`
	JSON    zoneSettingWebpListResponseJSON   `json:"-"`
}

// zoneSettingWebpListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingWebpListResponse]
type zoneSettingWebpListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebpListResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneSettingWebpListResponseErrorJSON `json:"-"`
}

// zoneSettingWebpListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingWebpListResponseError]
type zoneSettingWebpListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebpListResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingWebpListResponseMessageJSON `json:"-"`
}

// zoneSettingWebpListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingWebpListResponseMessage]
type zoneSettingWebpListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingWebpListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingWebpListResponseResultID `json:"id,required"`
	// Value of the zone setting.
	Value ZoneSettingWebpListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingWebpListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingWebpListResponseResultJSON `json:"-"`
}

// zoneSettingWebpListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingWebpListResponseResult]
type zoneSettingWebpListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingWebpListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingWebpListResponseResultID string

const (
	ZoneSettingWebpListResponseResultIDWebp ZoneSettingWebpListResponseResultID = "webp"
)

// Value of the zone setting.
type ZoneSettingWebpListResponseResultValue string

const (
	ZoneSettingWebpListResponseResultValueOff ZoneSettingWebpListResponseResultValue = "off"
	ZoneSettingWebpListResponseResultValueOn  ZoneSettingWebpListResponseResultValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingWebpListResponseResultEditable bool

const (
	ZoneSettingWebpListResponseResultEditableTrue  ZoneSettingWebpListResponseResultEditable = true
	ZoneSettingWebpListResponseResultEditableFalse ZoneSettingWebpListResponseResultEditable = false
)

type ZoneSettingWebpUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingWebpUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingWebpUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingWebpUpdateParamsValue string

const (
	ZoneSettingWebpUpdateParamsValueOff ZoneSettingWebpUpdateParamsValue = "off"
	ZoneSettingWebpUpdateParamsValueOn  ZoneSettingWebpUpdateParamsValue = "on"
)
