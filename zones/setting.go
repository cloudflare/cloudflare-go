// File generated from our OpenAPI spec by Stainless.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSettingService] method instead.
type SettingService struct {
	Options                       []option.RequestOption
	ZeroRTT                       *SettingZeroRTTService
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
	r.ZeroRTT = NewSettingZeroRTTService(opts...)
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

// Edit settings for a zone.
func (r *SettingService) Edit(ctx context.Context, params SettingEditParams, opts ...option.RequestOption) (res *[]SettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Available settings for your user in relation to a zone.
func (r *SettingService) Get(ctx context.Context, query SettingGetParams, opts ...option.RequestOption) (res *[]SettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingGetResponseEnvelope
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
// Union satisfied by [zones.SettingEditResponseZones0rtt],
// [zones.SettingEditResponseZonesAdvancedDDOS],
// [zones.SettingEditResponseZonesAlwaysOnline],
// [zones.SettingEditResponseZonesAlwaysUseHTTPS],
// [zones.SettingEditResponseZonesAutomaticHTTPSRewrites],
// [zones.SettingEditResponseZonesBrotli],
// [zones.SettingEditResponseZonesBrowserCacheTTL],
// [zones.SettingEditResponseZonesBrowserCheck],
// [zones.SettingEditResponseZonesCacheLevel],
// [zones.SettingEditResponseZonesChallengeTTL],
// [zones.SettingEditResponseZonesCiphers],
// [zones.SettingEditResponseZonesCNAMEFlattening],
// [zones.SettingEditResponseZonesDevelopmentMode],
// [zones.SettingEditResponseZonesEarlyHints],
// [zones.SettingEditResponseZonesEdgeCacheTTL],
// [zones.SettingEditResponseZonesEmailObfuscation],
// [zones.SettingEditResponseZonesH2Prioritization],
// [zones.SettingEditResponseZonesHotlinkProtection],
// [zones.SettingEditResponseZonesHTTP2], [zones.SettingEditResponseZonesHTTP3],
// [zones.SettingEditResponseZonesImageResizing],
// [zones.SettingEditResponseZonesIPGeolocation],
// [zones.SettingEditResponseZonesIPV6], [zones.SettingEditResponseZonesMaxUpload],
// [zones.SettingEditResponseZonesMinTLSVersion],
// [zones.SettingEditResponseZonesMinify], [zones.SettingEditResponseZonesMirage],
// [zones.SettingEditResponseZonesMobileRedirect],
// [zones.SettingEditResponseZonesNEL],
// [zones.SettingEditResponseZonesOpportunisticEncryption],
// [zones.SettingEditResponseZonesOpportunisticOnion],
// [zones.SettingEditResponseZonesOrangeToOrange],
// [zones.SettingEditResponseZonesOriginErrorPagePassThru],
// [zones.SettingEditResponseZonesPolish],
// [zones.SettingEditResponseZonesPrefetchPreload],
// [zones.SettingEditResponseZonesProxyReadTimeout],
// [zones.SettingEditResponseZonesPseudoIPV4],
// [zones.SettingEditResponseZonesResponseBuffering],
// [zones.SettingEditResponseZonesRocketLoader],
// [zones.SettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.SettingEditResponseZonesSecurityHeader],
// [zones.SettingEditResponseZonesSecurityLevel],
// [zones.SettingEditResponseZonesServerSideExclude],
// [zones.SettingEditResponseZonesSha1Support],
// [zones.SettingEditResponseZonesSortQueryStringForCache],
// [zones.SettingEditResponseZonesSSL],
// [zones.SettingEditResponseZonesSSLRecommender],
// [zones.SettingEditResponseZonesTLS1_2Only],
// [zones.SettingEditResponseZonesTLS1_3],
// [zones.SettingEditResponseZonesTLSClientAuth],
// [zones.SettingEditResponseZonesTrueClientIPHeader],
// [zones.SettingEditResponseZonesWAF], [zones.SettingEditResponseZonesWebp] or
// [zones.SettingEditResponseZonesWebsockets].
type SettingEditResponse interface {
	implementsZonesSettingEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZones0rtt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesAdvancedDDOS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesAlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesBrotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesBrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesCacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesCiphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesCNAMEFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesDevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesEarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesEmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesH2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesHotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesHTTP2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesHTTP3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesIPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesIPV6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesMinTLSVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesMinify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesMobileRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesNEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesOpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesOpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesOrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesPolish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesPrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesPseudoIPV4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesResponseBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesRocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSecurityHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesServerSideExclude{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSSL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSSLRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesTLS1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesTLS1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesTLSClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesWAF{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesWebp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesWebsockets{}),
		},
	)
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

func (r settingEditResponseZones0rttJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZones0rtt) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesAdvancedDDOSJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesAdvancedDDOS) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesAlwaysOnlineJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesAlwaysOnline) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesAlwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesAlwaysUseHTTPS) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesAutomaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesAutomaticHTTPSRewrites) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesBrotliJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesBrotli) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesBrowserCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesBrowserCacheTTL) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesBrowserCheckJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesBrowserCheck) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesCacheLevelJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesCacheLevel) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesChallengeTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesChallengeTTL) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesCiphersJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesCiphers) implementsZonesSettingEditResponse() {}

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

// Current value of the zone setting.
type SettingEditResponseZonesCNAMEFlatteningValue string

const (
	SettingEditResponseZonesCNAMEFlatteningValueFlattenAtRoot SettingEditResponseZonesCNAMEFlatteningValue = "flatten_at_root"
	SettingEditResponseZonesCNAMEFlatteningValueFlattenAll    SettingEditResponseZonesCNAMEFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesCNAMEFlatteningEditable bool

const (
	SettingEditResponseZonesCNAMEFlatteningEditableTrue  SettingEditResponseZonesCNAMEFlatteningEditable = true
	SettingEditResponseZonesCNAMEFlatteningEditableFalse SettingEditResponseZonesCNAMEFlatteningEditable = false
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

func (r settingEditResponseZonesDevelopmentModeJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesDevelopmentMode) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesEarlyHintsJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesEarlyHints) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesEdgeCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesEdgeCacheTTL) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesEmailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesEmailObfuscation) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesH2PrioritizationJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesH2Prioritization) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesHotlinkProtectionJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesHotlinkProtection) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesHTTP2JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesHTTP2) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesHTTP3JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesHTTP3) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesImageResizingJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesImageResizing) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesIPGeolocationJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesIPGeolocation) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesIPV6JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesIPV6) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesMaxUploadJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesMaxUpload) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesMinTLSVersionJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesMinTLSVersion) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesMinifyJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesMinify) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesMinifyValueJSON) RawJSON() string {
	return r.raw
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

func (r settingEditResponseZonesMirageJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesMirage) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesMobileRedirectJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesMobileRedirect) implementsZonesSettingEditResponse() {}

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
	StripURI bool                                            `json:"strip_uri"`
	JSON     settingEditResponseZonesMobileRedirectValueJSON `json:"-"`
}

// settingEditResponseZonesMobileRedirectValueJSON contains the JSON metadata for
// the struct [SettingEditResponseZonesMobileRedirectValue]
type settingEditResponseZonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripURI        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingEditResponseZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseZonesMobileRedirectValueJSON) RawJSON() string {
	return r.raw
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

func (r settingEditResponseZonesNELJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesNEL) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesNELValueJSON) RawJSON() string {
	return r.raw
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

func (r settingEditResponseZonesOpportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesOpportunisticEncryption) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesOpportunisticOnionJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesOpportunisticOnion) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesOrangeToOrangeJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesOrangeToOrange) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesOriginErrorPagePassThruJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesOriginErrorPagePassThru) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesPolishJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesPolish) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesPrefetchPreloadJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesPrefetchPreload) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesProxyReadTimeoutJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesProxyReadTimeout) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesPseudoIPV4JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesPseudoIPV4) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesResponseBufferingJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesResponseBuffering) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesRocketLoaderJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesRocketLoader) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesSchemasAutomaticPlatformOptimizationValueJSON) RawJSON() string {
	return r.raw
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

func (r settingEditResponseZonesSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSecurityHeader) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesSecurityHeaderValueJSON) RawJSON() string {
	return r.raw
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

func (r settingEditResponseZonesSecurityHeaderValueStrictTransportSecurityJSON) RawJSON() string {
	return r.raw
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

func (r settingEditResponseZonesSecurityLevelJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSecurityLevel) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesServerSideExcludeJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesServerSideExclude) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesSha1SupportJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSha1Support) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesSortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSortQueryStringForCache) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesSSLJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSSL) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesSSLRecommenderJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesSSLRecommender) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesTls1_2OnlyJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesTLS1_2Only) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesTls1_3JSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesTLS1_3) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesTLSClientAuthJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesTLSClientAuth) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesTrueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesTrueClientIPHeader) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesWAFJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesWAF) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesWebpJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesWebp) implementsZonesSettingEditResponse() {}

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

func (r settingEditResponseZonesWebsocketsJSON) RawJSON() string {
	return r.raw
}

