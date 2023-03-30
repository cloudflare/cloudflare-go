package requests

import (
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type ZeroRtt struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable fields.Field[ZeroRttEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[ZeroRttID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the 0-RTT setting.
	Value fields.Field[ZeroRttValue] `json:"value"`
}

// MarshalJSON serializes ZeroRtt into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *ZeroRtt) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ZeroRtt) String() (result string) {
	return fmt.Sprintf("&ZeroRtt{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool]                           `json:"success"`
	Result  fields.Field[[]ZoneSettingsCollectionResult] `json:"result"`
}

// MarshalJSON serializes ZoneSettingsCollection into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ZoneSettingsCollection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ZoneSettingsCollection) String() (result string) {
	return fmt.Sprintf("&ZoneSettingsCollection{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, core.Fmt(r.Result))
}

// 🚧 Undiscriminated unions are still being implemented.
type ZoneSettingsCollectionResult struct{}

func (r *ZoneSettingsCollectionResult) MarshalJSON() ([]byte, error) { return nil, nil }

type AdvancedDdos struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable fields.Field[AdvancedDdosEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[AdvancedDdosID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value fields.Field[AdvancedDdosValue] `json:"value"`
}

// MarshalJSON serializes AdvancedDdos into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AdvancedDdos) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AdvancedDdos) String() (result string) {
	return fmt.Sprintf("&AdvancedDdos{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[AlwaysOnlineEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[AlwaysOnlineID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[AlwaysOnlineValue] `json:"value"`
}

// MarshalJSON serializes AlwaysOnline into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AlwaysOnline) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AlwaysOnline) String() (result string) {
	return fmt.Sprintf("&AlwaysOnline{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[AlwaysUseHTTPsEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[AlwaysUseHTTPsID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[AlwaysUseHTTPsValue] `json:"value"`
}

// MarshalJSON serializes AlwaysUseHTTPs into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AlwaysUseHTTPs) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AlwaysUseHTTPs) String() (result string) {
	return fmt.Sprintf("&AlwaysUseHTTPs{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[AutomaticHTTPsRewritesEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[AutomaticHTTPsRewritesID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value fields.Field[AutomaticHTTPsRewritesValue] `json:"value"`
}

// MarshalJSON serializes AutomaticHTTPsRewrites into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AutomaticHTTPsRewrites) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AutomaticHTTPsRewrites) String() (result string) {
	return fmt.Sprintf("&AutomaticHTTPsRewrites{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[BrotliEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[BrotliID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[BrotliValue] `json:"value"`
}

// MarshalJSON serializes Brotli into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Brotli) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Brotli) String() (result string) {
	return fmt.Sprintf("&Brotli{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[BrowserCacheTtlEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[BrowserCacheTtlID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value fields.Field[BrowserCacheTtlValue] `json:"value"`
}

// MarshalJSON serializes BrowserCacheTtl into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *BrowserCacheTtl) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrowserCacheTtl) String() (result string) {
	return fmt.Sprintf("&BrowserCacheTtl{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[BrowserCheckEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[BrowserCheckID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[BrowserCheckValue] `json:"value"`
}

// MarshalJSON serializes BrowserCheck into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *BrowserCheck) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrowserCheck) String() (result string) {
	return fmt.Sprintf("&BrowserCheck{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[CacheLevelEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[CacheLevelID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[CacheLevelValue] `json:"value"`
}

// MarshalJSON serializes CacheLevel into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CacheLevel) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CacheLevel) String() (result string) {
	return fmt.Sprintf("&CacheLevel{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[ChallengeTtlEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[ChallengeTtlID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[ChallengeTtlValue] `json:"value"`
}

// MarshalJSON serializes ChallengeTtl into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ChallengeTtl) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ChallengeTtl) String() (result string) {
	return fmt.Sprintf("&ChallengeTtl{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[CiphersEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[CiphersID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[[]string] `json:"value"`
}

// MarshalJSON serializes Ciphers into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Ciphers) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Ciphers) String() (result string) {
	return fmt.Sprintf("&Ciphers{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, core.Fmt(r.Value))
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
	Editable fields.Field[CnameFlatteningEditable] `json:"editable"`
	// How to flatten the cname destination.
	ID fields.Field[CnameFlatteningID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the cname flattening setting.
	Value fields.Field[CnameFlatteningValue] `json:"value"`
}

// MarshalJSON serializes CnameFlattening into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *CnameFlattening) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CnameFlattening) String() (result string) {
	return fmt.Sprintf("&CnameFlattening{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[DevelopmentModeEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[DevelopmentModeID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[DevelopmentModeValue] `json:"value"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining fields.Field[float64] `json:"time_remaining"`
}

// MarshalJSON serializes DevelopmentMode into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DevelopmentMode) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DevelopmentMode) String() (result string) {
	return fmt.Sprintf("&DevelopmentMode{Editable:%s ID:%s ModifiedOn:%s Value:%s TimeRemaining:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value, r.TimeRemaining)
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
	Editable fields.Field[EarlyHintsEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[EarlyHintsID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[EarlyHintsValue] `json:"value"`
}

// MarshalJSON serializes EarlyHints into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EarlyHints) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EarlyHints) String() (result string) {
	return fmt.Sprintf("&EarlyHints{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[EdgeCacheTtlEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[EdgeCacheTtlID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value fields.Field[EdgeCacheTtlValue] `json:"value"`
}

// MarshalJSON serializes EdgeCacheTtl into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EdgeCacheTtl) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EdgeCacheTtl) String() (result string) {
	return fmt.Sprintf("&EdgeCacheTtl{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[EmailObfuscationEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[EmailObfuscationID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[EmailObfuscationValue] `json:"value"`
}

// MarshalJSON serializes EmailObfuscation into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *EmailObfuscation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r EmailObfuscation) String() (result string) {
	return fmt.Sprintf("&EmailObfuscation{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[H2PrioritizationEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[H2PrioritizationID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[H2PrioritizationValue] `json:"value"`
}

// MarshalJSON serializes H2Prioritization into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *H2Prioritization) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r H2Prioritization) String() (result string) {
	return fmt.Sprintf("&H2Prioritization{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[HotlinkProtectionEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[HotlinkProtectionID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[HotlinkProtectionValue] `json:"value"`
}

// MarshalJSON serializes HotlinkProtection into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *HotlinkProtection) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r HotlinkProtection) String() (result string) {
	return fmt.Sprintf("&HotlinkProtection{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[Http2Editable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[Http2ID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value fields.Field[Http2Value] `json:"value"`
}

// MarshalJSON serializes Http2 into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Http2) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Http2) String() (result string) {
	return fmt.Sprintf("&Http2{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[Http3Editable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[Http3ID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP3 setting.
	Value fields.Field[Http3Value] `json:"value"`
}

// MarshalJSON serializes Http3 into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Http3) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Http3) String() (result string) {
	return fmt.Sprintf("&Http3{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[ImageResizingEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[ImageResizingID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value fields.Field[ImageResizingValue] `json:"value"`
}

// MarshalJSON serializes ImageResizing into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ImageResizing) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ImageResizing) String() (result string) {
	return fmt.Sprintf("&ImageResizing{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[IpGeolocationEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[IpGeolocationID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[IpGeolocationValue] `json:"value"`
}

// MarshalJSON serializes IpGeolocation into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *IpGeolocation) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r IpGeolocation) String() (result string) {
	return fmt.Sprintf("&IpGeolocation{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[Ipv6Editable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[Ipv6ID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[Ipv6Value] `json:"value"`
}

// MarshalJSON serializes Ipv6 into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Ipv6) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Ipv6) String() (result string) {
	return fmt.Sprintf("&Ipv6{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[MaxUploadEditable] `json:"editable"`
	// identifier of the zone setting.
	ID fields.Field[MaxUploadID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The size depends on the plan level of the
	// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
	Value fields.Field[MaxUploadValue] `json:"value"`
}

// MarshalJSON serializes MaxUpload into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *MaxUpload) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MaxUpload) String() (result string) {
	return fmt.Sprintf("&MaxUpload{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[MinTlsVersionEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[MinTlsVersionID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[MinTlsVersionValue] `json:"value"`
}

// MarshalJSON serializes MinTlsVersion into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *MinTlsVersion) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MinTlsVersion) String() (result string) {
	return fmt.Sprintf("&MinTlsVersion{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[MinifyEditable] `json:"editable"`
	// Zone setting identifier.
	ID fields.Field[MinifyID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[MinifyValue] `json:"value"`
}

// MarshalJSON serializes Minify into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Minify) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Minify) String() (result string) {
	return fmt.Sprintf("&Minify{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Css fields.Field[MinifyValueCss] `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML fields.Field[MinifyValueHTML] `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js fields.Field[MinifyValueJs] `json:"js"`
}

// MarshalJSON serializes MinifyValue into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *MinifyValue) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MinifyValue) String() (result string) {
	return fmt.Sprintf("&MinifyValue{Css:%s HTML:%s Js:%s}", r.Css, r.HTML, r.Js)
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
	Editable fields.Field[MirageEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[MirageID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[MirageValue] `json:"value"`
}

// MarshalJSON serializes Mirage into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Mirage) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Mirage) String() (result string) {
	return fmt.Sprintf("&Mirage{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[MobileRedirectEditable] `json:"editable"`
	// Identifier of the zone setting.
	ID fields.Field[MobileRedirectID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[MobileRedirectValue] `json:"value"`
}

// MarshalJSON serializes MobileRedirect into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *MobileRedirect) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MobileRedirect) String() (result string) {
	return fmt.Sprintf("&MobileRedirect{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	MobileSubdomain fields.Field[string] `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status fields.Field[MobileRedirectValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri fields.Field[bool] `json:"strip_uri"`
}

// MarshalJSON serializes MobileRedirectValue into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *MobileRedirectValue) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MobileRedirectValue) String() (result string) {
	return fmt.Sprintf("&MobileRedirectValue{MobileSubdomain:%s Status:%s StripUri:%s}", r.MobileSubdomain, r.Status, r.StripUri)
}

type MobileRedirectValueStatus string

const (
	MobileRedirectValueStatusOn  MobileRedirectValueStatus = "on"
	MobileRedirectValueStatusOff MobileRedirectValueStatus = "off"
)

type Nel struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable fields.Field[NelEditable] `json:"editable"`
	// Zone setting identifier.
	ID fields.Field[NelID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[NelValue] `json:"value"`
}

// MarshalJSON serializes Nel into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Nel) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Nel) String() (result string) {
	return fmt.Sprintf("&Nel{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Enabled fields.Field[bool] `json:"enabled"`
}

// MarshalJSON serializes NelValue into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *NelValue) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r NelValue) String() (result string) {
	return fmt.Sprintf("&NelValue{Enabled:%s}", r.Enabled)
}

type OpportunisticEncryption struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable fields.Field[OpportunisticEncryptionEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[OpportunisticEncryptionID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value fields.Field[OpportunisticEncryptionValue] `json:"value"`
}

// MarshalJSON serializes OpportunisticEncryption into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *OpportunisticEncryption) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OpportunisticEncryption) String() (result string) {
	return fmt.Sprintf("&OpportunisticEncryption{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[OpportunisticOnionEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[OpportunisticOnionID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value fields.Field[OpportunisticOnionValue] `json:"value"`
}

// MarshalJSON serializes OpportunisticOnion into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *OpportunisticOnion) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OpportunisticOnion) String() (result string) {
	return fmt.Sprintf("&OpportunisticOnion{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[OrangeToOrangeEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[OrangeToOrangeID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[OrangeToOrangeValue] `json:"value"`
}

// MarshalJSON serializes OrangeToOrange into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *OrangeToOrange) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OrangeToOrange) String() (result string) {
	return fmt.Sprintf("&OrangeToOrange{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[OriginErrorPagePassThruEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[OriginErrorPagePassThruID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[OriginErrorPagePassThruValue] `json:"value"`
}

// MarshalJSON serializes OriginErrorPagePassThru into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *OriginErrorPagePassThru) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OriginErrorPagePassThru) String() (result string) {
	return fmt.Sprintf("&OriginErrorPagePassThru{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	ID fields.Field[OriginMaxHTTPVersionID] `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
}

// MarshalJSON serializes OriginMaxHTTPVersion into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *OriginMaxHTTPVersion) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OriginMaxHTTPVersion) String() (result string) {
	return fmt.Sprintf("&OriginMaxHTTPVersion{ID:%s ModifiedOn:%s}", r.ID, r.ModifiedOn)
}

type OriginMaxHTTPVersionID string

const (
	OriginMaxHTTPVersionIDOriginMaxHTTPVersion OriginMaxHTTPVersionID = "origin_max_http_version"
)

type Polish struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable fields.Field[PolishEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[PolishID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[PolishValue] `json:"value"`
}

// MarshalJSON serializes Polish into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Polish) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Polish) String() (result string) {
	return fmt.Sprintf("&Polish{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[PrefetchPreloadEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[PrefetchPreloadID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[PrefetchPreloadValue] `json:"value"`
}

// MarshalJSON serializes PrefetchPreload into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *PrefetchPreload) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PrefetchPreload) String() (result string) {
	return fmt.Sprintf("&PrefetchPreload{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[PrivacyPassEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[PrivacyPassID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[PrivacyPassValue] `json:"value"`
}

// MarshalJSON serializes PrivacyPass into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *PrivacyPass) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PrivacyPass) String() (result string) {
	return fmt.Sprintf("&PrivacyPass{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[ProxyReadTimeoutEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[ProxyReadTimeoutID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value fields.Field[float64] `json:"value"`
}

// MarshalJSON serializes ProxyReadTimeout into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ProxyReadTimeout) String() (result string) {
	return fmt.Sprintf("&ProxyReadTimeout{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[PseudoIpv4Editable] `json:"editable"`
	// Value of the Pseudo IPv4 setting.
	ID fields.Field[PseudoIpv4ID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the Pseudo IPv4 setting.
	Value fields.Field[PseudoIpv4Value] `json:"value"`
}

// MarshalJSON serializes PseudoIpv4 into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *PseudoIpv4) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PseudoIpv4) String() (result string) {
	return fmt.Sprintf("&PseudoIpv4{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[ResponseBufferingEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[ResponseBufferingID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[ResponseBufferingValue] `json:"value"`
}

// MarshalJSON serializes ResponseBuffering into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ResponseBuffering) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ResponseBuffering) String() (result string) {
	return fmt.Sprintf("&ResponseBuffering{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[RocketLoaderEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[RocketLoaderID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[RocketLoaderValue] `json:"value"`
}

// MarshalJSON serializes RocketLoader into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *RocketLoader) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r RocketLoader) String() (result string) {
	return fmt.Sprintf("&RocketLoader{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[SchemasAutomaticPlatformOptimizationEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[SchemasAutomaticPlatformOptimizationID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time]                     `json:"modified_on,nullable" format:"date-time"`
	Value      fields.Field[AutomaticPlatformOptimization] `json:"value"`
}

// MarshalJSON serializes SchemasAutomaticPlatformOptimization into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *SchemasAutomaticPlatformOptimization) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SchemasAutomaticPlatformOptimization) String() (result string) {
	return fmt.Sprintf("&SchemasAutomaticPlatformOptimization{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	CacheByDeviceType fields.Field[bool] `json:"cache_by_device_type,required"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf fields.Field[bool] `json:"cf,required"`
	// Indicates whether or not Automatic Platform Optimization is enabled.
	Enabled fields.Field[bool] `json:"enabled,required"`
	// An array of hostnames where Automatic Platform Optimization for WordPress is
	// activated.
	Hostnames fields.Field[[]string] `json:"hostnames,required"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress fields.Field[bool] `json:"wordpress,required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin fields.Field[bool] `json:"wp_plugin,required"`
}

// MarshalJSON serializes AutomaticPlatformOptimization into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *AutomaticPlatformOptimization) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AutomaticPlatformOptimization) String() (result string) {
	return fmt.Sprintf("&AutomaticPlatformOptimization{CacheByDeviceType:%s Cf:%s Enabled:%s Hostnames:%s Wordpress:%s WpPlugin:%s}", r.CacheByDeviceType, r.Cf, r.Enabled, core.Fmt(r.Hostnames), r.Wordpress, r.WpPlugin)
}

type SecurityHeader struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable fields.Field[SecurityHeaderEditable] `json:"editable"`
	// ID of the zone's security header.
	ID fields.Field[SecurityHeaderID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time]           `json:"modified_on,nullable" format:"date-time"`
	Value      fields.Field[SecurityHeaderValue] `json:"value"`
}

// MarshalJSON serializes SecurityHeader into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *SecurityHeader) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SecurityHeader) String() (result string) {
	return fmt.Sprintf("&SecurityHeader{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	StrictTransportSecurity fields.Field[SecurityHeaderValueStrictTransportSecurity] `json:"strict_transport_security"`
}

// MarshalJSON serializes SecurityHeaderValue into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SecurityHeaderValue) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SecurityHeaderValue) String() (result string) {
	return fmt.Sprintf("&SecurityHeaderValue{StrictTransportSecurity:%s}", r.StrictTransportSecurity)
}

type SecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled fields.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains fields.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge fields.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff fields.Field[bool] `json:"nosniff"`
}

// MarshalJSON serializes SecurityHeaderValueStrictTransportSecurity into an array
// of bytes using the gjson library. Members of the `jsonFields` field are
// serialized into the top-level, and will overwrite known members of the same
// name.
func (r *SecurityHeaderValueStrictTransportSecurity) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SecurityHeaderValueStrictTransportSecurity) String() (result string) {
	return fmt.Sprintf("&SecurityHeaderValueStrictTransportSecurity{Enabled:%s IncludeSubdomains:%s MaxAge:%s Nosniff:%s}", r.Enabled, r.IncludeSubdomains, r.MaxAge, r.Nosniff)
}

type SecurityLevel struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable fields.Field[SecurityLevelEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[SecurityLevelID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[SecurityLevelValue] `json:"value"`
}

// MarshalJSON serializes SecurityLevel into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *SecurityLevel) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SecurityLevel) String() (result string) {
	return fmt.Sprintf("&SecurityLevel{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[ServerSideExcludeEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[ServerSideExcludeID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[ServerSideExcludeValue] `json:"value"`
}

// MarshalJSON serializes ServerSideExclude into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ServerSideExclude) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ServerSideExclude) String() (result string) {
	return fmt.Sprintf("&ServerSideExclude{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[Sha1SupportEditable] `json:"editable"`
	// Zone setting identifier.
	ID fields.Field[Sha1SupportID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[Sha1SupportValue] `json:"value"`
}

// MarshalJSON serializes Sha1Support into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Sha1Support) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Sha1Support) String() (result string) {
	return fmt.Sprintf("&Sha1Support{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[SortQueryStringForCacheEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[SortQueryStringForCacheID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[SortQueryStringForCacheValue] `json:"value"`
}

// MarshalJSON serializes SortQueryStringForCache into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *SortQueryStringForCache) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SortQueryStringForCache) String() (result string) {
	return fmt.Sprintf("&SortQueryStringForCache{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[SslEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[SslID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value fields.Field[SslValue] `json:"value"`
}

// MarshalJSON serializes Ssl into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Ssl) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Ssl) String() (result string) {
	return fmt.Sprintf("&Ssl{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Enabled fields.Field[bool] `json:"enabled"`
	// Enrollment value for SSL/TLS Recommender.
	ID fields.Field[SslRecommenderID] `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable fields.Field[Editable] `json:"editable"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
}

// MarshalJSON serializes SslRecommender into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *SslRecommender) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SslRecommender) String() (result string) {
	return fmt.Sprintf("&SslRecommender{Enabled:%s ID:%s Editable:%s ModifiedOn:%s}", r.Enabled, r.ID, r.Editable, r.ModifiedOn)
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
	Editable fields.Field[Tls_1_2OnlyEditable] `json:"editable"`
	// Zone setting identifier.
	ID fields.Field[Tls_1_2OnlyID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[Tls_1_2OnlyValue] `json:"value"`
}

// MarshalJSON serializes Tls_1_2Only into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Tls_1_2Only) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Tls_1_2Only) String() (result string) {
	return fmt.Sprintf("&Tls_1_2Only{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[Tls_1_3Editable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[Tls_1_3ID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value fields.Field[Tls_1_3Value] `json:"value"`
}

// MarshalJSON serializes Tls_1_3 into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Tls_1_3) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Tls_1_3) String() (result string) {
	return fmt.Sprintf("&Tls_1_3{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[TlsClientAuthEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[TlsClientAuthID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value fields.Field[TlsClientAuthValue] `json:"value"`
}

// MarshalJSON serializes TlsClientAuth into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *TlsClientAuth) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TlsClientAuth) String() (result string) {
	return fmt.Sprintf("&TlsClientAuth{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[TrueClientIpHeaderEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[TrueClientIpHeaderID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[TrueClientIpHeaderValue] `json:"value"`
}

// MarshalJSON serializes TrueClientIpHeader into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *TrueClientIpHeader) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r TrueClientIpHeader) String() (result string) {
	return fmt.Sprintf("&TrueClientIpHeader{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[WafEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[WafID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[WafValue] `json:"value"`
}

// MarshalJSON serializes Waf into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Waf) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Waf) String() (result string) {
	return fmt.Sprintf("&Waf{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[WebpEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[WebpID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[WebpValue] `json:"value"`
}

// MarshalJSON serializes Webp into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Webp) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Webp) String() (result string) {
	return fmt.Sprintf("&Webp{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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
	Editable fields.Field[WebsocketsEditable] `json:"editable"`
	// ID of the zone setting.
	ID fields.Field[WebsocketsID] `json:"id"`
	// last time this setting was modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value fields.Field[WebsocketsValue] `json:"value"`
}

// MarshalJSON serializes Websockets into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Websockets) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Websockets) String() (result string) {
	return fmt.Sprintf("&Websockets{Editable:%s ID:%s ModifiedOn:%s Value:%s}", r.Editable, r.ID, r.ModifiedOn, r.Value)
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

type SettingEditParams struct {
	// One or more zone setting objects. Must contain an ID and a value.
	Items fields.Field[[]Setting] `json:"items,required"`
}

// MarshalJSON serializes SettingEditParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *SettingEditParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SettingEditParams) String() (result string) {
	return fmt.Sprintf("&SettingEditParams{Items:%s}", core.Fmt(r.Items))
}

// 🚧 Undiscriminated unions are still being implemented.
type Setting struct{}

func (r *Setting) MarshalJSON() ([]byte, error) { return nil, nil }
