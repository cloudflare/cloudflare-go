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

// Available settings for your user in relation to a zone.
func (r *ZoneSettingService) List(ctx context.Context, query ZoneSettingListParams, opts ...option.RequestOption) (res *[]ZoneSettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZoneSettingListResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [ZoneSettingListResponseZones0rtt],
// [ZoneSettingListResponseZonesAdvancedDDOS],
// [ZoneSettingListResponseZonesAlwaysOnline],
// [ZoneSettingListResponseZonesAlwaysUseHTTPS],
// [ZoneSettingListResponseZonesAutomaticHTTPSRewrites],
// [ZoneSettingListResponseZonesBrotli],
// [ZoneSettingListResponseZonesBrowserCacheTTL],
// [ZoneSettingListResponseZonesBrowserCheck],
// [ZoneSettingListResponseZonesCacheLevel],
// [ZoneSettingListResponseZonesChallengeTTL],
// [ZoneSettingListResponseZonesCiphers],
// [ZoneSettingListResponseZonesCnameFlattening],
// [ZoneSettingListResponseZonesDevelopmentMode],
// [ZoneSettingListResponseZonesEarlyHints],
// [ZoneSettingListResponseZonesEdgeCacheTTL],
// [ZoneSettingListResponseZonesEmailObfuscation],
// [ZoneSettingListResponseZonesH2Prioritization],
// [ZoneSettingListResponseZonesHotlinkProtection],
// [ZoneSettingListResponseZonesHTTP2], [ZoneSettingListResponseZonesHTTP3],
// [ZoneSettingListResponseZonesImageResizing],
// [ZoneSettingListResponseZonesIPGeolocation], [ZoneSettingListResponseZonesIPV6],
// [ZoneSettingListResponseZonesMaxUpload],
// [ZoneSettingListResponseZonesMinTLSVersion],
// [ZoneSettingListResponseZonesMinify], [ZoneSettingListResponseZonesMirage],
// [ZoneSettingListResponseZonesMobileRedirect], [ZoneSettingListResponseZonesNEL],
// [ZoneSettingListResponseZonesOpportunisticEncryption],
// [ZoneSettingListResponseZonesOpportunisticOnion],
// [ZoneSettingListResponseZonesOrangeToOrange],
// [ZoneSettingListResponseZonesOriginErrorPagePassThru],
// [ZoneSettingListResponseZonesPolish],
// [ZoneSettingListResponseZonesPrefetchPreload],
// [ZoneSettingListResponseZonesProxyReadTimeout],
// [ZoneSettingListResponseZonesPseudoIPV4],
// [ZoneSettingListResponseZonesResponseBuffering],
// [ZoneSettingListResponseZonesRocketLoader],
// [ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimization],
// [ZoneSettingListResponseZonesSecurityHeader],
// [ZoneSettingListResponseZonesSecurityLevel],
// [ZoneSettingListResponseZonesServerSideExclude],
// [ZoneSettingListResponseZonesSha1Support],
// [ZoneSettingListResponseZonesSortQueryStringForCache],
// [ZoneSettingListResponseZonesSSL], [ZoneSettingListResponseZonesSSLRecommender],
// [ZoneSettingListResponseZonesTLS1_2Only], [ZoneSettingListResponseZonesTLS1_3],
// [ZoneSettingListResponseZonesTLSClientAuth],
// [ZoneSettingListResponseZonesTrueClientIPHeader],
// [ZoneSettingListResponseZonesWAF], [ZoneSettingListResponseZonesWebp] or
// [ZoneSettingListResponseZonesWebsockets].
type ZoneSettingListResponse interface {
	implementsZoneSettingListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneSettingListResponse)(nil)).Elem(), "")
}

// 0-RTT session resumption enabled for this zone.
type ZoneSettingListResponseZones0rtt struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZones0rttID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZones0rttValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZones0rttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZones0rttJSON `json:"-"`
}

// zoneSettingListResponseZones0rttJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZones0rtt]
type zoneSettingListResponseZones0rttJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZones0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZones0rtt) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZones0rttID string

const (
	ZoneSettingListResponseZones0rttID0rtt ZoneSettingListResponseZones0rttID = "0rtt"
)

// Current value of the zone setting.
type ZoneSettingListResponseZones0rttValue string

const (
	ZoneSettingListResponseZones0rttValueOn  ZoneSettingListResponseZones0rttValue = "on"
	ZoneSettingListResponseZones0rttValueOff ZoneSettingListResponseZones0rttValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZones0rttEditable bool

const (
	ZoneSettingListResponseZones0rttEditableTrue  ZoneSettingListResponseZones0rttEditable = true
	ZoneSettingListResponseZones0rttEditableFalse ZoneSettingListResponseZones0rttEditable = false
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingListResponseZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesAdvancedDDOSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesAdvancedDDOSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesAdvancedDDOSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesAdvancedDDOSJSON `json:"-"`
}

// zoneSettingListResponseZonesAdvancedDDOSJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesAdvancedDDOS]
type zoneSettingListResponseZonesAdvancedDDOSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesAdvancedDDOS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesAdvancedDDOS) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesAdvancedDDOSID string