func (r SettingEditResponseZonesWebsockets) implementsZonesSettingEditResponse() {}

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

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [zones.SettingGetResponseZones0rtt],
// [zones.SettingGetResponseZonesAdvancedDDOS],
// [zones.SettingGetResponseZonesAlwaysOnline],
// [zones.SettingGetResponseZonesAlwaysUseHTTPS],
// [zones.SettingGetResponseZonesAutomaticHTTPSRewrites],
// [zones.SettingGetResponseZonesBrotli],
// [zones.SettingGetResponseZonesBrowserCacheTTL],
// [zones.SettingGetResponseZonesBrowserCheck],
// [zones.SettingGetResponseZonesCacheLevel],
// [zones.SettingGetResponseZonesChallengeTTL],
// [zones.SettingGetResponseZonesCiphers],
// [zones.SettingGetResponseZonesCNAMEFlattening],
// [zones.SettingGetResponseZonesDevelopmentMode],
// [zones.SettingGetResponseZonesEarlyHints],
// [zones.SettingGetResponseZonesEdgeCacheTTL],
// [zones.SettingGetResponseZonesEmailObfuscation],
// [zones.SettingGetResponseZonesH2Prioritization],
// [zones.SettingGetResponseZonesHotlinkProtection],
// [zones.SettingGetResponseZonesHTTP2], [zones.SettingGetResponseZonesHTTP3],
// [zones.SettingGetResponseZonesImageResizing],
// [zones.SettingGetResponseZonesIPGeolocation],
// [zones.SettingGetResponseZonesIPV6], [zones.SettingGetResponseZonesMaxUpload],
// [zones.SettingGetResponseZonesMinTLSVersion],
// [zones.SettingGetResponseZonesMinify], [zones.SettingGetResponseZonesMirage],
// [zones.SettingGetResponseZonesMobileRedirect],
// [zones.SettingGetResponseZonesNEL],
// [zones.SettingGetResponseZonesOpportunisticEncryption],
// [zones.SettingGetResponseZonesOpportunisticOnion],
// [zones.SettingGetResponseZonesOrangeToOrange],
// [zones.SettingGetResponseZonesOriginErrorPagePassThru],
// [zones.SettingGetResponseZonesPolish],
// [zones.SettingGetResponseZonesPrefetchPreload],
// [zones.SettingGetResponseZonesProxyReadTimeout],
// [zones.SettingGetResponseZonesPseudoIPV4],
// [zones.SettingGetResponseZonesResponseBuffering],
// [zones.SettingGetResponseZonesRocketLoader],
// [zones.SettingGetResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.SettingGetResponseZonesSecurityHeader],
// [zones.SettingGetResponseZonesSecurityLevel],
// [zones.SettingGetResponseZonesServerSideExclude],
// [zones.SettingGetResponseZonesSha1Support],
// [zones.SettingGetResponseZonesSortQueryStringForCache],
// [zones.SettingGetResponseZonesSSL],
// [zones.SettingGetResponseZonesSSLRecommender],
// [zones.SettingGetResponseZonesTLS1_2Only],
// [zones.SettingGetResponseZonesTLS1_3],
// [zones.SettingGetResponseZonesTLSClientAuth],
// [zones.SettingGetResponseZonesTrueClientIPHeader],
// [zones.SettingGetResponseZonesWAF], [zones.SettingGetResponseZonesWebp] or
// [zones.SettingGetResponseZonesWebsockets].
type SettingGetResponse interface {
	implementsZonesSettingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZones0rtt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesAdvancedDDOS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesAlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesBrotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesBrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesCacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesCiphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesCNAMEFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesDevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesEarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesEmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesH2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesHotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesHTTP2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesHTTP3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesIPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesIPV6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesMinTLSVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesMinify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesMobileRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesNEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesOpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesOpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesOrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesPolish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesPrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesPseudoIPV4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesResponseBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesRocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSecurityHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesServerSideExclude{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSSL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSSLRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesTLS1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesTLS1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesTLSClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesWAF{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesWebp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesWebsockets{}),
		},
	)
}

// 0-RTT session resumption enabled for this zone.
type SettingGetResponseZones0rtt struct {
	// ID of the zone setting.
	ID SettingGetResponseZones0rttID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZones0rttValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZones0rttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZones0rttJSON `json:"-"`
}

// settingGetResponseZones0rttJSON contains the JSON metadata for the struct
// [SettingGetResponseZones0rtt]
type settingGetResponseZones0rttJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZones0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZones0rttJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZones0rtt) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZones0rttID string

const (
	SettingGetResponseZones0rttID0rtt SettingGetResponseZones0rttID = "0rtt"
)

// Current value of the zone setting.
type SettingGetResponseZones0rttValue string

const (
	SettingGetResponseZones0rttValueOn  SettingGetResponseZones0rttValue = "on"
	SettingGetResponseZones0rttValueOff SettingGetResponseZones0rttValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZones0rttEditable bool

const (
	SettingGetResponseZones0rttEditableTrue  SettingGetResponseZones0rttEditable = true
	SettingGetResponseZones0rttEditableFalse SettingGetResponseZones0rttEditable = false
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type SettingGetResponseZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesAdvancedDDOSID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesAdvancedDDOSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesAdvancedDDOSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesAdvancedDDOSJSON `json:"-"`
}

// settingGetResponseZonesAdvancedDDOSJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesAdvancedDDOS]
type settingGetResponseZonesAdvancedDDOSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesAdvancedDDOS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesAdvancedDDOSJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesAdvancedDDOS) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesAdvancedDDOSID string

const (
	SettingGetResponseZonesAdvancedDDOSIDAdvancedDDOS SettingGetResponseZonesAdvancedDDOSID = "advanced_ddos"
)

// Current value of the zone setting.
type SettingGetResponseZonesAdvancedDDOSValue string

const (
	SettingGetResponseZonesAdvancedDDOSValueOn  SettingGetResponseZonesAdvancedDDOSValue = "on"
	SettingGetResponseZonesAdvancedDDOSValueOff SettingGetResponseZonesAdvancedDDOSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesAdvancedDDOSEditable bool

const (
	SettingGetResponseZonesAdvancedDDOSEditableTrue  SettingGetResponseZonesAdvancedDDOSEditable = true
	SettingGetResponseZonesAdvancedDDOSEditableFalse SettingGetResponseZonesAdvancedDDOSEditable = false
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type SettingGetResponseZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesAlwaysOnlineID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesAlwaysOnlineValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesAlwaysOnlineJSON `json:"-"`
}

// settingGetResponseZonesAlwaysOnlineJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesAlwaysOnline]
type settingGetResponseZonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesAlwaysOnlineJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesAlwaysOnline) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesAlwaysOnlineID string

const (
	SettingGetResponseZonesAlwaysOnlineIDAlwaysOnline SettingGetResponseZonesAlwaysOnlineID = "always_online"
)

// Current value of the zone setting.
type SettingGetResponseZonesAlwaysOnlineValue string

const (
	SettingGetResponseZonesAlwaysOnlineValueOn  SettingGetResponseZonesAlwaysOnlineValue = "on"
	SettingGetResponseZonesAlwaysOnlineValueOff SettingGetResponseZonesAlwaysOnlineValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesAlwaysOnlineEditable bool

const (
	SettingGetResponseZonesAlwaysOnlineEditableTrue  SettingGetResponseZonesAlwaysOnlineEditable = true
	SettingGetResponseZonesAlwaysOnlineEditableFalse SettingGetResponseZonesAlwaysOnlineEditable = false
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type SettingGetResponseZonesAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesAlwaysUseHTTPSID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesAlwaysUseHTTPSValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesAlwaysUseHTTPSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesAlwaysUseHTTPSJSON `json:"-"`
}

// settingGetResponseZonesAlwaysUseHTTPSJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesAlwaysUseHTTPS]
type settingGetResponseZonesAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesAlwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesAlwaysUseHTTPS) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesAlwaysUseHTTPSID string

const (
	SettingGetResponseZonesAlwaysUseHTTPSIDAlwaysUseHTTPS SettingGetResponseZonesAlwaysUseHTTPSID = "always_use_https"
)

// Current value of the zone setting.
type SettingGetResponseZonesAlwaysUseHTTPSValue string

const (
	SettingGetResponseZonesAlwaysUseHTTPSValueOn  SettingGetResponseZonesAlwaysUseHTTPSValue = "on"
	SettingGetResponseZonesAlwaysUseHTTPSValueOff SettingGetResponseZonesAlwaysUseHTTPSValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesAlwaysUseHTTPSEditable bool

const (
	SettingGetResponseZonesAlwaysUseHTTPSEditableTrue  SettingGetResponseZonesAlwaysUseHTTPSEditable = true
	SettingGetResponseZonesAlwaysUseHTTPSEditableFalse SettingGetResponseZonesAlwaysUseHTTPSEditable = false
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type SettingGetResponseZonesAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesAutomaticHTTPSRewritesID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesAutomaticHTTPSRewritesValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesAutomaticHTTPSRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesAutomaticHTTPSRewritesJSON `json:"-"`
}

// settingGetResponseZonesAutomaticHTTPSRewritesJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesAutomaticHTTPSRewrites]
type settingGetResponseZonesAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesAutomaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesAutomaticHTTPSRewrites) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesAutomaticHTTPSRewritesID string

