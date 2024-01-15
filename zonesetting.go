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
	Options                        []option.RequestOption
	ZeroRtts                       *ZoneSettingZeroRttService
	AdvancedDdos                   *ZoneSettingAdvancedDdoService
	AlwaysOnlines                  *ZoneSettingAlwaysOnlineService
	AlwaysUseHTTPs                 *ZoneSettingAlwaysUseHTTPService
	AutomaticHTTPsRewrites         *ZoneSettingAutomaticHTTPsRewriteService
	AutomaticPlatformOptimizations *ZoneSettingAutomaticPlatformOptimizationService
	Brotli                         *ZoneSettingBrotliService
	BrowserCacheTtls               *ZoneSettingBrowserCacheTtlService
	BrowserChecks                  *ZoneSettingBrowserCheckService
	CacheLevels                    *ZoneSettingCacheLevelService
	ChallengeTtls                  *ZoneSettingChallengeTtlService
	Ciphers                        *ZoneSettingCipherService
	DevelopmentModes               *ZoneSettingDevelopmentModeService
	EarlyHints                     *ZoneSettingEarlyHintService
	EmailObfuscations              *ZoneSettingEmailObfuscationService
	H2Prioritizations              *ZoneSettingH2PrioritizationService
	HotlinkProtections             *ZoneSettingHotlinkProtectionService
	Http2s                         *ZoneSettingHttp2Service
	Http3s                         *ZoneSettingHttp3Service
	ImageResizings                 *ZoneSettingImageResizingService
	IPGeolocations                 *ZoneSettingIPGeolocationService
	Ipv6s                          *ZoneSettingIpv6Service
	MinTlsVersions                 *ZoneSettingMinTlsVersionService
	Minifies                       *ZoneSettingMinifyService
	Mirages                        *ZoneSettingMirageService
	MobileRedirects                *ZoneSettingMobileRedirectService
	Nels                           *ZoneSettingNelService
	OpportunisticEncryptions       *ZoneSettingOpportunisticEncryptionService
	OpportunisticOnions            *ZoneSettingOpportunisticOnionService
	OrangeToOranges                *ZoneSettingOrangeToOrangeService
	OriginErrorPagePassThrus       *ZoneSettingOriginErrorPagePassThrusService
	OriginMaxHTTPVersions          *ZoneSettingOriginMaxHTTPVersionService
	Polishes                       *ZoneSettingPolishService
	PrefetchPreloads               *ZoneSettingPrefetchPreloadService
	ProxyReadTimeouts              *ZoneSettingProxyReadTimeoutService
	PseudoIpv4s                    *ZoneSettingPseudoIpv4Service
	ResponseBufferings             *ZoneSettingResponseBufferingService
	RocketLoaders                  *ZoneSettingRocketLoaderService
	SecurityHeaders                *ZoneSettingSecurityHeaderService
	SecurityLevels                 *ZoneSettingSecurityLevelService
	ServerSideExcludes             *ZoneSettingServerSideExcludeService
	SortQueryStringForCaches       *ZoneSettingSortQueryStringForCachService
	Ssls                           *ZoneSettingSslService
	SslRecommenders                *ZoneSettingSslRecommenderService
	Tls1_3s                        *ZoneSettingTls1_3Service
	TlsClientAuths                 *ZoneSettingTlsClientAuthService
	TrueClientIPHeaders            *ZoneSettingTrueClientIPHeaderService
	Wafs                           *ZoneSettingWafService
	Webps                          *ZoneSettingWebpService
	Websockets                     *ZoneSettingWebsocketService
	Fonts                          *ZoneSettingFontService
	Zaraz                          *ZoneSettingZarazService
}

// NewZoneSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingService(opts ...option.RequestOption) (r *ZoneSettingService) {
	r = &ZoneSettingService{}
	r.Options = opts
	r.ZeroRtts = NewZoneSettingZeroRttService(opts...)
	r.AdvancedDdos = NewZoneSettingAdvancedDdoService(opts...)
	r.AlwaysOnlines = NewZoneSettingAlwaysOnlineService(opts...)
	r.AlwaysUseHTTPs = NewZoneSettingAlwaysUseHTTPService(opts...)
	r.AutomaticHTTPsRewrites = NewZoneSettingAutomaticHTTPsRewriteService(opts...)
	r.AutomaticPlatformOptimizations = NewZoneSettingAutomaticPlatformOptimizationService(opts...)
	r.Brotli = NewZoneSettingBrotliService(opts...)
	r.BrowserCacheTtls = NewZoneSettingBrowserCacheTtlService(opts...)
	r.BrowserChecks = NewZoneSettingBrowserCheckService(opts...)
	r.CacheLevels = NewZoneSettingCacheLevelService(opts...)
	r.ChallengeTtls = NewZoneSettingChallengeTtlService(opts...)
	r.Ciphers = NewZoneSettingCipherService(opts...)
	r.DevelopmentModes = NewZoneSettingDevelopmentModeService(opts...)
	r.EarlyHints = NewZoneSettingEarlyHintService(opts...)
	r.EmailObfuscations = NewZoneSettingEmailObfuscationService(opts...)
	r.H2Prioritizations = NewZoneSettingH2PrioritizationService(opts...)
	r.HotlinkProtections = NewZoneSettingHotlinkProtectionService(opts...)
	r.Http2s = NewZoneSettingHttp2Service(opts...)
	r.Http3s = NewZoneSettingHttp3Service(opts...)
	r.ImageResizings = NewZoneSettingImageResizingService(opts...)
	r.IPGeolocations = NewZoneSettingIPGeolocationService(opts...)
	r.Ipv6s = NewZoneSettingIpv6Service(opts...)
	r.MinTlsVersions = NewZoneSettingMinTlsVersionService(opts...)
	r.Minifies = NewZoneSettingMinifyService(opts...)
	r.Mirages = NewZoneSettingMirageService(opts...)
	r.MobileRedirects = NewZoneSettingMobileRedirectService(opts...)
	r.Nels = NewZoneSettingNelService(opts...)
	r.OpportunisticEncryptions = NewZoneSettingOpportunisticEncryptionService(opts...)
	r.OpportunisticOnions = NewZoneSettingOpportunisticOnionService(opts...)
	r.OrangeToOranges = NewZoneSettingOrangeToOrangeService(opts...)
	r.OriginErrorPagePassThrus = NewZoneSettingOriginErrorPagePassThrusService(opts...)
	r.OriginMaxHTTPVersions = NewZoneSettingOriginMaxHTTPVersionService(opts...)
	r.Polishes = NewZoneSettingPolishService(opts...)
	r.PrefetchPreloads = NewZoneSettingPrefetchPreloadService(opts...)
	r.ProxyReadTimeouts = NewZoneSettingProxyReadTimeoutService(opts...)
	r.PseudoIpv4s = NewZoneSettingPseudoIpv4Service(opts...)
	r.ResponseBufferings = NewZoneSettingResponseBufferingService(opts...)
	r.RocketLoaders = NewZoneSettingRocketLoaderService(opts...)
	r.SecurityHeaders = NewZoneSettingSecurityHeaderService(opts...)
	r.SecurityLevels = NewZoneSettingSecurityLevelService(opts...)
	r.ServerSideExcludes = NewZoneSettingServerSideExcludeService(opts...)
	r.SortQueryStringForCaches = NewZoneSettingSortQueryStringForCachService(opts...)
	r.Ssls = NewZoneSettingSslService(opts...)
	r.SslRecommenders = NewZoneSettingSslRecommenderService(opts...)
	r.Tls1_3s = NewZoneSettingTls1_3Service(opts...)
	r.TlsClientAuths = NewZoneSettingTlsClientAuthService(opts...)
	r.TrueClientIPHeaders = NewZoneSettingTrueClientIPHeaderService(opts...)
	r.Wafs = NewZoneSettingWafService(opts...)
	r.Webps = NewZoneSettingWebpService(opts...)
	r.Websockets = NewZoneSettingWebsocketService(opts...)
	r.Fonts = NewZoneSettingFontService(opts...)
	r.Zaraz = NewZoneSettingZarazService(opts...)
	return
}

// Available settings for your user in relation to a zone.
func (r *ZoneSettingService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit settings for a zone.
func (r *ZoneSettingService) Edit(ctx context.Context, zoneIdentifier string, body ZoneSettingEditParams, opts ...option.RequestOption) (res *ZoneSettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneSettingListResponse struct {
	Errors   []ZoneSettingListResponseError   `json:"errors"`
	Messages []ZoneSettingListResponseMessage `json:"messages"`
	Result   []ZoneSettingListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success bool                        `json:"success"`
	JSON    zoneSettingListResponseJSON `json:"-"`
}

// zoneSettingListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingListResponse]
type zoneSettingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingListResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    zoneSettingListResponseErrorJSON `json:"-"`
}

// zoneSettingListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseError]
type zoneSettingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingListResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneSettingListResponseMessageJSON `json:"-"`
}

// zoneSettingListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingListResponseMessage]
type zoneSettingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [ZoneSettingListResponseResultZones0rtt],
// [ZoneSettingListResponseResultZonesAdvancedDdos],
// [ZoneSettingListResponseResultZonesAlwaysOnline],
// [ZoneSettingListResponseResultZonesAlwaysUseHTTPs],
// [ZoneSettingListResponseResultZonesAutomaticHTTPsRewrites],
// [ZoneSettingListResponseResultZonesBrotli],
// [ZoneSettingListResponseResultZonesBrowserCacheTtl],
// [ZoneSettingListResponseResultZonesBrowserCheck],
// [ZoneSettingListResponseResultZonesCacheLevel],
// [ZoneSettingListResponseResultZonesChallengeTtl],
// [ZoneSettingListResponseResultZonesCiphers],
// [ZoneSettingListResponseResultZonesCnameFlattening],
// [ZoneSettingListResponseResultZonesDevelopmentMode],
// [ZoneSettingListResponseResultZonesEarlyHints],
// [ZoneSettingListResponseResultZonesEdgeCacheTtl],
// [ZoneSettingListResponseResultZonesEmailObfuscation],
// [ZoneSettingListResponseResultZonesH2Prioritization],
// [ZoneSettingListResponseResultZonesHotlinkProtection],
// [ZoneSettingListResponseResultZonesHttp2],
// [ZoneSettingListResponseResultZonesHttp3],
// [ZoneSettingListResponseResultZonesImageResizing],
// [ZoneSettingListResponseResultZonesIPGeolocation],
// [ZoneSettingListResponseResultZonesIpv6],
// [ZoneSettingListResponseResultZonesMaxUpload],
// [ZoneSettingListResponseResultZonesMinTlsVersion],
// [ZoneSettingListResponseResultZonesMinify],
// [ZoneSettingListResponseResultZonesMirage],
// [ZoneSettingListResponseResultZonesMobileRedirect],
// [ZoneSettingListResponseResultZonesNel],
// [ZoneSettingListResponseResultZonesOpportunisticEncryption],
// [ZoneSettingListResponseResultZonesOpportunisticOnion],
// [ZoneSettingListResponseResultZonesOrangeToOrange],
// [ZoneSettingListResponseResultZonesOriginErrorPagePassThru],
// [ZoneSettingListResponseResultZonesOriginMaxHTTPVersion],
// [ZoneSettingListResponseResultZonesPolish],
// [ZoneSettingListResponseResultZonesPrefetchPreload],
// [ZoneSettingListResponseResultZonesProxyReadTimeout],
// [ZoneSettingListResponseResultZonesPseudoIpv4],
// [ZoneSettingListResponseResultZonesResponseBuffering],
// [ZoneSettingListResponseResultZonesRocketLoader],
// [ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimization],
// [ZoneSettingListResponseResultZonesSecurityHeader],
// [ZoneSettingListResponseResultZonesSecurityLevel],
// [ZoneSettingListResponseResultZonesServerSideExclude],
// [ZoneSettingListResponseResultZonesSha1Support],
// [ZoneSettingListResponseResultZonesSortQueryStringForCache],
// [ZoneSettingListResponseResultZonesSsl],
// [ZoneSettingListResponseResultZonesSslRecommender],
// [ZoneSettingListResponseResultZonesTls1_2Only],
// [ZoneSettingListResponseResultZonesTls1_3],
// [ZoneSettingListResponseResultZonesTlsClientAuth],
// [ZoneSettingListResponseResultZonesTrueClientIPHeader],
// [ZoneSettingListResponseResultZonesWaf],
// [ZoneSettingListResponseResultZonesWebp] or
// [ZoneSettingListResponseResultZonesWebsockets].
type ZoneSettingListResponseResult interface {
	implementsZoneSettingListResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneSettingListResponseResult)(nil)).Elem(), "")
}

// 0-RTT session resumption enabled for this zone.
type ZoneSettingListResponseResultZones0rtt struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZones0rttID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZones0rttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the 0-RTT setting.
	Value ZoneSettingListResponseResultZones0rttValue `json:"value"`
	JSON  zoneSettingListResponseResultZones0rttJSON  `json:"-"`
}

// zoneSettingListResponseResultZones0rttJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZones0rtt]
type zoneSettingListResponseResultZones0rttJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZones0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZones0rtt) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZones0rttID string

const (
	ZoneSettingListResponseResultZones0rttID0rtt ZoneSettingListResponseResultZones0rttID = "0rtt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZones0rttEditable bool

const (
	ZoneSettingListResponseResultZones0rttEditableTrue  ZoneSettingListResponseResultZones0rttEditable = true
	ZoneSettingListResponseResultZones0rttEditableFalse ZoneSettingListResponseResultZones0rttEditable = false
)

// Value of the 0-RTT setting.
type ZoneSettingListResponseResultZones0rttValue string

const (
	ZoneSettingListResponseResultZones0rttValueOn  ZoneSettingListResponseResultZones0rttValue = "on"
	ZoneSettingListResponseResultZones0rttValueOff ZoneSettingListResponseResultZones0rttValue = "off"
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingListResponseResultZonesAdvancedDdos struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesAdvancedDdosID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesAdvancedDdosEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value ZoneSettingListResponseResultZonesAdvancedDdosValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesAdvancedDdosJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesAdvancedDdosJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesAdvancedDdos]
type zoneSettingListResponseResultZonesAdvancedDdosJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesAdvancedDdos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesAdvancedDdos) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesAdvancedDdosID string

const (
	ZoneSettingListResponseResultZonesAdvancedDdosIDAdvancedDdos ZoneSettingListResponseResultZonesAdvancedDdosID = "advanced_ddos"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesAdvancedDdosEditable bool

const (
	ZoneSettingListResponseResultZonesAdvancedDdosEditableTrue  ZoneSettingListResponseResultZonesAdvancedDdosEditable = true
	ZoneSettingListResponseResultZonesAdvancedDdosEditableFalse ZoneSettingListResponseResultZonesAdvancedDdosEditable = false
)

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZoneSettingListResponseResultZonesAdvancedDdosValue string

const (
	ZoneSettingListResponseResultZonesAdvancedDdosValueOn  ZoneSettingListResponseResultZonesAdvancedDdosValue = "on"
	ZoneSettingListResponseResultZonesAdvancedDdosValueOff ZoneSettingListResponseResultZonesAdvancedDdosValue = "off"
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingListResponseResultZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesAlwaysOnlineID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesAlwaysOnlineValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesAlwaysOnlineJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesAlwaysOnlineJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesAlwaysOnline]
type zoneSettingListResponseResultZonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesAlwaysOnline) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesAlwaysOnlineID string

