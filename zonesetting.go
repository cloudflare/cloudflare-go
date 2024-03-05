// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSettingService] method
// instead.
type ZoneSettingService struct {
	Options                       []option.RequestOption
	ZeroRTT                       *ZoneSettingZeroRTTService
	AdvancedDDOS                  *ZoneSettingAdvancedDDOSService
	AlwaysOnline                  *ZoneSettingAlwaysOnlineService
	AlwaysUseHTTPS                *ZoneSettingAlwaysUseHTTPSService
	AutomaticHTTPSRewrites        *ZoneSettingAutomaticHTTPSRewriteService
	AutomaticPlatformOptimization *ZoneSettingAutomaticPlatformOptimizationService
	Brotli                        *ZoneSettingBrotliService
	BrowserCacheTTL               *ZoneSettingBrowserCacheTTLService
	BrowserCheck                  *ZoneSettingBrowserCheckService
	CacheLevel                    *ZoneSettingCacheLevelService
	ChallengeTTL                  *ZoneSettingChallengeTTLService
	Ciphers                       *ZoneSettingCipherService
	DevelopmentMode               *ZoneSettingDevelopmentModeService
	EarlyHints                    *ZoneSettingEarlyHintService
	EmailObfuscation              *ZoneSettingEmailObfuscationService
	H2Prioritization              *ZoneSettingH2PrioritizationService
	HotlinkProtection             *ZoneSettingHotlinkProtectionService
	HTTP2                         *ZoneSettingHTTP2Service
	HTTP3                         *ZoneSettingHTTP3Service
	ImageResizing                 *ZoneSettingImageResizingService
	IPGeolocation                 *ZoneSettingIPGeolocationService
	IPV6                          *ZoneSettingIPV6Service
	MinTLSVersion                 *ZoneSettingMinTLSVersionService
	Minify                        *ZoneSettingMinifyService
	Mirage                        *ZoneSettingMirageService
	MobileRedirect                *ZoneSettingMobileRedirectService
	NEL                           *ZoneSettingNELService
	OpportunisticEncryption       *ZoneSettingOpportunisticEncryptionService
	OpportunisticOnion            *ZoneSettingOpportunisticOnionService
	OrangeToOrange                *ZoneSettingOrangeToOrangeService
	OriginErrorPagePassThru       *ZoneSettingOriginErrorPagePassThruService
	OriginMaxHTTPVersion          *ZoneSettingOriginMaxHTTPVersionService
	Polish                        *ZoneSettingPolishService
	PrefetchPreload               *ZoneSettingPrefetchPreloadService
	ProxyReadTimeout              *ZoneSettingProxyReadTimeoutService
	PseudoIPV4                    *ZoneSettingPseudoIPV4Service
	ResponseBuffering             *ZoneSettingResponseBufferingService
	RocketLoader                  *ZoneSettingRocketLoaderService
	SecurityHeaders               *ZoneSettingSecurityHeaderService
	SecurityLevel                 *ZoneSettingSecurityLevelService
	ServerSideExcludes            *ZoneSettingServerSideExcludeService
	SortQueryStringForCache       *ZoneSettingSortQueryStringForCacheService
	SSL                           *ZoneSettingSSLService
	SSLRecommender                *ZoneSettingSSLRecommenderService
	TLS1_3                        *ZoneSettingTLS1_3Service
	TLSClientAuth                 *ZoneSettingTLSClientAuthService
	TrueClientIPHeader            *ZoneSettingTrueClientIPHeaderService
	WAF                           *ZoneSettingWAFService
	Webp                          *ZoneSettingWebpService
	Websocket                     *ZoneSettingWebsocketService
	FontSettings                  *ZoneSettingFontSettingService
}

// NewZoneSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingService(opts ...option.RequestOption) (r *ZoneSettingService) {
	r = &ZoneSettingService{}
	r.Options = opts
	r.ZeroRTT = NewZoneSettingZeroRTTService(opts...)
	r.AdvancedDDOS = NewZoneSettingAdvancedDDOSService(opts...)
	r.AlwaysOnline = NewZoneSettingAlwaysOnlineService(opts...)
	r.AlwaysUseHTTPS = NewZoneSettingAlwaysUseHTTPSService(opts...)
	r.AutomaticHTTPSRewrites = NewZoneSettingAutomaticHTTPSRewriteService(opts...)
	r.AutomaticPlatformOptimization = NewZoneSettingAutomaticPlatformOptimizationService(opts...)
	r.Brotli = NewZoneSettingBrotliService(opts...)
	r.BrowserCacheTTL = NewZoneSettingBrowserCacheTTLService(opts...)
	r.BrowserCheck = NewZoneSettingBrowserCheckService(opts...)
	r.CacheLevel = NewZoneSettingCacheLevelService(opts...)
	r.ChallengeTTL = NewZoneSettingChallengeTTLService(opts...)
	r.Ciphers = NewZoneSettingCipherService(opts...)
	r.DevelopmentMode = NewZoneSettingDevelopmentModeService(opts...)
	r.EarlyHints = NewZoneSettingEarlyHintService(opts...)
	r.EmailObfuscation = NewZoneSettingEmailObfuscationService(opts...)
	r.H2Prioritization = NewZoneSettingH2PrioritizationService(opts...)
	r.HotlinkProtection = NewZoneSettingHotlinkProtectionService(opts...)
	r.HTTP2 = NewZoneSettingHTTP2Service(opts...)
	r.HTTP3 = NewZoneSettingHTTP3Service(opts...)
	r.ImageResizing = NewZoneSettingImageResizingService(opts...)
	r.IPGeolocation = NewZoneSettingIPGeolocationService(opts...)
	r.IPV6 = NewZoneSettingIPV6Service(opts...)
	r.MinTLSVersion = NewZoneSettingMinTLSVersionService(opts...)
	r.Minify = NewZoneSettingMinifyService(opts...)
	r.Mirage = NewZoneSettingMirageService(opts...)
	r.MobileRedirect = NewZoneSettingMobileRedirectService(opts...)
	r.NEL = NewZoneSettingNELService(opts...)
	r.OpportunisticEncryption = NewZoneSettingOpportunisticEncryptionService(opts...)
	r.OpportunisticOnion = NewZoneSettingOpportunisticOnionService(opts...)
	r.OrangeToOrange = NewZoneSettingOrangeToOrangeService(opts...)
	r.OriginErrorPagePassThru = NewZoneSettingOriginErrorPagePassThruService(opts...)
	r.OriginMaxHTTPVersion = NewZoneSettingOriginMaxHTTPVersionService(opts...)
	r.Polish = NewZoneSettingPolishService(opts...)
	r.PrefetchPreload = NewZoneSettingPrefetchPreloadService(opts...)
	r.ProxyReadTimeout = NewZoneSettingProxyReadTimeoutService(opts...)
	r.PseudoIPV4 = NewZoneSettingPseudoIPV4Service(opts...)
	r.ResponseBuffering = NewZoneSettingResponseBufferingService(opts...)
	r.RocketLoader = NewZoneSettingRocketLoaderService(opts...)
	r.SecurityHeaders = NewZoneSettingSecurityHeaderService(opts...)
	r.SecurityLevel = NewZoneSettingSecurityLevelService(opts...)
	r.ServerSideExcludes = NewZoneSettingServerSideExcludeService(opts...)
	r.SortQueryStringForCache = NewZoneSettingSortQueryStringForCacheService(opts...)
	r.SSL = NewZoneSettingSSLService(opts...)
	r.SSLRecommender = NewZoneSettingSSLRecommenderService(opts...)
	r.TLS1_3 = NewZoneSettingTLS1_3Service(opts...)
	r.TLSClientAuth = NewZoneSettingTLSClientAuthService(opts...)
	r.TrueClientIPHeader = NewZoneSettingTrueClientIPHeaderService(opts...)
	r.WAF = NewZoneSettingWAFService(opts...)
	r.Webp = NewZoneSettingWebpService(opts...)
	r.Websocket = NewZoneSettingWebsocketService(opts...)
	r.FontSettings = NewZoneSettingFontSettingService(opts...)
	return
}

// Edit settings for a zone.
func (r *ZoneSettingService) Edit(ctx context.Context, params ZoneSettingEditParams, opts ...option.RequestOption) (res *[]ZoneSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Available settings for your user in relation to a zone.
func (r *ZoneSettingService) Get(ctx context.Context, query ZoneSettingGetParams, opts ...option.RequestOption) (res *[]ZoneSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [Zones0rtt], [ZonesAdvancedDDOS], [ZonesAlwaysOnline],
// [ZonesAlwaysUseHTTPS], [ZonesAutomaticHTTPSRewrites], [ZonesBrotli],
// [ZonesBrowserCacheTTL], [ZonesBrowserCheck], [ZonesCacheLevel],
// [ZonesChallengeTTL], [ZonesCiphers],
// [ZoneSettingEditResponseZonesCnameFlattening], [ZonesDevelopmentMode],
// [ZonesEarlyHints], [ZoneSettingEditResponseZonesEdgeCacheTTL],
// [ZonesEmailObfuscation], [ZonesH2Prioritization], [ZonesHotlinkProtection],
// [ZonesHTTP2], [ZonesHTTP3], [ZonesImageResizing], [ZonesIPGeolocation],
// [ZonesIPV6], [ZoneSettingEditResponseZonesMaxUpload], [ZonesMinTLSVersion],
// [ZonesMinify], [ZonesMirage], [ZonesMobileRedirect], [ZonesNEL],
// [ZonesOpportunisticEncryption], [ZonesOpportunisticOnion],
// [ZonesOrangeToOrange], [ZonesOriginErrorPagePassThru], [ZonesPolish],
// [ZonesPrefetchPreload], [ZonesProxyReadTimeout], [ZonesPseudoIPV4],
// [ZonesBuffering], [ZonesRocketLoader],
// [ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [ZonesSecurityHeader], [ZonesSecurityLevel], [ZonesServerSideExclude],
// [ZoneSettingEditResponseZonesSha1Support], [ZonesSortQueryStringForCache],
// [ZonesSSL], [ZonesSSLRecommender], [ZoneSettingEditResponseZonesTLS1_2Only],
// [ZonesTLS1_3], [ZonesTLSClientAuth], [ZonesTrueClientIPHeader], [ZonesWAF],
// [ZonesWebp] or [ZonesWebsockets].
type ZoneSettingEditResponse interface {
	implementsZoneSettingEditResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneSettingEditResponse)(nil)).Elem(), "")
}

// Whether or not cname flattening is on.
type ZoneSettingEditResponseZonesCnameFlattening struct {
	// How to flatten the cname destination.
	ID ZoneSettingEditResponseZonesCnameFlatteningID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesCnameFlatteningValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesCnameFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesCnameFlatteningJSON `json:"-"`
}

