// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
	AdvancedDDoS                  *SettingAdvancedDDoSService
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
	WebP                          *SettingWebPService
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
	r.AdvancedDDoS = NewSettingAdvancedDDoSService(opts...)
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
	r.WebP = NewSettingWebPService(opts...)
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
type SettingEditResponse struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingEditResponseEditable `json:"editable"`
	// ID of the zone setting.
	ID SettingEditResponseID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time   `json:"modified_on,nullable" format:"date-time"`
	Value      interface{} `json:"value,required"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// ssl-recommender enrollment setting.
	Enabled bool                    `json:"enabled"`
	JSON    settingEditResponseJSON `json:"-"`
	union   SettingEditResponseUnion
}

// settingEditResponseJSON contains the JSON metadata for the struct
// [SettingEditResponse]
type settingEditResponseJSON struct {
	Editable      apijson.Field
	ID            apijson.Field
	ModifiedOn    apijson.Field
	Value         apijson.Field
	TimeRemaining apijson.Field
	Enabled       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r settingEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r SettingEditResponse) AsUnion() SettingEditResponseUnion {
	return r.union
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [zones.ZoneSetting0rtt], [zones.ZoneSettingAdvancedDDoS],
// [zones.ZoneSettingAlwaysOnline], [zones.ZoneSettingAlwaysUseHTTPS],
// [zones.ZoneSettingAutomaticHTTPSRewrites], [zones.ZoneSettingBrotli],
// [zones.ZoneSettingBrowserCacheTTL], [zones.ZoneSettingBrowserCheck],
// [zones.ZoneSettingCacheLevel], [zones.ZoneSettingChallengeTTL],
// [zones.ZoneSettingCiphers], [zones.SettingEditResponseZonesCNAMEFlattening],
// [zones.ZoneSettingDevelopmentMode], [zones.ZoneSettingEarlyHints],
// [zones.SettingEditResponseZonesEdgeCacheTTL],
// [zones.ZoneSettingEmailObfuscation], [zones.ZoneSettingH2Prioritization],
// [zones.ZoneSettingHotlinkProtection], [zones.ZoneSettingHTTP2],
// [zones.ZoneSettingHTTP3], [zones.ZoneSettingImageResizing],
// [zones.ZoneSettingIPGeolocation], [zones.ZoneSettingIPV6],
// [zones.SettingEditResponseZonesMaxUpload], [zones.ZoneSettingMinTLSVersion],
// [zones.ZoneSettingMinify], [zones.ZoneSettingMirage],
// [zones.ZoneSettingMobileRedirect], [zones.ZoneSettingNEL],
// [zones.ZoneSettingOpportunisticEncryption],
// [zones.ZoneSettingOpportunisticOnion], [zones.ZoneSettingOrangeToOrange],
// [zones.ZoneSettingOriginErrorPagePassThru], [zones.ZoneSettingPolish],
// [zones.ZoneSettingPrefetchPreload], [zones.ZoneSettingProxyReadTimeout],
// [zones.ZoneSettingPseudoIPV4], [zones.ZoneSettingBuffering],
// [zones.ZoneSettingRocketLoader],
// [zones.SettingEditResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.ZoneSettingSecurityHeader], [zones.ZoneSettingSecurityLevel],
// [zones.ZoneSettingServerSideExclude],
// [zones.SettingEditResponseZonesSha1Support],
// [zones.ZoneSettingSortQueryStringForCache], [zones.ZoneSettingSSL],
// [zones.ZoneSettingSSLRecommender], [zones.SettingEditResponseZonesTLS1_2Only],
// [zones.ZoneSettingTLS1_3], [zones.ZoneSettingTLSClientAuth],
// [zones.ZoneSettingTrueClientIPHeader], [zones.ZoneSettingWAF],
// [zones.ZoneSettingWebP] or [zones.ZoneSettingWebsockets].
type SettingEditResponseUnion interface {
	implementsZonesSettingEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSetting0rtt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingAdvancedDDoS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingAlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingBrotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingBrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingCacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingCiphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesCNAMEFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingDevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingH2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingHotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingHTTP2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingHTTP3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingIPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingIPV6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingMinTLSVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingMinify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingMobileRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingNEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingOpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingOpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingOrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingPolish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingPrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingPseudoIPV4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingRocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingSecurityHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingSecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingServerSideExclude{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingSSL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingSSLRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingEditResponseZonesTLS1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingTLS1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingTLSClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingWAF{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingWebP{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingWebsockets{}),
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

func (r SettingEditResponseZonesEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

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

func (r SettingEditResponseZonesEdgeCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEdgeCacheTTLValue30, SettingEditResponseZonesEdgeCacheTTLValue60, SettingEditResponseZonesEdgeCacheTTLValue300, SettingEditResponseZonesEdgeCacheTTLValue1200, SettingEditResponseZonesEdgeCacheTTLValue1800, SettingEditResponseZonesEdgeCacheTTLValue3600, SettingEditResponseZonesEdgeCacheTTLValue7200, SettingEditResponseZonesEdgeCacheTTLValue10800, SettingEditResponseZonesEdgeCacheTTLValue14400, SettingEditResponseZonesEdgeCacheTTLValue18000, SettingEditResponseZonesEdgeCacheTTLValue28800, SettingEditResponseZonesEdgeCacheTTLValue43200, SettingEditResponseZonesEdgeCacheTTLValue57600, SettingEditResponseZonesEdgeCacheTTLValue72000, SettingEditResponseZonesEdgeCacheTTLValue86400, SettingEditResponseZonesEdgeCacheTTLValue172800, SettingEditResponseZonesEdgeCacheTTLValue259200, SettingEditResponseZonesEdgeCacheTTLValue345600, SettingEditResponseZonesEdgeCacheTTLValue432000, SettingEditResponseZonesEdgeCacheTTLValue518400, SettingEditResponseZonesEdgeCacheTTLValue604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditResponseZonesEdgeCacheTTLEditable bool

const (
	SettingEditResponseZonesEdgeCacheTTLEditableTrue  SettingEditResponseZonesEdgeCacheTTLEditable = true
	SettingEditResponseZonesEdgeCacheTTLEditableFalse SettingEditResponseZonesEdgeCacheTTLEditable = false
)

func (r SettingEditResponseZonesEdgeCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingEditResponseZonesEdgeCacheTTLEditableTrue, SettingEditResponseZonesEdgeCacheTTLEditableFalse:
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

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingEditResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingEditResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAutomaticPlatformOptimization `json:"value,required"`
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
	SettingEditResponseIDMinify                        SettingEditResponseID = "minify"
	SettingEditResponseIDMirage                        SettingEditResponseID = "mirage"
	SettingEditResponseIDMobileRedirect                SettingEditResponseID = "mobile_redirect"
	SettingEditResponseIDNEL                           SettingEditResponseID = "nel"
	SettingEditResponseIDOpportunisticEncryption       SettingEditResponseID = "opportunistic_encryption"
	SettingEditResponseIDOpportunisticOnion            SettingEditResponseID = "opportunistic_onion"
	SettingEditResponseIDOrangeToOrange                SettingEditResponseID = "orange_to_orange"
	SettingEditResponseIDOriginErrorPagePassThru       SettingEditResponseID = "origin_error_page_pass_thru"
	SettingEditResponseIDPolish                        SettingEditResponseID = "polish"
	SettingEditResponseIDPrefetchPreload               SettingEditResponseID = "prefetch_preload"
	SettingEditResponseIDProxyReadTimeout              SettingEditResponseID = "proxy_read_timeout"
	SettingEditResponseIDPseudoIPV4                    SettingEditResponseID = "pseudo_ipv4"
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
	case SettingEditResponseID0rtt, SettingEditResponseIDAdvancedDDoS, SettingEditResponseIDAlwaysOnline, SettingEditResponseIDAlwaysUseHTTPS, SettingEditResponseIDAutomaticHTTPSRewrites, SettingEditResponseIDBrotli, SettingEditResponseIDBrowserCacheTTL, SettingEditResponseIDBrowserCheck, SettingEditResponseIDCacheLevel, SettingEditResponseIDChallengeTTL, SettingEditResponseIDCiphers, SettingEditResponseIDCNAMEFlattening, SettingEditResponseIDDevelopmentMode, SettingEditResponseIDEarlyHints, SettingEditResponseIDEdgeCacheTTL, SettingEditResponseIDEmailObfuscation, SettingEditResponseIDH2Prioritization, SettingEditResponseIDHotlinkProtection, SettingEditResponseIDHTTP2, SettingEditResponseIDHTTP3, SettingEditResponseIDImageResizing, SettingEditResponseIDIPGeolocation, SettingEditResponseIDIPV6, SettingEditResponseIDMaxUpload, SettingEditResponseIDMinTLSVersion, SettingEditResponseIDMinify, SettingEditResponseIDMirage, SettingEditResponseIDMobileRedirect, SettingEditResponseIDNEL, SettingEditResponseIDOpportunisticEncryption, SettingEditResponseIDOpportunisticOnion, SettingEditResponseIDOrangeToOrange, SettingEditResponseIDOriginErrorPagePassThru, SettingEditResponseIDPolish, SettingEditResponseIDPrefetchPreload, SettingEditResponseIDProxyReadTimeout, SettingEditResponseIDPseudoIPV4, SettingEditResponseIDResponseBuffering, SettingEditResponseIDRocketLoader, SettingEditResponseIDAutomaticPlatformOptimization, SettingEditResponseIDSecurityHeader, SettingEditResponseIDSecurityLevel, SettingEditResponseIDServerSideExclude, SettingEditResponseIDSha1Support, SettingEditResponseIDSortQueryStringForCache, SettingEditResponseIDSSL, SettingEditResponseIDSSLRecommender, SettingEditResponseIDTLS1_2Only, SettingEditResponseIDTLS1_3, SettingEditResponseIDTLSClientAuth, SettingEditResponseIDTrueClientIPHeader, SettingEditResponseIDWAF, SettingEditResponseIDWebP, SettingEditResponseIDWebsockets:
		return true
	}
	return false
}

// 0-RTT session resumption enabled for this zone.
type SettingGetResponse struct {
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SettingGetResponseEditable `json:"editable"`
	// ID of the zone setting.
	ID SettingGetResponseID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time   `json:"modified_on,nullable" format:"date-time"`
	Value      interface{} `json:"value,required"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// ssl-recommender enrollment setting.
	Enabled bool                   `json:"enabled"`
	JSON    settingGetResponseJSON `json:"-"`
	union   SettingGetResponseUnion
}