const (
	ZoneSettingListResponseResultZonesAlwaysOnlineIDAlwaysOnline ZoneSettingListResponseResultZonesAlwaysOnlineID = "always_online"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesAlwaysOnlineEditable bool

const (
	ZoneSettingListResponseResultZonesAlwaysOnlineEditableTrue  ZoneSettingListResponseResultZonesAlwaysOnlineEditable = true
	ZoneSettingListResponseResultZonesAlwaysOnlineEditableFalse ZoneSettingListResponseResultZonesAlwaysOnlineEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesAlwaysOnlineValue string

const (
	ZoneSettingListResponseResultZonesAlwaysOnlineValueOn  ZoneSettingListResponseResultZonesAlwaysOnlineValue = "on"
	ZoneSettingListResponseResultZonesAlwaysOnlineValueOff ZoneSettingListResponseResultZonesAlwaysOnlineValue = "off"
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingListResponseResultZonesAlwaysUseHTTPs struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesAlwaysUseHTTPsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesAlwaysUseHTTPsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesAlwaysUseHTTPsValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesAlwaysUseHTTPsJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesAlwaysUseHTTPsJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesAlwaysUseHTTPs]
type zoneSettingListResponseResultZonesAlwaysUseHTTPsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesAlwaysUseHTTPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesAlwaysUseHTTPs) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesAlwaysUseHTTPsID string

const (
	ZoneSettingListResponseResultZonesAlwaysUseHTTPsIDAlwaysUseHTTPs ZoneSettingListResponseResultZonesAlwaysUseHTTPsID = "always_use_https"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesAlwaysUseHTTPsEditable bool

const (
	ZoneSettingListResponseResultZonesAlwaysUseHTTPsEditableTrue  ZoneSettingListResponseResultZonesAlwaysUseHTTPsEditable = true
	ZoneSettingListResponseResultZonesAlwaysUseHTTPsEditableFalse ZoneSettingListResponseResultZonesAlwaysUseHTTPsEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesAlwaysUseHTTPsValue string

const (
	ZoneSettingListResponseResultZonesAlwaysUseHTTPsValueOn  ZoneSettingListResponseResultZonesAlwaysUseHTTPsValue = "on"
	ZoneSettingListResponseResultZonesAlwaysUseHTTPsValueOff ZoneSettingListResponseResultZonesAlwaysUseHTTPsValue = "off"
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingListResponseResultZonesAutomaticHTTPsRewrites struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesAutomaticHTTPsRewritesJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesAutomaticHTTPsRewritesJSON contains the JSON
// metadata for the struct
// [ZoneSettingListResponseResultZonesAutomaticHTTPsRewrites]
type zoneSettingListResponseResultZonesAutomaticHTTPsRewritesJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesAutomaticHTTPsRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesAutomaticHTTPsRewrites) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesID string

const (
	ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesIDAutomaticHTTPsRewrites ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesID = "automatic_https_rewrites"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesEditable bool

const (
	ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesEditableTrue  ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesEditable = true
	ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesEditableFalse ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesValue string

const (
	ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesValueOn  ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesValue = "on"
	ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesValueOff ZoneSettingListResponseResultZonesAutomaticHTTPsRewritesValue = "off"
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingListResponseResultZonesBrotli struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesBrotliID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesBrotliValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesBrotliJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesBrotliJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesBrotli]
type zoneSettingListResponseResultZonesBrotliJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesBrotli) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesBrotliID string

const (
	ZoneSettingListResponseResultZonesBrotliIDBrotli ZoneSettingListResponseResultZonesBrotliID = "brotli"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesBrotliEditable bool

const (
	ZoneSettingListResponseResultZonesBrotliEditableTrue  ZoneSettingListResponseResultZonesBrotliEditable = true
	ZoneSettingListResponseResultZonesBrotliEditableFalse ZoneSettingListResponseResultZonesBrotliEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesBrotliValue string

const (
	ZoneSettingListResponseResultZonesBrotliValueOff ZoneSettingListResponseResultZonesBrotliValue = "off"
	ZoneSettingListResponseResultZonesBrotliValueOn  ZoneSettingListResponseResultZonesBrotliValue = "on"
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingListResponseResultZonesBrowserCacheTtl struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesBrowserCacheTtlID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesBrowserCacheTtlEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingListResponseResultZonesBrowserCacheTtlValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesBrowserCacheTtlJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesBrowserCacheTtlJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesBrowserCacheTtl]
type zoneSettingListResponseResultZonesBrowserCacheTtlJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesBrowserCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesBrowserCacheTtl) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesBrowserCacheTtlID string

const (
	ZoneSettingListResponseResultZonesBrowserCacheTtlIDBrowserCacheTtl ZoneSettingListResponseResultZonesBrowserCacheTtlID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesBrowserCacheTtlEditable bool

const (
	ZoneSettingListResponseResultZonesBrowserCacheTtlEditableTrue  ZoneSettingListResponseResultZonesBrowserCacheTtlEditable = true
	ZoneSettingListResponseResultZonesBrowserCacheTtlEditableFalse ZoneSettingListResponseResultZonesBrowserCacheTtlEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingListResponseResultZonesBrowserCacheTtlValue float64

const (
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue0        ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 0
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue30       ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 30
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue60       ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 60
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue120      ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 120
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue300      ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 300
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue1200     ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 1200
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue1800     ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 1800
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue3600     ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 3600
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue7200     ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 7200
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue10800    ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 10800
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue14400    ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 14400
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue18000    ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 18000
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue28800    ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 28800
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue43200    ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 43200
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue57600    ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 57600
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue72000    ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 72000
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue86400    ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 86400
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue172800   ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 172800
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue259200   ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 259200
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue345600   ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 345600
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue432000   ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 432000
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue691200   ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 691200
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue1382400  ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 1382400
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue2073600  ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 2073600
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue2678400  ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 2678400
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue5356800  ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 5356800
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue16070400 ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 16070400
	ZoneSettingListResponseResultZonesBrowserCacheTtlValue31536000 ZoneSettingListResponseResultZonesBrowserCacheTtlValue = 31536000
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingListResponseResultZonesBrowserCheck struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesBrowserCheckID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesBrowserCheckValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesBrowserCheckJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesBrowserCheckJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesBrowserCheck]
type zoneSettingListResponseResultZonesBrowserCheckJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesBrowserCheck) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesBrowserCheckID string

const (
	ZoneSettingListResponseResultZonesBrowserCheckIDBrowserCheck ZoneSettingListResponseResultZonesBrowserCheckID = "browser_check"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesBrowserCheckEditable bool

const (
	ZoneSettingListResponseResultZonesBrowserCheckEditableTrue  ZoneSettingListResponseResultZonesBrowserCheckEditable = true
	ZoneSettingListResponseResultZonesBrowserCheckEditableFalse ZoneSettingListResponseResultZonesBrowserCheckEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesBrowserCheckValue string

const (
	ZoneSettingListResponseResultZonesBrowserCheckValueOn  ZoneSettingListResponseResultZonesBrowserCheckValue = "on"
	ZoneSettingListResponseResultZonesBrowserCheckValueOff ZoneSettingListResponseResultZonesBrowserCheckValue = "off"
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingListResponseResultZonesCacheLevel struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesCacheLevelID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesCacheLevelValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesCacheLevelJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesCacheLevelJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesCacheLevel]
type zoneSettingListResponseResultZonesCacheLevelJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesCacheLevel) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesCacheLevelID string

const (
	ZoneSettingListResponseResultZonesCacheLevelIDCacheLevel ZoneSettingListResponseResultZonesCacheLevelID = "cache_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesCacheLevelEditable bool

const (
	ZoneSettingListResponseResultZonesCacheLevelEditableTrue  ZoneSettingListResponseResultZonesCacheLevelEditable = true
	ZoneSettingListResponseResultZonesCacheLevelEditableFalse ZoneSettingListResponseResultZonesCacheLevelEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesCacheLevelValue string

const (
	ZoneSettingListResponseResultZonesCacheLevelValueAggressive ZoneSettingListResponseResultZonesCacheLevelValue = "aggressive"
	ZoneSettingListResponseResultZonesCacheLevelValueBasic      ZoneSettingListResponseResultZonesCacheLevelValue = "basic"
	ZoneSettingListResponseResultZonesCacheLevelValueSimplified ZoneSettingListResponseResultZonesCacheLevelValue = "simplified"
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingListResponseResultZonesChallengeTtl struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesChallengeTtlID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesChallengeTtlEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesChallengeTtlValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesChallengeTtlJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesChallengeTtlJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesChallengeTtl]
type zoneSettingListResponseResultZonesChallengeTtlJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesChallengeTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesChallengeTtl) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesChallengeTtlID string

const (
	ZoneSettingListResponseResultZonesChallengeTtlIDChallengeTtl ZoneSettingListResponseResultZonesChallengeTtlID = "challenge_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesChallengeTtlEditable bool

const (
	ZoneSettingListResponseResultZonesChallengeTtlEditableTrue  ZoneSettingListResponseResultZonesChallengeTtlEditable = true
	ZoneSettingListResponseResultZonesChallengeTtlEditableFalse ZoneSettingListResponseResultZonesChallengeTtlEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesChallengeTtlValue float64

const (
	ZoneSettingListResponseResultZonesChallengeTtlValue300      ZoneSettingListResponseResultZonesChallengeTtlValue = 300
	ZoneSettingListResponseResultZonesChallengeTtlValue900      ZoneSettingListResponseResultZonesChallengeTtlValue = 900
	ZoneSettingListResponseResultZonesChallengeTtlValue1800     ZoneSettingListResponseResultZonesChallengeTtlValue = 1800
	ZoneSettingListResponseResultZonesChallengeTtlValue2700     ZoneSettingListResponseResultZonesChallengeTtlValue = 2700
	ZoneSettingListResponseResultZonesChallengeTtlValue3600     ZoneSettingListResponseResultZonesChallengeTtlValue = 3600
	ZoneSettingListResponseResultZonesChallengeTtlValue7200     ZoneSettingListResponseResultZonesChallengeTtlValue = 7200
	ZoneSettingListResponseResultZonesChallengeTtlValue10800    ZoneSettingListResponseResultZonesChallengeTtlValue = 10800
	ZoneSettingListResponseResultZonesChallengeTtlValue14400    ZoneSettingListResponseResultZonesChallengeTtlValue = 14400
	ZoneSettingListResponseResultZonesChallengeTtlValue28800    ZoneSettingListResponseResultZonesChallengeTtlValue = 28800
	ZoneSettingListResponseResultZonesChallengeTtlValue57600    ZoneSettingListResponseResultZonesChallengeTtlValue = 57600
	ZoneSettingListResponseResultZonesChallengeTtlValue86400    ZoneSettingListResponseResultZonesChallengeTtlValue = 86400
	ZoneSettingListResponseResultZonesChallengeTtlValue604800   ZoneSettingListResponseResultZonesChallengeTtlValue = 604800
	ZoneSettingListResponseResultZonesChallengeTtlValue2592000  ZoneSettingListResponseResultZonesChallengeTtlValue = 2592000
	ZoneSettingListResponseResultZonesChallengeTtlValue31536000 ZoneSettingListResponseResultZonesChallengeTtlValue = 31536000
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingListResponseResultZonesCiphers struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesCiphersID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesCiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value []string                                      `json:"value"`
	JSON  zoneSettingListResponseResultZonesCiphersJSON `json:"-"`
}

// zoneSettingListResponseResultZonesCiphersJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesCiphers]
type zoneSettingListResponseResultZonesCiphersJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesCiphers) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesCiphersID string

const (
	ZoneSettingListResponseResultZonesCiphersIDCiphers ZoneSettingListResponseResultZonesCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesCiphersEditable bool

const (
	ZoneSettingListResponseResultZonesCiphersEditableTrue  ZoneSettingListResponseResultZonesCiphersEditable = true
	ZoneSettingListResponseResultZonesCiphersEditableFalse ZoneSettingListResponseResultZonesCiphersEditable = false
)

// Whether or not cname flattening is on.
type ZoneSettingListResponseResultZonesCnameFlattening struct {
	// How to flatten the cname destination.
	ID ZoneSettingListResponseResultZonesCnameFlatteningID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesCnameFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the cname flattening setting.
	Value ZoneSettingListResponseResultZonesCnameFlatteningValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesCnameFlatteningJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesCnameFlatteningJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesCnameFlattening]
type zoneSettingListResponseResultZonesCnameFlatteningJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesCnameFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesCnameFlattening) implementsZoneSettingListResponseResult() {
}

// How to flatten the cname destination.
type ZoneSettingListResponseResultZonesCnameFlatteningID string

const (
	ZoneSettingListResponseResultZonesCnameFlatteningIDCnameFlattening ZoneSettingListResponseResultZonesCnameFlatteningID = "cname_flattening"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesCnameFlatteningEditable bool

const (
	ZoneSettingListResponseResultZonesCnameFlatteningEditableTrue  ZoneSettingListResponseResultZonesCnameFlatteningEditable = true
	ZoneSettingListResponseResultZonesCnameFlatteningEditableFalse ZoneSettingListResponseResultZonesCnameFlatteningEditable = false
)

// Value of the cname flattening setting.
type ZoneSettingListResponseResultZonesCnameFlatteningValue string

const (
	ZoneSettingListResponseResultZonesCnameFlatteningValueFlattenAtRoot ZoneSettingListResponseResultZonesCnameFlatteningValue = "flatten_at_root"
	ZoneSettingListResponseResultZonesCnameFlatteningValueFlattenAll    ZoneSettingListResponseResultZonesCnameFlatteningValue = "flatten_all"
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingListResponseResultZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesDevelopmentModeID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesDevelopmentModeValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesDevelopmentModeJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesDevelopmentModeJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesDevelopmentMode]
type zoneSettingListResponseResultZonesDevelopmentModeJSON struct {
	ID            apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesDevelopmentMode) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesDevelopmentModeID string

const (
	ZoneSettingListResponseResultZonesDevelopmentModeIDDevelopmentMode ZoneSettingListResponseResultZonesDevelopmentModeID = "development_mode"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesDevelopmentModeEditable bool

const (
	ZoneSettingListResponseResultZonesDevelopmentModeEditableTrue  ZoneSettingListResponseResultZonesDevelopmentModeEditable = true
	ZoneSettingListResponseResultZonesDevelopmentModeEditableFalse ZoneSettingListResponseResultZonesDevelopmentModeEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesDevelopmentModeValue string

const (
	ZoneSettingListResponseResultZonesDevelopmentModeValueOn  ZoneSettingListResponseResultZonesDevelopmentModeValue = "on"
	ZoneSettingListResponseResultZonesDevelopmentModeValueOff ZoneSettingListResponseResultZonesDevelopmentModeValue = "off"
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingListResponseResultZonesEarlyHints struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesEarlyHintsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesEarlyHintsValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesEarlyHintsJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesEarlyHintsJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesEarlyHints]
type zoneSettingListResponseResultZonesEarlyHintsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesEarlyHints) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesEarlyHintsID string

const (
	ZoneSettingListResponseResultZonesEarlyHintsIDEarlyHints ZoneSettingListResponseResultZonesEarlyHintsID = "early_hints"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesEarlyHintsEditable bool

const (
	ZoneSettingListResponseResultZonesEarlyHintsEditableTrue  ZoneSettingListResponseResultZonesEarlyHintsEditable = true
	ZoneSettingListResponseResultZonesEarlyHintsEditableFalse ZoneSettingListResponseResultZonesEarlyHintsEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesEarlyHintsValue string

const (
	ZoneSettingListResponseResultZonesEarlyHintsValueOn  ZoneSettingListResponseResultZonesEarlyHintsValue = "on"
	ZoneSettingListResponseResultZonesEarlyHintsValueOff ZoneSettingListResponseResultZonesEarlyHintsValue = "off"
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZoneSettingListResponseResultZonesEdgeCacheTtl struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesEdgeCacheTtlID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesEdgeCacheTtlEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value ZoneSettingListResponseResultZonesEdgeCacheTtlValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesEdgeCacheTtlJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesEdgeCacheTtlJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesEdgeCacheTtl]
type zoneSettingListResponseResultZonesEdgeCacheTtlJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesEdgeCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesEdgeCacheTtl) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesEdgeCacheTtlID string

const (
	ZoneSettingListResponseResultZonesEdgeCacheTtlIDEdgeCacheTtl ZoneSettingListResponseResultZonesEdgeCacheTtlID = "edge_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesEdgeCacheTtlEditable bool

const (
	ZoneSettingListResponseResultZonesEdgeCacheTtlEditableTrue  ZoneSettingListResponseResultZonesEdgeCacheTtlEditable = true
	ZoneSettingListResponseResultZonesEdgeCacheTtlEditableFalse ZoneSettingListResponseResultZonesEdgeCacheTtlEditable = false
)

// Value of the zone setting. Notes: The minimum TTL available depends on the plan
// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
type ZoneSettingListResponseResultZonesEdgeCacheTtlValue float64

const (
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue30     ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 30
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue60     ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 60
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue300    ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 300
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue1200   ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 1200
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue1800   ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 1800
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue3600   ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 3600
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue7200   ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 7200
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue10800  ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 10800
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue14400  ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 14400
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue18000  ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 18000
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue28800  ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 28800
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue43200  ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 43200
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue57600  ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 57600
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue72000  ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 72000
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue86400  ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 86400
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue172800 ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 172800
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue259200 ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 259200
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue345600 ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 345600
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue432000 ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 432000
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue518400 ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 518400
	ZoneSettingListResponseResultZonesEdgeCacheTtlValue604800 ZoneSettingListResponseResultZonesEdgeCacheTtlValue = 604800
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingListResponseResultZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesEmailObfuscationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesEmailObfuscationValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesEmailObfuscationJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesEmailObfuscationJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesEmailObfuscation]
type zoneSettingListResponseResultZonesEmailObfuscationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesEmailObfuscation) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesEmailObfuscationID string

const (
	ZoneSettingListResponseResultZonesEmailObfuscationIDEmailObfuscation ZoneSettingListResponseResultZonesEmailObfuscationID = "email_obfuscation"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesEmailObfuscationEditable bool

const (
	ZoneSettingListResponseResultZonesEmailObfuscationEditableTrue  ZoneSettingListResponseResultZonesEmailObfuscationEditable = true
	ZoneSettingListResponseResultZonesEmailObfuscationEditableFalse ZoneSettingListResponseResultZonesEmailObfuscationEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesEmailObfuscationValue string

const (
	ZoneSettingListResponseResultZonesEmailObfuscationValueOn  ZoneSettingListResponseResultZonesEmailObfuscationValue = "on"
	ZoneSettingListResponseResultZonesEmailObfuscationValueOff ZoneSettingListResponseResultZonesEmailObfuscationValue = "off"
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingListResponseResultZonesH2Prioritization struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesH2PrioritizationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesH2PrioritizationValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesH2PrioritizationJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesH2PrioritizationJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesH2Prioritization]
type zoneSettingListResponseResultZonesH2PrioritizationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesH2Prioritization) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesH2PrioritizationID string

const (
	ZoneSettingListResponseResultZonesH2PrioritizationIDH2Prioritization ZoneSettingListResponseResultZonesH2PrioritizationID = "h2_prioritization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesH2PrioritizationEditable bool

const (
	ZoneSettingListResponseResultZonesH2PrioritizationEditableTrue  ZoneSettingListResponseResultZonesH2PrioritizationEditable = true
	ZoneSettingListResponseResultZonesH2PrioritizationEditableFalse ZoneSettingListResponseResultZonesH2PrioritizationEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesH2PrioritizationValue string

const (
	ZoneSettingListResponseResultZonesH2PrioritizationValueOn     ZoneSettingListResponseResultZonesH2PrioritizationValue = "on"
	ZoneSettingListResponseResultZonesH2PrioritizationValueOff    ZoneSettingListResponseResultZonesH2PrioritizationValue = "off"
	ZoneSettingListResponseResultZonesH2PrioritizationValueCustom ZoneSettingListResponseResultZonesH2PrioritizationValue = "custom"
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZoneSettingListResponseResultZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesHotlinkProtectionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesHotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesHotlinkProtectionValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesHotlinkProtectionJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesHotlinkProtectionJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesHotlinkProtection]
type zoneSettingListResponseResultZonesHotlinkProtectionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesHotlinkProtection) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesHotlinkProtectionID string

const (
	ZoneSettingListResponseResultZonesHotlinkProtectionIDHotlinkProtection ZoneSettingListResponseResultZonesHotlinkProtectionID = "hotlink_protection"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesHotlinkProtectionEditable bool

const (
	ZoneSettingListResponseResultZonesHotlinkProtectionEditableTrue  ZoneSettingListResponseResultZonesHotlinkProtectionEditable = true
	ZoneSettingListResponseResultZonesHotlinkProtectionEditableFalse ZoneSettingListResponseResultZonesHotlinkProtectionEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesHotlinkProtectionValue string

const (
	ZoneSettingListResponseResultZonesHotlinkProtectionValueOn  ZoneSettingListResponseResultZonesHotlinkProtectionValue = "on"
	ZoneSettingListResponseResultZonesHotlinkProtectionValueOff ZoneSettingListResponseResultZonesHotlinkProtectionValue = "off"
)

// HTTP2 enabled for this zone.
type ZoneSettingListResponseResultZonesHttp2 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesHttp2ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesHttp2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value ZoneSettingListResponseResultZonesHttp2Value `json:"value"`
	JSON  zoneSettingListResponseResultZonesHttp2JSON  `json:"-"`
}

// zoneSettingListResponseResultZonesHttp2JSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesHttp2]
type zoneSettingListResponseResultZonesHttp2JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesHttp2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesHttp2) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesHttp2ID string

const (
	ZoneSettingListResponseResultZonesHttp2IDHttp2 ZoneSettingListResponseResultZonesHttp2ID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesHttp2Editable bool

const (
	ZoneSettingListResponseResultZonesHttp2EditableTrue  ZoneSettingListResponseResultZonesHttp2Editable = true
	ZoneSettingListResponseResultZonesHttp2EditableFalse ZoneSettingListResponseResultZonesHttp2Editable = false
)

// Value of the HTTP2 setting.
type ZoneSettingListResponseResultZonesHttp2Value string

const (
	ZoneSettingListResponseResultZonesHttp2ValueOn  ZoneSettingListResponseResultZonesHttp2Value = "on"
	ZoneSettingListResponseResultZonesHttp2ValueOff ZoneSettingListResponseResultZonesHttp2Value = "off"
)

// HTTP3 enabled for this zone.
type ZoneSettingListResponseResultZonesHttp3 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesHttp3ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesHttp3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP3 setting.
	Value ZoneSettingListResponseResultZonesHttp3Value `json:"value"`
	JSON  zoneSettingListResponseResultZonesHttp3JSON  `json:"-"`
}

// zoneSettingListResponseResultZonesHttp3JSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesHttp3]
type zoneSettingListResponseResultZonesHttp3JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesHttp3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesHttp3) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesHttp3ID string

const (
	ZoneSettingListResponseResultZonesHttp3IDHttp3 ZoneSettingListResponseResultZonesHttp3ID = "http3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesHttp3Editable bool

const (
	ZoneSettingListResponseResultZonesHttp3EditableTrue  ZoneSettingListResponseResultZonesHttp3Editable = true
	ZoneSettingListResponseResultZonesHttp3EditableFalse ZoneSettingListResponseResultZonesHttp3Editable = false
)

// Value of the HTTP3 setting.
type ZoneSettingListResponseResultZonesHttp3Value string

const (
	ZoneSettingListResponseResultZonesHttp3ValueOn  ZoneSettingListResponseResultZonesHttp3Value = "on"
	ZoneSettingListResponseResultZonesHttp3ValueOff ZoneSettingListResponseResultZonesHttp3Value = "off"
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingListResponseResultZonesImageResizing struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesImageResizingID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value ZoneSettingListResponseResultZonesImageResizingValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesImageResizingJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesImageResizingJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesImageResizing]
type zoneSettingListResponseResultZonesImageResizingJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesImageResizing) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesImageResizingID string

const (
	ZoneSettingListResponseResultZonesImageResizingIDImageResizing ZoneSettingListResponseResultZonesImageResizingID = "image_resizing"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesImageResizingEditable bool

const (
	ZoneSettingListResponseResultZonesImageResizingEditableTrue  ZoneSettingListResponseResultZonesImageResizingEditable = true
	ZoneSettingListResponseResultZonesImageResizingEditableFalse ZoneSettingListResponseResultZonesImageResizingEditable = false
)

// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
type ZoneSettingListResponseResultZonesImageResizingValue string

const (
	ZoneSettingListResponseResultZonesImageResizingValueOn   ZoneSettingListResponseResultZonesImageResizingValue = "on"
	ZoneSettingListResponseResultZonesImageResizingValueOff  ZoneSettingListResponseResultZonesImageResizingValue = "off"
	ZoneSettingListResponseResultZonesImageResizingValueOpen ZoneSettingListResponseResultZonesImageResizingValue = "open"
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingListResponseResultZonesIPGeolocation struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesIPGeolocationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesIPGeolocationValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesIPGeolocationJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesIPGeolocationJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesIPGeolocation]
type zoneSettingListResponseResultZonesIPGeolocationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesIPGeolocation) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesIPGeolocationID string

const (
	ZoneSettingListResponseResultZonesIPGeolocationIDIPGeolocation ZoneSettingListResponseResultZonesIPGeolocationID = "ip_geolocation"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesIPGeolocationEditable bool

const (
	ZoneSettingListResponseResultZonesIPGeolocationEditableTrue  ZoneSettingListResponseResultZonesIPGeolocationEditable = true
	ZoneSettingListResponseResultZonesIPGeolocationEditableFalse ZoneSettingListResponseResultZonesIPGeolocationEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesIPGeolocationValue string

const (
	ZoneSettingListResponseResultZonesIPGeolocationValueOn  ZoneSettingListResponseResultZonesIPGeolocationValue = "on"
	ZoneSettingListResponseResultZonesIPGeolocationValueOff ZoneSettingListResponseResultZonesIPGeolocationValue = "off"
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingListResponseResultZonesIpv6 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesIpv6ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesIpv6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesIpv6Value `json:"value"`
	JSON  zoneSettingListResponseResultZonesIpv6JSON  `json:"-"`
}

// zoneSettingListResponseResultZonesIpv6JSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesIpv6]
type zoneSettingListResponseResultZonesIpv6JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesIpv6) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesIpv6ID string

const (
	ZoneSettingListResponseResultZonesIpv6IDIpv6 ZoneSettingListResponseResultZonesIpv6ID = "ipv6"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesIpv6Editable bool

const (
	ZoneSettingListResponseResultZonesIpv6EditableTrue  ZoneSettingListResponseResultZonesIpv6Editable = true
	ZoneSettingListResponseResultZonesIpv6EditableFalse ZoneSettingListResponseResultZonesIpv6Editable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesIpv6Value string

const (
	ZoneSettingListResponseResultZonesIpv6ValueOff ZoneSettingListResponseResultZonesIpv6Value = "off"
	ZoneSettingListResponseResultZonesIpv6ValueOn  ZoneSettingListResponseResultZonesIpv6Value = "on"
)

// Maximum size of an allowable upload.
type ZoneSettingListResponseResultZonesMaxUpload struct {
	// identifier of the zone setting.
	ID ZoneSettingListResponseResultZonesMaxUploadID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The size depends on the plan level of the
	// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
	Value ZoneSettingListResponseResultZonesMaxUploadValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesMaxUploadJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesMaxUploadJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesMaxUpload]
type zoneSettingListResponseResultZonesMaxUploadJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesMaxUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesMaxUpload) implementsZoneSettingListResponseResult() {}

// identifier of the zone setting.
type ZoneSettingListResponseResultZonesMaxUploadID string

const (
	ZoneSettingListResponseResultZonesMaxUploadIDMaxUpload ZoneSettingListResponseResultZonesMaxUploadID = "max_upload"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesMaxUploadEditable bool

const (
	ZoneSettingListResponseResultZonesMaxUploadEditableTrue  ZoneSettingListResponseResultZonesMaxUploadEditable = true
	ZoneSettingListResponseResultZonesMaxUploadEditableFalse ZoneSettingListResponseResultZonesMaxUploadEditable = false
)

// Value of the zone setting. Notes: The size depends on the plan level of the
// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
type ZoneSettingListResponseResultZonesMaxUploadValue float64

const (
	ZoneSettingListResponseResultZonesMaxUploadValue100 ZoneSettingListResponseResultZonesMaxUploadValue = 100
	ZoneSettingListResponseResultZonesMaxUploadValue200 ZoneSettingListResponseResultZonesMaxUploadValue = 200
	ZoneSettingListResponseResultZonesMaxUploadValue500 ZoneSettingListResponseResultZonesMaxUploadValue = 500
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingListResponseResultZonesMinTlsVersion struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesMinTlsVersionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesMinTlsVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesMinTlsVersionValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesMinTlsVersionJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesMinTlsVersionJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesMinTlsVersion]
type zoneSettingListResponseResultZonesMinTlsVersionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesMinTlsVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesMinTlsVersion) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesMinTlsVersionID string

const (
	ZoneSettingListResponseResultZonesMinTlsVersionIDMinTlsVersion ZoneSettingListResponseResultZonesMinTlsVersionID = "min_tls_version"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesMinTlsVersionEditable bool

const (
	ZoneSettingListResponseResultZonesMinTlsVersionEditableTrue  ZoneSettingListResponseResultZonesMinTlsVersionEditable = true
	ZoneSettingListResponseResultZonesMinTlsVersionEditableFalse ZoneSettingListResponseResultZonesMinTlsVersionEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesMinTlsVersionValue string

const (
	ZoneSettingListResponseResultZonesMinTlsVersionValue1_0 ZoneSettingListResponseResultZonesMinTlsVersionValue = "1.0"
	ZoneSettingListResponseResultZonesMinTlsVersionValue1_1 ZoneSettingListResponseResultZonesMinTlsVersionValue = "1.1"
	ZoneSettingListResponseResultZonesMinTlsVersionValue1_2 ZoneSettingListResponseResultZonesMinTlsVersionValue = "1.2"
	ZoneSettingListResponseResultZonesMinTlsVersionValue1_3 ZoneSettingListResponseResultZonesMinTlsVersionValue = "1.3"
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingListResponseResultZonesMinify struct {
	// Zone setting identifier.
	ID ZoneSettingListResponseResultZonesMinifyID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesMinifyValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesMinifyJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesMinifyJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesMinify]
type zoneSettingListResponseResultZonesMinifyJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesMinify) implementsZoneSettingListResponseResult() {}

// Zone setting identifier.
type ZoneSettingListResponseResultZonesMinifyID string

const (
	ZoneSettingListResponseResultZonesMinifyIDMinify ZoneSettingListResponseResultZonesMinifyID = "minify"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesMinifyEditable bool

const (
	ZoneSettingListResponseResultZonesMinifyEditableTrue  ZoneSettingListResponseResultZonesMinifyEditable = true
	ZoneSettingListResponseResultZonesMinifyEditableFalse ZoneSettingListResponseResultZonesMinifyEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css ZoneSettingListResponseResultZonesMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML ZoneSettingListResponseResultZonesMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   ZoneSettingListResponseResultZonesMinifyValueJs   `json:"js"`
	JSON zoneSettingListResponseResultZonesMinifyValueJSON `json:"-"`
}

// zoneSettingListResponseResultZonesMinifyValueJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesMinifyValue]
type zoneSettingListResponseResultZonesMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingListResponseResultZonesMinifyValueCss string

const (
	ZoneSettingListResponseResultZonesMinifyValueCssOn  ZoneSettingListResponseResultZonesMinifyValueCss = "on"
	ZoneSettingListResponseResultZonesMinifyValueCssOff ZoneSettingListResponseResultZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingListResponseResultZonesMinifyValueHTML string

const (
	ZoneSettingListResponseResultZonesMinifyValueHTMLOn  ZoneSettingListResponseResultZonesMinifyValueHTML = "on"
	ZoneSettingListResponseResultZonesMinifyValueHTMLOff ZoneSettingListResponseResultZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingListResponseResultZonesMinifyValueJs string

const (
	ZoneSettingListResponseResultZonesMinifyValueJsOn  ZoneSettingListResponseResultZonesMinifyValueJs = "on"
	ZoneSettingListResponseResultZonesMinifyValueJsOff ZoneSettingListResponseResultZonesMinifyValueJs = "off"
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingListResponseResultZonesMirage struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesMirageID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesMirageValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesMirageJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesMirageJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesMirage]
type zoneSettingListResponseResultZonesMirageJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesMirage) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesMirageID string

const (
	ZoneSettingListResponseResultZonesMirageIDMirage ZoneSettingListResponseResultZonesMirageID = "mirage"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesMirageEditable bool

const (
	ZoneSettingListResponseResultZonesMirageEditableTrue  ZoneSettingListResponseResultZonesMirageEditable = true
	ZoneSettingListResponseResultZonesMirageEditableFalse ZoneSettingListResponseResultZonesMirageEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesMirageValue string

const (
	ZoneSettingListResponseResultZonesMirageValueOn  ZoneSettingListResponseResultZonesMirageValue = "on"
	ZoneSettingListResponseResultZonesMirageValueOff ZoneSettingListResponseResultZonesMirageValue = "off"
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingListResponseResultZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID ZoneSettingListResponseResultZonesMobileRedirectID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesMobileRedirectValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesMobileRedirectJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesMobileRedirectJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesMobileRedirect]
type zoneSettingListResponseResultZonesMobileRedirectJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesMobileRedirect) implementsZoneSettingListResponseResult() {}

// Identifier of the zone setting.
type ZoneSettingListResponseResultZonesMobileRedirectID string

const (
	ZoneSettingListResponseResultZonesMobileRedirectIDMobileRedirect ZoneSettingListResponseResultZonesMobileRedirectID = "mobile_redirect"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesMobileRedirectEditable bool

const (
	ZoneSettingListResponseResultZonesMobileRedirectEditableTrue  ZoneSettingListResponseResultZonesMobileRedirectEditable = true
	ZoneSettingListResponseResultZonesMobileRedirectEditableFalse ZoneSettingListResponseResultZonesMobileRedirectEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingListResponseResultZonesMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                                      `json:"strip_uri"`
	JSON     zoneSettingListResponseResultZonesMobileRedirectValueJSON `json:"-"`
}

// zoneSettingListResponseResultZonesMobileRedirectValueJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesMobileRedirectValue]
type zoneSettingListResponseResultZonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingListResponseResultZonesMobileRedirectValueStatus string

const (
	ZoneSettingListResponseResultZonesMobileRedirectValueStatusOn  ZoneSettingListResponseResultZonesMobileRedirectValueStatus = "on"
	ZoneSettingListResponseResultZonesMobileRedirectValueStatusOff ZoneSettingListResponseResultZonesMobileRedirectValueStatus = "off"
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingListResponseResultZonesNel struct {
	// Zone setting identifier.
	ID ZoneSettingListResponseResultZonesNelID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesNelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesNelValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesNelJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesNelJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesNel]
type zoneSettingListResponseResultZonesNelJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesNel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesNel) implementsZoneSettingListResponseResult() {}

// Zone setting identifier.
type ZoneSettingListResponseResultZonesNelID string

const (
	ZoneSettingListResponseResultZonesNelIDNel ZoneSettingListResponseResultZonesNelID = "nel"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesNelEditable bool

const (
	ZoneSettingListResponseResultZonesNelEditableTrue  ZoneSettingListResponseResultZonesNelEditable = true
	ZoneSettingListResponseResultZonesNelEditableFalse ZoneSettingListResponseResultZonesNelEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesNelValue struct {
	Enabled bool                                           `json:"enabled"`
	JSON    zoneSettingListResponseResultZonesNelValueJSON `json:"-"`
}

// zoneSettingListResponseResultZonesNelValueJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesNelValue]
type zoneSettingListResponseResultZonesNelValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesNelValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingListResponseResultZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesOpportunisticEncryptionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingListResponseResultZonesOpportunisticEncryptionValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesOpportunisticEncryptionJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesOpportunisticEncryptionJSON contains the JSON
// metadata for the struct
// [ZoneSettingListResponseResultZonesOpportunisticEncryption]
type zoneSettingListResponseResultZonesOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesOpportunisticEncryption) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesOpportunisticEncryptionID string

const (
	ZoneSettingListResponseResultZonesOpportunisticEncryptionIDOpportunisticEncryption ZoneSettingListResponseResultZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesOpportunisticEncryptionEditable bool

const (
	ZoneSettingListResponseResultZonesOpportunisticEncryptionEditableTrue  ZoneSettingListResponseResultZonesOpportunisticEncryptionEditable = true
	ZoneSettingListResponseResultZonesOpportunisticEncryptionEditableFalse ZoneSettingListResponseResultZonesOpportunisticEncryptionEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingListResponseResultZonesOpportunisticEncryptionValue string

const (
	ZoneSettingListResponseResultZonesOpportunisticEncryptionValueOn  ZoneSettingListResponseResultZonesOpportunisticEncryptionValue = "on"
	ZoneSettingListResponseResultZonesOpportunisticEncryptionValueOff ZoneSettingListResponseResultZonesOpportunisticEncryptionValue = "off"
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingListResponseResultZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesOpportunisticOnionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingListResponseResultZonesOpportunisticOnionValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesOpportunisticOnionJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesOpportunisticOnionJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesOpportunisticOnion]
type zoneSettingListResponseResultZonesOpportunisticOnionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesOpportunisticOnion) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesOpportunisticOnionID string

const (
	ZoneSettingListResponseResultZonesOpportunisticOnionIDOpportunisticOnion ZoneSettingListResponseResultZonesOpportunisticOnionID = "opportunistic_onion"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesOpportunisticOnionEditable bool

const (
	ZoneSettingListResponseResultZonesOpportunisticOnionEditableTrue  ZoneSettingListResponseResultZonesOpportunisticOnionEditable = true
	ZoneSettingListResponseResultZonesOpportunisticOnionEditableFalse ZoneSettingListResponseResultZonesOpportunisticOnionEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingListResponseResultZonesOpportunisticOnionValue string

const (
	ZoneSettingListResponseResultZonesOpportunisticOnionValueOn  ZoneSettingListResponseResultZonesOpportunisticOnionValue = "on"
	ZoneSettingListResponseResultZonesOpportunisticOnionValueOff ZoneSettingListResponseResultZonesOpportunisticOnionValue = "off"
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingListResponseResultZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesOrangeToOrangeID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesOrangeToOrangeValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesOrangeToOrangeJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesOrangeToOrangeJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesOrangeToOrange]
type zoneSettingListResponseResultZonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesOrangeToOrange) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesOrangeToOrangeID string

const (
	ZoneSettingListResponseResultZonesOrangeToOrangeIDOrangeToOrange ZoneSettingListResponseResultZonesOrangeToOrangeID = "orange_to_orange"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesOrangeToOrangeEditable bool

const (
	ZoneSettingListResponseResultZonesOrangeToOrangeEditableTrue  ZoneSettingListResponseResultZonesOrangeToOrangeEditable = true
	ZoneSettingListResponseResultZonesOrangeToOrangeEditableFalse ZoneSettingListResponseResultZonesOrangeToOrangeEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesOrangeToOrangeValue string

const (
	ZoneSettingListResponseResultZonesOrangeToOrangeValueOn  ZoneSettingListResponseResultZonesOrangeToOrangeValue = "on"
	ZoneSettingListResponseResultZonesOrangeToOrangeValueOff ZoneSettingListResponseResultZonesOrangeToOrangeValue = "off"
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZoneSettingListResponseResultZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesOriginErrorPagePassThruID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesOriginErrorPagePassThruValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesOriginErrorPagePassThruJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesOriginErrorPagePassThruJSON contains the JSON
// metadata for the struct
// [ZoneSettingListResponseResultZonesOriginErrorPagePassThru]
type zoneSettingListResponseResultZonesOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesOriginErrorPagePassThru) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesOriginErrorPagePassThruID string

const (
	ZoneSettingListResponseResultZonesOriginErrorPagePassThruIDOriginErrorPagePassThru ZoneSettingListResponseResultZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesOriginErrorPagePassThruEditable bool

const (
	ZoneSettingListResponseResultZonesOriginErrorPagePassThruEditableTrue  ZoneSettingListResponseResultZonesOriginErrorPagePassThruEditable = true
	ZoneSettingListResponseResultZonesOriginErrorPagePassThruEditableFalse ZoneSettingListResponseResultZonesOriginErrorPagePassThruEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesOriginErrorPagePassThruValue string

const (
	ZoneSettingListResponseResultZonesOriginErrorPagePassThruValueOn  ZoneSettingListResponseResultZonesOriginErrorPagePassThruValue = "on"
	ZoneSettingListResponseResultZonesOriginErrorPagePassThruValueOff ZoneSettingListResponseResultZonesOriginErrorPagePassThruValue = "off"
)

type ZoneSettingListResponseResultZonesOriginMaxHTTPVersion struct {
	// Identifier of the zone setting.
	ID ZoneSettingListResponseResultZonesOriginMaxHTTPVersionID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingListResponseResultZonesOriginMaxHTTPVersionJSON `json:"-"`
}

// zoneSettingListResponseResultZonesOriginMaxHTTPVersionJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesOriginMaxHTTPVersion]
type zoneSettingListResponseResultZonesOriginMaxHTTPVersionJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesOriginMaxHTTPVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesOriginMaxHTTPVersion) implementsZoneSettingListResponseResult() {
}

// Identifier of the zone setting.
type ZoneSettingListResponseResultZonesOriginMaxHTTPVersionID string

const (
	ZoneSettingListResponseResultZonesOriginMaxHTTPVersionIDOriginMaxHTTPVersion ZoneSettingListResponseResultZonesOriginMaxHTTPVersionID = "origin_max_http_version"
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingListResponseResultZonesPolish struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesPolishID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesPolishValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesPolishJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesPolishJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesPolish]
type zoneSettingListResponseResultZonesPolishJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesPolish) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesPolishID string

const (
	ZoneSettingListResponseResultZonesPolishIDPolish ZoneSettingListResponseResultZonesPolishID = "polish"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesPolishEditable bool

const (
	ZoneSettingListResponseResultZonesPolishEditableTrue  ZoneSettingListResponseResultZonesPolishEditable = true
	ZoneSettingListResponseResultZonesPolishEditableFalse ZoneSettingListResponseResultZonesPolishEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesPolishValue string

const (
	ZoneSettingListResponseResultZonesPolishValueOff      ZoneSettingListResponseResultZonesPolishValue = "off"
	ZoneSettingListResponseResultZonesPolishValueLossless ZoneSettingListResponseResultZonesPolishValue = "lossless"
	ZoneSettingListResponseResultZonesPolishValueLossy    ZoneSettingListResponseResultZonesPolishValue = "lossy"
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingListResponseResultZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesPrefetchPreloadID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesPrefetchPreloadValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesPrefetchPreloadJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesPrefetchPreloadJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesPrefetchPreload]
type zoneSettingListResponseResultZonesPrefetchPreloadJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesPrefetchPreload) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesPrefetchPreloadID string

const (
	ZoneSettingListResponseResultZonesPrefetchPreloadIDPrefetchPreload ZoneSettingListResponseResultZonesPrefetchPreloadID = "prefetch_preload"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesPrefetchPreloadEditable bool

const (
	ZoneSettingListResponseResultZonesPrefetchPreloadEditableTrue  ZoneSettingListResponseResultZonesPrefetchPreloadEditable = true
	ZoneSettingListResponseResultZonesPrefetchPreloadEditableFalse ZoneSettingListResponseResultZonesPrefetchPreloadEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesPrefetchPreloadValue string

const (
	ZoneSettingListResponseResultZonesPrefetchPreloadValueOn  ZoneSettingListResponseResultZonesPrefetchPreloadValue = "on"
	ZoneSettingListResponseResultZonesPrefetchPreloadValueOff ZoneSettingListResponseResultZonesPrefetchPreloadValue = "off"
)

// Maximum time between two read operations from origin.
type ZoneSettingListResponseResultZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesProxyReadTimeoutID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value float64                                                `json:"value"`
	JSON  zoneSettingListResponseResultZonesProxyReadTimeoutJSON `json:"-"`
}

// zoneSettingListResponseResultZonesProxyReadTimeoutJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesProxyReadTimeout]
type zoneSettingListResponseResultZonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesProxyReadTimeout) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesProxyReadTimeoutID string

const (
	ZoneSettingListResponseResultZonesProxyReadTimeoutIDProxyReadTimeout ZoneSettingListResponseResultZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesProxyReadTimeoutEditable bool

const (
	ZoneSettingListResponseResultZonesProxyReadTimeoutEditableTrue  ZoneSettingListResponseResultZonesProxyReadTimeoutEditable = true
	ZoneSettingListResponseResultZonesProxyReadTimeoutEditableFalse ZoneSettingListResponseResultZonesProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type ZoneSettingListResponseResultZonesPseudoIpv4 struct {
	// Value of the Pseudo IPv4 setting.
	ID ZoneSettingListResponseResultZonesPseudoIpv4ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesPseudoIpv4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the Pseudo IPv4 setting.
	Value ZoneSettingListResponseResultZonesPseudoIpv4Value `json:"value"`
	JSON  zoneSettingListResponseResultZonesPseudoIpv4JSON  `json:"-"`
}

// zoneSettingListResponseResultZonesPseudoIpv4JSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesPseudoIpv4]
type zoneSettingListResponseResultZonesPseudoIpv4JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesPseudoIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesPseudoIpv4) implementsZoneSettingListResponseResult() {}

// Value of the Pseudo IPv4 setting.
type ZoneSettingListResponseResultZonesPseudoIpv4ID string

const (
	ZoneSettingListResponseResultZonesPseudoIpv4IDPseudoIpv4 ZoneSettingListResponseResultZonesPseudoIpv4ID = "pseudo_ipv4"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesPseudoIpv4Editable bool

const (
	ZoneSettingListResponseResultZonesPseudoIpv4EditableTrue  ZoneSettingListResponseResultZonesPseudoIpv4Editable = true
	ZoneSettingListResponseResultZonesPseudoIpv4EditableFalse ZoneSettingListResponseResultZonesPseudoIpv4Editable = false
)

// Value of the Pseudo IPv4 setting.
type ZoneSettingListResponseResultZonesPseudoIpv4Value string

const (
	ZoneSettingListResponseResultZonesPseudoIpv4ValueOff             ZoneSettingListResponseResultZonesPseudoIpv4Value = "off"
	ZoneSettingListResponseResultZonesPseudoIpv4ValueAddHeader       ZoneSettingListResponseResultZonesPseudoIpv4Value = "add_header"
	ZoneSettingListResponseResultZonesPseudoIpv4ValueOverwriteHeader ZoneSettingListResponseResultZonesPseudoIpv4Value = "overwrite_header"
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingListResponseResultZonesResponseBuffering struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesResponseBufferingID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesResponseBufferingValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesResponseBufferingJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesResponseBufferingJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesResponseBuffering]
type zoneSettingListResponseResultZonesResponseBufferingJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesResponseBuffering) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesResponseBufferingID string

const (
	ZoneSettingListResponseResultZonesResponseBufferingIDResponseBuffering ZoneSettingListResponseResultZonesResponseBufferingID = "response_buffering"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesResponseBufferingEditable bool

const (
	ZoneSettingListResponseResultZonesResponseBufferingEditableTrue  ZoneSettingListResponseResultZonesResponseBufferingEditable = true
	ZoneSettingListResponseResultZonesResponseBufferingEditableFalse ZoneSettingListResponseResultZonesResponseBufferingEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesResponseBufferingValue string

const (
	ZoneSettingListResponseResultZonesResponseBufferingValueOn  ZoneSettingListResponseResultZonesResponseBufferingValue = "on"
	ZoneSettingListResponseResultZonesResponseBufferingValueOff ZoneSettingListResponseResultZonesResponseBufferingValue = "off"
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
type ZoneSettingListResponseResultZonesRocketLoader struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesRocketLoaderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesRocketLoaderValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesRocketLoaderJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesRocketLoaderJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesRocketLoader]
type zoneSettingListResponseResultZonesRocketLoaderJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesRocketLoader) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesRocketLoaderID string

const (
	ZoneSettingListResponseResultZonesRocketLoaderIDRocketLoader ZoneSettingListResponseResultZonesRocketLoaderID = "rocket_loader"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesRocketLoaderEditable bool

const (
	ZoneSettingListResponseResultZonesRocketLoaderEditableTrue  ZoneSettingListResponseResultZonesRocketLoaderEditable = true
	ZoneSettingListResponseResultZonesRocketLoaderEditableFalse ZoneSettingListResponseResultZonesRocketLoaderEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesRocketLoaderValue string

const (
	ZoneSettingListResponseResultZonesRocketLoaderValueOn  ZoneSettingListResponseResultZonesRocketLoaderValue = "on"
	ZoneSettingListResponseResultZonesRocketLoaderValueOff ZoneSettingListResponseResultZonesRocketLoaderValue = "off"
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                                   `json:"modified_on,nullable" format:"date-time"`
	Value      ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationValue `json:"value"`
	JSON       zoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationJSON
// contains the JSON metadata for the struct
// [ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimization]
type zoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimization) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationID string

const (
	ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationEditable = false
)

type ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationValue struct {
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
	WpPlugin bool                                                                            `json:"wp_plugin,required"`
	JSON     zoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationValueJSON `json:"-"`
}

// zoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationValueJSON
// contains the JSON metadata for the struct
// [ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationValue]
type zoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationValueJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSchemasAutomaticPlatformOptimizationValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare security header for a zone.
type ZoneSettingListResponseResultZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID ZoneSettingListResponseResultZonesSecurityHeaderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesSecurityHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                             `json:"modified_on,nullable" format:"date-time"`
	Value      ZoneSettingListResponseResultZonesSecurityHeaderValue `json:"value"`
	JSON       zoneSettingListResponseResultZonesSecurityHeaderJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesSecurityHeaderJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesSecurityHeader]
type zoneSettingListResponseResultZonesSecurityHeaderJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesSecurityHeader) implementsZoneSettingListResponseResult() {}

// ID of the zone's security header.
type ZoneSettingListResponseResultZonesSecurityHeaderID string

const (
	ZoneSettingListResponseResultZonesSecurityHeaderIDSecurityHeader ZoneSettingListResponseResultZonesSecurityHeaderID = "security_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesSecurityHeaderEditable bool

const (
	ZoneSettingListResponseResultZonesSecurityHeaderEditableTrue  ZoneSettingListResponseResultZonesSecurityHeaderEditable = true
	ZoneSettingListResponseResultZonesSecurityHeaderEditableFalse ZoneSettingListResponseResultZonesSecurityHeaderEditable = false
)

type ZoneSettingListResponseResultZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZoneSettingListResponseResultZonesSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zoneSettingListResponseResultZonesSecurityHeaderValueJSON                    `json:"-"`
}

// zoneSettingListResponseResultZonesSecurityHeaderValueJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesSecurityHeaderValue]
type zoneSettingListResponseResultZonesSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type ZoneSettingListResponseResultZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                             `json:"nosniff"`
	JSON    zoneSettingListResponseResultZonesSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// zoneSettingListResponseResultZonesSecurityHeaderValueStrictTransportSecurityJSON
// contains the JSON metadata for the struct
// [ZoneSettingListResponseResultZonesSecurityHeaderValueStrictTransportSecurity]
type zoneSettingListResponseResultZonesSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingListResponseResultZonesSecurityLevel struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesSecurityLevelID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesSecurityLevelValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesSecurityLevelJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesSecurityLevelJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesSecurityLevel]
type zoneSettingListResponseResultZonesSecurityLevelJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesSecurityLevel) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesSecurityLevelID string

const (
	ZoneSettingListResponseResultZonesSecurityLevelIDSecurityLevel ZoneSettingListResponseResultZonesSecurityLevelID = "security_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesSecurityLevelEditable bool

const (
	ZoneSettingListResponseResultZonesSecurityLevelEditableTrue  ZoneSettingListResponseResultZonesSecurityLevelEditable = true
	ZoneSettingListResponseResultZonesSecurityLevelEditableFalse ZoneSettingListResponseResultZonesSecurityLevelEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesSecurityLevelValue string

const (
	ZoneSettingListResponseResultZonesSecurityLevelValueOff            ZoneSettingListResponseResultZonesSecurityLevelValue = "off"
	ZoneSettingListResponseResultZonesSecurityLevelValueEssentiallyOff ZoneSettingListResponseResultZonesSecurityLevelValue = "essentially_off"
	ZoneSettingListResponseResultZonesSecurityLevelValueLow            ZoneSettingListResponseResultZonesSecurityLevelValue = "low"
	ZoneSettingListResponseResultZonesSecurityLevelValueMedium         ZoneSettingListResponseResultZonesSecurityLevelValue = "medium"
	ZoneSettingListResponseResultZonesSecurityLevelValueHigh           ZoneSettingListResponseResultZonesSecurityLevelValue = "high"
	ZoneSettingListResponseResultZonesSecurityLevelValueUnderAttack    ZoneSettingListResponseResultZonesSecurityLevelValue = "under_attack"
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
type ZoneSettingListResponseResultZonesServerSideExclude struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesServerSideExcludeID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesServerSideExcludeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesServerSideExcludeValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesServerSideExcludeJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesServerSideExcludeJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesServerSideExclude]
type zoneSettingListResponseResultZonesServerSideExcludeJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesServerSideExclude) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesServerSideExcludeID string

const (
	ZoneSettingListResponseResultZonesServerSideExcludeIDServerSideExclude ZoneSettingListResponseResultZonesServerSideExcludeID = "server_side_exclude"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesServerSideExcludeEditable bool

const (
	ZoneSettingListResponseResultZonesServerSideExcludeEditableTrue  ZoneSettingListResponseResultZonesServerSideExcludeEditable = true
	ZoneSettingListResponseResultZonesServerSideExcludeEditableFalse ZoneSettingListResponseResultZonesServerSideExcludeEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesServerSideExcludeValue string

const (
	ZoneSettingListResponseResultZonesServerSideExcludeValueOn  ZoneSettingListResponseResultZonesServerSideExcludeValue = "on"
	ZoneSettingListResponseResultZonesServerSideExcludeValueOff ZoneSettingListResponseResultZonesServerSideExcludeValue = "off"
)

// Allow SHA1 support.
type ZoneSettingListResponseResultZonesSha1Support struct {
	// Zone setting identifier.
	ID ZoneSettingListResponseResultZonesSha1SupportID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesSha1SupportValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesSha1SupportJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesSha1SupportJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesSha1Support]
type zoneSettingListResponseResultZonesSha1SupportJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSha1Support) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesSha1Support) implementsZoneSettingListResponseResult() {}

// Zone setting identifier.
type ZoneSettingListResponseResultZonesSha1SupportID string

const (
	ZoneSettingListResponseResultZonesSha1SupportIDSha1Support ZoneSettingListResponseResultZonesSha1SupportID = "sha1_support"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesSha1SupportEditable bool

const (
	ZoneSettingListResponseResultZonesSha1SupportEditableTrue  ZoneSettingListResponseResultZonesSha1SupportEditable = true
	ZoneSettingListResponseResultZonesSha1SupportEditableFalse ZoneSettingListResponseResultZonesSha1SupportEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesSha1SupportValue string

const (
	ZoneSettingListResponseResultZonesSha1SupportValueOff ZoneSettingListResponseResultZonesSha1SupportValue = "off"
	ZoneSettingListResponseResultZonesSha1SupportValueOn  ZoneSettingListResponseResultZonesSha1SupportValue = "on"
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingListResponseResultZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesSortQueryStringForCacheID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesSortQueryStringForCacheValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesSortQueryStringForCacheJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesSortQueryStringForCacheJSON contains the JSON
// metadata for the struct
// [ZoneSettingListResponseResultZonesSortQueryStringForCache]
type zoneSettingListResponseResultZonesSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesSortQueryStringForCache) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesSortQueryStringForCacheID string

const (
	ZoneSettingListResponseResultZonesSortQueryStringForCacheIDSortQueryStringForCache ZoneSettingListResponseResultZonesSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesSortQueryStringForCacheEditable bool

const (
	ZoneSettingListResponseResultZonesSortQueryStringForCacheEditableTrue  ZoneSettingListResponseResultZonesSortQueryStringForCacheEditable = true
	ZoneSettingListResponseResultZonesSortQueryStringForCacheEditableFalse ZoneSettingListResponseResultZonesSortQueryStringForCacheEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesSortQueryStringForCacheValue string

const (
	ZoneSettingListResponseResultZonesSortQueryStringForCacheValueOn  ZoneSettingListResponseResultZonesSortQueryStringForCacheValue = "on"
	ZoneSettingListResponseResultZonesSortQueryStringForCacheValueOff ZoneSettingListResponseResultZonesSortQueryStringForCacheValue = "off"
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
type ZoneSettingListResponseResultZonesSsl struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesSslID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesSslEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value ZoneSettingListResponseResultZonesSslValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesSslJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesSslJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesSsl]
type zoneSettingListResponseResultZonesSslJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesSsl) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesSslID string

const (
	ZoneSettingListResponseResultZonesSslIDSsl ZoneSettingListResponseResultZonesSslID = "ssl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesSslEditable bool

const (
	ZoneSettingListResponseResultZonesSslEditableTrue  ZoneSettingListResponseResultZonesSslEditable = true
	ZoneSettingListResponseResultZonesSslEditableFalse ZoneSettingListResponseResultZonesSslEditable = false
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingListResponseResultZonesSslValue string

const (
	ZoneSettingListResponseResultZonesSslValueOff      ZoneSettingListResponseResultZonesSslValue = "off"
	ZoneSettingListResponseResultZonesSslValueFlexible ZoneSettingListResponseResultZonesSslValue = "flexible"
	ZoneSettingListResponseResultZonesSslValueFull     ZoneSettingListResponseResultZonesSslValue = "full"
	ZoneSettingListResponseResultZonesSslValueStrict   ZoneSettingListResponseResultZonesSslValue = "strict"
)

type ZoneSettingListResponseResultZonesSslRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingListResponseResultZonesSslRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                                 `json:"enabled"`
	JSON    zoneSettingListResponseResultZonesSslRecommenderJSON `json:"-"`
}

// zoneSettingListResponseResultZonesSslRecommenderJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesSslRecommender]
type zoneSettingListResponseResultZonesSslRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSslRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesSslRecommender) implementsZoneSettingListResponseResult() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingListResponseResultZonesSslRecommenderID string

const (
	ZoneSettingListResponseResultZonesSslRecommenderIDSslRecommender ZoneSettingListResponseResultZonesSslRecommenderID = "ssl_recommender"
)

// Only allows TLS1.2.
type ZoneSettingListResponseResultZonesTls1_2Only struct {
	// Zone setting identifier.
	ID ZoneSettingListResponseResultZonesTls1_2OnlyID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesTls1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesTls1_2OnlyValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesTls1_2OnlyJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesTls1_2OnlyJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesTls1_2Only]
type zoneSettingListResponseResultZonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesTls1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesTls1_2Only) implementsZoneSettingListResponseResult() {}

// Zone setting identifier.
type ZoneSettingListResponseResultZonesTls1_2OnlyID string

const (
	ZoneSettingListResponseResultZonesTls1_2OnlyIDTls1_2Only ZoneSettingListResponseResultZonesTls1_2OnlyID = "tls_1_2_only"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesTls1_2OnlyEditable bool

const (
	ZoneSettingListResponseResultZonesTls1_2OnlyEditableTrue  ZoneSettingListResponseResultZonesTls1_2OnlyEditable = true
	ZoneSettingListResponseResultZonesTls1_2OnlyEditableFalse ZoneSettingListResponseResultZonesTls1_2OnlyEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesTls1_2OnlyValue string

const (
	ZoneSettingListResponseResultZonesTls1_2OnlyValueOff ZoneSettingListResponseResultZonesTls1_2OnlyValue = "off"
	ZoneSettingListResponseResultZonesTls1_2OnlyValueOn  ZoneSettingListResponseResultZonesTls1_2OnlyValue = "on"
)

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingListResponseResultZonesTls1_3 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesTls1_3ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesTls1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingListResponseResultZonesTls1_3Value `json:"value"`
	JSON  zoneSettingListResponseResultZonesTls1_3JSON  `json:"-"`
}

// zoneSettingListResponseResultZonesTls1_3JSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesTls1_3]
type zoneSettingListResponseResultZonesTls1_3JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesTls1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesTls1_3) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesTls1_3ID string

const (
	ZoneSettingListResponseResultZonesTls1_3IDTls1_3 ZoneSettingListResponseResultZonesTls1_3ID = "tls_1_3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesTls1_3Editable bool

const (
	ZoneSettingListResponseResultZonesTls1_3EditableTrue  ZoneSettingListResponseResultZonesTls1_3Editable = true
	ZoneSettingListResponseResultZonesTls1_3EditableFalse ZoneSettingListResponseResultZonesTls1_3Editable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingListResponseResultZonesTls1_3Value string

const (
	ZoneSettingListResponseResultZonesTls1_3ValueOn  ZoneSettingListResponseResultZonesTls1_3Value = "on"
	ZoneSettingListResponseResultZonesTls1_3ValueOff ZoneSettingListResponseResultZonesTls1_3Value = "off"
	ZoneSettingListResponseResultZonesTls1_3ValueZrt ZoneSettingListResponseResultZonesTls1_3Value = "zrt"
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingListResponseResultZonesTlsClientAuth struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesTlsClientAuthID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesTlsClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value ZoneSettingListResponseResultZonesTlsClientAuthValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesTlsClientAuthJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesTlsClientAuthJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesTlsClientAuth]
type zoneSettingListResponseResultZonesTlsClientAuthJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesTlsClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesTlsClientAuth) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesTlsClientAuthID string

const (
	ZoneSettingListResponseResultZonesTlsClientAuthIDTlsClientAuth ZoneSettingListResponseResultZonesTlsClientAuthID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesTlsClientAuthEditable bool

const (
	ZoneSettingListResponseResultZonesTlsClientAuthEditableTrue  ZoneSettingListResponseResultZonesTlsClientAuthEditable = true
	ZoneSettingListResponseResultZonesTlsClientAuthEditableFalse ZoneSettingListResponseResultZonesTlsClientAuthEditable = false
)

// value of the zone setting.
type ZoneSettingListResponseResultZonesTlsClientAuthValue string

const (
	ZoneSettingListResponseResultZonesTlsClientAuthValueOn  ZoneSettingListResponseResultZonesTlsClientAuthValue = "on"
	ZoneSettingListResponseResultZonesTlsClientAuthValueOff ZoneSettingListResponseResultZonesTlsClientAuthValue = "off"
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingListResponseResultZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesTrueClientIPHeaderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesTrueClientIPHeaderValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesTrueClientIPHeaderJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesTrueClientIPHeaderJSON contains the JSON
// metadata for the struct [ZoneSettingListResponseResultZonesTrueClientIPHeader]
type zoneSettingListResponseResultZonesTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesTrueClientIPHeader) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesTrueClientIPHeaderID string

const (
	ZoneSettingListResponseResultZonesTrueClientIPHeaderIDTrueClientIPHeader ZoneSettingListResponseResultZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesTrueClientIPHeaderEditable bool

const (
	ZoneSettingListResponseResultZonesTrueClientIPHeaderEditableTrue  ZoneSettingListResponseResultZonesTrueClientIPHeaderEditable = true
	ZoneSettingListResponseResultZonesTrueClientIPHeaderEditableFalse ZoneSettingListResponseResultZonesTrueClientIPHeaderEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesTrueClientIPHeaderValue string

const (
	ZoneSettingListResponseResultZonesTrueClientIPHeaderValueOn  ZoneSettingListResponseResultZonesTrueClientIPHeaderValue = "on"
	ZoneSettingListResponseResultZonesTrueClientIPHeaderValueOff ZoneSettingListResponseResultZonesTrueClientIPHeaderValue = "off"
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
type ZoneSettingListResponseResultZonesWaf struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesWafID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesWafEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesWafValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesWafJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesWafJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesWaf]
type zoneSettingListResponseResultZonesWafJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesWaf) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesWafID string

const (
	ZoneSettingListResponseResultZonesWafIDWaf ZoneSettingListResponseResultZonesWafID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesWafEditable bool

const (
	ZoneSettingListResponseResultZonesWafEditableTrue  ZoneSettingListResponseResultZonesWafEditable = true
	ZoneSettingListResponseResultZonesWafEditableFalse ZoneSettingListResponseResultZonesWafEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesWafValue string

const (
	ZoneSettingListResponseResultZonesWafValueOn  ZoneSettingListResponseResultZonesWafValue = "on"
	ZoneSettingListResponseResultZonesWafValueOff ZoneSettingListResponseResultZonesWafValue = "off"
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingListResponseResultZonesWebp struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesWebpID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesWebpValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesWebpJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesWebpJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesWebp]
type zoneSettingListResponseResultZonesWebpJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesWebp) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesWebpID string

const (
	ZoneSettingListResponseResultZonesWebpIDWebp ZoneSettingListResponseResultZonesWebpID = "webp"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesWebpEditable bool

const (
	ZoneSettingListResponseResultZonesWebpEditableTrue  ZoneSettingListResponseResultZonesWebpEditable = true
	ZoneSettingListResponseResultZonesWebpEditableFalse ZoneSettingListResponseResultZonesWebpEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesWebpValue string

const (
	ZoneSettingListResponseResultZonesWebpValueOff ZoneSettingListResponseResultZonesWebpValue = "off"
	ZoneSettingListResponseResultZonesWebpValueOn  ZoneSettingListResponseResultZonesWebpValue = "on"
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZoneSettingListResponseResultZonesWebsockets struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesWebsocketsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesWebsocketsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesWebsocketsValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesWebsocketsJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesWebsocketsJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesWebsockets]
type zoneSettingListResponseResultZonesWebsocketsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesWebsockets) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesWebsocketsID string

const (
	ZoneSettingListResponseResultZonesWebsocketsIDWebsockets ZoneSettingListResponseResultZonesWebsocketsID = "websockets"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesWebsocketsEditable bool

const (
	ZoneSettingListResponseResultZonesWebsocketsEditableTrue  ZoneSettingListResponseResultZonesWebsocketsEditable = true
	ZoneSettingListResponseResultZonesWebsocketsEditableFalse ZoneSettingListResponseResultZonesWebsocketsEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesWebsocketsValue string

const (
	ZoneSettingListResponseResultZonesWebsocketsValueOff ZoneSettingListResponseResultZonesWebsocketsValue = "off"
	ZoneSettingListResponseResultZonesWebsocketsValueOn  ZoneSettingListResponseResultZonesWebsocketsValue = "on"
)

type ZoneSettingEditResponse struct {
	Errors   []ZoneSettingEditResponseError   `json:"errors"`
	Messages []ZoneSettingEditResponseMessage `json:"messages"`
	Result   []ZoneSettingEditResponseResult  `json:"result"`
	// Whether the API call was successful
	Success bool                        `json:"success"`
	JSON    zoneSettingEditResponseJSON `json:"-"`
}

// zoneSettingEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponse]
type zoneSettingEditResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEditResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    zoneSettingEditResponseErrorJSON `json:"-"`
}

// zoneSettingEditResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseError]
type zoneSettingEditResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingEditResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneSettingEditResponseMessageJSON `json:"-"`
}

// zoneSettingEditResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponseMessage]
type zoneSettingEditResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [ZoneSettingEditResponseResultZones0rtt],
// [ZoneSettingEditResponseResultZonesAdvancedDdos],
// [ZoneSettingEditResponseResultZonesAlwaysOnline],
// [ZoneSettingEditResponseResultZonesAlwaysUseHTTPs],
// [ZoneSettingEditResponseResultZonesAutomaticHTTPsRewrites],
// [ZoneSettingEditResponseResultZonesBrotli],
// [ZoneSettingEditResponseResultZonesBrowserCacheTtl],
// [ZoneSettingEditResponseResultZonesBrowserCheck],
// [ZoneSettingEditResponseResultZonesCacheLevel],
// [ZoneSettingEditResponseResultZonesChallengeTtl],
// [ZoneSettingEditResponseResultZonesCiphers],
// [ZoneSettingEditResponseResultZonesCnameFlattening],
// [ZoneSettingEditResponseResultZonesDevelopmentMode],
// [ZoneSettingEditResponseResultZonesEarlyHints],
// [ZoneSettingEditResponseResultZonesEdgeCacheTtl],
// [ZoneSettingEditResponseResultZonesEmailObfuscation],
// [ZoneSettingEditResponseResultZonesH2Prioritization],
// [ZoneSettingEditResponseResultZonesHotlinkProtection],
// [ZoneSettingEditResponseResultZonesHttp2],
// [ZoneSettingEditResponseResultZonesHttp3],
// [ZoneSettingEditResponseResultZonesImageResizing],
// [ZoneSettingEditResponseResultZonesIPGeolocation],
// [ZoneSettingEditResponseResultZonesIpv6],
// [ZoneSettingEditResponseResultZonesMaxUpload],
// [ZoneSettingEditResponseResultZonesMinTlsVersion],
// [ZoneSettingEditResponseResultZonesMinify],
// [ZoneSettingEditResponseResultZonesMirage],
// [ZoneSettingEditResponseResultZonesMobileRedirect],
// [ZoneSettingEditResponseResultZonesNel],
// [ZoneSettingEditResponseResultZonesOpportunisticEncryption],
// [ZoneSettingEditResponseResultZonesOpportunisticOnion],
// [ZoneSettingEditResponseResultZonesOrangeToOrange],
// [ZoneSettingEditResponseResultZonesOriginErrorPagePassThru],
// [ZoneSettingEditResponseResultZonesOriginMaxHTTPVersion],
// [ZoneSettingEditResponseResultZonesPolish],
// [ZoneSettingEditResponseResultZonesPrefetchPreload],
// [ZoneSettingEditResponseResultZonesProxyReadTimeout],
// [ZoneSettingEditResponseResultZonesPseudoIpv4],
// [ZoneSettingEditResponseResultZonesResponseBuffering],
// [ZoneSettingEditResponseResultZonesRocketLoader],
// [ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimization],
// [ZoneSettingEditResponseResultZonesSecurityHeader],
// [ZoneSettingEditResponseResultZonesSecurityLevel],
// [ZoneSettingEditResponseResultZonesServerSideExclude],
// [ZoneSettingEditResponseResultZonesSha1Support],
// [ZoneSettingEditResponseResultZonesSortQueryStringForCache],
// [ZoneSettingEditResponseResultZonesSsl],
// [ZoneSettingEditResponseResultZonesSslRecommender],
// [ZoneSettingEditResponseResultZonesTls1_2Only],
// [ZoneSettingEditResponseResultZonesTls1_3],
// [ZoneSettingEditResponseResultZonesTlsClientAuth],
// [ZoneSettingEditResponseResultZonesTrueClientIPHeader],
// [ZoneSettingEditResponseResultZonesWaf],
// [ZoneSettingEditResponseResultZonesWebp] or
// [ZoneSettingEditResponseResultZonesWebsockets].
type ZoneSettingEditResponseResult interface {
	implementsZoneSettingEditResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneSettingEditResponseResult)(nil)).Elem(), "")
}

// 0-RTT session resumption enabled for this zone.
type ZoneSettingEditResponseResultZones0rtt struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZones0rttID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZones0rttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the 0-RTT setting.
	Value ZoneSettingEditResponseResultZones0rttValue `json:"value"`
	JSON  zoneSettingEditResponseResultZones0rttJSON  `json:"-"`
}

// zoneSettingEditResponseResultZones0rttJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZones0rtt]
type zoneSettingEditResponseResultZones0rttJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZones0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZones0rtt) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZones0rttID string

const (
	ZoneSettingEditResponseResultZones0rttID0rtt ZoneSettingEditResponseResultZones0rttID = "0rtt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZones0rttEditable bool

const (
	ZoneSettingEditResponseResultZones0rttEditableTrue  ZoneSettingEditResponseResultZones0rttEditable = true
	ZoneSettingEditResponseResultZones0rttEditableFalse ZoneSettingEditResponseResultZones0rttEditable = false
)

// Value of the 0-RTT setting.
type ZoneSettingEditResponseResultZones0rttValue string

const (
	ZoneSettingEditResponseResultZones0rttValueOn  ZoneSettingEditResponseResultZones0rttValue = "on"
	ZoneSettingEditResponseResultZones0rttValueOff ZoneSettingEditResponseResultZones0rttValue = "off"
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingEditResponseResultZonesAdvancedDdos struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesAdvancedDdosID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesAdvancedDdosEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value ZoneSettingEditResponseResultZonesAdvancedDdosValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesAdvancedDdosJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesAdvancedDdosJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesAdvancedDdos]
type zoneSettingEditResponseResultZonesAdvancedDdosJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesAdvancedDdos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesAdvancedDdos) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesAdvancedDdosID string

const (
	ZoneSettingEditResponseResultZonesAdvancedDdosIDAdvancedDdos ZoneSettingEditResponseResultZonesAdvancedDdosID = "advanced_ddos"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesAdvancedDdosEditable bool

const (
	ZoneSettingEditResponseResultZonesAdvancedDdosEditableTrue  ZoneSettingEditResponseResultZonesAdvancedDdosEditable = true
	ZoneSettingEditResponseResultZonesAdvancedDdosEditableFalse ZoneSettingEditResponseResultZonesAdvancedDdosEditable = false
)

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZoneSettingEditResponseResultZonesAdvancedDdosValue string

const (
	ZoneSettingEditResponseResultZonesAdvancedDdosValueOn  ZoneSettingEditResponseResultZonesAdvancedDdosValue = "on"
	ZoneSettingEditResponseResultZonesAdvancedDdosValueOff ZoneSettingEditResponseResultZonesAdvancedDdosValue = "off"
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingEditResponseResultZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesAlwaysOnlineID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesAlwaysOnlineValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesAlwaysOnlineJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesAlwaysOnlineJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesAlwaysOnline]
type zoneSettingEditResponseResultZonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesAlwaysOnline) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesAlwaysOnlineID string

const (
	ZoneSettingEditResponseResultZonesAlwaysOnlineIDAlwaysOnline ZoneSettingEditResponseResultZonesAlwaysOnlineID = "always_online"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesAlwaysOnlineEditable bool

const (
	ZoneSettingEditResponseResultZonesAlwaysOnlineEditableTrue  ZoneSettingEditResponseResultZonesAlwaysOnlineEditable = true
	ZoneSettingEditResponseResultZonesAlwaysOnlineEditableFalse ZoneSettingEditResponseResultZonesAlwaysOnlineEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesAlwaysOnlineValue string

const (
	ZoneSettingEditResponseResultZonesAlwaysOnlineValueOn  ZoneSettingEditResponseResultZonesAlwaysOnlineValue = "on"
	ZoneSettingEditResponseResultZonesAlwaysOnlineValueOff ZoneSettingEditResponseResultZonesAlwaysOnlineValue = "off"
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingEditResponseResultZonesAlwaysUseHTTPs struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesAlwaysUseHTTPsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesAlwaysUseHTTPsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesAlwaysUseHTTPsValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesAlwaysUseHTTPsJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesAlwaysUseHTTPsJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesAlwaysUseHTTPs]
type zoneSettingEditResponseResultZonesAlwaysUseHTTPsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesAlwaysUseHTTPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesAlwaysUseHTTPs) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesAlwaysUseHTTPsID string

const (
	ZoneSettingEditResponseResultZonesAlwaysUseHTTPsIDAlwaysUseHTTPs ZoneSettingEditResponseResultZonesAlwaysUseHTTPsID = "always_use_https"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesAlwaysUseHTTPsEditable bool

const (
	ZoneSettingEditResponseResultZonesAlwaysUseHTTPsEditableTrue  ZoneSettingEditResponseResultZonesAlwaysUseHTTPsEditable = true
	ZoneSettingEditResponseResultZonesAlwaysUseHTTPsEditableFalse ZoneSettingEditResponseResultZonesAlwaysUseHTTPsEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesAlwaysUseHTTPsValue string

const (
	ZoneSettingEditResponseResultZonesAlwaysUseHTTPsValueOn  ZoneSettingEditResponseResultZonesAlwaysUseHTTPsValue = "on"
	ZoneSettingEditResponseResultZonesAlwaysUseHTTPsValueOff ZoneSettingEditResponseResultZonesAlwaysUseHTTPsValue = "off"
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingEditResponseResultZonesAutomaticHTTPsRewrites struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesAutomaticHTTPsRewritesJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesAutomaticHTTPsRewritesJSON contains the JSON
// metadata for the struct
// [ZoneSettingEditResponseResultZonesAutomaticHTTPsRewrites]
type zoneSettingEditResponseResultZonesAutomaticHTTPsRewritesJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesAutomaticHTTPsRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesAutomaticHTTPsRewrites) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesID string

const (
	ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesIDAutomaticHTTPsRewrites ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesID = "automatic_https_rewrites"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesEditable bool

const (
	ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesEditableTrue  ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesEditable = true
	ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesEditableFalse ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesValue string

const (
	ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesValueOn  ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesValue = "on"
	ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesValueOff ZoneSettingEditResponseResultZonesAutomaticHTTPsRewritesValue = "off"
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingEditResponseResultZonesBrotli struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesBrotliID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesBrotliValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesBrotliJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesBrotliJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesBrotli]
type zoneSettingEditResponseResultZonesBrotliJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesBrotli) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesBrotliID string

const (
	ZoneSettingEditResponseResultZonesBrotliIDBrotli ZoneSettingEditResponseResultZonesBrotliID = "brotli"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesBrotliEditable bool

const (
	ZoneSettingEditResponseResultZonesBrotliEditableTrue  ZoneSettingEditResponseResultZonesBrotliEditable = true
	ZoneSettingEditResponseResultZonesBrotliEditableFalse ZoneSettingEditResponseResultZonesBrotliEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesBrotliValue string

const (
	ZoneSettingEditResponseResultZonesBrotliValueOff ZoneSettingEditResponseResultZonesBrotliValue = "off"
	ZoneSettingEditResponseResultZonesBrotliValueOn  ZoneSettingEditResponseResultZonesBrotliValue = "on"
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingEditResponseResultZonesBrowserCacheTtl struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesBrowserCacheTtlID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesBrowserCacheTtlEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingEditResponseResultZonesBrowserCacheTtlValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesBrowserCacheTtlJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesBrowserCacheTtlJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesBrowserCacheTtl]
type zoneSettingEditResponseResultZonesBrowserCacheTtlJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesBrowserCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesBrowserCacheTtl) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesBrowserCacheTtlID string

const (
	ZoneSettingEditResponseResultZonesBrowserCacheTtlIDBrowserCacheTtl ZoneSettingEditResponseResultZonesBrowserCacheTtlID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesBrowserCacheTtlEditable bool

const (
	ZoneSettingEditResponseResultZonesBrowserCacheTtlEditableTrue  ZoneSettingEditResponseResultZonesBrowserCacheTtlEditable = true
	ZoneSettingEditResponseResultZonesBrowserCacheTtlEditableFalse ZoneSettingEditResponseResultZonesBrowserCacheTtlEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingEditResponseResultZonesBrowserCacheTtlValue float64

const (
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue0        ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 0
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue30       ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 30
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue60       ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 60
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue120      ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 120
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue300      ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 300
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue1200     ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 1200
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue1800     ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 1800
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue3600     ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 3600
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue7200     ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 7200
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue10800    ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 10800
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue14400    ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 14400
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue18000    ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 18000
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue28800    ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 28800
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue43200    ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 43200
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue57600    ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 57600
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue72000    ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 72000
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue86400    ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 86400
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue172800   ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 172800
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue259200   ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 259200
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue345600   ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 345600
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue432000   ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 432000
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue691200   ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 691200
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue1382400  ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 1382400
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue2073600  ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 2073600
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue2678400  ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 2678400
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue5356800  ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 5356800
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue16070400 ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 16070400
	ZoneSettingEditResponseResultZonesBrowserCacheTtlValue31536000 ZoneSettingEditResponseResultZonesBrowserCacheTtlValue = 31536000
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingEditResponseResultZonesBrowserCheck struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesBrowserCheckID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesBrowserCheckValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesBrowserCheckJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesBrowserCheckJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesBrowserCheck]
type zoneSettingEditResponseResultZonesBrowserCheckJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesBrowserCheck) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesBrowserCheckID string

const (
	ZoneSettingEditResponseResultZonesBrowserCheckIDBrowserCheck ZoneSettingEditResponseResultZonesBrowserCheckID = "browser_check"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesBrowserCheckEditable bool

const (
	ZoneSettingEditResponseResultZonesBrowserCheckEditableTrue  ZoneSettingEditResponseResultZonesBrowserCheckEditable = true
	ZoneSettingEditResponseResultZonesBrowserCheckEditableFalse ZoneSettingEditResponseResultZonesBrowserCheckEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesBrowserCheckValue string

const (
	ZoneSettingEditResponseResultZonesBrowserCheckValueOn  ZoneSettingEditResponseResultZonesBrowserCheckValue = "on"
	ZoneSettingEditResponseResultZonesBrowserCheckValueOff ZoneSettingEditResponseResultZonesBrowserCheckValue = "off"
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingEditResponseResultZonesCacheLevel struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesCacheLevelID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesCacheLevelValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesCacheLevelJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesCacheLevelJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesCacheLevel]
type zoneSettingEditResponseResultZonesCacheLevelJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesCacheLevel) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesCacheLevelID string

const (
	ZoneSettingEditResponseResultZonesCacheLevelIDCacheLevel ZoneSettingEditResponseResultZonesCacheLevelID = "cache_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesCacheLevelEditable bool

const (
	ZoneSettingEditResponseResultZonesCacheLevelEditableTrue  ZoneSettingEditResponseResultZonesCacheLevelEditable = true
	ZoneSettingEditResponseResultZonesCacheLevelEditableFalse ZoneSettingEditResponseResultZonesCacheLevelEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesCacheLevelValue string

const (
	ZoneSettingEditResponseResultZonesCacheLevelValueAggressive ZoneSettingEditResponseResultZonesCacheLevelValue = "aggressive"
	ZoneSettingEditResponseResultZonesCacheLevelValueBasic      ZoneSettingEditResponseResultZonesCacheLevelValue = "basic"
	ZoneSettingEditResponseResultZonesCacheLevelValueSimplified ZoneSettingEditResponseResultZonesCacheLevelValue = "simplified"
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingEditResponseResultZonesChallengeTtl struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesChallengeTtlID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesChallengeTtlEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesChallengeTtlValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesChallengeTtlJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesChallengeTtlJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesChallengeTtl]
type zoneSettingEditResponseResultZonesChallengeTtlJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesChallengeTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesChallengeTtl) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesChallengeTtlID string

const (
	ZoneSettingEditResponseResultZonesChallengeTtlIDChallengeTtl ZoneSettingEditResponseResultZonesChallengeTtlID = "challenge_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesChallengeTtlEditable bool

const (
	ZoneSettingEditResponseResultZonesChallengeTtlEditableTrue  ZoneSettingEditResponseResultZonesChallengeTtlEditable = true
	ZoneSettingEditResponseResultZonesChallengeTtlEditableFalse ZoneSettingEditResponseResultZonesChallengeTtlEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesChallengeTtlValue float64

const (
	ZoneSettingEditResponseResultZonesChallengeTtlValue300      ZoneSettingEditResponseResultZonesChallengeTtlValue = 300
	ZoneSettingEditResponseResultZonesChallengeTtlValue900      ZoneSettingEditResponseResultZonesChallengeTtlValue = 900
	ZoneSettingEditResponseResultZonesChallengeTtlValue1800     ZoneSettingEditResponseResultZonesChallengeTtlValue = 1800
	ZoneSettingEditResponseResultZonesChallengeTtlValue2700     ZoneSettingEditResponseResultZonesChallengeTtlValue = 2700
	ZoneSettingEditResponseResultZonesChallengeTtlValue3600     ZoneSettingEditResponseResultZonesChallengeTtlValue = 3600
	ZoneSettingEditResponseResultZonesChallengeTtlValue7200     ZoneSettingEditResponseResultZonesChallengeTtlValue = 7200
	ZoneSettingEditResponseResultZonesChallengeTtlValue10800    ZoneSettingEditResponseResultZonesChallengeTtlValue = 10800
	ZoneSettingEditResponseResultZonesChallengeTtlValue14400    ZoneSettingEditResponseResultZonesChallengeTtlValue = 14400
	ZoneSettingEditResponseResultZonesChallengeTtlValue28800    ZoneSettingEditResponseResultZonesChallengeTtlValue = 28800
	ZoneSettingEditResponseResultZonesChallengeTtlValue57600    ZoneSettingEditResponseResultZonesChallengeTtlValue = 57600
	ZoneSettingEditResponseResultZonesChallengeTtlValue86400    ZoneSettingEditResponseResultZonesChallengeTtlValue = 86400
	ZoneSettingEditResponseResultZonesChallengeTtlValue604800   ZoneSettingEditResponseResultZonesChallengeTtlValue = 604800
	ZoneSettingEditResponseResultZonesChallengeTtlValue2592000  ZoneSettingEditResponseResultZonesChallengeTtlValue = 2592000
	ZoneSettingEditResponseResultZonesChallengeTtlValue31536000 ZoneSettingEditResponseResultZonesChallengeTtlValue = 31536000
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingEditResponseResultZonesCiphers struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesCiphersID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesCiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value []string                                      `json:"value"`
	JSON  zoneSettingEditResponseResultZonesCiphersJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesCiphersJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesCiphers]
type zoneSettingEditResponseResultZonesCiphersJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesCiphers) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesCiphersID string

const (
	ZoneSettingEditResponseResultZonesCiphersIDCiphers ZoneSettingEditResponseResultZonesCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesCiphersEditable bool

const (
	ZoneSettingEditResponseResultZonesCiphersEditableTrue  ZoneSettingEditResponseResultZonesCiphersEditable = true
	ZoneSettingEditResponseResultZonesCiphersEditableFalse ZoneSettingEditResponseResultZonesCiphersEditable = false
)

// Whether or not cname flattening is on.
type ZoneSettingEditResponseResultZonesCnameFlattening struct {
	// How to flatten the cname destination.
	ID ZoneSettingEditResponseResultZonesCnameFlatteningID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesCnameFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the cname flattening setting.
	Value ZoneSettingEditResponseResultZonesCnameFlatteningValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesCnameFlatteningJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesCnameFlatteningJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesCnameFlattening]
type zoneSettingEditResponseResultZonesCnameFlatteningJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesCnameFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesCnameFlattening) implementsZoneSettingEditResponseResult() {
}

// How to flatten the cname destination.
type ZoneSettingEditResponseResultZonesCnameFlatteningID string

const (
	ZoneSettingEditResponseResultZonesCnameFlatteningIDCnameFlattening ZoneSettingEditResponseResultZonesCnameFlatteningID = "cname_flattening"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesCnameFlatteningEditable bool

const (
	ZoneSettingEditResponseResultZonesCnameFlatteningEditableTrue  ZoneSettingEditResponseResultZonesCnameFlatteningEditable = true
	ZoneSettingEditResponseResultZonesCnameFlatteningEditableFalse ZoneSettingEditResponseResultZonesCnameFlatteningEditable = false
)

// Value of the cname flattening setting.
type ZoneSettingEditResponseResultZonesCnameFlatteningValue string

const (
	ZoneSettingEditResponseResultZonesCnameFlatteningValueFlattenAtRoot ZoneSettingEditResponseResultZonesCnameFlatteningValue = "flatten_at_root"
	ZoneSettingEditResponseResultZonesCnameFlatteningValueFlattenAll    ZoneSettingEditResponseResultZonesCnameFlatteningValue = "flatten_all"
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingEditResponseResultZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesDevelopmentModeID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesDevelopmentModeValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesDevelopmentModeJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesDevelopmentModeJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesDevelopmentMode]
type zoneSettingEditResponseResultZonesDevelopmentModeJSON struct {
	ID            apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesDevelopmentMode) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesDevelopmentModeID string

const (
	ZoneSettingEditResponseResultZonesDevelopmentModeIDDevelopmentMode ZoneSettingEditResponseResultZonesDevelopmentModeID = "development_mode"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesDevelopmentModeEditable bool

const (
	ZoneSettingEditResponseResultZonesDevelopmentModeEditableTrue  ZoneSettingEditResponseResultZonesDevelopmentModeEditable = true
	ZoneSettingEditResponseResultZonesDevelopmentModeEditableFalse ZoneSettingEditResponseResultZonesDevelopmentModeEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesDevelopmentModeValue string

const (
	ZoneSettingEditResponseResultZonesDevelopmentModeValueOn  ZoneSettingEditResponseResultZonesDevelopmentModeValue = "on"
	ZoneSettingEditResponseResultZonesDevelopmentModeValueOff ZoneSettingEditResponseResultZonesDevelopmentModeValue = "off"
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingEditResponseResultZonesEarlyHints struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesEarlyHintsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesEarlyHintsValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesEarlyHintsJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesEarlyHintsJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesEarlyHints]
type zoneSettingEditResponseResultZonesEarlyHintsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesEarlyHints) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesEarlyHintsID string

const (
	ZoneSettingEditResponseResultZonesEarlyHintsIDEarlyHints ZoneSettingEditResponseResultZonesEarlyHintsID = "early_hints"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesEarlyHintsEditable bool

const (
	ZoneSettingEditResponseResultZonesEarlyHintsEditableTrue  ZoneSettingEditResponseResultZonesEarlyHintsEditable = true
	ZoneSettingEditResponseResultZonesEarlyHintsEditableFalse ZoneSettingEditResponseResultZonesEarlyHintsEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesEarlyHintsValue string

const (
	ZoneSettingEditResponseResultZonesEarlyHintsValueOn  ZoneSettingEditResponseResultZonesEarlyHintsValue = "on"
	ZoneSettingEditResponseResultZonesEarlyHintsValueOff ZoneSettingEditResponseResultZonesEarlyHintsValue = "off"
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZoneSettingEditResponseResultZonesEdgeCacheTtl struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesEdgeCacheTtlID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesEdgeCacheTtlEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value ZoneSettingEditResponseResultZonesEdgeCacheTtlValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesEdgeCacheTtlJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesEdgeCacheTtlJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesEdgeCacheTtl]
type zoneSettingEditResponseResultZonesEdgeCacheTtlJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesEdgeCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesEdgeCacheTtl) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesEdgeCacheTtlID string

const (
	ZoneSettingEditResponseResultZonesEdgeCacheTtlIDEdgeCacheTtl ZoneSettingEditResponseResultZonesEdgeCacheTtlID = "edge_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesEdgeCacheTtlEditable bool

const (
	ZoneSettingEditResponseResultZonesEdgeCacheTtlEditableTrue  ZoneSettingEditResponseResultZonesEdgeCacheTtlEditable = true
	ZoneSettingEditResponseResultZonesEdgeCacheTtlEditableFalse ZoneSettingEditResponseResultZonesEdgeCacheTtlEditable = false
)

// Value of the zone setting. Notes: The minimum TTL available depends on the plan
// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
type ZoneSettingEditResponseResultZonesEdgeCacheTtlValue float64

const (
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue30     ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 30
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue60     ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 60
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue300    ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 300
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue1200   ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 1200
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue1800   ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 1800
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue3600   ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 3600
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue7200   ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 7200
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue10800  ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 10800
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue14400  ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 14400
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue18000  ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 18000
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue28800  ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 28800
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue43200  ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 43200
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue57600  ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 57600
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue72000  ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 72000
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue86400  ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 86400
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue172800 ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 172800
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue259200 ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 259200
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue345600 ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 345600
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue432000 ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 432000
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue518400 ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 518400
	ZoneSettingEditResponseResultZonesEdgeCacheTtlValue604800 ZoneSettingEditResponseResultZonesEdgeCacheTtlValue = 604800
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingEditResponseResultZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesEmailObfuscationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesEmailObfuscationValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesEmailObfuscationJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesEmailObfuscationJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesEmailObfuscation]
type zoneSettingEditResponseResultZonesEmailObfuscationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesEmailObfuscation) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesEmailObfuscationID string

const (
	ZoneSettingEditResponseResultZonesEmailObfuscationIDEmailObfuscation ZoneSettingEditResponseResultZonesEmailObfuscationID = "email_obfuscation"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesEmailObfuscationEditable bool

const (
	ZoneSettingEditResponseResultZonesEmailObfuscationEditableTrue  ZoneSettingEditResponseResultZonesEmailObfuscationEditable = true
	ZoneSettingEditResponseResultZonesEmailObfuscationEditableFalse ZoneSettingEditResponseResultZonesEmailObfuscationEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesEmailObfuscationValue string

const (
	ZoneSettingEditResponseResultZonesEmailObfuscationValueOn  ZoneSettingEditResponseResultZonesEmailObfuscationValue = "on"
	ZoneSettingEditResponseResultZonesEmailObfuscationValueOff ZoneSettingEditResponseResultZonesEmailObfuscationValue = "off"
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingEditResponseResultZonesH2Prioritization struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesH2PrioritizationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesH2PrioritizationValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesH2PrioritizationJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesH2PrioritizationJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesH2Prioritization]
type zoneSettingEditResponseResultZonesH2PrioritizationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesH2Prioritization) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesH2PrioritizationID string

const (
	ZoneSettingEditResponseResultZonesH2PrioritizationIDH2Prioritization ZoneSettingEditResponseResultZonesH2PrioritizationID = "h2_prioritization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesH2PrioritizationEditable bool

const (
	ZoneSettingEditResponseResultZonesH2PrioritizationEditableTrue  ZoneSettingEditResponseResultZonesH2PrioritizationEditable = true
	ZoneSettingEditResponseResultZonesH2PrioritizationEditableFalse ZoneSettingEditResponseResultZonesH2PrioritizationEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesH2PrioritizationValue string

const (
	ZoneSettingEditResponseResultZonesH2PrioritizationValueOn     ZoneSettingEditResponseResultZonesH2PrioritizationValue = "on"
	ZoneSettingEditResponseResultZonesH2PrioritizationValueOff    ZoneSettingEditResponseResultZonesH2PrioritizationValue = "off"
	ZoneSettingEditResponseResultZonesH2PrioritizationValueCustom ZoneSettingEditResponseResultZonesH2PrioritizationValue = "custom"
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZoneSettingEditResponseResultZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesHotlinkProtectionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesHotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesHotlinkProtectionValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesHotlinkProtectionJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesHotlinkProtectionJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesHotlinkProtection]
type zoneSettingEditResponseResultZonesHotlinkProtectionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesHotlinkProtection) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesHotlinkProtectionID string

const (
	ZoneSettingEditResponseResultZonesHotlinkProtectionIDHotlinkProtection ZoneSettingEditResponseResultZonesHotlinkProtectionID = "hotlink_protection"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesHotlinkProtectionEditable bool

const (
	ZoneSettingEditResponseResultZonesHotlinkProtectionEditableTrue  ZoneSettingEditResponseResultZonesHotlinkProtectionEditable = true
	ZoneSettingEditResponseResultZonesHotlinkProtectionEditableFalse ZoneSettingEditResponseResultZonesHotlinkProtectionEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesHotlinkProtectionValue string

const (
	ZoneSettingEditResponseResultZonesHotlinkProtectionValueOn  ZoneSettingEditResponseResultZonesHotlinkProtectionValue = "on"
	ZoneSettingEditResponseResultZonesHotlinkProtectionValueOff ZoneSettingEditResponseResultZonesHotlinkProtectionValue = "off"
)

// HTTP2 enabled for this zone.
type ZoneSettingEditResponseResultZonesHttp2 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesHttp2ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesHttp2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value ZoneSettingEditResponseResultZonesHttp2Value `json:"value"`
	JSON  zoneSettingEditResponseResultZonesHttp2JSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesHttp2JSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesHttp2]
type zoneSettingEditResponseResultZonesHttp2JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesHttp2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesHttp2) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesHttp2ID string

const (
	ZoneSettingEditResponseResultZonesHttp2IDHttp2 ZoneSettingEditResponseResultZonesHttp2ID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesHttp2Editable bool

const (
	ZoneSettingEditResponseResultZonesHttp2EditableTrue  ZoneSettingEditResponseResultZonesHttp2Editable = true
	ZoneSettingEditResponseResultZonesHttp2EditableFalse ZoneSettingEditResponseResultZonesHttp2Editable = false
)

// Value of the HTTP2 setting.
type ZoneSettingEditResponseResultZonesHttp2Value string

const (
	ZoneSettingEditResponseResultZonesHttp2ValueOn  ZoneSettingEditResponseResultZonesHttp2Value = "on"
	ZoneSettingEditResponseResultZonesHttp2ValueOff ZoneSettingEditResponseResultZonesHttp2Value = "off"
)

// HTTP3 enabled for this zone.
type ZoneSettingEditResponseResultZonesHttp3 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesHttp3ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesHttp3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP3 setting.
	Value ZoneSettingEditResponseResultZonesHttp3Value `json:"value"`
	JSON  zoneSettingEditResponseResultZonesHttp3JSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesHttp3JSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesHttp3]
type zoneSettingEditResponseResultZonesHttp3JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesHttp3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesHttp3) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesHttp3ID string

const (
	ZoneSettingEditResponseResultZonesHttp3IDHttp3 ZoneSettingEditResponseResultZonesHttp3ID = "http3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesHttp3Editable bool

const (
	ZoneSettingEditResponseResultZonesHttp3EditableTrue  ZoneSettingEditResponseResultZonesHttp3Editable = true
	ZoneSettingEditResponseResultZonesHttp3EditableFalse ZoneSettingEditResponseResultZonesHttp3Editable = false
)

// Value of the HTTP3 setting.
type ZoneSettingEditResponseResultZonesHttp3Value string

const (
	ZoneSettingEditResponseResultZonesHttp3ValueOn  ZoneSettingEditResponseResultZonesHttp3Value = "on"
	ZoneSettingEditResponseResultZonesHttp3ValueOff ZoneSettingEditResponseResultZonesHttp3Value = "off"
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingEditResponseResultZonesImageResizing struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesImageResizingID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value ZoneSettingEditResponseResultZonesImageResizingValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesImageResizingJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesImageResizingJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesImageResizing]
type zoneSettingEditResponseResultZonesImageResizingJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesImageResizing) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesImageResizingID string

const (
	ZoneSettingEditResponseResultZonesImageResizingIDImageResizing ZoneSettingEditResponseResultZonesImageResizingID = "image_resizing"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesImageResizingEditable bool

const (
	ZoneSettingEditResponseResultZonesImageResizingEditableTrue  ZoneSettingEditResponseResultZonesImageResizingEditable = true
	ZoneSettingEditResponseResultZonesImageResizingEditableFalse ZoneSettingEditResponseResultZonesImageResizingEditable = false
)

// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
type ZoneSettingEditResponseResultZonesImageResizingValue string

const (
	ZoneSettingEditResponseResultZonesImageResizingValueOn   ZoneSettingEditResponseResultZonesImageResizingValue = "on"
	ZoneSettingEditResponseResultZonesImageResizingValueOff  ZoneSettingEditResponseResultZonesImageResizingValue = "off"
	ZoneSettingEditResponseResultZonesImageResizingValueOpen ZoneSettingEditResponseResultZonesImageResizingValue = "open"
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingEditResponseResultZonesIPGeolocation struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesIPGeolocationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesIPGeolocationValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesIPGeolocationJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesIPGeolocationJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesIPGeolocation]
type zoneSettingEditResponseResultZonesIPGeolocationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesIPGeolocation) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesIPGeolocationID string

const (
	ZoneSettingEditResponseResultZonesIPGeolocationIDIPGeolocation ZoneSettingEditResponseResultZonesIPGeolocationID = "ip_geolocation"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesIPGeolocationEditable bool

const (
	ZoneSettingEditResponseResultZonesIPGeolocationEditableTrue  ZoneSettingEditResponseResultZonesIPGeolocationEditable = true
	ZoneSettingEditResponseResultZonesIPGeolocationEditableFalse ZoneSettingEditResponseResultZonesIPGeolocationEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesIPGeolocationValue string

const (
	ZoneSettingEditResponseResultZonesIPGeolocationValueOn  ZoneSettingEditResponseResultZonesIPGeolocationValue = "on"
	ZoneSettingEditResponseResultZonesIPGeolocationValueOff ZoneSettingEditResponseResultZonesIPGeolocationValue = "off"
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingEditResponseResultZonesIpv6 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesIpv6ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesIpv6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesIpv6Value `json:"value"`
	JSON  zoneSettingEditResponseResultZonesIpv6JSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesIpv6JSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesIpv6]
type zoneSettingEditResponseResultZonesIpv6JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesIpv6) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesIpv6ID string

const (
	ZoneSettingEditResponseResultZonesIpv6IDIpv6 ZoneSettingEditResponseResultZonesIpv6ID = "ipv6"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesIpv6Editable bool

const (
	ZoneSettingEditResponseResultZonesIpv6EditableTrue  ZoneSettingEditResponseResultZonesIpv6Editable = true
	ZoneSettingEditResponseResultZonesIpv6EditableFalse ZoneSettingEditResponseResultZonesIpv6Editable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesIpv6Value string

const (
	ZoneSettingEditResponseResultZonesIpv6ValueOff ZoneSettingEditResponseResultZonesIpv6Value = "off"
	ZoneSettingEditResponseResultZonesIpv6ValueOn  ZoneSettingEditResponseResultZonesIpv6Value = "on"
)

// Maximum size of an allowable upload.
type ZoneSettingEditResponseResultZonesMaxUpload struct {
	// identifier of the zone setting.
	ID ZoneSettingEditResponseResultZonesMaxUploadID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The size depends on the plan level of the
	// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
	Value ZoneSettingEditResponseResultZonesMaxUploadValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesMaxUploadJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesMaxUploadJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesMaxUpload]
type zoneSettingEditResponseResultZonesMaxUploadJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesMaxUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesMaxUpload) implementsZoneSettingEditResponseResult() {}

// identifier of the zone setting.
type ZoneSettingEditResponseResultZonesMaxUploadID string

const (
	ZoneSettingEditResponseResultZonesMaxUploadIDMaxUpload ZoneSettingEditResponseResultZonesMaxUploadID = "max_upload"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesMaxUploadEditable bool

const (
	ZoneSettingEditResponseResultZonesMaxUploadEditableTrue  ZoneSettingEditResponseResultZonesMaxUploadEditable = true
	ZoneSettingEditResponseResultZonesMaxUploadEditableFalse ZoneSettingEditResponseResultZonesMaxUploadEditable = false
)

// Value of the zone setting. Notes: The size depends on the plan level of the
// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
type ZoneSettingEditResponseResultZonesMaxUploadValue float64

const (
	ZoneSettingEditResponseResultZonesMaxUploadValue100 ZoneSettingEditResponseResultZonesMaxUploadValue = 100
	ZoneSettingEditResponseResultZonesMaxUploadValue200 ZoneSettingEditResponseResultZonesMaxUploadValue = 200
	ZoneSettingEditResponseResultZonesMaxUploadValue500 ZoneSettingEditResponseResultZonesMaxUploadValue = 500
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingEditResponseResultZonesMinTlsVersion struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesMinTlsVersionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesMinTlsVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesMinTlsVersionValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesMinTlsVersionJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesMinTlsVersionJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesMinTlsVersion]
type zoneSettingEditResponseResultZonesMinTlsVersionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesMinTlsVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesMinTlsVersion) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesMinTlsVersionID string

const (
	ZoneSettingEditResponseResultZonesMinTlsVersionIDMinTlsVersion ZoneSettingEditResponseResultZonesMinTlsVersionID = "min_tls_version"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesMinTlsVersionEditable bool

const (
	ZoneSettingEditResponseResultZonesMinTlsVersionEditableTrue  ZoneSettingEditResponseResultZonesMinTlsVersionEditable = true
	ZoneSettingEditResponseResultZonesMinTlsVersionEditableFalse ZoneSettingEditResponseResultZonesMinTlsVersionEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesMinTlsVersionValue string

const (
	ZoneSettingEditResponseResultZonesMinTlsVersionValue1_0 ZoneSettingEditResponseResultZonesMinTlsVersionValue = "1.0"
	ZoneSettingEditResponseResultZonesMinTlsVersionValue1_1 ZoneSettingEditResponseResultZonesMinTlsVersionValue = "1.1"
	ZoneSettingEditResponseResultZonesMinTlsVersionValue1_2 ZoneSettingEditResponseResultZonesMinTlsVersionValue = "1.2"
	ZoneSettingEditResponseResultZonesMinTlsVersionValue1_3 ZoneSettingEditResponseResultZonesMinTlsVersionValue = "1.3"
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingEditResponseResultZonesMinify struct {
	// Zone setting identifier.
	ID ZoneSettingEditResponseResultZonesMinifyID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesMinifyValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesMinifyJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesMinifyJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesMinify]
type zoneSettingEditResponseResultZonesMinifyJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesMinify) implementsZoneSettingEditResponseResult() {}

// Zone setting identifier.
type ZoneSettingEditResponseResultZonesMinifyID string

const (
	ZoneSettingEditResponseResultZonesMinifyIDMinify ZoneSettingEditResponseResultZonesMinifyID = "minify"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesMinifyEditable bool

const (
	ZoneSettingEditResponseResultZonesMinifyEditableTrue  ZoneSettingEditResponseResultZonesMinifyEditable = true
	ZoneSettingEditResponseResultZonesMinifyEditableFalse ZoneSettingEditResponseResultZonesMinifyEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css ZoneSettingEditResponseResultZonesMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML ZoneSettingEditResponseResultZonesMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   ZoneSettingEditResponseResultZonesMinifyValueJs   `json:"js"`
	JSON zoneSettingEditResponseResultZonesMinifyValueJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesMinifyValueJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesMinifyValue]
type zoneSettingEditResponseResultZonesMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingEditResponseResultZonesMinifyValueCss string

const (
	ZoneSettingEditResponseResultZonesMinifyValueCssOn  ZoneSettingEditResponseResultZonesMinifyValueCss = "on"
	ZoneSettingEditResponseResultZonesMinifyValueCssOff ZoneSettingEditResponseResultZonesMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingEditResponseResultZonesMinifyValueHTML string

const (
	ZoneSettingEditResponseResultZonesMinifyValueHTMLOn  ZoneSettingEditResponseResultZonesMinifyValueHTML = "on"
	ZoneSettingEditResponseResultZonesMinifyValueHTMLOff ZoneSettingEditResponseResultZonesMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingEditResponseResultZonesMinifyValueJs string

const (
	ZoneSettingEditResponseResultZonesMinifyValueJsOn  ZoneSettingEditResponseResultZonesMinifyValueJs = "on"
	ZoneSettingEditResponseResultZonesMinifyValueJsOff ZoneSettingEditResponseResultZonesMinifyValueJs = "off"
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingEditResponseResultZonesMirage struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesMirageID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesMirageValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesMirageJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesMirageJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesMirage]
type zoneSettingEditResponseResultZonesMirageJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesMirage) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesMirageID string

const (
	ZoneSettingEditResponseResultZonesMirageIDMirage ZoneSettingEditResponseResultZonesMirageID = "mirage"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesMirageEditable bool

const (
	ZoneSettingEditResponseResultZonesMirageEditableTrue  ZoneSettingEditResponseResultZonesMirageEditable = true
	ZoneSettingEditResponseResultZonesMirageEditableFalse ZoneSettingEditResponseResultZonesMirageEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesMirageValue string

const (
	ZoneSettingEditResponseResultZonesMirageValueOn  ZoneSettingEditResponseResultZonesMirageValue = "on"
	ZoneSettingEditResponseResultZonesMirageValueOff ZoneSettingEditResponseResultZonesMirageValue = "off"
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingEditResponseResultZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID ZoneSettingEditResponseResultZonesMobileRedirectID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesMobileRedirectValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesMobileRedirectJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesMobileRedirectJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesMobileRedirect]
type zoneSettingEditResponseResultZonesMobileRedirectJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesMobileRedirect) implementsZoneSettingEditResponseResult() {}

// Identifier of the zone setting.
type ZoneSettingEditResponseResultZonesMobileRedirectID string

const (
	ZoneSettingEditResponseResultZonesMobileRedirectIDMobileRedirect ZoneSettingEditResponseResultZonesMobileRedirectID = "mobile_redirect"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesMobileRedirectEditable bool

const (
	ZoneSettingEditResponseResultZonesMobileRedirectEditableTrue  ZoneSettingEditResponseResultZonesMobileRedirectEditable = true
	ZoneSettingEditResponseResultZonesMobileRedirectEditableFalse ZoneSettingEditResponseResultZonesMobileRedirectEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingEditResponseResultZonesMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                                      `json:"strip_uri"`
	JSON     zoneSettingEditResponseResultZonesMobileRedirectValueJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesMobileRedirectValueJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesMobileRedirectValue]
type zoneSettingEditResponseResultZonesMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingEditResponseResultZonesMobileRedirectValueStatus string

const (
	ZoneSettingEditResponseResultZonesMobileRedirectValueStatusOn  ZoneSettingEditResponseResultZonesMobileRedirectValueStatus = "on"
	ZoneSettingEditResponseResultZonesMobileRedirectValueStatusOff ZoneSettingEditResponseResultZonesMobileRedirectValueStatus = "off"
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingEditResponseResultZonesNel struct {
	// Zone setting identifier.
	ID ZoneSettingEditResponseResultZonesNelID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesNelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesNelValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesNelJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesNelJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesNel]
type zoneSettingEditResponseResultZonesNelJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesNel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesNel) implementsZoneSettingEditResponseResult() {}

// Zone setting identifier.
type ZoneSettingEditResponseResultZonesNelID string

const (
	ZoneSettingEditResponseResultZonesNelIDNel ZoneSettingEditResponseResultZonesNelID = "nel"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesNelEditable bool

const (
	ZoneSettingEditResponseResultZonesNelEditableTrue  ZoneSettingEditResponseResultZonesNelEditable = true
	ZoneSettingEditResponseResultZonesNelEditableFalse ZoneSettingEditResponseResultZonesNelEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesNelValue struct {
	Enabled bool                                           `json:"enabled"`
	JSON    zoneSettingEditResponseResultZonesNelValueJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesNelValueJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesNelValue]
type zoneSettingEditResponseResultZonesNelValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesNelValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingEditResponseResultZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesOpportunisticEncryptionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingEditResponseResultZonesOpportunisticEncryptionValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesOpportunisticEncryptionJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesOpportunisticEncryptionJSON contains the JSON
// metadata for the struct
// [ZoneSettingEditResponseResultZonesOpportunisticEncryption]
type zoneSettingEditResponseResultZonesOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesOpportunisticEncryption) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesOpportunisticEncryptionID string

const (
	ZoneSettingEditResponseResultZonesOpportunisticEncryptionIDOpportunisticEncryption ZoneSettingEditResponseResultZonesOpportunisticEncryptionID = "opportunistic_encryption"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesOpportunisticEncryptionEditable bool

const (
	ZoneSettingEditResponseResultZonesOpportunisticEncryptionEditableTrue  ZoneSettingEditResponseResultZonesOpportunisticEncryptionEditable = true
	ZoneSettingEditResponseResultZonesOpportunisticEncryptionEditableFalse ZoneSettingEditResponseResultZonesOpportunisticEncryptionEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditResponseResultZonesOpportunisticEncryptionValue string

const (
	ZoneSettingEditResponseResultZonesOpportunisticEncryptionValueOn  ZoneSettingEditResponseResultZonesOpportunisticEncryptionValue = "on"
	ZoneSettingEditResponseResultZonesOpportunisticEncryptionValueOff ZoneSettingEditResponseResultZonesOpportunisticEncryptionValue = "off"
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingEditResponseResultZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesOpportunisticOnionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingEditResponseResultZonesOpportunisticOnionValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesOpportunisticOnionJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesOpportunisticOnionJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesOpportunisticOnion]
type zoneSettingEditResponseResultZonesOpportunisticOnionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesOpportunisticOnion) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesOpportunisticOnionID string

const (
	ZoneSettingEditResponseResultZonesOpportunisticOnionIDOpportunisticOnion ZoneSettingEditResponseResultZonesOpportunisticOnionID = "opportunistic_onion"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesOpportunisticOnionEditable bool

const (
	ZoneSettingEditResponseResultZonesOpportunisticOnionEditableTrue  ZoneSettingEditResponseResultZonesOpportunisticOnionEditable = true
	ZoneSettingEditResponseResultZonesOpportunisticOnionEditableFalse ZoneSettingEditResponseResultZonesOpportunisticOnionEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditResponseResultZonesOpportunisticOnionValue string

const (
	ZoneSettingEditResponseResultZonesOpportunisticOnionValueOn  ZoneSettingEditResponseResultZonesOpportunisticOnionValue = "on"
	ZoneSettingEditResponseResultZonesOpportunisticOnionValueOff ZoneSettingEditResponseResultZonesOpportunisticOnionValue = "off"
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingEditResponseResultZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesOrangeToOrangeID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesOrangeToOrangeValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesOrangeToOrangeJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesOrangeToOrangeJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesOrangeToOrange]
type zoneSettingEditResponseResultZonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesOrangeToOrange) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesOrangeToOrangeID string

const (
	ZoneSettingEditResponseResultZonesOrangeToOrangeIDOrangeToOrange ZoneSettingEditResponseResultZonesOrangeToOrangeID = "orange_to_orange"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesOrangeToOrangeEditable bool

const (
	ZoneSettingEditResponseResultZonesOrangeToOrangeEditableTrue  ZoneSettingEditResponseResultZonesOrangeToOrangeEditable = true
	ZoneSettingEditResponseResultZonesOrangeToOrangeEditableFalse ZoneSettingEditResponseResultZonesOrangeToOrangeEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesOrangeToOrangeValue string

const (
	ZoneSettingEditResponseResultZonesOrangeToOrangeValueOn  ZoneSettingEditResponseResultZonesOrangeToOrangeValue = "on"
	ZoneSettingEditResponseResultZonesOrangeToOrangeValueOff ZoneSettingEditResponseResultZonesOrangeToOrangeValue = "off"
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZoneSettingEditResponseResultZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesOriginErrorPagePassThruID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesOriginErrorPagePassThruValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesOriginErrorPagePassThruJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesOriginErrorPagePassThruJSON contains the JSON
// metadata for the struct
// [ZoneSettingEditResponseResultZonesOriginErrorPagePassThru]
type zoneSettingEditResponseResultZonesOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesOriginErrorPagePassThru) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesOriginErrorPagePassThruID string

const (
	ZoneSettingEditResponseResultZonesOriginErrorPagePassThruIDOriginErrorPagePassThru ZoneSettingEditResponseResultZonesOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesOriginErrorPagePassThruEditable bool

const (
	ZoneSettingEditResponseResultZonesOriginErrorPagePassThruEditableTrue  ZoneSettingEditResponseResultZonesOriginErrorPagePassThruEditable = true
	ZoneSettingEditResponseResultZonesOriginErrorPagePassThruEditableFalse ZoneSettingEditResponseResultZonesOriginErrorPagePassThruEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesOriginErrorPagePassThruValue string

const (
	ZoneSettingEditResponseResultZonesOriginErrorPagePassThruValueOn  ZoneSettingEditResponseResultZonesOriginErrorPagePassThruValue = "on"
	ZoneSettingEditResponseResultZonesOriginErrorPagePassThruValueOff ZoneSettingEditResponseResultZonesOriginErrorPagePassThruValue = "off"
)

type ZoneSettingEditResponseResultZonesOriginMaxHTTPVersion struct {
	// Identifier of the zone setting.
	ID ZoneSettingEditResponseResultZonesOriginMaxHTTPVersionID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingEditResponseResultZonesOriginMaxHTTPVersionJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesOriginMaxHTTPVersionJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesOriginMaxHTTPVersion]
type zoneSettingEditResponseResultZonesOriginMaxHTTPVersionJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesOriginMaxHTTPVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesOriginMaxHTTPVersion) implementsZoneSettingEditResponseResult() {
}

// Identifier of the zone setting.
type ZoneSettingEditResponseResultZonesOriginMaxHTTPVersionID string

const (
	ZoneSettingEditResponseResultZonesOriginMaxHTTPVersionIDOriginMaxHTTPVersion ZoneSettingEditResponseResultZonesOriginMaxHTTPVersionID = "origin_max_http_version"
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingEditResponseResultZonesPolish struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesPolishID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesPolishValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesPolishJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesPolishJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesPolish]
type zoneSettingEditResponseResultZonesPolishJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesPolish) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesPolishID string

const (
	ZoneSettingEditResponseResultZonesPolishIDPolish ZoneSettingEditResponseResultZonesPolishID = "polish"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesPolishEditable bool

const (
	ZoneSettingEditResponseResultZonesPolishEditableTrue  ZoneSettingEditResponseResultZonesPolishEditable = true
	ZoneSettingEditResponseResultZonesPolishEditableFalse ZoneSettingEditResponseResultZonesPolishEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesPolishValue string

const (
	ZoneSettingEditResponseResultZonesPolishValueOff      ZoneSettingEditResponseResultZonesPolishValue = "off"
	ZoneSettingEditResponseResultZonesPolishValueLossless ZoneSettingEditResponseResultZonesPolishValue = "lossless"
	ZoneSettingEditResponseResultZonesPolishValueLossy    ZoneSettingEditResponseResultZonesPolishValue = "lossy"
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingEditResponseResultZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesPrefetchPreloadID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesPrefetchPreloadValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesPrefetchPreloadJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesPrefetchPreloadJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesPrefetchPreload]
type zoneSettingEditResponseResultZonesPrefetchPreloadJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesPrefetchPreload) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesPrefetchPreloadID string

const (
	ZoneSettingEditResponseResultZonesPrefetchPreloadIDPrefetchPreload ZoneSettingEditResponseResultZonesPrefetchPreloadID = "prefetch_preload"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesPrefetchPreloadEditable bool

const (
	ZoneSettingEditResponseResultZonesPrefetchPreloadEditableTrue  ZoneSettingEditResponseResultZonesPrefetchPreloadEditable = true
	ZoneSettingEditResponseResultZonesPrefetchPreloadEditableFalse ZoneSettingEditResponseResultZonesPrefetchPreloadEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesPrefetchPreloadValue string

const (
	ZoneSettingEditResponseResultZonesPrefetchPreloadValueOn  ZoneSettingEditResponseResultZonesPrefetchPreloadValue = "on"
	ZoneSettingEditResponseResultZonesPrefetchPreloadValueOff ZoneSettingEditResponseResultZonesPrefetchPreloadValue = "off"
)

// Maximum time between two read operations from origin.
type ZoneSettingEditResponseResultZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesProxyReadTimeoutID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value float64                                                `json:"value"`
	JSON  zoneSettingEditResponseResultZonesProxyReadTimeoutJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesProxyReadTimeoutJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesProxyReadTimeout]
type zoneSettingEditResponseResultZonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesProxyReadTimeout) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesProxyReadTimeoutID string

const (
	ZoneSettingEditResponseResultZonesProxyReadTimeoutIDProxyReadTimeout ZoneSettingEditResponseResultZonesProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesProxyReadTimeoutEditable bool

const (
	ZoneSettingEditResponseResultZonesProxyReadTimeoutEditableTrue  ZoneSettingEditResponseResultZonesProxyReadTimeoutEditable = true
	ZoneSettingEditResponseResultZonesProxyReadTimeoutEditableFalse ZoneSettingEditResponseResultZonesProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type ZoneSettingEditResponseResultZonesPseudoIpv4 struct {
	// Value of the Pseudo IPv4 setting.
	ID ZoneSettingEditResponseResultZonesPseudoIpv4ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesPseudoIpv4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the Pseudo IPv4 setting.
	Value ZoneSettingEditResponseResultZonesPseudoIpv4Value `json:"value"`
	JSON  zoneSettingEditResponseResultZonesPseudoIpv4JSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesPseudoIpv4JSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesPseudoIpv4]
type zoneSettingEditResponseResultZonesPseudoIpv4JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesPseudoIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesPseudoIpv4) implementsZoneSettingEditResponseResult() {}

// Value of the Pseudo IPv4 setting.
type ZoneSettingEditResponseResultZonesPseudoIpv4ID string

const (
	ZoneSettingEditResponseResultZonesPseudoIpv4IDPseudoIpv4 ZoneSettingEditResponseResultZonesPseudoIpv4ID = "pseudo_ipv4"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesPseudoIpv4Editable bool

const (
	ZoneSettingEditResponseResultZonesPseudoIpv4EditableTrue  ZoneSettingEditResponseResultZonesPseudoIpv4Editable = true
	ZoneSettingEditResponseResultZonesPseudoIpv4EditableFalse ZoneSettingEditResponseResultZonesPseudoIpv4Editable = false
)

// Value of the Pseudo IPv4 setting.
type ZoneSettingEditResponseResultZonesPseudoIpv4Value string

const (
	ZoneSettingEditResponseResultZonesPseudoIpv4ValueOff             ZoneSettingEditResponseResultZonesPseudoIpv4Value = "off"
	ZoneSettingEditResponseResultZonesPseudoIpv4ValueAddHeader       ZoneSettingEditResponseResultZonesPseudoIpv4Value = "add_header"
	ZoneSettingEditResponseResultZonesPseudoIpv4ValueOverwriteHeader ZoneSettingEditResponseResultZonesPseudoIpv4Value = "overwrite_header"
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingEditResponseResultZonesResponseBuffering struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesResponseBufferingID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesResponseBufferingValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesResponseBufferingJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesResponseBufferingJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesResponseBuffering]
type zoneSettingEditResponseResultZonesResponseBufferingJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesResponseBuffering) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesResponseBufferingID string

const (
	ZoneSettingEditResponseResultZonesResponseBufferingIDResponseBuffering ZoneSettingEditResponseResultZonesResponseBufferingID = "response_buffering"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesResponseBufferingEditable bool

const (
	ZoneSettingEditResponseResultZonesResponseBufferingEditableTrue  ZoneSettingEditResponseResultZonesResponseBufferingEditable = true
	ZoneSettingEditResponseResultZonesResponseBufferingEditableFalse ZoneSettingEditResponseResultZonesResponseBufferingEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesResponseBufferingValue string

const (
	ZoneSettingEditResponseResultZonesResponseBufferingValueOn  ZoneSettingEditResponseResultZonesResponseBufferingValue = "on"
	ZoneSettingEditResponseResultZonesResponseBufferingValueOff ZoneSettingEditResponseResultZonesResponseBufferingValue = "off"
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
type ZoneSettingEditResponseResultZonesRocketLoader struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesRocketLoaderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesRocketLoaderValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesRocketLoaderJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesRocketLoaderJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesRocketLoader]
type zoneSettingEditResponseResultZonesRocketLoaderJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesRocketLoader) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesRocketLoaderID string

const (
	ZoneSettingEditResponseResultZonesRocketLoaderIDRocketLoader ZoneSettingEditResponseResultZonesRocketLoaderID = "rocket_loader"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesRocketLoaderEditable bool

const (
	ZoneSettingEditResponseResultZonesRocketLoaderEditableTrue  ZoneSettingEditResponseResultZonesRocketLoaderEditable = true
	ZoneSettingEditResponseResultZonesRocketLoaderEditableFalse ZoneSettingEditResponseResultZonesRocketLoaderEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesRocketLoaderValue string

const (
	ZoneSettingEditResponseResultZonesRocketLoaderValueOn  ZoneSettingEditResponseResultZonesRocketLoaderValue = "on"
	ZoneSettingEditResponseResultZonesRocketLoaderValueOff ZoneSettingEditResponseResultZonesRocketLoaderValue = "off"
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                                   `json:"modified_on,nullable" format:"date-time"`
	Value      ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationValue `json:"value"`
	JSON       zoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationJSON
// contains the JSON metadata for the struct
// [ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimization]
type zoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimization) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationID string

const (
	ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationEditable = false
)

type ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationValue struct {
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
	WpPlugin bool                                                                            `json:"wp_plugin,required"`
	JSON     zoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationValueJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationValueJSON
// contains the JSON metadata for the struct
// [ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationValue]
type zoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationValueJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSchemasAutomaticPlatformOptimizationValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare security header for a zone.
type ZoneSettingEditResponseResultZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID ZoneSettingEditResponseResultZonesSecurityHeaderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesSecurityHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                             `json:"modified_on,nullable" format:"date-time"`
	Value      ZoneSettingEditResponseResultZonesSecurityHeaderValue `json:"value"`
	JSON       zoneSettingEditResponseResultZonesSecurityHeaderJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesSecurityHeaderJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesSecurityHeader]
type zoneSettingEditResponseResultZonesSecurityHeaderJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesSecurityHeader) implementsZoneSettingEditResponseResult() {}

// ID of the zone's security header.
type ZoneSettingEditResponseResultZonesSecurityHeaderID string

const (
	ZoneSettingEditResponseResultZonesSecurityHeaderIDSecurityHeader ZoneSettingEditResponseResultZonesSecurityHeaderID = "security_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesSecurityHeaderEditable bool

const (
	ZoneSettingEditResponseResultZonesSecurityHeaderEditableTrue  ZoneSettingEditResponseResultZonesSecurityHeaderEditable = true
	ZoneSettingEditResponseResultZonesSecurityHeaderEditableFalse ZoneSettingEditResponseResultZonesSecurityHeaderEditable = false
)

type ZoneSettingEditResponseResultZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZoneSettingEditResponseResultZonesSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zoneSettingEditResponseResultZonesSecurityHeaderValueJSON                    `json:"-"`
}

// zoneSettingEditResponseResultZonesSecurityHeaderValueJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesSecurityHeaderValue]
type zoneSettingEditResponseResultZonesSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type ZoneSettingEditResponseResultZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                             `json:"nosniff"`
	JSON    zoneSettingEditResponseResultZonesSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesSecurityHeaderValueStrictTransportSecurityJSON
// contains the JSON metadata for the struct
// [ZoneSettingEditResponseResultZonesSecurityHeaderValueStrictTransportSecurity]
type zoneSettingEditResponseResultZonesSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingEditResponseResultZonesSecurityLevel struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesSecurityLevelID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesSecurityLevelValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesSecurityLevelJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesSecurityLevelJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesSecurityLevel]
type zoneSettingEditResponseResultZonesSecurityLevelJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesSecurityLevel) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesSecurityLevelID string

const (
	ZoneSettingEditResponseResultZonesSecurityLevelIDSecurityLevel ZoneSettingEditResponseResultZonesSecurityLevelID = "security_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesSecurityLevelEditable bool

const (
	ZoneSettingEditResponseResultZonesSecurityLevelEditableTrue  ZoneSettingEditResponseResultZonesSecurityLevelEditable = true
	ZoneSettingEditResponseResultZonesSecurityLevelEditableFalse ZoneSettingEditResponseResultZonesSecurityLevelEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesSecurityLevelValue string

const (
	ZoneSettingEditResponseResultZonesSecurityLevelValueOff            ZoneSettingEditResponseResultZonesSecurityLevelValue = "off"
	ZoneSettingEditResponseResultZonesSecurityLevelValueEssentiallyOff ZoneSettingEditResponseResultZonesSecurityLevelValue = "essentially_off"
	ZoneSettingEditResponseResultZonesSecurityLevelValueLow            ZoneSettingEditResponseResultZonesSecurityLevelValue = "low"
	ZoneSettingEditResponseResultZonesSecurityLevelValueMedium         ZoneSettingEditResponseResultZonesSecurityLevelValue = "medium"
	ZoneSettingEditResponseResultZonesSecurityLevelValueHigh           ZoneSettingEditResponseResultZonesSecurityLevelValue = "high"
	ZoneSettingEditResponseResultZonesSecurityLevelValueUnderAttack    ZoneSettingEditResponseResultZonesSecurityLevelValue = "under_attack"
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
type ZoneSettingEditResponseResultZonesServerSideExclude struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesServerSideExcludeID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesServerSideExcludeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesServerSideExcludeValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesServerSideExcludeJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesServerSideExcludeJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesServerSideExclude]
type zoneSettingEditResponseResultZonesServerSideExcludeJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesServerSideExclude) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesServerSideExcludeID string

const (
	ZoneSettingEditResponseResultZonesServerSideExcludeIDServerSideExclude ZoneSettingEditResponseResultZonesServerSideExcludeID = "server_side_exclude"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesServerSideExcludeEditable bool

const (
	ZoneSettingEditResponseResultZonesServerSideExcludeEditableTrue  ZoneSettingEditResponseResultZonesServerSideExcludeEditable = true
	ZoneSettingEditResponseResultZonesServerSideExcludeEditableFalse ZoneSettingEditResponseResultZonesServerSideExcludeEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesServerSideExcludeValue string

const (
	ZoneSettingEditResponseResultZonesServerSideExcludeValueOn  ZoneSettingEditResponseResultZonesServerSideExcludeValue = "on"
	ZoneSettingEditResponseResultZonesServerSideExcludeValueOff ZoneSettingEditResponseResultZonesServerSideExcludeValue = "off"
)

// Allow SHA1 support.
type ZoneSettingEditResponseResultZonesSha1Support struct {
	// Zone setting identifier.
	ID ZoneSettingEditResponseResultZonesSha1SupportID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesSha1SupportValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesSha1SupportJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesSha1SupportJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesSha1Support]
type zoneSettingEditResponseResultZonesSha1SupportJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSha1Support) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesSha1Support) implementsZoneSettingEditResponseResult() {}

// Zone setting identifier.
type ZoneSettingEditResponseResultZonesSha1SupportID string

const (
	ZoneSettingEditResponseResultZonesSha1SupportIDSha1Support ZoneSettingEditResponseResultZonesSha1SupportID = "sha1_support"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesSha1SupportEditable bool

const (
	ZoneSettingEditResponseResultZonesSha1SupportEditableTrue  ZoneSettingEditResponseResultZonesSha1SupportEditable = true
	ZoneSettingEditResponseResultZonesSha1SupportEditableFalse ZoneSettingEditResponseResultZonesSha1SupportEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesSha1SupportValue string

const (
	ZoneSettingEditResponseResultZonesSha1SupportValueOff ZoneSettingEditResponseResultZonesSha1SupportValue = "off"
	ZoneSettingEditResponseResultZonesSha1SupportValueOn  ZoneSettingEditResponseResultZonesSha1SupportValue = "on"
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingEditResponseResultZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesSortQueryStringForCacheID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesSortQueryStringForCacheValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesSortQueryStringForCacheJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesSortQueryStringForCacheJSON contains the JSON
// metadata for the struct
// [ZoneSettingEditResponseResultZonesSortQueryStringForCache]
type zoneSettingEditResponseResultZonesSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesSortQueryStringForCache) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesSortQueryStringForCacheID string

const (
	ZoneSettingEditResponseResultZonesSortQueryStringForCacheIDSortQueryStringForCache ZoneSettingEditResponseResultZonesSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesSortQueryStringForCacheEditable bool

const (
	ZoneSettingEditResponseResultZonesSortQueryStringForCacheEditableTrue  ZoneSettingEditResponseResultZonesSortQueryStringForCacheEditable = true
	ZoneSettingEditResponseResultZonesSortQueryStringForCacheEditableFalse ZoneSettingEditResponseResultZonesSortQueryStringForCacheEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesSortQueryStringForCacheValue string

const (
	ZoneSettingEditResponseResultZonesSortQueryStringForCacheValueOn  ZoneSettingEditResponseResultZonesSortQueryStringForCacheValue = "on"
	ZoneSettingEditResponseResultZonesSortQueryStringForCacheValueOff ZoneSettingEditResponseResultZonesSortQueryStringForCacheValue = "off"
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
type ZoneSettingEditResponseResultZonesSsl struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesSslID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesSslEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value ZoneSettingEditResponseResultZonesSslValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesSslJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesSslJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesSsl]
type zoneSettingEditResponseResultZonesSslJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesSsl) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesSslID string

const (
	ZoneSettingEditResponseResultZonesSslIDSsl ZoneSettingEditResponseResultZonesSslID = "ssl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesSslEditable bool

const (
	ZoneSettingEditResponseResultZonesSslEditableTrue  ZoneSettingEditResponseResultZonesSslEditable = true
	ZoneSettingEditResponseResultZonesSslEditableFalse ZoneSettingEditResponseResultZonesSslEditable = false
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingEditResponseResultZonesSslValue string

const (
	ZoneSettingEditResponseResultZonesSslValueOff      ZoneSettingEditResponseResultZonesSslValue = "off"
	ZoneSettingEditResponseResultZonesSslValueFlexible ZoneSettingEditResponseResultZonesSslValue = "flexible"
	ZoneSettingEditResponseResultZonesSslValueFull     ZoneSettingEditResponseResultZonesSslValue = "full"
	ZoneSettingEditResponseResultZonesSslValueStrict   ZoneSettingEditResponseResultZonesSslValue = "strict"
)

type ZoneSettingEditResponseResultZonesSslRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingEditResponseResultZonesSslRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                                 `json:"enabled"`
	JSON    zoneSettingEditResponseResultZonesSslRecommenderJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesSslRecommenderJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesSslRecommender]
type zoneSettingEditResponseResultZonesSslRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSslRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesSslRecommender) implementsZoneSettingEditResponseResult() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingEditResponseResultZonesSslRecommenderID string

const (
	ZoneSettingEditResponseResultZonesSslRecommenderIDSslRecommender ZoneSettingEditResponseResultZonesSslRecommenderID = "ssl_recommender"
)

// Only allows TLS1.2.
type ZoneSettingEditResponseResultZonesTls1_2Only struct {
	// Zone setting identifier.
	ID ZoneSettingEditResponseResultZonesTls1_2OnlyID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesTls1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesTls1_2OnlyValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesTls1_2OnlyJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesTls1_2OnlyJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesTls1_2Only]
type zoneSettingEditResponseResultZonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesTls1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesTls1_2Only) implementsZoneSettingEditResponseResult() {}

// Zone setting identifier.
type ZoneSettingEditResponseResultZonesTls1_2OnlyID string

const (
	ZoneSettingEditResponseResultZonesTls1_2OnlyIDTls1_2Only ZoneSettingEditResponseResultZonesTls1_2OnlyID = "tls_1_2_only"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesTls1_2OnlyEditable bool

const (
	ZoneSettingEditResponseResultZonesTls1_2OnlyEditableTrue  ZoneSettingEditResponseResultZonesTls1_2OnlyEditable = true
	ZoneSettingEditResponseResultZonesTls1_2OnlyEditableFalse ZoneSettingEditResponseResultZonesTls1_2OnlyEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesTls1_2OnlyValue string

const (
	ZoneSettingEditResponseResultZonesTls1_2OnlyValueOff ZoneSettingEditResponseResultZonesTls1_2OnlyValue = "off"
	ZoneSettingEditResponseResultZonesTls1_2OnlyValueOn  ZoneSettingEditResponseResultZonesTls1_2OnlyValue = "on"
)

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingEditResponseResultZonesTls1_3 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesTls1_3ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesTls1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingEditResponseResultZonesTls1_3Value `json:"value"`
	JSON  zoneSettingEditResponseResultZonesTls1_3JSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesTls1_3JSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesTls1_3]
type zoneSettingEditResponseResultZonesTls1_3JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesTls1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesTls1_3) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesTls1_3ID string

const (
	ZoneSettingEditResponseResultZonesTls1_3IDTls1_3 ZoneSettingEditResponseResultZonesTls1_3ID = "tls_1_3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesTls1_3Editable bool

const (
	ZoneSettingEditResponseResultZonesTls1_3EditableTrue  ZoneSettingEditResponseResultZonesTls1_3Editable = true
	ZoneSettingEditResponseResultZonesTls1_3EditableFalse ZoneSettingEditResponseResultZonesTls1_3Editable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditResponseResultZonesTls1_3Value string

const (
	ZoneSettingEditResponseResultZonesTls1_3ValueOn  ZoneSettingEditResponseResultZonesTls1_3Value = "on"
	ZoneSettingEditResponseResultZonesTls1_3ValueOff ZoneSettingEditResponseResultZonesTls1_3Value = "off"
	ZoneSettingEditResponseResultZonesTls1_3ValueZrt ZoneSettingEditResponseResultZonesTls1_3Value = "zrt"
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingEditResponseResultZonesTlsClientAuth struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesTlsClientAuthID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesTlsClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value ZoneSettingEditResponseResultZonesTlsClientAuthValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesTlsClientAuthJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesTlsClientAuthJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesTlsClientAuth]
type zoneSettingEditResponseResultZonesTlsClientAuthJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesTlsClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesTlsClientAuth) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesTlsClientAuthID string

const (
	ZoneSettingEditResponseResultZonesTlsClientAuthIDTlsClientAuth ZoneSettingEditResponseResultZonesTlsClientAuthID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesTlsClientAuthEditable bool

const (
	ZoneSettingEditResponseResultZonesTlsClientAuthEditableTrue  ZoneSettingEditResponseResultZonesTlsClientAuthEditable = true
	ZoneSettingEditResponseResultZonesTlsClientAuthEditableFalse ZoneSettingEditResponseResultZonesTlsClientAuthEditable = false
)

// value of the zone setting.
type ZoneSettingEditResponseResultZonesTlsClientAuthValue string

const (
	ZoneSettingEditResponseResultZonesTlsClientAuthValueOn  ZoneSettingEditResponseResultZonesTlsClientAuthValue = "on"
	ZoneSettingEditResponseResultZonesTlsClientAuthValueOff ZoneSettingEditResponseResultZonesTlsClientAuthValue = "off"
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingEditResponseResultZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesTrueClientIPHeaderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesTrueClientIPHeaderValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesTrueClientIPHeaderJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesTrueClientIPHeaderJSON contains the JSON
// metadata for the struct [ZoneSettingEditResponseResultZonesTrueClientIPHeader]
type zoneSettingEditResponseResultZonesTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesTrueClientIPHeader) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesTrueClientIPHeaderID string

const (
	ZoneSettingEditResponseResultZonesTrueClientIPHeaderIDTrueClientIPHeader ZoneSettingEditResponseResultZonesTrueClientIPHeaderID = "true_client_ip_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesTrueClientIPHeaderEditable bool

const (
	ZoneSettingEditResponseResultZonesTrueClientIPHeaderEditableTrue  ZoneSettingEditResponseResultZonesTrueClientIPHeaderEditable = true
	ZoneSettingEditResponseResultZonesTrueClientIPHeaderEditableFalse ZoneSettingEditResponseResultZonesTrueClientIPHeaderEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesTrueClientIPHeaderValue string

const (
	ZoneSettingEditResponseResultZonesTrueClientIPHeaderValueOn  ZoneSettingEditResponseResultZonesTrueClientIPHeaderValue = "on"
	ZoneSettingEditResponseResultZonesTrueClientIPHeaderValueOff ZoneSettingEditResponseResultZonesTrueClientIPHeaderValue = "off"
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
type ZoneSettingEditResponseResultZonesWaf struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesWafID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesWafEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesWafValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesWafJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesWafJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesWaf]
type zoneSettingEditResponseResultZonesWafJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesWaf) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesWafID string

const (
	ZoneSettingEditResponseResultZonesWafIDWaf ZoneSettingEditResponseResultZonesWafID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesWafEditable bool

const (
	ZoneSettingEditResponseResultZonesWafEditableTrue  ZoneSettingEditResponseResultZonesWafEditable = true
	ZoneSettingEditResponseResultZonesWafEditableFalse ZoneSettingEditResponseResultZonesWafEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesWafValue string

const (
	ZoneSettingEditResponseResultZonesWafValueOn  ZoneSettingEditResponseResultZonesWafValue = "on"
	ZoneSettingEditResponseResultZonesWafValueOff ZoneSettingEditResponseResultZonesWafValue = "off"
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingEditResponseResultZonesWebp struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesWebpID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesWebpValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesWebpJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesWebpJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesWebp]
type zoneSettingEditResponseResultZonesWebpJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesWebp) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesWebpID string

const (
	ZoneSettingEditResponseResultZonesWebpIDWebp ZoneSettingEditResponseResultZonesWebpID = "webp"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesWebpEditable bool

const (
	ZoneSettingEditResponseResultZonesWebpEditableTrue  ZoneSettingEditResponseResultZonesWebpEditable = true
	ZoneSettingEditResponseResultZonesWebpEditableFalse ZoneSettingEditResponseResultZonesWebpEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesWebpValue string

const (
	ZoneSettingEditResponseResultZonesWebpValueOff ZoneSettingEditResponseResultZonesWebpValue = "off"
	ZoneSettingEditResponseResultZonesWebpValueOn  ZoneSettingEditResponseResultZonesWebpValue = "on"
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZoneSettingEditResponseResultZonesWebsockets struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesWebsocketsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesWebsocketsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesWebsocketsValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesWebsocketsJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesWebsocketsJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesWebsockets]
type zoneSettingEditResponseResultZonesWebsocketsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesWebsockets) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesWebsocketsID string

const (
	ZoneSettingEditResponseResultZonesWebsocketsIDWebsockets ZoneSettingEditResponseResultZonesWebsocketsID = "websockets"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesWebsocketsEditable bool

const (
	ZoneSettingEditResponseResultZonesWebsocketsEditableTrue  ZoneSettingEditResponseResultZonesWebsocketsEditable = true
	ZoneSettingEditResponseResultZonesWebsocketsEditableFalse ZoneSettingEditResponseResultZonesWebsocketsEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesWebsocketsValue string

const (
	ZoneSettingEditResponseResultZonesWebsocketsValueOff ZoneSettingEditResponseResultZonesWebsocketsValue = "off"
	ZoneSettingEditResponseResultZonesWebsocketsValueOn  ZoneSettingEditResponseResultZonesWebsocketsValue = "on"
)

type ZoneSettingEditParams struct {
	// One or more zone setting objects. Must contain an ID and a value.
	Items param.Field[[]ZoneSettingEditParamsItem] `json:"items,required"`
}

func (r ZoneSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// 0-RTT session resumption enabled for this zone.
//
// Satisfied by [ZoneSettingEditParamsItemsZones0rtt],
// [ZoneSettingEditParamsItemsZonesAdvancedDdos],
// [ZoneSettingEditParamsItemsZonesAlwaysOnline],
// [ZoneSettingEditParamsItemsZonesAlwaysUseHTTPs],
// [ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewrites],
// [ZoneSettingEditParamsItemsZonesBrotli],
// [ZoneSettingEditParamsItemsZonesBrowserCacheTtl],
// [ZoneSettingEditParamsItemsZonesBrowserCheck],
// [ZoneSettingEditParamsItemsZonesCacheLevel],
// [ZoneSettingEditParamsItemsZonesChallengeTtl],
// [ZoneSettingEditParamsItemsZonesCiphers],
// [ZoneSettingEditParamsItemsZonesCnameFlattening],
// [ZoneSettingEditParamsItemsZonesDevelopmentMode],
// [ZoneSettingEditParamsItemsZonesEarlyHints],
// [ZoneSettingEditParamsItemsZonesEdgeCacheTtl],
// [ZoneSettingEditParamsItemsZonesEmailObfuscation],
// [ZoneSettingEditParamsItemsZonesH2Prioritization],
// [ZoneSettingEditParamsItemsZonesHotlinkProtection],
// [ZoneSettingEditParamsItemsZonesHttp2], [ZoneSettingEditParamsItemsZonesHttp3],
// [ZoneSettingEditParamsItemsZonesImageResizing],
// [ZoneSettingEditParamsItemsZonesIPGeolocation],
// [ZoneSettingEditParamsItemsZonesIpv6],
// [ZoneSettingEditParamsItemsZonesMaxUpload],
// [ZoneSettingEditParamsItemsZonesMinTlsVersion],
// [ZoneSettingEditParamsItemsZonesMinify],
// [ZoneSettingEditParamsItemsZonesMirage],
// [ZoneSettingEditParamsItemsZonesMobileRedirect],
// [ZoneSettingEditParamsItemsZonesNel],
// [ZoneSettingEditParamsItemsZonesOpportunisticEncryption],
// [ZoneSettingEditParamsItemsZonesOpportunisticOnion],
// [ZoneSettingEditParamsItemsZonesOrangeToOrange],
// [ZoneSettingEditParamsItemsZonesOriginErrorPagePassThru],
// [ZoneSettingEditParamsItemsZonesOriginMaxHTTPVersion],
// [ZoneSettingEditParamsItemsZonesPolish],
// [ZoneSettingEditParamsItemsZonesPrefetchPreload],
// [ZoneSettingEditParamsItemsZonesProxyReadTimeout],
// [ZoneSettingEditParamsItemsZonesPseudoIpv4],
// [ZoneSettingEditParamsItemsZonesResponseBuffering],
// [ZoneSettingEditParamsItemsZonesRocketLoader],
// [ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization],
// [ZoneSettingEditParamsItemsZonesSecurityHeader],
// [ZoneSettingEditParamsItemsZonesSecurityLevel],
// [ZoneSettingEditParamsItemsZonesServerSideExclude],
// [ZoneSettingEditParamsItemsZonesSha1Support],
// [ZoneSettingEditParamsItemsZonesSortQueryStringForCache],
// [ZoneSettingEditParamsItemsZonesSsl],
// [ZoneSettingEditParamsItemsZonesSslRecommender],
// [ZoneSettingEditParamsItemsZonesTls1_2Only],
// [ZoneSettingEditParamsItemsZonesTls1_3],
// [ZoneSettingEditParamsItemsZonesTlsClientAuth],
// [ZoneSettingEditParamsItemsZonesTrueClientIPHeader],
// [ZoneSettingEditParamsItemsZonesWaf], [ZoneSettingEditParamsItemsZonesWebp],
// [ZoneSettingEditParamsItemsZonesWebsockets].
type ZoneSettingEditParamsItem interface {
	implementsZoneSettingEditParamsItem()
}

// 0-RTT session resumption enabled for this zone.
type ZoneSettingEditParamsItemsZones0rtt struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZones0rttID] `json:"id"`
	// Value of the 0-RTT setting.
	Value param.Field[ZoneSettingEditParamsItemsZones0rttValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZones0rttEditable bool

const (
	ZoneSettingEditParamsItemsZones0rttEditableTrue  ZoneSettingEditParamsItemsZones0rttEditable = true
	ZoneSettingEditParamsItemsZones0rttEditableFalse ZoneSettingEditParamsItemsZones0rttEditable = false
)

// Value of the 0-RTT setting.
type ZoneSettingEditParamsItemsZones0rttValue string

const (
	ZoneSettingEditParamsItemsZones0rttValueOn  ZoneSettingEditParamsItemsZones0rttValue = "on"
	ZoneSettingEditParamsItemsZones0rttValueOff ZoneSettingEditParamsItemsZones0rttValue = "off"
)

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingEditParamsItemsZonesAdvancedDdos struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesAdvancedDdosID] `json:"id"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value param.Field[ZoneSettingEditParamsItemsZonesAdvancedDdosValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesAdvancedDdos) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesAdvancedDdos) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesAdvancedDdosID string

const (
	ZoneSettingEditParamsItemsZonesAdvancedDdosIDAdvancedDdos ZoneSettingEditParamsItemsZonesAdvancedDdosID = "advanced_ddos"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesAdvancedDdosEditable bool

const (
	ZoneSettingEditParamsItemsZonesAdvancedDdosEditableTrue  ZoneSettingEditParamsItemsZonesAdvancedDdosEditable = true
	ZoneSettingEditParamsItemsZonesAdvancedDdosEditableFalse ZoneSettingEditParamsItemsZonesAdvancedDdosEditable = false
)

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZoneSettingEditParamsItemsZonesAdvancedDdosValue string

const (
	ZoneSettingEditParamsItemsZonesAdvancedDdosValueOn  ZoneSettingEditParamsItemsZonesAdvancedDdosValue = "on"
	ZoneSettingEditParamsItemsZonesAdvancedDdosValueOff ZoneSettingEditParamsItemsZonesAdvancedDdosValue = "off"
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingEditParamsItemsZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesAlwaysOnlineID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesAlwaysOnlineValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesAlwaysOnlineEditable bool

const (
	ZoneSettingEditParamsItemsZonesAlwaysOnlineEditableTrue  ZoneSettingEditParamsItemsZonesAlwaysOnlineEditable = true
	ZoneSettingEditParamsItemsZonesAlwaysOnlineEditableFalse ZoneSettingEditParamsItemsZonesAlwaysOnlineEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesAlwaysOnlineValue string

const (
	ZoneSettingEditParamsItemsZonesAlwaysOnlineValueOn  ZoneSettingEditParamsItemsZonesAlwaysOnlineValue = "on"
	ZoneSettingEditParamsItemsZonesAlwaysOnlineValueOff ZoneSettingEditParamsItemsZonesAlwaysOnlineValue = "off"
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingEditParamsItemsZonesAlwaysUseHTTPs struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesAlwaysUseHTTPs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesAlwaysUseHTTPs) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsID string

const (
	ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsIDAlwaysUseHTTPs ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsID = "always_use_https"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsEditable bool

const (
	ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsEditableTrue  ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsEditable = true
	ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsEditableFalse ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsValue string

const (
	ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsValueOn  ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsValue = "on"
	ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsValueOff ZoneSettingEditParamsItemsZonesAlwaysUseHTTPsValue = "off"
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewrites struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewrites) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewrites) implementsZoneSettingEditParamsItem() {
}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesID string

const (
	ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesIDAutomaticHTTPsRewrites ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesID = "automatic_https_rewrites"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesEditable bool

const (
	ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesEditableTrue  ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesEditable = true
	ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesEditableFalse ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesValue string

const (
	ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesValueOn  ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesValue = "on"
	ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesValueOff ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewritesValue = "off"
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingEditParamsItemsZonesBrotli struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesBrotliID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesBrotliValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesBrotliEditable bool

const (
	ZoneSettingEditParamsItemsZonesBrotliEditableTrue  ZoneSettingEditParamsItemsZonesBrotliEditable = true
	ZoneSettingEditParamsItemsZonesBrotliEditableFalse ZoneSettingEditParamsItemsZonesBrotliEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesBrotliValue string

const (
	ZoneSettingEditParamsItemsZonesBrotliValueOff ZoneSettingEditParamsItemsZonesBrotliValue = "off"
	ZoneSettingEditParamsItemsZonesBrotliValueOn  ZoneSettingEditParamsItemsZonesBrotliValue = "on"
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingEditParamsItemsZonesBrowserCacheTtl struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesBrowserCacheTtlID] `json:"id"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesBrowserCacheTtl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesBrowserCacheTtl) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesBrowserCacheTtlID string

const (
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlIDBrowserCacheTtl ZoneSettingEditParamsItemsZonesBrowserCacheTtlID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesBrowserCacheTtlEditable bool

const (
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlEditableTrue  ZoneSettingEditParamsItemsZonesBrowserCacheTtlEditable = true
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlEditableFalse ZoneSettingEditParamsItemsZonesBrowserCacheTtlEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue float64

const (
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue0        ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 0
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue30       ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 30
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue60       ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 60
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue120      ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 120
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue300      ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 300
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue1200     ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 1200
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue1800     ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 1800
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue3600     ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 3600
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue7200     ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 7200
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue10800    ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 10800
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue14400    ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 14400
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue18000    ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 18000
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue28800    ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 28800
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue43200    ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 43200
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue57600    ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 57600
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue72000    ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 72000
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue86400    ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 86400
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue172800   ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 172800
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue259200   ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 259200
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue345600   ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 345600
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue432000   ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 432000
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue691200   ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 691200
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue1382400  ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 1382400
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue2073600  ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 2073600
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue2678400  ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 2678400
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue5356800  ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 5356800
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue16070400 ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 16070400
	ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue31536000 ZoneSettingEditParamsItemsZonesBrowserCacheTtlValue = 31536000
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingEditParamsItemsZonesBrowserCheck struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesBrowserCheckID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesBrowserCheckValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesBrowserCheckEditable bool

const (
	ZoneSettingEditParamsItemsZonesBrowserCheckEditableTrue  ZoneSettingEditParamsItemsZonesBrowserCheckEditable = true
	ZoneSettingEditParamsItemsZonesBrowserCheckEditableFalse ZoneSettingEditParamsItemsZonesBrowserCheckEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesBrowserCheckValue string

const (
	ZoneSettingEditParamsItemsZonesBrowserCheckValueOn  ZoneSettingEditParamsItemsZonesBrowserCheckValue = "on"
	ZoneSettingEditParamsItemsZonesBrowserCheckValueOff ZoneSettingEditParamsItemsZonesBrowserCheckValue = "off"
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingEditParamsItemsZonesCacheLevel struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesCacheLevelID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesCacheLevelValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesCacheLevelEditable bool

const (
	ZoneSettingEditParamsItemsZonesCacheLevelEditableTrue  ZoneSettingEditParamsItemsZonesCacheLevelEditable = true
	ZoneSettingEditParamsItemsZonesCacheLevelEditableFalse ZoneSettingEditParamsItemsZonesCacheLevelEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesCacheLevelValue string

const (
	ZoneSettingEditParamsItemsZonesCacheLevelValueAggressive ZoneSettingEditParamsItemsZonesCacheLevelValue = "aggressive"
	ZoneSettingEditParamsItemsZonesCacheLevelValueBasic      ZoneSettingEditParamsItemsZonesCacheLevelValue = "basic"
	ZoneSettingEditParamsItemsZonesCacheLevelValueSimplified ZoneSettingEditParamsItemsZonesCacheLevelValue = "simplified"
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingEditParamsItemsZonesChallengeTtl struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesChallengeTtlID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesChallengeTtlValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesChallengeTtl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesChallengeTtl) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesChallengeTtlID string

const (
	ZoneSettingEditParamsItemsZonesChallengeTtlIDChallengeTtl ZoneSettingEditParamsItemsZonesChallengeTtlID = "challenge_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesChallengeTtlEditable bool

const (
	ZoneSettingEditParamsItemsZonesChallengeTtlEditableTrue  ZoneSettingEditParamsItemsZonesChallengeTtlEditable = true
	ZoneSettingEditParamsItemsZonesChallengeTtlEditableFalse ZoneSettingEditParamsItemsZonesChallengeTtlEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesChallengeTtlValue float64

const (
	ZoneSettingEditParamsItemsZonesChallengeTtlValue300      ZoneSettingEditParamsItemsZonesChallengeTtlValue = 300
	ZoneSettingEditParamsItemsZonesChallengeTtlValue900      ZoneSettingEditParamsItemsZonesChallengeTtlValue = 900
	ZoneSettingEditParamsItemsZonesChallengeTtlValue1800     ZoneSettingEditParamsItemsZonesChallengeTtlValue = 1800
	ZoneSettingEditParamsItemsZonesChallengeTtlValue2700     ZoneSettingEditParamsItemsZonesChallengeTtlValue = 2700
	ZoneSettingEditParamsItemsZonesChallengeTtlValue3600     ZoneSettingEditParamsItemsZonesChallengeTtlValue = 3600
	ZoneSettingEditParamsItemsZonesChallengeTtlValue7200     ZoneSettingEditParamsItemsZonesChallengeTtlValue = 7200
	ZoneSettingEditParamsItemsZonesChallengeTtlValue10800    ZoneSettingEditParamsItemsZonesChallengeTtlValue = 10800
	ZoneSettingEditParamsItemsZonesChallengeTtlValue14400    ZoneSettingEditParamsItemsZonesChallengeTtlValue = 14400
	ZoneSettingEditParamsItemsZonesChallengeTtlValue28800    ZoneSettingEditParamsItemsZonesChallengeTtlValue = 28800
	ZoneSettingEditParamsItemsZonesChallengeTtlValue57600    ZoneSettingEditParamsItemsZonesChallengeTtlValue = 57600
	ZoneSettingEditParamsItemsZonesChallengeTtlValue86400    ZoneSettingEditParamsItemsZonesChallengeTtlValue = 86400
	ZoneSettingEditParamsItemsZonesChallengeTtlValue604800   ZoneSettingEditParamsItemsZonesChallengeTtlValue = 604800
	ZoneSettingEditParamsItemsZonesChallengeTtlValue2592000  ZoneSettingEditParamsItemsZonesChallengeTtlValue = 2592000
	ZoneSettingEditParamsItemsZonesChallengeTtlValue31536000 ZoneSettingEditParamsItemsZonesChallengeTtlValue = 31536000
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingEditParamsItemsZonesCiphers struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesCiphersID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[[]string] `json:"value"`
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
	ID param.Field[ZoneSettingEditParamsItemsZonesCnameFlatteningID] `json:"id"`
	// Value of the cname flattening setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesCnameFlatteningValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesCnameFlatteningEditable bool

const (
	ZoneSettingEditParamsItemsZonesCnameFlatteningEditableTrue  ZoneSettingEditParamsItemsZonesCnameFlatteningEditable = true
	ZoneSettingEditParamsItemsZonesCnameFlatteningEditableFalse ZoneSettingEditParamsItemsZonesCnameFlatteningEditable = false
)

// Value of the cname flattening setting.
type ZoneSettingEditParamsItemsZonesCnameFlatteningValue string

const (
	ZoneSettingEditParamsItemsZonesCnameFlatteningValueFlattenAtRoot ZoneSettingEditParamsItemsZonesCnameFlatteningValue = "flatten_at_root"
	ZoneSettingEditParamsItemsZonesCnameFlatteningValueFlattenAll    ZoneSettingEditParamsItemsZonesCnameFlatteningValue = "flatten_all"
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingEditParamsItemsZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesDevelopmentModeID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesDevelopmentModeValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesDevelopmentModeEditable bool

const (
	ZoneSettingEditParamsItemsZonesDevelopmentModeEditableTrue  ZoneSettingEditParamsItemsZonesDevelopmentModeEditable = true
	ZoneSettingEditParamsItemsZonesDevelopmentModeEditableFalse ZoneSettingEditParamsItemsZonesDevelopmentModeEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesDevelopmentModeValue string

const (
	ZoneSettingEditParamsItemsZonesDevelopmentModeValueOn  ZoneSettingEditParamsItemsZonesDevelopmentModeValue = "on"
	ZoneSettingEditParamsItemsZonesDevelopmentModeValueOff ZoneSettingEditParamsItemsZonesDevelopmentModeValue = "off"
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingEditParamsItemsZonesEarlyHints struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesEarlyHintsID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesEarlyHintsValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesEarlyHintsEditable bool

const (
	ZoneSettingEditParamsItemsZonesEarlyHintsEditableTrue  ZoneSettingEditParamsItemsZonesEarlyHintsEditable = true
	ZoneSettingEditParamsItemsZonesEarlyHintsEditableFalse ZoneSettingEditParamsItemsZonesEarlyHintsEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesEarlyHintsValue string

const (
	ZoneSettingEditParamsItemsZonesEarlyHintsValueOn  ZoneSettingEditParamsItemsZonesEarlyHintsValue = "on"
	ZoneSettingEditParamsItemsZonesEarlyHintsValueOff ZoneSettingEditParamsItemsZonesEarlyHintsValue = "off"
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZoneSettingEditParamsItemsZonesEdgeCacheTtl struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesEdgeCacheTtlID] `json:"id"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value param.Field[ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesEdgeCacheTtl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesEdgeCacheTtl) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesEdgeCacheTtlID string

const (
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlIDEdgeCacheTtl ZoneSettingEditParamsItemsZonesEdgeCacheTtlID = "edge_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesEdgeCacheTtlEditable bool

const (
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlEditableTrue  ZoneSettingEditParamsItemsZonesEdgeCacheTtlEditable = true
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlEditableFalse ZoneSettingEditParamsItemsZonesEdgeCacheTtlEditable = false
)

// Value of the zone setting. Notes: The minimum TTL available depends on the plan
// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
type ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue float64

const (
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue30     ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 30
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue60     ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 60
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue300    ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 300
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue1200   ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 1200
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue1800   ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 1800
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue3600   ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 3600
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue7200   ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 7200
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue10800  ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 10800
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue14400  ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 14400
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue18000  ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 18000
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue28800  ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 28800
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue43200  ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 43200
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue57600  ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 57600
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue72000  ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 72000
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue86400  ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 86400
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue172800 ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 172800
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue259200 ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 259200
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue345600 ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 345600
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue432000 ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 432000
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue518400 ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 518400
	ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue604800 ZoneSettingEditParamsItemsZonesEdgeCacheTtlValue = 604800
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingEditParamsItemsZonesEmailObfuscation struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesEmailObfuscationID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesEmailObfuscationValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesEmailObfuscationEditable bool

const (
	ZoneSettingEditParamsItemsZonesEmailObfuscationEditableTrue  ZoneSettingEditParamsItemsZonesEmailObfuscationEditable = true
	ZoneSettingEditParamsItemsZonesEmailObfuscationEditableFalse ZoneSettingEditParamsItemsZonesEmailObfuscationEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesEmailObfuscationValue string

const (
	ZoneSettingEditParamsItemsZonesEmailObfuscationValueOn  ZoneSettingEditParamsItemsZonesEmailObfuscationValue = "on"
	ZoneSettingEditParamsItemsZonesEmailObfuscationValueOff ZoneSettingEditParamsItemsZonesEmailObfuscationValue = "off"
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingEditParamsItemsZonesH2Prioritization struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesH2PrioritizationID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesH2PrioritizationValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesH2PrioritizationEditable bool

const (
	ZoneSettingEditParamsItemsZonesH2PrioritizationEditableTrue  ZoneSettingEditParamsItemsZonesH2PrioritizationEditable = true
	ZoneSettingEditParamsItemsZonesH2PrioritizationEditableFalse ZoneSettingEditParamsItemsZonesH2PrioritizationEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesH2PrioritizationValue string

const (
	ZoneSettingEditParamsItemsZonesH2PrioritizationValueOn     ZoneSettingEditParamsItemsZonesH2PrioritizationValue = "on"
	ZoneSettingEditParamsItemsZonesH2PrioritizationValueOff    ZoneSettingEditParamsItemsZonesH2PrioritizationValue = "off"
	ZoneSettingEditParamsItemsZonesH2PrioritizationValueCustom ZoneSettingEditParamsItemsZonesH2PrioritizationValue = "custom"
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
	ID param.Field[ZoneSettingEditParamsItemsZonesHotlinkProtectionID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesHotlinkProtectionValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesHotlinkProtectionEditable bool

const (
	ZoneSettingEditParamsItemsZonesHotlinkProtectionEditableTrue  ZoneSettingEditParamsItemsZonesHotlinkProtectionEditable = true
	ZoneSettingEditParamsItemsZonesHotlinkProtectionEditableFalse ZoneSettingEditParamsItemsZonesHotlinkProtectionEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesHotlinkProtectionValue string

const (
	ZoneSettingEditParamsItemsZonesHotlinkProtectionValueOn  ZoneSettingEditParamsItemsZonesHotlinkProtectionValue = "on"
	ZoneSettingEditParamsItemsZonesHotlinkProtectionValueOff ZoneSettingEditParamsItemsZonesHotlinkProtectionValue = "off"
)

// HTTP2 enabled for this zone.
type ZoneSettingEditParamsItemsZonesHttp2 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesHttp2ID] `json:"id"`
	// Value of the HTTP2 setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesHttp2Value] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesHttp2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesHttp2) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesHttp2ID string

const (
	ZoneSettingEditParamsItemsZonesHttp2IDHttp2 ZoneSettingEditParamsItemsZonesHttp2ID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesHttp2Editable bool

const (
	ZoneSettingEditParamsItemsZonesHttp2EditableTrue  ZoneSettingEditParamsItemsZonesHttp2Editable = true
	ZoneSettingEditParamsItemsZonesHttp2EditableFalse ZoneSettingEditParamsItemsZonesHttp2Editable = false
)

// Value of the HTTP2 setting.
type ZoneSettingEditParamsItemsZonesHttp2Value string

const (
	ZoneSettingEditParamsItemsZonesHttp2ValueOn  ZoneSettingEditParamsItemsZonesHttp2Value = "on"
	ZoneSettingEditParamsItemsZonesHttp2ValueOff ZoneSettingEditParamsItemsZonesHttp2Value = "off"
)

// HTTP3 enabled for this zone.
type ZoneSettingEditParamsItemsZonesHttp3 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesHttp3ID] `json:"id"`
	// Value of the HTTP3 setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesHttp3Value] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesHttp3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesHttp3) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesHttp3ID string

const (
	ZoneSettingEditParamsItemsZonesHttp3IDHttp3 ZoneSettingEditParamsItemsZonesHttp3ID = "http3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesHttp3Editable bool

const (
	ZoneSettingEditParamsItemsZonesHttp3EditableTrue  ZoneSettingEditParamsItemsZonesHttp3Editable = true
	ZoneSettingEditParamsItemsZonesHttp3EditableFalse ZoneSettingEditParamsItemsZonesHttp3Editable = false
)

// Value of the HTTP3 setting.
type ZoneSettingEditParamsItemsZonesHttp3Value string

const (
	ZoneSettingEditParamsItemsZonesHttp3ValueOn  ZoneSettingEditParamsItemsZonesHttp3Value = "on"
	ZoneSettingEditParamsItemsZonesHttp3ValueOff ZoneSettingEditParamsItemsZonesHttp3Value = "off"
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingEditParamsItemsZonesImageResizing struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesImageResizingID] `json:"id"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value param.Field[ZoneSettingEditParamsItemsZonesImageResizingValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesImageResizingEditable bool

const (
	ZoneSettingEditParamsItemsZonesImageResizingEditableTrue  ZoneSettingEditParamsItemsZonesImageResizingEditable = true
	ZoneSettingEditParamsItemsZonesImageResizingEditableFalse ZoneSettingEditParamsItemsZonesImageResizingEditable = false
)

// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
type ZoneSettingEditParamsItemsZonesImageResizingValue string

const (
	ZoneSettingEditParamsItemsZonesImageResizingValueOn   ZoneSettingEditParamsItemsZonesImageResizingValue = "on"
	ZoneSettingEditParamsItemsZonesImageResizingValueOff  ZoneSettingEditParamsItemsZonesImageResizingValue = "off"
	ZoneSettingEditParamsItemsZonesImageResizingValueOpen ZoneSettingEditParamsItemsZonesImageResizingValue = "open"
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingEditParamsItemsZonesIPGeolocation struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesIPGeolocationID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesIPGeolocationValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesIPGeolocationEditable bool

const (
	ZoneSettingEditParamsItemsZonesIPGeolocationEditableTrue  ZoneSettingEditParamsItemsZonesIPGeolocationEditable = true
	ZoneSettingEditParamsItemsZonesIPGeolocationEditableFalse ZoneSettingEditParamsItemsZonesIPGeolocationEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesIPGeolocationValue string

const (
	ZoneSettingEditParamsItemsZonesIPGeolocationValueOn  ZoneSettingEditParamsItemsZonesIPGeolocationValue = "on"
	ZoneSettingEditParamsItemsZonesIPGeolocationValueOff ZoneSettingEditParamsItemsZonesIPGeolocationValue = "off"
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingEditParamsItemsZonesIpv6 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesIpv6ID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesIpv6Value] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesIpv6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesIpv6) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesIpv6ID string

const (
	ZoneSettingEditParamsItemsZonesIpv6IDIpv6 ZoneSettingEditParamsItemsZonesIpv6ID = "ipv6"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesIpv6Editable bool

const (
	ZoneSettingEditParamsItemsZonesIpv6EditableTrue  ZoneSettingEditParamsItemsZonesIpv6Editable = true
	ZoneSettingEditParamsItemsZonesIpv6EditableFalse ZoneSettingEditParamsItemsZonesIpv6Editable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesIpv6Value string

const (
	ZoneSettingEditParamsItemsZonesIpv6ValueOff ZoneSettingEditParamsItemsZonesIpv6Value = "off"
	ZoneSettingEditParamsItemsZonesIpv6ValueOn  ZoneSettingEditParamsItemsZonesIpv6Value = "on"
)

// Maximum size of an allowable upload.
type ZoneSettingEditParamsItemsZonesMaxUpload struct {
	// identifier of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesMaxUploadID] `json:"id"`
	// Value of the zone setting. Notes: The size depends on the plan level of the
	// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
	Value param.Field[ZoneSettingEditParamsItemsZonesMaxUploadValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMaxUploadEditable bool

const (
	ZoneSettingEditParamsItemsZonesMaxUploadEditableTrue  ZoneSettingEditParamsItemsZonesMaxUploadEditable = true
	ZoneSettingEditParamsItemsZonesMaxUploadEditableFalse ZoneSettingEditParamsItemsZonesMaxUploadEditable = false
)

// Value of the zone setting. Notes: The size depends on the plan level of the
// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
type ZoneSettingEditParamsItemsZonesMaxUploadValue float64

const (
	ZoneSettingEditParamsItemsZonesMaxUploadValue100 ZoneSettingEditParamsItemsZonesMaxUploadValue = 100
	ZoneSettingEditParamsItemsZonesMaxUploadValue200 ZoneSettingEditParamsItemsZonesMaxUploadValue = 200
	ZoneSettingEditParamsItemsZonesMaxUploadValue500 ZoneSettingEditParamsItemsZonesMaxUploadValue = 500
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingEditParamsItemsZonesMinTlsVersion struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesMinTlsVersionID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesMinTlsVersionValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesMinTlsVersion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesMinTlsVersion) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesMinTlsVersionID string

const (
	ZoneSettingEditParamsItemsZonesMinTlsVersionIDMinTlsVersion ZoneSettingEditParamsItemsZonesMinTlsVersionID = "min_tls_version"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMinTlsVersionEditable bool

const (
	ZoneSettingEditParamsItemsZonesMinTlsVersionEditableTrue  ZoneSettingEditParamsItemsZonesMinTlsVersionEditable = true
	ZoneSettingEditParamsItemsZonesMinTlsVersionEditableFalse ZoneSettingEditParamsItemsZonesMinTlsVersionEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesMinTlsVersionValue string

const (
	ZoneSettingEditParamsItemsZonesMinTlsVersionValue1_0 ZoneSettingEditParamsItemsZonesMinTlsVersionValue = "1.0"
	ZoneSettingEditParamsItemsZonesMinTlsVersionValue1_1 ZoneSettingEditParamsItemsZonesMinTlsVersionValue = "1.1"
	ZoneSettingEditParamsItemsZonesMinTlsVersionValue1_2 ZoneSettingEditParamsItemsZonesMinTlsVersionValue = "1.2"
	ZoneSettingEditParamsItemsZonesMinTlsVersionValue1_3 ZoneSettingEditParamsItemsZonesMinTlsVersionValue = "1.3"
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingEditParamsItemsZonesMinify struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsZonesMinifyID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesMinifyValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMinifyEditable bool

const (
	ZoneSettingEditParamsItemsZonesMinifyEditableTrue  ZoneSettingEditParamsItemsZonesMinifyEditable = true
	ZoneSettingEditParamsItemsZonesMinifyEditableFalse ZoneSettingEditParamsItemsZonesMinifyEditable = false
)

// Value of the zone setting.
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

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingEditParamsItemsZonesMirage struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesMirageID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesMirageValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMirageEditable bool

const (
	ZoneSettingEditParamsItemsZonesMirageEditableTrue  ZoneSettingEditParamsItemsZonesMirageEditable = true
	ZoneSettingEditParamsItemsZonesMirageEditableFalse ZoneSettingEditParamsItemsZonesMirageEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesMirageValue string

const (
	ZoneSettingEditParamsItemsZonesMirageValueOn  ZoneSettingEditParamsItemsZonesMirageValue = "on"
	ZoneSettingEditParamsItemsZonesMirageValueOff ZoneSettingEditParamsItemsZonesMirageValue = "off"
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingEditParamsItemsZonesMobileRedirect struct {
	// Identifier of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesMobileRedirectID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesMobileRedirectValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMobileRedirectEditable bool

const (
	ZoneSettingEditParamsItemsZonesMobileRedirectEditableTrue  ZoneSettingEditParamsItemsZonesMobileRedirectEditable = true
	ZoneSettingEditParamsItemsZonesMobileRedirectEditableFalse ZoneSettingEditParamsItemsZonesMobileRedirectEditable = false
)

// Value of the zone setting.
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

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingEditParamsItemsZonesNel struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsZonesNelID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesNelValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesNel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesNel) implementsZoneSettingEditParamsItem() {}

// Zone setting identifier.
type ZoneSettingEditParamsItemsZonesNelID string

const (
	ZoneSettingEditParamsItemsZonesNelIDNel ZoneSettingEditParamsItemsZonesNelID = "nel"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesNelEditable bool

const (
	ZoneSettingEditParamsItemsZonesNelEditableTrue  ZoneSettingEditParamsItemsZonesNelEditable = true
	ZoneSettingEditParamsItemsZonesNelEditableFalse ZoneSettingEditParamsItemsZonesNelEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesNelValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingEditParamsItemsZonesNelValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingEditParamsItemsZonesOpportunisticEncryption struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesOpportunisticEncryptionID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesOpportunisticEncryptionEditable bool

const (
	ZoneSettingEditParamsItemsZonesOpportunisticEncryptionEditableTrue  ZoneSettingEditParamsItemsZonesOpportunisticEncryptionEditable = true
	ZoneSettingEditParamsItemsZonesOpportunisticEncryptionEditableFalse ZoneSettingEditParamsItemsZonesOpportunisticEncryptionEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValue string

const (
	ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValueOn  ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValue = "on"
	ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValueOff ZoneSettingEditParamsItemsZonesOpportunisticEncryptionValue = "off"
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingEditParamsItemsZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesOpportunisticOnionID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingEditParamsItemsZonesOpportunisticOnionValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesOpportunisticOnionEditable bool

const (
	ZoneSettingEditParamsItemsZonesOpportunisticOnionEditableTrue  ZoneSettingEditParamsItemsZonesOpportunisticOnionEditable = true
	ZoneSettingEditParamsItemsZonesOpportunisticOnionEditableFalse ZoneSettingEditParamsItemsZonesOpportunisticOnionEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditParamsItemsZonesOpportunisticOnionValue string

const (
	ZoneSettingEditParamsItemsZonesOpportunisticOnionValueOn  ZoneSettingEditParamsItemsZonesOpportunisticOnionValue = "on"
	ZoneSettingEditParamsItemsZonesOpportunisticOnionValueOff ZoneSettingEditParamsItemsZonesOpportunisticOnionValue = "off"
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingEditParamsItemsZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesOrangeToOrangeID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesOrangeToOrangeValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesOrangeToOrangeEditable bool

const (
	ZoneSettingEditParamsItemsZonesOrangeToOrangeEditableTrue  ZoneSettingEditParamsItemsZonesOrangeToOrangeEditable = true
	ZoneSettingEditParamsItemsZonesOrangeToOrangeEditableFalse ZoneSettingEditParamsItemsZonesOrangeToOrangeEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesOrangeToOrangeValue string

const (
	ZoneSettingEditParamsItemsZonesOrangeToOrangeValueOn  ZoneSettingEditParamsItemsZonesOrangeToOrangeValue = "on"
	ZoneSettingEditParamsItemsZonesOrangeToOrangeValueOff ZoneSettingEditParamsItemsZonesOrangeToOrangeValue = "off"
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsZonesOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruEditable bool

const (
	ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruEditableTrue  ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruEditable = true
	ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruEditableFalse ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValue string

const (
	ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValueOn  ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValue = "on"
	ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValueOff ZoneSettingEditParamsItemsZonesOriginErrorPagePassThruValue = "off"
)

type ZoneSettingEditParamsItemsZonesOriginMaxHTTPVersion struct {
	// Identifier of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesOriginMaxHTTPVersionID] `json:"id,required"`
}

func (r ZoneSettingEditParamsItemsZonesOriginMaxHTTPVersion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesOriginMaxHTTPVersion) implementsZoneSettingEditParamsItem() {}

// Identifier of the zone setting.
type ZoneSettingEditParamsItemsZonesOriginMaxHTTPVersionID string

const (
	ZoneSettingEditParamsItemsZonesOriginMaxHTTPVersionIDOriginMaxHTTPVersion ZoneSettingEditParamsItemsZonesOriginMaxHTTPVersionID = "origin_max_http_version"
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingEditParamsItemsZonesPolish struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesPolishID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesPolishValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesPolishEditable bool

const (
	ZoneSettingEditParamsItemsZonesPolishEditableTrue  ZoneSettingEditParamsItemsZonesPolishEditable = true
	ZoneSettingEditParamsItemsZonesPolishEditableFalse ZoneSettingEditParamsItemsZonesPolishEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesPolishValue string

const (
	ZoneSettingEditParamsItemsZonesPolishValueOff      ZoneSettingEditParamsItemsZonesPolishValue = "off"
	ZoneSettingEditParamsItemsZonesPolishValueLossless ZoneSettingEditParamsItemsZonesPolishValue = "lossless"
	ZoneSettingEditParamsItemsZonesPolishValueLossy    ZoneSettingEditParamsItemsZonesPolishValue = "lossy"
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesPrefetchPreloadID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesPrefetchPreloadValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesPrefetchPreloadEditable bool

const (
	ZoneSettingEditParamsItemsZonesPrefetchPreloadEditableTrue  ZoneSettingEditParamsItemsZonesPrefetchPreloadEditable = true
	ZoneSettingEditParamsItemsZonesPrefetchPreloadEditableFalse ZoneSettingEditParamsItemsZonesPrefetchPreloadEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesPrefetchPreloadValue string

const (
	ZoneSettingEditParamsItemsZonesPrefetchPreloadValueOn  ZoneSettingEditParamsItemsZonesPrefetchPreloadValue = "on"
	ZoneSettingEditParamsItemsZonesPrefetchPreloadValueOff ZoneSettingEditParamsItemsZonesPrefetchPreloadValue = "off"
)

// Maximum time between two read operations from origin.
type ZoneSettingEditParamsItemsZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesProxyReadTimeoutID] `json:"id"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value param.Field[float64] `json:"value"`
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
type ZoneSettingEditParamsItemsZonesPseudoIpv4 struct {
	// Value of the Pseudo IPv4 setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesPseudoIpv4ID] `json:"id"`
	// Value of the Pseudo IPv4 setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesPseudoIpv4Value] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesPseudoIpv4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesPseudoIpv4) implementsZoneSettingEditParamsItem() {}

// Value of the Pseudo IPv4 setting.
type ZoneSettingEditParamsItemsZonesPseudoIpv4ID string

const (
	ZoneSettingEditParamsItemsZonesPseudoIpv4IDPseudoIpv4 ZoneSettingEditParamsItemsZonesPseudoIpv4ID = "pseudo_ipv4"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesPseudoIpv4Editable bool

const (
	ZoneSettingEditParamsItemsZonesPseudoIpv4EditableTrue  ZoneSettingEditParamsItemsZonesPseudoIpv4Editable = true
	ZoneSettingEditParamsItemsZonesPseudoIpv4EditableFalse ZoneSettingEditParamsItemsZonesPseudoIpv4Editable = false
)

// Value of the Pseudo IPv4 setting.
type ZoneSettingEditParamsItemsZonesPseudoIpv4Value string

const (
	ZoneSettingEditParamsItemsZonesPseudoIpv4ValueOff             ZoneSettingEditParamsItemsZonesPseudoIpv4Value = "off"
	ZoneSettingEditParamsItemsZonesPseudoIpv4ValueAddHeader       ZoneSettingEditParamsItemsZonesPseudoIpv4Value = "add_header"
	ZoneSettingEditParamsItemsZonesPseudoIpv4ValueOverwriteHeader ZoneSettingEditParamsItemsZonesPseudoIpv4Value = "overwrite_header"
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsZonesResponseBuffering struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesResponseBufferingID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesResponseBufferingValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesResponseBufferingEditable bool

const (
	ZoneSettingEditParamsItemsZonesResponseBufferingEditableTrue  ZoneSettingEditParamsItemsZonesResponseBufferingEditable = true
	ZoneSettingEditParamsItemsZonesResponseBufferingEditableFalse ZoneSettingEditParamsItemsZonesResponseBufferingEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesResponseBufferingValue string

const (
	ZoneSettingEditParamsItemsZonesResponseBufferingValueOn  ZoneSettingEditParamsItemsZonesResponseBufferingValue = "on"
	ZoneSettingEditParamsItemsZonesResponseBufferingValueOff ZoneSettingEditParamsItemsZonesResponseBufferingValue = "off"
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
	ID param.Field[ZoneSettingEditParamsItemsZonesRocketLoaderID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesRocketLoaderValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesRocketLoaderEditable bool

const (
	ZoneSettingEditParamsItemsZonesRocketLoaderEditableTrue  ZoneSettingEditParamsItemsZonesRocketLoaderEditable = true
	ZoneSettingEditParamsItemsZonesRocketLoaderEditableFalse ZoneSettingEditParamsItemsZonesRocketLoaderEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesRocketLoaderValue string

const (
	ZoneSettingEditParamsItemsZonesRocketLoaderValueOn  ZoneSettingEditParamsItemsZonesRocketLoaderValue = "on"
	ZoneSettingEditParamsItemsZonesRocketLoaderValueOff ZoneSettingEditParamsItemsZonesRocketLoaderValue = "off"
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID    param.Field[ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationID]    `json:"id"`
	Value param.Field[ZoneSettingEditParamsItemsZonesSchemasAutomaticPlatformOptimizationValue] `json:"value"`
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

// Cloudflare security header for a zone.
type ZoneSettingEditParamsItemsZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID    param.Field[ZoneSettingEditParamsItemsZonesSecurityHeaderID]    `json:"id"`
	Value param.Field[ZoneSettingEditParamsItemsZonesSecurityHeaderValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSecurityHeaderEditable bool

const (
	ZoneSettingEditParamsItemsZonesSecurityHeaderEditableTrue  ZoneSettingEditParamsItemsZonesSecurityHeaderEditable = true
	ZoneSettingEditParamsItemsZonesSecurityHeaderEditableFalse ZoneSettingEditParamsItemsZonesSecurityHeaderEditable = false
)

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

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingEditParamsItemsZonesSecurityLevel struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesSecurityLevelID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesSecurityLevelValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSecurityLevelEditable bool

const (
	ZoneSettingEditParamsItemsZonesSecurityLevelEditableTrue  ZoneSettingEditParamsItemsZonesSecurityLevelEditable = true
	ZoneSettingEditParamsItemsZonesSecurityLevelEditableFalse ZoneSettingEditParamsItemsZonesSecurityLevelEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesSecurityLevelValue string

const (
	ZoneSettingEditParamsItemsZonesSecurityLevelValueOff            ZoneSettingEditParamsItemsZonesSecurityLevelValue = "off"
	ZoneSettingEditParamsItemsZonesSecurityLevelValueEssentiallyOff ZoneSettingEditParamsItemsZonesSecurityLevelValue = "essentially_off"
	ZoneSettingEditParamsItemsZonesSecurityLevelValueLow            ZoneSettingEditParamsItemsZonesSecurityLevelValue = "low"
	ZoneSettingEditParamsItemsZonesSecurityLevelValueMedium         ZoneSettingEditParamsItemsZonesSecurityLevelValue = "medium"
	ZoneSettingEditParamsItemsZonesSecurityLevelValueHigh           ZoneSettingEditParamsItemsZonesSecurityLevelValue = "high"
	ZoneSettingEditParamsItemsZonesSecurityLevelValueUnderAttack    ZoneSettingEditParamsItemsZonesSecurityLevelValue = "under_attack"
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
	ID param.Field[ZoneSettingEditParamsItemsZonesServerSideExcludeID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesServerSideExcludeValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesServerSideExcludeEditable bool

const (
	ZoneSettingEditParamsItemsZonesServerSideExcludeEditableTrue  ZoneSettingEditParamsItemsZonesServerSideExcludeEditable = true
	ZoneSettingEditParamsItemsZonesServerSideExcludeEditableFalse ZoneSettingEditParamsItemsZonesServerSideExcludeEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesServerSideExcludeValue string

const (
	ZoneSettingEditParamsItemsZonesServerSideExcludeValueOn  ZoneSettingEditParamsItemsZonesServerSideExcludeValue = "on"
	ZoneSettingEditParamsItemsZonesServerSideExcludeValueOff ZoneSettingEditParamsItemsZonesServerSideExcludeValue = "off"
)

// Allow SHA1 support.
type ZoneSettingEditParamsItemsZonesSha1Support struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsZonesSha1SupportID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesSha1SupportValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSha1SupportEditable bool

const (
	ZoneSettingEditParamsItemsZonesSha1SupportEditableTrue  ZoneSettingEditParamsItemsZonesSha1SupportEditable = true
	ZoneSettingEditParamsItemsZonesSha1SupportEditableFalse ZoneSettingEditParamsItemsZonesSha1SupportEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesSha1SupportValue string

const (
	ZoneSettingEditParamsItemsZonesSha1SupportValueOff ZoneSettingEditParamsItemsZonesSha1SupportValue = "off"
	ZoneSettingEditParamsItemsZonesSha1SupportValueOn  ZoneSettingEditParamsItemsZonesSha1SupportValue = "on"
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingEditParamsItemsZonesSortQueryStringForCache struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesSortQueryStringForCacheID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSortQueryStringForCacheEditable bool

const (
	ZoneSettingEditParamsItemsZonesSortQueryStringForCacheEditableTrue  ZoneSettingEditParamsItemsZonesSortQueryStringForCacheEditable = true
	ZoneSettingEditParamsItemsZonesSortQueryStringForCacheEditableFalse ZoneSettingEditParamsItemsZonesSortQueryStringForCacheEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValue string

const (
	ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValueOn  ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValue = "on"
	ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValueOff ZoneSettingEditParamsItemsZonesSortQueryStringForCacheValue = "off"
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
type ZoneSettingEditParamsItemsZonesSsl struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesSslID] `json:"id"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value param.Field[ZoneSettingEditParamsItemsZonesSslValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesSsl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesSsl) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesSslID string

const (
	ZoneSettingEditParamsItemsZonesSslIDSsl ZoneSettingEditParamsItemsZonesSslID = "ssl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSslEditable bool

const (
	ZoneSettingEditParamsItemsZonesSslEditableTrue  ZoneSettingEditParamsItemsZonesSslEditable = true
	ZoneSettingEditParamsItemsZonesSslEditableFalse ZoneSettingEditParamsItemsZonesSslEditable = false
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingEditParamsItemsZonesSslValue string

const (
	ZoneSettingEditParamsItemsZonesSslValueOff      ZoneSettingEditParamsItemsZonesSslValue = "off"
	ZoneSettingEditParamsItemsZonesSslValueFlexible ZoneSettingEditParamsItemsZonesSslValue = "flexible"
	ZoneSettingEditParamsItemsZonesSslValueFull     ZoneSettingEditParamsItemsZonesSslValue = "full"
	ZoneSettingEditParamsItemsZonesSslValueStrict   ZoneSettingEditParamsItemsZonesSslValue = "strict"
)

type ZoneSettingEditParamsItemsZonesSslRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[ZoneSettingEditParamsItemsZonesSslRecommenderID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingEditParamsItemsZonesSslRecommender) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesSslRecommender) implementsZoneSettingEditParamsItem() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingEditParamsItemsZonesSslRecommenderID string

const (
	ZoneSettingEditParamsItemsZonesSslRecommenderIDSslRecommender ZoneSettingEditParamsItemsZonesSslRecommenderID = "ssl_recommender"
)

// Only allows TLS1.2.
type ZoneSettingEditParamsItemsZonesTls1_2Only struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsZonesTls1_2OnlyID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesTls1_2OnlyValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesTls1_2Only) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesTls1_2Only) implementsZoneSettingEditParamsItem() {}

// Zone setting identifier.
type ZoneSettingEditParamsItemsZonesTls1_2OnlyID string

const (
	ZoneSettingEditParamsItemsZonesTls1_2OnlyIDTls1_2Only ZoneSettingEditParamsItemsZonesTls1_2OnlyID = "tls_1_2_only"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTls1_2OnlyEditable bool

const (
	ZoneSettingEditParamsItemsZonesTls1_2OnlyEditableTrue  ZoneSettingEditParamsItemsZonesTls1_2OnlyEditable = true
	ZoneSettingEditParamsItemsZonesTls1_2OnlyEditableFalse ZoneSettingEditParamsItemsZonesTls1_2OnlyEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesTls1_2OnlyValue string

const (
	ZoneSettingEditParamsItemsZonesTls1_2OnlyValueOff ZoneSettingEditParamsItemsZonesTls1_2OnlyValue = "off"
	ZoneSettingEditParamsItemsZonesTls1_2OnlyValueOn  ZoneSettingEditParamsItemsZonesTls1_2OnlyValue = "on"
)

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingEditParamsItemsZonesTls1_3 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesTls1_3ID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingEditParamsItemsZonesTls1_3Value] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesTls1_3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesTls1_3) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesTls1_3ID string

const (
	ZoneSettingEditParamsItemsZonesTls1_3IDTls1_3 ZoneSettingEditParamsItemsZonesTls1_3ID = "tls_1_3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTls1_3Editable bool

const (
	ZoneSettingEditParamsItemsZonesTls1_3EditableTrue  ZoneSettingEditParamsItemsZonesTls1_3Editable = true
	ZoneSettingEditParamsItemsZonesTls1_3EditableFalse ZoneSettingEditParamsItemsZonesTls1_3Editable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditParamsItemsZonesTls1_3Value string

const (
	ZoneSettingEditParamsItemsZonesTls1_3ValueOn  ZoneSettingEditParamsItemsZonesTls1_3Value = "on"
	ZoneSettingEditParamsItemsZonesTls1_3ValueOff ZoneSettingEditParamsItemsZonesTls1_3Value = "off"
	ZoneSettingEditParamsItemsZonesTls1_3ValueZrt ZoneSettingEditParamsItemsZonesTls1_3Value = "zrt"
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingEditParamsItemsZonesTlsClientAuth struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesTlsClientAuthID] `json:"id"`
	// value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesTlsClientAuthValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesTlsClientAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesTlsClientAuth) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesTlsClientAuthID string

const (
	ZoneSettingEditParamsItemsZonesTlsClientAuthIDTlsClientAuth ZoneSettingEditParamsItemsZonesTlsClientAuthID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTlsClientAuthEditable bool

const (
	ZoneSettingEditParamsItemsZonesTlsClientAuthEditableTrue  ZoneSettingEditParamsItemsZonesTlsClientAuthEditable = true
	ZoneSettingEditParamsItemsZonesTlsClientAuthEditableFalse ZoneSettingEditParamsItemsZonesTlsClientAuthEditable = false
)

// value of the zone setting.
type ZoneSettingEditParamsItemsZonesTlsClientAuthValue string

const (
	ZoneSettingEditParamsItemsZonesTlsClientAuthValueOn  ZoneSettingEditParamsItemsZonesTlsClientAuthValue = "on"
	ZoneSettingEditParamsItemsZonesTlsClientAuthValueOff ZoneSettingEditParamsItemsZonesTlsClientAuthValue = "off"
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsZonesTrueClientIPHeader struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesTrueClientIPHeaderID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTrueClientIPHeaderEditable bool

const (
	ZoneSettingEditParamsItemsZonesTrueClientIPHeaderEditableTrue  ZoneSettingEditParamsItemsZonesTrueClientIPHeaderEditable = true
	ZoneSettingEditParamsItemsZonesTrueClientIPHeaderEditableFalse ZoneSettingEditParamsItemsZonesTrueClientIPHeaderEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValue string

const (
	ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValueOn  ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValue = "on"
	ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValueOff ZoneSettingEditParamsItemsZonesTrueClientIPHeaderValue = "off"
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
type ZoneSettingEditParamsItemsZonesWaf struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesWafID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesWafValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsZonesWaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsZonesWaf) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsZonesWafID string

const (
	ZoneSettingEditParamsItemsZonesWafIDWaf ZoneSettingEditParamsItemsZonesWafID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesWafEditable bool

const (
	ZoneSettingEditParamsItemsZonesWafEditableTrue  ZoneSettingEditParamsItemsZonesWafEditable = true
	ZoneSettingEditParamsItemsZonesWafEditableFalse ZoneSettingEditParamsItemsZonesWafEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesWafValue string

const (
	ZoneSettingEditParamsItemsZonesWafValueOn  ZoneSettingEditParamsItemsZonesWafValue = "on"
	ZoneSettingEditParamsItemsZonesWafValueOff ZoneSettingEditParamsItemsZonesWafValue = "off"
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingEditParamsItemsZonesWebp struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesWebpID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesWebpValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesWebpEditable bool

const (
	ZoneSettingEditParamsItemsZonesWebpEditableTrue  ZoneSettingEditParamsItemsZonesWebpEditable = true
	ZoneSettingEditParamsItemsZonesWebpEditableFalse ZoneSettingEditParamsItemsZonesWebpEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesWebpValue string

const (
	ZoneSettingEditParamsItemsZonesWebpValueOff ZoneSettingEditParamsItemsZonesWebpValue = "off"
	ZoneSettingEditParamsItemsZonesWebpValueOn  ZoneSettingEditParamsItemsZonesWebpValue = "on"
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
	ID param.Field[ZoneSettingEditParamsItemsZonesWebsocketsID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesWebsocketsValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesWebsocketsEditable bool

const (
	ZoneSettingEditParamsItemsZonesWebsocketsEditableTrue  ZoneSettingEditParamsItemsZonesWebsocketsEditable = true
	ZoneSettingEditParamsItemsZonesWebsocketsEditableFalse ZoneSettingEditParamsItemsZonesWebsocketsEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesWebsocketsValue string

const (
	ZoneSettingEditParamsItemsZonesWebsocketsValueOff ZoneSettingEditParamsItemsZonesWebsocketsValue = "off"
	ZoneSettingEditParamsItemsZonesWebsocketsValueOn  ZoneSettingEditParamsItemsZonesWebsocketsValue = "on"
)
