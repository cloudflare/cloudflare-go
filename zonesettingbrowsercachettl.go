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
func (r *ZoneSettingBrowserCacheTTLService) Edit(ctx context.Context, params ZoneSettingBrowserCacheTTLEditParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCacheTTLEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingBrowserCacheTTLEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *ZoneSettingBrowserCacheTTLService) Get(ctx context.Context, query ZoneSettingBrowserCacheTTLGetParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCacheTTLGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingBrowserCacheTTLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingBrowserCacheTTLEditResponse struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCacheTTLEditResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingBrowserCacheTTLEditResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCacheTTLEditResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrowserCacheTTLEditResponseJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLEditResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCacheTTLEditResponse]
type zoneSettingBrowserCacheTTLEditResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrowserCacheTTLEditResponseID string

const (
	ZoneSettingBrowserCacheTTLEditResponseIDBrowserCacheTTL ZoneSettingBrowserCacheTTLEditResponseID = "browser_cache_ttl"
)

// Current value of the zone setting.
type ZoneSettingBrowserCacheTTLEditResponseValue float64

const (
	ZoneSettingBrowserCacheTTLEditResponseValue0        ZoneSettingBrowserCacheTTLEditResponseValue = 0
	ZoneSettingBrowserCacheTTLEditResponseValue30       ZoneSettingBrowserCacheTTLEditResponseValue = 30
	ZoneSettingBrowserCacheTTLEditResponseValue60       ZoneSettingBrowserCacheTTLEditResponseValue = 60
	ZoneSettingBrowserCacheTTLEditResponseValue120      ZoneSettingBrowserCacheTTLEditResponseValue = 120
	ZoneSettingBrowserCacheTTLEditResponseValue300      ZoneSettingBrowserCacheTTLEditResponseValue = 300
	ZoneSettingBrowserCacheTTLEditResponseValue1200     ZoneSettingBrowserCacheTTLEditResponseValue = 1200
	ZoneSettingBrowserCacheTTLEditResponseValue1800     ZoneSettingBrowserCacheTTLEditResponseValue = 1800
	ZoneSettingBrowserCacheTTLEditResponseValue3600     ZoneSettingBrowserCacheTTLEditResponseValue = 3600
	ZoneSettingBrowserCacheTTLEditResponseValue7200     ZoneSettingBrowserCacheTTLEditResponseValue = 7200
	ZoneSettingBrowserCacheTTLEditResponseValue10800    ZoneSettingBrowserCacheTTLEditResponseValue = 10800
	ZoneSettingBrowserCacheTTLEditResponseValue14400    ZoneSettingBrowserCacheTTLEditResponseValue = 14400
	ZoneSettingBrowserCacheTTLEditResponseValue18000    ZoneSettingBrowserCacheTTLEditResponseValue = 18000
	ZoneSettingBrowserCacheTTLEditResponseValue28800    ZoneSettingBrowserCacheTTLEditResponseValue = 28800
	ZoneSettingBrowserCacheTTLEditResponseValue43200    ZoneSettingBrowserCacheTTLEditResponseValue = 43200
	ZoneSettingBrowserCacheTTLEditResponseValue57600    ZoneSettingBrowserCacheTTLEditResponseValue = 57600
	ZoneSettingBrowserCacheTTLEditResponseValue72000    ZoneSettingBrowserCacheTTLEditResponseValue = 72000
	ZoneSettingBrowserCacheTTLEditResponseValue86400    ZoneSettingBrowserCacheTTLEditResponseValue = 86400
	ZoneSettingBrowserCacheTTLEditResponseValue172800   ZoneSettingBrowserCacheTTLEditResponseValue = 172800
	ZoneSettingBrowserCacheTTLEditResponseValue259200   ZoneSettingBrowserCacheTTLEditResponseValue = 259200
	ZoneSettingBrowserCacheTTLEditResponseValue345600   ZoneSettingBrowserCacheTTLEditResponseValue = 345600
	ZoneSettingBrowserCacheTTLEditResponseValue432000   ZoneSettingBrowserCacheTTLEditResponseValue = 432000
	ZoneSettingBrowserCacheTTLEditResponseValue691200   ZoneSettingBrowserCacheTTLEditResponseValue = 691200
	ZoneSettingBrowserCacheTTLEditResponseValue1382400  ZoneSettingBrowserCacheTTLEditResponseValue = 1382400
	ZoneSettingBrowserCacheTTLEditResponseValue2073600  ZoneSettingBrowserCacheTTLEditResponseValue = 2073600
	ZoneSettingBrowserCacheTTLEditResponseValue2678400  ZoneSettingBrowserCacheTTLEditResponseValue = 2678400
	ZoneSettingBrowserCacheTTLEditResponseValue5356800  ZoneSettingBrowserCacheTTLEditResponseValue = 5356800
	ZoneSettingBrowserCacheTTLEditResponseValue16070400 ZoneSettingBrowserCacheTTLEditResponseValue = 16070400
	ZoneSettingBrowserCacheTTLEditResponseValue31536000 ZoneSettingBrowserCacheTTLEditResponseValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCacheTTLEditResponseEditable bool

const (
	ZoneSettingBrowserCacheTTLEditResponseEditableTrue  ZoneSettingBrowserCacheTTLEditResponseEditable = true
	ZoneSettingBrowserCacheTTLEditResponseEditableFalse ZoneSettingBrowserCacheTTLEditResponseEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingBrowserCacheTTLGetResponse struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCacheTTLGetResponseID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingBrowserCacheTTLGetResponseValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCacheTTLGetResponseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrowserCacheTTLGetResponseJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingBrowserCacheTTLGetResponse]
type zoneSettingBrowserCacheTTLGetResponseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneSettingBrowserCacheTTLGetResponseID string

const (
	ZoneSettingBrowserCacheTTLGetResponseIDBrowserCacheTTL ZoneSettingBrowserCacheTTLGetResponseID = "browser_cache_ttl"
)

// Current value of the zone setting.
type ZoneSettingBrowserCacheTTLGetResponseValue float64

const (
	ZoneSettingBrowserCacheTTLGetResponseValue0        ZoneSettingBrowserCacheTTLGetResponseValue = 0
	ZoneSettingBrowserCacheTTLGetResponseValue30       ZoneSettingBrowserCacheTTLGetResponseValue = 30
	ZoneSettingBrowserCacheTTLGetResponseValue60       ZoneSettingBrowserCacheTTLGetResponseValue = 60
	ZoneSettingBrowserCacheTTLGetResponseValue120      ZoneSettingBrowserCacheTTLGetResponseValue = 120
	ZoneSettingBrowserCacheTTLGetResponseValue300      ZoneSettingBrowserCacheTTLGetResponseValue = 300
	ZoneSettingBrowserCacheTTLGetResponseValue1200     ZoneSettingBrowserCacheTTLGetResponseValue = 1200
	ZoneSettingBrowserCacheTTLGetResponseValue1800     ZoneSettingBrowserCacheTTLGetResponseValue = 1800
	ZoneSettingBrowserCacheTTLGetResponseValue3600     ZoneSettingBrowserCacheTTLGetResponseValue = 3600
	ZoneSettingBrowserCacheTTLGetResponseValue7200     ZoneSettingBrowserCacheTTLGetResponseValue = 7200
	ZoneSettingBrowserCacheTTLGetResponseValue10800    ZoneSettingBrowserCacheTTLGetResponseValue = 10800
	ZoneSettingBrowserCacheTTLGetResponseValue14400    ZoneSettingBrowserCacheTTLGetResponseValue = 14400
	ZoneSettingBrowserCacheTTLGetResponseValue18000    ZoneSettingBrowserCacheTTLGetResponseValue = 18000
	ZoneSettingBrowserCacheTTLGetResponseValue28800    ZoneSettingBrowserCacheTTLGetResponseValue = 28800
	ZoneSettingBrowserCacheTTLGetResponseValue43200    ZoneSettingBrowserCacheTTLGetResponseValue = 43200
	ZoneSettingBrowserCacheTTLGetResponseValue57600    ZoneSettingBrowserCacheTTLGetResponseValue = 57600
	ZoneSettingBrowserCacheTTLGetResponseValue72000    ZoneSettingBrowserCacheTTLGetResponseValue = 72000
	ZoneSettingBrowserCacheTTLGetResponseValue86400    ZoneSettingBrowserCacheTTLGetResponseValue = 86400
	ZoneSettingBrowserCacheTTLGetResponseValue172800   ZoneSettingBrowserCacheTTLGetResponseValue = 172800
	ZoneSettingBrowserCacheTTLGetResponseValue259200   ZoneSettingBrowserCacheTTLGetResponseValue = 259200
	ZoneSettingBrowserCacheTTLGetResponseValue345600   ZoneSettingBrowserCacheTTLGetResponseValue = 345600
	ZoneSettingBrowserCacheTTLGetResponseValue432000   ZoneSettingBrowserCacheTTLGetResponseValue = 432000
	ZoneSettingBrowserCacheTTLGetResponseValue691200   ZoneSettingBrowserCacheTTLGetResponseValue = 691200
	ZoneSettingBrowserCacheTTLGetResponseValue1382400  ZoneSettingBrowserCacheTTLGetResponseValue = 1382400
	ZoneSettingBrowserCacheTTLGetResponseValue2073600  ZoneSettingBrowserCacheTTLGetResponseValue = 2073600
	ZoneSettingBrowserCacheTTLGetResponseValue2678400  ZoneSettingBrowserCacheTTLGetResponseValue = 2678400
	ZoneSettingBrowserCacheTTLGetResponseValue5356800  ZoneSettingBrowserCacheTTLGetResponseValue = 5356800
	ZoneSettingBrowserCacheTTLGetResponseValue16070400 ZoneSettingBrowserCacheTTLGetResponseValue = 16070400
	ZoneSettingBrowserCacheTTLGetResponseValue31536000 ZoneSettingBrowserCacheTTLGetResponseValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCacheTTLGetResponseEditable bool

const (
	ZoneSettingBrowserCacheTTLGetResponseEditableTrue  ZoneSettingBrowserCacheTTLGetResponseEditable = true
	ZoneSettingBrowserCacheTTLGetResponseEditableFalse ZoneSettingBrowserCacheTTLGetResponseEditable = false
)

type ZoneSettingBrowserCacheTTLEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[ZoneSettingBrowserCacheTTLEditParamsValue] `json:"value,required"`
}

func (r ZoneSettingBrowserCacheTTLEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingBrowserCacheTTLEditParamsValue float64

const (
	ZoneSettingBrowserCacheTTLEditParamsValue0        ZoneSettingBrowserCacheTTLEditParamsValue = 0
	ZoneSettingBrowserCacheTTLEditParamsValue30       ZoneSettingBrowserCacheTTLEditParamsValue = 30
	ZoneSettingBrowserCacheTTLEditParamsValue60       ZoneSettingBrowserCacheTTLEditParamsValue = 60
	ZoneSettingBrowserCacheTTLEditParamsValue120      ZoneSettingBrowserCacheTTLEditParamsValue = 120
	ZoneSettingBrowserCacheTTLEditParamsValue300      ZoneSettingBrowserCacheTTLEditParamsValue = 300
	ZoneSettingBrowserCacheTTLEditParamsValue1200     ZoneSettingBrowserCacheTTLEditParamsValue = 1200
	ZoneSettingBrowserCacheTTLEditParamsValue1800     ZoneSettingBrowserCacheTTLEditParamsValue = 1800
	ZoneSettingBrowserCacheTTLEditParamsValue3600     ZoneSettingBrowserCacheTTLEditParamsValue = 3600
	ZoneSettingBrowserCacheTTLEditParamsValue7200     ZoneSettingBrowserCacheTTLEditParamsValue = 7200
	ZoneSettingBrowserCacheTTLEditParamsValue10800    ZoneSettingBrowserCacheTTLEditParamsValue = 10800
	ZoneSettingBrowserCacheTTLEditParamsValue14400    ZoneSettingBrowserCacheTTLEditParamsValue = 14400
	ZoneSettingBrowserCacheTTLEditParamsValue18000    ZoneSettingBrowserCacheTTLEditParamsValue = 18000
	ZoneSettingBrowserCacheTTLEditParamsValue28800    ZoneSettingBrowserCacheTTLEditParamsValue = 28800
	ZoneSettingBrowserCacheTTLEditParamsValue43200    ZoneSettingBrowserCacheTTLEditParamsValue = 43200
	ZoneSettingBrowserCacheTTLEditParamsValue57600    ZoneSettingBrowserCacheTTLEditParamsValue = 57600
	ZoneSettingBrowserCacheTTLEditParamsValue72000    ZoneSettingBrowserCacheTTLEditParamsValue = 72000
	ZoneSettingBrowserCacheTTLEditParamsValue86400    ZoneSettingBrowserCacheTTLEditParamsValue = 86400
	ZoneSettingBrowserCacheTTLEditParamsValue172800   ZoneSettingBrowserCacheTTLEditParamsValue = 172800
	ZoneSettingBrowserCacheTTLEditParamsValue259200   ZoneSettingBrowserCacheTTLEditParamsValue = 259200
	ZoneSettingBrowserCacheTTLEditParamsValue345600   ZoneSettingBrowserCacheTTLEditParamsValue = 345600
	ZoneSettingBrowserCacheTTLEditParamsValue432000   ZoneSettingBrowserCacheTTLEditParamsValue = 432000
	ZoneSettingBrowserCacheTTLEditParamsValue691200   ZoneSettingBrowserCacheTTLEditParamsValue = 691200
	ZoneSettingBrowserCacheTTLEditParamsValue1382400  ZoneSettingBrowserCacheTTLEditParamsValue = 1382400
	ZoneSettingBrowserCacheTTLEditParamsValue2073600  ZoneSettingBrowserCacheTTLEditParamsValue = 2073600
	ZoneSettingBrowserCacheTTLEditParamsValue2678400  ZoneSettingBrowserCacheTTLEditParamsValue = 2678400
	ZoneSettingBrowserCacheTTLEditParamsValue5356800  ZoneSettingBrowserCacheTTLEditParamsValue = 5356800
	ZoneSettingBrowserCacheTTLEditParamsValue16070400 ZoneSettingBrowserCacheTTLEditParamsValue = 16070400
	ZoneSettingBrowserCacheTTLEditParamsValue31536000 ZoneSettingBrowserCacheTTLEditParamsValue = 31536000
)

type ZoneSettingBrowserCacheTTLEditResponseEnvelope struct {
	Errors   []ZoneSettingBrowserCacheTTLEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingBrowserCacheTTLEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result ZoneSettingBrowserCacheTTLEditResponse             `json:"result"`
	JSON   zoneSettingBrowserCacheTTLEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneSettingBrowserCacheTTLEditResponseEnvelope]
type zoneSettingBrowserCacheTTLEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLEditResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingBrowserCacheTTLEditResponseEnvelopeErrors]
type zoneSettingBrowserCacheTTLEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLEditResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingBrowserCacheTTLEditResponseEnvelopeMessages]
type zoneSettingBrowserCacheTTLEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingBrowserCacheTTLGetResponseEnvelope struct {
	Errors   []ZoneSettingBrowserCacheTTLGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingBrowserCacheTTLGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result ZoneSettingBrowserCacheTTLGetResponse             `json:"result"`
	JSON   zoneSettingBrowserCacheTTLGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneSettingBrowserCacheTTLGetResponseEnvelope]
type zoneSettingBrowserCacheTTLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneSettingBrowserCacheTTLGetResponseEnvelopeErrors]
type zoneSettingBrowserCacheTTLGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingBrowserCacheTTLGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zoneSettingBrowserCacheTTLGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneSettingBrowserCacheTTLGetResponseEnvelopeMessages]
type zoneSettingBrowserCacheTTLGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