const (
	SettingGetResponseZonesAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites SettingGetResponseZonesAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

// Current value of the zone setting.
type SettingGetResponseZonesAutomaticHTTPSRewritesValue string

const (
	SettingGetResponseZonesAutomaticHTTPSRewritesValueOn  SettingGetResponseZonesAutomaticHTTPSRewritesValue = "on"
	SettingGetResponseZonesAutomaticHTTPSRewritesValueOff SettingGetResponseZonesAutomaticHTTPSRewritesValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesAutomaticHTTPSRewritesEditable bool

const (
	SettingGetResponseZonesAutomaticHTTPSRewritesEditableTrue  SettingGetResponseZonesAutomaticHTTPSRewritesEditable = true
	SettingGetResponseZonesAutomaticHTTPSRewritesEditableFalse SettingGetResponseZonesAutomaticHTTPSRewritesEditable = false
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type SettingGetResponseZonesBrotli struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesBrotliID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesBrotliValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesBrotliJSON `json:"-"`
}

// settingGetResponseZonesBrotliJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesBrotli]
type settingGetResponseZonesBrotliJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesBrotliJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesBrotli) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesBrotliID string

const (
	SettingGetResponseZonesBrotliIDBrotli SettingGetResponseZonesBrotliID = "brotli"
)

// Current value of the zone setting.
type SettingGetResponseZonesBrotliValue string

const (
	SettingGetResponseZonesBrotliValueOff SettingGetResponseZonesBrotliValue = "off"
	SettingGetResponseZonesBrotliValueOn  SettingGetResponseZonesBrotliValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesBrotliEditable bool

const (
	SettingGetResponseZonesBrotliEditableTrue  SettingGetResponseZonesBrotliEditable = true
	SettingGetResponseZonesBrotliEditableFalse SettingGetResponseZonesBrotliEditable = false
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type SettingGetResponseZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesBrowserCacheTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesBrowserCacheTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesBrowserCacheTTLJSON `json:"-"`
}

// settingGetResponseZonesBrowserCacheTTLJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesBrowserCacheTTL]
type settingGetResponseZonesBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesBrowserCacheTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesBrowserCacheTTL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesBrowserCacheTTLID string

const (
	SettingGetResponseZonesBrowserCacheTTLIDBrowserCacheTTL SettingGetResponseZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Current value of the zone setting.
type SettingGetResponseZonesBrowserCacheTTLValue float64

const (
	SettingGetResponseZonesBrowserCacheTTLValue0        SettingGetResponseZonesBrowserCacheTTLValue = 0
	SettingGetResponseZonesBrowserCacheTTLValue30       SettingGetResponseZonesBrowserCacheTTLValue = 30
	SettingGetResponseZonesBrowserCacheTTLValue60       SettingGetResponseZonesBrowserCacheTTLValue = 60
	SettingGetResponseZonesBrowserCacheTTLValue120      SettingGetResponseZonesBrowserCacheTTLValue = 120
	SettingGetResponseZonesBrowserCacheTTLValue300      SettingGetResponseZonesBrowserCacheTTLValue = 300
	SettingGetResponseZonesBrowserCacheTTLValue1200     SettingGetResponseZonesBrowserCacheTTLValue = 1200
	SettingGetResponseZonesBrowserCacheTTLValue1800     SettingGetResponseZonesBrowserCacheTTLValue = 1800
	SettingGetResponseZonesBrowserCacheTTLValue3600     SettingGetResponseZonesBrowserCacheTTLValue = 3600
	SettingGetResponseZonesBrowserCacheTTLValue7200     SettingGetResponseZonesBrowserCacheTTLValue = 7200
	SettingGetResponseZonesBrowserCacheTTLValue10800    SettingGetResponseZonesBrowserCacheTTLValue = 10800
	SettingGetResponseZonesBrowserCacheTTLValue14400    SettingGetResponseZonesBrowserCacheTTLValue = 14400
	SettingGetResponseZonesBrowserCacheTTLValue18000    SettingGetResponseZonesBrowserCacheTTLValue = 18000
	SettingGetResponseZonesBrowserCacheTTLValue28800    SettingGetResponseZonesBrowserCacheTTLValue = 28800
	SettingGetResponseZonesBrowserCacheTTLValue43200    SettingGetResponseZonesBrowserCacheTTLValue = 43200
	SettingGetResponseZonesBrowserCacheTTLValue57600    SettingGetResponseZonesBrowserCacheTTLValue = 57600
	SettingGetResponseZonesBrowserCacheTTLValue72000    SettingGetResponseZonesBrowserCacheTTLValue = 72000
	SettingGetResponseZonesBrowserCacheTTLValue86400    SettingGetResponseZonesBrowserCacheTTLValue = 86400
	SettingGetResponseZonesBrowserCacheTTLValue172800   SettingGetResponseZonesBrowserCacheTTLValue = 172800
	SettingGetResponseZonesBrowserCacheTTLValue259200   SettingGetResponseZonesBrowserCacheTTLValue = 259200
	SettingGetResponseZonesBrowserCacheTTLValue345600   SettingGetResponseZonesBrowserCacheTTLValue = 345600
	SettingGetResponseZonesBrowserCacheTTLValue432000   SettingGetResponseZonesBrowserCacheTTLValue = 432000
	SettingGetResponseZonesBrowserCacheTTLValue691200   SettingGetResponseZonesBrowserCacheTTLValue = 691200
	SettingGetResponseZonesBrowserCacheTTLValue1382400  SettingGetResponseZonesBrowserCacheTTLValue = 1382400
	SettingGetResponseZonesBrowserCacheTTLValue2073600  SettingGetResponseZonesBrowserCacheTTLValue = 2073600
	SettingGetResponseZonesBrowserCacheTTLValue2678400  SettingGetResponseZonesBrowserCacheTTLValue = 2678400
	SettingGetResponseZonesBrowserCacheTTLValue5356800  SettingGetResponseZonesBrowserCacheTTLValue = 5356800
	SettingGetResponseZonesBrowserCacheTTLValue16070400 SettingGetResponseZonesBrowserCacheTTLValue = 16070400
	SettingGetResponseZonesBrowserCacheTTLValue31536000 SettingGetResponseZonesBrowserCacheTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesBrowserCacheTTLEditable bool

const (
	SettingGetResponseZonesBrowserCacheTTLEditableTrue  SettingGetResponseZonesBrowserCacheTTLEditable = true
	SettingGetResponseZonesBrowserCacheTTLEditableFalse SettingGetResponseZonesBrowserCacheTTLEditable = false
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type SettingGetResponseZonesBrowserCheck struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesBrowserCheckID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesBrowserCheckValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesBrowserCheckJSON `json:"-"`
}

// settingGetResponseZonesBrowserCheckJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesBrowserCheck]
type settingGetResponseZonesBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesBrowserCheckJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesBrowserCheck) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesBrowserCheckID string

const (
	SettingGetResponseZonesBrowserCheckIDBrowserCheck SettingGetResponseZonesBrowserCheckID = "browser_check"
)

// Current value of the zone setting.
type SettingGetResponseZonesBrowserCheckValue string

const (
	SettingGetResponseZonesBrowserCheckValueOn  SettingGetResponseZonesBrowserCheckValue = "on"
	SettingGetResponseZonesBrowserCheckValueOff SettingGetResponseZonesBrowserCheckValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesBrowserCheckEditable bool

const (
	SettingGetResponseZonesBrowserCheckEditableTrue  SettingGetResponseZonesBrowserCheckEditable = true
	SettingGetResponseZonesBrowserCheckEditableFalse SettingGetResponseZonesBrowserCheckEditable = false
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type SettingGetResponseZonesCacheLevel struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesCacheLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesCacheLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesCacheLevelJSON `json:"-"`
}

// settingGetResponseZonesCacheLevelJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesCacheLevel]
type settingGetResponseZonesCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesCacheLevelJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesCacheLevel) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesCacheLevelID string

const (
	SettingGetResponseZonesCacheLevelIDCacheLevel SettingGetResponseZonesCacheLevelID = "cache_level"
)

// Current value of the zone setting.
type SettingGetResponseZonesCacheLevelValue string

const (
	SettingGetResponseZonesCacheLevelValueAggressive SettingGetResponseZonesCacheLevelValue = "aggressive"
	SettingGetResponseZonesCacheLevelValueBasic      SettingGetResponseZonesCacheLevelValue = "basic"
	SettingGetResponseZonesCacheLevelValueSimplified SettingGetResponseZonesCacheLevelValue = "simplified"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesCacheLevelEditable bool

const (
	SettingGetResponseZonesCacheLevelEditableTrue  SettingGetResponseZonesCacheLevelEditable = true
	SettingGetResponseZonesCacheLevelEditableFalse SettingGetResponseZonesCacheLevelEditable = false
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type SettingGetResponseZonesChallengeTTL struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesChallengeTTLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesChallengeTTLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesChallengeTTLJSON `json:"-"`
}

// settingGetResponseZonesChallengeTTLJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesChallengeTTL]
type settingGetResponseZonesChallengeTTLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesChallengeTTLJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesChallengeTTL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesChallengeTTLID string

const (
	SettingGetResponseZonesChallengeTTLIDChallengeTTL SettingGetResponseZonesChallengeTTLID = "challenge_ttl"
)

// Current value of the zone setting.
type SettingGetResponseZonesChallengeTTLValue float64

const (
	SettingGetResponseZonesChallengeTTLValue300      SettingGetResponseZonesChallengeTTLValue = 300
	SettingGetResponseZonesChallengeTTLValue900      SettingGetResponseZonesChallengeTTLValue = 900
	SettingGetResponseZonesChallengeTTLValue1800     SettingGetResponseZonesChallengeTTLValue = 1800
	SettingGetResponseZonesChallengeTTLValue2700     SettingGetResponseZonesChallengeTTLValue = 2700
	SettingGetResponseZonesChallengeTTLValue3600     SettingGetResponseZonesChallengeTTLValue = 3600
	SettingGetResponseZonesChallengeTTLValue7200     SettingGetResponseZonesChallengeTTLValue = 7200
	SettingGetResponseZonesChallengeTTLValue10800    SettingGetResponseZonesChallengeTTLValue = 10800
	SettingGetResponseZonesChallengeTTLValue14400    SettingGetResponseZonesChallengeTTLValue = 14400
	SettingGetResponseZonesChallengeTTLValue28800    SettingGetResponseZonesChallengeTTLValue = 28800
	SettingGetResponseZonesChallengeTTLValue57600    SettingGetResponseZonesChallengeTTLValue = 57600
	SettingGetResponseZonesChallengeTTLValue86400    SettingGetResponseZonesChallengeTTLValue = 86400
	SettingGetResponseZonesChallengeTTLValue604800   SettingGetResponseZonesChallengeTTLValue = 604800
	SettingGetResponseZonesChallengeTTLValue2592000  SettingGetResponseZonesChallengeTTLValue = 2592000
	SettingGetResponseZonesChallengeTTLValue31536000 SettingGetResponseZonesChallengeTTLValue = 31536000
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesChallengeTTLEditable bool

const (
	SettingGetResponseZonesChallengeTTLEditableTrue  SettingGetResponseZonesChallengeTTLEditable = true
	SettingGetResponseZonesChallengeTTLEditableFalse SettingGetResponseZonesChallengeTTLEditable = false
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type SettingGetResponseZonesCiphers struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesCiphersID `json:"id,required"`
	// Current value of the zone setting.
	Value []string `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesCiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesCiphersJSON `json:"-"`
}

// settingGetResponseZonesCiphersJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesCiphers]
type settingGetResponseZonesCiphersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesCiphersJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesCiphers) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesCiphersID string

const (
	SettingGetResponseZonesCiphersIDCiphers SettingGetResponseZonesCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesCiphersEditable bool

const (
	SettingGetResponseZonesCiphersEditableTrue  SettingGetResponseZonesCiphersEditable = true
	SettingGetResponseZonesCiphersEditableFalse SettingGetResponseZonesCiphersEditable = false
)

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

// Current value of the zone setting.
type SettingGetResponseZonesCNAMEFlatteningValue string

const (
	SettingGetResponseZonesCNAMEFlatteningValueFlattenAtRoot SettingGetResponseZonesCNAMEFlatteningValue = "flatten_at_root"
	SettingGetResponseZonesCNAMEFlatteningValueFlattenAll    SettingGetResponseZonesCNAMEFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesCNAMEFlatteningEditable bool

const (
	SettingGetResponseZonesCNAMEFlatteningEditableTrue  SettingGetResponseZonesCNAMEFlatteningEditable = true
	SettingGetResponseZonesCNAMEFlatteningEditableFalse SettingGetResponseZonesCNAMEFlatteningEditable = false
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type SettingGetResponseZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesDevelopmentModeID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesDevelopmentModeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64                                    `json:"time_remaining"`
	JSON          settingGetResponseZonesDevelopmentModeJSON `json:"-"`
}

// settingGetResponseZonesDevelopmentModeJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesDevelopmentMode]
type settingGetResponseZonesDevelopmentModeJSON struct {
	ID            apijson.Field
	Value         apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingGetResponseZonesDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesDevelopmentModeJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesDevelopmentMode) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesDevelopmentModeID string

const (
	SettingGetResponseZonesDevelopmentModeIDDevelopmentMode SettingGetResponseZonesDevelopmentModeID = "development_mode"
)

// Current value of the zone setting.
type SettingGetResponseZonesDevelopmentModeValue string

const (
	SettingGetResponseZonesDevelopmentModeValueOn  SettingGetResponseZonesDevelopmentModeValue = "on"
	SettingGetResponseZonesDevelopmentModeValueOff SettingGetResponseZonesDevelopmentModeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesDevelopmentModeEditable bool

const (
	SettingGetResponseZonesDevelopmentModeEditableTrue  SettingGetResponseZonesDevelopmentModeEditable = true
	SettingGetResponseZonesDevelopmentModeEditableFalse SettingGetResponseZonesDevelopmentModeEditable = false
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type SettingGetResponseZonesEarlyHints struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesEarlyHintsID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesEarlyHintsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesEarlyHintsJSON `json:"-"`
}

// settingGetResponseZonesEarlyHintsJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesEarlyHints]
type settingGetResponseZonesEarlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesEarlyHintsJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesEarlyHints) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesEarlyHintsID string

const (
	SettingGetResponseZonesEarlyHintsIDEarlyHints SettingGetResponseZonesEarlyHintsID = "early_hints"
)

// Current value of the zone setting.
type SettingGetResponseZonesEarlyHintsValue string

const (
	SettingGetResponseZonesEarlyHintsValueOn  SettingGetResponseZonesEarlyHintsValue = "on"
	SettingGetResponseZonesEarlyHintsValueOff SettingGetResponseZonesEarlyHintsValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesEarlyHintsEditable bool

const (
	SettingGetResponseZonesEarlyHintsEditableTrue  SettingGetResponseZonesEarlyHintsEditable = true
	SettingGetResponseZonesEarlyHintsEditableFalse SettingGetResponseZonesEarlyHintsEditable = false
)

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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesEdgeCacheTTLEditable bool

const (
	SettingGetResponseZonesEdgeCacheTTLEditableTrue  SettingGetResponseZonesEdgeCacheTTLEditable = true
	SettingGetResponseZonesEdgeCacheTTLEditableFalse SettingGetResponseZonesEdgeCacheTTLEditable = false
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type SettingGetResponseZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesEmailObfuscationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesEmailObfuscationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesEmailObfuscationJSON `json:"-"`
}

// settingGetResponseZonesEmailObfuscationJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesEmailObfuscation]
type settingGetResponseZonesEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesEmailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesEmailObfuscation) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesEmailObfuscationID string

const (
	SettingGetResponseZonesEmailObfuscationIDEmailObfuscation SettingGetResponseZonesEmailObfuscationID = "email_obfuscation"
)

// Current value of the zone setting.
type SettingGetResponseZonesEmailObfuscationValue string

const (
	SettingGetResponseZonesEmailObfuscationValueOn  SettingGetResponseZonesEmailObfuscationValue = "on"
	SettingGetResponseZonesEmailObfuscationValueOff SettingGetResponseZonesEmailObfuscationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesEmailObfuscationEditable bool

const (
	SettingGetResponseZonesEmailObfuscationEditableTrue  SettingGetResponseZonesEmailObfuscationEditable = true
	SettingGetResponseZonesEmailObfuscationEditableFalse SettingGetResponseZonesEmailObfuscationEditable = false
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type SettingGetResponseZonesH2Prioritization struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesH2PrioritizationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesH2PrioritizationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesH2PrioritizationJSON `json:"-"`
}

// settingGetResponseZonesH2PrioritizationJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesH2Prioritization]
type settingGetResponseZonesH2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesH2PrioritizationJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesH2Prioritization) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesH2PrioritizationID string

const (
	SettingGetResponseZonesH2PrioritizationIDH2Prioritization SettingGetResponseZonesH2PrioritizationID = "h2_prioritization"
)

// Current value of the zone setting.
type SettingGetResponseZonesH2PrioritizationValue string

const (
	SettingGetResponseZonesH2PrioritizationValueOn     SettingGetResponseZonesH2PrioritizationValue = "on"
	SettingGetResponseZonesH2PrioritizationValueOff    SettingGetResponseZonesH2PrioritizationValue = "off"
	SettingGetResponseZonesH2PrioritizationValueCustom SettingGetResponseZonesH2PrioritizationValue = "custom"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesH2PrioritizationEditable bool

const (
	SettingGetResponseZonesH2PrioritizationEditableTrue  SettingGetResponseZonesH2PrioritizationEditable = true
	SettingGetResponseZonesH2PrioritizationEditableFalse SettingGetResponseZonesH2PrioritizationEditable = false
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type SettingGetResponseZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesHotlinkProtectionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesHotlinkProtectionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesHotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesHotlinkProtectionJSON `json:"-"`
}

// settingGetResponseZonesHotlinkProtectionJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesHotlinkProtection]
type settingGetResponseZonesHotlinkProtectionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesHotlinkProtectionJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesHotlinkProtection) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesHotlinkProtectionID string

const (
	SettingGetResponseZonesHotlinkProtectionIDHotlinkProtection SettingGetResponseZonesHotlinkProtectionID = "hotlink_protection"
)

// Current value of the zone setting.
type SettingGetResponseZonesHotlinkProtectionValue string

const (
	SettingGetResponseZonesHotlinkProtectionValueOn  SettingGetResponseZonesHotlinkProtectionValue = "on"
	SettingGetResponseZonesHotlinkProtectionValueOff SettingGetResponseZonesHotlinkProtectionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesHotlinkProtectionEditable bool

const (
	SettingGetResponseZonesHotlinkProtectionEditableTrue  SettingGetResponseZonesHotlinkProtectionEditable = true
	SettingGetResponseZonesHotlinkProtectionEditableFalse SettingGetResponseZonesHotlinkProtectionEditable = false
)

// HTTP2 enabled for this zone.
type SettingGetResponseZonesHTTP2 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesHTTP2ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesHTTP2Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesHTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesHTTP2JSON `json:"-"`
}

// settingGetResponseZonesHTTP2JSON contains the JSON metadata for the struct
// [SettingGetResponseZonesHTTP2]
type settingGetResponseZonesHTTP2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesHTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesHTTP2JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesHTTP2) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesHTTP2ID string

const (
	SettingGetResponseZonesHTTP2IDHTTP2 SettingGetResponseZonesHTTP2ID = "http2"
)

// Current value of the zone setting.
type SettingGetResponseZonesHTTP2Value string

const (
	SettingGetResponseZonesHTTP2ValueOn  SettingGetResponseZonesHTTP2Value = "on"
	SettingGetResponseZonesHTTP2ValueOff SettingGetResponseZonesHTTP2Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesHTTP2Editable bool

const (
	SettingGetResponseZonesHTTP2EditableTrue  SettingGetResponseZonesHTTP2Editable = true
	SettingGetResponseZonesHTTP2EditableFalse SettingGetResponseZonesHTTP2Editable = false
)

// HTTP3 enabled for this zone.
type SettingGetResponseZonesHTTP3 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesHTTP3ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesHTTP3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesHTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesHTTP3JSON `json:"-"`
}

// settingGetResponseZonesHTTP3JSON contains the JSON metadata for the struct
// [SettingGetResponseZonesHTTP3]
type settingGetResponseZonesHTTP3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesHTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesHTTP3JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesHTTP3) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesHTTP3ID string

const (
	SettingGetResponseZonesHTTP3IDHTTP3 SettingGetResponseZonesHTTP3ID = "http3"
)

// Current value of the zone setting.
type SettingGetResponseZonesHTTP3Value string

const (
	SettingGetResponseZonesHTTP3ValueOn  SettingGetResponseZonesHTTP3Value = "on"
	SettingGetResponseZonesHTTP3ValueOff SettingGetResponseZonesHTTP3Value = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesHTTP3Editable bool

const (
	SettingGetResponseZonesHTTP3EditableTrue  SettingGetResponseZonesHTTP3Editable = true
	SettingGetResponseZonesHTTP3EditableFalse SettingGetResponseZonesHTTP3Editable = false
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type SettingGetResponseZonesImageResizing struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesImageResizingID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesImageResizingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesImageResizingJSON `json:"-"`
}

// settingGetResponseZonesImageResizingJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesImageResizing]
type settingGetResponseZonesImageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesImageResizingJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesImageResizing) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesImageResizingID string

const (
	SettingGetResponseZonesImageResizingIDImageResizing SettingGetResponseZonesImageResizingID = "image_resizing"
)

// Current value of the zone setting.
type SettingGetResponseZonesImageResizingValue string

const (
	SettingGetResponseZonesImageResizingValueOn   SettingGetResponseZonesImageResizingValue = "on"
	SettingGetResponseZonesImageResizingValueOff  SettingGetResponseZonesImageResizingValue = "off"
	SettingGetResponseZonesImageResizingValueOpen SettingGetResponseZonesImageResizingValue = "open"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesImageResizingEditable bool

const (
	SettingGetResponseZonesImageResizingEditableTrue  SettingGetResponseZonesImageResizingEditable = true
	SettingGetResponseZonesImageResizingEditableFalse SettingGetResponseZonesImageResizingEditable = false
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type SettingGetResponseZonesIPGeolocation struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesIPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesIPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesIPGeolocationJSON `json:"-"`
}

// settingGetResponseZonesIPGeolocationJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesIPGeolocation]
type settingGetResponseZonesIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesIPGeolocationJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesIPGeolocation) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesIPGeolocationID string

const (
	SettingGetResponseZonesIPGeolocationIDIPGeolocation SettingGetResponseZonesIPGeolocationID = "ip_geolocation"
)

// Current value of the zone setting.
type SettingGetResponseZonesIPGeolocationValue string

const (
	SettingGetResponseZonesIPGeolocationValueOn  SettingGetResponseZonesIPGeolocationValue = "on"
	SettingGetResponseZonesIPGeolocationValueOff SettingGetResponseZonesIPGeolocationValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesIPGeolocationEditable bool

const (
	SettingGetResponseZonesIPGeolocationEditableTrue  SettingGetResponseZonesIPGeolocationEditable = true
	SettingGetResponseZonesIPGeolocationEditableFalse SettingGetResponseZonesIPGeolocationEditable = false
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type SettingGetResponseZonesIPV6 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesIPV6ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesIPV6Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesIPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesIPV6JSON `json:"-"`
}

// settingGetResponseZonesIPV6JSON contains the JSON metadata for the struct
// [SettingGetResponseZonesIPV6]
type settingGetResponseZonesIPV6JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesIPV6JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesIPV6) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesIPV6ID string

const (
	SettingGetResponseZonesIPV6IDIPV6 SettingGetResponseZonesIPV6ID = "ipv6"
)

// Current value of the zone setting.
type SettingGetResponseZonesIPV6Value string

const (
	SettingGetResponseZonesIPV6ValueOff SettingGetResponseZonesIPV6Value = "off"
	SettingGetResponseZonesIPV6ValueOn  SettingGetResponseZonesIPV6Value = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesIPV6Editable bool

const (
	SettingGetResponseZonesIPV6EditableTrue  SettingGetResponseZonesIPV6Editable = true
	SettingGetResponseZonesIPV6EditableFalse SettingGetResponseZonesIPV6Editable = false
)

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

// Current value of the zone setting.
type SettingGetResponseZonesMaxUploadValue float64

const (
	SettingGetResponseZonesMaxUploadValue100 SettingGetResponseZonesMaxUploadValue = 100
	SettingGetResponseZonesMaxUploadValue200 SettingGetResponseZonesMaxUploadValue = 200
	SettingGetResponseZonesMaxUploadValue500 SettingGetResponseZonesMaxUploadValue = 500
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesMaxUploadEditable bool

const (
	SettingGetResponseZonesMaxUploadEditableTrue  SettingGetResponseZonesMaxUploadEditable = true
	SettingGetResponseZonesMaxUploadEditableFalse SettingGetResponseZonesMaxUploadEditable = false
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type SettingGetResponseZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesMinTLSVersionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesMinTLSVersionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesMinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesMinTLSVersionJSON `json:"-"`
}

// settingGetResponseZonesMinTLSVersionJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesMinTLSVersion]
type settingGetResponseZonesMinTLSVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesMinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesMinTLSVersionJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesMinTLSVersion) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesMinTLSVersionID string

const (
	SettingGetResponseZonesMinTLSVersionIDMinTLSVersion SettingGetResponseZonesMinTLSVersionID = "min_tls_version"
)

// Current value of the zone setting.
type SettingGetResponseZonesMinTLSVersionValue string

const (
	SettingGetResponseZonesMinTLSVersionValue1_0 SettingGetResponseZonesMinTLSVersionValue = "1.0"
	SettingGetResponseZonesMinTLSVersionValue1_1 SettingGetResponseZonesMinTLSVersionValue = "1.1"
	SettingGetResponseZonesMinTLSVersionValue1_2 SettingGetResponseZonesMinTLSVersionValue = "1.2"
	SettingGetResponseZonesMinTLSVersionValue1_3 SettingGetResponseZonesMinTLSVersionValue = "1.3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesMinTLSVersionEditable bool

const (
	SettingGetResponseZonesMinTLSVersionEditableTrue  SettingGetResponseZonesMinTLSVersionEditable = true
	SettingGetResponseZonesMinTLSVersionEditableFalse SettingGetResponseZonesMinTLSVersionEditable = false
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type SettingGetResponseZonesMinify struct {
	// Zone setting identifier.
	ID SettingGetResponseZonesMinifyID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesMinifyValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesMinifyJSON `json:"-"`
}

// settingGetResponseZonesMinifyJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesMinify]
type settingGetResponseZonesMinifyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesMinifyJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesMinify) implementsZonesSettingGetResponse() {}

// Zone setting identifier.
type SettingGetResponseZonesMinifyID string

const (
	SettingGetResponseZonesMinifyIDMinify SettingGetResponseZonesMinifyID = "minify"
)

// Current value of the zone setting.
type SettingGetResponseZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css SettingGetResponseZonesMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML SettingGetResponseZonesMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   SettingGetResponseZonesMinifyValueJs   `json:"js"`
	JSON settingGetResponseZonesMinifyValueJSON `json:"-"`
}

// settingGetResponseZonesMinifyValueJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesMinifyValue]
type settingGetResponseZonesMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesMinifyValueJSON) RawJSON() string {
	return r.raw
}

// Automatically minify all CSS files for your website.
type SettingGetResponseZonesMinifyValueCss string

const (
	SettingGetResponseZonesMinifyValueCssOn  SettingGetResponseZonesMinifyValueCss = "on"
	SettingGetResponseZonesMinifyValueCssOff SettingGetResponseZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type SettingGetResponseZonesMinifyValueHTML string

const (
	SettingGetResponseZonesMinifyValueHTMLOn  SettingGetResponseZonesMinifyValueHTML = "on"
	SettingGetResponseZonesMinifyValueHTMLOff SettingGetResponseZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type SettingGetResponseZonesMinifyValueJs string

const (
	SettingGetResponseZonesMinifyValueJsOn  SettingGetResponseZonesMinifyValueJs = "on"
	SettingGetResponseZonesMinifyValueJsOff SettingGetResponseZonesMinifyValueJs = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesMinifyEditable bool

const (
	SettingGetResponseZonesMinifyEditableTrue  SettingGetResponseZonesMinifyEditable = true
	SettingGetResponseZonesMinifyEditableFalse SettingGetResponseZonesMinifyEditable = false
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type SettingGetResponseZonesMirage struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesMirageID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesMirageValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesMirageJSON `json:"-"`
}

// settingGetResponseZonesMirageJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesMirage]
type settingGetResponseZonesMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesMirageJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesMirage) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesMirageID string

const (
	SettingGetResponseZonesMirageIDMirage SettingGetResponseZonesMirageID = "mirage"
)

// Current value of the zone setting.
type SettingGetResponseZonesMirageValue string

const (
	SettingGetResponseZonesMirageValueOn  SettingGetResponseZonesMirageValue = "on"
	SettingGetResponseZonesMirageValueOff SettingGetResponseZonesMirageValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesMirageEditable bool

const (
	SettingGetResponseZonesMirageEditableTrue  SettingGetResponseZonesMirageEditable = true
	SettingGetResponseZonesMirageEditableFalse SettingGetResponseZonesMirageEditable = false
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type SettingGetResponseZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID SettingGetResponseZonesMobileRedirectID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesMobileRedirectValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesMobileRedirectJSON `json:"-"`
}

// settingGetResponseZonesMobileRedirectJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesMobileRedirect]
type settingGetResponseZonesMobileRedirectJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesMobileRedirectJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesMobileRedirect) implementsZonesSettingGetResponse() {}

// Identifier of the zone setting.
type SettingGetResponseZonesMobileRedirectID string

const (
	SettingGetResponseZonesMobileRedirectIDMobileRedirect SettingGetResponseZonesMobileRedirectID = "mobile_redirect"
)

// Current value of the zone setting.
type SettingGetResponseZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status SettingGetResponseZonesMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripURI bool                                           `json:"strip_uri"`
	JSON     settingGetResponseZonesMobileRedirectValueJSON `json:"-"`
}

// settingGetResponseZonesMobileRedirectValueJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesMobileRedirectValue]
type settingGetResponseZonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripURI        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingGetResponseZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesMobileRedirectValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not mobile redirect is enabled.
type SettingGetResponseZonesMobileRedirectValueStatus string

const (
	SettingGetResponseZonesMobileRedirectValueStatusOn  SettingGetResponseZonesMobileRedirectValueStatus = "on"
	SettingGetResponseZonesMobileRedirectValueStatusOff SettingGetResponseZonesMobileRedirectValueStatus = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesMobileRedirectEditable bool

const (
	SettingGetResponseZonesMobileRedirectEditableTrue  SettingGetResponseZonesMobileRedirectEditable = true
	SettingGetResponseZonesMobileRedirectEditableFalse SettingGetResponseZonesMobileRedirectEditable = false
)

// Enable Network Error Logging reporting on your zone. (Beta)
type SettingGetResponseZonesNEL struct {
	// Zone setting identifier.
	ID SettingGetResponseZonesNELID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesNELValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesNELJSON `json:"-"`
}

// settingGetResponseZonesNELJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesNEL]
type settingGetResponseZonesNELJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesNELJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesNEL) implementsZonesSettingGetResponse() {}

// Zone setting identifier.
type SettingGetResponseZonesNELID string

const (
	SettingGetResponseZonesNELIDNEL SettingGetResponseZonesNELID = "nel"
)

// Current value of the zone setting.
type SettingGetResponseZonesNELValue struct {
	Enabled bool                                `json:"enabled"`
	JSON    settingGetResponseZonesNELValueJSON `json:"-"`
}

// settingGetResponseZonesNELValueJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesNELValue]
type settingGetResponseZonesNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesNELValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesNELValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesNELEditable bool

const (
	SettingGetResponseZonesNELEditableTrue  SettingGetResponseZonesNELEditable = true
	SettingGetResponseZonesNELEditableFalse SettingGetResponseZonesNELEditable = false
)

// Enables the Opportunistic Encryption feature for a zone.
type SettingGetResponseZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesOpportunisticEncryptionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesOpportunisticEncryptionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesOpportunisticEncryptionJSON `json:"-"`
}

// settingGetResponseZonesOpportunisticEncryptionJSON contains the JSON metadata
// for the struct [SettingGetResponseZonesOpportunisticEncryption]
type settingGetResponseZonesOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesOpportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesOpportunisticEncryption) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesOpportunisticEncryptionID string

const (
	SettingGetResponseZonesOpportunisticEncryptionIDOpportunisticEncryption SettingGetResponseZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Current value of the zone setting.
type SettingGetResponseZonesOpportunisticEncryptionValue string

const (
	SettingGetResponseZonesOpportunisticEncryptionValueOn  SettingGetResponseZonesOpportunisticEncryptionValue = "on"
	SettingGetResponseZonesOpportunisticEncryptionValueOff SettingGetResponseZonesOpportunisticEncryptionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesOpportunisticEncryptionEditable bool

const (
	SettingGetResponseZonesOpportunisticEncryptionEditableTrue  SettingGetResponseZonesOpportunisticEncryptionEditable = true
	SettingGetResponseZonesOpportunisticEncryptionEditableFalse SettingGetResponseZonesOpportunisticEncryptionEditable = false
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type SettingGetResponseZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesOpportunisticOnionID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesOpportunisticOnionValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesOpportunisticOnionJSON `json:"-"`
}

// settingGetResponseZonesOpportunisticOnionJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesOpportunisticOnion]
type settingGetResponseZonesOpportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesOpportunisticOnionJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesOpportunisticOnion) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesOpportunisticOnionID string

const (
	SettingGetResponseZonesOpportunisticOnionIDOpportunisticOnion SettingGetResponseZonesOpportunisticOnionID = "opportunistic_onion"
)

// Current value of the zone setting.
type SettingGetResponseZonesOpportunisticOnionValue string

const (
	SettingGetResponseZonesOpportunisticOnionValueOn  SettingGetResponseZonesOpportunisticOnionValue = "on"
	SettingGetResponseZonesOpportunisticOnionValueOff SettingGetResponseZonesOpportunisticOnionValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesOpportunisticOnionEditable bool

const (
	SettingGetResponseZonesOpportunisticOnionEditableTrue  SettingGetResponseZonesOpportunisticOnionEditable = true
	SettingGetResponseZonesOpportunisticOnionEditableFalse SettingGetResponseZonesOpportunisticOnionEditable = false
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type SettingGetResponseZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesOrangeToOrangeID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesOrangeToOrangeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesOrangeToOrangeJSON `json:"-"`
}

// settingGetResponseZonesOrangeToOrangeJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesOrangeToOrange]
type settingGetResponseZonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesOrangeToOrangeJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesOrangeToOrange) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesOrangeToOrangeID string

const (
	SettingGetResponseZonesOrangeToOrangeIDOrangeToOrange SettingGetResponseZonesOrangeToOrangeID = "orange_to_orange"
)

// Current value of the zone setting.
type SettingGetResponseZonesOrangeToOrangeValue string

const (
	SettingGetResponseZonesOrangeToOrangeValueOn  SettingGetResponseZonesOrangeToOrangeValue = "on"
	SettingGetResponseZonesOrangeToOrangeValueOff SettingGetResponseZonesOrangeToOrangeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesOrangeToOrangeEditable bool

const (
	SettingGetResponseZonesOrangeToOrangeEditableTrue  SettingGetResponseZonesOrangeToOrangeEditable = true
	SettingGetResponseZonesOrangeToOrangeEditableFalse SettingGetResponseZonesOrangeToOrangeEditable = false
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type SettingGetResponseZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesOriginErrorPagePassThruID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesOriginErrorPagePassThruValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesOriginErrorPagePassThruJSON `json:"-"`
}

// settingGetResponseZonesOriginErrorPagePassThruJSON contains the JSON metadata
// for the struct [SettingGetResponseZonesOriginErrorPagePassThru]
type settingGetResponseZonesOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesOriginErrorPagePassThruJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesOriginErrorPagePassThru) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesOriginErrorPagePassThruID string

const (
	SettingGetResponseZonesOriginErrorPagePassThruIDOriginErrorPagePassThru SettingGetResponseZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Current value of the zone setting.
type SettingGetResponseZonesOriginErrorPagePassThruValue string

const (
	SettingGetResponseZonesOriginErrorPagePassThruValueOn  SettingGetResponseZonesOriginErrorPagePassThruValue = "on"
	SettingGetResponseZonesOriginErrorPagePassThruValueOff SettingGetResponseZonesOriginErrorPagePassThruValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesOriginErrorPagePassThruEditable bool

const (
	SettingGetResponseZonesOriginErrorPagePassThruEditableTrue  SettingGetResponseZonesOriginErrorPagePassThruEditable = true
	SettingGetResponseZonesOriginErrorPagePassThruEditableFalse SettingGetResponseZonesOriginErrorPagePassThruEditable = false
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type SettingGetResponseZonesPolish struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesPolishID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesPolishValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesPolishJSON `json:"-"`
}

// settingGetResponseZonesPolishJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesPolish]
type settingGetResponseZonesPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesPolishJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesPolish) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesPolishID string

const (
	SettingGetResponseZonesPolishIDPolish SettingGetResponseZonesPolishID = "polish"
)

// Current value of the zone setting.
type SettingGetResponseZonesPolishValue string

const (
	SettingGetResponseZonesPolishValueOff      SettingGetResponseZonesPolishValue = "off"
	SettingGetResponseZonesPolishValueLossless SettingGetResponseZonesPolishValue = "lossless"
	SettingGetResponseZonesPolishValueLossy    SettingGetResponseZonesPolishValue = "lossy"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesPolishEditable bool

const (
	SettingGetResponseZonesPolishEditableTrue  SettingGetResponseZonesPolishEditable = true
	SettingGetResponseZonesPolishEditableFalse SettingGetResponseZonesPolishEditable = false
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type SettingGetResponseZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesPrefetchPreloadID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesPrefetchPreloadValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesPrefetchPreloadJSON `json:"-"`
}

// settingGetResponseZonesPrefetchPreloadJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesPrefetchPreload]
type settingGetResponseZonesPrefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesPrefetchPreloadJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesPrefetchPreload) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesPrefetchPreloadID string

const (
	SettingGetResponseZonesPrefetchPreloadIDPrefetchPreload SettingGetResponseZonesPrefetchPreloadID = "prefetch_preload"
)

// Current value of the zone setting.
type SettingGetResponseZonesPrefetchPreloadValue string

const (
	SettingGetResponseZonesPrefetchPreloadValueOn  SettingGetResponseZonesPrefetchPreloadValue = "on"
	SettingGetResponseZonesPrefetchPreloadValueOff SettingGetResponseZonesPrefetchPreloadValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesPrefetchPreloadEditable bool

const (
	SettingGetResponseZonesPrefetchPreloadEditableTrue  SettingGetResponseZonesPrefetchPreloadEditable = true
	SettingGetResponseZonesPrefetchPreloadEditableFalse SettingGetResponseZonesPrefetchPreloadEditable = false
)

// Maximum time between two read operations from origin.
type SettingGetResponseZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesProxyReadTimeoutID `json:"id,required"`
	// Current value of the zone setting.
	Value float64 `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                   `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesProxyReadTimeoutJSON `json:"-"`
}

// settingGetResponseZonesProxyReadTimeoutJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesProxyReadTimeout]
type settingGetResponseZonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesProxyReadTimeoutJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesProxyReadTimeout) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesProxyReadTimeoutID string

const (
	SettingGetResponseZonesProxyReadTimeoutIDProxyReadTimeout SettingGetResponseZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesProxyReadTimeoutEditable bool

const (
	SettingGetResponseZonesProxyReadTimeoutEditableTrue  SettingGetResponseZonesProxyReadTimeoutEditable = true
	SettingGetResponseZonesProxyReadTimeoutEditableFalse SettingGetResponseZonesProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type SettingGetResponseZonesPseudoIPV4 struct {
	// Value of the Pseudo IPv4 setting.
	ID SettingGetResponseZonesPseudoIPV4ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesPseudoIPV4Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesPseudoIPV4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesPseudoIPV4JSON `json:"-"`
}

// settingGetResponseZonesPseudoIPV4JSON contains the JSON metadata for the struct
// [SettingGetResponseZonesPseudoIPV4]
type settingGetResponseZonesPseudoIPV4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesPseudoIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesPseudoIPV4JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesPseudoIPV4) implementsZonesSettingGetResponse() {}

// Value of the Pseudo IPv4 setting.
type SettingGetResponseZonesPseudoIPV4ID string

const (
	SettingGetResponseZonesPseudoIPV4IDPseudoIPV4 SettingGetResponseZonesPseudoIPV4ID = "pseudo_ipv4"
)

// Current value of the zone setting.
type SettingGetResponseZonesPseudoIPV4Value string

const (
	SettingGetResponseZonesPseudoIPV4ValueOff             SettingGetResponseZonesPseudoIPV4Value = "off"
	SettingGetResponseZonesPseudoIPV4ValueAddHeader       SettingGetResponseZonesPseudoIPV4Value = "add_header"
	SettingGetResponseZonesPseudoIPV4ValueOverwriteHeader SettingGetResponseZonesPseudoIPV4Value = "overwrite_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesPseudoIPV4Editable bool

const (
	SettingGetResponseZonesPseudoIPV4EditableTrue  SettingGetResponseZonesPseudoIPV4Editable = true
	SettingGetResponseZonesPseudoIPV4EditableFalse SettingGetResponseZonesPseudoIPV4Editable = false
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type SettingGetResponseZonesResponseBuffering struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesResponseBufferingID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesResponseBufferingValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesResponseBufferingJSON `json:"-"`
}

// settingGetResponseZonesResponseBufferingJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesResponseBuffering]
type settingGetResponseZonesResponseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesResponseBufferingJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesResponseBuffering) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesResponseBufferingID string

const (
	SettingGetResponseZonesResponseBufferingIDResponseBuffering SettingGetResponseZonesResponseBufferingID = "response_buffering"
)

// Current value of the zone setting.
type SettingGetResponseZonesResponseBufferingValue string

const (
	SettingGetResponseZonesResponseBufferingValueOn  SettingGetResponseZonesResponseBufferingValue = "on"
	SettingGetResponseZonesResponseBufferingValueOff SettingGetResponseZonesResponseBufferingValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesResponseBufferingEditable bool

const (
	SettingGetResponseZonesResponseBufferingEditableTrue  SettingGetResponseZonesResponseBufferingEditable = true
	SettingGetResponseZonesResponseBufferingEditableFalse SettingGetResponseZonesResponseBufferingEditable = false
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
type SettingGetResponseZonesRocketLoader struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesRocketLoaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesRocketLoaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesRocketLoaderJSON `json:"-"`
}

// settingGetResponseZonesRocketLoaderJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesRocketLoader]
type settingGetResponseZonesRocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesRocketLoaderJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesRocketLoader) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesRocketLoaderID string

const (
	SettingGetResponseZonesRocketLoaderIDRocketLoader SettingGetResponseZonesRocketLoaderID = "rocket_loader"
)

// Current value of the zone setting.
type SettingGetResponseZonesRocketLoaderValue string

const (
	SettingGetResponseZonesRocketLoaderValueOn  SettingGetResponseZonesRocketLoaderValue = "on"
	SettingGetResponseZonesRocketLoaderValueOff SettingGetResponseZonesRocketLoaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesRocketLoaderEditable bool

const (
	SettingGetResponseZonesRocketLoaderEditableTrue  SettingGetResponseZonesRocketLoaderEditable = true
	SettingGetResponseZonesRocketLoaderEditableFalse SettingGetResponseZonesRocketLoaderEditable = false
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingGetResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSchemasAutomaticPlatformOptimizationValue `json:"value,required"`
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

// Current value of the zone setting.
type SettingGetResponseZonesSchemasAutomaticPlatformOptimizationValue struct {
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
	WpPlugin bool                                                                 `json:"wp_plugin,required"`
	JSON     settingGetResponseZonesSchemasAutomaticPlatformOptimizationValueJSON `json:"-"`
}

// settingGetResponseZonesSchemasAutomaticPlatformOptimizationValueJSON contains
// the JSON metadata for the struct
// [SettingGetResponseZonesSchemasAutomaticPlatformOptimizationValue]
type settingGetResponseZonesSchemasAutomaticPlatformOptimizationValueJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SettingGetResponseZonesSchemasAutomaticPlatformOptimizationValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSchemasAutomaticPlatformOptimizationValueJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
)

// Cloudflare security header for a zone.
type SettingGetResponseZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID SettingGetResponseZonesSecurityHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSecurityHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSecurityHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                 `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSecurityHeaderJSON `json:"-"`
}

// settingGetResponseZonesSecurityHeaderJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesSecurityHeader]
type settingGetResponseZonesSecurityHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSecurityHeader) implementsZonesSettingGetResponse() {}

// ID of the zone's security header.
type SettingGetResponseZonesSecurityHeaderID string

const (
	SettingGetResponseZonesSecurityHeaderIDSecurityHeader SettingGetResponseZonesSecurityHeaderID = "security_header"
)

// Current value of the zone setting.
type SettingGetResponseZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity SettingGetResponseZonesSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    settingGetResponseZonesSecurityHeaderValueJSON                    `json:"-"`
}

// settingGetResponseZonesSecurityHeaderValueJSON contains the JSON metadata for
// the struct [SettingGetResponseZonesSecurityHeaderValue]
type settingGetResponseZonesSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SettingGetResponseZonesSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSecurityHeaderValueJSON) RawJSON() string {
	return r.raw
}

// Strict Transport Security.
type SettingGetResponseZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                  `json:"nosniff"`
	JSON    settingGetResponseZonesSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// settingGetResponseZonesSecurityHeaderValueStrictTransportSecurityJSON contains
// the JSON metadata for the struct
// [SettingGetResponseZonesSecurityHeaderValueStrictTransportSecurity]
type settingGetResponseZonesSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SettingGetResponseZonesSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSecurityHeaderValueStrictTransportSecurityJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSecurityHeaderEditable bool

const (
	SettingGetResponseZonesSecurityHeaderEditableTrue  SettingGetResponseZonesSecurityHeaderEditable = true
	SettingGetResponseZonesSecurityHeaderEditableFalse SettingGetResponseZonesSecurityHeaderEditable = false
)

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type SettingGetResponseZonesSecurityLevel struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSecurityLevelID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSecurityLevelValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSecurityLevelJSON `json:"-"`
}

// settingGetResponseZonesSecurityLevelJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesSecurityLevel]
type settingGetResponseZonesSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSecurityLevelJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSecurityLevel) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSecurityLevelID string