const (
	ZoneSettingListResponseZonesAdvancedDDOSIDAdvancedDDOS ZoneSettingListResponseZonesAdvancedDDOSID = "advanced_ddos"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesAdvancedDDOSValue string

const (
	ZoneSettingListResponseZonesAdvancedDDOSValueOn  ZoneSettingListResponseZonesAdvancedDDOSValue = "on"
	ZoneSettingListResponseZonesAdvancedDDOSValueOff ZoneSettingListResponseZonesAdvancedDDOSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesAdvancedDDOSEditable bool

const (
	ZoneSettingListResponseZonesAdvancedDDOSEditableTrue  ZoneSettingListResponseZonesAdvancedDDOSEditable = true
	ZoneSettingListResponseZonesAdvancedDDOSEditableFalse ZoneSettingListResponseZonesAdvancedDDOSEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingListResponseZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesAlwaysOnlineID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesAlwaysOnlineValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesAlwaysOnlineJSON `json:"-"`
}

// zoneSettingListResponseZonesAlwaysOnlineJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesAlwaysOnline]
type zoneSettingListResponseZonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesAlwaysOnline) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesAlwaysOnlineID string

const (
	ZoneSettingListResponseZonesAlwaysOnlineIDAlwaysOnline ZoneSettingListResponseZonesAlwaysOnlineID = "always_online"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesAlwaysOnlineValue string

const (
	ZoneSettingListResponseZonesAlwaysOnlineValueOn  ZoneSettingListResponseZonesAlwaysOnlineValue = "on"
	ZoneSettingListResponseZonesAlwaysOnlineValueOff ZoneSettingListResponseZonesAlwaysOnlineValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesAlwaysOnlineEditable bool

const (
	ZoneSettingListResponseZonesAlwaysOnlineEditableTrue  ZoneSettingListResponseZonesAlwaysOnlineEditable = true
	ZoneSettingListResponseZonesAlwaysOnlineEditableFalse ZoneSettingListResponseZonesAlwaysOnlineEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingListResponseZonesAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesAlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesAlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesAlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesAlwaysUseHTTPSJSON `json:"-"`
}

// zoneSettingListResponseZonesAlwaysUseHTTPSJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesAlwaysUseHTTPS]
type zoneSettingListResponseZonesAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesAlwaysUseHTTPS) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesAlwaysUseHTTPSID string

const (
	ZoneSettingListResponseZonesAlwaysUseHTTPSIDAlwaysUseHTTPS ZoneSettingListResponseZonesAlwaysUseHTTPSID = "always_use_https"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesAlwaysUseHTTPSValue string

const (
	ZoneSettingListResponseZonesAlwaysUseHTTPSValueOn  ZoneSettingListResponseZonesAlwaysUseHTTPSValue = "on"
	ZoneSettingListResponseZonesAlwaysUseHTTPSValueOff ZoneSettingListResponseZonesAlwaysUseHTTPSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesAlwaysUseHTTPSEditable bool

const (
	ZoneSettingListResponseZonesAlwaysUseHTTPSEditableTrue  ZoneSettingListResponseZonesAlwaysUseHTTPSEditable = true
	ZoneSettingListResponseZonesAlwaysUseHTTPSEditableFalse ZoneSettingListResponseZonesAlwaysUseHTTPSEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingListResponseZonesAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesAutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesAutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesAutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesAutomaticHTTPSRewritesJSON `json:"-"`
}

// zoneSettingListResponseZonesAutomaticHTTPSRewritesJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseZonesAutomaticHTTPSRewrites]
type zoneSettingListResponseZonesAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesAutomaticHTTPSRewrites) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesAutomaticHTTPSRewritesID string

const (
	ZoneSettingListResponseZonesAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites ZoneSettingListResponseZonesAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesAutomaticHTTPSRewritesValue string

const (
	ZoneSettingListResponseZonesAutomaticHTTPSRewritesValueOn  ZoneSettingListResponseZonesAutomaticHTTPSRewritesValue = "on"
	ZoneSettingListResponseZonesAutomaticHTTPSRewritesValueOff ZoneSettingListResponseZonesAutomaticHTTPSRewritesValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesAutomaticHTTPSRewritesEditable bool

const (
	ZoneSettingListResponseZonesAutomaticHTTPSRewritesEditableTrue  ZoneSettingListResponseZonesAutomaticHTTPSRewritesEditable = true
	ZoneSettingListResponseZonesAutomaticHTTPSRewritesEditableFalse ZoneSettingListResponseZonesAutomaticHTTPSRewritesEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingListResponseZonesBrotli struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesBrotliID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesBrotliValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesBrotliJSON `json:"-"`
}

// zoneSettingListResponseZonesBrotliJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesBrotli]
type zoneSettingListResponseZonesBrotliJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesBrotli) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesBrotliID string

const (
	ZoneSettingListResponseZonesBrotliIDBrotli ZoneSettingListResponseZonesBrotliID = "brotli"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesBrotliValue string

const (
	ZoneSettingListResponseZonesBrotliValueOff ZoneSettingListResponseZonesBrotliValue = "off"
	ZoneSettingListResponseZonesBrotliValueOn  ZoneSettingListResponseZonesBrotliValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesBrotliEditable bool

const (
	ZoneSettingListResponseZonesBrotliEditableTrue  ZoneSettingListResponseZonesBrotliEditable = true
	ZoneSettingListResponseZonesBrotliEditableFalse ZoneSettingListResponseZonesBrotliEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingListResponseZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesBrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesBrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesBrowserCacheTTLJSON `json:"-"`
}

// zoneSettingListResponseZonesBrowserCacheTTLJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesBrowserCacheTTL]
type zoneSettingListResponseZonesBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesBrowserCacheTTL) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesBrowserCacheTTLID string

const (
	ZoneSettingListResponseZonesBrowserCacheTTLIDBrowserCacheTTL ZoneSettingListResponseZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesBrowserCacheTTLValue float64

const (
	ZoneSettingListResponseZonesBrowserCacheTTLValue0        ZoneSettingListResponseZonesBrowserCacheTTLValue = 0
	ZoneSettingListResponseZonesBrowserCacheTTLValue30       ZoneSettingListResponseZonesBrowserCacheTTLValue = 30
	ZoneSettingListResponseZonesBrowserCacheTTLValue60       ZoneSettingListResponseZonesBrowserCacheTTLValue = 60
	ZoneSettingListResponseZonesBrowserCacheTTLValue120      ZoneSettingListResponseZonesBrowserCacheTTLValue = 120
	ZoneSettingListResponseZonesBrowserCacheTTLValue300      ZoneSettingListResponseZonesBrowserCacheTTLValue = 300
	ZoneSettingListResponseZonesBrowserCacheTTLValue1200     ZoneSettingListResponseZonesBrowserCacheTTLValue = 1200
	ZoneSettingListResponseZonesBrowserCacheTTLValue1800     ZoneSettingListResponseZonesBrowserCacheTTLValue = 1800
	ZoneSettingListResponseZonesBrowserCacheTTLValue3600     ZoneSettingListResponseZonesBrowserCacheTTLValue = 3600
	ZoneSettingListResponseZonesBrowserCacheTTLValue7200     ZoneSettingListResponseZonesBrowserCacheTTLValue = 7200
	ZoneSettingListResponseZonesBrowserCacheTTLValue10800    ZoneSettingListResponseZonesBrowserCacheTTLValue = 10800
	ZoneSettingListResponseZonesBrowserCacheTTLValue14400    ZoneSettingListResponseZonesBrowserCacheTTLValue = 14400
	ZoneSettingListResponseZonesBrowserCacheTTLValue18000    ZoneSettingListResponseZonesBrowserCacheTTLValue = 18000
	ZoneSettingListResponseZonesBrowserCacheTTLValue28800    ZoneSettingListResponseZonesBrowserCacheTTLValue = 28800
	ZoneSettingListResponseZonesBrowserCacheTTLValue43200    ZoneSettingListResponseZonesBrowserCacheTTLValue = 43200
	ZoneSettingListResponseZonesBrowserCacheTTLValue57600    ZoneSettingListResponseZonesBrowserCacheTTLValue = 57600
	ZoneSettingListResponseZonesBrowserCacheTTLValue72000    ZoneSettingListResponseZonesBrowserCacheTTLValue = 72000
	ZoneSettingListResponseZonesBrowserCacheTTLValue86400    ZoneSettingListResponseZonesBrowserCacheTTLValue = 86400
	ZoneSettingListResponseZonesBrowserCacheTTLValue172800   ZoneSettingListResponseZonesBrowserCacheTTLValue = 172800
	ZoneSettingListResponseZonesBrowserCacheTTLValue259200   ZoneSettingListResponseZonesBrowserCacheTTLValue = 259200
	ZoneSettingListResponseZonesBrowserCacheTTLValue345600   ZoneSettingListResponseZonesBrowserCacheTTLValue = 345600
	ZoneSettingListResponseZonesBrowserCacheTTLValue432000   ZoneSettingListResponseZonesBrowserCacheTTLValue = 432000
	ZoneSettingListResponseZonesBrowserCacheTTLValue691200   ZoneSettingListResponseZonesBrowserCacheTTLValue = 691200
	ZoneSettingListResponseZonesBrowserCacheTTLValue1382400  ZoneSettingListResponseZonesBrowserCacheTTLValue = 1382400
	ZoneSettingListResponseZonesBrowserCacheTTLValue2073600  ZoneSettingListResponseZonesBrowserCacheTTLValue = 2073600
	ZoneSettingListResponseZonesBrowserCacheTTLValue2678400  ZoneSettingListResponseZonesBrowserCacheTTLValue = 2678400
	ZoneSettingListResponseZonesBrowserCacheTTLValue5356800  ZoneSettingListResponseZonesBrowserCacheTTLValue = 5356800
	ZoneSettingListResponseZonesBrowserCacheTTLValue16070400 ZoneSettingListResponseZonesBrowserCacheTTLValue = 16070400
	ZoneSettingListResponseZonesBrowserCacheTTLValue31536000 ZoneSettingListResponseZonesBrowserCacheTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesBrowserCacheTTLEditable bool

const (
	ZoneSettingListResponseZonesBrowserCacheTTLEditableTrue  ZoneSettingListResponseZonesBrowserCacheTTLEditable = true
	ZoneSettingListResponseZonesBrowserCacheTTLEditableFalse ZoneSettingListResponseZonesBrowserCacheTTLEditable = false
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingListResponseZonesBrowserCheck struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesBrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesBrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesBrowserCheckJSON `json:"-"`
}

// zoneSettingListResponseZonesBrowserCheckJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesBrowserCheck]
type zoneSettingListResponseZonesBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesBrowserCheck) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesBrowserCheckID string

const (
	ZoneSettingListResponseZonesBrowserCheckIDBrowserCheck ZoneSettingListResponseZonesBrowserCheckID = "browser_check"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesBrowserCheckValue string

const (
	ZoneSettingListResponseZonesBrowserCheckValueOn  ZoneSettingListResponseZonesBrowserCheckValue = "on"
	ZoneSettingListResponseZonesBrowserCheckValueOff ZoneSettingListResponseZonesBrowserCheckValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesBrowserCheckEditable bool

const (
	ZoneSettingListResponseZonesBrowserCheckEditableTrue  ZoneSettingListResponseZonesBrowserCheckEditable = true
	ZoneSettingListResponseZonesBrowserCheckEditableFalse ZoneSettingListResponseZonesBrowserCheckEditable = false
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingListResponseZonesCacheLevel struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesCacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesCacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesCacheLevelJSON `json:"-"`
}

// zoneSettingListResponseZonesCacheLevelJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesCacheLevel]
type zoneSettingListResponseZonesCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesCacheLevel) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesCacheLevelID string

const (
	ZoneSettingListResponseZonesCacheLevelIDCacheLevel ZoneSettingListResponseZonesCacheLevelID = "cache_level"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesCacheLevelValue string

const (
	ZoneSettingListResponseZonesCacheLevelValueAggressive ZoneSettingListResponseZonesCacheLevelValue = "aggressive"
	ZoneSettingListResponseZonesCacheLevelValueBasic      ZoneSettingListResponseZonesCacheLevelValue = "basic"
	ZoneSettingListResponseZonesCacheLevelValueSimplified ZoneSettingListResponseZonesCacheLevelValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesCacheLevelEditable bool

const (
	ZoneSettingListResponseZonesCacheLevelEditableTrue  ZoneSettingListResponseZonesCacheLevelEditable = true
	ZoneSettingListResponseZonesCacheLevelEditableFalse ZoneSettingListResponseZonesCacheLevelEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingListResponseZonesChallengeTTL struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesChallengeTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesChallengeTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesChallengeTTLJSON `json:"-"`
}

// zoneSettingListResponseZonesChallengeTTLJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesChallengeTTL]
type zoneSettingListResponseZonesChallengeTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesChallengeTTL) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesChallengeTTLID string

const (
	ZoneSettingListResponseZonesChallengeTTLIDChallengeTTL ZoneSettingListResponseZonesChallengeTTLID = "challenge_ttl"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesChallengeTTLValue float64

const (
	ZoneSettingListResponseZonesChallengeTTLValue300      ZoneSettingListResponseZonesChallengeTTLValue = 300
	ZoneSettingListResponseZonesChallengeTTLValue900      ZoneSettingListResponseZonesChallengeTTLValue = 900
	ZoneSettingListResponseZonesChallengeTTLValue1800     ZoneSettingListResponseZonesChallengeTTLValue = 1800
	ZoneSettingListResponseZonesChallengeTTLValue2700     ZoneSettingListResponseZonesChallengeTTLValue = 2700
	ZoneSettingListResponseZonesChallengeTTLValue3600     ZoneSettingListResponseZonesChallengeTTLValue = 3600
	ZoneSettingListResponseZonesChallengeTTLValue7200     ZoneSettingListResponseZonesChallengeTTLValue = 7200
	ZoneSettingListResponseZonesChallengeTTLValue10800    ZoneSettingListResponseZonesChallengeTTLValue = 10800
	ZoneSettingListResponseZonesChallengeTTLValue14400    ZoneSettingListResponseZonesChallengeTTLValue = 14400
	ZoneSettingListResponseZonesChallengeTTLValue28800    ZoneSettingListResponseZonesChallengeTTLValue = 28800
	ZoneSettingListResponseZonesChallengeTTLValue57600    ZoneSettingListResponseZonesChallengeTTLValue = 57600
	ZoneSettingListResponseZonesChallengeTTLValue86400    ZoneSettingListResponseZonesChallengeTTLValue = 86400
	ZoneSettingListResponseZonesChallengeTTLValue604800   ZoneSettingListResponseZonesChallengeTTLValue = 604800
	ZoneSettingListResponseZonesChallengeTTLValue2592000  ZoneSettingListResponseZonesChallengeTTLValue = 2592000
	ZoneSettingListResponseZonesChallengeTTLValue31536000 ZoneSettingListResponseZonesChallengeTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesChallengeTTLEditable bool

const (
	ZoneSettingListResponseZonesChallengeTTLEditableTrue  ZoneSettingListResponseZonesChallengeTTLEditable = true
	ZoneSettingListResponseZonesChallengeTTLEditableFalse ZoneSettingListResponseZonesChallengeTTLEditable = false
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingListResponseZonesCiphers struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesCiphersID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesCiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesCiphersJSON `json:"-"`
}

// zoneSettingListResponseZonesCiphersJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesCiphers]
type zoneSettingListResponseZonesCiphersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesCiphers) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesCiphersID string

const (
	ZoneSettingListResponseZonesCiphersIDCiphers ZoneSettingListResponseZonesCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesCiphersEditable bool

const (
	ZoneSettingListResponseZonesCiphersEditableTrue  ZoneSettingListResponseZonesCiphersEditable = true
	ZoneSettingListResponseZonesCiphersEditableFalse ZoneSettingListResponseZonesCiphersEditable = false
)

// Whether or not cname flattening is on.
type ZoneSettingListResponseZonesCnameFlattening struct {
	// How to flatten the cname destination.
	ID ZoneSettingListResponseZonesCnameFlatteningID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesCnameFlatteningValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesCnameFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesCnameFlatteningJSON `json:"-"`
}

// zoneSettingListResponseZonesCnameFlatteningJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesCnameFlattening]
type zoneSettingListResponseZonesCnameFlatteningJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesCnameFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesCnameFlattening) implementsZoneSettingListResponse() {}

// How to flatten the cname destination.
type ZoneSettingListResponseZonesCnameFlatteningID string

const (
	ZoneSettingListResponseZonesCnameFlatteningIDCnameFlattening ZoneSettingListResponseZonesCnameFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesCnameFlatteningValue string

const (
	ZoneSettingListResponseZonesCnameFlatteningValueFlattenAtRoot ZoneSettingListResponseZonesCnameFlatteningValue = "flatten_at_root"
	ZoneSettingListResponseZonesCnameFlatteningValueFlattenAll    ZoneSettingListResponseZonesCnameFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesCnameFlatteningEditable bool

const (
	ZoneSettingListResponseZonesCnameFlatteningEditableTrue  ZoneSettingListResponseZonesCnameFlatteningEditable = true
	ZoneSettingListResponseZonesCnameFlatteningEditableFalse ZoneSettingListResponseZonesCnameFlatteningEditable = false
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingListResponseZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesDevelopmentModeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesDevelopmentModeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                                         `json:"time_remaining"`
	JSON          zoneSettingListResponseZonesDevelopmentModeJSON `json:"-"`
}

// zoneSettingListResponseZonesDevelopmentModeJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesDevelopmentMode]
type zoneSettingListResponseZonesDevelopmentModeJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesDevelopmentMode) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesDevelopmentModeID string

const (
	ZoneSettingListResponseZonesDevelopmentModeIDDevelopmentMode ZoneSettingListResponseZonesDevelopmentModeID = "development_mode"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesDevelopmentModeValue string

const (
	ZoneSettingListResponseZonesDevelopmentModeValueOn  ZoneSettingListResponseZonesDevelopmentModeValue = "on"
	ZoneSettingListResponseZonesDevelopmentModeValueOff ZoneSettingListResponseZonesDevelopmentModeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesDevelopmentModeEditable bool

const (
	ZoneSettingListResponseZonesDevelopmentModeEditableTrue  ZoneSettingListResponseZonesDevelopmentModeEditable = true
	ZoneSettingListResponseZonesDevelopmentModeEditableFalse ZoneSettingListResponseZonesDevelopmentModeEditable = false
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingListResponseZonesEarlyHints struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesEarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesEarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesEarlyHintsJSON `json:"-"`
}

// zoneSettingListResponseZonesEarlyHintsJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesEarlyHints]
type zoneSettingListResponseZonesEarlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesEarlyHints) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesEarlyHintsID string

const (
	ZoneSettingListResponseZonesEarlyHintsIDEarlyHints ZoneSettingListResponseZonesEarlyHintsID = "early_hints"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesEarlyHintsValue string

const (
	ZoneSettingListResponseZonesEarlyHintsValueOn  ZoneSettingListResponseZonesEarlyHintsValue = "on"
	ZoneSettingListResponseZonesEarlyHintsValueOff ZoneSettingListResponseZonesEarlyHintsValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesEarlyHintsEditable bool

const (
	ZoneSettingListResponseZonesEarlyHintsEditableTrue  ZoneSettingListResponseZonesEarlyHintsEditable = true
	ZoneSettingListResponseZonesEarlyHintsEditableFalse ZoneSettingListResponseZonesEarlyHintsEditable = false
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZoneSettingListResponseZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesEdgeCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesEdgeCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesEdgeCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesEdgeCacheTTLJSON `json:"-"`
}

// zoneSettingListResponseZonesEdgeCacheTTLJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesEdgeCacheTTL]
type zoneSettingListResponseZonesEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesEdgeCacheTTL) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesEdgeCacheTTLID string

const (
	ZoneSettingListResponseZonesEdgeCacheTTLIDEdgeCacheTTL ZoneSettingListResponseZonesEdgeCacheTTLID = "edge_cache_ttl"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesEdgeCacheTTLValue float64

const (
	ZoneSettingListResponseZonesEdgeCacheTTLValue30     ZoneSettingListResponseZonesEdgeCacheTTLValue = 30
	ZoneSettingListResponseZonesEdgeCacheTTLValue60     ZoneSettingListResponseZonesEdgeCacheTTLValue = 60
	ZoneSettingListResponseZonesEdgeCacheTTLValue300    ZoneSettingListResponseZonesEdgeCacheTTLValue = 300
	ZoneSettingListResponseZonesEdgeCacheTTLValue1200   ZoneSettingListResponseZonesEdgeCacheTTLValue = 1200
	ZoneSettingListResponseZonesEdgeCacheTTLValue1800   ZoneSettingListResponseZonesEdgeCacheTTLValue = 1800
	ZoneSettingListResponseZonesEdgeCacheTTLValue3600   ZoneSettingListResponseZonesEdgeCacheTTLValue = 3600
	ZoneSettingListResponseZonesEdgeCacheTTLValue7200   ZoneSettingListResponseZonesEdgeCacheTTLValue = 7200
	ZoneSettingListResponseZonesEdgeCacheTTLValue10800  ZoneSettingListResponseZonesEdgeCacheTTLValue = 10800
	ZoneSettingListResponseZonesEdgeCacheTTLValue14400  ZoneSettingListResponseZonesEdgeCacheTTLValue = 14400
	ZoneSettingListResponseZonesEdgeCacheTTLValue18000  ZoneSettingListResponseZonesEdgeCacheTTLValue = 18000
	ZoneSettingListResponseZonesEdgeCacheTTLValue28800  ZoneSettingListResponseZonesEdgeCacheTTLValue = 28800
	ZoneSettingListResponseZonesEdgeCacheTTLValue43200  ZoneSettingListResponseZonesEdgeCacheTTLValue = 43200
	ZoneSettingListResponseZonesEdgeCacheTTLValue57600  ZoneSettingListResponseZonesEdgeCacheTTLValue = 57600
	ZoneSettingListResponseZonesEdgeCacheTTLValue72000  ZoneSettingListResponseZonesEdgeCacheTTLValue = 72000
	ZoneSettingListResponseZonesEdgeCacheTTLValue86400  ZoneSettingListResponseZonesEdgeCacheTTLValue = 86400
	ZoneSettingListResponseZonesEdgeCacheTTLValue172800 ZoneSettingListResponseZonesEdgeCacheTTLValue = 172800
	ZoneSettingListResponseZonesEdgeCacheTTLValue259200 ZoneSettingListResponseZonesEdgeCacheTTLValue = 259200
	ZoneSettingListResponseZonesEdgeCacheTTLValue345600 ZoneSettingListResponseZonesEdgeCacheTTLValue = 345600
	ZoneSettingListResponseZonesEdgeCacheTTLValue432000 ZoneSettingListResponseZonesEdgeCacheTTLValue = 432000
	ZoneSettingListResponseZonesEdgeCacheTTLValue518400 ZoneSettingListResponseZonesEdgeCacheTTLValue = 518400
	ZoneSettingListResponseZonesEdgeCacheTTLValue604800 ZoneSettingListResponseZonesEdgeCacheTTLValue = 604800
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesEdgeCacheTTLEditable bool

const (
	ZoneSettingListResponseZonesEdgeCacheTTLEditableTrue  ZoneSettingListResponseZonesEdgeCacheTTLEditable = true
	ZoneSettingListResponseZonesEdgeCacheTTLEditableFalse ZoneSettingListResponseZonesEdgeCacheTTLEditable = false
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingListResponseZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesEmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesEmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesEmailObfuscationJSON `json:"-"`
}

// zoneSettingListResponseZonesEmailObfuscationJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesEmailObfuscation]
type zoneSettingListResponseZonesEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesEmailObfuscation) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesEmailObfuscationID string

const (
	ZoneSettingListResponseZonesEmailObfuscationIDEmailObfuscation ZoneSettingListResponseZonesEmailObfuscationID = "email_obfuscation"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesEmailObfuscationValue string

const (
	ZoneSettingListResponseZonesEmailObfuscationValueOn  ZoneSettingListResponseZonesEmailObfuscationValue = "on"
	ZoneSettingListResponseZonesEmailObfuscationValueOff ZoneSettingListResponseZonesEmailObfuscationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesEmailObfuscationEditable bool

const (
	ZoneSettingListResponseZonesEmailObfuscationEditableTrue  ZoneSettingListResponseZonesEmailObfuscationEditable = true
	ZoneSettingListResponseZonesEmailObfuscationEditableFalse ZoneSettingListResponseZonesEmailObfuscationEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingListResponseZonesH2Prioritization struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesH2PrioritizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesH2PrioritizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesH2PrioritizationJSON `json:"-"`
}

// zoneSettingListResponseZonesH2PrioritizationJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesH2Prioritization]
type zoneSettingListResponseZonesH2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesH2Prioritization) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesH2PrioritizationID string

const (
	ZoneSettingListResponseZonesH2PrioritizationIDH2Prioritization ZoneSettingListResponseZonesH2PrioritizationID = "h2_prioritization"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesH2PrioritizationValue string

const (
	ZoneSettingListResponseZonesH2PrioritizationValueOn     ZoneSettingListResponseZonesH2PrioritizationValue = "on"
	ZoneSettingListResponseZonesH2PrioritizationValueOff    ZoneSettingListResponseZonesH2PrioritizationValue = "off"
	ZoneSettingListResponseZonesH2PrioritizationValueCustom ZoneSettingListResponseZonesH2PrioritizationValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesH2PrioritizationEditable bool

const (
	ZoneSettingListResponseZonesH2PrioritizationEditableTrue  ZoneSettingListResponseZonesH2PrioritizationEditable = true
	ZoneSettingListResponseZonesH2PrioritizationEditableFalse ZoneSettingListResponseZonesH2PrioritizationEditable = false
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZoneSettingListResponseZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesHotlinkProtectionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesHotlinkProtectionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesHotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesHotlinkProtectionJSON `json:"-"`
}

// zoneSettingListResponseZonesHotlinkProtectionJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesHotlinkProtection]
type zoneSettingListResponseZonesHotlinkProtectionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesHotlinkProtection) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesHotlinkProtectionID string

const (
	ZoneSettingListResponseZonesHotlinkProtectionIDHotlinkProtection ZoneSettingListResponseZonesHotlinkProtectionID = "hotlink_protection"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesHotlinkProtectionValue string

const (
	ZoneSettingListResponseZonesHotlinkProtectionValueOn  ZoneSettingListResponseZonesHotlinkProtectionValue = "on"
	ZoneSettingListResponseZonesHotlinkProtectionValueOff ZoneSettingListResponseZonesHotlinkProtectionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesHotlinkProtectionEditable bool

const (
	ZoneSettingListResponseZonesHotlinkProtectionEditableTrue  ZoneSettingListResponseZonesHotlinkProtectionEditable = true
	ZoneSettingListResponseZonesHotlinkProtectionEditableFalse ZoneSettingListResponseZonesHotlinkProtectionEditable = false
)

// HTTP2 enabled for this zone.
type ZoneSettingListResponseZonesHTTP2 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesHTTP2ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesHTTP2Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesHTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesHTTP2JSON `json:"-"`
}

// zoneSettingListResponseZonesHTTP2JSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesHTTP2]
type zoneSettingListResponseZonesHTTP2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesHTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesHTTP2) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesHTTP2ID string

const (
	ZoneSettingListResponseZonesHTTP2IDHTTP2 ZoneSettingListResponseZonesHTTP2ID = "http2"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesHTTP2Value string

const (
	ZoneSettingListResponseZonesHTTP2ValueOn  ZoneSettingListResponseZonesHTTP2Value = "on"
	ZoneSettingListResponseZonesHTTP2ValueOff ZoneSettingListResponseZonesHTTP2Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesHTTP2Editable bool

const (
	ZoneSettingListResponseZonesHTTP2EditableTrue  ZoneSettingListResponseZonesHTTP2Editable = true
	ZoneSettingListResponseZonesHTTP2EditableFalse ZoneSettingListResponseZonesHTTP2Editable = false
)

// HTTP3 enabled for this zone.
type ZoneSettingListResponseZonesHTTP3 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesHTTP3ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesHTTP3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesHTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesHTTP3JSON `json:"-"`
}

// zoneSettingListResponseZonesHTTP3JSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesHTTP3]
type zoneSettingListResponseZonesHTTP3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesHTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesHTTP3) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesHTTP3ID string

const (
	ZoneSettingListResponseZonesHTTP3IDHTTP3 ZoneSettingListResponseZonesHTTP3ID = "http3"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesHTTP3Value string

const (
	ZoneSettingListResponseZonesHTTP3ValueOn  ZoneSettingListResponseZonesHTTP3Value = "on"
	ZoneSettingListResponseZonesHTTP3ValueOff ZoneSettingListResponseZonesHTTP3Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesHTTP3Editable bool

const (
	ZoneSettingListResponseZonesHTTP3EditableTrue  ZoneSettingListResponseZonesHTTP3Editable = true
	ZoneSettingListResponseZonesHTTP3EditableFalse ZoneSettingListResponseZonesHTTP3Editable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingListResponseZonesImageResizing struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesImageResizingID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesImageResizingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesImageResizingJSON `json:"-"`
}

// zoneSettingListResponseZonesImageResizingJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesImageResizing]
type zoneSettingListResponseZonesImageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesImageResizing) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesImageResizingID string

const (
	ZoneSettingListResponseZonesImageResizingIDImageResizing ZoneSettingListResponseZonesImageResizingID = "image_resizing"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesImageResizingValue string

const (
	ZoneSettingListResponseZonesImageResizingValueOn   ZoneSettingListResponseZonesImageResizingValue = "on"
	ZoneSettingListResponseZonesImageResizingValueOff  ZoneSettingListResponseZonesImageResizingValue = "off"
	ZoneSettingListResponseZonesImageResizingValueOpen ZoneSettingListResponseZonesImageResizingValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesImageResizingEditable bool

const (
	ZoneSettingListResponseZonesImageResizingEditableTrue  ZoneSettingListResponseZonesImageResizingEditable = true
	ZoneSettingListResponseZonesImageResizingEditableFalse ZoneSettingListResponseZonesImageResizingEditable = false
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingListResponseZonesIPGeolocation struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesIPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesIPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesIPGeolocationJSON `json:"-"`
}

// zoneSettingListResponseZonesIPGeolocationJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesIPGeolocation]
type zoneSettingListResponseZonesIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesIPGeolocation) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesIPGeolocationID string

const (
	ZoneSettingListResponseZonesIPGeolocationIDIPGeolocation ZoneSettingListResponseZonesIPGeolocationID = "ip_geolocation"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesIPGeolocationValue string

const (
	ZoneSettingListResponseZonesIPGeolocationValueOn  ZoneSettingListResponseZonesIPGeolocationValue = "on"
	ZoneSettingListResponseZonesIPGeolocationValueOff ZoneSettingListResponseZonesIPGeolocationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesIPGeolocationEditable bool

const (
	ZoneSettingListResponseZonesIPGeolocationEditableTrue  ZoneSettingListResponseZonesIPGeolocationEditable = true
	ZoneSettingListResponseZonesIPGeolocationEditableFalse ZoneSettingListResponseZonesIPGeolocationEditable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingListResponseZonesIPV6 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesIPV6ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesIPV6Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesIPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesIPV6JSON `json:"-"`
}

// zoneSettingListResponseZonesIPV6JSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesIPV6]
type zoneSettingListResponseZonesIPV6JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesIPV6) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesIPV6ID string

const (
	ZoneSettingListResponseZonesIPV6IDIPV6 ZoneSettingListResponseZonesIPV6ID = "ipv6"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesIPV6Value string

const (
	ZoneSettingListResponseZonesIPV6ValueOff ZoneSettingListResponseZonesIPV6Value = "off"
	ZoneSettingListResponseZonesIPV6ValueOn  ZoneSettingListResponseZonesIPV6Value = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesIPV6Editable bool

const (
	ZoneSettingListResponseZonesIPV6EditableTrue  ZoneSettingListResponseZonesIPV6Editable = true
	ZoneSettingListResponseZonesIPV6EditableFalse ZoneSettingListResponseZonesIPV6Editable = false
)

// Maximum size of an allowable upload.
type ZoneSettingListResponseZonesMaxUpload struct {
	// identifier of the zone setting.
	ID ZoneSettingListResponseZonesMaxUploadID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesMaxUploadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesMaxUploadJSON `json:"-"`
}

// zoneSettingListResponseZonesMaxUploadJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesMaxUpload]
type zoneSettingListResponseZonesMaxUploadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesMaxUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesMaxUpload) implementsZoneSettingListResponse() {}

// identifier of the zone setting.
type ZoneSettingListResponseZonesMaxUploadID string

const (
	ZoneSettingListResponseZonesMaxUploadIDMaxUpload ZoneSettingListResponseZonesMaxUploadID = "max_upload"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesMaxUploadValue float64

const (
	ZoneSettingListResponseZonesMaxUploadValue100 ZoneSettingListResponseZonesMaxUploadValue = 100
	ZoneSettingListResponseZonesMaxUploadValue200 ZoneSettingListResponseZonesMaxUploadValue = 200
	ZoneSettingListResponseZonesMaxUploadValue500 ZoneSettingListResponseZonesMaxUploadValue = 500
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesMaxUploadEditable bool

const (
	ZoneSettingListResponseZonesMaxUploadEditableTrue  ZoneSettingListResponseZonesMaxUploadEditable = true
	ZoneSettingListResponseZonesMaxUploadEditableFalse ZoneSettingListResponseZonesMaxUploadEditable = false
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingListResponseZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesMinTLSVersionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesMinTLSVersionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesMinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesMinTLSVersionJSON `json:"-"`
}

// zoneSettingListResponseZonesMinTLSVersionJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesMinTLSVersion]
type zoneSettingListResponseZonesMinTLSVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesMinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesMinTLSVersion) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesMinTLSVersionID string

const (
	ZoneSettingListResponseZonesMinTLSVersionIDMinTLSVersion ZoneSettingListResponseZonesMinTLSVersionID = "min_tls_version"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesMinTLSVersionValue string

const (
	ZoneSettingListResponseZonesMinTLSVersionValue1_0 ZoneSettingListResponseZonesMinTLSVersionValue = "1.0"
	ZoneSettingListResponseZonesMinTLSVersionValue1_1 ZoneSettingListResponseZonesMinTLSVersionValue = "1.1"
	ZoneSettingListResponseZonesMinTLSVersionValue1_2 ZoneSettingListResponseZonesMinTLSVersionValue = "1.2"
	ZoneSettingListResponseZonesMinTLSVersionValue1_3 ZoneSettingListResponseZonesMinTLSVersionValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesMinTLSVersionEditable bool

const (
	ZoneSettingListResponseZonesMinTLSVersionEditableTrue  ZoneSettingListResponseZonesMinTLSVersionEditable = true
	ZoneSettingListResponseZonesMinTLSVersionEditableFalse ZoneSettingListResponseZonesMinTLSVersionEditable = false
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingListResponseZonesMinify struct {
	// Zone setting identifier.
	ID ZoneSettingListResponseZonesMinifyID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesMinifyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesMinifyJSON `json:"-"`
}

// zoneSettingListResponseZonesMinifyJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesMinify]
type zoneSettingListResponseZonesMinifyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesMinify) implementsZoneSettingListResponse() {}

// Zone setting identifier.
type ZoneSettingListResponseZonesMinifyID string

const (
	ZoneSettingListResponseZonesMinifyIDMinify ZoneSettingListResponseZonesMinifyID = "minify"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css ZoneSettingListResponseZonesMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML ZoneSettingListResponseZonesMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   ZoneSettingListResponseZonesMinifyValueJs   `json:"js"`
	JSON zoneSettingListResponseZonesMinifyValueJSON `json:"-"`
}

// zoneSettingListResponseZonesMinifyValueJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesMinifyValue]
type zoneSettingListResponseZonesMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingListResponseZonesMinifyValueCss string

const (
	ZoneSettingListResponseZonesMinifyValueCssOn  ZoneSettingListResponseZonesMinifyValueCss = "on"
	ZoneSettingListResponseZonesMinifyValueCssOff ZoneSettingListResponseZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingListResponseZonesMinifyValueHTML string

const (
	ZoneSettingListResponseZonesMinifyValueHTMLOn  ZoneSettingListResponseZonesMinifyValueHTML = "on"
	ZoneSettingListResponseZonesMinifyValueHTMLOff ZoneSettingListResponseZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingListResponseZonesMinifyValueJs string

const (
	ZoneSettingListResponseZonesMinifyValueJsOn  ZoneSettingListResponseZonesMinifyValueJs = "on"
	ZoneSettingListResponseZonesMinifyValueJsOff ZoneSettingListResponseZonesMinifyValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesMinifyEditable bool

const (
	ZoneSettingListResponseZonesMinifyEditableTrue  ZoneSettingListResponseZonesMinifyEditable = true
	ZoneSettingListResponseZonesMinifyEditableFalse ZoneSettingListResponseZonesMinifyEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingListResponseZonesMirage struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesMirageID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesMirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesMirageJSON `json:"-"`
}

// zoneSettingListResponseZonesMirageJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesMirage]
type zoneSettingListResponseZonesMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesMirage) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesMirageID string

const (
	ZoneSettingListResponseZonesMirageIDMirage ZoneSettingListResponseZonesMirageID = "mirage"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesMirageValue string

const (
	ZoneSettingListResponseZonesMirageValueOn  ZoneSettingListResponseZonesMirageValue = "on"
	ZoneSettingListResponseZonesMirageValueOff ZoneSettingListResponseZonesMirageValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesMirageEditable bool

const (
	ZoneSettingListResponseZonesMirageEditableTrue  ZoneSettingListResponseZonesMirageEditable = true
	ZoneSettingListResponseZonesMirageEditableFalse ZoneSettingListResponseZonesMirageEditable = false
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingListResponseZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID ZoneSettingListResponseZonesMobileRedirectID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesMobileRedirectValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesMobileRedirectJSON `json:"-"`
}

// zoneSettingListResponseZonesMobileRedirectJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesMobileRedirect]
type zoneSettingListResponseZonesMobileRedirectJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesMobileRedirect) implementsZoneSettingListResponse() {}

// Identifier of the zone setting.
type ZoneSettingListResponseZonesMobileRedirectID string

const (
	ZoneSettingListResponseZonesMobileRedirectIDMobileRedirect ZoneSettingListResponseZonesMobileRedirectID = "mobile_redirect"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingListResponseZonesMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                                `json:"strip_uri"`
	JSON     zoneSettingListResponseZonesMobileRedirectValueJSON `json:"-"`
}

// zoneSettingListResponseZonesMobileRedirectValueJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseZonesMobileRedirectValue]
type zoneSettingListResponseZonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingListResponseZonesMobileRedirectValueStatus string

const (
	ZoneSettingListResponseZonesMobileRedirectValueStatusOn  ZoneSettingListResponseZonesMobileRedirectValueStatus = "on"
	ZoneSettingListResponseZonesMobileRedirectValueStatusOff ZoneSettingListResponseZonesMobileRedirectValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesMobileRedirectEditable bool

const (
	ZoneSettingListResponseZonesMobileRedirectEditableTrue  ZoneSettingListResponseZonesMobileRedirectEditable = true
	ZoneSettingListResponseZonesMobileRedirectEditableFalse ZoneSettingListResponseZonesMobileRedirectEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingListResponseZonesNEL struct {
	// Zone setting identifier.
	ID ZoneSettingListResponseZonesNELID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesNELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesNELJSON `json:"-"`
}

// zoneSettingListResponseZonesNELJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesNEL]
type zoneSettingListResponseZonesNELJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesNEL) implementsZoneSettingListResponse() {}

// Zone setting identifier.
type ZoneSettingListResponseZonesNELID string

const (
	ZoneSettingListResponseZonesNELIDNEL ZoneSettingListResponseZonesNELID = "nel"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesNELValue struct {
	Enabled bool                                     `json:"enabled"`
	JSON    zoneSettingListResponseZonesNELValueJSON `json:"-"`
}

// zoneSettingListResponseZonesNELValueJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesNELValue]
type zoneSettingListResponseZonesNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesNELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesNELEditable bool

const (
	ZoneSettingListResponseZonesNELEditableTrue  ZoneSettingListResponseZonesNELEditable = true
	ZoneSettingListResponseZonesNELEditableFalse ZoneSettingListResponseZonesNELEditable = false
)

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingListResponseZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesOpportunisticEncryptionJSON `json:"-"`
}

// zoneSettingListResponseZonesOpportunisticEncryptionJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseZonesOpportunisticEncryption]
type zoneSettingListResponseZonesOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesOpportunisticEncryption) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesOpportunisticEncryptionID string

const (
	ZoneSettingListResponseZonesOpportunisticEncryptionIDOpportunisticEncryption ZoneSettingListResponseZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesOpportunisticEncryptionValue string

const (
	ZoneSettingListResponseZonesOpportunisticEncryptionValueOn  ZoneSettingListResponseZonesOpportunisticEncryptionValue = "on"
	ZoneSettingListResponseZonesOpportunisticEncryptionValueOff ZoneSettingListResponseZonesOpportunisticEncryptionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesOpportunisticEncryptionEditable bool

const (
	ZoneSettingListResponseZonesOpportunisticEncryptionEditableTrue  ZoneSettingListResponseZonesOpportunisticEncryptionEditable = true
	ZoneSettingListResponseZonesOpportunisticEncryptionEditableFalse ZoneSettingListResponseZonesOpportunisticEncryptionEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingListResponseZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesOpportunisticOnionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesOpportunisticOnionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesOpportunisticOnionJSON `json:"-"`
}

// zoneSettingListResponseZonesOpportunisticOnionJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseZonesOpportunisticOnion]
type zoneSettingListResponseZonesOpportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesOpportunisticOnion) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesOpportunisticOnionID string

const (
	ZoneSettingListResponseZonesOpportunisticOnionIDOpportunisticOnion ZoneSettingListResponseZonesOpportunisticOnionID = "opportunistic_onion"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesOpportunisticOnionValue string

const (
	ZoneSettingListResponseZonesOpportunisticOnionValueOn  ZoneSettingListResponseZonesOpportunisticOnionValue = "on"
	ZoneSettingListResponseZonesOpportunisticOnionValueOff ZoneSettingListResponseZonesOpportunisticOnionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesOpportunisticOnionEditable bool

const (
	ZoneSettingListResponseZonesOpportunisticOnionEditableTrue  ZoneSettingListResponseZonesOpportunisticOnionEditable = true
	ZoneSettingListResponseZonesOpportunisticOnionEditableFalse ZoneSettingListResponseZonesOpportunisticOnionEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingListResponseZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesOrangeToOrangeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesOrangeToOrangeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesOrangeToOrangeJSON `json:"-"`
}

// zoneSettingListResponseZonesOrangeToOrangeJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesOrangeToOrange]
type zoneSettingListResponseZonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesOrangeToOrange) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesOrangeToOrangeID string

const (
	ZoneSettingListResponseZonesOrangeToOrangeIDOrangeToOrange ZoneSettingListResponseZonesOrangeToOrangeID = "orange_to_orange"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesOrangeToOrangeValue string

const (
	ZoneSettingListResponseZonesOrangeToOrangeValueOn  ZoneSettingListResponseZonesOrangeToOrangeValue = "on"
	ZoneSettingListResponseZonesOrangeToOrangeValueOff ZoneSettingListResponseZonesOrangeToOrangeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesOrangeToOrangeEditable bool

const (
	ZoneSettingListResponseZonesOrangeToOrangeEditableTrue  ZoneSettingListResponseZonesOrangeToOrangeEditable = true
	ZoneSettingListResponseZonesOrangeToOrangeEditableFalse ZoneSettingListResponseZonesOrangeToOrangeEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZoneSettingListResponseZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesOriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesOriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesOriginErrorPagePassThruJSON `json:"-"`
}

// zoneSettingListResponseZonesOriginErrorPagePassThruJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseZonesOriginErrorPagePassThru]
type zoneSettingListResponseZonesOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesOriginErrorPagePassThru) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesOriginErrorPagePassThruID string

const (
	ZoneSettingListResponseZonesOriginErrorPagePassThruIDOriginErrorPagePassThru ZoneSettingListResponseZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesOriginErrorPagePassThruValue string

const (
	ZoneSettingListResponseZonesOriginErrorPagePassThruValueOn  ZoneSettingListResponseZonesOriginErrorPagePassThruValue = "on"
	ZoneSettingListResponseZonesOriginErrorPagePassThruValueOff ZoneSettingListResponseZonesOriginErrorPagePassThruValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesOriginErrorPagePassThruEditable bool

const (
	ZoneSettingListResponseZonesOriginErrorPagePassThruEditableTrue  ZoneSettingListResponseZonesOriginErrorPagePassThruEditable = true
	ZoneSettingListResponseZonesOriginErrorPagePassThruEditableFalse ZoneSettingListResponseZonesOriginErrorPagePassThruEditable = false
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingListResponseZonesPolish struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesPolishID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesPolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesPolishJSON `json:"-"`
}

// zoneSettingListResponseZonesPolishJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesPolish]
type zoneSettingListResponseZonesPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesPolish) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesPolishID string

const (
	ZoneSettingListResponseZonesPolishIDPolish ZoneSettingListResponseZonesPolishID = "polish"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesPolishValue string

const (
	ZoneSettingListResponseZonesPolishValueOff      ZoneSettingListResponseZonesPolishValue = "off"
	ZoneSettingListResponseZonesPolishValueLossless ZoneSettingListResponseZonesPolishValue = "lossless"
	ZoneSettingListResponseZonesPolishValueLossy    ZoneSettingListResponseZonesPolishValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesPolishEditable bool

const (
	ZoneSettingListResponseZonesPolishEditableTrue  ZoneSettingListResponseZonesPolishEditable = true
	ZoneSettingListResponseZonesPolishEditableFalse ZoneSettingListResponseZonesPolishEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingListResponseZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesPrefetchPreloadID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesPrefetchPreloadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesPrefetchPreloadJSON `json:"-"`
}

// zoneSettingListResponseZonesPrefetchPreloadJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesPrefetchPreload]
type zoneSettingListResponseZonesPrefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesPrefetchPreload) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesPrefetchPreloadID string

const (
	ZoneSettingListResponseZonesPrefetchPreloadIDPrefetchPreload ZoneSettingListResponseZonesPrefetchPreloadID = "prefetch_preload"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesPrefetchPreloadValue string

const (
	ZoneSettingListResponseZonesPrefetchPreloadValueOn  ZoneSettingListResponseZonesPrefetchPreloadValue = "on"
	ZoneSettingListResponseZonesPrefetchPreloadValueOff ZoneSettingListResponseZonesPrefetchPreloadValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesPrefetchPreloadEditable bool

const (
	ZoneSettingListResponseZonesPrefetchPreloadEditableTrue  ZoneSettingListResponseZonesPrefetchPreloadEditable = true
	ZoneSettingListResponseZonesPrefetchPreloadEditableFalse ZoneSettingListResponseZonesPrefetchPreloadEditable = false
)

// Maximum time between two read operations from origin.
type ZoneSettingListResponseZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesProxyReadTimeoutJSON `json:"-"`
}

// zoneSettingListResponseZonesProxyReadTimeoutJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesProxyReadTimeout]
type zoneSettingListResponseZonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesProxyReadTimeout) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesProxyReadTimeoutID string

const (
	ZoneSettingListResponseZonesProxyReadTimeoutIDProxyReadTimeout ZoneSettingListResponseZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesProxyReadTimeoutEditable bool

const (
	ZoneSettingListResponseZonesProxyReadTimeoutEditableTrue  ZoneSettingListResponseZonesProxyReadTimeoutEditable = true
	ZoneSettingListResponseZonesProxyReadTimeoutEditableFalse ZoneSettingListResponseZonesProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type ZoneSettingListResponseZonesPseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID ZoneSettingListResponseZonesPseudoIPV4ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesPseudoIPV4Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesPseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesPseudoIPV4JSON `json:"-"`
}

// zoneSettingListResponseZonesPseudoIPV4JSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesPseudoIPV4]
type zoneSettingListResponseZonesPseudoIPV4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesPseudoIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesPseudoIPV4) implementsZoneSettingListResponse() {}

// Value of the Pseudo IPv4 setting.
type ZoneSettingListResponseZonesPseudoIPV4ID string

const (
	ZoneSettingListResponseZonesPseudoIPV4IDPseudoIPV4 ZoneSettingListResponseZonesPseudoIPV4ID = "pseudo_ipv4"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesPseudoIPV4Value string

const (
	ZoneSettingListResponseZonesPseudoIPV4ValueOff             ZoneSettingListResponseZonesPseudoIPV4Value = "off"
	ZoneSettingListResponseZonesPseudoIPV4ValueAddHeader       ZoneSettingListResponseZonesPseudoIPV4Value = "add_header"
	ZoneSettingListResponseZonesPseudoIPV4ValueOverwriteHeader ZoneSettingListResponseZonesPseudoIPV4Value = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesPseudoIPV4Editable bool

const (
	ZoneSettingListResponseZonesPseudoIPV4EditableTrue  ZoneSettingListResponseZonesPseudoIPV4Editable = true
	ZoneSettingListResponseZonesPseudoIPV4EditableFalse ZoneSettingListResponseZonesPseudoIPV4Editable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingListResponseZonesResponseBuffering struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesResponseBufferingID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesResponseBufferingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesResponseBufferingJSON `json:"-"`
}

// zoneSettingListResponseZonesResponseBufferingJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesResponseBuffering]
type zoneSettingListResponseZonesResponseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesResponseBuffering) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesResponseBufferingID string

const (
	ZoneSettingListResponseZonesResponseBufferingIDResponseBuffering ZoneSettingListResponseZonesResponseBufferingID = "response_buffering"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesResponseBufferingValue string

const (
	ZoneSettingListResponseZonesResponseBufferingValueOn  ZoneSettingListResponseZonesResponseBufferingValue = "on"
	ZoneSettingListResponseZonesResponseBufferingValueOff ZoneSettingListResponseZonesResponseBufferingValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesResponseBufferingEditable bool

const (
	ZoneSettingListResponseZonesResponseBufferingEditableTrue  ZoneSettingListResponseZonesResponseBufferingEditable = true
	ZoneSettingListResponseZonesResponseBufferingEditableFalse ZoneSettingListResponseZonesResponseBufferingEditable = false
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
type ZoneSettingListResponseZonesRocketLoader struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesRocketLoaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesRocketLoaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesRocketLoaderJSON `json:"-"`
}

// zoneSettingListResponseZonesRocketLoaderJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesRocketLoader]
type zoneSettingListResponseZonesRocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesRocketLoader) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesRocketLoaderID string

const (
	ZoneSettingListResponseZonesRocketLoaderIDRocketLoader ZoneSettingListResponseZonesRocketLoaderID = "rocket_loader"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesRocketLoaderValue string

const (
	ZoneSettingListResponseZonesRocketLoaderValueOn  ZoneSettingListResponseZonesRocketLoaderValue = "on"
	ZoneSettingListResponseZonesRocketLoaderValueOff ZoneSettingListResponseZonesRocketLoaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesRocketLoaderEditable bool

const (
	ZoneSettingListResponseZonesRocketLoaderEditableTrue  ZoneSettingListResponseZonesRocketLoaderEditable = true
	ZoneSettingListResponseZonesRocketLoaderEditableFalse ZoneSettingListResponseZonesRocketLoaderEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationJSON `json:"-"`
}

// zoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationJSON contains
// the JSON metadata for the struct
// [ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimization]
type zoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimization) implementsZoneSettingListResponse() {
}

// ID of the zone setting.
type ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationID string

const (
	ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationValue struct {
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
	WpPlugin bool                                                                      `json:"wp_plugin,required"`
	JSON     zoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationValueJSON `json:"-"`
}

// zoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationValueJSON
// contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationValue]
type zoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationValueJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingListResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Cloudflare security header for a zone.
type ZoneSettingListResponseZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID ZoneSettingListResponseZonesSecurityHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesSecurityHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesSecurityHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesSecurityHeaderJSON `json:"-"`
}

// zoneSettingListResponseZonesSecurityHeaderJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesSecurityHeader]
type zoneSettingListResponseZonesSecurityHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesSecurityHeader) implementsZoneSettingListResponse() {}

// ID of the zone's security header.
type ZoneSettingListResponseZonesSecurityHeaderID string

const (
	ZoneSettingListResponseZonesSecurityHeaderIDSecurityHeader ZoneSettingListResponseZonesSecurityHeaderID = "security_header"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZoneSettingListResponseZonesSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zoneSettingListResponseZonesSecurityHeaderValueJSON                    `json:"-"`
}

// zoneSettingListResponseZonesSecurityHeaderValueJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseZonesSecurityHeaderValue]
type zoneSettingListResponseZonesSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type ZoneSettingListResponseZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                       `json:"nosniff"`
	JSON    zoneSettingListResponseZonesSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// zoneSettingListResponseZonesSecurityHeaderValueStrictTransportSecurityJSON
// contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesSecurityHeaderValueStrictTransportSecurity]
type zoneSettingListResponseZonesSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesSecurityHeaderEditable bool

const (
	ZoneSettingListResponseZonesSecurityHeaderEditableTrue  ZoneSettingListResponseZonesSecurityHeaderEditable = true
	ZoneSettingListResponseZonesSecurityHeaderEditableFalse ZoneSettingListResponseZonesSecurityHeaderEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingListResponseZonesSecurityLevel struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesSecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesSecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesSecurityLevelJSON `json:"-"`
}

// zoneSettingListResponseZonesSecurityLevelJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesSecurityLevel]
type zoneSettingListResponseZonesSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesSecurityLevel) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesSecurityLevelID string

const (
	ZoneSettingListResponseZonesSecurityLevelIDSecurityLevel ZoneSettingListResponseZonesSecurityLevelID = "security_level"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesSecurityLevelValue string

const (
	ZoneSettingListResponseZonesSecurityLevelValueOff            ZoneSettingListResponseZonesSecurityLevelValue = "off"
	ZoneSettingListResponseZonesSecurityLevelValueEssentiallyOff ZoneSettingListResponseZonesSecurityLevelValue = "essentially_off"
	ZoneSettingListResponseZonesSecurityLevelValueLow            ZoneSettingListResponseZonesSecurityLevelValue = "low"
	ZoneSettingListResponseZonesSecurityLevelValueMedium         ZoneSettingListResponseZonesSecurityLevelValue = "medium"
	ZoneSettingListResponseZonesSecurityLevelValueHigh           ZoneSettingListResponseZonesSecurityLevelValue = "high"
	ZoneSettingListResponseZonesSecurityLevelValueUnderAttack    ZoneSettingListResponseZonesSecurityLevelValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesSecurityLevelEditable bool

const (
	ZoneSettingListResponseZonesSecurityLevelEditableTrue  ZoneSettingListResponseZonesSecurityLevelEditable = true
	ZoneSettingListResponseZonesSecurityLevelEditableFalse ZoneSettingListResponseZonesSecurityLevelEditable = false
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
type ZoneSettingListResponseZonesServerSideExclude struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesServerSideExcludeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesServerSideExcludeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesServerSideExcludeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesServerSideExcludeJSON `json:"-"`
}

// zoneSettingListResponseZonesServerSideExcludeJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesServerSideExclude]
type zoneSettingListResponseZonesServerSideExcludeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesServerSideExclude) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesServerSideExcludeID string

const (
	ZoneSettingListResponseZonesServerSideExcludeIDServerSideExclude ZoneSettingListResponseZonesServerSideExcludeID = "server_side_exclude"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesServerSideExcludeValue string

const (
	ZoneSettingListResponseZonesServerSideExcludeValueOn  ZoneSettingListResponseZonesServerSideExcludeValue = "on"
	ZoneSettingListResponseZonesServerSideExcludeValueOff ZoneSettingListResponseZonesServerSideExcludeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesServerSideExcludeEditable bool

const (
	ZoneSettingListResponseZonesServerSideExcludeEditableTrue  ZoneSettingListResponseZonesServerSideExcludeEditable = true
	ZoneSettingListResponseZonesServerSideExcludeEditableFalse ZoneSettingListResponseZonesServerSideExcludeEditable = false
)

// Allow SHA1 support.
type ZoneSettingListResponseZonesSha1Support struct {
	// Zone setting identifier.
	ID ZoneSettingListResponseZonesSha1SupportID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesSha1SupportValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesSha1SupportJSON `json:"-"`
}

// zoneSettingListResponseZonesSha1SupportJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesSha1Support]
type zoneSettingListResponseZonesSha1SupportJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesSha1Support) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesSha1Support) implementsZoneSettingListResponse() {}

// Zone setting identifier.
type ZoneSettingListResponseZonesSha1SupportID string

const (
	ZoneSettingListResponseZonesSha1SupportIDSha1Support ZoneSettingListResponseZonesSha1SupportID = "sha1_support"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesSha1SupportValue string

const (
	ZoneSettingListResponseZonesSha1SupportValueOff ZoneSettingListResponseZonesSha1SupportValue = "off"
	ZoneSettingListResponseZonesSha1SupportValueOn  ZoneSettingListResponseZonesSha1SupportValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesSha1SupportEditable bool

const (
	ZoneSettingListResponseZonesSha1SupportEditableTrue  ZoneSettingListResponseZonesSha1SupportEditable = true
	ZoneSettingListResponseZonesSha1SupportEditableFalse ZoneSettingListResponseZonesSha1SupportEditable = false
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingListResponseZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesSortQueryStringForCacheID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesSortQueryStringForCacheValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesSortQueryStringForCacheJSON `json:"-"`
}

// zoneSettingListResponseZonesSortQueryStringForCacheJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseZonesSortQueryStringForCache]
type zoneSettingListResponseZonesSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesSortQueryStringForCache) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesSortQueryStringForCacheID string

const (
	ZoneSettingListResponseZonesSortQueryStringForCacheIDSortQueryStringForCache ZoneSettingListResponseZonesSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesSortQueryStringForCacheValue string

const (
	ZoneSettingListResponseZonesSortQueryStringForCacheValueOn  ZoneSettingListResponseZonesSortQueryStringForCacheValue = "on"
	ZoneSettingListResponseZonesSortQueryStringForCacheValueOff ZoneSettingListResponseZonesSortQueryStringForCacheValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesSortQueryStringForCacheEditable bool

const (
	ZoneSettingListResponseZonesSortQueryStringForCacheEditableTrue  ZoneSettingListResponseZonesSortQueryStringForCacheEditable = true
	ZoneSettingListResponseZonesSortQueryStringForCacheEditableFalse ZoneSettingListResponseZonesSortQueryStringForCacheEditable = false
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
type ZoneSettingListResponseZonesSSL struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesSSLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesSSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesSSLJSON `json:"-"`
}

// zoneSettingListResponseZonesSSLJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesSSL]
type zoneSettingListResponseZonesSSLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesSSL) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesSSLID string

const (
	ZoneSettingListResponseZonesSSLIDSSL ZoneSettingListResponseZonesSSLID = "ssl"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesSSLValue string

const (
	ZoneSettingListResponseZonesSSLValueOff      ZoneSettingListResponseZonesSSLValue = "off"
	ZoneSettingListResponseZonesSSLValueFlexible ZoneSettingListResponseZonesSSLValue = "flexible"
	ZoneSettingListResponseZonesSSLValueFull     ZoneSettingListResponseZonesSSLValue = "full"
	ZoneSettingListResponseZonesSSLValueStrict   ZoneSettingListResponseZonesSSLValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesSSLEditable bool

const (
	ZoneSettingListResponseZonesSSLEditableTrue  ZoneSettingListResponseZonesSSLEditable = true
	ZoneSettingListResponseZonesSSLEditableFalse ZoneSettingListResponseZonesSSLEditable = false
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingListResponseZonesSSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingListResponseZonesSSLRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                           `json:"enabled"`
	JSON    zoneSettingListResponseZonesSSLRecommenderJSON `json:"-"`
}

// zoneSettingListResponseZonesSSLRecommenderJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseZonesSSLRecommender]
type zoneSettingListResponseZonesSSLRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesSSLRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesSSLRecommender) implementsZoneSettingListResponse() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingListResponseZonesSSLRecommenderID string

const (
	ZoneSettingListResponseZonesSSLRecommenderIDSSLRecommender ZoneSettingListResponseZonesSSLRecommenderID = "ssl_recommender"
)

// Only allows TLS1.2.
type ZoneSettingListResponseZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID ZoneSettingListResponseZonesTLS1_2OnlyID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesTLS1_2OnlyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesTLS1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesTls1_2OnlyJSON `json:"-"`
}

// zoneSettingListResponseZonesTls1_2OnlyJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesTLS1_2Only]
type zoneSettingListResponseZonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesTLS1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesTLS1_2Only) implementsZoneSettingListResponse() {}

// Zone setting identifier.
type ZoneSettingListResponseZonesTLS1_2OnlyID string

const (
	ZoneSettingListResponseZonesTLS1_2OnlyIDTLS1_2Only ZoneSettingListResponseZonesTLS1_2OnlyID = "tls_1_2_only"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesTLS1_2OnlyValue string

const (
	ZoneSettingListResponseZonesTLS1_2OnlyValueOff ZoneSettingListResponseZonesTLS1_2OnlyValue = "off"
	ZoneSettingListResponseZonesTLS1_2OnlyValueOn  ZoneSettingListResponseZonesTLS1_2OnlyValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesTLS1_2OnlyEditable bool

const (
	ZoneSettingListResponseZonesTLS1_2OnlyEditableTrue  ZoneSettingListResponseZonesTLS1_2OnlyEditable = true
	ZoneSettingListResponseZonesTLS1_2OnlyEditableFalse ZoneSettingListResponseZonesTLS1_2OnlyEditable = false
)

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingListResponseZonesTLS1_3 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesTLS1_3ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesTLS1_3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesTLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesTls1_3JSON `json:"-"`
}

// zoneSettingListResponseZonesTls1_3JSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesTLS1_3]
type zoneSettingListResponseZonesTls1_3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesTLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesTLS1_3) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesTLS1_3ID string

const (
	ZoneSettingListResponseZonesTLS1_3IDTLS1_3 ZoneSettingListResponseZonesTLS1_3ID = "tls_1_3"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesTLS1_3Value string

const (
	ZoneSettingListResponseZonesTLS1_3ValueOn  ZoneSettingListResponseZonesTLS1_3Value = "on"
	ZoneSettingListResponseZonesTLS1_3ValueOff ZoneSettingListResponseZonesTLS1_3Value = "off"
	ZoneSettingListResponseZonesTLS1_3ValueZrt ZoneSettingListResponseZonesTLS1_3Value = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesTLS1_3Editable bool

const (
	ZoneSettingListResponseZonesTLS1_3EditableTrue  ZoneSettingListResponseZonesTLS1_3Editable = true
	ZoneSettingListResponseZonesTLS1_3EditableFalse ZoneSettingListResponseZonesTLS1_3Editable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingListResponseZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesTLSClientAuthID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesTLSClientAuthValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesTLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesTLSClientAuthJSON `json:"-"`
}

// zoneSettingListResponseZonesTLSClientAuthJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesTLSClientAuth]
type zoneSettingListResponseZonesTLSClientAuthJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesTLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesTLSClientAuth) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesTLSClientAuthID string

const (
	ZoneSettingListResponseZonesTLSClientAuthIDTLSClientAuth ZoneSettingListResponseZonesTLSClientAuthID = "tls_client_auth"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesTLSClientAuthValue string

const (
	ZoneSettingListResponseZonesTLSClientAuthValueOn  ZoneSettingListResponseZonesTLSClientAuthValue = "on"
	ZoneSettingListResponseZonesTLSClientAuthValueOff ZoneSettingListResponseZonesTLSClientAuthValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesTLSClientAuthEditable bool

const (
	ZoneSettingListResponseZonesTLSClientAuthEditableTrue  ZoneSettingListResponseZonesTLSClientAuthEditable = true
	ZoneSettingListResponseZonesTLSClientAuthEditableFalse ZoneSettingListResponseZonesTLSClientAuthEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingListResponseZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesTrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesTrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesTrueClientIPHeaderJSON `json:"-"`
}

// zoneSettingListResponseZonesTrueClientIPHeaderJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseZonesTrueClientIPHeader]
type zoneSettingListResponseZonesTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesTrueClientIPHeader) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesTrueClientIPHeaderID string

const (
	ZoneSettingListResponseZonesTrueClientIPHeaderIDTrueClientIPHeader ZoneSettingListResponseZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesTrueClientIPHeaderValue string

const (
	ZoneSettingListResponseZonesTrueClientIPHeaderValueOn  ZoneSettingListResponseZonesTrueClientIPHeaderValue = "on"
	ZoneSettingListResponseZonesTrueClientIPHeaderValueOff ZoneSettingListResponseZonesTrueClientIPHeaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesTrueClientIPHeaderEditable bool

const (
	ZoneSettingListResponseZonesTrueClientIPHeaderEditableTrue  ZoneSettingListResponseZonesTrueClientIPHeaderEditable = true
	ZoneSettingListResponseZonesTrueClientIPHeaderEditableFalse ZoneSettingListResponseZonesTrueClientIPHeaderEditable = false
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
type ZoneSettingListResponseZonesWAF struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesWAFID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesWAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesWAFJSON `json:"-"`
}

// zoneSettingListResponseZonesWAFJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesWAF]
type zoneSettingListResponseZonesWAFJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesWAF) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesWAFID string

const (
	ZoneSettingListResponseZonesWAFIDWAF ZoneSettingListResponseZonesWAFID = "waf"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesWAFValue string

const (
	ZoneSettingListResponseZonesWAFValueOn  ZoneSettingListResponseZonesWAFValue = "on"
	ZoneSettingListResponseZonesWAFValueOff ZoneSettingListResponseZonesWAFValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesWAFEditable bool

const (
	ZoneSettingListResponseZonesWAFEditableTrue  ZoneSettingListResponseZonesWAFEditable = true
	ZoneSettingListResponseZonesWAFEditableFalse ZoneSettingListResponseZonesWAFEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingListResponseZonesWebp struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesWebpID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesWebpValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesWebpJSON `json:"-"`
}

// zoneSettingListResponseZonesWebpJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseZonesWebp]
type zoneSettingListResponseZonesWebpJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesWebp) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesWebpID string

const (
	ZoneSettingListResponseZonesWebpIDWebp ZoneSettingListResponseZonesWebpID = "webp"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesWebpValue string

const (
	ZoneSettingListResponseZonesWebpValueOff ZoneSettingListResponseZonesWebpValue = "off"
	ZoneSettingListResponseZonesWebpValueOn  ZoneSettingListResponseZonesWebpValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesWebpEditable bool

const (
	ZoneSettingListResponseZonesWebpEditableTrue  ZoneSettingListResponseZonesWebpEditable = true
	ZoneSettingListResponseZonesWebpEditableFalse ZoneSettingListResponseZonesWebpEditable = false
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZoneSettingListResponseZonesWebsockets struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseZonesWebsocketsID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingListResponseZonesWebsocketsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseZonesWebsocketsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseZonesWebsocketsJSON `json:"-"`
}

// zoneSettingListResponseZonesWebsocketsJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseZonesWebsockets]
type zoneSettingListResponseZonesWebsocketsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseZonesWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseZonesWebsockets) implementsZoneSettingListResponse() {}

// ID of the zone setting.
type ZoneSettingListResponseZonesWebsocketsID string

const (
	ZoneSettingListResponseZonesWebsocketsIDWebsockets ZoneSettingListResponseZonesWebsocketsID = "websockets"
)

// Current value of the zone setting.
type ZoneSettingListResponseZonesWebsocketsValue string

const (
	ZoneSettingListResponseZonesWebsocketsValueOff ZoneSettingListResponseZonesWebsocketsValue = "off"
	ZoneSettingListResponseZonesWebsocketsValueOn  ZoneSettingListResponseZonesWebsocketsValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseZonesWebsocketsEditable bool

const (
	ZoneSettingListResponseZonesWebsocketsEditableTrue  ZoneSettingListResponseZonesWebsocketsEditable = true
	ZoneSettingListResponseZonesWebsocketsEditableFalse ZoneSettingListResponseZonesWebsocketsEditable = false
)

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [ZoneSettingEditResponseZones0rtt],
// [ZoneSettingEditResponseZonesAdvancedDDOS],
// [ZoneSettingEditResponseZonesAlwaysOnline],
// [ZoneSettingEditResponseZonesAlwaysUseHTTPS],
// [ZoneSettingEditResponseZonesAutomaticHTTPSRewrites],
// [ZoneSettingEditResponseZonesBrotli],
// [ZoneSettingEditResponseZonesBrowserCacheTTL],
// [ZoneSettingEditResponseZonesBrowserCheck],
// [ZoneSettingEditResponseZonesCacheLevel],
// [ZoneSettingEditResponseZonesChallengeTTL],
// [ZoneSettingEditResponseZonesCiphers],
// [ZoneSettingEditResponseZonesCnameFlattening],
// [ZoneSettingEditResponseZonesDevelopmentMode],
// [ZoneSettingEditResponseZonesEarlyHints],
// [ZoneSettingEditResponseZonesEdgeCacheTTL],
// [ZoneSettingEditResponseZonesEmailObfuscation],
// [ZoneSettingEditResponseZonesH2Prioritization],
// [ZoneSettingEditResponseZonesHotlinkProtection],
// [ZoneSettingEditResponseZonesHTTP2], [ZoneSettingEditResponseZonesHTTP3],
// [ZoneSettingEditResponseZonesImageResizing],
// [ZoneSettingEditResponseZonesIPGeolocation], [ZoneSettingEditResponseZonesIPV6],
// [ZoneSettingEditResponseZonesMaxUpload],
// [ZoneSettingEditResponseZonesMinTLSVersion],
// [ZoneSettingEditResponseZonesMinify], [ZoneSettingEditResponseZonesMirage],
// [ZoneSettingEditResponseZonesMobileRedirect], [ZoneSettingEditResponseZonesNEL],
// [ZoneSettingEditResponseZonesOpportunisticEncryption],
// [ZoneSettingEditResponseZonesOpportunisticOnion],
// [ZoneSettingEditResponseZonesOrangeToOrange],
// [ZoneSettingEditResponseZonesOriginErrorPagePassThru],
// [ZoneSettingEditResponseZonesPolish],
// [ZoneSettingEditResponseZonesPrefetchPreload],
// [ZoneSettingEditResponseZonesProxyReadTimeout],
// [ZoneSettingEditResponseZonesPseudoIPV4],
// [ZoneSettingEditResponseZonesResponseBuffering],
// [ZoneSettingEditResponseZonesRocketLoader],
// [ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [ZoneSettingEditResponseZonesSecurityHeader],
// [ZoneSettingEditResponseZonesSecurityLevel],
// [ZoneSettingEditResponseZonesServerSideExclude],
// [ZoneSettingEditResponseZonesSha1Support],
// [ZoneSettingEditResponseZonesSortQueryStringForCache],
// [ZoneSettingEditResponseZonesSSL], [ZoneSettingEditResponseZonesSSLRecommender],
// [ZoneSettingEditResponseZonesTLS1_2Only], [ZoneSettingEditResponseZonesTLS1_3],
// [ZoneSettingEditResponseZonesTLSClientAuth],
// [ZoneSettingEditResponseZonesTrueClientIPHeader],
// [ZoneSettingEditResponseZonesWAF], [ZoneSettingEditResponseZonesWebp] or
// [ZoneSettingEditResponseZonesWebsockets].
type ZoneSettingEditResponse interface {
	implementsZoneSettingEditResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneSettingEditResponse)(nil)).Elem(), "")
}

// 0-RTT session resumption enabled for this zone.
type ZoneSettingEditResponseZones0rtt struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZones0rttID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZones0rttValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZones0rttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZones0rttJSON `json:"-"`
}

// zoneSettingEditResponseZones0rttJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZones0rtt]
type zoneSettingEditResponseZones0rttJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZones0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZones0rtt) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZones0rttID string

const (
	ZoneSettingEditResponseZones0rttID0rtt ZoneSettingEditResponseZones0rttID = "0rtt"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZones0rttValue string

const (
	ZoneSettingEditResponseZones0rttValueOn  ZoneSettingEditResponseZones0rttValue = "on"
	ZoneSettingEditResponseZones0rttValueOff ZoneSettingEditResponseZones0rttValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZones0rttEditable bool

const (
	ZoneSettingEditResponseZones0rttEditableTrue  ZoneSettingEditResponseZones0rttEditable = true
	ZoneSettingEditResponseZones0rttEditableFalse ZoneSettingEditResponseZones0rttEditable = false
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingEditResponseZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesAdvancedDDOSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesAdvancedDDOSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesAdvancedDDOSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesAdvancedDDOSJSON `json:"-"`
}

// zoneSettingEditResponseZonesAdvancedDDOSJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesAdvancedDDOS]
type zoneSettingEditResponseZonesAdvancedDDOSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesAdvancedDDOS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesAdvancedDDOS) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesAdvancedDDOSID string

const (
	ZoneSettingEditResponseZonesAdvancedDDOSIDAdvancedDDOS ZoneSettingEditResponseZonesAdvancedDDOSID = "advanced_ddos"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesAdvancedDDOSValue string

const (
	ZoneSettingEditResponseZonesAdvancedDDOSValueOn  ZoneSettingEditResponseZonesAdvancedDDOSValue = "on"
	ZoneSettingEditResponseZonesAdvancedDDOSValueOff ZoneSettingEditResponseZonesAdvancedDDOSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesAdvancedDDOSEditable bool

const (
	ZoneSettingEditResponseZonesAdvancedDDOSEditableTrue  ZoneSettingEditResponseZonesAdvancedDDOSEditable = true
	ZoneSettingEditResponseZonesAdvancedDDOSEditableFalse ZoneSettingEditResponseZonesAdvancedDDOSEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingEditResponseZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesAlwaysOnlineID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesAlwaysOnlineValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesAlwaysOnlineJSON `json:"-"`
}

// zoneSettingEditResponseZonesAlwaysOnlineJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesAlwaysOnline]
type zoneSettingEditResponseZonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesAlwaysOnline) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesAlwaysOnlineID string

const (
	ZoneSettingEditResponseZonesAlwaysOnlineIDAlwaysOnline ZoneSettingEditResponseZonesAlwaysOnlineID = "always_online"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesAlwaysOnlineValue string

const (
	ZoneSettingEditResponseZonesAlwaysOnlineValueOn  ZoneSettingEditResponseZonesAlwaysOnlineValue = "on"
	ZoneSettingEditResponseZonesAlwaysOnlineValueOff ZoneSettingEditResponseZonesAlwaysOnlineValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesAlwaysOnlineEditable bool

const (
	ZoneSettingEditResponseZonesAlwaysOnlineEditableTrue  ZoneSettingEditResponseZonesAlwaysOnlineEditable = true
	ZoneSettingEditResponseZonesAlwaysOnlineEditableFalse ZoneSettingEditResponseZonesAlwaysOnlineEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingEditResponseZonesAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesAlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesAlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesAlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesAlwaysUseHTTPSJSON `json:"-"`
}

// zoneSettingEditResponseZonesAlwaysUseHTTPSJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesAlwaysUseHTTPS]
type zoneSettingEditResponseZonesAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesAlwaysUseHTTPS) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesAlwaysUseHTTPSID string

const (
	ZoneSettingEditResponseZonesAlwaysUseHTTPSIDAlwaysUseHTTPS ZoneSettingEditResponseZonesAlwaysUseHTTPSID = "always_use_https"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesAlwaysUseHTTPSValue string

const (
	ZoneSettingEditResponseZonesAlwaysUseHTTPSValueOn  ZoneSettingEditResponseZonesAlwaysUseHTTPSValue = "on"
	ZoneSettingEditResponseZonesAlwaysUseHTTPSValueOff ZoneSettingEditResponseZonesAlwaysUseHTTPSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesAlwaysUseHTTPSEditable bool

const (
	ZoneSettingEditResponseZonesAlwaysUseHTTPSEditableTrue  ZoneSettingEditResponseZonesAlwaysUseHTTPSEditable = true
	ZoneSettingEditResponseZonesAlwaysUseHTTPSEditableFalse ZoneSettingEditResponseZonesAlwaysUseHTTPSEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingEditResponseZonesAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesAutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesAutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesAutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesAutomaticHTTPSRewritesJSON `json:"-"`
}

// zoneSettingEditResponseZonesAutomaticHTTPSRewritesJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseZonesAutomaticHTTPSRewrites]
type zoneSettingEditResponseZonesAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesAutomaticHTTPSRewrites) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesAutomaticHTTPSRewritesID string

const (
	ZoneSettingEditResponseZonesAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites ZoneSettingEditResponseZonesAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesAutomaticHTTPSRewritesValue string

const (
	ZoneSettingEditResponseZonesAutomaticHTTPSRewritesValueOn  ZoneSettingEditResponseZonesAutomaticHTTPSRewritesValue = "on"
	ZoneSettingEditResponseZonesAutomaticHTTPSRewritesValueOff ZoneSettingEditResponseZonesAutomaticHTTPSRewritesValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesAutomaticHTTPSRewritesEditable bool

const (
	ZoneSettingEditResponseZonesAutomaticHTTPSRewritesEditableTrue  ZoneSettingEditResponseZonesAutomaticHTTPSRewritesEditable = true
	ZoneSettingEditResponseZonesAutomaticHTTPSRewritesEditableFalse ZoneSettingEditResponseZonesAutomaticHTTPSRewritesEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingEditResponseZonesBrotli struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesBrotliID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesBrotliValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesBrotliJSON `json:"-"`
}

// zoneSettingEditResponseZonesBrotliJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesBrotli]
type zoneSettingEditResponseZonesBrotliJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesBrotli) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesBrotliID string

const (
	ZoneSettingEditResponseZonesBrotliIDBrotli ZoneSettingEditResponseZonesBrotliID = "brotli"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesBrotliValue string

const (
	ZoneSettingEditResponseZonesBrotliValueOff ZoneSettingEditResponseZonesBrotliValue = "off"
	ZoneSettingEditResponseZonesBrotliValueOn  ZoneSettingEditResponseZonesBrotliValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesBrotliEditable bool

const (
	ZoneSettingEditResponseZonesBrotliEditableTrue  ZoneSettingEditResponseZonesBrotliEditable = true
	ZoneSettingEditResponseZonesBrotliEditableFalse ZoneSettingEditResponseZonesBrotliEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingEditResponseZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesBrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesBrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesBrowserCacheTTLJSON `json:"-"`
}

// zoneSettingEditResponseZonesBrowserCacheTTLJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesBrowserCacheTTL]
type zoneSettingEditResponseZonesBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesBrowserCacheTTL) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesBrowserCacheTTLID string

const (
	ZoneSettingEditResponseZonesBrowserCacheTTLIDBrowserCacheTTL ZoneSettingEditResponseZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesBrowserCacheTTLValue float64

const (
	ZoneSettingEditResponseZonesBrowserCacheTTLValue0        ZoneSettingEditResponseZonesBrowserCacheTTLValue = 0
	ZoneSettingEditResponseZonesBrowserCacheTTLValue30       ZoneSettingEditResponseZonesBrowserCacheTTLValue = 30
	ZoneSettingEditResponseZonesBrowserCacheTTLValue60       ZoneSettingEditResponseZonesBrowserCacheTTLValue = 60
	ZoneSettingEditResponseZonesBrowserCacheTTLValue120      ZoneSettingEditResponseZonesBrowserCacheTTLValue = 120
	ZoneSettingEditResponseZonesBrowserCacheTTLValue300      ZoneSettingEditResponseZonesBrowserCacheTTLValue = 300
	ZoneSettingEditResponseZonesBrowserCacheTTLValue1200     ZoneSettingEditResponseZonesBrowserCacheTTLValue = 1200
	ZoneSettingEditResponseZonesBrowserCacheTTLValue1800     ZoneSettingEditResponseZonesBrowserCacheTTLValue = 1800
	ZoneSettingEditResponseZonesBrowserCacheTTLValue3600     ZoneSettingEditResponseZonesBrowserCacheTTLValue = 3600
	ZoneSettingEditResponseZonesBrowserCacheTTLValue7200     ZoneSettingEditResponseZonesBrowserCacheTTLValue = 7200
	ZoneSettingEditResponseZonesBrowserCacheTTLValue10800    ZoneSettingEditResponseZonesBrowserCacheTTLValue = 10800
	ZoneSettingEditResponseZonesBrowserCacheTTLValue14400    ZoneSettingEditResponseZonesBrowserCacheTTLValue = 14400
	ZoneSettingEditResponseZonesBrowserCacheTTLValue18000    ZoneSettingEditResponseZonesBrowserCacheTTLValue = 18000
	ZoneSettingEditResponseZonesBrowserCacheTTLValue28800    ZoneSettingEditResponseZonesBrowserCacheTTLValue = 28800
	ZoneSettingEditResponseZonesBrowserCacheTTLValue43200    ZoneSettingEditResponseZonesBrowserCacheTTLValue = 43200
	ZoneSettingEditResponseZonesBrowserCacheTTLValue57600    ZoneSettingEditResponseZonesBrowserCacheTTLValue = 57600
	ZoneSettingEditResponseZonesBrowserCacheTTLValue72000    ZoneSettingEditResponseZonesBrowserCacheTTLValue = 72000
	ZoneSettingEditResponseZonesBrowserCacheTTLValue86400    ZoneSettingEditResponseZonesBrowserCacheTTLValue = 86400
	ZoneSettingEditResponseZonesBrowserCacheTTLValue172800   ZoneSettingEditResponseZonesBrowserCacheTTLValue = 172800
	ZoneSettingEditResponseZonesBrowserCacheTTLValue259200   ZoneSettingEditResponseZonesBrowserCacheTTLValue = 259200
	ZoneSettingEditResponseZonesBrowserCacheTTLValue345600   ZoneSettingEditResponseZonesBrowserCacheTTLValue = 345600
	ZoneSettingEditResponseZonesBrowserCacheTTLValue432000   ZoneSettingEditResponseZonesBrowserCacheTTLValue = 432000
	ZoneSettingEditResponseZonesBrowserCacheTTLValue691200   ZoneSettingEditResponseZonesBrowserCacheTTLValue = 691200
	ZoneSettingEditResponseZonesBrowserCacheTTLValue1382400  ZoneSettingEditResponseZonesBrowserCacheTTLValue = 1382400
	ZoneSettingEditResponseZonesBrowserCacheTTLValue2073600  ZoneSettingEditResponseZonesBrowserCacheTTLValue = 2073600
	ZoneSettingEditResponseZonesBrowserCacheTTLValue2678400  ZoneSettingEditResponseZonesBrowserCacheTTLValue = 2678400
	ZoneSettingEditResponseZonesBrowserCacheTTLValue5356800  ZoneSettingEditResponseZonesBrowserCacheTTLValue = 5356800
	ZoneSettingEditResponseZonesBrowserCacheTTLValue16070400 ZoneSettingEditResponseZonesBrowserCacheTTLValue = 16070400
	ZoneSettingEditResponseZonesBrowserCacheTTLValue31536000 ZoneSettingEditResponseZonesBrowserCacheTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesBrowserCacheTTLEditable bool

const (
	ZoneSettingEditResponseZonesBrowserCacheTTLEditableTrue  ZoneSettingEditResponseZonesBrowserCacheTTLEditable = true
	ZoneSettingEditResponseZonesBrowserCacheTTLEditableFalse ZoneSettingEditResponseZonesBrowserCacheTTLEditable = false
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingEditResponseZonesBrowserCheck struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesBrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesBrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesBrowserCheckJSON `json:"-"`
}

// zoneSettingEditResponseZonesBrowserCheckJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesBrowserCheck]
type zoneSettingEditResponseZonesBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesBrowserCheck) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesBrowserCheckID string

const (
	ZoneSettingEditResponseZonesBrowserCheckIDBrowserCheck ZoneSettingEditResponseZonesBrowserCheckID = "browser_check"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesBrowserCheckValue string

const (
	ZoneSettingEditResponseZonesBrowserCheckValueOn  ZoneSettingEditResponseZonesBrowserCheckValue = "on"
	ZoneSettingEditResponseZonesBrowserCheckValueOff ZoneSettingEditResponseZonesBrowserCheckValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesBrowserCheckEditable bool

const (
	ZoneSettingEditResponseZonesBrowserCheckEditableTrue  ZoneSettingEditResponseZonesBrowserCheckEditable = true
	ZoneSettingEditResponseZonesBrowserCheckEditableFalse ZoneSettingEditResponseZonesBrowserCheckEditable = false
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingEditResponseZonesCacheLevel struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesCacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesCacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesCacheLevelJSON `json:"-"`
}

// zoneSettingEditResponseZonesCacheLevelJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesCacheLevel]
type zoneSettingEditResponseZonesCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesCacheLevel) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesCacheLevelID string

const (
	ZoneSettingEditResponseZonesCacheLevelIDCacheLevel ZoneSettingEditResponseZonesCacheLevelID = "cache_level"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesCacheLevelValue string

const (
	ZoneSettingEditResponseZonesCacheLevelValueAggressive ZoneSettingEditResponseZonesCacheLevelValue = "aggressive"
	ZoneSettingEditResponseZonesCacheLevelValueBasic      ZoneSettingEditResponseZonesCacheLevelValue = "basic"
	ZoneSettingEditResponseZonesCacheLevelValueSimplified ZoneSettingEditResponseZonesCacheLevelValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesCacheLevelEditable bool

const (
	ZoneSettingEditResponseZonesCacheLevelEditableTrue  ZoneSettingEditResponseZonesCacheLevelEditable = true
	ZoneSettingEditResponseZonesCacheLevelEditableFalse ZoneSettingEditResponseZonesCacheLevelEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingEditResponseZonesChallengeTTL struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesChallengeTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesChallengeTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesChallengeTTLJSON `json:"-"`
}

// zoneSettingEditResponseZonesChallengeTTLJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesChallengeTTL]
type zoneSettingEditResponseZonesChallengeTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesChallengeTTL) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesChallengeTTLID string

const (
	ZoneSettingEditResponseZonesChallengeTTLIDChallengeTTL ZoneSettingEditResponseZonesChallengeTTLID = "challenge_ttl"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesChallengeTTLValue float64

const (
	ZoneSettingEditResponseZonesChallengeTTLValue300      ZoneSettingEditResponseZonesChallengeTTLValue = 300
	ZoneSettingEditResponseZonesChallengeTTLValue900      ZoneSettingEditResponseZonesChallengeTTLValue = 900
	ZoneSettingEditResponseZonesChallengeTTLValue1800     ZoneSettingEditResponseZonesChallengeTTLValue = 1800
	ZoneSettingEditResponseZonesChallengeTTLValue2700     ZoneSettingEditResponseZonesChallengeTTLValue = 2700
	ZoneSettingEditResponseZonesChallengeTTLValue3600     ZoneSettingEditResponseZonesChallengeTTLValue = 3600
	ZoneSettingEditResponseZonesChallengeTTLValue7200     ZoneSettingEditResponseZonesChallengeTTLValue = 7200
	ZoneSettingEditResponseZonesChallengeTTLValue10800    ZoneSettingEditResponseZonesChallengeTTLValue = 10800
	ZoneSettingEditResponseZonesChallengeTTLValue14400    ZoneSettingEditResponseZonesChallengeTTLValue = 14400
	ZoneSettingEditResponseZonesChallengeTTLValue28800    ZoneSettingEditResponseZonesChallengeTTLValue = 28800
	ZoneSettingEditResponseZonesChallengeTTLValue57600    ZoneSettingEditResponseZonesChallengeTTLValue = 57600
	ZoneSettingEditResponseZonesChallengeTTLValue86400    ZoneSettingEditResponseZonesChallengeTTLValue = 86400
	ZoneSettingEditResponseZonesChallengeTTLValue604800   ZoneSettingEditResponseZonesChallengeTTLValue = 604800
	ZoneSettingEditResponseZonesChallengeTTLValue2592000  ZoneSettingEditResponseZonesChallengeTTLValue = 2592000
	ZoneSettingEditResponseZonesChallengeTTLValue31536000 ZoneSettingEditResponseZonesChallengeTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesChallengeTTLEditable bool

const (
	ZoneSettingEditResponseZonesChallengeTTLEditableTrue  ZoneSettingEditResponseZonesChallengeTTLEditable = true
	ZoneSettingEditResponseZonesChallengeTTLEditableFalse ZoneSettingEditResponseZonesChallengeTTLEditable = false
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingEditResponseZonesCiphers struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesCiphersID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesCiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesCiphersJSON `json:"-"`
}

// zoneSettingEditResponseZonesCiphersJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesCiphers]
type zoneSettingEditResponseZonesCiphersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesCiphers) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesCiphersID string

const (
	ZoneSettingEditResponseZonesCiphersIDCiphers ZoneSettingEditResponseZonesCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesCiphersEditable bool

const (
	ZoneSettingEditResponseZonesCiphersEditableTrue  ZoneSettingEditResponseZonesCiphersEditable = true
	ZoneSettingEditResponseZonesCiphersEditableFalse ZoneSettingEditResponseZonesCiphersEditable = false
)

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

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingEditResponseZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesDevelopmentModeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesDevelopmentModeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                                         `json:"time_remaining"`
	JSON          zoneSettingEditResponseZonesDevelopmentModeJSON `json:"-"`
}

// zoneSettingEditResponseZonesDevelopmentModeJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesDevelopmentMode]
type zoneSettingEditResponseZonesDevelopmentModeJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesDevelopmentMode) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesDevelopmentModeID string

const (
	ZoneSettingEditResponseZonesDevelopmentModeIDDevelopmentMode ZoneSettingEditResponseZonesDevelopmentModeID = "development_mode"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesDevelopmentModeValue string

const (
	ZoneSettingEditResponseZonesDevelopmentModeValueOn  ZoneSettingEditResponseZonesDevelopmentModeValue = "on"
	ZoneSettingEditResponseZonesDevelopmentModeValueOff ZoneSettingEditResponseZonesDevelopmentModeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesDevelopmentModeEditable bool

const (
	ZoneSettingEditResponseZonesDevelopmentModeEditableTrue  ZoneSettingEditResponseZonesDevelopmentModeEditable = true
	ZoneSettingEditResponseZonesDevelopmentModeEditableFalse ZoneSettingEditResponseZonesDevelopmentModeEditable = false
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingEditResponseZonesEarlyHints struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesEarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesEarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesEarlyHintsJSON `json:"-"`
}

// zoneSettingEditResponseZonesEarlyHintsJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesEarlyHints]
type zoneSettingEditResponseZonesEarlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesEarlyHints) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesEarlyHintsID string

const (
	ZoneSettingEditResponseZonesEarlyHintsIDEarlyHints ZoneSettingEditResponseZonesEarlyHintsID = "early_hints"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesEarlyHintsValue string

const (
	ZoneSettingEditResponseZonesEarlyHintsValueOn  ZoneSettingEditResponseZonesEarlyHintsValue = "on"
	ZoneSettingEditResponseZonesEarlyHintsValueOff ZoneSettingEditResponseZonesEarlyHintsValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesEarlyHintsEditable bool

const (
	ZoneSettingEditResponseZonesEarlyHintsEditableTrue  ZoneSettingEditResponseZonesEarlyHintsEditable = true
	ZoneSettingEditResponseZonesEarlyHintsEditableFalse ZoneSettingEditResponseZonesEarlyHintsEditable = false
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

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingEditResponseZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesEmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesEmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesEmailObfuscationJSON `json:"-"`
}

// zoneSettingEditResponseZonesEmailObfuscationJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesEmailObfuscation]
type zoneSettingEditResponseZonesEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesEmailObfuscation) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesEmailObfuscationID string

const (
	ZoneSettingEditResponseZonesEmailObfuscationIDEmailObfuscation ZoneSettingEditResponseZonesEmailObfuscationID = "email_obfuscation"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesEmailObfuscationValue string

const (
	ZoneSettingEditResponseZonesEmailObfuscationValueOn  ZoneSettingEditResponseZonesEmailObfuscationValue = "on"
	ZoneSettingEditResponseZonesEmailObfuscationValueOff ZoneSettingEditResponseZonesEmailObfuscationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesEmailObfuscationEditable bool

const (
	ZoneSettingEditResponseZonesEmailObfuscationEditableTrue  ZoneSettingEditResponseZonesEmailObfuscationEditable = true
	ZoneSettingEditResponseZonesEmailObfuscationEditableFalse ZoneSettingEditResponseZonesEmailObfuscationEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingEditResponseZonesH2Prioritization struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesH2PrioritizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesH2PrioritizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesH2PrioritizationJSON `json:"-"`
}

// zoneSettingEditResponseZonesH2PrioritizationJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesH2Prioritization]
type zoneSettingEditResponseZonesH2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesH2Prioritization) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesH2PrioritizationID string

const (
	ZoneSettingEditResponseZonesH2PrioritizationIDH2Prioritization ZoneSettingEditResponseZonesH2PrioritizationID = "h2_prioritization"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesH2PrioritizationValue string

const (
	ZoneSettingEditResponseZonesH2PrioritizationValueOn     ZoneSettingEditResponseZonesH2PrioritizationValue = "on"
	ZoneSettingEditResponseZonesH2PrioritizationValueOff    ZoneSettingEditResponseZonesH2PrioritizationValue = "off"
	ZoneSettingEditResponseZonesH2PrioritizationValueCustom ZoneSettingEditResponseZonesH2PrioritizationValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesH2PrioritizationEditable bool

const (
	ZoneSettingEditResponseZonesH2PrioritizationEditableTrue  ZoneSettingEditResponseZonesH2PrioritizationEditable = true
	ZoneSettingEditResponseZonesH2PrioritizationEditableFalse ZoneSettingEditResponseZonesH2PrioritizationEditable = false
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZoneSettingEditResponseZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesHotlinkProtectionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesHotlinkProtectionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesHotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesHotlinkProtectionJSON `json:"-"`
}

// zoneSettingEditResponseZonesHotlinkProtectionJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesHotlinkProtection]
type zoneSettingEditResponseZonesHotlinkProtectionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesHotlinkProtection) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesHotlinkProtectionID string

const (
	ZoneSettingEditResponseZonesHotlinkProtectionIDHotlinkProtection ZoneSettingEditResponseZonesHotlinkProtectionID = "hotlink_protection"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesHotlinkProtectionValue string

const (
	ZoneSettingEditResponseZonesHotlinkProtectionValueOn  ZoneSettingEditResponseZonesHotlinkProtectionValue = "on"
	ZoneSettingEditResponseZonesHotlinkProtectionValueOff ZoneSettingEditResponseZonesHotlinkProtectionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesHotlinkProtectionEditable bool

const (
	ZoneSettingEditResponseZonesHotlinkProtectionEditableTrue  ZoneSettingEditResponseZonesHotlinkProtectionEditable = true
	ZoneSettingEditResponseZonesHotlinkProtectionEditableFalse ZoneSettingEditResponseZonesHotlinkProtectionEditable = false
)

// HTTP2 enabled for this zone.
type ZoneSettingEditResponseZonesHTTP2 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesHTTP2ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesHTTP2Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesHTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesHTTP2JSON `json:"-"`
}

// zoneSettingEditResponseZonesHTTP2JSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesHTTP2]
type zoneSettingEditResponseZonesHTTP2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesHTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesHTTP2) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesHTTP2ID string

const (
	ZoneSettingEditResponseZonesHTTP2IDHTTP2 ZoneSettingEditResponseZonesHTTP2ID = "http2"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesHTTP2Value string

const (
	ZoneSettingEditResponseZonesHTTP2ValueOn  ZoneSettingEditResponseZonesHTTP2Value = "on"
	ZoneSettingEditResponseZonesHTTP2ValueOff ZoneSettingEditResponseZonesHTTP2Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesHTTP2Editable bool

const (
	ZoneSettingEditResponseZonesHTTP2EditableTrue  ZoneSettingEditResponseZonesHTTP2Editable = true
	ZoneSettingEditResponseZonesHTTP2EditableFalse ZoneSettingEditResponseZonesHTTP2Editable = false
)

// HTTP3 enabled for this zone.
type ZoneSettingEditResponseZonesHTTP3 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesHTTP3ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesHTTP3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesHTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesHTTP3JSON `json:"-"`
}

// zoneSettingEditResponseZonesHTTP3JSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesHTTP3]
type zoneSettingEditResponseZonesHTTP3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesHTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesHTTP3) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesHTTP3ID string

const (
	ZoneSettingEditResponseZonesHTTP3IDHTTP3 ZoneSettingEditResponseZonesHTTP3ID = "http3"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesHTTP3Value string

const (
	ZoneSettingEditResponseZonesHTTP3ValueOn  ZoneSettingEditResponseZonesHTTP3Value = "on"
	ZoneSettingEditResponseZonesHTTP3ValueOff ZoneSettingEditResponseZonesHTTP3Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesHTTP3Editable bool

const (
	ZoneSettingEditResponseZonesHTTP3EditableTrue  ZoneSettingEditResponseZonesHTTP3Editable = true
	ZoneSettingEditResponseZonesHTTP3EditableFalse ZoneSettingEditResponseZonesHTTP3Editable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingEditResponseZonesImageResizing struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesImageResizingID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesImageResizingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesImageResizingJSON `json:"-"`
}

// zoneSettingEditResponseZonesImageResizingJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesImageResizing]
type zoneSettingEditResponseZonesImageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesImageResizing) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesImageResizingID string

const (
	ZoneSettingEditResponseZonesImageResizingIDImageResizing ZoneSettingEditResponseZonesImageResizingID = "image_resizing"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesImageResizingValue string

const (
	ZoneSettingEditResponseZonesImageResizingValueOn   ZoneSettingEditResponseZonesImageResizingValue = "on"
	ZoneSettingEditResponseZonesImageResizingValueOff  ZoneSettingEditResponseZonesImageResizingValue = "off"
	ZoneSettingEditResponseZonesImageResizingValueOpen ZoneSettingEditResponseZonesImageResizingValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesImageResizingEditable bool

const (
	ZoneSettingEditResponseZonesImageResizingEditableTrue  ZoneSettingEditResponseZonesImageResizingEditable = true
	ZoneSettingEditResponseZonesImageResizingEditableFalse ZoneSettingEditResponseZonesImageResizingEditable = false
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingEditResponseZonesIPGeolocation struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesIPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesIPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesIPGeolocationJSON `json:"-"`
}

// zoneSettingEditResponseZonesIPGeolocationJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesIPGeolocation]
type zoneSettingEditResponseZonesIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesIPGeolocation) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesIPGeolocationID string

const (
	ZoneSettingEditResponseZonesIPGeolocationIDIPGeolocation ZoneSettingEditResponseZonesIPGeolocationID = "ip_geolocation"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesIPGeolocationValue string

const (
	ZoneSettingEditResponseZonesIPGeolocationValueOn  ZoneSettingEditResponseZonesIPGeolocationValue = "on"
	ZoneSettingEditResponseZonesIPGeolocationValueOff ZoneSettingEditResponseZonesIPGeolocationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesIPGeolocationEditable bool

const (
	ZoneSettingEditResponseZonesIPGeolocationEditableTrue  ZoneSettingEditResponseZonesIPGeolocationEditable = true
	ZoneSettingEditResponseZonesIPGeolocationEditableFalse ZoneSettingEditResponseZonesIPGeolocationEditable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingEditResponseZonesIPV6 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesIPV6ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesIPV6Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesIPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesIPV6JSON `json:"-"`
}

// zoneSettingEditResponseZonesIPV6JSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesIPV6]
type zoneSettingEditResponseZonesIPV6JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesIPV6) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesIPV6ID string

const (
	ZoneSettingEditResponseZonesIPV6IDIPV6 ZoneSettingEditResponseZonesIPV6ID = "ipv6"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesIPV6Value string

const (
	ZoneSettingEditResponseZonesIPV6ValueOff ZoneSettingEditResponseZonesIPV6Value = "off"
	ZoneSettingEditResponseZonesIPV6ValueOn  ZoneSettingEditResponseZonesIPV6Value = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesIPV6Editable bool

const (
	ZoneSettingEditResponseZonesIPV6EditableTrue  ZoneSettingEditResponseZonesIPV6Editable = true
	ZoneSettingEditResponseZonesIPV6EditableFalse ZoneSettingEditResponseZonesIPV6Editable = false
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

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingEditResponseZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesMinTLSVersionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesMinTLSVersionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesMinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesMinTLSVersionJSON `json:"-"`
}

// zoneSettingEditResponseZonesMinTLSVersionJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesMinTLSVersion]
type zoneSettingEditResponseZonesMinTLSVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesMinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesMinTLSVersion) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesMinTLSVersionID string

const (
	ZoneSettingEditResponseZonesMinTLSVersionIDMinTLSVersion ZoneSettingEditResponseZonesMinTLSVersionID = "min_tls_version"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesMinTLSVersionValue string

const (
	ZoneSettingEditResponseZonesMinTLSVersionValue1_0 ZoneSettingEditResponseZonesMinTLSVersionValue = "1.0"
	ZoneSettingEditResponseZonesMinTLSVersionValue1_1 ZoneSettingEditResponseZonesMinTLSVersionValue = "1.1"
	ZoneSettingEditResponseZonesMinTLSVersionValue1_2 ZoneSettingEditResponseZonesMinTLSVersionValue = "1.2"
	ZoneSettingEditResponseZonesMinTLSVersionValue1_3 ZoneSettingEditResponseZonesMinTLSVersionValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesMinTLSVersionEditable bool

const (
	ZoneSettingEditResponseZonesMinTLSVersionEditableTrue  ZoneSettingEditResponseZonesMinTLSVersionEditable = true
	ZoneSettingEditResponseZonesMinTLSVersionEditableFalse ZoneSettingEditResponseZonesMinTLSVersionEditable = false
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingEditResponseZonesMinify struct {
	// Zone setting identifier.
	ID ZoneSettingEditResponseZonesMinifyID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesMinifyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesMinifyJSON `json:"-"`
}

// zoneSettingEditResponseZonesMinifyJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesMinify]
type zoneSettingEditResponseZonesMinifyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesMinify) implementsZoneSettingEditResponse() {}

// Zone setting identifier.
type ZoneSettingEditResponseZonesMinifyID string

const (
	ZoneSettingEditResponseZonesMinifyIDMinify ZoneSettingEditResponseZonesMinifyID = "minify"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css ZoneSettingEditResponseZonesMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML ZoneSettingEditResponseZonesMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   ZoneSettingEditResponseZonesMinifyValueJs   `json:"js"`
	JSON zoneSettingEditResponseZonesMinifyValueJSON `json:"-"`
}

// zoneSettingEditResponseZonesMinifyValueJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesMinifyValue]
type zoneSettingEditResponseZonesMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingEditResponseZonesMinifyValueCss string

const (
	ZoneSettingEditResponseZonesMinifyValueCssOn  ZoneSettingEditResponseZonesMinifyValueCss = "on"
	ZoneSettingEditResponseZonesMinifyValueCssOff ZoneSettingEditResponseZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingEditResponseZonesMinifyValueHTML string

const (
	ZoneSettingEditResponseZonesMinifyValueHTMLOn  ZoneSettingEditResponseZonesMinifyValueHTML = "on"
	ZoneSettingEditResponseZonesMinifyValueHTMLOff ZoneSettingEditResponseZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingEditResponseZonesMinifyValueJs string

const (
	ZoneSettingEditResponseZonesMinifyValueJsOn  ZoneSettingEditResponseZonesMinifyValueJs = "on"
	ZoneSettingEditResponseZonesMinifyValueJsOff ZoneSettingEditResponseZonesMinifyValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesMinifyEditable bool

const (
	ZoneSettingEditResponseZonesMinifyEditableTrue  ZoneSettingEditResponseZonesMinifyEditable = true
	ZoneSettingEditResponseZonesMinifyEditableFalse ZoneSettingEditResponseZonesMinifyEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingEditResponseZonesMirage struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesMirageID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesMirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesMirageJSON `json:"-"`
}

// zoneSettingEditResponseZonesMirageJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesMirage]
type zoneSettingEditResponseZonesMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesMirage) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesMirageID string

const (
	ZoneSettingEditResponseZonesMirageIDMirage ZoneSettingEditResponseZonesMirageID = "mirage"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesMirageValue string

const (
	ZoneSettingEditResponseZonesMirageValueOn  ZoneSettingEditResponseZonesMirageValue = "on"
	ZoneSettingEditResponseZonesMirageValueOff ZoneSettingEditResponseZonesMirageValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesMirageEditable bool

const (
	ZoneSettingEditResponseZonesMirageEditableTrue  ZoneSettingEditResponseZonesMirageEditable = true
	ZoneSettingEditResponseZonesMirageEditableFalse ZoneSettingEditResponseZonesMirageEditable = false
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingEditResponseZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID ZoneSettingEditResponseZonesMobileRedirectID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesMobileRedirectValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesMobileRedirectJSON `json:"-"`
}

// zoneSettingEditResponseZonesMobileRedirectJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesMobileRedirect]
type zoneSettingEditResponseZonesMobileRedirectJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesMobileRedirect) implementsZoneSettingEditResponse() {}

// Identifier of the zone setting.
type ZoneSettingEditResponseZonesMobileRedirectID string

const (
	ZoneSettingEditResponseZonesMobileRedirectIDMobileRedirect ZoneSettingEditResponseZonesMobileRedirectID = "mobile_redirect"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingEditResponseZonesMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                                `json:"strip_uri"`
	JSON     zoneSettingEditResponseZonesMobileRedirectValueJSON `json:"-"`
}

// zoneSettingEditResponseZonesMobileRedirectValueJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseZonesMobileRedirectValue]
type zoneSettingEditResponseZonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingEditResponseZonesMobileRedirectValueStatus string

const (
	ZoneSettingEditResponseZonesMobileRedirectValueStatusOn  ZoneSettingEditResponseZonesMobileRedirectValueStatus = "on"
	ZoneSettingEditResponseZonesMobileRedirectValueStatusOff ZoneSettingEditResponseZonesMobileRedirectValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesMobileRedirectEditable bool

const (
	ZoneSettingEditResponseZonesMobileRedirectEditableTrue  ZoneSettingEditResponseZonesMobileRedirectEditable = true
	ZoneSettingEditResponseZonesMobileRedirectEditableFalse ZoneSettingEditResponseZonesMobileRedirectEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingEditResponseZonesNEL struct {
	// Zone setting identifier.
	ID ZoneSettingEditResponseZonesNELID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesNELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesNELJSON `json:"-"`
}

// zoneSettingEditResponseZonesNELJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesNEL]
type zoneSettingEditResponseZonesNELJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesNEL) implementsZoneSettingEditResponse() {}

// Zone setting identifier.
type ZoneSettingEditResponseZonesNELID string

const (
	ZoneSettingEditResponseZonesNELIDNEL ZoneSettingEditResponseZonesNELID = "nel"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesNELValue struct {
	Enabled bool                                     `json:"enabled"`
	JSON    zoneSettingEditResponseZonesNELValueJSON `json:"-"`
}

// zoneSettingEditResponseZonesNELValueJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesNELValue]
type zoneSettingEditResponseZonesNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesNELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesNELEditable bool

const (
	ZoneSettingEditResponseZonesNELEditableTrue  ZoneSettingEditResponseZonesNELEditable = true
	ZoneSettingEditResponseZonesNELEditableFalse ZoneSettingEditResponseZonesNELEditable = false
)

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingEditResponseZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesOpportunisticEncryptionJSON `json:"-"`
}

// zoneSettingEditResponseZonesOpportunisticEncryptionJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseZonesOpportunisticEncryption]
type zoneSettingEditResponseZonesOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesOpportunisticEncryption) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesOpportunisticEncryptionID string

const (
	ZoneSettingEditResponseZonesOpportunisticEncryptionIDOpportunisticEncryption ZoneSettingEditResponseZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesOpportunisticEncryptionValue string

const (
	ZoneSettingEditResponseZonesOpportunisticEncryptionValueOn  ZoneSettingEditResponseZonesOpportunisticEncryptionValue = "on"
	ZoneSettingEditResponseZonesOpportunisticEncryptionValueOff ZoneSettingEditResponseZonesOpportunisticEncryptionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesOpportunisticEncryptionEditable bool

const (
	ZoneSettingEditResponseZonesOpportunisticEncryptionEditableTrue  ZoneSettingEditResponseZonesOpportunisticEncryptionEditable = true
	ZoneSettingEditResponseZonesOpportunisticEncryptionEditableFalse ZoneSettingEditResponseZonesOpportunisticEncryptionEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingEditResponseZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesOpportunisticOnionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesOpportunisticOnionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesOpportunisticOnionJSON `json:"-"`
}

// zoneSettingEditResponseZonesOpportunisticOnionJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseZonesOpportunisticOnion]
type zoneSettingEditResponseZonesOpportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesOpportunisticOnion) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesOpportunisticOnionID string

const (
	ZoneSettingEditResponseZonesOpportunisticOnionIDOpportunisticOnion ZoneSettingEditResponseZonesOpportunisticOnionID = "opportunistic_onion"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesOpportunisticOnionValue string

const (
	ZoneSettingEditResponseZonesOpportunisticOnionValueOn  ZoneSettingEditResponseZonesOpportunisticOnionValue = "on"
	ZoneSettingEditResponseZonesOpportunisticOnionValueOff ZoneSettingEditResponseZonesOpportunisticOnionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesOpportunisticOnionEditable bool

const (
	ZoneSettingEditResponseZonesOpportunisticOnionEditableTrue  ZoneSettingEditResponseZonesOpportunisticOnionEditable = true
	ZoneSettingEditResponseZonesOpportunisticOnionEditableFalse ZoneSettingEditResponseZonesOpportunisticOnionEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingEditResponseZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesOrangeToOrangeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesOrangeToOrangeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesOrangeToOrangeJSON `json:"-"`
}

// zoneSettingEditResponseZonesOrangeToOrangeJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesOrangeToOrange]
type zoneSettingEditResponseZonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesOrangeToOrange) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesOrangeToOrangeID string

const (
	ZoneSettingEditResponseZonesOrangeToOrangeIDOrangeToOrange ZoneSettingEditResponseZonesOrangeToOrangeID = "orange_to_orange"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesOrangeToOrangeValue string

const (
	ZoneSettingEditResponseZonesOrangeToOrangeValueOn  ZoneSettingEditResponseZonesOrangeToOrangeValue = "on"
	ZoneSettingEditResponseZonesOrangeToOrangeValueOff ZoneSettingEditResponseZonesOrangeToOrangeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesOrangeToOrangeEditable bool

const (
	ZoneSettingEditResponseZonesOrangeToOrangeEditableTrue  ZoneSettingEditResponseZonesOrangeToOrangeEditable = true
	ZoneSettingEditResponseZonesOrangeToOrangeEditableFalse ZoneSettingEditResponseZonesOrangeToOrangeEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZoneSettingEditResponseZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesOriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesOriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesOriginErrorPagePassThruJSON `json:"-"`
}

// zoneSettingEditResponseZonesOriginErrorPagePassThruJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseZonesOriginErrorPagePassThru]
type zoneSettingEditResponseZonesOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesOriginErrorPagePassThru) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesOriginErrorPagePassThruID string

const (
	ZoneSettingEditResponseZonesOriginErrorPagePassThruIDOriginErrorPagePassThru ZoneSettingEditResponseZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesOriginErrorPagePassThruValue string

const (
	ZoneSettingEditResponseZonesOriginErrorPagePassThruValueOn  ZoneSettingEditResponseZonesOriginErrorPagePassThruValue = "on"
	ZoneSettingEditResponseZonesOriginErrorPagePassThruValueOff ZoneSettingEditResponseZonesOriginErrorPagePassThruValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesOriginErrorPagePassThruEditable bool

const (
	ZoneSettingEditResponseZonesOriginErrorPagePassThruEditableTrue  ZoneSettingEditResponseZonesOriginErrorPagePassThruEditable = true
	ZoneSettingEditResponseZonesOriginErrorPagePassThruEditableFalse ZoneSettingEditResponseZonesOriginErrorPagePassThruEditable = false
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingEditResponseZonesPolish struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesPolishID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesPolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesPolishJSON `json:"-"`
}

// zoneSettingEditResponseZonesPolishJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesPolish]
type zoneSettingEditResponseZonesPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesPolish) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesPolishID string

const (
	ZoneSettingEditResponseZonesPolishIDPolish ZoneSettingEditResponseZonesPolishID = "polish"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesPolishValue string

const (
	ZoneSettingEditResponseZonesPolishValueOff      ZoneSettingEditResponseZonesPolishValue = "off"
	ZoneSettingEditResponseZonesPolishValueLossless ZoneSettingEditResponseZonesPolishValue = "lossless"
	ZoneSettingEditResponseZonesPolishValueLossy    ZoneSettingEditResponseZonesPolishValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesPolishEditable bool

const (
	ZoneSettingEditResponseZonesPolishEditableTrue  ZoneSettingEditResponseZonesPolishEditable = true
	ZoneSettingEditResponseZonesPolishEditableFalse ZoneSettingEditResponseZonesPolishEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingEditResponseZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesPrefetchPreloadID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesPrefetchPreloadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesPrefetchPreloadJSON `json:"-"`
}

// zoneSettingEditResponseZonesPrefetchPreloadJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesPrefetchPreload]
type zoneSettingEditResponseZonesPrefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesPrefetchPreload) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesPrefetchPreloadID string

const (
	ZoneSettingEditResponseZonesPrefetchPreloadIDPrefetchPreload ZoneSettingEditResponseZonesPrefetchPreloadID = "prefetch_preload"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesPrefetchPreloadValue string

const (
	ZoneSettingEditResponseZonesPrefetchPreloadValueOn  ZoneSettingEditResponseZonesPrefetchPreloadValue = "on"
	ZoneSettingEditResponseZonesPrefetchPreloadValueOff ZoneSettingEditResponseZonesPrefetchPreloadValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesPrefetchPreloadEditable bool

const (
	ZoneSettingEditResponseZonesPrefetchPreloadEditableTrue  ZoneSettingEditResponseZonesPrefetchPreloadEditable = true
	ZoneSettingEditResponseZonesPrefetchPreloadEditableFalse ZoneSettingEditResponseZonesPrefetchPreloadEditable = false
)

// Maximum time between two read operations from origin.
type ZoneSettingEditResponseZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesProxyReadTimeoutJSON `json:"-"`
}

// zoneSettingEditResponseZonesProxyReadTimeoutJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesProxyReadTimeout]
type zoneSettingEditResponseZonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesProxyReadTimeout) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesProxyReadTimeoutID string

const (
	ZoneSettingEditResponseZonesProxyReadTimeoutIDProxyReadTimeout ZoneSettingEditResponseZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesProxyReadTimeoutEditable bool

const (
	ZoneSettingEditResponseZonesProxyReadTimeoutEditableTrue  ZoneSettingEditResponseZonesProxyReadTimeoutEditable = true
	ZoneSettingEditResponseZonesProxyReadTimeoutEditableFalse ZoneSettingEditResponseZonesProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type ZoneSettingEditResponseZonesPseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID ZoneSettingEditResponseZonesPseudoIPV4ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesPseudoIPV4Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesPseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesPseudoIPV4JSON `json:"-"`
}

// zoneSettingEditResponseZonesPseudoIPV4JSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesPseudoIPV4]
type zoneSettingEditResponseZonesPseudoIPV4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesPseudoIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesPseudoIPV4) implementsZoneSettingEditResponse() {}

// Value of the Pseudo IPv4 setting.
type ZoneSettingEditResponseZonesPseudoIPV4ID string

const (
	ZoneSettingEditResponseZonesPseudoIPV4IDPseudoIPV4 ZoneSettingEditResponseZonesPseudoIPV4ID = "pseudo_ipv4"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesPseudoIPV4Value string

const (
	ZoneSettingEditResponseZonesPseudoIPV4ValueOff             ZoneSettingEditResponseZonesPseudoIPV4Value = "off"
	ZoneSettingEditResponseZonesPseudoIPV4ValueAddHeader       ZoneSettingEditResponseZonesPseudoIPV4Value = "add_header"
	ZoneSettingEditResponseZonesPseudoIPV4ValueOverwriteHeader ZoneSettingEditResponseZonesPseudoIPV4Value = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesPseudoIPV4Editable bool

const (
	ZoneSettingEditResponseZonesPseudoIPV4EditableTrue  ZoneSettingEditResponseZonesPseudoIPV4Editable = true
	ZoneSettingEditResponseZonesPseudoIPV4EditableFalse ZoneSettingEditResponseZonesPseudoIPV4Editable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingEditResponseZonesResponseBuffering struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesResponseBufferingID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesResponseBufferingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesResponseBufferingJSON `json:"-"`
}

// zoneSettingEditResponseZonesResponseBufferingJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesResponseBuffering]
type zoneSettingEditResponseZonesResponseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesResponseBuffering) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesResponseBufferingID string

const (
	ZoneSettingEditResponseZonesResponseBufferingIDResponseBuffering ZoneSettingEditResponseZonesResponseBufferingID = "response_buffering"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesResponseBufferingValue string

const (
	ZoneSettingEditResponseZonesResponseBufferingValueOn  ZoneSettingEditResponseZonesResponseBufferingValue = "on"
	ZoneSettingEditResponseZonesResponseBufferingValueOff ZoneSettingEditResponseZonesResponseBufferingValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesResponseBufferingEditable bool

const (
	ZoneSettingEditResponseZonesResponseBufferingEditableTrue  ZoneSettingEditResponseZonesResponseBufferingEditable = true
	ZoneSettingEditResponseZonesResponseBufferingEditableFalse ZoneSettingEditResponseZonesResponseBufferingEditable = false
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
type ZoneSettingEditResponseZonesRocketLoader struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesRocketLoaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesRocketLoaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesRocketLoaderJSON `json:"-"`
}

// zoneSettingEditResponseZonesRocketLoaderJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesRocketLoader]
type zoneSettingEditResponseZonesRocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesRocketLoader) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesRocketLoaderID string

const (
	ZoneSettingEditResponseZonesRocketLoaderIDRocketLoader ZoneSettingEditResponseZonesRocketLoaderID = "rocket_loader"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesRocketLoaderValue string

const (
	ZoneSettingEditResponseZonesRocketLoaderValueOn  ZoneSettingEditResponseZonesRocketLoaderValue = "on"
	ZoneSettingEditResponseZonesRocketLoaderValueOff ZoneSettingEditResponseZonesRocketLoaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesRocketLoaderEditable bool

const (
	ZoneSettingEditResponseZonesRocketLoaderEditableTrue  ZoneSettingEditResponseZonesRocketLoaderEditable = true
	ZoneSettingEditResponseZonesRocketLoaderEditableFalse ZoneSettingEditResponseZonesRocketLoaderEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationValue `json:"value,required"`
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

// Current value of the zone setting.
type ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationValue struct {
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
	WpPlugin bool                                                                      `json:"wp_plugin,required"`
	JSON     zoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationValueJSON `json:"-"`
}

// zoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationValueJSON
// contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationValue]
type zoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationValueJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Cloudflare security header for a zone.
type ZoneSettingEditResponseZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID ZoneSettingEditResponseZonesSecurityHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesSecurityHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesSecurityHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesSecurityHeaderJSON `json:"-"`
}

// zoneSettingEditResponseZonesSecurityHeaderJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesSecurityHeader]
type zoneSettingEditResponseZonesSecurityHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesSecurityHeader) implementsZoneSettingEditResponse() {}

// ID of the zone's security header.
type ZoneSettingEditResponseZonesSecurityHeaderID string

const (
	ZoneSettingEditResponseZonesSecurityHeaderIDSecurityHeader ZoneSettingEditResponseZonesSecurityHeaderID = "security_header"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZoneSettingEditResponseZonesSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zoneSettingEditResponseZonesSecurityHeaderValueJSON                    `json:"-"`
}

// zoneSettingEditResponseZonesSecurityHeaderValueJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseZonesSecurityHeaderValue]
type zoneSettingEditResponseZonesSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type ZoneSettingEditResponseZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                       `json:"nosniff"`
	JSON    zoneSettingEditResponseZonesSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// zoneSettingEditResponseZonesSecurityHeaderValueStrictTransportSecurityJSON
// contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesSecurityHeaderValueStrictTransportSecurity]
type zoneSettingEditResponseZonesSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesSecurityHeaderEditable bool

const (
	ZoneSettingEditResponseZonesSecurityHeaderEditableTrue  ZoneSettingEditResponseZonesSecurityHeaderEditable = true
	ZoneSettingEditResponseZonesSecurityHeaderEditableFalse ZoneSettingEditResponseZonesSecurityHeaderEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingEditResponseZonesSecurityLevel struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesSecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesSecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesSecurityLevelJSON `json:"-"`
}

// zoneSettingEditResponseZonesSecurityLevelJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesSecurityLevel]
type zoneSettingEditResponseZonesSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesSecurityLevel) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesSecurityLevelID string

const (
	ZoneSettingEditResponseZonesSecurityLevelIDSecurityLevel ZoneSettingEditResponseZonesSecurityLevelID = "security_level"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesSecurityLevelValue string

const (
	ZoneSettingEditResponseZonesSecurityLevelValueOff            ZoneSettingEditResponseZonesSecurityLevelValue = "off"
	ZoneSettingEditResponseZonesSecurityLevelValueEssentiallyOff ZoneSettingEditResponseZonesSecurityLevelValue = "essentially_off"
	ZoneSettingEditResponseZonesSecurityLevelValueLow            ZoneSettingEditResponseZonesSecurityLevelValue = "low"
	ZoneSettingEditResponseZonesSecurityLevelValueMedium         ZoneSettingEditResponseZonesSecurityLevelValue = "medium"
	ZoneSettingEditResponseZonesSecurityLevelValueHigh           ZoneSettingEditResponseZonesSecurityLevelValue = "high"
	ZoneSettingEditResponseZonesSecurityLevelValueUnderAttack    ZoneSettingEditResponseZonesSecurityLevelValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesSecurityLevelEditable bool

const (
	ZoneSettingEditResponseZonesSecurityLevelEditableTrue  ZoneSettingEditResponseZonesSecurityLevelEditable = true
	ZoneSettingEditResponseZonesSecurityLevelEditableFalse ZoneSettingEditResponseZonesSecurityLevelEditable = false
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
type ZoneSettingEditResponseZonesServerSideExclude struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesServerSideExcludeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesServerSideExcludeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesServerSideExcludeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesServerSideExcludeJSON `json:"-"`
}

// zoneSettingEditResponseZonesServerSideExcludeJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesServerSideExclude]
type zoneSettingEditResponseZonesServerSideExcludeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesServerSideExclude) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesServerSideExcludeID string

const (
	ZoneSettingEditResponseZonesServerSideExcludeIDServerSideExclude ZoneSettingEditResponseZonesServerSideExcludeID = "server_side_exclude"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesServerSideExcludeValue string

const (
	ZoneSettingEditResponseZonesServerSideExcludeValueOn  ZoneSettingEditResponseZonesServerSideExcludeValue = "on"
	ZoneSettingEditResponseZonesServerSideExcludeValueOff ZoneSettingEditResponseZonesServerSideExcludeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesServerSideExcludeEditable bool

const (
	ZoneSettingEditResponseZonesServerSideExcludeEditableTrue  ZoneSettingEditResponseZonesServerSideExcludeEditable = true
	ZoneSettingEditResponseZonesServerSideExcludeEditableFalse ZoneSettingEditResponseZonesServerSideExcludeEditable = false
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

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingEditResponseZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesSortQueryStringForCacheID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesSortQueryStringForCacheValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesSortQueryStringForCacheJSON `json:"-"`
}

// zoneSettingEditResponseZonesSortQueryStringForCacheJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseZonesSortQueryStringForCache]
type zoneSettingEditResponseZonesSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesSortQueryStringForCache) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesSortQueryStringForCacheID string

const (
	ZoneSettingEditResponseZonesSortQueryStringForCacheIDSortQueryStringForCache ZoneSettingEditResponseZonesSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesSortQueryStringForCacheValue string

const (
	ZoneSettingEditResponseZonesSortQueryStringForCacheValueOn  ZoneSettingEditResponseZonesSortQueryStringForCacheValue = "on"
	ZoneSettingEditResponseZonesSortQueryStringForCacheValueOff ZoneSettingEditResponseZonesSortQueryStringForCacheValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesSortQueryStringForCacheEditable bool

const (
	ZoneSettingEditResponseZonesSortQueryStringForCacheEditableTrue  ZoneSettingEditResponseZonesSortQueryStringForCacheEditable = true
	ZoneSettingEditResponseZonesSortQueryStringForCacheEditableFalse ZoneSettingEditResponseZonesSortQueryStringForCacheEditable = false
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
type ZoneSettingEditResponseZonesSSL struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesSSLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesSSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesSSLJSON `json:"-"`
}

// zoneSettingEditResponseZonesSSLJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesSSL]
type zoneSettingEditResponseZonesSSLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesSSL) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesSSLID string

const (
	ZoneSettingEditResponseZonesSSLIDSSL ZoneSettingEditResponseZonesSSLID = "ssl"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesSSLValue string

const (
	ZoneSettingEditResponseZonesSSLValueOff      ZoneSettingEditResponseZonesSSLValue = "off"
	ZoneSettingEditResponseZonesSSLValueFlexible ZoneSettingEditResponseZonesSSLValue = "flexible"
	ZoneSettingEditResponseZonesSSLValueFull     ZoneSettingEditResponseZonesSSLValue = "full"
	ZoneSettingEditResponseZonesSSLValueStrict   ZoneSettingEditResponseZonesSSLValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesSSLEditable bool

const (
	ZoneSettingEditResponseZonesSSLEditableTrue  ZoneSettingEditResponseZonesSSLEditable = true
	ZoneSettingEditResponseZonesSSLEditableFalse ZoneSettingEditResponseZonesSSLEditable = false
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingEditResponseZonesSSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingEditResponseZonesSSLRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                           `json:"enabled"`
	JSON    zoneSettingEditResponseZonesSSLRecommenderJSON `json:"-"`
}

// zoneSettingEditResponseZonesSSLRecommenderJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesSSLRecommender]
type zoneSettingEditResponseZonesSSLRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesSSLRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesSSLRecommender) implementsZoneSettingEditResponse() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingEditResponseZonesSSLRecommenderID string

const (
	ZoneSettingEditResponseZonesSSLRecommenderIDSSLRecommender ZoneSettingEditResponseZonesSSLRecommenderID = "ssl_recommender"
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

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingEditResponseZonesTLS1_3 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesTLS1_3ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesTLS1_3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesTLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesTls1_3JSON `json:"-"`
}

// zoneSettingEditResponseZonesTls1_3JSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesTLS1_3]
type zoneSettingEditResponseZonesTls1_3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesTLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesTLS1_3) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesTLS1_3ID string

const (
	ZoneSettingEditResponseZonesTLS1_3IDTLS1_3 ZoneSettingEditResponseZonesTLS1_3ID = "tls_1_3"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesTLS1_3Value string

const (
	ZoneSettingEditResponseZonesTLS1_3ValueOn  ZoneSettingEditResponseZonesTLS1_3Value = "on"
	ZoneSettingEditResponseZonesTLS1_3ValueOff ZoneSettingEditResponseZonesTLS1_3Value = "off"
	ZoneSettingEditResponseZonesTLS1_3ValueZrt ZoneSettingEditResponseZonesTLS1_3Value = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesTLS1_3Editable bool

const (
	ZoneSettingEditResponseZonesTLS1_3EditableTrue  ZoneSettingEditResponseZonesTLS1_3Editable = true
	ZoneSettingEditResponseZonesTLS1_3EditableFalse ZoneSettingEditResponseZonesTLS1_3Editable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingEditResponseZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesTLSClientAuthID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesTLSClientAuthValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesTLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesTLSClientAuthJSON `json:"-"`
}

// zoneSettingEditResponseZonesTLSClientAuthJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesTLSClientAuth]
type zoneSettingEditResponseZonesTLSClientAuthJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesTLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesTLSClientAuth) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesTLSClientAuthID string

const (
	ZoneSettingEditResponseZonesTLSClientAuthIDTLSClientAuth ZoneSettingEditResponseZonesTLSClientAuthID = "tls_client_auth"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesTLSClientAuthValue string

const (
	ZoneSettingEditResponseZonesTLSClientAuthValueOn  ZoneSettingEditResponseZonesTLSClientAuthValue = "on"
	ZoneSettingEditResponseZonesTLSClientAuthValueOff ZoneSettingEditResponseZonesTLSClientAuthValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesTLSClientAuthEditable bool

const (
	ZoneSettingEditResponseZonesTLSClientAuthEditableTrue  ZoneSettingEditResponseZonesTLSClientAuthEditable = true
	ZoneSettingEditResponseZonesTLSClientAuthEditableFalse ZoneSettingEditResponseZonesTLSClientAuthEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingEditResponseZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesTrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesTrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesTrueClientIPHeaderJSON `json:"-"`
}

// zoneSettingEditResponseZonesTrueClientIPHeaderJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseZonesTrueClientIPHeader]
type zoneSettingEditResponseZonesTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesTrueClientIPHeader) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesTrueClientIPHeaderID string

const (
	ZoneSettingEditResponseZonesTrueClientIPHeaderIDTrueClientIPHeader ZoneSettingEditResponseZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesTrueClientIPHeaderValue string

const (
	ZoneSettingEditResponseZonesTrueClientIPHeaderValueOn  ZoneSettingEditResponseZonesTrueClientIPHeaderValue = "on"
	ZoneSettingEditResponseZonesTrueClientIPHeaderValueOff ZoneSettingEditResponseZonesTrueClientIPHeaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesTrueClientIPHeaderEditable bool

const (
	ZoneSettingEditResponseZonesTrueClientIPHeaderEditableTrue  ZoneSettingEditResponseZonesTrueClientIPHeaderEditable = true
	ZoneSettingEditResponseZonesTrueClientIPHeaderEditableFalse ZoneSettingEditResponseZonesTrueClientIPHeaderEditable = false
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
type ZoneSettingEditResponseZonesWAF struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesWAFID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesWAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesWAFJSON `json:"-"`
}

// zoneSettingEditResponseZonesWAFJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesWAF]
type zoneSettingEditResponseZonesWAFJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesWAF) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesWAFID string

const (
	ZoneSettingEditResponseZonesWAFIDWAF ZoneSettingEditResponseZonesWAFID = "waf"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesWAFValue string

const (
	ZoneSettingEditResponseZonesWAFValueOn  ZoneSettingEditResponseZonesWAFValue = "on"
	ZoneSettingEditResponseZonesWAFValueOff ZoneSettingEditResponseZonesWAFValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesWAFEditable bool

const (
	ZoneSettingEditResponseZonesWAFEditableTrue  ZoneSettingEditResponseZonesWAFEditable = true
	ZoneSettingEditResponseZonesWAFEditableFalse ZoneSettingEditResponseZonesWAFEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingEditResponseZonesWebp struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesWebpID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesWebpValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesWebpJSON `json:"-"`
}

// zoneSettingEditResponseZonesWebpJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseZonesWebp]
type zoneSettingEditResponseZonesWebpJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesWebp) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesWebpID string

const (
	ZoneSettingEditResponseZonesWebpIDWebp ZoneSettingEditResponseZonesWebpID = "webp"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesWebpValue string

const (
	ZoneSettingEditResponseZonesWebpValueOff ZoneSettingEditResponseZonesWebpValue = "off"
	ZoneSettingEditResponseZonesWebpValueOn  ZoneSettingEditResponseZonesWebpValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesWebpEditable bool

const (
	ZoneSettingEditResponseZonesWebpEditableTrue  ZoneSettingEditResponseZonesWebpEditable = true
	ZoneSettingEditResponseZonesWebpEditableFalse ZoneSettingEditResponseZonesWebpEditable = false
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZoneSettingEditResponseZonesWebsockets struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseZonesWebsocketsID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesWebsocketsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesWebsocketsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesWebsocketsJSON `json:"-"`
}

// zoneSettingEditResponseZonesWebsocketsJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseZonesWebsockets]
type zoneSettingEditResponseZonesWebsocketsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseZonesWebsockets) implementsZoneSettingEditResponse() {}

// ID of the zone setting.
type ZoneSettingEditResponseZonesWebsocketsID string

const (
	ZoneSettingEditResponseZonesWebsocketsIDWebsockets ZoneSettingEditResponseZonesWebsocketsID = "websockets"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesWebsocketsValue string

const (
	ZoneSettingEditResponseZonesWebsocketsValueOff ZoneSettingEditResponseZonesWebsocketsValue = "off"
	ZoneSettingEditResponseZonesWebsocketsValueOn  ZoneSettingEditResponseZonesWebsocketsValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesWebsocketsEditable bool

const (
	ZoneSettingEditResponseZonesWebsocketsEditableTrue  ZoneSettingEditResponseZonesWebsocketsEditable = true
	ZoneSettingEditResponseZonesWebsocketsEditableFalse ZoneSettingEditResponseZonesWebsocketsEditable = false
)

type ZoneSettingListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneSettingListResponseEnvelope struct {
	Errors   []ZoneSettingListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneSettingListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                                `json:"success,required"`
	Result  []ZoneSettingListResponse           `json:"result"`
	JSON    zoneSettingListResponseEnvelopeJSON `json:"-"`
}

// zoneSettingListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseEnvelope]
type zoneSettingListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSettingListResponseEnvelopeErrorsJSON `json:"-"`
}

// zoneSettingListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseEnvelopeErrors]
type zoneSettingListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneSettingListResponseEnvelopeMessagesJSON `json:"-"`
}

// zoneSettingListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseEnvelopeMessages]
type zoneSettingListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
// Satisfied by [ZoneSettingEditParamsItemsZones0rtt],
// [ZoneSettingEditParamsItemsZonesAdvancedDDOS],
// [ZoneSettingEditParamsItemsZonesAlwaysOnline],
// [ZoneSettingEditParamsItemsZonesAlwaysUseHTTPS],
// [ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewrites],
// [ZoneSettingEditParamsItemsZonesBrotli],
// [ZoneSettingEditParamsItemsZonesBrowserCacheTTL],
// [ZoneSettingEditParamsItemsZonesBrowserCheck],
// [ZoneSettingEditParamsItemsZonesCacheLevel],
// [ZoneSettingEditParamsItemsZonesChallengeTTL],
// [ZoneSettingEditParamsItemsZonesCiphers],
// [ZoneSettingEditParamsItemsZonesCnameFlattening],
// [ZoneSettingEditParamsItemsZonesDevelopmentMode],
// [ZoneSettingEditParamsItemsZonesEarlyHints],
// [ZoneSettingEditParamsItemsZonesEdgeCacheTTL],
// [ZoneSettingEditParamsItemsZonesEmailObfuscation],
// [ZoneSettingEditParamsItemsZonesH2Prioritization],
// [ZoneSettingEditParamsItemsZonesHotlinkProtection],
// [ZoneSettingEditParamsItemsZonesHTTP2], [ZoneSettingEditParamsItemsZonesHTTP3],
// [ZoneSettingEditParamsItemsZonesImageResizing],
// [ZoneSettingEditParamsItemsZonesIPGeolocation],
// [ZoneSettingEditParamsItemsZonesIPV6],
// [ZoneSettingEditParamsItemsZonesMaxUpload],
// [ZoneSettingEditParamsItemsZonesMinTLSVersion],
// [ZoneSettingEditParamsItemsZonesMinify],
// [ZoneSettingEditParamsItemsZonesMirage],
// [ZoneSettingEditParamsItemsZonesMobileRedirect],
// [ZoneSettingEditParamsItemsZonesNEL],
// [ZoneSettingEditParamsItemsZonesOpportunisticEncryption],
// [ZoneSettingEditParamsItemsZonesOpportunisticOnion],
// [ZoneSettingEditParamsItemsZonesOrangeToOrange],
// [ZoneSettingEditParamsItemsZonesOriginErrorPagePassThru],
// [ZoneSettingEditParamsItemsZonesPolish],
// [ZoneSettingEditParamsItemsZonesPrefetchPreload],
// [ZoneSettingEditParamsItemsZonesProxyReadTimeout],
// [ZoneSettingEditParamsItemsZonesPseudoIPV4],
// [ZoneSettingEditParamsItemsZonesResponseBuffering],
// [ZoneSettingEditParamsItemsZonesRocketLoader],
// [ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization],
// [ZoneSettingEditParamsItemsZonesSecurityHeader],
// [ZoneSettingEditParamsItemsZonesSecurityLevel],
// [ZoneSettingEditParamsItemsZonesServerSideExclude],
// [ZoneSettingEditParamsItemsZonesSha1Support],
// [ZoneSettingEditParamsItemsZonesSortQueryStringForCache],
// [ZoneSettingEditParamsItemsZonesSSL],
// [ZoneSettingEditParamsItemsZonesSSLRecommender],
// [ZoneSettingEditParamsItemsZonesTLS1_2Only],
// [ZoneSettingEditParamsItemsZonesTLS1_3],
// [ZoneSettingEditParamsItemsZonesTLSClientAuth],
// [ZoneSettingEditParamsItemsZonesTrueClientIPHeader],
// [ZoneSettingEditParamsItemsZonesWAF], [ZoneSettingEditParamsItemsZonesWebp],
// [ZoneSettingEditParamsItemsZonesWebsockets].
type ZoneSettingEditParamsItem interface {
	implementsZoneSettingEditParamsItem()
}

// 0-RTT session resumption enabled for this zone.
type ZoneSettingEditParamsItemsZones0rtt struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZones0rttID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZones0rttValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZones0rtt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZones0rtt) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZones0rttID string

const (
	ZoneSettingEditParamsItemsZones0rttID0rtt ZoneSettingEditParamsItemsZones0rttID = "0rtt"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZones0rttValue string

const (
	ZoneSettingEditParamsItemsZones0rttValueOn  ZoneSettingEditParamsItemsZones0rttValue = "on"
	ZoneSettingEditParamsItemsZones0rttValueOff ZoneSettingEditParamsItemsZones0rttValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZones0rttEditable bool

const (
	ZoneSettingEditParamsItemsZones0rttEditableTrue  ZoneSettingEditParamsItemsZones0rttEditable = true
	ZoneSettingEditParamsItemsZones0rttEditableFalse ZoneSettingEditParamsItemsZones0rttEditable = false
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingEditParamsItemsZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesAdvancedDDOSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesAdvancedDDOSValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesAdvancedDDOS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesAdvancedDDOS) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesAdvancedDDOSID string

const (
	ZoneSettingEditParamsItemsZonesAdvancedDDOSIDAdvancedDDOS ZoneSettingEditParamsItemsZonesAdvancedDDOSID = "advanced_ddos"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesAdvancedDDOSValue string

const (
	ZoneSettingEditParamsItemsZonesAdvancedDDOSValueOn  ZoneSettingEditParamsItemsZonesAdvancedDDOSValue = "on"
	ZoneSettingEditParamsItemsZonesAdvancedDDOSValueOff ZoneSettingEditParamsItemsZonesAdvancedDDOSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesAdvancedDDOSEditable bool

const (
	ZoneSettingEditParamsItemsZonesAdvancedDDOSEditableTrue  ZoneSettingEditParamsItemsZonesAdvancedDDOSEditable = true
	ZoneSettingEditParamsItemsZonesAdvancedDDOSEditableFalse ZoneSettingEditParamsItemsZonesAdvancedDDOSEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingEditParamsItemsZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesAlwaysOnlineID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesAlwaysOnlineValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesAlwaysOnline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesAlwaysOnline) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesAlwaysOnlineID string

const (
	ZoneSettingEditParamsItemsZonesAlwaysOnlineIDAlwaysOnline ZoneSettingEditParamsItemsZonesAlwaysOnlineID = "always_online"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesAlwaysOnlineValue string

const (
	ZoneSettingEditParamsItemsZonesAlwaysOnlineValueOn  ZoneSettingEditParamsItemsZonesAlwaysOnlineValue = "on"
	ZoneSettingEditParamsItemsZonesAlwaysOnlineValueOff ZoneSettingEditParamsItemsZonesAlwaysOnlineValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesAlwaysOnlineEditable bool

const (
	ZoneSettingEditParamsItemsZonesAlwaysOnlineEditableTrue  ZoneSettingEditParamsItemsZonesAlwaysOnlineEditable = true
	ZoneSettingEditParamsItemsZonesAlwaysOnlineEditableFalse ZoneSettingEditParamsItemsZonesAlwaysOnlineEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingEditParamsItemsZonesAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesAlwaysUseHTTPS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesAlwaysUseHTTPS) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSID string

const (
	ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSIDAlwaysUseHTTPS ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSID = "always_use_https"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSValue string

const (
	ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSValueOn  ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSValue = "on"
	ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSValueOff ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSEditable bool

const (
	ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSEditableTrue  ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSEditable = true
	ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSEditableFalse ZoneSettingEditParamsItemsZonesAlwaysUseHTTPSEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewrites) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewrites) implementsZoneSettingEditParamsItem() {
}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesID string

const (
	ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesValue string

const (
	ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesValueOn  ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesValue = "on"
	ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesValueOff ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesEditable bool

const (
	ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesEditableTrue  ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesEditable = true
	ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesEditableFalse ZoneSettingEditParamsItemsZonesAutomaticHTTPSRewritesEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingEditParamsItemsZonesBrotli struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesBrotliID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesBrotliValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesBrotli) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesBrotli) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesBrotliID string

const (
	ZoneSettingEditParamsItemsZonesBrotliIDBrotli ZoneSettingEditParamsItemsZonesBrotliID = "brotli"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesBrotliValue string

const (
	ZoneSettingEditParamsItemsZonesBrotliValueOff ZoneSettingEditParamsItemsZonesBrotliValue = "off"
	ZoneSettingEditParamsItemsZonesBrotliValueOn  ZoneSettingEditParamsItemsZonesBrotliValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesBrotliEditable bool

const (
	ZoneSettingEditParamsItemsZonesBrotliEditableTrue  ZoneSettingEditParamsItemsZonesBrotliEditable = true
	ZoneSettingEditParamsItemsZonesBrotliEditableFalse ZoneSettingEditParamsItemsZonesBrotliEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingEditParamsItemsZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesBrowserCacheTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesBrowserCacheTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesBrowserCacheTTL) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesBrowserCacheTTLID string

const (
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLIDBrowserCacheTTL ZoneSettingEditParamsItemsZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue float64

const (
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue0        ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 0
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue30       ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 30
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue60       ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 60
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue120      ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 120
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue300      ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 300
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue1200     ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 1200
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue1800     ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 1800
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue3600     ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 3600
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue7200     ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 7200
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue10800    ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 10800
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue14400    ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 14400
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue18000    ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 18000
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue28800    ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 28800
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue43200    ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 43200
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue57600    ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 57600
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue72000    ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 72000
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue86400    ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 86400
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue172800   ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 172800
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue259200   ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 259200
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue345600   ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 345600
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue432000   ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 432000
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue691200   ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 691200
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue1382400  ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 1382400
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue2073600  ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 2073600
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue2678400  ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 2678400
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue5356800  ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 5356800
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue16070400 ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 16070400
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue31536000 ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesBrowserCacheTTLEditable bool

const (
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLEditableTrue  ZoneSettingEditParamsItemsZonesBrowserCacheTTLEditable = true
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLEditableFalse ZoneSettingEditParamsItemsZonesBrowserCacheTTLEditable = false
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingEditParamsItemsZonesBrowserCheck struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesBrowserCheckID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesBrowserCheckValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesBrowserCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesBrowserCheck) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesBrowserCheckID string

const (
	ZoneSettingEditParamsItemsZonesBrowserCheckIDBrowserCheck ZoneSettingEditParamsItemsZonesBrowserCheckID = "browser_check"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesBrowserCheckValue string

const (
	ZoneSettingEditParamsItemsZonesBrowserCheckValueOn  ZoneSettingEditParamsItemsZonesBrowserCheckValue = "on"
	ZoneSettingEditParamsItemsZonesBrowserCheckValueOff ZoneSettingEditParamsItemsZonesBrowserCheckValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesBrowserCheckEditable bool

const (
	ZoneSettingEditParamsItemsZonesBrowserCheckEditableTrue  ZoneSettingEditParamsItemsZonesBrowserCheckEditable = true
	ZoneSettingEditParamsItemsZonesBrowserCheckEditableFalse ZoneSettingEditParamsItemsZonesBrowserCheckEditable = false
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingEditParamsItemsZonesCacheLevel struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesCacheLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesCacheLevelValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesCacheLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesCacheLevel) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesCacheLevelID string

const (
	ZoneSettingEditParamsItemsZonesCacheLevelIDCacheLevel ZoneSettingEditParamsItemsZonesCacheLevelID = "cache_level"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesCacheLevelValue string

const (
	ZoneSettingEditParamsItemsZonesCacheLevelValueAggressive ZoneSettingEditParamsItemsZonesCacheLevelValue = "aggressive"
	ZoneSettingEditParamsItemsZonesCacheLevelValueBasic      ZoneSettingEditParamsItemsZonesCacheLevelValue = "basic"
	ZoneSettingEditParamsItemsZonesCacheLevelValueSimplified ZoneSettingEditParamsItemsZonesCacheLevelValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesCacheLevelEditable bool

const (
	ZoneSettingEditParamsItemsZonesCacheLevelEditableTrue  ZoneSettingEditParamsItemsZonesCacheLevelEditable = true
	ZoneSettingEditParamsItemsZonesCacheLevelEditableFalse ZoneSettingEditParamsItemsZonesCacheLevelEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingEditParamsItemsZonesChallengeTTL struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesChallengeTTLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesChallengeTTLValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesChallengeTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesChallengeTTL) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesChallengeTTLID string

const (
	ZoneSettingEditParamsItemsZonesChallengeTTLIDChallengeTTL ZoneSettingEditParamsItemsZonesChallengeTTLID = "challenge_ttl"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesChallengeTTLValue float64

const (
	ZoneSettingEditParamsItemsZonesChallengeTTLValue300      ZoneSettingEditParamsItemsZonesChallengeTTLValue = 300
	ZoneSettingEditParamsItemsZonesChallengeTTLValue900      ZoneSettingEditParamsItemsZonesChallengeTTLValue = 900
	ZoneSettingEditParamsItemsZonesChallengeTTLValue1800     ZoneSettingEditParamsItemsZonesChallengeTTLValue = 1800
	ZoneSettingEditParamsItemsZonesChallengeTTLValue2700     ZoneSettingEditParamsItemsZonesChallengeTTLValue = 2700
	ZoneSettingEditParamsItemsZonesChallengeTTLValue3600     ZoneSettingEditParamsItemsZonesChallengeTTLValue = 3600
	ZoneSettingEditParamsItemsZonesChallengeTTLValue7200     ZoneSettingEditParamsItemsZonesChallengeTTLValue = 7200
	ZoneSettingEditParamsItemsZonesChallengeTTLValue10800    ZoneSettingEditParamsItemsZonesChallengeTTLValue = 10800
	ZoneSettingEditParamsItemsZonesChallengeTTLValue14400    ZoneSettingEditParamsItemsZonesChallengeTTLValue = 14400
	ZoneSettingEditParamsItemsZonesChallengeTTLValue28800    ZoneSettingEditParamsItemsZonesChallengeTTLValue = 28800
	ZoneSettingEditParamsItemsZonesChallengeTTLValue57600    ZoneSettingEditParamsItemsZonesChallengeTTLValue = 57600
	ZoneSettingEditParamsItemsZonesChallengeTTLValue86400    ZoneSettingEditParamsItemsZonesChallengeTTLValue = 86400
	ZoneSettingEditParamsItemsZonesChallengeTTLValue604800   ZoneSettingEditParamsItemsZonesChallengeTTLValue = 604800
	ZoneSettingEditParamsItemsZonesChallengeTTLValue2592000  ZoneSettingEditParamsItemsZonesChallengeTTLValue = 2592000
	ZoneSettingEditParamsItemsZonesChallengeTTLValue31536000 ZoneSettingEditParamsItemsZonesChallengeTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesChallengeTTLEditable bool

const (
	ZoneSettingEditParamsItemsZonesChallengeTTLEditableTrue  ZoneSettingEditParamsItemsZonesChallengeTTLEditable = true
	ZoneSettingEditParamsItemsZonesChallengeTTLEditableFalse ZoneSettingEditParamsItemsZonesChallengeTTLEditable = false
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingEditParamsItemsZonesCiphers struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesCiphersID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[[]string] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesCiphers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesCiphers) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesCiphersID string

const (
	ZoneSettingEditParamsItemsZonesCiphersIDCiphers ZoneSettingEditParamsItemsZonesCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesCiphersEditable bool

const (
	ZoneSettingEditParamsItemsZonesCiphersEditableTrue  ZoneSettingEditParamsItemsZonesCiphersEditable = true
	ZoneSettingEditParamsItemsZonesCiphersEditableFalse ZoneSettingEditParamsItemsZonesCiphersEditable = false
)

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

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingEditParamsItemsZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesDevelopmentModeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesDevelopmentModeValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesDevelopmentMode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesDevelopmentMode) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesDevelopmentModeID string

const (
	ZoneSettingEditParamsItemsZonesDevelopmentModeIDDevelopmentMode ZoneSettingEditParamsItemsZonesDevelopmentModeID = "development_mode"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesDevelopmentModeValue string

const (
	ZoneSettingEditParamsItemsZonesDevelopmentModeValueOn  ZoneSettingEditParamsItemsZonesDevelopmentModeValue = "on"
	ZoneSettingEditParamsItemsZonesDevelopmentModeValueOff ZoneSettingEditParamsItemsZonesDevelopmentModeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesDevelopmentModeEditable bool

const (
	ZoneSettingEditParamsItemsZonesDevelopmentModeEditableTrue  ZoneSettingEditParamsItemsZonesDevelopmentModeEditable = true
	ZoneSettingEditParamsItemsZonesDevelopmentModeEditableFalse ZoneSettingEditParamsItemsZonesDevelopmentModeEditable = false
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingEditParamsItemsZonesEarlyHints struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesEarlyHintsID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesEarlyHintsValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesEarlyHints) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesEarlyHints) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesEarlyHintsID string

const (
	ZoneSettingEditParamsItemsZonesEarlyHintsIDEarlyHints ZoneSettingEditParamsItemsZonesEarlyHintsID = "early_hints"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesEarlyHintsValue string

const (
	ZoneSettingEditParamsItemsZonesEarlyHintsValueOn  ZoneSettingEditParamsItemsZonesEarlyHintsValue = "on"
	ZoneSettingEditParamsItemsZonesEarlyHintsValueOff ZoneSettingEditParamsItemsZonesEarlyHintsValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesEarlyHintsEditable bool

const (
	ZoneSettingEditParamsItemsZonesEarlyHintsEditableTrue  ZoneSettingEditParamsItemsZonesEarlyHintsEditable = true
	ZoneSettingEditParamsItemsZonesEarlyHintsEditableFalse ZoneSettingEditParamsItemsZonesEarlyHintsEditable = false
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

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingEditParamsItemsZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesEmailObfuscationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesEmailObfuscationValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesEmailObfuscation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesEmailObfuscation) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesEmailObfuscationID string

const (
	ZoneSettingEditParamsItemsZonesEmailObfuscationIDEmailObfuscation ZoneSettingEditParamsItemsZonesEmailObfuscationID = "email_obfuscation"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesEmailObfuscationValue string

const (
	ZoneSettingEditParamsItemsZonesEmailObfuscationValueOn  ZoneSettingEditParamsItemsZonesEmailObfuscationValue = "on"
	ZoneSettingEditParamsItemsZonesEmailObfuscationValueOff ZoneSettingEditParamsItemsZonesEmailObfuscationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesEmailObfuscationEditable bool

const (
	ZoneSettingEditParamsItemsZonesEmailObfuscationEditableTrue  ZoneSettingEditParamsItemsZonesEmailObfuscationEditable = true
	ZoneSettingEditParamsItemsZonesEmailObfuscationEditableFalse ZoneSettingEditParamsItemsZonesEmailObfuscationEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingEditParamsItemsZonesH2Prioritization struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesH2PrioritizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesH2PrioritizationValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesH2Prioritization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesH2Prioritization) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesH2PrioritizationID string

const (
	ZoneSettingEditParamsItemsZonesH2PrioritizationIDH2Prioritization ZoneSettingEditParamsItemsZonesH2PrioritizationID = "h2_prioritization"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesH2PrioritizationValue string

const (
	ZoneSettingEditParamsItemsZonesH2PrioritizationValueOn     ZoneSettingEditParamsItemsZonesH2PrioritizationValue = "on"
	ZoneSettingEditParamsItemsZonesH2PrioritizationValueOff    ZoneSettingEditParamsItemsZonesH2PrioritizationValue = "off"
	ZoneSettingEditParamsItemsZonesH2PrioritizationValueCustom ZoneSettingEditParamsItemsZonesH2PrioritizationValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesH2PrioritizationEditable bool

const (
	ZoneSettingEditParamsItemsZonesH2PrioritizationEditableTrue  ZoneSettingEditParamsItemsZonesH2PrioritizationEditable = true
	ZoneSettingEditParamsItemsZonesH2PrioritizationEditableFalse ZoneSettingEditParamsItemsZonesH2PrioritizationEditable = false
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZoneSettingEditParamsItemsZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesHotlinkProtectionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesHotlinkProtectionValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesHotlinkProtection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesHotlinkProtection) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesHotlinkProtectionID string

const (
	ZoneSettingEditParamsItemsZonesHotlinkProtectionIDHotlinkProtection ZoneSettingEditParamsItemsZonesHotlinkProtectionID = "hotlink_protection"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesHotlinkProtectionValue string

const (
	ZoneSettingEditParamsItemsZonesHotlinkProtectionValueOn  ZoneSettingEditParamsItemsZonesHotlinkProtectionValue = "on"
	ZoneSettingEditParamsItemsZonesHotlinkProtectionValueOff ZoneSettingEditParamsItemsZonesHotlinkProtectionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesHotlinkProtectionEditable bool

const (
	ZoneSettingEditParamsItemsZonesHotlinkProtectionEditableTrue  ZoneSettingEditParamsItemsZonesHotlinkProtectionEditable = true
	ZoneSettingEditParamsItemsZonesHotlinkProtectionEditableFalse ZoneSettingEditParamsItemsZonesHotlinkProtectionEditable = false
)

// HTTP2 enabled for this zone.
type ZoneSettingEditParamsItemsZonesHTTP2 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesHTTP2ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesHTTP2Value] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesHTTP2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesHTTP2) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesHTTP2ID string

const (
	ZoneSettingEditParamsItemsZonesHTTP2IDHTTP2 ZoneSettingEditParamsItemsZonesHTTP2ID = "http2"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesHTTP2Value string

const (
	ZoneSettingEditParamsItemsZonesHTTP2ValueOn  ZoneSettingEditParamsItemsZonesHTTP2Value = "on"
	ZoneSettingEditParamsItemsZonesHTTP2ValueOff ZoneSettingEditParamsItemsZonesHTTP2Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesHTTP2Editable bool

const (
	ZoneSettingEditParamsItemsZonesHTTP2EditableTrue  ZoneSettingEditParamsItemsZonesHTTP2Editable = true
	ZoneSettingEditParamsItemsZonesHTTP2EditableFalse ZoneSettingEditParamsItemsZonesHTTP2Editable = false
)

// HTTP3 enabled for this zone.
type ZoneSettingEditParamsItemsZonesHTTP3 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesHTTP3ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesHTTP3Value] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesHTTP3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesHTTP3) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesHTTP3ID string

const (
	ZoneSettingEditParamsItemsZonesHTTP3IDHTTP3 ZoneSettingEditParamsItemsZonesHTTP3ID = "http3"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesHTTP3Value string

const (
	ZoneSettingEditParamsItemsZonesHTTP3ValueOn  ZoneSettingEditParamsItemsZonesHTTP3Value = "on"
	ZoneSettingEditParamsItemsZonesHTTP3ValueOff ZoneSettingEditParamsItemsZonesHTTP3Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesHTTP3Editable bool

const (
	ZoneSettingEditParamsItemsZonesHTTP3EditableTrue  ZoneSettingEditParamsItemsZonesHTTP3Editable = true
	ZoneSettingEditParamsItemsZonesHTTP3EditableFalse ZoneSettingEditParamsItemsZonesHTTP3Editable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingEditParamsItemsZonesImageResizing struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesImageResizingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesImageResizingValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesImageResizing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesImageResizing) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesImageResizingID string

const (
	ZoneSettingEditParamsItemsZonesImageResizingIDImageResizing ZoneSettingEditParamsItemsZonesImageResizingID = "image_resizing"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesImageResizingValue string

const (
	ZoneSettingEditParamsItemsZonesImageResizingValueOn   ZoneSettingEditParamsItemsZonesImageResizingValue = "on"
	ZoneSettingEditParamsItemsZonesImageResizingValueOff  ZoneSettingEditParamsItemsZonesImageResizingValue = "off"
	ZoneSettingEditParamsItemsZonesImageResizingValueOpen ZoneSettingEditParamsItemsZonesImageResizingValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesImageResizingEditable bool

const (
	ZoneSettingEditParamsItemsZonesImageResizingEditableTrue  ZoneSettingEditParamsItemsZonesImageResizingEditable = true
	ZoneSettingEditParamsItemsZonesImageResizingEditableFalse ZoneSettingEditParamsItemsZonesImageResizingEditable = false
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingEditParamsItemsZonesIPGeolocation struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesIPGeolocationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesIPGeolocationValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesIPGeolocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesIPGeolocation) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesIPGeolocationID string

const (
	ZoneSettingEditParamsItemsZonesIPGeolocationIDIPGeolocation ZoneSettingEditParamsItemsZonesIPGeolocationID = "ip_geolocation"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesIPGeolocationValue string

const (
	ZoneSettingEditParamsItemsZonesIPGeolocationValueOn  ZoneSettingEditParamsItemsZonesIPGeolocationValue = "on"
	ZoneSettingEditParamsItemsZonesIPGeolocationValueOff ZoneSettingEditParamsItemsZonesIPGeolocationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesIPGeolocationEditable bool

const (
	ZoneSettingEditParamsItemsZonesIPGeolocationEditableTrue  ZoneSettingEditParamsItemsZonesIPGeolocationEditable = true
	ZoneSettingEditParamsItemsZonesIPGeolocationEditableFalse ZoneSettingEditParamsItemsZonesIPGeolocationEditable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingEditParamsItemsZonesIPV6 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesIPV6ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesIPV6Value] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesIPV6) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesIPV6ID string

const (
	ZoneSettingEditParamsItemsZonesIPV6IDIPV6 ZoneSettingEditParamsItemsZonesIPV6ID = "ipv6"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesIPV6Value string

const (
	ZoneSettingEditParamsItemsZonesIPV6ValueOff ZoneSettingEditParamsItemsZonesIPV6Value = "off"
	ZoneSettingEditParamsItemsZonesIPV6ValueOn  ZoneSettingEditParamsItemsZonesIPV6Value = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesIPV6Editable bool

const (
	ZoneSettingEditParamsItemsZonesIPV6EditableTrue  ZoneSettingEditParamsItemsZonesIPV6Editable = true
	ZoneSettingEditParamsItemsZonesIPV6EditableFalse ZoneSettingEditParamsItemsZonesIPV6Editable = false
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

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingEditParamsItemsZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesMinTLSVersionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesMinTLSVersionValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesMinTLSVersion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesMinTLSVersion) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesMinTLSVersionID string

const (
	ZoneSettingEditParamsItemsZonesMinTLSVersionIDMinTLSVersion ZoneSettingEditParamsItemsZonesMinTLSVersionID = "min_tls_version"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesMinTLSVersionValue string

const (
	ZoneSettingEditParamsItemsZonesMinTLSVersionValue1_0 ZoneSettingEditParamsItemsZonesMinTLSVersionValue = "1.0"
	ZoneSettingEditParamsItemsZonesMinTLSVersionValue1_1 ZoneSettingEditParamsItemsZonesMinTLSVersionValue = "1.1"
	ZoneSettingEditParamsItemsZonesMinTLSVersionValue1_2 ZoneSettingEditParamsItemsZonesMinTLSVersionValue = "1.2"
	ZoneSettingEditParamsItemsZonesMinTLSVersionValue1_3 ZoneSettingEditParamsItemsZonesMinTLSVersionValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMinTLSVersionEditable bool

const (
	ZoneSettingEditParamsItemsZonesMinTLSVersionEditableTrue  ZoneSettingEditParamsItemsZonesMinTLSVersionEditable = true
	ZoneSettingEditParamsItemsZonesMinTLSVersionEditableFalse ZoneSettingEditParamsItemsZonesMinTLSVersionEditable = false
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingEditParamsItemsZonesMinify struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsZonesMinifyID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesMinifyValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesMinify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesMinify) implementsZoneSettingEditParamsItem() {}

// Zone setting identifier.
type ZoneSettingEditParamsItemsZonesMinifyID string

const (
	ZoneSettingEditParamsItemsZonesMinifyIDMinify ZoneSettingEditParamsItemsZonesMinifyID = "minify"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css param.Field[ZoneSettingEditParamsItemsZonesMinifyValueCss] `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML param.Field[ZoneSettingEditParamsItemsZonesMinifyValueHTML] `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js param.Field[ZoneSettingEditParamsItemsZonesMinifyValueJs] `json:"js"`
}

func (r ZoneSettingEditParamsItemsZonesMinifyValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingEditParamsItemsZonesMinifyValueCss string

const (
	ZoneSettingEditParamsItemsZonesMinifyValueCssOn  ZoneSettingEditParamsItemsZonesMinifyValueCss = "on"
	ZoneSettingEditParamsItemsZonesMinifyValueCssOff ZoneSettingEditParamsItemsZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingEditParamsItemsZonesMinifyValueHTML string

const (
	ZoneSettingEditParamsItemsZonesMinifyValueHTMLOn  ZoneSettingEditParamsItemsZonesMinifyValueHTML = "on"
	ZoneSettingEditParamsItemsZonesMinifyValueHTMLOff ZoneSettingEditParamsItemsZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingEditParamsItemsZonesMinifyValueJs string

const (
	ZoneSettingEditParamsItemsZonesMinifyValueJsOn  ZoneSettingEditParamsItemsZonesMinifyValueJs = "on"
	ZoneSettingEditParamsItemsZonesMinifyValueJsOff ZoneSettingEditParamsItemsZonesMinifyValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMinifyEditable bool

const (
	ZoneSettingEditParamsItemsZonesMinifyEditableTrue  ZoneSettingEditParamsItemsZonesMinifyEditable = true
	ZoneSettingEditParamsItemsZonesMinifyEditableFalse ZoneSettingEditParamsItemsZonesMinifyEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingEditParamsItemsZonesMirage struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesMirageID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesMirageValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesMirage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesMirage) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesMirageID string

const (
	ZoneSettingEditParamsItemsZonesMirageIDMirage ZoneSettingEditParamsItemsZonesMirageID = "mirage"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesMirageValue string

const (
	ZoneSettingEditParamsItemsZonesMirageValueOn  ZoneSettingEditParamsItemsZonesMirageValue = "on"
	ZoneSettingEditParamsItemsZonesMirageValueOff ZoneSettingEditParamsItemsZonesMirageValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMirageEditable bool

const (
	ZoneSettingEditParamsItemsZonesMirageEditableTrue  ZoneSettingEditParamsItemsZonesMirageEditable = true
	ZoneSettingEditParamsItemsZonesMirageEditableFalse ZoneSettingEditParamsItemsZonesMirageEditable = false
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingEditParamsItemsZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesMobileRedirectID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesMobileRedirectValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesMobileRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesMobileRedirect) implementsZoneSettingEditParamsItem() {}

// Identifier of the zone setting.
type ZoneSettingEditParamsItemsZonesMobileRedirectID string

const (
	ZoneSettingEditParamsItemsZonesMobileRedirectIDMobileRedirect ZoneSettingEditParamsItemsZonesMobileRedirectID = "mobile_redirect"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[ZoneSettingEditParamsItemsZonesMobileRedirectValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri param.Field[bool] `json:"strip_uri"`
}

func (r ZoneSettingEditParamsItemsZonesMobileRedirectValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingEditParamsItemsZonesMobileRedirectValueStatus string

const (
	ZoneSettingEditParamsItemsZonesMobileRedirectValueStatusOn  ZoneSettingEditParamsItemsZonesMobileRedirectValueStatus = "on"
	ZoneSettingEditParamsItemsZonesMobileRedirectValueStatusOff ZoneSettingEditParamsItemsZonesMobileRedirectValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMobileRedirectEditable bool

const (
	ZoneSettingEditParamsItemsZonesMobileRedirectEditableTrue  ZoneSettingEditParamsItemsZonesMobileRedirectEditable = true
	ZoneSettingEditParamsItemsZonesMobileRedirectEditableFalse ZoneSettingEditParamsItemsZonesMobileRedirectEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingEditParamsItemsZonesNEL struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsZonesNELID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesNELValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesNEL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesNEL) implementsZoneSettingEditParamsItem() {}

// Zone setting identifier.
type ZoneSettingEditParamsItemsZonesNELID string

const (
	ZoneSettingEditParamsItemsZonesNELIDNEL ZoneSettingEditParamsItemsZonesNELID = "nel"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesNELValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingEditParamsItemsZonesNELValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesNELEditable bool

const (
	ZoneSettingEditParamsItemsZonesNELEditableTrue  ZoneSettingEditParamsItemsZonesNELEditable = true
	ZoneSettingEditParamsItemsZonesNELEditableFalse ZoneSettingEditParamsItemsZonesNELEditable = false
)

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingEditParamsItemsZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesOpportunisticEncryptionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesOpportunisticEncryption) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesOpportunisticEncryption) implementsZoneSettingEditParamsItem() {
}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesOpportunisticEncryptionID string

const (
	ZoneSettingEditParamsItemsZonesOpportunisticEncryptionIDOpportunisticEncryption ZoneSettingEditParamsItemsZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValue string

const (
	ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValueOn  ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValue = "on"
	ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValueOff ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesOpportunisticEncryptionEditable bool

const (
	ZoneSettingEditParamsItemsZonesOpportunisticEncryptionEditableTrue  ZoneSettingEditParamsItemsZonesOpportunisticEncryptionEditable = true
	ZoneSettingEditParamsItemsZonesOpportunisticEncryptionEditableFalse ZoneSettingEditParamsItemsZonesOpportunisticEncryptionEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingEditParamsItemsZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesOpportunisticOnionID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesOpportunisticOnionValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesOpportunisticOnion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesOpportunisticOnion) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesOpportunisticOnionID string

const (
	ZoneSettingEditParamsItemsZonesOpportunisticOnionIDOpportunisticOnion ZoneSettingEditParamsItemsZonesOpportunisticOnionID = "opportunistic_onion"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesOpportunisticOnionValue string

const (
	ZoneSettingEditParamsItemsZonesOpportunisticOnionValueOn  ZoneSettingEditParamsItemsZonesOpportunisticOnionValue = "on"
	ZoneSettingEditParamsItemsZonesOpportunisticOnionValueOff ZoneSettingEditParamsItemsZonesOpportunisticOnionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesOpportunisticOnionEditable bool

const (
	ZoneSettingEditParamsItemsZonesOpportunisticOnionEditableTrue  ZoneSettingEditParamsItemsZonesOpportunisticOnionEditable = true
	ZoneSettingEditParamsItemsZonesOpportunisticOnionEditableFalse ZoneSettingEditParamsItemsZonesOpportunisticOnionEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingEditParamsItemsZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesOrangeToOrangeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesOrangeToOrangeValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesOrangeToOrange) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesOrangeToOrange) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesOrangeToOrangeID string

const (
	ZoneSettingEditParamsItemsZonesOrangeToOrangeIDOrangeToOrange ZoneSettingEditParamsItemsZonesOrangeToOrangeID = "orange_to_orange"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesOrangeToOrangeValue string

const (
	ZoneSettingEditParamsItemsZonesOrangeToOrangeValueOn  ZoneSettingEditParamsItemsZonesOrangeToOrangeValue = "on"
	ZoneSettingEditParamsItemsZonesOrangeToOrangeValueOff ZoneSettingEditParamsItemsZonesOrangeToOrangeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesOrangeToOrangeEditable bool

const (
	ZoneSettingEditParamsItemsZonesOrangeToOrangeEditableTrue  ZoneSettingEditParamsItemsZonesOrangeToOrangeEditable = true
	ZoneSettingEditParamsItemsZonesOrangeToOrangeEditableFalse ZoneSettingEditParamsItemsZonesOrangeToOrangeEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesOriginErrorPagePassThru) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesOriginErrorPagePassThru) implementsZoneSettingEditParamsItem() {
}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruID string

const (
	ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruIDOriginErrorPagePassThru ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValue string

const (
	ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValueOn  ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValue = "on"
	ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValueOff ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruEditable bool

const (
	ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruEditableTrue  ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruEditable = true
	ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruEditableFalse ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruEditable = false
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingEditParamsItemsZonesPolish struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesPolishID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesPolishValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesPolish) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesPolish) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesPolishID string

const (
	ZoneSettingEditParamsItemsZonesPolishIDPolish ZoneSettingEditParamsItemsZonesPolishID = "polish"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesPolishValue string

const (
	ZoneSettingEditParamsItemsZonesPolishValueOff      ZoneSettingEditParamsItemsZonesPolishValue = "off"
	ZoneSettingEditParamsItemsZonesPolishValueLossless ZoneSettingEditParamsItemsZonesPolishValue = "lossless"
	ZoneSettingEditParamsItemsZonesPolishValueLossy    ZoneSettingEditParamsItemsZonesPolishValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesPolishEditable bool

const (
	ZoneSettingEditParamsItemsZonesPolishEditableTrue  ZoneSettingEditParamsItemsZonesPolishEditable = true
	ZoneSettingEditParamsItemsZonesPolishEditableFalse ZoneSettingEditParamsItemsZonesPolishEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesPrefetchPreloadID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesPrefetchPreloadValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesPrefetchPreload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesPrefetchPreload) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesPrefetchPreloadID string

const (
	ZoneSettingEditParamsItemsZonesPrefetchPreloadIDPrefetchPreload ZoneSettingEditParamsItemsZonesPrefetchPreloadID = "prefetch_preload"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesPrefetchPreloadValue string

const (
	ZoneSettingEditParamsItemsZonesPrefetchPreloadValueOn  ZoneSettingEditParamsItemsZonesPrefetchPreloadValue = "on"
	ZoneSettingEditParamsItemsZonesPrefetchPreloadValueOff ZoneSettingEditParamsItemsZonesPrefetchPreloadValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesPrefetchPreloadEditable bool

const (
	ZoneSettingEditParamsItemsZonesPrefetchPreloadEditableTrue  ZoneSettingEditParamsItemsZonesPrefetchPreloadEditable = true
	ZoneSettingEditParamsItemsZonesPrefetchPreloadEditableFalse ZoneSettingEditParamsItemsZonesPrefetchPreloadEditable = false
)

// Maximum time between two read operations from origin.
type ZoneSettingEditParamsItemsZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesProxyReadTimeoutID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[float64] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesProxyReadTimeout) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesProxyReadTimeoutID string

const (
	ZoneSettingEditParamsItemsZonesProxyReadTimeoutIDProxyReadTimeout ZoneSettingEditParamsItemsZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesProxyReadTimeoutEditable bool

const (
	ZoneSettingEditParamsItemsZonesProxyReadTimeoutEditableTrue  ZoneSettingEditParamsItemsZonesProxyReadTimeoutEditable = true
	ZoneSettingEditParamsItemsZonesProxyReadTimeoutEditableFalse ZoneSettingEditParamsItemsZonesProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type ZoneSettingEditParamsItemsZonesPseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesPseudoIPV4ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesPseudoIPV4Value] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesPseudoIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesPseudoIPV4) implementsZoneSettingEditParamsItem() {}

// Value of the Pseudo IPv4 setting.
type ZoneSettingEditParamsItemsZonesPseudoIPV4ID string

const (
	ZoneSettingEditParamsItemsZonesPseudoIPV4IDPseudoIPV4 ZoneSettingEditParamsItemsZonesPseudoIPV4ID = "pseudo_ipv4"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesPseudoIPV4Value string

const (
	ZoneSettingEditParamsItemsZonesPseudoIPV4ValueOff             ZoneSettingEditParamsItemsZonesPseudoIPV4Value = "off"
	ZoneSettingEditParamsItemsZonesPseudoIPV4ValueAddHeader       ZoneSettingEditParamsItemsZonesPseudoIPV4Value = "add_header"
	ZoneSettingEditParamsItemsZonesPseudoIPV4ValueOverwriteHeader ZoneSettingEditParamsItemsZonesPseudoIPV4Value = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesPseudoIPV4Editable bool

const (
	ZoneSettingEditParamsItemsZonesPseudoIPV4EditableTrue  ZoneSettingEditParamsItemsZonesPseudoIPV4Editable = true
	ZoneSettingEditParamsItemsZonesPseudoIPV4EditableFalse ZoneSettingEditParamsItemsZonesPseudoIPV4Editable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsZonesResponseBuffering struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesResponseBufferingID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesResponseBufferingValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesResponseBuffering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesResponseBuffering) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesResponseBufferingID string

const (
	ZoneSettingEditParamsItemsZonesResponseBufferingIDResponseBuffering ZoneSettingEditParamsItemsZonesResponseBufferingID = "response_buffering"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesResponseBufferingValue string

const (
	ZoneSettingEditParamsItemsZonesResponseBufferingValueOn  ZoneSettingEditParamsItemsZonesResponseBufferingValue = "on"
	ZoneSettingEditParamsItemsZonesResponseBufferingValueOff ZoneSettingEditParamsItemsZonesResponseBufferingValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesResponseBufferingEditable bool

const (
	ZoneSettingEditParamsItemsZonesResponseBufferingEditableTrue  ZoneSettingEditParamsItemsZonesResponseBufferingEditable = true
	ZoneSettingEditParamsItemsZonesResponseBufferingEditableFalse ZoneSettingEditParamsItemsZonesResponseBufferingEditable = false
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
type ZoneSettingEditParamsItemsZonesRocketLoader struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesRocketLoaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesRocketLoaderValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesRocketLoader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesRocketLoader) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesRocketLoaderID string

const (
	ZoneSettingEditParamsItemsZonesRocketLoaderIDRocketLoader ZoneSettingEditParamsItemsZonesRocketLoaderID = "rocket_loader"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesRocketLoaderValue string

const (
	ZoneSettingEditParamsItemsZonesRocketLoaderValueOn  ZoneSettingEditParamsItemsZonesRocketLoaderValue = "on"
	ZoneSettingEditParamsItemsZonesRocketLoaderValueOff ZoneSettingEditParamsItemsZonesRocketLoaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesRocketLoaderEditable bool

const (
	ZoneSettingEditParamsItemsZonesRocketLoaderEditableTrue  ZoneSettingEditParamsItemsZonesRocketLoaderEditable = true
	ZoneSettingEditParamsItemsZonesRocketLoaderEditableFalse ZoneSettingEditParamsItemsZonesRocketLoaderEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationValue] `json:"value,required"`
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

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationValue struct {
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

func (r ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Cloudflare security header for a zone.
type ZoneSettingEditParamsItemsZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID param.Field[ZoneSettingEditParamsItemsZonesSecurityHeaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesSecurityHeaderValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesSecurityHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesSecurityHeader) implementsZoneSettingEditParamsItem() {}

// ID of the zone's security header.
type ZoneSettingEditParamsItemsZonesSecurityHeaderID string

const (
	ZoneSettingEditParamsItemsZonesSecurityHeaderIDSecurityHeader ZoneSettingEditParamsItemsZonesSecurityHeaderID = "security_header"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[ZoneSettingEditParamsItemsZonesSecurityHeaderValueStrictTransportSecurity] `json:"strict_transport_security"`
}

func (r ZoneSettingEditParamsItemsZonesSecurityHeaderValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strict Transport Security.
type ZoneSettingEditParamsItemsZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge param.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff param.Field[bool] `json:"nosniff"`
}

func (r ZoneSettingEditParamsItemsZonesSecurityHeaderValueStrictTransportSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSecurityHeaderEditable bool

const (
	ZoneSettingEditParamsItemsZonesSecurityHeaderEditableTrue  ZoneSettingEditParamsItemsZonesSecurityHeaderEditable = true
	ZoneSettingEditParamsItemsZonesSecurityHeaderEditableFalse ZoneSettingEditParamsItemsZonesSecurityHeaderEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingEditParamsItemsZonesSecurityLevel struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesSecurityLevelID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesSecurityLevelValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesSecurityLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesSecurityLevel) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesSecurityLevelID string

const (
	ZoneSettingEditParamsItemsZonesSecurityLevelIDSecurityLevel ZoneSettingEditParamsItemsZonesSecurityLevelID = "security_level"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesSecurityLevelValue string

const (
	ZoneSettingEditParamsItemsZonesSecurityLevelValueOff            ZoneSettingEditParamsItemsZonesSecurityLevelValue = "off"
	ZoneSettingEditParamsItemsZonesSecurityLevelValueEssentiallyOff ZoneSettingEditParamsItemsZonesSecurityLevelValue = "essentially_off"
	ZoneSettingEditParamsItemsZonesSecurityLevelValueLow            ZoneSettingEditParamsItemsZonesSecurityLevelValue = "low"
	ZoneSettingEditParamsItemsZonesSecurityLevelValueMedium         ZoneSettingEditParamsItemsZonesSecurityLevelValue = "medium"
	ZoneSettingEditParamsItemsZonesSecurityLevelValueHigh           ZoneSettingEditParamsItemsZonesSecurityLevelValue = "high"
	ZoneSettingEditParamsItemsZonesSecurityLevelValueUnderAttack    ZoneSettingEditParamsItemsZonesSecurityLevelValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSecurityLevelEditable bool

const (
	ZoneSettingEditParamsItemsZonesSecurityLevelEditableTrue  ZoneSettingEditParamsItemsZonesSecurityLevelEditable = true
	ZoneSettingEditParamsItemsZonesSecurityLevelEditableFalse ZoneSettingEditParamsItemsZonesSecurityLevelEditable = false
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
type ZoneSettingEditParamsItemsZonesServerSideExclude struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesServerSideExcludeID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesServerSideExcludeValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesServerSideExclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesServerSideExclude) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesServerSideExcludeID string

const (
	ZoneSettingEditParamsItemsZonesServerSideExcludeIDServerSideExclude ZoneSettingEditParamsItemsZonesServerSideExcludeID = "server_side_exclude"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesServerSideExcludeValue string

const (
	ZoneSettingEditParamsItemsZonesServerSideExcludeValueOn  ZoneSettingEditParamsItemsZonesServerSideExcludeValue = "on"
	ZoneSettingEditParamsItemsZonesServerSideExcludeValueOff ZoneSettingEditParamsItemsZonesServerSideExcludeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesServerSideExcludeEditable bool

const (
	ZoneSettingEditParamsItemsZonesServerSideExcludeEditableTrue  ZoneSettingEditParamsItemsZonesServerSideExcludeEditable = true
	ZoneSettingEditParamsItemsZonesServerSideExcludeEditableFalse ZoneSettingEditParamsItemsZonesServerSideExcludeEditable = false
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

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingEditParamsItemsZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesSortQueryStringForCacheID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesSortQueryStringForCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesSortQueryStringForCache) implementsZoneSettingEditParamsItem() {
}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesSortQueryStringForCacheID string

const (
	ZoneSettingEditParamsItemsZonesSortQueryStringForCacheIDSortQueryStringForCache ZoneSettingEditParamsItemsZonesSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValue string

const (
	ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValueOn  ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValue = "on"
	ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValueOff ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSortQueryStringForCacheEditable bool

const (
	ZoneSettingEditParamsItemsZonesSortQueryStringForCacheEditableTrue  ZoneSettingEditParamsItemsZonesSortQueryStringForCacheEditable = true
	ZoneSettingEditParamsItemsZonesSortQueryStringForCacheEditableFalse ZoneSettingEditParamsItemsZonesSortQueryStringForCacheEditable = false
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
type ZoneSettingEditParamsItemsZonesSSL struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesSSLID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesSSLValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesSSL) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesSSLID string

const (
	ZoneSettingEditParamsItemsZonesSSLIDSSL ZoneSettingEditParamsItemsZonesSSLID = "ssl"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesSSLValue string

const (
	ZoneSettingEditParamsItemsZonesSSLValueOff      ZoneSettingEditParamsItemsZonesSSLValue = "off"
	ZoneSettingEditParamsItemsZonesSSLValueFlexible ZoneSettingEditParamsItemsZonesSSLValue = "flexible"
	ZoneSettingEditParamsItemsZonesSSLValueFull     ZoneSettingEditParamsItemsZonesSSLValue = "full"
	ZoneSettingEditParamsItemsZonesSSLValueStrict   ZoneSettingEditParamsItemsZonesSSLValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSSLEditable bool

const (
	ZoneSettingEditParamsItemsZonesSSLEditableTrue  ZoneSettingEditParamsItemsZonesSSLEditable = true
	ZoneSettingEditParamsItemsZonesSSLEditableFalse ZoneSettingEditParamsItemsZonesSSLEditable = false
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingEditParamsItemsZonesSSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[ZoneSettingEditParamsItemsZonesSSLRecommenderID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingEditParamsItemsZonesSSLRecommender) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesSSLRecommender) implementsZoneSettingEditParamsItem() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingEditParamsItemsZonesSSLRecommenderID string

const (
	ZoneSettingEditParamsItemsZonesSSLRecommenderIDSSLRecommender ZoneSettingEditParamsItemsZonesSSLRecommenderID = "ssl_recommender"
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

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingEditParamsItemsZonesTLS1_3 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesTLS1_3ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesTLS1_3Value] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesTLS1_3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesTLS1_3) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesTLS1_3ID string

const (
	ZoneSettingEditParamsItemsZonesTLS1_3IDTLS1_3 ZoneSettingEditParamsItemsZonesTLS1_3ID = "tls_1_3"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesTLS1_3Value string

const (
	ZoneSettingEditParamsItemsZonesTLS1_3ValueOn  ZoneSettingEditParamsItemsZonesTLS1_3Value = "on"
	ZoneSettingEditParamsItemsZonesTLS1_3ValueOff ZoneSettingEditParamsItemsZonesTLS1_3Value = "off"
	ZoneSettingEditParamsItemsZonesTLS1_3ValueZrt ZoneSettingEditParamsItemsZonesTLS1_3Value = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTLS1_3Editable bool

const (
	ZoneSettingEditParamsItemsZonesTLS1_3EditableTrue  ZoneSettingEditParamsItemsZonesTLS1_3Editable = true
	ZoneSettingEditParamsItemsZonesTLS1_3EditableFalse ZoneSettingEditParamsItemsZonesTLS1_3Editable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingEditParamsItemsZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesTLSClientAuthID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesTLSClientAuthValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesTLSClientAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesTLSClientAuth) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesTLSClientAuthID string

const (
	ZoneSettingEditParamsItemsZonesTLSClientAuthIDTLSClientAuth ZoneSettingEditParamsItemsZonesTLSClientAuthID = "tls_client_auth"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesTLSClientAuthValue string

const (
	ZoneSettingEditParamsItemsZonesTLSClientAuthValueOn  ZoneSettingEditParamsItemsZonesTLSClientAuthValue = "on"
	ZoneSettingEditParamsItemsZonesTLSClientAuthValueOff ZoneSettingEditParamsItemsZonesTLSClientAuthValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTLSClientAuthEditable bool

const (
	ZoneSettingEditParamsItemsZonesTLSClientAuthEditableTrue  ZoneSettingEditParamsItemsZonesTLSClientAuthEditable = true
	ZoneSettingEditParamsItemsZonesTLSClientAuthEditableFalse ZoneSettingEditParamsItemsZonesTLSClientAuthEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesTrueClientIPHeaderID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesTrueClientIPHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesTrueClientIPHeader) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesTrueClientIPHeaderID string

const (
	ZoneSettingEditParamsItemsZonesTrueClientIPHeaderIDTrueClientIPHeader ZoneSettingEditParamsItemsZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValue string

const (
	ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValueOn  ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValue = "on"
	ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValueOff ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTrueClientIPHeaderEditable bool

const (
	ZoneSettingEditParamsItemsZonesTrueClientIPHeaderEditableTrue  ZoneSettingEditParamsItemsZonesTrueClientIPHeaderEditable = true
	ZoneSettingEditParamsItemsZonesTrueClientIPHeaderEditableFalse ZoneSettingEditParamsItemsZonesTrueClientIPHeaderEditable = false
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
type ZoneSettingEditParamsItemsZonesWAF struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesWAFID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesWAFValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesWAF) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesWAF) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesWAFID string

const (
	ZoneSettingEditParamsItemsZonesWAFIDWAF ZoneSettingEditParamsItemsZonesWAFID = "waf"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesWAFValue string

const (
	ZoneSettingEditParamsItemsZonesWAFValueOn  ZoneSettingEditParamsItemsZonesWAFValue = "on"
	ZoneSettingEditParamsItemsZonesWAFValueOff ZoneSettingEditParamsItemsZonesWAFValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesWAFEditable bool

const (
	ZoneSettingEditParamsItemsZonesWAFEditableTrue  ZoneSettingEditParamsItemsZonesWAFEditable = true
	ZoneSettingEditParamsItemsZonesWAFEditableFalse ZoneSettingEditParamsItemsZonesWAFEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingEditParamsItemsZonesWebp struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesWebpID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesWebpValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesWebp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesWebp) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesWebpID string

const (
	ZoneSettingEditParamsItemsZonesWebpIDWebp ZoneSettingEditParamsItemsZonesWebpID = "webp"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesWebpValue string

const (
	ZoneSettingEditParamsItemsZonesWebpValueOff ZoneSettingEditParamsItemsZonesWebpValue = "off"
	ZoneSettingEditParamsItemsZonesWebpValueOn  ZoneSettingEditParamsItemsZonesWebpValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesWebpEditable bool

const (
	ZoneSettingEditParamsItemsZonesWebpEditableTrue  ZoneSettingEditParamsItemsZonesWebpEditable = true
	ZoneSettingEditParamsItemsZonesWebpEditableFalse ZoneSettingEditParamsItemsZonesWebpEditable = false
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZoneSettingEditParamsItemsZonesWebsockets struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesWebsocketsID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesWebsocketsValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesWebsockets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesWebsockets) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesWebsocketsID string

const (
	ZoneSettingEditParamsItemsZonesWebsocketsIDWebsockets ZoneSettingEditParamsItemsZonesWebsocketsID = "websockets"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesWebsocketsValue string

const (
	ZoneSettingEditParamsItemsZonesWebsocketsValueOff ZoneSettingEditParamsItemsZonesWebsocketsValue = "off"
	ZoneSettingEditParamsItemsZonesWebsocketsValueOn  ZoneSettingEditParamsItemsZonesWebsocketsValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesWebsocketsEditable bool

const (
	ZoneSettingEditParamsItemsZonesWebsocketsEditableTrue  ZoneSettingEditParamsItemsZonesWebsocketsEditable = true
	ZoneSettingEditParamsItemsZonesWebsocketsEditableFalse ZoneSettingEditParamsItemsZonesWebsocketsEditable = false
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
