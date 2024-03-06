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
	"github.com/tidwall/gjson"
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
// [ZoneSettingEditResponseZonesCNAMEFlattening],
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
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneSettingEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZones0rtt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesAdvancedDDOS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesAlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesBrotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesBrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesCacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesCiphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesCNAMEFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesDevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesEarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesEmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesH2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesHotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesHTTP2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesHTTP3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesIPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesIPV6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesMinTLSVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesMinify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesMobileRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesNEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesOpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesOpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesOrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesPolish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesPrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesPseudoIPV4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesResponseBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesRocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesSecurityHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesSecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesServerSideExclude{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesSSL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesSSLRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesTLS1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesTLS1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesTLSClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesWAF{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesWebp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEditResponseZonesWebsockets{}),
		},
	)
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

func (r zoneSettingEditResponseZones0rttJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesAdvancedDDOSJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesAlwaysOnlineJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesAlwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesAutomaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesBrotliJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesBrowserCacheTTLJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesBrowserCheckJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesCacheLevelJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesChallengeTTLJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesCiphersJSON) RawJSON() string {
	return r.raw
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
type ZoneSettingEditResponseZonesCNAMEFlattening struct {
	// How to flatten the cname destination.
	ID ZoneSettingEditResponseZonesCNAMEFlatteningID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingEditResponseZonesCNAMEFlatteningValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseZonesCNAMEFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseZonesCNAMEFlatteningJSON `json:"-"`
}

// zoneSettingEditResponseZonesCNAMEFlatteningJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseZonesCNAMEFlattening]
type zoneSettingEditResponseZonesCNAMEFlatteningJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesCNAMEFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingEditResponseZonesCNAMEFlatteningJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingEditResponseZonesCNAMEFlattening) implementsZoneSettingEditResponse() {}

// How to flatten the cname destination.
type ZoneSettingEditResponseZonesCNAMEFlatteningID string

const (
	ZoneSettingEditResponseZonesCNAMEFlatteningIDCNAMEFlattening ZoneSettingEditResponseZonesCNAMEFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type ZoneSettingEditResponseZonesCNAMEFlatteningValue string

const (
	ZoneSettingEditResponseZonesCNAMEFlatteningValueFlattenAtRoot ZoneSettingEditResponseZonesCNAMEFlatteningValue = "flatten_at_root"
	ZoneSettingEditResponseZonesCNAMEFlatteningValueFlattenAll    ZoneSettingEditResponseZonesCNAMEFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseZonesCNAMEFlatteningEditable bool

const (
	ZoneSettingEditResponseZonesCNAMEFlatteningEditableTrue  ZoneSettingEditResponseZonesCNAMEFlatteningEditable = true
	ZoneSettingEditResponseZonesCNAMEFlatteningEditableFalse ZoneSettingEditResponseZonesCNAMEFlatteningEditable = false
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

func (r zoneSettingEditResponseZonesDevelopmentModeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesEarlyHintsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesEdgeCacheTTLJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesEmailObfuscationJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesH2PrioritizationJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesHotlinkProtectionJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesHTTP2JSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesHTTP3JSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesImageResizingJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesIPGeolocationJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesIPV6JSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesMaxUploadJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesMinTLSVersionJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesMinifyJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesMinifyValueJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesMirageJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesMobileRedirectJSON) RawJSON() string {
	return r.raw
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
	StripURI bool                                                `json:"strip_uri"`
	JSON     zoneSettingEditResponseZonesMobileRedirectValueJSON `json:"-"`
}

// zoneSettingEditResponseZonesMobileRedirectValueJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseZonesMobileRedirectValue]
type zoneSettingEditResponseZonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripURI        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingEditResponseZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingEditResponseZonesMobileRedirectValueJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesNELJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesNELValueJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesOpportunisticEncryptionJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesOpportunisticOnionJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesOrangeToOrangeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesOriginErrorPagePassThruJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesPolishJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesPrefetchPreloadJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesProxyReadTimeoutJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesPseudoIPV4JSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesResponseBufferingJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesRocketLoaderJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesSchemasAutomaticPlatformOptimizationValueJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesSecurityHeaderJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesSecurityHeaderValueJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesSecurityHeaderValueStrictTransportSecurityJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesSecurityLevelJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesServerSideExcludeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesSha1SupportJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesSortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesSSLJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesSSLRecommenderJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesTls1_2OnlyJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesTls1_3JSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesTLSClientAuthJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesTrueClientIPHeaderJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesWAFJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesWebpJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseZonesWebsocketsJSON) RawJSON() string {
	return r.raw
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

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [ZoneSettingGetResponseZones0rtt],
// [ZoneSettingGetResponseZonesAdvancedDDOS],
// [ZoneSettingGetResponseZonesAlwaysOnline],
// [ZoneSettingGetResponseZonesAlwaysUseHTTPS],
// [ZoneSettingGetResponseZonesAutomaticHTTPSRewrites],
// [ZoneSettingGetResponseZonesBrotli],
// [ZoneSettingGetResponseZonesBrowserCacheTTL],
// [ZoneSettingGetResponseZonesBrowserCheck],
// [ZoneSettingGetResponseZonesCacheLevel],
// [ZoneSettingGetResponseZonesChallengeTTL], [ZoneSettingGetResponseZonesCiphers],
// [ZoneSettingGetResponseZonesCNAMEFlattening],
// [ZoneSettingGetResponseZonesDevelopmentMode],
// [ZoneSettingGetResponseZonesEarlyHints],
// [ZoneSettingGetResponseZonesEdgeCacheTTL],
// [ZoneSettingGetResponseZonesEmailObfuscation],
// [ZoneSettingGetResponseZonesH2Prioritization],
// [ZoneSettingGetResponseZonesHotlinkProtection],
// [ZoneSettingGetResponseZonesHTTP2], [ZoneSettingGetResponseZonesHTTP3],
// [ZoneSettingGetResponseZonesImageResizing],
// [ZoneSettingGetResponseZonesIPGeolocation], [ZoneSettingGetResponseZonesIPV6],
// [ZoneSettingGetResponseZonesMaxUpload],
// [ZoneSettingGetResponseZonesMinTLSVersion], [ZoneSettingGetResponseZonesMinify],
// [ZoneSettingGetResponseZonesMirage],
// [ZoneSettingGetResponseZonesMobileRedirect], [ZoneSettingGetResponseZonesNEL],
// [ZoneSettingGetResponseZonesOpportunisticEncryption],
// [ZoneSettingGetResponseZonesOpportunisticOnion],
// [ZoneSettingGetResponseZonesOrangeToOrange],
// [ZoneSettingGetResponseZonesOriginErrorPagePassThru],
// [ZoneSettingGetResponseZonesPolish],
// [ZoneSettingGetResponseZonesPrefetchPreload],
// [ZoneSettingGetResponseZonesProxyReadTimeout],
// [ZoneSettingGetResponseZonesPseudoIPV4],
// [ZoneSettingGetResponseZonesResponseBuffering],
// [ZoneSettingGetResponseZonesRocketLoader],
// [ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimization],
// [ZoneSettingGetResponseZonesSecurityHeader],
// [ZoneSettingGetResponseZonesSecurityLevel],
// [ZoneSettingGetResponseZonesServerSideExclude],
// [ZoneSettingGetResponseZonesSha1Support],
// [ZoneSettingGetResponseZonesSortQueryStringForCache],
// [ZoneSettingGetResponseZonesSSL], [ZoneSettingGetResponseZonesSSLRecommender],
// [ZoneSettingGetResponseZonesTLS1_2Only], [ZoneSettingGetResponseZonesTLS1_3],
// [ZoneSettingGetResponseZonesTLSClientAuth],
// [ZoneSettingGetResponseZonesTrueClientIPHeader],
// [ZoneSettingGetResponseZonesWAF], [ZoneSettingGetResponseZonesWebp] or
// [ZoneSettingGetResponseZonesWebsockets].
type ZoneSettingGetResponse interface {
	implementsZoneSettingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneSettingGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZones0rtt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesAdvancedDDOS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesAlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesBrotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesBrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesCacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesCiphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesCNAMEFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesDevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesEarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesEmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesH2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesHotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesHTTP2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesHTTP3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesIPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesIPV6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesMinTLSVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesMinify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesMobileRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesNEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesOpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesOpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesOrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesPolish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesPrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesPseudoIPV4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesResponseBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesRocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesSecurityHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesSecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesServerSideExclude{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesSSL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesSSLRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesTLS1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesTLS1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesTLSClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesWAF{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesWebp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingGetResponseZonesWebsockets{}),
		},
	)
}

// 0-RTT session resumption enabled for this zone.
type ZoneSettingGetResponseZones0rtt struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZones0rttID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZones0rttValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZones0rttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZones0rttJSON `json:"-"`
}

// zoneSettingGetResponseZones0rttJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZones0rtt]
type zoneSettingGetResponseZones0rttJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZones0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZones0rttJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZones0rtt) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZones0rttID string