const (
	SettingGetResponseZonesSecurityLevelIDSecurityLevel SettingGetResponseZonesSecurityLevelID = "security_level"
)

// Current value of the zone setting.
type SettingGetResponseZonesSecurityLevelValue string

const (
	SettingGetResponseZonesSecurityLevelValueOff            SettingGetResponseZonesSecurityLevelValue = "off"
	SettingGetResponseZonesSecurityLevelValueEssentiallyOff SettingGetResponseZonesSecurityLevelValue = "essentially_off"
	SettingGetResponseZonesSecurityLevelValueLow            SettingGetResponseZonesSecurityLevelValue = "low"
	SettingGetResponseZonesSecurityLevelValueMedium         SettingGetResponseZonesSecurityLevelValue = "medium"
	SettingGetResponseZonesSecurityLevelValueHigh           SettingGetResponseZonesSecurityLevelValue = "high"
	SettingGetResponseZonesSecurityLevelValueUnderAttack    SettingGetResponseZonesSecurityLevelValue = "under_attack"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSecurityLevelEditable bool

const (
	SettingGetResponseZonesSecurityLevelEditableTrue  SettingGetResponseZonesSecurityLevelEditable = true
	SettingGetResponseZonesSecurityLevelEditableFalse SettingGetResponseZonesSecurityLevelEditable = false
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
type SettingGetResponseZonesServerSideExclude struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesServerSideExcludeID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesServerSideExcludeValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesServerSideExcludeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesServerSideExcludeJSON `json:"-"`
}

// settingGetResponseZonesServerSideExcludeJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesServerSideExclude]
type settingGetResponseZonesServerSideExcludeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesServerSideExcludeJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesServerSideExclude) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesServerSideExcludeID string