// settingGetResponseJSON contains the JSON metadata for the struct
// [SettingGetResponse]
type settingGetResponseJSON struct {
	Editable      apijson.Field
	ID            apijson.Field
	ModifiedOn    apijson.Field
	Value         apijson.Field
	TimeRemaining apijson.Field
	Enabled       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r settingGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *SettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r SettingGetResponse) AsUnion() SettingGetResponseUnion {
	return r.union
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [zones.ZoneSetting0rtt], [zones.ZoneSettingAdvancedDDoS],
// [zones.ZoneSettingAlwaysOnline], [zones.ZoneSettingAlwaysUseHTTPS],
// [zones.ZoneSettingAutomaticHTTPSRewrites], [zones.ZoneSettingBrotli],
// [zones.ZoneSettingBrowserCacheTTL], [zones.ZoneSettingBrowserCheck],
// [zones.ZoneSettingCacheLevel], [zones.ZoneSettingChallengeTTL],
// [zones.ZoneSettingCiphers], [zones.SettingGetResponseZonesCNAMEFlattening],
// [zones.ZoneSettingDevelopmentMode], [zones.ZoneSettingEarlyHints],
// [zones.SettingGetResponseZonesEdgeCacheTTL],
// [zones.ZoneSettingEmailObfuscation], [zones.ZoneSettingH2Prioritization],
// [zones.ZoneSettingHotlinkProtection], [zones.ZoneSettingHTTP2],
// [zones.ZoneSettingHTTP3], [zones.ZoneSettingImageResizing],
// [zones.ZoneSettingIPGeolocation], [zones.ZoneSettingIPV6],
// [zones.SettingGetResponseZonesMaxUpload], [zones.ZoneSettingMinTLSVersion],
// [zones.ZoneSettingMinify], [zones.ZoneSettingMirage],
// [zones.ZoneSettingMobileRedirect], [zones.ZoneSettingNEL],
// [zones.ZoneSettingOpportunisticEncryption],
// [zones.ZoneSettingOpportunisticOnion], [zones.ZoneSettingOrangeToOrange],
// [zones.ZoneSettingOriginErrorPagePassThru], [zones.ZoneSettingPolish],
// [zones.ZoneSettingPrefetchPreload], [zones.ZoneSettingProxyReadTimeout],
// [zones.ZoneSettingPseudoIPV4], [zones.ZoneSettingBuffering],
// [zones.ZoneSettingRocketLoader],
// [zones.SettingGetResponseZonesSchemasAutomaticPlatformOptimization],
// [zones.ZoneSettingSecurityHeader], [zones.ZoneSettingSecurityLevel],
// [zones.ZoneSettingServerSideExclude],
// [zones.SettingGetResponseZonesSha1Support],
// [zones.ZoneSettingSortQueryStringForCache], [zones.ZoneSettingSSL],
// [zones.ZoneSettingSSLRecommender], [zones.SettingGetResponseZonesTLS1_2Only],
// [zones.ZoneSettingTLS1_3], [zones.ZoneSettingTLSClientAuth],
// [zones.ZoneSettingTrueClientIPHeader], [zones.ZoneSettingWAF],
// [zones.ZoneSettingWebP] or [zones.ZoneSettingWebsockets].
type SettingGetResponseUnion interface {
	implementsZonesSettingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSetting0rtt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingAdvancedDDoS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingAlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingBrotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingBrowserCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingCacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingChallengeTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingCiphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesCNAMEFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingDevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesEdgeCacheTTL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingEmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingH2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingHotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingHTTP2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingHTTP3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingIPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingIPV6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingMinTLSVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingMinify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingMobileRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingNEL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingOpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingOpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingOrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingPolish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingPrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingPseudoIPV4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingRocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingSecurityHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingSecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingServerSideExclude{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingSSL{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingSSLRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingGetResponseZonesTLS1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingTLS1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingTLSClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingWAF{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingWebP{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneSettingWebsockets{}),
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

func (r SettingGetResponseZonesEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

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

func (r SettingGetResponseZonesEdgeCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEdgeCacheTTLValue30, SettingGetResponseZonesEdgeCacheTTLValue60, SettingGetResponseZonesEdgeCacheTTLValue300, SettingGetResponseZonesEdgeCacheTTLValue1200, SettingGetResponseZonesEdgeCacheTTLValue1800, SettingGetResponseZonesEdgeCacheTTLValue3600, SettingGetResponseZonesEdgeCacheTTLValue7200, SettingGetResponseZonesEdgeCacheTTLValue10800, SettingGetResponseZonesEdgeCacheTTLValue14400, SettingGetResponseZonesEdgeCacheTTLValue18000, SettingGetResponseZonesEdgeCacheTTLValue28800, SettingGetResponseZonesEdgeCacheTTLValue43200, SettingGetResponseZonesEdgeCacheTTLValue57600, SettingGetResponseZonesEdgeCacheTTLValue72000, SettingGetResponseZonesEdgeCacheTTLValue86400, SettingGetResponseZonesEdgeCacheTTLValue172800, SettingGetResponseZonesEdgeCacheTTLValue259200, SettingGetResponseZonesEdgeCacheTTLValue345600, SettingGetResponseZonesEdgeCacheTTLValue432000, SettingGetResponseZonesEdgeCacheTTLValue518400, SettingGetResponseZonesEdgeCacheTTLValue604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingGetResponseZonesEdgeCacheTTLEditable bool

const (
	SettingGetResponseZonesEdgeCacheTTLEditableTrue  SettingGetResponseZonesEdgeCacheTTLEditable = true
	SettingGetResponseZonesEdgeCacheTTLEditableFalse SettingGetResponseZonesEdgeCacheTTLEditable = false
)

func (r SettingGetResponseZonesEdgeCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingGetResponseZonesEdgeCacheTTLEditableTrue, SettingGetResponseZonesEdgeCacheTTLEditableFalse:
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

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingGetResponseZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID SettingGetResponseZonesSchemasAutomaticPlatformOptimizationID `json:"id,required"`
	// Current value of the zone setting.
	Value ZoneSettingAutomaticPlatformOptimization `json:"value,required"`
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
	SettingGetResponseIDMinify                        SettingGetResponseID = "minify"
	SettingGetResponseIDMirage                        SettingGetResponseID = "mirage"
	SettingGetResponseIDMobileRedirect                SettingGetResponseID = "mobile_redirect"
	SettingGetResponseIDNEL                           SettingGetResponseID = "nel"
	SettingGetResponseIDOpportunisticEncryption       SettingGetResponseID = "opportunistic_encryption"
	SettingGetResponseIDOpportunisticOnion            SettingGetResponseID = "opportunistic_onion"
	SettingGetResponseIDOrangeToOrange                SettingGetResponseID = "orange_to_orange"
	SettingGetResponseIDOriginErrorPagePassThru       SettingGetResponseID = "origin_error_page_pass_thru"
	SettingGetResponseIDPolish                        SettingGetResponseID = "polish"
	SettingGetResponseIDPrefetchPreload               SettingGetResponseID = "prefetch_preload"
	SettingGetResponseIDProxyReadTimeout              SettingGetResponseID = "proxy_read_timeout"
	SettingGetResponseIDPseudoIPV4                    SettingGetResponseID = "pseudo_ipv4"
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
	case SettingGetResponseID0rtt, SettingGetResponseIDAdvancedDDoS, SettingGetResponseIDAlwaysOnline, SettingGetResponseIDAlwaysUseHTTPS, SettingGetResponseIDAutomaticHTTPSRewrites, SettingGetResponseIDBrotli, SettingGetResponseIDBrowserCacheTTL, SettingGetResponseIDBrowserCheck, SettingGetResponseIDCacheLevel, SettingGetResponseIDChallengeTTL, SettingGetResponseIDCiphers, SettingGetResponseIDCNAMEFlattening, SettingGetResponseIDDevelopmentMode, SettingGetResponseIDEarlyHints, SettingGetResponseIDEdgeCacheTTL, SettingGetResponseIDEmailObfuscation, SettingGetResponseIDH2Prioritization, SettingGetResponseIDHotlinkProtection, SettingGetResponseIDHTTP2, SettingGetResponseIDHTTP3, SettingGetResponseIDImageResizing, SettingGetResponseIDIPGeolocation, SettingGetResponseIDIPV6, SettingGetResponseIDMaxUpload, SettingGetResponseIDMinTLSVersion, SettingGetResponseIDMinify, SettingGetResponseIDMirage, SettingGetResponseIDMobileRedirect, SettingGetResponseIDNEL, SettingGetResponseIDOpportunisticEncryption, SettingGetResponseIDOpportunisticOnion, SettingGetResponseIDOrangeToOrange, SettingGetResponseIDOriginErrorPagePassThru, SettingGetResponseIDPolish, SettingGetResponseIDPrefetchPreload, SettingGetResponseIDProxyReadTimeout, SettingGetResponseIDPseudoIPV4, SettingGetResponseIDResponseBuffering, SettingGetResponseIDRocketLoader, SettingGetResponseIDAutomaticPlatformOptimization, SettingGetResponseIDSecurityHeader, SettingGetResponseIDSecurityLevel, SettingGetResponseIDServerSideExclude, SettingGetResponseIDSha1Support, SettingGetResponseIDSortQueryStringForCache, SettingGetResponseIDSSL, SettingGetResponseIDSSLRecommender, SettingGetResponseIDTLS1_2Only, SettingGetResponseIDTLS1_3, SettingGetResponseIDTLSClientAuth, SettingGetResponseIDTrueClientIPHeader, SettingGetResponseIDWAF, SettingGetResponseIDWebP, SettingGetResponseIDWebsockets:
		return true
	}
	return false
}

type SettingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// One or more zone setting objects. Must contain an ID and a value.
	Items param.Field[[]SettingEditParamsItemUnion] `json:"items,required"`
}

func (r SettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// 0-RTT session resumption enabled for this zone.
type SettingEditParamsItem struct {
	// ID of the zone setting.
	ID    param.Field[SettingEditParamsItemsID] `json:"id"`
	Value param.Field[interface{}]              `json:"value,required"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r SettingEditParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItem) implementsZonesSettingEditParamsItemUnion() {}

// 0-RTT session resumption enabled for this zone.
//
// Satisfied by [zones.ZoneSetting0rttParam], [zones.ZoneSettingAdvancedDDoSParam],
// [zones.ZoneSettingAlwaysOnlineParam], [zones.ZoneSettingAlwaysUseHTTPSParam],
// [zones.ZoneSettingAutomaticHTTPSRewritesParam], [zones.ZoneSettingBrotliParam],
// [zones.ZoneSettingBrowserCacheTTLParam], [zones.ZoneSettingBrowserCheckParam],
// [zones.ZoneSettingCacheLevelParam], [zones.ZoneSettingChallengeTTLParam],
// [zones.ZoneSettingCiphersParam],
// [zones.SettingEditParamsItemsZonesCNAMEFlattening],
// [zones.ZoneSettingDevelopmentModeParam], [zones.ZoneSettingEarlyHintsParam],
// [zones.SettingEditParamsItemsZonesEdgeCacheTTL],
// [zones.ZoneSettingEmailObfuscationParam],
// [zones.ZoneSettingH2PrioritizationParam],
// [zones.ZoneSettingHotlinkProtectionParam], [zones.ZoneSettingHTTP2Param],
// [zones.ZoneSettingHTTP3Param], [zones.ZoneSettingImageResizingParam],
// [zones.ZoneSettingIPGeolocationParam], [zones.ZoneSettingIPV6Param],
// [zones.SettingEditParamsItemsZonesMaxUpload],
// [zones.ZoneSettingMinTLSVersionParam], [zones.ZoneSettingMinifyParam],
// [zones.ZoneSettingMirageParam], [zones.ZoneSettingMobileRedirectParam],
// [zones.ZoneSettingNELParam], [zones.ZoneSettingOpportunisticEncryptionParam],
// [zones.ZoneSettingOpportunisticOnionParam],
// [zones.ZoneSettingOrangeToOrangeParam],
// [zones.ZoneSettingOriginErrorPagePassThruParam], [zones.ZoneSettingPolishParam],
// [zones.ZoneSettingPrefetchPreloadParam],
// [zones.ZoneSettingProxyReadTimeoutParam], [zones.ZoneSettingPseudoIPV4Param],
// [zones.ZoneSettingBufferingParam], [zones.ZoneSettingRocketLoaderParam],
// [zones.SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization],
// [zones.ZoneSettingSecurityHeaderParam], [zones.ZoneSettingSecurityLevelParam],
// [zones.ZoneSettingServerSideExcludeParam],
// [zones.SettingEditParamsItemsZonesSha1Support],
// [zones.ZoneSettingSortQueryStringForCacheParam], [zones.ZoneSettingSSLParam],
// [zones.ZoneSettingSSLRecommenderParam],
// [zones.SettingEditParamsItemsZonesTLS1_2Only], [zones.ZoneSettingTLS1_3Param],
// [zones.ZoneSettingTLSClientAuthParam],
// [zones.ZoneSettingTrueClientIPHeaderParam], [zones.ZoneSettingWAFParam],
// [zones.ZoneSettingWebPParam], [zones.ZoneSettingWebsocketsParam],
// [SettingEditParamsItem].
type SettingEditParamsItemUnion interface {
	implementsZonesSettingEditParamsItemUnion()
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

func (r SettingEditParamsItemsZonesCNAMEFlattening) implementsZonesSettingEditParamsItemUnion() {}

// How to flatten the cname destination.
type SettingEditParamsItemsZonesCNAMEFlatteningID string

const (
	SettingEditParamsItemsZonesCNAMEFlatteningIDCNAMEFlattening SettingEditParamsItemsZonesCNAMEFlatteningID = "cname_flattening"
)

func (r SettingEditParamsItemsZonesCNAMEFlatteningID) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesCNAMEFlatteningIDCNAMEFlattening:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsItemsZonesCNAMEFlatteningValue string

const (
	SettingEditParamsItemsZonesCNAMEFlatteningValueFlattenAtRoot SettingEditParamsItemsZonesCNAMEFlatteningValue = "flatten_at_root"
	SettingEditParamsItemsZonesCNAMEFlatteningValueFlattenAll    SettingEditParamsItemsZonesCNAMEFlatteningValue = "flatten_all"
)

func (r SettingEditParamsItemsZonesCNAMEFlatteningValue) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesCNAMEFlatteningValueFlattenAtRoot, SettingEditParamsItemsZonesCNAMEFlatteningValueFlattenAll:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesCNAMEFlatteningEditable bool

const (
	SettingEditParamsItemsZonesCNAMEFlatteningEditableTrue  SettingEditParamsItemsZonesCNAMEFlatteningEditable = true
	SettingEditParamsItemsZonesCNAMEFlatteningEditableFalse SettingEditParamsItemsZonesCNAMEFlatteningEditable = false
)

func (r SettingEditParamsItemsZonesCNAMEFlatteningEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesCNAMEFlatteningEditableTrue, SettingEditParamsItemsZonesCNAMEFlatteningEditableFalse:
		return true
	}
	return false
}

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

func (r SettingEditParamsItemsZonesEdgeCacheTTL) implementsZonesSettingEditParamsItemUnion() {}

// ID of the zone setting.
type SettingEditParamsItemsZonesEdgeCacheTTLID string

const (
	SettingEditParamsItemsZonesEdgeCacheTTLIDEdgeCacheTTL SettingEditParamsItemsZonesEdgeCacheTTLID = "edge_cache_ttl"
)

func (r SettingEditParamsItemsZonesEdgeCacheTTLID) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesEdgeCacheTTLIDEdgeCacheTTL:
		return true
	}
	return false
}

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

func (r SettingEditParamsItemsZonesEdgeCacheTTLValue) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesEdgeCacheTTLValue30, SettingEditParamsItemsZonesEdgeCacheTTLValue60, SettingEditParamsItemsZonesEdgeCacheTTLValue300, SettingEditParamsItemsZonesEdgeCacheTTLValue1200, SettingEditParamsItemsZonesEdgeCacheTTLValue1800, SettingEditParamsItemsZonesEdgeCacheTTLValue3600, SettingEditParamsItemsZonesEdgeCacheTTLValue7200, SettingEditParamsItemsZonesEdgeCacheTTLValue10800, SettingEditParamsItemsZonesEdgeCacheTTLValue14400, SettingEditParamsItemsZonesEdgeCacheTTLValue18000, SettingEditParamsItemsZonesEdgeCacheTTLValue28800, SettingEditParamsItemsZonesEdgeCacheTTLValue43200, SettingEditParamsItemsZonesEdgeCacheTTLValue57600, SettingEditParamsItemsZonesEdgeCacheTTLValue72000, SettingEditParamsItemsZonesEdgeCacheTTLValue86400, SettingEditParamsItemsZonesEdgeCacheTTLValue172800, SettingEditParamsItemsZonesEdgeCacheTTLValue259200, SettingEditParamsItemsZonesEdgeCacheTTLValue345600, SettingEditParamsItemsZonesEdgeCacheTTLValue432000, SettingEditParamsItemsZonesEdgeCacheTTLValue518400, SettingEditParamsItemsZonesEdgeCacheTTLValue604800:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesEdgeCacheTTLEditable bool

const (
	SettingEditParamsItemsZonesEdgeCacheTTLEditableTrue  SettingEditParamsItemsZonesEdgeCacheTTLEditable = true
	SettingEditParamsItemsZonesEdgeCacheTTLEditableFalse SettingEditParamsItemsZonesEdgeCacheTTLEditable = false
)

func (r SettingEditParamsItemsZonesEdgeCacheTTLEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesEdgeCacheTTLEditableTrue, SettingEditParamsItemsZonesEdgeCacheTTLEditableFalse:
		return true
	}
	return false
}

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

func (r SettingEditParamsItemsZonesMaxUpload) implementsZonesSettingEditParamsItemUnion() {}

// identifier of the zone setting.
type SettingEditParamsItemsZonesMaxUploadID string

const (
	SettingEditParamsItemsZonesMaxUploadIDMaxUpload SettingEditParamsItemsZonesMaxUploadID = "max_upload"
)

func (r SettingEditParamsItemsZonesMaxUploadID) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesMaxUploadIDMaxUpload:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsItemsZonesMaxUploadValue float64

const (
	SettingEditParamsItemsZonesMaxUploadValue100 SettingEditParamsItemsZonesMaxUploadValue = 100
	SettingEditParamsItemsZonesMaxUploadValue200 SettingEditParamsItemsZonesMaxUploadValue = 200
	SettingEditParamsItemsZonesMaxUploadValue500 SettingEditParamsItemsZonesMaxUploadValue = 500
)

func (r SettingEditParamsItemsZonesMaxUploadValue) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesMaxUploadValue100, SettingEditParamsItemsZonesMaxUploadValue200, SettingEditParamsItemsZonesMaxUploadValue500:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesMaxUploadEditable bool

const (
	SettingEditParamsItemsZonesMaxUploadEditableTrue  SettingEditParamsItemsZonesMaxUploadEditable = true
	SettingEditParamsItemsZonesMaxUploadEditableFalse SettingEditParamsItemsZonesMaxUploadEditable = false
)

func (r SettingEditParamsItemsZonesMaxUploadEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesMaxUploadEditableTrue, SettingEditParamsItemsZonesMaxUploadEditableFalse:
		return true
	}
	return false
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID param.Field[SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZoneSettingAutomaticPlatformOptimizationParam] `json:"value,required"`
}

func (r SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization) implementsZonesSettingEditParamsItemUnion() {
}

// ID of the zone setting.
type SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID string

const (
	SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

func (r SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableTrue  SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable = true
	SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableFalse SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable = false
)

func (r SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableTrue, SettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationEditableFalse:
		return true
	}
	return false
}

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

func (r SettingEditParamsItemsZonesSha1Support) implementsZonesSettingEditParamsItemUnion() {}

// Zone setting identifier.
type SettingEditParamsItemsZonesSha1SupportID string

const (
	SettingEditParamsItemsZonesSha1SupportIDSha1Support SettingEditParamsItemsZonesSha1SupportID = "sha1_support"
)

func (r SettingEditParamsItemsZonesSha1SupportID) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesSha1SupportIDSha1Support:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsItemsZonesSha1SupportValue string

const (
	SettingEditParamsItemsZonesSha1SupportValueOff SettingEditParamsItemsZonesSha1SupportValue = "off"
	SettingEditParamsItemsZonesSha1SupportValueOn  SettingEditParamsItemsZonesSha1SupportValue = "on"
)

func (r SettingEditParamsItemsZonesSha1SupportValue) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesSha1SupportValueOff, SettingEditParamsItemsZonesSha1SupportValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesSha1SupportEditable bool

const (
	SettingEditParamsItemsZonesSha1SupportEditableTrue  SettingEditParamsItemsZonesSha1SupportEditable = true
	SettingEditParamsItemsZonesSha1SupportEditableFalse SettingEditParamsItemsZonesSha1SupportEditable = false
)

func (r SettingEditParamsItemsZonesSha1SupportEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesSha1SupportEditableTrue, SettingEditParamsItemsZonesSha1SupportEditableFalse:
		return true
	}
	return false
}

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

func (r SettingEditParamsItemsZonesTLS1_2Only) implementsZonesSettingEditParamsItemUnion() {}

// Zone setting identifier.
type SettingEditParamsItemsZonesTLS1_2OnlyID string

const (
	SettingEditParamsItemsZonesTLS1_2OnlyIDTLS1_2Only SettingEditParamsItemsZonesTLS1_2OnlyID = "tls_1_2_only"
)

func (r SettingEditParamsItemsZonesTLS1_2OnlyID) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesTLS1_2OnlyIDTLS1_2Only:
		return true
	}
	return false
}

// Current value of the zone setting.
type SettingEditParamsItemsZonesTLS1_2OnlyValue string

const (
	SettingEditParamsItemsZonesTLS1_2OnlyValueOff SettingEditParamsItemsZonesTLS1_2OnlyValue = "off"
	SettingEditParamsItemsZonesTLS1_2OnlyValueOn  SettingEditParamsItemsZonesTLS1_2OnlyValue = "on"
)

func (r SettingEditParamsItemsZonesTLS1_2OnlyValue) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesTLS1_2OnlyValueOff, SettingEditParamsItemsZonesTLS1_2OnlyValueOn:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsZonesTLS1_2OnlyEditable bool

const (
	SettingEditParamsItemsZonesTLS1_2OnlyEditableTrue  SettingEditParamsItemsZonesTLS1_2OnlyEditable = true
	SettingEditParamsItemsZonesTLS1_2OnlyEditableFalse SettingEditParamsItemsZonesTLS1_2OnlyEditable = false
)

func (r SettingEditParamsItemsZonesTLS1_2OnlyEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsZonesTLS1_2OnlyEditableTrue, SettingEditParamsItemsZonesTLS1_2OnlyEditableFalse:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SettingEditParamsItemsEditable bool

const (
	SettingEditParamsItemsEditableTrue  SettingEditParamsItemsEditable = true
	SettingEditParamsItemsEditableFalse SettingEditParamsItemsEditable = false
)

func (r SettingEditParamsItemsEditable) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsEditableTrue, SettingEditParamsItemsEditableFalse:
		return true
	}
	return false
}

// ID of the zone setting.
type SettingEditParamsItemsID string

const (
	SettingEditParamsItemsID0rtt                          SettingEditParamsItemsID = "0rtt"
	SettingEditParamsItemsIDAdvancedDDoS                  SettingEditParamsItemsID = "advanced_ddos"
	SettingEditParamsItemsIDAlwaysOnline                  SettingEditParamsItemsID = "always_online"
	SettingEditParamsItemsIDAlwaysUseHTTPS                SettingEditParamsItemsID = "always_use_https"
	SettingEditParamsItemsIDAutomaticHTTPSRewrites        SettingEditParamsItemsID = "automatic_https_rewrites"
	SettingEditParamsItemsIDBrotli                        SettingEditParamsItemsID = "brotli"
	SettingEditParamsItemsIDBrowserCacheTTL               SettingEditParamsItemsID = "browser_cache_ttl"
	SettingEditParamsItemsIDBrowserCheck                  SettingEditParamsItemsID = "browser_check"
	SettingEditParamsItemsIDCacheLevel                    SettingEditParamsItemsID = "cache_level"
	SettingEditParamsItemsIDChallengeTTL                  SettingEditParamsItemsID = "challenge_ttl"
	SettingEditParamsItemsIDCiphers                       SettingEditParamsItemsID = "ciphers"
	SettingEditParamsItemsIDCNAMEFlattening               SettingEditParamsItemsID = "cname_flattening"
	SettingEditParamsItemsIDDevelopmentMode               SettingEditParamsItemsID = "development_mode"
	SettingEditParamsItemsIDEarlyHints                    SettingEditParamsItemsID = "early_hints"
	SettingEditParamsItemsIDEdgeCacheTTL                  SettingEditParamsItemsID = "edge_cache_ttl"
	SettingEditParamsItemsIDEmailObfuscation              SettingEditParamsItemsID = "email_obfuscation"
	SettingEditParamsItemsIDH2Prioritization              SettingEditParamsItemsID = "h2_prioritization"
	SettingEditParamsItemsIDHotlinkProtection             SettingEditParamsItemsID = "hotlink_protection"
	SettingEditParamsItemsIDHTTP2                         SettingEditParamsItemsID = "http2"
	SettingEditParamsItemsIDHTTP3                         SettingEditParamsItemsID = "http3"
	SettingEditParamsItemsIDImageResizing                 SettingEditParamsItemsID = "image_resizing"
	SettingEditParamsItemsIDIPGeolocation                 SettingEditParamsItemsID = "ip_geolocation"
	SettingEditParamsItemsIDIPV6                          SettingEditParamsItemsID = "ipv6"
	SettingEditParamsItemsIDMaxUpload                     SettingEditParamsItemsID = "max_upload"
	SettingEditParamsItemsIDMinTLSVersion                 SettingEditParamsItemsID = "min_tls_version"
	SettingEditParamsItemsIDMinify                        SettingEditParamsItemsID = "minify"
	SettingEditParamsItemsIDMirage                        SettingEditParamsItemsID = "mirage"
	SettingEditParamsItemsIDMobileRedirect                SettingEditParamsItemsID = "mobile_redirect"
	SettingEditParamsItemsIDNEL                           SettingEditParamsItemsID = "nel"
	SettingEditParamsItemsIDOpportunisticEncryption       SettingEditParamsItemsID = "opportunistic_encryption"
	SettingEditParamsItemsIDOpportunisticOnion            SettingEditParamsItemsID = "opportunistic_onion"
	SettingEditParamsItemsIDOrangeToOrange                SettingEditParamsItemsID = "orange_to_orange"
	SettingEditParamsItemsIDOriginErrorPagePassThru       SettingEditParamsItemsID = "origin_error_page_pass_thru"
	SettingEditParamsItemsIDPolish                        SettingEditParamsItemsID = "polish"
	SettingEditParamsItemsIDPrefetchPreload               SettingEditParamsItemsID = "prefetch_preload"
	SettingEditParamsItemsIDProxyReadTimeout              SettingEditParamsItemsID = "proxy_read_timeout"
	SettingEditParamsItemsIDPseudoIPV4                    SettingEditParamsItemsID = "pseudo_ipv4"
	SettingEditParamsItemsIDResponseBuffering             SettingEditParamsItemsID = "response_buffering"
	SettingEditParamsItemsIDRocketLoader                  SettingEditParamsItemsID = "rocket_loader"
	SettingEditParamsItemsIDAutomaticPlatformOptimization SettingEditParamsItemsID = "automatic_platform_optimization"
	SettingEditParamsItemsIDSecurityHeader                SettingEditParamsItemsID = "security_header"
	SettingEditParamsItemsIDSecurityLevel                 SettingEditParamsItemsID = "security_level"
	SettingEditParamsItemsIDServerSideExclude             SettingEditParamsItemsID = "server_side_exclude"
	SettingEditParamsItemsIDSha1Support                   SettingEditParamsItemsID = "sha1_support"
	SettingEditParamsItemsIDSortQueryStringForCache       SettingEditParamsItemsID = "sort_query_string_for_cache"
	SettingEditParamsItemsIDSSL                           SettingEditParamsItemsID = "ssl"
	SettingEditParamsItemsIDSSLRecommender                SettingEditParamsItemsID = "ssl_recommender"
	SettingEditParamsItemsIDTLS1_2Only                    SettingEditParamsItemsID = "tls_1_2_only"
	SettingEditParamsItemsIDTLS1_3                        SettingEditParamsItemsID = "tls_1_3"
	SettingEditParamsItemsIDTLSClientAuth                 SettingEditParamsItemsID = "tls_client_auth"
	SettingEditParamsItemsIDTrueClientIPHeader            SettingEditParamsItemsID = "true_client_ip_header"
	SettingEditParamsItemsIDWAF                           SettingEditParamsItemsID = "waf"
	SettingEditParamsItemsIDWebP                          SettingEditParamsItemsID = "webp"
	SettingEditParamsItemsIDWebsockets                    SettingEditParamsItemsID = "websockets"
)

func (r SettingEditParamsItemsID) IsKnown() bool {
	switch r {
	case SettingEditParamsItemsID0rtt, SettingEditParamsItemsIDAdvancedDDoS, SettingEditParamsItemsIDAlwaysOnline, SettingEditParamsItemsIDAlwaysUseHTTPS, SettingEditParamsItemsIDAutomaticHTTPSRewrites, SettingEditParamsItemsIDBrotli, SettingEditParamsItemsIDBrowserCacheTTL, SettingEditParamsItemsIDBrowserCheck, SettingEditParamsItemsIDCacheLevel, SettingEditParamsItemsIDChallengeTTL, SettingEditParamsItemsIDCiphers, SettingEditParamsItemsIDCNAMEFlattening, SettingEditParamsItemsIDDevelopmentMode, SettingEditParamsItemsIDEarlyHints, SettingEditParamsItemsIDEdgeCacheTTL, SettingEditParamsItemsIDEmailObfuscation, SettingEditParamsItemsIDH2Prioritization, SettingEditParamsItemsIDHotlinkProtection, SettingEditParamsItemsIDHTTP2, SettingEditParamsItemsIDHTTP3, SettingEditParamsItemsIDImageResizing, SettingEditParamsItemsIDIPGeolocation, SettingEditParamsItemsIDIPV6, SettingEditParamsItemsIDMaxUpload, SettingEditParamsItemsIDMinTLSVersion, SettingEditParamsItemsIDMinify, SettingEditParamsItemsIDMirage, SettingEditParamsItemsIDMobileRedirect, SettingEditParamsItemsIDNEL, SettingEditParamsItemsIDOpportunisticEncryption, SettingEditParamsItemsIDOpportunisticOnion, SettingEditParamsItemsIDOrangeToOrange, SettingEditParamsItemsIDOriginErrorPagePassThru, SettingEditParamsItemsIDPolish, SettingEditParamsItemsIDPrefetchPreload, SettingEditParamsItemsIDProxyReadTimeout, SettingEditParamsItemsIDPseudoIPV4, SettingEditParamsItemsIDResponseBuffering, SettingEditParamsItemsIDRocketLoader, SettingEditParamsItemsIDAutomaticPlatformOptimization, SettingEditParamsItemsIDSecurityHeader, SettingEditParamsItemsIDSecurityLevel, SettingEditParamsItemsIDServerSideExclude, SettingEditParamsItemsIDSha1Support, SettingEditParamsItemsIDSortQueryStringForCache, SettingEditParamsItemsIDSSL, SettingEditParamsItemsIDSSLRecommender, SettingEditParamsItemsIDTLS1_2Only, SettingEditParamsItemsIDTLS1_3, SettingEditParamsItemsIDTLSClientAuth, SettingEditParamsItemsIDTrueClientIPHeader, SettingEditParamsItemsIDWAF, SettingEditParamsItemsIDWebP, SettingEditParamsItemsIDWebsockets:
		return true
	}
	return false
}

type SettingEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
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

type SettingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
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
