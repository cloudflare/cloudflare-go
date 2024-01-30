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

// ZoneSettingBrowserCacheTTLSService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingBrowserCacheTTLSService] method instead.
type ZoneSettingBrowserCacheTTLSService struct {
	Options []option.RequestOption
}

// NewZoneSettingBrowserCacheTTLSService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingBrowserCacheTTLSService(opts ...option.RequestOption) (r *ZoneSettingBrowserCacheTTLSService) {
	r = &ZoneSettingBrowserCacheTTLSService{}
	r.Options = opts
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *ZoneSettingBrowserCacheTTLSService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingBrowserCacheTTLSUpdateParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCacheTTLSUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *ZoneSettingBrowserCacheTTLSService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingBrowserCacheTTLSListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingBrowserCacheTTLSUpdateResponse struct {
	Errors   []ZoneSettingBrowserCacheTTLSUpdateResponseError   `json:"errors,required"`
	Messages []ZoneSettingBrowserCacheTTLSUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                            `json:"success,required"`
	Result  ZoneSettingBrowserCacheTTLSUpdateResponseResult `json:"result"`
	JSON    zoneSettingBrowserCacheTTLSUpdateResponseJSON   `json:"-"`
}

// zoneSettingBrowserCacheTTLSUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCacheTTLSUpdateResponse]
type zoneSettingBrowserCacheTTLSUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLSUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLSUpdateResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLSUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLSUpdateResponseErrorJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCacheTTLSUpdateResponseError]
type zoneSettingBrowserCacheTTLSUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLSUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLSUpdateResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLSUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLSUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCacheTTLSUpdateResponseMessage]
type zoneSettingBrowserCacheTTLSUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLSUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLSUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCacheTTLSUpdateResponseResultID `json:"id,required"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingBrowserCacheTTLSUpdateResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCacheTTLSUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrowserCacheTTLSUpdateResponseResultJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLSUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCacheTTLSUpdateResponseResult]
type zoneSettingBrowserCacheTTLSUpdateResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLSUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrowserCacheTTLSUpdateResponseResultID string

const (
	ZoneSettingBrowserCacheTTLSUpdateResponseResultIDBrowserCacheTtl ZoneSettingBrowserCacheTTLSUpdateResponseResultID = "browser_cache_ttl"
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingBrowserCacheTTLSUpdateResponseResultValue float64

const (
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue0        ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 0
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue30       ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 30
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue60       ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 60
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue120      ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 120
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue300      ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 300
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue1200     ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 1200
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue1800     ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 1800
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue3600     ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 3600
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue7200     ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 7200
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue10800    ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 10800
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue14400    ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 14400
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue18000    ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 18000
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue28800    ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 28800
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue43200    ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 43200
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue57600    ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 57600
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue72000    ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 72000
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue86400    ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 86400
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue172800   ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 172800
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue259200   ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 259200
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue345600   ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 345600
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue432000   ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 432000
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue691200   ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 691200
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue1382400  ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 1382400
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue2073600  ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 2073600
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue2678400  ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 2678400
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue5356800  ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 5356800
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue16070400 ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 16070400
	ZoneSettingBrowserCacheTTLSUpdateResponseResultValue31536000 ZoneSettingBrowserCacheTTLSUpdateResponseResultValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCacheTTLSUpdateResponseResultEditable bool

const (
	ZoneSettingBrowserCacheTTLSUpdateResponseResultEditableTrue  ZoneSettingBrowserCacheTTLSUpdateResponseResultEditable = true
	ZoneSettingBrowserCacheTTLSUpdateResponseResultEditableFalse ZoneSettingBrowserCacheTTLSUpdateResponseResultEditable = false
)

type ZoneSettingBrowserCacheTTLSListResponse struct {
	Errors   []ZoneSettingBrowserCacheTTLSListResponseError   `json:"errors,required"`
	Messages []ZoneSettingBrowserCacheTTLSListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                          `json:"success,required"`
	Result  ZoneSettingBrowserCacheTTLSListResponseResult `json:"result"`
	JSON    zoneSettingBrowserCacheTTLSListResponseJSON   `json:"-"`
}

// zoneSettingBrowserCacheTTLSListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCacheTTLSListResponse]
type zoneSettingBrowserCacheTTLSListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLSListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLSListResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLSListResponseErrorJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLSListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTTLSListResponseError]
type zoneSettingBrowserCacheTTLSListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLSListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLSListResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLSListResponseMessageJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLSListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCacheTTLSListResponseMessage]
type zoneSettingBrowserCacheTTLSListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLSListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLSListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCacheTTLSListResponseResultID `json:"id,required"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingBrowserCacheTTLSListResponseResultValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCacheTTLSListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrowserCacheTTLSListResponseResultJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLSListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTTLSListResponseResult]
type zoneSettingBrowserCacheTTLSListResponseResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLSListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrowserCacheTTLSListResponseResultID string

const (
	ZoneSettingBrowserCacheTTLSListResponseResultIDBrowserCacheTtl ZoneSettingBrowserCacheTTLSListResponseResultID = "browser_cache_ttl"
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingBrowserCacheTTLSListResponseResultValue float64

const (
	ZoneSettingBrowserCacheTTLSListResponseResultValue0        ZoneSettingBrowserCacheTTLSListResponseResultValue = 0
	ZoneSettingBrowserCacheTTLSListResponseResultValue30       ZoneSettingBrowserCacheTTLSListResponseResultValue = 30
	ZoneSettingBrowserCacheTTLSListResponseResultValue60       ZoneSettingBrowserCacheTTLSListResponseResultValue = 60
	ZoneSettingBrowserCacheTTLSListResponseResultValue120      ZoneSettingBrowserCacheTTLSListResponseResultValue = 120
	ZoneSettingBrowserCacheTTLSListResponseResultValue300      ZoneSettingBrowserCacheTTLSListResponseResultValue = 300
	ZoneSettingBrowserCacheTTLSListResponseResultValue1200     ZoneSettingBrowserCacheTTLSListResponseResultValue = 1200
	ZoneSettingBrowserCacheTTLSListResponseResultValue1800     ZoneSettingBrowserCacheTTLSListResponseResultValue = 1800
	ZoneSettingBrowserCacheTTLSListResponseResultValue3600     ZoneSettingBrowserCacheTTLSListResponseResultValue = 3600
	ZoneSettingBrowserCacheTTLSListResponseResultValue7200     ZoneSettingBrowserCacheTTLSListResponseResultValue = 7200
	ZoneSettingBrowserCacheTTLSListResponseResultValue10800    ZoneSettingBrowserCacheTTLSListResponseResultValue = 10800
	ZoneSettingBrowserCacheTTLSListResponseResultValue14400    ZoneSettingBrowserCacheTTLSListResponseResultValue = 14400
	ZoneSettingBrowserCacheTTLSListResponseResultValue18000    ZoneSettingBrowserCacheTTLSListResponseResultValue = 18000
	ZoneSettingBrowserCacheTTLSListResponseResultValue28800    ZoneSettingBrowserCacheTTLSListResponseResultValue = 28800
	ZoneSettingBrowserCacheTTLSListResponseResultValue43200    ZoneSettingBrowserCacheTTLSListResponseResultValue = 43200
	ZoneSettingBrowserCacheTTLSListResponseResultValue57600    ZoneSettingBrowserCacheTTLSListResponseResultValue = 57600
	ZoneSettingBrowserCacheTTLSListResponseResultValue72000    ZoneSettingBrowserCacheTTLSListResponseResultValue = 72000
	ZoneSettingBrowserCacheTTLSListResponseResultValue86400    ZoneSettingBrowserCacheTTLSListResponseResultValue = 86400
	ZoneSettingBrowserCacheTTLSListResponseResultValue172800   ZoneSettingBrowserCacheTTLSListResponseResultValue = 172800
	ZoneSettingBrowserCacheTTLSListResponseResultValue259200   ZoneSettingBrowserCacheTTLSListResponseResultValue = 259200
	ZoneSettingBrowserCacheTTLSListResponseResultValue345600   ZoneSettingBrowserCacheTTLSListResponseResultValue = 345600
	ZoneSettingBrowserCacheTTLSListResponseResultValue432000   ZoneSettingBrowserCacheTTLSListResponseResultValue = 432000
	ZoneSettingBrowserCacheTTLSListResponseResultValue691200   ZoneSettingBrowserCacheTTLSListResponseResultValue = 691200
	ZoneSettingBrowserCacheTTLSListResponseResultValue1382400  ZoneSettingBrowserCacheTTLSListResponseResultValue = 1382400
	ZoneSettingBrowserCacheTTLSListResponseResultValue2073600  ZoneSettingBrowserCacheTTLSListResponseResultValue = 2073600
	ZoneSettingBrowserCacheTTLSListResponseResultValue2678400  ZoneSettingBrowserCacheTTLSListResponseResultValue = 2678400
	ZoneSettingBrowserCacheTTLSListResponseResultValue5356800  ZoneSettingBrowserCacheTTLSListResponseResultValue = 5356800
	ZoneSettingBrowserCacheTTLSListResponseResultValue16070400 ZoneSettingBrowserCacheTTLSListResponseResultValue = 16070400
	ZoneSettingBrowserCacheTTLSListResponseResultValue31536000 ZoneSettingBrowserCacheTTLSListResponseResultValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCacheTTLSListResponseResultEditable bool

const (
	ZoneSettingBrowserCacheTTLSListResponseResultEditableTrue  ZoneSettingBrowserCacheTTLSListResponseResultEditable = true
	ZoneSettingBrowserCacheTTLSListResponseResultEditableFalse ZoneSettingBrowserCacheTTLSListResponseResultEditable = false
)

type ZoneSettingBrowserCacheTTLSUpdateParams struct {
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[ZoneSettingBrowserCacheTTLSUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingBrowserCacheTTLSUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingBrowserCacheTTLSUpdateParamsValue float64

const (
	ZoneSettingBrowserCacheTTLSUpdateParamsValue0        ZoneSettingBrowserCacheTTLSUpdateParamsValue = 0
	ZoneSettingBrowserCacheTTLSUpdateParamsValue30       ZoneSettingBrowserCacheTTLSUpdateParamsValue = 30
	ZoneSettingBrowserCacheTTLSUpdateParamsValue60       ZoneSettingBrowserCacheTTLSUpdateParamsValue = 60
	ZoneSettingBrowserCacheTTLSUpdateParamsValue120      ZoneSettingBrowserCacheTTLSUpdateParamsValue = 120
	ZoneSettingBrowserCacheTTLSUpdateParamsValue300      ZoneSettingBrowserCacheTTLSUpdateParamsValue = 300
	ZoneSettingBrowserCacheTTLSUpdateParamsValue1200     ZoneSettingBrowserCacheTTLSUpdateParamsValue = 1200
	ZoneSettingBrowserCacheTTLSUpdateParamsValue1800     ZoneSettingBrowserCacheTTLSUpdateParamsValue = 1800
	ZoneSettingBrowserCacheTTLSUpdateParamsValue3600     ZoneSettingBrowserCacheTTLSUpdateParamsValue = 3600
	ZoneSettingBrowserCacheTTLSUpdateParamsValue7200     ZoneSettingBrowserCacheTTLSUpdateParamsValue = 7200
	ZoneSettingBrowserCacheTTLSUpdateParamsValue10800    ZoneSettingBrowserCacheTTLSUpdateParamsValue = 10800
	ZoneSettingBrowserCacheTTLSUpdateParamsValue14400    ZoneSettingBrowserCacheTTLSUpdateParamsValue = 14400
	ZoneSettingBrowserCacheTTLSUpdateParamsValue18000    ZoneSettingBrowserCacheTTLSUpdateParamsValue = 18000
	ZoneSettingBrowserCacheTTLSUpdateParamsValue28800    ZoneSettingBrowserCacheTTLSUpdateParamsValue = 28800
	ZoneSettingBrowserCacheTTLSUpdateParamsValue43200    ZoneSettingBrowserCacheTTLSUpdateParamsValue = 43200
	ZoneSettingBrowserCacheTTLSUpdateParamsValue57600    ZoneSettingBrowserCacheTTLSUpdateParamsValue = 57600
	ZoneSettingBrowserCacheTTLSUpdateParamsValue72000    ZoneSettingBrowserCacheTTLSUpdateParamsValue = 72000
	ZoneSettingBrowserCacheTTLSUpdateParamsValue86400    ZoneSettingBrowserCacheTTLSUpdateParamsValue = 86400
	ZoneSettingBrowserCacheTTLSUpdateParamsValue172800   ZoneSettingBrowserCacheTTLSUpdateParamsValue = 172800
	ZoneSettingBrowserCacheTTLSUpdateParamsValue259200   ZoneSettingBrowserCacheTTLSUpdateParamsValue = 259200
	ZoneSettingBrowserCacheTTLSUpdateParamsValue345600   ZoneSettingBrowserCacheTTLSUpdateParamsValue = 345600
	ZoneSettingBrowserCacheTTLSUpdateParamsValue432000   ZoneSettingBrowserCacheTTLSUpdateParamsValue = 432000
	ZoneSettingBrowserCacheTTLSUpdateParamsValue691200   ZoneSettingBrowserCacheTTLSUpdateParamsValue = 691200
	ZoneSettingBrowserCacheTTLSUpdateParamsValue1382400  ZoneSettingBrowserCacheTTLSUpdateParamsValue = 1382400
	ZoneSettingBrowserCacheTTLSUpdateParamsValue2073600  ZoneSettingBrowserCacheTTLSUpdateParamsValue = 2073600
	ZoneSettingBrowserCacheTTLSUpdateParamsValue2678400  ZoneSettingBrowserCacheTTLSUpdateParamsValue = 2678400
	ZoneSettingBrowserCacheTTLSUpdateParamsValue5356800  ZoneSettingBrowserCacheTTLSUpdateParamsValue = 5356800
	ZoneSettingBrowserCacheTTLSUpdateParamsValue16070400 ZoneSettingBrowserCacheTTLSUpdateParamsValue = 16070400
	ZoneSettingBrowserCacheTTLSUpdateParamsValue31536000 ZoneSettingBrowserCacheTTLSUpdateParamsValue = 31536000
)