// zoneSettingEditResponseZonesCnameFlatteningJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesCnameFlattening]
type zoneSettingEditResponseZonesCnameFlatteningJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesCnameFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesCnameFlattening) implementsZoneSettingEditResponse() {}

// How to flatten the cname destination.
type ZoneSettingEditResponseZonesCnameFlatteningID string

const (
	ZoneSettingEditResponseZonesCnameFlatteningIDCnameFlattening ZoneSettingEditResponseZonesCnameFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesCnameFlatteningValue string

const (
	ZoneSettingEditResponseZonesCnameFlatteningValueFlattenAtRoot ZoneSettingEditResponseZonesCnameFlatteningValue = "flatten_at_root"
	ZoneSettingEditResponseZonesCnameFlatteningValueFlattenAll    ZoneSettingEditResponseZonesCnameFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesCnameFlatteningEditable bool

const (
	ZoneSettingEditResponseZonesCnameFlatteningEditableTrue  ZoneSettingEditResponseZonesCnameFlatteningEditable = true
	ZoneSettingEditResponseZonesCnameFlatteningEditableFalse ZoneSettingEditResponseZonesCnameFlatteningEditable = false
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZoneSettingEditResponseZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesEdgeCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesEdgeCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesEdgeCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesEdgeCacheTTLJSON `json:"-"`
}

// zoneSettingEditResponseZonesEdgeCacheTTLJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesEdgeCacheTTL]
type zoneSettingEditResponseZonesEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesEdgeCacheTTL) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesEdgeCacheTTLID string

const (
	ZoneSettingEditResponseZonesEdgeCacheTTLIDEdgeCacheTTL ZoneSettingEditResponseZonesEdgeCacheTTLID = "edge_cache_ttl"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesEdgeCacheTTLValue float64

const (
	ZoneSettingEditResponseZonesEdgeCacheTTLValue30     ZoneSettingEditResponseZonesEdgeCacheTTLValue = 30
	ZoneSettingEditResponseZonesEdgeCacheTTLValue60     ZoneSettingEditResponseZonesEdgeCacheTTLValue = 60
	ZoneSettingEditResponseZonesEdgeCacheTTLValue300    ZoneSettingEditResponseZonesEdgeCacheTTLValue = 300
	ZoneSettingEditResponseZonesEdgeCacheTTLValue1200   ZoneSettingEditResponseZonesEdgeCacheTTLValue = 1200
	ZoneSettingEditResponseZonesEdgeCacheTTLValue1800   ZoneSettingEditResponseZonesEdgeCacheTTLValue = 1800
	ZoneSettingEditResponseZonesEdgeCacheTTLValue3600   ZoneSettingEditResponseZonesEdgeCacheTTLValue = 3600
	ZoneSettingEditResponseZonesEdgeCacheTTLValue7200   ZoneSettingEditResponseZonesEdgeCacheTTLValue = 7200
	ZoneSettingEditResponseZonesEdgeCacheTTLValue10800  ZoneSettingEditResponseZonesEdgeCacheTTLValue = 10800
	ZoneSettingEditResponseZonesEdgeCacheTTLValue14400  ZoneSettingEditResponseZonesEdgeCacheTTLValue = 14400
	ZoneSettingEditResponseZonesEdgeCacheTTLValue18000  ZoneSettingEditResponseZonesEdgeCacheTTLValue = 18000
	ZoneSettingEditResponseZonesEdgeCacheTTLValue28800  ZoneSettingEditResponseZonesEdgeCacheTTLValue = 28800
	ZoneSettingEditResponseZonesEdgeCacheTTLValue43200  ZoneSettingEditResponseZonesEdgeCacheTTLValue = 43200
	ZoneSettingEditResponseZonesEdgeCacheTTLValue57600  ZoneSettingEditResponseZonesEdgeCacheTTLValue = 57600
	ZoneSettingEditResponseZonesEdgeCacheTTLValue72000  ZoneSettingEditResponseZonesEdgeCacheTTLValue = 72000
	ZoneSettingEditResponseZonesEdgeCacheTTLValue86400  ZoneSettingEditResponseZonesEdgeCacheTTLValue = 86400
	ZoneSettingEditResponseZonesEdgeCacheTTLValue172800 ZoneSettingEditResponseZonesEdgeCacheTTLValue = 172800
	ZoneSettingEditResponseZonesEdgeCacheTTLValue259200 ZoneSettingEditResponseZonesEdgeCacheTTLValue = 259200
	ZoneSettingEditResponseZonesEdgeCacheTTLValue345600 ZoneSettingEditResponseZonesEdgeCacheTTLValue = 345600
	ZoneSettingEditResponseZonesEdgeCacheTTLValue432000 ZoneSettingEditResponseZonesEdgeCacheTTLValue = 432000
	ZoneSettingEditResponseZonesEdgeCacheTTLValue518400 ZoneSettingEditResponseZonesEdgeCacheTTLValue = 518400
	ZoneSettingEditResponseZonesEdgeCacheTTLValue604800 ZoneSettingEditResponseZonesEdgeCacheTTLValue = 604800
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesEdgeCacheTTLEditable bool

const (
	ZoneSettingEditResponseZonesEdgeCacheTTLEditableTrue  ZoneSettingEditResponseZonesEdgeCacheTTLEditable = true
	ZoneSettingEditResponseZonesEdgeCacheTTLEditableFalse ZoneSettingEditResponseZonesEdgeCacheTTLEditable = false
)

// Maximum size of an allowable upload.
type ZoneSettingEditResponseZonesMaxUpload struct {
	// identifier of the zone setting.
	ID ZoneSettingEditResponseZonesMaxUploadID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesMaxUploadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesMaxUploadJSON `json:"-"`
}

// zoneSettingEditResponseZonesMaxUploadJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesMaxUpload]
type zoneSettingEditResponseZonesMaxUploadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesMaxUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesMaxUpload) implementsZoneSettingEditResponse() {}

// identifier of the zone setting.
type ZoneSettingEditResponseZonesMaxUploadID string

const (
	ZoneSettingEditResponseZonesMaxUploadIDMaxUpload ZoneSettingEditResponseZonesMaxUploadID = "max_upload"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesMaxUploadValue float64

const (
	ZoneSettingEditResponseZonesMaxUploadValue100 ZoneSettingEditResponseZonesMaxUploadValue = 100
	ZoneSettingEditResponseZonesMaxUploadValue200 ZoneSettingEditResponseZonesMaxUploadValue = 200
	ZoneSettingEditResponseZonesMaxUploadValue500 ZoneSettingEditResponseZonesMaxUploadValue = 500
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesMaxUploadEditable bool

const (
	ZoneSettingEditResponseZonesMaxUploadEditableTrue  ZoneSettingEditResponseZonesMaxUploadEditable = true
	ZoneSettingEditResponseZonesMaxUploadEditableFalse ZoneSettingEditResponseZonesMaxUploadEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesAutomaticPlatformOptimization `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationJSON `json:"-"`
}

// zoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationJSON contains
// the JSON metadata for the struct
// [ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimization]
type zoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimization) implementsZoneSettingEditResponse() {
}

// ID of the zone setting.
type ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationID string

const (
	ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Allow SHA1 support.
type ZoneSettingEditResponseZonesSha1Support struct {
	// Zone setting identifier.
	ID ZoneSettingEditResponseZonesSha1SupportID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesSha1SupportValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesSha1SupportJSON `json:"-"`
}

// zoneSettingEditResponseZonesSha1SupportJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesSha1Support]
type zoneSettingEditResponseZonesSha1SupportJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesSha1Support) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesSha1Support) implementsZoneSettingEditResponse() {}

// Zone setting identifier.
type ZoneSettingEditResponseZonesSha1SupportID string

const (
	ZoneSettingEditResponseZonesSha1SupportIDSha1Support ZoneSettingEditResponseZonesSha1SupportID = "sha1_support"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesSha1SupportValue string

const (
	ZoneSettingEditResponseZonesSha1SupportValueOff ZoneSettingEditResponseZonesSha1SupportValue = "off"
	ZoneSettingEditResponseZonesSha1SupportValueOn  ZoneSettingEditResponseZonesSha1SupportValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesSha1SupportEditable bool

const (
	ZoneSettingEditResponseZonesSha1SupportEditableTrue  ZoneSettingEditResponseZonesSha1SupportEditable = true
	ZoneSettingEditResponseZonesSha1SupportEditableFalse ZoneSettingEditResponseZonesSha1SupportEditable = false
)

// Only allows TLS1.2.
type ZoneSettingEditResponseZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID ZoneSettingEditResponseZonesTLS1_2OnlyID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesTLS1_2OnlyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesTLS1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesTls1_2OnlyJSON `json:"-"`
}

// zoneSettingEditResponseZonesTls1_2OnlyJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesTLS1_2Only]
type zoneSettingEditResponseZonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesTLS1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesTLS1_2Only) implementsZoneSettingEditResponse() {}

// Zone setting identifier.
type ZoneSettingEditResponseZonesTLS1_2OnlyID string

const (
	ZoneSettingEditResponseZonesTLS1_2OnlyIDTLS1_2Only ZoneSettingEditResponseZonesTLS1_2OnlyID = "tls_1_2_only"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesTLS1_2OnlyValue string

const (
	ZoneSettingEditResponseZonesTLS1_2OnlyValueOff ZoneSettingEditResponseZonesTLS1_2OnlyValue = "off"
	ZoneSettingEditResponseZonesTLS1_2OnlyValueOn  ZoneSettingEditResponseZonesTLS1_2OnlyValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesTLS1_2OnlyEditable bool

const (
	ZoneSettingEditResponseZonesTLS1_2OnlyEditableTrue  ZoneSettingEditResponseZonesTLS1_2OnlyEditable = true
	ZoneSettingEditResponseZonesTLS1_2OnlyEditableFalse ZoneSettingEditResponseZonesTLS1_2OnlyEditable = false
)

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [Zones0rtt], [ZonesAdvancedDDOS], [ZonesAlwaysOnline],
// [ZonesAlwaysUseHTTPS], [ZonesAutomaticHTTPSRewrites], [ZonesBrotli],
// [ZonesBrowserCacheTTL], [ZonesBrowserCheck], [ZonesCacheLevel],
// [ZonesChallengeTTL], [ZonesCiphers],
// [ZoneSettingGetResponseZonesCnameFlattening], [ZonesDevelopmentMode],
// [ZonesEarlyHints], [ZoneSettingGetResponseZonesEdgeCacheTTL],
// [ZonesEmailObfuscation], [ZonesH2Prioritization], [ZonesHotlinkProtection],
// [ZonesHTTP2], [ZonesHTTP3], [ZonesImageResizing], [ZonesIPGeolocation],
// [ZonesIPV6], [ZoneSettingGetResponseZonesMaxUpload], [ZonesMinTLSVersion],
// [ZonesMinify], [ZonesMirage], [ZonesMobileRedirect], [ZonesNEL],
// [ZonesOpportunisticEncryption], [ZonesOpportunisticOnion],
// [ZonesOrangeToOrange], [ZonesOriginErrorPagePassThru], [ZonesPolish],
// [ZonesPrefetchPreload], [ZonesProxyReadTimeout], [ZonesPseudoIPV4],
// [ZonesBuffering], [ZonesRocketLoader],
// [ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimization],
// [ZonesSecurityHeader], [ZonesSecurityLevel], [ZonesServerSideExclude],
// [ZoneSettingGetResponseZonesSha1Support], [ZonesSortQueryStringForCache],
// [ZonesSSL], [ZonesSSLRecommender], [ZoneSettingGetResponseZonesTLS1_2Only],
// [ZonesTLS1_3], [ZonesTLSClientAuth], [ZonesTrueClientIPHeader], [ZonesWAF],
// [ZonesWebp] or [ZonesWebsockets].
type ZoneSettingGetResponse interface {
	implementsZoneSettingGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneSettingGetResponse)(nil)).Elem(), "")
}

// Whether or not cname flattening is on.
type ZoneSettingGetResponseZonesCnameFlattening struct {
	// How to flatten the cname destination.
	ID ZoneSettingGetResponseZonesCnameFlatteningID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesCnameFlatteningValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesCnameFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesCnameFlatteningJSON `json:"-"`
}

// zoneSettingGetResponseZonesCnameFlatteningJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesCnameFlattening]
type zoneSettingGetResponseZonesCnameFlatteningJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesCnameFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingGetResponseZonesCnameFlattening) implementsZoneSettingGetResponse() {}

// How to flatten the cname destination.
type ZoneSettingGetResponseZonesCnameFlatteningID string

const (
	ZoneSettingGetResponseZonesCnameFlatteningIDCnameFlattening ZoneSettingGetResponseZonesCnameFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesCnameFlatteningValue string

const (
	ZoneSettingGetResponseZonesCnameFlatteningValueFlattenAtRoot ZoneSettingGetResponseZonesCnameFlatteningValue = "flatten_at_root"
	ZoneSettingGetResponseZonesCnameFlatteningValueFlattenAll    ZoneSettingGetResponseZonesCnameFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesCnameFlatteningEditable bool

const (
	ZoneSettingGetResponseZonesCnameFlatteningEditableTrue  ZoneSettingGetResponseZonesCnameFlatteningEditable = true
	ZoneSettingGetResponseZonesCnameFlatteningEditableFalse ZoneSettingGetResponseZonesCnameFlatteningEditable = false
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZoneSettingGetResponseZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesEdgeCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesEdgeCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesEdgeCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesEdgeCacheTTLJSON `json:"-"`
}

// zoneSettingGetResponseZonesEdgeCacheTTLJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesEdgeCacheTTL]
type zoneSettingGetResponseZonesEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingGetResponseZonesEdgeCacheTTL) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesEdgeCacheTTLID string

const (
	ZoneSettingGetResponseZonesEdgeCacheTTLIDEdgeCacheTTL ZoneSettingGetResponseZonesEdgeCacheTTLID = "edge_cache_ttl"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesEdgeCacheTTLValue float64

const (
	ZoneSettingGetResponseZonesEdgeCacheTTLValue30     ZoneSettingGetResponseZonesEdgeCacheTTLValue = 30
	ZoneSettingGetResponseZonesEdgeCacheTTLValue60     ZoneSettingGetResponseZonesEdgeCacheTTLValue = 60
	ZoneSettingGetResponseZonesEdgeCacheTTLValue300    ZoneSettingGetResponseZonesEdgeCacheTTLValue = 300
	ZoneSettingGetResponseZonesEdgeCacheTTLValue1200   ZoneSettingGetResponseZonesEdgeCacheTTLValue = 1200
	ZoneSettingGetResponseZonesEdgeCacheTTLValue1800   ZoneSettingGetResponseZonesEdgeCacheTTLValue = 1800
	ZoneSettingGetResponseZonesEdgeCacheTTLValue3600   ZoneSettingGetResponseZonesEdgeCacheTTLValue = 3600
	ZoneSettingGetResponseZonesEdgeCacheTTLValue7200   ZoneSettingGetResponseZonesEdgeCacheTTLValue = 7200
	ZoneSettingGetResponseZonesEdgeCacheTTLValue10800  ZoneSettingGetResponseZonesEdgeCacheTTLValue = 10800
	ZoneSettingGetResponseZonesEdgeCacheTTLValue14400  ZoneSettingGetResponseZonesEdgeCacheTTLValue = 14400
	ZoneSettingGetResponseZonesEdgeCacheTTLValue18000  ZoneSettingGetResponseZonesEdgeCacheTTLValue = 18000
	ZoneSettingGetResponseZonesEdgeCacheTTLValue28800  ZoneSettingGetResponseZonesEdgeCacheTTLValue = 28800
	ZoneSettingGetResponseZonesEdgeCacheTTLValue43200  ZoneSettingGetResponseZonesEdgeCacheTTLValue = 43200
	ZoneSettingGetResponseZonesEdgeCacheTTLValue57600  ZoneSettingGetResponseZonesEdgeCacheTTLValue = 57600
	ZoneSettingGetResponseZonesEdgeCacheTTLValue72000  ZoneSettingGetResponseZonesEdgeCacheTTLValue = 72000
	ZoneSettingGetResponseZonesEdgeCacheTTLValue86400  ZoneSettingGetResponseZonesEdgeCacheTTLValue = 86400
	ZoneSettingGetResponseZonesEdgeCacheTTLValue172800 ZoneSettingGetResponseZonesEdgeCacheTTLValue = 172800
	ZoneSettingGetResponseZonesEdgeCacheTTLValue259200 ZoneSettingGetResponseZonesEdgeCacheTTLValue = 259200
	ZoneSettingGetResponseZonesEdgeCacheTTLValue345600 ZoneSettingGetResponseZonesEdgeCacheTTLValue = 345600
	ZoneSettingGetResponseZonesEdgeCacheTTLValue432000 ZoneSettingGetResponseZonesEdgeCacheTTLValue = 432000
	ZoneSettingGetResponseZonesEdgeCacheTTLValue518400 ZoneSettingGetResponseZonesEdgeCacheTTLValue = 518400
	ZoneSettingGetResponseZonesEdgeCacheTTLValue604800 ZoneSettingGetResponseZonesEdgeCacheTTLValue = 604800
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesEdgeCacheTTLEditable bool

const (
	ZoneSettingGetResponseZonesEdgeCacheTTLEditableTrue  ZoneSettingGetResponseZonesEdgeCacheTTLEditable = true
	ZoneSettingGetResponseZonesEdgeCacheTTLEditableFalse ZoneSettingGetResponseZonesEdgeCacheTTLEditable = false
)

// Maximum size of an allowable upload.
type ZoneSettingGetResponseZonesMaxUpload struct {
	// identifier of the zone setting.
	ID ZoneSettingGetResponseZonesMaxUploadID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesMaxUploadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesMaxUploadJSON `json:"-"`
}

// zoneSettingGetResponseZonesMaxUploadJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesMaxUpload]
type zoneSettingGetResponseZonesMaxUploadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesMaxUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingGetResponseZonesMaxUpload) implementsZoneSettingGetResponse() {}

// identifier of the zone setting.
type ZoneSettingGetResponseZonesMaxUploadID string

const (
	ZoneSettingGetResponseZonesMaxUploadIDMaxUpload ZoneSettingGetResponseZonesMaxUploadID = "max_upload"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesMaxUploadValue float64

const (
	ZoneSettingGetResponseZonesMaxUploadValue100 ZoneSettingGetResponseZonesMaxUploadValue = 100
	ZoneSettingGetResponseZonesMaxUploadValue200 ZoneSettingGetResponseZonesMaxUploadValue = 200
	ZoneSettingGetResponseZonesMaxUploadValue500 ZoneSettingGetResponseZonesMaxUploadValue = 500
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesMaxUploadEditable bool

const (
	ZoneSettingGetResponseZonesMaxUploadEditableTrue  ZoneSettingGetResponseZonesMaxUploadEditable = true
	ZoneSettingGetResponseZonesMaxUploadEditableFalse ZoneSettingGetResponseZonesMaxUploadEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesAutomaticPlatformOptimization `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationJSON `json:"-"`
}

// zoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationJSON contains the
// JSON metadata for the struct
// [ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimization]
type zoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimization) implementsZoneSettingGetResponse() {
}

// ID of the zone setting.
type ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationID string

const (
	ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Allow SHA1 support.
type ZoneSettingGetResponseZonesSha1Support struct {
	// Zone setting identifier.
	ID ZoneSettingGetResponseZonesSha1SupportID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesSha1SupportValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesSha1SupportJSON `json:"-"`
}

// zoneSettingGetResponseZonesSha1SupportJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesSha1Support]
type zoneSettingGetResponseZonesSha1SupportJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesSha1Support) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingGetResponseZonesSha1Support) implementsZoneSettingGetResponse() {}

// Zone setting identifier.
type ZoneSettingGetResponseZonesSha1SupportID string

const (
	ZoneSettingGetResponseZonesSha1SupportIDSha1Support ZoneSettingGetResponseZonesSha1SupportID = "sha1_support"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesSha1SupportValue string

const (
	ZoneSettingGetResponseZonesSha1SupportValueOff ZoneSettingGetResponseZonesSha1SupportValue = "off"
	ZoneSettingGetResponseZonesSha1SupportValueOn  ZoneSettingGetResponseZonesSha1SupportValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesSha1SupportEditable bool

const (
	ZoneSettingGetResponseZonesSha1SupportEditableTrue  ZoneSettingGetResponseZonesSha1SupportEditable = true
	ZoneSettingGetResponseZonesSha1SupportEditableFalse ZoneSettingGetResponseZonesSha1SupportEditable = false
)

// Only allows TLS1.2.
type ZoneSettingGetResponseZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID ZoneSettingGetResponseZonesTLS1_2OnlyID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesTLS1_2OnlyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesTLS1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesTls1_2OnlyJSON `json:"-"`
}

// zoneSettingGetResponseZonesTls1_2OnlyJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesTLS1_2Only]
type zoneSettingGetResponseZonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesTLS1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingGetResponseZonesTLS1_2Only) implementsZoneSettingGetResponse() {}

// Zone setting identifier.
type ZoneSettingGetResponseZonesTLS1_2OnlyID string

const (
	ZoneSettingGetResponseZonesTLS1_2OnlyIDTLS1_2Only ZoneSettingGetResponseZonesTLS1_2OnlyID = "tls_1_2_only"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesTLS1_2OnlyValue string

const (
	ZoneSettingGetResponseZonesTLS1_2OnlyValueOff ZoneSettingGetResponseZonesTLS1_2OnlyValue = "off"
	ZoneSettingGetResponseZonesTLS1_2OnlyValueOn  ZoneSettingGetResponseZonesTLS1_2OnlyValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesTLS1_2OnlyEditable bool

const (
	ZoneSettingGetResponseZonesTLS1_2OnlyEditableTrue  ZoneSettingGetResponseZonesTLS1_2OnlyEditable = true
	ZoneSettingGetResponseZonesTLS1_2OnlyEditableFalse ZoneSettingGetResponseZonesTLS1_2OnlyEditable = false
)

type ZoneSettingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// One or more zone setting objects. Must contain an ID and a value.
	Items param.Field[[]ZoneSettingEditParamsItem] `json:"items,required"`
}

func (r ZoneSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// 0-RTT session resumption enabled for this zone.
//
// Satisfied by [Zones0rttParam], [ZonesAdvancedDDOSParam],
// [ZonesAlwaysOnlineParam], [ZonesAlwaysUseHTTPSParam],
// [ZonesAutomaticHTTPSRewritesParam], [ZonesBrotliParam],
// [ZonesBrowserCacheTTLParam], [ZonesBrowserCheckParam], [ZonesCacheLevelParam],
// [ZonesChallengeTTLParam], [ZonesCiphersParam],
// [ZoneSettingEditParamsItemsZonesCnameFlattening], [ZonesDevelopmentModeParam],
// [ZonesEarlyHintsParam], [ZoneSettingEditParamsItemsZonesEdgeCacheTTL],
// [ZonesEmailObfuscationParam], [ZonesH2PrioritizationParam],
// [ZonesHotlinkProtectionParam], [ZonesHTTP2Param], [ZonesHTTP3Param],
// [ZonesImageResizingParam], [ZonesIPGeolocationParam], [ZonesIPV6Param],
// [ZoneSettingEditParamsItemsZonesMaxUpload], [ZonesMinTLSVersionParam],
// [ZonesMinifyParam], [ZonesMirageParam], [ZonesMobileRedirectParam],
// [ZonesNELParam], [ZonesOpportunisticEncryptionParam],
// [ZonesOpportunisticOnionParam], [ZonesOrangeToOrangeParam],
// [ZonesOriginErrorPagePassThruParam], [ZonesPolishParam],
// [ZonesPrefetchPreloadParam], [ZonesProxyReadTimeoutParam],
// [ZonesPseudoIPV4Param], [ZonesBufferingParam], [ZonesRocketLoaderParam],
// [ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization],
// [ZonesSecurityHeaderParam], [ZonesSecurityLevelParam],
// [ZonesServerSideExcludeParam], [ZoneSettingEditParamsItemsZonesSha1Support],
// [ZonesSortQueryStringForCacheParam], [ZonesSSLParam],
// [ZonesSSLRecommenderParam], [ZoneSettingEditParamsItemsZonesTLS1_2Only],
// [ZonesTLS1_3Param], [ZonesTLSClientAuthParam], [ZonesTrueClientIPHeaderParam],
// [ZonesWAFParam], [ZonesWebpParam], [ZonesWebsocketsParam].
type ZoneSettingEditParamsItem interface {
	implementsZoneSettingEditParamsItem()
}

// Whether or not cname flattening is on.
type ZoneSettingEditParamsItemsZonesCnameFlattening struct {
	// How to flatten the cname destination.
	ID param.Field[ZoneSettingEditParamsItemsZonesCnameFlatteningID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesCnameFlatteningValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesCnameFlattening) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesCnameFlattening) implementsZoneSettingEditParamsItem() {}

// How to flatten the cname destination.
type ZoneSettingEditParamsItemsZonesCnameFlatteningID string

const (
	ZoneSettingEditParamsItemsZonesCnameFlatteningIDCnameFlattening ZoneSettingEditParamsItemsZonesCnameFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesCnameFlatteningValue string

const (
	ZoneSettingEditParamsItemsZonesCnameFlatteningValueFlattenAtRoot ZoneSettingEditParamsItemsZonesCnameFlatteningValue = "flatten_at_root"
	ZoneSettingEditParamsItemsZonesCnameFlatteningValueFlattenAll    ZoneSettingEditParamsItemsZonesCnameFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesCnameFlatteningEditable bool

const (
	ZoneSettingEditParamsItemsZonesCnameFlatteningEditableTrue  ZoneSettingEditParamsItemsZonesCnameFlatteningEditable = true
	ZoneSettingEditParamsItemsZonesCnameFlatteningEditableFalse ZoneSettingEditParamsItemsZonesCnameFlatteningEditable = false
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZoneSettingEditParamsItemsZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesEdgeCacheTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesEdgeCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesEdgeCacheTTL) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesEdgeCacheTTLID string

const (
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLIDEdgeCacheTTL ZoneSettingEditParamsItemsZonesEdgeCacheTTLID = "edge_cache_ttl"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue float64

const (
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue30     ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 30
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue60     ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 60
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue300    ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 300
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue1200   ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 1200
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue1800   ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 1800
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue3600   ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 3600
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue7200   ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 7200
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue10800  ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 10800
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue14400  ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 14400
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue18000  ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 18000
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue28800  ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 28800
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue43200  ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 43200
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue57600  ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 57600
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue72000  ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 72000
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue86400  ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 86400
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue172800 ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 172800
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue259200 ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 259200
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue345600 ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 345600
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue432000 ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 432000
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue518400 ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 518400
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue604800 ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue = 604800
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesEdgeCacheTTLEditable bool

const (
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLEditableTrue  ZoneSettingEditParamsItemsZonesEdgeCacheTTLEditable = true
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLEditableFalse ZoneSettingEditParamsItemsZonesEdgeCacheTTLEditable = false
)

// Maximum size of an allowable upload.
type ZoneSettingEditParamsItemsZonesMaxUpload struct {
	// identifier of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesMaxUploadID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesMaxUploadValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesMaxUpload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesMaxUpload) implementsZoneSettingEditParamsItem() {}

// identifier of the zone setting.
type ZoneSettingEditParamsItemsZonesMaxUploadID string

const (
	ZoneSettingEditParamsItemsZonesMaxUploadIDMaxUpload ZoneSettingEditParamsItemsZonesMaxUploadID = "max_upload"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesMaxUploadValue float64

const (
	ZoneSettingEditParamsItemsZonesMaxUploadValue100 ZoneSettingEditParamsItemsZonesMaxUploadValue = 100
	ZoneSettingEditParamsItemsZonesMaxUploadValue200 ZoneSettingEditParamsItemsZonesMaxUploadValue = 200
	ZoneSettingEditParamsItemsZonesMaxUploadValue500 ZoneSettingEditParamsItemsZonesMaxUploadValue = 500
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMaxUploadEditable bool

const (
	ZoneSettingEditParamsItemsZonesMaxUploadEditableTrue  ZoneSettingEditParamsItemsZonesMaxUploadEditable = true
	ZoneSettingEditParamsItemsZonesMaxUploadEditableFalse ZoneSettingEditParamsItemsZonesMaxUploadEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesAutomaticPlatformOptimizationParam] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization) implementsZoneSettingEditParamsItem() {
}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID string

const (
	ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Allow SHA1 support.
type ZoneSettingEditParamsItemsZonesSha1Support struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsZonesSha1SupportID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesSha1SupportValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesSha1Support) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesSha1Support) implementsZoneSettingEditParamsItem() {}

// Zone setting identifier.
type ZoneSettingEditParamsItemsZonesSha1SupportID string

const (
	ZoneSettingEditParamsItemsZonesSha1SupportIDSha1Support ZoneSettingEditParamsItemsZonesSha1SupportID = "sha1_support"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesSha1SupportValue string

const (
	ZoneSettingEditParamsItemsZonesSha1SupportValueOff ZoneSettingEditParamsItemsZonesSha1SupportValue = "off"
	ZoneSettingEditParamsItemsZonesSha1SupportValueOn  ZoneSettingEditParamsItemsZonesSha1SupportValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSha1SupportEditable bool

const (
	ZoneSettingEditParamsItemsZonesSha1SupportEditableTrue  ZoneSettingEditParamsItemsZonesSha1SupportEditable = true
	ZoneSettingEditParamsItemsZonesSha1SupportEditableFalse ZoneSettingEditParamsItemsZonesSha1SupportEditable = false
)

// Only allows TLS1.2.
type ZoneSettingEditParamsItemsZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsZonesTLS1_2OnlyID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesTLS1_2OnlyValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesTLS1_2Only) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesTLS1_2Only) implementsZoneSettingEditParamsItem() {}

// Zone setting identifier.
type ZoneSettingEditParamsItemsZonesTLS1_2OnlyID string

const (
	ZoneSettingEditParamsItemsZonesTLS1_2OnlyIDTLS1_2Only ZoneSettingEditParamsItemsZonesTLS1_2OnlyID = "tls_1_2_only"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesTLS1_2OnlyValue string

const (
	ZoneSettingEditParamsItemsZonesTLS1_2OnlyValueOff ZoneSettingEditParamsItemsZonesTLS1_2OnlyValue = "off"
	ZoneSettingEditParamsItemsZonesTLS1_2OnlyValueOn  ZoneSettingEditParamsItemsZonesTLS1_2OnlyValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTLS1_2OnlyEditable bool

const (
	ZoneSettingEditParamsItemsZonesTLS1_2OnlyEditableTrue  ZoneSettingEditParamsItemsZonesTLS1_2OnlyEditable = true
	ZoneSettingEditParamsItemsZonesTLS1_2OnlyEditableFalse ZoneSettingEditParamsItemsZonesTLS1_2OnlyEditable = false
)

type ZoneSettingEditResponseEnvelope struct {
	Errors   []ZoneSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                `json:"success,required"`
	Result  []ZoneSettingEditResponse           `json:"result"`
	JSON    zoneSettingEditResponseEnvelopeJSON `json:"-"`
}

// zoneSettingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseEnvelope]
type zoneSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEditResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseEnvelopeErrors]
type zoneSettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEditResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseEnvelopeMessages]
type zoneSettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingGetResponseEnvelope struct {
	Errors   []ZoneSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                               `json:"success,required"`
	Result  []ZoneSettingGetResponse           `json:"result"`
	JSON    zoneSettingGetResponseEnvelopeJSON `json:"-"`
}

// zoneSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseEnvelope]
type zoneSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseEnvelopeErrors]
type zoneSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseEnvelopeMessages]
type zoneSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
