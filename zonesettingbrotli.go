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

// ZoneSettingBrotliService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingBrotliService] method
// instead.
type ZoneSettingBrotliService struct {
	Options []option.RequestOption
}

// NewZoneSettingBrotliService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingBrotliService(opts ...option.RequestOption) (r *ZoneSettingBrotliService) {
	r = &ZoneSettingBrotliService{}
	r.Options = opts
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
func (r *ZoneSettingBrotliService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingBrotliUpdateParams, opts ...option.RequestOption) (res *ZoneSettingBrotliUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/brotli", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
func (r *ZoneSettingBrotliService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingBrotliListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/brotli", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingBrotliUpdateResponse struct {
	Errors   []ZoneSettingBrotliUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingBrotliUpdateResponseMessage `json:"messages"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result ZoneSettingBrotliUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                `json:"success"`
	JSON    zoneSettingBrotliUpdateResponseJSON `json:"-"`
}

// zoneSettingBrotliUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingBrotliUpdateResponse]
type zoneSettingBrotliUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrotliUpdateResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingBrotliUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingBrotliUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSettingBrotliUpdateResponseError]
type zoneSettingBrotliUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrotliUpdateResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSettingBrotliUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingBrotliUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingBrotliUpdateResponseMessage]
type zoneSettingBrotliUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingBrotliUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingBrotliUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrotliUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingBrotliUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingBrotliUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingBrotliUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingBrotliUpdateResponseResult]
type zoneSettingBrotliUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrotliUpdateResponseResultID string

const (
	ZoneSettingBrotliUpdateResponseResultIDBrotli ZoneSettingBrotliUpdateResponseResultID = "brotli"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrotliUpdateResponseResultEditable bool

const (
	ZoneSettingBrotliUpdateResponseResultEditableTrue  ZoneSettingBrotliUpdateResponseResultEditable = true
	ZoneSettingBrotliUpdateResponseResultEditableFalse ZoneSettingBrotliUpdateResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingBrotliUpdateResponseResultValue string

const (
	ZoneSettingBrotliUpdateResponseResultValueOff ZoneSettingBrotliUpdateResponseResultValue = "off"
	ZoneSettingBrotliUpdateResponseResultValueOn  ZoneSettingBrotliUpdateResponseResultValue = "on"
)

type ZoneSettingBrotliListResponse struct {
	Errors   []ZoneSettingBrotliListResponseError   `json:"errors"`
	Messages []ZoneSettingBrotliListResponseMessage `json:"messages"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result ZoneSettingBrotliListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                              `json:"success"`
	JSON    zoneSettingBrotliListResponseJSON `json:"-"`
}

// zoneSettingBrotliListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingBrotliListResponse]
type zoneSettingBrotliListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrotliListResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneSettingBrotliListResponseErrorJSON `json:"-"`
}

// zoneSettingBrotliListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingBrotliListResponseError]
type zoneSettingBrotliListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrotliListResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingBrotliListResponseMessageJSON `json:"-"`
}

// zoneSettingBrotliListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSettingBrotliListResponseMessage]
type zoneSettingBrotliListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingBrotliListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingBrotliListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrotliListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingBrotliListResponseResultValue `json:"value"`
	JSON  zoneSettingBrotliListResponseResultJSON  `json:"-"`
}

// zoneSettingBrotliListResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingBrotliListResponseResult]
type zoneSettingBrotliListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrotliListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrotliListResponseResultID string

const (
	ZoneSettingBrotliListResponseResultIDBrotli ZoneSettingBrotliListResponseResultID = "brotli"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrotliListResponseResultEditable bool

const (
	ZoneSettingBrotliListResponseResultEditableTrue  ZoneSettingBrotliListResponseResultEditable = true
	ZoneSettingBrotliListResponseResultEditableFalse ZoneSettingBrotliListResponseResultEditable = false
)

// Value of the zone setting.
type ZoneSettingBrotliListResponseResultValue string

const (
	ZoneSettingBrotliListResponseResultValueOff ZoneSettingBrotliListResponseResultValue = "off"
	ZoneSettingBrotliListResponseResultValueOn  ZoneSettingBrotliListResponseResultValue = "on"
)

type ZoneSettingBrotliUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[ZoneSettingBrotliUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingBrotliUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type ZoneSettingBrotliUpdateParamsValue string

const (
	ZoneSettingBrotliUpdateParamsValueOff ZoneSettingBrotliUpdateParamsValue = "off"
	ZoneSettingBrotliUpdateParamsValueOn  ZoneSettingBrotliUpdateParamsValue = "on"
)
