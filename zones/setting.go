// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
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
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if settingID == "" {
		err = errors.New("missing required setting_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/%s", params.ZoneID, settingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetch a single zone setting by name
func (r *SettingService) Get(ctx context.Context, settingID string, query SettingGetParams, opts ...option.RequestOption) (res *SettingGetResponse, err error) {
	var env SettingGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if settingID == "" {
		err = errors.New("missing required setting_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/settings/%s", query.ZoneID, settingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type AdvancedDDoS struct {
	// ID of the zone setting.
	ID AdvancedDDoSID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value AdvancedDDoSValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AdvancedDDoSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r AdvancedDDoS) implementsSettingEditResponse() {}

func (r AdvancedDDoS) implementsSettingGetResponse() {}

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

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type AlwaysOnline struct {
	// ID of the zone setting.
	ID AlwaysOnlineID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value AlwaysOnlineValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r AlwaysOnline) implementsSettingEditResponse() {}

func (r AlwaysOnline) implementsSettingGetResponse() {}

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

func (r AlwaysUseHTTPS) ImplementsPageRuleAction() {}

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

func (r AlwaysUseHTTPSParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r AlwaysUseHTTPSParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r AlwaysUseHTTPSParam) ImplementsPageRuleEditParamsActionUnion() {}

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

func (r AutomaticHTTPSRewrites) ImplementsPageRuleAction() {}

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

func (r AutomaticHTTPSRewritesParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r AutomaticHTTPSRewritesParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r AutomaticHTTPSRewritesParam) ImplementsPageRuleEditParamsActionUnion() {}

type AutomaticPlatformOptimization struct {
	// Indicates whether or not
	// [cache by device type](https://developers.cloudflare.com/automatic-platform-optimization/reference/cache-device-type/)
	// is enabled.
	CacheByDeviceType bool `json:"cache_by_device_type" api:"required"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf bool `json:"cf" api:"required"`
	// Indicates whether or not Automatic Platform Optimization is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// An array of hostnames where Automatic Platform Optimization for WordPress is
	// activated.
	Hostnames []string `json:"hostnames" api:"required" format:"hostname"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress bool `json:"wordpress" api:"required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin bool                              `json:"wp_plugin" api:"required"`
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
	CacheByDeviceType param.Field[bool] `json:"cache_by_device_type" api:"required"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf param.Field[bool] `json:"cf" api:"required"`
	// Indicates whether or not Automatic Platform Optimization is enabled.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// An array of hostnames where Automatic Platform Optimization for WordPress is
	// activated.
	Hostnames param.Field[[]string] `json:"hostnames" api:"required" format:"hostname"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress param.Field[bool] `json:"wordpress" api:"required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin param.Field[bool] `json:"wp_plugin" api:"required"`
}

func (r AutomaticPlatformOptimizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AutomaticPlatformOptimizationParam) ImplementsSettingEditParamsBodyValueValueUnion() {}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type Brotli struct {
	// ID of the zone setting.
	ID BrotliID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value BrotliValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable BrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time  `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r Brotli) implementsSettingEditResponse() {}

func (r Brotli) implementsSettingGetResponse() {}

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

type BrowserCacheTTL struct {
	// Control how long resources cached by client browsers remain valid.
	ID BrowserCacheTTLID `json:"id"`
	// The number of seconds to cache resources for. Setting this to 0 enables "Respect
	// Existing Headers".
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

func (r BrowserCacheTTL) ImplementsPageRuleAction() {}

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
	// The number of seconds to cache resources for. Setting this to 0 enables "Respect
	// Existing Headers".
	Value param.Field[int64] `json:"value"`
}

func (r BrowserCacheTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BrowserCacheTTLParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r BrowserCacheTTLParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r BrowserCacheTTLParam) ImplementsPageRuleEditParamsActionUnion() {}

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

func (r BrowserCheck) ImplementsPageRuleAction() {}

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

func (r BrowserCheckParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r BrowserCheckParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r BrowserCheckParam) ImplementsPageRuleEditParamsActionUnion() {}

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

func (r CacheLevel) ImplementsPageRuleAction() {}

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

func (r CacheLevelParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r CacheLevelParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r CacheLevelParam) ImplementsPageRuleEditParamsActionUnion() {}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ChallengeTTL struct {
	// ID of the zone setting.
	ID ChallengeTTLID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value ChallengeTTLValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time        `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r ChallengeTTL) implementsSettingEditResponse() {}

func (r ChallengeTTL) implementsSettingGetResponse() {}

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

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type Ciphers struct {
	// ID of the zone setting.
	ID CiphersID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value []string `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable CiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time   `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r Ciphers) implementsSettingEditResponse() {}

func (r Ciphers) implementsSettingGetResponse() {}

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

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type DevelopmentMode struct {
	// ID of the zone setting.
	ID DevelopmentModeID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value DevelopmentModeValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable DevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r DevelopmentMode) implementsSettingEditResponse() {}

func (r DevelopmentMode) implementsSettingGetResponse() {}

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

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type EarlyHints struct {
	// ID of the zone setting.
	ID EarlyHintsID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value EarlyHintsValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable EarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time      `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r EarlyHints) implementsSettingEditResponse() {}

func (r EarlyHints) implementsSettingGetResponse() {}

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

func (r EmailObfuscation) ImplementsPageRuleAction() {}

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

func (r EmailObfuscationParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r EmailObfuscationParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r EmailObfuscationParam) ImplementsPageRuleEditParamsActionUnion() {}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type H2Prioritization struct {
	// ID of the zone setting.
	ID H2PrioritizationID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value H2PrioritizationValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable H2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time            `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r H2Prioritization) implementsSettingEditResponse() {}

func (r H2Prioritization) implementsSettingGetResponse() {}

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

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type HotlinkProtection struct {
	// ID of the zone setting.
	ID HotlinkProtectionID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value HotlinkProtectionValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable HotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time             `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r HotlinkProtection) implementsSettingEditResponse() {}

func (r HotlinkProtection) implementsSettingGetResponse() {}

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

// HTTP2 enabled for this zone.
type HTTP2 struct {
	// ID of the zone setting.
	ID HTTP2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value HTTP2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable HTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r HTTP2) implementsSettingEditResponse() {}

func (r HTTP2) implementsSettingGetResponse() {}

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

// HTTP3 enabled for this zone.
type HTTP3 struct {
	// ID of the zone setting.
	ID HTTP3ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value HTTP3Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable HTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r HTTP3) implementsSettingEditResponse() {}

func (r HTTP3) implementsSettingGetResponse() {}

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

// Image Transformations provides on-demand resizing, conversion and optimization
// for images served through Cloudflare's network. Refer to the
// [Image Transformations documentation](https://developers.cloudflare.com/images/)
// for more information.
type ImageResizing struct {
	// ID of the zone setting.
	ID ImageResizingID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value ImageResizingValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r ImageResizing) implementsSettingEditResponse() {}

func (r ImageResizing) implementsSettingGetResponse() {}

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

func (r IPGeolocation) ImplementsPageRuleAction() {}

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

func (r IPGeolocationParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r IPGeolocationParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r IPGeolocationParam) ImplementsPageRuleEditParamsActionUnion() {}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type IPV6 struct {
	// ID of the zone setting.
	ID IPV6ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value IPV6Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable IPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r IPV6) implementsSettingEditResponse() {}

func (r IPV6) implementsSettingGetResponse() {}

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

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type MinTLSVersion struct {
	// ID of the zone setting.
	ID MinTLSVersionID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value MinTLSVersionValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r MinTLSVersion) implementsSettingEditResponse() {}

func (r MinTLSVersion) implementsSettingGetResponse() {}

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

// Deprecated: Mirage is deprecated. This functionality is no longer supported. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2025-11-03
// for further details.
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

func (r Mirage) ImplementsPageRuleAction() {}

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

// Deprecated: Mirage is deprecated. This functionality is no longer supported. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2025-11-03
// for further details.
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

func (r MirageParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r MirageParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r MirageParam) ImplementsPageRuleEditParamsActionUnion() {}

// Enable Network Error Logging reporting on your zone. (Beta)
type NEL struct {
	// Zone setting identifier.
	ID NELID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value NELValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable NELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r NEL) implementsSettingEditResponse() {}

func (r NEL) implementsSettingGetResponse() {}

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

func (r OpportunisticEncryption) ImplementsPageRuleAction() {}

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

func (r OpportunisticEncryptionParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r OpportunisticEncryptionParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r OpportunisticEncryptionParam) ImplementsPageRuleEditParamsActionUnion() {}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type OpportunisticOnion struct {
	// ID of the zone setting.
	ID OpportunisticOnionID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value OpportunisticOnionValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r OpportunisticOnion) implementsSettingEditResponse() {}

func (r OpportunisticOnion) implementsSettingGetResponse() {}

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

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type OrangeToOrange struct {
	// ID of the zone setting.
	ID OrangeToOrangeID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value OrangeToOrangeValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time          `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r OrangeToOrange) implementsSettingEditResponse() {}

func (r OrangeToOrange) implementsSettingGetResponse() {}

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

func (r OriginErrorPagePassThru) ImplementsPageRuleAction() {}

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

func (r OriginErrorPagePassThruParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r OriginErrorPagePassThruParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r OriginErrorPagePassThruParam) ImplementsPageRuleEditParamsActionUnion() {}

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

func (r Polish) ImplementsPageRuleAction() {}

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

func (r PolishParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r PolishParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r PolishParam) ImplementsPageRuleEditParamsActionUnion() {}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type PrefetchPreload struct {
	// ID of the zone setting.
	ID PrefetchPreloadID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value PrefetchPreloadValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r PrefetchPreload) implementsSettingEditResponse() {}

func (r PrefetchPreload) implementsSettingGetResponse() {}

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

// Maximum time between two read operations from origin.
type ProxyReadTimeout struct {
	// ID of the zone setting.
	ID ProxyReadTimeoutID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value float64 `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time            `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r ProxyReadTimeout) implementsSettingEditResponse() {}

func (r ProxyReadTimeout) implementsSettingGetResponse() {}

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

// The value set for the Pseudo IPv4 setting.
type PseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID PseudoIPV4ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value PseudoIPV4Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time      `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r PseudoIPV4) implementsSettingEditResponse() {}

func (r PseudoIPV4) implementsSettingGetResponse() {}

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

// Deprecated: This page rule is deprecated. This functionality is no longer
// supported.
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

func (r ResponseBuffering) ImplementsPageRuleAction() {}

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

// Deprecated: This page rule is deprecated. This functionality is no longer
// supported.
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

func (r ResponseBufferingParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r ResponseBufferingParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r ResponseBufferingParam) ImplementsPageRuleEditParamsActionUnion() {}

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

func (r RocketLoader) ImplementsPageRuleAction() {}

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

func (r RocketLoaderParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r RocketLoaderParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r RocketLoaderParam) ImplementsPageRuleEditParamsActionUnion() {}

// Cloudflare security header for a zone.
type SecurityHeaders struct {
	// ID of the zone's security header.
	ID SecurityHeadersID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SecurityHeadersValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SecurityHeadersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SecurityHeaders) implementsSettingEditResponse() {}

func (r SecurityHeaders) implementsSettingGetResponse() {}

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

func (r SecurityLevel) ImplementsPageRuleAction() {}

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

func (r SecurityLevelParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r SecurityLevelParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r SecurityLevelParam) ImplementsPageRuleEditParamsActionUnion() {}

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
	ID ServerSideExcludesID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value ServerSideExcludesValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ServerSideExcludesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time              `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r ServerSideExcludes) implementsSettingEditResponse() {}

func (r ServerSideExcludes) implementsSettingGetResponse() {}

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

func (r SortQueryStringForCache) ImplementsPageRuleAction() {}

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

func (r SortQueryStringForCacheParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r SortQueryStringForCacheParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r SortQueryStringForCacheParam) ImplementsPageRuleEditParamsActionUnion() {}

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

func (r SSL) ImplementsPageRuleAction() {}

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

func (r SSLParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r SSLParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r SSLParam) ImplementsPageRuleEditParamsActionUnion() {}

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

func (r SSLRecommender) implementsSettingEditResponse() {}

func (r SSLRecommender) implementsSettingGetResponse() {}

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

// Enables Crypto TLS 1.3 feature for a zone.
type TLS1_3 struct {
	// ID of the zone setting.
	ID TLS1_3ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value TLS1_3Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time  `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r TLS1_3) implementsSettingEditResponse() {}

func (r TLS1_3) implementsSettingGetResponse() {}

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

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type TLSClientAuth struct {
	// ID of the zone setting.
	ID TLSClientAuthID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value TLSClientAuthValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r TLSClientAuth) implementsSettingEditResponse() {}

func (r TLSClientAuth) implementsSettingGetResponse() {}

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

func (r TrueClientIPHeader) ImplementsPageRuleAction() {}

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

func (r TrueClientIPHeaderParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r TrueClientIPHeaderParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r TrueClientIPHeaderParam) ImplementsPageRuleEditParamsActionUnion() {}

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

func (r WAF) ImplementsPageRuleAction() {}

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

func (r WAFParam) ImplementsPageRuleNewParamsActionUnion() {}

func (r WAFParam) ImplementsPageRuleUpdateParamsActionUnion() {}

func (r WAFParam) ImplementsPageRuleEditParamsActionUnion() {}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type WebP struct {
	// ID of the zone setting.
	ID WebPID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value WebPValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WebPEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r WebP) implementsSettingEditResponse() {}

func (r WebP) implementsSettingGetResponse() {}

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

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type Websocket struct {
	// ID of the zone setting.
	ID WebsocketID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value WebsocketValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WebsocketEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time     `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r Websocket) implementsSettingEditResponse() {}

func (r Websocket) implementsSettingGetResponse() {}

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

// 0-RTT session resumption enabled for this zone.
type ZeroRTT struct {
	// ID of the zone setting.
	ID ZeroRTTID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value ZeroRTTValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZeroRTTEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time   `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r ZeroRTT) implementsSettingEditResponse() {}

func (r ZeroRTT) implementsSettingGetResponse() {}

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
type SettingEditResponse struct {
	// ID of the zone setting.
	ID SettingEditResponseID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseEditable `json:"editable"`
	// ssl-recommender enrollment setting.
	Enabled bool `json:"enabled"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// This field can have the runtime type of [ZeroRTTValue], [AdvancedDDoSValue],
	// [SettingEditResponseZonesCacheRulesAegisValue], [AlwaysOnlineValue],
	// [SettingEditResponseZonesAlwaysUseHTTPS2Value],
	// [SettingEditResponseZonesAutomaticHTTPSRewrites2Value], [BrotliValue], [int64],
	// [SettingEditResponseZonesBrowserCheck2Value],
	// [SettingEditResponseZonesCacheLevel2Value], [ChallengeTTLValue],
	// [SettingEditResponseZonesChinaNetworkEnabledValue],
	// [SettingEditResponseZonesContentConverterValue], [[]string],
	// [SettingEditResponseZonesCNAMEFlatteningValue], [DevelopmentModeValue],
	// [EarlyHintsValue], [SettingEditResponseZonesEdgeCacheTTL2Value],
	// [SettingEditResponseZonesEmailObfuscation2Value], [H2PrioritizationValue],
	// [HotlinkProtectionValue], [HTTP2Value], [HTTP3Value], [ImageResizingValue],
	// [SettingEditResponseZonesIPGeolocation2Value], [IPV6Value],
	// [SettingEditResponseZonesMaxUploadValue], [MinTLSVersionValue],
	// [SettingEditResponseZonesMirage2Value], [NELValue],
	// [SettingEditResponseZonesOpportunisticEncryption2Value],
	// [OpportunisticOnionValue], [OrangeToOrangeValue],
	// [SettingEditResponseZonesOriginErrorPagePassThru2Value],
	// [SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionValue],
	// [SettingEditResponseZonesPolish2Value], [PrefetchPreloadValue],
	// [SettingEditResponseZonesPrivacyPassValue], [float64], [PseudoIPV4Value],
	// [SettingEditResponseZonesRedirectsForAITrainingValue],
	// [SettingEditResponseZonesReplaceInsecureJSValue],
	// [SettingEditResponseZonesResponseBuffering2Value],
	// [SettingEditResponseZonesRocketLoader2Value], [AutomaticPlatformOptimization],
	// [SettingEditResponseZonesSearchForAgentsValue], [SecurityHeadersValue],
	// [SettingEditResponseZonesSecurityLevel2Value], [ServerSideExcludesValue],
	// [SettingEditResponseZonesSha1SupportValue],
	// [SettingEditResponseZonesSortQueryStringForCache2Value],
	// [SettingEditResponseZonesSSL2Value], [SettingEditResponseZonesTLS1_2OnlyValue],
	// [TLS1_3Value], [TLSClientAuthValue],
	// [SettingEditResponseZonesTransformationsValue], [string],
	// [SettingEditResponseZonesTrueClientIPHeader2Value],
	// [SettingEditResponseZonesWAF2Value], [WebPValue], [WebsocketValue].
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
// Possible runtime types of the union are [ZeroRTT], [AdvancedDDoS],
// [SettingEditResponseZonesCacheRulesAegis], [AlwaysOnline],
// [SettingEditResponseZonesAlwaysUseHTTPS2],
// [SettingEditResponseZonesAutomaticHTTPSRewrites2], [Brotli],
// [SettingEditResponseZonesBrowserCacheTTL2],
// [SettingEditResponseZonesBrowserCheck2], [SettingEditResponseZonesCacheLevel2],
// [ChallengeTTL], [SettingEditResponseZonesChinaNetworkEnabled],
// [SettingEditResponseZonesContentConverter], [Ciphers],
// [SettingEditResponseZonesCNAMEFlattening], [DevelopmentMode], [EarlyHints],
// [SettingEditResponseZonesEdgeCacheTTL2],
// [SettingEditResponseZonesEmailObfuscation2], [H2Prioritization],
// [HotlinkProtection], [HTTP2], [HTTP3], [ImageResizing],
// [SettingEditResponseZonesIPGeolocation2], [IPV6],
// [SettingEditResponseZonesMaxUpload], [MinTLSVersion],
// [SettingEditResponseZonesMirage2], [NEL],
// [SettingEditResponseZonesOpportunisticEncryption2], [OpportunisticOnion],
// [OrangeToOrange], [SettingEditResponseZonesOriginErrorPagePassThru2],
// [SettingEditResponseZonesCacheRulesOriginH2MaxStreams],
// [SettingEditResponseZonesCacheRulesOriginMaxHTTPVersion],
// [SettingEditResponseZonesPolish2], [PrefetchPreload],
// [SettingEditResponseZonesPrivacyPass], [ProxyReadTimeout], [PseudoIPV4],
// [SettingEditResponseZonesRedirectsForAITraining],
// [SettingEditResponseZonesReplaceInsecureJS],
// [SettingEditResponseZonesResponseBuffering2],
// [SettingEditResponseZonesRocketLoader2],
// [SettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [SettingEditResponseZonesSearchForAgents], [SecurityHeaders],
// [SettingEditResponseZonesSecurityLevel2], [ServerSideExcludes],
// [SettingEditResponseZonesSha1Support],
// [SettingEditResponseZonesSortQueryStringForCache2],
// [SettingEditResponseZonesSSL2], [SSLRecommender],
// [SettingEditResponseZonesTLS1_2Only], [TLS1_3], [TLSClientAuth],
// [SettingEditResponseZonesTransformations],
// [SettingEditResponseZonesTransformationsAllowedOrigins],
// [SettingEditResponseZonesTrueClientIPHeader2], [SettingEditResponseZonesWAF2],
// [WebP], [Websocket].
func (r SettingEditResponse) AsUnion() SettingEditResponseUnion {
	return r.union
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [ZeroRTT], [AdvancedDDoS],
// [SettingEditResponseZonesCacheRulesAegis], [AlwaysOnline],
// [SettingEditResponseZonesAlwaysUseHTTPS2],
// [SettingEditResponseZonesAutomaticHTTPSRewrites2], [Brotli],
// [SettingEditResponseZonesBrowserCacheTTL2],
// [SettingEditResponseZonesBrowserCheck2], [SettingEditResponseZonesCacheLevel2],
// [ChallengeTTL], [SettingEditResponseZonesChinaNetworkEnabled],
// [SettingEditResponseZonesContentConverter], [Ciphers],
// [SettingEditResponseZonesCNAMEFlattening], [DevelopmentMode], [EarlyHints],
// [SettingEditResponseZonesEdgeCacheTTL2],
// [SettingEditResponseZonesEmailObfuscation2], [H2Prioritization],
// [HotlinkProtection], [HTTP2], [HTTP3], [ImageResizing],
// [SettingEditResponseZonesIPGeolocation2], [IPV6],
// [SettingEditResponseZonesMaxUpload], [MinTLSVersion],
// [SettingEditResponseZonesMirage2], [NEL],
// [SettingEditResponseZonesOpportunisticEncryption2], [OpportunisticOnion],
// [OrangeToOrange], [SettingEditResponseZonesOriginErrorPagePassThru2],
// [SettingEditResponseZonesCacheRulesOriginH2MaxStreams],
// [SettingEditResponseZonesCacheRulesOriginMaxHTTPVersion],
// [SettingEditResponseZonesPolish2], [PrefetchPreload],
// [SettingEditResponseZonesPrivacyPass], [ProxyReadTimeout], [PseudoIPV4],
// [SettingEditResponseZonesRedirectsForAITraining],
// [SettingEditResponseZonesReplaceInsecureJS],
// [SettingEditResponseZonesResponseBuffering2],
// [SettingEditResponseZonesRocketLoader2],
// [SettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [SettingEditResponseZonesSearchForAgents], [SecurityHeaders],
// [SettingEditResponseZonesSecurityLevel2], [ServerSideExcludes],
// [SettingEditResponseZonesSha1Support],
// [SettingEditResponseZonesSortQueryStringForCache2],
// [SettingEditResponseZonesSSL2], [SSLRecommender],
// [SettingEditResponseZonesTLS1_2Only], [TLS1_3], [TLSClientAuth],
// [SettingEditResponseZonesTransformations],
// [SettingEditResponseZonesTransformationsAllowedOrigins],
// [SettingEditResponseZonesTrueClientIPHeader2], [SettingEditResponseZonesWAF2],
// [WebP] or [Websocket].
type SettingEditResponseUnion interface {
	implementsSettingEditResponse()
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
			Type:       reflect.TypeOf(SettingEditResponseZonesCacheRulesAegis{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesAlwaysUseHTTPS2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesAutomaticHTTPSRewrites2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Brotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesBrowserCacheTTL2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesBrowserCheck2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesCacheLevel2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesChinaNetworkEnabled{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesContentConverter{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesEdgeCacheTTL2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesEmailObfuscation2{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesIPGeolocation2{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesMirage2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesOpportunisticEncryption2{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesOriginErrorPagePassThru2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesCacheRulesOriginH2MaxStreams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesCacheRulesOriginMaxHTTPVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesPolish2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesPrivacyPass{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesRedirectsForAITraining{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesReplaceInsecureJS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesResponseBuffering2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesRocketLoader2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSearchForAgents{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecurityHeaders{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSecurityLevel2{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesSortQueryStringForCache2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSSL2{}),
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
			Type:       reflect.TypeOf(SettingEditResponseZonesTransformations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesTransformationsAllowedOrigins{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesTrueClientIPHeader2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesWAF2{}),
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

// Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your
// layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your
// account so that you can increase your origin security by only allowing traffic
// from a small list of IP addresses.
type SettingEditResponseZonesCacheRulesAegis struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesCacheRulesAegisID `json:"id" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Value of the zone setting.
	Value SettingEditResponseZonesCacheRulesAegisValue `json:"value"`
	JSON  settingEditResponseZonesCacheRulesAegisJSON  `json:"-"`
}

// settingEditResponseZonesCacheRulesAegisJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesCacheRulesAegis]
type settingEditResponseZonesCacheRulesAegisJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesCacheRulesAegis) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesCacheRulesAegisJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesCacheRulesAegis) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesCacheRulesAegisID string

const (
	SettingEditResponseZonesCacheRulesAegisIDAegis SettingEditResponseZonesCacheRulesAegisID = "aegis"
)

func (r SettingEditResponseZonesCacheRulesAegisID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesCacheRulesAegisIDAegis:
		return true
	}
	return false
}

// Value of the zone setting.
type SettingEditResponseZonesCacheRulesAegisValue struct {
	// Whether the feature is enabled or not.
	Enabled bool `json:"enabled"`
	// Egress pool id which refers to a grouping of dedicated egress IPs through which
	// Cloudflare will connect to origin.
	PoolID string                                           `json:"pool_id"`
	JSON   settingEditResponseZonesCacheRulesAegisValueJSON `json:"-"`
}

// settingEditResponseZonesCacheRulesAegisValueJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesCacheRulesAegisValue]
type settingEditResponseZonesCacheRulesAegisValueJSON struct {
	Enabled     apijson.Field
	PoolID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesCacheRulesAegisValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesCacheRulesAegisValueJSON) RawJSON() string {
	return r.raw
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingEditResponseZonesAlwaysUseHTTPS2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesAlwaysUseHTTPS2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesAlwaysUseHTTPS2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesAlwaysUseHTTPS2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesAlwaysUseHttps2JSON `json:"-"`
}

// settingEditResponseZonesAlwaysUseHttps2JSON contains the JSON metadata for the
// struct [SettingEditResponseZonesAlwaysUseHTTPS2]
type settingEditResponseZonesAlwaysUseHttps2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesAlwaysUseHTTPS2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesAlwaysUseHttps2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesAlwaysUseHTTPS2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesAlwaysUseHTTPS2ID string

const (
	SettingEditResponseZonesAlwaysUseHTTPS2IDAlwaysUseHTTPS SettingEditResponseZonesAlwaysUseHTTPS2ID = "always_use_https"
)

func (r SettingEditResponseZonesAlwaysUseHTTPS2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesAlwaysUseHTTPS2IDAlwaysUseHTTPS:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesAlwaysUseHTTPS2Value string

const (
	SettingEditResponseZonesAlwaysUseHTTPS2ValueOn  SettingEditResponseZonesAlwaysUseHTTPS2Value = "on"
	SettingEditResponseZonesAlwaysUseHTTPS2ValueOff SettingEditResponseZonesAlwaysUseHTTPS2Value = "off"
)

func (r SettingEditResponseZonesAlwaysUseHTTPS2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesAlwaysUseHTTPS2ValueOn, SettingEditResponseZonesAlwaysUseHTTPS2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesAlwaysUseHTTPS2Editable bool

const (
	SettingEditResponseZonesAlwaysUseHTTPS2EditableTrue  SettingEditResponseZonesAlwaysUseHTTPS2Editable = true
	SettingEditResponseZonesAlwaysUseHTTPS2EditableFalse SettingEditResponseZonesAlwaysUseHTTPS2Editable = false
)

func (r SettingEditResponseZonesAlwaysUseHTTPS2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesAlwaysUseHTTPS2EditableTrue, SettingEditResponseZonesAlwaysUseHTTPS2EditableFalse:
		return true
	}
	return false
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingEditResponseZonesAutomaticHTTPSRewrites2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesAutomaticHTTPSRewrites2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesAutomaticHTTPSRewrites2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesAutomaticHTTPSRewrites2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesAutomaticHTTPSRewrites2JSON `json:"-"`
}

// settingEditResponseZonesAutomaticHTTPSRewrites2JSON contains the JSON metadata
// for the struct [SettingEditResponseZonesAutomaticHTTPSRewrites2]
type settingEditResponseZonesAutomaticHTTPSRewrites2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesAutomaticHTTPSRewrites2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesAutomaticHTTPSRewrites2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesAutomaticHTTPSRewrites2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesAutomaticHTTPSRewrites2ID string

const (
	SettingEditResponseZonesAutomaticHTTPSRewrites2IDAutomaticHTTPSRewrites SettingEditResponseZonesAutomaticHTTPSRewrites2ID = "automatic_https_rewrites"
)

func (r SettingEditResponseZonesAutomaticHTTPSRewrites2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesAutomaticHTTPSRewrites2IDAutomaticHTTPSRewrites:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesAutomaticHTTPSRewrites2Value string

const (
	SettingEditResponseZonesAutomaticHTTPSRewrites2ValueOn  SettingEditResponseZonesAutomaticHTTPSRewrites2Value = "on"
	SettingEditResponseZonesAutomaticHTTPSRewrites2ValueOff SettingEditResponseZonesAutomaticHTTPSRewrites2Value = "off"
)

func (r SettingEditResponseZonesAutomaticHTTPSRewrites2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesAutomaticHTTPSRewrites2ValueOn, SettingEditResponseZonesAutomaticHTTPSRewrites2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesAutomaticHTTPSRewrites2Editable bool

const (
	SettingEditResponseZonesAutomaticHTTPSRewrites2EditableTrue  SettingEditResponseZonesAutomaticHTTPSRewrites2Editable = true
	SettingEditResponseZonesAutomaticHTTPSRewrites2EditableFalse SettingEditResponseZonesAutomaticHTTPSRewrites2Editable = false
)

func (r SettingEditResponseZonesAutomaticHTTPSRewrites2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesAutomaticHTTPSRewrites2EditableTrue, SettingEditResponseZonesAutomaticHTTPSRewrites2EditableFalse:
		return true
	}
	return false
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingEditResponseZonesBrowserCacheTTL2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesBrowserCacheTTL2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value int64 `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesBrowserCacheTTL2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesBrowserCacheTtl2JSON `json:"-"`
}

// settingEditResponseZonesBrowserCacheTtl2JSON contains the JSON metadata for the
// struct [SettingEditResponseZonesBrowserCacheTTL2]
type settingEditResponseZonesBrowserCacheTtl2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesBrowserCacheTTL2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesBrowserCacheTtl2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesBrowserCacheTTL2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesBrowserCacheTTL2ID string

const (
	SettingEditResponseZonesBrowserCacheTTL2IDBrowserCacheTTL SettingEditResponseZonesBrowserCacheTTL2ID = "browser_cache_ttl"
)

func (r SettingEditResponseZonesBrowserCacheTTL2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesBrowserCacheTTL2IDBrowserCacheTTL:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesBrowserCacheTTL2Editable bool

const (
	SettingEditResponseZonesBrowserCacheTTL2EditableTrue  SettingEditResponseZonesBrowserCacheTTL2Editable = true
	SettingEditResponseZonesBrowserCacheTTL2EditableFalse SettingEditResponseZonesBrowserCacheTTL2Editable = false
)

func (r SettingEditResponseZonesBrowserCacheTTL2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesBrowserCacheTTL2EditableTrue, SettingEditResponseZonesBrowserCacheTTL2EditableFalse:
		return true
	}
	return false
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type SettingEditResponseZonesBrowserCheck2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesBrowserCheck2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesBrowserCheck2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesBrowserCheck2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesBrowserCheck2JSON `json:"-"`
}

// settingEditResponseZonesBrowserCheck2JSON contains the JSON metadata for the
// struct [SettingEditResponseZonesBrowserCheck2]
type settingEditResponseZonesBrowserCheck2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesBrowserCheck2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesBrowserCheck2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesBrowserCheck2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesBrowserCheck2ID string

const (
	SettingEditResponseZonesBrowserCheck2IDBrowserCheck SettingEditResponseZonesBrowserCheck2ID = "browser_check"
)

func (r SettingEditResponseZonesBrowserCheck2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesBrowserCheck2IDBrowserCheck:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesBrowserCheck2Value string

const (
	SettingEditResponseZonesBrowserCheck2ValueOn  SettingEditResponseZonesBrowserCheck2Value = "on"
	SettingEditResponseZonesBrowserCheck2ValueOff SettingEditResponseZonesBrowserCheck2Value = "off"
)

func (r SettingEditResponseZonesBrowserCheck2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesBrowserCheck2ValueOn, SettingEditResponseZonesBrowserCheck2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesBrowserCheck2Editable bool

const (
	SettingEditResponseZonesBrowserCheck2EditableTrue  SettingEditResponseZonesBrowserCheck2Editable = true
	SettingEditResponseZonesBrowserCheck2EditableFalse SettingEditResponseZonesBrowserCheck2Editable = false
)

func (r SettingEditResponseZonesBrowserCheck2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesBrowserCheck2EditableTrue, SettingEditResponseZonesBrowserCheck2EditableFalse:
		return true
	}
	return false
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type SettingEditResponseZonesCacheLevel2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesCacheLevel2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesCacheLevel2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesCacheLevel2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesCacheLevel2JSON `json:"-"`
}

// settingEditResponseZonesCacheLevel2JSON contains the JSON metadata for the
// struct [SettingEditResponseZonesCacheLevel2]
type settingEditResponseZonesCacheLevel2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesCacheLevel2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesCacheLevel2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesCacheLevel2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesCacheLevel2ID string

const (
	SettingEditResponseZonesCacheLevel2IDCacheLevel SettingEditResponseZonesCacheLevel2ID = "cache_level"
)

func (r SettingEditResponseZonesCacheLevel2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesCacheLevel2IDCacheLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesCacheLevel2Value string

const (
	SettingEditResponseZonesCacheLevel2ValueAggressive SettingEditResponseZonesCacheLevel2Value = "aggressive"
	SettingEditResponseZonesCacheLevel2ValueBasic      SettingEditResponseZonesCacheLevel2Value = "basic"
	SettingEditResponseZonesCacheLevel2ValueSimplified SettingEditResponseZonesCacheLevel2Value = "simplified"
)

func (r SettingEditResponseZonesCacheLevel2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesCacheLevel2ValueAggressive, SettingEditResponseZonesCacheLevel2ValueBasic, SettingEditResponseZonesCacheLevel2ValueSimplified:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesCacheLevel2Editable bool

const (
	SettingEditResponseZonesCacheLevel2EditableTrue  SettingEditResponseZonesCacheLevel2Editable = true
	SettingEditResponseZonesCacheLevel2EditableFalse SettingEditResponseZonesCacheLevel2Editable = false
)

func (r SettingEditResponseZonesCacheLevel2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesCacheLevel2EditableTrue, SettingEditResponseZonesCacheLevel2EditableFalse:
		return true
	}
	return false
}

// Determines whether or not the china network is enabled.
type SettingEditResponseZonesChinaNetworkEnabled struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesChinaNetworkEnabledID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesChinaNetworkEnabledValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesChinaNetworkEnabledEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesChinaNetworkEnabledJSON `json:"-"`
}

// settingEditResponseZonesChinaNetworkEnabledJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesChinaNetworkEnabled]
type settingEditResponseZonesChinaNetworkEnabledJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesChinaNetworkEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesChinaNetworkEnabledJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesChinaNetworkEnabled) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesChinaNetworkEnabledID string

const (
	SettingEditResponseZonesChinaNetworkEnabledIDChinaNetworkEnabled SettingEditResponseZonesChinaNetworkEnabledID = "china_network_enabled"
)

func (r SettingEditResponseZonesChinaNetworkEnabledID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesChinaNetworkEnabledIDChinaNetworkEnabled:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesChinaNetworkEnabledValue string

const (
	SettingEditResponseZonesChinaNetworkEnabledValueOn  SettingEditResponseZonesChinaNetworkEnabledValue = "on"
	SettingEditResponseZonesChinaNetworkEnabledValueOff SettingEditResponseZonesChinaNetworkEnabledValue = "off"
)

func (r SettingEditResponseZonesChinaNetworkEnabledValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesChinaNetworkEnabledValueOn, SettingEditResponseZonesChinaNetworkEnabledValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesChinaNetworkEnabledEditable bool

const (
	SettingEditResponseZonesChinaNetworkEnabledEditableTrue  SettingEditResponseZonesChinaNetworkEnabledEditable = true
	SettingEditResponseZonesChinaNetworkEnabledEditableFalse SettingEditResponseZonesChinaNetworkEnabledEditable = false
)

func (r SettingEditResponseZonesChinaNetworkEnabledEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesChinaNetworkEnabledEditableTrue, SettingEditResponseZonesChinaNetworkEnabledEditableFalse:
		return true
	}
	return false
}

// When enabled and the client sends an Accept header requesting text/markdown,
// Cloudflare will convert HTML responses to Markdown format using the toMarkdown()
// service. Refer to the
// [developer documentation](https://developers.cloudflare.com/workers-ai/features/markdown-conversion/)
// for more information.
type SettingEditResponseZonesContentConverter struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesContentConverterID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesContentConverterValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesContentConverterEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesContentConverterJSON `json:"-"`
}

// settingEditResponseZonesContentConverterJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesContentConverter]
type settingEditResponseZonesContentConverterJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesContentConverter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesContentConverterJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesContentConverter) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesContentConverterID string

const (
	SettingEditResponseZonesContentConverterIDContentConverter SettingEditResponseZonesContentConverterID = "content_converter"
)

func (r SettingEditResponseZonesContentConverterID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesContentConverterIDContentConverter:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesContentConverterValue string

const (
	SettingEditResponseZonesContentConverterValueOff SettingEditResponseZonesContentConverterValue = "off"
	SettingEditResponseZonesContentConverterValueOn  SettingEditResponseZonesContentConverterValue = "on"
)

func (r SettingEditResponseZonesContentConverterValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesContentConverterValueOff, SettingEditResponseZonesContentConverterValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesContentConverterEditable bool

const (
	SettingEditResponseZonesContentConverterEditableTrue  SettingEditResponseZonesContentConverterEditable = true
	SettingEditResponseZonesContentConverterEditableFalse SettingEditResponseZonesContentConverterEditable = false
)

func (r SettingEditResponseZonesContentConverterEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesContentConverterEditableTrue, SettingEditResponseZonesContentConverterEditableFalse:
		return true
	}
	return false
}

// Whether or not cname flattening is on.
//
// Deprecated: This zone setting is deprecated; please use the DNS Settings route
// instead. More information at
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2025-03-21
type SettingEditResponseZonesCNAMEFlattening struct {
	// How to flatten the cname destination.
	ID SettingEditResponseZonesCNAMEFlatteningID `json:"id" api:"required"`
	// Current value of the zone setting.
	//
	// Deprecated: This zone setting is deprecated; please use the DNS Settings route
	// instead. More information at
	// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2025-03-21
	Value SettingEditResponseZonesCNAMEFlatteningValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesCNAMEFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingEditResponseZonesCNAMEFlattening) implementsSettingEditResponse() {}

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
type SettingEditResponseZonesEdgeCacheTTL2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesEdgeCacheTTL2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesEdgeCacheTTL2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesEdgeCacheTTL2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesEdgeCacheTtl2JSON `json:"-"`
}

// settingEditResponseZonesEdgeCacheTtl2JSON contains the JSON metadata for the
// struct [SettingEditResponseZonesEdgeCacheTTL2]
type settingEditResponseZonesEdgeCacheTtl2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesEdgeCacheTTL2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesEdgeCacheTtl2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesEdgeCacheTTL2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesEdgeCacheTTL2ID string

const (
	SettingEditResponseZonesEdgeCacheTTL2IDEdgeCacheTTL SettingEditResponseZonesEdgeCacheTTL2ID = "edge_cache_ttl"
)

func (r SettingEditResponseZonesEdgeCacheTTL2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEdgeCacheTTL2IDEdgeCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesEdgeCacheTTL2Value float64

const (
	SettingEditResponseZonesEdgeCacheTTL2Value30     SettingEditResponseZonesEdgeCacheTTL2Value = 30
	SettingEditResponseZonesEdgeCacheTTL2Value60     SettingEditResponseZonesEdgeCacheTTL2Value = 60
	SettingEditResponseZonesEdgeCacheTTL2Value300    SettingEditResponseZonesEdgeCacheTTL2Value = 300
	SettingEditResponseZonesEdgeCacheTTL2Value1200   SettingEditResponseZonesEdgeCacheTTL2Value = 1200
	SettingEditResponseZonesEdgeCacheTTL2Value1800   SettingEditResponseZonesEdgeCacheTTL2Value = 1800
	SettingEditResponseZonesEdgeCacheTTL2Value3600   SettingEditResponseZonesEdgeCacheTTL2Value = 3600
	SettingEditResponseZonesEdgeCacheTTL2Value7200   SettingEditResponseZonesEdgeCacheTTL2Value = 7200
	SettingEditResponseZonesEdgeCacheTTL2Value10800  SettingEditResponseZonesEdgeCacheTTL2Value = 10800
	SettingEditResponseZonesEdgeCacheTTL2Value14400  SettingEditResponseZonesEdgeCacheTTL2Value = 14400
	SettingEditResponseZonesEdgeCacheTTL2Value18000  SettingEditResponseZonesEdgeCacheTTL2Value = 18000
	SettingEditResponseZonesEdgeCacheTTL2Value28800  SettingEditResponseZonesEdgeCacheTTL2Value = 28800
	SettingEditResponseZonesEdgeCacheTTL2Value43200  SettingEditResponseZonesEdgeCacheTTL2Value = 43200
	SettingEditResponseZonesEdgeCacheTTL2Value57600  SettingEditResponseZonesEdgeCacheTTL2Value = 57600
	SettingEditResponseZonesEdgeCacheTTL2Value72000  SettingEditResponseZonesEdgeCacheTTL2Value = 72000
	SettingEditResponseZonesEdgeCacheTTL2Value86400  SettingEditResponseZonesEdgeCacheTTL2Value = 86400
	SettingEditResponseZonesEdgeCacheTTL2Value172800 SettingEditResponseZonesEdgeCacheTTL2Value = 172800
	SettingEditResponseZonesEdgeCacheTTL2Value259200 SettingEditResponseZonesEdgeCacheTTL2Value = 259200
	SettingEditResponseZonesEdgeCacheTTL2Value345600 SettingEditResponseZonesEdgeCacheTTL2Value = 345600
	SettingEditResponseZonesEdgeCacheTTL2Value432000 SettingEditResponseZonesEdgeCacheTTL2Value = 432000
	SettingEditResponseZonesEdgeCacheTTL2Value518400 SettingEditResponseZonesEdgeCacheTTL2Value = 518400
	SettingEditResponseZonesEdgeCacheTTL2Value604800 SettingEditResponseZonesEdgeCacheTTL2Value = 604800
)

func (r SettingEditResponseZonesEdgeCacheTTL2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEdgeCacheTTL2Value30, SettingEditResponseZonesEdgeCacheTTL2Value60, SettingEditResponseZonesEdgeCacheTTL2Value300, SettingEditResponseZonesEdgeCacheTTL2Value1200, SettingEditResponseZonesEdgeCacheTTL2Value1800, SettingEditResponseZonesEdgeCacheTTL2Value3600, SettingEditResponseZonesEdgeCacheTTL2Value7200, SettingEditResponseZonesEdgeCacheTTL2Value10800, SettingEditResponseZonesEdgeCacheTTL2Value14400, SettingEditResponseZonesEdgeCacheTTL2Value18000, SettingEditResponseZonesEdgeCacheTTL2Value28800, SettingEditResponseZonesEdgeCacheTTL2Value43200, SettingEditResponseZonesEdgeCacheTTL2Value57600, SettingEditResponseZonesEdgeCacheTTL2Value72000, SettingEditResponseZonesEdgeCacheTTL2Value86400, SettingEditResponseZonesEdgeCacheTTL2Value172800, SettingEditResponseZonesEdgeCacheTTL2Value259200, SettingEditResponseZonesEdgeCacheTTL2Value345600, SettingEditResponseZonesEdgeCacheTTL2Value432000, SettingEditResponseZonesEdgeCacheTTL2Value518400, SettingEditResponseZonesEdgeCacheTTL2Value604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesEdgeCacheTTL2Editable bool

const (
	SettingEditResponseZonesEdgeCacheTTL2EditableTrue  SettingEditResponseZonesEdgeCacheTTL2Editable = true
	SettingEditResponseZonesEdgeCacheTTL2EditableFalse SettingEditResponseZonesEdgeCacheTTL2Editable = false
)

func (r SettingEditResponseZonesEdgeCacheTTL2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEdgeCacheTTL2EditableTrue, SettingEditResponseZonesEdgeCacheTTL2EditableFalse:
		return true
	}
	return false
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingEditResponseZonesEmailObfuscation2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesEmailObfuscation2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesEmailObfuscation2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesEmailObfuscation2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesEmailObfuscation2JSON `json:"-"`
}

// settingEditResponseZonesEmailObfuscation2JSON contains the JSON metadata for the
// struct [SettingEditResponseZonesEmailObfuscation2]
type settingEditResponseZonesEmailObfuscation2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesEmailObfuscation2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesEmailObfuscation2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesEmailObfuscation2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesEmailObfuscation2ID string

const (
	SettingEditResponseZonesEmailObfuscation2IDEmailObfuscation SettingEditResponseZonesEmailObfuscation2ID = "email_obfuscation"
)

func (r SettingEditResponseZonesEmailObfuscation2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEmailObfuscation2IDEmailObfuscation:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesEmailObfuscation2Value string

const (
	SettingEditResponseZonesEmailObfuscation2ValueOn  SettingEditResponseZonesEmailObfuscation2Value = "on"
	SettingEditResponseZonesEmailObfuscation2ValueOff SettingEditResponseZonesEmailObfuscation2Value = "off"
)

func (r SettingEditResponseZonesEmailObfuscation2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEmailObfuscation2ValueOn, SettingEditResponseZonesEmailObfuscation2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesEmailObfuscation2Editable bool

const (
	SettingEditResponseZonesEmailObfuscation2EditableTrue  SettingEditResponseZonesEmailObfuscation2Editable = true
	SettingEditResponseZonesEmailObfuscation2EditableFalse SettingEditResponseZonesEmailObfuscation2Editable = false
)

func (r SettingEditResponseZonesEmailObfuscation2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEmailObfuscation2EditableTrue, SettingEditResponseZonesEmailObfuscation2EditableFalse:
		return true
	}
	return false
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingEditResponseZonesIPGeolocation2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesIPGeolocation2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesIPGeolocation2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesIPGeolocation2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesIPGeolocation2JSON `json:"-"`
}

// settingEditResponseZonesIPGeolocation2JSON contains the JSON metadata for the
// struct [SettingEditResponseZonesIPGeolocation2]
type settingEditResponseZonesIPGeolocation2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesIPGeolocation2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesIPGeolocation2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesIPGeolocation2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesIPGeolocation2ID string

const (
	SettingEditResponseZonesIPGeolocation2IDIPGeolocation SettingEditResponseZonesIPGeolocation2ID = "ip_geolocation"
)

func (r SettingEditResponseZonesIPGeolocation2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesIPGeolocation2IDIPGeolocation:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesIPGeolocation2Value string

const (
	SettingEditResponseZonesIPGeolocation2ValueOn  SettingEditResponseZonesIPGeolocation2Value = "on"
	SettingEditResponseZonesIPGeolocation2ValueOff SettingEditResponseZonesIPGeolocation2Value = "off"
)

func (r SettingEditResponseZonesIPGeolocation2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesIPGeolocation2ValueOn, SettingEditResponseZonesIPGeolocation2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesIPGeolocation2Editable bool

const (
	SettingEditResponseZonesIPGeolocation2EditableTrue  SettingEditResponseZonesIPGeolocation2Editable = true
	SettingEditResponseZonesIPGeolocation2EditableFalse SettingEditResponseZonesIPGeolocation2Editable = false
)

func (r SettingEditResponseZonesIPGeolocation2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesIPGeolocation2EditableTrue, SettingEditResponseZonesIPGeolocation2EditableFalse:
		return true
	}
	return false
}

// Maximum size of an allowable upload.
type SettingEditResponseZonesMaxUpload struct {
	// identifier of the zone setting.
	ID SettingEditResponseZonesMaxUploadID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesMaxUploadValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingEditResponseZonesMaxUpload) implementsSettingEditResponse() {}

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
type SettingEditResponseZonesMaxUploadValue int64

const (
	SettingEditResponseZonesMaxUploadValue100  SettingEditResponseZonesMaxUploadValue = 100
	SettingEditResponseZonesMaxUploadValue125  SettingEditResponseZonesMaxUploadValue = 125
	SettingEditResponseZonesMaxUploadValue150  SettingEditResponseZonesMaxUploadValue = 150
	SettingEditResponseZonesMaxUploadValue175  SettingEditResponseZonesMaxUploadValue = 175
	SettingEditResponseZonesMaxUploadValue200  SettingEditResponseZonesMaxUploadValue = 200
	SettingEditResponseZonesMaxUploadValue225  SettingEditResponseZonesMaxUploadValue = 225
	SettingEditResponseZonesMaxUploadValue250  SettingEditResponseZonesMaxUploadValue = 250
	SettingEditResponseZonesMaxUploadValue275  SettingEditResponseZonesMaxUploadValue = 275
	SettingEditResponseZonesMaxUploadValue300  SettingEditResponseZonesMaxUploadValue = 300
	SettingEditResponseZonesMaxUploadValue325  SettingEditResponseZonesMaxUploadValue = 325
	SettingEditResponseZonesMaxUploadValue350  SettingEditResponseZonesMaxUploadValue = 350
	SettingEditResponseZonesMaxUploadValue375  SettingEditResponseZonesMaxUploadValue = 375
	SettingEditResponseZonesMaxUploadValue400  SettingEditResponseZonesMaxUploadValue = 400
	SettingEditResponseZonesMaxUploadValue425  SettingEditResponseZonesMaxUploadValue = 425
	SettingEditResponseZonesMaxUploadValue450  SettingEditResponseZonesMaxUploadValue = 450
	SettingEditResponseZonesMaxUploadValue475  SettingEditResponseZonesMaxUploadValue = 475
	SettingEditResponseZonesMaxUploadValue500  SettingEditResponseZonesMaxUploadValue = 500
	SettingEditResponseZonesMaxUploadValue1000 SettingEditResponseZonesMaxUploadValue = 1000
)

func (r SettingEditResponseZonesMaxUploadValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesMaxUploadValue100, SettingEditResponseZonesMaxUploadValue125, SettingEditResponseZonesMaxUploadValue150, SettingEditResponseZonesMaxUploadValue175, SettingEditResponseZonesMaxUploadValue200, SettingEditResponseZonesMaxUploadValue225, SettingEditResponseZonesMaxUploadValue250, SettingEditResponseZonesMaxUploadValue275, SettingEditResponseZonesMaxUploadValue300, SettingEditResponseZonesMaxUploadValue325, SettingEditResponseZonesMaxUploadValue350, SettingEditResponseZonesMaxUploadValue375, SettingEditResponseZonesMaxUploadValue400, SettingEditResponseZonesMaxUploadValue425, SettingEditResponseZonesMaxUploadValue450, SettingEditResponseZonesMaxUploadValue475, SettingEditResponseZonesMaxUploadValue500, SettingEditResponseZonesMaxUploadValue1000:
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
//
// Deprecated: Mirage is being deprecated. More information at
// https://developers.cloudflare.com/speed/optimization/images/mirage/
type SettingEditResponseZonesMirage2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesMirage2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	//
	// Deprecated: Mirage is being deprecated. More information at
	// https://developers.cloudflare.com/speed/optimization/images/mirage/
	Value SettingEditResponseZonesMirage2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesMirage2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesMirage2JSON `json:"-"`
}

// settingEditResponseZonesMirage2JSON contains the JSON metadata for the struct
// [SettingEditResponseZonesMirage2]
type settingEditResponseZonesMirage2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesMirage2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesMirage2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesMirage2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesMirage2ID string

const (
	SettingEditResponseZonesMirage2IDMirage SettingEditResponseZonesMirage2ID = "mirage"
)

func (r SettingEditResponseZonesMirage2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesMirage2IDMirage:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesMirage2Value string

const (
	SettingEditResponseZonesMirage2ValueOn  SettingEditResponseZonesMirage2Value = "on"
	SettingEditResponseZonesMirage2ValueOff SettingEditResponseZonesMirage2Value = "off"
)

func (r SettingEditResponseZonesMirage2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesMirage2ValueOn, SettingEditResponseZonesMirage2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesMirage2Editable bool

const (
	SettingEditResponseZonesMirage2EditableTrue  SettingEditResponseZonesMirage2Editable = true
	SettingEditResponseZonesMirage2EditableFalse SettingEditResponseZonesMirage2Editable = false
)

func (r SettingEditResponseZonesMirage2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesMirage2EditableTrue, SettingEditResponseZonesMirage2EditableFalse:
		return true
	}
	return false
}

// Enables the Opportunistic Encryption feature for a zone.
type SettingEditResponseZonesOpportunisticEncryption2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesOpportunisticEncryption2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesOpportunisticEncryption2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesOpportunisticEncryption2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                            `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesOpportunisticEncryption2JSON `json:"-"`
}

// settingEditResponseZonesOpportunisticEncryption2JSON contains the JSON metadata
// for the struct [SettingEditResponseZonesOpportunisticEncryption2]
type settingEditResponseZonesOpportunisticEncryption2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesOpportunisticEncryption2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesOpportunisticEncryption2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesOpportunisticEncryption2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesOpportunisticEncryption2ID string

const (
	SettingEditResponseZonesOpportunisticEncryption2IDOpportunisticEncryption SettingEditResponseZonesOpportunisticEncryption2ID = "opportunistic_encryption"
)

func (r SettingEditResponseZonesOpportunisticEncryption2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesOpportunisticEncryption2IDOpportunisticEncryption:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesOpportunisticEncryption2Value string

const (
	SettingEditResponseZonesOpportunisticEncryption2ValueOn  SettingEditResponseZonesOpportunisticEncryption2Value = "on"
	SettingEditResponseZonesOpportunisticEncryption2ValueOff SettingEditResponseZonesOpportunisticEncryption2Value = "off"
)

func (r SettingEditResponseZonesOpportunisticEncryption2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesOpportunisticEncryption2ValueOn, SettingEditResponseZonesOpportunisticEncryption2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesOpportunisticEncryption2Editable bool

const (
	SettingEditResponseZonesOpportunisticEncryption2EditableTrue  SettingEditResponseZonesOpportunisticEncryption2Editable = true
	SettingEditResponseZonesOpportunisticEncryption2EditableFalse SettingEditResponseZonesOpportunisticEncryption2Editable = false
)

func (r SettingEditResponseZonesOpportunisticEncryption2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesOpportunisticEncryption2EditableTrue, SettingEditResponseZonesOpportunisticEncryption2EditableFalse:
		return true
	}
	return false
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingEditResponseZonesOriginErrorPagePassThru2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesOriginErrorPagePassThru2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesOriginErrorPagePassThru2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesOriginErrorPagePassThru2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                            `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesOriginErrorPagePassThru2JSON `json:"-"`
}

// settingEditResponseZonesOriginErrorPagePassThru2JSON contains the JSON metadata
// for the struct [SettingEditResponseZonesOriginErrorPagePassThru2]
type settingEditResponseZonesOriginErrorPagePassThru2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesOriginErrorPagePassThru2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesOriginErrorPagePassThru2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesOriginErrorPagePassThru2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesOriginErrorPagePassThru2ID string

const (
	SettingEditResponseZonesOriginErrorPagePassThru2IDOriginErrorPagePassThru SettingEditResponseZonesOriginErrorPagePassThru2ID = "origin_error_page_pass_thru"
)

func (r SettingEditResponseZonesOriginErrorPagePassThru2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesOriginErrorPagePassThru2IDOriginErrorPagePassThru:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesOriginErrorPagePassThru2Value string

const (
	SettingEditResponseZonesOriginErrorPagePassThru2ValueOn  SettingEditResponseZonesOriginErrorPagePassThru2Value = "on"
	SettingEditResponseZonesOriginErrorPagePassThru2ValueOff SettingEditResponseZonesOriginErrorPagePassThru2Value = "off"
)

func (r SettingEditResponseZonesOriginErrorPagePassThru2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesOriginErrorPagePassThru2ValueOn, SettingEditResponseZonesOriginErrorPagePassThru2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesOriginErrorPagePassThru2Editable bool

const (
	SettingEditResponseZonesOriginErrorPagePassThru2EditableTrue  SettingEditResponseZonesOriginErrorPagePassThru2Editable = true
	SettingEditResponseZonesOriginErrorPagePassThru2EditableFalse SettingEditResponseZonesOriginErrorPagePassThru2Editable = false
)

func (r SettingEditResponseZonesOriginErrorPagePassThru2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesOriginErrorPagePassThru2EditableTrue, SettingEditResponseZonesOriginErrorPagePassThru2EditableFalse:
		return true
	}
	return false
}

// Origin H2 Max Streams configures the max number of concurrent requests that
// Cloudflare will send within the same connection when communicating with the
// origin server, if the origin supports it. Note that if your origin does not
// support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also
// note that the default value is `100` for all plan types except Enterprise where
// it is `1`. `1` means that H2 multiplexing is disabled.
type SettingEditResponseZonesCacheRulesOriginH2MaxStreams struct {
	// Value of the zone setting.
	ID SettingEditResponseZonesCacheRulesOriginH2MaxStreamsID `json:"id" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Value of the Origin H2 Max Streams Setting.
	Value int64                                                    `json:"value"`
	JSON  settingEditResponseZonesCacheRulesOriginH2MaxStreamsJSON `json:"-"`
}

// settingEditResponseZonesCacheRulesOriginH2MaxStreamsJSON contains the JSON
// metadata for the struct [SettingEditResponseZonesCacheRulesOriginH2MaxStreams]
type settingEditResponseZonesCacheRulesOriginH2MaxStreamsJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesCacheRulesOriginH2MaxStreams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesCacheRulesOriginH2MaxStreamsJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesCacheRulesOriginH2MaxStreams) implementsSettingEditResponse() {}

// Value of the zone setting.
type SettingEditResponseZonesCacheRulesOriginH2MaxStreamsID string

const (
	SettingEditResponseZonesCacheRulesOriginH2MaxStreamsIDOriginH2MaxStreams SettingEditResponseZonesCacheRulesOriginH2MaxStreamsID = "origin_h2_max_streams"
)

func (r SettingEditResponseZonesCacheRulesOriginH2MaxStreamsID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesCacheRulesOriginH2MaxStreamsIDOriginH2MaxStreams:
		return true
	}
	return false
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except
// Enterprise where it is "1".
type SettingEditResponseZonesCacheRulesOriginMaxHTTPVersion struct {
	// Value of the zone setting.
	ID SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionID `json:"id" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Value of the Origin Max HTTP Version Setting.
	Value SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionValue `json:"value"`
	JSON  settingEditResponseZonesCacheRulesOriginMaxHTTPVersionJSON  `json:"-"`
}

// settingEditResponseZonesCacheRulesOriginMaxHTTPVersionJSON contains the JSON
// metadata for the struct [SettingEditResponseZonesCacheRulesOriginMaxHTTPVersion]
type settingEditResponseZonesCacheRulesOriginMaxHTTPVersionJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesCacheRulesOriginMaxHTTPVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesCacheRulesOriginMaxHTTPVersionJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesCacheRulesOriginMaxHTTPVersion) implementsSettingEditResponse() {}

// Value of the zone setting.
type SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionID string

const (
	SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionIDOriginMaxHTTPVersion SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionID = "origin_max_http_version"
)

func (r SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionIDOriginMaxHTTPVersion:
		return true
	}
	return false
}

// Value of the Origin Max HTTP Version Setting.
type SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionValue string

const (
	SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionValue2 SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionValue = "2"
	SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionValue1 SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionValue = "1"
)

func (r SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionValue2, SettingEditResponseZonesCacheRulesOriginMaxHTTPVersionValue1:
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
type SettingEditResponseZonesPolish2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesPolish2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesPolish2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesPolish2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesPolish2JSON `json:"-"`
}

// settingEditResponseZonesPolish2JSON contains the JSON metadata for the struct
// [SettingEditResponseZonesPolish2]
type settingEditResponseZonesPolish2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesPolish2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesPolish2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesPolish2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesPolish2ID string

const (
	SettingEditResponseZonesPolish2IDPolish SettingEditResponseZonesPolish2ID = "polish"
)

func (r SettingEditResponseZonesPolish2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesPolish2IDPolish:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesPolish2Value string

const (
	SettingEditResponseZonesPolish2ValueOff      SettingEditResponseZonesPolish2Value = "off"
	SettingEditResponseZonesPolish2ValueLossless SettingEditResponseZonesPolish2Value = "lossless"
	SettingEditResponseZonesPolish2ValueLossy    SettingEditResponseZonesPolish2Value = "lossy"
)

func (r SettingEditResponseZonesPolish2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesPolish2ValueOff, SettingEditResponseZonesPolish2ValueLossless, SettingEditResponseZonesPolish2ValueLossy:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesPolish2Editable bool

const (
	SettingEditResponseZonesPolish2EditableTrue  SettingEditResponseZonesPolish2Editable = true
	SettingEditResponseZonesPolish2EditableFalse SettingEditResponseZonesPolish2Editable = false
)

func (r SettingEditResponseZonesPolish2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesPolish2EditableTrue, SettingEditResponseZonesPolish2EditableFalse:
		return true
	}
	return false
}

// Privacy Pass v1 was a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors by allowing users to reduce
// the number of CAPTCHAs shown.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
//
// Deprecated: Privacy Pass v1 was deprecated in 2023. (Announcement -
// https://blog.cloudflare.com/privacy-pass-standard/) and (API deprecation
// details -
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2024-03-31)
type SettingEditResponseZonesPrivacyPass struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesPrivacyPassID `json:"id" api:"required"`
	// Current value of the zone setting.
	//
	// Deprecated: Privacy Pass v1 was deprecated in 2023. (Announcement -
	// https://blog.cloudflare.com/privacy-pass-standard/) and (API deprecation
	// details -
	// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2024-03-31)
	Value SettingEditResponseZonesPrivacyPassValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesPrivacyPassEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesPrivacyPassJSON `json:"-"`
}

// settingEditResponseZonesPrivacyPassJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesPrivacyPass]
type settingEditResponseZonesPrivacyPassJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesPrivacyPass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesPrivacyPassJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesPrivacyPass) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesPrivacyPassID string

const (
	SettingEditResponseZonesPrivacyPassIDPrivacyPass SettingEditResponseZonesPrivacyPassID = "privacy_pass"
)

func (r SettingEditResponseZonesPrivacyPassID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesPrivacyPassIDPrivacyPass:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesPrivacyPassValue string

const (
	SettingEditResponseZonesPrivacyPassValueOn  SettingEditResponseZonesPrivacyPassValue = "on"
	SettingEditResponseZonesPrivacyPassValueOff SettingEditResponseZonesPrivacyPassValue = "off"
)

func (r SettingEditResponseZonesPrivacyPassValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesPrivacyPassValueOn, SettingEditResponseZonesPrivacyPassValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesPrivacyPassEditable bool

const (
	SettingEditResponseZonesPrivacyPassEditableTrue  SettingEditResponseZonesPrivacyPassEditable = true
	SettingEditResponseZonesPrivacyPassEditableFalse SettingEditResponseZonesPrivacyPassEditable = false
)

func (r SettingEditResponseZonesPrivacyPassEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesPrivacyPassEditableTrue, SettingEditResponseZonesPrivacyPassEditableFalse:
		return true
	}
	return false
}

// When enabled, Cloudflare will redirect verified AI training crawlers to
// canonical URLs found in the HTML response, ensuring AI models train on
// authoritative content.
type SettingEditResponseZonesRedirectsForAITraining struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesRedirectsForAITrainingID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesRedirectsForAITrainingValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesRedirectsForAITrainingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesRedirectsForAITrainingJSON `json:"-"`
}

// settingEditResponseZonesRedirectsForAITrainingJSON contains the JSON metadata
// for the struct [SettingEditResponseZonesRedirectsForAITraining]
type settingEditResponseZonesRedirectsForAITrainingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesRedirectsForAITraining) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesRedirectsForAITrainingJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesRedirectsForAITraining) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesRedirectsForAITrainingID string

const (
	SettingEditResponseZonesRedirectsForAITrainingIDRedirectsForAITraining SettingEditResponseZonesRedirectsForAITrainingID = "redirects_for_ai_training"
)

func (r SettingEditResponseZonesRedirectsForAITrainingID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesRedirectsForAITrainingIDRedirectsForAITraining:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesRedirectsForAITrainingValue string

const (
	SettingEditResponseZonesRedirectsForAITrainingValueOff SettingEditResponseZonesRedirectsForAITrainingValue = "off"
	SettingEditResponseZonesRedirectsForAITrainingValueOn  SettingEditResponseZonesRedirectsForAITrainingValue = "on"
)

func (r SettingEditResponseZonesRedirectsForAITrainingValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesRedirectsForAITrainingValueOff, SettingEditResponseZonesRedirectsForAITrainingValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesRedirectsForAITrainingEditable bool

const (
	SettingEditResponseZonesRedirectsForAITrainingEditableTrue  SettingEditResponseZonesRedirectsForAITrainingEditable = true
	SettingEditResponseZonesRedirectsForAITrainingEditableFalse SettingEditResponseZonesRedirectsForAITrainingEditable = false
)

func (r SettingEditResponseZonesRedirectsForAITrainingEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesRedirectsForAITrainingEditableTrue, SettingEditResponseZonesRedirectsForAITrainingEditableFalse:
		return true
	}
	return false
}

// Automatically replace insecure JavaScript libraries with safer and faster
// alternatives provided under cdnjs and powered by Cloudflare. Currently supports
// the following libraries: Polyfill under polyfill.io.
type SettingEditResponseZonesReplaceInsecureJS struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesReplaceInsecureJSID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesReplaceInsecureJSValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesReplaceInsecureJSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingEditResponseZonesReplaceInsecureJS) implementsSettingEditResponse() {}

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
//
// Deprecated: This zone setting is deprecated. This functionality is no longer
// supported.
type SettingEditResponseZonesResponseBuffering2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesResponseBuffering2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesResponseBuffering2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesResponseBuffering2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesResponseBuffering2JSON `json:"-"`
}

// settingEditResponseZonesResponseBuffering2JSON contains the JSON metadata for
// the struct [SettingEditResponseZonesResponseBuffering2]
type settingEditResponseZonesResponseBuffering2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesResponseBuffering2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesResponseBuffering2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesResponseBuffering2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesResponseBuffering2ID string

const (
	SettingEditResponseZonesResponseBuffering2IDResponseBuffering SettingEditResponseZonesResponseBuffering2ID = "response_buffering"
)

func (r SettingEditResponseZonesResponseBuffering2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesResponseBuffering2IDResponseBuffering:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesResponseBuffering2Value string

const (
	SettingEditResponseZonesResponseBuffering2ValueOn  SettingEditResponseZonesResponseBuffering2Value = "on"
	SettingEditResponseZonesResponseBuffering2ValueOff SettingEditResponseZonesResponseBuffering2Value = "off"
)

func (r SettingEditResponseZonesResponseBuffering2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesResponseBuffering2ValueOn, SettingEditResponseZonesResponseBuffering2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesResponseBuffering2Editable bool

const (
	SettingEditResponseZonesResponseBuffering2EditableTrue  SettingEditResponseZonesResponseBuffering2Editable = true
	SettingEditResponseZonesResponseBuffering2EditableFalse SettingEditResponseZonesResponseBuffering2Editable = false
)

func (r SettingEditResponseZonesResponseBuffering2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesResponseBuffering2EditableTrue, SettingEditResponseZonesResponseBuffering2EditableFalse:
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
type SettingEditResponseZonesRocketLoader2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesRocketLoader2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesRocketLoader2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesRocketLoader2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesRocketLoader2JSON `json:"-"`
}

// settingEditResponseZonesRocketLoader2JSON contains the JSON metadata for the
// struct [SettingEditResponseZonesRocketLoader2]
type settingEditResponseZonesRocketLoader2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesRocketLoader2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesRocketLoader2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesRocketLoader2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesRocketLoader2ID string

const (
	SettingEditResponseZonesRocketLoader2IDRocketLoader SettingEditResponseZonesRocketLoader2ID = "rocket_loader"
)

func (r SettingEditResponseZonesRocketLoader2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesRocketLoader2IDRocketLoader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesRocketLoader2Value string

const (
	SettingEditResponseZonesRocketLoader2ValueOn  SettingEditResponseZonesRocketLoader2Value = "on"
	SettingEditResponseZonesRocketLoader2ValueOff SettingEditResponseZonesRocketLoader2Value = "off"
)

func (r SettingEditResponseZonesRocketLoader2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesRocketLoader2ValueOn, SettingEditResponseZonesRocketLoader2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesRocketLoader2Editable bool

const (
	SettingEditResponseZonesRocketLoader2EditableTrue  SettingEditResponseZonesRocketLoader2Editable = true
	SettingEditResponseZonesRocketLoader2EditableFalse SettingEditResponseZonesRocketLoader2Editable = false
)

func (r SettingEditResponseZonesRocketLoader2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesRocketLoader2EditableTrue, SettingEditResponseZonesRocketLoader2EditableFalse:
		return true
	}
	return false
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingEditResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value AutomaticPlatformOptimization `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                        `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingEditResponseZonesSchemasAutomaticPlatformOptimization) implementsSettingEditResponse() {
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

// When enabled, Cloudflare provisions an AI Search instance for the zone and
// exposes a /.well-known/ai-search endpoint that AI agents can query. Markdown
// responses also receive an agent: YAML capability block advertising the search
// endpoint.
type SettingEditResponseZonesSearchForAgents struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSearchForAgentsID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSearchForAgentsValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSearchForAgentsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesSearchForAgentsJSON `json:"-"`
}

// settingEditResponseZonesSearchForAgentsJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesSearchForAgents]
type settingEditResponseZonesSearchForAgentsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSearchForAgents) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSearchForAgentsJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSearchForAgents) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSearchForAgentsID string

const (
	SettingEditResponseZonesSearchForAgentsIDSearchForAgents SettingEditResponseZonesSearchForAgentsID = "search_for_agents"
)

func (r SettingEditResponseZonesSearchForAgentsID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSearchForAgentsIDSearchForAgents:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSearchForAgentsValue string

const (
	SettingEditResponseZonesSearchForAgentsValueOff SettingEditResponseZonesSearchForAgentsValue = "off"
	SettingEditResponseZonesSearchForAgentsValueOn  SettingEditResponseZonesSearchForAgentsValue = "on"
)

func (r SettingEditResponseZonesSearchForAgentsValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSearchForAgentsValueOff, SettingEditResponseZonesSearchForAgentsValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSearchForAgentsEditable bool

const (
	SettingEditResponseZonesSearchForAgentsEditableTrue  SettingEditResponseZonesSearchForAgentsEditable = true
	SettingEditResponseZonesSearchForAgentsEditableFalse SettingEditResponseZonesSearchForAgentsEditable = false
)

func (r SettingEditResponseZonesSearchForAgentsEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSearchForAgentsEditableTrue, SettingEditResponseZonesSearchForAgentsEditableFalse:
		return true
	}
	return false
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingEditResponseZonesSecurityLevel2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSecurityLevel2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSecurityLevel2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSecurityLevel2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesSecurityLevel2JSON `json:"-"`
}

// settingEditResponseZonesSecurityLevel2JSON contains the JSON metadata for the
// struct [SettingEditResponseZonesSecurityLevel2]
type settingEditResponseZonesSecurityLevel2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSecurityLevel2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSecurityLevel2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSecurityLevel2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSecurityLevel2ID string

const (
	SettingEditResponseZonesSecurityLevel2IDSecurityLevel SettingEditResponseZonesSecurityLevel2ID = "security_level"
)

func (r SettingEditResponseZonesSecurityLevel2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSecurityLevel2IDSecurityLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSecurityLevel2Value string

const (
	SettingEditResponseZonesSecurityLevel2ValueOff            SettingEditResponseZonesSecurityLevel2Value = "off"
	SettingEditResponseZonesSecurityLevel2ValueEssentiallyOff SettingEditResponseZonesSecurityLevel2Value = "essentially_off"
	SettingEditResponseZonesSecurityLevel2ValueLow            SettingEditResponseZonesSecurityLevel2Value = "low"
	SettingEditResponseZonesSecurityLevel2ValueMedium         SettingEditResponseZonesSecurityLevel2Value = "medium"
	SettingEditResponseZonesSecurityLevel2ValueHigh           SettingEditResponseZonesSecurityLevel2Value = "high"
	SettingEditResponseZonesSecurityLevel2ValueUnderAttack    SettingEditResponseZonesSecurityLevel2Value = "under_attack"
)

func (r SettingEditResponseZonesSecurityLevel2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSecurityLevel2ValueOff, SettingEditResponseZonesSecurityLevel2ValueEssentiallyOff, SettingEditResponseZonesSecurityLevel2ValueLow, SettingEditResponseZonesSecurityLevel2ValueMedium, SettingEditResponseZonesSecurityLevel2ValueHigh, SettingEditResponseZonesSecurityLevel2ValueUnderAttack:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSecurityLevel2Editable bool

const (
	SettingEditResponseZonesSecurityLevel2EditableTrue  SettingEditResponseZonesSecurityLevel2Editable = true
	SettingEditResponseZonesSecurityLevel2EditableFalse SettingEditResponseZonesSecurityLevel2Editable = false
)

func (r SettingEditResponseZonesSecurityLevel2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSecurityLevel2EditableTrue, SettingEditResponseZonesSecurityLevel2EditableFalse:
		return true
	}
	return false
}

// Allow SHA1 support.
type SettingEditResponseZonesSha1Support struct {
	// Zone setting identifier.
	ID SettingEditResponseZonesSha1SupportID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSha1SupportValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingEditResponseZonesSha1Support) implementsSettingEditResponse() {}

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
type SettingEditResponseZonesSortQueryStringForCache2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSortQueryStringForCache2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSortQueryStringForCache2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSortQueryStringForCache2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                            `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesSortQueryStringForCache2JSON `json:"-"`
}

// settingEditResponseZonesSortQueryStringForCache2JSON contains the JSON metadata
// for the struct [SettingEditResponseZonesSortQueryStringForCache2]
type settingEditResponseZonesSortQueryStringForCache2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSortQueryStringForCache2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSortQueryStringForCache2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSortQueryStringForCache2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSortQueryStringForCache2ID string

const (
	SettingEditResponseZonesSortQueryStringForCache2IDSortQueryStringForCache SettingEditResponseZonesSortQueryStringForCache2ID = "sort_query_string_for_cache"
)

func (r SettingEditResponseZonesSortQueryStringForCache2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSortQueryStringForCache2IDSortQueryStringForCache:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSortQueryStringForCache2Value string

const (
	SettingEditResponseZonesSortQueryStringForCache2ValueOn  SettingEditResponseZonesSortQueryStringForCache2Value = "on"
	SettingEditResponseZonesSortQueryStringForCache2ValueOff SettingEditResponseZonesSortQueryStringForCache2Value = "off"
)

func (r SettingEditResponseZonesSortQueryStringForCache2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSortQueryStringForCache2ValueOn, SettingEditResponseZonesSortQueryStringForCache2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSortQueryStringForCache2Editable bool

const (
	SettingEditResponseZonesSortQueryStringForCache2EditableTrue  SettingEditResponseZonesSortQueryStringForCache2Editable = true
	SettingEditResponseZonesSortQueryStringForCache2EditableFalse SettingEditResponseZonesSortQueryStringForCache2Editable = false
)

func (r SettingEditResponseZonesSortQueryStringForCache2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSortQueryStringForCache2EditableTrue, SettingEditResponseZonesSortQueryStringForCache2EditableFalse:
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
type SettingEditResponseZonesSSL2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSSL2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSSL2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSSL2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesSsl2JSON `json:"-"`
}

// settingEditResponseZonesSsl2JSON contains the JSON metadata for the struct
// [SettingEditResponseZonesSSL2]
type settingEditResponseZonesSsl2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSSL2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesSsl2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSSL2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSSL2ID string

const (
	SettingEditResponseZonesSSL2IDSSL SettingEditResponseZonesSSL2ID = "ssl"
)

func (r SettingEditResponseZonesSSL2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSSL2IDSSL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesSSL2Value string

const (
	SettingEditResponseZonesSSL2ValueOff      SettingEditResponseZonesSSL2Value = "off"
	SettingEditResponseZonesSSL2ValueFlexible SettingEditResponseZonesSSL2Value = "flexible"
	SettingEditResponseZonesSSL2ValueFull     SettingEditResponseZonesSSL2Value = "full"
	SettingEditResponseZonesSSL2ValueStrict   SettingEditResponseZonesSSL2Value = "strict"
)

func (r SettingEditResponseZonesSSL2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSSL2ValueOff, SettingEditResponseZonesSSL2ValueFlexible, SettingEditResponseZonesSSL2ValueFull, SettingEditResponseZonesSSL2ValueStrict:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSSL2Editable bool

const (
	SettingEditResponseZonesSSL2EditableTrue  SettingEditResponseZonesSSL2Editable = true
	SettingEditResponseZonesSSL2EditableFalse SettingEditResponseZonesSSL2Editable = false
)

func (r SettingEditResponseZonesSSL2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesSSL2EditableTrue, SettingEditResponseZonesSSL2EditableFalse:
		return true
	}
	return false
}

// Only allows TLS1.2.
type SettingEditResponseZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID SettingEditResponseZonesTLS1_2OnlyID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesTLS1_2OnlyValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesTLS1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingEditResponseZonesTLS1_2Only) implementsSettingEditResponse() {}

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

// Media Transformations provides on-demand resizing, conversion and optimization
// for images and video served through Cloudflare's network. Refer to the
// [Image Transformations](https://developers.cloudflare.com/images/) and
// [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started)
// documentation for more information.
type SettingEditResponseZonesTransformations struct {
	// ID of the zone setting. Shared between Image Transformations and Video
	// Transformations.
	ID SettingEditResponseZonesTransformationsID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesTransformationsValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesTransformationsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesTransformationsJSON `json:"-"`
}

// settingEditResponseZonesTransformationsJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesTransformations]
type settingEditResponseZonesTransformationsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesTransformations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesTransformationsJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesTransformations) implementsSettingEditResponse() {}

// ID of the zone setting. Shared between Image Transformations and Video
// Transformations.
type SettingEditResponseZonesTransformationsID string

const (
	SettingEditResponseZonesTransformationsIDTransformations SettingEditResponseZonesTransformationsID = "transformations"
)

func (r SettingEditResponseZonesTransformationsID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTransformationsIDTransformations:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesTransformationsValue string

const (
	SettingEditResponseZonesTransformationsValueOn   SettingEditResponseZonesTransformationsValue = "on"
	SettingEditResponseZonesTransformationsValueOff  SettingEditResponseZonesTransformationsValue = "off"
	SettingEditResponseZonesTransformationsValueOpen SettingEditResponseZonesTransformationsValue = "open"
)

func (r SettingEditResponseZonesTransformationsValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTransformationsValueOn, SettingEditResponseZonesTransformationsValueOff, SettingEditResponseZonesTransformationsValueOpen:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesTransformationsEditable bool

const (
	SettingEditResponseZonesTransformationsEditableTrue  SettingEditResponseZonesTransformationsEditable = true
	SettingEditResponseZonesTransformationsEditableFalse SettingEditResponseZonesTransformationsEditable = false
)

func (r SettingEditResponseZonesTransformationsEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTransformationsEditableTrue, SettingEditResponseZonesTransformationsEditableFalse:
		return true
	}
	return false
}

// Media Transformations Allowed Origins restricts transformations for images and
// video served through Cloudflare's network. Refer to the
// [Image Transformations](https://developers.cloudflare.com/images/) and
// [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started)
// documentation for more information.
type SettingEditResponseZonesTransformationsAllowedOrigins struct {
	// ID of the zone setting. Shared between Image Transformations and Video
	// Transformations.
	ID SettingEditResponseZonesTransformationsAllowedOriginsID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value string `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesTransformationsAllowedOriginsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                 `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesTransformationsAllowedOriginsJSON `json:"-"`
}

// settingEditResponseZonesTransformationsAllowedOriginsJSON contains the JSON
// metadata for the struct [SettingEditResponseZonesTransformationsAllowedOrigins]
type settingEditResponseZonesTransformationsAllowedOriginsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesTransformationsAllowedOrigins) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesTransformationsAllowedOriginsJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesTransformationsAllowedOrigins) implementsSettingEditResponse() {}

// ID of the zone setting. Shared between Image Transformations and Video
// Transformations.
type SettingEditResponseZonesTransformationsAllowedOriginsID string

const (
	SettingEditResponseZonesTransformationsAllowedOriginsIDTransformationsAllowedOrigins SettingEditResponseZonesTransformationsAllowedOriginsID = "transformations_allowed_origins"
)

func (r SettingEditResponseZonesTransformationsAllowedOriginsID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTransformationsAllowedOriginsIDTransformationsAllowedOrigins:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesTransformationsAllowedOriginsEditable bool

const (
	SettingEditResponseZonesTransformationsAllowedOriginsEditableTrue  SettingEditResponseZonesTransformationsAllowedOriginsEditable = true
	SettingEditResponseZonesTransformationsAllowedOriginsEditableFalse SettingEditResponseZonesTransformationsAllowedOriginsEditable = false
)

func (r SettingEditResponseZonesTransformationsAllowedOriginsEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTransformationsAllowedOriginsEditableTrue, SettingEditResponseZonesTransformationsAllowedOriginsEditableFalse:
		return true
	}
	return false
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingEditResponseZonesTrueClientIPHeader2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesTrueClientIPHeader2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesTrueClientIPHeader2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesTrueClientIPHeader2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesTrueClientIPHeader2JSON `json:"-"`
}

// settingEditResponseZonesTrueClientIPHeader2JSON contains the JSON metadata for
// the struct [SettingEditResponseZonesTrueClientIPHeader2]
type settingEditResponseZonesTrueClientIPHeader2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesTrueClientIPHeader2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesTrueClientIPHeader2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesTrueClientIPHeader2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesTrueClientIPHeader2ID string

const (
	SettingEditResponseZonesTrueClientIPHeader2IDTrueClientIPHeader SettingEditResponseZonesTrueClientIPHeader2ID = "true_client_ip_header"
)

func (r SettingEditResponseZonesTrueClientIPHeader2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTrueClientIPHeader2IDTrueClientIPHeader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesTrueClientIPHeader2Value string

const (
	SettingEditResponseZonesTrueClientIPHeader2ValueOn  SettingEditResponseZonesTrueClientIPHeader2Value = "on"
	SettingEditResponseZonesTrueClientIPHeader2ValueOff SettingEditResponseZonesTrueClientIPHeader2Value = "off"
)

func (r SettingEditResponseZonesTrueClientIPHeader2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTrueClientIPHeader2ValueOn, SettingEditResponseZonesTrueClientIPHeader2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesTrueClientIPHeader2Editable bool

const (
	SettingEditResponseZonesTrueClientIPHeader2EditableTrue  SettingEditResponseZonesTrueClientIPHeader2Editable = true
	SettingEditResponseZonesTrueClientIPHeader2EditableFalse SettingEditResponseZonesTrueClientIPHeader2Editable = false
)

func (r SettingEditResponseZonesTrueClientIPHeader2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesTrueClientIPHeader2EditableTrue, SettingEditResponseZonesTrueClientIPHeader2EditableFalse:
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
type SettingEditResponseZonesWAF2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesWAF2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesWAF2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesWAF2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingEditResponseZonesWaf2JSON `json:"-"`
}

// settingEditResponseZonesWaf2JSON contains the JSON metadata for the struct
// [SettingEditResponseZonesWAF2]
type settingEditResponseZonesWaf2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesWAF2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesWaf2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesWAF2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesWAF2ID string

const (
	SettingEditResponseZonesWAF2IDWAF SettingEditResponseZonesWAF2ID = "waf"
)

func (r SettingEditResponseZonesWAF2ID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesWAF2IDWAF:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditResponseZonesWAF2Value string

const (
	SettingEditResponseZonesWAF2ValueOn  SettingEditResponseZonesWAF2Value = "on"
	SettingEditResponseZonesWAF2ValueOff SettingEditResponseZonesWAF2Value = "off"
)

func (r SettingEditResponseZonesWAF2Value) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesWAF2ValueOn, SettingEditResponseZonesWAF2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesWAF2Editable bool

const (
	SettingEditResponseZonesWAF2EditableTrue  SettingEditResponseZonesWAF2Editable = true
	SettingEditResponseZonesWAF2EditableFalse SettingEditResponseZonesWAF2Editable = false
)

func (r SettingEditResponseZonesWAF2Editable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesWAF2EditableTrue, SettingEditResponseZonesWAF2EditableFalse:
		return true
	}
	return false
}

// ID of the zone setting.
type SettingEditResponseID string

const (
	SettingEditResponseID0rtt                          SettingEditResponseID = "0rtt"
	SettingEditResponseIDAdvancedDDoS                  SettingEditResponseID = "advanced_ddos"
	SettingEditResponseIDAegis                         SettingEditResponseID = "aegis"
	SettingEditResponseIDAlwaysOnline                  SettingEditResponseID = "always_online"
	SettingEditResponseIDAlwaysUseHTTPS                SettingEditResponseID = "always_use_https"
	SettingEditResponseIDAutomaticHTTPSRewrites        SettingEditResponseID = "automatic_https_rewrites"
	SettingEditResponseIDBrotli                        SettingEditResponseID = "brotli"
	SettingEditResponseIDBrowserCacheTTL               SettingEditResponseID = "browser_cache_ttl"
	SettingEditResponseIDBrowserCheck                  SettingEditResponseID = "browser_check"
	SettingEditResponseIDCacheLevel                    SettingEditResponseID = "cache_level"
	SettingEditResponseIDChallengeTTL                  SettingEditResponseID = "challenge_ttl"
	SettingEditResponseIDChinaNetworkEnabled           SettingEditResponseID = "china_network_enabled"
	SettingEditResponseIDContentConverter              SettingEditResponseID = "content_converter"
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
	SettingEditResponseIDOriginH2MaxStreams            SettingEditResponseID = "origin_h2_max_streams"
	SettingEditResponseIDOriginMaxHTTPVersion          SettingEditResponseID = "origin_max_http_version"
	SettingEditResponseIDPolish                        SettingEditResponseID = "polish"
	SettingEditResponseIDPrefetchPreload               SettingEditResponseID = "prefetch_preload"
	SettingEditResponseIDPrivacyPass                   SettingEditResponseID = "privacy_pass"
	SettingEditResponseIDProxyReadTimeout              SettingEditResponseID = "proxy_read_timeout"
	SettingEditResponseIDPseudoIPV4                    SettingEditResponseID = "pseudo_ipv4"
	SettingEditResponseIDRedirectsForAITraining        SettingEditResponseID = "redirects_for_ai_training"
	SettingEditResponseIDReplaceInsecureJS             SettingEditResponseID = "replace_insecure_js"
	SettingEditResponseIDResponseBuffering             SettingEditResponseID = "response_buffering"
	SettingEditResponseIDRocketLoader                  SettingEditResponseID = "rocket_loader"
	SettingEditResponseIDAutomaticPlatformOptimization SettingEditResponseID = "automatic_platform_optimization"
	SettingEditResponseIDSearchForAgents               SettingEditResponseID = "search_for_agents"
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
	SettingEditResponseIDTransformations               SettingEditResponseID = "transformations"
	SettingEditResponseIDTransformationsAllowedOrigins SettingEditResponseID = "transformations_allowed_origins"
	SettingEditResponseIDTrueClientIPHeader            SettingEditResponseID = "true_client_ip_header"
	SettingEditResponseIDWAF                           SettingEditResponseID = "waf"
	SettingEditResponseIDWebP                          SettingEditResponseID = "webp"
	SettingEditResponseIDWebsockets                    SettingEditResponseID = "websockets"
)

func (r SettingEditResponseID) IsKnown() bool {
	switch r {
	case SettingEditResponseID0rtt, SettingEditResponseIDAdvancedDDoS, SettingEditResponseIDAegis, SettingEditResponseIDAlwaysOnline, SettingEditResponseIDAlwaysUseHTTPS, SettingEditResponseIDAutomaticHTTPSRewrites, SettingEditResponseIDBrotli, SettingEditResponseIDBrowserCacheTTL, SettingEditResponseIDBrowserCheck, SettingEditResponseIDCacheLevel, SettingEditResponseIDChallengeTTL, SettingEditResponseIDChinaNetworkEnabled, SettingEditResponseIDContentConverter, SettingEditResponseIDCiphers, SettingEditResponseIDCNAMEFlattening, SettingEditResponseIDDevelopmentMode, SettingEditResponseIDEarlyHints, SettingEditResponseIDEdgeCacheTTL, SettingEditResponseIDEmailObfuscation, SettingEditResponseIDH2Prioritization, SettingEditResponseIDHotlinkProtection, SettingEditResponseIDHTTP2, SettingEditResponseIDHTTP3, SettingEditResponseIDImageResizing, SettingEditResponseIDIPGeolocation, SettingEditResponseIDIPV6, SettingEditResponseIDMaxUpload, SettingEditResponseIDMinTLSVersion, SettingEditResponseIDMirage, SettingEditResponseIDNEL, SettingEditResponseIDOpportunisticEncryption, SettingEditResponseIDOpportunisticOnion, SettingEditResponseIDOrangeToOrange, SettingEditResponseIDOriginErrorPagePassThru, SettingEditResponseIDOriginH2MaxStreams, SettingEditResponseIDOriginMaxHTTPVersion, SettingEditResponseIDPolish, SettingEditResponseIDPrefetchPreload, SettingEditResponseIDPrivacyPass, SettingEditResponseIDProxyReadTimeout, SettingEditResponseIDPseudoIPV4, SettingEditResponseIDRedirectsForAITraining, SettingEditResponseIDReplaceInsecureJS, SettingEditResponseIDResponseBuffering, SettingEditResponseIDRocketLoader, SettingEditResponseIDAutomaticPlatformOptimization, SettingEditResponseIDSearchForAgents, SettingEditResponseIDSecurityHeader, SettingEditResponseIDSecurityLevel, SettingEditResponseIDServerSideExclude, SettingEditResponseIDSha1Support, SettingEditResponseIDSortQueryStringForCache, SettingEditResponseIDSSL, SettingEditResponseIDSSLRecommender, SettingEditResponseIDTLS1_2Only, SettingEditResponseIDTLS1_3, SettingEditResponseIDTLSClientAuth, SettingEditResponseIDTransformations, SettingEditResponseIDTransformationsAllowedOrigins, SettingEditResponseIDTrueClientIPHeader, SettingEditResponseIDWAF, SettingEditResponseIDWebP, SettingEditResponseIDWebsockets:
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
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// This field can have the runtime type of [ZeroRTTValue], [AdvancedDDoSValue],
	// [SettingGetResponseZonesCacheRulesAegisValue], [AlwaysOnlineValue],
	// [SettingGetResponseZonesAlwaysUseHTTPS2Value],
	// [SettingGetResponseZonesAutomaticHTTPSRewrites2Value], [BrotliValue], [int64],
	// [SettingGetResponseZonesBrowserCheck2Value],
	// [SettingGetResponseZonesCacheLevel2Value], [ChallengeTTLValue],
	// [SettingGetResponseZonesChinaNetworkEnabledValue],
	// [SettingGetResponseZonesContentConverterValue], [[]string],
	// [SettingGetResponseZonesCNAMEFlatteningValue], [DevelopmentModeValue],
	// [EarlyHintsValue], [SettingGetResponseZonesEdgeCacheTTL2Value],
	// [SettingGetResponseZonesEmailObfuscation2Value], [H2PrioritizationValue],
	// [HotlinkProtectionValue], [HTTP2Value], [HTTP3Value], [ImageResizingValue],
	// [SettingGetResponseZonesIPGeolocation2Value], [IPV6Value],
	// [SettingGetResponseZonesMaxUploadValue], [MinTLSVersionValue],
	// [SettingGetResponseZonesMirage2Value], [NELValue],
	// [SettingGetResponseZonesOpportunisticEncryption2Value],
	// [OpportunisticOnionValue], [OrangeToOrangeValue],
	// [SettingGetResponseZonesOriginErrorPagePassThru2Value],
	// [SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionValue],
	// [SettingGetResponseZonesPolish2Value], [PrefetchPreloadValue],
	// [SettingGetResponseZonesPrivacyPassValue], [float64], [PseudoIPV4Value],
	// [SettingGetResponseZonesRedirectsForAITrainingValue],
	// [SettingGetResponseZonesReplaceInsecureJSValue],
	// [SettingGetResponseZonesResponseBuffering2Value],
	// [SettingGetResponseZonesRocketLoader2Value], [AutomaticPlatformOptimization],
	// [SettingGetResponseZonesSearchForAgentsValue], [SecurityHeadersValue],
	// [SettingGetResponseZonesSecurityLevel2Value], [ServerSideExcludesValue],
	// [SettingGetResponseZonesSha1SupportValue],
	// [SettingGetResponseZonesSortQueryStringForCache2Value],
	// [SettingGetResponseZonesSSL2Value], [SettingGetResponseZonesTLS1_2OnlyValue],
	// [TLS1_3Value], [TLSClientAuthValue],
	// [SettingGetResponseZonesTransformationsValue], [string],
	// [SettingGetResponseZonesTrueClientIPHeader2Value],
	// [SettingGetResponseZonesWAF2Value], [WebPValue], [WebsocketValue].
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
// Possible runtime types of the union are [ZeroRTT], [AdvancedDDoS],
// [SettingGetResponseZonesCacheRulesAegis], [AlwaysOnline],
// [SettingGetResponseZonesAlwaysUseHTTPS2],
// [SettingGetResponseZonesAutomaticHTTPSRewrites2], [Brotli],
// [SettingGetResponseZonesBrowserCacheTTL2],
// [SettingGetResponseZonesBrowserCheck2], [SettingGetResponseZonesCacheLevel2],
// [ChallengeTTL], [SettingGetResponseZonesChinaNetworkEnabled],
// [SettingGetResponseZonesContentConverter], [Ciphers],
// [SettingGetResponseZonesCNAMEFlattening], [DevelopmentMode], [EarlyHints],
// [SettingGetResponseZonesEdgeCacheTTL2],
// [SettingGetResponseZonesEmailObfuscation2], [H2Prioritization],
// [HotlinkProtection], [HTTP2], [HTTP3], [ImageResizing],
// [SettingGetResponseZonesIPGeolocation2], [IPV6],
// [SettingGetResponseZonesMaxUpload], [MinTLSVersion],
// [SettingGetResponseZonesMirage2], [NEL],
// [SettingGetResponseZonesOpportunisticEncryption2], [OpportunisticOnion],
// [OrangeToOrange], [SettingGetResponseZonesOriginErrorPagePassThru2],
// [SettingGetResponseZonesCacheRulesOriginH2MaxStreams],
// [SettingGetResponseZonesCacheRulesOriginMaxHTTPVersion],
// [SettingGetResponseZonesPolish2], [PrefetchPreload],
// [SettingGetResponseZonesPrivacyPass], [ProxyReadTimeout], [PseudoIPV4],
// [SettingGetResponseZonesRedirectsForAITraining],
// [SettingGetResponseZonesReplaceInsecureJS],
// [SettingGetResponseZonesResponseBuffering2],
// [SettingGetResponseZonesRocketLoader2],
// [SettingGetResponseZonesSchemasAutomaticPlatformOptimization],
// [SettingGetResponseZonesSearchForAgents], [SecurityHeaders],
// [SettingGetResponseZonesSecurityLevel2], [ServerSideExcludes],
// [SettingGetResponseZonesSha1Support],
// [SettingGetResponseZonesSortQueryStringForCache2],
// [SettingGetResponseZonesSSL2], [SSLRecommender],
// [SettingGetResponseZonesTLS1_2Only], [TLS1_3], [TLSClientAuth],
// [SettingGetResponseZonesTransformations],
// [SettingGetResponseZonesTransformationsAllowedOrigins],
// [SettingGetResponseZonesTrueClientIPHeader2], [SettingGetResponseZonesWAF2],
// [WebP], [Websocket].
func (r SettingGetResponse) AsUnion() SettingGetResponseUnion {
	return r.union
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [ZeroRTT], [AdvancedDDoS],
// [SettingGetResponseZonesCacheRulesAegis], [AlwaysOnline],
// [SettingGetResponseZonesAlwaysUseHTTPS2],
// [SettingGetResponseZonesAutomaticHTTPSRewrites2], [Brotli],
// [SettingGetResponseZonesBrowserCacheTTL2],
// [SettingGetResponseZonesBrowserCheck2], [SettingGetResponseZonesCacheLevel2],
// [ChallengeTTL], [SettingGetResponseZonesChinaNetworkEnabled],
// [SettingGetResponseZonesContentConverter], [Ciphers],
// [SettingGetResponseZonesCNAMEFlattening], [DevelopmentMode], [EarlyHints],
// [SettingGetResponseZonesEdgeCacheTTL2],
// [SettingGetResponseZonesEmailObfuscation2], [H2Prioritization],
// [HotlinkProtection], [HTTP2], [HTTP3], [ImageResizing],
// [SettingGetResponseZonesIPGeolocation2], [IPV6],
// [SettingGetResponseZonesMaxUpload], [MinTLSVersion],
// [SettingGetResponseZonesMirage2], [NEL],
// [SettingGetResponseZonesOpportunisticEncryption2], [OpportunisticOnion],
// [OrangeToOrange], [SettingGetResponseZonesOriginErrorPagePassThru2],
// [SettingGetResponseZonesCacheRulesOriginH2MaxStreams],
// [SettingGetResponseZonesCacheRulesOriginMaxHTTPVersion],
// [SettingGetResponseZonesPolish2], [PrefetchPreload],
// [SettingGetResponseZonesPrivacyPass], [ProxyReadTimeout], [PseudoIPV4],
// [SettingGetResponseZonesRedirectsForAITraining],
// [SettingGetResponseZonesReplaceInsecureJS],
// [SettingGetResponseZonesResponseBuffering2],
// [SettingGetResponseZonesRocketLoader2],
// [SettingGetResponseZonesSchemasAutomaticPlatformOptimization],
// [SettingGetResponseZonesSearchForAgents], [SecurityHeaders],
// [SettingGetResponseZonesSecurityLevel2], [ServerSideExcludes],
// [SettingGetResponseZonesSha1Support],
// [SettingGetResponseZonesSortQueryStringForCache2],
// [SettingGetResponseZonesSSL2], [SSLRecommender],
// [SettingGetResponseZonesTLS1_2Only], [TLS1_3], [TLSClientAuth],
// [SettingGetResponseZonesTransformations],
// [SettingGetResponseZonesTransformationsAllowedOrigins],
// [SettingGetResponseZonesTrueClientIPHeader2], [SettingGetResponseZonesWAF2],
// [WebP] or [Websocket].
type SettingGetResponseUnion interface {
	implementsSettingGetResponse()
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
			Type:       reflect.TypeOf(SettingGetResponseZonesCacheRulesAegis{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesAlwaysUseHTTPS2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesAutomaticHTTPSRewrites2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Brotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesBrowserCacheTTL2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesBrowserCheck2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesCacheLevel2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesChinaNetworkEnabled{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesContentConverter{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesEdgeCacheTTL2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesEmailObfuscation2{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesIPGeolocation2{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesMirage2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesOpportunisticEncryption2{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesOriginErrorPagePassThru2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesCacheRulesOriginH2MaxStreams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesCacheRulesOriginMaxHTTPVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesPolish2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesPrivacyPass{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesRedirectsForAITraining{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesReplaceInsecureJS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesResponseBuffering2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesRocketLoader2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSearchForAgents{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SecurityHeaders{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSecurityLevel2{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesSortQueryStringForCache2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSSL2{}),
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
			Type:       reflect.TypeOf(SettingGetResponseZonesTransformations{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesTransformationsAllowedOrigins{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesTrueClientIPHeader2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesWAF2{}),
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

// Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your
// layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your
// account so that you can increase your origin security by only allowing traffic
// from a small list of IP addresses.
type SettingGetResponseZonesCacheRulesAegis struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesCacheRulesAegisID `json:"id" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Value of the zone setting.
	Value SettingGetResponseZonesCacheRulesAegisValue `json:"value"`
	JSON  settingGetResponseZonesCacheRulesAegisJSON  `json:"-"`
}

// settingGetResponseZonesCacheRulesAegisJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesCacheRulesAegis]
type settingGetResponseZonesCacheRulesAegisJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesCacheRulesAegis) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesCacheRulesAegisJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesCacheRulesAegis) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesCacheRulesAegisID string

const (
	SettingGetResponseZonesCacheRulesAegisIDAegis SettingGetResponseZonesCacheRulesAegisID = "aegis"
)

func (r SettingGetResponseZonesCacheRulesAegisID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesCacheRulesAegisIDAegis:
		return true
	}
	return false
}

// Value of the zone setting.
type SettingGetResponseZonesCacheRulesAegisValue struct {
	// Whether the feature is enabled or not.
	Enabled bool `json:"enabled"`
	// Egress pool id which refers to a grouping of dedicated egress IPs through which
	// Cloudflare will connect to origin.
	PoolID string                                          `json:"pool_id"`
	JSON   settingGetResponseZonesCacheRulesAegisValueJSON `json:"-"`
}

// settingGetResponseZonesCacheRulesAegisValueJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesCacheRulesAegisValue]
type settingGetResponseZonesCacheRulesAegisValueJSON struct {
	Enabled     apijson.Field
	PoolID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesCacheRulesAegisValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesCacheRulesAegisValueJSON) RawJSON() string {
	return r.raw
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingGetResponseZonesAlwaysUseHTTPS2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesAlwaysUseHTTPS2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesAlwaysUseHTTPS2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesAlwaysUseHTTPS2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesAlwaysUseHttps2JSON `json:"-"`
}

// settingGetResponseZonesAlwaysUseHttps2JSON contains the JSON metadata for the
// struct [SettingGetResponseZonesAlwaysUseHTTPS2]
type settingGetResponseZonesAlwaysUseHttps2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesAlwaysUseHTTPS2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesAlwaysUseHttps2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesAlwaysUseHTTPS2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesAlwaysUseHTTPS2ID string

const (
	SettingGetResponseZonesAlwaysUseHTTPS2IDAlwaysUseHTTPS SettingGetResponseZonesAlwaysUseHTTPS2ID = "always_use_https"
)

func (r SettingGetResponseZonesAlwaysUseHTTPS2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesAlwaysUseHTTPS2IDAlwaysUseHTTPS:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesAlwaysUseHTTPS2Value string

const (
	SettingGetResponseZonesAlwaysUseHTTPS2ValueOn  SettingGetResponseZonesAlwaysUseHTTPS2Value = "on"
	SettingGetResponseZonesAlwaysUseHTTPS2ValueOff SettingGetResponseZonesAlwaysUseHTTPS2Value = "off"
)

func (r SettingGetResponseZonesAlwaysUseHTTPS2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesAlwaysUseHTTPS2ValueOn, SettingGetResponseZonesAlwaysUseHTTPS2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesAlwaysUseHTTPS2Editable bool

const (
	SettingGetResponseZonesAlwaysUseHTTPS2EditableTrue  SettingGetResponseZonesAlwaysUseHTTPS2Editable = true
	SettingGetResponseZonesAlwaysUseHTTPS2EditableFalse SettingGetResponseZonesAlwaysUseHTTPS2Editable = false
)

func (r SettingGetResponseZonesAlwaysUseHTTPS2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesAlwaysUseHTTPS2EditableTrue, SettingGetResponseZonesAlwaysUseHTTPS2EditableFalse:
		return true
	}
	return false
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingGetResponseZonesAutomaticHTTPSRewrites2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesAutomaticHTTPSRewrites2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesAutomaticHTTPSRewrites2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesAutomaticHTTPSRewrites2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesAutomaticHTTPSRewrites2JSON `json:"-"`
}

// settingGetResponseZonesAutomaticHTTPSRewrites2JSON contains the JSON metadata
// for the struct [SettingGetResponseZonesAutomaticHTTPSRewrites2]
type settingGetResponseZonesAutomaticHTTPSRewrites2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesAutomaticHTTPSRewrites2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesAutomaticHTTPSRewrites2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesAutomaticHTTPSRewrites2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesAutomaticHTTPSRewrites2ID string

const (
	SettingGetResponseZonesAutomaticHTTPSRewrites2IDAutomaticHTTPSRewrites SettingGetResponseZonesAutomaticHTTPSRewrites2ID = "automatic_https_rewrites"
)

func (r SettingGetResponseZonesAutomaticHTTPSRewrites2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesAutomaticHTTPSRewrites2IDAutomaticHTTPSRewrites:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesAutomaticHTTPSRewrites2Value string

const (
	SettingGetResponseZonesAutomaticHTTPSRewrites2ValueOn  SettingGetResponseZonesAutomaticHTTPSRewrites2Value = "on"
	SettingGetResponseZonesAutomaticHTTPSRewrites2ValueOff SettingGetResponseZonesAutomaticHTTPSRewrites2Value = "off"
)

func (r SettingGetResponseZonesAutomaticHTTPSRewrites2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesAutomaticHTTPSRewrites2ValueOn, SettingGetResponseZonesAutomaticHTTPSRewrites2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesAutomaticHTTPSRewrites2Editable bool

const (
	SettingGetResponseZonesAutomaticHTTPSRewrites2EditableTrue  SettingGetResponseZonesAutomaticHTTPSRewrites2Editable = true
	SettingGetResponseZonesAutomaticHTTPSRewrites2EditableFalse SettingGetResponseZonesAutomaticHTTPSRewrites2Editable = false
)

func (r SettingGetResponseZonesAutomaticHTTPSRewrites2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesAutomaticHTTPSRewrites2EditableTrue, SettingGetResponseZonesAutomaticHTTPSRewrites2EditableFalse:
		return true
	}
	return false
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingGetResponseZonesBrowserCacheTTL2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesBrowserCacheTTL2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value int64 `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesBrowserCacheTTL2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesBrowserCacheTtl2JSON `json:"-"`
}

// settingGetResponseZonesBrowserCacheTtl2JSON contains the JSON metadata for the
// struct [SettingGetResponseZonesBrowserCacheTTL2]
type settingGetResponseZonesBrowserCacheTtl2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesBrowserCacheTTL2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesBrowserCacheTtl2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesBrowserCacheTTL2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesBrowserCacheTTL2ID string

const (
	SettingGetResponseZonesBrowserCacheTTL2IDBrowserCacheTTL SettingGetResponseZonesBrowserCacheTTL2ID = "browser_cache_ttl"
)

func (r SettingGetResponseZonesBrowserCacheTTL2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesBrowserCacheTTL2IDBrowserCacheTTL:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesBrowserCacheTTL2Editable bool

const (
	SettingGetResponseZonesBrowserCacheTTL2EditableTrue  SettingGetResponseZonesBrowserCacheTTL2Editable = true
	SettingGetResponseZonesBrowserCacheTTL2EditableFalse SettingGetResponseZonesBrowserCacheTTL2Editable = false
)

func (r SettingGetResponseZonesBrowserCacheTTL2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesBrowserCacheTTL2EditableTrue, SettingGetResponseZonesBrowserCacheTTL2EditableFalse:
		return true
	}
	return false
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type SettingGetResponseZonesBrowserCheck2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesBrowserCheck2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesBrowserCheck2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesBrowserCheck2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesBrowserCheck2JSON `json:"-"`
}

// settingGetResponseZonesBrowserCheck2JSON contains the JSON metadata for the
// struct [SettingGetResponseZonesBrowserCheck2]
type settingGetResponseZonesBrowserCheck2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesBrowserCheck2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesBrowserCheck2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesBrowserCheck2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesBrowserCheck2ID string

const (
	SettingGetResponseZonesBrowserCheck2IDBrowserCheck SettingGetResponseZonesBrowserCheck2ID = "browser_check"
)

func (r SettingGetResponseZonesBrowserCheck2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesBrowserCheck2IDBrowserCheck:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesBrowserCheck2Value string

const (
	SettingGetResponseZonesBrowserCheck2ValueOn  SettingGetResponseZonesBrowserCheck2Value = "on"
	SettingGetResponseZonesBrowserCheck2ValueOff SettingGetResponseZonesBrowserCheck2Value = "off"
)

func (r SettingGetResponseZonesBrowserCheck2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesBrowserCheck2ValueOn, SettingGetResponseZonesBrowserCheck2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesBrowserCheck2Editable bool

const (
	SettingGetResponseZonesBrowserCheck2EditableTrue  SettingGetResponseZonesBrowserCheck2Editable = true
	SettingGetResponseZonesBrowserCheck2EditableFalse SettingGetResponseZonesBrowserCheck2Editable = false
)

func (r SettingGetResponseZonesBrowserCheck2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesBrowserCheck2EditableTrue, SettingGetResponseZonesBrowserCheck2EditableFalse:
		return true
	}
	return false
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type SettingGetResponseZonesCacheLevel2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesCacheLevel2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesCacheLevel2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesCacheLevel2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesCacheLevel2JSON `json:"-"`
}

// settingGetResponseZonesCacheLevel2JSON contains the JSON metadata for the struct
// [SettingGetResponseZonesCacheLevel2]
type settingGetResponseZonesCacheLevel2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesCacheLevel2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesCacheLevel2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesCacheLevel2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesCacheLevel2ID string

const (
	SettingGetResponseZonesCacheLevel2IDCacheLevel SettingGetResponseZonesCacheLevel2ID = "cache_level"
)

func (r SettingGetResponseZonesCacheLevel2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesCacheLevel2IDCacheLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesCacheLevel2Value string

const (
	SettingGetResponseZonesCacheLevel2ValueAggressive SettingGetResponseZonesCacheLevel2Value = "aggressive"
	SettingGetResponseZonesCacheLevel2ValueBasic      SettingGetResponseZonesCacheLevel2Value = "basic"
	SettingGetResponseZonesCacheLevel2ValueSimplified SettingGetResponseZonesCacheLevel2Value = "simplified"
)

func (r SettingGetResponseZonesCacheLevel2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesCacheLevel2ValueAggressive, SettingGetResponseZonesCacheLevel2ValueBasic, SettingGetResponseZonesCacheLevel2ValueSimplified:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesCacheLevel2Editable bool

const (
	SettingGetResponseZonesCacheLevel2EditableTrue  SettingGetResponseZonesCacheLevel2Editable = true
	SettingGetResponseZonesCacheLevel2EditableFalse SettingGetResponseZonesCacheLevel2Editable = false
)

func (r SettingGetResponseZonesCacheLevel2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesCacheLevel2EditableTrue, SettingGetResponseZonesCacheLevel2EditableFalse:
		return true
	}
	return false
}

// Determines whether or not the china network is enabled.
type SettingGetResponseZonesChinaNetworkEnabled struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesChinaNetworkEnabledID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesChinaNetworkEnabledValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesChinaNetworkEnabledEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesChinaNetworkEnabledJSON `json:"-"`
}

// settingGetResponseZonesChinaNetworkEnabledJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesChinaNetworkEnabled]
type settingGetResponseZonesChinaNetworkEnabledJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesChinaNetworkEnabled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesChinaNetworkEnabledJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesChinaNetworkEnabled) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesChinaNetworkEnabledID string

const (
	SettingGetResponseZonesChinaNetworkEnabledIDChinaNetworkEnabled SettingGetResponseZonesChinaNetworkEnabledID = "china_network_enabled"
)

func (r SettingGetResponseZonesChinaNetworkEnabledID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesChinaNetworkEnabledIDChinaNetworkEnabled:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesChinaNetworkEnabledValue string

const (
	SettingGetResponseZonesChinaNetworkEnabledValueOn  SettingGetResponseZonesChinaNetworkEnabledValue = "on"
	SettingGetResponseZonesChinaNetworkEnabledValueOff SettingGetResponseZonesChinaNetworkEnabledValue = "off"
)

func (r SettingGetResponseZonesChinaNetworkEnabledValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesChinaNetworkEnabledValueOn, SettingGetResponseZonesChinaNetworkEnabledValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesChinaNetworkEnabledEditable bool

const (
	SettingGetResponseZonesChinaNetworkEnabledEditableTrue  SettingGetResponseZonesChinaNetworkEnabledEditable = true
	SettingGetResponseZonesChinaNetworkEnabledEditableFalse SettingGetResponseZonesChinaNetworkEnabledEditable = false
)

func (r SettingGetResponseZonesChinaNetworkEnabledEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesChinaNetworkEnabledEditableTrue, SettingGetResponseZonesChinaNetworkEnabledEditableFalse:
		return true
	}
	return false
}

// When enabled and the client sends an Accept header requesting text/markdown,
// Cloudflare will convert HTML responses to Markdown format using the toMarkdown()
// service. Refer to the
// [developer documentation](https://developers.cloudflare.com/workers-ai/features/markdown-conversion/)
// for more information.
type SettingGetResponseZonesContentConverter struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesContentConverterID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesContentConverterValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesContentConverterEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesContentConverterJSON `json:"-"`
}

// settingGetResponseZonesContentConverterJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesContentConverter]
type settingGetResponseZonesContentConverterJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesContentConverter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesContentConverterJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesContentConverter) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesContentConverterID string

const (
	SettingGetResponseZonesContentConverterIDContentConverter SettingGetResponseZonesContentConverterID = "content_converter"
)

func (r SettingGetResponseZonesContentConverterID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesContentConverterIDContentConverter:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesContentConverterValue string

const (
	SettingGetResponseZonesContentConverterValueOff SettingGetResponseZonesContentConverterValue = "off"
	SettingGetResponseZonesContentConverterValueOn  SettingGetResponseZonesContentConverterValue = "on"
)

func (r SettingGetResponseZonesContentConverterValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesContentConverterValueOff, SettingGetResponseZonesContentConverterValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesContentConverterEditable bool

const (
	SettingGetResponseZonesContentConverterEditableTrue  SettingGetResponseZonesContentConverterEditable = true
	SettingGetResponseZonesContentConverterEditableFalse SettingGetResponseZonesContentConverterEditable = false
)

func (r SettingGetResponseZonesContentConverterEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesContentConverterEditableTrue, SettingGetResponseZonesContentConverterEditableFalse:
		return true
	}
	return false
}

// Whether or not cname flattening is on.
//
// Deprecated: This zone setting is deprecated; please use the DNS Settings route
// instead. More information at
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2025-03-21
type SettingGetResponseZonesCNAMEFlattening struct {
	// How to flatten the cname destination.
	ID SettingGetResponseZonesCNAMEFlatteningID `json:"id" api:"required"`
	// Current value of the zone setting.
	//
	// Deprecated: This zone setting is deprecated; please use the DNS Settings route
	// instead. More information at
	// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2025-03-21
	Value SettingGetResponseZonesCNAMEFlatteningValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesCNAMEFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingGetResponseZonesCNAMEFlattening) implementsSettingGetResponse() {}

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
type SettingGetResponseZonesEdgeCacheTTL2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesEdgeCacheTTL2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesEdgeCacheTTL2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesEdgeCacheTTL2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesEdgeCacheTtl2JSON `json:"-"`
}

// settingGetResponseZonesEdgeCacheTtl2JSON contains the JSON metadata for the
// struct [SettingGetResponseZonesEdgeCacheTTL2]
type settingGetResponseZonesEdgeCacheTtl2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesEdgeCacheTTL2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesEdgeCacheTtl2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesEdgeCacheTTL2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesEdgeCacheTTL2ID string

const (
	SettingGetResponseZonesEdgeCacheTTL2IDEdgeCacheTTL SettingGetResponseZonesEdgeCacheTTL2ID = "edge_cache_ttl"
)

func (r SettingGetResponseZonesEdgeCacheTTL2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEdgeCacheTTL2IDEdgeCacheTTL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesEdgeCacheTTL2Value float64

const (
	SettingGetResponseZonesEdgeCacheTTL2Value30     SettingGetResponseZonesEdgeCacheTTL2Value = 30
	SettingGetResponseZonesEdgeCacheTTL2Value60     SettingGetResponseZonesEdgeCacheTTL2Value = 60
	SettingGetResponseZonesEdgeCacheTTL2Value300    SettingGetResponseZonesEdgeCacheTTL2Value = 300
	SettingGetResponseZonesEdgeCacheTTL2Value1200   SettingGetResponseZonesEdgeCacheTTL2Value = 1200
	SettingGetResponseZonesEdgeCacheTTL2Value1800   SettingGetResponseZonesEdgeCacheTTL2Value = 1800
	SettingGetResponseZonesEdgeCacheTTL2Value3600   SettingGetResponseZonesEdgeCacheTTL2Value = 3600
	SettingGetResponseZonesEdgeCacheTTL2Value7200   SettingGetResponseZonesEdgeCacheTTL2Value = 7200
	SettingGetResponseZonesEdgeCacheTTL2Value10800  SettingGetResponseZonesEdgeCacheTTL2Value = 10800
	SettingGetResponseZonesEdgeCacheTTL2Value14400  SettingGetResponseZonesEdgeCacheTTL2Value = 14400
	SettingGetResponseZonesEdgeCacheTTL2Value18000  SettingGetResponseZonesEdgeCacheTTL2Value = 18000
	SettingGetResponseZonesEdgeCacheTTL2Value28800  SettingGetResponseZonesEdgeCacheTTL2Value = 28800
	SettingGetResponseZonesEdgeCacheTTL2Value43200  SettingGetResponseZonesEdgeCacheTTL2Value = 43200
	SettingGetResponseZonesEdgeCacheTTL2Value57600  SettingGetResponseZonesEdgeCacheTTL2Value = 57600
	SettingGetResponseZonesEdgeCacheTTL2Value72000  SettingGetResponseZonesEdgeCacheTTL2Value = 72000
	SettingGetResponseZonesEdgeCacheTTL2Value86400  SettingGetResponseZonesEdgeCacheTTL2Value = 86400
	SettingGetResponseZonesEdgeCacheTTL2Value172800 SettingGetResponseZonesEdgeCacheTTL2Value = 172800
	SettingGetResponseZonesEdgeCacheTTL2Value259200 SettingGetResponseZonesEdgeCacheTTL2Value = 259200
	SettingGetResponseZonesEdgeCacheTTL2Value345600 SettingGetResponseZonesEdgeCacheTTL2Value = 345600
	SettingGetResponseZonesEdgeCacheTTL2Value432000 SettingGetResponseZonesEdgeCacheTTL2Value = 432000
	SettingGetResponseZonesEdgeCacheTTL2Value518400 SettingGetResponseZonesEdgeCacheTTL2Value = 518400
	SettingGetResponseZonesEdgeCacheTTL2Value604800 SettingGetResponseZonesEdgeCacheTTL2Value = 604800
)

func (r SettingGetResponseZonesEdgeCacheTTL2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEdgeCacheTTL2Value30, SettingGetResponseZonesEdgeCacheTTL2Value60, SettingGetResponseZonesEdgeCacheTTL2Value300, SettingGetResponseZonesEdgeCacheTTL2Value1200, SettingGetResponseZonesEdgeCacheTTL2Value1800, SettingGetResponseZonesEdgeCacheTTL2Value3600, SettingGetResponseZonesEdgeCacheTTL2Value7200, SettingGetResponseZonesEdgeCacheTTL2Value10800, SettingGetResponseZonesEdgeCacheTTL2Value14400, SettingGetResponseZonesEdgeCacheTTL2Value18000, SettingGetResponseZonesEdgeCacheTTL2Value28800, SettingGetResponseZonesEdgeCacheTTL2Value43200, SettingGetResponseZonesEdgeCacheTTL2Value57600, SettingGetResponseZonesEdgeCacheTTL2Value72000, SettingGetResponseZonesEdgeCacheTTL2Value86400, SettingGetResponseZonesEdgeCacheTTL2Value172800, SettingGetResponseZonesEdgeCacheTTL2Value259200, SettingGetResponseZonesEdgeCacheTTL2Value345600, SettingGetResponseZonesEdgeCacheTTL2Value432000, SettingGetResponseZonesEdgeCacheTTL2Value518400, SettingGetResponseZonesEdgeCacheTTL2Value604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesEdgeCacheTTL2Editable bool

const (
	SettingGetResponseZonesEdgeCacheTTL2EditableTrue  SettingGetResponseZonesEdgeCacheTTL2Editable = true
	SettingGetResponseZonesEdgeCacheTTL2EditableFalse SettingGetResponseZonesEdgeCacheTTL2Editable = false
)

func (r SettingGetResponseZonesEdgeCacheTTL2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEdgeCacheTTL2EditableTrue, SettingGetResponseZonesEdgeCacheTTL2EditableFalse:
		return true
	}
	return false
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingGetResponseZonesEmailObfuscation2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesEmailObfuscation2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesEmailObfuscation2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesEmailObfuscation2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesEmailObfuscation2JSON `json:"-"`
}

// settingGetResponseZonesEmailObfuscation2JSON contains the JSON metadata for the
// struct [SettingGetResponseZonesEmailObfuscation2]
type settingGetResponseZonesEmailObfuscation2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesEmailObfuscation2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesEmailObfuscation2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesEmailObfuscation2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesEmailObfuscation2ID string

const (
	SettingGetResponseZonesEmailObfuscation2IDEmailObfuscation SettingGetResponseZonesEmailObfuscation2ID = "email_obfuscation"
)

func (r SettingGetResponseZonesEmailObfuscation2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEmailObfuscation2IDEmailObfuscation:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesEmailObfuscation2Value string

const (
	SettingGetResponseZonesEmailObfuscation2ValueOn  SettingGetResponseZonesEmailObfuscation2Value = "on"
	SettingGetResponseZonesEmailObfuscation2ValueOff SettingGetResponseZonesEmailObfuscation2Value = "off"
)

func (r SettingGetResponseZonesEmailObfuscation2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEmailObfuscation2ValueOn, SettingGetResponseZonesEmailObfuscation2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesEmailObfuscation2Editable bool

const (
	SettingGetResponseZonesEmailObfuscation2EditableTrue  SettingGetResponseZonesEmailObfuscation2Editable = true
	SettingGetResponseZonesEmailObfuscation2EditableFalse SettingGetResponseZonesEmailObfuscation2Editable = false
)

func (r SettingGetResponseZonesEmailObfuscation2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEmailObfuscation2EditableTrue, SettingGetResponseZonesEmailObfuscation2EditableFalse:
		return true
	}
	return false
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingGetResponseZonesIPGeolocation2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesIPGeolocation2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesIPGeolocation2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesIPGeolocation2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesIPGeolocation2JSON `json:"-"`
}

// settingGetResponseZonesIPGeolocation2JSON contains the JSON metadata for the
// struct [SettingGetResponseZonesIPGeolocation2]
type settingGetResponseZonesIPGeolocation2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesIPGeolocation2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesIPGeolocation2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesIPGeolocation2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesIPGeolocation2ID string

const (
	SettingGetResponseZonesIPGeolocation2IDIPGeolocation SettingGetResponseZonesIPGeolocation2ID = "ip_geolocation"
)

func (r SettingGetResponseZonesIPGeolocation2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesIPGeolocation2IDIPGeolocation:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesIPGeolocation2Value string

const (
	SettingGetResponseZonesIPGeolocation2ValueOn  SettingGetResponseZonesIPGeolocation2Value = "on"
	SettingGetResponseZonesIPGeolocation2ValueOff SettingGetResponseZonesIPGeolocation2Value = "off"
)

func (r SettingGetResponseZonesIPGeolocation2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesIPGeolocation2ValueOn, SettingGetResponseZonesIPGeolocation2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesIPGeolocation2Editable bool

const (
	SettingGetResponseZonesIPGeolocation2EditableTrue  SettingGetResponseZonesIPGeolocation2Editable = true
	SettingGetResponseZonesIPGeolocation2EditableFalse SettingGetResponseZonesIPGeolocation2Editable = false
)

func (r SettingGetResponseZonesIPGeolocation2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesIPGeolocation2EditableTrue, SettingGetResponseZonesIPGeolocation2EditableFalse:
		return true
	}
	return false
}

// Maximum size of an allowable upload.
type SettingGetResponseZonesMaxUpload struct {
	// identifier of the zone setting.
	ID SettingGetResponseZonesMaxUploadID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesMaxUploadValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingGetResponseZonesMaxUpload) implementsSettingGetResponse() {}

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
type SettingGetResponseZonesMaxUploadValue int64

const (
	SettingGetResponseZonesMaxUploadValue100  SettingGetResponseZonesMaxUploadValue = 100
	SettingGetResponseZonesMaxUploadValue125  SettingGetResponseZonesMaxUploadValue = 125
	SettingGetResponseZonesMaxUploadValue150  SettingGetResponseZonesMaxUploadValue = 150
	SettingGetResponseZonesMaxUploadValue175  SettingGetResponseZonesMaxUploadValue = 175
	SettingGetResponseZonesMaxUploadValue200  SettingGetResponseZonesMaxUploadValue = 200
	SettingGetResponseZonesMaxUploadValue225  SettingGetResponseZonesMaxUploadValue = 225
	SettingGetResponseZonesMaxUploadValue250  SettingGetResponseZonesMaxUploadValue = 250
	SettingGetResponseZonesMaxUploadValue275  SettingGetResponseZonesMaxUploadValue = 275
	SettingGetResponseZonesMaxUploadValue300  SettingGetResponseZonesMaxUploadValue = 300
	SettingGetResponseZonesMaxUploadValue325  SettingGetResponseZonesMaxUploadValue = 325
	SettingGetResponseZonesMaxUploadValue350  SettingGetResponseZonesMaxUploadValue = 350
	SettingGetResponseZonesMaxUploadValue375  SettingGetResponseZonesMaxUploadValue = 375
	SettingGetResponseZonesMaxUploadValue400  SettingGetResponseZonesMaxUploadValue = 400
	SettingGetResponseZonesMaxUploadValue425  SettingGetResponseZonesMaxUploadValue = 425
	SettingGetResponseZonesMaxUploadValue450  SettingGetResponseZonesMaxUploadValue = 450
	SettingGetResponseZonesMaxUploadValue475  SettingGetResponseZonesMaxUploadValue = 475
	SettingGetResponseZonesMaxUploadValue500  SettingGetResponseZonesMaxUploadValue = 500
	SettingGetResponseZonesMaxUploadValue1000 SettingGetResponseZonesMaxUploadValue = 1000
)

func (r SettingGetResponseZonesMaxUploadValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesMaxUploadValue100, SettingGetResponseZonesMaxUploadValue125, SettingGetResponseZonesMaxUploadValue150, SettingGetResponseZonesMaxUploadValue175, SettingGetResponseZonesMaxUploadValue200, SettingGetResponseZonesMaxUploadValue225, SettingGetResponseZonesMaxUploadValue250, SettingGetResponseZonesMaxUploadValue275, SettingGetResponseZonesMaxUploadValue300, SettingGetResponseZonesMaxUploadValue325, SettingGetResponseZonesMaxUploadValue350, SettingGetResponseZonesMaxUploadValue375, SettingGetResponseZonesMaxUploadValue400, SettingGetResponseZonesMaxUploadValue425, SettingGetResponseZonesMaxUploadValue450, SettingGetResponseZonesMaxUploadValue475, SettingGetResponseZonesMaxUploadValue500, SettingGetResponseZonesMaxUploadValue1000:
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
//
// Deprecated: Mirage is being deprecated. More information at
// https://developers.cloudflare.com/speed/optimization/images/mirage/
type SettingGetResponseZonesMirage2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesMirage2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	//
	// Deprecated: Mirage is being deprecated. More information at
	// https://developers.cloudflare.com/speed/optimization/images/mirage/
	Value SettingGetResponseZonesMirage2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesMirage2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesMirage2JSON `json:"-"`
}

// settingGetResponseZonesMirage2JSON contains the JSON metadata for the struct
// [SettingGetResponseZonesMirage2]
type settingGetResponseZonesMirage2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesMirage2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesMirage2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesMirage2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesMirage2ID string

const (
	SettingGetResponseZonesMirage2IDMirage SettingGetResponseZonesMirage2ID = "mirage"
)

func (r SettingGetResponseZonesMirage2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesMirage2IDMirage:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesMirage2Value string

const (
	SettingGetResponseZonesMirage2ValueOn  SettingGetResponseZonesMirage2Value = "on"
	SettingGetResponseZonesMirage2ValueOff SettingGetResponseZonesMirage2Value = "off"
)

func (r SettingGetResponseZonesMirage2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesMirage2ValueOn, SettingGetResponseZonesMirage2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesMirage2Editable bool

const (
	SettingGetResponseZonesMirage2EditableTrue  SettingGetResponseZonesMirage2Editable = true
	SettingGetResponseZonesMirage2EditableFalse SettingGetResponseZonesMirage2Editable = false
)

func (r SettingGetResponseZonesMirage2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesMirage2EditableTrue, SettingGetResponseZonesMirage2EditableFalse:
		return true
	}
	return false
}

// Enables the Opportunistic Encryption feature for a zone.
type SettingGetResponseZonesOpportunisticEncryption2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesOpportunisticEncryption2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesOpportunisticEncryption2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesOpportunisticEncryption2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesOpportunisticEncryption2JSON `json:"-"`
}

// settingGetResponseZonesOpportunisticEncryption2JSON contains the JSON metadata
// for the struct [SettingGetResponseZonesOpportunisticEncryption2]
type settingGetResponseZonesOpportunisticEncryption2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesOpportunisticEncryption2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesOpportunisticEncryption2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesOpportunisticEncryption2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesOpportunisticEncryption2ID string

const (
	SettingGetResponseZonesOpportunisticEncryption2IDOpportunisticEncryption SettingGetResponseZonesOpportunisticEncryption2ID = "opportunistic_encryption"
)

func (r SettingGetResponseZonesOpportunisticEncryption2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesOpportunisticEncryption2IDOpportunisticEncryption:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesOpportunisticEncryption2Value string

const (
	SettingGetResponseZonesOpportunisticEncryption2ValueOn  SettingGetResponseZonesOpportunisticEncryption2Value = "on"
	SettingGetResponseZonesOpportunisticEncryption2ValueOff SettingGetResponseZonesOpportunisticEncryption2Value = "off"
)

func (r SettingGetResponseZonesOpportunisticEncryption2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesOpportunisticEncryption2ValueOn, SettingGetResponseZonesOpportunisticEncryption2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesOpportunisticEncryption2Editable bool

const (
	SettingGetResponseZonesOpportunisticEncryption2EditableTrue  SettingGetResponseZonesOpportunisticEncryption2Editable = true
	SettingGetResponseZonesOpportunisticEncryption2EditableFalse SettingGetResponseZonesOpportunisticEncryption2Editable = false
)

func (r SettingGetResponseZonesOpportunisticEncryption2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesOpportunisticEncryption2EditableTrue, SettingGetResponseZonesOpportunisticEncryption2EditableFalse:
		return true
	}
	return false
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingGetResponseZonesOriginErrorPagePassThru2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesOriginErrorPagePassThru2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesOriginErrorPagePassThru2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesOriginErrorPagePassThru2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesOriginErrorPagePassThru2JSON `json:"-"`
}

// settingGetResponseZonesOriginErrorPagePassThru2JSON contains the JSON metadata
// for the struct [SettingGetResponseZonesOriginErrorPagePassThru2]
type settingGetResponseZonesOriginErrorPagePassThru2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesOriginErrorPagePassThru2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesOriginErrorPagePassThru2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesOriginErrorPagePassThru2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesOriginErrorPagePassThru2ID string

const (
	SettingGetResponseZonesOriginErrorPagePassThru2IDOriginErrorPagePassThru SettingGetResponseZonesOriginErrorPagePassThru2ID = "origin_error_page_pass_thru"
)

func (r SettingGetResponseZonesOriginErrorPagePassThru2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesOriginErrorPagePassThru2IDOriginErrorPagePassThru:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesOriginErrorPagePassThru2Value string

const (
	SettingGetResponseZonesOriginErrorPagePassThru2ValueOn  SettingGetResponseZonesOriginErrorPagePassThru2Value = "on"
	SettingGetResponseZonesOriginErrorPagePassThru2ValueOff SettingGetResponseZonesOriginErrorPagePassThru2Value = "off"
)

func (r SettingGetResponseZonesOriginErrorPagePassThru2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesOriginErrorPagePassThru2ValueOn, SettingGetResponseZonesOriginErrorPagePassThru2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesOriginErrorPagePassThru2Editable bool

const (
	SettingGetResponseZonesOriginErrorPagePassThru2EditableTrue  SettingGetResponseZonesOriginErrorPagePassThru2Editable = true
	SettingGetResponseZonesOriginErrorPagePassThru2EditableFalse SettingGetResponseZonesOriginErrorPagePassThru2Editable = false
)

func (r SettingGetResponseZonesOriginErrorPagePassThru2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesOriginErrorPagePassThru2EditableTrue, SettingGetResponseZonesOriginErrorPagePassThru2EditableFalse:
		return true
	}
	return false
}

// Origin H2 Max Streams configures the max number of concurrent requests that
// Cloudflare will send within the same connection when communicating with the
// origin server, if the origin supports it. Note that if your origin does not
// support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also
// note that the default value is `100` for all plan types except Enterprise where
// it is `1`. `1` means that H2 multiplexing is disabled.
type SettingGetResponseZonesCacheRulesOriginH2MaxStreams struct {
	// Value of the zone setting.
	ID SettingGetResponseZonesCacheRulesOriginH2MaxStreamsID `json:"id" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Value of the Origin H2 Max Streams Setting.
	Value int64                                                   `json:"value"`
	JSON  settingGetResponseZonesCacheRulesOriginH2MaxStreamsJSON `json:"-"`
}

// settingGetResponseZonesCacheRulesOriginH2MaxStreamsJSON contains the JSON
// metadata for the struct [SettingGetResponseZonesCacheRulesOriginH2MaxStreams]
type settingGetResponseZonesCacheRulesOriginH2MaxStreamsJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesCacheRulesOriginH2MaxStreams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesCacheRulesOriginH2MaxStreamsJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesCacheRulesOriginH2MaxStreams) implementsSettingGetResponse() {}

// Value of the zone setting.
type SettingGetResponseZonesCacheRulesOriginH2MaxStreamsID string

const (
	SettingGetResponseZonesCacheRulesOriginH2MaxStreamsIDOriginH2MaxStreams SettingGetResponseZonesCacheRulesOriginH2MaxStreamsID = "origin_h2_max_streams"
)

func (r SettingGetResponseZonesCacheRulesOriginH2MaxStreamsID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesCacheRulesOriginH2MaxStreamsIDOriginH2MaxStreams:
		return true
	}
	return false
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except
// Enterprise where it is "1".
type SettingGetResponseZonesCacheRulesOriginMaxHTTPVersion struct {
	// Value of the zone setting.
	ID SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionID `json:"id" api:"required"`
	// Last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on" api:"nullable" format:"date-time"`
	// Value of the Origin Max HTTP Version Setting.
	Value SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionValue `json:"value"`
	JSON  settingGetResponseZonesCacheRulesOriginMaxHTTPVersionJSON  `json:"-"`
}

// settingGetResponseZonesCacheRulesOriginMaxHTTPVersionJSON contains the JSON
// metadata for the struct [SettingGetResponseZonesCacheRulesOriginMaxHTTPVersion]
type settingGetResponseZonesCacheRulesOriginMaxHTTPVersionJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesCacheRulesOriginMaxHTTPVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesCacheRulesOriginMaxHTTPVersionJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesCacheRulesOriginMaxHTTPVersion) implementsSettingGetResponse() {}

// Value of the zone setting.
type SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionID string

const (
	SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionIDOriginMaxHTTPVersion SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionID = "origin_max_http_version"
)

func (r SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionIDOriginMaxHTTPVersion:
		return true
	}
	return false
}

// Value of the Origin Max HTTP Version Setting.
type SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionValue string

const (
	SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionValue2 SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionValue = "2"
	SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionValue1 SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionValue = "1"
)

func (r SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionValue2, SettingGetResponseZonesCacheRulesOriginMaxHTTPVersionValue1:
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
type SettingGetResponseZonesPolish2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesPolish2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesPolish2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesPolish2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesPolish2JSON `json:"-"`
}

// settingGetResponseZonesPolish2JSON contains the JSON metadata for the struct
// [SettingGetResponseZonesPolish2]
type settingGetResponseZonesPolish2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesPolish2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesPolish2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesPolish2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesPolish2ID string

const (
	SettingGetResponseZonesPolish2IDPolish SettingGetResponseZonesPolish2ID = "polish"
)

func (r SettingGetResponseZonesPolish2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesPolish2IDPolish:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesPolish2Value string

const (
	SettingGetResponseZonesPolish2ValueOff      SettingGetResponseZonesPolish2Value = "off"
	SettingGetResponseZonesPolish2ValueLossless SettingGetResponseZonesPolish2Value = "lossless"
	SettingGetResponseZonesPolish2ValueLossy    SettingGetResponseZonesPolish2Value = "lossy"
)

func (r SettingGetResponseZonesPolish2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesPolish2ValueOff, SettingGetResponseZonesPolish2ValueLossless, SettingGetResponseZonesPolish2ValueLossy:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesPolish2Editable bool

const (
	SettingGetResponseZonesPolish2EditableTrue  SettingGetResponseZonesPolish2Editable = true
	SettingGetResponseZonesPolish2EditableFalse SettingGetResponseZonesPolish2Editable = false
)

func (r SettingGetResponseZonesPolish2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesPolish2EditableTrue, SettingGetResponseZonesPolish2EditableFalse:
		return true
	}
	return false
}

// Privacy Pass v1 was a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors by allowing users to reduce
// the number of CAPTCHAs shown.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
//
// Deprecated: Privacy Pass v1 was deprecated in 2023. (Announcement -
// https://blog.cloudflare.com/privacy-pass-standard/) and (API deprecation
// details -
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2024-03-31)
type SettingGetResponseZonesPrivacyPass struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesPrivacyPassID `json:"id" api:"required"`
	// Current value of the zone setting.
	//
	// Deprecated: Privacy Pass v1 was deprecated in 2023. (Announcement -
	// https://blog.cloudflare.com/privacy-pass-standard/) and (API deprecation
	// details -
	// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2024-03-31)
	Value SettingGetResponseZonesPrivacyPassValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesPrivacyPassEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesPrivacyPassJSON `json:"-"`
}

// settingGetResponseZonesPrivacyPassJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesPrivacyPass]
type settingGetResponseZonesPrivacyPassJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesPrivacyPass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesPrivacyPassJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesPrivacyPass) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesPrivacyPassID string

const (
	SettingGetResponseZonesPrivacyPassIDPrivacyPass SettingGetResponseZonesPrivacyPassID = "privacy_pass"
)

func (r SettingGetResponseZonesPrivacyPassID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesPrivacyPassIDPrivacyPass:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesPrivacyPassValue string

const (
	SettingGetResponseZonesPrivacyPassValueOn  SettingGetResponseZonesPrivacyPassValue = "on"
	SettingGetResponseZonesPrivacyPassValueOff SettingGetResponseZonesPrivacyPassValue = "off"
)

func (r SettingGetResponseZonesPrivacyPassValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesPrivacyPassValueOn, SettingGetResponseZonesPrivacyPassValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesPrivacyPassEditable bool

const (
	SettingGetResponseZonesPrivacyPassEditableTrue  SettingGetResponseZonesPrivacyPassEditable = true
	SettingGetResponseZonesPrivacyPassEditableFalse SettingGetResponseZonesPrivacyPassEditable = false
)

func (r SettingGetResponseZonesPrivacyPassEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesPrivacyPassEditableTrue, SettingGetResponseZonesPrivacyPassEditableFalse:
		return true
	}
	return false
}

// When enabled, Cloudflare will redirect verified AI training crawlers to
// canonical URLs found in the HTML response, ensuring AI models train on
// authoritative content.
type SettingGetResponseZonesRedirectsForAITraining struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesRedirectsForAITrainingID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesRedirectsForAITrainingValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesRedirectsForAITrainingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesRedirectsForAITrainingJSON `json:"-"`
}

// settingGetResponseZonesRedirectsForAITrainingJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesRedirectsForAITraining]
type settingGetResponseZonesRedirectsForAITrainingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesRedirectsForAITraining) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesRedirectsForAITrainingJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesRedirectsForAITraining) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesRedirectsForAITrainingID string

const (
	SettingGetResponseZonesRedirectsForAITrainingIDRedirectsForAITraining SettingGetResponseZonesRedirectsForAITrainingID = "redirects_for_ai_training"
)

func (r SettingGetResponseZonesRedirectsForAITrainingID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesRedirectsForAITrainingIDRedirectsForAITraining:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesRedirectsForAITrainingValue string

const (
	SettingGetResponseZonesRedirectsForAITrainingValueOff SettingGetResponseZonesRedirectsForAITrainingValue = "off"
	SettingGetResponseZonesRedirectsForAITrainingValueOn  SettingGetResponseZonesRedirectsForAITrainingValue = "on"
)

func (r SettingGetResponseZonesRedirectsForAITrainingValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesRedirectsForAITrainingValueOff, SettingGetResponseZonesRedirectsForAITrainingValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesRedirectsForAITrainingEditable bool

const (
	SettingGetResponseZonesRedirectsForAITrainingEditableTrue  SettingGetResponseZonesRedirectsForAITrainingEditable = true
	SettingGetResponseZonesRedirectsForAITrainingEditableFalse SettingGetResponseZonesRedirectsForAITrainingEditable = false
)

func (r SettingGetResponseZonesRedirectsForAITrainingEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesRedirectsForAITrainingEditableTrue, SettingGetResponseZonesRedirectsForAITrainingEditableFalse:
		return true
	}
	return false
}

// Automatically replace insecure JavaScript libraries with safer and faster
// alternatives provided under cdnjs and powered by Cloudflare. Currently supports
// the following libraries: Polyfill under polyfill.io.
type SettingGetResponseZonesReplaceInsecureJS struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesReplaceInsecureJSID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesReplaceInsecureJSValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesReplaceInsecureJSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingGetResponseZonesReplaceInsecureJS) implementsSettingGetResponse() {}

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
//
// Deprecated: This zone setting is deprecated. This functionality is no longer
// supported.
type SettingGetResponseZonesResponseBuffering2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesResponseBuffering2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesResponseBuffering2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesResponseBuffering2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesResponseBuffering2JSON `json:"-"`
}

// settingGetResponseZonesResponseBuffering2JSON contains the JSON metadata for the
// struct [SettingGetResponseZonesResponseBuffering2]
type settingGetResponseZonesResponseBuffering2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesResponseBuffering2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesResponseBuffering2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesResponseBuffering2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesResponseBuffering2ID string

const (
	SettingGetResponseZonesResponseBuffering2IDResponseBuffering SettingGetResponseZonesResponseBuffering2ID = "response_buffering"
)

func (r SettingGetResponseZonesResponseBuffering2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesResponseBuffering2IDResponseBuffering:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesResponseBuffering2Value string

const (
	SettingGetResponseZonesResponseBuffering2ValueOn  SettingGetResponseZonesResponseBuffering2Value = "on"
	SettingGetResponseZonesResponseBuffering2ValueOff SettingGetResponseZonesResponseBuffering2Value = "off"
)

func (r SettingGetResponseZonesResponseBuffering2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesResponseBuffering2ValueOn, SettingGetResponseZonesResponseBuffering2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesResponseBuffering2Editable bool

const (
	SettingGetResponseZonesResponseBuffering2EditableTrue  SettingGetResponseZonesResponseBuffering2Editable = true
	SettingGetResponseZonesResponseBuffering2EditableFalse SettingGetResponseZonesResponseBuffering2Editable = false
)

func (r SettingGetResponseZonesResponseBuffering2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesResponseBuffering2EditableTrue, SettingGetResponseZonesResponseBuffering2EditableFalse:
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
type SettingGetResponseZonesRocketLoader2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesRocketLoader2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesRocketLoader2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesRocketLoader2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesRocketLoader2JSON `json:"-"`
}

// settingGetResponseZonesRocketLoader2JSON contains the JSON metadata for the
// struct [SettingGetResponseZonesRocketLoader2]
type settingGetResponseZonesRocketLoader2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesRocketLoader2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesRocketLoader2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesRocketLoader2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesRocketLoader2ID string

const (
	SettingGetResponseZonesRocketLoader2IDRocketLoader SettingGetResponseZonesRocketLoader2ID = "rocket_loader"
)

func (r SettingGetResponseZonesRocketLoader2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesRocketLoader2IDRocketLoader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesRocketLoader2Value string

const (
	SettingGetResponseZonesRocketLoader2ValueOn  SettingGetResponseZonesRocketLoader2Value = "on"
	SettingGetResponseZonesRocketLoader2ValueOff SettingGetResponseZonesRocketLoader2Value = "off"
)

func (r SettingGetResponseZonesRocketLoader2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesRocketLoader2ValueOn, SettingGetResponseZonesRocketLoader2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesRocketLoader2Editable bool

const (
	SettingGetResponseZonesRocketLoader2EditableTrue  SettingGetResponseZonesRocketLoader2Editable = true
	SettingGetResponseZonesRocketLoader2EditableFalse SettingGetResponseZonesRocketLoader2Editable = false
)

func (r SettingGetResponseZonesRocketLoader2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesRocketLoader2EditableTrue, SettingGetResponseZonesRocketLoader2EditableFalse:
		return true
	}
	return false
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingGetResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value AutomaticPlatformOptimization `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                       `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingGetResponseZonesSchemasAutomaticPlatformOptimization) implementsSettingGetResponse() {}

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

// When enabled, Cloudflare provisions an AI Search instance for the zone and
// exposes a /.well-known/ai-search endpoint that AI agents can query. Markdown
// responses also receive an agent: YAML capability block advertising the search
// endpoint.
type SettingGetResponseZonesSearchForAgents struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSearchForAgentsID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSearchForAgentsValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSearchForAgentsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesSearchForAgentsJSON `json:"-"`
}

// settingGetResponseZonesSearchForAgentsJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesSearchForAgents]
type settingGetResponseZonesSearchForAgentsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSearchForAgents) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSearchForAgentsJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSearchForAgents) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSearchForAgentsID string

const (
	SettingGetResponseZonesSearchForAgentsIDSearchForAgents SettingGetResponseZonesSearchForAgentsID = "search_for_agents"
)

func (r SettingGetResponseZonesSearchForAgentsID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSearchForAgentsIDSearchForAgents:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSearchForAgentsValue string

const (
	SettingGetResponseZonesSearchForAgentsValueOff SettingGetResponseZonesSearchForAgentsValue = "off"
	SettingGetResponseZonesSearchForAgentsValueOn  SettingGetResponseZonesSearchForAgentsValue = "on"
)

func (r SettingGetResponseZonesSearchForAgentsValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSearchForAgentsValueOff, SettingGetResponseZonesSearchForAgentsValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSearchForAgentsEditable bool

const (
	SettingGetResponseZonesSearchForAgentsEditableTrue  SettingGetResponseZonesSearchForAgentsEditable = true
	SettingGetResponseZonesSearchForAgentsEditableFalse SettingGetResponseZonesSearchForAgentsEditable = false
)

func (r SettingGetResponseZonesSearchForAgentsEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSearchForAgentsEditableTrue, SettingGetResponseZonesSearchForAgentsEditableFalse:
		return true
	}
	return false
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingGetResponseZonesSecurityLevel2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSecurityLevel2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSecurityLevel2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSecurityLevel2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesSecurityLevel2JSON `json:"-"`
}

// settingGetResponseZonesSecurityLevel2JSON contains the JSON metadata for the
// struct [SettingGetResponseZonesSecurityLevel2]
type settingGetResponseZonesSecurityLevel2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSecurityLevel2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSecurityLevel2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSecurityLevel2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSecurityLevel2ID string

const (
	SettingGetResponseZonesSecurityLevel2IDSecurityLevel SettingGetResponseZonesSecurityLevel2ID = "security_level"
)

func (r SettingGetResponseZonesSecurityLevel2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSecurityLevel2IDSecurityLevel:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSecurityLevel2Value string

const (
	SettingGetResponseZonesSecurityLevel2ValueOff            SettingGetResponseZonesSecurityLevel2Value = "off"
	SettingGetResponseZonesSecurityLevel2ValueEssentiallyOff SettingGetResponseZonesSecurityLevel2Value = "essentially_off"
	SettingGetResponseZonesSecurityLevel2ValueLow            SettingGetResponseZonesSecurityLevel2Value = "low"
	SettingGetResponseZonesSecurityLevel2ValueMedium         SettingGetResponseZonesSecurityLevel2Value = "medium"
	SettingGetResponseZonesSecurityLevel2ValueHigh           SettingGetResponseZonesSecurityLevel2Value = "high"
	SettingGetResponseZonesSecurityLevel2ValueUnderAttack    SettingGetResponseZonesSecurityLevel2Value = "under_attack"
)

func (r SettingGetResponseZonesSecurityLevel2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSecurityLevel2ValueOff, SettingGetResponseZonesSecurityLevel2ValueEssentiallyOff, SettingGetResponseZonesSecurityLevel2ValueLow, SettingGetResponseZonesSecurityLevel2ValueMedium, SettingGetResponseZonesSecurityLevel2ValueHigh, SettingGetResponseZonesSecurityLevel2ValueUnderAttack:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSecurityLevel2Editable bool

const (
	SettingGetResponseZonesSecurityLevel2EditableTrue  SettingGetResponseZonesSecurityLevel2Editable = true
	SettingGetResponseZonesSecurityLevel2EditableFalse SettingGetResponseZonesSecurityLevel2Editable = false
)

func (r SettingGetResponseZonesSecurityLevel2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSecurityLevel2EditableTrue, SettingGetResponseZonesSecurityLevel2EditableFalse:
		return true
	}
	return false
}

// Allow SHA1 support.
type SettingGetResponseZonesSha1Support struct {
	// Zone setting identifier.
	ID SettingGetResponseZonesSha1SupportID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSha1SupportValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingGetResponseZonesSha1Support) implementsSettingGetResponse() {}

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
type SettingGetResponseZonesSortQueryStringForCache2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSortQueryStringForCache2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSortQueryStringForCache2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSortQueryStringForCache2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesSortQueryStringForCache2JSON `json:"-"`
}

// settingGetResponseZonesSortQueryStringForCache2JSON contains the JSON metadata
// for the struct [SettingGetResponseZonesSortQueryStringForCache2]
type settingGetResponseZonesSortQueryStringForCache2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSortQueryStringForCache2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSortQueryStringForCache2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSortQueryStringForCache2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSortQueryStringForCache2ID string

const (
	SettingGetResponseZonesSortQueryStringForCache2IDSortQueryStringForCache SettingGetResponseZonesSortQueryStringForCache2ID = "sort_query_string_for_cache"
)

func (r SettingGetResponseZonesSortQueryStringForCache2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSortQueryStringForCache2IDSortQueryStringForCache:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSortQueryStringForCache2Value string

const (
	SettingGetResponseZonesSortQueryStringForCache2ValueOn  SettingGetResponseZonesSortQueryStringForCache2Value = "on"
	SettingGetResponseZonesSortQueryStringForCache2ValueOff SettingGetResponseZonesSortQueryStringForCache2Value = "off"
)

func (r SettingGetResponseZonesSortQueryStringForCache2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSortQueryStringForCache2ValueOn, SettingGetResponseZonesSortQueryStringForCache2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSortQueryStringForCache2Editable bool

const (
	SettingGetResponseZonesSortQueryStringForCache2EditableTrue  SettingGetResponseZonesSortQueryStringForCache2Editable = true
	SettingGetResponseZonesSortQueryStringForCache2EditableFalse SettingGetResponseZonesSortQueryStringForCache2Editable = false
)

func (r SettingGetResponseZonesSortQueryStringForCache2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSortQueryStringForCache2EditableTrue, SettingGetResponseZonesSortQueryStringForCache2EditableFalse:
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
type SettingGetResponseZonesSSL2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSSL2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSSL2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSSL2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesSsl2JSON `json:"-"`
}

// settingGetResponseZonesSsl2JSON contains the JSON metadata for the struct
// [SettingGetResponseZonesSSL2]
type settingGetResponseZonesSsl2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSSL2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSsl2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSSL2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSSL2ID string

const (
	SettingGetResponseZonesSSL2IDSSL SettingGetResponseZonesSSL2ID = "ssl"
)

func (r SettingGetResponseZonesSSL2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSSL2IDSSL:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesSSL2Value string

const (
	SettingGetResponseZonesSSL2ValueOff      SettingGetResponseZonesSSL2Value = "off"
	SettingGetResponseZonesSSL2ValueFlexible SettingGetResponseZonesSSL2Value = "flexible"
	SettingGetResponseZonesSSL2ValueFull     SettingGetResponseZonesSSL2Value = "full"
	SettingGetResponseZonesSSL2ValueStrict   SettingGetResponseZonesSSL2Value = "strict"
)

func (r SettingGetResponseZonesSSL2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSSL2ValueOff, SettingGetResponseZonesSSL2ValueFlexible, SettingGetResponseZonesSSL2ValueFull, SettingGetResponseZonesSSL2ValueStrict:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSSL2Editable bool

const (
	SettingGetResponseZonesSSL2EditableTrue  SettingGetResponseZonesSSL2Editable = true
	SettingGetResponseZonesSSL2EditableFalse SettingGetResponseZonesSSL2Editable = false
)

func (r SettingGetResponseZonesSSL2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesSSL2EditableTrue, SettingGetResponseZonesSSL2EditableFalse:
		return true
	}
	return false
}

// Only allows TLS1.2.
type SettingGetResponseZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID SettingGetResponseZonesTLS1_2OnlyID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesTLS1_2OnlyValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesTLS1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on" api:"nullable" format:"date-time"`
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

func (r SettingGetResponseZonesTLS1_2Only) implementsSettingGetResponse() {}

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

// Media Transformations provides on-demand resizing, conversion and optimization
// for images and video served through Cloudflare's network. Refer to the
// [Image Transformations](https://developers.cloudflare.com/images/) and
// [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started)
// documentation for more information.
type SettingGetResponseZonesTransformations struct {
	// ID of the zone setting. Shared between Image Transformations and Video
	// Transformations.
	ID SettingGetResponseZonesTransformationsID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesTransformationsValue `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesTransformationsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesTransformationsJSON `json:"-"`
}

// settingGetResponseZonesTransformationsJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesTransformations]
type settingGetResponseZonesTransformationsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesTransformations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesTransformationsJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesTransformations) implementsSettingGetResponse() {}

// ID of the zone setting. Shared between Image Transformations and Video
// Transformations.
type SettingGetResponseZonesTransformationsID string

const (
	SettingGetResponseZonesTransformationsIDTransformations SettingGetResponseZonesTransformationsID = "transformations"
)

func (r SettingGetResponseZonesTransformationsID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTransformationsIDTransformations:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesTransformationsValue string

const (
	SettingGetResponseZonesTransformationsValueOn   SettingGetResponseZonesTransformationsValue = "on"
	SettingGetResponseZonesTransformationsValueOff  SettingGetResponseZonesTransformationsValue = "off"
	SettingGetResponseZonesTransformationsValueOpen SettingGetResponseZonesTransformationsValue = "open"
)

func (r SettingGetResponseZonesTransformationsValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTransformationsValueOn, SettingGetResponseZonesTransformationsValueOff, SettingGetResponseZonesTransformationsValueOpen:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesTransformationsEditable bool

const (
	SettingGetResponseZonesTransformationsEditableTrue  SettingGetResponseZonesTransformationsEditable = true
	SettingGetResponseZonesTransformationsEditableFalse SettingGetResponseZonesTransformationsEditable = false
)

func (r SettingGetResponseZonesTransformationsEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTransformationsEditableTrue, SettingGetResponseZonesTransformationsEditableFalse:
		return true
	}
	return false
}

// Media Transformations Allowed Origins restricts transformations for images and
// video served through Cloudflare's network. Refer to the
// [Image Transformations](https://developers.cloudflare.com/images/) and
// [Video Transformations](https://developers.cloudflare.com/stream/transform-videos/#getting-started)
// documentation for more information.
type SettingGetResponseZonesTransformationsAllowedOrigins struct {
	// ID of the zone setting. Shared between Image Transformations and Video
	// Transformations.
	ID SettingGetResponseZonesTransformationsAllowedOriginsID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value string `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesTransformationsAllowedOriginsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesTransformationsAllowedOriginsJSON `json:"-"`
}

// settingGetResponseZonesTransformationsAllowedOriginsJSON contains the JSON
// metadata for the struct [SettingGetResponseZonesTransformationsAllowedOrigins]
type settingGetResponseZonesTransformationsAllowedOriginsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesTransformationsAllowedOrigins) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesTransformationsAllowedOriginsJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesTransformationsAllowedOrigins) implementsSettingGetResponse() {}

// ID of the zone setting. Shared between Image Transformations and Video
// Transformations.
type SettingGetResponseZonesTransformationsAllowedOriginsID string

const (
	SettingGetResponseZonesTransformationsAllowedOriginsIDTransformationsAllowedOrigins SettingGetResponseZonesTransformationsAllowedOriginsID = "transformations_allowed_origins"
)

func (r SettingGetResponseZonesTransformationsAllowedOriginsID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTransformationsAllowedOriginsIDTransformationsAllowedOrigins:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesTransformationsAllowedOriginsEditable bool

const (
	SettingGetResponseZonesTransformationsAllowedOriginsEditableTrue  SettingGetResponseZonesTransformationsAllowedOriginsEditable = true
	SettingGetResponseZonesTransformationsAllowedOriginsEditableFalse SettingGetResponseZonesTransformationsAllowedOriginsEditable = false
)

func (r SettingGetResponseZonesTransformationsAllowedOriginsEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTransformationsAllowedOriginsEditableTrue, SettingGetResponseZonesTransformationsAllowedOriginsEditableFalse:
		return true
	}
	return false
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingGetResponseZonesTrueClientIPHeader2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesTrueClientIPHeader2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesTrueClientIPHeader2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesTrueClientIPHeader2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesTrueClientIPHeader2JSON `json:"-"`
}

// settingGetResponseZonesTrueClientIPHeader2JSON contains the JSON metadata for
// the struct [SettingGetResponseZonesTrueClientIPHeader2]
type settingGetResponseZonesTrueClientIPHeader2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesTrueClientIPHeader2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesTrueClientIPHeader2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesTrueClientIPHeader2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesTrueClientIPHeader2ID string

const (
	SettingGetResponseZonesTrueClientIPHeader2IDTrueClientIPHeader SettingGetResponseZonesTrueClientIPHeader2ID = "true_client_ip_header"
)

func (r SettingGetResponseZonesTrueClientIPHeader2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTrueClientIPHeader2IDTrueClientIPHeader:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesTrueClientIPHeader2Value string

const (
	SettingGetResponseZonesTrueClientIPHeader2ValueOn  SettingGetResponseZonesTrueClientIPHeader2Value = "on"
	SettingGetResponseZonesTrueClientIPHeader2ValueOff SettingGetResponseZonesTrueClientIPHeader2Value = "off"
)

func (r SettingGetResponseZonesTrueClientIPHeader2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTrueClientIPHeader2ValueOn, SettingGetResponseZonesTrueClientIPHeader2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesTrueClientIPHeader2Editable bool

const (
	SettingGetResponseZonesTrueClientIPHeader2EditableTrue  SettingGetResponseZonesTrueClientIPHeader2Editable = true
	SettingGetResponseZonesTrueClientIPHeader2EditableFalse SettingGetResponseZonesTrueClientIPHeader2Editable = false
)

func (r SettingGetResponseZonesTrueClientIPHeader2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesTrueClientIPHeader2EditableTrue, SettingGetResponseZonesTrueClientIPHeader2EditableFalse:
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
type SettingGetResponseZonesWAF2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesWAF2ID `json:"id" api:"required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesWAF2Value `json:"value" api:"required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesWAF2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       settingGetResponseZonesWaf2JSON `json:"-"`
}

// settingGetResponseZonesWaf2JSON contains the JSON metadata for the struct
// [SettingGetResponseZonesWAF2]
type settingGetResponseZonesWaf2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesWAF2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesWaf2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesWAF2) implementsSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesWAF2ID string

const (
	SettingGetResponseZonesWAF2IDWAF SettingGetResponseZonesWAF2ID = "waf"
)

func (r SettingGetResponseZonesWAF2ID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesWAF2IDWAF:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingGetResponseZonesWAF2Value string

const (
	SettingGetResponseZonesWAF2ValueOn  SettingGetResponseZonesWAF2Value = "on"
	SettingGetResponseZonesWAF2ValueOff SettingGetResponseZonesWAF2Value = "off"
)

func (r SettingGetResponseZonesWAF2Value) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesWAF2ValueOn, SettingGetResponseZonesWAF2ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesWAF2Editable bool

const (
	SettingGetResponseZonesWAF2EditableTrue  SettingGetResponseZonesWAF2Editable = true
	SettingGetResponseZonesWAF2EditableFalse SettingGetResponseZonesWAF2Editable = false
)

func (r SettingGetResponseZonesWAF2Editable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesWAF2EditableTrue, SettingGetResponseZonesWAF2EditableFalse:
		return true
	}
	return false
}

// ID of the zone setting.
type SettingGetResponseID string

const (
	SettingGetResponseID0rtt                          SettingGetResponseID = "0rtt"
	SettingGetResponseIDAdvancedDDoS                  SettingGetResponseID = "advanced_ddos"
	SettingGetResponseIDAegis                         SettingGetResponseID = "aegis"
	SettingGetResponseIDAlwaysOnline                  SettingGetResponseID = "always_online"
	SettingGetResponseIDAlwaysUseHTTPS                SettingGetResponseID = "always_use_https"
	SettingGetResponseIDAutomaticHTTPSRewrites        SettingGetResponseID = "automatic_https_rewrites"
	SettingGetResponseIDBrotli                        SettingGetResponseID = "brotli"
	SettingGetResponseIDBrowserCacheTTL               SettingGetResponseID = "browser_cache_ttl"
	SettingGetResponseIDBrowserCheck                  SettingGetResponseID = "browser_check"
	SettingGetResponseIDCacheLevel                    SettingGetResponseID = "cache_level"
	SettingGetResponseIDChallengeTTL                  SettingGetResponseID = "challenge_ttl"
	SettingGetResponseIDChinaNetworkEnabled           SettingGetResponseID = "china_network_enabled"
	SettingGetResponseIDContentConverter              SettingGetResponseID = "content_converter"
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
	SettingGetResponseIDOriginH2MaxStreams            SettingGetResponseID = "origin_h2_max_streams"
	SettingGetResponseIDOriginMaxHTTPVersion          SettingGetResponseID = "origin_max_http_version"
	SettingGetResponseIDPolish                        SettingGetResponseID = "polish"
	SettingGetResponseIDPrefetchPreload               SettingGetResponseID = "prefetch_preload"
	SettingGetResponseIDPrivacyPass                   SettingGetResponseID = "privacy_pass"
	SettingGetResponseIDProxyReadTimeout              SettingGetResponseID = "proxy_read_timeout"
	SettingGetResponseIDPseudoIPV4                    SettingGetResponseID = "pseudo_ipv4"
	SettingGetResponseIDRedirectsForAITraining        SettingGetResponseID = "redirects_for_ai_training"
	SettingGetResponseIDReplaceInsecureJS             SettingGetResponseID = "replace_insecure_js"
	SettingGetResponseIDResponseBuffering             SettingGetResponseID = "response_buffering"
	SettingGetResponseIDRocketLoader                  SettingGetResponseID = "rocket_loader"
	SettingGetResponseIDAutomaticPlatformOptimization SettingGetResponseID = "automatic_platform_optimization"
	SettingGetResponseIDSearchForAgents               SettingGetResponseID = "search_for_agents"
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
	SettingGetResponseIDTransformations               SettingGetResponseID = "transformations"
	SettingGetResponseIDTransformationsAllowedOrigins SettingGetResponseID = "transformations_allowed_origins"
	SettingGetResponseIDTrueClientIPHeader            SettingGetResponseID = "true_client_ip_header"
	SettingGetResponseIDWAF                           SettingGetResponseID = "waf"
	SettingGetResponseIDWebP                          SettingGetResponseID = "webp"
	SettingGetResponseIDWebsockets                    SettingGetResponseID = "websockets"
)

func (r SettingGetResponseID) IsKnown() bool {
	switch r {
	case SettingGetResponseID0rtt, SettingGetResponseIDAdvancedDDoS, SettingGetResponseIDAegis, SettingGetResponseIDAlwaysOnline, SettingGetResponseIDAlwaysUseHTTPS, SettingGetResponseIDAutomaticHTTPSRewrites, SettingGetResponseIDBrotli, SettingGetResponseIDBrowserCacheTTL, SettingGetResponseIDBrowserCheck, SettingGetResponseIDCacheLevel, SettingGetResponseIDChallengeTTL, SettingGetResponseIDChinaNetworkEnabled, SettingGetResponseIDContentConverter, SettingGetResponseIDCiphers, SettingGetResponseIDCNAMEFlattening, SettingGetResponseIDDevelopmentMode, SettingGetResponseIDEarlyHints, SettingGetResponseIDEdgeCacheTTL, SettingGetResponseIDEmailObfuscation, SettingGetResponseIDH2Prioritization, SettingGetResponseIDHotlinkProtection, SettingGetResponseIDHTTP2, SettingGetResponseIDHTTP3, SettingGetResponseIDImageResizing, SettingGetResponseIDIPGeolocation, SettingGetResponseIDIPV6, SettingGetResponseIDMaxUpload, SettingGetResponseIDMinTLSVersion, SettingGetResponseIDMirage, SettingGetResponseIDNEL, SettingGetResponseIDOpportunisticEncryption, SettingGetResponseIDOpportunisticOnion, SettingGetResponseIDOrangeToOrange, SettingGetResponseIDOriginErrorPagePassThru, SettingGetResponseIDOriginH2MaxStreams, SettingGetResponseIDOriginMaxHTTPVersion, SettingGetResponseIDPolish, SettingGetResponseIDPrefetchPreload, SettingGetResponseIDPrivacyPass, SettingGetResponseIDProxyReadTimeout, SettingGetResponseIDPseudoIPV4, SettingGetResponseIDRedirectsForAITraining, SettingGetResponseIDReplaceInsecureJS, SettingGetResponseIDResponseBuffering, SettingGetResponseIDRocketLoader, SettingGetResponseIDAutomaticPlatformOptimization, SettingGetResponseIDSearchForAgents, SettingGetResponseIDSecurityHeader, SettingGetResponseIDSecurityLevel, SettingGetResponseIDServerSideExclude, SettingGetResponseIDSha1Support, SettingGetResponseIDSortQueryStringForCache, SettingGetResponseIDSSL, SettingGetResponseIDSSLRecommender, SettingGetResponseIDTLS1_2Only, SettingGetResponseIDTLS1_3, SettingGetResponseIDTLSClientAuth, SettingGetResponseIDTransformations, SettingGetResponseIDTransformationsAllowedOrigins, SettingGetResponseIDTrueClientIPHeader, SettingGetResponseIDWAF, SettingGetResponseIDWebP, SettingGetResponseIDWebsockets:
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
	ZoneID param.Field[string]        `path:"zone_id" api:"required"`
	Body   SettingEditParamsBodyUnion `json:"body" api:"required"`
}

func (r SettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SettingEditParamsBody struct {
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool]        `json:"enabled"`
	Value   param.Field[interface{}] `json:"value"`
}

func (r SettingEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBody) implementsSettingEditParamsBodyUnion() {}

// Satisfied by [zones.SettingEditParamsBodyEnabled],
// [zones.SettingEditParamsBodyValue], [SettingEditParamsBody].
type SettingEditParamsBodyUnion interface {
	implementsSettingEditParamsBodyUnion()
}

type SettingEditParamsBodyEnabled struct {
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SettingEditParamsBodyEnabled) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyEnabled) implementsSettingEditParamsBodyUnion() {}

type SettingEditParamsBodyValue struct {
	// Value of the zone setting.
	Value param.Field[SettingEditParamsBodyValueValueUnion] `json:"value"`
}

func (r SettingEditParamsBodyValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyValue) implementsSettingEditParamsBodyUnion() {}

// Value of the zone setting.
type SettingEditParamsBodyValueValue struct {
	// Indicates whether or not
	// [cache by device type](https://developers.cloudflare.com/automatic-platform-optimization/reference/cache-device-type/)
	// is enabled.
	CacheByDeviceType param.Field[bool] `json:"cache_by_device_type"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf param.Field[bool] `json:"cf"`
	// Whether the feature is enabled or not.
	Enabled   param.Field[bool]        `json:"enabled"`
	Hostnames param.Field[interface{}] `json:"hostnames"`
	// Egress pool id which refers to a grouping of dedicated egress IPs through which
	// Cloudflare will connect to origin.
	PoolID                  param.Field[string]      `json:"pool_id"`
	StrictTransportSecurity param.Field[interface{}] `json:"strict_transport_security"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress param.Field[bool] `json:"wordpress"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin param.Field[bool] `json:"wp_plugin"`
}

func (r SettingEditParamsBodyValueValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyValueValue) ImplementsSettingEditParamsBodyValueValueUnion() {}

// Value of the zone setting.
//
// Satisfied by [zones.SettingEditParamsBodyValueValueZonesCacheRulesAegisValue],
// [zones.SettingEditParamsBodyValueValueZonesCiphersValue],
// [zones.SettingEditParamsBodyValueValueZonesNELValue], [shared.UnionFloat],
// [zones.AutomaticPlatformOptimizationParam],
// [zones.SettingEditParamsBodyValueValueZonesSecurityHeaderValue],
// [SettingEditParamsBodyValueValue].
//
// Use [Raw()] to specify an arbitrary value for this param
type SettingEditParamsBodyValueValueUnion interface {
	ImplementsSettingEditParamsBodyValueValueUnion()
}

// Value of the zone setting.
type SettingEditParamsBodyValueValueZonesCacheRulesAegisValue struct {
	// Whether the feature is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
	// Egress pool id which refers to a grouping of dedicated egress IPs through which
	// Cloudflare will connect to origin.
	PoolID param.Field[string] `json:"pool_id"`
}

func (r SettingEditParamsBodyValueValueZonesCacheRulesAegisValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyValueValueZonesCacheRulesAegisValue) ImplementsSettingEditParamsBodyValueValueUnion() {
}

type SettingEditParamsBodyValueValueZonesCiphersValue []string

func (r SettingEditParamsBodyValueValueZonesCiphersValue) ImplementsSettingEditParamsBodyValueValueUnion() {
}

// Value of the zone setting.
type SettingEditParamsBodyValueValueZonesNELValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SettingEditParamsBodyValueValueZonesNELValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyValueValueZonesNELValue) ImplementsSettingEditParamsBodyValueValueUnion() {
}

type SettingEditParamsBodyValueValueZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[SettingEditParamsBodyValueValueZonesSecurityHeaderValueStrictTransportSecurity] `json:"strict_transport_security"`
}

func (r SettingEditParamsBodyValueValueZonesSecurityHeaderValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsBodyValueValueZonesSecurityHeaderValue) ImplementsSettingEditParamsBodyValueValueUnion() {
}

// Strict Transport Security.
type SettingEditParamsBodyValueValueZonesSecurityHeaderValueStrictTransportSecurity struct {
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

func (r SettingEditParamsBodyValueValueZonesSecurityHeaderValueStrictTransportSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful
	Success bool `json:"success" api:"required"`
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
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type SettingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful
	Success bool `json:"success" api:"required"`
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
