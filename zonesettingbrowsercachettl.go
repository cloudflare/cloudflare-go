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

// ZoneSettingBrowserCacheTtlService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingBrowserCacheTtlService] method instead.
type ZoneSettingBrowserCacheTtlService struct {
	Options []option.RequestOption
}

// NewZoneSettingBrowserCacheTtlService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingBrowserCacheTtlService(opts ...option.RequestOption) (r *ZoneSettingBrowserCacheTtlService) {
	r = &ZoneSettingBrowserCacheTtlService{}
	r.Options = opts
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *ZoneSettingBrowserCacheTtlService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingBrowserCacheTtlUpdateParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCacheTtlUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *ZoneSettingBrowserCacheTtlService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingBrowserCacheTtlListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingBrowserCacheTtlUpdateResponse struct {
	Errors   []ZoneSettingBrowserCacheTtlUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingBrowserCacheTtlUpdateResponseMessage `json:"messages"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result ZoneSettingBrowserCacheTtlUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                         `json:"success"`
	JSON    zoneSettingBrowserCacheTtlUpdateResponseJSON `json:"-"`
}

// zoneSettingBrowserCacheTtlUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCacheTtlUpdateResponse]
type zoneSettingBrowserCacheTtlUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTtlUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTtlUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingBrowserCacheTtlUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingBrowserCacheTtlUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTtlUpdateResponseError]
type zoneSettingBrowserCacheTtlUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTtlUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTtlUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingBrowserCacheTtlUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingBrowserCacheTtlUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCacheTtlUpdateResponseMessage]
type zoneSettingBrowserCacheTtlUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTtlUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingBrowserCacheTtlUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCacheTtlUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCacheTtlUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingBrowserCacheTtlUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingBrowserCacheTtlUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingBrowserCacheTtlUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCacheTtlUpdateResponseResult]
type zoneSettingBrowserCacheTtlUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTtlUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrowserCacheTtlUpdateResponseResultID string

const (
	ZoneSettingBrowserCacheTtlUpdateResponseResultIDBrowserCacheTtl ZoneSettingBrowserCacheTtlUpdateResponseResultID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCacheTtlUpdateResponseResultEditable bool

const (
	ZoneSettingBrowserCacheTtlUpdateResponseResultEditableTrue  ZoneSettingBrowserCacheTtlUpdateResponseResultEditable = true
	ZoneSettingBrowserCacheTtlUpdateResponseResultEditableFalse ZoneSettingBrowserCacheTtlUpdateResponseResultEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingBrowserCacheTtlUpdateResponseResultValue float64

const (
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue0        ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 0
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue30       ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 30
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue60       ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 60
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue120      ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 120
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue300      ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 300
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue1200     ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 1200
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue1800     ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 1800
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue3600     ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 3600
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue7200     ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 7200
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue10800    ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 10800
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue14400    ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 14400
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue18000    ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 18000
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue28800    ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 28800
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue43200    ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 43200
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue57600    ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 57600
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue72000    ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 72000
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue86400    ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 86400
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue172800   ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 172800
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue259200   ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 259200
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue345600   ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 345600
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue432000   ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 432000
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue691200   ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 691200
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue1382400  ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 1382400
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue2073600  ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 2073600
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue2678400  ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 2678400
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue5356800  ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 5356800
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue16070400 ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 16070400
	ZoneSettingBrowserCacheTtlUpdateResponseResultValue31536000 ZoneSettingBrowserCacheTtlUpdateResponseResultValue = 31536000
)

type ZoneSettingBrowserCacheTtlListResponse struct {
	Errors   []ZoneSettingBrowserCacheTtlListResponseError   `json:"errors"`
	Messages []ZoneSettingBrowserCacheTtlListResponseMessage `json:"messages"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result ZoneSettingBrowserCacheTtlListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                       `json:"success"`
	JSON    zoneSettingBrowserCacheTtlListResponseJSON `json:"-"`
}

// zoneSettingBrowserCacheTtlListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCacheTtlListResponse]
type zoneSettingBrowserCacheTtlListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTtlListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTtlListResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingBrowserCacheTtlListResponseErrorJSON `json:"-"`
}

// zoneSettingBrowserCacheTtlListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTtlListResponseError]
type zoneSettingBrowserCacheTtlListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTtlListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTtlListResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingBrowserCacheTtlListResponseMessageJSON `json:"-"`
}

// zoneSettingBrowserCacheTtlListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTtlListResponseMessage]
type zoneSettingBrowserCacheTtlListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTtlListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingBrowserCacheTtlListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCacheTtlListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCacheTtlListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingBrowserCacheTtlListResponseResultValue `json:"value"`
	JSON  zoneSettingBrowserCacheTtlListResponseResultJSON  `json:"-"`
}

// zoneSettingBrowserCacheTtlListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTtlListResponseResult]
type zoneSettingBrowserCacheTtlListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTtlListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrowserCacheTtlListResponseResultID string

const (
	ZoneSettingBrowserCacheTtlListResponseResultIDBrowserCacheTtl ZoneSettingBrowserCacheTtlListResponseResultID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCacheTtlListResponseResultEditable bool

const (
	ZoneSettingBrowserCacheTtlListResponseResultEditableTrue  ZoneSettingBrowserCacheTtlListResponseResultEditable = true
	ZoneSettingBrowserCacheTtlListResponseResultEditableFalse ZoneSettingBrowserCacheTtlListResponseResultEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingBrowserCacheTtlListResponseResultValue float64

const (
	ZoneSettingBrowserCacheTtlListResponseResultValue0        ZoneSettingBrowserCacheTtlListResponseResultValue = 0
	ZoneSettingBrowserCacheTtlListResponseResultValue30       ZoneSettingBrowserCacheTtlListResponseResultValue = 30
	ZoneSettingBrowserCacheTtlListResponseResultValue60       ZoneSettingBrowserCacheTtlListResponseResultValue = 60
	ZoneSettingBrowserCacheTtlListResponseResultValue120      ZoneSettingBrowserCacheTtlListResponseResultValue = 120
	ZoneSettingBrowserCacheTtlListResponseResultValue300      ZoneSettingBrowserCacheTtlListResponseResultValue = 300
	ZoneSettingBrowserCacheTtlListResponseResultValue1200     ZoneSettingBrowserCacheTtlListResponseResultValue = 1200
	ZoneSettingBrowserCacheTtlListResponseResultValue1800     ZoneSettingBrowserCacheTtlListResponseResultValue = 1800
	ZoneSettingBrowserCacheTtlListResponseResultValue3600     ZoneSettingBrowserCacheTtlListResponseResultValue = 3600
	ZoneSettingBrowserCacheTtlListResponseResultValue7200     ZoneSettingBrowserCacheTtlListResponseResultValue = 7200
	ZoneSettingBrowserCacheTtlListResponseResultValue10800    ZoneSettingBrowserCacheTtlListResponseResultValue = 10800
	ZoneSettingBrowserCacheTtlListResponseResultValue14400    ZoneSettingBrowserCacheTtlListResponseResultValue = 14400
	ZoneSettingBrowserCacheTtlListResponseResultValue18000    ZoneSettingBrowserCacheTtlListResponseResultValue = 18000
	ZoneSettingBrowserCacheTtlListResponseResultValue28800    ZoneSettingBrowserCacheTtlListResponseResultValue = 28800
	ZoneSettingBrowserCacheTtlListResponseResultValue43200    ZoneSettingBrowserCacheTtlListResponseResultValue = 43200
	ZoneSettingBrowserCacheTtlListResponseResultValue57600    ZoneSettingBrowserCacheTtlListResponseResultValue = 57600
	ZoneSettingBrowserCacheTtlListResponseResultValue72000    ZoneSettingBrowserCacheTtlListResponseResultValue = 72000
	ZoneSettingBrowserCacheTtlListResponseResultValue86400    ZoneSettingBrowserCacheTtlListResponseResultValue = 86400
	ZoneSettingBrowserCacheTtlListResponseResultValue172800   ZoneSettingBrowserCacheTtlListResponseResultValue = 172800
	ZoneSettingBrowserCacheTtlListResponseResultValue259200   ZoneSettingBrowserCacheTtlListResponseResultValue = 259200
	ZoneSettingBrowserCacheTtlListResponseResultValue345600   ZoneSettingBrowserCacheTtlListResponseResultValue = 345600
	ZoneSettingBrowserCacheTtlListResponseResultValue432000   ZoneSettingBrowserCacheTtlListResponseResultValue = 432000
	ZoneSettingBrowserCacheTtlListResponseResultValue691200   ZoneSettingBrowserCacheTtlListResponseResultValue = 691200
	ZoneSettingBrowserCacheTtlListResponseResultValue1382400  ZoneSettingBrowserCacheTtlListResponseResultValue = 1382400
	ZoneSettingBrowserCacheTtlListResponseResultValue2073600  ZoneSettingBrowserCacheTtlListResponseResultValue = 2073600
	ZoneSettingBrowserCacheTtlListResponseResultValue2678400  ZoneSettingBrowserCacheTtlListResponseResultValue = 2678400
	ZoneSettingBrowserCacheTtlListResponseResultValue5356800  ZoneSettingBrowserCacheTtlListResponseResultValue = 5356800
	ZoneSettingBrowserCacheTtlListResponseResultValue16070400 ZoneSettingBrowserCacheTtlListResponseResultValue = 16070400
	ZoneSettingBrowserCacheTtlListResponseResultValue31536000 ZoneSettingBrowserCacheTtlListResponseResultValue = 31536000
)

type ZoneSettingBrowserCacheTtlUpdateParams struct {
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[ZoneSettingBrowserCacheTtlUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingBrowserCacheTtlUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingBrowserCacheTtlUpdateParamsValue float64

const (
	ZoneSettingBrowserCacheTtlUpdateParamsValue0        ZoneSettingBrowserCacheTtlUpdateParamsValue = 0
	ZoneSettingBrowserCacheTtlUpdateParamsValue30       ZoneSettingBrowserCacheTtlUpdateParamsValue = 30
	ZoneSettingBrowserCacheTtlUpdateParamsValue60       ZoneSettingBrowserCacheTtlUpdateParamsValue = 60
	ZoneSettingBrowserCacheTtlUpdateParamsValue120      ZoneSettingBrowserCacheTtlUpdateParamsValue = 120
	ZoneSettingBrowserCacheTtlUpdateParamsValue300      ZoneSettingBrowserCacheTtlUpdateParamsValue = 300
	ZoneSettingBrowserCacheTtlUpdateParamsValue1200     ZoneSettingBrowserCacheTtlUpdateParamsValue = 1200
	ZoneSettingBrowserCacheTtlUpdateParamsValue1800     ZoneSettingBrowserCacheTtlUpdateParamsValue = 1800
	ZoneSettingBrowserCacheTtlUpdateParamsValue3600     ZoneSettingBrowserCacheTtlUpdateParamsValue = 3600
	ZoneSettingBrowserCacheTtlUpdateParamsValue7200     ZoneSettingBrowserCacheTtlUpdateParamsValue = 7200
	ZoneSettingBrowserCacheTtlUpdateParamsValue10800    ZoneSettingBrowserCacheTtlUpdateParamsValue = 10800
	ZoneSettingBrowserCacheTtlUpdateParamsValue14400    ZoneSettingBrowserCacheTtlUpdateParamsValue = 14400
	ZoneSettingBrowserCacheTtlUpdateParamsValue18000    ZoneSettingBrowserCacheTtlUpdateParamsValue = 18000
	ZoneSettingBrowserCacheTtlUpdateParamsValue28800    ZoneSettingBrowserCacheTtlUpdateParamsValue = 28800
	ZoneSettingBrowserCacheTtlUpdateParamsValue43200    ZoneSettingBrowserCacheTtlUpdateParamsValue = 43200
	ZoneSettingBrowserCacheTtlUpdateParamsValue57600    ZoneSettingBrowserCacheTtlUpdateParamsValue = 57600
	ZoneSettingBrowserCacheTtlUpdateParamsValue72000    ZoneSettingBrowserCacheTtlUpdateParamsValue = 72000
	ZoneSettingBrowserCacheTtlUpdateParamsValue86400    ZoneSettingBrowserCacheTtlUpdateParamsValue = 86400
	ZoneSettingBrowserCacheTtlUpdateParamsValue172800   ZoneSettingBrowserCacheTtlUpdateParamsValue = 172800
	ZoneSettingBrowserCacheTtlUpdateParamsValue259200   ZoneSettingBrowserCacheTtlUpdateParamsValue = 259200
	ZoneSettingBrowserCacheTtlUpdateParamsValue345600   ZoneSettingBrowserCacheTtlUpdateParamsValue = 345600
	ZoneSettingBrowserCacheTtlUpdateParamsValue432000   ZoneSettingBrowserCacheTtlUpdateParamsValue = 432000
	ZoneSettingBrowserCacheTtlUpdateParamsValue691200   ZoneSettingBrowserCacheTtlUpdateParamsValue = 691200
	ZoneSettingBrowserCacheTtlUpdateParamsValue1382400  ZoneSettingBrowserCacheTtlUpdateParamsValue = 1382400
	ZoneSettingBrowserCacheTtlUpdateParamsValue2073600  ZoneSettingBrowserCacheTtlUpdateParamsValue = 2073600
	ZoneSettingBrowserCacheTtlUpdateParamsValue2678400  ZoneSettingBrowserCacheTtlUpdateParamsValue = 2678400
	ZoneSettingBrowserCacheTtlUpdateParamsValue5356800  ZoneSettingBrowserCacheTtlUpdateParamsValue = 5356800
	ZoneSettingBrowserCacheTtlUpdateParamsValue16070400 ZoneSettingBrowserCacheTtlUpdateParamsValue = 16070400
	ZoneSettingBrowserCacheTtlUpdateParamsValue31536000 ZoneSettingBrowserCacheTtlUpdateParamsValue = 31536000
)
