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

// SettingService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSettingService] method instead.
type SettingService struct {
	Options                       []option.RequestOption
	ZeroRtt                       *SettingZeroRttService
	AdvancedDDOS                  *SettingAdvancedDDOSService
	AlwaysOnline                  *SettingAlwaysOnlineService
	AlwaysUseHTTPS                *SettingAlwaysUseHTTPSService
	AutomaticHTTPSRewrites        *SettingAutomaticHTTPSRewriteService
	AutomaticPlatformOptimization *SettingAutomaticPlatformOptimizationService
	Brotli                        *SettingBrotliService
	BrowserCacheTTL               *SettingBrowserCacheTTLService
	BrowserCheck                  *SettingBrowserCheckService
	CacheLevel                    *SettingCacheLevelService
	ChallengeTTL                  *SettingChallengeTTLService
	Ciphers                       *SettingCipherService
	DevelopmentMode               *SettingDevelopmentModeService
	EarlyHints                    *SettingEarlyHintService
	EmailObfuscation              *SettingEmailObfuscationService
	H2Prioritization              *SettingH2PrioritizationService
	HotlinkProtection             *SettingHotlinkProtectionService
	HTTP2                         *SettingHTTP2Service
	HTTP3                         *SettingHTTP3Service
	ImageResizing                 *SettingImageResizingService
	IPGeolocation                 *SettingIPGeolocationService
	IPV6                          *SettingIPV6Service
	MinTLSVersion                 *SettingMinTLSVersionService
	Minify                        *SettingMinifyService
	Mirage                        *SettingMirageService
	MobileRedirect                *SettingMobileRedirectService
	NEL                           *SettingNELService
	OpportunisticEncryption       *SettingOpportunisticEncryptionService
	OpportunisticOnion            *SettingOpportunisticOnionService
	OrangeToOrange                *SettingOrangeToOrangeService
	OriginErrorPagePassThru       *SettingOriginErrorPagePassThruService
	OriginMaxHTTPVersion          *SettingOriginMaxHTTPVersionService
	Polish                        *SettingPolishService
	PrefetchPreload               *SettingPrefetchPreloadService
	ProxyReadTimeout              *SettingProxyReadTimeoutService
	PseudoIPV4                    *SettingPseudoIPV4Service
	ResponseBuffering             *SettingResponseBufferingService
	RocketLoader                  *SettingRocketLoaderService
	SecurityHeaders               *SettingSecurityHeaderService
	SecurityLevel                 *SettingSecurityLevelService
	ServerSideExcludes            *SettingServerSideExcludeService
	SortQueryStringForCache       *SettingSortQueryStringForCacheService
	SSL                           *SettingSSLService
	SSLRecommender                *SettingSSLRecommenderService
	TLS1_3                        *SettingTLS1_3Service
	TLSClientAuth                 *SettingTLSClientAuthService
	TrueClientIPHeader            *SettingTrueClientIPHeaderService
	WAF                           *SettingWAFService
	Webp                          *SettingWebpService
	Websocket                     *SettingWebsocketService
	FontSettings                  *SettingFontSettingService
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	r.ZeroRtt = NewSettingZeroRttService(opts...)
	r.AdvancedDDOS = NewSettingAdvancedDDOSService(opts...)
	r.AlwaysOnline = NewSettingAlwaysOnlineService(opts...)
	r.AlwaysUseHTTPS = NewSettingAlwaysUseHTTPSService(opts...)
	r.AutomaticHTTPSRewrites = NewSettingAutomaticHTTPSRewriteService(opts...)
	r.AutomaticPlatformOptimization = NewSettingAutomaticPlatformOptimizationService(opts...)
	r.Brotli = NewSettingBrotliService(opts...)
	r.BrowserCacheTTL = NewSettingBrowserCacheTTLService(opts...)
	r.BrowserCheck = NewSettingBrowserCheckService(opts...)
	r.CacheLevel = NewSettingCacheLevelService(opts...)
	r.ChallengeTTL = NewSettingChallengeTTLService(opts...)
	r.Ciphers = NewSettingCipherService(opts...)
	r.DevelopmentMode = NewSettingDevelopmentModeService(opts...)
	r.EarlyHints = NewSettingEarlyHintService(opts...)
	r.EmailObfuscation = NewSettingEmailObfuscationService(opts...)
	r.H2Prioritization = NewSettingH2PrioritizationService(opts...)
	r.HotlinkProtection = NewSettingHotlinkProtectionService(opts...)
	r.HTTP2 = NewSettingHTTP2Service(opts...)
	r.HTTP3 = NewSettingHTTP3Service(opts...)
	r.ImageResizing = NewSettingImageResizingService(opts...)
	r.IPGeolocation = NewSettingIPGeolocationService(opts...)
	r.IPV6 = NewSettingIPV6Service(opts...)
	r.MinTLSVersion = NewSettingMinTLSVersionService(opts...)
	r.Minify = NewSettingMinifyService(opts...)
	r.Mirage = NewSettingMirageService(opts...)
	r.MobileRedirect = NewSettingMobileRedirectService(opts...)
	r.NEL = NewSettingNELService(opts...)
	r.OpportunisticEncryption = NewSettingOpportunisticEncryptionService(opts...)
	r.OpportunisticOnion = NewSettingOpportunisticOnionService(opts...)
	r.OrangeToOrange = NewSettingOrangeToOrangeService(opts...)
	r.OriginErrorPagePassThru = NewSettingOriginErrorPagePassThruService(opts...)
	r.OriginMaxHTTPVersion = NewSettingOriginMaxHTTPVersionService(opts...)
	r.Polish = NewSettingPolishService(opts...)
	r.PrefetchPreload = NewSettingPrefetchPreloadService(opts...)
	r.ProxyReadTimeout = NewSettingProxyReadTimeoutService(opts...)
	r.PseudoIPV4 = NewSettingPseudoIPV4Service(opts...)
	r.ResponseBuffering = NewSettingResponseBufferingService(opts...)
	r.RocketLoader = NewSettingRocketLoaderService(opts...)
	r.SecurityHeaders = NewSettingSecurityHeaderService(opts...)
	r.SecurityLevel = NewSettingSecurityLevelService(opts...)
	r.ServerSideExcludes = NewSettingServerSideExcludeService(opts...)
	r.SortQueryStringForCache = NewSettingSortQueryStringForCacheService(opts...)
	r.SSL = NewSettingSSLService(opts...)
	r.SSLRecommender = NewSettingSSLRecommenderService(opts...)
	r.TLS1_3 = NewSettingTLS1_3Service(opts...)
	r.TLSClientAuth = NewSettingTLSClientAuthService(opts...)
	r.TrueClientIPHeader = NewSettingTrueClientIPHeaderService(opts...)
	r.WAF = NewSettingWAFService(opts...)
	r.Webp = NewSettingWebpService(opts...)
	r.Websocket = NewSettingWebsocketService(opts...)
	r.FontSettings = NewSettingFontSettingService(opts...)
	return
}

// Available settings for your user in relation to a zone.
func (r *SettingService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *[]SettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingListResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit settings for a zone.
func (r *SettingService) Edit(ctx context.Context, zoneID string, body SettingEditParams, opts ...option.RequestOption) (res *[]SettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [SettingListResponseZones0rtt],
// [SettingListResponseZonesAdvancedDDOS], [SettingListResponseZonesAlwaysOnline],
// [SettingListResponseZonesAlwaysUseHTTPS],
// [SettingListResponseZonesAutomaticHTTPSRewrites],
// [SettingListResponseZonesBrotli], [SettingListResponseZonesBrowserCacheTTL],
// [SettingListResponseZonesBrowserCheck], [SettingListResponseZonesCacheLevel],
// [SettingListResponseZonesChallengeTTL], [SettingListResponseZonesCiphers],
// [SettingListResponseZonesCnameFlattening],
// [SettingListResponseZonesDevelopmentMode], [SettingListResponseZonesEarlyHints],
// [SettingListResponseZonesEdgeCacheTTL],
// [SettingListResponseZonesEmailObfuscation],
// [SettingListResponseZonesH2Prioritization],
// [SettingListResponseZonesHotlinkProtection], [SettingListResponseZonesHTTP2],
// [SettingListResponseZonesHTTP3], [SettingListResponseZonesImageResizing],
// [SettingListResponseZonesIPGeolocation], [SettingListResponseZonesIPV6],
// [SettingListResponseZonesMaxUpload], [SettingListResponseZonesMinTLSVersion],
// [SettingListResponseZonesMinify], [SettingListResponseZonesMirage],
// [SettingListResponseZonesMobileRedirect], [SettingListResponseZonesNEL],
// [SettingListResponseZonesOpportunisticEncryption],
// [SettingListResponseZonesOpportunisticOnion],
// [SettingListResponseZonesOrangeToOrange],
// [SettingListResponseZonesOriginErrorPagePassThru],
// [SettingListResponseZonesPolish], [SettingListResponseZonesPrefetchPreload],
// [SettingListResponseZonesProxyReadTimeout],
// [SettingListResponseZonesPseudoIPV4],
// [SettingListResponseZonesResponseBuffering],
// [SettingListResponseZonesRocketLoader],
// [SettingListResponseZonesSchemasAutomaticPlatformOptimization],
// [SettingListResponseZonesSecurityHeader],
// [SettingListResponseZonesSecurityLevel],
// [SettingListResponseZonesServerSideExclude],
// [SettingListResponseZonesSha1Support],
// [SettingListResponseZonesSortQueryStringForCache],
// [SettingListResponseZonesSSL], [SettingListResponseZonesSSLRecommender],
// [SettingListResponseZonesTLS1_2Only], [SettingListResponseZonesTLS1_3],
// [SettingListResponseZonesTLSClientAuth],
// [SettingListResponseZonesTrueClientIPHeader], [SettingListResponseZonesWAF],
// [SettingListResponseZonesWebp] or [SettingListResponseZonesWebsockets].
type SettingListResponse interface {
	implementsSettingListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SettingListResponse)(nil)).Elem(), "")
}

// 0-RTT session resumption enabled for this zone.
type SettingListResponseZones0rtt struct {
	// ID of the zone setting.
	ID SettingListResponseZones0rttID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZones0rttValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZones0rttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZones0rttJSON `json:"-"`
}

// settingListResponseZones0rttJSON contains the JSON metadata for the struct
// [SettingListResponseZones0rtt]
type settingListResponseZones0rttJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZones0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZones0rtt) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZones0rttID string

const (
	SettingListResponseZones0rttID0rtt SettingListResponseZones0rttID = "0rtt"
)

// Current value of the zone setting.
type SettingListResponseZones0rttValue string

const (
	SettingListResponseZones0rttValueOn  SettingListResponseZones0rttValue = "on"
	SettingListResponseZones0rttValueOff SettingListResponseZones0rttValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZones0rttEditable bool

const (
	SettingListResponseZones0rttEditableTrue  SettingListResponseZones0rttEditable = true
	SettingListResponseZones0rttEditableFalse SettingListResponseZones0rttEditable = false
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type SettingListResponseZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID SettingListResponseZonesAdvancedDDOSID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesAdvancedDDOSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesAdvancedDDOSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesAdvancedDDOSJSON `json:"-"`
}

// settingListResponseZonesAdvancedDDOSJSON contains the JSON metadata for the
// struct [SettingListResponseZonesAdvancedDDOS]
type settingListResponseZonesAdvancedDDOSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesAdvancedDDOS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesAdvancedDDOS) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesAdvancedDDOSID string

const (
	SettingListResponseZonesAdvancedDDOSIDAdvancedDDOS SettingListResponseZonesAdvancedDDOSID = "advanced_ddos"
)

// Current value of the zone setting.
type SettingListResponseZonesAdvancedDDOSValue string

const (
	SettingListResponseZonesAdvancedDDOSValueOn  SettingListResponseZonesAdvancedDDOSValue = "on"
	SettingListResponseZonesAdvancedDDOSValueOff SettingListResponseZonesAdvancedDDOSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesAdvancedDDOSEditable bool

const (
	SettingListResponseZonesAdvancedDDOSEditableTrue  SettingListResponseZonesAdvancedDDOSEditable = true
	SettingListResponseZonesAdvancedDDOSEditableFalse SettingListResponseZonesAdvancedDDOSEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type SettingListResponseZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID SettingListResponseZonesAlwaysOnlineID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesAlwaysOnlineValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesAlwaysOnlineJSON `json:"-"`
}

// settingListResponseZonesAlwaysOnlineJSON contains the JSON metadata for the
// struct [SettingListResponseZonesAlwaysOnline]
type settingListResponseZonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesAlwaysOnline) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesAlwaysOnlineID string

const (
	SettingListResponseZonesAlwaysOnlineIDAlwaysOnline SettingListResponseZonesAlwaysOnlineID = "always_online"
)

// Current value of the zone setting.
type SettingListResponseZonesAlwaysOnlineValue string

const (
	SettingListResponseZonesAlwaysOnlineValueOn  SettingListResponseZonesAlwaysOnlineValue = "on"
	SettingListResponseZonesAlwaysOnlineValueOff SettingListResponseZonesAlwaysOnlineValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesAlwaysOnlineEditable bool

const (
	SettingListResponseZonesAlwaysOnlineEditableTrue  SettingListResponseZonesAlwaysOnlineEditable = true
	SettingListResponseZonesAlwaysOnlineEditableFalse SettingListResponseZonesAlwaysOnlineEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingListResponseZonesAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID SettingListResponseZonesAlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesAlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesAlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesAlwaysUseHTTPSJSON `json:"-"`
}

// settingListResponseZonesAlwaysUseHTTPSJSON contains the JSON metadata for the
// struct [SettingListResponseZonesAlwaysUseHTTPS]
type settingListResponseZonesAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesAlwaysUseHTTPS) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesAlwaysUseHTTPSID string

const (
	SettingListResponseZonesAlwaysUseHTTPSIDAlwaysUseHTTPS SettingListResponseZonesAlwaysUseHTTPSID = "always_use_https"
)

// Current value of the zone setting.
type SettingListResponseZonesAlwaysUseHTTPSValue string

const (
	SettingListResponseZonesAlwaysUseHTTPSValueOn  SettingListResponseZonesAlwaysUseHTTPSValue = "on"
	SettingListResponseZonesAlwaysUseHTTPSValueOff SettingListResponseZonesAlwaysUseHTTPSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesAlwaysUseHTTPSEditable bool

const (
	SettingListResponseZonesAlwaysUseHTTPSEditableTrue  SettingListResponseZonesAlwaysUseHTTPSEditable = true
	SettingListResponseZonesAlwaysUseHTTPSEditableFalse SettingListResponseZonesAlwaysUseHTTPSEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingListResponseZonesAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID SettingListResponseZonesAutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesAutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesAutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesAutomaticHTTPSRewritesJSON `json:"-"`
}

// settingListResponseZonesAutomaticHTTPSRewritesJSON contains the JSON metadata
// for the struct [SettingListResponseZonesAutomaticHTTPSRewrites]
type settingListResponseZonesAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesAutomaticHTTPSRewrites) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesAutomaticHTTPSRewritesID string

