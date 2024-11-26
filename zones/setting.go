// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
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
	var env SettingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	var env SettingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
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

type AlwaysUseHTTPS struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID   AlwaysUseHTTPSID   `json:"id"`
	JSON alwaysUseHTTPSJSON `json:"-"`
}

// alwaysUseHTTPSJSON contains the JSON metadata for the struct [AlwaysUseHTTPS]
type alwaysUseHTTPSJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r alwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r AlwaysUseHTTPS) ImplementsPagerulesPageRuleAction() {}

// If enabled, any ` http://“ URL is converted to  `https://` through a 301
// redirect.
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

type AlwaysUseHTTPSParam struct {
	// If enabled, any ` http://“ URL is converted to  `https://` through a 301
	// redirect.
	ID param.Field[AlwaysUseHTTPSID] `json:"id"`
}

func (r AlwaysUseHTTPSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AlwaysUseHTTPSParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r AlwaysUseHTTPSParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r AlwaysUseHTTPSParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

type AutomaticHTTPSRewrites struct {
	// Turn on or off Automatic HTTPS Rewrites.
	ID AutomaticHTTPSRewritesID `json:"id"`
	// The status of Automatic HTTPS Rewrites.
	Value AutomaticHTTPSRewritesValue `json:"value"`
	JSON  automaticHTTPSRewritesJSON  `json:"-"`
}

// automaticHTTPSRewritesJSON contains the JSON metadata for the struct
// [AutomaticHTTPSRewrites]
type automaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r automaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

func (r AutomaticHTTPSRewrites) ImplementsPagerulesPageRuleAction() {}

// Turn on or off Automatic HTTPS Rewrites.
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

// The status of Automatic HTTPS Rewrites.
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

type AutomaticHTTPSRewritesParam struct {
	// Turn on or off Automatic HTTPS Rewrites.
	ID param.Field[AutomaticHTTPSRewritesID] `json:"id"`
	// The status of Automatic HTTPS Rewrites.
	Value param.Field[AutomaticHTTPSRewritesValue] `json:"value"`
}

func (r AutomaticHTTPSRewritesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AutomaticHTTPSRewritesParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r AutomaticHTTPSRewritesParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r AutomaticHTTPSRewritesParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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

type BrowserCacheTTL struct {
	// Control how long resources cached by client browsers remain valid.
	ID BrowserCacheTTLID `json:"id"`
	// The number of seconds to cache resources for. The API prohibits setting this to
	// 0 for non-Enterprise domains.
	Value int64               `json:"value"`
	JSON  browserCacheTTLJSON `json:"-"`
}

// browserCacheTTLJSON contains the JSON metadata for the struct [BrowserCacheTTL]
type browserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r BrowserCacheTTL) ImplementsPagerulesPageRuleAction() {}

// Control how long resources cached by client browsers remain valid.
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

type BrowserCacheTTLParam struct {
	// Control how long resources cached by client browsers remain valid.
	ID param.Field[BrowserCacheTTLID] `json:"id"`
	// The number of seconds to cache resources for. The API prohibits setting this to
	// 0 for non-Enterprise domains.
	Value param.Field[int64] `json:"value"`
}

func (r BrowserCacheTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BrowserCacheTTLParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r BrowserCacheTTLParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r BrowserCacheTTLParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

type BrowserCheck struct {
	// Inspect the visitor's browser for headers commonly associated with spammers and
	// certain bots.
	ID BrowserCheckID `json:"id"`
	// The status of Browser Integrity Check.
	Value BrowserCheckValue `json:"value"`
	JSON  browserCheckJSON  `json:"-"`
}

// browserCheckJSON contains the JSON metadata for the struct [BrowserCheck]
type browserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserCheckJSON) RawJSON() string {
	return r.raw
}

func (r BrowserCheck) ImplementsPagerulesPageRuleAction() {}

// Inspect the visitor's browser for headers commonly associated with spammers and
// certain bots.
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

// The status of Browser Integrity Check.
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

type BrowserCheckParam struct {
	// Inspect the visitor's browser for headers commonly associated with spammers and
	// certain bots.
	ID param.Field[BrowserCheckID] `json:"id"`
	// The status of Browser Integrity Check.
	Value param.Field[BrowserCheckValue] `json:"value"`
}

func (r BrowserCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BrowserCheckParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r BrowserCheckParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r BrowserCheckParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

type CacheLevel struct {
	// Apply custom caching based on the option selected.
	ID CacheLevelID `json:"id"`
	//   - `bypass`: Cloudflare does not cache.
	//   - `basic`: Delivers resources from cache when there is no query string.
	//   - `simplified`: Delivers the same resource to everyone independent of the query
	//     string.
	//   - `aggressive`: Caches all static content that has a query string.
	//   - `cache_everything`: Treats all content as static and caches all file types
	//     beyond the
	//     [Cloudflare default cached content](https://developers.cloudflare.com/cache/concepts/default-cache-behavior/#default-cached-file-extensions).
	Value CacheLevelValue `json:"value"`
	JSON  cacheLevelJSON  `json:"-"`
}

// cacheLevelJSON contains the JSON metadata for the struct [CacheLevel]
type cacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheLevelJSON) RawJSON() string {
	return r.raw
}

func (r CacheLevel) ImplementsPagerulesPageRuleAction() {}

// Apply custom caching based on the option selected.
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

//   - `bypass`: Cloudflare does not cache.
//   - `basic`: Delivers resources from cache when there is no query string.
//   - `simplified`: Delivers the same resource to everyone independent of the query
//     string.
//   - `aggressive`: Caches all static content that has a query string.
//   - `cache_everything`: Treats all content as static and caches all file types
//     beyond the
//     [Cloudflare default cached content](https://developers.cloudflare.com/cache/concepts/default-cache-behavior/#default-cached-file-extensions).
type CacheLevelValue string

const (
	CacheLevelValueBypass          CacheLevelValue = "bypass"
	CacheLevelValueBasic           CacheLevelValue = "basic"
	CacheLevelValueSimplified      CacheLevelValue = "simplified"
	CacheLevelValueAggressive      CacheLevelValue = "aggressive"
	CacheLevelValueCacheEverything CacheLevelValue = "cache_everything"
)

func (r CacheLevelValue) IsKnown() bool {
	switch r {
	case CacheLevelValueBypass, CacheLevelValueBasic, CacheLevelValueSimplified, CacheLevelValueAggressive, CacheLevelValueCacheEverything:
		return true
	}
	return false
}

type CacheLevelParam struct {
	// Apply custom caching based on the option selected.
	ID param.Field[CacheLevelID] `json:"id"`
	//   - `bypass`: Cloudflare does not cache.
	//   - `basic`: Delivers resources from cache when there is no query string.
	//   - `simplified`: Delivers the same resource to everyone independent of the query
	//     string.
	//   - `aggressive`: Caches all static content that has a query string.
	//   - `cache_everything`: Treats all content as static and caches all file types
	//     beyond the
	//     [Cloudflare default cached content](https://developers.cloudflare.com/cache/concepts/default-cache-behavior/#default-cached-file-extensions).
	Value param.Field[CacheLevelValue] `json:"value"`
}

func (r CacheLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CacheLevelParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r CacheLevelParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r CacheLevelParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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

type EmailObfuscation struct {
	// Turn on or off **Email Obfuscation**.
	ID EmailObfuscationID `json:"id"`
	// The status of Email Obfuscation.
	Value EmailObfuscationValue `json:"value"`
	JSON  emailObfuscationJSON  `json:"-"`
}

// emailObfuscationJSON contains the JSON metadata for the struct
// [EmailObfuscation]
type emailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r EmailObfuscation) ImplementsPagerulesPageRuleAction() {}

// Turn on or off **Email Obfuscation**.
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

// The status of Email Obfuscation.
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

type EmailObfuscationParam struct {
	// Turn on or off **Email Obfuscation**.
	ID param.Field[EmailObfuscationID] `json:"id"`
	// The status of Email Obfuscation.
	Value param.Field[EmailObfuscationValue] `json:"value"`
}

func (r EmailObfuscationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EmailObfuscationParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r EmailObfuscationParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r EmailObfuscationParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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

type IPGeolocation struct {
	// Cloudflare adds a CF-IPCountry HTTP header containing the country code that
	// corresponds to the visitor.
	ID IPGeolocationID `json:"id"`
	// The status of adding the IP Geolocation Header.
	Value IPGeolocationValue `json:"value"`
	JSON  ipGeolocationJSON  `json:"-"`
}

// ipGeolocationJSON contains the JSON metadata for the struct [IPGeolocation]
type ipGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipGeolocationJSON) RawJSON() string {
	return r.raw
}

func (r IPGeolocation) ImplementsPagerulesPageRuleAction() {}

// Cloudflare adds a CF-IPCountry HTTP header containing the country code that
// corresponds to the visitor.
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

// The status of adding the IP Geolocation Header.
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

type IPGeolocationParam struct {
	// Cloudflare adds a CF-IPCountry HTTP header containing the country code that
	// corresponds to the visitor.
	ID param.Field[IPGeolocationID] `json:"id"`
	// The status of adding the IP Geolocation Header.
	Value param.Field[IPGeolocationValue] `json:"value"`
}

func (r IPGeolocationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPGeolocationParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r IPGeolocationParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r IPGeolocationParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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

type Mirage struct {
	// Cloudflare Mirage reduces bandwidth used by images in mobile browsers. It can
	// accelerate loading of image-heavy websites on very slow mobile connections and
	// HTTP/1.
	ID MirageID `json:"id"`
	// The status of Mirage.
	Value MirageValue `json:"value"`
	JSON  mirageJSON  `json:"-"`
}

// mirageJSON contains the JSON metadata for the struct [Mirage]
type mirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Mirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mirageJSON) RawJSON() string {
	return r.raw
}

func (r Mirage) ImplementsPagerulesPageRuleAction() {}

// Cloudflare Mirage reduces bandwidth used by images in mobile browsers. It can
// accelerate loading of image-heavy websites on very slow mobile connections and
// HTTP/1.
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

// The status of Mirage.
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

type MirageParam struct {
	// Cloudflare Mirage reduces bandwidth used by images in mobile browsers. It can
	// accelerate loading of image-heavy websites on very slow mobile connections and
	// HTTP/1.
	ID param.Field[MirageID] `json:"id"`
	// The status of Mirage.
	Value param.Field[MirageValue] `json:"value"`
}

func (r MirageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MirageParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r MirageParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r MirageParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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

type OpportunisticEncryption struct {
	// Opportunistic Encryption allows browsers to access HTTP URIs over an encrypted
	// TLS channel. It's not a substitute for HTTPS, but provides additional security
	// for otherwise vulnerable requests.
	ID OpportunisticEncryptionID `json:"id"`
	// The status of Opportunistic Encryption.
	Value OpportunisticEncryptionValue `json:"value"`
	JSON  opportunisticEncryptionJSON  `json:"-"`
}

// opportunisticEncryptionJSON contains the JSON metadata for the struct
// [OpportunisticEncryption]
type opportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r opportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r OpportunisticEncryption) ImplementsPagerulesPageRuleAction() {}

// Opportunistic Encryption allows browsers to access HTTP URIs over an encrypted
// TLS channel. It's not a substitute for HTTPS, but provides additional security
// for otherwise vulnerable requests.
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

// The status of Opportunistic Encryption.
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

type OpportunisticEncryptionParam struct {
	// Opportunistic Encryption allows browsers to access HTTP URIs over an encrypted
	// TLS channel. It's not a substitute for HTTPS, but provides additional security
	// for otherwise vulnerable requests.
	ID param.Field[OpportunisticEncryptionID] `json:"id"`
	// The status of Opportunistic Encryption.
	Value param.Field[OpportunisticEncryptionValue] `json:"value"`
}

func (r OpportunisticEncryptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OpportunisticEncryptionParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r OpportunisticEncryptionParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r OpportunisticEncryptionParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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

type OriginErrorPagePassThru struct {
	// Turn on or off Cloudflare error pages generated from issues sent from the origin
	// server. If enabled, this setting triggers error pages issued by the origin.
	ID OriginErrorPagePassThruID `json:"id"`
	// The status of Origin Error Page Passthru.
	Value OriginErrorPagePassThruValue `json:"value"`
	JSON  originErrorPagePassThruJSON  `json:"-"`
}

// originErrorPagePassThruJSON contains the JSON metadata for the struct
// [OriginErrorPagePassThru]
type originErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originErrorPagePassThruJSON) RawJSON() string {
	return r.raw
}

func (r OriginErrorPagePassThru) ImplementsPagerulesPageRuleAction() {}

// Turn on or off Cloudflare error pages generated from issues sent from the origin
// server. If enabled, this setting triggers error pages issued by the origin.
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

// The status of Origin Error Page Passthru.
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

type OriginErrorPagePassThruParam struct {
	// Turn on or off Cloudflare error pages generated from issues sent from the origin
	// server. If enabled, this setting triggers error pages issued by the origin.
	ID param.Field[OriginErrorPagePassThruID] `json:"id"`
	// The status of Origin Error Page Passthru.
	Value param.Field[OriginErrorPagePassThruValue] `json:"value"`
}

func (r OriginErrorPagePassThruParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OriginErrorPagePassThruParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r OriginErrorPagePassThruParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r OriginErrorPagePassThruParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

type Polish struct {
	// Apply options from the Polish feature of the Cloudflare Speed app.
	ID PolishID `json:"id"`
	// The level of Polish you want applied to your origin.
	Value PolishValue `json:"value"`
	JSON  polishJSON  `json:"-"`
}

// polishJSON contains the JSON metadata for the struct [Polish]
type polishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Polish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r polishJSON) RawJSON() string {
	return r.raw
}

func (r Polish) ImplementsPagerulesPageRuleAction() {}

// Apply options from the Polish feature of the Cloudflare Speed app.
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

// The level of Polish you want applied to your origin.
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

type PolishParam struct {
	// Apply options from the Polish feature of the Cloudflare Speed app.
	ID param.Field[PolishID] `json:"id"`
	// The level of Polish you want applied to your origin.
	Value param.Field[PolishValue] `json:"value"`
}

func (r PolishParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PolishParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r PolishParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r PolishParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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

type ResponseBuffering struct {
	// Turn on or off whether Cloudflare should wait for an entire file from the origin
	// server before forwarding it to the site visitor. By default, Cloudflare sends
	// packets to the client as they arrive from the origin server.
	ID ResponseBufferingID `json:"id"`
	// The status of Response Buffering
	Value ResponseBufferingValue `json:"value"`
	JSON  responseBufferingJSON  `json:"-"`
}

// responseBufferingJSON contains the JSON metadata for the struct
// [ResponseBuffering]
type responseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseBufferingJSON) RawJSON() string {
	return r.raw
}

func (r ResponseBuffering) ImplementsPagerulesPageRuleAction() {}

// Turn on or off whether Cloudflare should wait for an entire file from the origin
// server before forwarding it to the site visitor. By default, Cloudflare sends
// packets to the client as they arrive from the origin server.
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

// The status of Response Buffering
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

type ResponseBufferingParam struct {
	// Turn on or off whether Cloudflare should wait for an entire file from the origin
	// server before forwarding it to the site visitor. By default, Cloudflare sends
	// packets to the client as they arrive from the origin server.
	ID param.Field[ResponseBufferingID] `json:"id"`
	// The status of Response Buffering
	Value param.Field[ResponseBufferingValue] `json:"value"`
}

func (r ResponseBufferingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ResponseBufferingParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r ResponseBufferingParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r ResponseBufferingParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

type RocketLoader struct {
	// Turn on or off Rocket Loader in the Cloudflare Speed app.
	ID RocketLoaderID `json:"id"`
	// The status of Rocket Loader
	Value RocketLoaderValue `json:"value"`
	JSON  rocketLoaderJSON  `json:"-"`
}

// rocketLoaderJSON contains the JSON metadata for the struct [RocketLoader]
type rocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rocketLoaderJSON) RawJSON() string {
	return r.raw
}

func (r RocketLoader) ImplementsPagerulesPageRuleAction() {}

// Turn on or off Rocket Loader in the Cloudflare Speed app.
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

// The status of Rocket Loader
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

type RocketLoaderParam struct {
	// Turn on or off Rocket Loader in the Cloudflare Speed app.
	ID param.Field[RocketLoaderID] `json:"id"`
	// The status of Rocket Loader
	Value param.Field[RocketLoaderValue] `json:"value"`
}

func (r RocketLoaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RocketLoaderParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r RocketLoaderParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r RocketLoaderParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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
	Nosniff bool `json:"nosniff"`
	// Enable automatic preload of the HSTS configuration.
	Preload bool                                            `json:"preload"`
	JSON    securityHeadersValueStrictTransportSecurityJSON `json:"-"`
}

// securityHeadersValueStrictTransportSecurityJSON contains the JSON metadata for
// the struct [SecurityHeadersValueStrictTransportSecurity]
type securityHeadersValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	Preload           apijson.Field
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
	// Enable automatic preload of the HSTS configuration.
	Preload param.Field[bool] `json:"preload"`
}

func (r SecurityHeadersValueStrictTransportSecurityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecurityLevel struct {
	// Control options for the **Security Level** feature from the **Security** app.
	ID    SecurityLevelID    `json:"id"`
	Value SecurityLevelValue `json:"value"`
	JSON  securityLevelJSON  `json:"-"`
}

// securityLevelJSON contains the JSON metadata for the struct [SecurityLevel]
type securityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r securityLevelJSON) RawJSON() string {
	return r.raw
}

func (r SecurityLevel) ImplementsPagerulesPageRuleAction() {}

// Control options for the **Security Level** feature from the **Security** app.
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

type SecurityLevelParam struct {
	// Control options for the **Security Level** feature from the **Security** app.
	ID    param.Field[SecurityLevelID]    `json:"id"`
	Value param.Field[SecurityLevelValue] `json:"value"`
}

func (r SecurityLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SecurityLevelParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r SecurityLevelParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r SecurityLevelParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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

type SortQueryStringForCache struct {
	// Turn on or off the reordering of query strings. When query strings have the same
	// structure, caching improves.
	ID SortQueryStringForCacheID `json:"id"`
	// The status of Query String Sort
	Value SortQueryStringForCacheValue `json:"value"`
	JSON  sortQueryStringForCacheJSON  `json:"-"`
}

// sortQueryStringForCacheJSON contains the JSON metadata for the struct
// [SortQueryStringForCache]
type sortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
}

func (r SortQueryStringForCache) ImplementsPagerulesPageRuleAction() {}

// Turn on or off the reordering of query strings. When query strings have the same
// structure, caching improves.
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

// The status of Query String Sort
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

type SortQueryStringForCacheParam struct {
	// Turn on or off the reordering of query strings. When query strings have the same
	// structure, caching improves.
	ID param.Field[SortQueryStringForCacheID] `json:"id"`
	// The status of Query String Sort
	Value param.Field[SortQueryStringForCacheValue] `json:"value"`
}

func (r SortQueryStringForCacheParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SortQueryStringForCacheParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r SortQueryStringForCacheParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r SortQueryStringForCacheParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

type SSL struct {
	// Control options for the SSL feature of the Edge Certificates tab in the
	// Cloudflare SSL/TLS app.
	ID SSLID `json:"id"`
	// The encryption mode that Cloudflare uses to connect to your origin server.
	Value SSLValue `json:"value"`
	JSON  sslJSON  `json:"-"`
}

// sslJSON contains the JSON metadata for the struct [SSL]
type sslJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslJSON) RawJSON() string {
	return r.raw
}

func (r SSL) ImplementsPagerulesPageRuleAction() {}

// Control options for the SSL feature of the Edge Certificates tab in the
// Cloudflare SSL/TLS app.
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

// The encryption mode that Cloudflare uses to connect to your origin server.
type SSLValue string

const (
	SSLValueOff        SSLValue = "off"
	SSLValueFlexible   SSLValue = "flexible"
	SSLValueFull       SSLValue = "full"
	SSLValueStrict     SSLValue = "strict"
	SSLValueOriginPull SSLValue = "origin_pull"
)

func (r SSLValue) IsKnown() bool {
	switch r {
	case SSLValueOff, SSLValueFlexible, SSLValueFull, SSLValueStrict, SSLValueOriginPull:
		return true
	}
	return false
}

type SSLParam struct {
	// Control options for the SSL feature of the Edge Certificates tab in the
	// Cloudflare SSL/TLS app.
	ID param.Field[SSLID] `json:"id"`
	// The encryption mode that Cloudflare uses to connect to your origin server.
	Value param.Field[SSLValue] `json:"value"`
}

