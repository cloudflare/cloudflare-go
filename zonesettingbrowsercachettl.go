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
func (r *ZoneSettingBrowserCacheTTLService) Edit(ctx context.Context, params ZoneSettingBrowserCacheTTLEditParams, opts ...option.RequestOption) (res *ZonesBrowserCacheTTL, err error) {
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
func (r *ZoneSettingBrowserCacheTTLService) Get(ctx context.Context, query ZoneSettingBrowserCacheTTLGetParams, opts ...option.RequestOption) (res *ZonesBrowserCacheTTL, err error) {
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
type ZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID ZonesBrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesBrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesBrowserCacheTTLJSON `json:"-"`
}

// zonesBrowserCacheTTLJSON contains the JSON metadata for the struct
// [ZonesBrowserCacheTTL]
type zonesBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonesBrowserCacheTTL) implementsZoneSettingEditResponse() {}

func (r ZonesBrowserCacheTTL) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZonesBrowserCacheTTLID string

const (
	ZonesBrowserCacheTTLIDBrowserCacheTTL ZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Current value of the zone setting.
type ZonesBrowserCacheTTLValue float64

const (
	ZonesBrowserCacheTTLValue0        ZonesBrowserCacheTTLValue = 0
	ZonesBrowserCacheTTLValue30       ZonesBrowserCacheTTLValue = 30
	ZonesBrowserCacheTTLValue60       ZonesBrowserCacheTTLValue = 60
	ZonesBrowserCacheTTLValue120      ZonesBrowserCacheTTLValue = 120
	ZonesBrowserCacheTTLValue300      ZonesBrowserCacheTTLValue = 300
	ZonesBrowserCacheTTLValue1200     ZonesBrowserCacheTTLValue = 1200
	ZonesBrowserCacheTTLValue1800     ZonesBrowserCacheTTLValue = 1800
	ZonesBrowserCacheTTLValue3600     ZonesBrowserCacheTTLValue = 3600
	ZonesBrowserCacheTTLValue7200     ZonesBrowserCacheTTLValue = 7200
	ZonesBrowserCacheTTLValue10800    ZonesBrowserCacheTTLValue = 10800
	ZonesBrowserCacheTTLValue14400    ZonesBrowserCacheTTLValue = 14400
	ZonesBrowserCacheTTLValue18000    ZonesBrowserCacheTTLValue = 18000
	ZonesBrowserCacheTTLValue28800    ZonesBrowserCacheTTLValue = 28800
	ZonesBrowserCacheTTLValue43200    ZonesBrowserCacheTTLValue = 43200
	ZonesBrowserCacheTTLValue57600    ZonesBrowserCacheTTLValue = 57600
	ZonesBrowserCacheTTLValue72000    ZonesBrowserCacheTTLValue = 72000
	ZonesBrowserCacheTTLValue86400    ZonesBrowserCacheTTLValue = 86400
	ZonesBrowserCacheTTLValue172800   ZonesBrowserCacheTTLValue = 172800
	ZonesBrowserCacheTTLValue259200   ZonesBrowserCacheTTLValue = 259200
	ZonesBrowserCacheTTLValue345600   ZonesBrowserCacheTTLValue = 345600
	ZonesBrowserCacheTTLValue432000   ZonesBrowserCacheTTLValue = 432000
	ZonesBrowserCacheTTLValue691200   ZonesBrowserCacheTTLValue = 691200
	ZonesBrowserCacheTTLValue1382400  ZonesBrowserCacheTTLValue = 1382400
	ZonesBrowserCacheTTLValue2073600  ZonesBrowserCacheTTLValue = 2073600
	ZonesBrowserCacheTTLValue2678400  ZonesBrowserCacheTTLValue = 2678400
	ZonesBrowserCacheTTLValue5356800  ZonesBrowserCacheTTLValue = 5356800
	ZonesBrowserCacheTTLValue16070400 ZonesBrowserCacheTTLValue = 16070400
	ZonesBrowserCacheTTLValue31536000 ZonesBrowserCacheTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesBrowserCacheTTLEditable bool

const (
	ZonesBrowserCacheTTLEditableTrue  ZonesBrowserCacheTTLEditable = true
	ZonesBrowserCacheTTLEditableFalse ZonesBrowserCacheTTLEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZonesBrowserCacheTTLParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesBrowserCacheTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesBrowserCacheTTLValue] `json:"value,required"`
}

func (r ZonesBrowserCacheTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesBrowserCacheTTLParam) implementsZoneSettingEditParamsItem() {}

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
	Result ZonesBrowserCacheTTL                               `json:"result"`
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
	Result ZonesBrowserCacheTTL                              `json:"result"`
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
