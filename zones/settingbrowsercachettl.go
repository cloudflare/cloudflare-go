// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *SettingBrowserCacheTTLService) Edit(ctx context.Context, params SettingBrowserCacheTTLEditParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCacheTTL, err error) {
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
func (r *SettingBrowserCacheTTLService) Get(ctx context.Context, query SettingBrowserCacheTTLGetParams, opts ...option.RequestOption) (res *ZoneSettingBrowserCacheTTL, err error) {
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
type ZoneSettingBrowserCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingBrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingBrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingBrowserCacheTTLJSON `json:"-"`
}

// zoneSettingBrowserCacheTTLJSON contains the JSON metadata for the struct
// [ZoneSettingBrowserCacheTTL]
type zoneSettingBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingBrowserCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingBrowserCacheTTL) implementsZonesSettingEditResponse() {}

func (r ZoneSettingBrowserCacheTTL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingBrowserCacheTTLID string

const (
	ZoneSettingBrowserCacheTTLIDBrowserCacheTTL ZoneSettingBrowserCacheTTLID = "browser_cache_ttl"
)

func (r ZoneSettingBrowserCacheTTLID) IsKnown() bool {
	switch r {
	case ZoneSettingBrowserCacheTTLIDBrowserCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZoneSettingBrowserCacheTTLValue float64

const (
	ZoneSettingBrowserCacheTTLValue0        ZoneSettingBrowserCacheTTLValue = 0
	ZoneSettingBrowserCacheTTLValue30       ZoneSettingBrowserCacheTTLValue = 30
	ZoneSettingBrowserCacheTTLValue60       ZoneSettingBrowserCacheTTLValue = 60
	ZoneSettingBrowserCacheTTLValue120      ZoneSettingBrowserCacheTTLValue = 120
	ZoneSettingBrowserCacheTTLValue300      ZoneSettingBrowserCacheTTLValue = 300
	ZoneSettingBrowserCacheTTLValue1200     ZoneSettingBrowserCacheTTLValue = 1200
	ZoneSettingBrowserCacheTTLValue1800     ZoneSettingBrowserCacheTTLValue = 1800
	ZoneSettingBrowserCacheTTLValue3600     ZoneSettingBrowserCacheTTLValue = 3600
	ZoneSettingBrowserCacheTTLValue7200     ZoneSettingBrowserCacheTTLValue = 7200
	ZoneSettingBrowserCacheTTLValue10800    ZoneSettingBrowserCacheTTLValue = 10800
	ZoneSettingBrowserCacheTTLValue14400    ZoneSettingBrowserCacheTTLValue = 14400
	ZoneSettingBrowserCacheTTLValue18000    ZoneSettingBrowserCacheTTLValue = 18000
	ZoneSettingBrowserCacheTTLValue28800    ZoneSettingBrowserCacheTTLValue = 28800
	ZoneSettingBrowserCacheTTLValue43200    ZoneSettingBrowserCacheTTLValue = 43200
	ZoneSettingBrowserCacheTTLValue57600    ZoneSettingBrowserCacheTTLValue = 57600
	ZoneSettingBrowserCacheTTLValue72000    ZoneSettingBrowserCacheTTLValue = 72000
	ZoneSettingBrowserCacheTTLValue86400    ZoneSettingBrowserCacheTTLValue = 86400
	ZoneSettingBrowserCacheTTLValue172800   ZoneSettingBrowserCacheTTLValue = 172800
	ZoneSettingBrowserCacheTTLValue259200   ZoneSettingBrowserCacheTTLValue = 259200
	ZoneSettingBrowserCacheTTLValue345600   ZoneSettingBrowserCacheTTLValue = 345600
	ZoneSettingBrowserCacheTTLValue432000   ZoneSettingBrowserCacheTTLValue = 432000
	ZoneSettingBrowserCacheTTLValue691200   ZoneSettingBrowserCacheTTLValue = 691200
	ZoneSettingBrowserCacheTTLValue1382400  ZoneSettingBrowserCacheTTLValue = 1382400
	ZoneSettingBrowserCacheTTLValue2073600  ZoneSettingBrowserCacheTTLValue = 2073600
	ZoneSettingBrowserCacheTTLValue2678400  ZoneSettingBrowserCacheTTLValue = 2678400
	ZoneSettingBrowserCacheTTLValue5356800  ZoneSettingBrowserCacheTTLValue = 5356800
	ZoneSettingBrowserCacheTTLValue16070400 ZoneSettingBrowserCacheTTLValue = 16070400
	ZoneSettingBrowserCacheTTLValue31536000 ZoneSettingBrowserCacheTTLValue = 31536000
)

func (r ZoneSettingBrowserCacheTTLValue) IsKnown() bool {
	switch r {
	case ZoneSettingBrowserCacheTTLValue0, ZoneSettingBrowserCacheTTLValue30, ZoneSettingBrowserCacheTTLValue60, ZoneSettingBrowserCacheTTLValue120, ZoneSettingBrowserCacheTTLValue300, ZoneSettingBrowserCacheTTLValue1200, ZoneSettingBrowserCacheTTLValue1800, ZoneSettingBrowserCacheTTLValue3600, ZoneSettingBrowserCacheTTLValue7200, ZoneSettingBrowserCacheTTLValue10800, ZoneSettingBrowserCacheTTLValue14400, ZoneSettingBrowserCacheTTLValue18000, ZoneSettingBrowserCacheTTLValue28800, ZoneSettingBrowserCacheTTLValue43200, ZoneSettingBrowserCacheTTLValue57600, ZoneSettingBrowserCacheTTLValue72000, ZoneSettingBrowserCacheTTLValue86400, ZoneSettingBrowserCacheTTLValue172800, ZoneSettingBrowserCacheTTLValue259200, ZoneSettingBrowserCacheTTLValue345600, ZoneSettingBrowserCacheTTLValue432000, ZoneSettingBrowserCacheTTLValue691200, ZoneSettingBrowserCacheTTLValue1382400, ZoneSettingBrowserCacheTTLValue2073600, ZoneSettingBrowserCacheTTLValue2678400, ZoneSettingBrowserCacheTTLValue5356800, ZoneSettingBrowserCacheTTLValue16070400, ZoneSettingBrowserCacheTTLValue31536000:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingBrowserCacheTTLEditable bool

const (
	ZoneSettingBrowserCacheTTLEditableTrue  ZoneSettingBrowserCacheTTLEditable = true
	ZoneSettingBrowserCacheTTLEditableFalse ZoneSettingBrowserCacheTTLEditable = false
)

func (r ZoneSettingBrowserCacheTTLEditable) IsKnown() bool {
	switch r {
	case ZoneSettingBrowserCacheTTLEditableTrue, ZoneSettingBrowserCacheTTLEditableFalse:
		return true
	}
	return false
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingBrowserCacheTTLParam struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingBrowserCacheTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingBrowserCacheTTLValue] `json:"value,required"`
}

func (r ZoneSettingBrowserCacheTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingBrowserCacheTTLParam) implementsZonesSettingEditParamsItem() {}

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

func (r SettingBrowserCacheTTLEditParamsValue) IsKnown() bool {
	switch r {
	case SettingBrowserCacheTTLEditParamsValue0, SettingBrowserCacheTTLEditParamsValue30, SettingBrowserCacheTTLEditParamsValue60, SettingBrowserCacheTTLEditParamsValue120, SettingBrowserCacheTTLEditParamsValue300, SettingBrowserCacheTTLEditParamsValue1200, SettingBrowserCacheTTLEditParamsValue1800, SettingBrowserCacheTTLEditParamsValue3600, SettingBrowserCacheTTLEditParamsValue7200, SettingBrowserCacheTTLEditParamsValue10800, SettingBrowserCacheTTLEditParamsValue14400, SettingBrowserCacheTTLEditParamsValue18000, SettingBrowserCacheTTLEditParamsValue28800, SettingBrowserCacheTTLEditParamsValue43200, SettingBrowserCacheTTLEditParamsValue57600, SettingBrowserCacheTTLEditParamsValue72000, SettingBrowserCacheTTLEditParamsValue86400, SettingBrowserCacheTTLEditParamsValue172800, SettingBrowserCacheTTLEditParamsValue259200, SettingBrowserCacheTTLEditParamsValue345600, SettingBrowserCacheTTLEditParamsValue432000, SettingBrowserCacheTTLEditParamsValue691200, SettingBrowserCacheTTLEditParamsValue1382400, SettingBrowserCacheTTLEditParamsValue2073600, SettingBrowserCacheTTLEditParamsValue2678400, SettingBrowserCacheTTLEditParamsValue5356800, SettingBrowserCacheTTLEditParamsValue16070400, SettingBrowserCacheTTLEditParamsValue31536000:
		return true
	}
	return false
}

type SettingBrowserCacheTTLEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result ZoneSettingBrowserCacheTTL                     `json:"result"`
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

type SettingBrowserCacheTTLGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingBrowserCacheTTLGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result ZoneSettingBrowserCacheTTL                    `json:"result"`
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
