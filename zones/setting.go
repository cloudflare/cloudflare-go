// File generated from our OpenAPI spec by Stainless.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [zones.Zones0rtt], [zones.ZonesAdvancedDDOS],
// [zones.ZonesAlwaysOnline], [zones.ZonesAlwaysUseHTTPS],
// [zones.ZonesAutomaticHTTPSRewrites], [zones.ZonesBrotli],
// [zones.ZonesBrowserCacheTTL], [zones.ZonesBrowserCheck],
// [zones.ZonesCacheLevel], [zones.ZonesChallengeTTL], [zones.ZonesCiphers],
// [zones.SettingEditResponseZonesCNAMEFlattening], [zones.ZonesDevelopmentMode],
// [zones.ZonesEarlyHints], [zones.SettingEditResponseZonesEdgeCacheTTL],
// [zones.ZonesEmailObfuscation], [zones.ZonesH2Prioritization],
// [zones.ZonesHotlinkProtection], [zones.ZonesHTTP2], [zones.ZonesHTTP3],
// [zones.ZonesImageResizing], [zones.ZonesIPGeolocation], [zones.ZonesIPV6],
// [zones.SettingEditResponseZonesMaxUpload], [zones.ZonesMinTLSVersion],
// [zones.ZonesMinify], [zones.ZonesMirage], [zones.ZonesMobileRedirect],
// [zones.ZonesNEL], [zones.ZonesOpportunisticEncryption],
// [zones.ZonesOpportunisticOnion], [zones.ZonesOrangeToOrange],
// [zones.ZonesOriginErrorPagePassThru], [zones.ZonesPolish],
// [zones.ZonesPrefetchPreload], [zones.ZonesProxyReadTimeout],
// [zones.ZonesPseudoIPV4], [zones.ZonesBuffering], [zones.ZonesRocketLoader],
// [zones.SettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.ZonesSecurityHeader], [zones.ZonesSecurityLevel],
// [zones.ZonesServerSideExclude], [zones.SettingEditResponseZonesSha1Support],
// [zones.ZonesSortQueryStringForCache], [zones.ZonesSSL],
// [zones.ZonesSSLRecommender], [zones.SettingEditResponseZonesTLS1_2Only],
// [zones.ZonesTLS1_3], [zones.ZonesTLSClientAuth],
// [zones.ZonesTrueClientIPHeader], [zones.ZonesWAF], [zones.ZonesWebp] or
// [zones.ZonesWebsockets].
type SettingEditResponse interface {
	implementsZonesSettingEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Zones0rtt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAdvancedDDOS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesBrotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesBrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCiphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesCNAMEFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesDevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesEarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesEmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesH2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHTTP2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHTTP3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesIPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesIPV6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMinTLSVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMinify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMobileRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesNEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPolish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPseudoIPV4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesRocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSecurityHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesServerSideExclude{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSSL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSSLRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesTLS1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTLS1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTLSClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesWAF{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesWebp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesWebsockets{}),
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

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingEditResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesAutomaticPlatformOptimization `json:"value,required"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingEditResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
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

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [zones.Zones0rtt], [zones.ZonesAdvancedDDOS],
// [zones.ZonesAlwaysOnline], [zones.ZonesAlwaysUseHTTPS],
// [zones.ZonesAutomaticHTTPSRewrites], [zones.ZonesBrotli],
// [zones.ZonesBrowserCacheTTL], [zones.ZonesBrowserCheck],
// [zones.ZonesCacheLevel], [zones.ZonesChallengeTTL], [zones.ZonesCiphers],
// [zones.SettingGetResponseZonesCNAMEFlattening], [zones.ZonesDevelopmentMode],
// [zones.ZonesEarlyHints], [zones.SettingGetResponseZonesEdgeCacheTTL],
// [zones.ZonesEmailObfuscation], [zones.ZonesH2Prioritization],
// [zones.ZonesHotlinkProtection], [zones.ZonesHTTP2], [zones.ZonesHTTP3],
// [zones.ZonesImageResizing], [zones.ZonesIPGeolocation], [zones.ZonesIPV6],
// [zones.SettingGetResponseZonesMaxUpload], [zones.ZonesMinTLSVersion],
// [zones.ZonesMinify], [zones.ZonesMirage], [zones.ZonesMobileRedirect],
// [zones.ZonesNEL], [zones.ZonesOpportunisticEncryption],
// [zones.ZonesOpportunisticOnion], [zones.ZonesOrangeToOrange],
// [zones.ZonesOriginErrorPagePassThru], [zones.ZonesPolish],
// [zones.ZonesPrefetchPreload], [zones.ZonesProxyReadTimeout],
// [zones.ZonesPseudoIPV4], [zones.ZonesBuffering], [zones.ZonesRocketLoader],
// [zones.SettingGetResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.ZonesSecurityHeader], [zones.ZonesSecurityLevel],
// [zones.ZonesServerSideExclude], [zones.SettingGetResponseZonesSha1Support],
// [zones.ZonesSortQueryStringForCache], [zones.ZonesSSL],
// [zones.ZonesSSLRecommender], [zones.SettingGetResponseZonesTLS1_2Only],
// [zones.ZonesTLS1_3], [zones.ZonesTLSClientAuth],
// [zones.ZonesTrueClientIPHeader], [zones.ZonesWAF], [zones.ZonesWebp] or
// [zones.ZonesWebsockets].
type SettingGetResponse interface {
	implementsZonesSettingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Zones0rtt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAdvancedDDOS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesBrotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesBrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCiphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesCNAMEFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesDevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesEarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesEmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesH2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHTTP2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHTTP3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesIPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesIPV6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMinTLSVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMinify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMobileRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesNEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPolish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPseudoIPV4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesRocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSecurityHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesServerSideExclude{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSSL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSSLRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesTLS1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTLS1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTLSClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesWAF{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesWebp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesWebsockets{}),
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

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingGetResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesAutomaticPlatformOptimization `json:"value,required"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingGetResponseZonesSchemasAutomaticPlatformOptimizationEditable = false
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
// Satisfied by [zones.Zones0rttParam], [zones.ZonesAdvancedDDOSParam],
// [zones.ZonesAlwaysOnlineParam], [zones.ZonesAlwaysUseHTTPSParam],
// [zones.ZonesAutomaticHTTPSRewritesParam], [zones.ZonesBrotliParam],
// [zones.ZonesBrowserCacheTTLParam], [zones.ZonesBrowserCheckParam],
// [zones.ZonesCacheLevelParam], [zones.ZonesChallengeTTLParam],
// [zones.ZonesCiphersParam], [zones.SettingEditParamsItemsZonesCNAMEFlattening],
// [zones.ZonesDevelopmentModeParam], [zones.ZonesEarlyHintsParam],
// [zones.SettingEditParamsItemsZonesEdgeCacheTTL],
// [zones.ZonesEmailObfuscationParam], [zones.ZonesH2PrioritizationParam],
// [zones.ZonesHotlinkProtectionParam], [zones.ZonesHTTP2Param],
// [zones.ZonesHTTP3Param], [zones.ZonesImageResizingParam],
// [zones.ZonesIPGeolocationParam], [zones.ZonesIPV6Param],
// [zones.SettingEditParamsItemsZonesMaxUpload], [zones.ZonesMinTLSVersionParam],
// [zones.ZonesMinifyParam], [zones.ZonesMirageParam],
// [zones.ZonesMobileRedirectParam], [zones.ZonesNELParam],
// [zones.ZonesOpportunisticEncryptionParam], [zones.ZonesOpportunisticOnionParam],
// [zones.ZonesOrangeToOrangeParam], [zones.ZonesOriginErrorPagePassThruParam],
// [zones.ZonesPolishParam], [zones.ZonesPrefetchPreloadParam],
// [zones.ZonesProxyReadTimeoutParam], [zones.ZonesPseudoIPV4Param],
// [zones.ZonesBufferingParam], [zones.ZonesRocketLoaderParam],
// [zones.SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization],
// [zones.ZonesSecurityHeaderParam], [zones.ZonesSecurityLevelParam],
// [zones.ZonesServerSideExcludeParam],
// [zones.SettingEditParamsItemsZonesSha1Support],
// [zones.ZonesSortQueryStringForCacheParam], [zones.ZonesSSLParam],
// [zones.ZonesSSLRecommenderParam], [zones.SettingEditParamsItemsZonesTLS1_2Only],
// [zones.ZonesTLS1_3Param], [zones.ZonesTLSClientAuthParam],
// [zones.ZonesTrueClientIPHeaderParam], [zones.ZonesWAFParam],
// [zones.ZonesWebpParam], [zones.ZonesWebsocketsParam].
type SettingEditParamsItem interface {
	implementsZonesSettingEditParamsItem()
}

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

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesAutomaticPlatformOptimizationParam] `json:"value,required"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable = false
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