const (
	ZoneSettingGetResponseZones0rttID0rtt ZoneSettingGetResponseZones0rttID = "0rtt"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZones0rttValue string

const (
	ZoneSettingGetResponseZones0rttValueOn  ZoneSettingGetResponseZones0rttValue = "on"
	ZoneSettingGetResponseZones0rttValueOff ZoneSettingGetResponseZones0rttValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZones0rttEditable bool

const (
	ZoneSettingGetResponseZones0rttEditableTrue  ZoneSettingGetResponseZones0rttEditable = true
	ZoneSettingGetResponseZones0rttEditableFalse ZoneSettingGetResponseZones0rttEditable = false
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingGetResponseZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesAdvancedDDOSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesAdvancedDDOSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesAdvancedDDOSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesAdvancedDDOSJSON `json:"-"`
}

// zoneSettingGetResponseZonesAdvancedDDOSJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesAdvancedDDOS]
type zoneSettingGetResponseZonesAdvancedDDOSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesAdvancedDDOS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesAdvancedDDOSJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesAdvancedDDOS) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesAdvancedDDOSID string

const (
	ZoneSettingGetResponseZonesAdvancedDDOSIDAdvancedDDOS ZoneSettingGetResponseZonesAdvancedDDOSID = "advanced_ddos"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesAdvancedDDOSValue string

const (
	ZoneSettingGetResponseZonesAdvancedDDOSValueOn  ZoneSettingGetResponseZonesAdvancedDDOSValue = "on"
	ZoneSettingGetResponseZonesAdvancedDDOSValueOff ZoneSettingGetResponseZonesAdvancedDDOSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesAdvancedDDOSEditable bool

const (
	ZoneSettingGetResponseZonesAdvancedDDOSEditableTrue  ZoneSettingGetResponseZonesAdvancedDDOSEditable = true
	ZoneSettingGetResponseZonesAdvancedDDOSEditableFalse ZoneSettingGetResponseZonesAdvancedDDOSEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingGetResponseZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesAlwaysOnlineID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesAlwaysOnlineValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesAlwaysOnlineJSON `json:"-"`
}

// zoneSettingGetResponseZonesAlwaysOnlineJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesAlwaysOnline]
type zoneSettingGetResponseZonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesAlwaysOnlineJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesAlwaysOnline) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesAlwaysOnlineID string

const (
	ZoneSettingGetResponseZonesAlwaysOnlineIDAlwaysOnline ZoneSettingGetResponseZonesAlwaysOnlineID = "always_online"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesAlwaysOnlineValue string

const (
	ZoneSettingGetResponseZonesAlwaysOnlineValueOn  ZoneSettingGetResponseZonesAlwaysOnlineValue = "on"
	ZoneSettingGetResponseZonesAlwaysOnlineValueOff ZoneSettingGetResponseZonesAlwaysOnlineValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesAlwaysOnlineEditable bool

const (
	ZoneSettingGetResponseZonesAlwaysOnlineEditableTrue  ZoneSettingGetResponseZonesAlwaysOnlineEditable = true
	ZoneSettingGetResponseZonesAlwaysOnlineEditableFalse ZoneSettingGetResponseZonesAlwaysOnlineEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingGetResponseZonesAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesAlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesAlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesAlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesAlwaysUseHTTPSJSON `json:"-"`
}

// zoneSettingGetResponseZonesAlwaysUseHTTPSJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesAlwaysUseHTTPS]
type zoneSettingGetResponseZonesAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesAlwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesAlwaysUseHTTPS) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesAlwaysUseHTTPSID string

const (
	ZoneSettingGetResponseZonesAlwaysUseHTTPSIDAlwaysUseHTTPS ZoneSettingGetResponseZonesAlwaysUseHTTPSID = "always_use_https"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesAlwaysUseHTTPSValue string

const (
	ZoneSettingGetResponseZonesAlwaysUseHTTPSValueOn  ZoneSettingGetResponseZonesAlwaysUseHTTPSValue = "on"
	ZoneSettingGetResponseZonesAlwaysUseHTTPSValueOff ZoneSettingGetResponseZonesAlwaysUseHTTPSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesAlwaysUseHTTPSEditable bool

const (
	ZoneSettingGetResponseZonesAlwaysUseHTTPSEditableTrue  ZoneSettingGetResponseZonesAlwaysUseHTTPSEditable = true
	ZoneSettingGetResponseZonesAlwaysUseHTTPSEditableFalse ZoneSettingGetResponseZonesAlwaysUseHTTPSEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingGetResponseZonesAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesAutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesAutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesAutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesAutomaticHTTPSRewritesJSON `json:"-"`
}

// zoneSettingGetResponseZonesAutomaticHTTPSRewritesJSON contains the JSON metadata
// for the struct [ZoneSettingGetResponseZonesAutomaticHTTPSRewrites]
type zoneSettingGetResponseZonesAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesAutomaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesAutomaticHTTPSRewrites) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesAutomaticHTTPSRewritesID string

const (
	ZoneSettingGetResponseZonesAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites ZoneSettingGetResponseZonesAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesAutomaticHTTPSRewritesValue string

const (
	ZoneSettingGetResponseZonesAutomaticHTTPSRewritesValueOn  ZoneSettingGetResponseZonesAutomaticHTTPSRewritesValue = "on"
	ZoneSettingGetResponseZonesAutomaticHTTPSRewritesValueOff ZoneSettingGetResponseZonesAutomaticHTTPSRewritesValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesAutomaticHTTPSRewritesEditable bool

const (
	ZoneSettingGetResponseZonesAutomaticHTTPSRewritesEditableTrue  ZoneSettingGetResponseZonesAutomaticHTTPSRewritesEditable = true
	ZoneSettingGetResponseZonesAutomaticHTTPSRewritesEditableFalse ZoneSettingGetResponseZonesAutomaticHTTPSRewritesEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingGetResponseZonesBrotli struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesBrotliID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesBrotliValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesBrotliJSON `json:"-"`
}

// zoneSettingGetResponseZonesBrotliJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesBrotli]
type zoneSettingGetResponseZonesBrotliJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesBrotliJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesBrotli) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesBrotliID string

const (
	ZoneSettingGetResponseZonesBrotliIDBrotli ZoneSettingGetResponseZonesBrotliID = "brotli"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesBrotliValue string

const (
	ZoneSettingGetResponseZonesBrotliValueOff ZoneSettingGetResponseZonesBrotliValue = "off"
	ZoneSettingGetResponseZonesBrotliValueOn  ZoneSettingGetResponseZonesBrotliValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesBrotliEditable bool

const (
	ZoneSettingGetResponseZonesBrotliEditableTrue  ZoneSettingGetResponseZonesBrotliEditable = true
	ZoneSettingGetResponseZonesBrotliEditableFalse ZoneSettingGetResponseZonesBrotliEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingGetResponseZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesBrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesBrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesBrowserCacheTTLJSON `json:"-"`
}

// zoneSettingGetResponseZonesBrowserCacheTTLJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesBrowserCacheTTL]
type zoneSettingGetResponseZonesBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesBrowserCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesBrowserCacheTTL) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesBrowserCacheTTLID string

const (
	ZoneSettingGetResponseZonesBrowserCacheTTLIDBrowserCacheTTL ZoneSettingGetResponseZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesBrowserCacheTTLValue float64

const (
	ZoneSettingGetResponseZonesBrowserCacheTTLValue0        ZoneSettingGetResponseZonesBrowserCacheTTLValue = 0
	ZoneSettingGetResponseZonesBrowserCacheTTLValue30       ZoneSettingGetResponseZonesBrowserCacheTTLValue = 30
	ZoneSettingGetResponseZonesBrowserCacheTTLValue60       ZoneSettingGetResponseZonesBrowserCacheTTLValue = 60
	ZoneSettingGetResponseZonesBrowserCacheTTLValue120      ZoneSettingGetResponseZonesBrowserCacheTTLValue = 120
	ZoneSettingGetResponseZonesBrowserCacheTTLValue300      ZoneSettingGetResponseZonesBrowserCacheTTLValue = 300
	ZoneSettingGetResponseZonesBrowserCacheTTLValue1200     ZoneSettingGetResponseZonesBrowserCacheTTLValue = 1200
	ZoneSettingGetResponseZonesBrowserCacheTTLValue1800     ZoneSettingGetResponseZonesBrowserCacheTTLValue = 1800
	ZoneSettingGetResponseZonesBrowserCacheTTLValue3600     ZoneSettingGetResponseZonesBrowserCacheTTLValue = 3600
	ZoneSettingGetResponseZonesBrowserCacheTTLValue7200     ZoneSettingGetResponseZonesBrowserCacheTTLValue = 7200
	ZoneSettingGetResponseZonesBrowserCacheTTLValue10800    ZoneSettingGetResponseZonesBrowserCacheTTLValue = 10800
	ZoneSettingGetResponseZonesBrowserCacheTTLValue14400    ZoneSettingGetResponseZonesBrowserCacheTTLValue = 14400
	ZoneSettingGetResponseZonesBrowserCacheTTLValue18000    ZoneSettingGetResponseZonesBrowserCacheTTLValue = 18000
	ZoneSettingGetResponseZonesBrowserCacheTTLValue28800    ZoneSettingGetResponseZonesBrowserCacheTTLValue = 28800
	ZoneSettingGetResponseZonesBrowserCacheTTLValue43200    ZoneSettingGetResponseZonesBrowserCacheTTLValue = 43200
	ZoneSettingGetResponseZonesBrowserCacheTTLValue57600    ZoneSettingGetResponseZonesBrowserCacheTTLValue = 57600
	ZoneSettingGetResponseZonesBrowserCacheTTLValue72000    ZoneSettingGetResponseZonesBrowserCacheTTLValue = 72000
	ZoneSettingGetResponseZonesBrowserCacheTTLValue86400    ZoneSettingGetResponseZonesBrowserCacheTTLValue = 86400
	ZoneSettingGetResponseZonesBrowserCacheTTLValue172800   ZoneSettingGetResponseZonesBrowserCacheTTLValue = 172800
	ZoneSettingGetResponseZonesBrowserCacheTTLValue259200   ZoneSettingGetResponseZonesBrowserCacheTTLValue = 259200
	ZoneSettingGetResponseZonesBrowserCacheTTLValue345600   ZoneSettingGetResponseZonesBrowserCacheTTLValue = 345600
	ZoneSettingGetResponseZonesBrowserCacheTTLValue432000   ZoneSettingGetResponseZonesBrowserCacheTTLValue = 432000
	ZoneSettingGetResponseZonesBrowserCacheTTLValue691200   ZoneSettingGetResponseZonesBrowserCacheTTLValue = 691200
	ZoneSettingGetResponseZonesBrowserCacheTTLValue1382400  ZoneSettingGetResponseZonesBrowserCacheTTLValue = 1382400
	ZoneSettingGetResponseZonesBrowserCacheTTLValue2073600  ZoneSettingGetResponseZonesBrowserCacheTTLValue = 2073600
	ZoneSettingGetResponseZonesBrowserCacheTTLValue2678400  ZoneSettingGetResponseZonesBrowserCacheTTLValue = 2678400
	ZoneSettingGetResponseZonesBrowserCacheTTLValue5356800  ZoneSettingGetResponseZonesBrowserCacheTTLValue = 5356800
	ZoneSettingGetResponseZonesBrowserCacheTTLValue16070400 ZoneSettingGetResponseZonesBrowserCacheTTLValue = 16070400
	ZoneSettingGetResponseZonesBrowserCacheTTLValue31536000 ZoneSettingGetResponseZonesBrowserCacheTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesBrowserCacheTTLEditable bool

const (
	ZoneSettingGetResponseZonesBrowserCacheTTLEditableTrue  ZoneSettingGetResponseZonesBrowserCacheTTLEditable = true
	ZoneSettingGetResponseZonesBrowserCacheTTLEditableFalse ZoneSettingGetResponseZonesBrowserCacheTTLEditable = false
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingGetResponseZonesBrowserCheck struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesBrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesBrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesBrowserCheckJSON `json:"-"`
}

// zoneSettingGetResponseZonesBrowserCheckJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesBrowserCheck]
type zoneSettingGetResponseZonesBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesBrowserCheckJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesBrowserCheck) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesBrowserCheckID string

const (
	ZoneSettingGetResponseZonesBrowserCheckIDBrowserCheck ZoneSettingGetResponseZonesBrowserCheckID = "browser_check"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesBrowserCheckValue string

const (
	ZoneSettingGetResponseZonesBrowserCheckValueOn  ZoneSettingGetResponseZonesBrowserCheckValue = "on"
	ZoneSettingGetResponseZonesBrowserCheckValueOff ZoneSettingGetResponseZonesBrowserCheckValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesBrowserCheckEditable bool

const (
	ZoneSettingGetResponseZonesBrowserCheckEditableTrue  ZoneSettingGetResponseZonesBrowserCheckEditable = true
	ZoneSettingGetResponseZonesBrowserCheckEditableFalse ZoneSettingGetResponseZonesBrowserCheckEditable = false
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingGetResponseZonesCacheLevel struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesCacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesCacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesCacheLevelJSON `json:"-"`
}

// zoneSettingGetResponseZonesCacheLevelJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesCacheLevel]
type zoneSettingGetResponseZonesCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesCacheLevelJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesCacheLevel) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesCacheLevelID string

const (
	ZoneSettingGetResponseZonesCacheLevelIDCacheLevel ZoneSettingGetResponseZonesCacheLevelID = "cache_level"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesCacheLevelValue string

const (
	ZoneSettingGetResponseZonesCacheLevelValueAggressive ZoneSettingGetResponseZonesCacheLevelValue = "aggressive"
	ZoneSettingGetResponseZonesCacheLevelValueBasic      ZoneSettingGetResponseZonesCacheLevelValue = "basic"
	ZoneSettingGetResponseZonesCacheLevelValueSimplified ZoneSettingGetResponseZonesCacheLevelValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesCacheLevelEditable bool

const (
	ZoneSettingGetResponseZonesCacheLevelEditableTrue  ZoneSettingGetResponseZonesCacheLevelEditable = true
	ZoneSettingGetResponseZonesCacheLevelEditableFalse ZoneSettingGetResponseZonesCacheLevelEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingGetResponseZonesChallengeTTL struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesChallengeTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesChallengeTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesChallengeTTLJSON `json:"-"`
}

// zoneSettingGetResponseZonesChallengeTTLJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesChallengeTTL]
type zoneSettingGetResponseZonesChallengeTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesChallengeTTLJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesChallengeTTL) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesChallengeTTLID string

const (
	ZoneSettingGetResponseZonesChallengeTTLIDChallengeTTL ZoneSettingGetResponseZonesChallengeTTLID = "challenge_ttl"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesChallengeTTLValue float64

const (
	ZoneSettingGetResponseZonesChallengeTTLValue300      ZoneSettingGetResponseZonesChallengeTTLValue = 300
	ZoneSettingGetResponseZonesChallengeTTLValue900      ZoneSettingGetResponseZonesChallengeTTLValue = 900
	ZoneSettingGetResponseZonesChallengeTTLValue1800     ZoneSettingGetResponseZonesChallengeTTLValue = 1800
	ZoneSettingGetResponseZonesChallengeTTLValue2700     ZoneSettingGetResponseZonesChallengeTTLValue = 2700
	ZoneSettingGetResponseZonesChallengeTTLValue3600     ZoneSettingGetResponseZonesChallengeTTLValue = 3600
	ZoneSettingGetResponseZonesChallengeTTLValue7200     ZoneSettingGetResponseZonesChallengeTTLValue = 7200
	ZoneSettingGetResponseZonesChallengeTTLValue10800    ZoneSettingGetResponseZonesChallengeTTLValue = 10800
	ZoneSettingGetResponseZonesChallengeTTLValue14400    ZoneSettingGetResponseZonesChallengeTTLValue = 14400
	ZoneSettingGetResponseZonesChallengeTTLValue28800    ZoneSettingGetResponseZonesChallengeTTLValue = 28800
	ZoneSettingGetResponseZonesChallengeTTLValue57600    ZoneSettingGetResponseZonesChallengeTTLValue = 57600
	ZoneSettingGetResponseZonesChallengeTTLValue86400    ZoneSettingGetResponseZonesChallengeTTLValue = 86400
	ZoneSettingGetResponseZonesChallengeTTLValue604800   ZoneSettingGetResponseZonesChallengeTTLValue = 604800
	ZoneSettingGetResponseZonesChallengeTTLValue2592000  ZoneSettingGetResponseZonesChallengeTTLValue = 2592000
	ZoneSettingGetResponseZonesChallengeTTLValue31536000 ZoneSettingGetResponseZonesChallengeTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesChallengeTTLEditable bool

const (
	ZoneSettingGetResponseZonesChallengeTTLEditableTrue  ZoneSettingGetResponseZonesChallengeTTLEditable = true
	ZoneSettingGetResponseZonesChallengeTTLEditableFalse ZoneSettingGetResponseZonesChallengeTTLEditable = false
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingGetResponseZonesCiphers struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesCiphersID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesCiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesCiphersJSON `json:"-"`
}

// zoneSettingGetResponseZonesCiphersJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesCiphers]
type zoneSettingGetResponseZonesCiphersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesCiphersJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesCiphers) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesCiphersID string

const (
	ZoneSettingGetResponseZonesCiphersIDCiphers ZoneSettingGetResponseZonesCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesCiphersEditable bool

const (
	ZoneSettingGetResponseZonesCiphersEditableTrue  ZoneSettingGetResponseZonesCiphersEditable = true
	ZoneSettingGetResponseZonesCiphersEditableFalse ZoneSettingGetResponseZonesCiphersEditable = false
)

// Whether or not cname flattening is on.
type ZoneSettingGetResponseZonesCNAMEFlattening struct {
	// How to flatten the cname destination.
	ID ZoneSettingGetResponseZonesCNAMEFlatteningID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesCNAMEFlatteningValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesCNAMEFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesCNAMEFlatteningJSON `json:"-"`
}

// zoneSettingGetResponseZonesCNAMEFlatteningJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesCNAMEFlattening]
type zoneSettingGetResponseZonesCNAMEFlatteningJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesCNAMEFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesCNAMEFlatteningJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesCNAMEFlattening) implementsZoneSettingGetResponse() {}

// How to flatten the cname destination.
type ZoneSettingGetResponseZonesCNAMEFlatteningID string

const (
	ZoneSettingGetResponseZonesCNAMEFlatteningIDCNAMEFlattening ZoneSettingGetResponseZonesCNAMEFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesCNAMEFlatteningValue string

const (
	ZoneSettingGetResponseZonesCNAMEFlatteningValueFlattenAtRoot ZoneSettingGetResponseZonesCNAMEFlatteningValue = "flatten_at_root"
	ZoneSettingGetResponseZonesCNAMEFlatteningValueFlattenAll    ZoneSettingGetResponseZonesCNAMEFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesCNAMEFlatteningEditable bool

const (
	ZoneSettingGetResponseZonesCNAMEFlatteningEditableTrue  ZoneSettingGetResponseZonesCNAMEFlatteningEditable = true
	ZoneSettingGetResponseZonesCNAMEFlatteningEditableFalse ZoneSettingGetResponseZonesCNAMEFlatteningEditable = false
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingGetResponseZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesDevelopmentModeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesDevelopmentModeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                                        `json:"time_remaining"`
	JSON          zoneSettingGetResponseZonesDevelopmentModeJSON `json:"-"`
}

// zoneSettingGetResponseZonesDevelopmentModeJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesDevelopmentMode]
type zoneSettingGetResponseZonesDevelopmentModeJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesDevelopmentModeJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesDevelopmentMode) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesDevelopmentModeID string

const (
	ZoneSettingGetResponseZonesDevelopmentModeIDDevelopmentMode ZoneSettingGetResponseZonesDevelopmentModeID = "development_mode"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesDevelopmentModeValue string

const (
	ZoneSettingGetResponseZonesDevelopmentModeValueOn  ZoneSettingGetResponseZonesDevelopmentModeValue = "on"
	ZoneSettingGetResponseZonesDevelopmentModeValueOff ZoneSettingGetResponseZonesDevelopmentModeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesDevelopmentModeEditable bool

const (
	ZoneSettingGetResponseZonesDevelopmentModeEditableTrue  ZoneSettingGetResponseZonesDevelopmentModeEditable = true
	ZoneSettingGetResponseZonesDevelopmentModeEditableFalse ZoneSettingGetResponseZonesDevelopmentModeEditable = false
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingGetResponseZonesEarlyHints struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesEarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesEarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesEarlyHintsJSON `json:"-"`
}

// zoneSettingGetResponseZonesEarlyHintsJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesEarlyHints]
type zoneSettingGetResponseZonesEarlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesEarlyHintsJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesEarlyHints) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesEarlyHintsID string

const (
	ZoneSettingGetResponseZonesEarlyHintsIDEarlyHints ZoneSettingGetResponseZonesEarlyHintsID = "early_hints"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesEarlyHintsValue string

const (
	ZoneSettingGetResponseZonesEarlyHintsValueOn  ZoneSettingGetResponseZonesEarlyHintsValue = "on"
	ZoneSettingGetResponseZonesEarlyHintsValueOff ZoneSettingGetResponseZonesEarlyHintsValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesEarlyHintsEditable bool

const (
	ZoneSettingGetResponseZonesEarlyHintsEditableTrue  ZoneSettingGetResponseZonesEarlyHintsEditable = true
	ZoneSettingGetResponseZonesEarlyHintsEditableFalse ZoneSettingGetResponseZonesEarlyHintsEditable = false
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

func (r zoneSettingGetResponseZonesEdgeCacheTTLJSON) RawJSON() string {
	return r.raw
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

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingGetResponseZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesEmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesEmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesEmailObfuscationJSON `json:"-"`
}

// zoneSettingGetResponseZonesEmailObfuscationJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesEmailObfuscation]
type zoneSettingGetResponseZonesEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesEmailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesEmailObfuscation) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesEmailObfuscationID string

const (
	ZoneSettingGetResponseZonesEmailObfuscationIDEmailObfuscation ZoneSettingGetResponseZonesEmailObfuscationID = "email_obfuscation"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesEmailObfuscationValue string

const (
	ZoneSettingGetResponseZonesEmailObfuscationValueOn  ZoneSettingGetResponseZonesEmailObfuscationValue = "on"
	ZoneSettingGetResponseZonesEmailObfuscationValueOff ZoneSettingGetResponseZonesEmailObfuscationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesEmailObfuscationEditable bool

const (
	ZoneSettingGetResponseZonesEmailObfuscationEditableTrue  ZoneSettingGetResponseZonesEmailObfuscationEditable = true
	ZoneSettingGetResponseZonesEmailObfuscationEditableFalse ZoneSettingGetResponseZonesEmailObfuscationEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingGetResponseZonesH2Prioritization struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesH2PrioritizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesH2PrioritizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesH2PrioritizationJSON `json:"-"`
}

// zoneSettingGetResponseZonesH2PrioritizationJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesH2Prioritization]
type zoneSettingGetResponseZonesH2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesH2PrioritizationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesH2Prioritization) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesH2PrioritizationID string

const (
	ZoneSettingGetResponseZonesH2PrioritizationIDH2Prioritization ZoneSettingGetResponseZonesH2PrioritizationID = "h2_prioritization"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesH2PrioritizationValue string

const (
	ZoneSettingGetResponseZonesH2PrioritizationValueOn     ZoneSettingGetResponseZonesH2PrioritizationValue = "on"
	ZoneSettingGetResponseZonesH2PrioritizationValueOff    ZoneSettingGetResponseZonesH2PrioritizationValue = "off"
	ZoneSettingGetResponseZonesH2PrioritizationValueCustom ZoneSettingGetResponseZonesH2PrioritizationValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesH2PrioritizationEditable bool

const (
	ZoneSettingGetResponseZonesH2PrioritizationEditableTrue  ZoneSettingGetResponseZonesH2PrioritizationEditable = true
	ZoneSettingGetResponseZonesH2PrioritizationEditableFalse ZoneSettingGetResponseZonesH2PrioritizationEditable = false
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZoneSettingGetResponseZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesHotlinkProtectionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesHotlinkProtectionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesHotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesHotlinkProtectionJSON `json:"-"`
}

// zoneSettingGetResponseZonesHotlinkProtectionJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesHotlinkProtection]
type zoneSettingGetResponseZonesHotlinkProtectionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesHotlinkProtectionJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesHotlinkProtection) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesHotlinkProtectionID string

const (
	ZoneSettingGetResponseZonesHotlinkProtectionIDHotlinkProtection ZoneSettingGetResponseZonesHotlinkProtectionID = "hotlink_protection"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesHotlinkProtectionValue string

const (
	ZoneSettingGetResponseZonesHotlinkProtectionValueOn  ZoneSettingGetResponseZonesHotlinkProtectionValue = "on"
	ZoneSettingGetResponseZonesHotlinkProtectionValueOff ZoneSettingGetResponseZonesHotlinkProtectionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesHotlinkProtectionEditable bool

const (
	ZoneSettingGetResponseZonesHotlinkProtectionEditableTrue  ZoneSettingGetResponseZonesHotlinkProtectionEditable = true
	ZoneSettingGetResponseZonesHotlinkProtectionEditableFalse ZoneSettingGetResponseZonesHotlinkProtectionEditable = false
)

// HTTP2 enabled for this zone.
type ZoneSettingGetResponseZonesHTTP2 struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesHTTP2ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesHTTP2Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesHTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesHTTP2JSON `json:"-"`
}

// zoneSettingGetResponseZonesHTTP2JSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesHTTP2]
type zoneSettingGetResponseZonesHTTP2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesHTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesHTTP2JSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesHTTP2) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesHTTP2ID string

const (
	ZoneSettingGetResponseZonesHTTP2IDHTTP2 ZoneSettingGetResponseZonesHTTP2ID = "http2"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesHTTP2Value string

const (
	ZoneSettingGetResponseZonesHTTP2ValueOn  ZoneSettingGetResponseZonesHTTP2Value = "on"
	ZoneSettingGetResponseZonesHTTP2ValueOff ZoneSettingGetResponseZonesHTTP2Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesHTTP2Editable bool

const (
	ZoneSettingGetResponseZonesHTTP2EditableTrue  ZoneSettingGetResponseZonesHTTP2Editable = true
	ZoneSettingGetResponseZonesHTTP2EditableFalse ZoneSettingGetResponseZonesHTTP2Editable = false
)

// HTTP3 enabled for this zone.
type ZoneSettingGetResponseZonesHTTP3 struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesHTTP3ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesHTTP3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesHTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesHTTP3JSON `json:"-"`
}

// zoneSettingGetResponseZonesHTTP3JSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesHTTP3]
type zoneSettingGetResponseZonesHTTP3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesHTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesHTTP3JSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesHTTP3) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesHTTP3ID string

const (
	ZoneSettingGetResponseZonesHTTP3IDHTTP3 ZoneSettingGetResponseZonesHTTP3ID = "http3"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesHTTP3Value string

const (
	ZoneSettingGetResponseZonesHTTP3ValueOn  ZoneSettingGetResponseZonesHTTP3Value = "on"
	ZoneSettingGetResponseZonesHTTP3ValueOff ZoneSettingGetResponseZonesHTTP3Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesHTTP3Editable bool

const (
	ZoneSettingGetResponseZonesHTTP3EditableTrue  ZoneSettingGetResponseZonesHTTP3Editable = true
	ZoneSettingGetResponseZonesHTTP3EditableFalse ZoneSettingGetResponseZonesHTTP3Editable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingGetResponseZonesImageResizing struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesImageResizingID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesImageResizingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesImageResizingJSON `json:"-"`
}

// zoneSettingGetResponseZonesImageResizingJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesImageResizing]
type zoneSettingGetResponseZonesImageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesImageResizingJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesImageResizing) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesImageResizingID string

const (
	ZoneSettingGetResponseZonesImageResizingIDImageResizing ZoneSettingGetResponseZonesImageResizingID = "image_resizing"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesImageResizingValue string

const (
	ZoneSettingGetResponseZonesImageResizingValueOn   ZoneSettingGetResponseZonesImageResizingValue = "on"
	ZoneSettingGetResponseZonesImageResizingValueOff  ZoneSettingGetResponseZonesImageResizingValue = "off"
	ZoneSettingGetResponseZonesImageResizingValueOpen ZoneSettingGetResponseZonesImageResizingValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesImageResizingEditable bool

const (
	ZoneSettingGetResponseZonesImageResizingEditableTrue  ZoneSettingGetResponseZonesImageResizingEditable = true
	ZoneSettingGetResponseZonesImageResizingEditableFalse ZoneSettingGetResponseZonesImageResizingEditable = false
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingGetResponseZonesIPGeolocation struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesIPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesIPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesIPGeolocationJSON `json:"-"`
}

// zoneSettingGetResponseZonesIPGeolocationJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesIPGeolocation]
type zoneSettingGetResponseZonesIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesIPGeolocationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesIPGeolocation) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesIPGeolocationID string

const (
	ZoneSettingGetResponseZonesIPGeolocationIDIPGeolocation ZoneSettingGetResponseZonesIPGeolocationID = "ip_geolocation"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesIPGeolocationValue string

const (
	ZoneSettingGetResponseZonesIPGeolocationValueOn  ZoneSettingGetResponseZonesIPGeolocationValue = "on"
	ZoneSettingGetResponseZonesIPGeolocationValueOff ZoneSettingGetResponseZonesIPGeolocationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesIPGeolocationEditable bool

const (
	ZoneSettingGetResponseZonesIPGeolocationEditableTrue  ZoneSettingGetResponseZonesIPGeolocationEditable = true
	ZoneSettingGetResponseZonesIPGeolocationEditableFalse ZoneSettingGetResponseZonesIPGeolocationEditable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingGetResponseZonesIPV6 struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesIPV6ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesIPV6Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesIPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesIPV6JSON `json:"-"`
}

// zoneSettingGetResponseZonesIPV6JSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesIPV6]
type zoneSettingGetResponseZonesIPV6JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesIPV6JSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesIPV6) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesIPV6ID string

const (
	ZoneSettingGetResponseZonesIPV6IDIPV6 ZoneSettingGetResponseZonesIPV6ID = "ipv6"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesIPV6Value string

const (
	ZoneSettingGetResponseZonesIPV6ValueOff ZoneSettingGetResponseZonesIPV6Value = "off"
	ZoneSettingGetResponseZonesIPV6ValueOn  ZoneSettingGetResponseZonesIPV6Value = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesIPV6Editable bool

const (
	ZoneSettingGetResponseZonesIPV6EditableTrue  ZoneSettingGetResponseZonesIPV6Editable = true
	ZoneSettingGetResponseZonesIPV6EditableFalse ZoneSettingGetResponseZonesIPV6Editable = false
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

func (r zoneSettingGetResponseZonesMaxUploadJSON) RawJSON() string {
	return r.raw
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

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingGetResponseZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesMinTLSVersionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesMinTLSVersionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesMinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesMinTLSVersionJSON `json:"-"`
}

// zoneSettingGetResponseZonesMinTLSVersionJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesMinTLSVersion]
type zoneSettingGetResponseZonesMinTLSVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesMinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesMinTLSVersionJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesMinTLSVersion) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesMinTLSVersionID string

const (
	ZoneSettingGetResponseZonesMinTLSVersionIDMinTLSVersion ZoneSettingGetResponseZonesMinTLSVersionID = "min_tls_version"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesMinTLSVersionValue string

const (
	ZoneSettingGetResponseZonesMinTLSVersionValue1_0 ZoneSettingGetResponseZonesMinTLSVersionValue = "1.0"
	ZoneSettingGetResponseZonesMinTLSVersionValue1_1 ZoneSettingGetResponseZonesMinTLSVersionValue = "1.1"
	ZoneSettingGetResponseZonesMinTLSVersionValue1_2 ZoneSettingGetResponseZonesMinTLSVersionValue = "1.2"
	ZoneSettingGetResponseZonesMinTLSVersionValue1_3 ZoneSettingGetResponseZonesMinTLSVersionValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesMinTLSVersionEditable bool

const (
	ZoneSettingGetResponseZonesMinTLSVersionEditableTrue  ZoneSettingGetResponseZonesMinTLSVersionEditable = true
	ZoneSettingGetResponseZonesMinTLSVersionEditableFalse ZoneSettingGetResponseZonesMinTLSVersionEditable = false
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingGetResponseZonesMinify struct {
	// Zone setting identifier.
	ID ZoneSettingGetResponseZonesMinifyID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesMinifyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesMinifyJSON `json:"-"`
}

// zoneSettingGetResponseZonesMinifyJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesMinify]
type zoneSettingGetResponseZonesMinifyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesMinifyJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesMinify) implementsZoneSettingGetResponse() {}

// Zone setting identifier.
type ZoneSettingGetResponseZonesMinifyID string

const (
	ZoneSettingGetResponseZonesMinifyIDMinify ZoneSettingGetResponseZonesMinifyID = "minify"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css ZoneSettingGetResponseZonesMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML ZoneSettingGetResponseZonesMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   ZoneSettingGetResponseZonesMinifyValueJs   `json:"js"`
	JSON zoneSettingGetResponseZonesMinifyValueJSON `json:"-"`
}

// zoneSettingGetResponseZonesMinifyValueJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesMinifyValue]
type zoneSettingGetResponseZonesMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesMinifyValueJSON) RawJSON() string {
	return r.raw
}

// Automatically minify all CSS files for your website.
type ZoneSettingGetResponseZonesMinifyValueCss string

const (
	ZoneSettingGetResponseZonesMinifyValueCssOn  ZoneSettingGetResponseZonesMinifyValueCss = "on"
	ZoneSettingGetResponseZonesMinifyValueCssOff ZoneSettingGetResponseZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingGetResponseZonesMinifyValueHTML string

const (
	ZoneSettingGetResponseZonesMinifyValueHTMLOn  ZoneSettingGetResponseZonesMinifyValueHTML = "on"
	ZoneSettingGetResponseZonesMinifyValueHTMLOff ZoneSettingGetResponseZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingGetResponseZonesMinifyValueJs string

const (
	ZoneSettingGetResponseZonesMinifyValueJsOn  ZoneSettingGetResponseZonesMinifyValueJs = "on"
	ZoneSettingGetResponseZonesMinifyValueJsOff ZoneSettingGetResponseZonesMinifyValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesMinifyEditable bool

const (
	ZoneSettingGetResponseZonesMinifyEditableTrue  ZoneSettingGetResponseZonesMinifyEditable = true
	ZoneSettingGetResponseZonesMinifyEditableFalse ZoneSettingGetResponseZonesMinifyEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingGetResponseZonesMirage struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesMirageID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesMirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesMirageJSON `json:"-"`
}

// zoneSettingGetResponseZonesMirageJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesMirage]
type zoneSettingGetResponseZonesMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesMirageJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesMirage) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesMirageID string

const (
	ZoneSettingGetResponseZonesMirageIDMirage ZoneSettingGetResponseZonesMirageID = "mirage"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesMirageValue string

const (
	ZoneSettingGetResponseZonesMirageValueOn  ZoneSettingGetResponseZonesMirageValue = "on"
	ZoneSettingGetResponseZonesMirageValueOff ZoneSettingGetResponseZonesMirageValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesMirageEditable bool

const (
	ZoneSettingGetResponseZonesMirageEditableTrue  ZoneSettingGetResponseZonesMirageEditable = true
	ZoneSettingGetResponseZonesMirageEditableFalse ZoneSettingGetResponseZonesMirageEditable = false
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingGetResponseZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID ZoneSettingGetResponseZonesMobileRedirectID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesMobileRedirectValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesMobileRedirectJSON `json:"-"`
}

// zoneSettingGetResponseZonesMobileRedirectJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesMobileRedirect]
type zoneSettingGetResponseZonesMobileRedirectJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesMobileRedirectJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesMobileRedirect) implementsZoneSettingGetResponse() {}

// Identifier of the zone setting.
type ZoneSettingGetResponseZonesMobileRedirectID string

const (
	ZoneSettingGetResponseZonesMobileRedirectIDMobileRedirect ZoneSettingGetResponseZonesMobileRedirectID = "mobile_redirect"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingGetResponseZonesMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripURI bool                                               `json:"strip_uri"`
	JSON     zoneSettingGetResponseZonesMobileRedirectValueJSON `json:"-"`
}

// zoneSettingGetResponseZonesMobileRedirectValueJSON contains the JSON metadata
// for the struct [ZoneSettingGetResponseZonesMobileRedirectValue]
type zoneSettingGetResponseZonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripURI        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesMobileRedirectValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not mobile redirect is enabled.
type ZoneSettingGetResponseZonesMobileRedirectValueStatus string

const (
	ZoneSettingGetResponseZonesMobileRedirectValueStatusOn  ZoneSettingGetResponseZonesMobileRedirectValueStatus = "on"
	ZoneSettingGetResponseZonesMobileRedirectValueStatusOff ZoneSettingGetResponseZonesMobileRedirectValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesMobileRedirectEditable bool

const (
	ZoneSettingGetResponseZonesMobileRedirectEditableTrue  ZoneSettingGetResponseZonesMobileRedirectEditable = true
	ZoneSettingGetResponseZonesMobileRedirectEditableFalse ZoneSettingGetResponseZonesMobileRedirectEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingGetResponseZonesNEL struct {
	// Zone setting identifier.
	ID ZoneSettingGetResponseZonesNELID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesNELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesNELJSON `json:"-"`
}

// zoneSettingGetResponseZonesNELJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesNEL]
type zoneSettingGetResponseZonesNELJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesNELJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesNEL) implementsZoneSettingGetResponse() {}

// Zone setting identifier.
type ZoneSettingGetResponseZonesNELID string

const (
	ZoneSettingGetResponseZonesNELIDNEL ZoneSettingGetResponseZonesNELID = "nel"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesNELValue struct {
	Enabled bool                                    `json:"enabled"`
	JSON    zoneSettingGetResponseZonesNELValueJSON `json:"-"`
}

// zoneSettingGetResponseZonesNELValueJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesNELValue]
type zoneSettingGetResponseZonesNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesNELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesNELValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesNELEditable bool

const (
	ZoneSettingGetResponseZonesNELEditableTrue  ZoneSettingGetResponseZonesNELEditable = true
	ZoneSettingGetResponseZonesNELEditableFalse ZoneSettingGetResponseZonesNELEditable = false
)

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingGetResponseZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesOpportunisticEncryptionJSON `json:"-"`
}

// zoneSettingGetResponseZonesOpportunisticEncryptionJSON contains the JSON
// metadata for the struct [ZoneSettingGetResponseZonesOpportunisticEncryption]
type zoneSettingGetResponseZonesOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesOpportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesOpportunisticEncryption) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesOpportunisticEncryptionID string

const (
	ZoneSettingGetResponseZonesOpportunisticEncryptionIDOpportunisticEncryption ZoneSettingGetResponseZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesOpportunisticEncryptionValue string

const (
	ZoneSettingGetResponseZonesOpportunisticEncryptionValueOn  ZoneSettingGetResponseZonesOpportunisticEncryptionValue = "on"
	ZoneSettingGetResponseZonesOpportunisticEncryptionValueOff ZoneSettingGetResponseZonesOpportunisticEncryptionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesOpportunisticEncryptionEditable bool

const (
	ZoneSettingGetResponseZonesOpportunisticEncryptionEditableTrue  ZoneSettingGetResponseZonesOpportunisticEncryptionEditable = true
	ZoneSettingGetResponseZonesOpportunisticEncryptionEditableFalse ZoneSettingGetResponseZonesOpportunisticEncryptionEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingGetResponseZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesOpportunisticOnionID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesOpportunisticOnionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesOpportunisticOnionJSON `json:"-"`
}

// zoneSettingGetResponseZonesOpportunisticOnionJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesOpportunisticOnion]
type zoneSettingGetResponseZonesOpportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesOpportunisticOnionJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesOpportunisticOnion) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesOpportunisticOnionID string

const (
	ZoneSettingGetResponseZonesOpportunisticOnionIDOpportunisticOnion ZoneSettingGetResponseZonesOpportunisticOnionID = "opportunistic_onion"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesOpportunisticOnionValue string

const (
	ZoneSettingGetResponseZonesOpportunisticOnionValueOn  ZoneSettingGetResponseZonesOpportunisticOnionValue = "on"
	ZoneSettingGetResponseZonesOpportunisticOnionValueOff ZoneSettingGetResponseZonesOpportunisticOnionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesOpportunisticOnionEditable bool

const (
	ZoneSettingGetResponseZonesOpportunisticOnionEditableTrue  ZoneSettingGetResponseZonesOpportunisticOnionEditable = true
	ZoneSettingGetResponseZonesOpportunisticOnionEditableFalse ZoneSettingGetResponseZonesOpportunisticOnionEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingGetResponseZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesOrangeToOrangeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesOrangeToOrangeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesOrangeToOrangeJSON `json:"-"`
}

// zoneSettingGetResponseZonesOrangeToOrangeJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesOrangeToOrange]
type zoneSettingGetResponseZonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesOrangeToOrangeJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesOrangeToOrange) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesOrangeToOrangeID string

const (
	ZoneSettingGetResponseZonesOrangeToOrangeIDOrangeToOrange ZoneSettingGetResponseZonesOrangeToOrangeID = "orange_to_orange"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesOrangeToOrangeValue string

const (
	ZoneSettingGetResponseZonesOrangeToOrangeValueOn  ZoneSettingGetResponseZonesOrangeToOrangeValue = "on"
	ZoneSettingGetResponseZonesOrangeToOrangeValueOff ZoneSettingGetResponseZonesOrangeToOrangeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesOrangeToOrangeEditable bool

const (
	ZoneSettingGetResponseZonesOrangeToOrangeEditableTrue  ZoneSettingGetResponseZonesOrangeToOrangeEditable = true
	ZoneSettingGetResponseZonesOrangeToOrangeEditableFalse ZoneSettingGetResponseZonesOrangeToOrangeEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZoneSettingGetResponseZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesOriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesOriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesOriginErrorPagePassThruJSON `json:"-"`
}

// zoneSettingGetResponseZonesOriginErrorPagePassThruJSON contains the JSON
// metadata for the struct [ZoneSettingGetResponseZonesOriginErrorPagePassThru]
type zoneSettingGetResponseZonesOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesOriginErrorPagePassThruJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesOriginErrorPagePassThru) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesOriginErrorPagePassThruID string

const (
	ZoneSettingGetResponseZonesOriginErrorPagePassThruIDOriginErrorPagePassThru ZoneSettingGetResponseZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesOriginErrorPagePassThruValue string

const (
	ZoneSettingGetResponseZonesOriginErrorPagePassThruValueOn  ZoneSettingGetResponseZonesOriginErrorPagePassThruValue = "on"
	ZoneSettingGetResponseZonesOriginErrorPagePassThruValueOff ZoneSettingGetResponseZonesOriginErrorPagePassThruValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesOriginErrorPagePassThruEditable bool

const (
	ZoneSettingGetResponseZonesOriginErrorPagePassThruEditableTrue  ZoneSettingGetResponseZonesOriginErrorPagePassThruEditable = true
	ZoneSettingGetResponseZonesOriginErrorPagePassThruEditableFalse ZoneSettingGetResponseZonesOriginErrorPagePassThruEditable = false
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingGetResponseZonesPolish struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesPolishID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesPolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesPolishJSON `json:"-"`
}

// zoneSettingGetResponseZonesPolishJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesPolish]
type zoneSettingGetResponseZonesPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesPolishJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesPolish) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesPolishID string

const (
	ZoneSettingGetResponseZonesPolishIDPolish ZoneSettingGetResponseZonesPolishID = "polish"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesPolishValue string

const (
	ZoneSettingGetResponseZonesPolishValueOff      ZoneSettingGetResponseZonesPolishValue = "off"
	ZoneSettingGetResponseZonesPolishValueLossless ZoneSettingGetResponseZonesPolishValue = "lossless"
	ZoneSettingGetResponseZonesPolishValueLossy    ZoneSettingGetResponseZonesPolishValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesPolishEditable bool

const (
	ZoneSettingGetResponseZonesPolishEditableTrue  ZoneSettingGetResponseZonesPolishEditable = true
	ZoneSettingGetResponseZonesPolishEditableFalse ZoneSettingGetResponseZonesPolishEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingGetResponseZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesPrefetchPreloadID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesPrefetchPreloadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesPrefetchPreloadJSON `json:"-"`
}

// zoneSettingGetResponseZonesPrefetchPreloadJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesPrefetchPreload]
type zoneSettingGetResponseZonesPrefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesPrefetchPreloadJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesPrefetchPreload) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesPrefetchPreloadID string

const (
	ZoneSettingGetResponseZonesPrefetchPreloadIDPrefetchPreload ZoneSettingGetResponseZonesPrefetchPreloadID = "prefetch_preload"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesPrefetchPreloadValue string

const (
	ZoneSettingGetResponseZonesPrefetchPreloadValueOn  ZoneSettingGetResponseZonesPrefetchPreloadValue = "on"
	ZoneSettingGetResponseZonesPrefetchPreloadValueOff ZoneSettingGetResponseZonesPrefetchPreloadValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesPrefetchPreloadEditable bool

const (
	ZoneSettingGetResponseZonesPrefetchPreloadEditableTrue  ZoneSettingGetResponseZonesPrefetchPreloadEditable = true
	ZoneSettingGetResponseZonesPrefetchPreloadEditableFalse ZoneSettingGetResponseZonesPrefetchPreloadEditable = false
)

// Maximum time between two read operations from origin.
type ZoneSettingGetResponseZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesProxyReadTimeoutJSON `json:"-"`
}

// zoneSettingGetResponseZonesProxyReadTimeoutJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesProxyReadTimeout]
type zoneSettingGetResponseZonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesProxyReadTimeoutJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesProxyReadTimeout) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesProxyReadTimeoutID string

const (
	ZoneSettingGetResponseZonesProxyReadTimeoutIDProxyReadTimeout ZoneSettingGetResponseZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesProxyReadTimeoutEditable bool

const (
	ZoneSettingGetResponseZonesProxyReadTimeoutEditableTrue  ZoneSettingGetResponseZonesProxyReadTimeoutEditable = true
	ZoneSettingGetResponseZonesProxyReadTimeoutEditableFalse ZoneSettingGetResponseZonesProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type ZoneSettingGetResponseZonesPseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID ZoneSettingGetResponseZonesPseudoIPV4ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesPseudoIPV4Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesPseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesPseudoIPV4JSON `json:"-"`
}

// zoneSettingGetResponseZonesPseudoIPV4JSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesPseudoIPV4]
type zoneSettingGetResponseZonesPseudoIPV4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesPseudoIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesPseudoIPV4JSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesPseudoIPV4) implementsZoneSettingGetResponse() {}

// Value of the Pseudo IPv4 setting.
type ZoneSettingGetResponseZonesPseudoIPV4ID string

const (
	ZoneSettingGetResponseZonesPseudoIPV4IDPseudoIPV4 ZoneSettingGetResponseZonesPseudoIPV4ID = "pseudo_ipv4"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesPseudoIPV4Value string

const (
	ZoneSettingGetResponseZonesPseudoIPV4ValueOff             ZoneSettingGetResponseZonesPseudoIPV4Value = "off"
	ZoneSettingGetResponseZonesPseudoIPV4ValueAddHeader       ZoneSettingGetResponseZonesPseudoIPV4Value = "add_header"
	ZoneSettingGetResponseZonesPseudoIPV4ValueOverwriteHeader ZoneSettingGetResponseZonesPseudoIPV4Value = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesPseudoIPV4Editable bool

const (
	ZoneSettingGetResponseZonesPseudoIPV4EditableTrue  ZoneSettingGetResponseZonesPseudoIPV4Editable = true
	ZoneSettingGetResponseZonesPseudoIPV4EditableFalse ZoneSettingGetResponseZonesPseudoIPV4Editable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingGetResponseZonesResponseBuffering struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesResponseBufferingID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesResponseBufferingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesResponseBufferingJSON `json:"-"`
}

// zoneSettingGetResponseZonesResponseBufferingJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesResponseBuffering]
type zoneSettingGetResponseZonesResponseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesResponseBufferingJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesResponseBuffering) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesResponseBufferingID string

const (
	ZoneSettingGetResponseZonesResponseBufferingIDResponseBuffering ZoneSettingGetResponseZonesResponseBufferingID = "response_buffering"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesResponseBufferingValue string

const (
	ZoneSettingGetResponseZonesResponseBufferingValueOn  ZoneSettingGetResponseZonesResponseBufferingValue = "on"
	ZoneSettingGetResponseZonesResponseBufferingValueOff ZoneSettingGetResponseZonesResponseBufferingValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesResponseBufferingEditable bool

const (
	ZoneSettingGetResponseZonesResponseBufferingEditableTrue  ZoneSettingGetResponseZonesResponseBufferingEditable = true
	ZoneSettingGetResponseZonesResponseBufferingEditableFalse ZoneSettingGetResponseZonesResponseBufferingEditable = false
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
type ZoneSettingGetResponseZonesRocketLoader struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesRocketLoaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesRocketLoaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesRocketLoaderJSON `json:"-"`
}

// zoneSettingGetResponseZonesRocketLoaderJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesRocketLoader]
type zoneSettingGetResponseZonesRocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesRocketLoaderJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesRocketLoader) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesRocketLoaderID string

const (
	ZoneSettingGetResponseZonesRocketLoaderIDRocketLoader ZoneSettingGetResponseZonesRocketLoaderID = "rocket_loader"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesRocketLoaderValue string

const (
	ZoneSettingGetResponseZonesRocketLoaderValueOn  ZoneSettingGetResponseZonesRocketLoaderValue = "on"
	ZoneSettingGetResponseZonesRocketLoaderValueOff ZoneSettingGetResponseZonesRocketLoaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesRocketLoaderEditable bool

const (
	ZoneSettingGetResponseZonesRocketLoaderEditableTrue  ZoneSettingGetResponseZonesRocketLoaderEditable = true
	ZoneSettingGetResponseZonesRocketLoaderEditableFalse ZoneSettingGetResponseZonesRocketLoaderEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationValue `json:"value,required"`
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

func (r zoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimization) implementsZoneSettingGetResponse() {
}

// ID of the zone setting.
type ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationID string

const (
	ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationValue struct {
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
	WpPlugin bool                                                                     `json:"wp_plugin,required"`
	JSON     zoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationValueJSON `json:"-"`
}

// zoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationValueJSON
// contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationValue]
type zoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationValueJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Cloudflare security header for a zone.
type ZoneSettingGetResponseZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID ZoneSettingGetResponseZonesSecurityHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesSecurityHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesSecurityHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesSecurityHeaderJSON `json:"-"`
}

// zoneSettingGetResponseZonesSecurityHeaderJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesSecurityHeader]
type zoneSettingGetResponseZonesSecurityHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesSecurityHeader) implementsZoneSettingGetResponse() {}

// ID of the zone's security header.
type ZoneSettingGetResponseZonesSecurityHeaderID string

const (
	ZoneSettingGetResponseZonesSecurityHeaderIDSecurityHeader ZoneSettingGetResponseZonesSecurityHeaderID = "security_header"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZoneSettingGetResponseZonesSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zoneSettingGetResponseZonesSecurityHeaderValueJSON                    `json:"-"`
}

// zoneSettingGetResponseZonesSecurityHeaderValueJSON contains the JSON metadata
// for the struct [ZoneSettingGetResponseZonesSecurityHeaderValue]
type zoneSettingGetResponseZonesSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesSecurityHeaderValueJSON) RawJSON() string {
	return r.raw
}

// Strict Transport Security.
type ZoneSettingGetResponseZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                      `json:"nosniff"`
	JSON    zoneSettingGetResponseZonesSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// zoneSettingGetResponseZonesSecurityHeaderValueStrictTransportSecurityJSON
// contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesSecurityHeaderValueStrictTransportSecurity]
type zoneSettingGetResponseZonesSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesSecurityHeaderValueStrictTransportSecurityJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesSecurityHeaderEditable bool

const (
	ZoneSettingGetResponseZonesSecurityHeaderEditableTrue  ZoneSettingGetResponseZonesSecurityHeaderEditable = true
	ZoneSettingGetResponseZonesSecurityHeaderEditableFalse ZoneSettingGetResponseZonesSecurityHeaderEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingGetResponseZonesSecurityLevel struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesSecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesSecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesSecurityLevelJSON `json:"-"`
}

// zoneSettingGetResponseZonesSecurityLevelJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesSecurityLevel]
type zoneSettingGetResponseZonesSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesSecurityLevelJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesSecurityLevel) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesSecurityLevelID string

const (
	ZoneSettingGetResponseZonesSecurityLevelIDSecurityLevel ZoneSettingGetResponseZonesSecurityLevelID = "security_level"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesSecurityLevelValue string

const (
	ZoneSettingGetResponseZonesSecurityLevelValueOff            ZoneSettingGetResponseZonesSecurityLevelValue = "off"
	ZoneSettingGetResponseZonesSecurityLevelValueEssentiallyOff ZoneSettingGetResponseZonesSecurityLevelValue = "essentially_off"
	ZoneSettingGetResponseZonesSecurityLevelValueLow            ZoneSettingGetResponseZonesSecurityLevelValue = "low"
	ZoneSettingGetResponseZonesSecurityLevelValueMedium         ZoneSettingGetResponseZonesSecurityLevelValue = "medium"
	ZoneSettingGetResponseZonesSecurityLevelValueHigh           ZoneSettingGetResponseZonesSecurityLevelValue = "high"
	ZoneSettingGetResponseZonesSecurityLevelValueUnderAttack    ZoneSettingGetResponseZonesSecurityLevelValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesSecurityLevelEditable bool

const (
	ZoneSettingGetResponseZonesSecurityLevelEditableTrue  ZoneSettingGetResponseZonesSecurityLevelEditable = true
	ZoneSettingGetResponseZonesSecurityLevelEditableFalse ZoneSettingGetResponseZonesSecurityLevelEditable = false
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
type ZoneSettingGetResponseZonesServerSideExclude struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesServerSideExcludeID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesServerSideExcludeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesServerSideExcludeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesServerSideExcludeJSON `json:"-"`
}

// zoneSettingGetResponseZonesServerSideExcludeJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesServerSideExclude]
type zoneSettingGetResponseZonesServerSideExcludeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesServerSideExcludeJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesServerSideExclude) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesServerSideExcludeID string

const (
	ZoneSettingGetResponseZonesServerSideExcludeIDServerSideExclude ZoneSettingGetResponseZonesServerSideExcludeID = "server_side_exclude"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesServerSideExcludeValue string

const (
	ZoneSettingGetResponseZonesServerSideExcludeValueOn  ZoneSettingGetResponseZonesServerSideExcludeValue = "on"
	ZoneSettingGetResponseZonesServerSideExcludeValueOff ZoneSettingGetResponseZonesServerSideExcludeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesServerSideExcludeEditable bool

const (
	ZoneSettingGetResponseZonesServerSideExcludeEditableTrue  ZoneSettingGetResponseZonesServerSideExcludeEditable = true
	ZoneSettingGetResponseZonesServerSideExcludeEditableFalse ZoneSettingGetResponseZonesServerSideExcludeEditable = false
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

func (r zoneSettingGetResponseZonesSha1SupportJSON) RawJSON() string {
	return r.raw
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

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingGetResponseZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesSortQueryStringForCacheID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesSortQueryStringForCacheValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesSortQueryStringForCacheJSON `json:"-"`
}

// zoneSettingGetResponseZonesSortQueryStringForCacheJSON contains the JSON
// metadata for the struct [ZoneSettingGetResponseZonesSortQueryStringForCache]
type zoneSettingGetResponseZonesSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesSortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesSortQueryStringForCache) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesSortQueryStringForCacheID string

const (
	ZoneSettingGetResponseZonesSortQueryStringForCacheIDSortQueryStringForCache ZoneSettingGetResponseZonesSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesSortQueryStringForCacheValue string

const (
	ZoneSettingGetResponseZonesSortQueryStringForCacheValueOn  ZoneSettingGetResponseZonesSortQueryStringForCacheValue = "on"
	ZoneSettingGetResponseZonesSortQueryStringForCacheValueOff ZoneSettingGetResponseZonesSortQueryStringForCacheValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesSortQueryStringForCacheEditable bool

const (
	ZoneSettingGetResponseZonesSortQueryStringForCacheEditableTrue  ZoneSettingGetResponseZonesSortQueryStringForCacheEditable = true
	ZoneSettingGetResponseZonesSortQueryStringForCacheEditableFalse ZoneSettingGetResponseZonesSortQueryStringForCacheEditable = false
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
type ZoneSettingGetResponseZonesSSL struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesSSLID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesSSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesSSLJSON `json:"-"`
}

// zoneSettingGetResponseZonesSSLJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesSSL]
type zoneSettingGetResponseZonesSSLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesSSLJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesSSL) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesSSLID string

const (
	ZoneSettingGetResponseZonesSSLIDSSL ZoneSettingGetResponseZonesSSLID = "ssl"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesSSLValue string

const (
	ZoneSettingGetResponseZonesSSLValueOff      ZoneSettingGetResponseZonesSSLValue = "off"
	ZoneSettingGetResponseZonesSSLValueFlexible ZoneSettingGetResponseZonesSSLValue = "flexible"
	ZoneSettingGetResponseZonesSSLValueFull     ZoneSettingGetResponseZonesSSLValue = "full"
	ZoneSettingGetResponseZonesSSLValueStrict   ZoneSettingGetResponseZonesSSLValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesSSLEditable bool

const (
	ZoneSettingGetResponseZonesSSLEditableTrue  ZoneSettingGetResponseZonesSSLEditable = true
	ZoneSettingGetResponseZonesSSLEditableFalse ZoneSettingGetResponseZonesSSLEditable = false
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingGetResponseZonesSSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingGetResponseZonesSSLRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                          `json:"enabled"`
	JSON    zoneSettingGetResponseZonesSSLRecommenderJSON `json:"-"`
}

// zoneSettingGetResponseZonesSSLRecommenderJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesSSLRecommender]
type zoneSettingGetResponseZonesSSLRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesSSLRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesSSLRecommenderJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesSSLRecommender) implementsZoneSettingGetResponse() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingGetResponseZonesSSLRecommenderID string

const (
	ZoneSettingGetResponseZonesSSLRecommenderIDSSLRecommender ZoneSettingGetResponseZonesSSLRecommenderID = "ssl_recommender"
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

func (r zoneSettingGetResponseZonesTls1_2OnlyJSON) RawJSON() string {
	return r.raw
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

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingGetResponseZonesTLS1_3 struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesTLS1_3ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesTLS1_3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesTLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesTls1_3JSON `json:"-"`
}

// zoneSettingGetResponseZonesTls1_3JSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesTLS1_3]
type zoneSettingGetResponseZonesTls1_3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesTLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesTls1_3JSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesTLS1_3) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesTLS1_3ID string

const (
	ZoneSettingGetResponseZonesTLS1_3IDTLS1_3 ZoneSettingGetResponseZonesTLS1_3ID = "tls_1_3"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesTLS1_3Value string

const (
	ZoneSettingGetResponseZonesTLS1_3ValueOn  ZoneSettingGetResponseZonesTLS1_3Value = "on"
	ZoneSettingGetResponseZonesTLS1_3ValueOff ZoneSettingGetResponseZonesTLS1_3Value = "off"
	ZoneSettingGetResponseZonesTLS1_3ValueZrt ZoneSettingGetResponseZonesTLS1_3Value = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesTLS1_3Editable bool

const (
	ZoneSettingGetResponseZonesTLS1_3EditableTrue  ZoneSettingGetResponseZonesTLS1_3Editable = true
	ZoneSettingGetResponseZonesTLS1_3EditableFalse ZoneSettingGetResponseZonesTLS1_3Editable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingGetResponseZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesTLSClientAuthID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesTLSClientAuthValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesTLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesTLSClientAuthJSON `json:"-"`
}

// zoneSettingGetResponseZonesTLSClientAuthJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesTLSClientAuth]
type zoneSettingGetResponseZonesTLSClientAuthJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesTLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesTLSClientAuthJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesTLSClientAuth) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesTLSClientAuthID string

const (
	ZoneSettingGetResponseZonesTLSClientAuthIDTLSClientAuth ZoneSettingGetResponseZonesTLSClientAuthID = "tls_client_auth"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesTLSClientAuthValue string

const (
	ZoneSettingGetResponseZonesTLSClientAuthValueOn  ZoneSettingGetResponseZonesTLSClientAuthValue = "on"
	ZoneSettingGetResponseZonesTLSClientAuthValueOff ZoneSettingGetResponseZonesTLSClientAuthValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesTLSClientAuthEditable bool

const (
	ZoneSettingGetResponseZonesTLSClientAuthEditableTrue  ZoneSettingGetResponseZonesTLSClientAuthEditable = true
	ZoneSettingGetResponseZonesTLSClientAuthEditableFalse ZoneSettingGetResponseZonesTLSClientAuthEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingGetResponseZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesTrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesTrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesTrueClientIPHeaderJSON `json:"-"`
}

// zoneSettingGetResponseZonesTrueClientIPHeaderJSON contains the JSON metadata for
// the struct [ZoneSettingGetResponseZonesTrueClientIPHeader]
type zoneSettingGetResponseZonesTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesTrueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesTrueClientIPHeader) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesTrueClientIPHeaderID string

const (
	ZoneSettingGetResponseZonesTrueClientIPHeaderIDTrueClientIPHeader ZoneSettingGetResponseZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesTrueClientIPHeaderValue string

const (
	ZoneSettingGetResponseZonesTrueClientIPHeaderValueOn  ZoneSettingGetResponseZonesTrueClientIPHeaderValue = "on"
	ZoneSettingGetResponseZonesTrueClientIPHeaderValueOff ZoneSettingGetResponseZonesTrueClientIPHeaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesTrueClientIPHeaderEditable bool

const (
	ZoneSettingGetResponseZonesTrueClientIPHeaderEditableTrue  ZoneSettingGetResponseZonesTrueClientIPHeaderEditable = true
	ZoneSettingGetResponseZonesTrueClientIPHeaderEditableFalse ZoneSettingGetResponseZonesTrueClientIPHeaderEditable = false
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
type ZoneSettingGetResponseZonesWAF struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesWAFID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesWAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesWAFJSON `json:"-"`
}

// zoneSettingGetResponseZonesWAFJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesWAF]
type zoneSettingGetResponseZonesWAFJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesWAFJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesWAF) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesWAFID string

const (
	ZoneSettingGetResponseZonesWAFIDWAF ZoneSettingGetResponseZonesWAFID = "waf"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesWAFValue string

const (
	ZoneSettingGetResponseZonesWAFValueOn  ZoneSettingGetResponseZonesWAFValue = "on"
	ZoneSettingGetResponseZonesWAFValueOff ZoneSettingGetResponseZonesWAFValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesWAFEditable bool

const (
	ZoneSettingGetResponseZonesWAFEditableTrue  ZoneSettingGetResponseZonesWAFEditable = true
	ZoneSettingGetResponseZonesWAFEditableFalse ZoneSettingGetResponseZonesWAFEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingGetResponseZonesWebp struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesWebpID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesWebpValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesWebpJSON `json:"-"`
}

// zoneSettingGetResponseZonesWebpJSON contains the JSON metadata for the struct
// [ZoneSettingGetResponseZonesWebp]
type zoneSettingGetResponseZonesWebpJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesWebpJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesWebp) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesWebpID string

const (
	ZoneSettingGetResponseZonesWebpIDWebp ZoneSettingGetResponseZonesWebpID = "webp"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesWebpValue string

const (
	ZoneSettingGetResponseZonesWebpValueOff ZoneSettingGetResponseZonesWebpValue = "off"
	ZoneSettingGetResponseZonesWebpValueOn  ZoneSettingGetResponseZonesWebpValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesWebpEditable bool

const (
	ZoneSettingGetResponseZonesWebpEditableTrue  ZoneSettingGetResponseZonesWebpEditable = true
	ZoneSettingGetResponseZonesWebpEditableFalse ZoneSettingGetResponseZonesWebpEditable = false
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZoneSettingGetResponseZonesWebsockets struct {
	// ID of the zone setting.
	ID ZoneSettingGetResponseZonesWebsocketsID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingGetResponseZonesWebsocketsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingGetResponseZonesWebsocketsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingGetResponseZonesWebsocketsJSON `json:"-"`
}

// zoneSettingGetResponseZonesWebsocketsJSON contains the JSON metadata for the
// struct [ZoneSettingGetResponseZonesWebsockets]
type zoneSettingGetResponseZonesWebsocketsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetResponseZonesWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetResponseZonesWebsocketsJSON) RawJSON() string {
	return r.raw
}

func (r ZoneSettingGetResponseZonesWebsockets) implementsZoneSettingGetResponse() {}

// ID of the zone setting.
type ZoneSettingGetResponseZonesWebsocketsID string

const (
	ZoneSettingGetResponseZonesWebsocketsIDWebsockets ZoneSettingGetResponseZonesWebsocketsID = "websockets"
)

// Current value of the zone setting.
type ZoneSettingGetResponseZonesWebsocketsValue string

const (
	ZoneSettingGetResponseZonesWebsocketsValueOff ZoneSettingGetResponseZonesWebsocketsValue = "off"
	ZoneSettingGetResponseZonesWebsocketsValueOn  ZoneSettingGetResponseZonesWebsocketsValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingGetResponseZonesWebsocketsEditable bool

const (
	ZoneSettingGetResponseZonesWebsocketsEditableTrue  ZoneSettingGetResponseZonesWebsocketsEditable = true
	ZoneSettingGetResponseZonesWebsocketsEditableFalse ZoneSettingGetResponseZonesWebsocketsEditable = false
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
// [ZoneSettingEditParamsItemsZonesCNAMEFlattening],
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
type ZoneSettingEditParamsItemsZonesCNAMEFlattening struct {
	// How to flatten the cname destination.
	ID param.Field[ZoneSettingEditParamsItemsZonesCNAMEFlatteningID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesCNAMEFlatteningValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsZonesCNAMEFlattening) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesCNAMEFlattening) implementsZoneSettingEditParamsItem() {}

// How to flatten the cname destination.
type ZoneSettingEditParamsItemsZonesCNAMEFlatteningID string

const (
	ZoneSettingEditParamsItemsZonesCNAMEFlatteningIDCNAMEFlattening ZoneSettingEditParamsItemsZonesCNAMEFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type ZoneSettingEditParamsItemsZonesCNAMEFlatteningValue string

const (
	ZoneSettingEditParamsItemsZonesCNAMEFlatteningValueFlattenAtRoot ZoneSettingEditParamsItemsZonesCNAMEFlatteningValue = "flatten_at_root"
	ZoneSettingEditParamsItemsZonesCNAMEFlatteningValueFlattenAll    ZoneSettingEditParamsItemsZonesCNAMEFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesCNAMEFlatteningEditable bool

const (
	ZoneSettingEditParamsItemsZonesCNAMEFlatteningEditableTrue  ZoneSettingEditParamsItemsZonesCNAMEFlatteningEditable = true
	ZoneSettingEditParamsItemsZonesCNAMEFlatteningEditableFalse ZoneSettingEditParamsItemsZonesCNAMEFlatteningEditable = false
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
	StripURI param.Field[bool] `json:"strip_uri"`
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

func (r zoneSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zoneSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
