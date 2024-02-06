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

// ZoneSettingBrowserCacheTTLService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSettingBrowserCacheTTLService] method instead.
type ZoneSettingBrowserCacheTTLService struct {
	Options []option.RequestOption
}

// NewZoneSettingBrowserCacheTTLService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingBrowserCacheTTLService(opts ...option.RequestOption) (r *ZoneSettingBrowserCacheTTLService) {
	r = &ZoneSettingBrowserCacheTTLService{}
	r.Options = opts
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *ZoneSettingBrowserCacheTTLService) Update(ctx context.Context, zoneIdentifier string, body ZoneSettingBrowserCacheTTLUpdateParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCacheTTLUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *ZoneSettingBrowserCacheTTLService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingBrowserCacheTTLListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSettingBrowserCacheTTLUpdateResponse struct {
	Errors   []ZoneSettingBrowserCacheTTLUpdateResponseError   `json:"errors"`
	Messages []ZoneSettingBrowserCacheTTLUpdateResponseMessage `json:"messages"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result ZoneSettingBrowserCacheTTLUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                         `json:"success"`
	JSON    zoneSettingBrowserCacheTTLUpdateResponseJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCacheTTLUpdateResponse]
type zoneSettingBrowserCacheTTLUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLUpdateResponseErrorJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTTLUpdateResponseError]
type zoneSettingBrowserCacheTTLUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLUpdateResponseMessageJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCacheTTLUpdateResponseMessage]
type zoneSettingBrowserCacheTTLUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingBrowserCacheTTLUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCacheTTLUpdateResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCacheTTLUpdateResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingBrowserCacheTTLUpdateResponseResultValue `json:"value"`
	JSON  zoneSettingBrowserCacheTTLUpdateResponseResultJSON  `json:"-"`
}

// zoneSettingBrowserCacheTTLUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCacheTTLUpdateResponseResult]
type zoneSettingBrowserCacheTTLUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrowserCacheTTLUpdateResponseResultID string

const (
	ZoneSettingBrowserCacheTTLUpdateResponseResultIDBrowserCacheTTL ZoneSettingBrowserCacheTTLUpdateResponseResultID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCacheTTLUpdateResponseResultEditable bool

const (
	ZoneSettingBrowserCacheTTLUpdateResponseResultEditableTrue  ZoneSettingBrowserCacheTTLUpdateResponseResultEditable = true
	ZoneSettingBrowserCacheTTLUpdateResponseResultEditableFalse ZoneSettingBrowserCacheTTLUpdateResponseResultEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingBrowserCacheTTLUpdateResponseResultValue float64

const (
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue0        ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 0
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue30       ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 30
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue60       ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 60
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue120      ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 120
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue300      ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 300
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue1200     ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 1200
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue1800     ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 1800
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue3600     ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 3600
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue7200     ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 7200
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue10800    ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 10800
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue14400    ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 14400
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue18000    ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 18000
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue28800    ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 28800
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue43200    ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 43200
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue57600    ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 57600
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue72000    ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 72000
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue86400    ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 86400
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue172800   ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 172800
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue259200   ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 259200
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue345600   ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 345600
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue432000   ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 432000
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue691200   ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 691200
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue1382400  ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 1382400
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue2073600  ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 2073600
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue2678400  ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 2678400
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue5356800  ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 5356800
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue16070400 ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 16070400
	ZoneSettingBrowserCacheTTLUpdateResponseResultValue31536000 ZoneSettingBrowserCacheTTLUpdateResponseResultValue = 31536000
)

type ZoneSettingBrowserCacheTTLListResponse struct {
	Errors   []ZoneSettingBrowserCacheTTLListResponseError   `json:"errors"`
	Messages []ZoneSettingBrowserCacheTTLListResponseMessage `json:"messages"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result ZoneSettingBrowserCacheTTLListResponseResult `json:"result"`
	// Whether the API call was successful
	Success bool                                       `json:"success"`
	JSON    zoneSettingBrowserCacheTTLListResponseJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLListResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCacheTTLListResponse]
type zoneSettingBrowserCacheTTLListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLListResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLListResponseErrorJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTTLListResponseError]
type zoneSettingBrowserCacheTTLListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLListResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLListResponseMessageJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTTLListResponseMessage]
type zoneSettingBrowserCacheTTLListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingBrowserCacheTTLListResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCacheTTLListResponseResultID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCacheTTLListResponseResultEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingBrowserCacheTTLListResponseResultValue `json:"value"`
	JSON  zoneSettingBrowserCacheTTLListResponseResultJSON  `json:"-"`
}

// zoneSettingBrowserCacheTTLListResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTTLListResponseResult]
type zoneSettingBrowserCacheTTLListResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrowserCacheTTLListResponseResultID string

const (
	ZoneSettingBrowserCacheTTLListResponseResultIDBrowserCacheTTL ZoneSettingBrowserCacheTTLListResponseResultID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCacheTTLListResponseResultEditable bool

const (
	ZoneSettingBrowserCacheTTLListResponseResultEditableTrue  ZoneSettingBrowserCacheTTLListResponseResultEditable = true
	ZoneSettingBrowserCacheTTLListResponseResultEditableFalse ZoneSettingBrowserCacheTTLListResponseResultEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingBrowserCacheTTLListResponseResultValue float64

const (
	ZoneSettingBrowserCacheTTLListResponseResultValue0        ZoneSettingBrowserCacheTTLListResponseResultValue = 0
	ZoneSettingBrowserCacheTTLListResponseResultValue30       ZoneSettingBrowserCacheTTLListResponseResultValue = 30
	ZoneSettingBrowserCacheTTLListResponseResultValue60       ZoneSettingBrowserCacheTTLListResponseResultValue = 60
	ZoneSettingBrowserCacheTTLListResponseResultValue120      ZoneSettingBrowserCacheTTLListResponseResultValue = 120
	ZoneSettingBrowserCacheTTLListResponseResultValue300      ZoneSettingBrowserCacheTTLListResponseResultValue = 300
	ZoneSettingBrowserCacheTTLListResponseResultValue1200     ZoneSettingBrowserCacheTTLListResponseResultValue = 1200
	ZoneSettingBrowserCacheTTLListResponseResultValue1800     ZoneSettingBrowserCacheTTLListResponseResultValue = 1800
	ZoneSettingBrowserCacheTTLListResponseResultValue3600     ZoneSettingBrowserCacheTTLListResponseResultValue = 3600
	ZoneSettingBrowserCacheTTLListResponseResultValue7200     ZoneSettingBrowserCacheTTLListResponseResultValue = 7200
	ZoneSettingBrowserCacheTTLListResponseResultValue10800    ZoneSettingBrowserCacheTTLListResponseResultValue = 10800
	ZoneSettingBrowserCacheTTLListResponseResultValue14400    ZoneSettingBrowserCacheTTLListResponseResultValue = 14400
	ZoneSettingBrowserCacheTTLListResponseResultValue18000    ZoneSettingBrowserCacheTTLListResponseResultValue = 18000
	ZoneSettingBrowserCacheTTLListResponseResultValue28800    ZoneSettingBrowserCacheTTLListResponseResultValue = 28800
	ZoneSettingBrowserCacheTTLListResponseResultValue43200    ZoneSettingBrowserCacheTTLListResponseResultValue = 43200
	ZoneSettingBrowserCacheTTLListResponseResultValue57600    ZoneSettingBrowserCacheTTLListResponseResultValue = 57600
	ZoneSettingBrowserCacheTTLListResponseResultValue72000    ZoneSettingBrowserCacheTTLListResponseResultValue = 72000
	ZoneSettingBrowserCacheTTLListResponseResultValue86400    ZoneSettingBrowserCacheTTLListResponseResultValue = 86400
	ZoneSettingBrowserCacheTTLListResponseResultValue172800   ZoneSettingBrowserCacheTTLListResponseResultValue = 172800
	ZoneSettingBrowserCacheTTLListResponseResultValue259200   ZoneSettingBrowserCacheTTLListResponseResultValue = 259200
	ZoneSettingBrowserCacheTTLListResponseResultValue345600   ZoneSettingBrowserCacheTTLListResponseResultValue = 345600
	ZoneSettingBrowserCacheTTLListResponseResultValue432000   ZoneSettingBrowserCacheTTLListResponseResultValue = 432000
	ZoneSettingBrowserCacheTTLListResponseResultValue691200   ZoneSettingBrowserCacheTTLListResponseResultValue = 691200
	ZoneSettingBrowserCacheTTLListResponseResultValue1382400  ZoneSettingBrowserCacheTTLListResponseResultValue = 1382400
	ZoneSettingBrowserCacheTTLListResponseResultValue2073600  ZoneSettingBrowserCacheTTLListResponseResultValue = 2073600
	ZoneSettingBrowserCacheTTLListResponseResultValue2678400  ZoneSettingBrowserCacheTTLListResponseResultValue = 2678400
	ZoneSettingBrowserCacheTTLListResponseResultValue5356800  ZoneSettingBrowserCacheTTLListResponseResultValue = 5356800
	ZoneSettingBrowserCacheTTLListResponseResultValue16070400 ZoneSettingBrowserCacheTTLListResponseResultValue = 16070400
	ZoneSettingBrowserCacheTTLListResponseResultValue31536000 ZoneSettingBrowserCacheTTLListResponseResultValue = 31536000
)

type ZoneSettingBrowserCacheTTLUpdateParams struct {
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[ZoneSettingBrowserCacheTTLUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingBrowserCacheTTLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingBrowserCacheTTLUpdateParamsValue float64

const (
	ZoneSettingBrowserCacheTTLUpdateParamsValue0        ZoneSettingBrowserCacheTTLUpdateParamsValue = 0
	ZoneSettingBrowserCacheTTLUpdateParamsValue30       ZoneSettingBrowserCacheTTLUpdateParamsValue = 30
	ZoneSettingBrowserCacheTTLUpdateParamsValue60       ZoneSettingBrowserCacheTTLUpdateParamsValue = 60
	ZoneSettingBrowserCacheTTLUpdateParamsValue120      ZoneSettingBrowserCacheTTLUpdateParamsValue = 120
	ZoneSettingBrowserCacheTTLUpdateParamsValue300      ZoneSettingBrowserCacheTTLUpdateParamsValue = 300
	ZoneSettingBrowserCacheTTLUpdateParamsValue1200     ZoneSettingBrowserCacheTTLUpdateParamsValue = 1200
	ZoneSettingBrowserCacheTTLUpdateParamsValue1800     ZoneSettingBrowserCacheTTLUpdateParamsValue = 1800
	ZoneSettingBrowserCacheTTLUpdateParamsValue3600     ZoneSettingBrowserCacheTTLUpdateParamsValue = 3600
	ZoneSettingBrowserCacheTTLUpdateParamsValue7200     ZoneSettingBrowserCacheTTLUpdateParamsValue = 7200
	ZoneSettingBrowserCacheTTLUpdateParamsValue10800    ZoneSettingBrowserCacheTTLUpdateParamsValue = 10800
	ZoneSettingBrowserCacheTTLUpdateParamsValue14400    ZoneSettingBrowserCacheTTLUpdateParamsValue = 14400
	ZoneSettingBrowserCacheTTLUpdateParamsValue18000    ZoneSettingBrowserCacheTTLUpdateParamsValue = 18000
	ZoneSettingBrowserCacheTTLUpdateParamsValue28800    ZoneSettingBrowserCacheTTLUpdateParamsValue = 28800
	ZoneSettingBrowserCacheTTLUpdateParamsValue43200    ZoneSettingBrowserCacheTTLUpdateParamsValue = 43200
	ZoneSettingBrowserCacheTTLUpdateParamsValue57600    ZoneSettingBrowserCacheTTLUpdateParamsValue = 57600
	ZoneSettingBrowserCacheTTLUpdateParamsValue72000    ZoneSettingBrowserCacheTTLUpdateParamsValue = 72000
	ZoneSettingBrowserCacheTTLUpdateParamsValue86400    ZoneSettingBrowserCacheTTLUpdateParamsValue = 86400
	ZoneSettingBrowserCacheTTLUpdateParamsValue172800   ZoneSettingBrowserCacheTTLUpdateParamsValue = 172800
	ZoneSettingBrowserCacheTTLUpdateParamsValue259200   ZoneSettingBrowserCacheTTLUpdateParamsValue = 259200
	ZoneSettingBrowserCacheTTLUpdateParamsValue345600   ZoneSettingBrowserCacheTTLUpdateParamsValue = 345600
	ZoneSettingBrowserCacheTTLUpdateParamsValue432000   ZoneSettingBrowserCacheTTLUpdateParamsValue = 432000
	ZoneSettingBrowserCacheTTLUpdateParamsValue691200   ZoneSettingBrowserCacheTTLUpdateParamsValue = 691200
	ZoneSettingBrowserCacheTTLUpdateParamsValue1382400  ZoneSettingBrowserCacheTTLUpdateParamsValue = 1382400
	ZoneSettingBrowserCacheTTLUpdateParamsValue2073600  ZoneSettingBrowserCacheTTLUpdateParamsValue = 2073600
	ZoneSettingBrowserCacheTTLUpdateParamsValue2678400  ZoneSettingBrowserCacheTTLUpdateParamsValue = 2678400
	ZoneSettingBrowserCacheTTLUpdateParamsValue5356800  ZoneSettingBrowserCacheTTLUpdateParamsValue = 5356800
	ZoneSettingBrowserCacheTTLUpdateParamsValue16070400 ZoneSettingBrowserCacheTTLUpdateParamsValue = 16070400
	ZoneSettingBrowserCacheTTLUpdateParamsValue31536000 ZoneSettingBrowserCacheTTLUpdateParamsValue = 31536000
)