const (
	SettingGetResponseZonesServerSideExcludeIDServerSideExclude SettingGetResponseZonesServerSideExcludeID = "server_side_exclude"
)

// Current value of the zone setting.
type SettingGetResponseZonesServerSideExcludeValue string

const (
	SettingGetResponseZonesServerSideExcludeValueOn  SettingGetResponseZonesServerSideExcludeValue = "on"
	SettingGetResponseZonesServerSideExcludeValueOff SettingGetResponseZonesServerSideExcludeValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesServerSideExcludeEditable bool

const (
	SettingGetResponseZonesServerSideExcludeEditableTrue  SettingGetResponseZonesServerSideExcludeEditable = true
	SettingGetResponseZonesServerSideExcludeEditableFalse SettingGetResponseZonesServerSideExcludeEditable = false
)

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

// Current value of the zone setting.
type SettingGetResponseZonesSha1SupportValue string

const (
	SettingGetResponseZonesSha1SupportValueOff SettingGetResponseZonesSha1SupportValue = "off"
	SettingGetResponseZonesSha1SupportValueOn  SettingGetResponseZonesSha1SupportValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSha1SupportEditable bool

const (
	SettingGetResponseZonesSha1SupportEditableTrue  SettingGetResponseZonesSha1SupportEditable = true
	SettingGetResponseZonesSha1SupportEditableFalse SettingGetResponseZonesSha1SupportEditable = false
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type SettingGetResponseZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSortQueryStringForCacheID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSortQueryStringForCacheValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                          `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSortQueryStringForCacheJSON `json:"-"`
}

// settingGetResponseZonesSortQueryStringForCacheJSON contains the JSON metadata
// for the struct [SettingGetResponseZonesSortQueryStringForCache]
type settingGetResponseZonesSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSortQueryStringForCache) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSortQueryStringForCacheID string

const (
	SettingGetResponseZonesSortQueryStringForCacheIDSortQueryStringForCache SettingGetResponseZonesSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Current value of the zone setting.
type SettingGetResponseZonesSortQueryStringForCacheValue string

const (
	SettingGetResponseZonesSortQueryStringForCacheValueOn  SettingGetResponseZonesSortQueryStringForCacheValue = "on"
	SettingGetResponseZonesSortQueryStringForCacheValueOff SettingGetResponseZonesSortQueryStringForCacheValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSortQueryStringForCacheEditable bool

const (
	SettingGetResponseZonesSortQueryStringForCacheEditableTrue  SettingGetResponseZonesSortQueryStringForCacheEditable = true
	SettingGetResponseZonesSortQueryStringForCacheEditableFalse SettingGetResponseZonesSortQueryStringForCacheEditable = false
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
type SettingGetResponseZonesSSL struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSSLID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesSSLValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesSSLJSON `json:"-"`
}

// settingGetResponseZonesSSLJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesSSL]
type settingGetResponseZonesSSLJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSSLJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSSL) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesSSLID string

const (
	SettingGetResponseZonesSSLIDSSL SettingGetResponseZonesSSLID = "ssl"
)

// Current value of the zone setting.
type SettingGetResponseZonesSSLValue string

const (
	SettingGetResponseZonesSSLValueOff      SettingGetResponseZonesSSLValue = "off"
	SettingGetResponseZonesSSLValueFlexible SettingGetResponseZonesSSLValue = "flexible"
	SettingGetResponseZonesSSLValueFull     SettingGetResponseZonesSSLValue = "full"
	SettingGetResponseZonesSSLValueStrict   SettingGetResponseZonesSSLValue = "strict"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSSLEditable bool

const (
	SettingGetResponseZonesSSLEditableTrue  SettingGetResponseZonesSSLEditable = true
	SettingGetResponseZonesSSLEditableFalse SettingGetResponseZonesSSLEditable = false
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type SettingGetResponseZonesSSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID SettingGetResponseZonesSSLRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                      `json:"enabled"`
	JSON    settingGetResponseZonesSSLRecommenderJSON `json:"-"`
}

// settingGetResponseZonesSSLRecommenderJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesSSLRecommender]
type settingGetResponseZonesSSLRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesSSLRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesSSLRecommenderJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesSSLRecommender) implementsZonesSettingGetResponse() {}

// Enrollment value for SSL/TLS Recommender.
type SettingGetResponseZonesSSLRecommenderID string

const (
	SettingGetResponseZonesSSLRecommenderIDSSLRecommender SettingGetResponseZonesSSLRecommenderID = "ssl_recommender"
)

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

// Current value of the zone setting.
type SettingGetResponseZonesTLS1_2OnlyValue string

const (
	SettingGetResponseZonesTLS1_2OnlyValueOff SettingGetResponseZonesTLS1_2OnlyValue = "off"
	SettingGetResponseZonesTLS1_2OnlyValueOn  SettingGetResponseZonesTLS1_2OnlyValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesTLS1_2OnlyEditable bool

const (
	SettingGetResponseZonesTLS1_2OnlyEditableTrue  SettingGetResponseZonesTLS1_2OnlyEditable = true
	SettingGetResponseZonesTLS1_2OnlyEditableFalse SettingGetResponseZonesTLS1_2OnlyEditable = false
)

// Enables Crypto TLS 1.3 feature for a zone.
type SettingGetResponseZonesTLS1_3 struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesTLS1_3ID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesTLS1_3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesTLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                         `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesTls1_3JSON `json:"-"`
}

// settingGetResponseZonesTls1_3JSON contains the JSON metadata for the struct
// [SettingGetResponseZonesTLS1_3]
type settingGetResponseZonesTls1_3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesTLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesTls1_3JSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesTLS1_3) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesTLS1_3ID string

const (
	SettingGetResponseZonesTLS1_3IDTLS1_3 SettingGetResponseZonesTLS1_3ID = "tls_1_3"
)

// Current value of the zone setting.
type SettingGetResponseZonesTLS1_3Value string

const (
	SettingGetResponseZonesTLS1_3ValueOn  SettingGetResponseZonesTLS1_3Value = "on"
	SettingGetResponseZonesTLS1_3ValueOff SettingGetResponseZonesTLS1_3Value = "off"
	SettingGetResponseZonesTLS1_3ValueZrt SettingGetResponseZonesTLS1_3Value = "zrt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesTLS1_3Editable bool

const (
	SettingGetResponseZonesTLS1_3EditableTrue  SettingGetResponseZonesTLS1_3Editable = true
	SettingGetResponseZonesTLS1_3EditableFalse SettingGetResponseZonesTLS1_3Editable = false
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type SettingGetResponseZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesTLSClientAuthID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesTLSClientAuthValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesTLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesTLSClientAuthJSON `json:"-"`
}

// settingGetResponseZonesTLSClientAuthJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesTLSClientAuth]
type settingGetResponseZonesTLSClientAuthJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesTLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesTLSClientAuthJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesTLSClientAuth) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesTLSClientAuthID string

const (
	SettingGetResponseZonesTLSClientAuthIDTLSClientAuth SettingGetResponseZonesTLSClientAuthID = "tls_client_auth"
)

// Current value of the zone setting.
type SettingGetResponseZonesTLSClientAuthValue string

const (
	SettingGetResponseZonesTLSClientAuthValueOn  SettingGetResponseZonesTLSClientAuthValue = "on"
	SettingGetResponseZonesTLSClientAuthValueOff SettingGetResponseZonesTLSClientAuthValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesTLSClientAuthEditable bool

const (
	SettingGetResponseZonesTLSClientAuthEditableTrue  SettingGetResponseZonesTLSClientAuthEditable = true
	SettingGetResponseZonesTLSClientAuthEditableFalse SettingGetResponseZonesTLSClientAuthEditable = false
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type SettingGetResponseZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesTrueClientIPHeaderID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesTrueClientIPHeaderValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesTrueClientIPHeaderJSON `json:"-"`
}

// settingGetResponseZonesTrueClientIPHeaderJSON contains the JSON metadata for the
// struct [SettingGetResponseZonesTrueClientIPHeader]
type settingGetResponseZonesTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesTrueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesTrueClientIPHeader) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesTrueClientIPHeaderID string

const (
	SettingGetResponseZonesTrueClientIPHeaderIDTrueClientIPHeader SettingGetResponseZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Current value of the zone setting.
type SettingGetResponseZonesTrueClientIPHeaderValue string

const (
	SettingGetResponseZonesTrueClientIPHeaderValueOn  SettingGetResponseZonesTrueClientIPHeaderValue = "on"
	SettingGetResponseZonesTrueClientIPHeaderValueOff SettingGetResponseZonesTrueClientIPHeaderValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesTrueClientIPHeaderEditable bool

const (
	SettingGetResponseZonesTrueClientIPHeaderEditableTrue  SettingGetResponseZonesTrueClientIPHeaderEditable = true
	SettingGetResponseZonesTrueClientIPHeaderEditableFalse SettingGetResponseZonesTrueClientIPHeaderEditable = false
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
type SettingGetResponseZonesWAF struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesWAFID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesWAFValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesWAFJSON `json:"-"`
}

// settingGetResponseZonesWAFJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesWAF]
type settingGetResponseZonesWAFJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesWAFJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesWAF) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesWAFID string

const (
	SettingGetResponseZonesWAFIDWAF SettingGetResponseZonesWAFID = "waf"
)

// Current value of the zone setting.
type SettingGetResponseZonesWAFValue string

const (
	SettingGetResponseZonesWAFValueOn  SettingGetResponseZonesWAFValue = "on"
	SettingGetResponseZonesWAFValueOff SettingGetResponseZonesWAFValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesWAFEditable bool

const (
	SettingGetResponseZonesWAFEditableTrue  SettingGetResponseZonesWAFEditable = true
	SettingGetResponseZonesWAFEditableFalse SettingGetResponseZonesWAFEditable = false
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type SettingGetResponseZonesWebp struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesWebpID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesWebpValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesWebpJSON `json:"-"`
}

// settingGetResponseZonesWebpJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesWebp]
type settingGetResponseZonesWebpJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesWebpJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesWebp) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesWebpID string

const (
	SettingGetResponseZonesWebpIDWebp SettingGetResponseZonesWebpID = "webp"
)

// Current value of the zone setting.
type SettingGetResponseZonesWebpValue string

const (
	SettingGetResponseZonesWebpValueOff SettingGetResponseZonesWebpValue = "off"
	SettingGetResponseZonesWebpValueOn  SettingGetResponseZonesWebpValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesWebpEditable bool

const (
	SettingGetResponseZonesWebpEditableTrue  SettingGetResponseZonesWebpEditable = true
	SettingGetResponseZonesWebpEditableFalse SettingGetResponseZonesWebpEditable = false
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type SettingGetResponseZonesWebsockets struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesWebsocketsID `json:"id,required"`
	// Current value of the zone setting.
	Value SettingGetResponseZonesWebsocketsValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseZonesWebsocketsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       settingGetResponseZonesWebsocketsJSON `json:"-"`
}

// settingGetResponseZonesWebsocketsJSON contains the JSON metadata for the struct
// [SettingGetResponseZonesWebsockets]
type settingGetResponseZonesWebsocketsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseZonesWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseZonesWebsocketsJSON) RawJSON() string {
	return r.raw
}

func (r SettingGetResponseZonesWebsockets) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type SettingGetResponseZonesWebsocketsID string

const (
	SettingGetResponseZonesWebsocketsIDWebsockets SettingGetResponseZonesWebsocketsID = "websockets"
)

// Current value of the zone setting.
type SettingGetResponseZonesWebsocketsValue string

const (
	SettingGetResponseZonesWebsocketsValueOff SettingGetResponseZonesWebsocketsValue = "off"
	SettingGetResponseZonesWebsocketsValueOn  SettingGetResponseZonesWebsocketsValue = "on"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesWebsocketsEditable bool

const (
	SettingGetResponseZonesWebsocketsEditableTrue  SettingGetResponseZonesWebsocketsEditable = true
	SettingGetResponseZonesWebsocketsEditableFalse SettingGetResponseZonesWebsocketsEditable = false
)

type SettingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// One or more zone setting objects. Must contain an ID and a value.
	Items param.Field[[]SettingEditParamsItem] `json:"items,required"`
}

func (r SettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// 0-RTT session resumption enabled for this zone.
//
// Satisfied by [zones.SettingEditParamsItemsZones0rtt],
// [zones.SettingEditParamsItemsZonesAdvancedDDOS],
// [zones.SettingEditParamsItemsZonesAlwaysOnline],
// [zones.SettingEditParamsItemsZonesAlwaysUseHTTPS],
// [zones.SettingEditParamsItemsZonesAutomaticHTTPSRewrites],
// [zones.SettingEditParamsItemsZonesBrotli],
// [zones.SettingEditParamsItemsZonesBrowserCacheTTL],
// [zones.SettingEditParamsItemsZonesBrowserCheck],
// [zones.SettingEditParamsItemsZonesCacheLevel],
// [zones.SettingEditParamsItemsZonesChallengeTTL],
// [zones.SettingEditParamsItemsZonesCiphers],
// [zones.SettingEditParamsItemsZonesCNAMEFlattening],
// [zones.SettingEditParamsItemsZonesDevelopmentMode],
// [zones.SettingEditParamsItemsZonesEarlyHints],
// [zones.SettingEditParamsItemsZonesEdgeCacheTTL],
// [zones.SettingEditParamsItemsZonesEmailObfuscation],
// [zones.SettingEditParamsItemsZonesH2Prioritization],
// [zones.SettingEditParamsItemsZonesHotlinkProtection],
// [zones.SettingEditParamsItemsZonesHTTP2],
// [zones.SettingEditParamsItemsZonesHTTP3],
// [zones.SettingEditParamsItemsZonesImageResizing],
// [zones.SettingEditParamsItemsZonesIPGeolocation],
// [zones.SettingEditParamsItemsZonesIPV6],
// [zones.SettingEditParamsItemsZonesMaxUpload],
// [zones.SettingEditParamsItemsZonesMinTLSVersion],
// [zones.SettingEditParamsItemsZonesMinify],
// [zones.SettingEditParamsItemsZonesMirage],
// [zones.SettingEditParamsItemsZonesMobileRedirect],
// [zones.SettingEditParamsItemsZonesNEL],
// [zones.SettingEditParamsItemsZonesOpportunisticEncryption],
// [zones.SettingEditParamsItemsZonesOpportunisticOnion],
// [zones.SettingEditParamsItemsZonesOrangeToOrange],
// [zones.SettingEditParamsItemsZonesOriginErrorPagePassThru],
// [zones.SettingEditParamsItemsZonesPolish],
// [zones.SettingEditParamsItemsZonesPrefetchPreload],
// [zones.SettingEditParamsItemsZonesProxyReadTimeout],
// [zones.SettingEditParamsItemsZonesPseudoIPV4],
// [zones.SettingEditParamsItemsZonesResponseBuffering],
// [zones.SettingEditParamsItemsZonesRocketLoader],
// [zones.SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization],
// [zones.SettingEditParamsItemsZonesSecurityHeader],
// [zones.SettingEditParamsItemsZonesSecurityLevel],
// [zones.SettingEditParamsItemsZonesServerSideExclude],
// [zones.SettingEditParamsItemsZonesSha1Support],
// [zones.SettingEditParamsItemsZonesSortQueryStringForCache],
// [zones.SettingEditParamsItemsZonesSSL],
// [zones.SettingEditParamsItemsZonesSSLRecommender],
// [zones.SettingEditParamsItemsZonesTLS1_2Only],
// [zones.SettingEditParamsItemsZonesTLS1_3],
// [zones.SettingEditParamsItemsZonesTLSClientAuth],
// [zones.SettingEditParamsItemsZonesTrueClientIPHeader],
// [zones.SettingEditParamsItemsZonesWAF], [zones.SettingEditParamsItemsZonesWebp],
// [zones.SettingEditParamsItemsZonesWebsockets].
type SettingEditParamsItem interface {
	implementsZonesSettingEditParamsItem()
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

func (r SettingEditParamsItemsZones0rtt) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesAdvancedDDOS) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesAlwaysOnline) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesAlwaysUseHTTPS) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesAutomaticHTTPSRewrites) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesBrotli) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesBrowserCacheTTL) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesBrowserCheck) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesCacheLevel) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesChallengeTTL) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesCiphers) implementsZonesSettingEditParamsItem() {}

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
type SettingEditParamsItemsZonesCNAMEFlattening struct {
	// How to flatten the cname destination.
	ID param.Field[SettingEditParamsItemsZonesCNAMEFlatteningID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[SettingEditParamsItemsZonesCNAMEFlatteningValue] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesCNAMEFlattening) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesCNAMEFlattening) implementsZonesSettingEditParamsItem() {}

