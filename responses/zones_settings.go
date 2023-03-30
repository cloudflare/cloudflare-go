package responses

import (
	"time"

	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type ZeroRtt struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZeroRttEditable `json:"editable"`
	// ID of the zone setting.
	ID ZeroRttID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the 0-RTT setting.
	Value ZeroRttValue `json:"value"`
	JSON  ZeroRttJSON
}

type ZeroRttJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ZeroRtt using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *ZeroRtt) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ZeroRttEditable bool

const (
	ZeroRttEditableTrue  ZeroRttEditable = true
	ZeroRttEditableFalse ZeroRttEditable = false
)

type ZeroRttID string

const (
	ZeroRttID_0rtt ZeroRttID = "0rtt"
)

type ZeroRttValue string

const (
	ZeroRttValueOn  ZeroRttValue = "on"
	ZeroRttValueOff ZeroRttValue = "off"
)

type ZoneSettingsCollection struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool                           `json:"success"`
	Result  []ZoneSettingsCollectionResult `json:"result"`
	JSON    ZoneSettingsCollectionJSON
}

type ZoneSettingsCollectionJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ZoneSettingsCollection using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ZoneSettingsCollection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

// 🚧 Undiscriminated unions are still being implemented.
type ZoneSettingsCollectionResult struct{}

func (r *ZoneSettingsCollectionResult) MarshalJSON() ([]byte, error) { return nil, nil }

type AdvancedDdos struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AdvancedDdosEditable `json:"editable"`
	// ID of the zone setting.
	ID AdvancedDdosID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value AdvancedDdosValue `json:"value"`
	JSON  AdvancedDdosJSON
}

type AdvancedDdosJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AdvancedDdos using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AdvancedDdos) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AdvancedDdosEditable bool

const (
	AdvancedDdosEditableTrue  AdvancedDdosEditable = true
	AdvancedDdosEditableFalse AdvancedDdosEditable = false
)

type AdvancedDdosID string

const (
	AdvancedDdosIDAdvancedDdos AdvancedDdosID = "advanced_ddos"
)

type AdvancedDdosValue string

const (
	AdvancedDdosValueOn  AdvancedDdosValue = "on"
	AdvancedDdosValueOff AdvancedDdosValue = "off"
)

type AlwaysOnline struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AlwaysOnlineEditable `json:"editable"`
	// ID of the zone setting.
	ID AlwaysOnlineID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value AlwaysOnlineValue `json:"value"`
	JSON  AlwaysOnlineJSON
}

type AlwaysOnlineJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AlwaysOnline using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AlwaysOnlineEditable bool

const (
	AlwaysOnlineEditableTrue  AlwaysOnlineEditable = true
	AlwaysOnlineEditableFalse AlwaysOnlineEditable = false
)

type AlwaysOnlineID string

const (
	AlwaysOnlineIDAlwaysOnline AlwaysOnlineID = "always_online"
)

type AlwaysOnlineValue string

const (
	AlwaysOnlineValueOn  AlwaysOnlineValue = "on"
	AlwaysOnlineValueOff AlwaysOnlineValue = "off"
)

type AlwaysUseHTTPs struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AlwaysUseHTTPsEditable `json:"editable"`
	// ID of the zone setting.
	ID AlwaysUseHTTPsID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value AlwaysUseHTTPsValue `json:"value"`
	JSON  AlwaysUseHTTPsJSON
}

type AlwaysUseHTTPsJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AlwaysUseHTTPs using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AlwaysUseHTTPs) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AlwaysUseHTTPsEditable bool

const (
	AlwaysUseHTTPsEditableTrue  AlwaysUseHTTPsEditable = true
	AlwaysUseHTTPsEditableFalse AlwaysUseHTTPsEditable = false
)

type AlwaysUseHTTPsID string

const (
	AlwaysUseHTTPsIDAlwaysUseHTTPs AlwaysUseHTTPsID = "always_use_https"
)

type AlwaysUseHTTPsValue string

const (
	AlwaysUseHTTPsValueOn  AlwaysUseHTTPsValue = "on"
	AlwaysUseHTTPsValueOff AlwaysUseHTTPsValue = "off"
)

type AutomaticHTTPsRewrites struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable AutomaticHTTPsRewritesEditable `json:"editable"`
	// ID of the zone setting.
	ID AutomaticHTTPsRewritesID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value AutomaticHTTPsRewritesValue `json:"value"`
	JSON  AutomaticHTTPsRewritesJSON
}

type AutomaticHTTPsRewritesJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AutomaticHTTPsRewrites using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AutomaticHTTPsRewrites) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AutomaticHTTPsRewritesEditable bool

const (
	AutomaticHTTPsRewritesEditableTrue  AutomaticHTTPsRewritesEditable = true
	AutomaticHTTPsRewritesEditableFalse AutomaticHTTPsRewritesEditable = false
)

type AutomaticHTTPsRewritesID string

const (
	AutomaticHTTPsRewritesIDAutomaticHTTPsRewrites AutomaticHTTPsRewritesID = "automatic_https_rewrites"
)

type AutomaticHTTPsRewritesValue string

const (
	AutomaticHTTPsRewritesValueOn  AutomaticHTTPsRewritesValue = "on"
	AutomaticHTTPsRewritesValueOff AutomaticHTTPsRewritesValue = "off"
)

type Brotli struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable BrotliEditable `json:"editable"`
	// ID of the zone setting.
	ID BrotliID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value BrotliValue `json:"value"`
	JSON  BrotliJSON
}

type BrotliJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Brotli using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Brotli) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BrotliEditable bool

const (
	BrotliEditableTrue  BrotliEditable = true
	BrotliEditableFalse BrotliEditable = false
)

type BrotliID string

const (
	BrotliIDBrotli BrotliID = "brotli"
)

type BrotliValue string

const (
	BrotliValueOff BrotliValue = "off"
	BrotliValueOn  BrotliValue = "on"
)

type BrowserCacheTtl struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable BrowserCacheTtlEditable `json:"editable"`
	// ID of the zone setting.
	ID BrowserCacheTtlID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value BrowserCacheTtlValue `json:"value"`
	JSON  BrowserCacheTtlJSON
}

type BrowserCacheTtlJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BrowserCacheTtl using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BrowserCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BrowserCacheTtlEditable bool

const (
	BrowserCacheTtlEditableTrue  BrowserCacheTtlEditable = true
	BrowserCacheTtlEditableFalse BrowserCacheTtlEditable = false
)

type BrowserCacheTtlID string

const (
	BrowserCacheTtlIDBrowserCacheTtl BrowserCacheTtlID = "browser_cache_ttl"
)

