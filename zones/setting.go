// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	return
}

// Updates a single zone setting by the identifier
func (r *SettingService) Edit(ctx context.Context, settingID string, params SettingEditParams, opts ...option.RequestOption) (res *SettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEditResponseEnvelope
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if settingID == "" {
		err = errors.New("missing required setting_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/%s", params.ZoneID, settingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single zone setting by name
func (r *SettingService) Get(ctx context.Context, settingID string, query SettingGetParams, opts ...option.RequestOption) (res *SettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingGetResponseEnvelope
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if settingID == "" {
		err = errors.New("missing required setting_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/%s", query.ZoneID, settingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type AdvancedDDoS struct {
	// ID of the zone setting.
	ID AdvancedDDoSID `json:"id,required"`
	// Current value of the zone setting.
	Value AdvancedDDoSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AdvancedDDoSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on,nullable" format:"date-time"`
	JSON       advancedDDoSJSON `json:"-"`
}

// advancedDDoSJSON contains the JSON metadata for the struct [AdvancedDDoS]
type advancedDDoSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvancedDDoS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advancedDDoSJSON) RawJSON() string {
	return r.raw
}

func (r AdvancedDDoS) implementsZonesSettingEditResponse() {}

func (r AdvancedDDoS) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type AdvancedDDoSID string

const (
	AdvancedDDoSIDAdvancedDDoS AdvancedDDoSID = "advanced_ddos"
)

func (r AdvancedDDoSID) IsKnown() bool {
	switch r {
	case AdvancedDDoSIDAdvancedDDoS:
		return true
	}
	return false
}

// Current value of the zone setting.
type AdvancedDDoSValue string

const (
	AdvancedDDoSValueOn  AdvancedDDoSValue = "on"
	AdvancedDDoSValueOff AdvancedDDoSValue = "off"
)

func (r AdvancedDDoSValue) IsKnown() bool {
	switch r {
	case AdvancedDDoSValueOn, AdvancedDDoSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type AdvancedDDoSEditable bool

const (
	AdvancedDDoSEditableTrue  AdvancedDDoSEditable = true
	AdvancedDDoSEditableFalse AdvancedDDoSEditable = false
)

func (r AdvancedDDoSEditable) IsKnown() bool {
	switch r {
	case AdvancedDDoSEditableTrue, AdvancedDDoSEditableFalse:
		return true
	}
	return false
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type AdvancedDDoSParam struct {
	// ID of the zone setting.
	ID param.Field[AdvancedDDoSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[AdvancedDDoSValue] `json:"value,required"`
}

func (r AdvancedDDoSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AdvancedDDoSParam) implementsZonesSettingEditParamsBodyUnion() {}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type AlwaysOnline struct {
	// ID of the zone setting.
	ID AlwaysOnlineID `json:"id,required"`
	// Current value of the zone setting.
	Value AlwaysOnlineValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on,nullable" format:"date-time"`
	JSON       alwaysOnlineJSON `json:"-"`
}

// alwaysOnlineJSON contains the JSON metadata for the struct [AlwaysOnline]
type alwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alwaysOnlineJSON) RawJSON() string {
	return r.raw
}

func (r AlwaysOnline) implementsZonesSettingEditResponse() {}

func (r AlwaysOnline) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type AlwaysOnlineID string

const (
	AlwaysOnlineIDAlwaysOnline AlwaysOnlineID = "always_online"
)

func (r AlwaysOnlineID) IsKnown() bool {
	switch r {
	case AlwaysOnlineIDAlwaysOnline:
		return true
	}
	return false
}

// Current value of the zone setting.
type AlwaysOnlineValue string

const (
	AlwaysOnlineValueOn  AlwaysOnlineValue = "on"
	AlwaysOnlineValueOff AlwaysOnlineValue = "off"
)

func (r AlwaysOnlineValue) IsKnown() bool {
	switch r {
	case AlwaysOnlineValueOn, AlwaysOnlineValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type AlwaysOnlineEditable bool

const (
	AlwaysOnlineEditableTrue  AlwaysOnlineEditable = true
	AlwaysOnlineEditableFalse AlwaysOnlineEditable = false
)

func (r AlwaysOnlineEditable) IsKnown() bool {
	switch r {
	case AlwaysOnlineEditableTrue, AlwaysOnlineEditableFalse:
		return true
	}
	return false
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type AlwaysOnlineParam struct {
	// ID of the zone setting.
	ID param.Field[AlwaysOnlineID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[AlwaysOnlineValue] `json:"value,required"`
}

func (r AlwaysOnlineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AlwaysOnlineParam) implementsZonesSettingEditParamsBodyUnion() {}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type AlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID AlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value AlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time          `json:"modified_on,nullable" format:"date-time"`
	JSON       alwaysUseHTTPSJSON `json:"-"`
}

// alwaysUseHTTPSJSON contains the JSON metadata for the struct [AlwaysUseHTTPS]
type alwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r AlwaysUseHTTPS) implementsZonesSettingEditResponse() {}

func (r AlwaysUseHTTPS) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type AlwaysUseHTTPSID string

const (
	AlwaysUseHTTPSIDAlwaysUseHTTPS AlwaysUseHTTPSID = "always_use_https"
)

func (r AlwaysUseHTTPSID) IsKnown() bool {
	switch r {
	case AlwaysUseHTTPSIDAlwaysUseHTTPS:
		return true
	}
	return false
}

// Current value of the zone setting.
type AlwaysUseHTTPSValue string

const (
	AlwaysUseHTTPSValueOn  AlwaysUseHTTPSValue = "on"
	AlwaysUseHTTPSValueOff AlwaysUseHTTPSValue = "off"
)

func (r AlwaysUseHTTPSValue) IsKnown() bool {
	switch r {
	case AlwaysUseHTTPSValueOn, AlwaysUseHTTPSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type AlwaysUseHTTPSEditable bool

const (
	AlwaysUseHTTPSEditableTrue  AlwaysUseHTTPSEditable = true
	AlwaysUseHTTPSEditableFalse AlwaysUseHTTPSEditable = false
)

func (r AlwaysUseHTTPSEditable) IsKnown() bool {
	switch r {
	case AlwaysUseHTTPSEditableTrue, AlwaysUseHTTPSEditableFalse:
		return true
	}
	return false
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type AlwaysUseHTTPSParam struct {
	// ID of the zone setting.
	ID param.Field[AlwaysUseHTTPSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[AlwaysUseHTTPSValue] `json:"value,required"`
}

func (r AlwaysUseHTTPSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AlwaysUseHTTPSParam) implementsZonesSettingEditParamsBodyUnion() {}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type AutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID AutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value AutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                  `json:"modified_on,nullable" format:"date-time"`
	JSON       automaticHTTPSRewritesJSON `json:"-"`
}

// automaticHTTPSRewritesJSON contains the JSON metadata for the struct
// [AutomaticHTTPSRewrites]
type automaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

func (r AutomaticHTTPSRewrites) implementsZonesSettingEditResponse() {}

func (r AutomaticHTTPSRewrites) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type AutomaticHTTPSRewritesID string

const (
	AutomaticHTTPSRewritesIDAutomaticHTTPSRewrites AutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

func (r AutomaticHTTPSRewritesID) IsKnown() bool {
	switch r {
	case AutomaticHTTPSRewritesIDAutomaticHTTPSRewrites:
		return true
	}
	return false
}

// Current value of the zone setting.
type AutomaticHTTPSRewritesValue string

const (
	AutomaticHTTPSRewritesValueOn  AutomaticHTTPSRewritesValue = "on"
	AutomaticHTTPSRewritesValueOff AutomaticHTTPSRewritesValue = "off"
)

func (r AutomaticHTTPSRewritesValue) IsKnown() bool {
	switch r {
	case AutomaticHTTPSRewritesValueOn, AutomaticHTTPSRewritesValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type AutomaticHTTPSRewritesEditable bool

const (
	AutomaticHTTPSRewritesEditableTrue  AutomaticHTTPSRewritesEditable = true
	AutomaticHTTPSRewritesEditableFalse AutomaticHTTPSRewritesEditable = false
)

func (r AutomaticHTTPSRewritesEditable) IsKnown() bool {
	switch r {
	case AutomaticHTTPSRewritesEditableTrue, AutomaticHTTPSRewritesEditableFalse:
		return true
	}
	return false
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type AutomaticHTTPSRewritesParam struct {
	// ID of the zone setting.
	ID param.Field[AutomaticHTTPSRewritesID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[AutomaticHTTPSRewritesValue] `json:"value,required"`
}

func (r AutomaticHTTPSRewritesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AutomaticHTTPSRewritesParam) implementsZonesSettingEditParamsBodyUnion() {}

type AutomaticPlatformOptimization struct {
	// Indicates whether or not
	// [cache by device type](https://developers.cloudflare.com/automatic-platform-optimization/reference/cache-device-type/)
	// is enabled.
	CacheByDeviceType bool `json:"cache_by_device_type,required"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf bool `json:"cf,required"`
	// Indicates whether or not Automatic Platform Optimization is enabled.
	Enabled bool `json:"enabled,required"`
	// An array of hostnames where Automatic Platform Optimization for WordPress is
	// activated.
	Hostnames []string `json:"hostnames,required" format:"hostname"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress bool `json:"wordpress,required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin bool                              `json:"wp_plugin,required"`
	JSON     automaticPlatformOptimizationJSON `json:"-"`
}

// automaticPlatformOptimizationJSON contains the JSON metadata for the struct
// [AutomaticPlatformOptimization]
type automaticPlatformOptimizationJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automaticPlatformOptimizationJSON) RawJSON() string {
	return r.raw
}

type AutomaticPlatformOptimizationParam struct {
	// Indicates whether or not
	// [cache by device type](https://developers.cloudflare.com/automatic-platform-optimization/reference/cache-device-type/)
	// is enabled.
	CacheByDeviceType param.Field[bool] `json:"cache_by_device_type,required"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf param.Field[bool] `json:"cf,required"`
	// Indicates whether or not Automatic Platform Optimization is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// An array of hostnames where Automatic Platform Optimization for WordPress is
	// activated.
	Hostnames param.Field[[]string] `json:"hostnames,required" format:"hostname"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress param.Field[bool] `json:"wordpress,required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin param.Field[bool] `json:"wp_plugin,required"`
}

func (r AutomaticPlatformOptimizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type Brotli struct {
	// ID of the zone setting.
	ID BrotliID `json:"id,required"`
	// Current value of the zone setting.
	Value BrotliValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable BrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time  `json:"modified_on,nullable" format:"date-time"`
	JSON       brotliJSON `json:"-"`
}

// brotliJSON contains the JSON metadata for the struct [Brotli]
type brotliJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Brotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brotliJSON) RawJSON() string {
	return r.raw
}

func (r Brotli) implementsZonesSettingEditResponse() {}

func (r Brotli) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type BrotliID string

const (
	BrotliIDBrotli BrotliID = "brotli"
)

func (r BrotliID) IsKnown() bool {
	switch r {
	case BrotliIDBrotli:
		return true
	}
	return false
}

// Current value of the zone setting.
type BrotliValue string

const (
	BrotliValueOff BrotliValue = "off"
	BrotliValueOn  BrotliValue = "on"
)

func (r BrotliValue) IsKnown() bool {
	switch r {
	case BrotliValueOff, BrotliValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type BrotliEditable bool

const (
	BrotliEditableTrue  BrotliEditable = true
	BrotliEditableFalse BrotliEditable = false
)

func (r BrotliEditable) IsKnown() bool {
	switch r {
	case BrotliEditableTrue, BrotliEditableFalse:
		return true
	}
	return false
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type BrotliParam struct {
	// ID of the zone setting.
	ID param.Field[BrotliID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[BrotliValue] `json:"value,required"`
}

func (r BrotliParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BrotliParam) implementsZonesSettingEditParamsBodyUnion() {}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type BrowserCacheTTL struct {
	// ID of the zone setting.
	ID BrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value BrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable BrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       browserCacheTTLJSON `json:"-"`
}

// browserCacheTTLJSON contains the JSON metadata for the struct [BrowserCacheTTL]
type browserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r BrowserCacheTTL) implementsZonesSettingEditResponse() {}

func (r BrowserCacheTTL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type BrowserCacheTTLID string

const (
	BrowserCacheTTLIDBrowserCacheTTL BrowserCacheTTLID = "browser_cache_ttl"
)

func (r BrowserCacheTTLID) IsKnown() bool {
	switch r {
	case BrowserCacheTTLIDBrowserCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type BrowserCacheTTLValue float64

const (
	BrowserCacheTTLValue0        BrowserCacheTTLValue = 0
	BrowserCacheTTLValue30       BrowserCacheTTLValue = 30
	BrowserCacheTTLValue60       BrowserCacheTTLValue = 60
	BrowserCacheTTLValue120      BrowserCacheTTLValue = 120
	BrowserCacheTTLValue300      BrowserCacheTTLValue = 300
	BrowserCacheTTLValue1200     BrowserCacheTTLValue = 1200
	BrowserCacheTTLValue1800     BrowserCacheTTLValue = 1800
	BrowserCacheTTLValue3600     BrowserCacheTTLValue = 3600
	BrowserCacheTTLValue7200     BrowserCacheTTLValue = 7200
	BrowserCacheTTLValue10800    BrowserCacheTTLValue = 10800
	BrowserCacheTTLValue14400    BrowserCacheTTLValue = 14400
	BrowserCacheTTLValue18000    BrowserCacheTTLValue = 18000
	BrowserCacheTTLValue28800    BrowserCacheTTLValue = 28800
	BrowserCacheTTLValue43200    BrowserCacheTTLValue = 43200
	BrowserCacheTTLValue57600    BrowserCacheTTLValue = 57600
	BrowserCacheTTLValue72000    BrowserCacheTTLValue = 72000
	BrowserCacheTTLValue86400    BrowserCacheTTLValue = 86400
	BrowserCacheTTLValue172800   BrowserCacheTTLValue = 172800
	BrowserCacheTTLValue259200   BrowserCacheTTLValue = 259200
	BrowserCacheTTLValue345600   BrowserCacheTTLValue = 345600
	BrowserCacheTTLValue432000   BrowserCacheTTLValue = 432000
	BrowserCacheTTLValue691200   BrowserCacheTTLValue = 691200
	BrowserCacheTTLValue1382400  BrowserCacheTTLValue = 1382400
	BrowserCacheTTLValue2073600  BrowserCacheTTLValue = 2073600
	BrowserCacheTTLValue2678400  BrowserCacheTTLValue = 2678400
	BrowserCacheTTLValue5356800  BrowserCacheTTLValue = 5356800
	BrowserCacheTTLValue16070400 BrowserCacheTTLValue = 16070400
	BrowserCacheTTLValue31536000 BrowserCacheTTLValue = 31536000
)

func (r BrowserCacheTTLValue) IsKnown() bool {
	switch r {
	case BrowserCacheTTLValue0, BrowserCacheTTLValue30, BrowserCacheTTLValue60, BrowserCacheTTLValue120, BrowserCacheTTLValue300, BrowserCacheTTLValue1200, BrowserCacheTTLValue1800, BrowserCacheTTLValue3600, BrowserCacheTTLValue7200, BrowserCacheTTLValue10800, BrowserCacheTTLValue14400, BrowserCacheTTLValue18000, BrowserCacheTTLValue28800, BrowserCacheTTLValue43200, BrowserCacheTTLValue57600, BrowserCacheTTLValue72000, BrowserCacheTTLValue86400, BrowserCacheTTLValue172800, BrowserCacheTTLValue259200, BrowserCacheTTLValue345600, BrowserCacheTTLValue432000, BrowserCacheTTLValue691200, BrowserCacheTTLValue1382400, BrowserCacheTTLValue2073600, BrowserCacheTTLValue2678400, BrowserCacheTTLValue5356800, BrowserCacheTTLValue16070400, BrowserCacheTTLValue31536000:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type BrowserCacheTTLEditable bool

const (
	BrowserCacheTTLEditableTrue  BrowserCacheTTLEditable = true
	BrowserCacheTTLEditableFalse BrowserCacheTTLEditable = false
)

func (r BrowserCacheTTLEditable) IsKnown() bool {
	switch r {
	case BrowserCacheTTLEditableTrue, BrowserCacheTTLEditableFalse:
		return true
	}
	return false
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type BrowserCacheTTLParam struct {
	// ID of the zone setting.
	ID param.Field[BrowserCacheTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[BrowserCacheTTLValue] `json:"value,required"`
}

func (r BrowserCacheTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BrowserCacheTTLParam) implementsZonesSettingEditParamsBodyUnion() {}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type BrowserCheck struct {
	// ID of the zone setting.
	ID BrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value BrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable BrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on,nullable" format:"date-time"`
	JSON       browserCheckJSON `json:"-"`
}

// browserCheckJSON contains the JSON metadata for the struct [BrowserCheck]
type browserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserCheckJSON) RawJSON() string {
	return r.raw
}

func (r BrowserCheck) implementsZonesSettingEditResponse() {}

func (r BrowserCheck) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type BrowserCheckID string

const (
	BrowserCheckIDBrowserCheck BrowserCheckID = "browser_check"
)

func (r BrowserCheckID) IsKnown() bool {
	switch r {
	case BrowserCheckIDBrowserCheck:
		return true
	}
	return false
}

// Current value of the zone setting.
type BrowserCheckValue string

const (
	BrowserCheckValueOn  BrowserCheckValue = "on"
	BrowserCheckValueOff BrowserCheckValue = "off"
)

func (r BrowserCheckValue) IsKnown() bool {
	switch r {
	case BrowserCheckValueOn, BrowserCheckValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type BrowserCheckEditable bool

const (
	BrowserCheckEditableTrue  BrowserCheckEditable = true
	BrowserCheckEditableFalse BrowserCheckEditable = false
)

func (r BrowserCheckEditable) IsKnown() bool {
	switch r {
	case BrowserCheckEditableTrue, BrowserCheckEditableFalse:
		return true
	}
	return false
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type BrowserCheckParam struct {
	// ID of the zone setting.
	ID param.Field[BrowserCheckID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[BrowserCheckValue] `json:"value,required"`
}

func (r BrowserCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BrowserCheckParam) implementsZonesSettingEditParamsBodyUnion() {}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type CacheLevel struct {
	// ID of the zone setting.
	ID CacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value CacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable CacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time      `json:"modified_on,nullable" format:"date-time"`
	JSON       cacheLevelJSON `json:"-"`
}

// cacheLevelJSON contains the JSON metadata for the struct [CacheLevel]
type cacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheLevelJSON) RawJSON() string {
	return r.raw
}

func (r CacheLevel) implementsZonesSettingEditResponse() {}

func (r CacheLevel) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type CacheLevelID string

const (
	CacheLevelIDCacheLevel CacheLevelID = "cache_level"
)

func (r CacheLevelID) IsKnown() bool {
	switch r {
	case CacheLevelIDCacheLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type CacheLevelValue string

const (
	CacheLevelValueAggressive CacheLevelValue = "aggressive"
	CacheLevelValueBasic      CacheLevelValue = "basic"
	CacheLevelValueSimplified CacheLevelValue = "simplified"
)

func (r CacheLevelValue) IsKnown() bool {
	switch r {
	case CacheLevelValueAggressive, CacheLevelValueBasic, CacheLevelValueSimplified:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type CacheLevelEditable bool

const (
	CacheLevelEditableTrue  CacheLevelEditable = true
	CacheLevelEditableFalse CacheLevelEditable = false
)

func (r CacheLevelEditable) IsKnown() bool {
	switch r {
	case CacheLevelEditableTrue, CacheLevelEditableFalse:
		return true
	}
	return false
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type CacheLevelParam struct {
	// ID of the zone setting.
	ID param.Field[CacheLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[CacheLevelValue] `json:"value,required"`
}

func (r CacheLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheLevelParam) implementsZonesSettingEditParamsBodyUnion() {}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ChallengeTTL struct {
	// ID of the zone setting.
	ID ChallengeTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ChallengeTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on,nullable" format:"date-time"`
	JSON       challengeTTLJSON `json:"-"`
}

// challengeTTLJSON contains the JSON metadata for the struct [ChallengeTTL]
type challengeTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r challengeTTLJSON) RawJSON() string {
	return r.raw
}

func (r ChallengeTTL) implementsZonesSettingEditResponse() {}

func (r ChallengeTTL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ChallengeTTLID string

const (
	ChallengeTTLIDChallengeTTL ChallengeTTLID = "challenge_ttl"
)

func (r ChallengeTTLID) IsKnown() bool {
	switch r {
	case ChallengeTTLIDChallengeTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type ChallengeTTLValue float64

const (
	ChallengeTTLValue300      ChallengeTTLValue = 300
	ChallengeTTLValue900      ChallengeTTLValue = 900
	ChallengeTTLValue1800     ChallengeTTLValue = 1800
	ChallengeTTLValue2700     ChallengeTTLValue = 2700
	ChallengeTTLValue3600     ChallengeTTLValue = 3600
	ChallengeTTLValue7200     ChallengeTTLValue = 7200
	ChallengeTTLValue10800    ChallengeTTLValue = 10800
	ChallengeTTLValue14400    ChallengeTTLValue = 14400
	ChallengeTTLValue28800    ChallengeTTLValue = 28800
	ChallengeTTLValue57600    ChallengeTTLValue = 57600
	ChallengeTTLValue86400    ChallengeTTLValue = 86400
	ChallengeTTLValue604800   ChallengeTTLValue = 604800
	ChallengeTTLValue2592000  ChallengeTTLValue = 2592000
	ChallengeTTLValue31536000 ChallengeTTLValue = 31536000
)

func (r ChallengeTTLValue) IsKnown() bool {
	switch r {
	case ChallengeTTLValue300, ChallengeTTLValue900, ChallengeTTLValue1800, ChallengeTTLValue2700, ChallengeTTLValue3600, ChallengeTTLValue7200, ChallengeTTLValue10800, ChallengeTTLValue14400, ChallengeTTLValue28800, ChallengeTTLValue57600, ChallengeTTLValue86400, ChallengeTTLValue604800, ChallengeTTLValue2592000, ChallengeTTLValue31536000:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ChallengeTTLEditable bool

const (
	ChallengeTTLEditableTrue  ChallengeTTLEditable = true
	ChallengeTTLEditableFalse ChallengeTTLEditable = false
)

func (r ChallengeTTLEditable) IsKnown() bool {
	switch r {
	case ChallengeTTLEditableTrue, ChallengeTTLEditableFalse:
		return true
	}
	return false
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ChallengeTTLParam struct {
	// ID of the zone setting.
	ID param.Field[ChallengeTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ChallengeTTLValue] `json:"value,required"`
}

func (r ChallengeTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChallengeTTLParam) implementsZonesSettingEditParamsBodyUnion() {}

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type Ciphers struct {
	// ID of the zone setting.
	ID CiphersID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable CiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time   `json:"modified_on,nullable" format:"date-time"`
	JSON       ciphersJSON `json:"-"`
}

// ciphersJSON contains the JSON metadata for the struct [Ciphers]
type ciphersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ciphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ciphersJSON) RawJSON() string {
	return r.raw
}

func (r Ciphers) implementsZonesSettingEditResponse() {}

func (r Ciphers) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type CiphersID string

const (
	CiphersIDCiphers CiphersID = "ciphers"
)

func (r CiphersID) IsKnown() bool {
	switch r {
	case CiphersIDCiphers:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type CiphersEditable bool

const (
	CiphersEditableTrue  CiphersEditable = true
	CiphersEditableFalse CiphersEditable = false
)

func (r CiphersEditable) IsKnown() bool {
	switch r {
	case CiphersEditableTrue, CiphersEditableFalse:
		return true
	}
	return false
}

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type CiphersParam struct {
	// ID of the zone setting.
	ID param.Field[CiphersID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[[]string] `json:"value,required"`
}

func (r CiphersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CiphersParam) implementsZonesSettingEditParamsBodyUnion() {}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type DevelopmentMode struct {
	// ID of the zone setting.
	ID DevelopmentModeID `json:"id,required"`
	// Current value of the zone setting.
	Value DevelopmentModeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable DevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64             `json:"time_remaining"`
	JSON          developmentModeJSON `json:"-"`
}

// developmentModeJSON contains the JSON metadata for the struct [DevelopmentMode]
type developmentModeJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r developmentModeJSON) RawJSON() string {
	return r.raw
}

func (r DevelopmentMode) implementsZonesSettingEditResponse() {}

func (r DevelopmentMode) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type DevelopmentModeID string

const (
	DevelopmentModeIDDevelopmentMode DevelopmentModeID = "development_mode"
)

func (r DevelopmentModeID) IsKnown() bool {
	switch r {
	case DevelopmentModeIDDevelopmentMode:
		return true
	}
	return false
}

// Current value of the zone setting.
type DevelopmentModeValue string

const (
	DevelopmentModeValueOn  DevelopmentModeValue = "on"
	DevelopmentModeValueOff DevelopmentModeValue = "off"
)

func (r DevelopmentModeValue) IsKnown() bool {
	switch r {
	case DevelopmentModeValueOn, DevelopmentModeValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type DevelopmentModeEditable bool

const (
	DevelopmentModeEditableTrue  DevelopmentModeEditable = true
	DevelopmentModeEditableFalse DevelopmentModeEditable = false
)

func (r DevelopmentModeEditable) IsKnown() bool {
	switch r {
	case DevelopmentModeEditableTrue, DevelopmentModeEditableFalse:
		return true
	}
	return false
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type DevelopmentModeParam struct {
	// ID of the zone setting.
	ID param.Field[DevelopmentModeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[DevelopmentModeValue] `json:"value,required"`
}

func (r DevelopmentModeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DevelopmentModeParam) implementsZonesSettingEditParamsBodyUnion() {}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type EarlyHints struct {
	// ID of the zone setting.
	ID EarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value EarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable EarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time      `json:"modified_on,nullable" format:"date-time"`
	JSON       earlyHintsJSON `json:"-"`
}

// earlyHintsJSON contains the JSON metadata for the struct [EarlyHints]
type earlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r earlyHintsJSON) RawJSON() string {
	return r.raw
}

func (r EarlyHints) implementsZonesSettingEditResponse() {}

func (r EarlyHints) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type EarlyHintsID string

const (
	EarlyHintsIDEarlyHints EarlyHintsID = "early_hints"
)

func (r EarlyHintsID) IsKnown() bool {
	switch r {
	case EarlyHintsIDEarlyHints:
		return true
	}
	return false
}

// Current value of the zone setting.
type EarlyHintsValue string

const (
	EarlyHintsValueOn  EarlyHintsValue = "on"
	EarlyHintsValueOff EarlyHintsValue = "off"
)

func (r EarlyHintsValue) IsKnown() bool {
	switch r {
	case EarlyHintsValueOn, EarlyHintsValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type EarlyHintsEditable bool

const (
	EarlyHintsEditableTrue  EarlyHintsEditable = true
	EarlyHintsEditableFalse EarlyHintsEditable = false
)

func (r EarlyHintsEditable) IsKnown() bool {
	switch r {
	case EarlyHintsEditableTrue, EarlyHintsEditableFalse:
		return true
	}
	return false
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type EarlyHintsParam struct {
	// ID of the zone setting.
	ID param.Field[EarlyHintsID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[EarlyHintsValue] `json:"value,required"`
}

func (r EarlyHintsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EarlyHintsParam) implementsZonesSettingEditParamsBodyUnion() {}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type EmailObfuscation struct {
	// ID of the zone setting.
	ID EmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value EmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable EmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time            `json:"modified_on,nullable" format:"date-time"`
	JSON       emailObfuscationJSON `json:"-"`
}

// emailObfuscationJSON contains the JSON metadata for the struct
// [EmailObfuscation]
type emailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r EmailObfuscation) implementsZonesSettingEditResponse() {}

func (r EmailObfuscation) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type EmailObfuscationID string

const (
	EmailObfuscationIDEmailObfuscation EmailObfuscationID = "email_obfuscation"
)

func (r EmailObfuscationID) IsKnown() bool {
	switch r {
	case EmailObfuscationIDEmailObfuscation:
		return true
	}
	return false
}

// Current value of the zone setting.
type EmailObfuscationValue string

const (
	EmailObfuscationValueOn  EmailObfuscationValue = "on"
	EmailObfuscationValueOff EmailObfuscationValue = "off"
)

func (r EmailObfuscationValue) IsKnown() bool {
	switch r {
	case EmailObfuscationValueOn, EmailObfuscationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type EmailObfuscationEditable bool

const (
	EmailObfuscationEditableTrue  EmailObfuscationEditable = true
	EmailObfuscationEditableFalse EmailObfuscationEditable = false
)

func (r EmailObfuscationEditable) IsKnown() bool {
	switch r {
	case EmailObfuscationEditableTrue, EmailObfuscationEditableFalse:
		return true
	}
	return false
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type EmailObfuscationParam struct {
	// ID of the zone setting.
	ID param.Field[EmailObfuscationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[EmailObfuscationValue] `json:"value,required"`
}

func (r EmailObfuscationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailObfuscationParam) implementsZonesSettingEditParamsBodyUnion() {}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type H2Prioritization struct {
	// ID of the zone setting.
	ID H2PrioritizationID `json:"id,required"`
	// Current value of the zone setting.
	Value H2PrioritizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable H2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time            `json:"modified_on,nullable" format:"date-time"`
	JSON       h2PrioritizationJSON `json:"-"`
}

// h2PrioritizationJSON contains the JSON metadata for the struct
// [H2Prioritization]
type h2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *H2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r h2PrioritizationJSON) RawJSON() string {
	return r.raw
}

func (r H2Prioritization) implementsZonesSettingEditResponse() {}

func (r H2Prioritization) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type H2PrioritizationID string

const (
	H2PrioritizationIDH2Prioritization H2PrioritizationID = "h2_prioritization"
)

func (r H2PrioritizationID) IsKnown() bool {
	switch r {
	case H2PrioritizationIDH2Prioritization:
		return true
	}
	return false
}

// Current value of the zone setting.
type H2PrioritizationValue string

const (
	H2PrioritizationValueOn     H2PrioritizationValue = "on"
	H2PrioritizationValueOff    H2PrioritizationValue = "off"
	H2PrioritizationValueCustom H2PrioritizationValue = "custom"
)

func (r H2PrioritizationValue) IsKnown() bool {
	switch r {
	case H2PrioritizationValueOn, H2PrioritizationValueOff, H2PrioritizationValueCustom:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type H2PrioritizationEditable bool

const (
	H2PrioritizationEditableTrue  H2PrioritizationEditable = true
	H2PrioritizationEditableFalse H2PrioritizationEditable = false
)

func (r H2PrioritizationEditable) IsKnown() bool {
	switch r {
	case H2PrioritizationEditableTrue, H2PrioritizationEditableFalse:
		return true
	}
	return false
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type H2PrioritizationParam struct {
	// ID of the zone setting.
	ID param.Field[H2PrioritizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[H2PrioritizationValue] `json:"value,required"`
}

func (r H2PrioritizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r H2PrioritizationParam) implementsZonesSettingEditParamsBodyUnion() {}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type HotlinkProtection struct {
	// ID of the zone setting.
	ID HotlinkProtectionID `json:"id,required"`
	// Current value of the zone setting.
	Value HotlinkProtectionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable HotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on,nullable" format:"date-time"`
	JSON       hotlinkProtectionJSON `json:"-"`
}

// hotlinkProtectionJSON contains the JSON metadata for the struct
// [HotlinkProtection]
type hotlinkProtectionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hotlinkProtectionJSON) RawJSON() string {
	return r.raw
}

func (r HotlinkProtection) implementsZonesSettingEditResponse() {}

func (r HotlinkProtection) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type HotlinkProtectionID string

const (
	HotlinkProtectionIDHotlinkProtection HotlinkProtectionID = "hotlink_protection"
)

func (r HotlinkProtectionID) IsKnown() bool {
	switch r {
	case HotlinkProtectionIDHotlinkProtection:
		return true
	}
	return false
}

// Current value of the zone setting.
type HotlinkProtectionValue string

const (
	HotlinkProtectionValueOn  HotlinkProtectionValue = "on"
	HotlinkProtectionValueOff HotlinkProtectionValue = "off"
)

func (r HotlinkProtectionValue) IsKnown() bool {
	switch r {
	case HotlinkProtectionValueOn, HotlinkProtectionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type HotlinkProtectionEditable bool

const (
	HotlinkProtectionEditableTrue  HotlinkProtectionEditable = true
	HotlinkProtectionEditableFalse HotlinkProtectionEditable = false
)

func (r HotlinkProtectionEditable) IsKnown() bool {
	switch r {
	case HotlinkProtectionEditableTrue, HotlinkProtectionEditableFalse:
		return true
	}
	return false
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type HotlinkProtectionParam struct {
	// ID of the zone setting.
	ID param.Field[HotlinkProtectionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[HotlinkProtectionValue] `json:"value,required"`
}

func (r HotlinkProtectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HotlinkProtectionParam) implementsZonesSettingEditParamsBodyUnion() {}

// HTTP2 enabled for this zone.
type HTTP2 struct {
	// ID of the zone setting.
	ID HTTP2ID `json:"id,required"`
	// Current value of the zone setting.
	Value HTTP2Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable HTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       http2JSON `json:"-"`
}

// http2JSON contains the JSON metadata for the struct [HTTP2]
type http2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r http2JSON) RawJSON() string {
	return r.raw
}

func (r HTTP2) implementsZonesSettingEditResponse() {}

func (r HTTP2) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type HTTP2ID string

const (
	HTTP2IDHTTP2 HTTP2ID = "http2"
)

func (r HTTP2ID) IsKnown() bool {
	switch r {
	case HTTP2IDHTTP2:
		return true
	}
	return false
}

// Current value of the zone setting.
type HTTP2Value string

const (
	HTTP2ValueOn  HTTP2Value = "on"
	HTTP2ValueOff HTTP2Value = "off"
)

func (r HTTP2Value) IsKnown() bool {
	switch r {
	case HTTP2ValueOn, HTTP2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type HTTP2Editable bool

const (
	HTTP2EditableTrue  HTTP2Editable = true
	HTTP2EditableFalse HTTP2Editable = false
)

func (r HTTP2Editable) IsKnown() bool {
	switch r {
	case HTTP2EditableTrue, HTTP2EditableFalse:
		return true
	}
	return false
}

// HTTP2 enabled for this zone.
type HTTP2Param struct {
	// ID of the zone setting.
	ID param.Field[HTTP2ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[HTTP2Value] `json:"value,required"`
}

func (r HTTP2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HTTP2Param) implementsZonesSettingEditParamsBodyUnion() {}

// HTTP3 enabled for this zone.
type HTTP3 struct {
	// ID of the zone setting.
	ID HTTP3ID `json:"id,required"`
	// Current value of the zone setting.
	Value HTTP3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable HTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       http3JSON `json:"-"`
}

// http3JSON contains the JSON metadata for the struct [HTTP3]
type http3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r http3JSON) RawJSON() string {
	return r.raw
}

func (r HTTP3) implementsZonesSettingEditResponse() {}

func (r HTTP3) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type HTTP3ID string

const (
	HTTP3IDHTTP3 HTTP3ID = "http3"
)

func (r HTTP3ID) IsKnown() bool {
	switch r {
	case HTTP3IDHTTP3:
		return true
	}
	return false
}

// Current value of the zone setting.
type HTTP3Value string

const (
	HTTP3ValueOn  HTTP3Value = "on"
	HTTP3ValueOff HTTP3Value = "off"
)

func (r HTTP3Value) IsKnown() bool {
	switch r {
	case HTTP3ValueOn, HTTP3ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type HTTP3Editable bool

const (
	HTTP3EditableTrue  HTTP3Editable = true
	HTTP3EditableFalse HTTP3Editable = false
)

func (r HTTP3Editable) IsKnown() bool {
	switch r {
	case HTTP3EditableTrue, HTTP3EditableFalse:
		return true
	}
	return false
}

// HTTP3 enabled for this zone.
type HTTP3Param struct {
	// ID of the zone setting.
	ID param.Field[HTTP3ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[HTTP3Value] `json:"value,required"`
}

func (r HTTP3Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HTTP3Param) implementsZonesSettingEditParamsBodyUnion() {}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ImageResizing struct {
	// ID of the zone setting.
	ID ImageResizingID `json:"id,required"`
	// Current value of the zone setting.
	Value ImageResizingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on,nullable" format:"date-time"`
	JSON       imageResizingJSON `json:"-"`
}

// imageResizingJSON contains the JSON metadata for the struct [ImageResizing]
type imageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageResizingJSON) RawJSON() string {
	return r.raw
}

func (r ImageResizing) implementsZonesSettingEditResponse() {}

func (r ImageResizing) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ImageResizingID string

const (
	ImageResizingIDImageResizing ImageResizingID = "image_resizing"
)

func (r ImageResizingID) IsKnown() bool {
	switch r {
	case ImageResizingIDImageResizing:
		return true
	}
	return false
}

// Current value of the zone setting.
type ImageResizingValue string

const (
	ImageResizingValueOn   ImageResizingValue = "on"
	ImageResizingValueOff  ImageResizingValue = "off"
	ImageResizingValueOpen ImageResizingValue = "open"
)

func (r ImageResizingValue) IsKnown() bool {
	switch r {
	case ImageResizingValueOn, ImageResizingValueOff, ImageResizingValueOpen:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ImageResizingEditable bool

const (
	ImageResizingEditableTrue  ImageResizingEditable = true
	ImageResizingEditableFalse ImageResizingEditable = false
)

func (r ImageResizingEditable) IsKnown() bool {
	switch r {
	case ImageResizingEditableTrue, ImageResizingEditableFalse:
		return true
	}
	return false
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ImageResizingParam struct {
	// ID of the zone setting.
	ID param.Field[ImageResizingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ImageResizingValue] `json:"value,required"`
}

func (r ImageResizingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ImageResizingParam) implementsZonesSettingEditParamsBodyUnion() {}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type IPGeolocation struct {
	// ID of the zone setting.
	ID IPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value IPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable IPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on,nullable" format:"date-time"`
	JSON       ipGeolocationJSON `json:"-"`
}

// ipGeolocationJSON contains the JSON metadata for the struct [IPGeolocation]
type ipGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipGeolocationJSON) RawJSON() string {
	return r.raw
}

func (r IPGeolocation) implementsZonesSettingEditResponse() {}

func (r IPGeolocation) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type IPGeolocationID string

const (
	IPGeolocationIDIPGeolocation IPGeolocationID = "ip_geolocation"
)

func (r IPGeolocationID) IsKnown() bool {
	switch r {
	case IPGeolocationIDIPGeolocation:
		return true
	}
	return false
}

// Current value of the zone setting.
type IPGeolocationValue string

const (
	IPGeolocationValueOn  IPGeolocationValue = "on"
	IPGeolocationValueOff IPGeolocationValue = "off"
)

func (r IPGeolocationValue) IsKnown() bool {
	switch r {
	case IPGeolocationValueOn, IPGeolocationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type IPGeolocationEditable bool

const (
	IPGeolocationEditableTrue  IPGeolocationEditable = true
	IPGeolocationEditableFalse IPGeolocationEditable = false
)

func (r IPGeolocationEditable) IsKnown() bool {
	switch r {
	case IPGeolocationEditableTrue, IPGeolocationEditableFalse:
		return true
	}
	return false
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type IPGeolocationParam struct {
	// ID of the zone setting.
	ID param.Field[IPGeolocationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[IPGeolocationValue] `json:"value,required"`
}

func (r IPGeolocationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPGeolocationParam) implementsZonesSettingEditParamsBodyUnion() {}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type IPV6 struct {
	// ID of the zone setting.
	ID IPV6ID `json:"id,required"`
	// Current value of the zone setting.
	Value IPV6Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable IPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       ipv6JSON  `json:"-"`
}

// ipv6JSON contains the JSON metadata for the struct [IPV6]
type ipv6JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipv6JSON) RawJSON() string {
	return r.raw
}

func (r IPV6) implementsZonesSettingEditResponse() {}

func (r IPV6) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type IPV6ID string

const (
	IPV6IDIPV6 IPV6ID = "ipv6"
)

func (r IPV6ID) IsKnown() bool {
	switch r {
	case IPV6IDIPV6:
		return true
	}
	return false
}

// Current value of the zone setting.
type IPV6Value string

const (
	IPV6ValueOff IPV6Value = "off"
	IPV6ValueOn  IPV6Value = "on"
)

func (r IPV6Value) IsKnown() bool {
	switch r {
	case IPV6ValueOff, IPV6ValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type IPV6Editable bool

const (
	IPV6EditableTrue  IPV6Editable = true
	IPV6EditableFalse IPV6Editable = false
)

func (r IPV6Editable) IsKnown() bool {
	switch r {
	case IPV6EditableTrue, IPV6EditableFalse:
		return true
	}
	return false
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type IPV6Param struct {
	// ID of the zone setting.
	ID param.Field[IPV6ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[IPV6Value] `json:"value,required"`
}

func (r IPV6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPV6Param) implementsZonesSettingEditParamsBodyUnion() {}

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type MinTLSVersion struct {
	// ID of the zone setting.
	ID MinTLSVersionID `json:"id,required"`
	// Current value of the zone setting.
	Value MinTLSVersionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on,nullable" format:"date-time"`
	JSON       minTLSVersionJSON `json:"-"`
}

// minTLSVersionJSON contains the JSON metadata for the struct [MinTLSVersion]
type minTLSVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minTLSVersionJSON) RawJSON() string {
	return r.raw
}

func (r MinTLSVersion) implementsZonesSettingEditResponse() {}

func (r MinTLSVersion) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type MinTLSVersionID string

const (
	MinTLSVersionIDMinTLSVersion MinTLSVersionID = "min_tls_version"
)

func (r MinTLSVersionID) IsKnown() bool {
	switch r {
	case MinTLSVersionIDMinTLSVersion:
		return true
	}
	return false
}

// Current value of the zone setting.
type MinTLSVersionValue string

const (
	MinTLSVersionValue1_0 MinTLSVersionValue = "1.0"
	MinTLSVersionValue1_1 MinTLSVersionValue = "1.1"
	MinTLSVersionValue1_2 MinTLSVersionValue = "1.2"
	MinTLSVersionValue1_3 MinTLSVersionValue = "1.3"
)

func (r MinTLSVersionValue) IsKnown() bool {
	switch r {
	case MinTLSVersionValue1_0, MinTLSVersionValue1_1, MinTLSVersionValue1_2, MinTLSVersionValue1_3:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type MinTLSVersionEditable bool

const (
	MinTLSVersionEditableTrue  MinTLSVersionEditable = true
	MinTLSVersionEditableFalse MinTLSVersionEditable = false
)

func (r MinTLSVersionEditable) IsKnown() bool {
	switch r {
	case MinTLSVersionEditableTrue, MinTLSVersionEditableFalse:
		return true
	}
	return false
}

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type MinTLSVersionParam struct {
	// ID of the zone setting.
	ID param.Field[MinTLSVersionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[MinTLSVersionValue] `json:"value,required"`
}

func (r MinTLSVersionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MinTLSVersionParam) implementsZonesSettingEditParamsBodyUnion() {}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type Minify struct {
	// Zone setting identifier.
	ID MinifyID `json:"id,required"`
	// Current value of the zone setting.
	Value MinifyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time  `json:"modified_on,nullable" format:"date-time"`
	JSON       minifyJSON `json:"-"`
}

// minifyJSON contains the JSON metadata for the struct [Minify]
type minifyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Minify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minifyJSON) RawJSON() string {
	return r.raw
}

func (r Minify) implementsZonesSettingEditResponse() {}

func (r Minify) implementsZonesSettingGetResponse() {}

// Zone setting identifier.
type MinifyID string

const (
	MinifyIDMinify MinifyID = "minify"
)

func (r MinifyID) IsKnown() bool {
	switch r {
	case MinifyIDMinify:
		return true
	}
	return false
}

// Current value of the zone setting.
type MinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css MinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML MinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	JS   MinifyValueJS   `json:"js"`
	JSON minifyValueJSON `json:"-"`
}

// minifyValueJSON contains the JSON metadata for the struct [MinifyValue]
type minifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	JS          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r minifyValueJSON) RawJSON() string {
	return r.raw
}

// Automatically minify all CSS files for your website.
type MinifyValueCss string

const (
	MinifyValueCssOn  MinifyValueCss = "on"
	MinifyValueCssOff MinifyValueCss = "off"
)

func (r MinifyValueCss) IsKnown() bool {
	switch r {
	case MinifyValueCssOn, MinifyValueCssOff:
		return true
	}
	return false
}

// Automatically minify all HTML files for your website.
type MinifyValueHTML string

const (
	MinifyValueHTMLOn  MinifyValueHTML = "on"
	MinifyValueHTMLOff MinifyValueHTML = "off"
)

func (r MinifyValueHTML) IsKnown() bool {
	switch r {
	case MinifyValueHTMLOn, MinifyValueHTMLOff:
		return true
	}
	return false
}

// Automatically minify all JavaScript files for your website.
type MinifyValueJS string

const (
	MinifyValueJSOn  MinifyValueJS = "on"
	MinifyValueJSOff MinifyValueJS = "off"
)

func (r MinifyValueJS) IsKnown() bool {
	switch r {
	case MinifyValueJSOn, MinifyValueJSOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type MinifyEditable bool

const (
	MinifyEditableTrue  MinifyEditable = true
	MinifyEditableFalse MinifyEditable = false
)

func (r MinifyEditable) IsKnown() bool {
	switch r {
	case MinifyEditableTrue, MinifyEditableFalse:
		return true
	}
	return false
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type MinifyParam struct {
	// Zone setting identifier.
	ID param.Field[MinifyID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[MinifyValueParam] `json:"value,required"`
}

func (r MinifyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MinifyParam) implementsZonesSettingEditParamsBodyUnion() {}

// Current value of the zone setting.
type MinifyValueParam struct {
	// Automatically minify all CSS files for your website.
	Css param.Field[MinifyValueCss] `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML param.Field[MinifyValueHTML] `json:"html"`
	// Automatically minify all JavaScript files for your website.
	JS param.Field[MinifyValueJS] `json:"js"`
}

func (r MinifyValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type Mirage struct {
	// ID of the zone setting.
	ID MirageID `json:"id,required"`
	// Current value of the zone setting.
	Value MirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time  `json:"modified_on,nullable" format:"date-time"`
	JSON       mirageJSON `json:"-"`
}

// mirageJSON contains the JSON metadata for the struct [Mirage]
type mirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Mirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mirageJSON) RawJSON() string {
	return r.raw
}

func (r Mirage) implementsZonesSettingEditResponse() {}

func (r Mirage) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type MirageID string

const (
	MirageIDMirage MirageID = "mirage"
)

func (r MirageID) IsKnown() bool {
	switch r {
	case MirageIDMirage:
		return true
	}
	return false
}

// Current value of the zone setting.
type MirageValue string

const (
	MirageValueOn  MirageValue = "on"
	MirageValueOff MirageValue = "off"
)

func (r MirageValue) IsKnown() bool {
	switch r {
	case MirageValueOn, MirageValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type MirageEditable bool

const (
	MirageEditableTrue  MirageEditable = true
	MirageEditableFalse MirageEditable = false
)

func (r MirageEditable) IsKnown() bool {
	switch r {
	case MirageEditableTrue, MirageEditableFalse:
		return true
	}
	return false
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type MirageParam struct {
	// ID of the zone setting.
	ID param.Field[MirageID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[MirageValue] `json:"value,required"`
}

func (r MirageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MirageParam) implementsZonesSettingEditParamsBodyUnion() {}

// Deprecated: Use Single Redirects instead
// https://developers.cloudflare.com/rules/url-forwarding/single-redirects/examples/#perform-mobile-redirects.
// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain.
type MobileRedirect struct {
	// Identifier of the zone setting.
	ID MobileRedirectID `json:"id,required"`
	// Current value of the zone setting.
	Value MobileRedirectValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time          `json:"modified_on,nullable" format:"date-time"`
	JSON       mobileRedirectJSON `json:"-"`
}

// mobileRedirectJSON contains the JSON metadata for the struct [MobileRedirect]
type mobileRedirectJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mobileRedirectJSON) RawJSON() string {
	return r.raw
}

func (r MobileRedirect) implementsZonesSettingEditResponse() {}

func (r MobileRedirect) implementsZonesSettingGetResponse() {}

// Identifier of the zone setting.
type MobileRedirectID string

const (
	MobileRedirectIDMobileRedirect MobileRedirectID = "mobile_redirect"
)

func (r MobileRedirectID) IsKnown() bool {
	switch r {
	case MobileRedirectIDMobileRedirect:
		return true
	}
	return false
}

// Current value of the zone setting.
type MobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Deprecated: Use Single Redirects instead
	// https://developers.cloudflare.com/rules/url-forwarding/single-redirects/examples/#perform-mobile-redirects.
	// Whether or not mobile redirect is enabled.
	Status MobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripURI bool                    `json:"strip_uri"`
	JSON     mobileRedirectValueJSON `json:"-"`
}

// mobileRedirectValueJSON contains the JSON metadata for the struct
// [MobileRedirectValue]
type mobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripURI        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mobileRedirectValueJSON) RawJSON() string {
	return r.raw
}

// Deprecated: Use Single Redirects instead
// https://developers.cloudflare.com/rules/url-forwarding/single-redirects/examples/#perform-mobile-redirects.
// Whether or not mobile redirect is enabled.
type MobileRedirectValueStatus string

const (
	MobileRedirectValueStatusOn  MobileRedirectValueStatus = "on"
	MobileRedirectValueStatusOff MobileRedirectValueStatus = "off"
)

func (r MobileRedirectValueStatus) IsKnown() bool {
	switch r {
	case MobileRedirectValueStatusOn, MobileRedirectValueStatusOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type MobileRedirectEditable bool

const (
	MobileRedirectEditableTrue  MobileRedirectEditable = true
	MobileRedirectEditableFalse MobileRedirectEditable = false
)

func (r MobileRedirectEditable) IsKnown() bool {
	switch r {
	case MobileRedirectEditableTrue, MobileRedirectEditableFalse:
		return true
	}
	return false
}

// Deprecated: Use Single Redirects instead
// https://developers.cloudflare.com/rules/url-forwarding/single-redirects/examples/#perform-mobile-redirects.
// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain.
type MobileRedirectParam struct {
	// Identifier of the zone setting.
	ID param.Field[MobileRedirectID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[MobileRedirectValueParam] `json:"value,required"`
}

func (r MobileRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MobileRedirectParam) implementsZonesSettingEditParamsBodyUnion() {}

// Current value of the zone setting.
type MobileRedirectValueParam struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Deprecated: Use Single Redirects instead
	// https://developers.cloudflare.com/rules/url-forwarding/single-redirects/examples/#perform-mobile-redirects.
	// Whether or not mobile redirect is enabled.
	Status param.Field[MobileRedirectValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripURI param.Field[bool] `json:"strip_uri"`
}

func (r MobileRedirectValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable Network Error Logging reporting on your zone. (Beta)
type NEL struct {
	// Zone setting identifier.
	ID NELID `json:"id,required"`
	// Current value of the zone setting.
	Value NELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable NELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       nelJSON   `json:"-"`
}

// nelJSON contains the JSON metadata for the struct [NEL]
type nelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nelJSON) RawJSON() string {
	return r.raw
}

func (r NEL) implementsZonesSettingEditResponse() {}

func (r NEL) implementsZonesSettingGetResponse() {}

// Zone setting identifier.
type NELID string

const (
	NELIDNEL NELID = "nel"
)

func (r NELID) IsKnown() bool {
	switch r {
	case NELIDNEL:
		return true
	}
	return false
}

// Current value of the zone setting.
type NELValue struct {
	Enabled bool         `json:"enabled"`
	JSON    nelValueJSON `json:"-"`
}

// nelValueJSON contains the JSON metadata for the struct [NELValue]
type nelValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nelValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type NELEditable bool

const (
	NELEditableTrue  NELEditable = true
	NELEditableFalse NELEditable = false
)

func (r NELEditable) IsKnown() bool {
	switch r {
	case NELEditableTrue, NELEditableFalse:
		return true
	}
	return false
}

// Enable Network Error Logging reporting on your zone. (Beta)
type NELParam struct {
	// Zone setting identifier.
	ID param.Field[NELID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[NELValueParam] `json:"value,required"`
}

func (r NELParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NELParam) implementsZonesSettingEditParamsBodyUnion() {}

// Current value of the zone setting.
type NELValueParam struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r NELValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables the Opportunistic Encryption feature for a zone.
type OpportunisticEncryption struct {
	// ID of the zone setting.
	ID OpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value OpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       opportunisticEncryptionJSON `json:"-"`
}

// opportunisticEncryptionJSON contains the JSON metadata for the struct
// [OpportunisticEncryption]
type opportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r opportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r OpportunisticEncryption) implementsZonesSettingEditResponse() {}

func (r OpportunisticEncryption) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type OpportunisticEncryptionID string

const (
	OpportunisticEncryptionIDOpportunisticEncryption OpportunisticEncryptionID = "opportunistic_encryption"
)

func (r OpportunisticEncryptionID) IsKnown() bool {
	switch r {
	case OpportunisticEncryptionIDOpportunisticEncryption:
		return true
	}
	return false
}

// Current value of the zone setting.
type OpportunisticEncryptionValue string

const (
	OpportunisticEncryptionValueOn  OpportunisticEncryptionValue = "on"
	OpportunisticEncryptionValueOff OpportunisticEncryptionValue = "off"
)

func (r OpportunisticEncryptionValue) IsKnown() bool {
	switch r {
	case OpportunisticEncryptionValueOn, OpportunisticEncryptionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type OpportunisticEncryptionEditable bool

const (
	OpportunisticEncryptionEditableTrue  OpportunisticEncryptionEditable = true
	OpportunisticEncryptionEditableFalse OpportunisticEncryptionEditable = false
)

func (r OpportunisticEncryptionEditable) IsKnown() bool {
	switch r {
	case OpportunisticEncryptionEditableTrue, OpportunisticEncryptionEditableFalse:
		return true
	}
	return false
}

// Enables the Opportunistic Encryption feature for a zone.
type OpportunisticEncryptionParam struct {
	// ID of the zone setting.
	ID param.Field[OpportunisticEncryptionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[OpportunisticEncryptionValue] `json:"value,required"`
}

func (r OpportunisticEncryptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpportunisticEncryptionParam) implementsZonesSettingEditParamsBodyUnion() {}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type OpportunisticOnion struct {
	// ID of the zone setting.
	ID OpportunisticOnionID `json:"id,required"`
	// Current value of the zone setting.
	Value OpportunisticOnionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       opportunisticOnionJSON `json:"-"`
}

// opportunisticOnionJSON contains the JSON metadata for the struct
// [OpportunisticOnion]
type opportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r opportunisticOnionJSON) RawJSON() string {
	return r.raw
}

func (r OpportunisticOnion) implementsZonesSettingEditResponse() {}

func (r OpportunisticOnion) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type OpportunisticOnionID string

const (
	OpportunisticOnionIDOpportunisticOnion OpportunisticOnionID = "opportunistic_onion"
)

func (r OpportunisticOnionID) IsKnown() bool {
	switch r {
	case OpportunisticOnionIDOpportunisticOnion:
		return true
	}
	return false
}

// Current value of the zone setting.
type OpportunisticOnionValue string

const (
	OpportunisticOnionValueOn  OpportunisticOnionValue = "on"
	OpportunisticOnionValueOff OpportunisticOnionValue = "off"
)

func (r OpportunisticOnionValue) IsKnown() bool {
	switch r {
	case OpportunisticOnionValueOn, OpportunisticOnionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type OpportunisticOnionEditable bool

const (
	OpportunisticOnionEditableTrue  OpportunisticOnionEditable = true
	OpportunisticOnionEditableFalse OpportunisticOnionEditable = false
)

func (r OpportunisticOnionEditable) IsKnown() bool {
	switch r {
	case OpportunisticOnionEditableTrue, OpportunisticOnionEditableFalse:
		return true
	}
	return false
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type OpportunisticOnionParam struct {
	// ID of the zone setting.
	ID param.Field[OpportunisticOnionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[OpportunisticOnionValue] `json:"value,required"`
}

func (r OpportunisticOnionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpportunisticOnionParam) implementsZonesSettingEditParamsBodyUnion() {}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type OrangeToOrange struct {
	// ID of the zone setting.
	ID OrangeToOrangeID `json:"id,required"`
	// Current value of the zone setting.
	Value OrangeToOrangeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time          `json:"modified_on,nullable" format:"date-time"`
	JSON       orangeToOrangeJSON `json:"-"`
}

// orangeToOrangeJSON contains the JSON metadata for the struct [OrangeToOrange]
type orangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orangeToOrangeJSON) RawJSON() string {
	return r.raw
}

func (r OrangeToOrange) implementsZonesSettingEditResponse() {}

func (r OrangeToOrange) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type OrangeToOrangeID string

const (
	OrangeToOrangeIDOrangeToOrange OrangeToOrangeID = "orange_to_orange"
)

func (r OrangeToOrangeID) IsKnown() bool {
	switch r {
	case OrangeToOrangeIDOrangeToOrange:
		return true
	}
	return false
}

// Current value of the zone setting.
type OrangeToOrangeValue string

const (
	OrangeToOrangeValueOn  OrangeToOrangeValue = "on"
	OrangeToOrangeValueOff OrangeToOrangeValue = "off"
)

func (r OrangeToOrangeValue) IsKnown() bool {
	switch r {
	case OrangeToOrangeValueOn, OrangeToOrangeValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type OrangeToOrangeEditable bool

const (
	OrangeToOrangeEditableTrue  OrangeToOrangeEditable = true
	OrangeToOrangeEditableFalse OrangeToOrangeEditable = false
)

func (r OrangeToOrangeEditable) IsKnown() bool {
	switch r {
	case OrangeToOrangeEditableTrue, OrangeToOrangeEditableFalse:
		return true
	}
	return false
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type OrangeToOrangeParam struct {
	// ID of the zone setting.
	ID param.Field[OrangeToOrangeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[OrangeToOrangeValue] `json:"value,required"`
}

func (r OrangeToOrangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrangeToOrangeParam) implementsZonesSettingEditParamsBodyUnion() {}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type OriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID OriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value OriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       originErrorPagePassThruJSON `json:"-"`
}

// originErrorPagePassThruJSON contains the JSON metadata for the struct
// [OriginErrorPagePassThru]
type originErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originErrorPagePassThruJSON) RawJSON() string {
	return r.raw
}

func (r OriginErrorPagePassThru) implementsZonesSettingEditResponse() {}

func (r OriginErrorPagePassThru) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type OriginErrorPagePassThruID string

const (
	OriginErrorPagePassThruIDOriginErrorPagePassThru OriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

func (r OriginErrorPagePassThruID) IsKnown() bool {
	switch r {
	case OriginErrorPagePassThruIDOriginErrorPagePassThru:
		return true
	}
	return false
}

// Current value of the zone setting.
type OriginErrorPagePassThruValue string

const (
	OriginErrorPagePassThruValueOn  OriginErrorPagePassThruValue = "on"
	OriginErrorPagePassThruValueOff OriginErrorPagePassThruValue = "off"
)

func (r OriginErrorPagePassThruValue) IsKnown() bool {
	switch r {
	case OriginErrorPagePassThruValueOn, OriginErrorPagePassThruValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type OriginErrorPagePassThruEditable bool

const (
	OriginErrorPagePassThruEditableTrue  OriginErrorPagePassThruEditable = true
	OriginErrorPagePassThruEditableFalse OriginErrorPagePassThruEditable = false
)

func (r OriginErrorPagePassThruEditable) IsKnown() bool {
	switch r {
	case OriginErrorPagePassThruEditableTrue, OriginErrorPagePassThruEditableFalse:
		return true
	}
	return false
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type OriginErrorPagePassThruParam struct {
	// ID of the zone setting.
	ID param.Field[OriginErrorPagePassThruID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[OriginErrorPagePassThruValue] `json:"value,required"`
}

func (r OriginErrorPagePassThruParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OriginErrorPagePassThruParam) implementsZonesSettingEditParamsBodyUnion() {}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type Polish struct {
	// ID of the zone setting.
	ID PolishID `json:"id,required"`
	// Current value of the zone setting.
	Value PolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time  `json:"modified_on,nullable" format:"date-time"`
	JSON       polishJSON `json:"-"`
}

// polishJSON contains the JSON metadata for the struct [Polish]
type polishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Polish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r polishJSON) RawJSON() string {
	return r.raw
}

func (r Polish) implementsZonesSettingEditResponse() {}

func (r Polish) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type PolishID string

const (
	PolishIDPolish PolishID = "polish"
)

func (r PolishID) IsKnown() bool {
	switch r {
	case PolishIDPolish:
		return true
	}
	return false
}

// Current value of the zone setting.
type PolishValue string

const (
	PolishValueOff      PolishValue = "off"
	PolishValueLossless PolishValue = "lossless"
	PolishValueLossy    PolishValue = "lossy"
)

func (r PolishValue) IsKnown() bool {
	switch r {
	case PolishValueOff, PolishValueLossless, PolishValueLossy:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type PolishEditable bool

const (
	PolishEditableTrue  PolishEditable = true
	PolishEditableFalse PolishEditable = false
)

func (r PolishEditable) IsKnown() bool {
	switch r {
	case PolishEditableTrue, PolishEditableFalse:
		return true
	}
	return false
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type PolishParam struct {
	// ID of the zone setting.
	ID param.Field[PolishID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[PolishValue] `json:"value,required"`
}

func (r PolishParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PolishParam) implementsZonesSettingEditParamsBodyUnion() {}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type PrefetchPreload struct {
	// ID of the zone setting.
	ID PrefetchPreloadID `json:"id,required"`
	// Current value of the zone setting.
	Value PrefetchPreloadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       prefetchPreloadJSON `json:"-"`
}

// prefetchPreloadJSON contains the JSON metadata for the struct [PrefetchPreload]
type prefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefetchPreloadJSON) RawJSON() string {
	return r.raw
}

func (r PrefetchPreload) implementsZonesSettingEditResponse() {}

func (r PrefetchPreload) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type PrefetchPreloadID string

const (
	PrefetchPreloadIDPrefetchPreload PrefetchPreloadID = "prefetch_preload"
)

func (r PrefetchPreloadID) IsKnown() bool {
	switch r {
	case PrefetchPreloadIDPrefetchPreload:
		return true
	}
	return false
}

// Current value of the zone setting.
type PrefetchPreloadValue string

const (
	PrefetchPreloadValueOn  PrefetchPreloadValue = "on"
	PrefetchPreloadValueOff PrefetchPreloadValue = "off"
)

func (r PrefetchPreloadValue) IsKnown() bool {
	switch r {
	case PrefetchPreloadValueOn, PrefetchPreloadValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type PrefetchPreloadEditable bool

const (
	PrefetchPreloadEditableTrue  PrefetchPreloadEditable = true
	PrefetchPreloadEditableFalse PrefetchPreloadEditable = false
)

func (r PrefetchPreloadEditable) IsKnown() bool {
	switch r {
	case PrefetchPreloadEditableTrue, PrefetchPreloadEditableFalse:
		return true
	}
	return false
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type PrefetchPreloadParam struct {
	// ID of the zone setting.
	ID param.Field[PrefetchPreloadID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[PrefetchPreloadValue] `json:"value,required"`
}

func (r PrefetchPreloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PrefetchPreloadParam) implementsZonesSettingEditParamsBodyUnion() {}

// Maximum time between two read operations from origin.
type ProxyReadTimeout struct {
	// ID of the zone setting.
	ID ProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time            `json:"modified_on,nullable" format:"date-time"`
	JSON       proxyReadTimeoutJSON `json:"-"`
}

// proxyReadTimeoutJSON contains the JSON metadata for the struct
// [ProxyReadTimeout]
type proxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proxyReadTimeoutJSON) RawJSON() string {
	return r.raw
}

func (r ProxyReadTimeout) implementsZonesSettingEditResponse() {}

func (r ProxyReadTimeout) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ProxyReadTimeoutID string

const (
	ProxyReadTimeoutIDProxyReadTimeout ProxyReadTimeoutID = "proxy_read_timeout"
)

func (r ProxyReadTimeoutID) IsKnown() bool {
	switch r {
	case ProxyReadTimeoutIDProxyReadTimeout:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ProxyReadTimeoutEditable bool

const (
	ProxyReadTimeoutEditableTrue  ProxyReadTimeoutEditable = true
	ProxyReadTimeoutEditableFalse ProxyReadTimeoutEditable = false
)

func (r ProxyReadTimeoutEditable) IsKnown() bool {
	switch r {
	case ProxyReadTimeoutEditableTrue, ProxyReadTimeoutEditableFalse:
		return true
	}
	return false
}

// Maximum time between two read operations from origin.
type ProxyReadTimeoutParam struct {
	// ID of the zone setting.
	ID param.Field[ProxyReadTimeoutID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[float64] `json:"value,required"`
}

func (r ProxyReadTimeoutParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProxyReadTimeoutParam) implementsZonesSettingEditParamsBodyUnion() {}

// The value set for the Pseudo IPv4 setting.
type PseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID PseudoIPV4ID `json:"id,required"`
	// Current value of the zone setting.
	Value PseudoIPV4Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time      `json:"modified_on,nullable" format:"date-time"`
	JSON       pseudoIPV4JSON `json:"-"`
}

// pseudoIPV4JSON contains the JSON metadata for the struct [PseudoIPV4]
type pseudoIPV4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PseudoIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pseudoIPV4JSON) RawJSON() string {
	return r.raw
}

func (r PseudoIPV4) implementsZonesSettingEditResponse() {}

func (r PseudoIPV4) implementsZonesSettingGetResponse() {}

// Value of the Pseudo IPv4 setting.
type PseudoIPV4ID string

const (
	PseudoIPV4IDPseudoIPV4 PseudoIPV4ID = "pseudo_ipv4"
)

func (r PseudoIPV4ID) IsKnown() bool {
	switch r {
	case PseudoIPV4IDPseudoIPV4:
		return true
	}
	return false
}

// Current value of the zone setting.
type PseudoIPV4Value string

const (
	PseudoIPV4ValueOff             PseudoIPV4Value = "off"
	PseudoIPV4ValueAddHeader       PseudoIPV4Value = "add_header"
	PseudoIPV4ValueOverwriteHeader PseudoIPV4Value = "overwrite_header"
)

func (r PseudoIPV4Value) IsKnown() bool {
	switch r {
	case PseudoIPV4ValueOff, PseudoIPV4ValueAddHeader, PseudoIPV4ValueOverwriteHeader:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type PseudoIPV4Editable bool

const (
	PseudoIPV4EditableTrue  PseudoIPV4Editable = true
	PseudoIPV4EditableFalse PseudoIPV4Editable = false
)

func (r PseudoIPV4Editable) IsKnown() bool {
	switch r {
	case PseudoIPV4EditableTrue, PseudoIPV4EditableFalse:
		return true
	}
	return false
}

// The value set for the Pseudo IPv4 setting.
type PseudoIPV4Param struct {
	// Value of the Pseudo IPv4 setting.
	ID param.Field[PseudoIPV4ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[PseudoIPV4Value] `json:"value,required"`
}

func (r PseudoIPV4Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PseudoIPV4Param) implementsZonesSettingEditParamsBodyUnion() {}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ResponseBuffering struct {
	// ID of the zone setting.
	ID ResponseBufferingID `json:"id,required"`
	// Current value of the zone setting.
	Value ResponseBufferingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on,nullable" format:"date-time"`
	JSON       responseBufferingJSON `json:"-"`
}

// responseBufferingJSON contains the JSON metadata for the struct
// [ResponseBuffering]
type responseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseBufferingJSON) RawJSON() string {
	return r.raw
}

func (r ResponseBuffering) implementsZonesSettingEditResponse() {}

func (r ResponseBuffering) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ResponseBufferingID string

const (
	ResponseBufferingIDResponseBuffering ResponseBufferingID = "response_buffering"
)

func (r ResponseBufferingID) IsKnown() bool {
	switch r {
	case ResponseBufferingIDResponseBuffering:
		return true
	}
	return false
}

// Current value of the zone setting.
type ResponseBufferingValue string

const (
	ResponseBufferingValueOn  ResponseBufferingValue = "on"
	ResponseBufferingValueOff ResponseBufferingValue = "off"
)

func (r ResponseBufferingValue) IsKnown() bool {
	switch r {
	case ResponseBufferingValueOn, ResponseBufferingValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ResponseBufferingEditable bool

const (
	ResponseBufferingEditableTrue  ResponseBufferingEditable = true
	ResponseBufferingEditableFalse ResponseBufferingEditable = false
)

func (r ResponseBufferingEditable) IsKnown() bool {
	switch r {
	case ResponseBufferingEditableTrue, ResponseBufferingEditableFalse:
		return true
	}
	return false
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ResponseBufferingParam struct {
	// ID of the zone setting.
	ID param.Field[ResponseBufferingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ResponseBufferingValue] `json:"value,required"`
}

func (r ResponseBufferingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ResponseBufferingParam) implementsZonesSettingEditParamsBodyUnion() {}

// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
// prioritises rendering your content while loading your site's Javascript
// asynchronously. Turning on Rocket Loader will immediately improve a web page's
// rendering time sometimes measured as Time to First Paint (TTFP), and also the
// `window.onload` time (assuming there is JavaScript on the page). This can have a
// positive impact on your Google search ranking. When turned on, Rocket Loader
// will automatically defer the loading of all Javascript referenced in your HTML,
// with no configuration required. Refer to
// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
// for more information.
type RocketLoader struct {
	// ID of the zone setting.
	ID RocketLoaderID `json:"id,required"`
	// Current value of the zone setting.
	Value RocketLoaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable RocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on,nullable" format:"date-time"`
	JSON       rocketLoaderJSON `json:"-"`
}

// rocketLoaderJSON contains the JSON metadata for the struct [RocketLoader]
type rocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rocketLoaderJSON) RawJSON() string {
	return r.raw
}

func (r RocketLoader) implementsZonesSettingEditResponse() {}

func (r RocketLoader) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type RocketLoaderID string

const (
	RocketLoaderIDRocketLoader RocketLoaderID = "rocket_loader"
)

func (r RocketLoaderID) IsKnown() bool {
	switch r {
	case RocketLoaderIDRocketLoader:
		return true
	}
	return false
}

// Current value of the zone setting.
type RocketLoaderValue string

const (
	RocketLoaderValueOn  RocketLoaderValue = "on"
	RocketLoaderValueOff RocketLoaderValue = "off"
)

func (r RocketLoaderValue) IsKnown() bool {
	switch r {
	case RocketLoaderValueOn, RocketLoaderValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type RocketLoaderEditable bool

const (
	RocketLoaderEditableTrue  RocketLoaderEditable = true
	RocketLoaderEditableFalse RocketLoaderEditable = false
)

func (r RocketLoaderEditable) IsKnown() bool {
	switch r {
	case RocketLoaderEditableTrue, RocketLoaderEditableFalse:
		return true
	}
	return false
}

// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
// prioritises rendering your content while loading your site's Javascript
// asynchronously. Turning on Rocket Loader will immediately improve a web page's
// rendering time sometimes measured as Time to First Paint (TTFP), and also the
// `window.onload` time (assuming there is JavaScript on the page). This can have a
// positive impact on your Google search ranking. When turned on, Rocket Loader
// will automatically defer the loading of all Javascript referenced in your HTML,
// with no configuration required. Refer to
// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
// for more information.
type RocketLoaderParam struct {
	// ID of the zone setting.
	ID param.Field[RocketLoaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[RocketLoaderValue] `json:"value,required"`
}

func (r RocketLoaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RocketLoaderParam) implementsZonesSettingEditParamsBodyUnion() {}

// Cloudflare security header for a zone.
type SecurityHeaders struct {
	// ID of the zone's security header.
	ID SecurityHeadersID `json:"id,required"`
	// Current value of the zone setting.
	Value SecurityHeadersValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SecurityHeadersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	JSON       securityHeadersJSON `json:"-"`
}

// securityHeadersJSON contains the JSON metadata for the struct [SecurityHeaders]
type securityHeadersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityHeaders) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityHeadersJSON) RawJSON() string {
	return r.raw
}

func (r SecurityHeaders) implementsZonesSettingEditResponse() {}

func (r SecurityHeaders) implementsZonesSettingGetResponse() {}

// ID of the zone's security header.
type SecurityHeadersID string

const (
	SecurityHeadersIDSecurityHeader SecurityHeadersID = "security_header"
)

func (r SecurityHeadersID) IsKnown() bool {
	switch r {
	case SecurityHeadersIDSecurityHeader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SecurityHeadersValue struct {
	// Strict Transport Security.
	StrictTransportSecurity SecurityHeadersValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    securityHeadersValueJSON                    `json:"-"`
}

// securityHeadersValueJSON contains the JSON metadata for the struct
// [SecurityHeadersValue]
type securityHeadersValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SecurityHeadersValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityHeadersValueJSON) RawJSON() string {
	return r.raw
}

// Strict Transport Security.
type SecurityHeadersValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                            `json:"nosniff"`
	JSON    securityHeadersValueStrictTransportSecurityJSON `json:"-"`
}

// securityHeadersValueStrictTransportSecurityJSON contains the JSON metadata for
// the struct [SecurityHeadersValueStrictTransportSecurity]
type securityHeadersValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SecurityHeadersValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityHeadersValueStrictTransportSecurityJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SecurityHeadersEditable bool

const (
	SecurityHeadersEditableTrue  SecurityHeadersEditable = true
	SecurityHeadersEditableFalse SecurityHeadersEditable = false
)

func (r SecurityHeadersEditable) IsKnown() bool {
	switch r {
	case SecurityHeadersEditableTrue, SecurityHeadersEditableFalse:
		return true
	}
	return false
}

// Cloudflare security header for a zone.
type SecurityHeadersParam struct {
	// ID of the zone's security header.
	ID param.Field[SecurityHeadersID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SecurityHeadersValueParam] `json:"value,required"`
}

func (r SecurityHeadersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SecurityHeadersParam) implementsZonesSettingEditParamsBodyUnion() {}

// Current value of the zone setting.
type SecurityHeadersValueParam struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[SecurityHeadersValueStrictTransportSecurityParam] `json:"strict_transport_security"`
}

func (r SecurityHeadersValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strict Transport Security.
type SecurityHeadersValueStrictTransportSecurityParam struct {
	// Whether or not strict transport security is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge param.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff param.Field[bool] `json:"nosniff"`
}

func (r SecurityHeadersValueStrictTransportSecurityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SecurityLevel struct {
	// ID of the zone setting.
	ID SecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on,nullable" format:"date-time"`
	JSON       securityLevelJSON `json:"-"`
}

// securityLevelJSON contains the JSON metadata for the struct [SecurityLevel]
type securityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityLevelJSON) RawJSON() string {
	return r.raw
}

func (r SecurityLevel) implementsZonesSettingEditResponse() {}

func (r SecurityLevel) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SecurityLevelID string

const (
	SecurityLevelIDSecurityLevel SecurityLevelID = "security_level"
)

func (r SecurityLevelID) IsKnown() bool {
	switch r {
	case SecurityLevelIDSecurityLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SecurityLevelValue string

const (
	SecurityLevelValueOff            SecurityLevelValue = "off"
	SecurityLevelValueEssentiallyOff SecurityLevelValue = "essentially_off"
	SecurityLevelValueLow            SecurityLevelValue = "low"
	SecurityLevelValueMedium         SecurityLevelValue = "medium"
	SecurityLevelValueHigh           SecurityLevelValue = "high"
	SecurityLevelValueUnderAttack    SecurityLevelValue = "under_attack"
)

func (r SecurityLevelValue) IsKnown() bool {
	switch r {
	case SecurityLevelValueOff, SecurityLevelValueEssentiallyOff, SecurityLevelValueLow, SecurityLevelValueMedium, SecurityLevelValueHigh, SecurityLevelValueUnderAttack:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SecurityLevelEditable bool

const (
	SecurityLevelEditableTrue  SecurityLevelEditable = true
	SecurityLevelEditableFalse SecurityLevelEditable = false
)

func (r SecurityLevelEditable) IsKnown() bool {
	switch r {
	case SecurityLevelEditableTrue, SecurityLevelEditableFalse:
		return true
	}
	return false
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SecurityLevelParam struct {
	// ID of the zone setting.
	ID param.Field[SecurityLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SecurityLevelValue] `json:"value,required"`
}

func (r SecurityLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SecurityLevelParam) implementsZonesSettingEditParamsBodyUnion() {}

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
type ServerSideExcludes struct {
	// ID of the zone setting.
	ID ServerSideExcludesID `json:"id,required"`
	// Current value of the zone setting.
	Value ServerSideExcludesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ServerSideExcludesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       serverSideExcludesJSON `json:"-"`
}

// serverSideExcludesJSON contains the JSON metadata for the struct
// [ServerSideExcludes]
type serverSideExcludesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerSideExcludes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serverSideExcludesJSON) RawJSON() string {
	return r.raw
}

func (r ServerSideExcludes) implementsZonesSettingEditResponse() {}

func (r ServerSideExcludes) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ServerSideExcludesID string

const (
	ServerSideExcludesIDServerSideExclude ServerSideExcludesID = "server_side_exclude"
)

func (r ServerSideExcludesID) IsKnown() bool {
	switch r {
	case ServerSideExcludesIDServerSideExclude:
		return true
	}
	return false
}

// Current value of the zone setting.
type ServerSideExcludesValue string

const (
	ServerSideExcludesValueOn  ServerSideExcludesValue = "on"
	ServerSideExcludesValueOff ServerSideExcludesValue = "off"
)

func (r ServerSideExcludesValue) IsKnown() bool {
	switch r {
	case ServerSideExcludesValueOn, ServerSideExcludesValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ServerSideExcludesEditable bool

const (
	ServerSideExcludesEditableTrue  ServerSideExcludesEditable = true
	ServerSideExcludesEditableFalse ServerSideExcludesEditable = false
)

func (r ServerSideExcludesEditable) IsKnown() bool {
	switch r {
	case ServerSideExcludesEditableTrue, ServerSideExcludesEditableFalse:
		return true
	}
	return false
}

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
type ServerSideExcludesParam struct {
	// ID of the zone setting.
	ID param.Field[ServerSideExcludesID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ServerSideExcludesValue] `json:"value,required"`
}

func (r ServerSideExcludesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServerSideExcludesParam) implementsZonesSettingEditParamsBodyUnion() {}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type SortQueryStringForCache struct {
	// ID of the zone setting.
	ID SortQueryStringForCacheID `json:"id,required"`
	// Current value of the zone setting.
	Value SortQueryStringForCacheValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                   `json:"modified_on,nullable" format:"date-time"`
	JSON       sortQueryStringForCacheJSON `json:"-"`
}

// sortQueryStringForCacheJSON contains the JSON metadata for the struct
// [SortQueryStringForCache]
type sortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
}

func (r SortQueryStringForCache) implementsZonesSettingEditResponse() {}

func (r SortQueryStringForCache) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SortQueryStringForCacheID string

const (
	SortQueryStringForCacheIDSortQueryStringForCache SortQueryStringForCacheID = "sort_query_string_for_cache"
)

func (r SortQueryStringForCacheID) IsKnown() bool {
	switch r {
	case SortQueryStringForCacheIDSortQueryStringForCache:
		return true
	}
	return false
}

// Current value of the zone setting.
type SortQueryStringForCacheValue string

const (
	SortQueryStringForCacheValueOn  SortQueryStringForCacheValue = "on"
	SortQueryStringForCacheValueOff SortQueryStringForCacheValue = "off"
)

func (r SortQueryStringForCacheValue) IsKnown() bool {
	switch r {
	case SortQueryStringForCacheValueOn, SortQueryStringForCacheValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SortQueryStringForCacheEditable bool

const (
	SortQueryStringForCacheEditableTrue  SortQueryStringForCacheEditable = true
	SortQueryStringForCacheEditableFalse SortQueryStringForCacheEditable = false
)

func (r SortQueryStringForCacheEditable) IsKnown() bool {
	switch r {
	case SortQueryStringForCacheEditableTrue, SortQueryStringForCacheEditableFalse:
		return true
	}
	return false
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type SortQueryStringForCacheParam struct {
	// ID of the zone setting.
	ID param.Field[SortQueryStringForCacheID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SortQueryStringForCacheValue] `json:"value,required"`
}

func (r SortQueryStringForCacheParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SortQueryStringForCacheParam) implementsZonesSettingEditParamsBodyUnion() {}

// SSL encrypts your visitor's connection and safeguards credit card numbers and
// other personal data to and from your website. SSL can take up to 5 minutes to
// fully activate. Requires Cloudflare active on your root domain or www domain.
// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
// your web server. You don't need to have an SSL cert on your web server, but your
// vistors will still see the site as being HTTPS enabled. Full: SSL between the
// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
// Cloudflare and your web server. You'll need to have your own SSL cert or
// self-signed cert at the very least. Full (Strict): SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
// your web server. You'll need to have a valid SSL certificate installed on your
// web server. This certificate must be signed by a certificate authority, have an
// expiration date in the future, and respond for the request domain name
// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
type SSL struct {
	// ID of the zone setting.
	ID SSLID `json:"id,required"`
	// Current value of the zone setting.
	Value SSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       sslJSON   `json:"-"`
}

// sslJSON contains the JSON metadata for the struct [SSL]
type sslJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslJSON) RawJSON() string {
	return r.raw
}

func (r SSL) implementsZonesSettingEditResponse() {}

func (r SSL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SSLID string

const (
	SSLIDSSL SSLID = "ssl"
)

func (r SSLID) IsKnown() bool {
	switch r {
	case SSLIDSSL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SSLValue string

const (
	SSLValueOff      SSLValue = "off"
	SSLValueFlexible SSLValue = "flexible"
	SSLValueFull     SSLValue = "full"
	SSLValueStrict   SSLValue = "strict"
)

func (r SSLValue) IsKnown() bool {
	switch r {
	case SSLValueOff, SSLValueFlexible, SSLValueFull, SSLValueStrict:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SSLEditable bool

const (
	SSLEditableTrue  SSLEditable = true
	SSLEditableFalse SSLEditable = false
)

func (r SSLEditable) IsKnown() bool {
	switch r {
	case SSLEditableTrue, SSLEditableFalse:
		return true
	}
	return false
}

// SSL encrypts your visitor's connection and safeguards credit card numbers and
// other personal data to and from your website. SSL can take up to 5 minutes to
// fully activate. Requires Cloudflare active on your root domain or www domain.
// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
// your web server. You don't need to have an SSL cert on your web server, but your
// vistors will still see the site as being HTTPS enabled. Full: SSL between the
// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
// Cloudflare and your web server. You'll need to have your own SSL cert or
// self-signed cert at the very least. Full (Strict): SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
// your web server. You'll need to have a valid SSL certificate installed on your
// web server. This certificate must be signed by a certificate authority, have an
// expiration date in the future, and respond for the request domain name
// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
type SSLParam struct {
	// ID of the zone setting.
	ID param.Field[SSLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SSLValue] `json:"value,required"`
}

func (r SSLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SSLParam) implementsZonesSettingEditParamsBodyUnion() {}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID SSLRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool               `json:"enabled"`
	JSON    sslRecommenderJSON `json:"-"`
}

// sslRecommenderJSON contains the JSON metadata for the struct [SSLRecommender]
type sslRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslRecommenderJSON) RawJSON() string {
	return r.raw
}

func (r SSLRecommender) implementsZonesSettingEditResponse() {}

func (r SSLRecommender) implementsZonesSettingGetResponse() {}

// Enrollment value for SSL/TLS Recommender.
type SSLRecommenderID string

const (
	SSLRecommenderIDSSLRecommender SSLRecommenderID = "ssl_recommender"
)

func (r SSLRecommenderID) IsKnown() bool {
	switch r {
	case SSLRecommenderIDSSLRecommender:
		return true
	}
	return false
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SSLRecommenderParam struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[SSLRecommenderID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SSLRecommenderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SSLRecommenderParam) implementsZonesSettingEditParamsBodyUnion() {}

// Enables Crypto TLS 1.3 feature for a zone.
type TLS1_3 struct {
	// ID of the zone setting.
	ID TLS1_3ID `json:"id,required"`
	// Current value of the zone setting.
	Value TLS1_3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time  `json:"modified_on,nullable" format:"date-time"`
	JSON       tls1_3JSON `json:"-"`
}

// tls1_3JSON contains the JSON metadata for the struct [TLS1_3]
type tls1_3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tls1_3JSON) RawJSON() string {
	return r.raw
}

func (r TLS1_3) implementsZonesSettingEditResponse() {}

func (r TLS1_3) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type TLS1_3ID string

const (
	TLS1_3IDTLS1_3 TLS1_3ID = "tls_1_3"
)

func (r TLS1_3ID) IsKnown() bool {
	switch r {
	case TLS1_3IDTLS1_3:
		return true
	}
	return false
}

// Current value of the zone setting.
type TLS1_3Value string

const (
	TLS1_3ValueOn  TLS1_3Value = "on"
	TLS1_3ValueOff TLS1_3Value = "off"
	TLS1_3ValueZrt TLS1_3Value = "zrt"
)

func (r TLS1_3Value) IsKnown() bool {
	switch r {
	case TLS1_3ValueOn, TLS1_3ValueOff, TLS1_3ValueZrt:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type TLS1_3Editable bool

const (
	TLS1_3EditableTrue  TLS1_3Editable = true
	TLS1_3EditableFalse TLS1_3Editable = false
)

func (r TLS1_3Editable) IsKnown() bool {
	switch r {
	case TLS1_3EditableTrue, TLS1_3EditableFalse:
		return true
	}
	return false
}

// Enables Crypto TLS 1.3 feature for a zone.
type TLS1_3Param struct {
	// ID of the zone setting.
	ID param.Field[TLS1_3ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[TLS1_3Value] `json:"value,required"`
}

func (r TLS1_3Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TLS1_3Param) implementsZonesSettingEditParamsBodyUnion() {}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type TLSClientAuth struct {
	// ID of the zone setting.
	ID TLSClientAuthID `json:"id,required"`
	// Current value of the zone setting.
	Value TLSClientAuthValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on,nullable" format:"date-time"`
	JSON       tlsClientAuthJSON `json:"-"`
}

// tlsClientAuthJSON contains the JSON metadata for the struct [TLSClientAuth]
type tlsClientAuthJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsClientAuthJSON) RawJSON() string {
	return r.raw
}

func (r TLSClientAuth) implementsZonesSettingEditResponse() {}

func (r TLSClientAuth) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type TLSClientAuthID string

const (
	TLSClientAuthIDTLSClientAuth TLSClientAuthID = "tls_client_auth"
)

func (r TLSClientAuthID) IsKnown() bool {
	switch r {
	case TLSClientAuthIDTLSClientAuth:
		return true
	}
	return false
}

// Current value of the zone setting.
type TLSClientAuthValue string

const (
	TLSClientAuthValueOn  TLSClientAuthValue = "on"
	TLSClientAuthValueOff TLSClientAuthValue = "off"
)

func (r TLSClientAuthValue) IsKnown() bool {
	switch r {
	case TLSClientAuthValueOn, TLSClientAuthValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type TLSClientAuthEditable bool

const (
	TLSClientAuthEditableTrue  TLSClientAuthEditable = true
	TLSClientAuthEditableFalse TLSClientAuthEditable = false
)

func (r TLSClientAuthEditable) IsKnown() bool {
	switch r {
	case TLSClientAuthEditableTrue, TLSClientAuthEditableFalse:
		return true
	}
	return false
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type TLSClientAuthParam struct {
	// ID of the zone setting.
	ID param.Field[TLSClientAuthID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[TLSClientAuthValue] `json:"value,required"`
}

func (r TLSClientAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TLSClientAuthParam) implementsZonesSettingEditParamsBodyUnion() {}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type TrueClientIPHeader struct {
	// ID of the zone setting.
	ID TrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value TrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on,nullable" format:"date-time"`
	JSON       trueClientIPHeaderJSON `json:"-"`
}

// trueClientIPHeaderJSON contains the JSON metadata for the struct
// [TrueClientIPHeader]
type trueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

func (r TrueClientIPHeader) implementsZonesSettingEditResponse() {}

func (r TrueClientIPHeader) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type TrueClientIPHeaderID string

const (
	TrueClientIPHeaderIDTrueClientIPHeader TrueClientIPHeaderID = "true_client_ip_header"
)

func (r TrueClientIPHeaderID) IsKnown() bool {
	switch r {
	case TrueClientIPHeaderIDTrueClientIPHeader:
		return true
	}
	return false
}

// Current value of the zone setting.
type TrueClientIPHeaderValue string

const (
	TrueClientIPHeaderValueOn  TrueClientIPHeaderValue = "on"
	TrueClientIPHeaderValueOff TrueClientIPHeaderValue = "off"
)

func (r TrueClientIPHeaderValue) IsKnown() bool {
	switch r {
	case TrueClientIPHeaderValueOn, TrueClientIPHeaderValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type TrueClientIPHeaderEditable bool

const (
	TrueClientIPHeaderEditableTrue  TrueClientIPHeaderEditable = true
	TrueClientIPHeaderEditableFalse TrueClientIPHeaderEditable = false
)

func (r TrueClientIPHeaderEditable) IsKnown() bool {
	switch r {
	case TrueClientIPHeaderEditableTrue, TrueClientIPHeaderEditableFalse:
		return true
	}
	return false
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type TrueClientIPHeaderParam struct {
	// ID of the zone setting.
	ID param.Field[TrueClientIPHeaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[TrueClientIPHeaderValue] `json:"value,required"`
}

func (r TrueClientIPHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TrueClientIPHeaderParam) implementsZonesSettingEditParamsBodyUnion() {}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
type WAF struct {
	// ID of the zone setting.
	ID WAFID `json:"id,required"`
	// Current value of the zone setting.
	Value WAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       wafJSON   `json:"-"`
}

// wafJSON contains the JSON metadata for the struct [WAF]
type wafJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafJSON) RawJSON() string {
	return r.raw
}

func (r WAF) implementsZonesSettingEditResponse() {}

func (r WAF) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type WAFID string

const (
	WAFIDWAF WAFID = "waf"
)

func (r WAFID) IsKnown() bool {
	switch r {
	case WAFIDWAF:
		return true
	}
	return false
}

// Current value of the zone setting.
type WAFValue string

const (
	WAFValueOn  WAFValue = "on"
	WAFValueOff WAFValue = "off"
)

func (r WAFValue) IsKnown() bool {
	switch r {
	case WAFValueOn, WAFValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type WAFEditable bool

const (
	WAFEditableTrue  WAFEditable = true
	WAFEditableFalse WAFEditable = false
)

func (r WAFEditable) IsKnown() bool {
	switch r {
	case WAFEditableTrue, WAFEditableFalse:
		return true
	}
	return false
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
type WAFParam struct {
	// ID of the zone setting.
	ID param.Field[WAFID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[WAFValue] `json:"value,required"`
}

func (r WAFParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WAFParam) implementsZonesSettingEditParamsBodyUnion() {}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type WebP struct {
	// ID of the zone setting.
	ID WebPID `json:"id,required"`
	// Current value of the zone setting.
	Value WebPValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WebPEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       WebPJSON  `json:"-"`
}

// WebPJSON contains the JSON metadata for the struct [WebP]
type WebPJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r WebPJSON) RawJSON() string {
	return r.raw
}

func (r WebP) implementsZonesSettingEditResponse() {}

func (r WebP) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type WebPID string

const (
	WebPIDWebP WebPID = "webp"
)

func (r WebPID) IsKnown() bool {
	switch r {
	case WebPIDWebP:
		return true
	}
	return false
}

// Current value of the zone setting.
type WebPValue string

const (
	WebPValueOff WebPValue = "off"
	WebPValueOn  WebPValue = "on"
)

func (r WebPValue) IsKnown() bool {
	switch r {
	case WebPValueOff, WebPValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type WebPEditable bool

const (
	WebPEditableTrue  WebPEditable = true
	WebPEditableFalse WebPEditable = false
)

func (r WebPEditable) IsKnown() bool {
	switch r {
	case WebPEditableTrue, WebPEditableFalse:
		return true
	}
	return false
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type WebPParam struct {
	// ID of the zone setting.
	ID param.Field[WebPID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[WebPValue] `json:"value,required"`
}

func (r WebPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebPParam) implementsZonesSettingEditParamsBodyUnion() {}

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type Websocket struct {
	// ID of the zone setting.
	ID WebsocketID `json:"id,required"`
	// Current value of the zone setting.
	Value WebsocketValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WebsocketEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time     `json:"modified_on,nullable" format:"date-time"`
	JSON       websocketJSON `json:"-"`
}

// websocketJSON contains the JSON metadata for the struct [Websocket]
type websocketJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Websocket) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r websocketJSON) RawJSON() string {
	return r.raw
}

func (r Websocket) implementsZonesSettingEditResponse() {}

func (r Websocket) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type WebsocketID string

const (
	WebsocketIDWebsockets WebsocketID = "websockets"
)

func (r WebsocketID) IsKnown() bool {
	switch r {
	case WebsocketIDWebsockets:
		return true
	}
	return false
}

// Current value of the zone setting.
type WebsocketValue string

const (
	WebsocketValueOff WebsocketValue = "off"
	WebsocketValueOn  WebsocketValue = "on"
)

func (r WebsocketValue) IsKnown() bool {
	switch r {
	case WebsocketValueOff, WebsocketValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type WebsocketEditable bool

const (
	WebsocketEditableTrue  WebsocketEditable = true
	WebsocketEditableFalse WebsocketEditable = false
)

func (r WebsocketEditable) IsKnown() bool {
	switch r {
	case WebsocketEditableTrue, WebsocketEditableFalse:
		return true
	}
	return false
}

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type WebsocketParam struct {
	// ID of the zone setting.
	ID param.Field[WebsocketID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[WebsocketValue] `json:"value,required"`
}

func (r WebsocketParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WebsocketParam) implementsZonesSettingEditParamsBodyUnion() {}

// 0-RTT session resumption enabled for this zone.
type ZeroRTT struct {
	// ID of the zone setting.
	ID ZeroRTTID `json:"id,required"`
	// Current value of the zone setting.
	Value ZeroRTTValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZeroRTTEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time   `json:"modified_on,nullable" format:"date-time"`
	JSON       zeroRTTJSON `json:"-"`
}

// zeroRTTJSON contains the JSON metadata for the struct [ZeroRTT]
type zeroRTTJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroRTT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroRTTJSON) RawJSON() string {
	return r.raw
}

func (r ZeroRTT) implementsZonesSettingEditResponse() {}

func (r ZeroRTT) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZeroRTTID string

const (
	ZeroRTTID0rtt ZeroRTTID = "0rtt"
)

func (r ZeroRTTID) IsKnown() bool {
	switch r {
	case ZeroRTTID0rtt:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZeroRTTValue string

const (
	ZeroRTTValueOn  ZeroRTTValue = "on"
	ZeroRTTValueOff ZeroRTTValue = "off"
)

func (r ZeroRTTValue) IsKnown() bool {
	switch r {
	case ZeroRTTValueOn, ZeroRTTValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZeroRTTEditable bool

const (
	ZeroRTTEditableTrue  ZeroRTTEditable = true
	ZeroRTTEditableFalse ZeroRTTEditable = false
)

func (r ZeroRTTEditable) IsKnown() bool {
	switch r {
	case ZeroRTTEditableTrue, ZeroRTTEditableFalse:
		return true
	}
	return false
}

// 0-RTT session resumption enabled for this zone.
type ZeroRTTParam struct {
	// ID of the zone setting.
	ID param.Field[ZeroRTTID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZeroRTTValue] `json:"value,required"`
}

func (r ZeroRTTParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroRTTParam) implementsZonesSettingEditParamsBodyUnion() {}

// 0-RTT session resumption enabled for this zone.
type SettingEditResponse struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseEditable `json:"editable"`
	// ID of the zone setting.
	ID SettingEditResponseID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time   `json:"modified_on,nullable" format:"date-time"`
	Value      interface{} `json:"value,required"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// ssl-recommender enrollment setting.
	Enabled bool                    `json:"enabled"`
	JSON    settingEditResponseJSON `json:"-"`
	union   SettingEditResponseUnion
}

// settingEditResponseJSON contains the JSON metadata for the struct
// [SettingEditResponse]
type settingEditResponseJSON struct {
	Editable      apijson.Field
	ID            apijson.Field
	ModifiedOn    apijson.Field
	Value         apijson.Field
	TimeRemaining apijson.Field
	Enabled       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r settingEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r SettingEditResponse) AsUnion() SettingEditResponseUnion {
	return r.union
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [zones.ZeroRTT], [zones.AdvancedDDoS], [zones.AlwaysOnline],
// [zones.AlwaysUseHTTPS], [zones.AutomaticHTTPSRewrites], [zones.Brotli],
// [zones.BrowserCacheTTL], [zones.BrowserCheck], [zones.CacheLevel],
// [zones.ChallengeTTL], [zones.Ciphers],
// [zones.SettingEditResponseZonesCNAMEFlattening], [zones.DevelopmentMode],
// [zones.EarlyHints], [zones.SettingEditResponseZonesEdgeCacheTTL],
// [zones.EmailObfuscation], [zones.H2Prioritization], [zones.HotlinkProtection],
// [zones.HTTP2], [zones.HTTP3], [zones.ImageResizing], [zones.IPGeolocation],
// [zones.IPV6], [zones.SettingEditResponseZonesMaxUpload], [zones.MinTLSVersion],
// [zones.Minify], [zones.Mirage], [zones.MobileRedirect], [zones.NEL],
// [zones.OpportunisticEncryption], [zones.OpportunisticOnion],
// [zones.OrangeToOrange], [zones.OriginErrorPagePassThru], [zones.Polish],
// [zones.PrefetchPreload], [zones.ProxyReadTimeout], [zones.PseudoIPV4],
// [zones.SettingEditResponseZonesReplaceInsecureJS], [zones.ResponseBuffering],
// [zones.RocketLoader],
// [zones.SettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.SecurityHeaders], [zones.SecurityLevel], [zones.ServerSideExcludes],
// [zones.SettingEditResponseZonesSha1Support], [zones.SortQueryStringForCache],
// [zones.SSL], [zones.SSLRecommender], [zones.SettingEditResponseZonesTLS1_2Only],
// [zones.TLS1_3], [zones.TLSClientAuth], [zones.TrueClientIPHeader], [zones.WAF],
// [zones.WebP] or [zones.Websocket].
type SettingEditResponseUnion interface {
	implementsZonesSettingEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroRTT{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AdvancedDDoS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Brotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Ciphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesCNAMEFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(H2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HTTP2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HTTP3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPV6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MinTLSVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Minify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Mirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MobileRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Polish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PseudoIPV4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesReplaceInsecureJS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResponseBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecurityHeaders{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServerSideExcludes{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SSL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SSLRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesTLS1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TLS1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TLSClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAF{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WebP{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Websocket{}),
		},
	)
}

// Whether or not cname flattening is on.
type SettingEditResponseZonesCNAMEFlattening struct {
	// How to flatten the cname destination.
	ID SettingEditResponseZonesCNAMEFlatteningID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesCNAMEFlatteningValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesCNAMEFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesCNAMEFlatteningJSON `json:"-"`
}

// settingEditResponseZonesCNAMEFlatteningJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesCNAMEFlattening]
type settingEditResponseZonesCNAMEFlatteningJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesCNAMEFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesCNAMEFlatteningJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesCNAMEFlattening) implementsZonesSettingEditResponse() {}

// How to flatten the cname destination.
type SettingEditResponseZonesCNAMEFlatteningID string

const (
	SettingEditResponseZonesCNAMEFlatteningIDCNAMEFlattening SettingEditResponseZonesCNAMEFlatteningID = "cname_flattening"
)

func (r SettingEditResponseZonesCNAMEFlatteningID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesCNAMEFlatteningIDCNAMEFlattening:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesCNAMEFlatteningValue string

const (
	SettingEditResponseZonesCNAMEFlatteningValueFlattenAtRoot SettingEditResponseZonesCNAMEFlatteningValue = "flatten_at_root"
	SettingEditResponseZonesCNAMEFlatteningValueFlattenAll    SettingEditResponseZonesCNAMEFlatteningValue = "flatten_all"
)

func (r SettingEditResponseZonesCNAMEFlatteningValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesCNAMEFlatteningValueFlattenAtRoot, SettingEditResponseZonesCNAMEFlatteningValueFlattenAll:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesCNAMEFlatteningEditable bool

const (
	SettingEditResponseZonesCNAMEFlatteningEditableTrue  SettingEditResponseZonesCNAMEFlatteningEditable = true
	SettingEditResponseZonesCNAMEFlatteningEditableFalse SettingEditResponseZonesCNAMEFlatteningEditable = false
)

func (r SettingEditResponseZonesCNAMEFlatteningEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesCNAMEFlatteningEditableTrue, SettingEditResponseZonesCNAMEFlatteningEditableFalse:
		return true
	}
	return false
}

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type SettingEditResponseZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesEdgeCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesEdgeCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesEdgeCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesEdgeCacheTTLJSON `json:"-"`
}

// settingEditResponseZonesEdgeCacheTTLJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesEdgeCacheTTL]
type settingEditResponseZonesEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesEdgeCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesEdgeCacheTTL) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesEdgeCacheTTLID string

const (
	SettingEditResponseZonesEdgeCacheTTLIDEdgeCacheTTL SettingEditResponseZonesEdgeCacheTTLID = "edge_cache_ttl"
)

func (r SettingEditResponseZonesEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesEdgeCacheTTLValue float64

const (
	SettingEditResponseZonesEdgeCacheTTLValue30     SettingEditResponseZonesEdgeCacheTTLValue = 30
	SettingEditResponseZonesEdgeCacheTTLValue60     SettingEditResponseZonesEdgeCacheTTLValue = 60
	SettingEditResponseZonesEdgeCacheTTLValue300    SettingEditResponseZonesEdgeCacheTTLValue = 300
	SettingEditResponseZonesEdgeCacheTTLValue1200   SettingEditResponseZonesEdgeCacheTTLValue = 1200
	SettingEditResponseZonesEdgeCacheTTLValue1800   SettingEditResponseZonesEdgeCacheTTLValue = 1800
	SettingEditResponseZonesEdgeCacheTTLValue3600   SettingEditResponseZonesEdgeCacheTTLValue = 3600
	SettingEditResponseZonesEdgeCacheTTLValue7200   SettingEditResponseZonesEdgeCacheTTLValue = 7200
	SettingEditResponseZonesEdgeCacheTTLValue10800  SettingEditResponseZonesEdgeCacheTTLValue = 10800
	SettingEditResponseZonesEdgeCacheTTLValue14400  SettingEditResponseZonesEdgeCacheTTLValue = 14400
	SettingEditResponseZonesEdgeCacheTTLValue18000  SettingEditResponseZonesEdgeCacheTTLValue = 18000
	SettingEditResponseZonesEdgeCacheTTLValue28800  SettingEditResponseZonesEdgeCacheTTLValue = 28800
	SettingEditResponseZonesEdgeCacheTTLValue43200  SettingEditResponseZonesEdgeCacheTTLValue = 43200
	SettingEditResponseZonesEdgeCacheTTLValue57600  SettingEditResponseZonesEdgeCacheTTLValue = 57600
	SettingEditResponseZonesEdgeCacheTTLValue72000  SettingEditResponseZonesEdgeCacheTTLValue = 72000
	SettingEditResponseZonesEdgeCacheTTLValue86400  SettingEditResponseZonesEdgeCacheTTLValue = 86400
	SettingEditResponseZonesEdgeCacheTTLValue172800 SettingEditResponseZonesEdgeCacheTTLValue = 172800
	SettingEditResponseZonesEdgeCacheTTLValue259200 SettingEditResponseZonesEdgeCacheTTLValue = 259200
	SettingEditResponseZonesEdgeCacheTTLValue345600 SettingEditResponseZonesEdgeCacheTTLValue = 345600
	SettingEditResponseZonesEdgeCacheTTLValue432000 SettingEditResponseZonesEdgeCacheTTLValue = 432000
	SettingEditResponseZonesEdgeCacheTTLValue518400 SettingEditResponseZonesEdgeCacheTTLValue = 518400
	SettingEditResponseZonesEdgeCacheTTLValue604800 SettingEditResponseZonesEdgeCacheTTLValue = 604800
)

func (r SettingEditResponseZonesEdgeCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEdgeCacheTTLValue30, SettingEditResponseZonesEdgeCacheTTLValue60, SettingEditResponseZonesEdgeCacheTTLValue300, SettingEditResponseZonesEdgeCacheTTLValue1200, SettingEditResponseZonesEdgeCacheTTLValue1800, SettingEditResponseZonesEdgeCacheTTLValue3600, SettingEditResponseZonesEdgeCacheTTLValue7200, SettingEditResponseZonesEdgeCacheTTLValue10800, SettingEditResponseZonesEdgeCacheTTLValue14400, SettingEditResponseZonesEdgeCacheTTLValue18000, SettingEditResponseZonesEdgeCacheTTLValue28800, SettingEditResponseZonesEdgeCacheTTLValue43200, SettingEditResponseZonesEdgeCacheTTLValue57600, SettingEditResponseZonesEdgeCacheTTLValue72000, SettingEditResponseZonesEdgeCacheTTLValue86400, SettingEditResponseZonesEdgeCacheTTLValue172800, SettingEditResponseZonesEdgeCacheTTLValue259200, SettingEditResponseZonesEdgeCacheTTLValue345600, SettingEditResponseZonesEdgeCacheTTLValue432000, SettingEditResponseZonesEdgeCacheTTLValue518400, SettingEditResponseZonesEdgeCacheTTLValue604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesEdgeCacheTTLEditable bool

const (
	SettingEditResponseZonesEdgeCacheTTLEditableTrue  SettingEditResponseZonesEdgeCacheTTLEditable = true
	SettingEditResponseZonesEdgeCacheTTLEditableFalse SettingEditResponseZonesEdgeCacheTTLEditable = false
)

func (r SettingEditResponseZonesEdgeCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEdgeCacheTTLEditableTrue, SettingEditResponseZonesEdgeCacheTTLEditableFalse:
		return true
	}
	return false
}

// Maximum size of an allowable upload.
type SettingEditResponseZonesMaxUpload struct {
	// identifier of the zone setting.
	ID SettingEditResponseZonesMaxUploadID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesMaxUploadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesMaxUploadJSON `json:"-"`
}

// settingEditResponseZonesMaxUploadJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesMaxUpload]
type settingEditResponseZonesMaxUploadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesMaxUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesMaxUploadJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesMaxUpload) implementsZonesSettingEditResponse() {}

// identifier of the zone setting.
type SettingEditResponseZonesMaxUploadID string

const (
	SettingEditResponseZonesMaxUploadIDMaxUpload SettingEditResponseZonesMaxUploadID = "max_upload"
)

func (r SettingEditResponseZonesMaxUploadID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesMaxUploadIDMaxUpload:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesMaxUploadValue float64

const (
	SettingEditResponseZonesMaxUploadValue100 SettingEditResponseZonesMaxUploadValue = 100
	SettingEditResponseZonesMaxUploadValue200 SettingEditResponseZonesMaxUploadValue = 200
	SettingEditResponseZonesMaxUploadValue500 SettingEditResponseZonesMaxUploadValue = 500
)

func (r SettingEditResponseZonesMaxUploadValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesMaxUploadValue100, SettingEditResponseZonesMaxUploadValue200, SettingEditResponseZonesMaxUploadValue500:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesMaxUploadEditable bool

const (
	SettingEditResponseZonesMaxUploadEditableTrue  SettingEditResponseZonesMaxUploadEditable = true
	SettingEditResponseZonesMaxUploadEditableFalse SettingEditResponseZonesMaxUploadEditable = false
)

func (r SettingEditResponseZonesMaxUploadEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesMaxUploadEditableTrue, SettingEditResponseZonesMaxUploadEditableFalse:
		return true
	}
	return false
}

// Automatically replace insecure JavaScript libraries with safer and faster
// alternatives provided under cdnjs and powered by Cloudflare. Currently supports
// the following libraries: Polyfill under polyfill.io.
type SettingEditResponseZonesReplaceInsecureJS struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesReplaceInsecureJSID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesReplaceInsecureJSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesReplaceInsecureJSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesReplaceInsecureJSJSON `json:"-"`
}

// settingEditResponseZonesReplaceInsecureJSJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesReplaceInsecureJS]
type settingEditResponseZonesReplaceInsecureJSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesReplaceInsecureJS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesReplaceInsecureJSJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesReplaceInsecureJS) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesReplaceInsecureJSID string

const (
	SettingEditResponseZonesReplaceInsecureJSIDReplaceInsecureJS SettingEditResponseZonesReplaceInsecureJSID = "replace_insecure_js"
)

func (r SettingEditResponseZonesReplaceInsecureJSID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesReplaceInsecureJSIDReplaceInsecureJS:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesReplaceInsecureJSValue string

const (
	SettingEditResponseZonesReplaceInsecureJSValueOn  SettingEditResponseZonesReplaceInsecureJSValue = "on"
	SettingEditResponseZonesReplaceInsecureJSValueOff SettingEditResponseZonesReplaceInsecureJSValue = "off"
)

func (r SettingEditResponseZonesReplaceInsecureJSValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesReplaceInsecureJSValueOn, SettingEditResponseZonesReplaceInsecureJSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesReplaceInsecureJSEditable bool

const (
	SettingEditResponseZonesReplaceInsecureJSEditableTrue  SettingEditResponseZonesReplaceInsecureJSEditable = true
	SettingEditResponseZonesReplaceInsecureJSEditableFalse SettingEditResponseZonesReplaceInsecureJSEditable = false
)

func (r SettingEditResponseZonesReplaceInsecureJSEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesReplaceInsecureJSEditableTrue, SettingEditResponseZonesReplaceInsecureJSEditableFalse:
		return true
	}
	return false
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingEditResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value AutomaticPlatformOptimization `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasAutomaticPlatformOptimizationJSON `json:"-"`
}

// settingEditResponseZonesSchemasAutomaticPlatformOptimizationJSON contains the
// JSON metadata for the struct
// [SettingEditResponseZonesSchemasAutomaticPlatformOptimization]
type settingEditResponseZonesSchemasAutomaticPlatformOptimizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasAutomaticPlatformOptimizationJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasAutomaticPlatformOptimization) implementsZonesSettingEditResponse() {
}

// ID of the zone setting.
type SettingEditResponseZonesSchemasAutomaticPlatformOptimizationID string

const (
	SettingEditResponseZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization SettingEditResponseZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

func (r SettingEditResponseZonesSchemasAutomaticPlatformOptimizationID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
)

func (r SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue, SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse:
		return true
	}
	return false
}

// Allow SHA1 support.
type SettingEditResponseZonesSha1Support struct {
	// Zone setting identifier.
	ID SettingEditResponseZonesSha1SupportID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSha1SupportValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSha1SupportJSON `json:"-"`
}

// settingEditResponseZonesSha1SupportJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesSha1Support]
type settingEditResponseZonesSha1SupportJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSha1Support) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSha1SupportJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSha1Support) implementsZonesSettingEditResponse() {}

// Zone setting identifier.
type SettingEditResponseZonesSha1SupportID string

const (
	SettingEditResponseZonesSha1SupportIDSha1Support SettingEditResponseZonesSha1SupportID = "sha1_support"
)

func (r SettingEditResponseZonesSha1SupportID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSha1SupportIDSha1Support:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSha1SupportValue string

const (
	SettingEditResponseZonesSha1SupportValueOff SettingEditResponseZonesSha1SupportValue = "off"
	SettingEditResponseZonesSha1SupportValueOn  SettingEditResponseZonesSha1SupportValue = "on"
)

func (r SettingEditResponseZonesSha1SupportValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSha1SupportValueOff, SettingEditResponseZonesSha1SupportValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSha1SupportEditable bool

const (
	SettingEditResponseZonesSha1SupportEditableTrue  SettingEditResponseZonesSha1SupportEditable = true
	SettingEditResponseZonesSha1SupportEditableFalse SettingEditResponseZonesSha1SupportEditable = false
)

func (r SettingEditResponseZonesSha1SupportEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSha1SupportEditableTrue, SettingEditResponseZonesSha1SupportEditableFalse:
		return true
	}
	return false
}

// Only allows TLS1.2.
type SettingEditResponseZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID SettingEditResponseZonesTLS1_2OnlyID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesTLS1_2OnlyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesTLS1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesTls1_2OnlyJSON `json:"-"`
}

// settingEditResponseZonesTls1_2OnlyJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesTLS1_2Only]
type settingEditResponseZonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesTLS1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesTls1_2OnlyJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesTLS1_2Only) implementsZonesSettingEditResponse() {}

// Zone setting identifier.
type SettingEditResponseZonesTLS1_2OnlyID string

const (
	SettingEditResponseZonesTLS1_2OnlyIDTLS1_2Only SettingEditResponseZonesTLS1_2OnlyID = "tls_1_2_only"
)

func (r SettingEditResponseZonesTLS1_2OnlyID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTLS1_2OnlyIDTLS1_2Only:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesTLS1_2OnlyValue string

const (
	SettingEditResponseZonesTLS1_2OnlyValueOff SettingEditResponseZonesTLS1_2OnlyValue = "off"
	SettingEditResponseZonesTLS1_2OnlyValueOn  SettingEditResponseZonesTLS1_2OnlyValue = "on"
)

func (r SettingEditResponseZonesTLS1_2OnlyValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTLS1_2OnlyValueOff, SettingEditResponseZonesTLS1_2OnlyValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesTLS1_2OnlyEditable bool

const (
	SettingEditResponseZonesTLS1_2OnlyEditableTrue  SettingEditResponseZonesTLS1_2OnlyEditable = true
	SettingEditResponseZonesTLS1_2OnlyEditableFalse SettingEditResponseZonesTLS1_2OnlyEditable = false
)

func (r SettingEditResponseZonesTLS1_2OnlyEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTLS1_2OnlyEditableTrue, SettingEditResponseZonesTLS1_2OnlyEditableFalse:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseEditable bool

const (
	SettingEditResponseEditableTrue  SettingEditResponseEditable = true
	SettingEditResponseEditableFalse SettingEditResponseEditable = false
)

func (r SettingEditResponseEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseEditableTrue, SettingEditResponseEditableFalse:
		return true
	}
	return false
}

// ID of the zone setting.
type SettingEditResponseID string

const (
	SettingEditResponseID0rtt                          SettingEditResponseID = "0rtt"
	SettingEditResponseIDAdvancedDDoS                  SettingEditResponseID = "advanced_ddos"
	SettingEditResponseIDAlwaysOnline                  SettingEditResponseID = "always_online"
	SettingEditResponseIDAlwaysUseHTTPS                SettingEditResponseID = "always_use_https"
	SettingEditResponseIDAutomaticHTTPSRewrites        SettingEditResponseID = "automatic_https_rewrites"
	SettingEditResponseIDBrotli                        SettingEditResponseID = "brotli"
	SettingEditResponseIDBrowserCacheTTL               SettingEditResponseID = "browser_cache_ttl"
	SettingEditResponseIDBrowserCheck                  SettingEditResponseID = "browser_check"
	SettingEditResponseIDCacheLevel                    SettingEditResponseID = "cache_level"
	SettingEditResponseIDChallengeTTL                  SettingEditResponseID = "challenge_ttl"
	SettingEditResponseIDCiphers                       SettingEditResponseID = "ciphers"
	SettingEditResponseIDCNAMEFlattening               SettingEditResponseID = "cname_flattening"
	SettingEditResponseIDDevelopmentMode               SettingEditResponseID = "development_mode"
	SettingEditResponseIDEarlyHints                    SettingEditResponseID = "early_hints"
	SettingEditResponseIDEdgeCacheTTL                  SettingEditResponseID = "edge_cache_ttl"
	SettingEditResponseIDEmailObfuscation              SettingEditResponseID = "email_obfuscation"
	SettingEditResponseIDH2Prioritization              SettingEditResponseID = "h2_prioritization"
	SettingEditResponseIDHotlinkProtection             SettingEditResponseID = "hotlink_protection"
	SettingEditResponseIDHTTP2                         SettingEditResponseID = "http2"
	SettingEditResponseIDHTTP3                         SettingEditResponseID = "http3"
	SettingEditResponseIDImageResizing                 SettingEditResponseID = "image_resizing"
	SettingEditResponseIDIPGeolocation                 SettingEditResponseID = "ip_geolocation"
	SettingEditResponseIDIPV6                          SettingEditResponseID = "ipv6"
	SettingEditResponseIDMaxUpload                     SettingEditResponseID = "max_upload"
	SettingEditResponseIDMinTLSVersion                 SettingEditResponseID = "min_tls_version"
	SettingEditResponseIDMinify                        SettingEditResponseID = "minify"
	SettingEditResponseIDMirage                        SettingEditResponseID = "mirage"
	SettingEditResponseIDMobileRedirect                SettingEditResponseID = "mobile_redirect"
	SettingEditResponseIDNEL                           SettingEditResponseID = "nel"
	SettingEditResponseIDOpportunisticEncryption       SettingEditResponseID = "opportunistic_encryption"
	SettingEditResponseIDOpportunisticOnion            SettingEditResponseID = "opportunistic_onion"
	SettingEditResponseIDOrangeToOrange                SettingEditResponseID = "orange_to_orange"
	SettingEditResponseIDOriginErrorPagePassThru       SettingEditResponseID = "origin_error_page_pass_thru"
	SettingEditResponseIDPolish                        SettingEditResponseID = "polish"
	SettingEditResponseIDPrefetchPreload               SettingEditResponseID = "prefetch_preload"
	SettingEditResponseIDProxyReadTimeout              SettingEditResponseID = "proxy_read_timeout"
	SettingEditResponseIDPseudoIPV4                    SettingEditResponseID = "pseudo_ipv4"
	SettingEditResponseIDReplaceInsecureJS             SettingEditResponseID = "replace_insecure_js"
	SettingEditResponseIDResponseBuffering             SettingEditResponseID = "response_buffering"
	SettingEditResponseIDRocketLoader                  SettingEditResponseID = "rocket_loader"
	SettingEditResponseIDAutomaticPlatformOptimization SettingEditResponseID = "automatic_platform_optimization"
	SettingEditResponseIDSecurityHeader                SettingEditResponseID = "security_header"
	SettingEditResponseIDSecurityLevel                 SettingEditResponseID = "security_level"
	SettingEditResponseIDServerSideExclude             SettingEditResponseID = "server_side_exclude"
	SettingEditResponseIDSha1Support                   SettingEditResponseID = "sha1_support"
	SettingEditResponseIDSortQueryStringForCache       SettingEditResponseID = "sort_query_string_for_cache"
	SettingEditResponseIDSSL                           SettingEditResponseID = "ssl"
	SettingEditResponseIDSSLRecommender                SettingEditResponseID = "ssl_recommender"
	SettingEditResponseIDTLS1_2Only                    SettingEditResponseID = "tls_1_2_only"
	SettingEditResponseIDTLS1_3                        SettingEditResponseID = "tls_1_3"
	SettingEditResponseIDTLSClientAuth                 SettingEditResponseID = "tls_client_auth"
	SettingEditResponseIDTrueClientIPHeader            SettingEditResponseID = "true_client_ip_header"
	SettingEditResponseIDWAF                           SettingEditResponseID = "waf"
	SettingEditResponseIDWebP                          SettingEditResponseID = "webp"
	SettingEditResponseIDWebsockets                    SettingEditResponseID = "websockets"
)

func (r SettingEditResponseID) IsKnown() bool {
	switch r {
	case SettingEditResponseID0rtt, SettingEditResponseIDAdvancedDDoS, SettingEditResponseIDAlwaysOnline, SettingEditResponseIDAlwaysUseHTTPS, SettingEditResponseIDAutomaticHTTPSRewrites, SettingEditResponseIDBrotli, SettingEditResponseIDBrowserCacheTTL, SettingEditResponseIDBrowserCheck, SettingEditResponseIDCacheLevel, SettingEditResponseIDChallengeTTL, SettingEditResponseIDCiphers, SettingEditResponseIDCNAMEFlattening, SettingEditResponseIDDevelopmentMode, SettingEditResponseIDEarlyHints, SettingEditResponseIDEdgeCacheTTL, SettingEditResponseIDEmailObfuscation, SettingEditResponseIDH2Prioritization, SettingEditResponseIDHotlinkProtection, SettingEditResponseIDHTTP2, SettingEditResponseIDHTTP3, SettingEditResponseIDImageResizing, SettingEditResponseIDIPGeolocation, SettingEditResponseIDIPV6, SettingEditResponseIDMaxUpload, SettingEditResponseIDMinTLSVersion, SettingEditResponseIDMinify, SettingEditResponseIDMirage, SettingEditResponseIDMobileRedirect, SettingEditResponseIDNEL, SettingEditResponseIDOpportunisticEncryption, SettingEditResponseIDOpportunisticOnion, SettingEditResponseIDOrangeToOrange, SettingEditResponseIDOriginErrorPagePassThru, SettingEditResponseIDPolish, SettingEditResponseIDPrefetchPreload, SettingEditResponseIDProxyReadTimeout, SettingEditResponseIDPseudoIPV4, SettingEditResponseIDReplaceInsecureJS, SettingEditResponseIDResponseBuffering, SettingEditResponseIDRocketLoader, SettingEditResponseIDAutomaticPlatformOptimization, SettingEditResponseIDSecurityHeader, SettingEditResponseIDSecurityLevel, SettingEditResponseIDServerSideExclude, SettingEditResponseIDSha1Support, SettingEditResponseIDSortQueryStringForCache, SettingEditResponseIDSSL, SettingEditResponseIDSSLRecommender, SettingEditResponseIDTLS1_2Only, SettingEditResponseIDTLS1_3, SettingEditResponseIDTLSClientAuth, SettingEditResponseIDTrueClientIPHeader, SettingEditResponseIDWAF, SettingEditResponseIDWebP, SettingEditResponseIDWebsockets:
		return true
	}
	return false
}

// 0-RTT session resumption enabled for this zone.
type SettingGetResponse struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseEditable `json:"editable"`
	// ID of the zone setting.
	ID SettingGetResponseID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time   `json:"modified_on,nullable" format:"date-time"`
	Value      interface{} `json:"value,required"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// ssl-recommender enrollment setting.
	Enabled bool                   `json:"enabled"`
	JSON    settingGetResponseJSON `json:"-"`
	union   SettingGetResponseUnion
}

// settingGetResponseJSON contains the JSON metadata for the struct
// [SettingGetResponse]
type settingGetResponseJSON struct {
	Editable      apijson.Field
	ID            apijson.Field
	ModifiedOn    apijson.Field
	Value         apijson.Field
	TimeRemaining apijson.Field
	Enabled       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r settingGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r SettingGetResponse) AsUnion() SettingGetResponseUnion {
	return r.union
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [zones.ZeroRTT], [zones.AdvancedDDoS], [zones.AlwaysOnline],
// [zones.AlwaysUseHTTPS], [zones.AutomaticHTTPSRewrites], [zones.Brotli],
// [zones.BrowserCacheTTL], [zones.BrowserCheck], [zones.CacheLevel],
// [zones.ChallengeTTL], [zones.Ciphers],
// [zones.SettingGetResponseZonesCNAMEFlattening], [zones.DevelopmentMode],
// [zones.EarlyHints], [zones.SettingGetResponseZonesEdgeCacheTTL],
// [zones.EmailObfuscation], [zones.H2Prioritization], [zones.HotlinkProtection],
// [zones.HTTP2], [zones.HTTP3], [zones.ImageResizing], [zones.IPGeolocation],
// [zones.IPV6], [zones.SettingGetResponseZonesMaxUpload], [zones.MinTLSVersion],
// [zones.Minify], [zones.Mirage], [zones.MobileRedirect], [zones.NEL],
// [zones.OpportunisticEncryption], [zones.OpportunisticOnion],
// [zones.OrangeToOrange], [zones.OriginErrorPagePassThru], [zones.Polish],
// [zones.PrefetchPreload], [zones.ProxyReadTimeout], [zones.PseudoIPV4],
// [zones.SettingGetResponseZonesReplaceInsecureJS], [zones.ResponseBuffering],
// [zones.RocketLoader],
// [zones.SettingGetResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.SecurityHeaders], [zones.SecurityLevel], [zones.ServerSideExcludes],
// [zones.SettingGetResponseZonesSha1Support], [zones.SortQueryStringForCache],
// [zones.SSL], [zones.SSLRecommender], [zones.SettingGetResponseZonesTLS1_2Only],
// [zones.TLS1_3], [zones.TLSClientAuth], [zones.TrueClientIPHeader], [zones.WAF],
// [zones.WebP] or [zones.Websocket].
type SettingGetResponseUnion interface {
	implementsZonesSettingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroRTT{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AdvancedDDoS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Brotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Ciphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesCNAMEFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(H2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HTTP2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(HTTP3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPV6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MinTLSVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Minify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Mirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MobileRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Polish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PseudoIPV4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesReplaceInsecureJS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ResponseBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecurityHeaders{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ServerSideExcludes{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SSL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SSLRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesTLS1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TLS1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TLSClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAF{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WebP{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Websocket{}),
		},
	)
}

// Whether or not cname flattening is on.
type SettingGetResponseZonesCNAMEFlattening struct {
	// How to flatten the cname destination.
	ID SettingGetResponseZonesCNAMEFlatteningID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesCNAMEFlatteningValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesCNAMEFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesCNAMEFlatteningJSON `json:"-"`
}

// settingGetResponseZonesCNAMEFlatteningJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesCNAMEFlattening]
type settingGetResponseZonesCNAMEFlatteningJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesCNAMEFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesCNAMEFlatteningJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesCNAMEFlattening) implementsZonesSettingGetResponse() {}

// How to flatten the cname destination.
type SettingGetResponseZonesCNAMEFlatteningID string

const (
	SettingGetResponseZonesCNAMEFlatteningIDCNAMEFlattening SettingGetResponseZonesCNAMEFlatteningID = "cname_flattening"
)

func (r SettingGetResponseZonesCNAMEFlatteningID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesCNAMEFlatteningIDCNAMEFlattening:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesCNAMEFlatteningValue string

const (
	SettingGetResponseZonesCNAMEFlatteningValueFlattenAtRoot SettingGetResponseZonesCNAMEFlatteningValue = "flatten_at_root"
	SettingGetResponseZonesCNAMEFlatteningValueFlattenAll    SettingGetResponseZonesCNAMEFlatteningValue = "flatten_all"
)

func (r SettingGetResponseZonesCNAMEFlatteningValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesCNAMEFlatteningValueFlattenAtRoot, SettingGetResponseZonesCNAMEFlatteningValueFlattenAll:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesCNAMEFlatteningEditable bool

const (
	SettingGetResponseZonesCNAMEFlatteningEditableTrue  SettingGetResponseZonesCNAMEFlatteningEditable = true
	SettingGetResponseZonesCNAMEFlatteningEditableFalse SettingGetResponseZonesCNAMEFlatteningEditable = false
)

func (r SettingGetResponseZonesCNAMEFlatteningEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesCNAMEFlatteningEditableTrue, SettingGetResponseZonesCNAMEFlatteningEditableFalse:
		return true
	}
	return false
}

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type SettingGetResponseZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesEdgeCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesEdgeCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesEdgeCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesEdgeCacheTTLJSON `json:"-"`
}

// settingGetResponseZonesEdgeCacheTTLJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesEdgeCacheTTL]
type settingGetResponseZonesEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesEdgeCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesEdgeCacheTTL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesEdgeCacheTTLID string

const (
	SettingGetResponseZonesEdgeCacheTTLIDEdgeCacheTTL SettingGetResponseZonesEdgeCacheTTLID = "edge_cache_ttl"
)

func (r SettingGetResponseZonesEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesEdgeCacheTTLValue float64

const (
	SettingGetResponseZonesEdgeCacheTTLValue30     SettingGetResponseZonesEdgeCacheTTLValue = 30
	SettingGetResponseZonesEdgeCacheTTLValue60     SettingGetResponseZonesEdgeCacheTTLValue = 60
	SettingGetResponseZonesEdgeCacheTTLValue300    SettingGetResponseZonesEdgeCacheTTLValue = 300
	SettingGetResponseZonesEdgeCacheTTLValue1200   SettingGetResponseZonesEdgeCacheTTLValue = 1200
	SettingGetResponseZonesEdgeCacheTTLValue1800   SettingGetResponseZonesEdgeCacheTTLValue = 1800
	SettingGetResponseZonesEdgeCacheTTLValue3600   SettingGetResponseZonesEdgeCacheTTLValue = 3600
	SettingGetResponseZonesEdgeCacheTTLValue7200   SettingGetResponseZonesEdgeCacheTTLValue = 7200
	SettingGetResponseZonesEdgeCacheTTLValue10800  SettingGetResponseZonesEdgeCacheTTLValue = 10800
	SettingGetResponseZonesEdgeCacheTTLValue14400  SettingGetResponseZonesEdgeCacheTTLValue = 14400
	SettingGetResponseZonesEdgeCacheTTLValue18000  SettingGetResponseZonesEdgeCacheTTLValue = 18000
	SettingGetResponseZonesEdgeCacheTTLValue28800  SettingGetResponseZonesEdgeCacheTTLValue = 28800
	SettingGetResponseZonesEdgeCacheTTLValue43200  SettingGetResponseZonesEdgeCacheTTLValue = 43200
	SettingGetResponseZonesEdgeCacheTTLValue57600  SettingGetResponseZonesEdgeCacheTTLValue = 57600
	SettingGetResponseZonesEdgeCacheTTLValue72000  SettingGetResponseZonesEdgeCacheTTLValue = 72000
	SettingGetResponseZonesEdgeCacheTTLValue86400  SettingGetResponseZonesEdgeCacheTTLValue = 86400
	SettingGetResponseZonesEdgeCacheTTLValue172800 SettingGetResponseZonesEdgeCacheTTLValue = 172800
	SettingGetResponseZonesEdgeCacheTTLValue259200 SettingGetResponseZonesEdgeCacheTTLValue = 259200
	SettingGetResponseZonesEdgeCacheTTLValue345600 SettingGetResponseZonesEdgeCacheTTLValue = 345600
	SettingGetResponseZonesEdgeCacheTTLValue432000 SettingGetResponseZonesEdgeCacheTTLValue = 432000
	SettingGetResponseZonesEdgeCacheTTLValue518400 SettingGetResponseZonesEdgeCacheTTLValue = 518400
	SettingGetResponseZonesEdgeCacheTTLValue604800 SettingGetResponseZonesEdgeCacheTTLValue = 604800
)

func (r SettingGetResponseZonesEdgeCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEdgeCacheTTLValue30, SettingGetResponseZonesEdgeCacheTTLValue60, SettingGetResponseZonesEdgeCacheTTLValue300, SettingGetResponseZonesEdgeCacheTTLValue1200, SettingGetResponseZonesEdgeCacheTTLValue1800, SettingGetResponseZonesEdgeCacheTTLValue3600, SettingGetResponseZonesEdgeCacheTTLValue7200, SettingGetResponseZonesEdgeCacheTTLValue10800, SettingGetResponseZonesEdgeCacheTTLValue14400, SettingGetResponseZonesEdgeCacheTTLValue18000, SettingGetResponseZonesEdgeCacheTTLValue28800, SettingGetResponseZonesEdgeCacheTTLValue43200, SettingGetResponseZonesEdgeCacheTTLValue57600, SettingGetResponseZonesEdgeCacheTTLValue72000, SettingGetResponseZonesEdgeCacheTTLValue86400, SettingGetResponseZonesEdgeCacheTTLValue172800, SettingGetResponseZonesEdgeCacheTTLValue259200, SettingGetResponseZonesEdgeCacheTTLValue345600, SettingGetResponseZonesEdgeCacheTTLValue432000, SettingGetResponseZonesEdgeCacheTTLValue518400, SettingGetResponseZonesEdgeCacheTTLValue604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesEdgeCacheTTLEditable bool

const (
	SettingGetResponseZonesEdgeCacheTTLEditableTrue  SettingGetResponseZonesEdgeCacheTTLEditable = true
	SettingGetResponseZonesEdgeCacheTTLEditableFalse SettingGetResponseZonesEdgeCacheTTLEditable = false
)

func (r SettingGetResponseZonesEdgeCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEdgeCacheTTLEditableTrue, SettingGetResponseZonesEdgeCacheTTLEditableFalse:
		return true
	}
	return false
}

// Maximum size of an allowable upload.
type SettingGetResponseZonesMaxUpload struct {
	// identifier of the zone setting.
	ID SettingGetResponseZonesMaxUploadID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesMaxUploadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesMaxUploadJSON `json:"-"`
}

// settingGetResponseZonesMaxUploadJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesMaxUpload]
type settingGetResponseZonesMaxUploadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesMaxUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesMaxUploadJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesMaxUpload) implementsZonesSettingGetResponse() {}

// identifier of the zone setting.
type SettingGetResponseZonesMaxUploadID string

const (
	SettingGetResponseZonesMaxUploadIDMaxUpload SettingGetResponseZonesMaxUploadID = "max_upload"
)

func (r SettingGetResponseZonesMaxUploadID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesMaxUploadIDMaxUpload:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesMaxUploadValue float64

const (
	SettingGetResponseZonesMaxUploadValue100 SettingGetResponseZonesMaxUploadValue = 100
	SettingGetResponseZonesMaxUploadValue200 SettingGetResponseZonesMaxUploadValue = 200
	SettingGetResponseZonesMaxUploadValue500 SettingGetResponseZonesMaxUploadValue = 500
)

func (r SettingGetResponseZonesMaxUploadValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesMaxUploadValue100, SettingGetResponseZonesMaxUploadValue200, SettingGetResponseZonesMaxUploadValue500:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesMaxUploadEditable bool

const (
	SettingGetResponseZonesMaxUploadEditableTrue  SettingGetResponseZonesMaxUploadEditable = true
	SettingGetResponseZonesMaxUploadEditableFalse SettingGetResponseZonesMaxUploadEditable = false
)

func (r SettingGetResponseZonesMaxUploadEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesMaxUploadEditableTrue, SettingGetResponseZonesMaxUploadEditableFalse:
		return true
	}
	return false
}

// Automatically replace insecure JavaScript libraries with safer and faster
// alternatives provided under cdnjs and powered by Cloudflare. Currently supports
// the following libraries: Polyfill under polyfill.io.
type SettingGetResponseZonesReplaceInsecureJS struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesReplaceInsecureJSID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesReplaceInsecureJSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesReplaceInsecureJSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesReplaceInsecureJSJSON `json:"-"`
}

// settingGetResponseZonesReplaceInsecureJSJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesReplaceInsecureJS]
type settingGetResponseZonesReplaceInsecureJSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesReplaceInsecureJS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesReplaceInsecureJSJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesReplaceInsecureJS) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesReplaceInsecureJSID string

const (
	SettingGetResponseZonesReplaceInsecureJSIDReplaceInsecureJS SettingGetResponseZonesReplaceInsecureJSID = "replace_insecure_js"
)

func (r SettingGetResponseZonesReplaceInsecureJSID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesReplaceInsecureJSIDReplaceInsecureJS:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesReplaceInsecureJSValue string

const (
	SettingGetResponseZonesReplaceInsecureJSValueOn  SettingGetResponseZonesReplaceInsecureJSValue = "on"
	SettingGetResponseZonesReplaceInsecureJSValueOff SettingGetResponseZonesReplaceInsecureJSValue = "off"
)

func (r SettingGetResponseZonesReplaceInsecureJSValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesReplaceInsecureJSValueOn, SettingGetResponseZonesReplaceInsecureJSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesReplaceInsecureJSEditable bool

const (
	SettingGetResponseZonesReplaceInsecureJSEditableTrue  SettingGetResponseZonesReplaceInsecureJSEditable = true
	SettingGetResponseZonesReplaceInsecureJSEditableFalse SettingGetResponseZonesReplaceInsecureJSEditable = false
)

func (r SettingGetResponseZonesReplaceInsecureJSEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesReplaceInsecureJSEditableTrue, SettingGetResponseZonesReplaceInsecureJSEditableFalse:
		return true
	}
	return false
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingGetResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value AutomaticPlatformOptimization `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasAutomaticPlatformOptimizationJSON `json:"-"`
}

// settingGetResponseZonesSchemasAutomaticPlatformOptimizationJSON contains the
// JSON metadata for the struct
// [SettingGetResponseZonesSchemasAutomaticPlatformOptimization]
type settingGetResponseZonesSchemasAutomaticPlatformOptimizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasAutomaticPlatformOptimizationJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasAutomaticPlatformOptimization) implementsZonesSettingGetResponse() {
}

// ID of the zone setting.
type SettingGetResponseZonesSchemasAutomaticPlatformOptimizationID string

const (
	SettingGetResponseZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization SettingGetResponseZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

func (r SettingGetResponseZonesSchemasAutomaticPlatformOptimizationID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
)

func (r SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue, SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse:
		return true
	}
	return false
}

// Allow SHA1 support.
type SettingGetResponseZonesSha1Support struct {
	// Zone setting identifier.
	ID SettingGetResponseZonesSha1SupportID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSha1SupportValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSha1SupportJSON `json:"-"`
}

// settingGetResponseZonesSha1SupportJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesSha1Support]
type settingGetResponseZonesSha1SupportJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSha1Support) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSha1SupportJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSha1Support) implementsZonesSettingGetResponse() {}

// Zone setting identifier.
type SettingGetResponseZonesSha1SupportID string

const (
	SettingGetResponseZonesSha1SupportIDSha1Support SettingGetResponseZonesSha1SupportID = "sha1_support"
)

func (r SettingGetResponseZonesSha1SupportID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSha1SupportIDSha1Support:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSha1SupportValue string

const (
	SettingGetResponseZonesSha1SupportValueOff SettingGetResponseZonesSha1SupportValue = "off"
	SettingGetResponseZonesSha1SupportValueOn  SettingGetResponseZonesSha1SupportValue = "on"
)

func (r SettingGetResponseZonesSha1SupportValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSha1SupportValueOff, SettingGetResponseZonesSha1SupportValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSha1SupportEditable bool

const (
	SettingGetResponseZonesSha1SupportEditableTrue  SettingGetResponseZonesSha1SupportEditable = true
	SettingGetResponseZonesSha1SupportEditableFalse SettingGetResponseZonesSha1SupportEditable = false
)

func (r SettingGetResponseZonesSha1SupportEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSha1SupportEditableTrue, SettingGetResponseZonesSha1SupportEditableFalse:
		return true
	}
	return false
}

// Only allows TLS1.2.
type SettingGetResponseZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID SettingGetResponseZonesTLS1_2OnlyID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesTLS1_2OnlyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesTLS1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesTls1_2OnlyJSON `json:"-"`
}

// settingGetResponseZonesTls1_2OnlyJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesTLS1_2Only]
type settingGetResponseZonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesTLS1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesTls1_2OnlyJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesTLS1_2Only) implementsZonesSettingGetResponse() {}

// Zone setting identifier.
type SettingGetResponseZonesTLS1_2OnlyID string

const (
	SettingGetResponseZonesTLS1_2OnlyIDTLS1_2Only SettingGetResponseZonesTLS1_2OnlyID = "tls_1_2_only"
)

func (r SettingGetResponseZonesTLS1_2OnlyID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTLS1_2OnlyIDTLS1_2Only:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesTLS1_2OnlyValue string

const (
	SettingGetResponseZonesTLS1_2OnlyValueOff SettingGetResponseZonesTLS1_2OnlyValue = "off"
	SettingGetResponseZonesTLS1_2OnlyValueOn  SettingGetResponseZonesTLS1_2OnlyValue = "on"
)

func (r SettingGetResponseZonesTLS1_2OnlyValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTLS1_2OnlyValueOff, SettingGetResponseZonesTLS1_2OnlyValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesTLS1_2OnlyEditable bool

const (
	SettingGetResponseZonesTLS1_2OnlyEditableTrue  SettingGetResponseZonesTLS1_2OnlyEditable = true
	SettingGetResponseZonesTLS1_2OnlyEditableFalse SettingGetResponseZonesTLS1_2OnlyEditable = false
)

func (r SettingGetResponseZonesTLS1_2OnlyEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTLS1_2OnlyEditableTrue, SettingGetResponseZonesTLS1_2OnlyEditableFalse:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseEditable bool

const (
	SettingGetResponseEditableTrue  SettingGetResponseEditable = true
	SettingGetResponseEditableFalse SettingGetResponseEditable = false
)

func (r SettingGetResponseEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseEditableTrue, SettingGetResponseEditableFalse:
		return true
	}
	return false
}

// ID of the zone setting.
type SettingGetResponseID string

const (
	SettingGetResponseID0rtt                          SettingGetResponseID = "0rtt"
	SettingGetResponseIDAdvancedDDoS                  SettingGetResponseID = "advanced_ddos"
	SettingGetResponseIDAlwaysOnline                  SettingGetResponseID = "always_online"
	SettingGetResponseIDAlwaysUseHTTPS                SettingGetResponseID = "always_use_https"
	SettingGetResponseIDAutomaticHTTPSRewrites        SettingGetResponseID = "automatic_https_rewrites"
	SettingGetResponseIDBrotli                        SettingGetResponseID = "brotli"
	SettingGetResponseIDBrowserCacheTTL               SettingGetResponseID = "browser_cache_ttl"
	SettingGetResponseIDBrowserCheck                  SettingGetResponseID = "browser_check"
	SettingGetResponseIDCacheLevel                    SettingGetResponseID = "cache_level"
	SettingGetResponseIDChallengeTTL                  SettingGetResponseID = "challenge_ttl"
	SettingGetResponseIDCiphers                       SettingGetResponseID = "ciphers"
	SettingGetResponseIDCNAMEFlattening               SettingGetResponseID = "cname_flattening"
	SettingGetResponseIDDevelopmentMode               SettingGetResponseID = "development_mode"
	SettingGetResponseIDEarlyHints                    SettingGetResponseID = "early_hints"
	SettingGetResponseIDEdgeCacheTTL                  SettingGetResponseID = "edge_cache_ttl"
	SettingGetResponseIDEmailObfuscation              SettingGetResponseID = "email_obfuscation"
	SettingGetResponseIDH2Prioritization              SettingGetResponseID = "h2_prioritization"
	SettingGetResponseIDHotlinkProtection             SettingGetResponseID = "hotlink_protection"
	SettingGetResponseIDHTTP2                         SettingGetResponseID = "http2"
	SettingGetResponseIDHTTP3                         SettingGetResponseID = "http3"
	SettingGetResponseIDImageResizing                 SettingGetResponseID = "image_resizing"
	SettingGetResponseIDIPGeolocation                 SettingGetResponseID = "ip_geolocation"
	SettingGetResponseIDIPV6                          SettingGetResponseID = "ipv6"
	SettingGetResponseIDMaxUpload                     SettingGetResponseID = "max_upload"
	SettingGetResponseIDMinTLSVersion                 SettingGetResponseID = "min_tls_version"
	SettingGetResponseIDMinify                        SettingGetResponseID = "minify"
	SettingGetResponseIDMirage                        SettingGetResponseID = "mirage"
	SettingGetResponseIDMobileRedirect                SettingGetResponseID = "mobile_redirect"
	SettingGetResponseIDNEL                           SettingGetResponseID = "nel"
	SettingGetResponseIDOpportunisticEncryption       SettingGetResponseID = "opportunistic_encryption"
	SettingGetResponseIDOpportunisticOnion            SettingGetResponseID = "opportunistic_onion"
	SettingGetResponseIDOrangeToOrange                SettingGetResponseID = "orange_to_orange"
	SettingGetResponseIDOriginErrorPagePassThru       SettingGetResponseID = "origin_error_page_pass_thru"
	SettingGetResponseIDPolish                        SettingGetResponseID = "polish"
	SettingGetResponseIDPrefetchPreload               SettingGetResponseID = "prefetch_preload"
	SettingGetResponseIDProxyReadTimeout              SettingGetResponseID = "proxy_read_timeout"
	SettingGetResponseIDPseudoIPV4                    SettingGetResponseID = "pseudo_ipv4"
	SettingGetResponseIDReplaceInsecureJS             SettingGetResponseID = "replace_insecure_js"
	SettingGetResponseIDResponseBuffering             SettingGetResponseID = "response_buffering"
	SettingGetResponseIDRocketLoader                  SettingGetResponseID = "rocket_loader"
	SettingGetResponseIDAutomaticPlatformOptimization SettingGetResponseID = "automatic_platform_optimization"
	SettingGetResponseIDSecurityHeader                SettingGetResponseID = "security_header"
	SettingGetResponseIDSecurityLevel                 SettingGetResponseID = "security_level"
	SettingGetResponseIDServerSideExclude             SettingGetResponseID = "server_side_exclude"
	SettingGetResponseIDSha1Support                   SettingGetResponseID = "sha1_support"
	SettingGetResponseIDSortQueryStringForCache       SettingGetResponseID = "sort_query_string_for_cache"
	SettingGetResponseIDSSL                           SettingGetResponseID = "ssl"
	SettingGetResponseIDSSLRecommender                SettingGetResponseID = "ssl_recommender"
	SettingGetResponseIDTLS1_2Only                    SettingGetResponseID = "tls_1_2_only"
	SettingGetResponseIDTLS1_3                        SettingGetResponseID = "tls_1_3"
	SettingGetResponseIDTLSClientAuth                 SettingGetResponseID = "tls_client_auth"
	SettingGetResponseIDTrueClientIPHeader            SettingGetResponseID = "true_client_ip_header"
	SettingGetResponseIDWAF                           SettingGetResponseID = "waf"
	SettingGetResponseIDWebP                          SettingGetResponseID = "webp"
	SettingGetResponseIDWebsockets                    SettingGetResponseID = "websockets"
)

func (r SettingGetResponseID) IsKnown() bool {
	switch r {
	case SettingGetResponseID0rtt, SettingGetResponseIDAdvancedDDoS, SettingGetResponseIDAlwaysOnline, SettingGetResponseIDAlwaysUseHTTPS, SettingGetResponseIDAutomaticHTTPSRewrites, SettingGetResponseIDBrotli, SettingGetResponseIDBrowserCacheTTL, SettingGetResponseIDBrowserCheck, SettingGetResponseIDCacheLevel, SettingGetResponseIDChallengeTTL, SettingGetResponseIDCiphers, SettingGetResponseIDCNAMEFlattening, SettingGetResponseIDDevelopmentMode, SettingGetResponseIDEarlyHints, SettingGetResponseIDEdgeCacheTTL, SettingGetResponseIDEmailObfuscation, SettingGetResponseIDH2Prioritization, SettingGetResponseIDHotlinkProtection, SettingGetResponseIDHTTP2, SettingGetResponseIDHTTP3, SettingGetResponseIDImageResizing, SettingGetResponseIDIPGeolocation, SettingGetResponseIDIPV6, SettingGetResponseIDMaxUpload, SettingGetResponseIDMinTLSVersion, SettingGetResponseIDMinify, SettingGetResponseIDMirage, SettingGetResponseIDMobileRedirect, SettingGetResponseIDNEL, SettingGetResponseIDOpportunisticEncryption, SettingGetResponseIDOpportunisticOnion, SettingGetResponseIDOrangeToOrange, SettingGetResponseIDOriginErrorPagePassThru, SettingGetResponseIDPolish, SettingGetResponseIDPrefetchPreload, SettingGetResponseIDProxyReadTimeout, SettingGetResponseIDPseudoIPV4, SettingGetResponseIDReplaceInsecureJS, SettingGetResponseIDResponseBuffering, SettingGetResponseIDRocketLoader, SettingGetResponseIDAutomaticPlatformOptimization, SettingGetResponseIDSecurityHeader, SettingGetResponseIDSecurityLevel, SettingGetResponseIDServerSideExclude, SettingGetResponseIDSha1Support, SettingGetResponseIDSortQueryStringForCache, SettingGetResponseIDSSL, SettingGetResponseIDSSLRecommender, SettingGetResponseIDTLS1_2Only, SettingGetResponseIDTLS1_3, SettingGetResponseIDTLSClientAuth, SettingGetResponseIDTrueClientIPHeader, SettingGetResponseIDWAF, SettingGetResponseIDWebP, SettingGetResponseIDWebsockets:
		return true
	}
	return false
}

type SettingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// 0-RTT session resumption enabled for this zone.
	Body SettingEditParamsBodyUnion `json:"body,required"`
}

func (r SettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// 0-RTT session resumption enabled for this zone.
type SettingEditParamsBody struct {
	// ID of the zone setting.
	ID    param.Field[SettingEditParamsBodyID] `json:"id"`
	Value param.Field[interface{}]             `json:"value,required"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SettingEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBody) implementsZonesSettingEditParamsBodyUnion() {}

// 0-RTT session resumption enabled for this zone.
//
// Satisfied by [zones.ZeroRTTParam], [zones.AdvancedDDoSParam],
// [zones.AlwaysOnlineParam], [zones.AlwaysUseHTTPSParam],
// [zones.AutomaticHTTPSRewritesParam], [zones.BrotliParam],
// [zones.BrowserCacheTTLParam], [zones.BrowserCheckParam],
// [zones.CacheLevelParam], [zones.ChallengeTTLParam], [zones.CiphersParam],
// [zones.SettingEditParamsBodyZonesCNAMEFlattening], [zones.DevelopmentModeParam],
// [zones.EarlyHintsParam], [zones.SettingEditParamsBodyZonesEdgeCacheTTL],
// [zones.EmailObfuscationParam], [zones.H2PrioritizationParam],
// [zones.HotlinkProtectionParam], [zones.HTTP2Param], [zones.HTTP3Param],
// [zones.ImageResizingParam], [zones.IPGeolocationParam], [zones.IPV6Param],
// [zones.SettingEditParamsBodyZonesMaxUpload], [zones.MinTLSVersionParam],
// [zones.MinifyParam], [zones.MirageParam], [zones.MobileRedirectParam],
// [zones.NELParam], [zones.OpportunisticEncryptionParam],
// [zones.OpportunisticOnionParam], [zones.OrangeToOrangeParam],
// [zones.OriginErrorPagePassThruParam], [zones.PolishParam],
// [zones.PrefetchPreloadParam], [zones.ProxyReadTimeoutParam],
// [zones.PseudoIPV4Param], [zones.SettingEditParamsBodyZonesReplaceInsecureJS],
// [zones.ResponseBufferingParam], [zones.RocketLoaderParam],
// [zones.SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimization],
// [zones.SecurityHeadersParam], [zones.SecurityLevelParam],
// [zones.ServerSideExcludesParam], [zones.SettingEditParamsBodyZonesSha1Support],
// [zones.SortQueryStringForCacheParam], [zones.SSLParam],
// [zones.SSLRecommenderParam], [zones.SettingEditParamsBodyZonesTLS1_2Only],
// [zones.TLS1_3Param], [zones.TLSClientAuthParam],
// [zones.TrueClientIPHeaderParam], [zones.WAFParam], [zones.WebPParam],
// [zones.WebsocketParam], [SettingEditParamsBody].
type SettingEditParamsBodyUnion interface {
	implementsZonesSettingEditParamsBodyUnion()
}

// Whether or not cname flattening is on.
type SettingEditParamsBodyZonesCNAMEFlattening struct {
	// How to flatten the cname destination.
	ID param.Field[SettingEditParamsBodyZonesCNAMEFlatteningID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesCNAMEFlatteningValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesCNAMEFlattening) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesCNAMEFlattening) implementsZonesSettingEditParamsBodyUnion() {}

// How to flatten the cname destination.
type SettingEditParamsBodyZonesCNAMEFlatteningID string

const (
	SettingEditParamsBodyZonesCNAMEFlatteningIDCNAMEFlattening SettingEditParamsBodyZonesCNAMEFlatteningID = "cname_flattening"
)

func (r SettingEditParamsBodyZonesCNAMEFlatteningID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesCNAMEFlatteningIDCNAMEFlattening:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesCNAMEFlatteningValue string

const (
	SettingEditParamsBodyZonesCNAMEFlatteningValueFlattenAtRoot SettingEditParamsBodyZonesCNAMEFlatteningValue = "flatten_at_root"
	SettingEditParamsBodyZonesCNAMEFlatteningValueFlattenAll    SettingEditParamsBodyZonesCNAMEFlatteningValue = "flatten_all"
)

func (r SettingEditParamsBodyZonesCNAMEFlatteningValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesCNAMEFlatteningValueFlattenAtRoot, SettingEditParamsBodyZonesCNAMEFlatteningValueFlattenAll:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesCNAMEFlatteningEditable bool

const (
	SettingEditParamsBodyZonesCNAMEFlatteningEditableTrue  SettingEditParamsBodyZonesCNAMEFlatteningEditable = true
	SettingEditParamsBodyZonesCNAMEFlatteningEditableFalse SettingEditParamsBodyZonesCNAMEFlatteningEditable = false
)

func (r SettingEditParamsBodyZonesCNAMEFlatteningEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesCNAMEFlatteningEditableTrue, SettingEditParamsBodyZonesCNAMEFlatteningEditableFalse:
		return true
	}
	return false
}

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type SettingEditParamsBodyZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesEdgeCacheTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesEdgeCacheTTLValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesEdgeCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesEdgeCacheTTL) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesEdgeCacheTTLID string

const (
	SettingEditParamsBodyZonesEdgeCacheTTLIDEdgeCacheTTL SettingEditParamsBodyZonesEdgeCacheTTLID = "edge_cache_ttl"
)

func (r SettingEditParamsBodyZonesEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesEdgeCacheTTLValue float64

const (
	SettingEditParamsBodyZonesEdgeCacheTTLValue30     SettingEditParamsBodyZonesEdgeCacheTTLValue = 30
	SettingEditParamsBodyZonesEdgeCacheTTLValue60     SettingEditParamsBodyZonesEdgeCacheTTLValue = 60
	SettingEditParamsBodyZonesEdgeCacheTTLValue300    SettingEditParamsBodyZonesEdgeCacheTTLValue = 300
	SettingEditParamsBodyZonesEdgeCacheTTLValue1200   SettingEditParamsBodyZonesEdgeCacheTTLValue = 1200
	SettingEditParamsBodyZonesEdgeCacheTTLValue1800   SettingEditParamsBodyZonesEdgeCacheTTLValue = 1800
	SettingEditParamsBodyZonesEdgeCacheTTLValue3600   SettingEditParamsBodyZonesEdgeCacheTTLValue = 3600
	SettingEditParamsBodyZonesEdgeCacheTTLValue7200   SettingEditParamsBodyZonesEdgeCacheTTLValue = 7200
	SettingEditParamsBodyZonesEdgeCacheTTLValue10800  SettingEditParamsBodyZonesEdgeCacheTTLValue = 10800
	SettingEditParamsBodyZonesEdgeCacheTTLValue14400  SettingEditParamsBodyZonesEdgeCacheTTLValue = 14400
	SettingEditParamsBodyZonesEdgeCacheTTLValue18000  SettingEditParamsBodyZonesEdgeCacheTTLValue = 18000
	SettingEditParamsBodyZonesEdgeCacheTTLValue28800  SettingEditParamsBodyZonesEdgeCacheTTLValue = 28800
	SettingEditParamsBodyZonesEdgeCacheTTLValue43200  SettingEditParamsBodyZonesEdgeCacheTTLValue = 43200
	SettingEditParamsBodyZonesEdgeCacheTTLValue57600  SettingEditParamsBodyZonesEdgeCacheTTLValue = 57600
	SettingEditParamsBodyZonesEdgeCacheTTLValue72000  SettingEditParamsBodyZonesEdgeCacheTTLValue = 72000
	SettingEditParamsBodyZonesEdgeCacheTTLValue86400  SettingEditParamsBodyZonesEdgeCacheTTLValue = 86400
	SettingEditParamsBodyZonesEdgeCacheTTLValue172800 SettingEditParamsBodyZonesEdgeCacheTTLValue = 172800
	SettingEditParamsBodyZonesEdgeCacheTTLValue259200 SettingEditParamsBodyZonesEdgeCacheTTLValue = 259200
	SettingEditParamsBodyZonesEdgeCacheTTLValue345600 SettingEditParamsBodyZonesEdgeCacheTTLValue = 345600
	SettingEditParamsBodyZonesEdgeCacheTTLValue432000 SettingEditParamsBodyZonesEdgeCacheTTLValue = 432000
	SettingEditParamsBodyZonesEdgeCacheTTLValue518400 SettingEditParamsBodyZonesEdgeCacheTTLValue = 518400
	SettingEditParamsBodyZonesEdgeCacheTTLValue604800 SettingEditParamsBodyZonesEdgeCacheTTLValue = 604800
)

func (r SettingEditParamsBodyZonesEdgeCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesEdgeCacheTTLValue30, SettingEditParamsBodyZonesEdgeCacheTTLValue60, SettingEditParamsBodyZonesEdgeCacheTTLValue300, SettingEditParamsBodyZonesEdgeCacheTTLValue1200, SettingEditParamsBodyZonesEdgeCacheTTLValue1800, SettingEditParamsBodyZonesEdgeCacheTTLValue3600, SettingEditParamsBodyZonesEdgeCacheTTLValue7200, SettingEditParamsBodyZonesEdgeCacheTTLValue10800, SettingEditParamsBodyZonesEdgeCacheTTLValue14400, SettingEditParamsBodyZonesEdgeCacheTTLValue18000, SettingEditParamsBodyZonesEdgeCacheTTLValue28800, SettingEditParamsBodyZonesEdgeCacheTTLValue43200, SettingEditParamsBodyZonesEdgeCacheTTLValue57600, SettingEditParamsBodyZonesEdgeCacheTTLValue72000, SettingEditParamsBodyZonesEdgeCacheTTLValue86400, SettingEditParamsBodyZonesEdgeCacheTTLValue172800, SettingEditParamsBodyZonesEdgeCacheTTLValue259200, SettingEditParamsBodyZonesEdgeCacheTTLValue345600, SettingEditParamsBodyZonesEdgeCacheTTLValue432000, SettingEditParamsBodyZonesEdgeCacheTTLValue518400, SettingEditParamsBodyZonesEdgeCacheTTLValue604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesEdgeCacheTTLEditable bool

const (
	SettingEditParamsBodyZonesEdgeCacheTTLEditableTrue  SettingEditParamsBodyZonesEdgeCacheTTLEditable = true
	SettingEditParamsBodyZonesEdgeCacheTTLEditableFalse SettingEditParamsBodyZonesEdgeCacheTTLEditable = false
)

func (r SettingEditParamsBodyZonesEdgeCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesEdgeCacheTTLEditableTrue, SettingEditParamsBodyZonesEdgeCacheTTLEditableFalse:
		return true
	}
	return false
}

// Maximum size of an allowable upload.
type SettingEditParamsBodyZonesMaxUpload struct {
	// identifier of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesMaxUploadID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesMaxUploadValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesMaxUpload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesMaxUpload) implementsZonesSettingEditParamsBodyUnion() {}

// identifier of the zone setting.
type SettingEditParamsBodyZonesMaxUploadID string

const (
	SettingEditParamsBodyZonesMaxUploadIDMaxUpload SettingEditParamsBodyZonesMaxUploadID = "max_upload"
)

func (r SettingEditParamsBodyZonesMaxUploadID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesMaxUploadIDMaxUpload:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesMaxUploadValue float64

const (
	SettingEditParamsBodyZonesMaxUploadValue100 SettingEditParamsBodyZonesMaxUploadValue = 100
	SettingEditParamsBodyZonesMaxUploadValue200 SettingEditParamsBodyZonesMaxUploadValue = 200
	SettingEditParamsBodyZonesMaxUploadValue500 SettingEditParamsBodyZonesMaxUploadValue = 500
)

func (r SettingEditParamsBodyZonesMaxUploadValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesMaxUploadValue100, SettingEditParamsBodyZonesMaxUploadValue200, SettingEditParamsBodyZonesMaxUploadValue500:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesMaxUploadEditable bool

const (
	SettingEditParamsBodyZonesMaxUploadEditableTrue  SettingEditParamsBodyZonesMaxUploadEditable = true
	SettingEditParamsBodyZonesMaxUploadEditableFalse SettingEditParamsBodyZonesMaxUploadEditable = false
)

func (r SettingEditParamsBodyZonesMaxUploadEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesMaxUploadEditableTrue, SettingEditParamsBodyZonesMaxUploadEditableFalse:
		return true
	}
	return false
}

// Automatically replace insecure JavaScript libraries with safer and faster
// alternatives provided under cdnjs and powered by Cloudflare. Currently supports
// the following libraries: Polyfill under polyfill.io.
type SettingEditParamsBodyZonesReplaceInsecureJS struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesReplaceInsecureJSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesReplaceInsecureJSValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesReplaceInsecureJS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesReplaceInsecureJS) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesReplaceInsecureJSID string

const (
	SettingEditParamsBodyZonesReplaceInsecureJSIDReplaceInsecureJS SettingEditParamsBodyZonesReplaceInsecureJSID = "replace_insecure_js"
)

func (r SettingEditParamsBodyZonesReplaceInsecureJSID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesReplaceInsecureJSIDReplaceInsecureJS:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesReplaceInsecureJSValue string

const (
	SettingEditParamsBodyZonesReplaceInsecureJSValueOn  SettingEditParamsBodyZonesReplaceInsecureJSValue = "on"
	SettingEditParamsBodyZonesReplaceInsecureJSValueOff SettingEditParamsBodyZonesReplaceInsecureJSValue = "off"
)

func (r SettingEditParamsBodyZonesReplaceInsecureJSValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesReplaceInsecureJSValueOn, SettingEditParamsBodyZonesReplaceInsecureJSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesReplaceInsecureJSEditable bool

const (
	SettingEditParamsBodyZonesReplaceInsecureJSEditableTrue  SettingEditParamsBodyZonesReplaceInsecureJSEditable = true
	SettingEditParamsBodyZonesReplaceInsecureJSEditableFalse SettingEditParamsBodyZonesReplaceInsecureJSEditable = false
)

func (r SettingEditParamsBodyZonesReplaceInsecureJSEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesReplaceInsecureJSEditableTrue, SettingEditParamsBodyZonesReplaceInsecureJSEditableFalse:
		return true
	}
	return false
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[AutomaticPlatformOptimizationParam] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimization) implementsZonesSettingEditParamsBodyUnion() {
}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationID string

const (
	SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

func (r SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationEditable = false
)

func (r SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationEditableTrue, SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimizationEditableFalse:
		return true
	}
	return false
}

// Allow SHA1 support.
type SettingEditParamsBodyZonesSha1Support struct {
	// Zone setting identifier.
	ID param.Field[SettingEditParamsBodyZonesSha1SupportID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSha1SupportValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSha1Support) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSha1Support) implementsZonesSettingEditParamsBodyUnion() {}

// Zone setting identifier.
type SettingEditParamsBodyZonesSha1SupportID string

const (
	SettingEditParamsBodyZonesSha1SupportIDSha1Support SettingEditParamsBodyZonesSha1SupportID = "sha1_support"
)

func (r SettingEditParamsBodyZonesSha1SupportID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSha1SupportIDSha1Support:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSha1SupportValue string

const (
	SettingEditParamsBodyZonesSha1SupportValueOff SettingEditParamsBodyZonesSha1SupportValue = "off"
	SettingEditParamsBodyZonesSha1SupportValueOn  SettingEditParamsBodyZonesSha1SupportValue = "on"
)

func (r SettingEditParamsBodyZonesSha1SupportValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSha1SupportValueOff, SettingEditParamsBodyZonesSha1SupportValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSha1SupportEditable bool

const (
	SettingEditParamsBodyZonesSha1SupportEditableTrue  SettingEditParamsBodyZonesSha1SupportEditable = true
	SettingEditParamsBodyZonesSha1SupportEditableFalse SettingEditParamsBodyZonesSha1SupportEditable = false
)

func (r SettingEditParamsBodyZonesSha1SupportEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSha1SupportEditableTrue, SettingEditParamsBodyZonesSha1SupportEditableFalse:
		return true
	}
	return false
}

// Only allows TLS1.2.
type SettingEditParamsBodyZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID param.Field[SettingEditParamsBodyZonesTLS1_2OnlyID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesTLS1_2OnlyValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesTLS1_2Only) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesTLS1_2Only) implementsZonesSettingEditParamsBodyUnion() {}

// Zone setting identifier.
type SettingEditParamsBodyZonesTLS1_2OnlyID string

const (
	SettingEditParamsBodyZonesTLS1_2OnlyIDTLS1_2Only SettingEditParamsBodyZonesTLS1_2OnlyID = "tls_1_2_only"
)

func (r SettingEditParamsBodyZonesTLS1_2OnlyID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesTLS1_2OnlyIDTLS1_2Only:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesTLS1_2OnlyValue string

const (
	SettingEditParamsBodyZonesTLS1_2OnlyValueOff SettingEditParamsBodyZonesTLS1_2OnlyValue = "off"
	SettingEditParamsBodyZonesTLS1_2OnlyValueOn  SettingEditParamsBodyZonesTLS1_2OnlyValue = "on"
)

func (r SettingEditParamsBodyZonesTLS1_2OnlyValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesTLS1_2OnlyValueOff, SettingEditParamsBodyZonesTLS1_2OnlyValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesTLS1_2OnlyEditable bool

const (
	SettingEditParamsBodyZonesTLS1_2OnlyEditableTrue  SettingEditParamsBodyZonesTLS1_2OnlyEditable = true
	SettingEditParamsBodyZonesTLS1_2OnlyEditableFalse SettingEditParamsBodyZonesTLS1_2OnlyEditable = false
)

func (r SettingEditParamsBodyZonesTLS1_2OnlyEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesTLS1_2OnlyEditableTrue, SettingEditParamsBodyZonesTLS1_2OnlyEditableFalse:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyEditable bool

const (
	SettingEditParamsBodyEditableTrue  SettingEditParamsBodyEditable = true
	SettingEditParamsBodyEditableFalse SettingEditParamsBodyEditable = false
)

func (r SettingEditParamsBodyEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyEditableTrue, SettingEditParamsBodyEditableFalse:
		return true
	}
	return false
}

// ID of the zone setting.
type SettingEditParamsBodyID string

const (
	SettingEditParamsBodyID0rtt                          SettingEditParamsBodyID = "0rtt"
	SettingEditParamsBodyIDAdvancedDDoS                  SettingEditParamsBodyID = "advanced_ddos"
	SettingEditParamsBodyIDAlwaysOnline                  SettingEditParamsBodyID = "always_online"
	SettingEditParamsBodyIDAlwaysUseHTTPS                SettingEditParamsBodyID = "always_use_https"
	SettingEditParamsBodyIDAutomaticHTTPSRewrites        SettingEditParamsBodyID = "automatic_https_rewrites"
	SettingEditParamsBodyIDBrotli                        SettingEditParamsBodyID = "brotli"
	SettingEditParamsBodyIDBrowserCacheTTL               SettingEditParamsBodyID = "browser_cache_ttl"
	SettingEditParamsBodyIDBrowserCheck                  SettingEditParamsBodyID = "browser_check"
	SettingEditParamsBodyIDCacheLevel                    SettingEditParamsBodyID = "cache_level"
	SettingEditParamsBodyIDChallengeTTL                  SettingEditParamsBodyID = "challenge_ttl"
	SettingEditParamsBodyIDCiphers                       SettingEditParamsBodyID = "ciphers"
	SettingEditParamsBodyIDCNAMEFlattening               SettingEditParamsBodyID = "cname_flattening"
	SettingEditParamsBodyIDDevelopmentMode               SettingEditParamsBodyID = "development_mode"
	SettingEditParamsBodyIDEarlyHints                    SettingEditParamsBodyID = "early_hints"
	SettingEditParamsBodyIDEdgeCacheTTL                  SettingEditParamsBodyID = "edge_cache_ttl"
	SettingEditParamsBodyIDEmailObfuscation              SettingEditParamsBodyID = "email_obfuscation"
	SettingEditParamsBodyIDH2Prioritization              SettingEditParamsBodyID = "h2_prioritization"
	SettingEditParamsBodyIDHotlinkProtection             SettingEditParamsBodyID = "hotlink_protection"
	SettingEditParamsBodyIDHTTP2                         SettingEditParamsBodyID = "http2"
	SettingEditParamsBodyIDHTTP3                         SettingEditParamsBodyID = "http3"
	SettingEditParamsBodyIDImageResizing                 SettingEditParamsBodyID = "image_resizing"
	SettingEditParamsBodyIDIPGeolocation                 SettingEditParamsBodyID = "ip_geolocation"
	SettingEditParamsBodyIDIPV6                          SettingEditParamsBodyID = "ipv6"
	SettingEditParamsBodyIDMaxUpload                     SettingEditParamsBodyID = "max_upload"
	SettingEditParamsBodyIDMinTLSVersion                 SettingEditParamsBodyID = "min_tls_version"
	SettingEditParamsBodyIDMinify                        SettingEditParamsBodyID = "minify"
	SettingEditParamsBodyIDMirage                        SettingEditParamsBodyID = "mirage"
	SettingEditParamsBodyIDMobileRedirect                SettingEditParamsBodyID = "mobile_redirect"
	SettingEditParamsBodyIDNEL                           SettingEditParamsBodyID = "nel"
	SettingEditParamsBodyIDOpportunisticEncryption       SettingEditParamsBodyID = "opportunistic_encryption"
	SettingEditParamsBodyIDOpportunisticOnion            SettingEditParamsBodyID = "opportunistic_onion"
	SettingEditParamsBodyIDOrangeToOrange                SettingEditParamsBodyID = "orange_to_orange"
	SettingEditParamsBodyIDOriginErrorPagePassThru       SettingEditParamsBodyID = "origin_error_page_pass_thru"
	SettingEditParamsBodyIDPolish                        SettingEditParamsBodyID = "polish"
	SettingEditParamsBodyIDPrefetchPreload               SettingEditParamsBodyID = "prefetch_preload"
	SettingEditParamsBodyIDProxyReadTimeout              SettingEditParamsBodyID = "proxy_read_timeout"
	SettingEditParamsBodyIDPseudoIPV4                    SettingEditParamsBodyID = "pseudo_ipv4"
	SettingEditParamsBodyIDReplaceInsecureJS             SettingEditParamsBodyID = "replace_insecure_js"
	SettingEditParamsBodyIDResponseBuffering             SettingEditParamsBodyID = "response_buffering"
	SettingEditParamsBodyIDRocketLoader                  SettingEditParamsBodyID = "rocket_loader"
	SettingEditParamsBodyIDAutomaticPlatformOptimization SettingEditParamsBodyID = "automatic_platform_optimization"
	SettingEditParamsBodyIDSecurityHeader                SettingEditParamsBodyID = "security_header"
	SettingEditParamsBodyIDSecurityLevel                 SettingEditParamsBodyID = "security_level"
	SettingEditParamsBodyIDServerSideExclude             SettingEditParamsBodyID = "server_side_exclude"
	SettingEditParamsBodyIDSha1Support                   SettingEditParamsBodyID = "sha1_support"
	SettingEditParamsBodyIDSortQueryStringForCache       SettingEditParamsBodyID = "sort_query_string_for_cache"
	SettingEditParamsBodyIDSSL                           SettingEditParamsBodyID = "ssl"
	SettingEditParamsBodyIDSSLRecommender                SettingEditParamsBodyID = "ssl_recommender"
	SettingEditParamsBodyIDTLS1_2Only                    SettingEditParamsBodyID = "tls_1_2_only"
	SettingEditParamsBodyIDTLS1_3                        SettingEditParamsBodyID = "tls_1_3"
	SettingEditParamsBodyIDTLSClientAuth                 SettingEditParamsBodyID = "tls_client_auth"
	SettingEditParamsBodyIDTrueClientIPHeader            SettingEditParamsBodyID = "true_client_ip_header"
	SettingEditParamsBodyIDWAF                           SettingEditParamsBodyID = "waf"
	SettingEditParamsBodyIDWebP                          SettingEditParamsBodyID = "webp"
	SettingEditParamsBodyIDWebsockets                    SettingEditParamsBodyID = "websockets"
)

func (r SettingEditParamsBodyID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyID0rtt, SettingEditParamsBodyIDAdvancedDDoS, SettingEditParamsBodyIDAlwaysOnline, SettingEditParamsBodyIDAlwaysUseHTTPS, SettingEditParamsBodyIDAutomaticHTTPSRewrites, SettingEditParamsBodyIDBrotli, SettingEditParamsBodyIDBrowserCacheTTL, SettingEditParamsBodyIDBrowserCheck, SettingEditParamsBodyIDCacheLevel, SettingEditParamsBodyIDChallengeTTL, SettingEditParamsBodyIDCiphers, SettingEditParamsBodyIDCNAMEFlattening, SettingEditParamsBodyIDDevelopmentMode, SettingEditParamsBodyIDEarlyHints, SettingEditParamsBodyIDEdgeCacheTTL, SettingEditParamsBodyIDEmailObfuscation, SettingEditParamsBodyIDH2Prioritization, SettingEditParamsBodyIDHotlinkProtection, SettingEditParamsBodyIDHTTP2, SettingEditParamsBodyIDHTTP3, SettingEditParamsBodyIDImageResizing, SettingEditParamsBodyIDIPGeolocation, SettingEditParamsBodyIDIPV6, SettingEditParamsBodyIDMaxUpload, SettingEditParamsBodyIDMinTLSVersion, SettingEditParamsBodyIDMinify, SettingEditParamsBodyIDMirage, SettingEditParamsBodyIDMobileRedirect, SettingEditParamsBodyIDNEL, SettingEditParamsBodyIDOpportunisticEncryption, SettingEditParamsBodyIDOpportunisticOnion, SettingEditParamsBodyIDOrangeToOrange, SettingEditParamsBodyIDOriginErrorPagePassThru, SettingEditParamsBodyIDPolish, SettingEditParamsBodyIDPrefetchPreload, SettingEditParamsBodyIDProxyReadTimeout, SettingEditParamsBodyIDPseudoIPV4, SettingEditParamsBodyIDReplaceInsecureJS, SettingEditParamsBodyIDResponseBuffering, SettingEditParamsBodyIDRocketLoader, SettingEditParamsBodyIDAutomaticPlatformOptimization, SettingEditParamsBodyIDSecurityHeader, SettingEditParamsBodyIDSecurityLevel, SettingEditParamsBodyIDServerSideExclude, SettingEditParamsBodyIDSha1Support, SettingEditParamsBodyIDSortQueryStringForCache, SettingEditParamsBodyIDSSL, SettingEditParamsBodyIDSSLRecommender, SettingEditParamsBodyIDTLS1_2Only, SettingEditParamsBodyIDTLS1_3, SettingEditParamsBodyIDTLSClientAuth, SettingEditParamsBodyIDTrueClientIPHeader, SettingEditParamsBodyIDWAF, SettingEditParamsBodyIDWebP, SettingEditParamsBodyIDWebsockets:
		return true
	}
	return false
}

type SettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result SettingEditResponse             `json:"result"`
	JSON   settingEditResponseEnvelopeJSON `json:"-"`
}

// settingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingEditResponseEnvelope]
type settingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// 0-RTT session resumption enabled for this zone.
	Result SettingGetResponse             `json:"result"`
	JSON   settingGetResponseEnvelopeJSON `json:"-"`
}

// settingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingGetResponseEnvelope]
type settingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