func (r SSLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SSLParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r SSLParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r SSLParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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

type TrueClientIPHeader struct {
	// Turn on or off the True-Client-IP Header feature of the Cloudflare Network app.
	ID TrueClientIPHeaderID `json:"id"`
	// The status of True Client IP Header.
	Value TrueClientIPHeaderValue `json:"value"`
	JSON  trueClientIPHeaderJSON  `json:"-"`
}

// trueClientIPHeaderJSON contains the JSON metadata for the struct
// [TrueClientIPHeader]
type trueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

func (r TrueClientIPHeader) ImplementsPagerulesPageRuleAction() {}

// Turn on or off the True-Client-IP Header feature of the Cloudflare Network app.
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

// The status of True Client IP Header.
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

type TrueClientIPHeaderParam struct {
	// Turn on or off the True-Client-IP Header feature of the Cloudflare Network app.
	ID param.Field[TrueClientIPHeaderID] `json:"id"`
	// The status of True Client IP Header.
	Value param.Field[TrueClientIPHeaderValue] `json:"value"`
}

func (r TrueClientIPHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TrueClientIPHeaderParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r TrueClientIPHeaderParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r TrueClientIPHeaderParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

type WAF struct {
	// Turn on or off
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	// You cannot enable or disable individual WAF managed rules via Page Rules.
	ID WAFID `json:"id"`
	// The status of WAF managed rules (previous version).
	Value WAFValue `json:"value"`
	JSON  wafJSON  `json:"-"`
}

// wafJSON contains the JSON metadata for the struct [WAF]
type wafJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafJSON) RawJSON() string {
	return r.raw
}

func (r WAF) ImplementsPagerulesPageRuleAction() {}

// Turn on or off
// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
// You cannot enable or disable individual WAF managed rules via Page Rules.
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

// The status of WAF managed rules (previous version).
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

type WAFParam struct {
	// Turn on or off
	// [WAF managed rules (previous version, deprecated)](https://developers.cloudflare.com/waf/reference/legacy/old-waf-managed-rules/).
	// You cannot enable or disable individual WAF managed rules via Page Rules.
	ID param.Field[WAFID] `json:"id"`
	// The status of WAF managed rules (previous version).
	Value param.Field[WAFValue] `json:"value"`
}

func (r WAFParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WAFParam) ImplementsPagerulesPageruleNewParamsActionUnion() {}

func (r WAFParam) ImplementsPagerulesPageruleUpdateParamsActionUnion() {}

func (r WAFParam) ImplementsPagerulesPageruleEditParamsActionUnion() {}

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
	// ID of the zone setting.
	ID SettingEditResponseID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseEditable `json:"editable"`
	// ssl-recommender enrollment setting.
	Enabled bool `json:"enabled"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// This field can have the runtime type of [ZeroRTTValue], [AdvancedDDoSValue],
	// [AlwaysOnlineValue], [SettingEditResponseZonesSchemasAlwaysUseHTTPSValue],
	// [SettingEditResponseZonesSchemasAutomaticHTTPSRewritesValue], [BrotliValue],
	// [SettingEditResponseZonesSchemasBrowserCacheTTLValue],
	// [SettingEditResponseZonesSchemasBrowserCheckValue],
	// [SettingEditResponseZonesSchemasCacheLevelValue], [ChallengeTTLValue],
	// [[]string], [SettingEditResponseZonesCNAMEFlatteningValue],
	// [DevelopmentModeValue], [EarlyHintsValue],
	// [SettingEditResponseZonesSchemasEdgeCacheTTLValue],
	// [SettingEditResponseZonesSchemasEmailObfuscationValue], [H2PrioritizationValue],
	// [HotlinkProtectionValue], [HTTP2Value], [HTTP3Value], [ImageResizingValue],
	// [SettingEditResponseZonesSchemasIPGeolocationValue], [IPV6Value],
	// [SettingEditResponseZonesMaxUploadValue], [MinTLSVersionValue],
	// [SettingEditResponseZonesSchemasMirageValue], [NELValue],
	// [SettingEditResponseZonesSchemasOpportunisticEncryptionValue],
	// [OpportunisticOnionValue], [OrangeToOrangeValue],
	// [SettingEditResponseZonesSchemasOriginErrorPagePassThruValue],
	// [SettingEditResponseZonesSchemasPolishValue], [PrefetchPreloadValue], [float64],
	// [PseudoIPV4Value], [SettingEditResponseZonesReplaceInsecureJSValue],
	// [SettingEditResponseZonesSchemasResponseBufferingValue],
	// [SettingEditResponseZonesSchemasRocketLoaderValue],
	// [AutomaticPlatformOptimization], [SecurityHeadersValue],
	// [SettingEditResponseZonesSchemasSecurityLevelValue], [ServerSideExcludesValue],
	// [SettingEditResponseZonesSha1SupportValue],
	// [SettingEditResponseZonesSchemasSortQueryStringForCacheValue],
	// [SettingEditResponseZonesSchemasSSLValue],
	// [SettingEditResponseZonesTLS1_2OnlyValue], [TLS1_3Value], [TLSClientAuthValue],
	// [SettingEditResponseZonesSchemasTrueClientIPHeaderValue],
	// [SettingEditResponseZonesSchemasWAFValue], [WebPValue], [WebsocketValue].
	Value interface{}             `json:"value"`
	JSON  settingEditResponseJSON `json:"-"`
	union SettingEditResponseUnion
}

// settingEditResponseJSON contains the JSON metadata for the struct
// [SettingEditResponse]
type settingEditResponseJSON struct {
	ID            apijson.Field
	Editable      apijson.Field
	Enabled       apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r settingEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	*r = SettingEditResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SettingEditResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [zones.ZeroRTT], [zones.AdvancedDDoS],
// [zones.AlwaysOnline], [zones.SettingEditResponseZonesSchemasAlwaysUseHTTPS],
// [zones.SettingEditResponseZonesSchemasAutomaticHTTPSRewrites], [zones.Brotli],
// [zones.SettingEditResponseZonesSchemasBrowserCacheTTL],
// [zones.SettingEditResponseZonesSchemasBrowserCheck],
// [zones.SettingEditResponseZonesSchemasCacheLevel], [zones.ChallengeTTL],
// [zones.Ciphers], [zones.SettingEditResponseZonesCNAMEFlattening],
// [zones.DevelopmentMode], [zones.EarlyHints],
// [zones.SettingEditResponseZonesSchemasEdgeCacheTTL],
// [zones.SettingEditResponseZonesSchemasEmailObfuscation],
// [zones.H2Prioritization], [zones.HotlinkProtection], [zones.HTTP2],
// [zones.HTTP3], [zones.ImageResizing],
// [zones.SettingEditResponseZonesSchemasIPGeolocation], [zones.IPV6],
// [zones.SettingEditResponseZonesMaxUpload], [zones.MinTLSVersion],
// [zones.SettingEditResponseZonesSchemasMirage], [zones.NEL],
// [zones.SettingEditResponseZonesSchemasOpportunisticEncryption],
// [zones.OpportunisticOnion], [zones.OrangeToOrange],
// [zones.SettingEditResponseZonesSchemasOriginErrorPagePassThru],
// [zones.SettingEditResponseZonesSchemasPolish], [zones.PrefetchPreload],
// [zones.ProxyReadTimeout], [zones.PseudoIPV4],
// [zones.SettingEditResponseZonesReplaceInsecureJS],
// [zones.SettingEditResponseZonesSchemasResponseBuffering],
// [zones.SettingEditResponseZonesSchemasRocketLoader],
// [zones.SettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.SecurityHeaders], [zones.SettingEditResponseZonesSchemasSecurityLevel],
// [zones.ServerSideExcludes], [zones.SettingEditResponseZonesSha1Support],
// [zones.SettingEditResponseZonesSchemasSortQueryStringForCache],
// [zones.SettingEditResponseZonesSchemasSSL], [zones.SSLRecommender],
// [zones.SettingEditResponseZonesTLS1_2Only], [zones.TLS1_3],
// [zones.TLSClientAuth],
// [zones.SettingEditResponseZonesSchemasTrueClientIPHeader],
// [zones.SettingEditResponseZonesSchemasWAF], [zones.WebP], [zones.Websocket].
func (r SettingEditResponse) AsUnion() SettingEditResponseUnion {
	return r.union
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [zones.ZeroRTT], [zones.AdvancedDDoS], [zones.AlwaysOnline],
// [zones.SettingEditResponseZonesSchemasAlwaysUseHTTPS],
// [zones.SettingEditResponseZonesSchemasAutomaticHTTPSRewrites], [zones.Brotli],
// [zones.SettingEditResponseZonesSchemasBrowserCacheTTL],
// [zones.SettingEditResponseZonesSchemasBrowserCheck],
// [zones.SettingEditResponseZonesSchemasCacheLevel], [zones.ChallengeTTL],
// [zones.Ciphers], [zones.SettingEditResponseZonesCNAMEFlattening],
// [zones.DevelopmentMode], [zones.EarlyHints],
// [zones.SettingEditResponseZonesSchemasEdgeCacheTTL],
// [zones.SettingEditResponseZonesSchemasEmailObfuscation],
// [zones.H2Prioritization], [zones.HotlinkProtection], [zones.HTTP2],
// [zones.HTTP3], [zones.ImageResizing],
// [zones.SettingEditResponseZonesSchemasIPGeolocation], [zones.IPV6],
// [zones.SettingEditResponseZonesMaxUpload], [zones.MinTLSVersion],
// [zones.SettingEditResponseZonesSchemasMirage], [zones.NEL],
// [zones.SettingEditResponseZonesSchemasOpportunisticEncryption],
// [zones.OpportunisticOnion], [zones.OrangeToOrange],
// [zones.SettingEditResponseZonesSchemasOriginErrorPagePassThru],
// [zones.SettingEditResponseZonesSchemasPolish], [zones.PrefetchPreload],
// [zones.ProxyReadTimeout], [zones.PseudoIPV4],
// [zones.SettingEditResponseZonesReplaceInsecureJS],
// [zones.SettingEditResponseZonesSchemasResponseBuffering],
// [zones.SettingEditResponseZonesSchemasRocketLoader],
// [zones.SettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.SecurityHeaders], [zones.SettingEditResponseZonesSchemasSecurityLevel],
// [zones.ServerSideExcludes], [zones.SettingEditResponseZonesSha1Support],
// [zones.SettingEditResponseZonesSchemasSortQueryStringForCache],
// [zones.SettingEditResponseZonesSchemasSSL], [zones.SSLRecommender],
// [zones.SettingEditResponseZonesTLS1_2Only], [zones.TLS1_3],
// [zones.TLSClientAuth],
// [zones.SettingEditResponseZonesSchemasTrueClientIPHeader],
// [zones.SettingEditResponseZonesSchemasWAF], [zones.WebP] or [zones.Websocket].
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
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Brotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasBrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasCacheLevel{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasEmailObfuscation{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasIPGeolocation{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasOpportunisticEncryption{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasPolish{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasResponseBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasRocketLoader{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasSecurityLevel{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasSSL{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasWAF{}),
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

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingEditResponseZonesSchemasAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasAlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasAlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasAlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasAlwaysUseHTTPSJSON `json:"-"`
}

// settingEditResponseZonesSchemasAlwaysUseHTTPSJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesSchemasAlwaysUseHTTPS]
type settingEditResponseZonesSchemasAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasAlwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasAlwaysUseHTTPS) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasAlwaysUseHTTPSID string

const (
	SettingEditResponseZonesSchemasAlwaysUseHTTPSIDAlwaysUseHTTPS SettingEditResponseZonesSchemasAlwaysUseHTTPSID = "always_use_https"
)

func (r SettingEditResponseZonesSchemasAlwaysUseHTTPSID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasAlwaysUseHTTPSIDAlwaysUseHTTPS:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasAlwaysUseHTTPSValue string

const (
	SettingEditResponseZonesSchemasAlwaysUseHTTPSValueOn  SettingEditResponseZonesSchemasAlwaysUseHTTPSValue = "on"
	SettingEditResponseZonesSchemasAlwaysUseHTTPSValueOff SettingEditResponseZonesSchemasAlwaysUseHTTPSValue = "off"
)

func (r SettingEditResponseZonesSchemasAlwaysUseHTTPSValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasAlwaysUseHTTPSValueOn, SettingEditResponseZonesSchemasAlwaysUseHTTPSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasAlwaysUseHTTPSEditable bool

const (
	SettingEditResponseZonesSchemasAlwaysUseHTTPSEditableTrue  SettingEditResponseZonesSchemasAlwaysUseHTTPSEditable = true
	SettingEditResponseZonesSchemasAlwaysUseHTTPSEditableFalse SettingEditResponseZonesSchemasAlwaysUseHTTPSEditable = false
)

func (r SettingEditResponseZonesSchemasAlwaysUseHTTPSEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasAlwaysUseHTTPSEditableTrue, SettingEditResponseZonesSchemasAlwaysUseHTTPSEditableFalse:
		return true
	}
	return false
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingEditResponseZonesSchemasAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasAutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasAutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasAutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasAutomaticHTTPSRewritesJSON `json:"-"`
}

// settingEditResponseZonesSchemasAutomaticHTTPSRewritesJSON contains the JSON
// metadata for the struct [SettingEditResponseZonesSchemasAutomaticHTTPSRewrites]
type settingEditResponseZonesSchemasAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasAutomaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasAutomaticHTTPSRewrites) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasAutomaticHTTPSRewritesID string

const (
	SettingEditResponseZonesSchemasAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites SettingEditResponseZonesSchemasAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

func (r SettingEditResponseZonesSchemasAutomaticHTTPSRewritesID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasAutomaticHTTPSRewritesValue string

const (
	SettingEditResponseZonesSchemasAutomaticHTTPSRewritesValueOn  SettingEditResponseZonesSchemasAutomaticHTTPSRewritesValue = "on"
	SettingEditResponseZonesSchemasAutomaticHTTPSRewritesValueOff SettingEditResponseZonesSchemasAutomaticHTTPSRewritesValue = "off"
)

func (r SettingEditResponseZonesSchemasAutomaticHTTPSRewritesValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasAutomaticHTTPSRewritesValueOn, SettingEditResponseZonesSchemasAutomaticHTTPSRewritesValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasAutomaticHTTPSRewritesEditable bool

const (
	SettingEditResponseZonesSchemasAutomaticHTTPSRewritesEditableTrue  SettingEditResponseZonesSchemasAutomaticHTTPSRewritesEditable = true
	SettingEditResponseZonesSchemasAutomaticHTTPSRewritesEditableFalse SettingEditResponseZonesSchemasAutomaticHTTPSRewritesEditable = false
)

func (r SettingEditResponseZonesSchemasAutomaticHTTPSRewritesEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasAutomaticHTTPSRewritesEditableTrue, SettingEditResponseZonesSchemasAutomaticHTTPSRewritesEditableFalse:
		return true
	}
	return false
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingEditResponseZonesSchemasBrowserCacheTTL struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasBrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasBrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasBrowserCacheTTLJSON `json:"-"`
}

// settingEditResponseZonesSchemasBrowserCacheTTLJSON contains the JSON metadata
// for the struct [SettingEditResponseZonesSchemasBrowserCacheTTL]
type settingEditResponseZonesSchemasBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasBrowserCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasBrowserCacheTTL) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasBrowserCacheTTLID string

const (
	SettingEditResponseZonesSchemasBrowserCacheTTLIDBrowserCacheTTL SettingEditResponseZonesSchemasBrowserCacheTTLID = "browser_cache_ttl"
)

func (r SettingEditResponseZonesSchemasBrowserCacheTTLID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasBrowserCacheTTLIDBrowserCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasBrowserCacheTTLValue float64

const (
	SettingEditResponseZonesSchemasBrowserCacheTTLValue0        SettingEditResponseZonesSchemasBrowserCacheTTLValue = 0
	SettingEditResponseZonesSchemasBrowserCacheTTLValue30       SettingEditResponseZonesSchemasBrowserCacheTTLValue = 30
	SettingEditResponseZonesSchemasBrowserCacheTTLValue60       SettingEditResponseZonesSchemasBrowserCacheTTLValue = 60
	SettingEditResponseZonesSchemasBrowserCacheTTLValue120      SettingEditResponseZonesSchemasBrowserCacheTTLValue = 120
	SettingEditResponseZonesSchemasBrowserCacheTTLValue300      SettingEditResponseZonesSchemasBrowserCacheTTLValue = 300
	SettingEditResponseZonesSchemasBrowserCacheTTLValue1200     SettingEditResponseZonesSchemasBrowserCacheTTLValue = 1200
	SettingEditResponseZonesSchemasBrowserCacheTTLValue1800     SettingEditResponseZonesSchemasBrowserCacheTTLValue = 1800
	SettingEditResponseZonesSchemasBrowserCacheTTLValue3600     SettingEditResponseZonesSchemasBrowserCacheTTLValue = 3600
	SettingEditResponseZonesSchemasBrowserCacheTTLValue7200     SettingEditResponseZonesSchemasBrowserCacheTTLValue = 7200
	SettingEditResponseZonesSchemasBrowserCacheTTLValue10800    SettingEditResponseZonesSchemasBrowserCacheTTLValue = 10800
	SettingEditResponseZonesSchemasBrowserCacheTTLValue14400    SettingEditResponseZonesSchemasBrowserCacheTTLValue = 14400
	SettingEditResponseZonesSchemasBrowserCacheTTLValue18000    SettingEditResponseZonesSchemasBrowserCacheTTLValue = 18000
	SettingEditResponseZonesSchemasBrowserCacheTTLValue28800    SettingEditResponseZonesSchemasBrowserCacheTTLValue = 28800
	SettingEditResponseZonesSchemasBrowserCacheTTLValue43200    SettingEditResponseZonesSchemasBrowserCacheTTLValue = 43200
	SettingEditResponseZonesSchemasBrowserCacheTTLValue57600    SettingEditResponseZonesSchemasBrowserCacheTTLValue = 57600
	SettingEditResponseZonesSchemasBrowserCacheTTLValue72000    SettingEditResponseZonesSchemasBrowserCacheTTLValue = 72000
	SettingEditResponseZonesSchemasBrowserCacheTTLValue86400    SettingEditResponseZonesSchemasBrowserCacheTTLValue = 86400
	SettingEditResponseZonesSchemasBrowserCacheTTLValue172800   SettingEditResponseZonesSchemasBrowserCacheTTLValue = 172800
	SettingEditResponseZonesSchemasBrowserCacheTTLValue259200   SettingEditResponseZonesSchemasBrowserCacheTTLValue = 259200
	SettingEditResponseZonesSchemasBrowserCacheTTLValue345600   SettingEditResponseZonesSchemasBrowserCacheTTLValue = 345600
	SettingEditResponseZonesSchemasBrowserCacheTTLValue432000   SettingEditResponseZonesSchemasBrowserCacheTTLValue = 432000
	SettingEditResponseZonesSchemasBrowserCacheTTLValue691200   SettingEditResponseZonesSchemasBrowserCacheTTLValue = 691200
	SettingEditResponseZonesSchemasBrowserCacheTTLValue1382400  SettingEditResponseZonesSchemasBrowserCacheTTLValue = 1382400
	SettingEditResponseZonesSchemasBrowserCacheTTLValue2073600  SettingEditResponseZonesSchemasBrowserCacheTTLValue = 2073600
	SettingEditResponseZonesSchemasBrowserCacheTTLValue2678400  SettingEditResponseZonesSchemasBrowserCacheTTLValue = 2678400
	SettingEditResponseZonesSchemasBrowserCacheTTLValue5356800  SettingEditResponseZonesSchemasBrowserCacheTTLValue = 5356800
	SettingEditResponseZonesSchemasBrowserCacheTTLValue16070400 SettingEditResponseZonesSchemasBrowserCacheTTLValue = 16070400
	SettingEditResponseZonesSchemasBrowserCacheTTLValue31536000 SettingEditResponseZonesSchemasBrowserCacheTTLValue = 31536000
)

func (r SettingEditResponseZonesSchemasBrowserCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasBrowserCacheTTLValue0, SettingEditResponseZonesSchemasBrowserCacheTTLValue30, SettingEditResponseZonesSchemasBrowserCacheTTLValue60, SettingEditResponseZonesSchemasBrowserCacheTTLValue120, SettingEditResponseZonesSchemasBrowserCacheTTLValue300, SettingEditResponseZonesSchemasBrowserCacheTTLValue1200, SettingEditResponseZonesSchemasBrowserCacheTTLValue1800, SettingEditResponseZonesSchemasBrowserCacheTTLValue3600, SettingEditResponseZonesSchemasBrowserCacheTTLValue7200, SettingEditResponseZonesSchemasBrowserCacheTTLValue10800, SettingEditResponseZonesSchemasBrowserCacheTTLValue14400, SettingEditResponseZonesSchemasBrowserCacheTTLValue18000, SettingEditResponseZonesSchemasBrowserCacheTTLValue28800, SettingEditResponseZonesSchemasBrowserCacheTTLValue43200, SettingEditResponseZonesSchemasBrowserCacheTTLValue57600, SettingEditResponseZonesSchemasBrowserCacheTTLValue72000, SettingEditResponseZonesSchemasBrowserCacheTTLValue86400, SettingEditResponseZonesSchemasBrowserCacheTTLValue172800, SettingEditResponseZonesSchemasBrowserCacheTTLValue259200, SettingEditResponseZonesSchemasBrowserCacheTTLValue345600, SettingEditResponseZonesSchemasBrowserCacheTTLValue432000, SettingEditResponseZonesSchemasBrowserCacheTTLValue691200, SettingEditResponseZonesSchemasBrowserCacheTTLValue1382400, SettingEditResponseZonesSchemasBrowserCacheTTLValue2073600, SettingEditResponseZonesSchemasBrowserCacheTTLValue2678400, SettingEditResponseZonesSchemasBrowserCacheTTLValue5356800, SettingEditResponseZonesSchemasBrowserCacheTTLValue16070400, SettingEditResponseZonesSchemasBrowserCacheTTLValue31536000:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasBrowserCacheTTLEditable bool

const (
	SettingEditResponseZonesSchemasBrowserCacheTTLEditableTrue  SettingEditResponseZonesSchemasBrowserCacheTTLEditable = true
	SettingEditResponseZonesSchemasBrowserCacheTTLEditableFalse SettingEditResponseZonesSchemasBrowserCacheTTLEditable = false
)

func (r SettingEditResponseZonesSchemasBrowserCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasBrowserCacheTTLEditableTrue, SettingEditResponseZonesSchemasBrowserCacheTTLEditableFalse:
		return true
	}
	return false
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type SettingEditResponseZonesSchemasBrowserCheck struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasBrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasBrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasBrowserCheckJSON `json:"-"`
}

// settingEditResponseZonesSchemasBrowserCheckJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesSchemasBrowserCheck]
type settingEditResponseZonesSchemasBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasBrowserCheckJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasBrowserCheck) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasBrowserCheckID string

const (
	SettingEditResponseZonesSchemasBrowserCheckIDBrowserCheck SettingEditResponseZonesSchemasBrowserCheckID = "browser_check"
)

func (r SettingEditResponseZonesSchemasBrowserCheckID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasBrowserCheckIDBrowserCheck:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasBrowserCheckValue string

const (
	SettingEditResponseZonesSchemasBrowserCheckValueOn  SettingEditResponseZonesSchemasBrowserCheckValue = "on"
	SettingEditResponseZonesSchemasBrowserCheckValueOff SettingEditResponseZonesSchemasBrowserCheckValue = "off"
)

func (r SettingEditResponseZonesSchemasBrowserCheckValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasBrowserCheckValueOn, SettingEditResponseZonesSchemasBrowserCheckValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasBrowserCheckEditable bool

const (
	SettingEditResponseZonesSchemasBrowserCheckEditableTrue  SettingEditResponseZonesSchemasBrowserCheckEditable = true
	SettingEditResponseZonesSchemasBrowserCheckEditableFalse SettingEditResponseZonesSchemasBrowserCheckEditable = false
)

func (r SettingEditResponseZonesSchemasBrowserCheckEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasBrowserCheckEditableTrue, SettingEditResponseZonesSchemasBrowserCheckEditableFalse:
		return true
	}
	return false
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type SettingEditResponseZonesSchemasCacheLevel struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasCacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasCacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasCacheLevelJSON `json:"-"`
}

// settingEditResponseZonesSchemasCacheLevelJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesSchemasCacheLevel]
type settingEditResponseZonesSchemasCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasCacheLevelJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasCacheLevel) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasCacheLevelID string

const (
	SettingEditResponseZonesSchemasCacheLevelIDCacheLevel SettingEditResponseZonesSchemasCacheLevelID = "cache_level"
)

func (r SettingEditResponseZonesSchemasCacheLevelID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasCacheLevelIDCacheLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasCacheLevelValue string

const (
	SettingEditResponseZonesSchemasCacheLevelValueAggressive SettingEditResponseZonesSchemasCacheLevelValue = "aggressive"
	SettingEditResponseZonesSchemasCacheLevelValueBasic      SettingEditResponseZonesSchemasCacheLevelValue = "basic"
	SettingEditResponseZonesSchemasCacheLevelValueSimplified SettingEditResponseZonesSchemasCacheLevelValue = "simplified"
)

func (r SettingEditResponseZonesSchemasCacheLevelValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasCacheLevelValueAggressive, SettingEditResponseZonesSchemasCacheLevelValueBasic, SettingEditResponseZonesSchemasCacheLevelValueSimplified:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasCacheLevelEditable bool

const (
	SettingEditResponseZonesSchemasCacheLevelEditableTrue  SettingEditResponseZonesSchemasCacheLevelEditable = true
	SettingEditResponseZonesSchemasCacheLevelEditableFalse SettingEditResponseZonesSchemasCacheLevelEditable = false
)

func (r SettingEditResponseZonesSchemasCacheLevelEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasCacheLevelEditableTrue, SettingEditResponseZonesSchemasCacheLevelEditableFalse:
		return true
	}
	return false
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
type SettingEditResponseZonesSchemasEdgeCacheTTL struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasEdgeCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasEdgeCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasEdgeCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasEdgeCacheTTLJSON `json:"-"`
}

// settingEditResponseZonesSchemasEdgeCacheTTLJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesSchemasEdgeCacheTTL]
type settingEditResponseZonesSchemasEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasEdgeCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasEdgeCacheTTL) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasEdgeCacheTTLID string

const (
	SettingEditResponseZonesSchemasEdgeCacheTTLIDEdgeCacheTTL SettingEditResponseZonesSchemasEdgeCacheTTLID = "edge_cache_ttl"
)

func (r SettingEditResponseZonesSchemasEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasEdgeCacheTTLValue float64

const (
	SettingEditResponseZonesSchemasEdgeCacheTTLValue30     SettingEditResponseZonesSchemasEdgeCacheTTLValue = 30
	SettingEditResponseZonesSchemasEdgeCacheTTLValue60     SettingEditResponseZonesSchemasEdgeCacheTTLValue = 60
	SettingEditResponseZonesSchemasEdgeCacheTTLValue300    SettingEditResponseZonesSchemasEdgeCacheTTLValue = 300
	SettingEditResponseZonesSchemasEdgeCacheTTLValue1200   SettingEditResponseZonesSchemasEdgeCacheTTLValue = 1200
	SettingEditResponseZonesSchemasEdgeCacheTTLValue1800   SettingEditResponseZonesSchemasEdgeCacheTTLValue = 1800
	SettingEditResponseZonesSchemasEdgeCacheTTLValue3600   SettingEditResponseZonesSchemasEdgeCacheTTLValue = 3600
	SettingEditResponseZonesSchemasEdgeCacheTTLValue7200   SettingEditResponseZonesSchemasEdgeCacheTTLValue = 7200
	SettingEditResponseZonesSchemasEdgeCacheTTLValue10800  SettingEditResponseZonesSchemasEdgeCacheTTLValue = 10800
	SettingEditResponseZonesSchemasEdgeCacheTTLValue14400  SettingEditResponseZonesSchemasEdgeCacheTTLValue = 14400
	SettingEditResponseZonesSchemasEdgeCacheTTLValue18000  SettingEditResponseZonesSchemasEdgeCacheTTLValue = 18000
	SettingEditResponseZonesSchemasEdgeCacheTTLValue28800  SettingEditResponseZonesSchemasEdgeCacheTTLValue = 28800
	SettingEditResponseZonesSchemasEdgeCacheTTLValue43200  SettingEditResponseZonesSchemasEdgeCacheTTLValue = 43200
	SettingEditResponseZonesSchemasEdgeCacheTTLValue57600  SettingEditResponseZonesSchemasEdgeCacheTTLValue = 57600
	SettingEditResponseZonesSchemasEdgeCacheTTLValue72000  SettingEditResponseZonesSchemasEdgeCacheTTLValue = 72000
	SettingEditResponseZonesSchemasEdgeCacheTTLValue86400  SettingEditResponseZonesSchemasEdgeCacheTTLValue = 86400
	SettingEditResponseZonesSchemasEdgeCacheTTLValue172800 SettingEditResponseZonesSchemasEdgeCacheTTLValue = 172800
	SettingEditResponseZonesSchemasEdgeCacheTTLValue259200 SettingEditResponseZonesSchemasEdgeCacheTTLValue = 259200
	SettingEditResponseZonesSchemasEdgeCacheTTLValue345600 SettingEditResponseZonesSchemasEdgeCacheTTLValue = 345600
	SettingEditResponseZonesSchemasEdgeCacheTTLValue432000 SettingEditResponseZonesSchemasEdgeCacheTTLValue = 432000
	SettingEditResponseZonesSchemasEdgeCacheTTLValue518400 SettingEditResponseZonesSchemasEdgeCacheTTLValue = 518400
	SettingEditResponseZonesSchemasEdgeCacheTTLValue604800 SettingEditResponseZonesSchemasEdgeCacheTTLValue = 604800
)

func (r SettingEditResponseZonesSchemasEdgeCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasEdgeCacheTTLValue30, SettingEditResponseZonesSchemasEdgeCacheTTLValue60, SettingEditResponseZonesSchemasEdgeCacheTTLValue300, SettingEditResponseZonesSchemasEdgeCacheTTLValue1200, SettingEditResponseZonesSchemasEdgeCacheTTLValue1800, SettingEditResponseZonesSchemasEdgeCacheTTLValue3600, SettingEditResponseZonesSchemasEdgeCacheTTLValue7200, SettingEditResponseZonesSchemasEdgeCacheTTLValue10800, SettingEditResponseZonesSchemasEdgeCacheTTLValue14400, SettingEditResponseZonesSchemasEdgeCacheTTLValue18000, SettingEditResponseZonesSchemasEdgeCacheTTLValue28800, SettingEditResponseZonesSchemasEdgeCacheTTLValue43200, SettingEditResponseZonesSchemasEdgeCacheTTLValue57600, SettingEditResponseZonesSchemasEdgeCacheTTLValue72000, SettingEditResponseZonesSchemasEdgeCacheTTLValue86400, SettingEditResponseZonesSchemasEdgeCacheTTLValue172800, SettingEditResponseZonesSchemasEdgeCacheTTLValue259200, SettingEditResponseZonesSchemasEdgeCacheTTLValue345600, SettingEditResponseZonesSchemasEdgeCacheTTLValue432000, SettingEditResponseZonesSchemasEdgeCacheTTLValue518400, SettingEditResponseZonesSchemasEdgeCacheTTLValue604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasEdgeCacheTTLEditable bool

const (
	SettingEditResponseZonesSchemasEdgeCacheTTLEditableTrue  SettingEditResponseZonesSchemasEdgeCacheTTLEditable = true
	SettingEditResponseZonesSchemasEdgeCacheTTLEditableFalse SettingEditResponseZonesSchemasEdgeCacheTTLEditable = false
)

func (r SettingEditResponseZonesSchemasEdgeCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasEdgeCacheTTLEditableTrue, SettingEditResponseZonesSchemasEdgeCacheTTLEditableFalse:
		return true
	}
	return false
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingEditResponseZonesSchemasEmailObfuscation struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasEmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasEmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasEmailObfuscationJSON `json:"-"`
}

// settingEditResponseZonesSchemasEmailObfuscationJSON contains the JSON metadata
// for the struct [SettingEditResponseZonesSchemasEmailObfuscation]
type settingEditResponseZonesSchemasEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasEmailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasEmailObfuscation) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasEmailObfuscationID string

const (
	SettingEditResponseZonesSchemasEmailObfuscationIDEmailObfuscation SettingEditResponseZonesSchemasEmailObfuscationID = "email_obfuscation"
)

func (r SettingEditResponseZonesSchemasEmailObfuscationID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasEmailObfuscationIDEmailObfuscation:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasEmailObfuscationValue string

const (
	SettingEditResponseZonesSchemasEmailObfuscationValueOn  SettingEditResponseZonesSchemasEmailObfuscationValue = "on"
	SettingEditResponseZonesSchemasEmailObfuscationValueOff SettingEditResponseZonesSchemasEmailObfuscationValue = "off"
)

func (r SettingEditResponseZonesSchemasEmailObfuscationValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasEmailObfuscationValueOn, SettingEditResponseZonesSchemasEmailObfuscationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasEmailObfuscationEditable bool

const (
	SettingEditResponseZonesSchemasEmailObfuscationEditableTrue  SettingEditResponseZonesSchemasEmailObfuscationEditable = true
	SettingEditResponseZonesSchemasEmailObfuscationEditableFalse SettingEditResponseZonesSchemasEmailObfuscationEditable = false
)

func (r SettingEditResponseZonesSchemasEmailObfuscationEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasEmailObfuscationEditableTrue, SettingEditResponseZonesSchemasEmailObfuscationEditableFalse:
		return true
	}
	return false
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingEditResponseZonesSchemasIPGeolocation struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasIPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasIPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasIPGeolocationJSON `json:"-"`
}

// settingEditResponseZonesSchemasIPGeolocationJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesSchemasIPGeolocation]
type settingEditResponseZonesSchemasIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasIPGeolocationJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasIPGeolocation) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasIPGeolocationID string

const (
	SettingEditResponseZonesSchemasIPGeolocationIDIPGeolocation SettingEditResponseZonesSchemasIPGeolocationID = "ip_geolocation"
)

func (r SettingEditResponseZonesSchemasIPGeolocationID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasIPGeolocationIDIPGeolocation:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasIPGeolocationValue string

const (
	SettingEditResponseZonesSchemasIPGeolocationValueOn  SettingEditResponseZonesSchemasIPGeolocationValue = "on"
	SettingEditResponseZonesSchemasIPGeolocationValueOff SettingEditResponseZonesSchemasIPGeolocationValue = "off"
)

func (r SettingEditResponseZonesSchemasIPGeolocationValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasIPGeolocationValueOn, SettingEditResponseZonesSchemasIPGeolocationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasIPGeolocationEditable bool

const (
	SettingEditResponseZonesSchemasIPGeolocationEditableTrue  SettingEditResponseZonesSchemasIPGeolocationEditable = true
	SettingEditResponseZonesSchemasIPGeolocationEditableFalse SettingEditResponseZonesSchemasIPGeolocationEditable = false
)

func (r SettingEditResponseZonesSchemasIPGeolocationEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasIPGeolocationEditableTrue, SettingEditResponseZonesSchemasIPGeolocationEditableFalse:
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

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type SettingEditResponseZonesSchemasMirage struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasMirageID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasMirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasMirageJSON `json:"-"`
}

// settingEditResponseZonesSchemasMirageJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesSchemasMirage]
type settingEditResponseZonesSchemasMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasMirageJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasMirage) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasMirageID string

const (
	SettingEditResponseZonesSchemasMirageIDMirage SettingEditResponseZonesSchemasMirageID = "mirage"
)

func (r SettingEditResponseZonesSchemasMirageID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasMirageIDMirage:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasMirageValue string

const (
	SettingEditResponseZonesSchemasMirageValueOn  SettingEditResponseZonesSchemasMirageValue = "on"
	SettingEditResponseZonesSchemasMirageValueOff SettingEditResponseZonesSchemasMirageValue = "off"
)

func (r SettingEditResponseZonesSchemasMirageValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasMirageValueOn, SettingEditResponseZonesSchemasMirageValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasMirageEditable bool

const (
	SettingEditResponseZonesSchemasMirageEditableTrue  SettingEditResponseZonesSchemasMirageEditable = true
	SettingEditResponseZonesSchemasMirageEditableFalse SettingEditResponseZonesSchemasMirageEditable = false
)

func (r SettingEditResponseZonesSchemasMirageEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasMirageEditableTrue, SettingEditResponseZonesSchemasMirageEditableFalse:
		return true
	}
	return false
}

// Enables the Opportunistic Encryption feature for a zone.
type SettingEditResponseZonesSchemasOpportunisticEncryption struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasOpportunisticEncryptionJSON `json:"-"`
}

// settingEditResponseZonesSchemasOpportunisticEncryptionJSON contains the JSON
// metadata for the struct [SettingEditResponseZonesSchemasOpportunisticEncryption]
type settingEditResponseZonesSchemasOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasOpportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasOpportunisticEncryption) implementsZonesSettingEditResponse() {
}

// ID of the zone setting.
type SettingEditResponseZonesSchemasOpportunisticEncryptionID string

const (
	SettingEditResponseZonesSchemasOpportunisticEncryptionIDOpportunisticEncryption SettingEditResponseZonesSchemasOpportunisticEncryptionID = "opportunistic_encryption"
)

func (r SettingEditResponseZonesSchemasOpportunisticEncryptionID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasOpportunisticEncryptionIDOpportunisticEncryption:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasOpportunisticEncryptionValue string

const (
	SettingEditResponseZonesSchemasOpportunisticEncryptionValueOn  SettingEditResponseZonesSchemasOpportunisticEncryptionValue = "on"
	SettingEditResponseZonesSchemasOpportunisticEncryptionValueOff SettingEditResponseZonesSchemasOpportunisticEncryptionValue = "off"
)

func (r SettingEditResponseZonesSchemasOpportunisticEncryptionValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasOpportunisticEncryptionValueOn, SettingEditResponseZonesSchemasOpportunisticEncryptionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasOpportunisticEncryptionEditable bool

const (
	SettingEditResponseZonesSchemasOpportunisticEncryptionEditableTrue  SettingEditResponseZonesSchemasOpportunisticEncryptionEditable = true
	SettingEditResponseZonesSchemasOpportunisticEncryptionEditableFalse SettingEditResponseZonesSchemasOpportunisticEncryptionEditable = false
)

func (r SettingEditResponseZonesSchemasOpportunisticEncryptionEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasOpportunisticEncryptionEditableTrue, SettingEditResponseZonesSchemasOpportunisticEncryptionEditableFalse:
		return true
	}
	return false
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingEditResponseZonesSchemasOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasOriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasOriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasOriginErrorPagePassThruJSON `json:"-"`
}

// settingEditResponseZonesSchemasOriginErrorPagePassThruJSON contains the JSON
// metadata for the struct [SettingEditResponseZonesSchemasOriginErrorPagePassThru]
type settingEditResponseZonesSchemasOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasOriginErrorPagePassThruJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasOriginErrorPagePassThru) implementsZonesSettingEditResponse() {
}

// ID of the zone setting.
type SettingEditResponseZonesSchemasOriginErrorPagePassThruID string

const (
	SettingEditResponseZonesSchemasOriginErrorPagePassThruIDOriginErrorPagePassThru SettingEditResponseZonesSchemasOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

func (r SettingEditResponseZonesSchemasOriginErrorPagePassThruID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasOriginErrorPagePassThruIDOriginErrorPagePassThru:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasOriginErrorPagePassThruValue string

const (
	SettingEditResponseZonesSchemasOriginErrorPagePassThruValueOn  SettingEditResponseZonesSchemasOriginErrorPagePassThruValue = "on"
	SettingEditResponseZonesSchemasOriginErrorPagePassThruValueOff SettingEditResponseZonesSchemasOriginErrorPagePassThruValue = "off"
)

func (r SettingEditResponseZonesSchemasOriginErrorPagePassThruValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasOriginErrorPagePassThruValueOn, SettingEditResponseZonesSchemasOriginErrorPagePassThruValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasOriginErrorPagePassThruEditable bool

const (
	SettingEditResponseZonesSchemasOriginErrorPagePassThruEditableTrue  SettingEditResponseZonesSchemasOriginErrorPagePassThruEditable = true
	SettingEditResponseZonesSchemasOriginErrorPagePassThruEditableFalse SettingEditResponseZonesSchemasOriginErrorPagePassThruEditable = false
)

func (r SettingEditResponseZonesSchemasOriginErrorPagePassThruEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasOriginErrorPagePassThruEditableTrue, SettingEditResponseZonesSchemasOriginErrorPagePassThruEditableFalse:
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
type SettingEditResponseZonesSchemasPolish struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasPolishID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasPolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasPolishJSON `json:"-"`
}

// settingEditResponseZonesSchemasPolishJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesSchemasPolish]
type settingEditResponseZonesSchemasPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasPolishJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasPolish) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasPolishID string

const (
	SettingEditResponseZonesSchemasPolishIDPolish SettingEditResponseZonesSchemasPolishID = "polish"
)

func (r SettingEditResponseZonesSchemasPolishID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasPolishIDPolish:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasPolishValue string

const (
	SettingEditResponseZonesSchemasPolishValueOff      SettingEditResponseZonesSchemasPolishValue = "off"
	SettingEditResponseZonesSchemasPolishValueLossless SettingEditResponseZonesSchemasPolishValue = "lossless"
	SettingEditResponseZonesSchemasPolishValueLossy    SettingEditResponseZonesSchemasPolishValue = "lossy"
)

func (r SettingEditResponseZonesSchemasPolishValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasPolishValueOff, SettingEditResponseZonesSchemasPolishValueLossless, SettingEditResponseZonesSchemasPolishValueLossy:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasPolishEditable bool

const (
	SettingEditResponseZonesSchemasPolishEditableTrue  SettingEditResponseZonesSchemasPolishEditable = true
	SettingEditResponseZonesSchemasPolishEditableFalse SettingEditResponseZonesSchemasPolishEditable = false
)

func (r SettingEditResponseZonesSchemasPolishEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasPolishEditableTrue, SettingEditResponseZonesSchemasPolishEditableFalse:
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

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingEditResponseZonesSchemasResponseBuffering struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasResponseBufferingID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasResponseBufferingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasResponseBufferingJSON `json:"-"`
}

// settingEditResponseZonesSchemasResponseBufferingJSON contains the JSON metadata
// for the struct [SettingEditResponseZonesSchemasResponseBuffering]
type settingEditResponseZonesSchemasResponseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasResponseBufferingJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasResponseBuffering) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasResponseBufferingID string

const (
	SettingEditResponseZonesSchemasResponseBufferingIDResponseBuffering SettingEditResponseZonesSchemasResponseBufferingID = "response_buffering"
)

func (r SettingEditResponseZonesSchemasResponseBufferingID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasResponseBufferingIDResponseBuffering:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasResponseBufferingValue string

const (
	SettingEditResponseZonesSchemasResponseBufferingValueOn  SettingEditResponseZonesSchemasResponseBufferingValue = "on"
	SettingEditResponseZonesSchemasResponseBufferingValueOff SettingEditResponseZonesSchemasResponseBufferingValue = "off"
)

func (r SettingEditResponseZonesSchemasResponseBufferingValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasResponseBufferingValueOn, SettingEditResponseZonesSchemasResponseBufferingValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasResponseBufferingEditable bool

const (
	SettingEditResponseZonesSchemasResponseBufferingEditableTrue  SettingEditResponseZonesSchemasResponseBufferingEditable = true
	SettingEditResponseZonesSchemasResponseBufferingEditableFalse SettingEditResponseZonesSchemasResponseBufferingEditable = false
)

func (r SettingEditResponseZonesSchemasResponseBufferingEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasResponseBufferingEditableTrue, SettingEditResponseZonesSchemasResponseBufferingEditableFalse:
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
type SettingEditResponseZonesSchemasRocketLoader struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasRocketLoaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasRocketLoaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasRocketLoaderJSON `json:"-"`
}

// settingEditResponseZonesSchemasRocketLoaderJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesSchemasRocketLoader]
type settingEditResponseZonesSchemasRocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasRocketLoaderJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasRocketLoader) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasRocketLoaderID string

const (
	SettingEditResponseZonesSchemasRocketLoaderIDRocketLoader SettingEditResponseZonesSchemasRocketLoaderID = "rocket_loader"
)

func (r SettingEditResponseZonesSchemasRocketLoaderID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasRocketLoaderIDRocketLoader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasRocketLoaderValue string

const (
	SettingEditResponseZonesSchemasRocketLoaderValueOn  SettingEditResponseZonesSchemasRocketLoaderValue = "on"
	SettingEditResponseZonesSchemasRocketLoaderValueOff SettingEditResponseZonesSchemasRocketLoaderValue = "off"
)

func (r SettingEditResponseZonesSchemasRocketLoaderValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasRocketLoaderValueOn, SettingEditResponseZonesSchemasRocketLoaderValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasRocketLoaderEditable bool

const (
	SettingEditResponseZonesSchemasRocketLoaderEditableTrue  SettingEditResponseZonesSchemasRocketLoaderEditable = true
	SettingEditResponseZonesSchemasRocketLoaderEditableFalse SettingEditResponseZonesSchemasRocketLoaderEditable = false
)

func (r SettingEditResponseZonesSchemasRocketLoaderEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasRocketLoaderEditableTrue, SettingEditResponseZonesSchemasRocketLoaderEditableFalse:
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

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingEditResponseZonesSchemasSecurityLevel struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasSecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasSecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasSecurityLevelJSON `json:"-"`
}

// settingEditResponseZonesSchemasSecurityLevelJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesSchemasSecurityLevel]
type settingEditResponseZonesSchemasSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasSecurityLevelJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasSecurityLevel) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasSecurityLevelID string

const (
	SettingEditResponseZonesSchemasSecurityLevelIDSecurityLevel SettingEditResponseZonesSchemasSecurityLevelID = "security_level"
)

func (r SettingEditResponseZonesSchemasSecurityLevelID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasSecurityLevelIDSecurityLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasSecurityLevelValue string

const (
	SettingEditResponseZonesSchemasSecurityLevelValueOff            SettingEditResponseZonesSchemasSecurityLevelValue = "off"
	SettingEditResponseZonesSchemasSecurityLevelValueEssentiallyOff SettingEditResponseZonesSchemasSecurityLevelValue = "essentially_off"
	SettingEditResponseZonesSchemasSecurityLevelValueLow            SettingEditResponseZonesSchemasSecurityLevelValue = "low"
	SettingEditResponseZonesSchemasSecurityLevelValueMedium         SettingEditResponseZonesSchemasSecurityLevelValue = "medium"
	SettingEditResponseZonesSchemasSecurityLevelValueHigh           SettingEditResponseZonesSchemasSecurityLevelValue = "high"
	SettingEditResponseZonesSchemasSecurityLevelValueUnderAttack    SettingEditResponseZonesSchemasSecurityLevelValue = "under_attack"
)

func (r SettingEditResponseZonesSchemasSecurityLevelValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasSecurityLevelValueOff, SettingEditResponseZonesSchemasSecurityLevelValueEssentiallyOff, SettingEditResponseZonesSchemasSecurityLevelValueLow, SettingEditResponseZonesSchemasSecurityLevelValueMedium, SettingEditResponseZonesSchemasSecurityLevelValueHigh, SettingEditResponseZonesSchemasSecurityLevelValueUnderAttack:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasSecurityLevelEditable bool

const (
	SettingEditResponseZonesSchemasSecurityLevelEditableTrue  SettingEditResponseZonesSchemasSecurityLevelEditable = true
	SettingEditResponseZonesSchemasSecurityLevelEditableFalse SettingEditResponseZonesSchemasSecurityLevelEditable = false
)

func (r SettingEditResponseZonesSchemasSecurityLevelEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasSecurityLevelEditableTrue, SettingEditResponseZonesSchemasSecurityLevelEditableFalse:
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

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type SettingEditResponseZonesSchemasSortQueryStringForCache struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasSortQueryStringForCacheID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasSortQueryStringForCacheValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasSortQueryStringForCacheJSON `json:"-"`
}

// settingEditResponseZonesSchemasSortQueryStringForCacheJSON contains the JSON
// metadata for the struct [SettingEditResponseZonesSchemasSortQueryStringForCache]
type settingEditResponseZonesSchemasSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasSortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasSortQueryStringForCache) implementsZonesSettingEditResponse() {
}

// ID of the zone setting.
type SettingEditResponseZonesSchemasSortQueryStringForCacheID string

const (
	SettingEditResponseZonesSchemasSortQueryStringForCacheIDSortQueryStringForCache SettingEditResponseZonesSchemasSortQueryStringForCacheID = "sort_query_string_for_cache"
)

func (r SettingEditResponseZonesSchemasSortQueryStringForCacheID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasSortQueryStringForCacheIDSortQueryStringForCache:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasSortQueryStringForCacheValue string

const (
	SettingEditResponseZonesSchemasSortQueryStringForCacheValueOn  SettingEditResponseZonesSchemasSortQueryStringForCacheValue = "on"
	SettingEditResponseZonesSchemasSortQueryStringForCacheValueOff SettingEditResponseZonesSchemasSortQueryStringForCacheValue = "off"
)

func (r SettingEditResponseZonesSchemasSortQueryStringForCacheValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasSortQueryStringForCacheValueOn, SettingEditResponseZonesSchemasSortQueryStringForCacheValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasSortQueryStringForCacheEditable bool

const (
	SettingEditResponseZonesSchemasSortQueryStringForCacheEditableTrue  SettingEditResponseZonesSchemasSortQueryStringForCacheEditable = true
	SettingEditResponseZonesSchemasSortQueryStringForCacheEditableFalse SettingEditResponseZonesSchemasSortQueryStringForCacheEditable = false
)

func (r SettingEditResponseZonesSchemasSortQueryStringForCacheEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasSortQueryStringForCacheEditableTrue, SettingEditResponseZonesSchemasSortQueryStringForCacheEditableFalse:
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
type SettingEditResponseZonesSchemasSSL struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasSSLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasSSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasSSLJSON `json:"-"`
}

// settingEditResponseZonesSchemasSSLJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesSchemasSSL]
type settingEditResponseZonesSchemasSSLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasSSLJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasSSL) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasSSLID string

const (
	SettingEditResponseZonesSchemasSSLIDSSL SettingEditResponseZonesSchemasSSLID = "ssl"
)

func (r SettingEditResponseZonesSchemasSSLID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasSSLIDSSL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasSSLValue string

const (
	SettingEditResponseZonesSchemasSSLValueOff      SettingEditResponseZonesSchemasSSLValue = "off"
	SettingEditResponseZonesSchemasSSLValueFlexible SettingEditResponseZonesSchemasSSLValue = "flexible"
	SettingEditResponseZonesSchemasSSLValueFull     SettingEditResponseZonesSchemasSSLValue = "full"
	SettingEditResponseZonesSchemasSSLValueStrict   SettingEditResponseZonesSchemasSSLValue = "strict"
)

func (r SettingEditResponseZonesSchemasSSLValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasSSLValueOff, SettingEditResponseZonesSchemasSSLValueFlexible, SettingEditResponseZonesSchemasSSLValueFull, SettingEditResponseZonesSchemasSSLValueStrict:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasSSLEditable bool

const (
	SettingEditResponseZonesSchemasSSLEditableTrue  SettingEditResponseZonesSchemasSSLEditable = true
	SettingEditResponseZonesSchemasSSLEditableFalse SettingEditResponseZonesSchemasSSLEditable = false
)

func (r SettingEditResponseZonesSchemasSSLEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasSSLEditableTrue, SettingEditResponseZonesSchemasSSLEditableFalse:
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

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingEditResponseZonesSchemasTrueClientIPHeader struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasTrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasTrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasTrueClientIPHeaderJSON `json:"-"`
}

// settingEditResponseZonesSchemasTrueClientIPHeaderJSON contains the JSON metadata
// for the struct [SettingEditResponseZonesSchemasTrueClientIPHeader]
type settingEditResponseZonesSchemasTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasTrueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasTrueClientIPHeader) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasTrueClientIPHeaderID string

const (
	SettingEditResponseZonesSchemasTrueClientIPHeaderIDTrueClientIPHeader SettingEditResponseZonesSchemasTrueClientIPHeaderID = "true_client_ip_header"
)

func (r SettingEditResponseZonesSchemasTrueClientIPHeaderID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasTrueClientIPHeaderIDTrueClientIPHeader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasTrueClientIPHeaderValue string

const (
	SettingEditResponseZonesSchemasTrueClientIPHeaderValueOn  SettingEditResponseZonesSchemasTrueClientIPHeaderValue = "on"
	SettingEditResponseZonesSchemasTrueClientIPHeaderValueOff SettingEditResponseZonesSchemasTrueClientIPHeaderValue = "off"
)

func (r SettingEditResponseZonesSchemasTrueClientIPHeaderValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasTrueClientIPHeaderValueOn, SettingEditResponseZonesSchemasTrueClientIPHeaderValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasTrueClientIPHeaderEditable bool

const (
	SettingEditResponseZonesSchemasTrueClientIPHeaderEditableTrue  SettingEditResponseZonesSchemasTrueClientIPHeaderEditable = true
	SettingEditResponseZonesSchemasTrueClientIPHeaderEditableFalse SettingEditResponseZonesSchemasTrueClientIPHeaderEditable = false
)

func (r SettingEditResponseZonesSchemasTrueClientIPHeaderEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasTrueClientIPHeaderEditableTrue, SettingEditResponseZonesSchemasTrueClientIPHeaderEditableFalse:
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
type SettingEditResponseZonesSchemasWAF struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasWAFID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasWAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSchemasWAFJSON `json:"-"`
}

// settingEditResponseZonesSchemasWAFJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesSchemasWAF]
type settingEditResponseZonesSchemasWAFJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSchemasWAFJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSchemasWAF) implementsZonesSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSchemasWAFID string

const (
	SettingEditResponseZonesSchemasWAFIDWAF SettingEditResponseZonesSchemasWAFID = "waf"
)

func (r SettingEditResponseZonesSchemasWAFID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasWAFIDWAF:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSchemasWAFValue string

const (
	SettingEditResponseZonesSchemasWAFValueOn  SettingEditResponseZonesSchemasWAFValue = "on"
	SettingEditResponseZonesSchemasWAFValueOff SettingEditResponseZonesSchemasWAFValue = "off"
)

func (r SettingEditResponseZonesSchemasWAFValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasWAFValueOn, SettingEditResponseZonesSchemasWAFValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasWAFEditable bool

const (
	SettingEditResponseZonesSchemasWAFEditableTrue  SettingEditResponseZonesSchemasWAFEditable = true
	SettingEditResponseZonesSchemasWAFEditableFalse SettingEditResponseZonesSchemasWAFEditable = false
)

func (r SettingEditResponseZonesSchemasWAFEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSchemasWAFEditableTrue, SettingEditResponseZonesSchemasWAFEditableFalse:
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
	SettingEditResponseIDMirage                        SettingEditResponseID = "mirage"
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
	case SettingEditResponseID0rtt, SettingEditResponseIDAdvancedDDoS, SettingEditResponseIDAlwaysOnline, SettingEditResponseIDAlwaysUseHTTPS, SettingEditResponseIDAutomaticHTTPSRewrites, SettingEditResponseIDBrotli, SettingEditResponseIDBrowserCacheTTL, SettingEditResponseIDBrowserCheck, SettingEditResponseIDCacheLevel, SettingEditResponseIDChallengeTTL, SettingEditResponseIDCiphers, SettingEditResponseIDCNAMEFlattening, SettingEditResponseIDDevelopmentMode, SettingEditResponseIDEarlyHints, SettingEditResponseIDEdgeCacheTTL, SettingEditResponseIDEmailObfuscation, SettingEditResponseIDH2Prioritization, SettingEditResponseIDHotlinkProtection, SettingEditResponseIDHTTP2, SettingEditResponseIDHTTP3, SettingEditResponseIDImageResizing, SettingEditResponseIDIPGeolocation, SettingEditResponseIDIPV6, SettingEditResponseIDMaxUpload, SettingEditResponseIDMinTLSVersion, SettingEditResponseIDMirage, SettingEditResponseIDNEL, SettingEditResponseIDOpportunisticEncryption, SettingEditResponseIDOpportunisticOnion, SettingEditResponseIDOrangeToOrange, SettingEditResponseIDOriginErrorPagePassThru, SettingEditResponseIDPolish, SettingEditResponseIDPrefetchPreload, SettingEditResponseIDProxyReadTimeout, SettingEditResponseIDPseudoIPV4, SettingEditResponseIDReplaceInsecureJS, SettingEditResponseIDResponseBuffering, SettingEditResponseIDRocketLoader, SettingEditResponseIDAutomaticPlatformOptimization, SettingEditResponseIDSecurityHeader, SettingEditResponseIDSecurityLevel, SettingEditResponseIDServerSideExclude, SettingEditResponseIDSha1Support, SettingEditResponseIDSortQueryStringForCache, SettingEditResponseIDSSL, SettingEditResponseIDSSLRecommender, SettingEditResponseIDTLS1_2Only, SettingEditResponseIDTLS1_3, SettingEditResponseIDTLSClientAuth, SettingEditResponseIDTrueClientIPHeader, SettingEditResponseIDWAF, SettingEditResponseIDWebP, SettingEditResponseIDWebsockets:
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

// 0-RTT session resumption enabled for this zone.
type SettingGetResponse struct {
	// ID of the zone setting.
	ID SettingGetResponseID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseEditable `json:"editable"`
	// ssl-recommender enrollment setting.
	Enabled bool `json:"enabled"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// This field can have the runtime type of [ZeroRTTValue], [AdvancedDDoSValue],
	// [AlwaysOnlineValue], [SettingGetResponseZonesSchemasAlwaysUseHTTPSValue],
	// [SettingGetResponseZonesSchemasAutomaticHTTPSRewritesValue], [BrotliValue],
	// [SettingGetResponseZonesSchemasBrowserCacheTTLValue],
	// [SettingGetResponseZonesSchemasBrowserCheckValue],
	// [SettingGetResponseZonesSchemasCacheLevelValue], [ChallengeTTLValue],
	// [[]string], [SettingGetResponseZonesCNAMEFlatteningValue],
	// [DevelopmentModeValue], [EarlyHintsValue],
	// [SettingGetResponseZonesSchemasEdgeCacheTTLValue],
	// [SettingGetResponseZonesSchemasEmailObfuscationValue], [H2PrioritizationValue],
	// [HotlinkProtectionValue], [HTTP2Value], [HTTP3Value], [ImageResizingValue],
	// [SettingGetResponseZonesSchemasIPGeolocationValue], [IPV6Value],
	// [SettingGetResponseZonesMaxUploadValue], [MinTLSVersionValue],
	// [SettingGetResponseZonesSchemasMirageValue], [NELValue],
	// [SettingGetResponseZonesSchemasOpportunisticEncryptionValue],
	// [OpportunisticOnionValue], [OrangeToOrangeValue],
	// [SettingGetResponseZonesSchemasOriginErrorPagePassThruValue],
	// [SettingGetResponseZonesSchemasPolishValue], [PrefetchPreloadValue], [float64],
	// [PseudoIPV4Value], [SettingGetResponseZonesReplaceInsecureJSValue],
	// [SettingGetResponseZonesSchemasResponseBufferingValue],
	// [SettingGetResponseZonesSchemasRocketLoaderValue],
	// [AutomaticPlatformOptimization], [SecurityHeadersValue],
	// [SettingGetResponseZonesSchemasSecurityLevelValue], [ServerSideExcludesValue],
	// [SettingGetResponseZonesSha1SupportValue],
	// [SettingGetResponseZonesSchemasSortQueryStringForCacheValue],
	// [SettingGetResponseZonesSchemasSSLValue],
	// [SettingGetResponseZonesTLS1_2OnlyValue], [TLS1_3Value], [TLSClientAuthValue],
	// [SettingGetResponseZonesSchemasTrueClientIPHeaderValue],
	// [SettingGetResponseZonesSchemasWAFValue], [WebPValue], [WebsocketValue].
	Value interface{}            `json:"value"`
	JSON  settingGetResponseJSON `json:"-"`
	union SettingGetResponseUnion
}