type BrowserCacheTtlValue float64

const (
	BrowserCacheTtlValue_0        BrowserCacheTtlValue = 0
	BrowserCacheTtlValue_30       BrowserCacheTtlValue = 30
	BrowserCacheTtlValue_60       BrowserCacheTtlValue = 60
	BrowserCacheTtlValue_120      BrowserCacheTtlValue = 120
	BrowserCacheTtlValue_300      BrowserCacheTtlValue = 300
	BrowserCacheTtlValue_1200     BrowserCacheTtlValue = 1200
	BrowserCacheTtlValue_1800     BrowserCacheTtlValue = 1800
	BrowserCacheTtlValue_3600     BrowserCacheTtlValue = 3600
	BrowserCacheTtlValue_7200     BrowserCacheTtlValue = 7200
	BrowserCacheTtlValue_10800    BrowserCacheTtlValue = 10800
	BrowserCacheTtlValue_14400    BrowserCacheTtlValue = 14400
	BrowserCacheTtlValue_18000    BrowserCacheTtlValue = 18000
	BrowserCacheTtlValue_28800    BrowserCacheTtlValue = 28800
	BrowserCacheTtlValue_43200    BrowserCacheTtlValue = 43200
	BrowserCacheTtlValue_57600    BrowserCacheTtlValue = 57600
	BrowserCacheTtlValue_72000    BrowserCacheTtlValue = 72000
	BrowserCacheTtlValue_86400    BrowserCacheTtlValue = 86400
	BrowserCacheTtlValue_172800   BrowserCacheTtlValue = 172800
	BrowserCacheTtlValue_259200   BrowserCacheTtlValue = 259200
	BrowserCacheTtlValue_345600   BrowserCacheTtlValue = 345600
	BrowserCacheTtlValue_432000   BrowserCacheTtlValue = 432000
	BrowserCacheTtlValue_691200   BrowserCacheTtlValue = 691200
	BrowserCacheTtlValue_1382400  BrowserCacheTtlValue = 1382400
	BrowserCacheTtlValue_2073600  BrowserCacheTtlValue = 2073600
	BrowserCacheTtlValue_2678400  BrowserCacheTtlValue = 2678400
	BrowserCacheTtlValue_5356800  BrowserCacheTtlValue = 5356800
	BrowserCacheTtlValue_16070400 BrowserCacheTtlValue = 16070400
	BrowserCacheTtlValue_31536000 BrowserCacheTtlValue = 31536000
)

type BrowserCheck struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable BrowserCheckEditable `json:"editable"`
	// ID of the zone setting.
	ID BrowserCheckID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value BrowserCheckValue `json:"value"`
	JSON  BrowserCheckJSON
}

type BrowserCheckJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BrowserCheck using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BrowserCheckEditable bool

const (
	BrowserCheckEditableTrue  BrowserCheckEditable = true
	BrowserCheckEditableFalse BrowserCheckEditable = false
)

type BrowserCheckID string

const (
	BrowserCheckIDBrowserCheck BrowserCheckID = "browser_check"
)

type BrowserCheckValue string

const (
	BrowserCheckValueOn  BrowserCheckValue = "on"
	BrowserCheckValueOff BrowserCheckValue = "off"
)

type CacheLevel struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable CacheLevelEditable `json:"editable"`
	// ID of the zone setting.
	ID CacheLevelID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value CacheLevelValue `json:"value"`
	JSON  CacheLevelJSON
}

type CacheLevelJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CacheLevel using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *CacheLevel) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CacheLevelEditable bool

const (
	CacheLevelEditableTrue  CacheLevelEditable = true
	CacheLevelEditableFalse CacheLevelEditable = false
)

type CacheLevelID string

const (
	CacheLevelIDCacheLevel CacheLevelID = "cache_level"
)

type CacheLevelValue string

const (
	CacheLevelValueAggressive CacheLevelValue = "aggressive"
	CacheLevelValueBasic      CacheLevelValue = "basic"
	CacheLevelValueSimplified CacheLevelValue = "simplified"
)

type ChallengeTtl struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ChallengeTtlEditable `json:"editable"`
	// ID of the zone setting.
	ID ChallengeTtlID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ChallengeTtlValue `json:"value"`
	JSON  ChallengeTtlJSON
}

type ChallengeTtlJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ChallengeTtl using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ChallengeTtl) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ChallengeTtlEditable bool

const (
	ChallengeTtlEditableTrue  ChallengeTtlEditable = true
	ChallengeTtlEditableFalse ChallengeTtlEditable = false
)

type ChallengeTtlID string

const (
	ChallengeTtlIDChallengeTtl ChallengeTtlID = "challenge_ttl"
)

type ChallengeTtlValue float64

const (
	ChallengeTtlValue_300      ChallengeTtlValue = 300
	ChallengeTtlValue_900      ChallengeTtlValue = 900
	ChallengeTtlValue_1800     ChallengeTtlValue = 1800
	ChallengeTtlValue_2700     ChallengeTtlValue = 2700
	ChallengeTtlValue_3600     ChallengeTtlValue = 3600
	ChallengeTtlValue_7200     ChallengeTtlValue = 7200
	ChallengeTtlValue_10800    ChallengeTtlValue = 10800
	ChallengeTtlValue_14400    ChallengeTtlValue = 14400
	ChallengeTtlValue_28800    ChallengeTtlValue = 28800
	ChallengeTtlValue_57600    ChallengeTtlValue = 57600
	ChallengeTtlValue_86400    ChallengeTtlValue = 86400
	ChallengeTtlValue_604800   ChallengeTtlValue = 604800
	ChallengeTtlValue_2592000  ChallengeTtlValue = 2592000
	ChallengeTtlValue_31536000 ChallengeTtlValue = 31536000
)

type Ciphers struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable CiphersEditable `json:"editable"`
	// ID of the zone setting.
	ID CiphersID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value []string `json:"value"`
	JSON  CiphersJSON
}

type CiphersJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Ciphers using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Ciphers) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CiphersEditable bool

const (
	CiphersEditableTrue  CiphersEditable = true
	CiphersEditableFalse CiphersEditable = false
)

type CiphersID string

const (
	CiphersIDCiphers CiphersID = "ciphers"
)

type CnameFlattening struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable CnameFlatteningEditable `json:"editable"`
	// How to flatten the cname destination.
	ID CnameFlatteningID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the cname flattening setting.
	Value CnameFlatteningValue `json:"value"`
	JSON  CnameFlatteningJSON
}

type CnameFlatteningJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into CnameFlattening using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *CnameFlattening) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type CnameFlatteningEditable bool

const (
	CnameFlatteningEditableTrue  CnameFlatteningEditable = true
	CnameFlatteningEditableFalse CnameFlatteningEditable = false
)

type CnameFlatteningID string

const (
	CnameFlatteningIDCnameFlattening CnameFlatteningID = "cname_flattening"
)

type CnameFlatteningValue string

const (
	CnameFlatteningValueFlattenAtRoot CnameFlatteningValue = "flatten_at_root"
	CnameFlatteningValueFlattenAll    CnameFlatteningValue = "flatten_all"
)

type DevelopmentMode struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable DevelopmentModeEditable `json:"editable"`
	// ID of the zone setting.
	ID DevelopmentModeID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value DevelopmentModeValue `json:"value"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	JSON          DevelopmentModeJSON
}

type DevelopmentModeJSON struct {
	Editable      pjson.Metadata
	ID            pjson.Metadata
	ModifiedOn    pjson.Metadata
	Value         pjson.Metadata
	TimeRemaining pjson.Metadata
	Raw           []byte
	Extras        map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DevelopmentMode using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DevelopmentModeEditable bool

const (
	DevelopmentModeEditableTrue  DevelopmentModeEditable = true
	DevelopmentModeEditableFalse DevelopmentModeEditable = false
)

type DevelopmentModeID string

const (
	DevelopmentModeIDDevelopmentMode DevelopmentModeID = "development_mode"
)

type DevelopmentModeValue string

const (
	DevelopmentModeValueOn  DevelopmentModeValue = "on"
	DevelopmentModeValueOff DevelopmentModeValue = "off"
)

type EarlyHints struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable EarlyHintsEditable `json:"editable"`
	// ID of the zone setting.
	ID EarlyHintsID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value EarlyHintsValue `json:"value"`
	JSON  EarlyHintsJSON
}

type EarlyHintsJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EarlyHints using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *EarlyHints) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type EarlyHintsEditable bool

const (
	EarlyHintsEditableTrue  EarlyHintsEditable = true
	EarlyHintsEditableFalse EarlyHintsEditable = false
)

type EarlyHintsID string

const (
	EarlyHintsIDEarlyHints EarlyHintsID = "early_hints"
)

type EarlyHintsValue string

const (
	EarlyHintsValueOn  EarlyHintsValue = "on"
	EarlyHintsValueOff EarlyHintsValue = "off"
)

type EdgeCacheTtl struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable EdgeCacheTtlEditable `json:"editable"`
	// ID of the zone setting.
	ID EdgeCacheTtlID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value EdgeCacheTtlValue `json:"value"`
	JSON  EdgeCacheTtlJSON
}

type EdgeCacheTtlJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EdgeCacheTtl using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EdgeCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type EdgeCacheTtlEditable bool

const (
	EdgeCacheTtlEditableTrue  EdgeCacheTtlEditable = true
	EdgeCacheTtlEditableFalse EdgeCacheTtlEditable = false
)

type EdgeCacheTtlID string

const (
	EdgeCacheTtlIDEdgeCacheTtl EdgeCacheTtlID = "edge_cache_ttl"
)

type EdgeCacheTtlValue float64

const (
	EdgeCacheTtlValue_30     EdgeCacheTtlValue = 30
	EdgeCacheTtlValue_60     EdgeCacheTtlValue = 60
	EdgeCacheTtlValue_300    EdgeCacheTtlValue = 300
	EdgeCacheTtlValue_1200   EdgeCacheTtlValue = 1200
	EdgeCacheTtlValue_1800   EdgeCacheTtlValue = 1800
	EdgeCacheTtlValue_3600   EdgeCacheTtlValue = 3600
	EdgeCacheTtlValue_7200   EdgeCacheTtlValue = 7200
	EdgeCacheTtlValue_10800  EdgeCacheTtlValue = 10800
	EdgeCacheTtlValue_14400  EdgeCacheTtlValue = 14400
	EdgeCacheTtlValue_18000  EdgeCacheTtlValue = 18000
	EdgeCacheTtlValue_28800  EdgeCacheTtlValue = 28800
	EdgeCacheTtlValue_43200  EdgeCacheTtlValue = 43200
	EdgeCacheTtlValue_57600  EdgeCacheTtlValue = 57600
	EdgeCacheTtlValue_72000  EdgeCacheTtlValue = 72000
	EdgeCacheTtlValue_86400  EdgeCacheTtlValue = 86400
	EdgeCacheTtlValue_172800 EdgeCacheTtlValue = 172800
	EdgeCacheTtlValue_259200 EdgeCacheTtlValue = 259200
	EdgeCacheTtlValue_345600 EdgeCacheTtlValue = 345600
	EdgeCacheTtlValue_432000 EdgeCacheTtlValue = 432000
	EdgeCacheTtlValue_518400 EdgeCacheTtlValue = 518400
	EdgeCacheTtlValue_604800 EdgeCacheTtlValue = 604800
)

type EmailObfuscation struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable EmailObfuscationEditable `json:"editable"`
	// ID of the zone setting.
	ID EmailObfuscationID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value EmailObfuscationValue `json:"value"`
	JSON  EmailObfuscationJSON
}

type EmailObfuscationJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EmailObfuscation using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type EmailObfuscationEditable bool

const (
	EmailObfuscationEditableTrue  EmailObfuscationEditable = true
	EmailObfuscationEditableFalse EmailObfuscationEditable = false
)

type EmailObfuscationID string

const (
	EmailObfuscationIDEmailObfuscation EmailObfuscationID = "email_obfuscation"
)

type EmailObfuscationValue string

const (
	EmailObfuscationValueOn  EmailObfuscationValue = "on"
	EmailObfuscationValueOff EmailObfuscationValue = "off"
)

type H2Prioritization struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable H2PrioritizationEditable `json:"editable"`
	// ID of the zone setting.
	ID H2PrioritizationID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value H2PrioritizationValue `json:"value"`
	JSON  H2PrioritizationJSON
}

type H2PrioritizationJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into H2Prioritization using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *H2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type H2PrioritizationEditable bool

const (
	H2PrioritizationEditableTrue  H2PrioritizationEditable = true
	H2PrioritizationEditableFalse H2PrioritizationEditable = false
)

type H2PrioritizationID string

const (
	H2PrioritizationIDH2Prioritization H2PrioritizationID = "h2_prioritization"
)

type H2PrioritizationValue string

const (
	H2PrioritizationValueOn     H2PrioritizationValue = "on"
	H2PrioritizationValueOff    H2PrioritizationValue = "off"
	H2PrioritizationValueCustom H2PrioritizationValue = "custom"
)

type HotlinkProtection struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable HotlinkProtectionEditable `json:"editable"`
	// ID of the zone setting.
	ID HotlinkProtectionID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value HotlinkProtectionValue `json:"value"`
	JSON  HotlinkProtectionJSON
}

type HotlinkProtectionJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into HotlinkProtection using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *HotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type HotlinkProtectionEditable bool

const (
	HotlinkProtectionEditableTrue  HotlinkProtectionEditable = true
	HotlinkProtectionEditableFalse HotlinkProtectionEditable = false
)

type HotlinkProtectionID string

const (
	HotlinkProtectionIDHotlinkProtection HotlinkProtectionID = "hotlink_protection"
)

type HotlinkProtectionValue string

const (
	HotlinkProtectionValueOn  HotlinkProtectionValue = "on"
	HotlinkProtectionValueOff HotlinkProtectionValue = "off"
)

type Http2 struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable Http2Editable `json:"editable"`
	// ID of the zone setting.
	ID Http2ID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value Http2Value `json:"value"`
	JSON  Http2JSON
}

type Http2JSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Http2 using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Http2) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Http2Editable bool

const (
	Http2EditableTrue  Http2Editable = true
	Http2EditableFalse Http2Editable = false
)

type Http2ID string

const (
	Http2IDHttp2 Http2ID = "http2"
)

type Http2Value string

const (
	Http2ValueOn  Http2Value = "on"
	Http2ValueOff Http2Value = "off"
)

type Http3 struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable Http3Editable `json:"editable"`
	// ID of the zone setting.
	ID Http3ID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP3 setting.
	Value Http3Value `json:"value"`
	JSON  Http3JSON
}

type Http3JSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Http3 using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Http3) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Http3Editable bool

const (
	Http3EditableTrue  Http3Editable = true
	Http3EditableFalse Http3Editable = false
)

type Http3ID string

const (
	Http3IDHttp3 Http3ID = "http3"
)

type Http3Value string

const (
	Http3ValueOn  Http3Value = "on"
	Http3ValueOff Http3Value = "off"
)

type ImageResizing struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ImageResizingEditable `json:"editable"`
	// ID of the zone setting.
	ID ImageResizingID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value ImageResizingValue `json:"value"`
	JSON  ImageResizingJSON
}

type ImageResizingJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ImageResizing using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ImageResizing) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ImageResizingEditable bool

const (
	ImageResizingEditableTrue  ImageResizingEditable = true
	ImageResizingEditableFalse ImageResizingEditable = false
)

type ImageResizingID string

const (
	ImageResizingIDImageResizing ImageResizingID = "image_resizing"
)

type ImageResizingValue string

const (
	ImageResizingValueOn   ImageResizingValue = "on"
	ImageResizingValueOff  ImageResizingValue = "off"
	ImageResizingValueOpen ImageResizingValue = "open"
)

type IpGeolocation struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable IpGeolocationEditable `json:"editable"`
	// ID of the zone setting.
	ID IpGeolocationID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value IpGeolocationValue `json:"value"`
	JSON  IpGeolocationJSON
}

type IpGeolocationJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into IpGeolocation using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *IpGeolocation) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type IpGeolocationEditable bool

const (
	IpGeolocationEditableTrue  IpGeolocationEditable = true
	IpGeolocationEditableFalse IpGeolocationEditable = false
)

type IpGeolocationID string

const (
	IpGeolocationIDIpGeolocation IpGeolocationID = "ip_geolocation"
)

type IpGeolocationValue string

const (
	IpGeolocationValueOn  IpGeolocationValue = "on"
	IpGeolocationValueOff IpGeolocationValue = "off"
)

type Ipv6 struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable Ipv6Editable `json:"editable"`
	// ID of the zone setting.
	ID Ipv6ID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value Ipv6Value `json:"value"`
	JSON  Ipv6JSON
}

type Ipv6JSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Ipv6 using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Ipv6) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Ipv6Editable bool

const (
	Ipv6EditableTrue  Ipv6Editable = true
	Ipv6EditableFalse Ipv6Editable = false
)

type Ipv6ID string

const (
	Ipv6IDIpv6 Ipv6ID = "ipv6"
)

type Ipv6Value string

const (
	Ipv6ValueOff Ipv6Value = "off"
	Ipv6ValueOn  Ipv6Value = "on"
)

type MaxUpload struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MaxUploadEditable `json:"editable"`
	// identifier of the zone setting.
	ID MaxUploadID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The size depends on the plan level of the
	// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
	Value MaxUploadValue `json:"value"`
	JSON  MaxUploadJSON
}

type MaxUploadJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MaxUpload using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *MaxUpload) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MaxUploadEditable bool

const (
	MaxUploadEditableTrue  MaxUploadEditable = true
	MaxUploadEditableFalse MaxUploadEditable = false
)

type MaxUploadID string

const (
	MaxUploadIDMaxUpload MaxUploadID = "max_upload"
)

type MaxUploadValue float64

const (
	MaxUploadValue_100 MaxUploadValue = 100
	MaxUploadValue_200 MaxUploadValue = 200
	MaxUploadValue_500 MaxUploadValue = 500
)

type MinTlsVersion struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MinTlsVersionEditable `json:"editable"`
	// ID of the zone setting.
	ID MinTlsVersionID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value MinTlsVersionValue `json:"value"`
	JSON  MinTlsVersionJSON
}

type MinTlsVersionJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MinTlsVersion using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MinTlsVersion) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MinTlsVersionEditable bool

const (
	MinTlsVersionEditableTrue  MinTlsVersionEditable = true
	MinTlsVersionEditableFalse MinTlsVersionEditable = false
)

type MinTlsVersionID string

const (
	MinTlsVersionIDMinTlsVersion MinTlsVersionID = "min_tls_version"
)

type MinTlsVersionValue string

const (
	MinTlsVersionValue_1_0 MinTlsVersionValue = "1.0"
	MinTlsVersionValue_1_1 MinTlsVersionValue = "1.1"
	MinTlsVersionValue_1_2 MinTlsVersionValue = "1.2"
	MinTlsVersionValue_1_3 MinTlsVersionValue = "1.3"
)

type Minify struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MinifyEditable `json:"editable"`
	// Zone setting identifier.
	ID MinifyID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value MinifyValue `json:"value"`
	JSON  MinifyJSON
}

type MinifyJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Minify using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Minify) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MinifyEditable bool

const (
	MinifyEditableTrue  MinifyEditable = true
	MinifyEditableFalse MinifyEditable = false
)

type MinifyID string

const (
	MinifyIDMinify MinifyID = "minify"
)

type MinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css MinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML MinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   MinifyValueJs `json:"js"`
	JSON MinifyValueJSON
}

type MinifyValueJSON struct {
	Css    pjson.Metadata
	HTML   pjson.Metadata
	Js     pjson.Metadata
	Raw    []byte
	Extras map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MinifyValue using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MinifyValue) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MinifyValueCss string

const (
	MinifyValueCssOn  MinifyValueCss = "on"
	MinifyValueCssOff MinifyValueCss = "off"
)

type MinifyValueHTML string

const (
	MinifyValueHTMLOn  MinifyValueHTML = "on"
	MinifyValueHTMLOff MinifyValueHTML = "off"
)

type MinifyValueJs string

const (
	MinifyValueJsOn  MinifyValueJs = "on"
	MinifyValueJsOff MinifyValueJs = "off"
)

type Mirage struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MirageEditable `json:"editable"`
	// ID of the zone setting.
	ID MirageID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value MirageValue `json:"value"`
	JSON  MirageJSON
}

type MirageJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Mirage using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Mirage) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MirageEditable bool

const (
	MirageEditableTrue  MirageEditable = true
	MirageEditableFalse MirageEditable = false
)

type MirageID string

const (
	MirageIDMirage MirageID = "mirage"
)

type MirageValue string

const (
	MirageValueOn  MirageValue = "on"
	MirageValueOff MirageValue = "off"
)

type MobileRedirect struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable MobileRedirectEditable `json:"editable"`
	// Identifier of the zone setting.
	ID MobileRedirectID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value MobileRedirectValue `json:"value"`
	JSON  MobileRedirectJSON
}

type MobileRedirectJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MobileRedirect using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MobileRedirectEditable bool

const (
	MobileRedirectEditableTrue  MobileRedirectEditable = true
	MobileRedirectEditableFalse MobileRedirectEditable = false
)

type MobileRedirectID string

const (
	MobileRedirectIDMobileRedirect MobileRedirectID = "mobile_redirect"
)

type MobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status MobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool `json:"strip_uri"`
	JSON     MobileRedirectValueJSON
}

type MobileRedirectValueJSON struct {
	MobileSubdomain pjson.Metadata
	Status          pjson.Metadata
	StripUri        pjson.Metadata
	Raw             []byte
	Extras          map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MobileRedirectValue using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MobileRedirectValueStatus string

const (
	MobileRedirectValueStatusOn  MobileRedirectValueStatus = "on"
	MobileRedirectValueStatusOff MobileRedirectValueStatus = "off"
)

type Nel struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable NelEditable `json:"editable"`
	// Zone setting identifier.
	ID NelID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value NelValue `json:"value"`
	JSON  NelJSON
}

type NelJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Nel using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Nel) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type NelEditable bool

const (
	NelEditableTrue  NelEditable = true
	NelEditableFalse NelEditable = false
)

type NelID string

const (
	NelIDNel NelID = "nel"
)

type NelValue struct {
	Enabled bool `json:"enabled"`
	JSON    NelValueJSON
}

type NelValueJSON struct {
	Enabled pjson.Metadata
	Raw     []byte
	Extras  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into NelValue using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *NelValue) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OpportunisticEncryption struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OpportunisticEncryptionEditable `json:"editable"`
	// ID of the zone setting.
	ID OpportunisticEncryptionID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value OpportunisticEncryptionValue `json:"value"`
	JSON  OpportunisticEncryptionJSON
}

type OpportunisticEncryptionJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OpportunisticEncryption using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *OpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OpportunisticEncryptionEditable bool

const (
	OpportunisticEncryptionEditableTrue  OpportunisticEncryptionEditable = true
	OpportunisticEncryptionEditableFalse OpportunisticEncryptionEditable = false
)

type OpportunisticEncryptionID string

const (
	OpportunisticEncryptionIDOpportunisticEncryption OpportunisticEncryptionID = "opportunistic_encryption"
)

type OpportunisticEncryptionValue string

const (
	OpportunisticEncryptionValueOn  OpportunisticEncryptionValue = "on"
	OpportunisticEncryptionValueOff OpportunisticEncryptionValue = "off"
)

type OpportunisticOnion struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OpportunisticOnionEditable `json:"editable"`
	// ID of the zone setting.
	ID OpportunisticOnionID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value OpportunisticOnionValue `json:"value"`
	JSON  OpportunisticOnionJSON
}

type OpportunisticOnionJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OpportunisticOnion using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *OpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OpportunisticOnionEditable bool

const (
	OpportunisticOnionEditableTrue  OpportunisticOnionEditable = true
	OpportunisticOnionEditableFalse OpportunisticOnionEditable = false
)

type OpportunisticOnionID string

const (
	OpportunisticOnionIDOpportunisticOnion OpportunisticOnionID = "opportunistic_onion"
)

type OpportunisticOnionValue string

const (
	OpportunisticOnionValueOn  OpportunisticOnionValue = "on"
	OpportunisticOnionValueOff OpportunisticOnionValue = "off"
)

type OrangeToOrange struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OrangeToOrangeEditable `json:"editable"`
	// ID of the zone setting.
	ID OrangeToOrangeID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value OrangeToOrangeValue `json:"value"`
	JSON  OrangeToOrangeJSON
}

type OrangeToOrangeJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OrangeToOrange using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *OrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OrangeToOrangeEditable bool

const (
	OrangeToOrangeEditableTrue  OrangeToOrangeEditable = true
	OrangeToOrangeEditableFalse OrangeToOrangeEditable = false
)

type OrangeToOrangeID string

const (
	OrangeToOrangeIDOrangeToOrange OrangeToOrangeID = "orange_to_orange"
)

type OrangeToOrangeValue string

const (
	OrangeToOrangeValueOn  OrangeToOrangeValue = "on"
	OrangeToOrangeValueOff OrangeToOrangeValue = "off"
)

type OriginErrorPagePassThru struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable OriginErrorPagePassThruEditable `json:"editable"`
	// ID of the zone setting.
	ID OriginErrorPagePassThruID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value OriginErrorPagePassThruValue `json:"value"`
	JSON  OriginErrorPagePassThruJSON
}

type OriginErrorPagePassThruJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OriginErrorPagePassThru using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *OriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OriginErrorPagePassThruEditable bool

const (
	OriginErrorPagePassThruEditableTrue  OriginErrorPagePassThruEditable = true
	OriginErrorPagePassThruEditableFalse OriginErrorPagePassThruEditable = false
)

type OriginErrorPagePassThruID string

const (
	OriginErrorPagePassThruIDOriginErrorPagePassThru OriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

type OriginErrorPagePassThruValue string

const (
	OriginErrorPagePassThruValueOn  OriginErrorPagePassThruValue = "on"
	OriginErrorPagePassThruValueOff OriginErrorPagePassThruValue = "off"
)

type OriginMaxHTTPVersion struct {
	// Identifier of the zone setting.
	ID OriginMaxHTTPVersionID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       OriginMaxHTTPVersionJSON
}

type OriginMaxHTTPVersionJSON struct {
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OriginMaxHTTPVersion using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *OriginMaxHTTPVersion) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OriginMaxHTTPVersionID string

const (
	OriginMaxHTTPVersionIDOriginMaxHTTPVersion OriginMaxHTTPVersionID = "origin_max_http_version"
)

type Polish struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PolishEditable `json:"editable"`
	// ID of the zone setting.
	ID PolishID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value PolishValue `json:"value"`
	JSON  PolishJSON
}

type PolishJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Polish using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Polish) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PolishEditable bool

const (
	PolishEditableTrue  PolishEditable = true
	PolishEditableFalse PolishEditable = false
)

type PolishID string

const (
	PolishIDPolish PolishID = "polish"
)

type PolishValue string

const (
	PolishValueOff      PolishValue = "off"
	PolishValueLossless PolishValue = "lossless"
	PolishValueLossy    PolishValue = "lossy"
)

type PrefetchPreload struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PrefetchPreloadEditable `json:"editable"`
	// ID of the zone setting.
	ID PrefetchPreloadID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value PrefetchPreloadValue `json:"value"`
	JSON  PrefetchPreloadJSON
}

type PrefetchPreloadJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PrefetchPreload using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PrefetchPreloadEditable bool

const (
	PrefetchPreloadEditableTrue  PrefetchPreloadEditable = true
	PrefetchPreloadEditableFalse PrefetchPreloadEditable = false
)

type PrefetchPreloadID string

const (
	PrefetchPreloadIDPrefetchPreload PrefetchPreloadID = "prefetch_preload"
)

type PrefetchPreloadValue string

const (
	PrefetchPreloadValueOn  PrefetchPreloadValue = "on"
	PrefetchPreloadValueOff PrefetchPreloadValue = "off"
)

type PrivacyPass struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PrivacyPassEditable `json:"editable"`
	// ID of the zone setting.
	ID PrivacyPassID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value PrivacyPassValue `json:"value"`
	JSON  PrivacyPassJSON
}

type PrivacyPassJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PrivacyPass using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PrivacyPass) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PrivacyPassEditable bool

const (
	PrivacyPassEditableTrue  PrivacyPassEditable = true
	PrivacyPassEditableFalse PrivacyPassEditable = false
)

type PrivacyPassID string

const (
	PrivacyPassIDPrivacyPass PrivacyPassID = "privacy_pass"
)

type PrivacyPassValue string

const (
	PrivacyPassValueOn  PrivacyPassValue = "on"
	PrivacyPassValueOff PrivacyPassValue = "off"
)

type ProxyReadTimeout struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ProxyReadTimeoutEditable `json:"editable"`
	// ID of the zone setting.
	ID ProxyReadTimeoutID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value float64 `json:"value"`
	JSON  ProxyReadTimeoutJSON
}

type ProxyReadTimeoutJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ProxyReadTimeout using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ProxyReadTimeoutEditable bool

const (
	ProxyReadTimeoutEditableTrue  ProxyReadTimeoutEditable = true
	ProxyReadTimeoutEditableFalse ProxyReadTimeoutEditable = false
)

type ProxyReadTimeoutID string

const (
	ProxyReadTimeoutIDProxyReadTimeout ProxyReadTimeoutID = "proxy_read_timeout"
)

type PseudoIpv4 struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable PseudoIpv4Editable `json:"editable"`
	// Value of the Pseudo IPv4 setting.
	ID PseudoIpv4ID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the Pseudo IPv4 setting.
	Value PseudoIpv4Value `json:"value"`
	JSON  PseudoIpv4JSON
}

type PseudoIpv4JSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PseudoIpv4 using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *PseudoIpv4) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PseudoIpv4Editable bool

const (
	PseudoIpv4EditableTrue  PseudoIpv4Editable = true
	PseudoIpv4EditableFalse PseudoIpv4Editable = false
)

type PseudoIpv4ID string

const (
	PseudoIpv4IDPseudoIpv4 PseudoIpv4ID = "pseudo_ipv4"
)

type PseudoIpv4Value string

const (
	PseudoIpv4ValueOff             PseudoIpv4Value = "off"
	PseudoIpv4ValueAddHeader       PseudoIpv4Value = "add_header"
	PseudoIpv4ValueOverwriteHeader PseudoIpv4Value = "overwrite_header"
)

type ResponseBuffering struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ResponseBufferingEditable `json:"editable"`
	// ID of the zone setting.
	ID ResponseBufferingID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ResponseBufferingValue `json:"value"`
	JSON  ResponseBufferingJSON
}

type ResponseBufferingJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ResponseBuffering using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ResponseBufferingEditable bool

const (
	ResponseBufferingEditableTrue  ResponseBufferingEditable = true
	ResponseBufferingEditableFalse ResponseBufferingEditable = false
)

type ResponseBufferingID string

const (
	ResponseBufferingIDResponseBuffering ResponseBufferingID = "response_buffering"
)

type ResponseBufferingValue string

const (
	ResponseBufferingValueOn  ResponseBufferingValue = "on"
	ResponseBufferingValueOff ResponseBufferingValue = "off"
)

type RocketLoader struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable RocketLoaderEditable `json:"editable"`
	// ID of the zone setting.
	ID RocketLoaderID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value RocketLoaderValue `json:"value"`
	JSON  RocketLoaderJSON
}

type RocketLoaderJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into RocketLoader using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *RocketLoader) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type RocketLoaderEditable bool

const (
	RocketLoaderEditableTrue  RocketLoaderEditable = true
	RocketLoaderEditableFalse RocketLoaderEditable = false
)

type RocketLoaderID string

const (
	RocketLoaderIDRocketLoader RocketLoaderID = "rocket_loader"
)

type RocketLoaderValue string

const (
	RocketLoaderValueOn  RocketLoaderValue = "on"
	RocketLoaderValueOff RocketLoaderValue = "off"
)

type SchemasAutomaticPlatformOptimization struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// ID of the zone setting.
	ID SchemasAutomaticPlatformOptimizationID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	Value      AutomaticPlatformOptimization `json:"value"`
	JSON       SchemasAutomaticPlatformOptimizationJSON
}

type SchemasAutomaticPlatformOptimizationJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// SchemasAutomaticPlatformOptimization using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *SchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SchemasAutomaticPlatformOptimizationEditable bool

const (
	SchemasAutomaticPlatformOptimizationEditableTrue  SchemasAutomaticPlatformOptimizationEditable = true
	SchemasAutomaticPlatformOptimizationEditableFalse SchemasAutomaticPlatformOptimizationEditable = false
)

type SchemasAutomaticPlatformOptimizationID string

const (
	SchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization SchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

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
	Hostnames []string `json:"hostnames,required"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress bool `json:"wordpress,required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin bool `json:"wp_plugin,required"`
	JSON     AutomaticPlatformOptimizationJSON
}

type AutomaticPlatformOptimizationJSON struct {
	CacheByDeviceType pjson.Metadata
	Cf                pjson.Metadata
	Enabled           pjson.Metadata
	Hostnames         pjson.Metadata
	Wordpress         pjson.Metadata
	WpPlugin          pjson.Metadata
	Raw               []byte
	Extras            map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AutomaticPlatformOptimization
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SecurityHeader struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SecurityHeaderEditable `json:"editable"`
	// ID of the zone's security header.
	ID SecurityHeaderID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time           `json:"modified_on,nullable" format:"date-time"`
	Value      SecurityHeaderValue `json:"value"`
	JSON       SecurityHeaderJSON
}

type SecurityHeaderJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SecurityHeader using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *SecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SecurityHeaderEditable bool

const (
	SecurityHeaderEditableTrue  SecurityHeaderEditable = true
	SecurityHeaderEditableFalse SecurityHeaderEditable = false
)

type SecurityHeaderID string

const (
	SecurityHeaderIDSecurityHeader SecurityHeaderID = "security_header"
)

type SecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity SecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    SecurityHeaderValueJSON
}

type SecurityHeaderValueJSON struct {
	StrictTransportSecurity pjson.Metadata
	Raw                     []byte
	Extras                  map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SecurityHeaderValue using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *SecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool `json:"nosniff"`
	JSON    SecurityHeaderValueStrictTransportSecurityJSON
}

type SecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           pjson.Metadata
	IncludeSubdomains pjson.Metadata
	MaxAge            pjson.Metadata
	Nosniff           pjson.Metadata
	Raw               []byte
	Extras            map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// SecurityHeaderValueStrictTransportSecurity using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *SecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SecurityLevel struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SecurityLevelEditable `json:"editable"`
	// ID of the zone setting.
	ID SecurityLevelID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value SecurityLevelValue `json:"value"`
	JSON  SecurityLevelJSON
}

type SecurityLevelJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SecurityLevel using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *SecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SecurityLevelEditable bool

const (
	SecurityLevelEditableTrue  SecurityLevelEditable = true
	SecurityLevelEditableFalse SecurityLevelEditable = false
)

type SecurityLevelID string

const (
	SecurityLevelIDSecurityLevel SecurityLevelID = "security_level"
)

type SecurityLevelValue string

const (
	SecurityLevelValueOff            SecurityLevelValue = "off"
	SecurityLevelValueEssentiallyOff SecurityLevelValue = "essentially_off"
	SecurityLevelValueLow            SecurityLevelValue = "low"
	SecurityLevelValueMedium         SecurityLevelValue = "medium"
	SecurityLevelValueHigh           SecurityLevelValue = "high"
	SecurityLevelValueUnderAttack    SecurityLevelValue = "under_attack"
)

type ServerSideExclude struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ServerSideExcludeEditable `json:"editable"`
	// ID of the zone setting.
	ID ServerSideExcludeID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ServerSideExcludeValue `json:"value"`
	JSON  ServerSideExcludeJSON
}

type ServerSideExcludeJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ServerSideExclude using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ServerSideExcludeEditable bool

const (
	ServerSideExcludeEditableTrue  ServerSideExcludeEditable = true
	ServerSideExcludeEditableFalse ServerSideExcludeEditable = false
)

type ServerSideExcludeID string

const (
	ServerSideExcludeIDServerSideExclude ServerSideExcludeID = "server_side_exclude"
)

type ServerSideExcludeValue string

const (
	ServerSideExcludeValueOn  ServerSideExcludeValue = "on"
	ServerSideExcludeValueOff ServerSideExcludeValue = "off"
)

type Sha1Support struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable Sha1SupportEditable `json:"editable"`
	// Zone setting identifier.
	ID Sha1SupportID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value Sha1SupportValue `json:"value"`
	JSON  Sha1SupportJSON
}

type Sha1SupportJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Sha1Support using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Sha1Support) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Sha1SupportEditable bool

const (
	Sha1SupportEditableTrue  Sha1SupportEditable = true
	Sha1SupportEditableFalse Sha1SupportEditable = false
)

type Sha1SupportID string

const (
	Sha1SupportIDSha1Support Sha1SupportID = "sha1_support"
)

type Sha1SupportValue string

const (
	Sha1SupportValueOff Sha1SupportValue = "off"
	Sha1SupportValueOn  Sha1SupportValue = "on"
)

type SortQueryStringForCache struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SortQueryStringForCacheEditable `json:"editable"`
	// ID of the zone setting.
	ID SortQueryStringForCacheID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value SortQueryStringForCacheValue `json:"value"`
	JSON  SortQueryStringForCacheJSON
}

type SortQueryStringForCacheJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SortQueryStringForCache using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *SortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SortQueryStringForCacheEditable bool

const (
	SortQueryStringForCacheEditableTrue  SortQueryStringForCacheEditable = true
	SortQueryStringForCacheEditableFalse SortQueryStringForCacheEditable = false
)

type SortQueryStringForCacheID string

const (
	SortQueryStringForCacheIDSortQueryStringForCache SortQueryStringForCacheID = "sort_query_string_for_cache"
)

type SortQueryStringForCacheValue string

const (
	SortQueryStringForCacheValueOn  SortQueryStringForCacheValue = "on"
	SortQueryStringForCacheValueOff SortQueryStringForCacheValue = "off"
)

type Ssl struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SslEditable `json:"editable"`
	// ID of the zone setting.
	ID SslID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value SslValue `json:"value"`
	JSON  SslJSON
}

type SslJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Ssl using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Ssl) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SslEditable bool

const (
	SslEditableTrue  SslEditable = true
	SslEditableFalse SslEditable = false
)

type SslID string

const (
	SslIDSsl SslID = "ssl"
)

type SslValue string

const (
	SslValueOff      SslValue = "off"
	SslValueFlexible SslValue = "flexible"
	SslValueFull     SslValue = "full"
	SslValueStrict   SslValue = "strict"
)

type SslRecommender struct {
	// ssl-recommender enrollment setting.
	Enabled bool `json:"enabled"`
	// Enrollment value for SSL/TLS Recommender.
	ID SslRecommenderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	JSON       SslRecommenderJSON
}

type SslRecommenderJSON struct {
	Enabled    pjson.Metadata
	ID         pjson.Metadata
	Editable   pjson.Metadata
	ModifiedOn pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SslRecommender using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *SslRecommender) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SslRecommenderID string

const (
	SslRecommenderIDSslRecommender SslRecommenderID = "ssl_recommender"
)

type Editable bool

const (
	EditableTrue  Editable = true
	EditableFalse Editable = false
)

type Tls_1_2Only struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable Tls_1_2OnlyEditable `json:"editable"`
	// Zone setting identifier.
	ID Tls_1_2OnlyID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value Tls_1_2OnlyValue `json:"value"`
	JSON  Tls_1_2OnlyJSON
}

type Tls_1_2OnlyJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Tls_1_2Only using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Tls_1_2Only) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Tls_1_2OnlyEditable bool

const (
	Tls_1_2OnlyEditableTrue  Tls_1_2OnlyEditable = true
	Tls_1_2OnlyEditableFalse Tls_1_2OnlyEditable = false
)

type Tls_1_2OnlyID string

const (
	Tls_1_2OnlyIDTls_1_2Only Tls_1_2OnlyID = "tls_1_2_only"
)

type Tls_1_2OnlyValue string

const (
	Tls_1_2OnlyValueOff Tls_1_2OnlyValue = "off"
	Tls_1_2OnlyValueOn  Tls_1_2OnlyValue = "on"
)

type Tls_1_3 struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable Tls_1_3Editable `json:"editable"`
	// ID of the zone setting.
	ID Tls_1_3ID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value Tls_1_3Value `json:"value"`
	JSON  Tls_1_3JSON
}

type Tls_1_3JSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Tls_1_3 using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Tls_1_3) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Tls_1_3Editable bool

const (
	Tls_1_3EditableTrue  Tls_1_3Editable = true
	Tls_1_3EditableFalse Tls_1_3Editable = false
)

type Tls_1_3ID string

const (
	Tls_1_3IDTls_1_3 Tls_1_3ID = "tls_1_3"
)

type Tls_1_3Value string

const (
	Tls_1_3ValueOn  Tls_1_3Value = "on"
	Tls_1_3ValueOff Tls_1_3Value = "off"
	Tls_1_3ValueZrt Tls_1_3Value = "zrt"
)

type TlsClientAuth struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TlsClientAuthEditable `json:"editable"`
	// ID of the zone setting.
	ID TlsClientAuthID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value TlsClientAuthValue `json:"value"`
	JSON  TlsClientAuthJSON
}

type TlsClientAuthJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TlsClientAuth using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TlsClientAuth) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TlsClientAuthEditable bool

const (
	TlsClientAuthEditableTrue  TlsClientAuthEditable = true
	TlsClientAuthEditableFalse TlsClientAuthEditable = false
)

type TlsClientAuthID string

const (
	TlsClientAuthIDTlsClientAuth TlsClientAuthID = "tls_client_auth"
)

type TlsClientAuthValue string

const (
	TlsClientAuthValueOn  TlsClientAuthValue = "on"
	TlsClientAuthValueOff TlsClientAuthValue = "off"
)

type TrueClientIpHeader struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable TrueClientIpHeaderEditable `json:"editable"`
	// ID of the zone setting.
	ID TrueClientIpHeaderID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value TrueClientIpHeaderValue `json:"value"`
	JSON  TrueClientIpHeaderJSON
}

type TrueClientIpHeaderJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into TrueClientIpHeader using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *TrueClientIpHeader) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TrueClientIpHeaderEditable bool

const (
	TrueClientIpHeaderEditableTrue  TrueClientIpHeaderEditable = true
	TrueClientIpHeaderEditableFalse TrueClientIpHeaderEditable = false
)

type TrueClientIpHeaderID string

const (
	TrueClientIpHeaderIDTrueClientIpHeader TrueClientIpHeaderID = "true_client_ip_header"
)

type TrueClientIpHeaderValue string

const (
	TrueClientIpHeaderValueOn  TrueClientIpHeaderValue = "on"
	TrueClientIpHeaderValueOff TrueClientIpHeaderValue = "off"
)

type Waf struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WafEditable `json:"editable"`
	// ID of the zone setting.
	ID WafID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value WafValue `json:"value"`
	JSON  WafJSON
}

type WafJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Waf using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Waf) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WafEditable bool

const (
	WafEditableTrue  WafEditable = true
	WafEditableFalse WafEditable = false
)

type WafID string

const (
	WafIDWaf WafID = "waf"
)

type WafValue string

const (
	WafValueOn  WafValue = "on"
	WafValueOff WafValue = "off"
)

type Webp struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WebpEditable `json:"editable"`
	// ID of the zone setting.
	ID WebpID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value WebpValue `json:"value"`
	JSON  WebpJSON
}

type WebpJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Webp using the internal pjson
// library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Webp) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WebpEditable bool

const (
	WebpEditableTrue  WebpEditable = true
	WebpEditableFalse WebpEditable = false
)

type WebpID string

const (
	WebpIDWebp WebpID = "webp"
)

type WebpValue string

const (
	WebpValueOff WebpValue = "off"
	WebpValueOn  WebpValue = "on"
)

type Websockets struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable WebsocketsEditable `json:"editable"`
	// ID of the zone setting.
	ID WebsocketsID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value WebsocketsValue `json:"value"`
	JSON  WebsocketsJSON
}

type WebsocketsJSON struct {
	Editable   pjson.Metadata
	ID         pjson.Metadata
	ModifiedOn pjson.Metadata
	Value      pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Websockets using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Websockets) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WebsocketsEditable bool

const (
	WebsocketsEditableTrue  WebsocketsEditable = true
	WebsocketsEditableFalse WebsocketsEditable = false
)

type WebsocketsID string

const (
	WebsocketsIDWebsockets WebsocketsID = "websockets"
)

type WebsocketsValue string

const (
	WebsocketsValueOff WebsocketsValue = "off"
	WebsocketsValueOn  WebsocketsValue = "on"
)
