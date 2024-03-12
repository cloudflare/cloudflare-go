// File generated from our OpenAPI spec by Stainless.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingBrowserCacheTTLService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingBrowserCacheTTLService]
// method instead.
type SettingBrowserCacheTTLService struct {
	Options []option.RequestOption
}

// NewSettingBrowserCacheTTLService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingBrowserCacheTTLService(opts ...option.RequestOption) (r *SettingBrowserCacheTTLService) {
	r = &SettingBrowserCacheTTLService{}
	r.Options = opts
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *SettingBrowserCacheTTLService) Edit(ctx context.Context, params SettingBrowserCacheTTLEditParams, opts ...option.RequestOption) (res *ZonesBrowserCacheTTL, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrowserCacheTTLEditResponseEnvelope
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
func (r *SettingBrowserCacheTTLService) Get(ctx context.Context, query SettingBrowserCacheTTLGetParams, opts ...option.RequestOption) (res *ZonesBrowserCacheTTL, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingBrowserCacheTTLGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

func (r zonesBrowserCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r ZonesBrowserCacheTTL) implementsZonesSettingEditResponse() {}

func (r ZonesBrowserCacheTTL) implementsZonesSettingGetResponse() {}

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

func (r ZonesBrowserCacheTTLParam) implementsZonesSettingEditParamsItem() {}

type SettingBrowserCacheTTLEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[SettingBrowserCacheTTLEditParamsValue] `json:"value,required"`
}

func (r SettingBrowserCacheTTLEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type SettingBrowserCacheTTLEditParamsValue float64

const (
	SettingBrowserCacheTTLEditParamsValue0        SettingBrowserCacheTTLEditParamsValue = 0
	SettingBrowserCacheTTLEditParamsValue30       SettingBrowserCacheTTLEditParamsValue = 30
	SettingBrowserCacheTTLEditParamsValue60       SettingBrowserCacheTTLEditParamsValue = 60
	SettingBrowserCacheTTLEditParamsValue120      SettingBrowserCacheTTLEditParamsValue = 120
	SettingBrowserCacheTTLEditParamsValue300      SettingBrowserCacheTTLEditParamsValue = 300
	SettingBrowserCacheTTLEditParamsValue1200     SettingBrowserCacheTTLEditParamsValue = 1200
	SettingBrowserCacheTTLEditParamsValue1800     SettingBrowserCacheTTLEditParamsValue = 1800
	SettingBrowserCacheTTLEditParamsValue3600     SettingBrowserCacheTTLEditParamsValue = 3600
	SettingBrowserCacheTTLEditParamsValue7200     SettingBrowserCacheTTLEditParamsValue = 7200
	SettingBrowserCacheTTLEditParamsValue10800    SettingBrowserCacheTTLEditParamsValue = 10800
	SettingBrowserCacheTTLEditParamsValue14400    SettingBrowserCacheTTLEditParamsValue = 14400
	SettingBrowserCacheTTLEditParamsValue18000    SettingBrowserCacheTTLEditParamsValue = 18000
	SettingBrowserCacheTTLEditParamsValue28800    SettingBrowserCacheTTLEditParamsValue = 28800
	SettingBrowserCacheTTLEditParamsValue43200    SettingBrowserCacheTTLEditParamsValue = 43200
	SettingBrowserCacheTTLEditParamsValue57600    SettingBrowserCacheTTLEditParamsValue = 57600
	SettingBrowserCacheTTLEditParamsValue72000    SettingBrowserCacheTTLEditParamsValue = 72000
	SettingBrowserCacheTTLEditParamsValue86400    SettingBrowserCacheTTLEditParamsValue = 86400
	SettingBrowserCacheTTLEditParamsValue172800   SettingBrowserCacheTTLEditParamsValue = 172800
	SettingBrowserCacheTTLEditParamsValue259200   SettingBrowserCacheTTLEditParamsValue = 259200
	SettingBrowserCacheTTLEditParamsValue345600   SettingBrowserCacheTTLEditParamsValue = 345600
	SettingBrowserCacheTTLEditParamsValue432000   SettingBrowserCacheTTLEditParamsValue = 432000
	SettingBrowserCacheTTLEditParamsValue691200   SettingBrowserCacheTTLEditParamsValue = 691200
	SettingBrowserCacheTTLEditParamsValue1382400  SettingBrowserCacheTTLEditParamsValue = 1382400
	SettingBrowserCacheTTLEditParamsValue2073600  SettingBrowserCacheTTLEditParamsValue = 2073600
	SettingBrowserCacheTTLEditParamsValue2678400  SettingBrowserCacheTTLEditParamsValue = 2678400
	SettingBrowserCacheTTLEditParamsValue5356800  SettingBrowserCacheTTLEditParamsValue = 5356800
	SettingBrowserCacheTTLEditParamsValue16070400 SettingBrowserCacheTTLEditParamsValue = 16070400
	SettingBrowserCacheTTLEditParamsValue31536000 SettingBrowserCacheTTLEditParamsValue = 31536000
)

type SettingBrowserCacheTTLEditResponseEnvelope struct {
	Errors   []SettingBrowserCacheTTLEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrowserCacheTTLEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result ZonesBrowserCacheTTL                           `json:"result"`
	JSON   settingBrowserCacheTTLEditResponseEnvelopeJSON `json:"-"`
}

// settingBrowserCacheTTLEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingBrowserCacheTTLEditResponseEnvelope]
type settingBrowserCacheTTLEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCacheTTLEditResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    settingBrowserCacheTTLEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrowserCacheTTLEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingBrowserCacheTTLEditResponseEnvelopeErrors]
type settingBrowserCacheTTLEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCacheTTLEditResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    settingBrowserCacheTTLEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrowserCacheTTLEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingBrowserCacheTTLEditResponseEnvelopeMessages]
type settingBrowserCacheTTLEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCacheTTLGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingBrowserCacheTTLGetResponseEnvelope struct {
	Errors   []SettingBrowserCacheTTLGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingBrowserCacheTTLGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result ZonesBrowserCacheTTL                          `json:"result"`
	JSON   settingBrowserCacheTTLGetResponseEnvelopeJSON `json:"-"`
}

// settingBrowserCacheTTLGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBrowserCacheTTLGetResponseEnvelope]
type settingBrowserCacheTTLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCacheTTLGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    settingBrowserCacheTTLGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingBrowserCacheTTLGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingBrowserCacheTTLGetResponseEnvelopeErrors]
type settingBrowserCacheTTLGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBrowserCacheTTLGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    settingBrowserCacheTTLGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingBrowserCacheTTLGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingBrowserCacheTTLGetResponseEnvelopeMessages]
type settingBrowserCacheTTLGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBrowserCacheTTLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBrowserCacheTTLGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