// How to flatten the cname destination.
type SettingEditParamsItemsZonesCNAMEFlatteningID string

const (
	SettingEditParamsItemsZonesCNAMEFlatteningIDCNAMEFlattening SettingEditParamsItemsZonesCNAMEFlatteningID = "cname_flattening"
)

// Current value of the zone setting.
type SettingEditParamsItemsZonesCNAMEFlatteningValue string

const (
	SettingEditParamsItemsZonesCNAMEFlatteningValueFlattenAtRoot SettingEditParamsItemsZonesCNAMEFlatteningValue = "flatten_at_root"
	SettingEditParamsItemsZonesCNAMEFlatteningValueFlattenAll    SettingEditParamsItemsZonesCNAMEFlatteningValue = "flatten_all"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesCNAMEFlatteningEditable bool

const (
	SettingEditParamsItemsZonesCNAMEFlatteningEditableTrue  SettingEditParamsItemsZonesCNAMEFlatteningEditable = true
	SettingEditParamsItemsZonesCNAMEFlatteningEditableFalse SettingEditParamsItemsZonesCNAMEFlatteningEditable = false
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

func (r SettingEditParamsItemsZonesDevelopmentMode) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesEarlyHints) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesEdgeCacheTTL) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesEmailObfuscation) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesH2Prioritization) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesHotlinkProtection) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesHTTP2) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesHTTP3) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesImageResizing) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesIPGeolocation) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesIPV6) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesMaxUpload) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesMinTLSVersion) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesMinify) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesMirage) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesMobileRedirect) implementsZonesSettingEditParamsItem() {}

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
	StripURI param.Field[bool] `json:"strip_uri"`
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