const (
	SettingListResponseZonesAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites SettingListResponseZonesAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type SettingListResponseZonesAutomaticHTTPSRewritesValue string

const (
	SettingListResponseZonesAutomaticHTTPSRewritesValueOn  SettingListResponseZonesAutomaticHTTPSRewritesValue = "on"
	SettingListResponseZonesAutomaticHTTPSRewritesValueOff SettingListResponseZonesAutomaticHTTPSRewritesValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesAutomaticHTTPSRewritesEditable bool

const (
	SettingListResponseZonesAutomaticHTTPSRewritesEditableTrue  SettingListResponseZonesAutomaticHTTPSRewritesEditable = true
	SettingListResponseZonesAutomaticHTTPSRewritesEditableFalse SettingListResponseZonesAutomaticHTTPSRewritesEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type SettingListResponseZonesBrotli struct {
	// ID of the zone setting.
	ID SettingListResponseZonesBrotliID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesBrotliValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesBrotliJSON `json:"-"`
}

// settingListResponseZonesBrotliJSON contains the JSON metadata for the struct
// [SettingListResponseZonesBrotli]
type settingListResponseZonesBrotliJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesBrotli) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesBrotliID string

const (
	SettingListResponseZonesBrotliIDBrotli SettingListResponseZonesBrotliID = "brotli"
)

// Current value of the zone setting.
type SettingListResponseZonesBrotliValue string

const (
	SettingListResponseZonesBrotliValueOff SettingListResponseZonesBrotliValue = "off"
	SettingListResponseZonesBrotliValueOn  SettingListResponseZonesBrotliValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesBrotliEditable bool

const (
	SettingListResponseZonesBrotliEditableTrue  SettingListResponseZonesBrotliEditable = true
	SettingListResponseZonesBrotliEditableFalse SettingListResponseZonesBrotliEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingListResponseZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID SettingListResponseZonesBrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesBrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesBrowserCacheTTLJSON `json:"-"`
}

// settingListResponseZonesBrowserCacheTTLJSON contains the JSON metadata for the
// struct [SettingListResponseZonesBrowserCacheTTL]
type settingListResponseZonesBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesBrowserCacheTTL) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesBrowserCacheTTLID string

const (
	SettingListResponseZonesBrowserCacheTTLIDBrowserCacheTTL SettingListResponseZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Current value of the zone setting.
type SettingListResponseZonesBrowserCacheTTLValue float64

const (
	SettingListResponseZonesBrowserCacheTTLValue0        SettingListResponseZonesBrowserCacheTTLValue = 0
	SettingListResponseZonesBrowserCacheTTLValue30       SettingListResponseZonesBrowserCacheTTLValue = 30
	SettingListResponseZonesBrowserCacheTTLValue60       SettingListResponseZonesBrowserCacheTTLValue = 60
	SettingListResponseZonesBrowserCacheTTLValue120      SettingListResponseZonesBrowserCacheTTLValue = 120
	SettingListResponseZonesBrowserCacheTTLValue300      SettingListResponseZonesBrowserCacheTTLValue = 300
	SettingListResponseZonesBrowserCacheTTLValue1200     SettingListResponseZonesBrowserCacheTTLValue = 1200
	SettingListResponseZonesBrowserCacheTTLValue1800     SettingListResponseZonesBrowserCacheTTLValue = 1800
	SettingListResponseZonesBrowserCacheTTLValue3600     SettingListResponseZonesBrowserCacheTTLValue = 3600
	SettingListResponseZonesBrowserCacheTTLValue7200     SettingListResponseZonesBrowserCacheTTLValue = 7200
	SettingListResponseZonesBrowserCacheTTLValue10800    SettingListResponseZonesBrowserCacheTTLValue = 10800
	SettingListResponseZonesBrowserCacheTTLValue14400    SettingListResponseZonesBrowserCacheTTLValue = 14400
	SettingListResponseZonesBrowserCacheTTLValue18000    SettingListResponseZonesBrowserCacheTTLValue = 18000
	SettingListResponseZonesBrowserCacheTTLValue28800    SettingListResponseZonesBrowserCacheTTLValue = 28800
	SettingListResponseZonesBrowserCacheTTLValue43200    SettingListResponseZonesBrowserCacheTTLValue = 43200
	SettingListResponseZonesBrowserCacheTTLValue57600    SettingListResponseZonesBrowserCacheTTLValue = 57600
	SettingListResponseZonesBrowserCacheTTLValue72000    SettingListResponseZonesBrowserCacheTTLValue = 72000
	SettingListResponseZonesBrowserCacheTTLValue86400    SettingListResponseZonesBrowserCacheTTLValue = 86400
	SettingListResponseZonesBrowserCacheTTLValue172800   SettingListResponseZonesBrowserCacheTTLValue = 172800
	SettingListResponseZonesBrowserCacheTTLValue259200   SettingListResponseZonesBrowserCacheTTLValue = 259200
	SettingListResponseZonesBrowserCacheTTLValue345600   SettingListResponseZonesBrowserCacheTTLValue = 345600
	SettingListResponseZonesBrowserCacheTTLValue432000   SettingListResponseZonesBrowserCacheTTLValue = 432000
	SettingListResponseZonesBrowserCacheTTLValue691200   SettingListResponseZonesBrowserCacheTTLValue = 691200
	SettingListResponseZonesBrowserCacheTTLValue1382400  SettingListResponseZonesBrowserCacheTTLValue = 1382400
	SettingListResponseZonesBrowserCacheTTLValue2073600  SettingListResponseZonesBrowserCacheTTLValue = 2073600
	SettingListResponseZonesBrowserCacheTTLValue2678400  SettingListResponseZonesBrowserCacheTTLValue = 2678400
	SettingListResponseZonesBrowserCacheTTLValue5356800  SettingListResponseZonesBrowserCacheTTLValue = 5356800
	SettingListResponseZonesBrowserCacheTTLValue16070400 SettingListResponseZonesBrowserCacheTTLValue = 16070400
	SettingListResponseZonesBrowserCacheTTLValue31536000 SettingListResponseZonesBrowserCacheTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesBrowserCacheTTLEditable bool

const (
	SettingListResponseZonesBrowserCacheTTLEditableTrue  SettingListResponseZonesBrowserCacheTTLEditable = true
	SettingListResponseZonesBrowserCacheTTLEditableFalse SettingListResponseZonesBrowserCacheTTLEditable = false
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type SettingListResponseZonesBrowserCheck struct {
	// ID of the zone setting.
	ID SettingListResponseZonesBrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesBrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesBrowserCheckJSON `json:"-"`
}

// settingListResponseZonesBrowserCheckJSON contains the JSON metadata for the
// struct [SettingListResponseZonesBrowserCheck]
type settingListResponseZonesBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesBrowserCheck) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesBrowserCheckID string

const (
	SettingListResponseZonesBrowserCheckIDBrowserCheck SettingListResponseZonesBrowserCheckID = "browser_check"
)

// Current value of the zone setting.
type SettingListResponseZonesBrowserCheckValue string

const (
	SettingListResponseZonesBrowserCheckValueOn  SettingListResponseZonesBrowserCheckValue = "on"
	SettingListResponseZonesBrowserCheckValueOff SettingListResponseZonesBrowserCheckValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesBrowserCheckEditable bool

const (
	SettingListResponseZonesBrowserCheckEditableTrue  SettingListResponseZonesBrowserCheckEditable = true
	SettingListResponseZonesBrowserCheckEditableFalse SettingListResponseZonesBrowserCheckEditable = false
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type SettingListResponseZonesCacheLevel struct {
	// ID of the zone setting.
	ID SettingListResponseZonesCacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesCacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesCacheLevelJSON `json:"-"`
}

// settingListResponseZonesCacheLevelJSON contains the JSON metadata for the struct
// [SettingListResponseZonesCacheLevel]
type settingListResponseZonesCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesCacheLevel) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesCacheLevelID string

const (
	SettingListResponseZonesCacheLevelIDCacheLevel SettingListResponseZonesCacheLevelID = "cache_level"
)

// Current value of the zone setting.
type SettingListResponseZonesCacheLevelValue string

const (
	SettingListResponseZonesCacheLevelValueAggressive SettingListResponseZonesCacheLevelValue = "aggressive"
	SettingListResponseZonesCacheLevelValueBasic      SettingListResponseZonesCacheLevelValue = "basic"
	SettingListResponseZonesCacheLevelValueSimplified SettingListResponseZonesCacheLevelValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesCacheLevelEditable bool

const (
	SettingListResponseZonesCacheLevelEditableTrue  SettingListResponseZonesCacheLevelEditable = true
	SettingListResponseZonesCacheLevelEditableFalse SettingListResponseZonesCacheLevelEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type SettingListResponseZonesChallengeTTL struct {
	// ID of the zone setting.
	ID SettingListResponseZonesChallengeTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesChallengeTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesChallengeTTLJSON `json:"-"`
}

// settingListResponseZonesChallengeTTLJSON contains the JSON metadata for the
// struct [SettingListResponseZonesChallengeTTL]
type settingListResponseZonesChallengeTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesChallengeTTL) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesChallengeTTLID string

const (
	SettingListResponseZonesChallengeTTLIDChallengeTTL SettingListResponseZonesChallengeTTLID = "challenge_ttl"
)

// Current value of the zone setting.
type SettingListResponseZonesChallengeTTLValue float64

const (
	SettingListResponseZonesChallengeTTLValue300      SettingListResponseZonesChallengeTTLValue = 300
	SettingListResponseZonesChallengeTTLValue900      SettingListResponseZonesChallengeTTLValue = 900
	SettingListResponseZonesChallengeTTLValue1800     SettingListResponseZonesChallengeTTLValue = 1800
	SettingListResponseZonesChallengeTTLValue2700     SettingListResponseZonesChallengeTTLValue = 2700
	SettingListResponseZonesChallengeTTLValue3600     SettingListResponseZonesChallengeTTLValue = 3600
	SettingListResponseZonesChallengeTTLValue7200     SettingListResponseZonesChallengeTTLValue = 7200
	SettingListResponseZonesChallengeTTLValue10800    SettingListResponseZonesChallengeTTLValue = 10800
	SettingListResponseZonesChallengeTTLValue14400    SettingListResponseZonesChallengeTTLValue = 14400
	SettingListResponseZonesChallengeTTLValue28800    SettingListResponseZonesChallengeTTLValue = 28800
	SettingListResponseZonesChallengeTTLValue57600    SettingListResponseZonesChallengeTTLValue = 57600
	SettingListResponseZonesChallengeTTLValue86400    SettingListResponseZonesChallengeTTLValue = 86400
	SettingListResponseZonesChallengeTTLValue604800   SettingListResponseZonesChallengeTTLValue = 604800
	SettingListResponseZonesChallengeTTLValue2592000  SettingListResponseZonesChallengeTTLValue = 2592000
	SettingListResponseZonesChallengeTTLValue31536000 SettingListResponseZonesChallengeTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesChallengeTTLEditable bool

const (
	SettingListResponseZonesChallengeTTLEditableTrue  SettingListResponseZonesChallengeTTLEditable = true
	SettingListResponseZonesChallengeTTLEditableFalse SettingListResponseZonesChallengeTTLEditable = false
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type SettingListResponseZonesCiphers struct {
	// ID of the zone setting.
	ID SettingListResponseZonesCiphersID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesCiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesCiphersJSON `json:"-"`
}

// settingListResponseZonesCiphersJSON contains the JSON metadata for the struct
// [SettingListResponseZonesCiphers]
type settingListResponseZonesCiphersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesCiphers) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesCiphersID string

const (
	SettingListResponseZonesCiphersIDCiphers SettingListResponseZonesCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesCiphersEditable bool

const (
	SettingListResponseZonesCiphersEditableTrue  SettingListResponseZonesCiphersEditable = true
	SettingListResponseZonesCiphersEditableFalse SettingListResponseZonesCiphersEditable = false
)

// Whether or not cname flattening is on.
type SettingListResponseZonesCnameFlattening struct {
	// How to flatten the cname destination.
	ID SettingListResponseZonesCnameFlatteningID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesCnameFlatteningValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesCnameFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesCnameFlatteningJSON `json:"-"`
}

// settingListResponseZonesCnameFlatteningJSON contains the JSON metadata for the
// struct [SettingListResponseZonesCnameFlattening]
type settingListResponseZonesCnameFlatteningJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesCnameFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesCnameFlattening) implementsSettingListResponse() {}

// How to flatten the cname destination.
type SettingListResponseZonesCnameFlatteningID string

const (
	SettingListResponseZonesCnameFlatteningIDCnameFlattening SettingListResponseZonesCnameFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type SettingListResponseZonesCnameFlatteningValue string

const (
	SettingListResponseZonesCnameFlatteningValueFlattenAtRoot SettingListResponseZonesCnameFlatteningValue = "flatten_at_root"
	SettingListResponseZonesCnameFlatteningValueFlattenAll    SettingListResponseZonesCnameFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesCnameFlatteningEditable bool

const (
	SettingListResponseZonesCnameFlatteningEditableTrue  SettingListResponseZonesCnameFlatteningEditable = true
	SettingListResponseZonesCnameFlatteningEditableFalse SettingListResponseZonesCnameFlatteningEditable = false
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type SettingListResponseZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID SettingListResponseZonesDevelopmentModeID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesDevelopmentModeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                                     `json:"time_remaining"`
	JSON          settingListResponseZonesDevelopmentModeJSON `json:"-"`
}

// settingListResponseZonesDevelopmentModeJSON contains the JSON metadata for the
// struct [SettingListResponseZonesDevelopmentMode]
type settingListResponseZonesDevelopmentModeJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingListResponseZonesDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesDevelopmentMode) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesDevelopmentModeID string

const (
	SettingListResponseZonesDevelopmentModeIDDevelopmentMode SettingListResponseZonesDevelopmentModeID = "development_mode"
)

// Current value of the zone setting.
type SettingListResponseZonesDevelopmentModeValue string

const (
	SettingListResponseZonesDevelopmentModeValueOn  SettingListResponseZonesDevelopmentModeValue = "on"
	SettingListResponseZonesDevelopmentModeValueOff SettingListResponseZonesDevelopmentModeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesDevelopmentModeEditable bool

const (
	SettingListResponseZonesDevelopmentModeEditableTrue  SettingListResponseZonesDevelopmentModeEditable = true
	SettingListResponseZonesDevelopmentModeEditableFalse SettingListResponseZonesDevelopmentModeEditable = false
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type SettingListResponseZonesEarlyHints struct {
	// ID of the zone setting.
	ID SettingListResponseZonesEarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesEarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesEarlyHintsJSON `json:"-"`
}

// settingListResponseZonesEarlyHintsJSON contains the JSON metadata for the struct
// [SettingListResponseZonesEarlyHints]
type settingListResponseZonesEarlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesEarlyHints) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesEarlyHintsID string

const (
	SettingListResponseZonesEarlyHintsIDEarlyHints SettingListResponseZonesEarlyHintsID = "early_hints"
)

// Current value of the zone setting.
type SettingListResponseZonesEarlyHintsValue string

const (
	SettingListResponseZonesEarlyHintsValueOn  SettingListResponseZonesEarlyHintsValue = "on"
	SettingListResponseZonesEarlyHintsValueOff SettingListResponseZonesEarlyHintsValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesEarlyHintsEditable bool

const (
	SettingListResponseZonesEarlyHintsEditableTrue  SettingListResponseZonesEarlyHintsEditable = true
	SettingListResponseZonesEarlyHintsEditableFalse SettingListResponseZonesEarlyHintsEditable = false
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type SettingListResponseZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID SettingListResponseZonesEdgeCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesEdgeCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesEdgeCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesEdgeCacheTTLJSON `json:"-"`
}

// settingListResponseZonesEdgeCacheTTLJSON contains the JSON metadata for the
// struct [SettingListResponseZonesEdgeCacheTTL]
type settingListResponseZonesEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesEdgeCacheTTL) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesEdgeCacheTTLID string

const (
	SettingListResponseZonesEdgeCacheTTLIDEdgeCacheTTL SettingListResponseZonesEdgeCacheTTLID = "edge_cache_ttl"
)

// Current value of the zone setting.
type SettingListResponseZonesEdgeCacheTTLValue float64

const (
	SettingListResponseZonesEdgeCacheTTLValue30     SettingListResponseZonesEdgeCacheTTLValue = 30
	SettingListResponseZonesEdgeCacheTTLValue60     SettingListResponseZonesEdgeCacheTTLValue = 60
	SettingListResponseZonesEdgeCacheTTLValue300    SettingListResponseZonesEdgeCacheTTLValue = 300
	SettingListResponseZonesEdgeCacheTTLValue1200   SettingListResponseZonesEdgeCacheTTLValue = 1200
	SettingListResponseZonesEdgeCacheTTLValue1800   SettingListResponseZonesEdgeCacheTTLValue = 1800
	SettingListResponseZonesEdgeCacheTTLValue3600   SettingListResponseZonesEdgeCacheTTLValue = 3600
	SettingListResponseZonesEdgeCacheTTLValue7200   SettingListResponseZonesEdgeCacheTTLValue = 7200
	SettingListResponseZonesEdgeCacheTTLValue10800  SettingListResponseZonesEdgeCacheTTLValue = 10800
	SettingListResponseZonesEdgeCacheTTLValue14400  SettingListResponseZonesEdgeCacheTTLValue = 14400
	SettingListResponseZonesEdgeCacheTTLValue18000  SettingListResponseZonesEdgeCacheTTLValue = 18000
	SettingListResponseZonesEdgeCacheTTLValue28800  SettingListResponseZonesEdgeCacheTTLValue = 28800
	SettingListResponseZonesEdgeCacheTTLValue43200  SettingListResponseZonesEdgeCacheTTLValue = 43200
	SettingListResponseZonesEdgeCacheTTLValue57600  SettingListResponseZonesEdgeCacheTTLValue = 57600
	SettingListResponseZonesEdgeCacheTTLValue72000  SettingListResponseZonesEdgeCacheTTLValue = 72000
	SettingListResponseZonesEdgeCacheTTLValue86400  SettingListResponseZonesEdgeCacheTTLValue = 86400
	SettingListResponseZonesEdgeCacheTTLValue172800 SettingListResponseZonesEdgeCacheTTLValue = 172800
	SettingListResponseZonesEdgeCacheTTLValue259200 SettingListResponseZonesEdgeCacheTTLValue = 259200
	SettingListResponseZonesEdgeCacheTTLValue345600 SettingListResponseZonesEdgeCacheTTLValue = 345600
	SettingListResponseZonesEdgeCacheTTLValue432000 SettingListResponseZonesEdgeCacheTTLValue = 432000
	SettingListResponseZonesEdgeCacheTTLValue518400 SettingListResponseZonesEdgeCacheTTLValue = 518400
	SettingListResponseZonesEdgeCacheTTLValue604800 SettingListResponseZonesEdgeCacheTTLValue = 604800
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesEdgeCacheTTLEditable bool

const (
	SettingListResponseZonesEdgeCacheTTLEditableTrue  SettingListResponseZonesEdgeCacheTTLEditable = true
	SettingListResponseZonesEdgeCacheTTLEditableFalse SettingListResponseZonesEdgeCacheTTLEditable = false
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingListResponseZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID SettingListResponseZonesEmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesEmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesEmailObfuscationJSON `json:"-"`
}

// settingListResponseZonesEmailObfuscationJSON contains the JSON metadata for the
// struct [SettingListResponseZonesEmailObfuscation]
type settingListResponseZonesEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesEmailObfuscation) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesEmailObfuscationID string

const (
	SettingListResponseZonesEmailObfuscationIDEmailObfuscation SettingListResponseZonesEmailObfuscationID = "email_obfuscation"
)

// Current value of the zone setting.
type SettingListResponseZonesEmailObfuscationValue string

const (
	SettingListResponseZonesEmailObfuscationValueOn  SettingListResponseZonesEmailObfuscationValue = "on"
	SettingListResponseZonesEmailObfuscationValueOff SettingListResponseZonesEmailObfuscationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesEmailObfuscationEditable bool

const (
	SettingListResponseZonesEmailObfuscationEditableTrue  SettingListResponseZonesEmailObfuscationEditable = true
	SettingListResponseZonesEmailObfuscationEditableFalse SettingListResponseZonesEmailObfuscationEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type SettingListResponseZonesH2Prioritization struct {
	// ID of the zone setting.
	ID SettingListResponseZonesH2PrioritizationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesH2PrioritizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesH2PrioritizationJSON `json:"-"`
}

// settingListResponseZonesH2PrioritizationJSON contains the JSON metadata for the
// struct [SettingListResponseZonesH2Prioritization]
type settingListResponseZonesH2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesH2Prioritization) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesH2PrioritizationID string

const (
	SettingListResponseZonesH2PrioritizationIDH2Prioritization SettingListResponseZonesH2PrioritizationID = "h2_prioritization"
)

// Current value of the zone setting.
type SettingListResponseZonesH2PrioritizationValue string

const (
	SettingListResponseZonesH2PrioritizationValueOn     SettingListResponseZonesH2PrioritizationValue = "on"
	SettingListResponseZonesH2PrioritizationValueOff    SettingListResponseZonesH2PrioritizationValue = "off"
	SettingListResponseZonesH2PrioritizationValueCustom SettingListResponseZonesH2PrioritizationValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesH2PrioritizationEditable bool

const (
	SettingListResponseZonesH2PrioritizationEditableTrue  SettingListResponseZonesH2PrioritizationEditable = true
	SettingListResponseZonesH2PrioritizationEditableFalse SettingListResponseZonesH2PrioritizationEditable = false
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type SettingListResponseZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID SettingListResponseZonesHotlinkProtectionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesHotlinkProtectionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesHotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesHotlinkProtectionJSON `json:"-"`
}

// settingListResponseZonesHotlinkProtectionJSON contains the JSON metadata for the
// struct [SettingListResponseZonesHotlinkProtection]
type settingListResponseZonesHotlinkProtectionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesHotlinkProtection) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesHotlinkProtectionID string

const (
	SettingListResponseZonesHotlinkProtectionIDHotlinkProtection SettingListResponseZonesHotlinkProtectionID = "hotlink_protection"
)

// Current value of the zone setting.
type SettingListResponseZonesHotlinkProtectionValue string

const (
	SettingListResponseZonesHotlinkProtectionValueOn  SettingListResponseZonesHotlinkProtectionValue = "on"
	SettingListResponseZonesHotlinkProtectionValueOff SettingListResponseZonesHotlinkProtectionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesHotlinkProtectionEditable bool

const (
	SettingListResponseZonesHotlinkProtectionEditableTrue  SettingListResponseZonesHotlinkProtectionEditable = true
	SettingListResponseZonesHotlinkProtectionEditableFalse SettingListResponseZonesHotlinkProtectionEditable = false
)

// HTTP2 enabled for this zone.
type SettingListResponseZonesHTTP2 struct {
	// ID of the zone setting.
	ID SettingListResponseZonesHTTP2ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesHTTP2Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesHTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesHTTP2JSON `json:"-"`
}

// settingListResponseZonesHTTP2JSON contains the JSON metadata for the struct
// [SettingListResponseZonesHTTP2]
type settingListResponseZonesHTTP2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesHTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesHTTP2) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesHTTP2ID string

const (
	SettingListResponseZonesHTTP2IDHTTP2 SettingListResponseZonesHTTP2ID = "http2"
)

// Current value of the zone setting.
type SettingListResponseZonesHTTP2Value string

const (
	SettingListResponseZonesHTTP2ValueOn  SettingListResponseZonesHTTP2Value = "on"
	SettingListResponseZonesHTTP2ValueOff SettingListResponseZonesHTTP2Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesHTTP2Editable bool

const (
	SettingListResponseZonesHTTP2EditableTrue  SettingListResponseZonesHTTP2Editable = true
	SettingListResponseZonesHTTP2EditableFalse SettingListResponseZonesHTTP2Editable = false
)

// HTTP3 enabled for this zone.
type SettingListResponseZonesHTTP3 struct {
	// ID of the zone setting.
	ID SettingListResponseZonesHTTP3ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesHTTP3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesHTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesHTTP3JSON `json:"-"`
}

// settingListResponseZonesHTTP3JSON contains the JSON metadata for the struct
// [SettingListResponseZonesHTTP3]
type settingListResponseZonesHTTP3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesHTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesHTTP3) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesHTTP3ID string

const (
	SettingListResponseZonesHTTP3IDHTTP3 SettingListResponseZonesHTTP3ID = "http3"
)

// Current value of the zone setting.
type SettingListResponseZonesHTTP3Value string

const (
	SettingListResponseZonesHTTP3ValueOn  SettingListResponseZonesHTTP3Value = "on"
	SettingListResponseZonesHTTP3ValueOff SettingListResponseZonesHTTP3Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesHTTP3Editable bool

const (
	SettingListResponseZonesHTTP3EditableTrue  SettingListResponseZonesHTTP3Editable = true
	SettingListResponseZonesHTTP3EditableFalse SettingListResponseZonesHTTP3Editable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type SettingListResponseZonesImageResizing struct {
	// ID of the zone setting.
	ID SettingListResponseZonesImageResizingID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesImageResizingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesImageResizingJSON `json:"-"`
}

// settingListResponseZonesImageResizingJSON contains the JSON metadata for the
// struct [SettingListResponseZonesImageResizing]
type settingListResponseZonesImageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesImageResizing) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesImageResizingID string

const (
	SettingListResponseZonesImageResizingIDImageResizing SettingListResponseZonesImageResizingID = "image_resizing"
)

// Current value of the zone setting.
type SettingListResponseZonesImageResizingValue string

const (
	SettingListResponseZonesImageResizingValueOn   SettingListResponseZonesImageResizingValue = "on"
	SettingListResponseZonesImageResizingValueOff  SettingListResponseZonesImageResizingValue = "off"
	SettingListResponseZonesImageResizingValueOpen SettingListResponseZonesImageResizingValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesImageResizingEditable bool

const (
	SettingListResponseZonesImageResizingEditableTrue  SettingListResponseZonesImageResizingEditable = true
	SettingListResponseZonesImageResizingEditableFalse SettingListResponseZonesImageResizingEditable = false
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingListResponseZonesIPGeolocation struct {
	// ID of the zone setting.
	ID SettingListResponseZonesIPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesIPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesIPGeolocationJSON `json:"-"`
}

// settingListResponseZonesIPGeolocationJSON contains the JSON metadata for the
// struct [SettingListResponseZonesIPGeolocation]
type settingListResponseZonesIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesIPGeolocation) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesIPGeolocationID string

const (
	SettingListResponseZonesIPGeolocationIDIPGeolocation SettingListResponseZonesIPGeolocationID = "ip_geolocation"
)

// Current value of the zone setting.
type SettingListResponseZonesIPGeolocationValue string

const (
	SettingListResponseZonesIPGeolocationValueOn  SettingListResponseZonesIPGeolocationValue = "on"
	SettingListResponseZonesIPGeolocationValueOff SettingListResponseZonesIPGeolocationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesIPGeolocationEditable bool

const (
	SettingListResponseZonesIPGeolocationEditableTrue  SettingListResponseZonesIPGeolocationEditable = true
	SettingListResponseZonesIPGeolocationEditableFalse SettingListResponseZonesIPGeolocationEditable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type SettingListResponseZonesIPV6 struct {
	// ID of the zone setting.
	ID SettingListResponseZonesIPV6ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesIPV6Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesIPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesIPV6JSON `json:"-"`
}

// settingListResponseZonesIPV6JSON contains the JSON metadata for the struct
// [SettingListResponseZonesIPV6]
type settingListResponseZonesIPV6JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesIPV6) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesIPV6ID string

const (
	SettingListResponseZonesIPV6IDIPV6 SettingListResponseZonesIPV6ID = "ipv6"
)

// Current value of the zone setting.
type SettingListResponseZonesIPV6Value string

const (
	SettingListResponseZonesIPV6ValueOff SettingListResponseZonesIPV6Value = "off"
	SettingListResponseZonesIPV6ValueOn  SettingListResponseZonesIPV6Value = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesIPV6Editable bool

const (
	SettingListResponseZonesIPV6EditableTrue  SettingListResponseZonesIPV6Editable = true
	SettingListResponseZonesIPV6EditableFalse SettingListResponseZonesIPV6Editable = false
)

// Maximum size of an allowable upload.
type SettingListResponseZonesMaxUpload struct {
	// identifier of the zone setting.
	ID SettingListResponseZonesMaxUploadID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesMaxUploadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesMaxUploadJSON `json:"-"`
}

// settingListResponseZonesMaxUploadJSON contains the JSON metadata for the struct
// [SettingListResponseZonesMaxUpload]
type settingListResponseZonesMaxUploadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesMaxUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesMaxUpload) implementsSettingListResponse() {}

// identifier of the zone setting.
type SettingListResponseZonesMaxUploadID string

const (
	SettingListResponseZonesMaxUploadIDMaxUpload SettingListResponseZonesMaxUploadID = "max_upload"
)

// Current value of the zone setting.
type SettingListResponseZonesMaxUploadValue float64

const (
	SettingListResponseZonesMaxUploadValue100 SettingListResponseZonesMaxUploadValue = 100
	SettingListResponseZonesMaxUploadValue200 SettingListResponseZonesMaxUploadValue = 200
	SettingListResponseZonesMaxUploadValue500 SettingListResponseZonesMaxUploadValue = 500
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesMaxUploadEditable bool

const (
	SettingListResponseZonesMaxUploadEditableTrue  SettingListResponseZonesMaxUploadEditable = true
	SettingListResponseZonesMaxUploadEditableFalse SettingListResponseZonesMaxUploadEditable = false
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type SettingListResponseZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID SettingListResponseZonesMinTLSVersionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesMinTLSVersionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesMinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesMinTLSVersionJSON `json:"-"`
}

// settingListResponseZonesMinTLSVersionJSON contains the JSON metadata for the
// struct [SettingListResponseZonesMinTLSVersion]
type settingListResponseZonesMinTLSVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesMinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesMinTLSVersion) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesMinTLSVersionID string

const (
	SettingListResponseZonesMinTLSVersionIDMinTLSVersion SettingListResponseZonesMinTLSVersionID = "min_tls_version"
)

// Current value of the zone setting.
type SettingListResponseZonesMinTLSVersionValue string

const (
	SettingListResponseZonesMinTLSVersionValue1_0 SettingListResponseZonesMinTLSVersionValue = "1.0"
	SettingListResponseZonesMinTLSVersionValue1_1 SettingListResponseZonesMinTLSVersionValue = "1.1"
	SettingListResponseZonesMinTLSVersionValue1_2 SettingListResponseZonesMinTLSVersionValue = "1.2"
	SettingListResponseZonesMinTLSVersionValue1_3 SettingListResponseZonesMinTLSVersionValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesMinTLSVersionEditable bool

const (
	SettingListResponseZonesMinTLSVersionEditableTrue  SettingListResponseZonesMinTLSVersionEditable = true
	SettingListResponseZonesMinTLSVersionEditableFalse SettingListResponseZonesMinTLSVersionEditable = false
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type SettingListResponseZonesMinify struct {
	// Zone setting identifier.
	ID SettingListResponseZonesMinifyID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesMinifyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesMinifyJSON `json:"-"`
}

// settingListResponseZonesMinifyJSON contains the JSON metadata for the struct
// [SettingListResponseZonesMinify]
type settingListResponseZonesMinifyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesMinify) implementsSettingListResponse() {}

// Zone setting identifier.
type SettingListResponseZonesMinifyID string

const (
	SettingListResponseZonesMinifyIDMinify SettingListResponseZonesMinifyID = "minify"
)

// Current value of the zone setting.
type SettingListResponseZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css SettingListResponseZonesMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML SettingListResponseZonesMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   SettingListResponseZonesMinifyValueJs   `json:"js"`
	JSON settingListResponseZonesMinifyValueJSON `json:"-"`
}

// settingListResponseZonesMinifyValueJSON contains the JSON metadata for the
// struct [SettingListResponseZonesMinifyValue]
type settingListResponseZonesMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type SettingListResponseZonesMinifyValueCss string

const (
	SettingListResponseZonesMinifyValueCssOn  SettingListResponseZonesMinifyValueCss = "on"
	SettingListResponseZonesMinifyValueCssOff SettingListResponseZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type SettingListResponseZonesMinifyValueHTML string

const (
	SettingListResponseZonesMinifyValueHTMLOn  SettingListResponseZonesMinifyValueHTML = "on"
	SettingListResponseZonesMinifyValueHTMLOff SettingListResponseZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type SettingListResponseZonesMinifyValueJs string

const (
	SettingListResponseZonesMinifyValueJsOn  SettingListResponseZonesMinifyValueJs = "on"
	SettingListResponseZonesMinifyValueJsOff SettingListResponseZonesMinifyValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesMinifyEditable bool

const (
	SettingListResponseZonesMinifyEditableTrue  SettingListResponseZonesMinifyEditable = true
	SettingListResponseZonesMinifyEditableFalse SettingListResponseZonesMinifyEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type SettingListResponseZonesMirage struct {
	// ID of the zone setting.
	ID SettingListResponseZonesMirageID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesMirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesMirageJSON `json:"-"`
}

// settingListResponseZonesMirageJSON contains the JSON metadata for the struct
// [SettingListResponseZonesMirage]
type settingListResponseZonesMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesMirage) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesMirageID string

const (
	SettingListResponseZonesMirageIDMirage SettingListResponseZonesMirageID = "mirage"
)

// Current value of the zone setting.
type SettingListResponseZonesMirageValue string

const (
	SettingListResponseZonesMirageValueOn  SettingListResponseZonesMirageValue = "on"
	SettingListResponseZonesMirageValueOff SettingListResponseZonesMirageValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesMirageEditable bool

const (
	SettingListResponseZonesMirageEditableTrue  SettingListResponseZonesMirageEditable = true
	SettingListResponseZonesMirageEditableFalse SettingListResponseZonesMirageEditable = false
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type SettingListResponseZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID SettingListResponseZonesMobileRedirectID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesMobileRedirectValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesMobileRedirectJSON `json:"-"`
}

// settingListResponseZonesMobileRedirectJSON contains the JSON metadata for the
// struct [SettingListResponseZonesMobileRedirect]
type settingListResponseZonesMobileRedirectJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesMobileRedirect) implementsSettingListResponse() {}

// Identifier of the zone setting.
type SettingListResponseZonesMobileRedirectID string

const (
	SettingListResponseZonesMobileRedirectIDMobileRedirect SettingListResponseZonesMobileRedirectID = "mobile_redirect"
)

// Current value of the zone setting.
type SettingListResponseZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status SettingListResponseZonesMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                            `json:"strip_uri"`
	JSON     settingListResponseZonesMobileRedirectValueJSON `json:"-"`
}

// settingListResponseZonesMobileRedirectValueJSON contains the JSON metadata for
// the struct [SettingListResponseZonesMobileRedirectValue]
type settingListResponseZonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingListResponseZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type SettingListResponseZonesMobileRedirectValueStatus string

const (
	SettingListResponseZonesMobileRedirectValueStatusOn  SettingListResponseZonesMobileRedirectValueStatus = "on"
	SettingListResponseZonesMobileRedirectValueStatusOff SettingListResponseZonesMobileRedirectValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesMobileRedirectEditable bool

const (
	SettingListResponseZonesMobileRedirectEditableTrue  SettingListResponseZonesMobileRedirectEditable = true
	SettingListResponseZonesMobileRedirectEditableFalse SettingListResponseZonesMobileRedirectEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type SettingListResponseZonesNEL struct {
	// Zone setting identifier.
	ID SettingListResponseZonesNELID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesNELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesNELJSON `json:"-"`
}

// settingListResponseZonesNELJSON contains the JSON metadata for the struct
// [SettingListResponseZonesNEL]
type settingListResponseZonesNELJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesNEL) implementsSettingListResponse() {}

// Zone setting identifier.
type SettingListResponseZonesNELID string

const (
	SettingListResponseZonesNELIDNEL SettingListResponseZonesNELID = "nel"
)

// Current value of the zone setting.
type SettingListResponseZonesNELValue struct {
	Enabled bool                                 `json:"enabled"`
	JSON    settingListResponseZonesNELValueJSON `json:"-"`
}

// settingListResponseZonesNELValueJSON contains the JSON metadata for the struct
// [SettingListResponseZonesNELValue]
type settingListResponseZonesNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesNELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesNELEditable bool

const (
	SettingListResponseZonesNELEditableTrue  SettingListResponseZonesNELEditable = true
	SettingListResponseZonesNELEditableFalse SettingListResponseZonesNELEditable = false
)

// Enables the Opportunistic Encryption feature for a zone.
type SettingListResponseZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID SettingListResponseZonesOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesOpportunisticEncryptionJSON `json:"-"`
}

// settingListResponseZonesOpportunisticEncryptionJSON contains the JSON metadata
// for the struct [SettingListResponseZonesOpportunisticEncryption]
type settingListResponseZonesOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesOpportunisticEncryption) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesOpportunisticEncryptionID string

const (
	SettingListResponseZonesOpportunisticEncryptionIDOpportunisticEncryption SettingListResponseZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Current value of the zone setting.
type SettingListResponseZonesOpportunisticEncryptionValue string

const (
	SettingListResponseZonesOpportunisticEncryptionValueOn  SettingListResponseZonesOpportunisticEncryptionValue = "on"
	SettingListResponseZonesOpportunisticEncryptionValueOff SettingListResponseZonesOpportunisticEncryptionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesOpportunisticEncryptionEditable bool

const (
	SettingListResponseZonesOpportunisticEncryptionEditableTrue  SettingListResponseZonesOpportunisticEncryptionEditable = true
	SettingListResponseZonesOpportunisticEncryptionEditableFalse SettingListResponseZonesOpportunisticEncryptionEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type SettingListResponseZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID SettingListResponseZonesOpportunisticOnionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesOpportunisticOnionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesOpportunisticOnionJSON `json:"-"`
}

// settingListResponseZonesOpportunisticOnionJSON contains the JSON metadata for
// the struct [SettingListResponseZonesOpportunisticOnion]
type settingListResponseZonesOpportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesOpportunisticOnion) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesOpportunisticOnionID string

const (
	SettingListResponseZonesOpportunisticOnionIDOpportunisticOnion SettingListResponseZonesOpportunisticOnionID = "opportunistic_onion"
)

// Current value of the zone setting.
type SettingListResponseZonesOpportunisticOnionValue string

const (
	SettingListResponseZonesOpportunisticOnionValueOn  SettingListResponseZonesOpportunisticOnionValue = "on"
	SettingListResponseZonesOpportunisticOnionValueOff SettingListResponseZonesOpportunisticOnionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesOpportunisticOnionEditable bool

const (
	SettingListResponseZonesOpportunisticOnionEditableTrue  SettingListResponseZonesOpportunisticOnionEditable = true
	SettingListResponseZonesOpportunisticOnionEditableFalse SettingListResponseZonesOpportunisticOnionEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type SettingListResponseZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID SettingListResponseZonesOrangeToOrangeID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesOrangeToOrangeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesOrangeToOrangeJSON `json:"-"`
}

// settingListResponseZonesOrangeToOrangeJSON contains the JSON metadata for the
// struct [SettingListResponseZonesOrangeToOrange]
type settingListResponseZonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesOrangeToOrange) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesOrangeToOrangeID string

const (
	SettingListResponseZonesOrangeToOrangeIDOrangeToOrange SettingListResponseZonesOrangeToOrangeID = "orange_to_orange"
)

// Current value of the zone setting.
type SettingListResponseZonesOrangeToOrangeValue string

const (
	SettingListResponseZonesOrangeToOrangeValueOn  SettingListResponseZonesOrangeToOrangeValue = "on"
	SettingListResponseZonesOrangeToOrangeValueOff SettingListResponseZonesOrangeToOrangeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesOrangeToOrangeEditable bool

const (
	SettingListResponseZonesOrangeToOrangeEditableTrue  SettingListResponseZonesOrangeToOrangeEditable = true
	SettingListResponseZonesOrangeToOrangeEditableFalse SettingListResponseZonesOrangeToOrangeEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingListResponseZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID SettingListResponseZonesOriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesOriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesOriginErrorPagePassThruJSON `json:"-"`
}

// settingListResponseZonesOriginErrorPagePassThruJSON contains the JSON metadata
// for the struct [SettingListResponseZonesOriginErrorPagePassThru]
type settingListResponseZonesOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesOriginErrorPagePassThru) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesOriginErrorPagePassThruID string

const (
	SettingListResponseZonesOriginErrorPagePassThruIDOriginErrorPagePassThru SettingListResponseZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type SettingListResponseZonesOriginErrorPagePassThruValue string

const (
	SettingListResponseZonesOriginErrorPagePassThruValueOn  SettingListResponseZonesOriginErrorPagePassThruValue = "on"
	SettingListResponseZonesOriginErrorPagePassThruValueOff SettingListResponseZonesOriginErrorPagePassThruValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesOriginErrorPagePassThruEditable bool

const (
	SettingListResponseZonesOriginErrorPagePassThruEditableTrue  SettingListResponseZonesOriginErrorPagePassThruEditable = true
	SettingListResponseZonesOriginErrorPagePassThruEditableFalse SettingListResponseZonesOriginErrorPagePassThruEditable = false
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type SettingListResponseZonesPolish struct {
	// ID of the zone setting.
	ID SettingListResponseZonesPolishID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesPolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesPolishJSON `json:"-"`
}

// settingListResponseZonesPolishJSON contains the JSON metadata for the struct
// [SettingListResponseZonesPolish]
type settingListResponseZonesPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesPolish) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesPolishID string

const (
	SettingListResponseZonesPolishIDPolish SettingListResponseZonesPolishID = "polish"
)

// Current value of the zone setting.
type SettingListResponseZonesPolishValue string

const (
	SettingListResponseZonesPolishValueOff      SettingListResponseZonesPolishValue = "off"
	SettingListResponseZonesPolishValueLossless SettingListResponseZonesPolishValue = "lossless"
	SettingListResponseZonesPolishValueLossy    SettingListResponseZonesPolishValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesPolishEditable bool

const (
	SettingListResponseZonesPolishEditableTrue  SettingListResponseZonesPolishEditable = true
	SettingListResponseZonesPolishEditableFalse SettingListResponseZonesPolishEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type SettingListResponseZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID SettingListResponseZonesPrefetchPreloadID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesPrefetchPreloadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesPrefetchPreloadJSON `json:"-"`
}

// settingListResponseZonesPrefetchPreloadJSON contains the JSON metadata for the
// struct [SettingListResponseZonesPrefetchPreload]
type settingListResponseZonesPrefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesPrefetchPreload) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesPrefetchPreloadID string

const (
	SettingListResponseZonesPrefetchPreloadIDPrefetchPreload SettingListResponseZonesPrefetchPreloadID = "prefetch_preload"
)

// Current value of the zone setting.
type SettingListResponseZonesPrefetchPreloadValue string

const (
	SettingListResponseZonesPrefetchPreloadValueOn  SettingListResponseZonesPrefetchPreloadValue = "on"
	SettingListResponseZonesPrefetchPreloadValueOff SettingListResponseZonesPrefetchPreloadValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesPrefetchPreloadEditable bool

const (
	SettingListResponseZonesPrefetchPreloadEditableTrue  SettingListResponseZonesPrefetchPreloadEditable = true
	SettingListResponseZonesPrefetchPreloadEditableFalse SettingListResponseZonesPrefetchPreloadEditable = false
)

// Maximum time between two read operations from origin.
type SettingListResponseZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID SettingListResponseZonesProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesProxyReadTimeoutJSON `json:"-"`
}

// settingListResponseZonesProxyReadTimeoutJSON contains the JSON metadata for the
// struct [SettingListResponseZonesProxyReadTimeout]
type settingListResponseZonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesProxyReadTimeout) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesProxyReadTimeoutID string

const (
	SettingListResponseZonesProxyReadTimeoutIDProxyReadTimeout SettingListResponseZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesProxyReadTimeoutEditable bool

const (
	SettingListResponseZonesProxyReadTimeoutEditableTrue  SettingListResponseZonesProxyReadTimeoutEditable = true
	SettingListResponseZonesProxyReadTimeoutEditableFalse SettingListResponseZonesProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type SettingListResponseZonesPseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID SettingListResponseZonesPseudoIPV4ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesPseudoIPV4Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesPseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesPseudoIPV4JSON `json:"-"`
}

// settingListResponseZonesPseudoIPV4JSON contains the JSON metadata for the struct
// [SettingListResponseZonesPseudoIPV4]
type settingListResponseZonesPseudoIPV4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesPseudoIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesPseudoIPV4) implementsSettingListResponse() {}

// Value of the Pseudo IPv4 setting.
type SettingListResponseZonesPseudoIPV4ID string

const (
	SettingListResponseZonesPseudoIPV4IDPseudoIPV4 SettingListResponseZonesPseudoIPV4ID = "pseudo_ipv4"
)

// Current value of the zone setting.
type SettingListResponseZonesPseudoIPV4Value string

const (
	SettingListResponseZonesPseudoIPV4ValueOff             SettingListResponseZonesPseudoIPV4Value = "off"
	SettingListResponseZonesPseudoIPV4ValueAddHeader       SettingListResponseZonesPseudoIPV4Value = "add_header"
	SettingListResponseZonesPseudoIPV4ValueOverwriteHeader SettingListResponseZonesPseudoIPV4Value = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesPseudoIPV4Editable bool

const (
	SettingListResponseZonesPseudoIPV4EditableTrue  SettingListResponseZonesPseudoIPV4Editable = true
	SettingListResponseZonesPseudoIPV4EditableFalse SettingListResponseZonesPseudoIPV4Editable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingListResponseZonesResponseBuffering struct {
	// ID of the zone setting.
	ID SettingListResponseZonesResponseBufferingID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesResponseBufferingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesResponseBufferingJSON `json:"-"`
}

// settingListResponseZonesResponseBufferingJSON contains the JSON metadata for the
// struct [SettingListResponseZonesResponseBuffering]
type settingListResponseZonesResponseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesResponseBuffering) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesResponseBufferingID string

const (
	SettingListResponseZonesResponseBufferingIDResponseBuffering SettingListResponseZonesResponseBufferingID = "response_buffering"
)

// Current value of the zone setting.
type SettingListResponseZonesResponseBufferingValue string

const (
	SettingListResponseZonesResponseBufferingValueOn  SettingListResponseZonesResponseBufferingValue = "on"
	SettingListResponseZonesResponseBufferingValueOff SettingListResponseZonesResponseBufferingValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesResponseBufferingEditable bool

const (
	SettingListResponseZonesResponseBufferingEditableTrue  SettingListResponseZonesResponseBufferingEditable = true
	SettingListResponseZonesResponseBufferingEditableFalse SettingListResponseZonesResponseBufferingEditable = false
)

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
type SettingListResponseZonesRocketLoader struct {
	// ID of the zone setting.
	ID SettingListResponseZonesRocketLoaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesRocketLoaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesRocketLoaderJSON `json:"-"`
}

// settingListResponseZonesRocketLoaderJSON contains the JSON metadata for the
// struct [SettingListResponseZonesRocketLoader]
type settingListResponseZonesRocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesRocketLoader) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesRocketLoaderID string

const (
	SettingListResponseZonesRocketLoaderIDRocketLoader SettingListResponseZonesRocketLoaderID = "rocket_loader"
)

// Current value of the zone setting.
type SettingListResponseZonesRocketLoaderValue string

const (
	SettingListResponseZonesRocketLoaderValueOn  SettingListResponseZonesRocketLoaderValue = "on"
	SettingListResponseZonesRocketLoaderValueOff SettingListResponseZonesRocketLoaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesRocketLoaderEditable bool

const (
	SettingListResponseZonesRocketLoaderEditableTrue  SettingListResponseZonesRocketLoaderEditable = true
	SettingListResponseZonesRocketLoaderEditableFalse SettingListResponseZonesRocketLoaderEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingListResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingListResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesSchemasAutomaticPlatformOptimizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesSchemasAutomaticPlatformOptimizationJSON `json:"-"`
}

// settingListResponseZonesSchemasAutomaticPlatformOptimizationJSON contains the
// JSON metadata for the struct
// [SettingListResponseZonesSchemasAutomaticPlatformOptimization]
type settingListResponseZonesSchemasAutomaticPlatformOptimizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesSchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesSchemasAutomaticPlatformOptimization) implementsSettingListResponse() {
}

// ID of the zone setting.
type SettingListResponseZonesSchemasAutomaticPlatformOptimizationID string

const (
	SettingListResponseZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization SettingListResponseZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Current value of the zone setting.
type SettingListResponseZonesSchemasAutomaticPlatformOptimizationValue struct {
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
	WpPlugin bool                                                                  `json:"wp_plugin,required"`
	JSON     settingListResponseZonesSchemasAutomaticPlatformOptimizationValueJSON `json:"-"`
}

// settingListResponseZonesSchemasAutomaticPlatformOptimizationValueJSON contains
// the JSON metadata for the struct
// [SettingListResponseZonesSchemasAutomaticPlatformOptimizationValue]
type settingListResponseZonesSchemasAutomaticPlatformOptimizationValueJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SettingListResponseZonesSchemasAutomaticPlatformOptimizationValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingListResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingListResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingListResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingListResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Cloudflare security header for a zone.
type SettingListResponseZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID SettingListResponseZonesSecurityHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesSecurityHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesSecurityHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesSecurityHeaderJSON `json:"-"`
}

// settingListResponseZonesSecurityHeaderJSON contains the JSON metadata for the
// struct [SettingListResponseZonesSecurityHeader]
type settingListResponseZonesSecurityHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesSecurityHeader) implementsSettingListResponse() {}

// ID of the zone's security header.
type SettingListResponseZonesSecurityHeaderID string

const (
	SettingListResponseZonesSecurityHeaderIDSecurityHeader SettingListResponseZonesSecurityHeaderID = "security_header"
)

// Current value of the zone setting.
type SettingListResponseZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity SettingListResponseZonesSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    settingListResponseZonesSecurityHeaderValueJSON                    `json:"-"`
}

// settingListResponseZonesSecurityHeaderValueJSON contains the JSON metadata for
// the struct [SettingListResponseZonesSecurityHeaderValue]
type settingListResponseZonesSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingListResponseZonesSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type SettingListResponseZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                   `json:"nosniff"`
	JSON    settingListResponseZonesSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// settingListResponseZonesSecurityHeaderValueStrictTransportSecurityJSON contains
// the JSON metadata for the struct
// [SettingListResponseZonesSecurityHeaderValueStrictTransportSecurity]
type settingListResponseZonesSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SettingListResponseZonesSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesSecurityHeaderEditable bool

const (
	SettingListResponseZonesSecurityHeaderEditableTrue  SettingListResponseZonesSecurityHeaderEditable = true
	SettingListResponseZonesSecurityHeaderEditableFalse SettingListResponseZonesSecurityHeaderEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingListResponseZonesSecurityLevel struct {
	// ID of the zone setting.
	ID SettingListResponseZonesSecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesSecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesSecurityLevelJSON `json:"-"`
}

// settingListResponseZonesSecurityLevelJSON contains the JSON metadata for the
// struct [SettingListResponseZonesSecurityLevel]
type settingListResponseZonesSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesSecurityLevel) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesSecurityLevelID string

const (
	SettingListResponseZonesSecurityLevelIDSecurityLevel SettingListResponseZonesSecurityLevelID = "security_level"
)

// Current value of the zone setting.
type SettingListResponseZonesSecurityLevelValue string

const (
	SettingListResponseZonesSecurityLevelValueOff            SettingListResponseZonesSecurityLevelValue = "off"
	SettingListResponseZonesSecurityLevelValueEssentiallyOff SettingListResponseZonesSecurityLevelValue = "essentially_off"
	SettingListResponseZonesSecurityLevelValueLow            SettingListResponseZonesSecurityLevelValue = "low"
	SettingListResponseZonesSecurityLevelValueMedium         SettingListResponseZonesSecurityLevelValue = "medium"
	SettingListResponseZonesSecurityLevelValueHigh           SettingListResponseZonesSecurityLevelValue = "high"
	SettingListResponseZonesSecurityLevelValueUnderAttack    SettingListResponseZonesSecurityLevelValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesSecurityLevelEditable bool

const (
	SettingListResponseZonesSecurityLevelEditableTrue  SettingListResponseZonesSecurityLevelEditable = true
	SettingListResponseZonesSecurityLevelEditableFalse SettingListResponseZonesSecurityLevelEditable = false
)

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
type SettingListResponseZonesServerSideExclude struct {
	// ID of the zone setting.
	ID SettingListResponseZonesServerSideExcludeID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesServerSideExcludeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesServerSideExcludeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesServerSideExcludeJSON `json:"-"`
}

// settingListResponseZonesServerSideExcludeJSON contains the JSON metadata for the
// struct [SettingListResponseZonesServerSideExclude]
type settingListResponseZonesServerSideExcludeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesServerSideExclude) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesServerSideExcludeID string

const (
	SettingListResponseZonesServerSideExcludeIDServerSideExclude SettingListResponseZonesServerSideExcludeID = "server_side_exclude"
)

// Current value of the zone setting.
type SettingListResponseZonesServerSideExcludeValue string

const (
	SettingListResponseZonesServerSideExcludeValueOn  SettingListResponseZonesServerSideExcludeValue = "on"
	SettingListResponseZonesServerSideExcludeValueOff SettingListResponseZonesServerSideExcludeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesServerSideExcludeEditable bool

const (
	SettingListResponseZonesServerSideExcludeEditableTrue  SettingListResponseZonesServerSideExcludeEditable = true
	SettingListResponseZonesServerSideExcludeEditableFalse SettingListResponseZonesServerSideExcludeEditable = false
)

// Allow SHA1 support.
type SettingListResponseZonesSha1Support struct {
	// Zone setting identifier.
	ID SettingListResponseZonesSha1SupportID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesSha1SupportValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesSha1SupportJSON `json:"-"`
}

// settingListResponseZonesSha1SupportJSON contains the JSON metadata for the
// struct [SettingListResponseZonesSha1Support]
type settingListResponseZonesSha1SupportJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesSha1Support) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesSha1Support) implementsSettingListResponse() {}

// Zone setting identifier.
type SettingListResponseZonesSha1SupportID string

const (
	SettingListResponseZonesSha1SupportIDSha1Support SettingListResponseZonesSha1SupportID = "sha1_support"
)

// Current value of the zone setting.
type SettingListResponseZonesSha1SupportValue string

const (
	SettingListResponseZonesSha1SupportValueOff SettingListResponseZonesSha1SupportValue = "off"
	SettingListResponseZonesSha1SupportValueOn  SettingListResponseZonesSha1SupportValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesSha1SupportEditable bool

const (
	SettingListResponseZonesSha1SupportEditableTrue  SettingListResponseZonesSha1SupportEditable = true
	SettingListResponseZonesSha1SupportEditableFalse SettingListResponseZonesSha1SupportEditable = false
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type SettingListResponseZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID SettingListResponseZonesSortQueryStringForCacheID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesSortQueryStringForCacheValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesSortQueryStringForCacheJSON `json:"-"`
}

// settingListResponseZonesSortQueryStringForCacheJSON contains the JSON metadata
// for the struct [SettingListResponseZonesSortQueryStringForCache]
type settingListResponseZonesSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesSortQueryStringForCache) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesSortQueryStringForCacheID string

const (
	SettingListResponseZonesSortQueryStringForCacheIDSortQueryStringForCache SettingListResponseZonesSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type SettingListResponseZonesSortQueryStringForCacheValue string

const (
	SettingListResponseZonesSortQueryStringForCacheValueOn  SettingListResponseZonesSortQueryStringForCacheValue = "on"
	SettingListResponseZonesSortQueryStringForCacheValueOff SettingListResponseZonesSortQueryStringForCacheValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesSortQueryStringForCacheEditable bool

const (
	SettingListResponseZonesSortQueryStringForCacheEditableTrue  SettingListResponseZonesSortQueryStringForCacheEditable = true
	SettingListResponseZonesSortQueryStringForCacheEditableFalse SettingListResponseZonesSortQueryStringForCacheEditable = false
)

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
type SettingListResponseZonesSSL struct {
	// ID of the zone setting.
	ID SettingListResponseZonesSSLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesSSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesSSLJSON `json:"-"`
}

// settingListResponseZonesSSLJSON contains the JSON metadata for the struct
// [SettingListResponseZonesSSL]
type settingListResponseZonesSSLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesSSL) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesSSLID string

const (
	SettingListResponseZonesSSLIDSSL SettingListResponseZonesSSLID = "ssl"
)

// Current value of the zone setting.
type SettingListResponseZonesSSLValue string

const (
	SettingListResponseZonesSSLValueOff      SettingListResponseZonesSSLValue = "off"
	SettingListResponseZonesSSLValueFlexible SettingListResponseZonesSSLValue = "flexible"
	SettingListResponseZonesSSLValueFull     SettingListResponseZonesSSLValue = "full"
	SettingListResponseZonesSSLValueStrict   SettingListResponseZonesSSLValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesSSLEditable bool

const (
	SettingListResponseZonesSSLEditableTrue  SettingListResponseZonesSSLEditable = true
	SettingListResponseZonesSSLEditableFalse SettingListResponseZonesSSLEditable = false
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SettingListResponseZonesSSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID SettingListResponseZonesSSLRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                       `json:"enabled"`
	JSON    settingListResponseZonesSSLRecommenderJSON `json:"-"`
}

// settingListResponseZonesSSLRecommenderJSON contains the JSON metadata for the
// struct [SettingListResponseZonesSSLRecommender]
type settingListResponseZonesSSLRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesSSLRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesSSLRecommender) implementsSettingListResponse() {}

// Enrollment value for SSL/TLS Recommender.
type SettingListResponseZonesSSLRecommenderID string

const (
	SettingListResponseZonesSSLRecommenderIDSSLRecommender SettingListResponseZonesSSLRecommenderID = "ssl_recommender"
)

// Only allows TLS1.2.
type SettingListResponseZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID SettingListResponseZonesTLS1_2OnlyID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesTLS1_2OnlyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesTLS1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesTls1_2OnlyJSON `json:"-"`
}

// settingListResponseZonesTls1_2OnlyJSON contains the JSON metadata for the struct
// [SettingListResponseZonesTLS1_2Only]
type settingListResponseZonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesTLS1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesTLS1_2Only) implementsSettingListResponse() {}

// Zone setting identifier.
type SettingListResponseZonesTLS1_2OnlyID string

const (
	SettingListResponseZonesTLS1_2OnlyIDTLS1_2Only SettingListResponseZonesTLS1_2OnlyID = "tls_1_2_only"
)

// Current value of the zone setting.
type SettingListResponseZonesTLS1_2OnlyValue string

const (
	SettingListResponseZonesTLS1_2OnlyValueOff SettingListResponseZonesTLS1_2OnlyValue = "off"
	SettingListResponseZonesTLS1_2OnlyValueOn  SettingListResponseZonesTLS1_2OnlyValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesTLS1_2OnlyEditable bool

const (
	SettingListResponseZonesTLS1_2OnlyEditableTrue  SettingListResponseZonesTLS1_2OnlyEditable = true
	SettingListResponseZonesTLS1_2OnlyEditableFalse SettingListResponseZonesTLS1_2OnlyEditable = false
)

// Enables Crypto TLS 1.3 feature for a zone.
type SettingListResponseZonesTLS1_3 struct {
	// ID of the zone setting.
	ID SettingListResponseZonesTLS1_3ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesTLS1_3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesTLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesTls1_3JSON `json:"-"`
}

// settingListResponseZonesTls1_3JSON contains the JSON metadata for the struct
// [SettingListResponseZonesTLS1_3]
type settingListResponseZonesTls1_3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesTLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesTLS1_3) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesTLS1_3ID string

const (
	SettingListResponseZonesTLS1_3IDTLS1_3 SettingListResponseZonesTLS1_3ID = "tls_1_3"
)

// Current value of the zone setting.
type SettingListResponseZonesTLS1_3Value string

const (
	SettingListResponseZonesTLS1_3ValueOn  SettingListResponseZonesTLS1_3Value = "on"
	SettingListResponseZonesTLS1_3ValueOff SettingListResponseZonesTLS1_3Value = "off"
	SettingListResponseZonesTLS1_3ValueZrt SettingListResponseZonesTLS1_3Value = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesTLS1_3Editable bool

const (
	SettingListResponseZonesTLS1_3EditableTrue  SettingListResponseZonesTLS1_3Editable = true
	SettingListResponseZonesTLS1_3EditableFalse SettingListResponseZonesTLS1_3Editable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type SettingListResponseZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID SettingListResponseZonesTLSClientAuthID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesTLSClientAuthValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesTLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesTLSClientAuthJSON `json:"-"`
}

// settingListResponseZonesTLSClientAuthJSON contains the JSON metadata for the
// struct [SettingListResponseZonesTLSClientAuth]
type settingListResponseZonesTLSClientAuthJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesTLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesTLSClientAuth) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesTLSClientAuthID string

const (
	SettingListResponseZonesTLSClientAuthIDTLSClientAuth SettingListResponseZonesTLSClientAuthID = "tls_client_auth"
)

// Current value of the zone setting.
type SettingListResponseZonesTLSClientAuthValue string

const (
	SettingListResponseZonesTLSClientAuthValueOn  SettingListResponseZonesTLSClientAuthValue = "on"
	SettingListResponseZonesTLSClientAuthValueOff SettingListResponseZonesTLSClientAuthValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesTLSClientAuthEditable bool

const (
	SettingListResponseZonesTLSClientAuthEditableTrue  SettingListResponseZonesTLSClientAuthEditable = true
	SettingListResponseZonesTLSClientAuthEditableFalse SettingListResponseZonesTLSClientAuthEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingListResponseZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID SettingListResponseZonesTrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesTrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesTrueClientIPHeaderJSON `json:"-"`
}

// settingListResponseZonesTrueClientIPHeaderJSON contains the JSON metadata for
// the struct [SettingListResponseZonesTrueClientIPHeader]
type settingListResponseZonesTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesTrueClientIPHeader) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesTrueClientIPHeaderID string

const (
	SettingListResponseZonesTrueClientIPHeaderIDTrueClientIPHeader SettingListResponseZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Current value of the zone setting.
type SettingListResponseZonesTrueClientIPHeaderValue string

const (
	SettingListResponseZonesTrueClientIPHeaderValueOn  SettingListResponseZonesTrueClientIPHeaderValue = "on"
	SettingListResponseZonesTrueClientIPHeaderValueOff SettingListResponseZonesTrueClientIPHeaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesTrueClientIPHeaderEditable bool

const (
	SettingListResponseZonesTrueClientIPHeaderEditableTrue  SettingListResponseZonesTrueClientIPHeaderEditable = true
	SettingListResponseZonesTrueClientIPHeaderEditableFalse SettingListResponseZonesTrueClientIPHeaderEditable = false
)

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
type SettingListResponseZonesWAF struct {
	// ID of the zone setting.
	ID SettingListResponseZonesWAFID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesWAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesWAFJSON `json:"-"`
}

// settingListResponseZonesWAFJSON contains the JSON metadata for the struct
// [SettingListResponseZonesWAF]
type settingListResponseZonesWAFJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesWAF) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesWAFID string

const (
	SettingListResponseZonesWAFIDWAF SettingListResponseZonesWAFID = "waf"
)

// Current value of the zone setting.
type SettingListResponseZonesWAFValue string

const (
	SettingListResponseZonesWAFValueOn  SettingListResponseZonesWAFValue = "on"
	SettingListResponseZonesWAFValueOff SettingListResponseZonesWAFValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesWAFEditable bool

const (
	SettingListResponseZonesWAFEditableTrue  SettingListResponseZonesWAFEditable = true
	SettingListResponseZonesWAFEditableFalse SettingListResponseZonesWAFEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type SettingListResponseZonesWebp struct {
	// ID of the zone setting.
	ID SettingListResponseZonesWebpID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesWebpValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesWebpJSON `json:"-"`
}

// settingListResponseZonesWebpJSON contains the JSON metadata for the struct
// [SettingListResponseZonesWebp]
type settingListResponseZonesWebpJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesWebp) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesWebpID string

const (
	SettingListResponseZonesWebpIDWebp SettingListResponseZonesWebpID = "webp"
)

// Current value of the zone setting.
type SettingListResponseZonesWebpValue string

const (
	SettingListResponseZonesWebpValueOff SettingListResponseZonesWebpValue = "off"
	SettingListResponseZonesWebpValueOn  SettingListResponseZonesWebpValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesWebpEditable bool

const (
	SettingListResponseZonesWebpEditableTrue  SettingListResponseZonesWebpEditable = true
	SettingListResponseZonesWebpEditableFalse SettingListResponseZonesWebpEditable = false
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type SettingListResponseZonesWebsockets struct {
	// ID of the zone setting.
	ID SettingListResponseZonesWebsocketsID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingListResponseZonesWebsocketsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingListResponseZonesWebsocketsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingListResponseZonesWebsocketsJSON `json:"-"`
}

// settingListResponseZonesWebsocketsJSON contains the JSON metadata for the struct
// [SettingListResponseZonesWebsockets]
type settingListResponseZonesWebsocketsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseZonesWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingListResponseZonesWebsockets) implementsSettingListResponse() {}

// ID of the zone setting.
type SettingListResponseZonesWebsocketsID string

const (
	SettingListResponseZonesWebsocketsIDWebsockets SettingListResponseZonesWebsocketsID = "websockets"
)

// Current value of the zone setting.
type SettingListResponseZonesWebsocketsValue string

const (
	SettingListResponseZonesWebsocketsValueOff SettingListResponseZonesWebsocketsValue = "off"
	SettingListResponseZonesWebsocketsValueOn  SettingListResponseZonesWebsocketsValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingListResponseZonesWebsocketsEditable bool

const (
	SettingListResponseZonesWebsocketsEditableTrue  SettingListResponseZonesWebsocketsEditable = true
	SettingListResponseZonesWebsocketsEditableFalse SettingListResponseZonesWebsocketsEditable = false
)

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [SettingEditResponseZones0rtt],
// [SettingEditResponseZonesAdvancedDDOS], [SettingEditResponseZonesAlwaysOnline],
// [SettingEditResponseZonesAlwaysUseHTTPS],
// [SettingEditResponseZonesAutomaticHTTPSRewrites],
// [SettingEditResponseZonesBrotli], [SettingEditResponseZonesBrowserCacheTTL],
// [SettingEditResponseZonesBrowserCheck], [SettingEditResponseZonesCacheLevel],
// [SettingEditResponseZonesChallengeTTL], [SettingEditResponseZonesCiphers],
// [SettingEditResponseZonesCnameFlattening],
// [SettingEditResponseZonesDevelopmentMode], [SettingEditResponseZonesEarlyHints],
// [SettingEditResponseZonesEdgeCacheTTL],
// [SettingEditResponseZonesEmailObfuscation],
// [SettingEditResponseZonesH2Prioritization],
// [SettingEditResponseZonesHotlinkProtection], [SettingEditResponseZonesHTTP2],
// [SettingEditResponseZonesHTTP3], [SettingEditResponseZonesImageResizing],
// [SettingEditResponseZonesIPGeolocation], [SettingEditResponseZonesIPV6],
// [SettingEditResponseZonesMaxUpload], [SettingEditResponseZonesMinTLSVersion],
// [SettingEditResponseZonesMinify], [SettingEditResponseZonesMirage],
// [SettingEditResponseZonesMobileRedirect], [SettingEditResponseZonesNEL],
// [SettingEditResponseZonesOpportunisticEncryption],
// [SettingEditResponseZonesOpportunisticOnion],
// [SettingEditResponseZonesOrangeToOrange],
// [SettingEditResponseZonesOriginErrorPagePassThru],
// [SettingEditResponseZonesPolish], [SettingEditResponseZonesPrefetchPreload],
// [SettingEditResponseZonesProxyReadTimeout],
// [SettingEditResponseZonesPseudoIPV4],
// [SettingEditResponseZonesResponseBuffering],
// [SettingEditResponseZonesRocketLoader],
// [SettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [SettingEditResponseZonesSecurityHeader],
// [SettingEditResponseZonesSecurityLevel],
// [SettingEditResponseZonesServerSideExclude],
// [SettingEditResponseZonesSha1Support],
// [SettingEditResponseZonesSortQueryStringForCache],
// [SettingEditResponseZonesSSL], [SettingEditResponseZonesSSLRecommender],
// [SettingEditResponseZonesTLS1_2Only], [SettingEditResponseZonesTLS1_3],
// [SettingEditResponseZonesTLSClientAuth],
// [SettingEditResponseZonesTrueClientIPHeader], [SettingEditResponseZonesWAF],
// [SettingEditResponseZonesWebp] or [SettingEditResponseZonesWebsockets].
type SettingEditResponse interface {
	implementsSettingEditResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*SettingEditResponse)(nil)).Elem(), "")
}

// 0-RTT session resumption enabled for this zone.
type SettingEditResponseZones0rtt struct {
	// ID of the zone setting.
	ID SettingEditResponseZones0rttID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZones0rttValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZones0rttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZones0rttJSON `json:"-"`
}

// settingEditResponseZones0rttJSON contains the JSON metadata for the struct
// [SettingEditResponseZones0rtt]
type settingEditResponseZones0rttJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZones0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZones0rtt) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZones0rttID string

const (
	SettingEditResponseZones0rttID0rtt SettingEditResponseZones0rttID = "0rtt"
)

// Current value of the zone setting.
type SettingEditResponseZones0rttValue string

const (
	SettingEditResponseZones0rttValueOn  SettingEditResponseZones0rttValue = "on"
	SettingEditResponseZones0rttValueOff SettingEditResponseZones0rttValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZones0rttEditable bool

const (
	SettingEditResponseZones0rttEditableTrue  SettingEditResponseZones0rttEditable = true
	SettingEditResponseZones0rttEditableFalse SettingEditResponseZones0rttEditable = false
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type SettingEditResponseZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesAdvancedDDOSID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesAdvancedDDOSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesAdvancedDDOSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesAdvancedDDOSJSON `json:"-"`
}

// settingEditResponseZonesAdvancedDDOSJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesAdvancedDDOS]
type settingEditResponseZonesAdvancedDDOSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesAdvancedDDOS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesAdvancedDDOS) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesAdvancedDDOSID string

const (
	SettingEditResponseZonesAdvancedDDOSIDAdvancedDDOS SettingEditResponseZonesAdvancedDDOSID = "advanced_ddos"
)

// Current value of the zone setting.
type SettingEditResponseZonesAdvancedDDOSValue string

const (
	SettingEditResponseZonesAdvancedDDOSValueOn  SettingEditResponseZonesAdvancedDDOSValue = "on"
	SettingEditResponseZonesAdvancedDDOSValueOff SettingEditResponseZonesAdvancedDDOSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesAdvancedDDOSEditable bool

const (
	SettingEditResponseZonesAdvancedDDOSEditableTrue  SettingEditResponseZonesAdvancedDDOSEditable = true
	SettingEditResponseZonesAdvancedDDOSEditableFalse SettingEditResponseZonesAdvancedDDOSEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type SettingEditResponseZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesAlwaysOnlineID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesAlwaysOnlineValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesAlwaysOnlineJSON `json:"-"`
}

// settingEditResponseZonesAlwaysOnlineJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesAlwaysOnline]
type settingEditResponseZonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesAlwaysOnline) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesAlwaysOnlineID string

const (
	SettingEditResponseZonesAlwaysOnlineIDAlwaysOnline SettingEditResponseZonesAlwaysOnlineID = "always_online"
)

// Current value of the zone setting.
type SettingEditResponseZonesAlwaysOnlineValue string

const (
	SettingEditResponseZonesAlwaysOnlineValueOn  SettingEditResponseZonesAlwaysOnlineValue = "on"
	SettingEditResponseZonesAlwaysOnlineValueOff SettingEditResponseZonesAlwaysOnlineValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesAlwaysOnlineEditable bool

const (
	SettingEditResponseZonesAlwaysOnlineEditableTrue  SettingEditResponseZonesAlwaysOnlineEditable = true
	SettingEditResponseZonesAlwaysOnlineEditableFalse SettingEditResponseZonesAlwaysOnlineEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingEditResponseZonesAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesAlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesAlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesAlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesAlwaysUseHTTPSJSON `json:"-"`
}

// settingEditResponseZonesAlwaysUseHTTPSJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesAlwaysUseHTTPS]
type settingEditResponseZonesAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesAlwaysUseHTTPS) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesAlwaysUseHTTPSID string

const (
	SettingEditResponseZonesAlwaysUseHTTPSIDAlwaysUseHTTPS SettingEditResponseZonesAlwaysUseHTTPSID = "always_use_https"
)

// Current value of the zone setting.
type SettingEditResponseZonesAlwaysUseHTTPSValue string

const (
	SettingEditResponseZonesAlwaysUseHTTPSValueOn  SettingEditResponseZonesAlwaysUseHTTPSValue = "on"
	SettingEditResponseZonesAlwaysUseHTTPSValueOff SettingEditResponseZonesAlwaysUseHTTPSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesAlwaysUseHTTPSEditable bool

const (
	SettingEditResponseZonesAlwaysUseHTTPSEditableTrue  SettingEditResponseZonesAlwaysUseHTTPSEditable = true
	SettingEditResponseZonesAlwaysUseHTTPSEditableFalse SettingEditResponseZonesAlwaysUseHTTPSEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingEditResponseZonesAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesAutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesAutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesAutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesAutomaticHTTPSRewritesJSON `json:"-"`
}

// settingEditResponseZonesAutomaticHTTPSRewritesJSON contains the JSON metadata
// for the struct [SettingEditResponseZonesAutomaticHTTPSRewrites]
type settingEditResponseZonesAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesAutomaticHTTPSRewrites) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesAutomaticHTTPSRewritesID string

const (
	SettingEditResponseZonesAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites SettingEditResponseZonesAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type SettingEditResponseZonesAutomaticHTTPSRewritesValue string

const (
	SettingEditResponseZonesAutomaticHTTPSRewritesValueOn  SettingEditResponseZonesAutomaticHTTPSRewritesValue = "on"
	SettingEditResponseZonesAutomaticHTTPSRewritesValueOff SettingEditResponseZonesAutomaticHTTPSRewritesValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesAutomaticHTTPSRewritesEditable bool

const (
	SettingEditResponseZonesAutomaticHTTPSRewritesEditableTrue  SettingEditResponseZonesAutomaticHTTPSRewritesEditable = true
	SettingEditResponseZonesAutomaticHTTPSRewritesEditableFalse SettingEditResponseZonesAutomaticHTTPSRewritesEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type SettingEditResponseZonesBrotli struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesBrotliID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesBrotliValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesBrotliJSON `json:"-"`
}

// settingEditResponseZonesBrotliJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesBrotli]
type settingEditResponseZonesBrotliJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesBrotli) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesBrotliID string

const (
	SettingEditResponseZonesBrotliIDBrotli SettingEditResponseZonesBrotliID = "brotli"
)

// Current value of the zone setting.
type SettingEditResponseZonesBrotliValue string

const (
	SettingEditResponseZonesBrotliValueOff SettingEditResponseZonesBrotliValue = "off"
	SettingEditResponseZonesBrotliValueOn  SettingEditResponseZonesBrotliValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesBrotliEditable bool

const (
	SettingEditResponseZonesBrotliEditableTrue  SettingEditResponseZonesBrotliEditable = true
	SettingEditResponseZonesBrotliEditableFalse SettingEditResponseZonesBrotliEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingEditResponseZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesBrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesBrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesBrowserCacheTTLJSON `json:"-"`
}

// settingEditResponseZonesBrowserCacheTTLJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesBrowserCacheTTL]
type settingEditResponseZonesBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesBrowserCacheTTL) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesBrowserCacheTTLID string

const (
	SettingEditResponseZonesBrowserCacheTTLIDBrowserCacheTTL SettingEditResponseZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Current value of the zone setting.
type SettingEditResponseZonesBrowserCacheTTLValue float64

const (
	SettingEditResponseZonesBrowserCacheTTLValue0        SettingEditResponseZonesBrowserCacheTTLValue = 0
	SettingEditResponseZonesBrowserCacheTTLValue30       SettingEditResponseZonesBrowserCacheTTLValue = 30
	SettingEditResponseZonesBrowserCacheTTLValue60       SettingEditResponseZonesBrowserCacheTTLValue = 60
	SettingEditResponseZonesBrowserCacheTTLValue120      SettingEditResponseZonesBrowserCacheTTLValue = 120
	SettingEditResponseZonesBrowserCacheTTLValue300      SettingEditResponseZonesBrowserCacheTTLValue = 300
	SettingEditResponseZonesBrowserCacheTTLValue1200     SettingEditResponseZonesBrowserCacheTTLValue = 1200
	SettingEditResponseZonesBrowserCacheTTLValue1800     SettingEditResponseZonesBrowserCacheTTLValue = 1800
	SettingEditResponseZonesBrowserCacheTTLValue3600     SettingEditResponseZonesBrowserCacheTTLValue = 3600
	SettingEditResponseZonesBrowserCacheTTLValue7200     SettingEditResponseZonesBrowserCacheTTLValue = 7200
	SettingEditResponseZonesBrowserCacheTTLValue10800    SettingEditResponseZonesBrowserCacheTTLValue = 10800
	SettingEditResponseZonesBrowserCacheTTLValue14400    SettingEditResponseZonesBrowserCacheTTLValue = 14400
	SettingEditResponseZonesBrowserCacheTTLValue18000    SettingEditResponseZonesBrowserCacheTTLValue = 18000
	SettingEditResponseZonesBrowserCacheTTLValue28800    SettingEditResponseZonesBrowserCacheTTLValue = 28800
	SettingEditResponseZonesBrowserCacheTTLValue43200    SettingEditResponseZonesBrowserCacheTTLValue = 43200
	SettingEditResponseZonesBrowserCacheTTLValue57600    SettingEditResponseZonesBrowserCacheTTLValue = 57600
	SettingEditResponseZonesBrowserCacheTTLValue72000    SettingEditResponseZonesBrowserCacheTTLValue = 72000
	SettingEditResponseZonesBrowserCacheTTLValue86400    SettingEditResponseZonesBrowserCacheTTLValue = 86400
	SettingEditResponseZonesBrowserCacheTTLValue172800   SettingEditResponseZonesBrowserCacheTTLValue = 172800
	SettingEditResponseZonesBrowserCacheTTLValue259200   SettingEditResponseZonesBrowserCacheTTLValue = 259200
	SettingEditResponseZonesBrowserCacheTTLValue345600   SettingEditResponseZonesBrowserCacheTTLValue = 345600
	SettingEditResponseZonesBrowserCacheTTLValue432000   SettingEditResponseZonesBrowserCacheTTLValue = 432000
	SettingEditResponseZonesBrowserCacheTTLValue691200   SettingEditResponseZonesBrowserCacheTTLValue = 691200
	SettingEditResponseZonesBrowserCacheTTLValue1382400  SettingEditResponseZonesBrowserCacheTTLValue = 1382400
	SettingEditResponseZonesBrowserCacheTTLValue2073600  SettingEditResponseZonesBrowserCacheTTLValue = 2073600
	SettingEditResponseZonesBrowserCacheTTLValue2678400  SettingEditResponseZonesBrowserCacheTTLValue = 2678400
	SettingEditResponseZonesBrowserCacheTTLValue5356800  SettingEditResponseZonesBrowserCacheTTLValue = 5356800
	SettingEditResponseZonesBrowserCacheTTLValue16070400 SettingEditResponseZonesBrowserCacheTTLValue = 16070400
	SettingEditResponseZonesBrowserCacheTTLValue31536000 SettingEditResponseZonesBrowserCacheTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesBrowserCacheTTLEditable bool

const (
	SettingEditResponseZonesBrowserCacheTTLEditableTrue  SettingEditResponseZonesBrowserCacheTTLEditable = true
	SettingEditResponseZonesBrowserCacheTTLEditableFalse SettingEditResponseZonesBrowserCacheTTLEditable = false
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type SettingEditResponseZonesBrowserCheck struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesBrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesBrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesBrowserCheckJSON `json:"-"`
}

// settingEditResponseZonesBrowserCheckJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesBrowserCheck]
type settingEditResponseZonesBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesBrowserCheck) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesBrowserCheckID string

const (
	SettingEditResponseZonesBrowserCheckIDBrowserCheck SettingEditResponseZonesBrowserCheckID = "browser_check"
)

// Current value of the zone setting.
type SettingEditResponseZonesBrowserCheckValue string

const (
	SettingEditResponseZonesBrowserCheckValueOn  SettingEditResponseZonesBrowserCheckValue = "on"
	SettingEditResponseZonesBrowserCheckValueOff SettingEditResponseZonesBrowserCheckValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesBrowserCheckEditable bool

const (
	SettingEditResponseZonesBrowserCheckEditableTrue  SettingEditResponseZonesBrowserCheckEditable = true
	SettingEditResponseZonesBrowserCheckEditableFalse SettingEditResponseZonesBrowserCheckEditable = false
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type SettingEditResponseZonesCacheLevel struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesCacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesCacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesCacheLevelJSON `json:"-"`
}

// settingEditResponseZonesCacheLevelJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesCacheLevel]
type settingEditResponseZonesCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesCacheLevel) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesCacheLevelID string

const (
	SettingEditResponseZonesCacheLevelIDCacheLevel SettingEditResponseZonesCacheLevelID = "cache_level"
)

// Current value of the zone setting.
type SettingEditResponseZonesCacheLevelValue string

const (
	SettingEditResponseZonesCacheLevelValueAggressive SettingEditResponseZonesCacheLevelValue = "aggressive"
	SettingEditResponseZonesCacheLevelValueBasic      SettingEditResponseZonesCacheLevelValue = "basic"
	SettingEditResponseZonesCacheLevelValueSimplified SettingEditResponseZonesCacheLevelValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesCacheLevelEditable bool

const (
	SettingEditResponseZonesCacheLevelEditableTrue  SettingEditResponseZonesCacheLevelEditable = true
	SettingEditResponseZonesCacheLevelEditableFalse SettingEditResponseZonesCacheLevelEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type SettingEditResponseZonesChallengeTTL struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesChallengeTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesChallengeTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesChallengeTTLJSON `json:"-"`
}

// settingEditResponseZonesChallengeTTLJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesChallengeTTL]
type settingEditResponseZonesChallengeTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesChallengeTTL) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesChallengeTTLID string

const (
	SettingEditResponseZonesChallengeTTLIDChallengeTTL SettingEditResponseZonesChallengeTTLID = "challenge_ttl"
)

// Current value of the zone setting.
type SettingEditResponseZonesChallengeTTLValue float64

const (
	SettingEditResponseZonesChallengeTTLValue300      SettingEditResponseZonesChallengeTTLValue = 300
	SettingEditResponseZonesChallengeTTLValue900      SettingEditResponseZonesChallengeTTLValue = 900
	SettingEditResponseZonesChallengeTTLValue1800     SettingEditResponseZonesChallengeTTLValue = 1800
	SettingEditResponseZonesChallengeTTLValue2700     SettingEditResponseZonesChallengeTTLValue = 2700
	SettingEditResponseZonesChallengeTTLValue3600     SettingEditResponseZonesChallengeTTLValue = 3600
	SettingEditResponseZonesChallengeTTLValue7200     SettingEditResponseZonesChallengeTTLValue = 7200
	SettingEditResponseZonesChallengeTTLValue10800    SettingEditResponseZonesChallengeTTLValue = 10800
	SettingEditResponseZonesChallengeTTLValue14400    SettingEditResponseZonesChallengeTTLValue = 14400
	SettingEditResponseZonesChallengeTTLValue28800    SettingEditResponseZonesChallengeTTLValue = 28800
	SettingEditResponseZonesChallengeTTLValue57600    SettingEditResponseZonesChallengeTTLValue = 57600
	SettingEditResponseZonesChallengeTTLValue86400    SettingEditResponseZonesChallengeTTLValue = 86400
	SettingEditResponseZonesChallengeTTLValue604800   SettingEditResponseZonesChallengeTTLValue = 604800
	SettingEditResponseZonesChallengeTTLValue2592000  SettingEditResponseZonesChallengeTTLValue = 2592000
	SettingEditResponseZonesChallengeTTLValue31536000 SettingEditResponseZonesChallengeTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesChallengeTTLEditable bool

const (
	SettingEditResponseZonesChallengeTTLEditableTrue  SettingEditResponseZonesChallengeTTLEditable = true
	SettingEditResponseZonesChallengeTTLEditableFalse SettingEditResponseZonesChallengeTTLEditable = false
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type SettingEditResponseZonesCiphers struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesCiphersID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesCiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesCiphersJSON `json:"-"`
}

// settingEditResponseZonesCiphersJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesCiphers]
type settingEditResponseZonesCiphersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesCiphers) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesCiphersID string

const (
	SettingEditResponseZonesCiphersIDCiphers SettingEditResponseZonesCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesCiphersEditable bool

const (
	SettingEditResponseZonesCiphersEditableTrue  SettingEditResponseZonesCiphersEditable = true
	SettingEditResponseZonesCiphersEditableFalse SettingEditResponseZonesCiphersEditable = false
)

// Whether or not cname flattening is on.
type SettingEditResponseZonesCnameFlattening struct {
	// How to flatten the cname destination.
	ID SettingEditResponseZonesCnameFlatteningID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesCnameFlatteningValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesCnameFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesCnameFlatteningJSON `json:"-"`
}

// settingEditResponseZonesCnameFlatteningJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesCnameFlattening]
type settingEditResponseZonesCnameFlatteningJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesCnameFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesCnameFlattening) implementsSettingEditResponse() {}

// How to flatten the cname destination.
type SettingEditResponseZonesCnameFlatteningID string

const (
	SettingEditResponseZonesCnameFlatteningIDCnameFlattening SettingEditResponseZonesCnameFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type SettingEditResponseZonesCnameFlatteningValue string

const (
	SettingEditResponseZonesCnameFlatteningValueFlattenAtRoot SettingEditResponseZonesCnameFlatteningValue = "flatten_at_root"
	SettingEditResponseZonesCnameFlatteningValueFlattenAll    SettingEditResponseZonesCnameFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesCnameFlatteningEditable bool

const (
	SettingEditResponseZonesCnameFlatteningEditableTrue  SettingEditResponseZonesCnameFlatteningEditable = true
	SettingEditResponseZonesCnameFlatteningEditableFalse SettingEditResponseZonesCnameFlatteningEditable = false
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type SettingEditResponseZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesDevelopmentModeID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesDevelopmentModeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                                     `json:"time_remaining"`
	JSON          settingEditResponseZonesDevelopmentModeJSON `json:"-"`
}

// settingEditResponseZonesDevelopmentModeJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesDevelopmentMode]
type settingEditResponseZonesDevelopmentModeJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingEditResponseZonesDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesDevelopmentMode) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesDevelopmentModeID string

const (
	SettingEditResponseZonesDevelopmentModeIDDevelopmentMode SettingEditResponseZonesDevelopmentModeID = "development_mode"
)

// Current value of the zone setting.
type SettingEditResponseZonesDevelopmentModeValue string

const (
	SettingEditResponseZonesDevelopmentModeValueOn  SettingEditResponseZonesDevelopmentModeValue = "on"
	SettingEditResponseZonesDevelopmentModeValueOff SettingEditResponseZonesDevelopmentModeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesDevelopmentModeEditable bool

const (
	SettingEditResponseZonesDevelopmentModeEditableTrue  SettingEditResponseZonesDevelopmentModeEditable = true
	SettingEditResponseZonesDevelopmentModeEditableFalse SettingEditResponseZonesDevelopmentModeEditable = false
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type SettingEditResponseZonesEarlyHints struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesEarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesEarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesEarlyHintsJSON `json:"-"`
}

// settingEditResponseZonesEarlyHintsJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesEarlyHints]
type settingEditResponseZonesEarlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesEarlyHints) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesEarlyHintsID string

const (
	SettingEditResponseZonesEarlyHintsIDEarlyHints SettingEditResponseZonesEarlyHintsID = "early_hints"
)

// Current value of the zone setting.
type SettingEditResponseZonesEarlyHintsValue string

const (
	SettingEditResponseZonesEarlyHintsValueOn  SettingEditResponseZonesEarlyHintsValue = "on"
	SettingEditResponseZonesEarlyHintsValueOff SettingEditResponseZonesEarlyHintsValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesEarlyHintsEditable bool

const (
	SettingEditResponseZonesEarlyHintsEditableTrue  SettingEditResponseZonesEarlyHintsEditable = true
	SettingEditResponseZonesEarlyHintsEditableFalse SettingEditResponseZonesEarlyHintsEditable = false
)

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

func (r SettingEditResponseZonesEdgeCacheTTL) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesEdgeCacheTTLID string

const (
	SettingEditResponseZonesEdgeCacheTTLIDEdgeCacheTTL SettingEditResponseZonesEdgeCacheTTLID = "edge_cache_ttl"
)

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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesEdgeCacheTTLEditable bool

const (
	SettingEditResponseZonesEdgeCacheTTLEditableTrue  SettingEditResponseZonesEdgeCacheTTLEditable = true
	SettingEditResponseZonesEdgeCacheTTLEditableFalse SettingEditResponseZonesEdgeCacheTTLEditable = false
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingEditResponseZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesEmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesEmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesEmailObfuscationJSON `json:"-"`
}

// settingEditResponseZonesEmailObfuscationJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesEmailObfuscation]
type settingEditResponseZonesEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesEmailObfuscation) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesEmailObfuscationID string

const (
	SettingEditResponseZonesEmailObfuscationIDEmailObfuscation SettingEditResponseZonesEmailObfuscationID = "email_obfuscation"
)

// Current value of the zone setting.
type SettingEditResponseZonesEmailObfuscationValue string

const (
	SettingEditResponseZonesEmailObfuscationValueOn  SettingEditResponseZonesEmailObfuscationValue = "on"
	SettingEditResponseZonesEmailObfuscationValueOff SettingEditResponseZonesEmailObfuscationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesEmailObfuscationEditable bool

const (
	SettingEditResponseZonesEmailObfuscationEditableTrue  SettingEditResponseZonesEmailObfuscationEditable = true
	SettingEditResponseZonesEmailObfuscationEditableFalse SettingEditResponseZonesEmailObfuscationEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type SettingEditResponseZonesH2Prioritization struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesH2PrioritizationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesH2PrioritizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesH2PrioritizationJSON `json:"-"`
}

// settingEditResponseZonesH2PrioritizationJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesH2Prioritization]
type settingEditResponseZonesH2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesH2Prioritization) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesH2PrioritizationID string

const (
	SettingEditResponseZonesH2PrioritizationIDH2Prioritization SettingEditResponseZonesH2PrioritizationID = "h2_prioritization"
)

// Current value of the zone setting.
type SettingEditResponseZonesH2PrioritizationValue string

const (
	SettingEditResponseZonesH2PrioritizationValueOn     SettingEditResponseZonesH2PrioritizationValue = "on"
	SettingEditResponseZonesH2PrioritizationValueOff    SettingEditResponseZonesH2PrioritizationValue = "off"
	SettingEditResponseZonesH2PrioritizationValueCustom SettingEditResponseZonesH2PrioritizationValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesH2PrioritizationEditable bool

const (
	SettingEditResponseZonesH2PrioritizationEditableTrue  SettingEditResponseZonesH2PrioritizationEditable = true
	SettingEditResponseZonesH2PrioritizationEditableFalse SettingEditResponseZonesH2PrioritizationEditable = false
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type SettingEditResponseZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesHotlinkProtectionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesHotlinkProtectionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesHotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesHotlinkProtectionJSON `json:"-"`
}

// settingEditResponseZonesHotlinkProtectionJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesHotlinkProtection]
type settingEditResponseZonesHotlinkProtectionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesHotlinkProtection) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesHotlinkProtectionID string

const (
	SettingEditResponseZonesHotlinkProtectionIDHotlinkProtection SettingEditResponseZonesHotlinkProtectionID = "hotlink_protection"
)

// Current value of the zone setting.
type SettingEditResponseZonesHotlinkProtectionValue string

const (
	SettingEditResponseZonesHotlinkProtectionValueOn  SettingEditResponseZonesHotlinkProtectionValue = "on"
	SettingEditResponseZonesHotlinkProtectionValueOff SettingEditResponseZonesHotlinkProtectionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesHotlinkProtectionEditable bool

const (
	SettingEditResponseZonesHotlinkProtectionEditableTrue  SettingEditResponseZonesHotlinkProtectionEditable = true
	SettingEditResponseZonesHotlinkProtectionEditableFalse SettingEditResponseZonesHotlinkProtectionEditable = false
)

// HTTP2 enabled for this zone.
type SettingEditResponseZonesHTTP2 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesHTTP2ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesHTTP2Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesHTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesHTTP2JSON `json:"-"`
}

// settingEditResponseZonesHTTP2JSON contains the JSON metadata for the struct
// [SettingEditResponseZonesHTTP2]
type settingEditResponseZonesHTTP2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesHTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesHTTP2) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesHTTP2ID string

const (
	SettingEditResponseZonesHTTP2IDHTTP2 SettingEditResponseZonesHTTP2ID = "http2"
)

// Current value of the zone setting.
type SettingEditResponseZonesHTTP2Value string

const (
	SettingEditResponseZonesHTTP2ValueOn  SettingEditResponseZonesHTTP2Value = "on"
	SettingEditResponseZonesHTTP2ValueOff SettingEditResponseZonesHTTP2Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesHTTP2Editable bool

const (
	SettingEditResponseZonesHTTP2EditableTrue  SettingEditResponseZonesHTTP2Editable = true
	SettingEditResponseZonesHTTP2EditableFalse SettingEditResponseZonesHTTP2Editable = false
)

// HTTP3 enabled for this zone.
type SettingEditResponseZonesHTTP3 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesHTTP3ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesHTTP3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesHTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesHTTP3JSON `json:"-"`
}

// settingEditResponseZonesHTTP3JSON contains the JSON metadata for the struct
// [SettingEditResponseZonesHTTP3]
type settingEditResponseZonesHTTP3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesHTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesHTTP3) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesHTTP3ID string

const (
	SettingEditResponseZonesHTTP3IDHTTP3 SettingEditResponseZonesHTTP3ID = "http3"
)

// Current value of the zone setting.
type SettingEditResponseZonesHTTP3Value string

const (
	SettingEditResponseZonesHTTP3ValueOn  SettingEditResponseZonesHTTP3Value = "on"
	SettingEditResponseZonesHTTP3ValueOff SettingEditResponseZonesHTTP3Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesHTTP3Editable bool

const (
	SettingEditResponseZonesHTTP3EditableTrue  SettingEditResponseZonesHTTP3Editable = true
	SettingEditResponseZonesHTTP3EditableFalse SettingEditResponseZonesHTTP3Editable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type SettingEditResponseZonesImageResizing struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesImageResizingID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesImageResizingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesImageResizingJSON `json:"-"`
}

// settingEditResponseZonesImageResizingJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesImageResizing]
type settingEditResponseZonesImageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesImageResizing) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesImageResizingID string

const (
	SettingEditResponseZonesImageResizingIDImageResizing SettingEditResponseZonesImageResizingID = "image_resizing"
)

// Current value of the zone setting.
type SettingEditResponseZonesImageResizingValue string

const (
	SettingEditResponseZonesImageResizingValueOn   SettingEditResponseZonesImageResizingValue = "on"
	SettingEditResponseZonesImageResizingValueOff  SettingEditResponseZonesImageResizingValue = "off"
	SettingEditResponseZonesImageResizingValueOpen SettingEditResponseZonesImageResizingValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesImageResizingEditable bool

const (
	SettingEditResponseZonesImageResizingEditableTrue  SettingEditResponseZonesImageResizingEditable = true
	SettingEditResponseZonesImageResizingEditableFalse SettingEditResponseZonesImageResizingEditable = false
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingEditResponseZonesIPGeolocation struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesIPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesIPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesIPGeolocationJSON `json:"-"`
}

// settingEditResponseZonesIPGeolocationJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesIPGeolocation]
type settingEditResponseZonesIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesIPGeolocation) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesIPGeolocationID string

const (
	SettingEditResponseZonesIPGeolocationIDIPGeolocation SettingEditResponseZonesIPGeolocationID = "ip_geolocation"
)

// Current value of the zone setting.
type SettingEditResponseZonesIPGeolocationValue string

const (
	SettingEditResponseZonesIPGeolocationValueOn  SettingEditResponseZonesIPGeolocationValue = "on"
	SettingEditResponseZonesIPGeolocationValueOff SettingEditResponseZonesIPGeolocationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesIPGeolocationEditable bool

const (
	SettingEditResponseZonesIPGeolocationEditableTrue  SettingEditResponseZonesIPGeolocationEditable = true
	SettingEditResponseZonesIPGeolocationEditableFalse SettingEditResponseZonesIPGeolocationEditable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type SettingEditResponseZonesIPV6 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesIPV6ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesIPV6Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesIPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesIPV6JSON `json:"-"`
}

// settingEditResponseZonesIPV6JSON contains the JSON metadata for the struct
// [SettingEditResponseZonesIPV6]
type settingEditResponseZonesIPV6JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesIPV6) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesIPV6ID string

const (
	SettingEditResponseZonesIPV6IDIPV6 SettingEditResponseZonesIPV6ID = "ipv6"
)

// Current value of the zone setting.
type SettingEditResponseZonesIPV6Value string

const (
	SettingEditResponseZonesIPV6ValueOff SettingEditResponseZonesIPV6Value = "off"
	SettingEditResponseZonesIPV6ValueOn  SettingEditResponseZonesIPV6Value = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesIPV6Editable bool

const (
	SettingEditResponseZonesIPV6EditableTrue  SettingEditResponseZonesIPV6Editable = true
	SettingEditResponseZonesIPV6EditableFalse SettingEditResponseZonesIPV6Editable = false
)

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

func (r SettingEditResponseZonesMaxUpload) implementsSettingEditResponse() {}

// identifier of the zone setting.
type SettingEditResponseZonesMaxUploadID string

const (
	SettingEditResponseZonesMaxUploadIDMaxUpload SettingEditResponseZonesMaxUploadID = "max_upload"
)

// Current value of the zone setting.
type SettingEditResponseZonesMaxUploadValue float64

const (
	SettingEditResponseZonesMaxUploadValue100 SettingEditResponseZonesMaxUploadValue = 100
	SettingEditResponseZonesMaxUploadValue200 SettingEditResponseZonesMaxUploadValue = 200
	SettingEditResponseZonesMaxUploadValue500 SettingEditResponseZonesMaxUploadValue = 500
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesMaxUploadEditable bool

const (
	SettingEditResponseZonesMaxUploadEditableTrue  SettingEditResponseZonesMaxUploadEditable = true
	SettingEditResponseZonesMaxUploadEditableFalse SettingEditResponseZonesMaxUploadEditable = false
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type SettingEditResponseZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesMinTLSVersionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesMinTLSVersionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesMinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesMinTLSVersionJSON `json:"-"`
}

// settingEditResponseZonesMinTLSVersionJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesMinTLSVersion]
type settingEditResponseZonesMinTLSVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesMinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesMinTLSVersion) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesMinTLSVersionID string

const (
	SettingEditResponseZonesMinTLSVersionIDMinTLSVersion SettingEditResponseZonesMinTLSVersionID = "min_tls_version"
)

// Current value of the zone setting.
type SettingEditResponseZonesMinTLSVersionValue string

const (
	SettingEditResponseZonesMinTLSVersionValue1_0 SettingEditResponseZonesMinTLSVersionValue = "1.0"
	SettingEditResponseZonesMinTLSVersionValue1_1 SettingEditResponseZonesMinTLSVersionValue = "1.1"
	SettingEditResponseZonesMinTLSVersionValue1_2 SettingEditResponseZonesMinTLSVersionValue = "1.2"
	SettingEditResponseZonesMinTLSVersionValue1_3 SettingEditResponseZonesMinTLSVersionValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesMinTLSVersionEditable bool

const (
	SettingEditResponseZonesMinTLSVersionEditableTrue  SettingEditResponseZonesMinTLSVersionEditable = true
	SettingEditResponseZonesMinTLSVersionEditableFalse SettingEditResponseZonesMinTLSVersionEditable = false
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type SettingEditResponseZonesMinify struct {
	// Zone setting identifier.
	ID SettingEditResponseZonesMinifyID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesMinifyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesMinifyJSON `json:"-"`
}

// settingEditResponseZonesMinifyJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesMinify]
type settingEditResponseZonesMinifyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesMinify) implementsSettingEditResponse() {}

// Zone setting identifier.
type SettingEditResponseZonesMinifyID string

const (
	SettingEditResponseZonesMinifyIDMinify SettingEditResponseZonesMinifyID = "minify"
)

// Current value of the zone setting.
type SettingEditResponseZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css SettingEditResponseZonesMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML SettingEditResponseZonesMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   SettingEditResponseZonesMinifyValueJs   `json:"js"`
	JSON settingEditResponseZonesMinifyValueJSON `json:"-"`
}

// settingEditResponseZonesMinifyValueJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesMinifyValue]
type settingEditResponseZonesMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type SettingEditResponseZonesMinifyValueCss string

const (
	SettingEditResponseZonesMinifyValueCssOn  SettingEditResponseZonesMinifyValueCss = "on"
	SettingEditResponseZonesMinifyValueCssOff SettingEditResponseZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type SettingEditResponseZonesMinifyValueHTML string

const (
	SettingEditResponseZonesMinifyValueHTMLOn  SettingEditResponseZonesMinifyValueHTML = "on"
	SettingEditResponseZonesMinifyValueHTMLOff SettingEditResponseZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type SettingEditResponseZonesMinifyValueJs string

const (
	SettingEditResponseZonesMinifyValueJsOn  SettingEditResponseZonesMinifyValueJs = "on"
	SettingEditResponseZonesMinifyValueJsOff SettingEditResponseZonesMinifyValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesMinifyEditable bool

const (
	SettingEditResponseZonesMinifyEditableTrue  SettingEditResponseZonesMinifyEditable = true
	SettingEditResponseZonesMinifyEditableFalse SettingEditResponseZonesMinifyEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type SettingEditResponseZonesMirage struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesMirageID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesMirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesMirageJSON `json:"-"`
}

// settingEditResponseZonesMirageJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesMirage]
type settingEditResponseZonesMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesMirage) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesMirageID string

const (
	SettingEditResponseZonesMirageIDMirage SettingEditResponseZonesMirageID = "mirage"
)

// Current value of the zone setting.
type SettingEditResponseZonesMirageValue string

const (
	SettingEditResponseZonesMirageValueOn  SettingEditResponseZonesMirageValue = "on"
	SettingEditResponseZonesMirageValueOff SettingEditResponseZonesMirageValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesMirageEditable bool

const (
	SettingEditResponseZonesMirageEditableTrue  SettingEditResponseZonesMirageEditable = true
	SettingEditResponseZonesMirageEditableFalse SettingEditResponseZonesMirageEditable = false
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type SettingEditResponseZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID SettingEditResponseZonesMobileRedirectID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesMobileRedirectValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesMobileRedirectJSON `json:"-"`
}

// settingEditResponseZonesMobileRedirectJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesMobileRedirect]
type settingEditResponseZonesMobileRedirectJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesMobileRedirect) implementsSettingEditResponse() {}

// Identifier of the zone setting.
type SettingEditResponseZonesMobileRedirectID string

const (
	SettingEditResponseZonesMobileRedirectIDMobileRedirect SettingEditResponseZonesMobileRedirectID = "mobile_redirect"
)

// Current value of the zone setting.
type SettingEditResponseZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status SettingEditResponseZonesMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                            `json:"strip_uri"`
	JSON     settingEditResponseZonesMobileRedirectValueJSON `json:"-"`
}

// settingEditResponseZonesMobileRedirectValueJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesMobileRedirectValue]
type settingEditResponseZonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingEditResponseZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type SettingEditResponseZonesMobileRedirectValueStatus string

const (
	SettingEditResponseZonesMobileRedirectValueStatusOn  SettingEditResponseZonesMobileRedirectValueStatus = "on"
	SettingEditResponseZonesMobileRedirectValueStatusOff SettingEditResponseZonesMobileRedirectValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesMobileRedirectEditable bool

const (
	SettingEditResponseZonesMobileRedirectEditableTrue  SettingEditResponseZonesMobileRedirectEditable = true
	SettingEditResponseZonesMobileRedirectEditableFalse SettingEditResponseZonesMobileRedirectEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type SettingEditResponseZonesNEL struct {
	// Zone setting identifier.
	ID SettingEditResponseZonesNELID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesNELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesNELJSON `json:"-"`
}

// settingEditResponseZonesNELJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesNEL]
type settingEditResponseZonesNELJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesNEL) implementsSettingEditResponse() {}

// Zone setting identifier.
type SettingEditResponseZonesNELID string

const (
	SettingEditResponseZonesNELIDNEL SettingEditResponseZonesNELID = "nel"
)

// Current value of the zone setting.
type SettingEditResponseZonesNELValue struct {
	Enabled bool                                 `json:"enabled"`
	JSON    settingEditResponseZonesNELValueJSON `json:"-"`
}

// settingEditResponseZonesNELValueJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesNELValue]
type settingEditResponseZonesNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesNELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesNELEditable bool

const (
	SettingEditResponseZonesNELEditableTrue  SettingEditResponseZonesNELEditable = true
	SettingEditResponseZonesNELEditableFalse SettingEditResponseZonesNELEditable = false
)

// Enables the Opportunistic Encryption feature for a zone.
type SettingEditResponseZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesOpportunisticEncryptionJSON `json:"-"`
}

// settingEditResponseZonesOpportunisticEncryptionJSON contains the JSON metadata
// for the struct [SettingEditResponseZonesOpportunisticEncryption]
type settingEditResponseZonesOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesOpportunisticEncryption) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesOpportunisticEncryptionID string

const (
	SettingEditResponseZonesOpportunisticEncryptionIDOpportunisticEncryption SettingEditResponseZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Current value of the zone setting.
type SettingEditResponseZonesOpportunisticEncryptionValue string

const (
	SettingEditResponseZonesOpportunisticEncryptionValueOn  SettingEditResponseZonesOpportunisticEncryptionValue = "on"
	SettingEditResponseZonesOpportunisticEncryptionValueOff SettingEditResponseZonesOpportunisticEncryptionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesOpportunisticEncryptionEditable bool

const (
	SettingEditResponseZonesOpportunisticEncryptionEditableTrue  SettingEditResponseZonesOpportunisticEncryptionEditable = true
	SettingEditResponseZonesOpportunisticEncryptionEditableFalse SettingEditResponseZonesOpportunisticEncryptionEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type SettingEditResponseZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesOpportunisticOnionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesOpportunisticOnionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesOpportunisticOnionJSON `json:"-"`
}

// settingEditResponseZonesOpportunisticOnionJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesOpportunisticOnion]
type settingEditResponseZonesOpportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesOpportunisticOnion) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesOpportunisticOnionID string

const (
	SettingEditResponseZonesOpportunisticOnionIDOpportunisticOnion SettingEditResponseZonesOpportunisticOnionID = "opportunistic_onion"
)

// Current value of the zone setting.
type SettingEditResponseZonesOpportunisticOnionValue string

const (
	SettingEditResponseZonesOpportunisticOnionValueOn  SettingEditResponseZonesOpportunisticOnionValue = "on"
	SettingEditResponseZonesOpportunisticOnionValueOff SettingEditResponseZonesOpportunisticOnionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesOpportunisticOnionEditable bool

const (
	SettingEditResponseZonesOpportunisticOnionEditableTrue  SettingEditResponseZonesOpportunisticOnionEditable = true
	SettingEditResponseZonesOpportunisticOnionEditableFalse SettingEditResponseZonesOpportunisticOnionEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type SettingEditResponseZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesOrangeToOrangeID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesOrangeToOrangeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesOrangeToOrangeJSON `json:"-"`
}

// settingEditResponseZonesOrangeToOrangeJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesOrangeToOrange]
type settingEditResponseZonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesOrangeToOrange) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesOrangeToOrangeID string

const (
	SettingEditResponseZonesOrangeToOrangeIDOrangeToOrange SettingEditResponseZonesOrangeToOrangeID = "orange_to_orange"
)

// Current value of the zone setting.
type SettingEditResponseZonesOrangeToOrangeValue string

const (
	SettingEditResponseZonesOrangeToOrangeValueOn  SettingEditResponseZonesOrangeToOrangeValue = "on"
	SettingEditResponseZonesOrangeToOrangeValueOff SettingEditResponseZonesOrangeToOrangeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesOrangeToOrangeEditable bool

const (
	SettingEditResponseZonesOrangeToOrangeEditableTrue  SettingEditResponseZonesOrangeToOrangeEditable = true
	SettingEditResponseZonesOrangeToOrangeEditableFalse SettingEditResponseZonesOrangeToOrangeEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingEditResponseZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesOriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesOriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesOriginErrorPagePassThruJSON `json:"-"`
}

// settingEditResponseZonesOriginErrorPagePassThruJSON contains the JSON metadata
// for the struct [SettingEditResponseZonesOriginErrorPagePassThru]
type settingEditResponseZonesOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesOriginErrorPagePassThru) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesOriginErrorPagePassThruID string

const (
	SettingEditResponseZonesOriginErrorPagePassThruIDOriginErrorPagePassThru SettingEditResponseZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type SettingEditResponseZonesOriginErrorPagePassThruValue string

const (
	SettingEditResponseZonesOriginErrorPagePassThruValueOn  SettingEditResponseZonesOriginErrorPagePassThruValue = "on"
	SettingEditResponseZonesOriginErrorPagePassThruValueOff SettingEditResponseZonesOriginErrorPagePassThruValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesOriginErrorPagePassThruEditable bool

const (
	SettingEditResponseZonesOriginErrorPagePassThruEditableTrue  SettingEditResponseZonesOriginErrorPagePassThruEditable = true
	SettingEditResponseZonesOriginErrorPagePassThruEditableFalse SettingEditResponseZonesOriginErrorPagePassThruEditable = false
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type SettingEditResponseZonesPolish struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesPolishID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesPolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesPolishJSON `json:"-"`
}

// settingEditResponseZonesPolishJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesPolish]
type settingEditResponseZonesPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesPolish) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesPolishID string

const (
	SettingEditResponseZonesPolishIDPolish SettingEditResponseZonesPolishID = "polish"
)

// Current value of the zone setting.
type SettingEditResponseZonesPolishValue string

const (
	SettingEditResponseZonesPolishValueOff      SettingEditResponseZonesPolishValue = "off"
	SettingEditResponseZonesPolishValueLossless SettingEditResponseZonesPolishValue = "lossless"
	SettingEditResponseZonesPolishValueLossy    SettingEditResponseZonesPolishValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesPolishEditable bool

const (
	SettingEditResponseZonesPolishEditableTrue  SettingEditResponseZonesPolishEditable = true
	SettingEditResponseZonesPolishEditableFalse SettingEditResponseZonesPolishEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type SettingEditResponseZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesPrefetchPreloadID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesPrefetchPreloadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesPrefetchPreloadJSON `json:"-"`
}

// settingEditResponseZonesPrefetchPreloadJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesPrefetchPreload]
type settingEditResponseZonesPrefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesPrefetchPreload) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesPrefetchPreloadID string

const (
	SettingEditResponseZonesPrefetchPreloadIDPrefetchPreload SettingEditResponseZonesPrefetchPreloadID = "prefetch_preload"
)

// Current value of the zone setting.
type SettingEditResponseZonesPrefetchPreloadValue string

const (
	SettingEditResponseZonesPrefetchPreloadValueOn  SettingEditResponseZonesPrefetchPreloadValue = "on"
	SettingEditResponseZonesPrefetchPreloadValueOff SettingEditResponseZonesPrefetchPreloadValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesPrefetchPreloadEditable bool

const (
	SettingEditResponseZonesPrefetchPreloadEditableTrue  SettingEditResponseZonesPrefetchPreloadEditable = true
	SettingEditResponseZonesPrefetchPreloadEditableFalse SettingEditResponseZonesPrefetchPreloadEditable = false
)

// Maximum time between two read operations from origin.
type SettingEditResponseZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesProxyReadTimeoutJSON `json:"-"`
}

// settingEditResponseZonesProxyReadTimeoutJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesProxyReadTimeout]
type settingEditResponseZonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesProxyReadTimeout) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesProxyReadTimeoutID string

const (
	SettingEditResponseZonesProxyReadTimeoutIDProxyReadTimeout SettingEditResponseZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesProxyReadTimeoutEditable bool

const (
	SettingEditResponseZonesProxyReadTimeoutEditableTrue  SettingEditResponseZonesProxyReadTimeoutEditable = true
	SettingEditResponseZonesProxyReadTimeoutEditableFalse SettingEditResponseZonesProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type SettingEditResponseZonesPseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID SettingEditResponseZonesPseudoIPV4ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesPseudoIPV4Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesPseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesPseudoIPV4JSON `json:"-"`
}

// settingEditResponseZonesPseudoIPV4JSON contains the JSON metadata for the struct
// [SettingEditResponseZonesPseudoIPV4]
type settingEditResponseZonesPseudoIPV4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesPseudoIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesPseudoIPV4) implementsSettingEditResponse() {}

// Value of the Pseudo IPv4 setting.
type SettingEditResponseZonesPseudoIPV4ID string

const (
	SettingEditResponseZonesPseudoIPV4IDPseudoIPV4 SettingEditResponseZonesPseudoIPV4ID = "pseudo_ipv4"
)

// Current value of the zone setting.
type SettingEditResponseZonesPseudoIPV4Value string

const (
	SettingEditResponseZonesPseudoIPV4ValueOff             SettingEditResponseZonesPseudoIPV4Value = "off"
	SettingEditResponseZonesPseudoIPV4ValueAddHeader       SettingEditResponseZonesPseudoIPV4Value = "add_header"
	SettingEditResponseZonesPseudoIPV4ValueOverwriteHeader SettingEditResponseZonesPseudoIPV4Value = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesPseudoIPV4Editable bool

const (
	SettingEditResponseZonesPseudoIPV4EditableTrue  SettingEditResponseZonesPseudoIPV4Editable = true
	SettingEditResponseZonesPseudoIPV4EditableFalse SettingEditResponseZonesPseudoIPV4Editable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingEditResponseZonesResponseBuffering struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesResponseBufferingID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesResponseBufferingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesResponseBufferingJSON `json:"-"`
}

// settingEditResponseZonesResponseBufferingJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesResponseBuffering]
type settingEditResponseZonesResponseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesResponseBuffering) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesResponseBufferingID string

const (
	SettingEditResponseZonesResponseBufferingIDResponseBuffering SettingEditResponseZonesResponseBufferingID = "response_buffering"
)

// Current value of the zone setting.
type SettingEditResponseZonesResponseBufferingValue string

const (
	SettingEditResponseZonesResponseBufferingValueOn  SettingEditResponseZonesResponseBufferingValue = "on"
	SettingEditResponseZonesResponseBufferingValueOff SettingEditResponseZonesResponseBufferingValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesResponseBufferingEditable bool

const (
	SettingEditResponseZonesResponseBufferingEditableTrue  SettingEditResponseZonesResponseBufferingEditable = true
	SettingEditResponseZonesResponseBufferingEditableFalse SettingEditResponseZonesResponseBufferingEditable = false
)

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
type SettingEditResponseZonesRocketLoader struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesRocketLoaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesRocketLoaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesRocketLoaderJSON `json:"-"`
}

// settingEditResponseZonesRocketLoaderJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesRocketLoader]
type settingEditResponseZonesRocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesRocketLoader) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesRocketLoaderID string

const (
	SettingEditResponseZonesRocketLoaderIDRocketLoader SettingEditResponseZonesRocketLoaderID = "rocket_loader"
)

// Current value of the zone setting.
type SettingEditResponseZonesRocketLoaderValue string

const (
	SettingEditResponseZonesRocketLoaderValueOn  SettingEditResponseZonesRocketLoaderValue = "on"
	SettingEditResponseZonesRocketLoaderValueOff SettingEditResponseZonesRocketLoaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesRocketLoaderEditable bool

const (
	SettingEditResponseZonesRocketLoaderEditableTrue  SettingEditResponseZonesRocketLoaderEditable = true
	SettingEditResponseZonesRocketLoaderEditableFalse SettingEditResponseZonesRocketLoaderEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingEditResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSchemasAutomaticPlatformOptimizationValue `json:"value,required"`
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

func (r SettingEditResponseZonesSchemasAutomaticPlatformOptimization) implementsSettingEditResponse() {
}

// ID of the zone setting.
type SettingEditResponseZonesSchemasAutomaticPlatformOptimizationID string

const (
	SettingEditResponseZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization SettingEditResponseZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Current value of the zone setting.
type SettingEditResponseZonesSchemasAutomaticPlatformOptimizationValue struct {
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
	WpPlugin bool                                                                  `json:"wp_plugin,required"`
	JSON     settingEditResponseZonesSchemasAutomaticPlatformOptimizationValueJSON `json:"-"`
}

// settingEditResponseZonesSchemasAutomaticPlatformOptimizationValueJSON contains
// the JSON metadata for the struct
// [SettingEditResponseZonesSchemasAutomaticPlatformOptimizationValue]
type settingEditResponseZonesSchemasAutomaticPlatformOptimizationValueJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SettingEditResponseZonesSchemasAutomaticPlatformOptimizationValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Cloudflare security header for a zone.
type SettingEditResponseZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID SettingEditResponseZonesSecurityHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSecurityHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSecurityHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSecurityHeaderJSON `json:"-"`
}

// settingEditResponseZonesSecurityHeaderJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesSecurityHeader]
type settingEditResponseZonesSecurityHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesSecurityHeader) implementsSettingEditResponse() {}

// ID of the zone's security header.
type SettingEditResponseZonesSecurityHeaderID string

const (
	SettingEditResponseZonesSecurityHeaderIDSecurityHeader SettingEditResponseZonesSecurityHeaderID = "security_header"
)

// Current value of the zone setting.
type SettingEditResponseZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity SettingEditResponseZonesSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    settingEditResponseZonesSecurityHeaderValueJSON                    `json:"-"`
}

// settingEditResponseZonesSecurityHeaderValueJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesSecurityHeaderValue]
type settingEditResponseZonesSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingEditResponseZonesSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type SettingEditResponseZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                   `json:"nosniff"`
	JSON    settingEditResponseZonesSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// settingEditResponseZonesSecurityHeaderValueStrictTransportSecurityJSON contains
// the JSON metadata for the struct
// [SettingEditResponseZonesSecurityHeaderValueStrictTransportSecurity]
type settingEditResponseZonesSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SettingEditResponseZonesSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSecurityHeaderEditable bool

const (
	SettingEditResponseZonesSecurityHeaderEditableTrue  SettingEditResponseZonesSecurityHeaderEditable = true
	SettingEditResponseZonesSecurityHeaderEditableFalse SettingEditResponseZonesSecurityHeaderEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingEditResponseZonesSecurityLevel struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSecurityLevelJSON `json:"-"`
}

// settingEditResponseZonesSecurityLevelJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesSecurityLevel]
type settingEditResponseZonesSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesSecurityLevel) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSecurityLevelID string

const (
	SettingEditResponseZonesSecurityLevelIDSecurityLevel SettingEditResponseZonesSecurityLevelID = "security_level"
)

// Current value of the zone setting.
type SettingEditResponseZonesSecurityLevelValue string

const (
	SettingEditResponseZonesSecurityLevelValueOff            SettingEditResponseZonesSecurityLevelValue = "off"
	SettingEditResponseZonesSecurityLevelValueEssentiallyOff SettingEditResponseZonesSecurityLevelValue = "essentially_off"
	SettingEditResponseZonesSecurityLevelValueLow            SettingEditResponseZonesSecurityLevelValue = "low"
	SettingEditResponseZonesSecurityLevelValueMedium         SettingEditResponseZonesSecurityLevelValue = "medium"
	SettingEditResponseZonesSecurityLevelValueHigh           SettingEditResponseZonesSecurityLevelValue = "high"
	SettingEditResponseZonesSecurityLevelValueUnderAttack    SettingEditResponseZonesSecurityLevelValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSecurityLevelEditable bool

const (
	SettingEditResponseZonesSecurityLevelEditableTrue  SettingEditResponseZonesSecurityLevelEditable = true
	SettingEditResponseZonesSecurityLevelEditableFalse SettingEditResponseZonesSecurityLevelEditable = false
)

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
type SettingEditResponseZonesServerSideExclude struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesServerSideExcludeID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesServerSideExcludeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesServerSideExcludeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesServerSideExcludeJSON `json:"-"`
}

// settingEditResponseZonesServerSideExcludeJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesServerSideExclude]
type settingEditResponseZonesServerSideExcludeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesServerSideExclude) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesServerSideExcludeID string

const (
	SettingEditResponseZonesServerSideExcludeIDServerSideExclude SettingEditResponseZonesServerSideExcludeID = "server_side_exclude"
)

// Current value of the zone setting.
type SettingEditResponseZonesServerSideExcludeValue string

const (
	SettingEditResponseZonesServerSideExcludeValueOn  SettingEditResponseZonesServerSideExcludeValue = "on"
	SettingEditResponseZonesServerSideExcludeValueOff SettingEditResponseZonesServerSideExcludeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesServerSideExcludeEditable bool

const (
	SettingEditResponseZonesServerSideExcludeEditableTrue  SettingEditResponseZonesServerSideExcludeEditable = true
	SettingEditResponseZonesServerSideExcludeEditableFalse SettingEditResponseZonesServerSideExcludeEditable = false
)

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

func (r SettingEditResponseZonesSha1Support) implementsSettingEditResponse() {}

// Zone setting identifier.
type SettingEditResponseZonesSha1SupportID string

const (
	SettingEditResponseZonesSha1SupportIDSha1Support SettingEditResponseZonesSha1SupportID = "sha1_support"
)

// Current value of the zone setting.
type SettingEditResponseZonesSha1SupportValue string

const (
	SettingEditResponseZonesSha1SupportValueOff SettingEditResponseZonesSha1SupportValue = "off"
	SettingEditResponseZonesSha1SupportValueOn  SettingEditResponseZonesSha1SupportValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSha1SupportEditable bool

const (
	SettingEditResponseZonesSha1SupportEditableTrue  SettingEditResponseZonesSha1SupportEditable = true
	SettingEditResponseZonesSha1SupportEditableFalse SettingEditResponseZonesSha1SupportEditable = false
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type SettingEditResponseZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSortQueryStringForCacheID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSortQueryStringForCacheValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                           `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSortQueryStringForCacheJSON `json:"-"`
}

// settingEditResponseZonesSortQueryStringForCacheJSON contains the JSON metadata
// for the struct [SettingEditResponseZonesSortQueryStringForCache]
type settingEditResponseZonesSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesSortQueryStringForCache) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSortQueryStringForCacheID string

const (
	SettingEditResponseZonesSortQueryStringForCacheIDSortQueryStringForCache SettingEditResponseZonesSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type SettingEditResponseZonesSortQueryStringForCacheValue string

const (
	SettingEditResponseZonesSortQueryStringForCacheValueOn  SettingEditResponseZonesSortQueryStringForCacheValue = "on"
	SettingEditResponseZonesSortQueryStringForCacheValueOff SettingEditResponseZonesSortQueryStringForCacheValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSortQueryStringForCacheEditable bool

const (
	SettingEditResponseZonesSortQueryStringForCacheEditableTrue  SettingEditResponseZonesSortQueryStringForCacheEditable = true
	SettingEditResponseZonesSortQueryStringForCacheEditableFalse SettingEditResponseZonesSortQueryStringForCacheEditable = false
)

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
type SettingEditResponseZonesSSL struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSSLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesSSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesSSLJSON `json:"-"`
}

// settingEditResponseZonesSSLJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesSSL]
type settingEditResponseZonesSSLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesSSL) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesSSLID string

const (
	SettingEditResponseZonesSSLIDSSL SettingEditResponseZonesSSLID = "ssl"
)

// Current value of the zone setting.
type SettingEditResponseZonesSSLValue string

const (
	SettingEditResponseZonesSSLValueOff      SettingEditResponseZonesSSLValue = "off"
	SettingEditResponseZonesSSLValueFlexible SettingEditResponseZonesSSLValue = "flexible"
	SettingEditResponseZonesSSLValueFull     SettingEditResponseZonesSSLValue = "full"
	SettingEditResponseZonesSSLValueStrict   SettingEditResponseZonesSSLValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSSLEditable bool

const (
	SettingEditResponseZonesSSLEditableTrue  SettingEditResponseZonesSSLEditable = true
	SettingEditResponseZonesSSLEditableFalse SettingEditResponseZonesSSLEditable = false
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SettingEditResponseZonesSSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID SettingEditResponseZonesSSLRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                       `json:"enabled"`
	JSON    settingEditResponseZonesSSLRecommenderJSON `json:"-"`
}

// settingEditResponseZonesSSLRecommenderJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesSSLRecommender]
type settingEditResponseZonesSSLRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesSSLRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesSSLRecommender) implementsSettingEditResponse() {}

// Enrollment value for SSL/TLS Recommender.
type SettingEditResponseZonesSSLRecommenderID string

const (
	SettingEditResponseZonesSSLRecommenderIDSSLRecommender SettingEditResponseZonesSSLRecommenderID = "ssl_recommender"
)

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

func (r SettingEditResponseZonesTLS1_2Only) implementsSettingEditResponse() {}

// Zone setting identifier.
type SettingEditResponseZonesTLS1_2OnlyID string

const (
	SettingEditResponseZonesTLS1_2OnlyIDTLS1_2Only SettingEditResponseZonesTLS1_2OnlyID = "tls_1_2_only"
)

// Current value of the zone setting.
type SettingEditResponseZonesTLS1_2OnlyValue string

const (
	SettingEditResponseZonesTLS1_2OnlyValueOff SettingEditResponseZonesTLS1_2OnlyValue = "off"
	SettingEditResponseZonesTLS1_2OnlyValueOn  SettingEditResponseZonesTLS1_2OnlyValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesTLS1_2OnlyEditable bool

const (
	SettingEditResponseZonesTLS1_2OnlyEditableTrue  SettingEditResponseZonesTLS1_2OnlyEditable = true
	SettingEditResponseZonesTLS1_2OnlyEditableFalse SettingEditResponseZonesTLS1_2OnlyEditable = false
)

// Enables Crypto TLS 1.3 feature for a zone.
type SettingEditResponseZonesTLS1_3 struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesTLS1_3ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesTLS1_3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesTLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesTls1_3JSON `json:"-"`
}

// settingEditResponseZonesTls1_3JSON contains the JSON metadata for the struct
// [SettingEditResponseZonesTLS1_3]
type settingEditResponseZonesTls1_3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesTLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesTLS1_3) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesTLS1_3ID string

const (
	SettingEditResponseZonesTLS1_3IDTLS1_3 SettingEditResponseZonesTLS1_3ID = "tls_1_3"
)

// Current value of the zone setting.
type SettingEditResponseZonesTLS1_3Value string

const (
	SettingEditResponseZonesTLS1_3ValueOn  SettingEditResponseZonesTLS1_3Value = "on"
	SettingEditResponseZonesTLS1_3ValueOff SettingEditResponseZonesTLS1_3Value = "off"
	SettingEditResponseZonesTLS1_3ValueZrt SettingEditResponseZonesTLS1_3Value = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesTLS1_3Editable bool

const (
	SettingEditResponseZonesTLS1_3EditableTrue  SettingEditResponseZonesTLS1_3Editable = true
	SettingEditResponseZonesTLS1_3EditableFalse SettingEditResponseZonesTLS1_3Editable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type SettingEditResponseZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesTLSClientAuthID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesTLSClientAuthValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesTLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesTLSClientAuthJSON `json:"-"`
}

// settingEditResponseZonesTLSClientAuthJSON contains the JSON metadata for the
// struct [SettingEditResponseZonesTLSClientAuth]
type settingEditResponseZonesTLSClientAuthJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesTLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesTLSClientAuth) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesTLSClientAuthID string

const (
	SettingEditResponseZonesTLSClientAuthIDTLSClientAuth SettingEditResponseZonesTLSClientAuthID = "tls_client_auth"
)

// Current value of the zone setting.
type SettingEditResponseZonesTLSClientAuthValue string

const (
	SettingEditResponseZonesTLSClientAuthValueOn  SettingEditResponseZonesTLSClientAuthValue = "on"
	SettingEditResponseZonesTLSClientAuthValueOff SettingEditResponseZonesTLSClientAuthValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesTLSClientAuthEditable bool

const (
	SettingEditResponseZonesTLSClientAuthEditableTrue  SettingEditResponseZonesTLSClientAuthEditable = true
	SettingEditResponseZonesTLSClientAuthEditableFalse SettingEditResponseZonesTLSClientAuthEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingEditResponseZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesTrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesTrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesTrueClientIPHeaderJSON `json:"-"`
}

// settingEditResponseZonesTrueClientIPHeaderJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesTrueClientIPHeader]
type settingEditResponseZonesTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesTrueClientIPHeader) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesTrueClientIPHeaderID string

const (
	SettingEditResponseZonesTrueClientIPHeaderIDTrueClientIPHeader SettingEditResponseZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Current value of the zone setting.
type SettingEditResponseZonesTrueClientIPHeaderValue string

const (
	SettingEditResponseZonesTrueClientIPHeaderValueOn  SettingEditResponseZonesTrueClientIPHeaderValue = "on"
	SettingEditResponseZonesTrueClientIPHeaderValueOff SettingEditResponseZonesTrueClientIPHeaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesTrueClientIPHeaderEditable bool

const (
	SettingEditResponseZonesTrueClientIPHeaderEditableTrue  SettingEditResponseZonesTrueClientIPHeaderEditable = true
	SettingEditResponseZonesTrueClientIPHeaderEditableFalse SettingEditResponseZonesTrueClientIPHeaderEditable = false
)

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
type SettingEditResponseZonesWAF struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesWAFID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesWAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesWAFJSON `json:"-"`
}

// settingEditResponseZonesWAFJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesWAF]
type settingEditResponseZonesWAFJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesWAF) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesWAFID string

const (
	SettingEditResponseZonesWAFIDWAF SettingEditResponseZonesWAFID = "waf"
)

// Current value of the zone setting.
type SettingEditResponseZonesWAFValue string

const (
	SettingEditResponseZonesWAFValueOn  SettingEditResponseZonesWAFValue = "on"
	SettingEditResponseZonesWAFValueOff SettingEditResponseZonesWAFValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesWAFEditable bool

const (
	SettingEditResponseZonesWAFEditableTrue  SettingEditResponseZonesWAFEditable = true
	SettingEditResponseZonesWAFEditableFalse SettingEditResponseZonesWAFEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type SettingEditResponseZonesWebp struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesWebpID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesWebpValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesWebpJSON `json:"-"`
}

// settingEditResponseZonesWebpJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesWebp]
type settingEditResponseZonesWebpJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesWebp) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesWebpID string

const (
	SettingEditResponseZonesWebpIDWebp SettingEditResponseZonesWebpID = "webp"
)

// Current value of the zone setting.
type SettingEditResponseZonesWebpValue string

const (
	SettingEditResponseZonesWebpValueOff SettingEditResponseZonesWebpValue = "off"
	SettingEditResponseZonesWebpValueOn  SettingEditResponseZonesWebpValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesWebpEditable bool

const (
	SettingEditResponseZonesWebpEditableTrue  SettingEditResponseZonesWebpEditable = true
	SettingEditResponseZonesWebpEditableFalse SettingEditResponseZonesWebpEditable = false
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type SettingEditResponseZonesWebsockets struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesWebsocketsID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingEditResponseZonesWebsocketsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseZonesWebsocketsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       settingEditResponseZonesWebsocketsJSON `json:"-"`
}

// settingEditResponseZonesWebsocketsJSON contains the JSON metadata for the struct
// [SettingEditResponseZonesWebsockets]
type settingEditResponseZonesWebsocketsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseZonesWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r SettingEditResponseZonesWebsockets) implementsSettingEditResponse() {}

// ID of the zone setting.
type SettingEditResponseZonesWebsocketsID string

const (
	SettingEditResponseZonesWebsocketsIDWebsockets SettingEditResponseZonesWebsocketsID = "websockets"
)

// Current value of the zone setting.
type SettingEditResponseZonesWebsocketsValue string

const (
	SettingEditResponseZonesWebsocketsValueOff SettingEditResponseZonesWebsocketsValue = "off"
	SettingEditResponseZonesWebsocketsValueOn  SettingEditResponseZonesWebsocketsValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesWebsocketsEditable bool

const (
	SettingEditResponseZonesWebsocketsEditableTrue  SettingEditResponseZonesWebsocketsEditable = true
	SettingEditResponseZonesWebsocketsEditableFalse SettingEditResponseZonesWebsocketsEditable = false
)

type SettingListResponseEnvelope struct {
	Errors   []SettingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                            `json:"success,required"`
	Result  []SettingListResponse           `json:"result"`
	JSON    settingListResponseEnvelopeJSON `json:"-"`
}

// settingListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingListResponseEnvelope]
type settingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    settingListResponseEnvelopeErrorsJSON `json:"-"`
}

// settingListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SettingListResponseEnvelopeErrors]
type settingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    settingListResponseEnvelopeMessagesJSON `json:"-"`
}

// settingListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingListResponseEnvelopeMessages]
type settingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEditParams struct {
	// One or more zone setting objects. Must contain an ID and a value.
	Items param.Field[[]SettingEditParamsItem] `json:"items,required"`
}

func (r SettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// 0-RTT session resumption enabled for this zone.
//
// Satisfied by [SettingEditParamsItemsZones0rtt],
// [SettingEditParamsItemsZonesAdvancedDDOS],
// [SettingEditParamsItemsZonesAlwaysOnline],
// [SettingEditParamsItemsZonesAlwaysUseHTTPS],
// [SettingEditParamsItemsZonesAutomaticHTTPSRewrites],
// [SettingEditParamsItemsZonesBrotli],
// [SettingEditParamsItemsZonesBrowserCacheTTL],
// [SettingEditParamsItemsZonesBrowserCheck],
// [SettingEditParamsItemsZonesCacheLevel],
// [SettingEditParamsItemsZonesChallengeTTL], [SettingEditParamsItemsZonesCiphers],
// [SettingEditParamsItemsZonesCnameFlattening],
// [SettingEditParamsItemsZonesDevelopmentMode],
// [SettingEditParamsItemsZonesEarlyHints],
// [SettingEditParamsItemsZonesEdgeCacheTTL],
// [SettingEditParamsItemsZonesEmailObfuscation],
// [SettingEditParamsItemsZonesH2Prioritization],
// [SettingEditParamsItemsZonesHotlinkProtection],
// [SettingEditParamsItemsZonesHTTP2], [SettingEditParamsItemsZonesHTTP3],
// [SettingEditParamsItemsZonesImageResizing],
// [SettingEditParamsItemsZonesIPGeolocation], [SettingEditParamsItemsZonesIPV6],
// [SettingEditParamsItemsZonesMaxUpload],
// [SettingEditParamsItemsZonesMinTLSVersion], [SettingEditParamsItemsZonesMinify],
// [SettingEditParamsItemsZonesMirage],
// [SettingEditParamsItemsZonesMobileRedirect], [SettingEditParamsItemsZonesNEL],
// [SettingEditParamsItemsZonesOpportunisticEncryption],
// [SettingEditParamsItemsZonesOpportunisticOnion],
// [SettingEditParamsItemsZonesOrangeToOrange],
// [SettingEditParamsItemsZonesOriginErrorPagePassThru],
// [SettingEditParamsItemsZonesPolish],
// [SettingEditParamsItemsZonesPrefetchPreload],
// [SettingEditParamsItemsZonesProxyReadTimeout],
// [SettingEditParamsItemsZonesPseudoIPV4],
// [SettingEditParamsItemsZonesResponseBuffering],
// [SettingEditParamsItemsZonesRocketLoader],
// [SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization],
// [SettingEditParamsItemsZonesSecurityHeader],
// [SettingEditParamsItemsZonesSecurityLevel],
// [SettingEditParamsItemsZonesServerSideExclude],
// [SettingEditParamsItemsZonesSha1Support],
// [SettingEditParamsItemsZonesSortQueryStringForCache],
// [SettingEditParamsItemsZonesSSL], [SettingEditParamsItemsZonesSSLRecommender],
// [SettingEditParamsItemsZonesTLS1_2Only], [SettingEditParamsItemsZonesTLS1_3],
// [SettingEditParamsItemsZonesTLSClientAuth],
// [SettingEditParamsItemsZonesTrueClientIPHeader],
// [SettingEditParamsItemsZonesWAF], [SettingEditParamsItemsZonesWebp],
// [SettingEditParamsItemsZonesWebsockets].
type SettingEditParamsItem interface {
	implementsSettingEditParamsItem()
}

// 0-RTT session resumption enabled for this zone.
type SettingEditParamsItemsZones0rtt struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZones0rttID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZones0rttValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZones0rtt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZones0rtt) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZones0rttID string

const (
	SettingEditParamsItemsZones0rttID0rtt SettingEditParamsItemsZones0rttID = "0rtt"
)

// Current value of the zone setting.
type SettingEditParamsItemsZones0rttValue string

const (
	SettingEditParamsItemsZones0rttValueOn  SettingEditParamsItemsZones0rttValue = "on"
	SettingEditParamsItemsZones0rttValueOff SettingEditParamsItemsZones0rttValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZones0rttEditable bool

const (
	SettingEditParamsItemsZones0rttEditableTrue  SettingEditParamsItemsZones0rttEditable = true
	SettingEditParamsItemsZones0rttEditableFalse SettingEditParamsItemsZones0rttEditable = false
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type SettingEditParamsItemsZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesAdvancedDDOSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesAdvancedDDOSValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesAdvancedDDOS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesAdvancedDDOS) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesAdvancedDDOSID string

const (
	SettingEditParamsItemsZonesAdvancedDDOSIDAdvancedDDOS SettingEditParamsItemsZonesAdvancedDDOSID = "advanced_ddos"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesAdvancedDDOSValue string

const (
	SettingEditParamsItemsZonesAdvancedDDOSValueOn  SettingEditParamsItemsZonesAdvancedDDOSValue = "on"
	SettingEditParamsItemsZonesAdvancedDDOSValueOff SettingEditParamsItemsZonesAdvancedDDOSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesAdvancedDDOSEditable bool

const (
	SettingEditParamsItemsZonesAdvancedDDOSEditableTrue  SettingEditParamsItemsZonesAdvancedDDOSEditable = true
	SettingEditParamsItemsZonesAdvancedDDOSEditableFalse SettingEditParamsItemsZonesAdvancedDDOSEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type SettingEditParamsItemsZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesAlwaysOnlineID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesAlwaysOnlineValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesAlwaysOnline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesAlwaysOnline) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesAlwaysOnlineID string

const (
	SettingEditParamsItemsZonesAlwaysOnlineIDAlwaysOnline SettingEditParamsItemsZonesAlwaysOnlineID = "always_online"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesAlwaysOnlineValue string

const (
	SettingEditParamsItemsZonesAlwaysOnlineValueOn  SettingEditParamsItemsZonesAlwaysOnlineValue = "on"
	SettingEditParamsItemsZonesAlwaysOnlineValueOff SettingEditParamsItemsZonesAlwaysOnlineValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesAlwaysOnlineEditable bool

const (
	SettingEditParamsItemsZonesAlwaysOnlineEditableTrue  SettingEditParamsItemsZonesAlwaysOnlineEditable = true
	SettingEditParamsItemsZonesAlwaysOnlineEditableFalse SettingEditParamsItemsZonesAlwaysOnlineEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingEditParamsItemsZonesAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesAlwaysUseHTTPSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesAlwaysUseHTTPSValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesAlwaysUseHTTPS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesAlwaysUseHTTPS) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesAlwaysUseHTTPSID string

const (
	SettingEditParamsItemsZonesAlwaysUseHTTPSIDAlwaysUseHTTPS SettingEditParamsItemsZonesAlwaysUseHTTPSID = "always_use_https"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesAlwaysUseHTTPSValue string

const (
	SettingEditParamsItemsZonesAlwaysUseHTTPSValueOn  SettingEditParamsItemsZonesAlwaysUseHTTPSValue = "on"
	SettingEditParamsItemsZonesAlwaysUseHTTPSValueOff SettingEditParamsItemsZonesAlwaysUseHTTPSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesAlwaysUseHTTPSEditable bool

const (
	SettingEditParamsItemsZonesAlwaysUseHTTPSEditableTrue  SettingEditParamsItemsZonesAlwaysUseHTTPSEditable = true
	SettingEditParamsItemsZonesAlwaysUseHTTPSEditableFalse SettingEditParamsItemsZonesAlwaysUseHTTPSEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingEditParamsItemsZonesAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesAutomaticHTTPSRewritesID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesAutomaticHTTPSRewritesValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesAutomaticHTTPSRewrites) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesAutomaticHTTPSRewrites) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesAutomaticHTTPSRewritesID string

const (
	SettingEditParamsItemsZonesAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites SettingEditParamsItemsZonesAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesAutomaticHTTPSRewritesValue string

const (
	SettingEditParamsItemsZonesAutomaticHTTPSRewritesValueOn  SettingEditParamsItemsZonesAutomaticHTTPSRewritesValue = "on"
	SettingEditParamsItemsZonesAutomaticHTTPSRewritesValueOff SettingEditParamsItemsZonesAutomaticHTTPSRewritesValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesAutomaticHTTPSRewritesEditable bool

const (
	SettingEditParamsItemsZonesAutomaticHTTPSRewritesEditableTrue  SettingEditParamsItemsZonesAutomaticHTTPSRewritesEditable = true
	SettingEditParamsItemsZonesAutomaticHTTPSRewritesEditableFalse SettingEditParamsItemsZonesAutomaticHTTPSRewritesEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type SettingEditParamsItemsZonesBrotli struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesBrotliID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesBrotliValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesBrotli) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesBrotli) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesBrotliID string

const (
	SettingEditParamsItemsZonesBrotliIDBrotli SettingEditParamsItemsZonesBrotliID = "brotli"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesBrotliValue string

const (
	SettingEditParamsItemsZonesBrotliValueOff SettingEditParamsItemsZonesBrotliValue = "off"
	SettingEditParamsItemsZonesBrotliValueOn  SettingEditParamsItemsZonesBrotliValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesBrotliEditable bool

const (
	SettingEditParamsItemsZonesBrotliEditableTrue  SettingEditParamsItemsZonesBrotliEditable = true
	SettingEditParamsItemsZonesBrotliEditableFalse SettingEditParamsItemsZonesBrotliEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingEditParamsItemsZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesBrowserCacheTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesBrowserCacheTTLValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesBrowserCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesBrowserCacheTTL) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesBrowserCacheTTLID string

const (
	SettingEditParamsItemsZonesBrowserCacheTTLIDBrowserCacheTTL SettingEditParamsItemsZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesBrowserCacheTTLValue float64

const (
	SettingEditParamsItemsZonesBrowserCacheTTLValue0        SettingEditParamsItemsZonesBrowserCacheTTLValue = 0
	SettingEditParamsItemsZonesBrowserCacheTTLValue30       SettingEditParamsItemsZonesBrowserCacheTTLValue = 30
	SettingEditParamsItemsZonesBrowserCacheTTLValue60       SettingEditParamsItemsZonesBrowserCacheTTLValue = 60
	SettingEditParamsItemsZonesBrowserCacheTTLValue120      SettingEditParamsItemsZonesBrowserCacheTTLValue = 120
	SettingEditParamsItemsZonesBrowserCacheTTLValue300      SettingEditParamsItemsZonesBrowserCacheTTLValue = 300
	SettingEditParamsItemsZonesBrowserCacheTTLValue1200     SettingEditParamsItemsZonesBrowserCacheTTLValue = 1200
	SettingEditParamsItemsZonesBrowserCacheTTLValue1800     SettingEditParamsItemsZonesBrowserCacheTTLValue = 1800
	SettingEditParamsItemsZonesBrowserCacheTTLValue3600     SettingEditParamsItemsZonesBrowserCacheTTLValue = 3600
	SettingEditParamsItemsZonesBrowserCacheTTLValue7200     SettingEditParamsItemsZonesBrowserCacheTTLValue = 7200
	SettingEditParamsItemsZonesBrowserCacheTTLValue10800    SettingEditParamsItemsZonesBrowserCacheTTLValue = 10800
	SettingEditParamsItemsZonesBrowserCacheTTLValue14400    SettingEditParamsItemsZonesBrowserCacheTTLValue = 14400
	SettingEditParamsItemsZonesBrowserCacheTTLValue18000    SettingEditParamsItemsZonesBrowserCacheTTLValue = 18000
	SettingEditParamsItemsZonesBrowserCacheTTLValue28800    SettingEditParamsItemsZonesBrowserCacheTTLValue = 28800
	SettingEditParamsItemsZonesBrowserCacheTTLValue43200    SettingEditParamsItemsZonesBrowserCacheTTLValue = 43200
	SettingEditParamsItemsZonesBrowserCacheTTLValue57600    SettingEditParamsItemsZonesBrowserCacheTTLValue = 57600
	SettingEditParamsItemsZonesBrowserCacheTTLValue72000    SettingEditParamsItemsZonesBrowserCacheTTLValue = 72000
	SettingEditParamsItemsZonesBrowserCacheTTLValue86400    SettingEditParamsItemsZonesBrowserCacheTTLValue = 86400
	SettingEditParamsItemsZonesBrowserCacheTTLValue172800   SettingEditParamsItemsZonesBrowserCacheTTLValue = 172800
	SettingEditParamsItemsZonesBrowserCacheTTLValue259200   SettingEditParamsItemsZonesBrowserCacheTTLValue = 259200
	SettingEditParamsItemsZonesBrowserCacheTTLValue345600   SettingEditParamsItemsZonesBrowserCacheTTLValue = 345600
	SettingEditParamsItemsZonesBrowserCacheTTLValue432000   SettingEditParamsItemsZonesBrowserCacheTTLValue = 432000
	SettingEditParamsItemsZonesBrowserCacheTTLValue691200   SettingEditParamsItemsZonesBrowserCacheTTLValue = 691200
	SettingEditParamsItemsZonesBrowserCacheTTLValue1382400  SettingEditParamsItemsZonesBrowserCacheTTLValue = 1382400
	SettingEditParamsItemsZonesBrowserCacheTTLValue2073600  SettingEditParamsItemsZonesBrowserCacheTTLValue = 2073600
	SettingEditParamsItemsZonesBrowserCacheTTLValue2678400  SettingEditParamsItemsZonesBrowserCacheTTLValue = 2678400
	SettingEditParamsItemsZonesBrowserCacheTTLValue5356800  SettingEditParamsItemsZonesBrowserCacheTTLValue = 5356800
	SettingEditParamsItemsZonesBrowserCacheTTLValue16070400 SettingEditParamsItemsZonesBrowserCacheTTLValue = 16070400
	SettingEditParamsItemsZonesBrowserCacheTTLValue31536000 SettingEditParamsItemsZonesBrowserCacheTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesBrowserCacheTTLEditable bool

const (
	SettingEditParamsItemsZonesBrowserCacheTTLEditableTrue  SettingEditParamsItemsZonesBrowserCacheTTLEditable = true
	SettingEditParamsItemsZonesBrowserCacheTTLEditableFalse SettingEditParamsItemsZonesBrowserCacheTTLEditable = false
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type SettingEditParamsItemsZonesBrowserCheck struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesBrowserCheckID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesBrowserCheckValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesBrowserCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesBrowserCheck) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesBrowserCheckID string

const (
	SettingEditParamsItemsZonesBrowserCheckIDBrowserCheck SettingEditParamsItemsZonesBrowserCheckID = "browser_check"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesBrowserCheckValue string

const (
	SettingEditParamsItemsZonesBrowserCheckValueOn  SettingEditParamsItemsZonesBrowserCheckValue = "on"
	SettingEditParamsItemsZonesBrowserCheckValueOff SettingEditParamsItemsZonesBrowserCheckValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesBrowserCheckEditable bool

const (
	SettingEditParamsItemsZonesBrowserCheckEditableTrue  SettingEditParamsItemsZonesBrowserCheckEditable = true
	SettingEditParamsItemsZonesBrowserCheckEditableFalse SettingEditParamsItemsZonesBrowserCheckEditable = false
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type SettingEditParamsItemsZonesCacheLevel struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesCacheLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesCacheLevelValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesCacheLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesCacheLevel) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesCacheLevelID string

const (
	SettingEditParamsItemsZonesCacheLevelIDCacheLevel SettingEditParamsItemsZonesCacheLevelID = "cache_level"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesCacheLevelValue string

const (
	SettingEditParamsItemsZonesCacheLevelValueAggressive SettingEditParamsItemsZonesCacheLevelValue = "aggressive"
	SettingEditParamsItemsZonesCacheLevelValueBasic      SettingEditParamsItemsZonesCacheLevelValue = "basic"
	SettingEditParamsItemsZonesCacheLevelValueSimplified SettingEditParamsItemsZonesCacheLevelValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesCacheLevelEditable bool

const (
	SettingEditParamsItemsZonesCacheLevelEditableTrue  SettingEditParamsItemsZonesCacheLevelEditable = true
	SettingEditParamsItemsZonesCacheLevelEditableFalse SettingEditParamsItemsZonesCacheLevelEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type SettingEditParamsItemsZonesChallengeTTL struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesChallengeTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesChallengeTTLValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesChallengeTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesChallengeTTL) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesChallengeTTLID string

const (
	SettingEditParamsItemsZonesChallengeTTLIDChallengeTTL SettingEditParamsItemsZonesChallengeTTLID = "challenge_ttl"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesChallengeTTLValue float64

const (
	SettingEditParamsItemsZonesChallengeTTLValue300      SettingEditParamsItemsZonesChallengeTTLValue = 300
	SettingEditParamsItemsZonesChallengeTTLValue900      SettingEditParamsItemsZonesChallengeTTLValue = 900
	SettingEditParamsItemsZonesChallengeTTLValue1800     SettingEditParamsItemsZonesChallengeTTLValue = 1800
	SettingEditParamsItemsZonesChallengeTTLValue2700     SettingEditParamsItemsZonesChallengeTTLValue = 2700
	SettingEditParamsItemsZonesChallengeTTLValue3600     SettingEditParamsItemsZonesChallengeTTLValue = 3600
	SettingEditParamsItemsZonesChallengeTTLValue7200     SettingEditParamsItemsZonesChallengeTTLValue = 7200
	SettingEditParamsItemsZonesChallengeTTLValue10800    SettingEditParamsItemsZonesChallengeTTLValue = 10800
	SettingEditParamsItemsZonesChallengeTTLValue14400    SettingEditParamsItemsZonesChallengeTTLValue = 14400
	SettingEditParamsItemsZonesChallengeTTLValue28800    SettingEditParamsItemsZonesChallengeTTLValue = 28800
	SettingEditParamsItemsZonesChallengeTTLValue57600    SettingEditParamsItemsZonesChallengeTTLValue = 57600
	SettingEditParamsItemsZonesChallengeTTLValue86400    SettingEditParamsItemsZonesChallengeTTLValue = 86400
	SettingEditParamsItemsZonesChallengeTTLValue604800   SettingEditParamsItemsZonesChallengeTTLValue = 604800
	SettingEditParamsItemsZonesChallengeTTLValue2592000  SettingEditParamsItemsZonesChallengeTTLValue = 2592000
	SettingEditParamsItemsZonesChallengeTTLValue31536000 SettingEditParamsItemsZonesChallengeTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesChallengeTTLEditable bool

const (
	SettingEditParamsItemsZonesChallengeTTLEditableTrue  SettingEditParamsItemsZonesChallengeTTLEditable = true
	SettingEditParamsItemsZonesChallengeTTLEditableFalse SettingEditParamsItemsZonesChallengeTTLEditable = false
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type SettingEditParamsItemsZonesCiphers struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesCiphersID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[[]string] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesCiphers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesCiphers) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesCiphersID string

const (
	SettingEditParamsItemsZonesCiphersIDCiphers SettingEditParamsItemsZonesCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesCiphersEditable bool

const (
	SettingEditParamsItemsZonesCiphersEditableTrue  SettingEditParamsItemsZonesCiphersEditable = true
	SettingEditParamsItemsZonesCiphersEditableFalse SettingEditParamsItemsZonesCiphersEditable = false
)

// Whether or not cname flattening is on.
type SettingEditParamsItemsZonesCnameFlattening struct {
	// How to flatten the cname destination.
	ID param.Field[SettingEditParamsItemsZonesCnameFlatteningID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesCnameFlatteningValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesCnameFlattening) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesCnameFlattening) implementsSettingEditParamsItem() {}

// How to flatten the cname destination.
type SettingEditParamsItemsZonesCnameFlatteningID string

const (
	SettingEditParamsItemsZonesCnameFlatteningIDCnameFlattening SettingEditParamsItemsZonesCnameFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesCnameFlatteningValue string

const (
	SettingEditParamsItemsZonesCnameFlatteningValueFlattenAtRoot SettingEditParamsItemsZonesCnameFlatteningValue = "flatten_at_root"
	SettingEditParamsItemsZonesCnameFlatteningValueFlattenAll    SettingEditParamsItemsZonesCnameFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesCnameFlatteningEditable bool

const (
	SettingEditParamsItemsZonesCnameFlatteningEditableTrue  SettingEditParamsItemsZonesCnameFlatteningEditable = true
	SettingEditParamsItemsZonesCnameFlatteningEditableFalse SettingEditParamsItemsZonesCnameFlatteningEditable = false
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type SettingEditParamsItemsZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesDevelopmentModeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesDevelopmentModeValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesDevelopmentMode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesDevelopmentMode) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesDevelopmentModeID string

const (
	SettingEditParamsItemsZonesDevelopmentModeIDDevelopmentMode SettingEditParamsItemsZonesDevelopmentModeID = "development_mode"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesDevelopmentModeValue string

const (
	SettingEditParamsItemsZonesDevelopmentModeValueOn  SettingEditParamsItemsZonesDevelopmentModeValue = "on"
	SettingEditParamsItemsZonesDevelopmentModeValueOff SettingEditParamsItemsZonesDevelopmentModeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesDevelopmentModeEditable bool

const (
	SettingEditParamsItemsZonesDevelopmentModeEditableTrue  SettingEditParamsItemsZonesDevelopmentModeEditable = true
	SettingEditParamsItemsZonesDevelopmentModeEditableFalse SettingEditParamsItemsZonesDevelopmentModeEditable = false
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type SettingEditParamsItemsZonesEarlyHints struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesEarlyHintsID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesEarlyHintsValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesEarlyHints) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesEarlyHints) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesEarlyHintsID string

const (
	SettingEditParamsItemsZonesEarlyHintsIDEarlyHints SettingEditParamsItemsZonesEarlyHintsID = "early_hints"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesEarlyHintsValue string

const (
	SettingEditParamsItemsZonesEarlyHintsValueOn  SettingEditParamsItemsZonesEarlyHintsValue = "on"
	SettingEditParamsItemsZonesEarlyHintsValueOff SettingEditParamsItemsZonesEarlyHintsValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesEarlyHintsEditable bool

const (
	SettingEditParamsItemsZonesEarlyHintsEditableTrue  SettingEditParamsItemsZonesEarlyHintsEditable = true
	SettingEditParamsItemsZonesEarlyHintsEditableFalse SettingEditParamsItemsZonesEarlyHintsEditable = false
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type SettingEditParamsItemsZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesEdgeCacheTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesEdgeCacheTTLValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesEdgeCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesEdgeCacheTTL) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesEdgeCacheTTLID string

const (
	SettingEditParamsItemsZonesEdgeCacheTTLIDEdgeCacheTTL SettingEditParamsItemsZonesEdgeCacheTTLID = "edge_cache_ttl"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesEdgeCacheTTLValue float64

const (
	SettingEditParamsItemsZonesEdgeCacheTTLValue30     SettingEditParamsItemsZonesEdgeCacheTTLValue = 30
	SettingEditParamsItemsZonesEdgeCacheTTLValue60     SettingEditParamsItemsZonesEdgeCacheTTLValue = 60
	SettingEditParamsItemsZonesEdgeCacheTTLValue300    SettingEditParamsItemsZonesEdgeCacheTTLValue = 300
	SettingEditParamsItemsZonesEdgeCacheTTLValue1200   SettingEditParamsItemsZonesEdgeCacheTTLValue = 1200
	SettingEditParamsItemsZonesEdgeCacheTTLValue1800   SettingEditParamsItemsZonesEdgeCacheTTLValue = 1800
	SettingEditParamsItemsZonesEdgeCacheTTLValue3600   SettingEditParamsItemsZonesEdgeCacheTTLValue = 3600
	SettingEditParamsItemsZonesEdgeCacheTTLValue7200   SettingEditParamsItemsZonesEdgeCacheTTLValue = 7200
	SettingEditParamsItemsZonesEdgeCacheTTLValue10800  SettingEditParamsItemsZonesEdgeCacheTTLValue = 10800
	SettingEditParamsItemsZonesEdgeCacheTTLValue14400  SettingEditParamsItemsZonesEdgeCacheTTLValue = 14400
	SettingEditParamsItemsZonesEdgeCacheTTLValue18000  SettingEditParamsItemsZonesEdgeCacheTTLValue = 18000
	SettingEditParamsItemsZonesEdgeCacheTTLValue28800  SettingEditParamsItemsZonesEdgeCacheTTLValue = 28800
	SettingEditParamsItemsZonesEdgeCacheTTLValue43200  SettingEditParamsItemsZonesEdgeCacheTTLValue = 43200
	SettingEditParamsItemsZonesEdgeCacheTTLValue57600  SettingEditParamsItemsZonesEdgeCacheTTLValue = 57600
	SettingEditParamsItemsZonesEdgeCacheTTLValue72000  SettingEditParamsItemsZonesEdgeCacheTTLValue = 72000
	SettingEditParamsItemsZonesEdgeCacheTTLValue86400  SettingEditParamsItemsZonesEdgeCacheTTLValue = 86400
	SettingEditParamsItemsZonesEdgeCacheTTLValue172800 SettingEditParamsItemsZonesEdgeCacheTTLValue = 172800
	SettingEditParamsItemsZonesEdgeCacheTTLValue259200 SettingEditParamsItemsZonesEdgeCacheTTLValue = 259200
	SettingEditParamsItemsZonesEdgeCacheTTLValue345600 SettingEditParamsItemsZonesEdgeCacheTTLValue = 345600
	SettingEditParamsItemsZonesEdgeCacheTTLValue432000 SettingEditParamsItemsZonesEdgeCacheTTLValue = 432000
	SettingEditParamsItemsZonesEdgeCacheTTLValue518400 SettingEditParamsItemsZonesEdgeCacheTTLValue = 518400
	SettingEditParamsItemsZonesEdgeCacheTTLValue604800 SettingEditParamsItemsZonesEdgeCacheTTLValue = 604800
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesEdgeCacheTTLEditable bool

const (
	SettingEditParamsItemsZonesEdgeCacheTTLEditableTrue  SettingEditParamsItemsZonesEdgeCacheTTLEditable = true
	SettingEditParamsItemsZonesEdgeCacheTTLEditableFalse SettingEditParamsItemsZonesEdgeCacheTTLEditable = false
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingEditParamsItemsZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesEmailObfuscationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesEmailObfuscationValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesEmailObfuscation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesEmailObfuscation) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesEmailObfuscationID string

const (
	SettingEditParamsItemsZonesEmailObfuscationIDEmailObfuscation SettingEditParamsItemsZonesEmailObfuscationID = "email_obfuscation"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesEmailObfuscationValue string

const (
	SettingEditParamsItemsZonesEmailObfuscationValueOn  SettingEditParamsItemsZonesEmailObfuscationValue = "on"
	SettingEditParamsItemsZonesEmailObfuscationValueOff SettingEditParamsItemsZonesEmailObfuscationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesEmailObfuscationEditable bool

const (
	SettingEditParamsItemsZonesEmailObfuscationEditableTrue  SettingEditParamsItemsZonesEmailObfuscationEditable = true
	SettingEditParamsItemsZonesEmailObfuscationEditableFalse SettingEditParamsItemsZonesEmailObfuscationEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type SettingEditParamsItemsZonesH2Prioritization struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesH2PrioritizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesH2PrioritizationValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesH2Prioritization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesH2Prioritization) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesH2PrioritizationID string

const (
	SettingEditParamsItemsZonesH2PrioritizationIDH2Prioritization SettingEditParamsItemsZonesH2PrioritizationID = "h2_prioritization"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesH2PrioritizationValue string

const (
	SettingEditParamsItemsZonesH2PrioritizationValueOn     SettingEditParamsItemsZonesH2PrioritizationValue = "on"
	SettingEditParamsItemsZonesH2PrioritizationValueOff    SettingEditParamsItemsZonesH2PrioritizationValue = "off"
	SettingEditParamsItemsZonesH2PrioritizationValueCustom SettingEditParamsItemsZonesH2PrioritizationValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesH2PrioritizationEditable bool

const (
	SettingEditParamsItemsZonesH2PrioritizationEditableTrue  SettingEditParamsItemsZonesH2PrioritizationEditable = true
	SettingEditParamsItemsZonesH2PrioritizationEditableFalse SettingEditParamsItemsZonesH2PrioritizationEditable = false
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type SettingEditParamsItemsZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesHotlinkProtectionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesHotlinkProtectionValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesHotlinkProtection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesHotlinkProtection) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesHotlinkProtectionID string

const (
	SettingEditParamsItemsZonesHotlinkProtectionIDHotlinkProtection SettingEditParamsItemsZonesHotlinkProtectionID = "hotlink_protection"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesHotlinkProtectionValue string

const (
	SettingEditParamsItemsZonesHotlinkProtectionValueOn  SettingEditParamsItemsZonesHotlinkProtectionValue = "on"
	SettingEditParamsItemsZonesHotlinkProtectionValueOff SettingEditParamsItemsZonesHotlinkProtectionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesHotlinkProtectionEditable bool

const (
	SettingEditParamsItemsZonesHotlinkProtectionEditableTrue  SettingEditParamsItemsZonesHotlinkProtectionEditable = true
	SettingEditParamsItemsZonesHotlinkProtectionEditableFalse SettingEditParamsItemsZonesHotlinkProtectionEditable = false
)

// HTTP2 enabled for this zone.
type SettingEditParamsItemsZonesHTTP2 struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesHTTP2ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesHTTP2Value] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesHTTP2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesHTTP2) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesHTTP2ID string

const (
	SettingEditParamsItemsZonesHTTP2IDHTTP2 SettingEditParamsItemsZonesHTTP2ID = "http2"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesHTTP2Value string

const (
	SettingEditParamsItemsZonesHTTP2ValueOn  SettingEditParamsItemsZonesHTTP2Value = "on"
	SettingEditParamsItemsZonesHTTP2ValueOff SettingEditParamsItemsZonesHTTP2Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesHTTP2Editable bool

const (
	SettingEditParamsItemsZonesHTTP2EditableTrue  SettingEditParamsItemsZonesHTTP2Editable = true
	SettingEditParamsItemsZonesHTTP2EditableFalse SettingEditParamsItemsZonesHTTP2Editable = false
)

// HTTP3 enabled for this zone.
type SettingEditParamsItemsZonesHTTP3 struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesHTTP3ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesHTTP3Value] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesHTTP3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesHTTP3) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesHTTP3ID string

const (
	SettingEditParamsItemsZonesHTTP3IDHTTP3 SettingEditParamsItemsZonesHTTP3ID = "http3"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesHTTP3Value string

const (
	SettingEditParamsItemsZonesHTTP3ValueOn  SettingEditParamsItemsZonesHTTP3Value = "on"
	SettingEditParamsItemsZonesHTTP3ValueOff SettingEditParamsItemsZonesHTTP3Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesHTTP3Editable bool

const (
	SettingEditParamsItemsZonesHTTP3EditableTrue  SettingEditParamsItemsZonesHTTP3Editable = true
	SettingEditParamsItemsZonesHTTP3EditableFalse SettingEditParamsItemsZonesHTTP3Editable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type SettingEditParamsItemsZonesImageResizing struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesImageResizingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesImageResizingValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesImageResizing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesImageResizing) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesImageResizingID string

const (
	SettingEditParamsItemsZonesImageResizingIDImageResizing SettingEditParamsItemsZonesImageResizingID = "image_resizing"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesImageResizingValue string

const (
	SettingEditParamsItemsZonesImageResizingValueOn   SettingEditParamsItemsZonesImageResizingValue = "on"
	SettingEditParamsItemsZonesImageResizingValueOff  SettingEditParamsItemsZonesImageResizingValue = "off"
	SettingEditParamsItemsZonesImageResizingValueOpen SettingEditParamsItemsZonesImageResizingValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesImageResizingEditable bool

const (
	SettingEditParamsItemsZonesImageResizingEditableTrue  SettingEditParamsItemsZonesImageResizingEditable = true
	SettingEditParamsItemsZonesImageResizingEditableFalse SettingEditParamsItemsZonesImageResizingEditable = false
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingEditParamsItemsZonesIPGeolocation struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesIPGeolocationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesIPGeolocationValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesIPGeolocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesIPGeolocation) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesIPGeolocationID string

const (
	SettingEditParamsItemsZonesIPGeolocationIDIPGeolocation SettingEditParamsItemsZonesIPGeolocationID = "ip_geolocation"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesIPGeolocationValue string

const (
	SettingEditParamsItemsZonesIPGeolocationValueOn  SettingEditParamsItemsZonesIPGeolocationValue = "on"
	SettingEditParamsItemsZonesIPGeolocationValueOff SettingEditParamsItemsZonesIPGeolocationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesIPGeolocationEditable bool

const (
	SettingEditParamsItemsZonesIPGeolocationEditableTrue  SettingEditParamsItemsZonesIPGeolocationEditable = true
	SettingEditParamsItemsZonesIPGeolocationEditableFalse SettingEditParamsItemsZonesIPGeolocationEditable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type SettingEditParamsItemsZonesIPV6 struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesIPV6ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesIPV6Value] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesIPV6) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesIPV6ID string

const (
	SettingEditParamsItemsZonesIPV6IDIPV6 SettingEditParamsItemsZonesIPV6ID = "ipv6"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesIPV6Value string

const (
	SettingEditParamsItemsZonesIPV6ValueOff SettingEditParamsItemsZonesIPV6Value = "off"
	SettingEditParamsItemsZonesIPV6ValueOn  SettingEditParamsItemsZonesIPV6Value = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesIPV6Editable bool

const (
	SettingEditParamsItemsZonesIPV6EditableTrue  SettingEditParamsItemsZonesIPV6Editable = true
	SettingEditParamsItemsZonesIPV6EditableFalse SettingEditParamsItemsZonesIPV6Editable = false
)

// Maximum size of an allowable upload.
type SettingEditParamsItemsZonesMaxUpload struct {
	// identifier of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesMaxUploadID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesMaxUploadValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesMaxUpload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesMaxUpload) implementsSettingEditParamsItem() {}

// identifier of the zone setting.
type SettingEditParamsItemsZonesMaxUploadID string

const (
	SettingEditParamsItemsZonesMaxUploadIDMaxUpload SettingEditParamsItemsZonesMaxUploadID = "max_upload"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesMaxUploadValue float64

const (
	SettingEditParamsItemsZonesMaxUploadValue100 SettingEditParamsItemsZonesMaxUploadValue = 100
	SettingEditParamsItemsZonesMaxUploadValue200 SettingEditParamsItemsZonesMaxUploadValue = 200
	SettingEditParamsItemsZonesMaxUploadValue500 SettingEditParamsItemsZonesMaxUploadValue = 500
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesMaxUploadEditable bool

const (
	SettingEditParamsItemsZonesMaxUploadEditableTrue  SettingEditParamsItemsZonesMaxUploadEditable = true
	SettingEditParamsItemsZonesMaxUploadEditableFalse SettingEditParamsItemsZonesMaxUploadEditable = false
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type SettingEditParamsItemsZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesMinTLSVersionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesMinTLSVersionValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesMinTLSVersion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesMinTLSVersion) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesMinTLSVersionID string

const (
	SettingEditParamsItemsZonesMinTLSVersionIDMinTLSVersion SettingEditParamsItemsZonesMinTLSVersionID = "min_tls_version"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesMinTLSVersionValue string

const (
	SettingEditParamsItemsZonesMinTLSVersionValue1_0 SettingEditParamsItemsZonesMinTLSVersionValue = "1.0"
	SettingEditParamsItemsZonesMinTLSVersionValue1_1 SettingEditParamsItemsZonesMinTLSVersionValue = "1.1"
	SettingEditParamsItemsZonesMinTLSVersionValue1_2 SettingEditParamsItemsZonesMinTLSVersionValue = "1.2"
	SettingEditParamsItemsZonesMinTLSVersionValue1_3 SettingEditParamsItemsZonesMinTLSVersionValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesMinTLSVersionEditable bool

const (
	SettingEditParamsItemsZonesMinTLSVersionEditableTrue  SettingEditParamsItemsZonesMinTLSVersionEditable = true
	SettingEditParamsItemsZonesMinTLSVersionEditableFalse SettingEditParamsItemsZonesMinTLSVersionEditable = false
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type SettingEditParamsItemsZonesMinify struct {
	// Zone setting identifier.
	ID param.Field[SettingEditParamsItemsZonesMinifyID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesMinifyValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesMinify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesMinify) implementsSettingEditParamsItem() {}

// Zone setting identifier.
type SettingEditParamsItemsZonesMinifyID string

const (
	SettingEditParamsItemsZonesMinifyIDMinify SettingEditParamsItemsZonesMinifyID = "minify"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css param.Field[SettingEditParamsItemsZonesMinifyValueCss] `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML param.Field[SettingEditParamsItemsZonesMinifyValueHTML] `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js param.Field[SettingEditParamsItemsZonesMinifyValueJs] `json:"js"`
}

func (r SettingEditParamsItemsZonesMinifyValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Automatically minify all CSS files for your website.
type SettingEditParamsItemsZonesMinifyValueCss string

const (
	SettingEditParamsItemsZonesMinifyValueCssOn  SettingEditParamsItemsZonesMinifyValueCss = "on"
	SettingEditParamsItemsZonesMinifyValueCssOff SettingEditParamsItemsZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type SettingEditParamsItemsZonesMinifyValueHTML string

const (
	SettingEditParamsItemsZonesMinifyValueHTMLOn  SettingEditParamsItemsZonesMinifyValueHTML = "on"
	SettingEditParamsItemsZonesMinifyValueHTMLOff SettingEditParamsItemsZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type SettingEditParamsItemsZonesMinifyValueJs string

const (
	SettingEditParamsItemsZonesMinifyValueJsOn  SettingEditParamsItemsZonesMinifyValueJs = "on"
	SettingEditParamsItemsZonesMinifyValueJsOff SettingEditParamsItemsZonesMinifyValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesMinifyEditable bool

const (
	SettingEditParamsItemsZonesMinifyEditableTrue  SettingEditParamsItemsZonesMinifyEditable = true
	SettingEditParamsItemsZonesMinifyEditableFalse SettingEditParamsItemsZonesMinifyEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type SettingEditParamsItemsZonesMirage struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesMirageID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesMirageValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesMirage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesMirage) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesMirageID string

const (
	SettingEditParamsItemsZonesMirageIDMirage SettingEditParamsItemsZonesMirageID = "mirage"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesMirageValue string

const (
	SettingEditParamsItemsZonesMirageValueOn  SettingEditParamsItemsZonesMirageValue = "on"
	SettingEditParamsItemsZonesMirageValueOff SettingEditParamsItemsZonesMirageValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesMirageEditable bool

const (
	SettingEditParamsItemsZonesMirageEditableTrue  SettingEditParamsItemsZonesMirageEditable = true
	SettingEditParamsItemsZonesMirageEditableFalse SettingEditParamsItemsZonesMirageEditable = false
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type SettingEditParamsItemsZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesMobileRedirectID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesMobileRedirectValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesMobileRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesMobileRedirect) implementsSettingEditParamsItem() {}

// Identifier of the zone setting.
type SettingEditParamsItemsZonesMobileRedirectID string

const (
	SettingEditParamsItemsZonesMobileRedirectIDMobileRedirect SettingEditParamsItemsZonesMobileRedirectID = "mobile_redirect"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[SettingEditParamsItemsZonesMobileRedirectValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri param.Field[bool] `json:"strip_uri"`
}

func (r SettingEditParamsItemsZonesMobileRedirectValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not mobile redirect is enabled.
type SettingEditParamsItemsZonesMobileRedirectValueStatus string

const (
	SettingEditParamsItemsZonesMobileRedirectValueStatusOn  SettingEditParamsItemsZonesMobileRedirectValueStatus = "on"
	SettingEditParamsItemsZonesMobileRedirectValueStatusOff SettingEditParamsItemsZonesMobileRedirectValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesMobileRedirectEditable bool

const (
	SettingEditParamsItemsZonesMobileRedirectEditableTrue  SettingEditParamsItemsZonesMobileRedirectEditable = true
	SettingEditParamsItemsZonesMobileRedirectEditableFalse SettingEditParamsItemsZonesMobileRedirectEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type SettingEditParamsItemsZonesNEL struct {
	// Zone setting identifier.
	ID param.Field[SettingEditParamsItemsZonesNELID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesNELValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesNEL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesNEL) implementsSettingEditParamsItem() {}

// Zone setting identifier.
type SettingEditParamsItemsZonesNELID string

const (
	SettingEditParamsItemsZonesNELIDNEL SettingEditParamsItemsZonesNELID = "nel"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesNELValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SettingEditParamsItemsZonesNELValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesNELEditable bool

const (
	SettingEditParamsItemsZonesNELEditableTrue  SettingEditParamsItemsZonesNELEditable = true
	SettingEditParamsItemsZonesNELEditableFalse SettingEditParamsItemsZonesNELEditable = false
)

// Enables the Opportunistic Encryption feature for a zone.
type SettingEditParamsItemsZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesOpportunisticEncryptionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesOpportunisticEncryptionValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesOpportunisticEncryption) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesOpportunisticEncryption) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesOpportunisticEncryptionID string

const (
	SettingEditParamsItemsZonesOpportunisticEncryptionIDOpportunisticEncryption SettingEditParamsItemsZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesOpportunisticEncryptionValue string

const (
	SettingEditParamsItemsZonesOpportunisticEncryptionValueOn  SettingEditParamsItemsZonesOpportunisticEncryptionValue = "on"
	SettingEditParamsItemsZonesOpportunisticEncryptionValueOff SettingEditParamsItemsZonesOpportunisticEncryptionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesOpportunisticEncryptionEditable bool

const (
	SettingEditParamsItemsZonesOpportunisticEncryptionEditableTrue  SettingEditParamsItemsZonesOpportunisticEncryptionEditable = true
	SettingEditParamsItemsZonesOpportunisticEncryptionEditableFalse SettingEditParamsItemsZonesOpportunisticEncryptionEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type SettingEditParamsItemsZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesOpportunisticOnionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesOpportunisticOnionValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesOpportunisticOnion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesOpportunisticOnion) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesOpportunisticOnionID string

const (
	SettingEditParamsItemsZonesOpportunisticOnionIDOpportunisticOnion SettingEditParamsItemsZonesOpportunisticOnionID = "opportunistic_onion"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesOpportunisticOnionValue string

const (
	SettingEditParamsItemsZonesOpportunisticOnionValueOn  SettingEditParamsItemsZonesOpportunisticOnionValue = "on"
	SettingEditParamsItemsZonesOpportunisticOnionValueOff SettingEditParamsItemsZonesOpportunisticOnionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesOpportunisticOnionEditable bool

const (
	SettingEditParamsItemsZonesOpportunisticOnionEditableTrue  SettingEditParamsItemsZonesOpportunisticOnionEditable = true
	SettingEditParamsItemsZonesOpportunisticOnionEditableFalse SettingEditParamsItemsZonesOpportunisticOnionEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type SettingEditParamsItemsZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesOrangeToOrangeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesOrangeToOrangeValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesOrangeToOrange) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesOrangeToOrange) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesOrangeToOrangeID string

const (
	SettingEditParamsItemsZonesOrangeToOrangeIDOrangeToOrange SettingEditParamsItemsZonesOrangeToOrangeID = "orange_to_orange"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesOrangeToOrangeValue string

const (
	SettingEditParamsItemsZonesOrangeToOrangeValueOn  SettingEditParamsItemsZonesOrangeToOrangeValue = "on"
	SettingEditParamsItemsZonesOrangeToOrangeValueOff SettingEditParamsItemsZonesOrangeToOrangeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesOrangeToOrangeEditable bool

const (
	SettingEditParamsItemsZonesOrangeToOrangeEditableTrue  SettingEditParamsItemsZonesOrangeToOrangeEditable = true
	SettingEditParamsItemsZonesOrangeToOrangeEditableFalse SettingEditParamsItemsZonesOrangeToOrangeEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingEditParamsItemsZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesOriginErrorPagePassThruID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesOriginErrorPagePassThruValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesOriginErrorPagePassThru) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesOriginErrorPagePassThru) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesOriginErrorPagePassThruID string

const (
	SettingEditParamsItemsZonesOriginErrorPagePassThruIDOriginErrorPagePassThru SettingEditParamsItemsZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesOriginErrorPagePassThruValue string

const (
	SettingEditParamsItemsZonesOriginErrorPagePassThruValueOn  SettingEditParamsItemsZonesOriginErrorPagePassThruValue = "on"
	SettingEditParamsItemsZonesOriginErrorPagePassThruValueOff SettingEditParamsItemsZonesOriginErrorPagePassThruValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesOriginErrorPagePassThruEditable bool

const (
	SettingEditParamsItemsZonesOriginErrorPagePassThruEditableTrue  SettingEditParamsItemsZonesOriginErrorPagePassThruEditable = true
	SettingEditParamsItemsZonesOriginErrorPagePassThruEditableFalse SettingEditParamsItemsZonesOriginErrorPagePassThruEditable = false
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type SettingEditParamsItemsZonesPolish struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesPolishID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesPolishValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesPolish) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesPolish) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesPolishID string

const (
	SettingEditParamsItemsZonesPolishIDPolish SettingEditParamsItemsZonesPolishID = "polish"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesPolishValue string

const (
	SettingEditParamsItemsZonesPolishValueOff      SettingEditParamsItemsZonesPolishValue = "off"
	SettingEditParamsItemsZonesPolishValueLossless SettingEditParamsItemsZonesPolishValue = "lossless"
	SettingEditParamsItemsZonesPolishValueLossy    SettingEditParamsItemsZonesPolishValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesPolishEditable bool

const (
	SettingEditParamsItemsZonesPolishEditableTrue  SettingEditParamsItemsZonesPolishEditable = true
	SettingEditParamsItemsZonesPolishEditableFalse SettingEditParamsItemsZonesPolishEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type SettingEditParamsItemsZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesPrefetchPreloadID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesPrefetchPreloadValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesPrefetchPreload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesPrefetchPreload) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesPrefetchPreloadID string

const (
	SettingEditParamsItemsZonesPrefetchPreloadIDPrefetchPreload SettingEditParamsItemsZonesPrefetchPreloadID = "prefetch_preload"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesPrefetchPreloadValue string

const (
	SettingEditParamsItemsZonesPrefetchPreloadValueOn  SettingEditParamsItemsZonesPrefetchPreloadValue = "on"
	SettingEditParamsItemsZonesPrefetchPreloadValueOff SettingEditParamsItemsZonesPrefetchPreloadValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesPrefetchPreloadEditable bool

const (
	SettingEditParamsItemsZonesPrefetchPreloadEditableTrue  SettingEditParamsItemsZonesPrefetchPreloadEditable = true
	SettingEditParamsItemsZonesPrefetchPreloadEditableFalse SettingEditParamsItemsZonesPrefetchPreloadEditable = false
)

// Maximum time between two read operations from origin.
type SettingEditParamsItemsZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesProxyReadTimeoutID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[float64] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesProxyReadTimeout) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesProxyReadTimeoutID string

const (
	SettingEditParamsItemsZonesProxyReadTimeoutIDProxyReadTimeout SettingEditParamsItemsZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesProxyReadTimeoutEditable bool

const (
	SettingEditParamsItemsZonesProxyReadTimeoutEditableTrue  SettingEditParamsItemsZonesProxyReadTimeoutEditable = true
	SettingEditParamsItemsZonesProxyReadTimeoutEditableFalse SettingEditParamsItemsZonesProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type SettingEditParamsItemsZonesPseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID param.Field[SettingEditParamsItemsZonesPseudoIPV4ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesPseudoIPV4Value] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesPseudoIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesPseudoIPV4) implementsSettingEditParamsItem() {}

// Value of the Pseudo IPv4 setting.
type SettingEditParamsItemsZonesPseudoIPV4ID string

const (
	SettingEditParamsItemsZonesPseudoIPV4IDPseudoIPV4 SettingEditParamsItemsZonesPseudoIPV4ID = "pseudo_ipv4"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesPseudoIPV4Value string

const (
	SettingEditParamsItemsZonesPseudoIPV4ValueOff             SettingEditParamsItemsZonesPseudoIPV4Value = "off"
	SettingEditParamsItemsZonesPseudoIPV4ValueAddHeader       SettingEditParamsItemsZonesPseudoIPV4Value = "add_header"
	SettingEditParamsItemsZonesPseudoIPV4ValueOverwriteHeader SettingEditParamsItemsZonesPseudoIPV4Value = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesPseudoIPV4Editable bool

const (
	SettingEditParamsItemsZonesPseudoIPV4EditableTrue  SettingEditParamsItemsZonesPseudoIPV4Editable = true
	SettingEditParamsItemsZonesPseudoIPV4EditableFalse SettingEditParamsItemsZonesPseudoIPV4Editable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingEditParamsItemsZonesResponseBuffering struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesResponseBufferingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesResponseBufferingValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesResponseBuffering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesResponseBuffering) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesResponseBufferingID string

const (
	SettingEditParamsItemsZonesResponseBufferingIDResponseBuffering SettingEditParamsItemsZonesResponseBufferingID = "response_buffering"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesResponseBufferingValue string

const (
	SettingEditParamsItemsZonesResponseBufferingValueOn  SettingEditParamsItemsZonesResponseBufferingValue = "on"
	SettingEditParamsItemsZonesResponseBufferingValueOff SettingEditParamsItemsZonesResponseBufferingValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesResponseBufferingEditable bool

const (
	SettingEditParamsItemsZonesResponseBufferingEditableTrue  SettingEditParamsItemsZonesResponseBufferingEditable = true
	SettingEditParamsItemsZonesResponseBufferingEditableFalse SettingEditParamsItemsZonesResponseBufferingEditable = false
)

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
type SettingEditParamsItemsZonesRocketLoader struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesRocketLoaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesRocketLoaderValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesRocketLoader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesRocketLoader) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesRocketLoaderID string

const (
	SettingEditParamsItemsZonesRocketLoaderIDRocketLoader SettingEditParamsItemsZonesRocketLoaderID = "rocket_loader"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesRocketLoaderValue string

const (
	SettingEditParamsItemsZonesRocketLoaderValueOn  SettingEditParamsItemsZonesRocketLoaderValue = "on"
	SettingEditParamsItemsZonesRocketLoaderValueOff SettingEditParamsItemsZonesRocketLoaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesRocketLoaderEditable bool

const (
	SettingEditParamsItemsZonesRocketLoaderEditableTrue  SettingEditParamsItemsZonesRocketLoaderEditable = true
	SettingEditParamsItemsZonesRocketLoaderEditableFalse SettingEditParamsItemsZonesRocketLoaderEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization) implementsSettingEditParamsItem() {
}

// ID of the zone setting.
type SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID string

const (
	SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationValue struct {
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

func (r SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Cloudflare security header for a zone.
type SettingEditParamsItemsZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID param.Field[SettingEditParamsItemsZonesSecurityHeaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesSecurityHeaderValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesSecurityHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesSecurityHeader) implementsSettingEditParamsItem() {}

// ID of the zone's security header.
type SettingEditParamsItemsZonesSecurityHeaderID string

const (
	SettingEditParamsItemsZonesSecurityHeaderIDSecurityHeader SettingEditParamsItemsZonesSecurityHeaderID = "security_header"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[SettingEditParamsItemsZonesSecurityHeaderValueStrictTransportSecurity] `json:"strict_transport_security"`
}

func (r SettingEditParamsItemsZonesSecurityHeaderValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strict Transport Security.
type SettingEditParamsItemsZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge param.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff param.Field[bool] `json:"nosniff"`
}

func (r SettingEditParamsItemsZonesSecurityHeaderValueStrictTransportSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesSecurityHeaderEditable bool

const (
	SettingEditParamsItemsZonesSecurityHeaderEditableTrue  SettingEditParamsItemsZonesSecurityHeaderEditable = true
	SettingEditParamsItemsZonesSecurityHeaderEditableFalse SettingEditParamsItemsZonesSecurityHeaderEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingEditParamsItemsZonesSecurityLevel struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesSecurityLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesSecurityLevelValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesSecurityLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesSecurityLevel) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesSecurityLevelID string

const (
	SettingEditParamsItemsZonesSecurityLevelIDSecurityLevel SettingEditParamsItemsZonesSecurityLevelID = "security_level"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesSecurityLevelValue string

const (
	SettingEditParamsItemsZonesSecurityLevelValueOff            SettingEditParamsItemsZonesSecurityLevelValue = "off"
	SettingEditParamsItemsZonesSecurityLevelValueEssentiallyOff SettingEditParamsItemsZonesSecurityLevelValue = "essentially_off"
	SettingEditParamsItemsZonesSecurityLevelValueLow            SettingEditParamsItemsZonesSecurityLevelValue = "low"
	SettingEditParamsItemsZonesSecurityLevelValueMedium         SettingEditParamsItemsZonesSecurityLevelValue = "medium"
	SettingEditParamsItemsZonesSecurityLevelValueHigh           SettingEditParamsItemsZonesSecurityLevelValue = "high"
	SettingEditParamsItemsZonesSecurityLevelValueUnderAttack    SettingEditParamsItemsZonesSecurityLevelValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesSecurityLevelEditable bool

const (
	SettingEditParamsItemsZonesSecurityLevelEditableTrue  SettingEditParamsItemsZonesSecurityLevelEditable = true
	SettingEditParamsItemsZonesSecurityLevelEditableFalse SettingEditParamsItemsZonesSecurityLevelEditable = false
)

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
type SettingEditParamsItemsZonesServerSideExclude struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesServerSideExcludeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesServerSideExcludeValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesServerSideExclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesServerSideExclude) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesServerSideExcludeID string

const (
	SettingEditParamsItemsZonesServerSideExcludeIDServerSideExclude SettingEditParamsItemsZonesServerSideExcludeID = "server_side_exclude"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesServerSideExcludeValue string

const (
	SettingEditParamsItemsZonesServerSideExcludeValueOn  SettingEditParamsItemsZonesServerSideExcludeValue = "on"
	SettingEditParamsItemsZonesServerSideExcludeValueOff SettingEditParamsItemsZonesServerSideExcludeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesServerSideExcludeEditable bool

const (
	SettingEditParamsItemsZonesServerSideExcludeEditableTrue  SettingEditParamsItemsZonesServerSideExcludeEditable = true
	SettingEditParamsItemsZonesServerSideExcludeEditableFalse SettingEditParamsItemsZonesServerSideExcludeEditable = false
)

// Allow SHA1 support.
type SettingEditParamsItemsZonesSha1Support struct {
	// Zone setting identifier.
	ID param.Field[SettingEditParamsItemsZonesSha1SupportID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesSha1SupportValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesSha1Support) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesSha1Support) implementsSettingEditParamsItem() {}

// Zone setting identifier.
type SettingEditParamsItemsZonesSha1SupportID string

const (
	SettingEditParamsItemsZonesSha1SupportIDSha1Support SettingEditParamsItemsZonesSha1SupportID = "sha1_support"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesSha1SupportValue string

const (
	SettingEditParamsItemsZonesSha1SupportValueOff SettingEditParamsItemsZonesSha1SupportValue = "off"
	SettingEditParamsItemsZonesSha1SupportValueOn  SettingEditParamsItemsZonesSha1SupportValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesSha1SupportEditable bool

const (
	SettingEditParamsItemsZonesSha1SupportEditableTrue  SettingEditParamsItemsZonesSha1SupportEditable = true
	SettingEditParamsItemsZonesSha1SupportEditableFalse SettingEditParamsItemsZonesSha1SupportEditable = false
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type SettingEditParamsItemsZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesSortQueryStringForCacheID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesSortQueryStringForCacheValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesSortQueryStringForCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesSortQueryStringForCache) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesSortQueryStringForCacheID string

const (
	SettingEditParamsItemsZonesSortQueryStringForCacheIDSortQueryStringForCache SettingEditParamsItemsZonesSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesSortQueryStringForCacheValue string

const (
	SettingEditParamsItemsZonesSortQueryStringForCacheValueOn  SettingEditParamsItemsZonesSortQueryStringForCacheValue = "on"
	SettingEditParamsItemsZonesSortQueryStringForCacheValueOff SettingEditParamsItemsZonesSortQueryStringForCacheValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesSortQueryStringForCacheEditable bool

const (
	SettingEditParamsItemsZonesSortQueryStringForCacheEditableTrue  SettingEditParamsItemsZonesSortQueryStringForCacheEditable = true
	SettingEditParamsItemsZonesSortQueryStringForCacheEditableFalse SettingEditParamsItemsZonesSortQueryStringForCacheEditable = false
)

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
type SettingEditParamsItemsZonesSSL struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesSSLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesSSLValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesSSL) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesSSLID string

const (
	SettingEditParamsItemsZonesSSLIDSSL SettingEditParamsItemsZonesSSLID = "ssl"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesSSLValue string

const (
	SettingEditParamsItemsZonesSSLValueOff      SettingEditParamsItemsZonesSSLValue = "off"
	SettingEditParamsItemsZonesSSLValueFlexible SettingEditParamsItemsZonesSSLValue = "flexible"
	SettingEditParamsItemsZonesSSLValueFull     SettingEditParamsItemsZonesSSLValue = "full"
	SettingEditParamsItemsZonesSSLValueStrict   SettingEditParamsItemsZonesSSLValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesSSLEditable bool

const (
	SettingEditParamsItemsZonesSSLEditableTrue  SettingEditParamsItemsZonesSSLEditable = true
	SettingEditParamsItemsZonesSSLEditableFalse SettingEditParamsItemsZonesSSLEditable = false
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SettingEditParamsItemsZonesSSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[SettingEditParamsItemsZonesSSLRecommenderID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SettingEditParamsItemsZonesSSLRecommender) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesSSLRecommender) implementsSettingEditParamsItem() {}

// Enrollment value for SSL/TLS Recommender.
type SettingEditParamsItemsZonesSSLRecommenderID string

const (
	SettingEditParamsItemsZonesSSLRecommenderIDSSLRecommender SettingEditParamsItemsZonesSSLRecommenderID = "ssl_recommender"
)

// Only allows TLS1.2.
type SettingEditParamsItemsZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID param.Field[SettingEditParamsItemsZonesTLS1_2OnlyID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesTLS1_2OnlyValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesTLS1_2Only) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesTLS1_2Only) implementsSettingEditParamsItem() {}

// Zone setting identifier.
type SettingEditParamsItemsZonesTLS1_2OnlyID string

const (
	SettingEditParamsItemsZonesTLS1_2OnlyIDTLS1_2Only SettingEditParamsItemsZonesTLS1_2OnlyID = "tls_1_2_only"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesTLS1_2OnlyValue string

const (
	SettingEditParamsItemsZonesTLS1_2OnlyValueOff SettingEditParamsItemsZonesTLS1_2OnlyValue = "off"
	SettingEditParamsItemsZonesTLS1_2OnlyValueOn  SettingEditParamsItemsZonesTLS1_2OnlyValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesTLS1_2OnlyEditable bool

const (
	SettingEditParamsItemsZonesTLS1_2OnlyEditableTrue  SettingEditParamsItemsZonesTLS1_2OnlyEditable = true
	SettingEditParamsItemsZonesTLS1_2OnlyEditableFalse SettingEditParamsItemsZonesTLS1_2OnlyEditable = false
)

// Enables Crypto TLS 1.3 feature for a zone.
type SettingEditParamsItemsZonesTLS1_3 struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesTLS1_3ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesTLS1_3Value] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesTLS1_3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesTLS1_3) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesTLS1_3ID string

const (
	SettingEditParamsItemsZonesTLS1_3IDTLS1_3 SettingEditParamsItemsZonesTLS1_3ID = "tls_1_3"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesTLS1_3Value string

const (
	SettingEditParamsItemsZonesTLS1_3ValueOn  SettingEditParamsItemsZonesTLS1_3Value = "on"
	SettingEditParamsItemsZonesTLS1_3ValueOff SettingEditParamsItemsZonesTLS1_3Value = "off"
	SettingEditParamsItemsZonesTLS1_3ValueZrt SettingEditParamsItemsZonesTLS1_3Value = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesTLS1_3Editable bool

const (
	SettingEditParamsItemsZonesTLS1_3EditableTrue  SettingEditParamsItemsZonesTLS1_3Editable = true
	SettingEditParamsItemsZonesTLS1_3EditableFalse SettingEditParamsItemsZonesTLS1_3Editable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type SettingEditParamsItemsZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesTLSClientAuthID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesTLSClientAuthValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesTLSClientAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesTLSClientAuth) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesTLSClientAuthID string

const (
	SettingEditParamsItemsZonesTLSClientAuthIDTLSClientAuth SettingEditParamsItemsZonesTLSClientAuthID = "tls_client_auth"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesTLSClientAuthValue string

const (
	SettingEditParamsItemsZonesTLSClientAuthValueOn  SettingEditParamsItemsZonesTLSClientAuthValue = "on"
	SettingEditParamsItemsZonesTLSClientAuthValueOff SettingEditParamsItemsZonesTLSClientAuthValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesTLSClientAuthEditable bool

const (
	SettingEditParamsItemsZonesTLSClientAuthEditableTrue  SettingEditParamsItemsZonesTLSClientAuthEditable = true
	SettingEditParamsItemsZonesTLSClientAuthEditableFalse SettingEditParamsItemsZonesTLSClientAuthEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingEditParamsItemsZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesTrueClientIPHeaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesTrueClientIPHeaderValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesTrueClientIPHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesTrueClientIPHeader) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesTrueClientIPHeaderID string

const (
	SettingEditParamsItemsZonesTrueClientIPHeaderIDTrueClientIPHeader SettingEditParamsItemsZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesTrueClientIPHeaderValue string

const (
	SettingEditParamsItemsZonesTrueClientIPHeaderValueOn  SettingEditParamsItemsZonesTrueClientIPHeaderValue = "on"
	SettingEditParamsItemsZonesTrueClientIPHeaderValueOff SettingEditParamsItemsZonesTrueClientIPHeaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesTrueClientIPHeaderEditable bool

const (
	SettingEditParamsItemsZonesTrueClientIPHeaderEditableTrue  SettingEditParamsItemsZonesTrueClientIPHeaderEditable = true
	SettingEditParamsItemsZonesTrueClientIPHeaderEditableFalse SettingEditParamsItemsZonesTrueClientIPHeaderEditable = false
)

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
type SettingEditParamsItemsZonesWAF struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesWAFID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesWAFValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesWAF) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesWAF) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesWAFID string

const (
	SettingEditParamsItemsZonesWAFIDWAF SettingEditParamsItemsZonesWAFID = "waf"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesWAFValue string

const (
	SettingEditParamsItemsZonesWAFValueOn  SettingEditParamsItemsZonesWAFValue = "on"
	SettingEditParamsItemsZonesWAFValueOff SettingEditParamsItemsZonesWAFValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesWAFEditable bool

const (
	SettingEditParamsItemsZonesWAFEditableTrue  SettingEditParamsItemsZonesWAFEditable = true
	SettingEditParamsItemsZonesWAFEditableFalse SettingEditParamsItemsZonesWAFEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type SettingEditParamsItemsZonesWebp struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesWebpID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesWebpValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesWebp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesWebp) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesWebpID string

const (
	SettingEditParamsItemsZonesWebpIDWebp SettingEditParamsItemsZonesWebpID = "webp"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesWebpValue string

const (
	SettingEditParamsItemsZonesWebpValueOff SettingEditParamsItemsZonesWebpValue = "off"
	SettingEditParamsItemsZonesWebpValueOn  SettingEditParamsItemsZonesWebpValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesWebpEditable bool

const (
	SettingEditParamsItemsZonesWebpEditableTrue  SettingEditParamsItemsZonesWebpEditable = true
	SettingEditParamsItemsZonesWebpEditableFalse SettingEditParamsItemsZonesWebpEditable = false
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type SettingEditParamsItemsZonesWebsockets struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesWebsocketsID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesWebsocketsValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesWebsockets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesWebsockets) implementsSettingEditParamsItem() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesWebsocketsID string

const (
	SettingEditParamsItemsZonesWebsocketsIDWebsockets SettingEditParamsItemsZonesWebsocketsID = "websockets"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesWebsocketsValue string

const (
	SettingEditParamsItemsZonesWebsocketsValueOff SettingEditParamsItemsZonesWebsocketsValue = "off"
	SettingEditParamsItemsZonesWebsocketsValueOn  SettingEditParamsItemsZonesWebsocketsValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesWebsocketsEditable bool

const (
	SettingEditParamsItemsZonesWebsocketsEditableTrue  SettingEditParamsItemsZonesWebsocketsEditable = true
	SettingEditParamsItemsZonesWebsocketsEditableFalse SettingEditParamsItemsZonesWebsocketsEditable = false
)

type SettingEditResponseEnvelope struct {
	Errors   []SettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                            `json:"success,required"`
	Result  []SettingEditResponse           `json:"result"`
	JSON    settingEditResponseEnvelopeJSON `json:"-"`
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

type SettingEditResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    settingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SettingEditResponseEnvelopeErrors]
type settingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SettingEditResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    settingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingEditResponseEnvelopeMessages]
type settingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