// settingGetResponseJSON contains the JSON metadata for the struct
// [SettingGetResponse]
type settingGetResponseJSON struct {
	ID            apijson.Field
	Editable      apijson.Field
	Enabled       apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r settingGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = SettingGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [SettingGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [zones.ZeroRTT], [zones.AdvancedDDoS],
// [zones.AlwaysOnline], [zones.SettingGetResponseZonesSchemasAlwaysUseHTTPS],
// [zones.SettingGetResponseZonesSchemasAutomaticHTTPSRewrites], [zones.Brotli],
// [zones.SettingGetResponseZonesSchemasBrowserCacheTTL],
// [zones.SettingGetResponseZonesSchemasBrowserCheck],
// [zones.SettingGetResponseZonesSchemasCacheLevel], [zones.ChallengeTTL],
// [zones.Ciphers], [zones.SettingGetResponseZonesCNAMEFlattening],
// [zones.DevelopmentMode], [zones.EarlyHints],
// [zones.SettingGetResponseZonesSchemasEdgeCacheTTL],
// [zones.SettingGetResponseZonesSchemasEmailObfuscation],
// [zones.H2Prioritization], [zones.HotlinkProtection], [zones.HTTP2],
// [zones.HTTP3], [zones.ImageResizing],
// [zones.SettingGetResponseZonesSchemasIPGeolocation], [zones.IPV6],
// [zones.SettingGetResponseZonesMaxUpload], [zones.MinTLSVersion],
// [zones.SettingGetResponseZonesSchemasMirage], [zones.NEL],
// [zones.SettingGetResponseZonesSchemasOpportunisticEncryption],
// [zones.OpportunisticOnion], [zones.OrangeToOrange],
// [zones.SettingGetResponseZonesSchemasOriginErrorPagePassThru],
// [zones.SettingGetResponseZonesSchemasPolish], [zones.PrefetchPreload],
// [zones.ProxyReadTimeout], [zones.PseudoIPV4],
// [zones.SettingGetResponseZonesReplaceInsecureJS],
// [zones.SettingGetResponseZonesSchemasResponseBuffering],
// [zones.SettingGetResponseZonesSchemasRocketLoader],
// [zones.SettingGetResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.SecurityHeaders], [zones.SettingGetResponseZonesSchemasSecurityLevel],
// [zones.ServerSideExcludes], [zones.SettingGetResponseZonesSha1Support],
// [zones.SettingGetResponseZonesSchemasSortQueryStringForCache],
// [zones.SettingGetResponseZonesSchemasSSL], [zones.SSLRecommender],
// [zones.SettingGetResponseZonesTLS1_2Only], [zones.TLS1_3],
// [zones.TLSClientAuth], [zones.SettingGetResponseZonesSchemasTrueClientIPHeader],
// [zones.SettingGetResponseZonesSchemasWAF], [zones.WebP], [zones.Websocket].
func (r SettingGetResponse) AsUnion() SettingGetResponseUnion {
	return r.union
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [zones.ZeroRTT], [zones.AdvancedDDoS], [zones.AlwaysOnline],
// [zones.SettingGetResponseZonesSchemasAlwaysUseHTTPS],
// [zones.SettingGetResponseZonesSchemasAutomaticHTTPSRewrites], [zones.Brotli],
// [zones.SettingGetResponseZonesSchemasBrowserCacheTTL],
// [zones.SettingGetResponseZonesSchemasBrowserCheck],
// [zones.SettingGetResponseZonesSchemasCacheLevel], [zones.ChallengeTTL],
// [zones.Ciphers], [zones.SettingGetResponseZonesCNAMEFlattening],
// [zones.DevelopmentMode], [zones.EarlyHints],
// [zones.SettingGetResponseZonesSchemasEdgeCacheTTL],
// [zones.SettingGetResponseZonesSchemasEmailObfuscation],
// [zones.H2Prioritization], [zones.HotlinkProtection], [zones.HTTP2],
// [zones.HTTP3], [zones.ImageResizing],
// [zones.SettingGetResponseZonesSchemasIPGeolocation], [zones.IPV6],
// [zones.SettingGetResponseZonesMaxUpload], [zones.MinTLSVersion],
// [zones.SettingGetResponseZonesSchemasMirage], [zones.NEL],
// [zones.SettingGetResponseZonesSchemasOpportunisticEncryption],
// [zones.OpportunisticOnion], [zones.OrangeToOrange],
// [zones.SettingGetResponseZonesSchemasOriginErrorPagePassThru],
// [zones.SettingGetResponseZonesSchemasPolish], [zones.PrefetchPreload],
// [zones.ProxyReadTimeout], [zones.PseudoIPV4],
// [zones.SettingGetResponseZonesReplaceInsecureJS],
// [zones.SettingGetResponseZonesSchemasResponseBuffering],
// [zones.SettingGetResponseZonesSchemasRocketLoader],
// [zones.SettingGetResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.SecurityHeaders], [zones.SettingGetResponseZonesSchemasSecurityLevel],
// [zones.ServerSideExcludes], [zones.SettingGetResponseZonesSha1Support],
// [zones.SettingGetResponseZonesSchemasSortQueryStringForCache],
// [zones.SettingGetResponseZonesSchemasSSL], [zones.SSLRecommender],
// [zones.SettingGetResponseZonesTLS1_2Only], [zones.TLS1_3],
// [zones.TLSClientAuth], [zones.SettingGetResponseZonesSchemasTrueClientIPHeader],
// [zones.SettingGetResponseZonesSchemasWAF], [zones.WebP] or [zones.Websocket].
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
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Brotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasBrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasCacheLevel{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasEmailObfuscation{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasIPGeolocation{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasOpportunisticEncryption{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasPolish{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasResponseBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasRocketLoader{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasSecurityLevel{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasSSL{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasWAF{}),
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

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingGetResponseZonesSchemasAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasAlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasAlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasAlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasAlwaysUseHTTPSJSON `json:"-"`
}

// settingGetResponseZonesSchemasAlwaysUseHTTPSJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesSchemasAlwaysUseHTTPS]
type settingGetResponseZonesSchemasAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasAlwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasAlwaysUseHTTPS) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasAlwaysUseHTTPSID string

const (
	SettingGetResponseZonesSchemasAlwaysUseHTTPSIDAlwaysUseHTTPS SettingGetResponseZonesSchemasAlwaysUseHTTPSID = "always_use_https"
)

func (r SettingGetResponseZonesSchemasAlwaysUseHTTPSID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasAlwaysUseHTTPSIDAlwaysUseHTTPS:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasAlwaysUseHTTPSValue string

const (
	SettingGetResponseZonesSchemasAlwaysUseHTTPSValueOn  SettingGetResponseZonesSchemasAlwaysUseHTTPSValue = "on"
	SettingGetResponseZonesSchemasAlwaysUseHTTPSValueOff SettingGetResponseZonesSchemasAlwaysUseHTTPSValue = "off"
)

func (r SettingGetResponseZonesSchemasAlwaysUseHTTPSValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasAlwaysUseHTTPSValueOn, SettingGetResponseZonesSchemasAlwaysUseHTTPSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasAlwaysUseHTTPSEditable bool

const (
	SettingGetResponseZonesSchemasAlwaysUseHTTPSEditableTrue  SettingGetResponseZonesSchemasAlwaysUseHTTPSEditable = true
	SettingGetResponseZonesSchemasAlwaysUseHTTPSEditableFalse SettingGetResponseZonesSchemasAlwaysUseHTTPSEditable = false
)

func (r SettingGetResponseZonesSchemasAlwaysUseHTTPSEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasAlwaysUseHTTPSEditableTrue, SettingGetResponseZonesSchemasAlwaysUseHTTPSEditableFalse:
		return true
	}
	return false
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingGetResponseZonesSchemasAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasAutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasAutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasAutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasAutomaticHTTPSRewritesJSON `json:"-"`
}

// settingGetResponseZonesSchemasAutomaticHTTPSRewritesJSON contains the JSON
// metadata for the struct [SettingGetResponseZonesSchemasAutomaticHTTPSRewrites]
type settingGetResponseZonesSchemasAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasAutomaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasAutomaticHTTPSRewrites) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasAutomaticHTTPSRewritesID string

const (
	SettingGetResponseZonesSchemasAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites SettingGetResponseZonesSchemasAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

func (r SettingGetResponseZonesSchemasAutomaticHTTPSRewritesID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasAutomaticHTTPSRewritesValue string

const (
	SettingGetResponseZonesSchemasAutomaticHTTPSRewritesValueOn  SettingGetResponseZonesSchemasAutomaticHTTPSRewritesValue = "on"
	SettingGetResponseZonesSchemasAutomaticHTTPSRewritesValueOff SettingGetResponseZonesSchemasAutomaticHTTPSRewritesValue = "off"
)

func (r SettingGetResponseZonesSchemasAutomaticHTTPSRewritesValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasAutomaticHTTPSRewritesValueOn, SettingGetResponseZonesSchemasAutomaticHTTPSRewritesValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasAutomaticHTTPSRewritesEditable bool

const (
	SettingGetResponseZonesSchemasAutomaticHTTPSRewritesEditableTrue  SettingGetResponseZonesSchemasAutomaticHTTPSRewritesEditable = true
	SettingGetResponseZonesSchemasAutomaticHTTPSRewritesEditableFalse SettingGetResponseZonesSchemasAutomaticHTTPSRewritesEditable = false
)

func (r SettingGetResponseZonesSchemasAutomaticHTTPSRewritesEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasAutomaticHTTPSRewritesEditableTrue, SettingGetResponseZonesSchemasAutomaticHTTPSRewritesEditableFalse:
		return true
	}
	return false
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingGetResponseZonesSchemasBrowserCacheTTL struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasBrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasBrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasBrowserCacheTTLJSON `json:"-"`
}

// settingGetResponseZonesSchemasBrowserCacheTTLJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesSchemasBrowserCacheTTL]
type settingGetResponseZonesSchemasBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasBrowserCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasBrowserCacheTTL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasBrowserCacheTTLID string

const (
	SettingGetResponseZonesSchemasBrowserCacheTTLIDBrowserCacheTTL SettingGetResponseZonesSchemasBrowserCacheTTLID = "browser_cache_ttl"
)

func (r SettingGetResponseZonesSchemasBrowserCacheTTLID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasBrowserCacheTTLIDBrowserCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasBrowserCacheTTLValue float64

const (
	SettingGetResponseZonesSchemasBrowserCacheTTLValue0        SettingGetResponseZonesSchemasBrowserCacheTTLValue = 0
	SettingGetResponseZonesSchemasBrowserCacheTTLValue30       SettingGetResponseZonesSchemasBrowserCacheTTLValue = 30
	SettingGetResponseZonesSchemasBrowserCacheTTLValue60       SettingGetResponseZonesSchemasBrowserCacheTTLValue = 60
	SettingGetResponseZonesSchemasBrowserCacheTTLValue120      SettingGetResponseZonesSchemasBrowserCacheTTLValue = 120
	SettingGetResponseZonesSchemasBrowserCacheTTLValue300      SettingGetResponseZonesSchemasBrowserCacheTTLValue = 300
	SettingGetResponseZonesSchemasBrowserCacheTTLValue1200     SettingGetResponseZonesSchemasBrowserCacheTTLValue = 1200
	SettingGetResponseZonesSchemasBrowserCacheTTLValue1800     SettingGetResponseZonesSchemasBrowserCacheTTLValue = 1800
	SettingGetResponseZonesSchemasBrowserCacheTTLValue3600     SettingGetResponseZonesSchemasBrowserCacheTTLValue = 3600
	SettingGetResponseZonesSchemasBrowserCacheTTLValue7200     SettingGetResponseZonesSchemasBrowserCacheTTLValue = 7200
	SettingGetResponseZonesSchemasBrowserCacheTTLValue10800    SettingGetResponseZonesSchemasBrowserCacheTTLValue = 10800
	SettingGetResponseZonesSchemasBrowserCacheTTLValue14400    SettingGetResponseZonesSchemasBrowserCacheTTLValue = 14400
	SettingGetResponseZonesSchemasBrowserCacheTTLValue18000    SettingGetResponseZonesSchemasBrowserCacheTTLValue = 18000
	SettingGetResponseZonesSchemasBrowserCacheTTLValue28800    SettingGetResponseZonesSchemasBrowserCacheTTLValue = 28800
	SettingGetResponseZonesSchemasBrowserCacheTTLValue43200    SettingGetResponseZonesSchemasBrowserCacheTTLValue = 43200
	SettingGetResponseZonesSchemasBrowserCacheTTLValue57600    SettingGetResponseZonesSchemasBrowserCacheTTLValue = 57600
	SettingGetResponseZonesSchemasBrowserCacheTTLValue72000    SettingGetResponseZonesSchemasBrowserCacheTTLValue = 72000
	SettingGetResponseZonesSchemasBrowserCacheTTLValue86400    SettingGetResponseZonesSchemasBrowserCacheTTLValue = 86400
	SettingGetResponseZonesSchemasBrowserCacheTTLValue172800   SettingGetResponseZonesSchemasBrowserCacheTTLValue = 172800
	SettingGetResponseZonesSchemasBrowserCacheTTLValue259200   SettingGetResponseZonesSchemasBrowserCacheTTLValue = 259200
	SettingGetResponseZonesSchemasBrowserCacheTTLValue345600   SettingGetResponseZonesSchemasBrowserCacheTTLValue = 345600
	SettingGetResponseZonesSchemasBrowserCacheTTLValue432000   SettingGetResponseZonesSchemasBrowserCacheTTLValue = 432000
	SettingGetResponseZonesSchemasBrowserCacheTTLValue691200   SettingGetResponseZonesSchemasBrowserCacheTTLValue = 691200
	SettingGetResponseZonesSchemasBrowserCacheTTLValue1382400  SettingGetResponseZonesSchemasBrowserCacheTTLValue = 1382400
	SettingGetResponseZonesSchemasBrowserCacheTTLValue2073600  SettingGetResponseZonesSchemasBrowserCacheTTLValue = 2073600
	SettingGetResponseZonesSchemasBrowserCacheTTLValue2678400  SettingGetResponseZonesSchemasBrowserCacheTTLValue = 2678400
	SettingGetResponseZonesSchemasBrowserCacheTTLValue5356800  SettingGetResponseZonesSchemasBrowserCacheTTLValue = 5356800
	SettingGetResponseZonesSchemasBrowserCacheTTLValue16070400 SettingGetResponseZonesSchemasBrowserCacheTTLValue = 16070400
	SettingGetResponseZonesSchemasBrowserCacheTTLValue31536000 SettingGetResponseZonesSchemasBrowserCacheTTLValue = 31536000
)

func (r SettingGetResponseZonesSchemasBrowserCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasBrowserCacheTTLValue0, SettingGetResponseZonesSchemasBrowserCacheTTLValue30, SettingGetResponseZonesSchemasBrowserCacheTTLValue60, SettingGetResponseZonesSchemasBrowserCacheTTLValue120, SettingGetResponseZonesSchemasBrowserCacheTTLValue300, SettingGetResponseZonesSchemasBrowserCacheTTLValue1200, SettingGetResponseZonesSchemasBrowserCacheTTLValue1800, SettingGetResponseZonesSchemasBrowserCacheTTLValue3600, SettingGetResponseZonesSchemasBrowserCacheTTLValue7200, SettingGetResponseZonesSchemasBrowserCacheTTLValue10800, SettingGetResponseZonesSchemasBrowserCacheTTLValue14400, SettingGetResponseZonesSchemasBrowserCacheTTLValue18000, SettingGetResponseZonesSchemasBrowserCacheTTLValue28800, SettingGetResponseZonesSchemasBrowserCacheTTLValue43200, SettingGetResponseZonesSchemasBrowserCacheTTLValue57600, SettingGetResponseZonesSchemasBrowserCacheTTLValue72000, SettingGetResponseZonesSchemasBrowserCacheTTLValue86400, SettingGetResponseZonesSchemasBrowserCacheTTLValue172800, SettingGetResponseZonesSchemasBrowserCacheTTLValue259200, SettingGetResponseZonesSchemasBrowserCacheTTLValue345600, SettingGetResponseZonesSchemasBrowserCacheTTLValue432000, SettingGetResponseZonesSchemasBrowserCacheTTLValue691200, SettingGetResponseZonesSchemasBrowserCacheTTLValue1382400, SettingGetResponseZonesSchemasBrowserCacheTTLValue2073600, SettingGetResponseZonesSchemasBrowserCacheTTLValue2678400, SettingGetResponseZonesSchemasBrowserCacheTTLValue5356800, SettingGetResponseZonesSchemasBrowserCacheTTLValue16070400, SettingGetResponseZonesSchemasBrowserCacheTTLValue31536000:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasBrowserCacheTTLEditable bool

const (
	SettingGetResponseZonesSchemasBrowserCacheTTLEditableTrue  SettingGetResponseZonesSchemasBrowserCacheTTLEditable = true
	SettingGetResponseZonesSchemasBrowserCacheTTLEditableFalse SettingGetResponseZonesSchemasBrowserCacheTTLEditable = false
)

func (r SettingGetResponseZonesSchemasBrowserCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasBrowserCacheTTLEditableTrue, SettingGetResponseZonesSchemasBrowserCacheTTLEditableFalse:
		return true
	}
	return false
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type SettingGetResponseZonesSchemasBrowserCheck struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasBrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasBrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasBrowserCheckJSON `json:"-"`
}

// settingGetResponseZonesSchemasBrowserCheckJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesSchemasBrowserCheck]
type settingGetResponseZonesSchemasBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasBrowserCheckJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasBrowserCheck) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasBrowserCheckID string

const (
	SettingGetResponseZonesSchemasBrowserCheckIDBrowserCheck SettingGetResponseZonesSchemasBrowserCheckID = "browser_check"
)

func (r SettingGetResponseZonesSchemasBrowserCheckID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasBrowserCheckIDBrowserCheck:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasBrowserCheckValue string

const (
	SettingGetResponseZonesSchemasBrowserCheckValueOn  SettingGetResponseZonesSchemasBrowserCheckValue = "on"
	SettingGetResponseZonesSchemasBrowserCheckValueOff SettingGetResponseZonesSchemasBrowserCheckValue = "off"
)

func (r SettingGetResponseZonesSchemasBrowserCheckValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasBrowserCheckValueOn, SettingGetResponseZonesSchemasBrowserCheckValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasBrowserCheckEditable bool

const (
	SettingGetResponseZonesSchemasBrowserCheckEditableTrue  SettingGetResponseZonesSchemasBrowserCheckEditable = true
	SettingGetResponseZonesSchemasBrowserCheckEditableFalse SettingGetResponseZonesSchemasBrowserCheckEditable = false
)

func (r SettingGetResponseZonesSchemasBrowserCheckEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasBrowserCheckEditableTrue, SettingGetResponseZonesSchemasBrowserCheckEditableFalse:
		return true
	}
	return false
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type SettingGetResponseZonesSchemasCacheLevel struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasCacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasCacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasCacheLevelJSON `json:"-"`
}

// settingGetResponseZonesSchemasCacheLevelJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesSchemasCacheLevel]
type settingGetResponseZonesSchemasCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasCacheLevelJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasCacheLevel) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasCacheLevelID string

const (
	SettingGetResponseZonesSchemasCacheLevelIDCacheLevel SettingGetResponseZonesSchemasCacheLevelID = "cache_level"
)

func (r SettingGetResponseZonesSchemasCacheLevelID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasCacheLevelIDCacheLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasCacheLevelValue string

const (
	SettingGetResponseZonesSchemasCacheLevelValueAggressive SettingGetResponseZonesSchemasCacheLevelValue = "aggressive"
	SettingGetResponseZonesSchemasCacheLevelValueBasic      SettingGetResponseZonesSchemasCacheLevelValue = "basic"
	SettingGetResponseZonesSchemasCacheLevelValueSimplified SettingGetResponseZonesSchemasCacheLevelValue = "simplified"
)

func (r SettingGetResponseZonesSchemasCacheLevelValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasCacheLevelValueAggressive, SettingGetResponseZonesSchemasCacheLevelValueBasic, SettingGetResponseZonesSchemasCacheLevelValueSimplified:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasCacheLevelEditable bool

const (
	SettingGetResponseZonesSchemasCacheLevelEditableTrue  SettingGetResponseZonesSchemasCacheLevelEditable = true
	SettingGetResponseZonesSchemasCacheLevelEditableFalse SettingGetResponseZonesSchemasCacheLevelEditable = false
)

func (r SettingGetResponseZonesSchemasCacheLevelEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasCacheLevelEditableTrue, SettingGetResponseZonesSchemasCacheLevelEditableFalse:
		return true
	}
	return false
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
type SettingGetResponseZonesSchemasEdgeCacheTTL struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasEdgeCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasEdgeCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasEdgeCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasEdgeCacheTTLJSON `json:"-"`
}

// settingGetResponseZonesSchemasEdgeCacheTTLJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesSchemasEdgeCacheTTL]
type settingGetResponseZonesSchemasEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasEdgeCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasEdgeCacheTTL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasEdgeCacheTTLID string

const (
	SettingGetResponseZonesSchemasEdgeCacheTTLIDEdgeCacheTTL SettingGetResponseZonesSchemasEdgeCacheTTLID = "edge_cache_ttl"
)

func (r SettingGetResponseZonesSchemasEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasEdgeCacheTTLValue float64

const (
	SettingGetResponseZonesSchemasEdgeCacheTTLValue30     SettingGetResponseZonesSchemasEdgeCacheTTLValue = 30
	SettingGetResponseZonesSchemasEdgeCacheTTLValue60     SettingGetResponseZonesSchemasEdgeCacheTTLValue = 60
	SettingGetResponseZonesSchemasEdgeCacheTTLValue300    SettingGetResponseZonesSchemasEdgeCacheTTLValue = 300
	SettingGetResponseZonesSchemasEdgeCacheTTLValue1200   SettingGetResponseZonesSchemasEdgeCacheTTLValue = 1200
	SettingGetResponseZonesSchemasEdgeCacheTTLValue1800   SettingGetResponseZonesSchemasEdgeCacheTTLValue = 1800
	SettingGetResponseZonesSchemasEdgeCacheTTLValue3600   SettingGetResponseZonesSchemasEdgeCacheTTLValue = 3600
	SettingGetResponseZonesSchemasEdgeCacheTTLValue7200   SettingGetResponseZonesSchemasEdgeCacheTTLValue = 7200
	SettingGetResponseZonesSchemasEdgeCacheTTLValue10800  SettingGetResponseZonesSchemasEdgeCacheTTLValue = 10800
	SettingGetResponseZonesSchemasEdgeCacheTTLValue14400  SettingGetResponseZonesSchemasEdgeCacheTTLValue = 14400
	SettingGetResponseZonesSchemasEdgeCacheTTLValue18000  SettingGetResponseZonesSchemasEdgeCacheTTLValue = 18000
	SettingGetResponseZonesSchemasEdgeCacheTTLValue28800  SettingGetResponseZonesSchemasEdgeCacheTTLValue = 28800
	SettingGetResponseZonesSchemasEdgeCacheTTLValue43200  SettingGetResponseZonesSchemasEdgeCacheTTLValue = 43200
	SettingGetResponseZonesSchemasEdgeCacheTTLValue57600  SettingGetResponseZonesSchemasEdgeCacheTTLValue = 57600
	SettingGetResponseZonesSchemasEdgeCacheTTLValue72000  SettingGetResponseZonesSchemasEdgeCacheTTLValue = 72000
	SettingGetResponseZonesSchemasEdgeCacheTTLValue86400  SettingGetResponseZonesSchemasEdgeCacheTTLValue = 86400
	SettingGetResponseZonesSchemasEdgeCacheTTLValue172800 SettingGetResponseZonesSchemasEdgeCacheTTLValue = 172800
	SettingGetResponseZonesSchemasEdgeCacheTTLValue259200 SettingGetResponseZonesSchemasEdgeCacheTTLValue = 259200
	SettingGetResponseZonesSchemasEdgeCacheTTLValue345600 SettingGetResponseZonesSchemasEdgeCacheTTLValue = 345600
	SettingGetResponseZonesSchemasEdgeCacheTTLValue432000 SettingGetResponseZonesSchemasEdgeCacheTTLValue = 432000
	SettingGetResponseZonesSchemasEdgeCacheTTLValue518400 SettingGetResponseZonesSchemasEdgeCacheTTLValue = 518400
	SettingGetResponseZonesSchemasEdgeCacheTTLValue604800 SettingGetResponseZonesSchemasEdgeCacheTTLValue = 604800
)

func (r SettingGetResponseZonesSchemasEdgeCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasEdgeCacheTTLValue30, SettingGetResponseZonesSchemasEdgeCacheTTLValue60, SettingGetResponseZonesSchemasEdgeCacheTTLValue300, SettingGetResponseZonesSchemasEdgeCacheTTLValue1200, SettingGetResponseZonesSchemasEdgeCacheTTLValue1800, SettingGetResponseZonesSchemasEdgeCacheTTLValue3600, SettingGetResponseZonesSchemasEdgeCacheTTLValue7200, SettingGetResponseZonesSchemasEdgeCacheTTLValue10800, SettingGetResponseZonesSchemasEdgeCacheTTLValue14400, SettingGetResponseZonesSchemasEdgeCacheTTLValue18000, SettingGetResponseZonesSchemasEdgeCacheTTLValue28800, SettingGetResponseZonesSchemasEdgeCacheTTLValue43200, SettingGetResponseZonesSchemasEdgeCacheTTLValue57600, SettingGetResponseZonesSchemasEdgeCacheTTLValue72000, SettingGetResponseZonesSchemasEdgeCacheTTLValue86400, SettingGetResponseZonesSchemasEdgeCacheTTLValue172800, SettingGetResponseZonesSchemasEdgeCacheTTLValue259200, SettingGetResponseZonesSchemasEdgeCacheTTLValue345600, SettingGetResponseZonesSchemasEdgeCacheTTLValue432000, SettingGetResponseZonesSchemasEdgeCacheTTLValue518400, SettingGetResponseZonesSchemasEdgeCacheTTLValue604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasEdgeCacheTTLEditable bool

const (
	SettingGetResponseZonesSchemasEdgeCacheTTLEditableTrue  SettingGetResponseZonesSchemasEdgeCacheTTLEditable = true
	SettingGetResponseZonesSchemasEdgeCacheTTLEditableFalse SettingGetResponseZonesSchemasEdgeCacheTTLEditable = false
)

func (r SettingGetResponseZonesSchemasEdgeCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasEdgeCacheTTLEditableTrue, SettingGetResponseZonesSchemasEdgeCacheTTLEditableFalse:
		return true
	}
	return false
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingGetResponseZonesSchemasEmailObfuscation struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasEmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasEmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasEmailObfuscationJSON `json:"-"`
}

// settingGetResponseZonesSchemasEmailObfuscationJSON contains the JSON metadata
// for the struct [SettingGetResponseZonesSchemasEmailObfuscation]
type settingGetResponseZonesSchemasEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasEmailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasEmailObfuscation) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasEmailObfuscationID string

const (
	SettingGetResponseZonesSchemasEmailObfuscationIDEmailObfuscation SettingGetResponseZonesSchemasEmailObfuscationID = "email_obfuscation"
)

func (r SettingGetResponseZonesSchemasEmailObfuscationID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasEmailObfuscationIDEmailObfuscation:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasEmailObfuscationValue string

const (
	SettingGetResponseZonesSchemasEmailObfuscationValueOn  SettingGetResponseZonesSchemasEmailObfuscationValue = "on"
	SettingGetResponseZonesSchemasEmailObfuscationValueOff SettingGetResponseZonesSchemasEmailObfuscationValue = "off"
)

func (r SettingGetResponseZonesSchemasEmailObfuscationValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasEmailObfuscationValueOn, SettingGetResponseZonesSchemasEmailObfuscationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasEmailObfuscationEditable bool

const (
	SettingGetResponseZonesSchemasEmailObfuscationEditableTrue  SettingGetResponseZonesSchemasEmailObfuscationEditable = true
	SettingGetResponseZonesSchemasEmailObfuscationEditableFalse SettingGetResponseZonesSchemasEmailObfuscationEditable = false
)

func (r SettingGetResponseZonesSchemasEmailObfuscationEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasEmailObfuscationEditableTrue, SettingGetResponseZonesSchemasEmailObfuscationEditableFalse:
		return true
	}
	return false
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingGetResponseZonesSchemasIPGeolocation struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasIPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasIPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasIPGeolocationJSON `json:"-"`
}

// settingGetResponseZonesSchemasIPGeolocationJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesSchemasIPGeolocation]
type settingGetResponseZonesSchemasIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasIPGeolocationJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasIPGeolocation) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasIPGeolocationID string

const (
	SettingGetResponseZonesSchemasIPGeolocationIDIPGeolocation SettingGetResponseZonesSchemasIPGeolocationID = "ip_geolocation"
)

func (r SettingGetResponseZonesSchemasIPGeolocationID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasIPGeolocationIDIPGeolocation:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasIPGeolocationValue string

const (
	SettingGetResponseZonesSchemasIPGeolocationValueOn  SettingGetResponseZonesSchemasIPGeolocationValue = "on"
	SettingGetResponseZonesSchemasIPGeolocationValueOff SettingGetResponseZonesSchemasIPGeolocationValue = "off"
)

func (r SettingGetResponseZonesSchemasIPGeolocationValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasIPGeolocationValueOn, SettingGetResponseZonesSchemasIPGeolocationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasIPGeolocationEditable bool

const (
	SettingGetResponseZonesSchemasIPGeolocationEditableTrue  SettingGetResponseZonesSchemasIPGeolocationEditable = true
	SettingGetResponseZonesSchemasIPGeolocationEditableFalse SettingGetResponseZonesSchemasIPGeolocationEditable = false
)

func (r SettingGetResponseZonesSchemasIPGeolocationEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasIPGeolocationEditableTrue, SettingGetResponseZonesSchemasIPGeolocationEditableFalse:
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

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type SettingGetResponseZonesSchemasMirage struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasMirageID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasMirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasMirageJSON `json:"-"`
}

// settingGetResponseZonesSchemasMirageJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesSchemasMirage]
type settingGetResponseZonesSchemasMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasMirageJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasMirage) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasMirageID string

const (
	SettingGetResponseZonesSchemasMirageIDMirage SettingGetResponseZonesSchemasMirageID = "mirage"
)

func (r SettingGetResponseZonesSchemasMirageID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasMirageIDMirage:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasMirageValue string

const (
	SettingGetResponseZonesSchemasMirageValueOn  SettingGetResponseZonesSchemasMirageValue = "on"
	SettingGetResponseZonesSchemasMirageValueOff SettingGetResponseZonesSchemasMirageValue = "off"
)

func (r SettingGetResponseZonesSchemasMirageValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasMirageValueOn, SettingGetResponseZonesSchemasMirageValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasMirageEditable bool

const (
	SettingGetResponseZonesSchemasMirageEditableTrue  SettingGetResponseZonesSchemasMirageEditable = true
	SettingGetResponseZonesSchemasMirageEditableFalse SettingGetResponseZonesSchemasMirageEditable = false
)

func (r SettingGetResponseZonesSchemasMirageEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasMirageEditableTrue, SettingGetResponseZonesSchemasMirageEditableFalse:
		return true
	}
	return false
}

// Enables the Opportunistic Encryption feature for a zone.
type SettingGetResponseZonesSchemasOpportunisticEncryption struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasOpportunisticEncryptionJSON `json:"-"`
}

// settingGetResponseZonesSchemasOpportunisticEncryptionJSON contains the JSON
// metadata for the struct [SettingGetResponseZonesSchemasOpportunisticEncryption]
type settingGetResponseZonesSchemasOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasOpportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasOpportunisticEncryption) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasOpportunisticEncryptionID string

const (
	SettingGetResponseZonesSchemasOpportunisticEncryptionIDOpportunisticEncryption SettingGetResponseZonesSchemasOpportunisticEncryptionID = "opportunistic_encryption"
)

func (r SettingGetResponseZonesSchemasOpportunisticEncryptionID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasOpportunisticEncryptionIDOpportunisticEncryption:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasOpportunisticEncryptionValue string

const (
	SettingGetResponseZonesSchemasOpportunisticEncryptionValueOn  SettingGetResponseZonesSchemasOpportunisticEncryptionValue = "on"
	SettingGetResponseZonesSchemasOpportunisticEncryptionValueOff SettingGetResponseZonesSchemasOpportunisticEncryptionValue = "off"
)

func (r SettingGetResponseZonesSchemasOpportunisticEncryptionValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasOpportunisticEncryptionValueOn, SettingGetResponseZonesSchemasOpportunisticEncryptionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasOpportunisticEncryptionEditable bool

const (
	SettingGetResponseZonesSchemasOpportunisticEncryptionEditableTrue  SettingGetResponseZonesSchemasOpportunisticEncryptionEditable = true
	SettingGetResponseZonesSchemasOpportunisticEncryptionEditableFalse SettingGetResponseZonesSchemasOpportunisticEncryptionEditable = false
)

func (r SettingGetResponseZonesSchemasOpportunisticEncryptionEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasOpportunisticEncryptionEditableTrue, SettingGetResponseZonesSchemasOpportunisticEncryptionEditableFalse:
		return true
	}
	return false
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingGetResponseZonesSchemasOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasOriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasOriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasOriginErrorPagePassThruJSON `json:"-"`
}

// settingGetResponseZonesSchemasOriginErrorPagePassThruJSON contains the JSON
// metadata for the struct [SettingGetResponseZonesSchemasOriginErrorPagePassThru]
type settingGetResponseZonesSchemasOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasOriginErrorPagePassThruJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasOriginErrorPagePassThru) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasOriginErrorPagePassThruID string

const (
	SettingGetResponseZonesSchemasOriginErrorPagePassThruIDOriginErrorPagePassThru SettingGetResponseZonesSchemasOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

func (r SettingGetResponseZonesSchemasOriginErrorPagePassThruID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasOriginErrorPagePassThruIDOriginErrorPagePassThru:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasOriginErrorPagePassThruValue string

const (
	SettingGetResponseZonesSchemasOriginErrorPagePassThruValueOn  SettingGetResponseZonesSchemasOriginErrorPagePassThruValue = "on"
	SettingGetResponseZonesSchemasOriginErrorPagePassThruValueOff SettingGetResponseZonesSchemasOriginErrorPagePassThruValue = "off"
)

func (r SettingGetResponseZonesSchemasOriginErrorPagePassThruValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasOriginErrorPagePassThruValueOn, SettingGetResponseZonesSchemasOriginErrorPagePassThruValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasOriginErrorPagePassThruEditable bool

const (
	SettingGetResponseZonesSchemasOriginErrorPagePassThruEditableTrue  SettingGetResponseZonesSchemasOriginErrorPagePassThruEditable = true
	SettingGetResponseZonesSchemasOriginErrorPagePassThruEditableFalse SettingGetResponseZonesSchemasOriginErrorPagePassThruEditable = false
)

func (r SettingGetResponseZonesSchemasOriginErrorPagePassThruEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasOriginErrorPagePassThruEditableTrue, SettingGetResponseZonesSchemasOriginErrorPagePassThruEditableFalse:
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
type SettingGetResponseZonesSchemasPolish struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasPolishID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasPolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasPolishJSON `json:"-"`
}

// settingGetResponseZonesSchemasPolishJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesSchemasPolish]
type settingGetResponseZonesSchemasPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasPolishJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasPolish) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasPolishID string

const (
	SettingGetResponseZonesSchemasPolishIDPolish SettingGetResponseZonesSchemasPolishID = "polish"
)

func (r SettingGetResponseZonesSchemasPolishID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasPolishIDPolish:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasPolishValue string

const (
	SettingGetResponseZonesSchemasPolishValueOff      SettingGetResponseZonesSchemasPolishValue = "off"
	SettingGetResponseZonesSchemasPolishValueLossless SettingGetResponseZonesSchemasPolishValue = "lossless"
	SettingGetResponseZonesSchemasPolishValueLossy    SettingGetResponseZonesSchemasPolishValue = "lossy"
)

func (r SettingGetResponseZonesSchemasPolishValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasPolishValueOff, SettingGetResponseZonesSchemasPolishValueLossless, SettingGetResponseZonesSchemasPolishValueLossy:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasPolishEditable bool

const (
	SettingGetResponseZonesSchemasPolishEditableTrue  SettingGetResponseZonesSchemasPolishEditable = true
	SettingGetResponseZonesSchemasPolishEditableFalse SettingGetResponseZonesSchemasPolishEditable = false
)

func (r SettingGetResponseZonesSchemasPolishEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasPolishEditableTrue, SettingGetResponseZonesSchemasPolishEditableFalse:
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

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingGetResponseZonesSchemasResponseBuffering struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasResponseBufferingID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasResponseBufferingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasResponseBufferingJSON `json:"-"`
}

// settingGetResponseZonesSchemasResponseBufferingJSON contains the JSON metadata
// for the struct [SettingGetResponseZonesSchemasResponseBuffering]
type settingGetResponseZonesSchemasResponseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasResponseBufferingJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasResponseBuffering) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasResponseBufferingID string

const (
	SettingGetResponseZonesSchemasResponseBufferingIDResponseBuffering SettingGetResponseZonesSchemasResponseBufferingID = "response_buffering"
)

func (r SettingGetResponseZonesSchemasResponseBufferingID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasResponseBufferingIDResponseBuffering:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasResponseBufferingValue string

const (
	SettingGetResponseZonesSchemasResponseBufferingValueOn  SettingGetResponseZonesSchemasResponseBufferingValue = "on"
	SettingGetResponseZonesSchemasResponseBufferingValueOff SettingGetResponseZonesSchemasResponseBufferingValue = "off"
)

func (r SettingGetResponseZonesSchemasResponseBufferingValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasResponseBufferingValueOn, SettingGetResponseZonesSchemasResponseBufferingValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasResponseBufferingEditable bool

const (
	SettingGetResponseZonesSchemasResponseBufferingEditableTrue  SettingGetResponseZonesSchemasResponseBufferingEditable = true
	SettingGetResponseZonesSchemasResponseBufferingEditableFalse SettingGetResponseZonesSchemasResponseBufferingEditable = false
)

func (r SettingGetResponseZonesSchemasResponseBufferingEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasResponseBufferingEditableTrue, SettingGetResponseZonesSchemasResponseBufferingEditableFalse:
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
type SettingGetResponseZonesSchemasRocketLoader struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasRocketLoaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasRocketLoaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasRocketLoaderJSON `json:"-"`
}

// settingGetResponseZonesSchemasRocketLoaderJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesSchemasRocketLoader]
type settingGetResponseZonesSchemasRocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasRocketLoaderJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasRocketLoader) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasRocketLoaderID string

const (
	SettingGetResponseZonesSchemasRocketLoaderIDRocketLoader SettingGetResponseZonesSchemasRocketLoaderID = "rocket_loader"
)

func (r SettingGetResponseZonesSchemasRocketLoaderID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasRocketLoaderIDRocketLoader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasRocketLoaderValue string

const (
	SettingGetResponseZonesSchemasRocketLoaderValueOn  SettingGetResponseZonesSchemasRocketLoaderValue = "on"
	SettingGetResponseZonesSchemasRocketLoaderValueOff SettingGetResponseZonesSchemasRocketLoaderValue = "off"
)

func (r SettingGetResponseZonesSchemasRocketLoaderValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasRocketLoaderValueOn, SettingGetResponseZonesSchemasRocketLoaderValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasRocketLoaderEditable bool

const (
	SettingGetResponseZonesSchemasRocketLoaderEditableTrue  SettingGetResponseZonesSchemasRocketLoaderEditable = true
	SettingGetResponseZonesSchemasRocketLoaderEditableFalse SettingGetResponseZonesSchemasRocketLoaderEditable = false
)

func (r SettingGetResponseZonesSchemasRocketLoaderEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasRocketLoaderEditableTrue, SettingGetResponseZonesSchemasRocketLoaderEditableFalse:
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

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingGetResponseZonesSchemasSecurityLevel struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasSecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasSecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasSecurityLevelJSON `json:"-"`
}

// settingGetResponseZonesSchemasSecurityLevelJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesSchemasSecurityLevel]
type settingGetResponseZonesSchemasSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasSecurityLevelJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasSecurityLevel) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasSecurityLevelID string

const (
	SettingGetResponseZonesSchemasSecurityLevelIDSecurityLevel SettingGetResponseZonesSchemasSecurityLevelID = "security_level"
)

func (r SettingGetResponseZonesSchemasSecurityLevelID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasSecurityLevelIDSecurityLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasSecurityLevelValue string

const (
	SettingGetResponseZonesSchemasSecurityLevelValueOff            SettingGetResponseZonesSchemasSecurityLevelValue = "off"
	SettingGetResponseZonesSchemasSecurityLevelValueEssentiallyOff SettingGetResponseZonesSchemasSecurityLevelValue = "essentially_off"
	SettingGetResponseZonesSchemasSecurityLevelValueLow            SettingGetResponseZonesSchemasSecurityLevelValue = "low"
	SettingGetResponseZonesSchemasSecurityLevelValueMedium         SettingGetResponseZonesSchemasSecurityLevelValue = "medium"
	SettingGetResponseZonesSchemasSecurityLevelValueHigh           SettingGetResponseZonesSchemasSecurityLevelValue = "high"
	SettingGetResponseZonesSchemasSecurityLevelValueUnderAttack    SettingGetResponseZonesSchemasSecurityLevelValue = "under_attack"
)

func (r SettingGetResponseZonesSchemasSecurityLevelValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasSecurityLevelValueOff, SettingGetResponseZonesSchemasSecurityLevelValueEssentiallyOff, SettingGetResponseZonesSchemasSecurityLevelValueLow, SettingGetResponseZonesSchemasSecurityLevelValueMedium, SettingGetResponseZonesSchemasSecurityLevelValueHigh, SettingGetResponseZonesSchemasSecurityLevelValueUnderAttack:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasSecurityLevelEditable bool

const (
	SettingGetResponseZonesSchemasSecurityLevelEditableTrue  SettingGetResponseZonesSchemasSecurityLevelEditable = true
	SettingGetResponseZonesSchemasSecurityLevelEditableFalse SettingGetResponseZonesSchemasSecurityLevelEditable = false
)

func (r SettingGetResponseZonesSchemasSecurityLevelEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasSecurityLevelEditableTrue, SettingGetResponseZonesSchemasSecurityLevelEditableFalse:
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

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type SettingGetResponseZonesSchemasSortQueryStringForCache struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasSortQueryStringForCacheID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasSortQueryStringForCacheValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasSortQueryStringForCacheJSON `json:"-"`
}

// settingGetResponseZonesSchemasSortQueryStringForCacheJSON contains the JSON
// metadata for the struct [SettingGetResponseZonesSchemasSortQueryStringForCache]
type settingGetResponseZonesSchemasSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasSortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasSortQueryStringForCache) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasSortQueryStringForCacheID string

const (
	SettingGetResponseZonesSchemasSortQueryStringForCacheIDSortQueryStringForCache SettingGetResponseZonesSchemasSortQueryStringForCacheID = "sort_query_string_for_cache"
)

func (r SettingGetResponseZonesSchemasSortQueryStringForCacheID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasSortQueryStringForCacheIDSortQueryStringForCache:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasSortQueryStringForCacheValue string

const (
	SettingGetResponseZonesSchemasSortQueryStringForCacheValueOn  SettingGetResponseZonesSchemasSortQueryStringForCacheValue = "on"
	SettingGetResponseZonesSchemasSortQueryStringForCacheValueOff SettingGetResponseZonesSchemasSortQueryStringForCacheValue = "off"
)

func (r SettingGetResponseZonesSchemasSortQueryStringForCacheValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasSortQueryStringForCacheValueOn, SettingGetResponseZonesSchemasSortQueryStringForCacheValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasSortQueryStringForCacheEditable bool

const (
	SettingGetResponseZonesSchemasSortQueryStringForCacheEditableTrue  SettingGetResponseZonesSchemasSortQueryStringForCacheEditable = true
	SettingGetResponseZonesSchemasSortQueryStringForCacheEditableFalse SettingGetResponseZonesSchemasSortQueryStringForCacheEditable = false
)

func (r SettingGetResponseZonesSchemasSortQueryStringForCacheEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasSortQueryStringForCacheEditableTrue, SettingGetResponseZonesSchemasSortQueryStringForCacheEditableFalse:
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
type SettingGetResponseZonesSchemasSSL struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasSSLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasSSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasSSLJSON `json:"-"`
}

// settingGetResponseZonesSchemasSSLJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesSchemasSSL]
type settingGetResponseZonesSchemasSSLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasSSLJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasSSL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasSSLID string

const (
	SettingGetResponseZonesSchemasSSLIDSSL SettingGetResponseZonesSchemasSSLID = "ssl"
)

func (r SettingGetResponseZonesSchemasSSLID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasSSLIDSSL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasSSLValue string

const (
	SettingGetResponseZonesSchemasSSLValueOff      SettingGetResponseZonesSchemasSSLValue = "off"
	SettingGetResponseZonesSchemasSSLValueFlexible SettingGetResponseZonesSchemasSSLValue = "flexible"
	SettingGetResponseZonesSchemasSSLValueFull     SettingGetResponseZonesSchemasSSLValue = "full"
	SettingGetResponseZonesSchemasSSLValueStrict   SettingGetResponseZonesSchemasSSLValue = "strict"
)

func (r SettingGetResponseZonesSchemasSSLValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasSSLValueOff, SettingGetResponseZonesSchemasSSLValueFlexible, SettingGetResponseZonesSchemasSSLValueFull, SettingGetResponseZonesSchemasSSLValueStrict:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasSSLEditable bool

const (
	SettingGetResponseZonesSchemasSSLEditableTrue  SettingGetResponseZonesSchemasSSLEditable = true
	SettingGetResponseZonesSchemasSSLEditableFalse SettingGetResponseZonesSchemasSSLEditable = false
)

func (r SettingGetResponseZonesSchemasSSLEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasSSLEditableTrue, SettingGetResponseZonesSchemasSSLEditableFalse:
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

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingGetResponseZonesSchemasTrueClientIPHeader struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasTrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasTrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                            `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasTrueClientIPHeaderJSON `json:"-"`
}

// settingGetResponseZonesSchemasTrueClientIPHeaderJSON contains the JSON metadata
// for the struct [SettingGetResponseZonesSchemasTrueClientIPHeader]
type settingGetResponseZonesSchemasTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasTrueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasTrueClientIPHeader) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasTrueClientIPHeaderID string

const (
	SettingGetResponseZonesSchemasTrueClientIPHeaderIDTrueClientIPHeader SettingGetResponseZonesSchemasTrueClientIPHeaderID = "true_client_ip_header"
)

func (r SettingGetResponseZonesSchemasTrueClientIPHeaderID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasTrueClientIPHeaderIDTrueClientIPHeader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasTrueClientIPHeaderValue string

const (
	SettingGetResponseZonesSchemasTrueClientIPHeaderValueOn  SettingGetResponseZonesSchemasTrueClientIPHeaderValue = "on"
	SettingGetResponseZonesSchemasTrueClientIPHeaderValueOff SettingGetResponseZonesSchemasTrueClientIPHeaderValue = "off"
)

func (r SettingGetResponseZonesSchemasTrueClientIPHeaderValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasTrueClientIPHeaderValueOn, SettingGetResponseZonesSchemasTrueClientIPHeaderValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasTrueClientIPHeaderEditable bool

const (
	SettingGetResponseZonesSchemasTrueClientIPHeaderEditableTrue  SettingGetResponseZonesSchemasTrueClientIPHeaderEditable = true
	SettingGetResponseZonesSchemasTrueClientIPHeaderEditableFalse SettingGetResponseZonesSchemasTrueClientIPHeaderEditable = false
)

func (r SettingGetResponseZonesSchemasTrueClientIPHeaderEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasTrueClientIPHeaderEditableTrue, SettingGetResponseZonesSchemasTrueClientIPHeaderEditableFalse:
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
type SettingGetResponseZonesSchemasWAF struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasWAFID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasWAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSchemasWAFJSON `json:"-"`
}

// settingGetResponseZonesSchemasWAFJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesSchemasWAF]
type settingGetResponseZonesSchemasWAFJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasWAFJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSchemasWAF) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSchemasWAFID string

const (
	SettingGetResponseZonesSchemasWAFIDWAF SettingGetResponseZonesSchemasWAFID = "waf"
)

func (r SettingGetResponseZonesSchemasWAFID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasWAFIDWAF:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSchemasWAFValue string

const (
	SettingGetResponseZonesSchemasWAFValueOn  SettingGetResponseZonesSchemasWAFValue = "on"
	SettingGetResponseZonesSchemasWAFValueOff SettingGetResponseZonesSchemasWAFValue = "off"
)

func (r SettingGetResponseZonesSchemasWAFValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasWAFValueOn, SettingGetResponseZonesSchemasWAFValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasWAFEditable bool

const (
	SettingGetResponseZonesSchemasWAFEditableTrue  SettingGetResponseZonesSchemasWAFEditable = true
	SettingGetResponseZonesSchemasWAFEditableFalse SettingGetResponseZonesSchemasWAFEditable = false
)

func (r SettingGetResponseZonesSchemasWAFEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSchemasWAFEditableTrue, SettingGetResponseZonesSchemasWAFEditableFalse:
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
	SettingGetResponseIDMirage                        SettingGetResponseID = "mirage"
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
	case SettingGetResponseID0rtt, SettingGetResponseIDAdvancedDDoS, SettingGetResponseIDAlwaysOnline, SettingGetResponseIDAlwaysUseHTTPS, SettingGetResponseIDAutomaticHTTPSRewrites, SettingGetResponseIDBrotli, SettingGetResponseIDBrowserCacheTTL, SettingGetResponseIDBrowserCheck, SettingGetResponseIDCacheLevel, SettingGetResponseIDChallengeTTL, SettingGetResponseIDCiphers, SettingGetResponseIDCNAMEFlattening, SettingGetResponseIDDevelopmentMode, SettingGetResponseIDEarlyHints, SettingGetResponseIDEdgeCacheTTL, SettingGetResponseIDEmailObfuscation, SettingGetResponseIDH2Prioritization, SettingGetResponseIDHotlinkProtection, SettingGetResponseIDHTTP2, SettingGetResponseIDHTTP3, SettingGetResponseIDImageResizing, SettingGetResponseIDIPGeolocation, SettingGetResponseIDIPV6, SettingGetResponseIDMaxUpload, SettingGetResponseIDMinTLSVersion, SettingGetResponseIDMirage, SettingGetResponseIDNEL, SettingGetResponseIDOpportunisticEncryption, SettingGetResponseIDOpportunisticOnion, SettingGetResponseIDOrangeToOrange, SettingGetResponseIDOriginErrorPagePassThru, SettingGetResponseIDPolish, SettingGetResponseIDPrefetchPreload, SettingGetResponseIDProxyReadTimeout, SettingGetResponseIDPseudoIPV4, SettingGetResponseIDReplaceInsecureJS, SettingGetResponseIDResponseBuffering, SettingGetResponseIDRocketLoader, SettingGetResponseIDAutomaticPlatformOptimization, SettingGetResponseIDSecurityHeader, SettingGetResponseIDSecurityLevel, SettingGetResponseIDServerSideExclude, SettingGetResponseIDSha1Support, SettingGetResponseIDSortQueryStringForCache, SettingGetResponseIDSSL, SettingGetResponseIDSSLRecommender, SettingGetResponseIDTLS1_2Only, SettingGetResponseIDTLS1_3, SettingGetResponseIDTLSClientAuth, SettingGetResponseIDTrueClientIPHeader, SettingGetResponseIDWAF, SettingGetResponseIDWebP, SettingGetResponseIDWebsockets:
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
	ID param.Field[SettingEditParamsBodyID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool]        `json:"enabled"`
	Value   param.Field[interface{}] `json:"value"`
}

func (r SettingEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBody) implementsZonesSettingEditParamsBodyUnion() {}

// 0-RTT session resumption enabled for this zone.
//
// Satisfied by [zones.ZeroRTTParam], [zones.AdvancedDDoSParam],
// [zones.AlwaysOnlineParam],
// [zones.SettingEditParamsBodyZonesSchemasAlwaysUseHTTPS],
// [zones.SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewrites],
// [zones.BrotliParam], [zones.SettingEditParamsBodyZonesSchemasBrowserCacheTTL],
// [zones.SettingEditParamsBodyZonesSchemasBrowserCheck],
// [zones.SettingEditParamsBodyZonesSchemasCacheLevel], [zones.ChallengeTTLParam],
// [zones.CiphersParam], [zones.SettingEditParamsBodyZonesCNAMEFlattening],
// [zones.DevelopmentModeParam], [zones.EarlyHintsParam],
// [zones.SettingEditParamsBodyZonesSchemasEdgeCacheTTL],
// [zones.SettingEditParamsBodyZonesSchemasEmailObfuscation],
// [zones.H2PrioritizationParam], [zones.HotlinkProtectionParam],
// [zones.HTTP2Param], [zones.HTTP3Param], [zones.ImageResizingParam],
// [zones.SettingEditParamsBodyZonesSchemasIPGeolocation], [zones.IPV6Param],
// [zones.SettingEditParamsBodyZonesMaxUpload], [zones.MinTLSVersionParam],
// [zones.SettingEditParamsBodyZonesSchemasMirage], [zones.NELParam],
// [zones.SettingEditParamsBodyZonesSchemasOpportunisticEncryption],
// [zones.OpportunisticOnionParam], [zones.OrangeToOrangeParam],
// [zones.SettingEditParamsBodyZonesSchemasOriginErrorPagePassThru],
// [zones.SettingEditParamsBodyZonesSchemasPolish], [zones.PrefetchPreloadParam],
// [zones.ProxyReadTimeoutParam], [zones.PseudoIPV4Param],
// [zones.SettingEditParamsBodyZonesReplaceInsecureJS],
// [zones.SettingEditParamsBodyZonesSchemasResponseBuffering],
// [zones.SettingEditParamsBodyZonesSchemasRocketLoader],
// [zones.SettingEditParamsBodyZonesSchemasAutomaticPlatformOptimization],
// [zones.SecurityHeadersParam],
// [zones.SettingEditParamsBodyZonesSchemasSecurityLevel],
// [zones.ServerSideExcludesParam], [zones.SettingEditParamsBodyZonesSha1Support],
// [zones.SettingEditParamsBodyZonesSchemasSortQueryStringForCache],
// [zones.SettingEditParamsBodyZonesSchemasSSL], [zones.SSLRecommenderParam],
// [zones.SettingEditParamsBodyZonesTLS1_2Only], [zones.TLS1_3Param],
// [zones.TLSClientAuthParam],
// [zones.SettingEditParamsBodyZonesSchemasTrueClientIPHeader],
// [zones.SettingEditParamsBodyZonesSchemasWAF], [zones.WebPParam],
// [zones.WebsocketParam], [SettingEditParamsBody].
type SettingEditParamsBodyUnion interface {
	implementsZonesSettingEditParamsBodyUnion()
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingEditParamsBodyZonesSchemasAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasAlwaysUseHTTPS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasAlwaysUseHTTPS) implementsZonesSettingEditParamsBodyUnion() {
}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSID string

const (
	SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSIDAlwaysUseHTTPS SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSID = "always_use_https"
)

func (r SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSIDAlwaysUseHTTPS:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSValue string

const (
	SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSValueOn  SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSValue = "on"
	SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSValueOff SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSValueOn, SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSEditable bool

const (
	SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSEditableTrue  SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSEditable = true
	SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSEditableFalse SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSEditable = false
)

func (r SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSEditableTrue, SettingEditParamsBodyZonesSchemasAlwaysUseHTTPSEditableFalse:
		return true
	}
	return false
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewrites) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewrites) implementsZonesSettingEditParamsBodyUnion() {
}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesID string

const (
	SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

func (r SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesValue string

const (
	SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesValueOn  SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesValue = "on"
	SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesValueOff SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesValueOn, SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesEditable bool

const (
	SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesEditableTrue  SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesEditable = true
	SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesEditableFalse SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesEditable = false
)

func (r SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesEditableTrue, SettingEditParamsBodyZonesSchemasAutomaticHTTPSRewritesEditableFalse:
		return true
	}
	return false
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingEditParamsBodyZonesSchemasBrowserCacheTTL struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasBrowserCacheTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasBrowserCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasBrowserCacheTTL) implementsZonesSettingEditParamsBodyUnion() {
}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasBrowserCacheTTLID string

const (
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLIDBrowserCacheTTL SettingEditParamsBodyZonesSchemasBrowserCacheTTLID = "browser_cache_ttl"
)

func (r SettingEditParamsBodyZonesSchemasBrowserCacheTTLID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasBrowserCacheTTLIDBrowserCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue float64

const (
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue0        SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 0
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue30       SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 30
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue60       SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 60
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue120      SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 120
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue300      SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 300
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue1200     SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 1200
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue1800     SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 1800
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue3600     SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 3600
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue7200     SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 7200
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue10800    SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 10800
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue14400    SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 14400
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue18000    SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 18000
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue28800    SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 28800
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue43200    SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 43200
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue57600    SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 57600
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue72000    SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 72000
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue86400    SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 86400
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue172800   SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 172800
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue259200   SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 259200
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue345600   SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 345600
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue432000   SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 432000
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue691200   SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 691200
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue1382400  SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 1382400
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue2073600  SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 2073600
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue2678400  SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 2678400
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue5356800  SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 5356800
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue16070400 SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 16070400
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue31536000 SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue = 31536000
)

func (r SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue0, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue30, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue60, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue120, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue300, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue1200, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue1800, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue3600, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue7200, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue10800, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue14400, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue18000, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue28800, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue43200, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue57600, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue72000, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue86400, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue172800, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue259200, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue345600, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue432000, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue691200, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue1382400, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue2073600, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue2678400, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue5356800, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue16070400, SettingEditParamsBodyZonesSchemasBrowserCacheTTLValue31536000:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasBrowserCacheTTLEditable bool

const (
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLEditableTrue  SettingEditParamsBodyZonesSchemasBrowserCacheTTLEditable = true
	SettingEditParamsBodyZonesSchemasBrowserCacheTTLEditableFalse SettingEditParamsBodyZonesSchemasBrowserCacheTTLEditable = false
)

func (r SettingEditParamsBodyZonesSchemasBrowserCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasBrowserCacheTTLEditableTrue, SettingEditParamsBodyZonesSchemasBrowserCacheTTLEditableFalse:
		return true
	}
	return false
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type SettingEditParamsBodyZonesSchemasBrowserCheck struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasBrowserCheckID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasBrowserCheckValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasBrowserCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasBrowserCheck) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasBrowserCheckID string

const (
	SettingEditParamsBodyZonesSchemasBrowserCheckIDBrowserCheck SettingEditParamsBodyZonesSchemasBrowserCheckID = "browser_check"
)

func (r SettingEditParamsBodyZonesSchemasBrowserCheckID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasBrowserCheckIDBrowserCheck:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasBrowserCheckValue string

const (
	SettingEditParamsBodyZonesSchemasBrowserCheckValueOn  SettingEditParamsBodyZonesSchemasBrowserCheckValue = "on"
	SettingEditParamsBodyZonesSchemasBrowserCheckValueOff SettingEditParamsBodyZonesSchemasBrowserCheckValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasBrowserCheckValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasBrowserCheckValueOn, SettingEditParamsBodyZonesSchemasBrowserCheckValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasBrowserCheckEditable bool

const (
	SettingEditParamsBodyZonesSchemasBrowserCheckEditableTrue  SettingEditParamsBodyZonesSchemasBrowserCheckEditable = true
	SettingEditParamsBodyZonesSchemasBrowserCheckEditableFalse SettingEditParamsBodyZonesSchemasBrowserCheckEditable = false
)

func (r SettingEditParamsBodyZonesSchemasBrowserCheckEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasBrowserCheckEditableTrue, SettingEditParamsBodyZonesSchemasBrowserCheckEditableFalse:
		return true
	}
	return false
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type SettingEditParamsBodyZonesSchemasCacheLevel struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasCacheLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasCacheLevelValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasCacheLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasCacheLevel) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasCacheLevelID string

const (
	SettingEditParamsBodyZonesSchemasCacheLevelIDCacheLevel SettingEditParamsBodyZonesSchemasCacheLevelID = "cache_level"
)

func (r SettingEditParamsBodyZonesSchemasCacheLevelID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasCacheLevelIDCacheLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasCacheLevelValue string

const (
	SettingEditParamsBodyZonesSchemasCacheLevelValueAggressive SettingEditParamsBodyZonesSchemasCacheLevelValue = "aggressive"
	SettingEditParamsBodyZonesSchemasCacheLevelValueBasic      SettingEditParamsBodyZonesSchemasCacheLevelValue = "basic"
	SettingEditParamsBodyZonesSchemasCacheLevelValueSimplified SettingEditParamsBodyZonesSchemasCacheLevelValue = "simplified"
)

func (r SettingEditParamsBodyZonesSchemasCacheLevelValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasCacheLevelValueAggressive, SettingEditParamsBodyZonesSchemasCacheLevelValueBasic, SettingEditParamsBodyZonesSchemasCacheLevelValueSimplified:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasCacheLevelEditable bool

const (
	SettingEditParamsBodyZonesSchemasCacheLevelEditableTrue  SettingEditParamsBodyZonesSchemasCacheLevelEditable = true
	SettingEditParamsBodyZonesSchemasCacheLevelEditableFalse SettingEditParamsBodyZonesSchemasCacheLevelEditable = false
)

func (r SettingEditParamsBodyZonesSchemasCacheLevelEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasCacheLevelEditableTrue, SettingEditParamsBodyZonesSchemasCacheLevelEditableFalse:
		return true
	}
	return false
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
type SettingEditParamsBodyZonesSchemasEdgeCacheTTL struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasEdgeCacheTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasEdgeCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasEdgeCacheTTL) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasEdgeCacheTTLID string

const (
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLIDEdgeCacheTTL SettingEditParamsBodyZonesSchemasEdgeCacheTTLID = "edge_cache_ttl"
)

func (r SettingEditParamsBodyZonesSchemasEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue float64

const (
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue30     SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 30
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue60     SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 60
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue300    SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 300
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue1200   SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 1200
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue1800   SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 1800
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue3600   SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 3600
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue7200   SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 7200
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue10800  SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 10800
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue14400  SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 14400
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue18000  SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 18000
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue28800  SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 28800
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue43200  SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 43200
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue57600  SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 57600
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue72000  SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 72000
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue86400  SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 86400
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue172800 SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 172800
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue259200 SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 259200
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue345600 SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 345600
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue432000 SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 432000
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue518400 SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 518400
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue604800 SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue = 604800
)

func (r SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue30, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue60, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue300, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue1200, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue1800, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue3600, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue7200, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue10800, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue14400, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue18000, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue28800, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue43200, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue57600, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue72000, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue86400, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue172800, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue259200, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue345600, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue432000, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue518400, SettingEditParamsBodyZonesSchemasEdgeCacheTTLValue604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasEdgeCacheTTLEditable bool

const (
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLEditableTrue  SettingEditParamsBodyZonesSchemasEdgeCacheTTLEditable = true
	SettingEditParamsBodyZonesSchemasEdgeCacheTTLEditableFalse SettingEditParamsBodyZonesSchemasEdgeCacheTTLEditable = false
)

func (r SettingEditParamsBodyZonesSchemasEdgeCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasEdgeCacheTTLEditableTrue, SettingEditParamsBodyZonesSchemasEdgeCacheTTLEditableFalse:
		return true
	}
	return false
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingEditParamsBodyZonesSchemasEmailObfuscation struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasEmailObfuscationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasEmailObfuscationValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasEmailObfuscation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasEmailObfuscation) implementsZonesSettingEditParamsBodyUnion() {
}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasEmailObfuscationID string

const (
	SettingEditParamsBodyZonesSchemasEmailObfuscationIDEmailObfuscation SettingEditParamsBodyZonesSchemasEmailObfuscationID = "email_obfuscation"
)

func (r SettingEditParamsBodyZonesSchemasEmailObfuscationID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasEmailObfuscationIDEmailObfuscation:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasEmailObfuscationValue string

const (
	SettingEditParamsBodyZonesSchemasEmailObfuscationValueOn  SettingEditParamsBodyZonesSchemasEmailObfuscationValue = "on"
	SettingEditParamsBodyZonesSchemasEmailObfuscationValueOff SettingEditParamsBodyZonesSchemasEmailObfuscationValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasEmailObfuscationValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasEmailObfuscationValueOn, SettingEditParamsBodyZonesSchemasEmailObfuscationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasEmailObfuscationEditable bool

const (
	SettingEditParamsBodyZonesSchemasEmailObfuscationEditableTrue  SettingEditParamsBodyZonesSchemasEmailObfuscationEditable = true
	SettingEditParamsBodyZonesSchemasEmailObfuscationEditableFalse SettingEditParamsBodyZonesSchemasEmailObfuscationEditable = false
)

func (r SettingEditParamsBodyZonesSchemasEmailObfuscationEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasEmailObfuscationEditableTrue, SettingEditParamsBodyZonesSchemasEmailObfuscationEditableFalse:
		return true
	}
	return false
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingEditParamsBodyZonesSchemasIPGeolocation struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasIPGeolocationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasIPGeolocationValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasIPGeolocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasIPGeolocation) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasIPGeolocationID string

const (
	SettingEditParamsBodyZonesSchemasIPGeolocationIDIPGeolocation SettingEditParamsBodyZonesSchemasIPGeolocationID = "ip_geolocation"
)

func (r SettingEditParamsBodyZonesSchemasIPGeolocationID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasIPGeolocationIDIPGeolocation:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasIPGeolocationValue string

const (
	SettingEditParamsBodyZonesSchemasIPGeolocationValueOn  SettingEditParamsBodyZonesSchemasIPGeolocationValue = "on"
	SettingEditParamsBodyZonesSchemasIPGeolocationValueOff SettingEditParamsBodyZonesSchemasIPGeolocationValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasIPGeolocationValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasIPGeolocationValueOn, SettingEditParamsBodyZonesSchemasIPGeolocationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasIPGeolocationEditable bool

const (
	SettingEditParamsBodyZonesSchemasIPGeolocationEditableTrue  SettingEditParamsBodyZonesSchemasIPGeolocationEditable = true
	SettingEditParamsBodyZonesSchemasIPGeolocationEditableFalse SettingEditParamsBodyZonesSchemasIPGeolocationEditable = false
)

func (r SettingEditParamsBodyZonesSchemasIPGeolocationEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasIPGeolocationEditableTrue, SettingEditParamsBodyZonesSchemasIPGeolocationEditableFalse:
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

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type SettingEditParamsBodyZonesSchemasMirage struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasMirageID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasMirageValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasMirage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasMirage) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasMirageID string

const (
	SettingEditParamsBodyZonesSchemasMirageIDMirage SettingEditParamsBodyZonesSchemasMirageID = "mirage"
)

func (r SettingEditParamsBodyZonesSchemasMirageID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasMirageIDMirage:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasMirageValue string

const (
	SettingEditParamsBodyZonesSchemasMirageValueOn  SettingEditParamsBodyZonesSchemasMirageValue = "on"
	SettingEditParamsBodyZonesSchemasMirageValueOff SettingEditParamsBodyZonesSchemasMirageValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasMirageValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasMirageValueOn, SettingEditParamsBodyZonesSchemasMirageValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasMirageEditable bool

const (
	SettingEditParamsBodyZonesSchemasMirageEditableTrue  SettingEditParamsBodyZonesSchemasMirageEditable = true
	SettingEditParamsBodyZonesSchemasMirageEditableFalse SettingEditParamsBodyZonesSchemasMirageEditable = false
)

func (r SettingEditParamsBodyZonesSchemasMirageEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasMirageEditableTrue, SettingEditParamsBodyZonesSchemasMirageEditableFalse:
		return true
	}
	return false
}

// Enables the Opportunistic Encryption feature for a zone.
type SettingEditParamsBodyZonesSchemasOpportunisticEncryption struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasOpportunisticEncryptionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasOpportunisticEncryptionValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasOpportunisticEncryption) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasOpportunisticEncryption) implementsZonesSettingEditParamsBodyUnion() {
}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasOpportunisticEncryptionID string

const (
	SettingEditParamsBodyZonesSchemasOpportunisticEncryptionIDOpportunisticEncryption SettingEditParamsBodyZonesSchemasOpportunisticEncryptionID = "opportunistic_encryption"
)

func (r SettingEditParamsBodyZonesSchemasOpportunisticEncryptionID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasOpportunisticEncryptionIDOpportunisticEncryption:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasOpportunisticEncryptionValue string

const (
	SettingEditParamsBodyZonesSchemasOpportunisticEncryptionValueOn  SettingEditParamsBodyZonesSchemasOpportunisticEncryptionValue = "on"
	SettingEditParamsBodyZonesSchemasOpportunisticEncryptionValueOff SettingEditParamsBodyZonesSchemasOpportunisticEncryptionValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasOpportunisticEncryptionValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasOpportunisticEncryptionValueOn, SettingEditParamsBodyZonesSchemasOpportunisticEncryptionValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasOpportunisticEncryptionEditable bool

const (
	SettingEditParamsBodyZonesSchemasOpportunisticEncryptionEditableTrue  SettingEditParamsBodyZonesSchemasOpportunisticEncryptionEditable = true
	SettingEditParamsBodyZonesSchemasOpportunisticEncryptionEditableFalse SettingEditParamsBodyZonesSchemasOpportunisticEncryptionEditable = false
)

func (r SettingEditParamsBodyZonesSchemasOpportunisticEncryptionEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasOpportunisticEncryptionEditableTrue, SettingEditParamsBodyZonesSchemasOpportunisticEncryptionEditableFalse:
		return true
	}
	return false
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingEditParamsBodyZonesSchemasOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasOriginErrorPagePassThru) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasOriginErrorPagePassThru) implementsZonesSettingEditParamsBodyUnion() {
}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruID string

const (
	SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruIDOriginErrorPagePassThru SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

func (r SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruIDOriginErrorPagePassThru:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruValue string

const (
	SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruValueOn  SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruValue = "on"
	SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruValueOff SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruValueOn, SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruEditable bool

const (
	SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruEditableTrue  SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruEditable = true
	SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruEditableFalse SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruEditable = false
)

func (r SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruEditableTrue, SettingEditParamsBodyZonesSchemasOriginErrorPagePassThruEditableFalse:
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
type SettingEditParamsBodyZonesSchemasPolish struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasPolishID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasPolishValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasPolish) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasPolish) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasPolishID string

const (
	SettingEditParamsBodyZonesSchemasPolishIDPolish SettingEditParamsBodyZonesSchemasPolishID = "polish"
)

func (r SettingEditParamsBodyZonesSchemasPolishID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasPolishIDPolish:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasPolishValue string

const (
	SettingEditParamsBodyZonesSchemasPolishValueOff      SettingEditParamsBodyZonesSchemasPolishValue = "off"
	SettingEditParamsBodyZonesSchemasPolishValueLossless SettingEditParamsBodyZonesSchemasPolishValue = "lossless"
	SettingEditParamsBodyZonesSchemasPolishValueLossy    SettingEditParamsBodyZonesSchemasPolishValue = "lossy"
)

func (r SettingEditParamsBodyZonesSchemasPolishValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasPolishValueOff, SettingEditParamsBodyZonesSchemasPolishValueLossless, SettingEditParamsBodyZonesSchemasPolishValueLossy:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasPolishEditable bool

const (
	SettingEditParamsBodyZonesSchemasPolishEditableTrue  SettingEditParamsBodyZonesSchemasPolishEditable = true
	SettingEditParamsBodyZonesSchemasPolishEditableFalse SettingEditParamsBodyZonesSchemasPolishEditable = false
)

func (r SettingEditParamsBodyZonesSchemasPolishEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasPolishEditableTrue, SettingEditParamsBodyZonesSchemasPolishEditableFalse:
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

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingEditParamsBodyZonesSchemasResponseBuffering struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasResponseBufferingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasResponseBufferingValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasResponseBuffering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasResponseBuffering) implementsZonesSettingEditParamsBodyUnion() {
}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasResponseBufferingID string

const (
	SettingEditParamsBodyZonesSchemasResponseBufferingIDResponseBuffering SettingEditParamsBodyZonesSchemasResponseBufferingID = "response_buffering"
)

func (r SettingEditParamsBodyZonesSchemasResponseBufferingID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasResponseBufferingIDResponseBuffering:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasResponseBufferingValue string

const (
	SettingEditParamsBodyZonesSchemasResponseBufferingValueOn  SettingEditParamsBodyZonesSchemasResponseBufferingValue = "on"
	SettingEditParamsBodyZonesSchemasResponseBufferingValueOff SettingEditParamsBodyZonesSchemasResponseBufferingValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasResponseBufferingValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasResponseBufferingValueOn, SettingEditParamsBodyZonesSchemasResponseBufferingValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasResponseBufferingEditable bool

const (
	SettingEditParamsBodyZonesSchemasResponseBufferingEditableTrue  SettingEditParamsBodyZonesSchemasResponseBufferingEditable = true
	SettingEditParamsBodyZonesSchemasResponseBufferingEditableFalse SettingEditParamsBodyZonesSchemasResponseBufferingEditable = false
)

func (r SettingEditParamsBodyZonesSchemasResponseBufferingEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasResponseBufferingEditableTrue, SettingEditParamsBodyZonesSchemasResponseBufferingEditableFalse:
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
type SettingEditParamsBodyZonesSchemasRocketLoader struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasRocketLoaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasRocketLoaderValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasRocketLoader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasRocketLoader) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasRocketLoaderID string

const (
	SettingEditParamsBodyZonesSchemasRocketLoaderIDRocketLoader SettingEditParamsBodyZonesSchemasRocketLoaderID = "rocket_loader"
)

func (r SettingEditParamsBodyZonesSchemasRocketLoaderID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasRocketLoaderIDRocketLoader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasRocketLoaderValue string

const (
	SettingEditParamsBodyZonesSchemasRocketLoaderValueOn  SettingEditParamsBodyZonesSchemasRocketLoaderValue = "on"
	SettingEditParamsBodyZonesSchemasRocketLoaderValueOff SettingEditParamsBodyZonesSchemasRocketLoaderValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasRocketLoaderValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasRocketLoaderValueOn, SettingEditParamsBodyZonesSchemasRocketLoaderValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasRocketLoaderEditable bool

const (
	SettingEditParamsBodyZonesSchemasRocketLoaderEditableTrue  SettingEditParamsBodyZonesSchemasRocketLoaderEditable = true
	SettingEditParamsBodyZonesSchemasRocketLoaderEditableFalse SettingEditParamsBodyZonesSchemasRocketLoaderEditable = false
)

func (r SettingEditParamsBodyZonesSchemasRocketLoaderEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasRocketLoaderEditableTrue, SettingEditParamsBodyZonesSchemasRocketLoaderEditableFalse:
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

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingEditParamsBodyZonesSchemasSecurityLevel struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasSecurityLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasSecurityLevelValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasSecurityLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasSecurityLevel) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasSecurityLevelID string

const (
	SettingEditParamsBodyZonesSchemasSecurityLevelIDSecurityLevel SettingEditParamsBodyZonesSchemasSecurityLevelID = "security_level"
)

func (r SettingEditParamsBodyZonesSchemasSecurityLevelID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasSecurityLevelIDSecurityLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasSecurityLevelValue string

const (
	SettingEditParamsBodyZonesSchemasSecurityLevelValueOff            SettingEditParamsBodyZonesSchemasSecurityLevelValue = "off"
	SettingEditParamsBodyZonesSchemasSecurityLevelValueEssentiallyOff SettingEditParamsBodyZonesSchemasSecurityLevelValue = "essentially_off"
	SettingEditParamsBodyZonesSchemasSecurityLevelValueLow            SettingEditParamsBodyZonesSchemasSecurityLevelValue = "low"
	SettingEditParamsBodyZonesSchemasSecurityLevelValueMedium         SettingEditParamsBodyZonesSchemasSecurityLevelValue = "medium"
	SettingEditParamsBodyZonesSchemasSecurityLevelValueHigh           SettingEditParamsBodyZonesSchemasSecurityLevelValue = "high"
	SettingEditParamsBodyZonesSchemasSecurityLevelValueUnderAttack    SettingEditParamsBodyZonesSchemasSecurityLevelValue = "under_attack"
)

func (r SettingEditParamsBodyZonesSchemasSecurityLevelValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasSecurityLevelValueOff, SettingEditParamsBodyZonesSchemasSecurityLevelValueEssentiallyOff, SettingEditParamsBodyZonesSchemasSecurityLevelValueLow, SettingEditParamsBodyZonesSchemasSecurityLevelValueMedium, SettingEditParamsBodyZonesSchemasSecurityLevelValueHigh, SettingEditParamsBodyZonesSchemasSecurityLevelValueUnderAttack:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasSecurityLevelEditable bool

const (
	SettingEditParamsBodyZonesSchemasSecurityLevelEditableTrue  SettingEditParamsBodyZonesSchemasSecurityLevelEditable = true
	SettingEditParamsBodyZonesSchemasSecurityLevelEditableFalse SettingEditParamsBodyZonesSchemasSecurityLevelEditable = false
)

func (r SettingEditParamsBodyZonesSchemasSecurityLevelEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasSecurityLevelEditableTrue, SettingEditParamsBodyZonesSchemasSecurityLevelEditableFalse:
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

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type SettingEditParamsBodyZonesSchemasSortQueryStringForCache struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasSortQueryStringForCacheID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasSortQueryStringForCacheValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasSortQueryStringForCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasSortQueryStringForCache) implementsZonesSettingEditParamsBodyUnion() {
}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasSortQueryStringForCacheID string

const (
	SettingEditParamsBodyZonesSchemasSortQueryStringForCacheIDSortQueryStringForCache SettingEditParamsBodyZonesSchemasSortQueryStringForCacheID = "sort_query_string_for_cache"
)

func (r SettingEditParamsBodyZonesSchemasSortQueryStringForCacheID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasSortQueryStringForCacheIDSortQueryStringForCache:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasSortQueryStringForCacheValue string

const (
	SettingEditParamsBodyZonesSchemasSortQueryStringForCacheValueOn  SettingEditParamsBodyZonesSchemasSortQueryStringForCacheValue = "on"
	SettingEditParamsBodyZonesSchemasSortQueryStringForCacheValueOff SettingEditParamsBodyZonesSchemasSortQueryStringForCacheValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasSortQueryStringForCacheValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasSortQueryStringForCacheValueOn, SettingEditParamsBodyZonesSchemasSortQueryStringForCacheValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasSortQueryStringForCacheEditable bool

const (
	SettingEditParamsBodyZonesSchemasSortQueryStringForCacheEditableTrue  SettingEditParamsBodyZonesSchemasSortQueryStringForCacheEditable = true
	SettingEditParamsBodyZonesSchemasSortQueryStringForCacheEditableFalse SettingEditParamsBodyZonesSchemasSortQueryStringForCacheEditable = false
)

func (r SettingEditParamsBodyZonesSchemasSortQueryStringForCacheEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasSortQueryStringForCacheEditableTrue, SettingEditParamsBodyZonesSchemasSortQueryStringForCacheEditableFalse:
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
type SettingEditParamsBodyZonesSchemasSSL struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasSSLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasSSLValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasSSL) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasSSLID string

const (
	SettingEditParamsBodyZonesSchemasSSLIDSSL SettingEditParamsBodyZonesSchemasSSLID = "ssl"
)

func (r SettingEditParamsBodyZonesSchemasSSLID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasSSLIDSSL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasSSLValue string

const (
	SettingEditParamsBodyZonesSchemasSSLValueOff      SettingEditParamsBodyZonesSchemasSSLValue = "off"
	SettingEditParamsBodyZonesSchemasSSLValueFlexible SettingEditParamsBodyZonesSchemasSSLValue = "flexible"
	SettingEditParamsBodyZonesSchemasSSLValueFull     SettingEditParamsBodyZonesSchemasSSLValue = "full"
	SettingEditParamsBodyZonesSchemasSSLValueStrict   SettingEditParamsBodyZonesSchemasSSLValue = "strict"
)

func (r SettingEditParamsBodyZonesSchemasSSLValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasSSLValueOff, SettingEditParamsBodyZonesSchemasSSLValueFlexible, SettingEditParamsBodyZonesSchemasSSLValueFull, SettingEditParamsBodyZonesSchemasSSLValueStrict:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasSSLEditable bool

const (
	SettingEditParamsBodyZonesSchemasSSLEditableTrue  SettingEditParamsBodyZonesSchemasSSLEditable = true
	SettingEditParamsBodyZonesSchemasSSLEditableFalse SettingEditParamsBodyZonesSchemasSSLEditable = false
)

func (r SettingEditParamsBodyZonesSchemasSSLEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasSSLEditableTrue, SettingEditParamsBodyZonesSchemasSSLEditableFalse:
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

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingEditParamsBodyZonesSchemasTrueClientIPHeader struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasTrueClientIPHeaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasTrueClientIPHeaderValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasTrueClientIPHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasTrueClientIPHeader) implementsZonesSettingEditParamsBodyUnion() {
}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasTrueClientIPHeaderID string

const (
	SettingEditParamsBodyZonesSchemasTrueClientIPHeaderIDTrueClientIPHeader SettingEditParamsBodyZonesSchemasTrueClientIPHeaderID = "true_client_ip_header"
)

func (r SettingEditParamsBodyZonesSchemasTrueClientIPHeaderID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasTrueClientIPHeaderIDTrueClientIPHeader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasTrueClientIPHeaderValue string

const (
	SettingEditParamsBodyZonesSchemasTrueClientIPHeaderValueOn  SettingEditParamsBodyZonesSchemasTrueClientIPHeaderValue = "on"
	SettingEditParamsBodyZonesSchemasTrueClientIPHeaderValueOff SettingEditParamsBodyZonesSchemasTrueClientIPHeaderValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasTrueClientIPHeaderValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasTrueClientIPHeaderValueOn, SettingEditParamsBodyZonesSchemasTrueClientIPHeaderValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasTrueClientIPHeaderEditable bool

const (
	SettingEditParamsBodyZonesSchemasTrueClientIPHeaderEditableTrue  SettingEditParamsBodyZonesSchemasTrueClientIPHeaderEditable = true
	SettingEditParamsBodyZonesSchemasTrueClientIPHeaderEditableFalse SettingEditParamsBodyZonesSchemasTrueClientIPHeaderEditable = false
)

func (r SettingEditParamsBodyZonesSchemasTrueClientIPHeaderEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasTrueClientIPHeaderEditableTrue, SettingEditParamsBodyZonesSchemasTrueClientIPHeaderEditableFalse:
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
type SettingEditParamsBodyZonesSchemasWAF struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsBodyZonesSchemasWAFID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsBodyZonesSchemasWAFValue] `json:"value,required"`
}

func (r SettingEditParamsBodyZonesSchemasWAF) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyZonesSchemasWAF) implementsZonesSettingEditParamsBodyUnion() {}

// ID of the zone setting.
type SettingEditParamsBodyZonesSchemasWAFID string

const (
	SettingEditParamsBodyZonesSchemasWAFIDWAF SettingEditParamsBodyZonesSchemasWAFID = "waf"
)

func (r SettingEditParamsBodyZonesSchemasWAFID) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasWAFIDWAF:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsBodyZonesSchemasWAFValue string

const (
	SettingEditParamsBodyZonesSchemasWAFValueOn  SettingEditParamsBodyZonesSchemasWAFValue = "on"
	SettingEditParamsBodyZonesSchemasWAFValueOff SettingEditParamsBodyZonesSchemasWAFValue = "off"
)

func (r SettingEditParamsBodyZonesSchemasWAFValue) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasWAFValueOn, SettingEditParamsBodyZonesSchemasWAFValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsBodyZonesSchemasWAFEditable bool

const (
	SettingEditParamsBodyZonesSchemasWAFEditableTrue  SettingEditParamsBodyZonesSchemasWAFEditable = true
	SettingEditParamsBodyZonesSchemasWAFEditableFalse SettingEditParamsBodyZonesSchemasWAFEditable = false
)

func (r SettingEditParamsBodyZonesSchemasWAFEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsBodyZonesSchemasWAFEditableTrue, SettingEditParamsBodyZonesSchemasWAFEditableFalse:
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
	SettingEditParamsBodyIDMirage                        SettingEditParamsBodyID = "mirage"
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
	case SettingEditParamsBodyID0rtt, SettingEditParamsBodyIDAdvancedDDoS, SettingEditParamsBodyIDAlwaysOnline, SettingEditParamsBodyIDAlwaysUseHTTPS, SettingEditParamsBodyIDAutomaticHTTPSRewrites, SettingEditParamsBodyIDBrotli, SettingEditParamsBodyIDBrowserCacheTTL, SettingEditParamsBodyIDBrowserCheck, SettingEditParamsBodyIDCacheLevel, SettingEditParamsBodyIDChallengeTTL, SettingEditParamsBodyIDCiphers, SettingEditParamsBodyIDCNAMEFlattening, SettingEditParamsBodyIDDevelopmentMode, SettingEditParamsBodyIDEarlyHints, SettingEditParamsBodyIDEdgeCacheTTL, SettingEditParamsBodyIDEmailObfuscation, SettingEditParamsBodyIDH2Prioritization, SettingEditParamsBodyIDHotlinkProtection, SettingEditParamsBodyIDHTTP2, SettingEditParamsBodyIDHTTP3, SettingEditParamsBodyIDImageResizing, SettingEditParamsBodyIDIPGeolocation, SettingEditParamsBodyIDIPV6, SettingEditParamsBodyIDMaxUpload, SettingEditParamsBodyIDMinTLSVersion, SettingEditParamsBodyIDMirage, SettingEditParamsBodyIDNEL, SettingEditParamsBodyIDOpportunisticEncryption, SettingEditParamsBodyIDOpportunisticOnion, SettingEditParamsBodyIDOrangeToOrange, SettingEditParamsBodyIDOriginErrorPagePassThru, SettingEditParamsBodyIDPolish, SettingEditParamsBodyIDPrefetchPreload, SettingEditParamsBodyIDProxyReadTimeout, SettingEditParamsBodyIDPseudoIPV4, SettingEditParamsBodyIDReplaceInsecureJS, SettingEditParamsBodyIDResponseBuffering, SettingEditParamsBodyIDRocketLoader, SettingEditParamsBodyIDAutomaticPlatformOptimization, SettingEditParamsBodyIDSecurityHeader, SettingEditParamsBodyIDSecurityLevel, SettingEditParamsBodyIDServerSideExclude, SettingEditParamsBodyIDSha1Support, SettingEditParamsBodyIDSortQueryStringForCache, SettingEditParamsBodyIDSSL, SettingEditParamsBodyIDSSLRecommender, SettingEditParamsBodyIDTLS1_2Only, SettingEditParamsBodyIDTLS1_3, SettingEditParamsBodyIDTLSClientAuth, SettingEditParamsBodyIDTrueClientIPHeader, SettingEditParamsBodyIDWAF, SettingEditParamsBodyIDWebP, SettingEditParamsBodyIDWebsockets:
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