func (r SettingEditParamsItemsZonesNEL) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesOpportunisticEncryption) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesOpportunisticOnion) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesOrangeToOrange) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesOriginErrorPagePassThru) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesPolish) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesPrefetchPreload) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesProxyReadTimeout) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesPseudoIPV4) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesResponseBuffering) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesRocketLoader) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization) implementsZonesSettingEditParamsItem() {
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

func (r SettingEditParamsItemsZonesSecurityHeader) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesSecurityLevel) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesServerSideExclude) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesSha1Support) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesSortQueryStringForCache) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesSSL) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesSSLRecommender) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesTLS1_2Only) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesTLS1_3) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesTLSClientAuth) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesTrueClientIPHeader) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesWAF) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesWebp) implementsZonesSettingEditParamsItem() {}

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

func (r SettingEditParamsItemsZonesWebsockets) implementsZonesSettingEditParamsItem() {}

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

func (r settingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r settingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r settingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingGetResponseEnvelope struct {
	Errors   []SettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool                           `json:"success,required"`
	Result  []SettingGetResponse           `json:"result"`
	JSON    settingGetResponseEnvelopeJSON `json:"-"`
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

type SettingGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    settingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SettingGetResponseEnvelopeErrors]
type settingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    settingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SettingGetResponseEnvelopeMessages]
type settingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
